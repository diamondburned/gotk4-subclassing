// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

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
		{T: externglib.Type(C.gtk_border_get_type()), F: marshalBorder},
	})
}

// Border: struct that specifies a border around a rectangular area.
//
// Each side can have different width.
type Border struct {
	native C.GtkBorder
}

// WrapBorder wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapBorder(ptr unsafe.Pointer) *Border {
	return (*Border)(ptr)
}

func marshalBorder(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Border)(unsafe.Pointer(b)), nil
}

// NewBorder constructs a struct Border.
func NewBorder() *Border {
	var _cret *C.GtkBorder // in

	_cret = C.gtk_border_new()

	var _border *Border // out

	_border = (*Border)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_border, func(v *Border) {
		C.free(unsafe.Pointer(v))
	})

	return _border
}

// Native returns the underlying C source pointer.
func (b *Border) Native() unsafe.Pointer {
	return unsafe.Pointer(&b.native)
}

// Copy copies a Border-struct.
func (b *Border) Copy() *Border {
	var _arg0 *C.GtkBorder // out
	var _cret *C.GtkBorder // in

	_arg0 = (*C.GtkBorder)(unsafe.Pointer(b))

	_cret = C.gtk_border_copy(_arg0)

	var _border *Border // out

	_border = (*Border)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_border, func(v *Border) {
		C.free(unsafe.Pointer(v))
	})

	return _border
}

// Free frees a Border-struct.
func (b *Border) free() {
	var _arg0 *C.GtkBorder // out

	_arg0 = (*C.GtkBorder)(unsafe.Pointer(b))

	C.gtk_border_free(_arg0)
}
