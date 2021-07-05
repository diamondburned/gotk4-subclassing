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
		{T: externglib.Type(C.gtk_link_button_accessible_get_type()), F: marshalLinkButtonAccessible},
	})
}

type LinkButtonAccessible interface {
	ButtonAccessible

	// AsAction casts the class to the atk.Action interface.
	AsAction() atk.Action
	// AsImage casts the class to the atk.Image interface.
	AsImage() atk.Image
}

// linkButtonAccessible implements the LinkButtonAccessible class.
type linkButtonAccessible struct {
	ButtonAccessible
}

// WrapLinkButtonAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapLinkButtonAccessible(obj *externglib.Object) LinkButtonAccessible {
	return linkButtonAccessible{
		ButtonAccessible: WrapButtonAccessible(obj),
	}
}

func marshalLinkButtonAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapLinkButtonAccessible(obj), nil
}

func (l linkButtonAccessible) AsAction() atk.Action {
	return atk.WrapAction(gextras.InternObject(l))
}

func (l linkButtonAccessible) AsImage() atk.Image {
	return atk.WrapImage(gextras.InternObject(l))
}