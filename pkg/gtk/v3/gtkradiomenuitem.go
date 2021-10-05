// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_radio_menu_item_get_type()), F: marshalRadioMenuItemmer},
	})
}

// RadioMenuItemOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type RadioMenuItemOverrider interface {
	GroupChanged()
}

// RadioMenuItem: radio menu item is a check menu item that belongs to a group.
// At each instant exactly one of the radio menu items from a group is selected.
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
type RadioMenuItem struct {
	CheckMenuItem
}

func wrapRadioMenuItem(obj *externglib.Object) *RadioMenuItem {
	return &RadioMenuItem{
		CheckMenuItem: CheckMenuItem{
			MenuItem: MenuItem{
				Bin: Bin{
					Container: Container{
						Widget: Widget{
							InitiallyUnowned: externglib.InitiallyUnowned{
								Object: obj,
							},
							ImplementorIface: atk.ImplementorIface{
								Object: obj,
							},
							Buildable: Buildable{
								Object: obj,
							},
							Object: obj,
						},
					},
				},
				Actionable: Actionable{
					Widget: Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
						Object: obj,
					},
				},
				Activatable: Activatable{
					Object: obj,
				},
				Object: obj,
			},
		},
	}
}

func marshalRadioMenuItemmer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRadioMenuItem(obj), nil
}

// NewRadioMenuItem creates a new RadioMenuItem.
func NewRadioMenuItem(group []RadioMenuItem) *RadioMenuItem {
	var _arg1 *C.GSList    // out
	var _cret *C.GtkWidget // in

	if group != nil {
		for i := len(group) - 1; i >= 0; i-- {
			src := group[i]
			var dst *C.GtkRadioMenuItem // out
			dst = (*C.GtkRadioMenuItem)(unsafe.Pointer((&src).Native()))
			_arg1 = C.g_slist_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
		}
		defer C.g_slist_free(_arg1)
	}

	_cret = C.gtk_radio_menu_item_new(_arg1)
	runtime.KeepAlive(group)

	var _radioMenuItem *RadioMenuItem // out

	_radioMenuItem = wrapRadioMenuItem(externglib.Take(unsafe.Pointer(_cret)))

	return _radioMenuItem
}

// NewRadioMenuItemFromWidget creates a new RadioMenuItem adding it to the same
// group as group.
func NewRadioMenuItemFromWidget(group *RadioMenuItem) *RadioMenuItem {
	var _arg1 *C.GtkRadioMenuItem // out
	var _cret *C.GtkWidget        // in

	if group != nil {
		_arg1 = (*C.GtkRadioMenuItem)(unsafe.Pointer(group.Native()))
	}

	_cret = C.gtk_radio_menu_item_new_from_widget(_arg1)
	runtime.KeepAlive(group)

	var _radioMenuItem *RadioMenuItem // out

	_radioMenuItem = wrapRadioMenuItem(externglib.Take(unsafe.Pointer(_cret)))

	return _radioMenuItem
}

// NewRadioMenuItemWithLabel creates a new RadioMenuItem whose child is a simple
// Label.
func NewRadioMenuItemWithLabel(group []RadioMenuItem, label string) *RadioMenuItem {
	var _arg1 *C.GSList    // out
	var _arg2 *C.gchar     // out
	var _cret *C.GtkWidget // in

	if group != nil {
		for i := len(group) - 1; i >= 0; i-- {
			src := group[i]
			var dst *C.GtkRadioMenuItem // out
			dst = (*C.GtkRadioMenuItem)(unsafe.Pointer((&src).Native()))
			_arg1 = C.g_slist_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
		}
		defer C.g_slist_free(_arg1)
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_radio_menu_item_new_with_label(_arg1, _arg2)
	runtime.KeepAlive(group)
	runtime.KeepAlive(label)

	var _radioMenuItem *RadioMenuItem // out

	_radioMenuItem = wrapRadioMenuItem(externglib.Take(unsafe.Pointer(_cret)))

	return _radioMenuItem
}

// NewRadioMenuItemWithLabelFromWidget creates a new GtkRadioMenuItem whose
// child is a simple GtkLabel. The new RadioMenuItem is added to the same group
// as group.
func NewRadioMenuItemWithLabelFromWidget(group *RadioMenuItem, label string) *RadioMenuItem {
	var _arg1 *C.GtkRadioMenuItem // out
	var _arg2 *C.gchar            // out
	var _cret *C.GtkWidget        // in

	if group != nil {
		_arg1 = (*C.GtkRadioMenuItem)(unsafe.Pointer(group.Native()))
	}
	if label != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.gtk_radio_menu_item_new_with_label_from_widget(_arg1, _arg2)
	runtime.KeepAlive(group)
	runtime.KeepAlive(label)

	var _radioMenuItem *RadioMenuItem // out

	_radioMenuItem = wrapRadioMenuItem(externglib.Take(unsafe.Pointer(_cret)))

	return _radioMenuItem
}

// NewRadioMenuItemWithMnemonic creates a new RadioMenuItem containing a label.
// The label will be created using gtk_label_new_with_mnemonic(), so underscores
// in label indicate the mnemonic for the menu item.
func NewRadioMenuItemWithMnemonic(group []RadioMenuItem, label string) *RadioMenuItem {
	var _arg1 *C.GSList    // out
	var _arg2 *C.gchar     // out
	var _cret *C.GtkWidget // in

	if group != nil {
		for i := len(group) - 1; i >= 0; i-- {
			src := group[i]
			var dst *C.GtkRadioMenuItem // out
			dst = (*C.GtkRadioMenuItem)(unsafe.Pointer((&src).Native()))
			_arg1 = C.g_slist_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
		}
		defer C.g_slist_free(_arg1)
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_radio_menu_item_new_with_mnemonic(_arg1, _arg2)
	runtime.KeepAlive(group)
	runtime.KeepAlive(label)

	var _radioMenuItem *RadioMenuItem // out

	_radioMenuItem = wrapRadioMenuItem(externglib.Take(unsafe.Pointer(_cret)))

	return _radioMenuItem
}

// NewRadioMenuItemWithMnemonicFromWidget creates a new GtkRadioMenuItem
// containing a label. The label will be created using
// gtk_label_new_with_mnemonic(), so underscores in label indicate the mnemonic
// for the menu item.
//
// The new RadioMenuItem is added to the same group as group.
func NewRadioMenuItemWithMnemonicFromWidget(group *RadioMenuItem, label string) *RadioMenuItem {
	var _arg1 *C.GtkRadioMenuItem // out
	var _arg2 *C.gchar            // out
	var _cret *C.GtkWidget        // in

	if group != nil {
		_arg1 = (*C.GtkRadioMenuItem)(unsafe.Pointer(group.Native()))
	}
	if label != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.gtk_radio_menu_item_new_with_mnemonic_from_widget(_arg1, _arg2)
	runtime.KeepAlive(group)
	runtime.KeepAlive(label)

	var _radioMenuItem *RadioMenuItem // out

	_radioMenuItem = wrapRadioMenuItem(externglib.Take(unsafe.Pointer(_cret)))

	return _radioMenuItem
}

// Group returns the group to which the radio menu item belongs, as a #GList of
// RadioMenuItem. The list belongs to GTK+ and should not be freed.
func (radioMenuItem *RadioMenuItem) Group() []RadioMenuItem {
	var _arg0 *C.GtkRadioMenuItem // out
	var _cret *C.GSList           // in

	_arg0 = (*C.GtkRadioMenuItem)(unsafe.Pointer(radioMenuItem.Native()))

	_cret = C.gtk_radio_menu_item_get_group(_arg0)
	runtime.KeepAlive(radioMenuItem)

	var _sList []RadioMenuItem // out

	_sList = make([]RadioMenuItem, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.GtkRadioMenuItem)(v)
		var dst RadioMenuItem // out
		dst = *wrapRadioMenuItem(externglib.Take(unsafe.Pointer(src)))
		_sList = append(_sList, dst)
	})

	return _sList
}

// JoinGroup joins a RadioMenuItem object to the group of another RadioMenuItem
// object.
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
//        }.
func (radioMenuItem *RadioMenuItem) JoinGroup(groupSource *RadioMenuItem) {
	var _arg0 *C.GtkRadioMenuItem // out
	var _arg1 *C.GtkRadioMenuItem // out

	_arg0 = (*C.GtkRadioMenuItem)(unsafe.Pointer(radioMenuItem.Native()))
	if groupSource != nil {
		_arg1 = (*C.GtkRadioMenuItem)(unsafe.Pointer(groupSource.Native()))
	}

	C.gtk_radio_menu_item_join_group(_arg0, _arg1)
	runtime.KeepAlive(radioMenuItem)
	runtime.KeepAlive(groupSource)
}

// SetGroup sets the group of a radio menu item, or changes it.
func (radioMenuItem *RadioMenuItem) SetGroup(group []RadioMenuItem) {
	var _arg0 *C.GtkRadioMenuItem // out
	var _arg1 *C.GSList           // out

	_arg0 = (*C.GtkRadioMenuItem)(unsafe.Pointer(radioMenuItem.Native()))
	if group != nil {
		for i := len(group) - 1; i >= 0; i-- {
			src := group[i]
			var dst *C.GtkRadioMenuItem // out
			dst = (*C.GtkRadioMenuItem)(unsafe.Pointer((&src).Native()))
			_arg1 = C.g_slist_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
		}
		defer C.g_slist_free(_arg1)
	}

	C.gtk_radio_menu_item_set_group(_arg0, _arg1)
	runtime.KeepAlive(radioMenuItem)
	runtime.KeepAlive(group)
}

func (r *RadioMenuItem) ConnectGroupChanged(f func()) glib.SignalHandle {
	return r.Connect("group-changed", f)
}
