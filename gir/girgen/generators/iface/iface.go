// Package iface provides an interface generator.
package iface

import (
	"fmt"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/girgen/cmt"
	"github.com/diamondburned/gotk4/gir/girgen/file"
	"github.com/diamondburned/gotk4/gir/girgen/generators/callable"
	"github.com/diamondburned/gotk4/gir/girgen/logger"
	"github.com/diamondburned/gotk4/gir/girgen/types"
)

// CanGenerate checks if the given class or interface can be generated.
func CanGenerate(gen types.FileGenerator, v interface{}) bool {
	ifaceGen := NewGenerator(gen)
	return ifaceGen.init(v)
}

type Generator struct {
	Root gir.TypeFindResult

	Name           string
	CType          string
	GLibGetType    string
	GLibTypeStruct *TypeStruct
	InfoAttrs      *gir.InfoAttrs
	InfoElements   *gir.InfoElements

	InterfaceName string
	StructName    string

	Virtuals     Methods // deprecated
	Methods      Methods // for big interface
	Constructors Methods
	Signals      []Signal

	Tree types.Tree

	// Abstract bool

	methods  []gir.Method
	virtuals []gir.VirtualMethod

	header file.Header
	gen    types.FileGenerator
	cgen   callable.Generator
}

// Signal describes a GLib signal in minimal function form.
type Signal struct {
	Name string // kebab-cased
	Tail string

	InfoElements gir.InfoElements
}

var _ = cmt.EnsureInfoFields((*Generator)(nil))

// NewGenerator creates a new interface generator instance.
func NewGenerator(gen types.FileGenerator) Generator {
	return Generator{
		gen:  gen,
		cgen: callable.NewGenerator(gen),
		Tree: types.NewTree(gen),
	}
}

func (g *Generator) Logln(lvl logger.Level, v ...interface{}) {
	name := "class"
	if g.IsInterface() {
		name = "interface"
	}
	p := fmt.Sprintf("%s %s (C.%s):", name, g.InterfaceName, g.CType)
	g.gen.Logln(lvl, logger.Prefix(v, p)...)
}

// Reset resets the callback generator.
func (g *Generator) Reset() {
	g.header.Reset()
	g.cgen.Reset()

	g.Tree.Reset()
	g.Methods.reset(0)
	g.Virtuals.reset(0)

	*g = Generator{
		Tree:     g.Tree,
		Methods:  g.Methods,
		Virtuals: g.Virtuals,

		header: g.header,
		cgen:   g.cgen,
		gen:    g.gen,
	}
}

// Abstract returns true if the generator is generating an interface or abstract
// class.
func (g *Generator) Abstract() bool {
	switch v := g.Root.Type.(type) {
	case *gir.Class:
		return v.Abstract
	case *gir.Interface:
		return true
	default:
		panic("unknown root type")
	}
}

// IsClass returns true if the generator is generating a class.
func (g *Generator) IsClass() bool {
	switch g.Root.Type.(type) {
	case *gir.Class:
		return true
	case *gir.Interface:
		return false
	default:
		panic("unknown root type")
	}
}

// IsInterface returns true if the generator is generating an interface.
func (g *Generator) IsInterface() bool { return !g.IsClass() }

// OverriderName returns the name of the overrider interface.
func (g *Generator) OverriderName() string {
	return g.StructName + "Overrider"
}

// Header returns the callback generator's current header.
func (g *Generator) Header() *file.Header {
	return &g.header
}

func (g *Generator) init(typ interface{}) bool {
	resolved := types.TypeFromResult(g.gen, typ)
	if resolved == nil || resolved.Extern == nil {
		return false
	}

	g.Root.Type = typ
	g.Root.NamespaceFindResult = g.gen.Namespace()

	g.CType = resolved.CType
	g.StructName = resolved.ImplName()
	g.InterfaceName = resolved.PublicName()

	switch typ := typ.(type) {
	case *gir.Class:
		g.Name = typ.Name
		g.GLibGetType = typ.GLibGetType
		g.InfoAttrs = &typ.InfoAttrs
		g.InfoElements = &typ.InfoElements

		g.methods = typ.Methods
		g.virtuals = typ.VirtualMethods

		if !typ.IsIntrospectable() || types.Filter(g.gen, typ.Name, typ.CType) {
			return false
		}

		if !g.Tree.ResolveFromType(resolved) {
			g.Logln(logger.Debug, "class cannot be type-resolved")
			return false
		}

	case *gir.Interface:
		g.Name = typ.Name
		g.GLibGetType = typ.GLibGetType
		g.InfoAttrs = &typ.InfoAttrs
		g.InfoElements = &typ.InfoElements

		g.methods = typ.Methods
		g.virtuals = typ.VirtualMethods

		if !typ.IsIntrospectable() || types.Filter(g.gen, typ.Name, typ.CType) {
			return false
		}

		if !g.Tree.ResolveFromType(resolved) {
			g.Logln(logger.Debug, "interface cannot be type-resolved")
			return false
		}

	default:
		return false
	}

	return true
}

// Use accepts either a *gir.Class or a *gir.Interface; any other type will make
// it panic.
func (g *Generator) Use(typ interface{}) bool {
	g.Reset()

	if !g.init(typ) {
		return false
	}

	g.header.NeedsExternGLib()

	g.Methods.setMethods(g, g.methods)
	g.Virtuals.setVirtuals(g, g.virtuals)

	g.renameGetters(g.Methods)
	g.renameGetters(g.Virtuals)

	var signals []gir.Signal

	switch v := g.Root.Type.(type) {
	case *gir.Interface:
		signals = v.Signals
		g.checkTypeStruct(v.GLibTypeStruct)

	case *gir.Class:
		signals = v.Signals
		g.checkTypeStruct(v.GLibTypeStruct)

		for _, ctor := range v.Constructors {
			// Copy and bodge this so the constructors and stuff are named properly.
			// This copies things safely, so class is not modified.
			ctor := bodgeClassCtor(v, ctor)
			if !g.cgen.UseConstructor(&g.Root, &ctor.CallableAttrs) {
				continue
			}

			// file.ApplyHeader(g, &g.cgen)
			g.Constructors = append(g.Constructors, newMethod(&g.cgen))
		}
	}

	for _, sig := range signals {
		if !g.cgen.Use(&g.Root, &gir.CallableAttrs{
			Parameters:  sig.Parameters,
			ReturnValue: sig.ReturnValue,
		}) {
			g.Logln(logger.Debug, "skipping signal", sig.Name)
			continue
		}

		for _, res := range g.cgen.Results {
			res.Import(&g.header, true)
		}

		g.Signals = append(g.Signals, Signal{
			Name:         sig.Name,
			Tail:         callable.CoalesceTail(g.cgen.Tail),
			InfoElements: sig.InfoElements,
		})
	}

	return true
}

func (g *Generator) checkTypeStruct(girName string) {
	result := types.Find(g.gen, girName)
	if result == nil {
		return
	}

	if types.Filter(g.gen, result.Name(), result.CType()) {
		g.Logln(logger.Skip, "class/interface struct", result.Name())
		return
	}

	g.GLibTypeStruct = newTypeStruct(g, result)
}

func (g *Generator) ImplInterfaces() []string {
	impls := g.Tree.ImplInterfaces()
	names := make([]string, len(impls))

	for i, resolved := range impls {
		namespace := resolved.NeedsNamespace(g.gen.Namespace())
		if namespace {
			g.header.ImportPubl(resolved)
		}
		names[i] = resolved.PublicType(namespace)
	}

	return names
}

// IsInSameFile returns true if the given GIR item is in the same file. It's
// guessed using the InfoElements field in the given value.
func (g *Generator) IsInSameFile(v interface{}) bool {
	if g.InfoElements == nil {
		// Maybe?
		return true
	}

	ifields := cmt.GetInfoFields(v)
	if ifields.Elements == nil {
		return true
	}

	el1 := g.InfoElements
	el2 := ifields.Elements

	if el1.SourcePosition != nil && el2.SourcePosition != nil {
		return el1.SourcePosition.Filename == el2.SourcePosition.Filename
	}
	if el1.Doc != nil && el2.Doc != nil {
		return el1.Doc.Filename == el2.Doc.Filename
	}

	// Maybe?
	return true
}

// Wrap creates a wrapper around the given object variable.
func (g *Generator) Wrap(obj string) string {
	return g.Tree.Wrap(obj, &g.header)
}

func (g *Generator) renameGetters(methods Methods) {
	for i, method := range methods {
		newName, ok := callable.RenameGetter(method.Name)

		if g.hasMethod(newName) || methods.find(newName, i) {
			if ok {
				// Duplicate getter field; just skip renaming.
				continue
			}

			// Field is duplicate even when it wasn't renamed; work around this.
			newName += g.StructName
		}

		methods[i].Name = newName
	}
}

func (g *Generator) hasMethod(name string) bool {
	for _, parent := range g.Tree.Requires {
		if parent.Name() == name {
			return true
		}
	}

	// Don't override the Go name's base method, if any.
	if name == "Base"+g.StructName {
		return true
	}

	// Never override Object's methods.
	return types.IsObjectorMethod(name)
}

// bodgeClassCtor bodges the given constructor to return exactly the class type
// instead of any other. It returns the original ctor if the conditions don't
// match for bodging.
//
// We have to do this to work around some cases where widget constructors would
// return the widget class instead of the actual class.
func bodgeClassCtor(class *gir.Class, ctor gir.Constructor) gir.Constructor {
	if ctor.ReturnValue == nil || ctor.ReturnValue.Type == nil {
		return ctor
	}

	retVal := *ctor.ReturnValue
	retTyp := *retVal.AnyType.Type

	// Note: this has caused me quite a lot of trouble. It's probably wrong as
	// well. The whole point is to work around the C API's weird class typing.
	retTyp.CType = types.MoveCPtr(class.CType, retTyp.CType)

	retTyp.Name = class.Name
	retTyp.Introspectable = class.Introspectable

	retVal.AnyType.Type = &retTyp
	ctor.ReturnValue = &retVal

	name := ctor.Name
	if ctor.Shadows != "" {
		name = ctor.Shadows
		ctor.Shadows = "" // prevent callable from changing
	}

	name = strings.TrimPrefix(name, "new")
	name = strings.TrimPrefix(name, "_")
	if name != "" {
		name = "_" + name
	}

	name = "new_" + class.Name + name
	ctor.Name = name

	return ctor
}
