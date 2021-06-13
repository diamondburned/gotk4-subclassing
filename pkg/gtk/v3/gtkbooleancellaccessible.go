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
		{T: externglib.Type(C.gtk_boolean_cell_accessible_get_type()), F: marshalBooleanCellAccessible},
	})
}

type BooleanCellAccessible interface {
	RendererCellAccessible
}

// booleanCellAccessible implements the BooleanCellAccessible interface.
type booleanCellAccessible struct {
	RendererCellAccessible
}

var _ BooleanCellAccessible = (*booleanCellAccessible)(nil)

// WrapBooleanCellAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapBooleanCellAccessible(obj *externglib.Object) BooleanCellAccessible {
	return BooleanCellAccessible{
		RendererCellAccessible: WrapRendererCellAccessible(obj),
	}
}

func marshalBooleanCellAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBooleanCellAccessible(obj), nil
}
