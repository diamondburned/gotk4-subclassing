// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_password_entry_get_type()), F: marshalPasswordEntrier},
	})
}

// PasswordEntry: GtkPasswordEntry is an entry that has been tailored for
// entering secrets.
//
// !An example GtkPasswordEntry (password-entry.png)
//
// It does not show its contents in clear text, does not allow to copy it to the
// clipboard, and it shows a warning when Caps Lock is engaged. If the
// underlying platform allows it, GtkPasswordEntry will also place the text in a
// non-pageable memory area, to avoid it being written out to disk by the
// operating system.
//
// Optionally, it can offer a way to reveal the contents in clear text.
//
// GtkPasswordEntry provides only minimal API and should be used with the
// gtk.Editable API.
//
// CSS Nodes
//
//    entry.password
//    ╰── text
//        ├── image.caps-lock-indicator
//        ┊
//
//
// GtkPasswordEntry has a single CSS node with name entry that carries a
// .passwordstyle class. The text Css node below it has a child with name image
// and style class .caps-lock-indicator for the Caps Lock icon, and possibly
// other children.
//
//
// Accessibility
//
// GtkPasswordEntry uses the GTK_ACCESSIBLE_ROLE_TEXT_BOX role.
type PasswordEntry struct {
	_ [0]func() // equal guard
	Widget

	*externglib.Object
	Editable
}

var (
	_ Widgetter           = (*PasswordEntry)(nil)
	_ externglib.Objector = (*PasswordEntry)(nil)
)

func wrapPasswordEntry(obj *externglib.Object) *PasswordEntry {
	return &PasswordEntry{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
		Object: obj,
		Editable: Editable{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
		},
	}
}

func marshalPasswordEntrier(p uintptr) (interface{}, error) {
	return wrapPasswordEntry(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectActivate: emitted when the entry is activated.
//
// The keybindings for this signal are all forms of the Enter key.
func (entry *PasswordEntry) ConnectActivate(f func()) externglib.SignalHandle {
	return entry.Connect("activate", externglib.GeneratedClosure{Func: f})
}

// NewPasswordEntry creates a GtkPasswordEntry.
//
// The function returns the following values:
//
//    - passwordEntry: new GtkPasswordEntry.
//
func NewPasswordEntry() *PasswordEntry {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_password_entry_new()

	var _passwordEntry *PasswordEntry // out

	_passwordEntry = wrapPasswordEntry(externglib.Take(unsafe.Pointer(_cret)))

	return _passwordEntry
}

// ExtraMenu gets the menu model set with gtk_password_entry_set_extra_menu().
//
// The function returns the following values:
//
//    - menuModel: (nullable): the menu model.
//
func (entry *PasswordEntry) ExtraMenu() gio.MenuModeller {
	var _arg0 *C.GtkPasswordEntry // out
	var _cret *C.GMenuModel       // in

	_arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_password_entry_get_extra_menu(_arg0)
	runtime.KeepAlive(entry)

	var _menuModel gio.MenuModeller // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.MenuModeller is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(gio.MenuModeller)
			return ok
		})
		rv, ok := casted.(gio.MenuModeller)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.MenuModeller")
		}
		_menuModel = rv
	}

	return _menuModel
}

// ShowPeekIcon returns whether the entry is showing an icon to reveal the
// contents.
//
// The function returns the following values:
//
//    - ok: TRUE if an icon is shown.
//
func (entry *PasswordEntry) ShowPeekIcon() bool {
	var _arg0 *C.GtkPasswordEntry // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(entry.Native()))

	_cret = C.gtk_password_entry_get_show_peek_icon(_arg0)
	runtime.KeepAlive(entry)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetExtraMenu sets a menu model to add when constructing the context menu for
// entry.
//
// The function takes the following parameters:
//
//    - model (optional): GMenuModel.
//
func (entry *PasswordEntry) SetExtraMenu(model gio.MenuModeller) {
	var _arg0 *C.GtkPasswordEntry // out
	var _arg1 *C.GMenuModel       // out

	_arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(entry.Native()))
	if model != nil {
		_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))
	}

	C.gtk_password_entry_set_extra_menu(_arg0, _arg1)
	runtime.KeepAlive(entry)
	runtime.KeepAlive(model)
}

// SetShowPeekIcon sets whether the entry should have a clickable icon to reveal
// the contents.
//
// Setting this to FALSE also hides the text again.
//
// The function takes the following parameters:
//
//    - showPeekIcon: whether to show the peek icon.
//
func (entry *PasswordEntry) SetShowPeekIcon(showPeekIcon bool) {
	var _arg0 *C.GtkPasswordEntry // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPasswordEntry)(unsafe.Pointer(entry.Native()))
	if showPeekIcon {
		_arg1 = C.TRUE
	}

	C.gtk_password_entry_set_show_peek_icon(_arg0, _arg1)
	runtime.KeepAlive(entry)
	runtime.KeepAlive(showPeekIcon)
}
