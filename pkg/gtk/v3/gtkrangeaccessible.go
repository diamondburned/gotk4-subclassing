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
		{T: externglib.Type(C.gtk_range_accessible_get_type()), F: marshalRangeAccessible},
	})
}

type RangeAccessible interface {
	WidgetAccessible
}

// rangeAccessible implements the RangeAccessible interface.
type rangeAccessible struct {
	WidgetAccessible
}

var _ RangeAccessible = (*rangeAccessible)(nil)

// WrapRangeAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapRangeAccessible(obj *externglib.Object) RangeAccessible {
	return RangeAccessible{
		WidgetAccessible: WrapWidgetAccessible(obj),
	}
}

func marshalRangeAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRangeAccessible(obj), nil
}
