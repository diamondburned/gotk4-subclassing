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
		{T: externglib.Type(C.gtk_gesture_rotate_get_type()), F: marshalGestureRotate},
	})
}

// GestureRotate is a Gesture implementation able to recognize 2-finger
// rotations, whenever the angle between both handled sequences changes, the
// GestureRotate::angle-changed signal is emitted.
type GestureRotate interface {
	Gesture

	// AngleDelta: if @gesture is active, this function returns the angle
	// difference in radians since the gesture was first recognized. If @gesture
	// is not active, 0 is returned.
	AngleDelta() float64
}

// gestureRotate implements the GestureRotate interface.
type gestureRotate struct {
	Gesture
}

var _ GestureRotate = (*gestureRotate)(nil)

// WrapGestureRotate wraps a GObject to the right type. It is
// primarily used internally.
func WrapGestureRotate(obj *externglib.Object) GestureRotate {
	return GestureRotate{
		Gesture: WrapGesture(obj),
	}
}

func marshalGestureRotate(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureRotate(obj), nil
}

// AngleDelta: if @gesture is active, this function returns the angle
// difference in radians since the gesture was first recognized. If @gesture
// is not active, 0 is returned.
func (g gestureRotate) AngleDelta() float64 {
	var _arg0 *C.GtkGestureRotate // out

	_arg0 = (*C.GtkGestureRotate)(unsafe.Pointer(g.Native()))

	var _cret C.gdouble // in

	_cret = C.gtk_gesture_rotate_get_angle_delta(_arg0)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}
