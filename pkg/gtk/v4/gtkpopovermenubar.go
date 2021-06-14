// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_popover_menu_bar_get_type()), F: marshalPopoverMenuBar},
	})
}

// PopoverMenuBar: `GtkPopoverMenuBar` presents a horizontal bar of items that
// pop up popover menus when clicked.
//
// !An example GtkPopoverMenuBar (menubar.png)
//
// The only way to create instances of `GtkPopoverMenuBar` is from a
// `GMenuModel`.
//
//
// CSS nodes
//
// “` menubar ├── item[.active] ┊ ╰── popover ╰── item ╰── popover “`
//
// `GtkPopoverMenuBar` has a single CSS node with name menubar, below which each
// item has its CSS node, and below that the corresponding popover.
//
// The item whose popover is currently open gets the .active style class.
//
//
// Accessibility
//
// `GtkPopoverMenuBar` uses the GTK_ACCESSIBLE_ROLE_MENU_BAR role, the menu
// items use the GTK_ACCESSIBLE_ROLE_MENU_ITEM role and the menus use the
// GTK_ACCESSIBLE_ROLE_MENU role.
type PopoverMenuBar interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget

	// AddChild adds a custom widget to a generated menubar.
	//
	// For this to work, the menu model of @bar must have an item with a
	// `custom` attribute that matches @id.
	AddChild(child Widget, id string) bool
	// MenuModel returns the model from which the contents of @bar are taken.
	MenuModel() gio.MenuModel
	// RemoveChild removes a widget that has previously been added with
	// gtk_popover_menu_bar_add_child().
	RemoveChild(child Widget) bool
	// SetMenuModel sets a menu model from which @bar should take its contents.
	SetMenuModel(model gio.MenuModel)
}

// popoverMenuBar implements the PopoverMenuBar class.
type popoverMenuBar struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
}

var _ PopoverMenuBar = (*popoverMenuBar)(nil)

// WrapPopoverMenuBar wraps a GObject to the right type. It is
// primarily used internally.
func WrapPopoverMenuBar(obj *externglib.Object) PopoverMenuBar {
	return popoverMenuBar{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalPopoverMenuBar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPopoverMenuBar(obj), nil
}

// NewPopoverMenuBarFromModel constructs a class PopoverMenuBar.
func NewPopoverMenuBarFromModel(model gio.MenuModel) PopoverMenuBar {
	var _arg1 *C.GMenuModel // out

	_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	var _cret C.GtkPopoverMenuBar // in

	_cret = C.gtk_popover_menu_bar_new_from_model(_arg1)

	var _popoverMenuBar PopoverMenuBar // out

	_popoverMenuBar = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(PopoverMenuBar)

	return _popoverMenuBar
}

// AddChild adds a custom widget to a generated menubar.
//
// For this to work, the menu model of @bar must have an item with a
// `custom` attribute that matches @id.
func (b popoverMenuBar) AddChild(child Widget, id string) bool {
	var _arg0 *C.GtkPopoverMenuBar // out
	var _arg1 *C.GtkWidget         // out
	var _arg2 *C.char              // out

	_arg0 = (*C.GtkPopoverMenuBar)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.char)(C.CString(id))
	defer C.free(unsafe.Pointer(_arg2))

	var _cret C.gboolean // in

	_cret = C.gtk_popover_menu_bar_add_child(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MenuModel returns the model from which the contents of @bar are taken.
func (b popoverMenuBar) MenuModel() gio.MenuModel {
	var _arg0 *C.GtkPopoverMenuBar // out

	_arg0 = (*C.GtkPopoverMenuBar)(unsafe.Pointer(b.Native()))

	var _cret *C.GMenuModel // in

	_cret = C.gtk_popover_menu_bar_get_menu_model(_arg0)

	var _menuModel gio.MenuModel // out

	_menuModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gio.MenuModel)

	return _menuModel
}

// RemoveChild removes a widget that has previously been added with
// gtk_popover_menu_bar_add_child().
func (b popoverMenuBar) RemoveChild(child Widget) bool {
	var _arg0 *C.GtkPopoverMenuBar // out
	var _arg1 *C.GtkWidget         // out

	_arg0 = (*C.GtkPopoverMenuBar)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_popover_menu_bar_remove_child(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetMenuModel sets a menu model from which @bar should take its contents.
func (b popoverMenuBar) SetMenuModel(model gio.MenuModel) {
	var _arg0 *C.GtkPopoverMenuBar // out
	var _arg1 *C.GMenuModel        // out

	_arg0 = (*C.GtkPopoverMenuBar)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))

	C.gtk_popover_menu_bar_set_menu_model(_arg0, _arg1)
}
