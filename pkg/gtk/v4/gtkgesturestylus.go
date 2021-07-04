// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
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
		{T: externglib.Type(C.gtk_gesture_stylus_get_type()), F: marshalGestureStylus},
	})
}

// GestureStylus: `GtkGestureStylus` is a `GtkGesture` specific to stylus input.
//
// The provided signals just relay the basic information of the stylus events.
type GestureStylus interface {
	GestureSingle

	// Axis:
	Axis(axis gdk.AxisUse) (float64, bool)
	// Backlog:
	Backlog() ([]gdk.TimeCoord, bool)
	// DeviceTool:
	DeviceTool() gdk.DeviceTool
}

// gestureStylus implements the GestureStylus class.
type gestureStylus struct {
	GestureSingle
}

// WrapGestureStylus wraps a GObject to the right type. It is
// primarily used internally.
func WrapGestureStylus(obj *externglib.Object) GestureStylus {
	return gestureStylus{
		GestureSingle: WrapGestureSingle(obj),
	}
}

func marshalGestureStylus(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureStylus(obj), nil
}

// NewGestureStylus:
func NewGestureStylus() GestureStylus {
	var _cret *C.GtkGesture // in

	_cret = C.gtk_gesture_stylus_new()

	var _gestureStylus GestureStylus // out

	_gestureStylus = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(GestureStylus)

	return _gestureStylus
}

func (g gestureStylus) Axis(axis gdk.AxisUse) (float64, bool) {
	var _arg0 *C.GtkGestureStylus // out
	var _arg1 C.GdkAxisUse        // out
	var _arg2 C.double            // in
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkGestureStylus)(unsafe.Pointer(g.Native()))
	_arg1 = C.GdkAxisUse(axis)

	_cret = C.gtk_gesture_stylus_get_axis(_arg0, _arg1, &_arg2)

	var _value float64 // out
	var _ok bool       // out

	_value = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

func (g gestureStylus) Backlog() ([]gdk.TimeCoord, bool) {
	var _arg0 *C.GtkGestureStylus // out
	var _arg1 *C.GdkTimeCoord
	var _arg2 C.guint    // in
	var _cret C.gboolean // in

	_arg0 = (*C.GtkGestureStylus)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_stylus_get_backlog(_arg0, &_arg1, &_arg2)

	var _backlog []gdk.TimeCoord
	var _ok bool // out

	_backlog = unsafe.Slice((*gdk.TimeCoord)(unsafe.Pointer(_arg1)), _arg2)
	runtime.SetFinalizer(&_backlog, func(v *[]gdk.TimeCoord) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})
	if _cret != 0 {
		_ok = true
	}

	return _backlog, _ok
}

func (g gestureStylus) DeviceTool() gdk.DeviceTool {
	var _arg0 *C.GtkGestureStylus // out
	var _cret *C.GdkDeviceTool    // in

	_arg0 = (*C.GtkGestureStylus)(unsafe.Pointer(g.Native()))

	_cret = C.gtk_gesture_stylus_get_device_tool(_arg0)

	var _deviceTool gdk.DeviceTool // out

	_deviceTool = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdk.DeviceTool)

	return _deviceTool
}