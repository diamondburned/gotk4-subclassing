// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_hscale_get_type()), F: marshalHScale},
	})
}

// HScale: the HScale widget is used to allow the user to select a value using a
// horizontal slider. To create one, use gtk_hscale_new_with_range().
//
// The position to show the current value, and the number of decimal places
// shown can be set using the parent Scale class’s functions.
//
// GtkHScale has been deprecated, use Scale instead.
type HScale interface {
	Scale
	Buildable
	Orientable
}

// hScale implements the HScale interface.
type hScale struct {
	Scale
	Buildable
	Orientable
}

var _ HScale = (*hScale)(nil)

// WrapHScale wraps a GObject to the right type. It is
// primarily used internally.
func WrapHScale(obj *externglib.Object) HScale {
	return HScale{
		Scale:      WrapScale(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalHScale(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapHScale(obj), nil
}
