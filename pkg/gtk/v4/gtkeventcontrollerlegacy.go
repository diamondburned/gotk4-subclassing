// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_event_controller_legacy_get_type()), F: marshalEventControllerLegacy},
	})
}

// EventControllerLegacy is an event controller that gives you direct access to
// the event stream. It should only be used as a last resort if none of the
// other event controllers or gestures do the job.
type EventControllerLegacy interface {
	EventController
}

// eventControllerLegacy implements the EventControllerLegacy interface.
type eventControllerLegacy struct {
	EventController
}

var _ EventControllerLegacy = (*eventControllerLegacy)(nil)

// WrapEventControllerLegacy wraps a GObject to the right type. It is
// primarily used internally.
func WrapEventControllerLegacy(obj *externglib.Object) EventControllerLegacy {
	return EventControllerLegacy{
		EventController: WrapEventController(obj),
	}
}

func marshalEventControllerLegacy(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventControllerLegacy(obj), nil
}

// NewEventControllerLegacy constructs a class EventControllerLegacy.
func NewEventControllerLegacy() {
	C.gtk_event_controller_legacy_new()
}
