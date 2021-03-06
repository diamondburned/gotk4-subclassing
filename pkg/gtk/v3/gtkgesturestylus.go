// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_stylus_get_type()), F: marshalGestureStylusser},
	})
}

// GestureStylusOverrider contains methods that are overridable.
type GestureStylusOverrider interface {
}

// GestureStylus is a Gesture implementation specific to stylus input. The
// provided signals just provide the basic information.
type GestureStylus struct {
	_ [0]func() // equal guard
	GestureSingle
}

var (
	_ Gesturer = (*GestureStylus)(nil)
)

func classInitGestureStylusser(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapGestureStylus(obj *externglib.Object) *GestureStylus {
	return &GestureStylus{
		GestureSingle: GestureSingle{
			Gesture: Gesture{
				EventController: EventController{
					Object: obj,
				},
			},
		},
	}
}

func marshalGestureStylusser(p uintptr) (interface{}, error) {
	return wrapGestureStylus(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (gesture *GestureStylus) ConnectDown(f func(object, p0 float64)) externglib.SignalHandle {
	return gesture.Connect("down", externglib.GeneratedClosure{Func: f})
}

func (gesture *GestureStylus) ConnectMotion(f func(object, p0 float64)) externglib.SignalHandle {
	return gesture.Connect("motion", externglib.GeneratedClosure{Func: f})
}

func (gesture *GestureStylus) ConnectProximity(f func(object, p0 float64)) externglib.SignalHandle {
	return gesture.Connect("proximity", externglib.GeneratedClosure{Func: f})
}

func (gesture *GestureStylus) ConnectUp(f func(object, p0 float64)) externglib.SignalHandle {
	return gesture.Connect("up", externglib.GeneratedClosure{Func: f})
}

// NewGestureStylus creates a new GestureStylus.
//
// The function takes the following parameters:
//
//    - widget: Widget.
//
// The function returns the following values:
//
//    - gestureStylus: newly created stylus gesture.
//
func NewGestureStylus(widget Widgetter) *GestureStylus {
	var _arg1 *C.GtkWidget  // out
	var _cret *C.GtkGesture // in

	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_gesture_stylus_new(_arg1)
	runtime.KeepAlive(widget)

	var _gestureStylus *GestureStylus // out

	_gestureStylus = wrapGestureStylus(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _gestureStylus
}

// Axis returns the current value for the requested axis. This function must be
// called from either the GestureStylus:down, GestureStylus:motion,
// GestureStylus:up or GestureStylus:proximity signals.
//
// The function takes the following parameters:
//
//    - axis: requested device axis.
//
// The function returns the following values:
//
//    - value: return location for the axis value.
//    - ok if there is a current value for the axis.
//
func (gesture *GestureStylus) Axis(axis gdk.AxisUse) (float64, bool) {
	var _arg0 *C.GtkGestureStylus // out
	var _arg1 C.GdkAxisUse        // out
	var _arg2 C.gdouble           // in
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkGestureStylus)(unsafe.Pointer(gesture.Native()))
	_arg1 = C.GdkAxisUse(axis)

	_cret = C.gtk_gesture_stylus_get_axis(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(axis)

	var _value float64 // out
	var _ok bool       // out

	_value = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

// DeviceTool returns the DeviceTool currently driving input through this
// gesture. This function must be called from either the GestureStylus::down,
// GestureStylus::motion, GestureStylus::up or GestureStylus::proximity signal
// handlers.
//
// The function returns the following values:
//
//    - deviceTool (optional): current stylus tool.
//
func (gesture *GestureStylus) DeviceTool() *gdk.DeviceTool {
	var _arg0 *C.GtkGestureStylus // out
	var _cret *C.GdkDeviceTool    // in

	_arg0 = (*C.GtkGestureStylus)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_stylus_get_device_tool(_arg0)
	runtime.KeepAlive(gesture)

	var _deviceTool *gdk.DeviceTool // out

	if _cret != nil {
		{
			obj := externglib.Take(unsafe.Pointer(_cret))
			_deviceTool = &gdk.DeviceTool{
				Object: obj,
			}
		}
	}

	return _deviceTool
}
