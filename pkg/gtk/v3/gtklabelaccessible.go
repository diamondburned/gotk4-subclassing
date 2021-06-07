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
		{T: externglib.Type(C.gtk_label_accessible_get_type()), F: marshalLabelAccessible},
	})
}

type LabelAccessible interface {
	WidgetAccessible
}

// labelAccessible implements the LabelAccessible interface.
type labelAccessible struct {
	WidgetAccessible
}

var _ LabelAccessible = (*labelAccessible)(nil)

// WrapLabelAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapLabelAccessible(obj *externglib.Object) LabelAccessible {
	return LabelAccessible{
		WidgetAccessible: WrapWidgetAccessible(obj),
	}
}

func marshalLabelAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapLabelAccessible(obj), nil
}