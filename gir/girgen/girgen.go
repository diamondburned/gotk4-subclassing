package girgen

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"sync"
	"text/template"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/internal/pen"
)

func newGoTemplate(block string) *template.Template {
	t := template.New("")
	t.Funcs(template.FuncMap{
		"PascalToGo":               PascalToGo,
		"SnakeToGo":                SnakeToGo,
		"FirstChar":                FirstChar,
		"GoDoc":                    GoDoc,
		"CommentReflowLinesIndent": CommentReflowLinesIndent,
	})
	t = template.Must(t.Parse(block))
	return t
}

// Generator is a big generator that manages multiple repositories.
type Generator struct {
	// KnownTypes contains a list of type checks that return true if the given
	// type matches a known type.
	KnownTypes []func(string) bool
	// NoMarshalPkgs contains a list of namespace package names that don't have
	// marshalers generated. GLib is by default in this list, because it
	// shouldn't have any marshalers.
	NoMarshalPkgs []string

	Repos gir.Repositories

	logger *log.Logger

	unknownTypes map[string]struct{}
	uTypesMut    sync.Mutex
}

// NewGenerator creates a new generator with sane defaults.
func NewGenerator(repos gir.Repositories) *Generator {
	return &Generator{
		Repos:        repos,
		unknownTypes: map[string]struct{}{},
	}
}

// WithLogger sets the generator's logger.
func (g *Generator) WithLogger(logger *log.Logger) {
	g.logger = logger
}

func (g *Generator) debugln(v ...interface{}) {
	if g.logger != nil {
		g.logger.Println(v...)
	}
}

func (g *Generator) warnUnknownType(typ string) {
	if g.logger != nil {
		// g.uTypesMut.Lock()
		// defer g.uTypesMut.Unlock()

		// _, warned := g.unknownTypes[typ]
		// if warned {
		// 	return
		// }

		g.debugln("unknown gir type", typ)
		// g.unknownTypes[typ] = struct{}{}
	}
}

// UseNamespace creates a new namespace generator using the given namespace.
func (g *Generator) UseNamespace(namespace string) *NamespaceGenerator {
	res := g.Repos.FindNamespace(namespace, "")
	if res == nil {
		return nil
	}

	buf := bytes.Buffer{}
	buf.Grow(4 * 1024 * 1024) // 4MB

	return &NamespaceGenerator{
		pen:  pen.New(&buf),
		body: &buf,

		imports: map[string]struct{}{},
		gen:     g,
		current: res,
	}
}

// NamespaceGenerator is a generator for a specific namespace.
type NamespaceGenerator struct {
	pen  *pen.Pen
	body *bytes.Buffer

	imports map[string]struct{}

	gen     *Generator
	current *gir.NamespaceFindResult
}

// Generate generates the current namespace into the given writer.
func (ng *NamespaceGenerator) Generate(w io.Writer) error {
	ng.addImport("unsafe")
	ng.addImport("github.com/gotk3/gotk3/glib")
	ng.addImport("github.com/diamondburned/gotk4/internal/gextras")
	ng.addImport("github.com/diamondburned/gotk4/internal/callback")

	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	// CALL GENERATION FUNCTIONS HERE !!!
	ng.generateInit()
	ng.generateAliases()
	ng.generateEnums()
	ng.generateBitfields()
	ng.generateFuncs()
	ng.generateRecords()

	if err := ng.pen.Flush(); err != nil {
		return err
	}

	pen := pen.New(w)
	pen.Words("// Code generated by girgen. DO NOT EDIT.")
	pen.Line()

	pen.Words("package", ng.PackageName())
	pen.Line()

	if len(ng.imports) > 0 {
		pen.Words("import (")
		for imp := range ng.imports {
			pen.Words(strconv.Quote(imp))
		}
		pen.Block(")")
		pen.Line()
	}

	pkgs := []string{"// #cgo pkg-config:"}
	for _, pkg := range ng.current.Repository.Packages {
		pkgs = append(pkgs, pkg.Name)
	}

	pen.Words(pkgs...)
	pen.Words("// #cgo CFLAGS: -Wno-deprecated-declarations") // opt to warn over comments
	for _, cIncl := range ng.current.Repository.CIncludes {
		pen.Words("// #include", fmt.Sprintf("<%s>", cIncl.Name))
	}
	// pen.Words("// extern void callbackDelete(gpointer);")
	pen.Words(`import "C"`)
	pen.Line()

	// TODO: detect when this is needed.
	// pen.Words("//export callbackDelete")
	// pen.Block(`func callbackDelete(ptr C.gpointer) { callback.Delete(uintptr(ptr)) }`)

	pen.Write(ng.body.Bytes())

	return pen.Flush()
}

// PackageName returns the current namespace's package name.
func (ng *NamespaceGenerator) PackageName() string {
	return gir.GoNamespace(ng.current.Namespace)
}

// Namespace returns the generator's namespace.
func (ng *NamespaceGenerator) Namespace() gir.Namespace {
	return ng.current.Namespace
}

// Repository returns the generator's repository.
func (ng *NamespaceGenerator) Repository() gir.PkgRepository {
	return ng.current.Repository
}

func (ng *NamespaceGenerator) debugln(v ...interface{}) {
	if ng.gen.logger != nil {
		prefix := []interface{}{"package", ng.current.Namespace.Name + ":"}
		prefix = append(prefix, v...)

		ng.gen.debugln(prefix...)
	}
}

func (ng *NamespaceGenerator) addImport(pkgPath string) {
	_, ok := ng.imports[pkgPath]
	if ok {
		return
	}

	ng.imports[pkgPath] = struct{}{}
}

// resolveType resolves the given type from the GIR type field. It returns nil
// if the type is not known. It does not recursively traverse the type.
func (ng *NamespaceGenerator) resolveType(typ gir.Type) *resolvedType {
	resolved := ng._resolveType(typ)
	if resolved != nil && resolved.Import != "" {
		ng.addImport(resolved.Import)
	}

	return resolved
}

func (ng *NamespaceGenerator) _resolveType(typ gir.Type) *resolvedType {
	// Resolve the unknown namespace that is GLib and primitive types.
	switch typ.Name {
	case "none":
		return builtinType("", typ)
	case "gboolean":
		return builtinType("bool", typ)
	case "gfloat":
		return builtinType("float32", typ)
	case "gdouble":
		return builtinType("float64", typ)
	case "long double": // pain
		return builtinType("float64", typ)
	case "gint", "gssize":
		return builtinType("int", typ)
	case "gint8":
		return builtinType("int8", typ)
	case "gint16", "gshort":
		return builtinType("int16", typ)
	case "gint32", "glong", "int32":
		return builtinType("int32", typ)
	case "gint64":
		return builtinType("int64", typ)
	case "guint", "gsize":
		return builtinType("uint", typ)
	case "guchar", "gchar":
		return builtinType("byte", typ)
	case "guint8":
		return builtinType("uint8", typ)
	case "guint16", "gushort":
		return builtinType("uint16", typ)
	case "guint32", "gulong", "gunichar": // pain pain pain pain
		return builtinType("uint32", typ)
	case "guint64":
		return builtinType("uint64", typ)
	case "utf8", "filename": // filename is probably UTF-16 hybrid ???
		return builtinType("string", typ)
	case "gpointer":
		// TODO: ignore field
		// TODO: aaaaaaaaaaaaaaaaaaaaaaa
		return builtinType("unsafe.Pointer", typ)
	case "GLib.DestroyNotify": // This should be handled externally.
		return builtinType("unsafe.Pointer", typ)
	case "GType":
		return builtinType("glib.Type", typ)
	case "GObject.GValue", "GObject.Value": // inconsistency???
		return builtinType("*glib.Value", typ)
	case "GObject.Object":
		return builtinType("*glib.Object", typ)
	case "GObject.Closure":
		return builtinType("*glib.Closure", typ)
	case "GObject.GInitiallyUnowned":
		return builtinType("glib.InitiallyUnowned", typ)

	case "GObject.Callback":
		// Callback is a special func(Any) Any type, so we treat it as
		// interface{} similarly to object.Connect(). We can use glib's Closure
		// APIs to parse this interface{}.
		return builtinType("interface{}", typ)

	case "va_list":
		// CGo cannot handle variadic argument lists.
		return nil

	// We don't know what these types translates to.
	// TODO: Find a way to map EnumValue type.
	// TODO: Add _full function support.
	case "GObject.EnumValue", "DestroyNotify":
		return nil
	}

	// Types that aren't in the switch tree that match any of these patterns are
	// types that must be in the switch tree, so them not being in there is a
	// bug.
	for _, check := range ng.gen.KnownTypes {
		if check(typ.Name) {
			log.Fatalf("missing gir type %s in the type tree\n", typ.Name)
		}
	}

	result := ng.gen.Repos.FindType(
		ng.current.Namespace.Name,
		ng.current.Namespace.Version,
		typ.Name,
	)
	if result == nil {
		ng.gen.warnUnknownType(typ.Name)
		return nil
	}

	return typeFromResult(typ, result)
}

type resolvedType struct {
	GoType string
	Type   gir.Type
	Import string // optional
}

// builtinType is a convenient function to make a new resolvedType.
func builtinType(goType string, typ gir.Type) *resolvedType {
	resolved := &resolvedType{GoType: goType, Type: typ}
	resolved.setPtr(nil)
	return resolved
}

// typeFromResult creates a resolved type from the given type result.
func typeFromResult(typ gir.Type, result *gir.TypeFindResult) *resolvedType {
	name, _ := result.Info()

	// same namespace, no package qual required.
	if result.SameNamespace {
		resolved := &resolvedType{name, typ, ""}
		resolved.setPtr(result)
		return resolved
	}

	// different namespace.
	pkg := gir.GoNamespace(result.Namespace)
	path := gir.ImportPath(pkg)

	resolved := &resolvedType{pkg + "." + name, typ, path}
	resolved.setPtr(result)
	return resolved
}

func (typ *resolvedType) setPtr(result *gir.TypeFindResult) {
	ptr := strings.Count(typ.Type.CType, "*")

	// Edge case: interfaces must not be pointers. We should still sometimes
	// allow for pointers to interfaces, if needed, but this likely won't work.
	if result != nil && result.Interface != nil && ptr > 0 {
		ptr--
	}
	// Edge case: a string is a gchar*, so we don't need a pointer.
	if typ.Type.Name == "utf8" && ptr > 0 {
		ptr--
	}

	typ.GoType = strings.Repeat("*", ptr) + typ.GoType
}
