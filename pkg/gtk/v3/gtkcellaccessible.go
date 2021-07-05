// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_accessible_get_type()), F: marshalCellAccessible},
	})
}

type CellAccessible interface {
	Accessible

	// AsAction casts the class to the atk.Action interface.
	AsAction() atk.Action
}

// cellAccessible implements the CellAccessible class.
type cellAccessible struct {
	Accessible
}

// WrapCellAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapCellAccessible(obj *externglib.Object) CellAccessible {
	return cellAccessible{
		Accessible: WrapAccessible(obj),
	}
}

func marshalCellAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellAccessible(obj), nil
}

func (c cellAccessible) AsAction() atk.Action {
	return atk.WrapAction(gextras.InternObject(c))
}