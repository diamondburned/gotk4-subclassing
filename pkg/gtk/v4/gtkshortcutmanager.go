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
		{T: externglib.Type(C.gtk_shortcut_manager_get_type()), F: marshalShortcutManager},
	})
}

// ShortcutManagerOverrider contains methods that are overridable. This
// interface is a subset of the interface ShortcutManager.
type ShortcutManagerOverrider interface {
	AddController(s ShortcutManager, controller ShortcutController)

	RemoveController(s ShortcutManager, controller ShortcutController)
}

// ShortcutManager: the GtkShortcutManager interface is used to implement
// shortcut scopes.
//
// This is important for Native widgets that have their own surface, since the
// event controllers that are used to implement managed and global scopes are
// limited to the same native.
//
// Examples for widgets implementing ShortcutManager are Window and Popover.
type ShortcutManager interface {
	gextras.Objector
	ShortcutManagerOverrider
}

// shortcutManager implements the ShortcutManager interface.
type shortcutManager struct {
	gextras.Objector
}

var _ ShortcutManager = (*shortcutManager)(nil)

// WrapShortcutManager wraps a GObject to a type that implements interface
// ShortcutManager. It is primarily used internally.
func WrapShortcutManager(obj *externglib.Object) ShortcutManager {
	return ShortcutManager{
		Objector: obj,
	}
}

func marshalShortcutManager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapShortcutManager(obj), nil
}
