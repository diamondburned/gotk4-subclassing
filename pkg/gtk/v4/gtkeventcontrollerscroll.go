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
		{T: externglib.Type(C.gtk_event_controller_scroll_flags_get_type()), F: marshalEventControllerScrollFlags},
		{T: externglib.Type(C.gtk_event_controller_scroll_get_type()), F: marshalEventControllerScroll},
	})
}

// EventControllerScrollFlags describes the behavior of a
// `GtkEventControllerScroll`.
type EventControllerScrollFlags int

const (
	// EventControllerScrollFlagsNone: don't emit scroll.
	EventControllerScrollFlagsNone EventControllerScrollFlags = 0b0
	// EventControllerScrollFlagsVertical: emit scroll with vertical deltas.
	EventControllerScrollFlagsVertical EventControllerScrollFlags = 0b1
	// EventControllerScrollFlagsHorizontal: emit scroll with horizontal deltas.
	EventControllerScrollFlagsHorizontal EventControllerScrollFlags = 0b10
	// EventControllerScrollFlagsDiscrete: only emit deltas that are multiples
	// of 1.
	EventControllerScrollFlagsDiscrete EventControllerScrollFlags = 0b100
	// EventControllerScrollFlagsKinetic: emit ::decelerate after continuous
	// scroll finishes.
	EventControllerScrollFlagsKinetic EventControllerScrollFlags = 0b1000
	// EventControllerScrollFlagsBothAxes: emit scroll on both axes.
	EventControllerScrollFlagsBothAxes EventControllerScrollFlags = 0b11
)

func marshalEventControllerScrollFlags(p uintptr) (interface{}, error) {
	return EventControllerScrollFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// EventControllerScroll: `GtkEventControllerScroll` is an event controller that
// handles scroll events.
//
// It is capable of handling both discrete and continuous scroll events from
// mice or touchpads, abstracting them both with the
// [signal@Gtk.EventControllerScroll::scroll] signal. Deltas in the discrete
// case are multiples of 1.
//
// In the case of continuous scroll events, `GtkEventControllerScroll` encloses
// all [signal@Gtk.EventControllerScroll::scroll] emissions between two
// [signal@Gtk.EventControllerScroll::scroll-begin] and
// [signal@Gtk.EventControllerScroll::scroll-end] signals.
//
// The behavior of the event controller can be modified by the flags given at
// creation time, or modified at a later point through
// [method@Gtk.EventControllerScroll.set_flags] (e.g. because the scrolling
// conditions of the widget changed).
//
// The controller can be set up to emit motion for either/both vertical and
// horizontal scroll events through GTK_EVENT_CONTROLLER_SCROLL_VERTICAL,
// GTK_EVENT_CONTROLLER_SCROLL_HORIZONTAL and
// GTK_EVENT_CONTROLLER_SCROLL_BOTH_AXES. If any axis is disabled, the
// respective [signal@Gtk.EventControllerScroll::scroll] delta will be 0.
// Vertical scroll events will be translated to horizontal motion for the
// devices incapable of horizontal scrolling.
//
// The event controller can also be forced to emit discrete events on all
// devices through GTK_EVENT_CONTROLLER_SCROLL_DISCRETE. This can be used to
// implement discrete actions triggered through scroll events (e.g. switching
// across combobox options).
//
// The GTK_EVENT_CONTROLLER_SCROLL_KINETIC flag toggles the emission of the
// [signal@Gtk.EventControllerScroll::decelerate] signal, emitted at the end of
// scrolling with two X/Y velocity arguments that are consistent with the motion
// that was received.
type EventControllerScroll interface {
	gextras.Objector

	// Flags gets the flags conditioning the scroll controller behavior.
	Flags() EventControllerScrollFlags
}

// EventControllerScrollClass implements the EventControllerScroll interface.
type EventControllerScrollClass struct {
	EventControllerClass
}

var _ EventControllerScroll = (*EventControllerScrollClass)(nil)

func wrapEventControllerScroll(obj *externglib.Object) EventControllerScroll {
	return &EventControllerScrollClass{
		EventControllerClass: EventControllerClass{
			Object: obj,
		},
	}
}

func marshalEventControllerScroll(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEventControllerScroll(obj), nil
}

// Flags gets the flags conditioning the scroll controller behavior.
func (s *EventControllerScrollClass) Flags() EventControllerScrollFlags {
	var _arg0 *C.GtkEventControllerScroll     // out
	var _cret C.GtkEventControllerScrollFlags // in

	_arg0 = (*C.GtkEventControllerScroll)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_event_controller_scroll_get_flags(_arg0)

	var _eventControllerScrollFlags EventControllerScrollFlags // out

	_eventControllerScrollFlags = (EventControllerScrollFlags)(_cret)

	return _eventControllerScrollFlags
}
