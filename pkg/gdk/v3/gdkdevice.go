// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

// TimeCoord: a TimeCoord stores a single event in a motion history.
type TimeCoord struct {
	native C.GdkTimeCoord
}

// WrapTimeCoord wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTimeCoord(ptr unsafe.Pointer) *TimeCoord {
	if ptr == nil {
		return nil
	}

	return (*TimeCoord)(ptr)
}

func marshalTimeCoord(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTimeCoord(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TimeCoord) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// Time gets the field inside the struct.
func (t *TimeCoord) Time() uint32 {
	var v uint32
	v = (uint32)(t.native.time)
	return v
}

// Axes gets the field inside the struct.
func (t *TimeCoord) Axes() [128]float64 {
	var v [128]float64
	v = *(*[128]float64)(unsafe.Pointer(t.native.axes))
	return v
}
