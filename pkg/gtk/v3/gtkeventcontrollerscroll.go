// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_scroll_get_type()), F: marshalEventControllerScroll},
	})
}

// EventControllerScroll is an event controller meant to handle scroll events
// from mice and touchpads. It is capable of handling both discrete and
// continuous scroll events, abstracting them both on the
// EventControllerScroll::scroll signal (deltas in the discrete case are
// multiples of 1).
//
// In the case of continuous scroll events, EventControllerScroll encloses all
// EventControllerScroll::scroll events between two
// EventControllerScroll::scroll-begin and EventControllerScroll::scroll-end
// signals.
//
// The behavior of the event controller can be modified by the flags given at
// creation time, or modified at a later point through
// gtk_event_controller_scroll_set_flags() (e.g. because the scrolling
// conditions of the widget changed).
//
// The controller can be set up to emit motion for either/both vertical and
// horizontal scroll events through K_EVENT_CONTROLLER_SCROLL_VERTICAL,
// K_EVENT_CONTROLLER_SCROLL_HORIZONTAL and K_EVENT_CONTROLLER_SCROLL_BOTH. If
// any axis is disabled, the respective EventControllerScroll::scroll delta will
// be 0. Vertical scroll events will be translated to horizontal motion for the
// devices incapable of horizontal scrolling.
//
// The event controller can also be forced to emit discrete events on all
// devices through K_EVENT_CONTROLLER_SCROLL_DISCRETE. This can be used to
// implement discrete actions triggered through scroll events (e.g. switching
// across combobox options).
//
// The K_EVENT_CONTROLLER_SCROLL_KINETIC flag toggles the emission of the
// EventControllerScroll::decelerate signal, emitted at the end of scrolling
// with two X/Y velocity arguments that are consistent with the motion that was
// received.
//
// This object was added in 3.24.
type EventControllerScroll interface {
	EventController

	// Flags gets the flags conditioning the scroll controller behavior.
	Flags(c EventControllerScroll)
	// SetFlags sets the flags conditioning scroll controller behavior.
	SetFlags(c EventControllerScroll, flags EventControllerScrollFlags)
}

// eventControllerScroll implements the EventControllerScroll interface.
type eventControllerScroll struct {
	EventController
}

var _ EventControllerScroll = (*eventControllerScroll)(nil)

// WrapEventControllerScroll wraps a GObject to the right type. It is
// primarily used internally.
func WrapEventControllerScroll(obj *externglib.Object) EventControllerScroll {
	return EventControllerScroll{
		EventController: WrapEventController(obj),
	}
}

func marshalEventControllerScroll(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventControllerScroll(obj), nil
}

// NewEventControllerScroll constructs a class EventControllerScroll.
func NewEventControllerScroll(widget Widget, flags EventControllerScrollFlags) {
	var arg1 *C.GtkWidget
	var arg2 C.GtkEventControllerScrollFlags

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 = (C.GtkEventControllerScrollFlags)(flags)

	C.gtk_event_controller_scroll_new(arg1, arg2)
}

// Flags gets the flags conditioning the scroll controller behavior.
func (c eventControllerScroll) Flags(c EventControllerScroll) {
	var arg0 *C.GtkEventControllerScroll

	arg0 = (*C.GtkEventControllerScroll)(unsafe.Pointer(c.Native()))

	C.gtk_event_controller_scroll_get_flags(arg0)
}

// SetFlags sets the flags conditioning scroll controller behavior.
func (c eventControllerScroll) SetFlags(c EventControllerScroll, flags EventControllerScrollFlags) {
	var arg0 *C.GtkEventControllerScroll
	var arg1 C.GtkEventControllerScrollFlags

	arg0 = (*C.GtkEventControllerScroll)(unsafe.Pointer(c.Native()))
	arg1 = (C.GtkEventControllerScrollFlags)(flags)

	C.gtk_event_controller_scroll_set_flags(arg0, arg1)
}
