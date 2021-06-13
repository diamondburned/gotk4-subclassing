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
		{T: externglib.Type(C.gtk_gesture_single_get_type()), F: marshalGestureSingle},
	})
}

// GestureSingle is a subclass of Gesture, optimized (although not restricted)
// for dealing with mouse and single-touch gestures. Under interaction, these
// gestures stick to the first interacting sequence, which is accessible through
// gtk_gesture_single_get_current_sequence() while the gesture is being
// interacted with.
//
// By default gestures react to both GDK_BUTTON_PRIMARY and touch events,
// gtk_gesture_single_set_touch_only() can be used to change the touch behavior.
// Callers may also specify a different mouse button number to interact with
// through gtk_gesture_single_set_button(), or react to any mouse button by
// setting 0. While the gesture is active, the button being currently pressed
// can be known through gtk_gesture_single_get_current_button().
type GestureSingle interface {
	Gesture

	// Button returns the button number @gesture listens for, or 0 if @gesture
	// reacts to any button press.
	Button() uint
	// CurrentButton returns the button number currently interacting with
	// @gesture, or 0 if there is none.
	CurrentButton() uint
	// Exclusive gets whether a gesture is exclusive. For more information, see
	// gtk_gesture_single_set_exclusive().
	Exclusive() bool
	// TouchOnly returns true if the gesture is only triggered by touch events.
	TouchOnly() bool
	// SetButton sets the button number @gesture listens to. If non-0, every
	// button press from a different button number will be ignored. Touch events
	// implicitly match with button 1.
	SetButton(button uint)
	// SetExclusive sets whether @gesture is exclusive. An exclusive gesture
	// will only handle pointer and "pointer emulated" touch events, so at any
	// given time, there is only one sequence able to interact with those.
	SetExclusive(exclusive bool)
	// SetTouchOnly: if @touch_only is true, @gesture will only handle events of
	// type K_TOUCH_BEGIN, K_TOUCH_UPDATE or K_TOUCH_END. If false, mouse events
	// will be handled too.
	SetTouchOnly(touchOnly bool)
}

// gestureSingle implements the GestureSingle interface.
type gestureSingle struct {
	Gesture
}

var _ GestureSingle = (*gestureSingle)(nil)

// WrapGestureSingle wraps a GObject to the right type. It is
// primarily used internally.
func WrapGestureSingle(obj *externglib.Object) GestureSingle {
	return GestureSingle{
		Gesture: WrapGesture(obj),
	}
}

func marshalGestureSingle(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureSingle(obj), nil
}

// Button returns the button number @gesture listens for, or 0 if @gesture
// reacts to any button press.
func (g gestureSingle) Button() uint {
	var _arg0 *C.GtkGestureSingle // out

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(g.Native()))

	var _cret C.guint // in

	_cret = C.gtk_gesture_single_get_button(_arg0)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// CurrentButton returns the button number currently interacting with
// @gesture, or 0 if there is none.
func (g gestureSingle) CurrentButton() uint {
	var _arg0 *C.GtkGestureSingle // out

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(g.Native()))

	var _cret C.guint // in

	_cret = C.gtk_gesture_single_get_current_button(_arg0)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// Exclusive gets whether a gesture is exclusive. For more information, see
// gtk_gesture_single_set_exclusive().
func (g gestureSingle) Exclusive() bool {
	var _arg0 *C.GtkGestureSingle // out

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(g.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_gesture_single_get_exclusive(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// TouchOnly returns true if the gesture is only triggered by touch events.
func (g gestureSingle) TouchOnly() bool {
	var _arg0 *C.GtkGestureSingle // out

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(g.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_gesture_single_get_touch_only(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// SetButton sets the button number @gesture listens to. If non-0, every
// button press from a different button number will be ignored. Touch events
// implicitly match with button 1.
func (g gestureSingle) SetButton(button uint) {
	var _arg0 *C.GtkGestureSingle // out
	var _arg1 C.guint             // out

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(g.Native()))
	_arg1 = C.guint(button)

	C.gtk_gesture_single_set_button(_arg0, _arg1)
}

// SetExclusive sets whether @gesture is exclusive. An exclusive gesture
// will only handle pointer and "pointer emulated" touch events, so at any
// given time, there is only one sequence able to interact with those.
func (g gestureSingle) SetExclusive(exclusive bool) {
	var _arg0 *C.GtkGestureSingle // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(g.Native()))
	if exclusive {
		_arg1 = C.gboolean(1)
	}

	C.gtk_gesture_single_set_exclusive(_arg0, _arg1)
}

// SetTouchOnly: if @touch_only is true, @gesture will only handle events of
// type K_TOUCH_BEGIN, K_TOUCH_UPDATE or K_TOUCH_END. If false, mouse events
// will be handled too.
func (g gestureSingle) SetTouchOnly(touchOnly bool) {
	var _arg0 *C.GtkGestureSingle // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkGestureSingle)(unsafe.Pointer(g.Native()))
	if touchOnly {
		_arg1 = C.gboolean(1)
	}

	C.gtk_gesture_single_set_touch_only(_arg0, _arg1)
}
