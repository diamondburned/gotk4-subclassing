// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
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
		{T: externglib.Type(C.gtk_event_controller_get_type()), F: marshalEventController},
	})
}

// EventController is a base, low-level implementation for event controllers.
// Those react to a series of Events, and possibly trigger actions as a
// consequence of those.
type EventController interface {
	gextras.Objector

	// PropagationPhase gets the propagation phase at which @controller handles
	// events.
	PropagationPhase() PropagationPhase
	// Widget returns the Widget this controller relates to.
	Widget() Widget
	// Reset resets the @controller to a clean state. Every interaction the
	// controller did through EventController::handle-event will be dropped at
	// this point.
	Reset()
	// SetPropagationPhase sets the propagation phase at which a controller
	// handles events.
	//
	// If @phase is GTK_PHASE_NONE, no automatic event handling will be
	// performed, but other additional gesture maintenance will. In that phase,
	// the events can be managed by calling gtk_event_controller_handle_event().
	SetPropagationPhase(phase PropagationPhase)
}

// eventController implements the EventController interface.
type eventController struct {
	gextras.Objector
}

var _ EventController = (*eventController)(nil)

// WrapEventController wraps a GObject to the right type. It is
// primarily used internally.
func WrapEventController(obj *externglib.Object) EventController {
	return EventController{
		Objector: obj,
	}
}

func marshalEventController(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventController(obj), nil
}

// PropagationPhase gets the propagation phase at which @controller handles
// events.
func (c eventController) PropagationPhase() PropagationPhase {
	var arg0 *C.GtkEventController

	arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	var cret C.GtkPropagationPhase

	cret = C.gtk_event_controller_get_propagation_phase(arg0)

	var propagationPhase PropagationPhase

	propagationPhase = PropagationPhase(cret)

	return propagationPhase
}

// Widget returns the Widget this controller relates to.
func (c eventController) Widget() Widget {
	var arg0 *C.GtkEventController

	arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	var cret *C.GtkWidget

	cret = C.gtk_event_controller_get_widget(arg0)

	var widget Widget

	widget = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return widget
}

// Reset resets the @controller to a clean state. Every interaction the
// controller did through EventController::handle-event will be dropped at
// this point.
func (c eventController) Reset() {
	var arg0 *C.GtkEventController

	arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	C.gtk_event_controller_reset(arg0)
}

// SetPropagationPhase sets the propagation phase at which a controller
// handles events.
//
// If @phase is GTK_PHASE_NONE, no automatic event handling will be
// performed, but other additional gesture maintenance will. In that phase,
// the events can be managed by calling gtk_event_controller_handle_event().
func (c eventController) SetPropagationPhase(phase PropagationPhase) {
	var arg0 *C.GtkEventController
	var arg1 C.GtkPropagationPhase

	arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))
	arg1 = (C.GtkPropagationPhase)(phase)

	C.gtk_event_controller_set_propagation_phase(arg0, arg1)
}
