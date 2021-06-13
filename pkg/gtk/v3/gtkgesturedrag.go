// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_drag_get_type()), F: marshalGestureDrag},
	})
}

// GestureDrag is a Gesture implementation that recognizes drag operations. The
// drag operation itself can be tracked throught the GestureDrag::drag-begin,
// GestureDrag::drag-update and GestureDrag::drag-end signals, or the relevant
// coordinates be extracted through gtk_gesture_drag_get_offset() and
// gtk_gesture_drag_get_start_point().
type GestureDrag interface {
	GestureSingle

	// Offset: if the @gesture is active, this function returns true and fills
	// in @x and @y with the coordinates of the current point, as an offset to
	// the starting drag point.
	Offset() (x float64, y float64, ok bool)
	// StartPoint: if the @gesture is active, this function returns true and
	// fills in @x and @y with the drag start coordinates, in window-relative
	// coordinates.
	StartPoint() (x float64, y float64, ok bool)
}

// gestureDrag implements the GestureDrag interface.
type gestureDrag struct {
	GestureSingle
}

var _ GestureDrag = (*gestureDrag)(nil)

// WrapGestureDrag wraps a GObject to the right type. It is
// primarily used internally.
func WrapGestureDrag(obj *externglib.Object) GestureDrag {
	return GestureDrag{
		GestureSingle: WrapGestureSingle(obj),
	}
}

func marshalGestureDrag(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureDrag(obj), nil
}

// Offset: if the @gesture is active, this function returns true and fills
// in @x and @y with the coordinates of the current point, as an offset to
// the starting drag point.
func (g gestureDrag) Offset() (x float64, y float64, ok bool) {
	var _arg0 *C.GtkGestureDrag // out

	_arg0 = (*C.GtkGestureDrag)(unsafe.Pointer(g.Native()))

	var _arg1 C.gdouble  // in
	var _arg2 C.gdouble  // in
	var _cret C.gboolean // in

	_cret = C.gtk_gesture_drag_get_offset(_arg0, &_arg1, &_arg2)

	var _x float64 // out
	var _y float64 // out
	var _ok bool   // out

	_x = (float64)(_arg1)
	_y = (float64)(_arg2)
	if _cret {
		_ok = true
	}

	return _x, _y, _ok
}

// StartPoint: if the @gesture is active, this function returns true and
// fills in @x and @y with the drag start coordinates, in window-relative
// coordinates.
func (g gestureDrag) StartPoint() (x float64, y float64, ok bool) {
	var _arg0 *C.GtkGestureDrag // out

	_arg0 = (*C.GtkGestureDrag)(unsafe.Pointer(g.Native()))

	var _arg1 C.gdouble  // in
	var _arg2 C.gdouble  // in
	var _cret C.gboolean // in

	_cret = C.gtk_gesture_drag_get_start_point(_arg0, &_arg1, &_arg2)

	var _x float64 // out
	var _y float64 // out
	var _ok bool   // out

	_x = (float64)(_arg1)
	_y = (float64)(_arg2)
	if _cret {
		_ok = true
	}

	return _x, _y, _ok
}
