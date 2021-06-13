// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_drawing_context_get_type()), F: marshalDrawingContext},
	})
}

// DrawingContext is an object that represents the current drawing state of a
// Window.
//
// It's possible to use a DrawingContext to draw on a Window via rendering API
// like Cairo or OpenGL.
//
// A DrawingContext can only be created by calling gdk_window_begin_draw_frame()
// and will be valid until a call to gdk_window_end_draw_frame().
//
// DrawingContext is available since GDK 3.22
type DrawingContext interface {
	gextras.Objector

	// IsValid checks whether the given DrawingContext is valid.
	IsValid() bool
}

// drawingContext implements the DrawingContext interface.
type drawingContext struct {
	gextras.Objector
}

var _ DrawingContext = (*drawingContext)(nil)

// WrapDrawingContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapDrawingContext(obj *externglib.Object) DrawingContext {
	return DrawingContext{
		Objector: obj,
	}
}

func marshalDrawingContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDrawingContext(obj), nil
}

// IsValid checks whether the given DrawingContext is valid.
func (c drawingContext) IsValid() bool {
	var _arg0 *C.GdkDrawingContext // out

	_arg0 = (*C.GdkDrawingContext)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean // in

	_cret = C.gdk_drawing_context_is_valid(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}
