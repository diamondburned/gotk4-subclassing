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
		{T: externglib.Type(C.gtk_orientable_get_type()), F: marshalOrientable},
	})
}

// Orientable: the `GtkOrientable` interface is implemented by all widgets that
// can be oriented horizontally or vertically.
//
// `GtkOrientable` is more flexible in that it allows the orientation to be
// changed at runtime, allowing the widgets to “flip”.
type Orientable interface {
	gextras.Objector

	// Orientation retrieves the orientation of the @orientable.
	Orientation() Orientation
}

// OrientableInterface implements the Orientable interface.
type OrientableInterface struct {
	*externglib.Object
}

var _ Orientable = (*OrientableInterface)(nil)

func wrapOrientable(obj *externglib.Object) Orientable {
	return &OrientableInterface{
		Object: obj,
	}
}

func marshalOrientable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapOrientable(obj), nil
}

// Orientation retrieves the orientation of the @orientable.
func (o *OrientableInterface) Orientation() Orientation {
	var _arg0 *C.GtkOrientable // out
	var _cret C.GtkOrientation // in

	_arg0 = (*C.GtkOrientable)(unsafe.Pointer(o.Native()))

	_cret = C.gtk_orientable_get_orientation(_arg0)

	var _orientation Orientation // out

	_orientation = (Orientation)(_cret)

	return _orientation
}
