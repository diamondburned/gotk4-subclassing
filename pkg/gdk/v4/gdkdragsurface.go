// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_drag_surface_get_type()), F: marshalDragSurface},
	})
}

// DragSurface: a DragSurface is an interface for surfaces used during DND.
type DragSurface interface {
	Surface

	// Present: present @drag_surface.
	Present(width int, height int) bool
}

// dragSurface implements the DragSurface interface.
type dragSurface struct {
	Surface
}

var _ DragSurface = (*dragSurface)(nil)

// WrapDragSurface wraps a GObject to a type that implements interface
// DragSurface. It is primarily used internally.
func WrapDragSurface(obj *externglib.Object) DragSurface {
	return DragSurface{
		Surface: WrapSurface(obj),
	}
}

func marshalDragSurface(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDragSurface(obj), nil
}

// Present: present @drag_surface.
func (d dragSurface) Present(width int, height int) bool {
	var _arg0 *C.GdkDragSurface // out
	var _arg1 C.int             // out
	var _arg2 C.int             // out

	_arg0 = (*C.GdkDragSurface)(unsafe.Pointer(d.Native()))
	_arg1 = C.int(width)
	_arg2 = C.int(height)

	var _cret C.gboolean // in

	_cret = C.gdk_drag_surface_present(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}
