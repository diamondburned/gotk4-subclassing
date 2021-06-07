// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_check_button_get_type()), F: marshalCheckButton},
	})
}

// CheckButton: a CheckButton places a label next to an indicator.
//
// CSS nodes
//
//    checkbutton[.text-button]
//    ├── check
//    ╰── [label]
//
// A CheckButton has a main node with name checkbutton. If the CheckButton:label
// property is set, it contains a label child. The indicator node is named check
// when no group is set, and radio if the checkbutton is grouped together with
// other checkbuttons.
//
//
// Accessibility
//
// GtkCheckButton uses the K_ACCESSIBLE_ROLE_CHECKBOX role.
type CheckButton interface {
	Widget
	Accessible
	Actionable
	Buildable
	ConstraintTarget

	// Active returns the current value of the CheckButton:active property.
	Active(s CheckButton) bool
	// Inconsistent returns whether the check button is in an inconsistent
	// state.
	Inconsistent(c CheckButton) bool
	// Label returns the label of the checkbutton.
	Label(s CheckButton)
	// UseUnderline returns the current value of the CheckButton:use-underline
	// property.
	UseUnderline(s CheckButton) bool
	// SetActive sets the new value of the CheckButton:active property. See also
	// gtk_check_button_get_active().
	//
	// Setting CheckButton:active to true will add the `:checked:` state to both
	// the checkbutton and the indicator CSS node.
	SetActive(s CheckButton, setting bool)
	// SetGroup adds @self to the group of @group. In a group of multiple check
	// buttons, only one button can be active at a time.
	//
	// Setting the group of a check button also changes the css name of the
	// indicator widget's CSS node to 'radio'.
	//
	// The behavior of a checkbutton in a group is also commonly known as a
	// 'radio button'.
	//
	// Note that the same effect can be achieved via the Actionable api, by
	// using the same action with parameter type and state type 's' for all
	// buttons in the group, and giving each button its own target value.
	SetGroup(s CheckButton, group CheckButton)
	// SetInconsistent: if the user has selected a range of elements (such as
	// some text or spreadsheet cells) that are affected by a check button, and
	// the current values in that range are inconsistent, you may want to
	// display the toggle in an "in between" state. Normally you would turn off
	// the inconsistent state again if the user checks the check button. This
	// has to be done manually, gtk_check_button_set_inconsistent only affects
	// visual appearance, not the semantics of the button.
	SetInconsistent(c CheckButton, inconsistent bool)
	// SetLabel sets the text of @self. If CheckButton:use-underline is true,
	// the underscore in @label is interpreted as mnemonic indicator, see
	// gtk_check_button_set_use_underline() for details on this behavior.
	SetLabel(s CheckButton, label string)
	// SetUseUnderline sets the new value of the CheckButton:use-underline
	// property. See also gtk_check_button_get_use_underline().
	//
	// If @setting is true, an underscore character in @self's label indicates a
	// mnemonic accelerator key. This behavior is similar to
	// Label:use-underline.
	SetUseUnderline(s CheckButton, setting bool)
}

// checkButton implements the CheckButton interface.
type checkButton struct {
	Widget
	Accessible
	Actionable
	Buildable
	ConstraintTarget
}

var _ CheckButton = (*checkButton)(nil)

// WrapCheckButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapCheckButton(obj *externglib.Object) CheckButton {
	return CheckButton{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Actionable:       WrapActionable(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalCheckButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCheckButton(obj), nil
}

// NewCheckButton constructs a class CheckButton.
func NewCheckButton() {
	C.gtk_check_button_new()
}

// NewCheckButtonWithLabel constructs a class CheckButton.
func NewCheckButtonWithLabel(label string) {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_check_button_new_with_label(arg1)
}

// NewCheckButtonWithMnemonic constructs a class CheckButton.
func NewCheckButtonWithMnemonic(label string) {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_check_button_new_with_mnemonic(arg1)
}

// Active returns the current value of the CheckButton:active property.
func (s checkButton) Active(s CheckButton) bool {
	var arg0 *C.GtkCheckButton

	arg0 = (*C.GtkCheckButton)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_check_button_get_active(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Inconsistent returns whether the check button is in an inconsistent
// state.
func (c checkButton) Inconsistent(c CheckButton) bool {
	var arg0 *C.GtkCheckButton

	arg0 = (*C.GtkCheckButton)(unsafe.Pointer(c.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_check_button_get_inconsistent(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Label returns the label of the checkbutton.
func (s checkButton) Label(s CheckButton) {
	var arg0 *C.GtkCheckButton

	arg0 = (*C.GtkCheckButton)(unsafe.Pointer(s.Native()))

	C.gtk_check_button_get_label(arg0)
}

// UseUnderline returns the current value of the CheckButton:use-underline
// property.
func (s checkButton) UseUnderline(s CheckButton) bool {
	var arg0 *C.GtkCheckButton

	arg0 = (*C.GtkCheckButton)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_check_button_get_use_underline(arg0)

	if cret {
		ok = true
	}

	return ok
}

// SetActive sets the new value of the CheckButton:active property. See also
// gtk_check_button_get_active().
//
// Setting CheckButton:active to true will add the `:checked:` state to both
// the checkbutton and the indicator CSS node.
func (s checkButton) SetActive(s CheckButton, setting bool) {
	var arg0 *C.GtkCheckButton
	var arg1 C.gboolean

	arg0 = (*C.GtkCheckButton)(unsafe.Pointer(s.Native()))
	if setting {
		arg1 = C.gboolean(1)
	}

	C.gtk_check_button_set_active(arg0, arg1)
}

// SetGroup adds @self to the group of @group. In a group of multiple check
// buttons, only one button can be active at a time.
//
// Setting the group of a check button also changes the css name of the
// indicator widget's CSS node to 'radio'.
//
// The behavior of a checkbutton in a group is also commonly known as a
// 'radio button'.
//
// Note that the same effect can be achieved via the Actionable api, by
// using the same action with parameter type and state type 's' for all
// buttons in the group, and giving each button its own target value.
func (s checkButton) SetGroup(s CheckButton, group CheckButton) {
	var arg0 *C.GtkCheckButton
	var arg1 *C.GtkCheckButton

	arg0 = (*C.GtkCheckButton)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkCheckButton)(unsafe.Pointer(group.Native()))

	C.gtk_check_button_set_group(arg0, arg1)
}

// SetInconsistent: if the user has selected a range of elements (such as
// some text or spreadsheet cells) that are affected by a check button, and
// the current values in that range are inconsistent, you may want to
// display the toggle in an "in between" state. Normally you would turn off
// the inconsistent state again if the user checks the check button. This
// has to be done manually, gtk_check_button_set_inconsistent only affects
// visual appearance, not the semantics of the button.
func (c checkButton) SetInconsistent(c CheckButton, inconsistent bool) {
	var arg0 *C.GtkCheckButton
	var arg1 C.gboolean

	arg0 = (*C.GtkCheckButton)(unsafe.Pointer(c.Native()))
	if inconsistent {
		arg1 = C.gboolean(1)
	}

	C.gtk_check_button_set_inconsistent(arg0, arg1)
}

// SetLabel sets the text of @self. If CheckButton:use-underline is true,
// the underscore in @label is interpreted as mnemonic indicator, see
// gtk_check_button_set_use_underline() for details on this behavior.
func (s checkButton) SetLabel(s CheckButton, label string) {
	var arg0 *C.GtkCheckButton
	var arg1 *C.char

	arg0 = (*C.GtkCheckButton)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_check_button_set_label(arg0, arg1)
}

// SetUseUnderline sets the new value of the CheckButton:use-underline
// property. See also gtk_check_button_get_use_underline().
//
// If @setting is true, an underscore character in @self's label indicates a
// mnemonic accelerator key. This behavior is similar to
// Label:use-underline.
func (s checkButton) SetUseUnderline(s CheckButton, setting bool) {
	var arg0 *C.GtkCheckButton
	var arg1 C.gboolean

	arg0 = (*C.GtkCheckButton)(unsafe.Pointer(s.Native()))
	if setting {
		arg1 = C.gboolean(1)
	}

	C.gtk_check_button_set_use_underline(arg0, arg1)
}