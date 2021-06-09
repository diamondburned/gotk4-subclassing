// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
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
	ColorSelection() Widget
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
func NewColorSelectionDialog(title string) ColorSelectionDialog {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkColorSelectionDialog

	cret = C.gtk_color_selection_dialog_new(arg1)

	var colorSelectionDialog ColorSelectionDialog

	colorSelectionDialog = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ColorSelectionDialog)

	return colorSelectionDialog
}

// ColorSelection retrieves the ColorSelection widget embedded in the
// dialog.
func (c colorSelectionDialog) ColorSelection() Widget {
	var arg0 *C.GtkColorSelectionDialog

	arg0 = (*C.GtkColorSelectionDialog)(unsafe.Pointer(c.Native()))

	var cret *C.GtkWidget

	cret = C.gtk_color_selection_dialog_get_color_selection(arg0)

	var widget Widget

	widget = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return widget
}
