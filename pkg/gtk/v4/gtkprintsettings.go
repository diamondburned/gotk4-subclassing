// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

type PrintSettingsFunc func(key string, value string)

//export gotk4_PrintSettingsFunc
func gotk4_PrintSettingsFunc(arg0 *C.char, arg1 *C.char, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(PrintSettingsFunc)
	fn(key, value, userData)
}

// PageRange: see also gtk_print_settings_set_page_ranges().
type PageRange struct {
	native C.GtkPageRange
}

// WrapPageRange wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPageRange(ptr unsafe.Pointer) *PageRange {
	if ptr == nil {
		return nil
	}

	return (*PageRange)(ptr)
}

func marshalPageRange(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapPageRange(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (p *PageRange) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}

// Start gets the field inside the struct.
func (p *PageRange) Start() int {
	var v int
	v = int(p.native.start)
	return v
}

// End gets the field inside the struct.
func (p *PageRange) End() int {
	var v int
	v = int(p.native.end)
	return v
}
