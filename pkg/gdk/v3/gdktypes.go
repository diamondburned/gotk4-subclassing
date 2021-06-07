// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_rectangle_get_type()), F: marshalRectangle},
	})
}

// Atom: an opaque type representing a string as an index into a table of
// strings on the X server.
type Atom struct {
	native C.GdkAtom
}

// WrapAtom wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAtom(ptr unsafe.Pointer) *Atom {
	if ptr == nil {
		return nil
	}

	return (*Atom)(ptr)
}

func marshalAtom(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapAtom(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *Atom) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// Name determines the string corresponding to an atom.
func (a *Atom) Name(a Atom) {
	var arg0 C.GdkAtom

	arg0 = (C.GdkAtom)(unsafe.Pointer(a.Native()))

	C.gdk_atom_name(arg0)
}

// Point defines the x and y coordinates of a point.
type Point struct {
	native C.GdkPoint
}

// WrapPoint wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPoint(ptr unsafe.Pointer) *Point {
	if ptr == nil {
		return nil
	}

	return (*Point)(ptr)
}

func marshalPoint(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapPoint(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (p *Point) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}

// X gets the field inside the struct.
func (p *Point) X() int {
	var v int
	v = int(p.native.x)
	return v
}

// Y gets the field inside the struct.
func (p *Point) Y() int {
	var v int
	v = int(p.native.y)
	return v
}

// Rectangle defines the position and size of a rectangle. It is identical to
// #cairo_rectangle_int_t.
type Rectangle struct {
	native C.GdkRectangle
}

// WrapRectangle wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRectangle(ptr unsafe.Pointer) *Rectangle {
	if ptr == nil {
		return nil
	}

	return (*Rectangle)(ptr)
}

func marshalRectangle(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRectangle(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *Rectangle) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// X gets the field inside the struct.
func (r *Rectangle) X() int {
	var v int
	v = int(r.native.x)
	return v
}

// Y gets the field inside the struct.
func (r *Rectangle) Y() int {
	var v int
	v = int(r.native.y)
	return v
}

// Width gets the field inside the struct.
func (r *Rectangle) Width() int {
	var v int
	v = int(r.native.width)
	return v
}

// Height gets the field inside the struct.
func (r *Rectangle) Height() int {
	var v int
	v = int(r.native.height)
	return v
}

// Equal checks if the two given rectangles are equal.
func (r *Rectangle) Equal(r *Rectangle, rect2 *Rectangle) bool {
	var arg0 *C.GdkRectangle
	var arg1 *C.GdkRectangle

	arg0 = (*C.GdkRectangle)(unsafe.Pointer(r.Native()))
	arg1 = (*C.GdkRectangle)(unsafe.Pointer(rect2.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gdk_rectangle_equal(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// Intersect calculates the intersection of two rectangles. It is allowed for
// @dest to be the same as either @src1 or @src2. If the rectangles do not
// intersect, @dest’s width and height is set to 0 and its x and y values are
// undefined. If you are only interested in whether the rectangles intersect,
// but not in the intersecting area itself, pass nil for @dest.
func (s *Rectangle) Intersect(s *Rectangle, src2 *Rectangle) (dest *Rectangle, ok bool) {
	var arg0 *C.GdkRectangle
	var arg1 *C.GdkRectangle

	arg0 = (*C.GdkRectangle)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GdkRectangle)(unsafe.Pointer(src2.Native()))

	var arg2 C.GdkRectangle
	var dest *Rectangle
	var cret C.gboolean
	var ok bool

	cret = C.gdk_rectangle_intersect(arg0, arg1, &arg2)

	dest = WrapRectangle(unsafe.Pointer(&arg2))
	if cret {
		ok = true
	}

	return dest, ok
}

// Union calculates the union of two rectangles. The union of rectangles @src1
// and @src2 is the smallest rectangle which includes both @src1 and @src2
// within it. It is allowed for @dest to be the same as either @src1 or @src2.
//
// Note that this function does not ignore 'empty' rectangles (ie. with zero
// width or height).
func (s *Rectangle) Union(s *Rectangle, src2 *Rectangle) *Rectangle {
	var arg0 *C.GdkRectangle
	var arg1 *C.GdkRectangle

	arg0 = (*C.GdkRectangle)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GdkRectangle)(unsafe.Pointer(src2.Native()))

	var arg2 C.GdkRectangle
	var dest *Rectangle

	C.gdk_rectangle_union(arg0, arg1, &arg2)

	dest = WrapRectangle(unsafe.Pointer(&arg2))

	return dest
}
