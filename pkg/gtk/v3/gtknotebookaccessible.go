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
		{T: externglib.Type(C.gtk_notebook_accessible_get_type()), F: marshalNotebookAccessible},
	})
}

type NotebookAccessible interface {
	ContainerAccessible
}

// notebookAccessible implements the NotebookAccessible interface.
type notebookAccessible struct {
	ContainerAccessible
}

var _ NotebookAccessible = (*notebookAccessible)(nil)

// WrapNotebookAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapNotebookAccessible(obj *externglib.Object) NotebookAccessible {
	return NotebookAccessible{
		ContainerAccessible: WrapContainerAccessible(obj),
	}
}

func marshalNotebookAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNotebookAccessible(obj), nil
}
