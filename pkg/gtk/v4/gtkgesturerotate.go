// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_rotate_get_type()), F: marshalGestureRotate},
	})
}

// GestureRotate: `GtkGestureRotate` is a `GtkGesture` for 2-finger rotations.
//
// Whenever the angle between both handled sequences changes, the
// [signal@Gtk.GestureRotate::angle-changed] signal is emitted.
type GestureRotate interface {
	gextras.Objector

	// AngleDelta gets the angle delta in radians.
	//
	// If @gesture is active, this function returns the angle difference in
	// radians since the gesture was first recognized. If @gesture is not
	// active, 0 is returned.
	AngleDelta() float64
}

// GestureRotateClass implements the GestureRotate interface.
type GestureRotateClass struct {
	GestureClass
}

var _ GestureRotate = (*GestureRotateClass)(nil)

func wrapGestureRotate(obj *externglib.Object) GestureRotate {
	return &GestureRotateClass{
		GestureClass: GestureClass{
			EventControllerClass: EventControllerClass{
				Object: obj,
			},
		},
	}
}

func marshalGestureRotate(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGestureRotate(obj), nil
}

// NewGestureRotate returns a newly created `GtkGesture` that recognizes 2-touch
// rotation gestures.
func NewGestureRotate() *GestureRotateClass {
	var _cret *C.GtkGesture // in

	_cret = C.gtk_gesture_rotate_new()

	var _gestureRotate *GestureRotateClass // out

	_gestureRotate = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*GestureRotateClass)

	return _gestureRotate
}

// AngleDelta gets the angle delta in radians.
//
// If @gesture is active, this function returns the angle difference in radians
// since the gesture was first recognized. If @gesture is not active, 0 is
// returned.
func (g *GestureRotateClass) AngleDelta() float64 {
	var _arg0 *C.GtkGestureRotate // out
	var _cret C.double            // in

	_arg0 = (*C.GtkGestureRotate)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_rotate_get_angle_delta(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}
