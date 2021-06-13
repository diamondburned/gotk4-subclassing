// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
		{T: externglib.Type(C.gtk_tearoff_menu_item_get_type()), F: marshalTearoffMenuItem},
	})
}

// TearoffMenuItem: a TearoffMenuItem is a special MenuItem which is used to
// tear off and reattach its menu.
//
// When its menu is shown normally, the TearoffMenuItem is drawn as a dotted
// line indicating that the menu can be torn off. Activating it causes its menu
// to be torn off and displayed in its own window as a tearoff menu.
//
// When its menu is shown as a tearoff menu, the TearoffMenuItem is drawn as a
// dotted line which has a left pointing arrow graphic indicating that the
// tearoff menu can be reattached. Activating it will erase the tearoff menu
// window.
//
// > TearoffMenuItem is deprecated and should not be used in newly > written
// code. Menus are not meant to be torn around.
type TearoffMenuItem interface {
	MenuItem
	Actionable
	Activatable
	Buildable
}

// tearoffMenuItem implements the TearoffMenuItem interface.
type tearoffMenuItem struct {
	MenuItem
	Actionable
	Activatable
	Buildable
}

var _ TearoffMenuItem = (*tearoffMenuItem)(nil)

// WrapTearoffMenuItem wraps a GObject to the right type. It is
// primarily used internally.
func WrapTearoffMenuItem(obj *externglib.Object) TearoffMenuItem {
	return TearoffMenuItem{
		MenuItem:    WrapMenuItem(obj),
		Actionable:  WrapActionable(obj),
		Activatable: WrapActivatable(obj),
		Buildable:   WrapBuildable(obj),
	}
}

func marshalTearoffMenuItem(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTearoffMenuItem(obj), nil
}
