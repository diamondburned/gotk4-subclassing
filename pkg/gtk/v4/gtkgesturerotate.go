// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_rotate_get_type()), F: marshalGestureRotater},
	})
}

// GestureRotateOverrider contains methods that are overridable.
type GestureRotateOverrider interface {
}

// GestureRotate: GtkGestureRotate is a GtkGesture for 2-finger rotations.
//
// Whenever the angle between both handled sequences changes, the
// gtk.GestureRotate::angle-changed signal is emitted.
type GestureRotate struct {
	_ [0]func() // equal guard
	Gesture
}

var (
	_ Gesturer = (*GestureRotate)(nil)
)

func classInitGestureRotater(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapGestureRotate(obj *externglib.Object) *GestureRotate {
	return &GestureRotate{
		Gesture: Gesture{
			EventController: EventController{
				Object: obj,
			},
		},
	}
}

func marshalGestureRotater(p uintptr) (interface{}, error) {
	return wrapGestureRotate(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectAngleChanged: emitted when the angle between both tracked points
// changes.
func (gesture *GestureRotate) ConnectAngleChanged(f func(angle, angleDelta float64)) externglib.SignalHandle {
	return gesture.Connect("angle-changed", externglib.GeneratedClosure{Func: f})
}

// NewGestureRotate returns a newly created GtkGesture that recognizes 2-touch
// rotation gestures.
//
// The function returns the following values:
//
//    - gestureRotate: newly created GtkGestureRotate.
//
func NewGestureRotate() *GestureRotate {
	var _cret *C.GtkGesture // in

	_cret = C.gtk_gesture_rotate_new()

	var _gestureRotate *GestureRotate // out

	_gestureRotate = wrapGestureRotate(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureRotate
}

// AngleDelta gets the angle delta in radians.
//
// If gesture is active, this function returns the angle difference in radians
// since the gesture was first recognized. If gesture is not active, 0 is
// returned.
//
// The function returns the following values:
//
//    - gdouble: angle delta in radians.
//
func (gesture *GestureRotate) AngleDelta() float64 {
	var _arg0 *C.GtkGestureRotate // out
	var _cret C.double            // in

	_arg0 = (*C.GtkGestureRotate)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_rotate_get_angle_delta(_arg0)
	runtime.KeepAlive(gesture)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}
