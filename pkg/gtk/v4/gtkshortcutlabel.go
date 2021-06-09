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
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcut_label_get_type()), F: marshalShortcutLabel},
	})
}

// ShortcutLabel is a widget that represents a single keyboard shortcut or
// gesture in the user interface.
type ShortcutLabel interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget

	// Accelerator retrieves the current accelerator of @self.
	Accelerator() string
	// DisabledText retrieves the text that is displayed when no accelerator is
	// set.
	DisabledText() string
	// SetAccelerator sets the accelerator to be displayed by @self.
	SetAccelerator(accelerator string)
	// SetDisabledText sets the text to be displayed by @self when no
	// accelerator is set.
	SetDisabledText(disabledText string)
}

// shortcutLabel implements the ShortcutLabel interface.
type shortcutLabel struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
}

var _ ShortcutLabel = (*shortcutLabel)(nil)

// WrapShortcutLabel wraps a GObject to the right type. It is
// primarily used internally.
func WrapShortcutLabel(obj *externglib.Object) ShortcutLabel {
	return ShortcutLabel{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalShortcutLabel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapShortcutLabel(obj), nil
}

// NewShortcutLabel constructs a class ShortcutLabel.
func NewShortcutLabel(accelerator string) ShortcutLabel {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(accelerator))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkShortcutLabel

	cret = C.gtk_shortcut_label_new(arg1)

	var shortcutLabel ShortcutLabel

	shortcutLabel = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ShortcutLabel)

	return shortcutLabel
}

// Accelerator retrieves the current accelerator of @self.
func (s shortcutLabel) Accelerator() string {
	var arg0 *C.GtkShortcutLabel

	arg0 = (*C.GtkShortcutLabel)(unsafe.Pointer(s.Native()))

	var cret *C.char

	cret = C.gtk_shortcut_label_get_accelerator(arg0)

	var utf8 string

	utf8 = C.GoString(cret)

	return utf8
}

// DisabledText retrieves the text that is displayed when no accelerator is
// set.
func (s shortcutLabel) DisabledText() string {
	var arg0 *C.GtkShortcutLabel

	arg0 = (*C.GtkShortcutLabel)(unsafe.Pointer(s.Native()))

	var cret *C.char

	cret = C.gtk_shortcut_label_get_disabled_text(arg0)

	var utf8 string

	utf8 = C.GoString(cret)

	return utf8
}

// SetAccelerator sets the accelerator to be displayed by @self.
func (s shortcutLabel) SetAccelerator(accelerator string) {
	var arg0 *C.GtkShortcutLabel
	var arg1 *C.char

	arg0 = (*C.GtkShortcutLabel)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(accelerator))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_shortcut_label_set_accelerator(arg0, arg1)
}

// SetDisabledText sets the text to be displayed by @self when no
// accelerator is set.
func (s shortcutLabel) SetDisabledText(disabledText string) {
	var arg0 *C.GtkShortcutLabel
	var arg1 *C.char

	arg0 = (*C.GtkShortcutLabel)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(disabledText))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_shortcut_label_set_disabled_text(arg0, arg1)
}
