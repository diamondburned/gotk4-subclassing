package generators

import (
	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/gotmpl"
	"github.com/diamondburned/gotk4/gir/girgen/types"

	ifacegen "github.com/diamondburned/gotk4/gir/girgen/generators/iface"
)

var classInterfaceTmpl = gotmpl.NewGoTemplate(`
	{{ if .GLibTypeStruct }}
	// {{ .StructName }}Overrider contains methods that are overridable.
	type {{ .StructName }}Overrider interface {
		{{ range .GLibTypeStruct.Methods -}}
		{{ if $.IsInSameFile . -}}
		{{- GoDoc .Go 1 TrailingNewLine -}}
		{{- .Go.Name }}{{ .Go.Tail }}
		{{ end -}}
		{{ end -}}
	}
	{{ end }}

	{{ GoDoc . 0 (OverrideSelfName .StructName) }}
	type {{ .StructName }} struct {
		_ [0]func() // equal guard
		{{ index .Tree.ImplTypes 0 }}

		{{ range (slice .Tree.ImplTypes 1) -}}
		{{ . }}
		{{ end }}
	}

	var (
		{{ range .ImplInterfaces -}}
		_ {{ . }} = (*{{ $.StructName }})(nil)
		{{ end }}
	)

	{{ $needsPrivate := false }}

	{{ if .Abstract }}

	{{ if .IsClass }}
	// {{ .InterfaceName }} describes types inherited from class {{ .StructName }}.
	{{- $needsPrivate = true }}
	//
	// To get the original type, the caller must assert this to an interface or
	// another type.
	type {{ .InterfaceName }} interface {
		externglib.Objector
		base{{ .StructName }}() *{{ .StructName }}
	}
	{{ else }}
	// {{ .InterfaceName }} describes {{ .StructName }}'s interface methods.
	type {{ .InterfaceName }} interface {
		externglib.Objector

		{{ range .Methods -}}
		{{- if $.IsInSameFile . }}
		{{- Synopsis . 1 TrailingNewLine -}}
		{{- .Name }}{{ .Tail }}
		{{- end }}
		{{ else }}
		{{ $needsPrivate = true }}
		base{{ .StructName }}() *{{ .StructName }}
		{{ end -}}
	}
	{{ end }}

	var _ {{ .InterfaceName }} = (*{{ .StructName }})(nil)
	{{ end }}

	{{ if .GLibTypeStruct }}
	{{ if .IsClass }}
	func classInit{{ .InterfaceName }}(gclassPtr, data C.gpointer) {
		C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

		goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
		*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

		{{ if .GLibTypeStruct.Methods }}
		goval := gbox.Get(uintptr(data))
		pclass := (*C.{{ .GLibTypeStruct.CType }})(unsafe.Pointer(gclassPtr))
		// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
		// pclass := (*C.{{ .GLibTypeStruct.CType }})(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

		{{ range .GLibTypeStruct.Methods }}
		if _, ok := goval.(interface{ {{ .Go.Name }}{{ .Go.Tail }} }); ok {
			pclass.{{ .FieldName }} = (*[0]byte)(C.{{ .C.Name }})
		}
		{{ end -}}
		{{ end -}}
	}
	{{ else }}
	func ifaceInit{{ .InterfaceName }}(gifacePtr, data C.gpointer) {
		{{- if .GLibTypeStruct.Methods }}
		iface := (*C.{{ .GLibTypeStruct.CType }})(unsafe.Pointer(gifacePtr))
		{{- range .GLibTypeStruct.Methods }}
		iface.{{ .FieldName }} = (*[0]byte)(C.{{ .C.Name }})
		{{- end }}
		{{- end }}
	}
	{{ end }}

	{{ range .GLibTypeStruct.Methods }}
	//export {{ .C.Name }}
	func {{ .C.Name }}{{ .C.Tail }} {{ .C.Block }}
	{{ end }}
	{{ end }}

	{{ $wrapper := .Tree.WrapName false }}
	func {{ $wrapper }}(obj *externglib.Object) *{{ .StructName }} {
		return {{ .Wrap "obj" }}
	}

	{{ if .HasMarshaler }}
	func marshal{{ .InterfaceName }}(p uintptr) (interface{}, error) {
		return {{ $wrapper }}(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
	}
	{{ end }}

	{{ if $needsPrivate }}
	func ({{ .Recv }} *{{ .StructName }}) base{{ .StructName }}() *{{ .StructName }} {
		return {{ .Recv }}
	}

	// Base{{ .StructName }} returns the underlying base object.
	func Base{{ .StructName }}(obj {{ .InterfaceName }}) *{{ .StructName }} {
		return obj.base{{ .StructName }}()
	}
	{{ end }}

	{{ range .Signals }}
	{{ if $.IsInSameFile . }}
	{{ $name := printf "Connect%s" (KebabToGo true .Name) }}
	{{ GoDoc . 0 (OverrideSelfName $name) }}
	func ({{ $.Recv }} *{{ $.StructName }}) {{ $name }}(f func{{ .Tail }}) externglib.SignalHandle {
		return {{ $.Recv }}.Connect({{ Quote .Name }}, externglib.GeneratedClosure{Func: f})
	}
	{{ end }}
	{{ end }}
`)

var constructorInterfaceImpl = gotmpl.NewGoTemplate(`
	{{ GoDoc . 0 }}
	func {{ .Name }}{{ .Tail }} {{ .Block }}
`)

// methodInterfaceTmpl needs the following type:
//
//    struct {
//        Method
//        StructName string
//    }
//
var methodInterfaceTmpl = gotmpl.NewGoTemplate(`
	{{ with .Method }}
	{{ GoDoc . 0 }}
	func ({{ .Recv }} *{{ $.StructName }}) {{ .Name }}{{ .Tail }} {{ .Block }}
	{{ end }}
`)

// GenerateInterface generates a public interface declaration, optionally
// another one for overriding, and the private struct that implements the
// interface specifically for wrapping opaque C interfaces.
func GenerateInterface(gen FileGeneratorWriter, iface *gir.Interface) bool {
	igen := ifacegen.NewGenerator(gen)
	if !igen.Use(iface) {
		return false
	}

	generateInterfaceGenerator(gen, &igen)
	return true
}

// GenerateClass generates the given class into files.
func GenerateClass(gen FileGeneratorWriter, class *gir.Class) bool {
	igen := ifacegen.NewGenerator(gen)
	if !igen.Use(class) {
		return false
	}

	generateInterfaceGenerator(gen, &igen)
	return true
}

type ifacegenData struct {
	*ifacegen.Generator
	ImplInterfaces []string
	HasMarshaler   bool
}

func (d ifacegenData) Recv() string {
	if len(d.Methods) > 0 {
		return d.Methods[0].Recv
	}
	if len(d.Virtuals) > 0 {
		return d.Virtuals[0].Recv
	}
	return "v"
}

func generateInterfaceGenerator(gen FileGeneratorWriter, igen *ifacegen.Generator) {
	writer := FileWriterFromType(gen, igen)
	writer.Header().NeedsExternGLib()
	writer.Header().Import("unsafe")
	// TOOD: add gbox

	// Import for implementation types.
	for _, parent := range igen.Tree.Requires {
		writer.Header().ImportImpl(parent.Resolved)
	}

	// These conditions should follow what's in the template.
	if igen.IsClass() && igen.GLibTypeStruct != nil && len(igen.GLibTypeStruct.Methods) > 0 {
		writer.Header().ImportCore("gbox")
	}

	data := ifacegenData{
		Generator:      igen,
		ImplInterfaces: igen.ImplInterfaces(),
		HasMarshaler:   false,
	}

	if igen.GLibGetType != "" && !types.FilterCType(gen, igen.GLibGetType) {
		data.HasMarshaler = true
		writer.Header().AddMarshaler(igen.GLibGetType, igen.InterfaceName)
	}

	writer.Pen().WriteTmpl(classInterfaceTmpl, data)
	file.ApplyHeader(writer, igen)

	for _, ctor := range igen.Constructors {
		writer := FileWriterFromType(gen, ctor)
		writer.Header().ApplyFrom(&ctor.Header)

		writer.Pen().WriteTmpl(constructorInterfaceImpl, ctor)
	}

	for _, method := range igen.Methods {
		writer := FileWriterFromType(gen, method)
		writer.Header().ApplyFrom(&method.Header)

		writer.Pen().WriteTmpl(methodInterfaceTmpl, gotmpl.M{
			"Method":     method,
			"StructName": igen.StructName,
		})
	}
}
