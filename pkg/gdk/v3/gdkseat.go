// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"github.com/diamondburned/gotk4/internal/box"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// SeatGrabPrepareFunc: type of the callback used to set up @window so it can be
// grabbed. A typical action would be ensuring the window is visible, although
// there's room for other initialization actions.
type SeatGrabPrepareFunc func(seat Seat, window Window)

//export gotk4_SeatGrabPrepareFunc
func gotk4_SeatGrabPrepareFunc(arg0 *C.GdkSeat, arg1 *C.GdkWindow, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(SeatGrabPrepareFunc)
	fn(seat, window, userData)
}