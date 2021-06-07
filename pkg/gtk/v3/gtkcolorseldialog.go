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
		{T: externglib.Type(C.gtk_color_selection_dialog_get_type()), F: marshalColorSelectionDialog},
	})
}

type ColorSelectionDialog interface {
	Dialog
	Buildable

	// ColorSelection retrieves the ColorSelection widget embedded in the
	// dialog.
	ColorSelection(c ColorSelectionDialog)
}

// colorSelectionDialog implements the ColorSelectionDialog interface.
type colorSelectionDialog struct {
	Dialog
	Buildable
}

var _ ColorSelectionDialog = (*colorSelectionDialog)(nil)

// WrapColorSelectionDialog wraps a GObject to the right type. It is
// primarily used internally.
func WrapColorSelectionDialog(obj *externglib.Object) ColorSelectionDialog {
	return ColorSelectionDialog{
		Dialog:    WrapDialog(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalColorSelectionDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapColorSelectionDialog(obj), nil
}

// NewColorSelectionDialog constructs a class ColorSelectionDialog.
func NewColorSelectionDialog(title string) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_color_selection_dialog_new(arg1)
}

// ColorSelection retrieves the ColorSelection widget embedded in the
// dialog.
func (c colorSelectionDialog) ColorSelection(c ColorSelectionDialog) {
	var arg0 *C.GtkColorSelectionDialog

	arg0 = (*C.GtkColorSelectionDialog)(unsafe.Pointer(c.Native()))

	C.gtk_color_selection_dialog_get_color_selection(arg0)
}
