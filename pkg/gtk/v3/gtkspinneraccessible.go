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
		{T: externglib.Type(C.gtk_spinner_accessible_get_type()), F: marshalSpinnerAccessible},
	})
}

type SpinnerAccessible interface {
	WidgetAccessible

	// AsImage casts the class to the atk.Image interface.
	AsImage() atk.Image
}

// spinnerAccessible implements the SpinnerAccessible class.
type spinnerAccessible struct {
	WidgetAccessible
}

// WrapSpinnerAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapSpinnerAccessible(obj *externglib.Object) SpinnerAccessible {
	return spinnerAccessible{
		WidgetAccessible: WrapWidgetAccessible(obj),
	}
}

func marshalSpinnerAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSpinnerAccessible(obj), nil
}

func (s spinnerAccessible) AsImage() atk.Image {
	return atk.WrapImage(gextras.InternObject(s))
}