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
		{T: externglib.Type(C.gtk_shortcut_label_get_type()), F: marshalShortcutLabel},
	})
}

// ShortcutLabel is a widget that represents a single keyboard shortcut or
// gesture in the user interface.
type ShortcutLabel interface {
	Box
	Buildable
	Orientable

	// Accelerator retrieves the current accelerator of @self.
	Accelerator(s ShortcutLabel)
	// DisabledText retrieves the text that is displayed when no accelerator is
	// set.
	DisabledText(s ShortcutLabel)
	// SetAccelerator sets the accelerator to be displayed by @self.
	SetAccelerator(s ShortcutLabel, accelerator string)
	// SetDisabledText sets the text to be displayed by @self when no
	// accelerator is set.
	SetDisabledText(s ShortcutLabel, disabledText string)
}

// shortcutLabel implements the ShortcutLabel interface.
type shortcutLabel struct {
	Box
	Buildable
	Orientable
}

var _ ShortcutLabel = (*shortcutLabel)(nil)

// WrapShortcutLabel wraps a GObject to the right type. It is
// primarily used internally.
func WrapShortcutLabel(obj *externglib.Object) ShortcutLabel {
	return ShortcutLabel{
		Box:        WrapBox(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalShortcutLabel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapShortcutLabel(obj), nil
}

// NewShortcutLabel constructs a class ShortcutLabel.
func NewShortcutLabel(accelerator string) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(accelerator))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_shortcut_label_new(arg1)
}

// Accelerator retrieves the current accelerator of @self.
func (s shortcutLabel) Accelerator(s ShortcutLabel) {
	var arg0 *C.GtkShortcutLabel

	arg0 = (*C.GtkShortcutLabel)(unsafe.Pointer(s.Native()))

	C.gtk_shortcut_label_get_accelerator(arg0)
}

// DisabledText retrieves the text that is displayed when no accelerator is
// set.
func (s shortcutLabel) DisabledText(s ShortcutLabel) {
	var arg0 *C.GtkShortcutLabel

	arg0 = (*C.GtkShortcutLabel)(unsafe.Pointer(s.Native()))

	C.gtk_shortcut_label_get_disabled_text(arg0)
}

// SetAccelerator sets the accelerator to be displayed by @self.
func (s shortcutLabel) SetAccelerator(s ShortcutLabel, accelerator string) {
	var arg0 *C.GtkShortcutLabel
	var arg1 *C.gchar

	arg0 = (*C.GtkShortcutLabel)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(accelerator))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_shortcut_label_set_accelerator(arg0, arg1)
}

// SetDisabledText sets the text to be displayed by @self when no
// accelerator is set.
func (s shortcutLabel) SetDisabledText(s ShortcutLabel, disabledText string) {
	var arg0 *C.GtkShortcutLabel
	var arg1 *C.gchar

	arg0 = (*C.GtkShortcutLabel)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(disabledText))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_shortcut_label_set_disabled_text(arg0, arg1)
}
