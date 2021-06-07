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
		{T: externglib.Type(C.gtk_plug_accessible_get_type()), F: marshalPlugAccessible},
	})
}

type PlugAccessible interface {
	WindowAccessible

	ID(p PlugAccessible)
}

// plugAccessible implements the PlugAccessible interface.
type plugAccessible struct {
	WindowAccessible
}

var _ PlugAccessible = (*plugAccessible)(nil)

// WrapPlugAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapPlugAccessible(obj *externglib.Object) PlugAccessible {
	return PlugAccessible{
		WindowAccessible: WrapWindowAccessible(obj),
	}
}

func marshalPlugAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPlugAccessible(obj), nil
}

func (p plugAccessible) ID(p PlugAccessible) {
	var arg0 *C.GtkPlugAccessible

	arg0 = (*C.GtkPlugAccessible)(unsafe.Pointer(p.Native()))

	C.gtk_plug_accessible_get_id(arg0)
}