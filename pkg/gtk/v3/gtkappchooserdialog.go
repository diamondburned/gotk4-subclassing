// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_app_chooser_dialog_get_type()), F: marshalAppChooserDialog},
	})
}

// AppChooserDialog shows a AppChooserWidget inside a Dialog.
//
// Note that AppChooserDialog does not have any interesting methods of its own.
// Instead, you should get the embedded AppChooserWidget using
// gtk_app_chooser_dialog_get_widget() and call its methods if the generic
// AppChooser interface is not sufficient for your needs.
//
// To set the heading that is shown above the AppChooserWidget, use
// gtk_app_chooser_dialog_set_heading().
type AppChooserDialog interface {
	Dialog
	AppChooser
	Buildable

	// Heading returns the text to display at the top of the dialog.
	Heading() string
	// SetHeading sets the text to display at the top of the dialog. If the
	// heading is not set, the dialog displays a default text.
	SetHeading(heading string)
}

// appChooserDialog implements the AppChooserDialog interface.
type appChooserDialog struct {
	Dialog
	AppChooser
	Buildable
}

var _ AppChooserDialog = (*appChooserDialog)(nil)

// WrapAppChooserDialog wraps a GObject to the right type. It is
// primarily used internally.
func WrapAppChooserDialog(obj *externglib.Object) AppChooserDialog {
	return AppChooserDialog{
		Dialog:     WrapDialog(obj),
		AppChooser: WrapAppChooser(obj),
		Buildable:  WrapBuildable(obj),
	}
}

func marshalAppChooserDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAppChooserDialog(obj), nil
}

// Heading returns the text to display at the top of the dialog.
func (s appChooserDialog) Heading() string {
	var _arg0 *C.GtkAppChooserDialog // out

	_arg0 = (*C.GtkAppChooserDialog)(unsafe.Pointer(s.Native()))

	var _cret *C.gchar // in

	_cret = C.gtk_app_chooser_dialog_get_heading(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// SetHeading sets the text to display at the top of the dialog. If the
// heading is not set, the dialog displays a default text.
func (s appChooserDialog) SetHeading(heading string) {
	var _arg0 *C.GtkAppChooserDialog // out
	var _arg1 *C.gchar               // out

	_arg0 = (*C.GtkAppChooserDialog)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(heading))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_app_chooser_dialog_set_heading(_arg0, _arg1)
}
