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
		{T: externglib.Type(C.gtk_app_chooser_button_get_type()), F: marshalAppChooserButton},
	})
}

// AppChooserButton: the AppChooserButton is a widget that lets the user select
// an application. It implements the AppChooser interface.
//
// Initially, a AppChooserButton selects the first application in its list,
// which will either be the most-recently used application or, if
// AppChooserButton:show-default-item is true, the default application.
//
// The list of applications shown in a AppChooserButton includes the recommended
// applications for the given content type. When
// AppChooserButton:show-default-item is set, the default application is also
// included. To let the user chooser other applications, you can set the
// AppChooserButton:show-dialog-item property, which allows to open a full
// AppChooserDialog.
//
// It is possible to add custom items to the list, using
// gtk_app_chooser_button_append_custom_item(). These items cause the
// AppChooserButton::custom-item-activated signal to be emitted when they are
// selected.
//
// To track changes in the selected application, use the ComboBox::changed
// signal.
type AppChooserButton interface {
	ComboBox
	AppChooser
	Buildable
	CellEditable
	CellLayout

	// AppendCustomItem appends a custom item to the list of applications that
	// is shown in the popup; the item name must be unique per-widget. Clients
	// can use the provided name as a detail for the
	// AppChooserButton::custom-item-activated signal, to add a callback for the
	// activation of a particular custom item in the list. See also
	// gtk_app_chooser_button_append_separator().
	AppendCustomItem(s AppChooserButton, name string, label string, icon gio.Icon)
	// AppendSeparator appends a separator to the list of applications that is
	// shown in the popup.
	AppendSeparator(s AppChooserButton)
	// Heading returns the text to display at the top of the dialog.
	Heading(s AppChooserButton)
	// ShowDefaultItem returns the current value of the
	// AppChooserButton:show-default-item property.
	ShowDefaultItem(s AppChooserButton) bool
	// ShowDialogItem returns the current value of the
	// AppChooserButton:show-dialog-item property.
	ShowDialogItem(s AppChooserButton) bool
	// SetActiveCustomItem selects a custom item previously added with
	// gtk_app_chooser_button_append_custom_item().
	//
	// Use gtk_app_chooser_refresh() to bring the selection to its initial
	// state.
	SetActiveCustomItem(s AppChooserButton, name string)
	// SetHeading sets the text to display at the top of the dialog. If the
	// heading is not set, the dialog displays a default text.
	SetHeading(s AppChooserButton, heading string)
	// SetShowDefaultItem sets whether the dropdown menu of this button should
	// show the default application for the given content type at top.
	SetShowDefaultItem(s AppChooserButton, setting bool)
	// SetShowDialogItem sets whether the dropdown menu of this button should
	// show an entry to trigger a AppChooserDialog.
	SetShowDialogItem(s AppChooserButton, setting bool)
}

// appChooserButton implements the AppChooserButton interface.
type appChooserButton struct {
	ComboBox
	AppChooser
	Buildable
	CellEditable
	CellLayout
}

var _ AppChooserButton = (*appChooserButton)(nil)

// WrapAppChooserButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapAppChooserButton(obj *externglib.Object) AppChooserButton {
	return AppChooserButton{
		ComboBox:     WrapComboBox(obj),
		AppChooser:   WrapAppChooser(obj),
		Buildable:    WrapBuildable(obj),
		CellEditable: WrapCellEditable(obj),
		CellLayout:   WrapCellLayout(obj),
	}
}

func marshalAppChooserButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAppChooserButton(obj), nil
}

// NewAppChooserButton constructs a class AppChooserButton.
func NewAppChooserButton(contentType string) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(contentType))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_app_chooser_button_new(arg1)
}

// AppendCustomItem appends a custom item to the list of applications that
// is shown in the popup; the item name must be unique per-widget. Clients
// can use the provided name as a detail for the
// AppChooserButton::custom-item-activated signal, to add a callback for the
// activation of a particular custom item in the list. See also
// gtk_app_chooser_button_append_separator().
func (s appChooserButton) AppendCustomItem(s AppChooserButton, name string, label string, icon gio.Icon) {
	var arg0 *C.GtkAppChooserButton
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.GIcon

	arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.GIcon)(unsafe.Pointer(icon.Native()))

	C.gtk_app_chooser_button_append_custom_item(arg0, arg1, arg2, arg3)
}

// AppendSeparator appends a separator to the list of applications that is
// shown in the popup.
func (s appChooserButton) AppendSeparator(s AppChooserButton) {
	var arg0 *C.GtkAppChooserButton

	arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	C.gtk_app_chooser_button_append_separator(arg0)
}

// Heading returns the text to display at the top of the dialog.
func (s appChooserButton) Heading(s AppChooserButton) {
	var arg0 *C.GtkAppChooserButton

	arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	C.gtk_app_chooser_button_get_heading(arg0)
}

// ShowDefaultItem returns the current value of the
// AppChooserButton:show-default-item property.
func (s appChooserButton) ShowDefaultItem(s AppChooserButton) bool {
	var arg0 *C.GtkAppChooserButton

	arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_app_chooser_button_get_show_default_item(arg0)

	if cret {
		ok = true
	}

	return ok
}

// ShowDialogItem returns the current value of the
// AppChooserButton:show-dialog-item property.
func (s appChooserButton) ShowDialogItem(s AppChooserButton) bool {
	var arg0 *C.GtkAppChooserButton

	arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_app_chooser_button_get_show_dialog_item(arg0)

	if cret {
		ok = true
	}

	return ok
}

// SetActiveCustomItem selects a custom item previously added with
// gtk_app_chooser_button_append_custom_item().
//
// Use gtk_app_chooser_refresh() to bring the selection to its initial
// state.
func (s appChooserButton) SetActiveCustomItem(s AppChooserButton, name string) {
	var arg0 *C.GtkAppChooserButton
	var arg1 *C.gchar

	arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_app_chooser_button_set_active_custom_item(arg0, arg1)
}

// SetHeading sets the text to display at the top of the dialog. If the
// heading is not set, the dialog displays a default text.
func (s appChooserButton) SetHeading(s AppChooserButton, heading string) {
	var arg0 *C.GtkAppChooserButton
	var arg1 *C.gchar

	arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(heading))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_app_chooser_button_set_heading(arg0, arg1)
}

// SetShowDefaultItem sets whether the dropdown menu of this button should
// show the default application for the given content type at top.
func (s appChooserButton) SetShowDefaultItem(s AppChooserButton, setting bool) {
	var arg0 *C.GtkAppChooserButton
	var arg1 C.gboolean

	arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	if setting {
		arg1 = C.gboolean(1)
	}

	C.gtk_app_chooser_button_set_show_default_item(arg0, arg1)
}

// SetShowDialogItem sets whether the dropdown menu of this button should
// show an entry to trigger a AppChooserDialog.
func (s appChooserButton) SetShowDialogItem(s AppChooserButton, setting bool) {
	var arg0 *C.GtkAppChooserButton
	var arg1 C.gboolean

	arg0 = (*C.GtkAppChooserButton)(unsafe.Pointer(s.Native()))
	if setting {
		arg1 = C.gboolean(1)
	}

	C.gtk_app_chooser_button_set_show_dialog_item(arg0, arg1)
}