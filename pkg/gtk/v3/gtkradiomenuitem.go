// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.gtk_radio_menu_item_get_type()), F: marshalRadioMenuItem},
	})
}

// RadioMenuItem: a radio menu item is a check menu item that belongs to a
// group. At each instant exactly one of the radio menu items from a group is
// selected.
//
// The group list does not need to be freed, as each RadioMenuItem will remove
// itself and its list item when it is destroyed.
//
// The correct way to create a group of radio menu items is approximatively
// this:
//
// How to create a group of radio menu items.
//
//    menuitem
//    ├── radio.left
//    ╰── <child>
//
// GtkRadioMenuItem has a main CSS node with name menuitem, and a subnode with
// name radio, which gets the .left or .right style class.
type RadioMenuItem interface {
	CheckMenuItem
	Actionable
	Activatable
	Buildable

	// JoinGroup joins a RadioMenuItem object to the group of another
	// RadioMenuItem object.
	//
	// This function should be used by language bindings to avoid the memory
	// manangement of the opaque List of gtk_radio_menu_item_get_group() and
	// gtk_radio_menu_item_set_group().
	//
	// A common way to set up a group of RadioMenuItem instances is:
	//
	//      GtkRadioMenuItem *last_item = NULL;
	//
	//      while ( ...more items to add... )
	//        {
	//          GtkRadioMenuItem *radio_item;
	//
	//          radio_item = gtk_radio_menu_item_new (...);
	//
	//          gtk_radio_menu_item_join_group (radio_item, last_item);
	//          last_item = radio_item;
	//        }
	JoinGroup(groupSource RadioMenuItem)
	// SetGroup sets the group of a radio menu item, or changes it.
	SetGroup(group *glib.SList)
}

// radioMenuItem implements the RadioMenuItem interface.
type radioMenuItem struct {
	CheckMenuItem
	Actionable
	Activatable
	Buildable
}

var _ RadioMenuItem = (*radioMenuItem)(nil)

// WrapRadioMenuItem wraps a GObject to the right type. It is
// primarily used internally.
func WrapRadioMenuItem(obj *externglib.Object) RadioMenuItem {
	return RadioMenuItem{
		CheckMenuItem: WrapCheckMenuItem(obj),
		Actionable:    WrapActionable(obj),
		Activatable:   WrapActivatable(obj),
		Buildable:     WrapBuildable(obj),
	}
}

func marshalRadioMenuItem(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRadioMenuItem(obj), nil
}

// JoinGroup joins a RadioMenuItem object to the group of another
// RadioMenuItem object.
//
// This function should be used by language bindings to avoid the memory
// manangement of the opaque List of gtk_radio_menu_item_get_group() and
// gtk_radio_menu_item_set_group().
//
// A common way to set up a group of RadioMenuItem instances is:
//
//      GtkRadioMenuItem *last_item = NULL;
//
//      while ( ...more items to add... )
//        {
//          GtkRadioMenuItem *radio_item;
//
//          radio_item = gtk_radio_menu_item_new (...);
//
//          gtk_radio_menu_item_join_group (radio_item, last_item);
//          last_item = radio_item;
//        }
func (r radioMenuItem) JoinGroup(groupSource RadioMenuItem) {
	var _arg0 *C.GtkRadioMenuItem // out
	var _arg1 *C.GtkRadioMenuItem // out

	_arg0 = (*C.GtkRadioMenuItem)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.GtkRadioMenuItem)(unsafe.Pointer(groupSource.Native()))

	C.gtk_radio_menu_item_join_group(_arg0, _arg1)
}

// SetGroup sets the group of a radio menu item, or changes it.
func (r radioMenuItem) SetGroup(group *glib.SList) {
	var _arg0 *C.GtkRadioMenuItem // out
	var _arg1 *C.GSList           // out

	_arg0 = (*C.GtkRadioMenuItem)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.GSList)(unsafe.Pointer(group.Native()))

	C.gtk_radio_menu_item_set_group(_arg0, _arg1)
}
