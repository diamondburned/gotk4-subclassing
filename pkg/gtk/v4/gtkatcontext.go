// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_at_context_get_type()), F: marshalATContext},
	})
}

// ATContext: `GtkATContext` is an abstract class provided by GTK to communicate
// to platform-specific assistive technologies API.
//
// Each platform supported by GTK implements a `GtkATContext` subclass, and is
// responsible for updating the accessible state in response to state changes in
// `GtkAccessible`.
type ATContext interface {
	gextras.Objector
}

// atContext implements the ATContext interface.
type atContext struct {
	gextras.Objector
}

var _ ATContext = (*atContext)(nil)

// WrapATContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapATContext(obj *externglib.Object) ATContext {
	return ATContext{
		Objector: obj,
	}
}

func marshalATContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapATContext(obj), nil
}
