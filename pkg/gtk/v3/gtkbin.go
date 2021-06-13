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
		{T: externglib.Type(C.gtk_bin_get_type()), F: marshalBin},
	})
}

// Bin: the Bin widget is a container with just one child. It is not very useful
// itself, but it is useful for deriving subclasses, since it provides common
// code needed for handling a single child widget.
//
// Many GTK+ widgets are subclasses of Bin, including Window, Button, Frame,
// HandleBox or ScrolledWindow.
type Bin interface {
	Container
	Buildable
}

// bin implements the Bin interface.
type bin struct {
	Container
	Buildable
}

var _ Bin = (*bin)(nil)

// WrapBin wraps a GObject to the right type. It is
// primarily used internally.
func WrapBin(obj *externglib.Object) Bin {
	return Bin{
		Container: WrapContainer(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalBin(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBin(obj), nil
}
