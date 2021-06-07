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
		{T: externglib.Type(C.gtk_password_entry_get_type()), F: marshalPasswordEntry},
	})
}

// PasswordEntry is entry that has been tailored for entering secrets. It does
// not show its contents in clear text, does not allow to copy it to the
// clipboard, and it shows a warning when Caps Lock is engaged. If the
// underlying platform allows it, GtkPasswordEntry will also place the text in a
// non-pageable memory area, to avoid it being written out to disk by the
// operating system.
//
// Optionally, it can offer a way to reveal the contents in clear text.
//
// GtkPasswordEntry provides only minimal API and should be used with the
// Editable API.
//
// CSS Nodes
//
//    entry.password
//    ╰── text
//        ├── image.caps-lock-indicator
//        ┊
//
// GtkPasswordEntry has a single CSS node with name entry that carries a
// .passwordstyle class. The text Css node below it has a child with name image
// and style class .caps-lock-indicator for the Caps Lock icon, and possibly
// other children.
//
//
// Accessibility
//
// GtkPasswordEntry uses the K_ACCESSIBLE_ROLE_TEXT_BOX role.
type PasswordEntry interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Editable

	// ExtraMenu gets the menu model set with
	// gtk_password_entry_set_extra_menu().
	ExtraMenu(e PasswordEntry)
	// ShowPeekIcon returns whether the entry is showing a clickable icon to
	// reveal the contents of the entry in clear text.
	ShowPeekIcon(e PasswordEntry) bool
	// SetExtraMenu sets a menu model to add when constructing the context menu
	// for @entry.
	SetExtraMenu(e PasswordEntry, model gio.MenuModel)
	// SetShowPeekIcon sets whether the entry should have a clickable icon to
	// show the contents of the entry in clear text.
	//
	// Setting this to false also hides the text again.
	SetShowPeekIcon(e PasswordEntry, showPeekIcon bool)
}

// passwordEntry implements the PasswordEntry interface.
type passwordEntry struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Editable
}

var _ PasswordEntry = (*passwordEntry)(nil)

// WrapPasswordEntry wraps a GObject to the right type. It is
// primarily used internally.
func WrapPasswordEntry(obj *externglib.Object) PasswordEntry {
	return PasswordEntry{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		Editable:         WrapEditable(obj),
	}
}

func marshalPasswordEntry(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPasswordEntry(obj), nil
}

// NewPasswordEntry constructs a class PasswordEntry.
func NewPasswordEntry() {
	C.gtk_password_entry_new()
}

// ExtraMenu gets the menu model set with
// gtk_password_entry_set_extra_menu().
func (e passwordEntry) ExtraMenu(e PasswordEntry) {
	var arg0 *C.GtkPasswordEntry

	arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(e.Native()))

	C.gtk_password_entry_get_extra_menu(arg0)
}

// ShowPeekIcon returns whether the entry is showing a clickable icon to
// reveal the contents of the entry in clear text.
func (e passwordEntry) ShowPeekIcon(e PasswordEntry) bool {
	var arg0 *C.GtkPasswordEntry

	arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(e.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_password_entry_get_show_peek_icon(arg0)

	if cret {
		ok = true
	}

	return ok
}

// SetExtraMenu sets a menu model to add when constructing the context menu
// for @entry.
func (e passwordEntry) SetExtraMenu(e PasswordEntry, model gio.MenuModel) {
	var arg0 *C.GtkPasswordEntry
	var arg1 *C.GMenuModel

	arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(e.Native()))
	arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	C.gtk_password_entry_set_extra_menu(arg0, arg1)
}

// SetShowPeekIcon sets whether the entry should have a clickable icon to
// show the contents of the entry in clear text.
//
// Setting this to false also hides the text again.
func (e passwordEntry) SetShowPeekIcon(e PasswordEntry, showPeekIcon bool) {
	var arg0 *C.GtkPasswordEntry
	var arg1 C.gboolean

	arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(e.Native()))
	if showPeekIcon {
		arg1 = C.gboolean(1)
	}

	C.gtk_password_entry_set_show_peek_icon(arg0, arg1)
}