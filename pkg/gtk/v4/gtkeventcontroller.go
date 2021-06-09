// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
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

	// CurrentEvent returns the event that is currently being handled by the
	// controller, and nil at other times.
	CurrentEvent() gdk.Event
	// CurrentEventDevice returns the device of the event that is currently
	// being handled by the controller, and nil otherwise.
	CurrentEventDevice() gdk.Device
	// CurrentEventState returns the modifier state of the event that is
	// currently being handled by the controller, and 0 otherwise.
	CurrentEventState() gdk.ModifierType
	// CurrentEventTime returns the timestamp of the event that is currently
	// being handled by the controller, and 0 otherwise.
	CurrentEventTime() uint32
	// Name gets the name of @controller.
	Name() string
	// PropagationLimit gets the propagation limit of the event controller.
	PropagationLimit() PropagationLimit
	// PropagationPhase gets the propagation phase at which @controller handles
	// events.
	PropagationPhase() PropagationPhase
	// Widget returns the Widget this controller relates to.
	Widget() Widget
	// Reset resets the @controller to a clean state. Every interaction the
	// controller did through gtk_event_controller_handle_event() will be
	// dropped at this point.
	Reset()
	// SetName sets a name on the controller that can be used for debugging.
	SetName(name string)
	// SetPropagationLimit sets the event propagation limit on the event
	// controller.
	//
	// If the limit is set to GTK_LIMIT_SAME_NATIVE, the controller won't handle
	// events that are targeted at widgets on a different surface, such as
	// popovers.
	SetPropagationLimit(limit PropagationLimit)
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

// CurrentEvent returns the event that is currently being handled by the
// controller, and nil at other times.
func (c eventController) CurrentEvent() gdk.Event {
	var arg0 *C.GtkEventController

	arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	var cret *C.GdkEvent

	cret = C.gtk_event_controller_get_current_event(arg0)

	var event gdk.Event

	event = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdk.Event)

	return event
}

// CurrentEventDevice returns the device of the event that is currently
// being handled by the controller, and nil otherwise.
func (c eventController) CurrentEventDevice() gdk.Device {
	var arg0 *C.GtkEventController

	arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	var cret *C.GdkDevice

	cret = C.gtk_event_controller_get_current_event_device(arg0)

	var device gdk.Device

	device = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdk.Device)

	return device
}

// CurrentEventState returns the modifier state of the event that is
// currently being handled by the controller, and 0 otherwise.
func (c eventController) CurrentEventState() gdk.ModifierType {
	var arg0 *C.GtkEventController

	arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	var cret C.GdkModifierType

	cret = C.gtk_event_controller_get_current_event_state(arg0)

	var modifierType gdk.ModifierType

	modifierType = gdk.ModifierType(cret)

	return modifierType
}

// CurrentEventTime returns the timestamp of the event that is currently
// being handled by the controller, and 0 otherwise.
func (c eventController) CurrentEventTime() uint32 {
	var arg0 *C.GtkEventController

	arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	var cret C.guint32

	cret = C.gtk_event_controller_get_current_event_time(arg0)

	var guint32 uint32

	guint32 = (uint32)(cret)

	return guint32
}

// Name gets the name of @controller.
func (c eventController) Name() string {
	var arg0 *C.GtkEventController

	arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	var cret *C.char

	cret = C.gtk_event_controller_get_name(arg0)

	var utf8 string

	utf8 = C.GoString(cret)

	return utf8
}

// PropagationLimit gets the propagation limit of the event controller.
func (c eventController) PropagationLimit() PropagationLimit {
	var arg0 *C.GtkEventController

	arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	var cret C.GtkPropagationLimit

	cret = C.gtk_event_controller_get_propagation_limit(arg0)

	var propagationLimit PropagationLimit

	propagationLimit = PropagationLimit(cret)

	return propagationLimit
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
// controller did through gtk_event_controller_handle_event() will be
// dropped at this point.
func (c eventController) Reset() {
	var arg0 *C.GtkEventController

	arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))

	C.gtk_event_controller_reset(arg0)
}

// SetName sets a name on the controller that can be used for debugging.
func (c eventController) SetName(name string) {
	var arg0 *C.GtkEventController
	var arg1 *C.char

	arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))
	arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_event_controller_set_name(arg0, arg1)
}

// SetPropagationLimit sets the event propagation limit on the event
// controller.
//
// If the limit is set to GTK_LIMIT_SAME_NATIVE, the controller won't handle
// events that are targeted at widgets on a different surface, such as
// popovers.
func (c eventController) SetPropagationLimit(limit PropagationLimit) {
	var arg0 *C.GtkEventController
	var arg1 C.GtkPropagationLimit

	arg0 = (*C.GtkEventController)(unsafe.Pointer(c.Native()))
	arg1 = (C.GtkPropagationLimit)(limit)

	C.gtk_event_controller_set_propagation_limit(arg0, arg1)
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
