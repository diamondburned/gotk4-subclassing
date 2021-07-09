// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_seat_capabilities_get_type()), F: marshalSeatCapabilities},
		{T: externglib.Type(C.gdk_seat_get_type()), F: marshalSeat},
	})
}

// SeatCapabilities flags describing the seat capabilities.
type SeatCapabilities int

const (
	// SeatCapabilitiesNone: no input capabilities
	SeatCapabilitiesNone SeatCapabilities = 0b0
	// SeatCapabilitiesPointer: the seat has a pointer (e.g. mouse)
	SeatCapabilitiesPointer SeatCapabilities = 0b1
	// SeatCapabilitiesTouch: the seat has touchscreen(s) attached
	SeatCapabilitiesTouch SeatCapabilities = 0b10
	// SeatCapabilitiesTabletStylus: the seat has drawing tablet(s) attached
	SeatCapabilitiesTabletStylus SeatCapabilities = 0b100
	// SeatCapabilitiesKeyboard: the seat has keyboard(s) attached
	SeatCapabilitiesKeyboard SeatCapabilities = 0b1000
	// SeatCapabilitiesAllPointing: the union of all pointing capabilities
	SeatCapabilitiesAllPointing SeatCapabilities = 0b111
	// SeatCapabilitiesAll: the union of all capabilities
	SeatCapabilitiesAll SeatCapabilities = 0b1111
)

func marshalSeatCapabilities(p uintptr) (interface{}, error) {
	return SeatCapabilities(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// SeatGrabPrepareFunc: type of the callback used to set up @window so it can be
// grabbed. A typical action would be ensuring the window is visible, although
// there's room for other initialization actions.
type SeatGrabPrepareFunc func(seat *SeatClass, window *WindowClass, userData interface{})

//export gotk4_SeatGrabPrepareFunc
func gotk4_SeatGrabPrepareFunc(arg0 *C.GdkSeat, arg1 *C.GdkWindow, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var seat *SeatClass      // out
	var window *WindowClass  // out
	var userData interface{} // out

	seat = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*SeatClass)
	window = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg1)))).(*WindowClass)
	userData = box.Get(uintptr(arg2))

	fn := v.(SeatGrabPrepareFunc)
	fn(seat, window, userData)
}

// Seat: the Seat object represents a collection of input devices that belong to
// a user.
type Seat interface {
	gextras.Objector

	// Capabilities returns the capabilities this Seat currently has.
	Capabilities() SeatCapabilities
	// Display returns the Display this seat belongs to.
	Display() *DisplayClass
	// Keyboard returns the master device that routes keyboard events.
	Keyboard() *DeviceClass
	// Pointer returns the master device that routes pointer events.
	Pointer() *DeviceClass
	// Ungrab releases a grab added through gdk_seat_grab().
	Ungrab()
}

// SeatClass implements the Seat interface.
type SeatClass struct {
	*externglib.Object
}

var _ Seat = (*SeatClass)(nil)

func wrapSeat(obj *externglib.Object) Seat {
	return &SeatClass{
		Object: obj,
	}
}

func marshalSeat(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSeat(obj), nil
}

// Capabilities returns the capabilities this Seat currently has.
func (s *SeatClass) Capabilities() SeatCapabilities {
	var _arg0 *C.GdkSeat            // out
	var _cret C.GdkSeatCapabilities // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_seat_get_capabilities(_arg0)

	var _seatCapabilities SeatCapabilities // out

	_seatCapabilities = (SeatCapabilities)(_cret)

	return _seatCapabilities
}

// Display returns the Display this seat belongs to.
func (s *SeatClass) Display() *DisplayClass {
	var _arg0 *C.GdkSeat    // out
	var _cret *C.GdkDisplay // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_seat_get_display(_arg0)

	var _display *DisplayClass // out

	_display = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*DisplayClass)

	return _display
}

// Keyboard returns the master device that routes keyboard events.
func (s *SeatClass) Keyboard() *DeviceClass {
	var _arg0 *C.GdkSeat   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_seat_get_keyboard(_arg0)

	var _device *DeviceClass // out

	_device = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*DeviceClass)

	return _device
}

// Pointer returns the master device that routes pointer events.
func (s *SeatClass) Pointer() *DeviceClass {
	var _arg0 *C.GdkSeat   // out
	var _cret *C.GdkDevice // in

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(s.Native()))

	_cret = C.gdk_seat_get_pointer(_arg0)

	var _device *DeviceClass // out

	_device = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*DeviceClass)

	return _device
}

// Ungrab releases a grab added through gdk_seat_grab().
func (s *SeatClass) Ungrab() {
	var _arg0 *C.GdkSeat // out

	_arg0 = (*C.GdkSeat)(unsafe.Pointer(s.Native()))

	C.gdk_seat_ungrab(_arg0)
}
