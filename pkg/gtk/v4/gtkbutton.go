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
		{T: externglib.Type(C.gtk_button_get_type()), F: marshalButton},
	})
}

// Button: the Button widget is generally used to trigger a callback function
// that is called when the button is pressed. The various signals and how to use
// them are outlined below.
//
// The Button widget can hold any valid child widget. That is, it can hold
// almost any other standard Widget. The most commonly used child is the Label.
//
//
// CSS nodes
//
// GtkButton has a single CSS node with name button. The node will get the style
// classes .image-button or .text-button, if the content is just an image or
// label, respectively. It may also receive the .flat style class.
//
// Other style classes that are commonly used with GtkButton include
// .suggested-action and .destructive-action. In special cases, buttons can be
// made round by adding the .circular style class.
//
// Button-like widgets like ToggleButton, MenuButton, VolumeButton, LockButton,
// ColorButton or FontButton use style classes such as .toggle, .popup, .scale,
// .lock, .color on the button node to differentiate themselves from a plain
// GtkButton.
//
//
// Accessibility
//
// GtkButton uses the K_ACCESSIBLE_ROLE_BUTTON role.
type Button interface {
	Widget
	Accessible
	Actionable
	Buildable
	ConstraintTarget

	// Child gets the child widget of @button.
	Child(b Button)
	// HasFrame returns whether the button has a frame.
	HasFrame(b Button) bool
	// IconName returns the icon name set via gtk_button_set_icon_name().
	IconName(b Button)
	// Label fetches the text from the label of the button, as set by
	// gtk_button_set_label(). If the label text has not been set the return
	// value will be nil. This will be the case if you create an empty button
	// with gtk_button_new() to use as a container.
	Label(b Button)
	// UseUnderline returns whether an embedded underline in the button label
	// indicates a mnemonic. See gtk_button_set_use_underline().
	UseUnderline(b Button) bool
	// SetChild sets the child widget of @button.
	SetChild(b Button, child Widget)
	// SetHasFrame sets the style of the button. Buttons can has a flat
	// appearance or have a frame drawn around them.
	SetHasFrame(b Button, hasFrame bool)
	// SetIconName adds a Image with the given icon name as a child. If @button
	// already contains a child widget, that child widget will be removed and
	// replaced with the image.
	SetIconName(b Button, iconName string)
	// SetLabel sets the text of the label of the button to @label.
	//
	// This will also clear any previously set labels.
	SetLabel(b Button, label string)
	// SetUseUnderline: if true, an underline in the text of the button label
	// indicates the next character should be used for the mnemonic accelerator
	// key.
	SetUseUnderline(b Button, useUnderline bool)
}

// button implements the Button interface.
type button struct {
	Widget
	Accessible
	Actionable
	Buildable
	ConstraintTarget
}

var _ Button = (*button)(nil)

// WrapButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapButton(obj *externglib.Object) Button {
	return Button{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Actionable:       WrapActionable(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapButton(obj), nil
}

// NewButton constructs a class Button.
func NewButton() {
	C.gtk_button_new()
}

// NewButtonFromIconName constructs a class Button.
func NewButtonFromIconName(iconName string) {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(iconName))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_button_new_from_icon_name(arg1)
}

// NewButtonWithLabel constructs a class Button.
func NewButtonWithLabel(label string) {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_button_new_with_label(arg1)
}

// NewButtonWithMnemonic constructs a class Button.
func NewButtonWithMnemonic(label string) {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_button_new_with_mnemonic(arg1)
}

// Child gets the child widget of @button.
func (b button) Child(b Button) {
	var arg0 *C.GtkButton

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	C.gtk_button_get_child(arg0)
}

// HasFrame returns whether the button has a frame.
func (b button) HasFrame(b Button) bool {
	var arg0 *C.GtkButton

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_button_get_has_frame(arg0)

	if cret {
		ok = true
	}

	return ok
}

// IconName returns the icon name set via gtk_button_set_icon_name().
func (b button) IconName(b Button) {
	var arg0 *C.GtkButton

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	C.gtk_button_get_icon_name(arg0)
}

// Label fetches the text from the label of the button, as set by
// gtk_button_set_label(). If the label text has not been set the return
// value will be nil. This will be the case if you create an empty button
// with gtk_button_new() to use as a container.
func (b button) Label(b Button) {
	var arg0 *C.GtkButton

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	C.gtk_button_get_label(arg0)
}

// UseUnderline returns whether an embedded underline in the button label
// indicates a mnemonic. See gtk_button_set_use_underline().
func (b button) UseUnderline(b Button) bool {
	var arg0 *C.GtkButton

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_button_get_use_underline(arg0)

	if cret {
		ok = true
	}

	return ok
}

// SetChild sets the child widget of @button.
func (b button) SetChild(b Button, child Widget) {
	var arg0 *C.GtkButton
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_button_set_child(arg0, arg1)
}

// SetHasFrame sets the style of the button. Buttons can has a flat
// appearance or have a frame drawn around them.
func (b button) SetHasFrame(b Button, hasFrame bool) {
	var arg0 *C.GtkButton
	var arg1 C.gboolean

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))
	if hasFrame {
		arg1 = C.gboolean(1)
	}

	C.gtk_button_set_has_frame(arg0, arg1)
}

// SetIconName adds a Image with the given icon name as a child. If @button
// already contains a child widget, that child widget will be removed and
// replaced with the image.
func (b button) SetIconName(b Button, iconName string) {
	var arg0 *C.GtkButton
	var arg1 *C.char

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(iconName))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_button_set_icon_name(arg0, arg1)
}

// SetLabel sets the text of the label of the button to @label.
//
// This will also clear any previously set labels.
func (b button) SetLabel(b Button, label string) {
	var arg0 *C.GtkButton
	var arg1 *C.char

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_button_set_label(arg0, arg1)
}

// SetUseUnderline: if true, an underline in the text of the button label
// indicates the next character should be used for the mnemonic accelerator
// key.
func (b button) SetUseUnderline(b Button, useUnderline bool) {
	var arg0 *C.GtkButton
	var arg1 C.gboolean

	arg0 = (*C.GtkButton)(unsafe.Pointer(b.Native()))
	if useUnderline {
		arg1 = C.gboolean(1)
	}

	C.gtk_button_set_use_underline(arg0, arg1)
}