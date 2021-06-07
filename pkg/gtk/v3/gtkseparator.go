// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
		{T: externglib.Type(C.gtk_separator_get_type()), F: marshalSeparator},
	})
}

// Separator: gtkSeparator is a horizontal or vertical separator widget,
// depending on the value of the Orientable:orientation property, used to group
// the widgets within a window. It displays a line with a shadow to make it
// appear sunken into the interface.
//
//
// CSS nodes
//
// GtkSeparator has a single CSS node with name separator. The node gets one of
// the .horizontal or .vertical style classes.
type Separator interface {
	Widget
	Buildable
	Orientable
}

// separator implements the Separator interface.
type separator struct {
	Widget
	Buildable
	Orientable
}

var _ Separator = (*separator)(nil)

// WrapSeparator wraps a GObject to the right type. It is
// primarily used internally.
func WrapSeparator(obj *externglib.Object) Separator {
	return Separator{
		Widget:     WrapWidget(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalSeparator(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSeparator(obj), nil
}

// NewSeparator constructs a class Separator.
func NewSeparator(orientation Orientation) {
	var arg1 C.GtkOrientation

	arg1 = (C.GtkOrientation)(orientation)

	C.gtk_separator_new(arg1)
}