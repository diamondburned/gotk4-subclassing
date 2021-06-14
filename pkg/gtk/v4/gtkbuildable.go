// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_buildable_get_type()), F: marshalBuildable},
	})
}

// BuildableOverrider contains methods that are overridable. This
// interface is a subset of the interface Buildable.
type BuildableOverrider interface {
	// AddChild adds a child to @buildable. @type is an optional string
	// describing how the child should be added.
	AddChild(builder Builder, child gextras.Objector, typ string)

	ID() string
	// InternalChild retrieves the internal child called @childname of the
	// @buildable object.
	InternalChild(builder Builder, childname string) gextras.Objector

	ParserFinished(builder Builder)

	SetBuildableProperty(builder Builder, name string, value **externglib.Value)

	SetID(id string)
}

// Buildable: `GtkBuildable` allows objects to extend and customize their
// deserialization from ui files.
//
// The interface includes methods for setting names and properties of objects,
// parsing custom tags and constructing child objects.
//
// The `GtkBuildable` interface is implemented by all widgets and many of the
// non-widget objects that are provided by GTK. The main user of this interface
// is [class@Gtk.Builder]. There should be very little need for applications to
// call any of these functions directly.
//
// An object only needs to implement this interface if it needs to extend the
// `GtkBuilder` XML format or run any extra routines at deserialization time.
type Buildable interface {
	gextras.Objector
	BuildableOverrider

	// BuildableID gets the ID of the @buildable object.
	//
	// `GtkBuilder` sets the name based on the ID attribute of the <object> tag
	// used to construct the @buildable.
	BuildableID() string
}

// buildable implements the Buildable interface.
type buildable struct {
	gextras.Objector
}

var _ Buildable = (*buildable)(nil)

// WrapBuildable wraps a GObject to a type that implements interface
// Buildable. It is primarily used internally.
func WrapBuildable(obj *externglib.Object) Buildable {
	return Buildable{
		Objector: obj,
	}
}

func marshalBuildable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBuildable(obj), nil
}

// BuildableID gets the ID of the @buildable object.
//
// `GtkBuilder` sets the name based on the ID attribute of the <object> tag
// used to construct the @buildable.
func (b buildable) BuildableID() string {
	var _arg0 *C.GtkBuildable // out

	_arg0 = (*C.GtkBuildable)(unsafe.Pointer(b.Native()))

	var _cret *C.char // in

	_cret = C.gtk_buildable_get_buildable_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// BuildableParseContext: an opaque context struct for `GtkBuildableParser`.
type BuildableParseContext struct {
	native C.GtkBuildableParseContext
}

// WrapBuildableParseContext wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapBuildableParseContext(ptr unsafe.Pointer) *BuildableParseContext {
	if ptr == nil {
		return nil
	}

	return (*BuildableParseContext)(ptr)
}

// Native returns the underlying C source pointer.
func (b *BuildableParseContext) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}

// Element retrieves the name of the currently open element.
//
// If called from the start_element or end_element handlers this will give the
// element_name as passed to those functions. For the parent elements, see
// gtk_buildable_parse_context_get_element_stack().
func (c *BuildableParseContext) Element() string {
	var _arg0 *C.GtkBuildableParseContext // out

	_arg0 = (*C.GtkBuildableParseContext)(unsafe.Pointer(c.Native()))

	var _cret *C.char // in

	_cret = C.gtk_buildable_parse_context_get_element(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// ElementStack retrieves the element stack from the internal state of the
// parser.
//
// The returned Array is an array of strings where the last item is the
// currently open tag (as would be returned by
// gtk_buildable_parse_context_get_element()) and the previous item is its
// immediate parent.
//
// This function is intended to be used in the start_element and end_element
// handlers where gtk_buildable_parse_context_get_element() would merely return
// the name of the element that is being processed.
func (c *BuildableParseContext) ElementStack() []string {
	var _arg0 *C.GtkBuildableParseContext // out

	_arg0 = (*C.GtkBuildableParseContext)(unsafe.Pointer(c.Native()))

	var _cret *C.GPtrArray

	_cret = C.gtk_buildable_parse_context_get_element_stack(_arg0)

	var _utf8s []string

	{
		var i int
		for p := _cret; *p != nil; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
		}
	}

	return _utf8s
}

// Position retrieves the current line number and the number of the character on
// that line. Intended for use in error messages; there are no strict semantics
// for what constitutes the "current" line number other than "the best number we
// could come up with for error messages."
func (c *BuildableParseContext) Position() (lineNumber int, charNumber int) {
	var _arg0 *C.GtkBuildableParseContext // out

	_arg0 = (*C.GtkBuildableParseContext)(unsafe.Pointer(c.Native()))

	var _arg1 C.int // in
	var _arg2 C.int // in

	C.gtk_buildable_parse_context_get_position(_arg0, &_arg1, &_arg2)

	var _lineNumber int // out
	var _charNumber int // out

	_lineNumber = (int)(_arg1)
	_charNumber = (int)(_arg2)

	return _lineNumber, _charNumber
}
