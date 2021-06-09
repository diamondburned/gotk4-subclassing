// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_menu_shell_get_type()), F: marshalMenuShell},
	})
}

// MenuShell: a MenuShell is the abstract base class used to derive the Menu and
// MenuBar subclasses.
//
// A MenuShell is a container of MenuItem objects arranged in a list which can
// be navigated, selected, and activated by the user to perform application
// functions. A MenuItem can have a submenu associated with it, allowing for
// nested hierarchical menus.
//
//
// Terminology
//
// A menu item can be “selected”, this means that it is displayed in the
// prelight state, and if it has a submenu, that submenu will be popped up.
//
// A menu is “active” when it is visible onscreen and the user is selecting from
// it. A menubar is not active until the user clicks on one of its menuitems.
// When a menu is active, passing the mouse over a submenu will pop it up.
//
// There is also is a concept of the current menu and a current menu item. The
// current menu item is the selected menu item that is furthest down in the
// hierarchy. (Every active menu shell does not necessarily contain a selected
// menu item, but if it does, then the parent menu shell must also contain a
// selected menu item.) The current menu is the menu that contains the current
// menu item. It will always have a GTK grab and receive all key presses.
type MenuShell interface {
	Container
	Buildable

	// ActivateItem activates the menu item within the menu shell.
	ActivateItem(menuItem Widget, forceDeactivate bool)
	// Append adds a new MenuItem to the end of the menu shell's item list.
	Append(child MenuItem)
	// BindModel establishes a binding between a MenuShell and a Model.
	//
	// The contents of @shell are removed and then refilled with menu items
	// according to @model. When @model changes, @shell is updated. Calling this
	// function twice on @shell with different @model will cause the first
	// binding to be replaced with a binding to the new model. If @model is nil
	// then any previous binding is undone and all children are removed.
	//
	// @with_separators determines if toplevel items (eg: sections) have
	// separators inserted between them. This is typically desired for menus but
	// doesn’t make sense for menubars.
	//
	// If @action_namespace is non-nil then the effect is as if all actions
	// mentioned in the @model have their names prefixed with the namespace,
	// plus a dot. For example, if the action “quit” is mentioned and
	// @action_namespace is “app” then the effective action name is “app.quit”.
	//
	// This function uses Actionable to define the action name and target values
	// on the created menu items. If you want to use an action group other than
	// “app” and “win”, or if you want to use a MenuShell outside of a
	// ApplicationWindow, then you will need to attach your own action group to
	// the widget hierarchy using gtk_widget_insert_action_group(). As an
	// example, if you created a group with a “quit” action and inserted it with
	// the name “mygroup” then you would use the action name “mygroup.quit” in
	// your Model.
	//
	// For most cases you are probably better off using
	// gtk_menu_new_from_model() or gtk_menu_bar_new_from_model() or just
	// directly passing the Model to gtk_application_set_app_menu() or
	// gtk_application_set_menubar().
	BindModel(model gio.MenuModel, actionNamespace string, withSeparators bool)
	// Cancel cancels the selection within the menu shell.
	Cancel()
	// Deactivate deactivates the menu shell.
	//
	// Typically this results in the menu shell being erased from the screen.
	Deactivate()
	// Deselect deselects the currently selected item from the menu shell, if
	// any.
	Deselect()
	// ParentShell gets the parent menu shell.
	//
	// The parent menu shell of a submenu is the Menu or MenuBar from which it
	// was opened up.
	ParentShell() Widget
	// SelectedItem gets the currently selected item.
	SelectedItem() Widget
	// TakeFocus returns true if the menu shell will take the keyboard focus on
	// popup.
	TakeFocus() bool
	// Insert adds a new MenuItem to the menu shell’s item list at the position
	// indicated by @position.
	Insert(child Widget, position int)
	// Prepend adds a new MenuItem to the beginning of the menu shell's item
	// list.
	Prepend(child Widget)
	// SelectFirst: select the first visible or selectable child of the menu
	// shell; don’t select tearoff items unless the only item is a tearoff item.
	SelectFirst(searchSensitive bool)
	// SelectItem selects the menu item from the menu shell.
	SelectItem(menuItem Widget)
	// SetTakeFocus: if @take_focus is true (the default) the menu shell will
	// take the keyboard focus so that it will receive all keyboard events which
	// is needed to enable keyboard navigation in menus.
	//
	// Setting @take_focus to false is useful only for special applications like
	// virtual keyboard implementations which should not take keyboard focus.
	//
	// The @take_focus state of a menu or menu bar is automatically propagated
	// to submenus whenever a submenu is popped up, so you don’t have to worry
	// about recursively setting it for your entire menu hierarchy. Only when
	// programmatically picking a submenu and popping it up manually, the
	// @take_focus property of the submenu needs to be set explicitly.
	//
	// Note that setting it to false has side-effects:
	//
	// If the focus is in some other app, it keeps the focus and keynav in the
	// menu doesn’t work. Consequently, keynav on the menu will only work if the
	// focus is on some toplevel owned by the onscreen keyboard.
	//
	// To avoid confusing the user, menus with @take_focus set to false should
	// not display mnemonics or accelerators, since it cannot be guaranteed that
	// they will work.
	//
	// See also gdk_keyboard_grab()
	SetTakeFocus(takeFocus bool)
}

// menuShell implements the MenuShell interface.
type menuShell struct {
	Container
	Buildable
}

var _ MenuShell = (*menuShell)(nil)

// WrapMenuShell wraps a GObject to the right type. It is
// primarily used internally.
func WrapMenuShell(obj *externglib.Object) MenuShell {
	return MenuShell{
		Container: WrapContainer(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalMenuShell(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMenuShell(obj), nil
}

// ActivateItem activates the menu item within the menu shell.
func (m menuShell) ActivateItem(menuItem Widget, forceDeactivate bool) {
	var arg0 *C.GtkMenuShell
	var arg1 *C.GtkWidget
	var arg2 C.gboolean

	arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(menuItem.Native()))
	if forceDeactivate {
		arg2 = C.gboolean(1)
	}

	C.gtk_menu_shell_activate_item(arg0, arg1, arg2)
}

// Append adds a new MenuItem to the end of the menu shell's item list.
func (m menuShell) Append(child MenuItem) {
	var arg0 *C.GtkMenuShell
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_menu_shell_append(arg0, arg1)
}

// BindModel establishes a binding between a MenuShell and a Model.
//
// The contents of @shell are removed and then refilled with menu items
// according to @model. When @model changes, @shell is updated. Calling this
// function twice on @shell with different @model will cause the first
// binding to be replaced with a binding to the new model. If @model is nil
// then any previous binding is undone and all children are removed.
//
// @with_separators determines if toplevel items (eg: sections) have
// separators inserted between them. This is typically desired for menus but
// doesn’t make sense for menubars.
//
// If @action_namespace is non-nil then the effect is as if all actions
// mentioned in the @model have their names prefixed with the namespace,
// plus a dot. For example, if the action “quit” is mentioned and
// @action_namespace is “app” then the effective action name is “app.quit”.
//
// This function uses Actionable to define the action name and target values
// on the created menu items. If you want to use an action group other than
// “app” and “win”, or if you want to use a MenuShell outside of a
// ApplicationWindow, then you will need to attach your own action group to
// the widget hierarchy using gtk_widget_insert_action_group(). As an
// example, if you created a group with a “quit” action and inserted it with
// the name “mygroup” then you would use the action name “mygroup.quit” in
// your Model.
//
// For most cases you are probably better off using
// gtk_menu_new_from_model() or gtk_menu_bar_new_from_model() or just
// directly passing the Model to gtk_application_set_app_menu() or
// gtk_application_set_menubar().
func (m menuShell) BindModel(model gio.MenuModel, actionNamespace string, withSeparators bool) {
	var arg0 *C.GtkMenuShell
	var arg1 *C.GMenuModel
	var arg2 *C.gchar
	var arg3 C.gboolean

	arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GMenuModel)(unsafe.Pointer(model.Native()))
	arg2 = (*C.gchar)(C.CString(actionNamespace))
	defer C.free(unsafe.Pointer(arg2))
	if withSeparators {
		arg3 = C.gboolean(1)
	}

	C.gtk_menu_shell_bind_model(arg0, arg1, arg2, arg3)
}

// Cancel cancels the selection within the menu shell.
func (m menuShell) Cancel() {
	var arg0 *C.GtkMenuShell

	arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))

	C.gtk_menu_shell_cancel(arg0)
}

// Deactivate deactivates the menu shell.
//
// Typically this results in the menu shell being erased from the screen.
func (m menuShell) Deactivate() {
	var arg0 *C.GtkMenuShell

	arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))

	C.gtk_menu_shell_deactivate(arg0)
}

// Deselect deselects the currently selected item from the menu shell, if
// any.
func (m menuShell) Deselect() {
	var arg0 *C.GtkMenuShell

	arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))

	C.gtk_menu_shell_deselect(arg0)
}

// ParentShell gets the parent menu shell.
//
// The parent menu shell of a submenu is the Menu or MenuBar from which it
// was opened up.
func (m menuShell) ParentShell() Widget {
	var arg0 *C.GtkMenuShell

	arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))

	var cret *C.GtkWidget

	cret = C.gtk_menu_shell_get_parent_shell(arg0)

	var widget Widget

	widget = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return widget
}

// SelectedItem gets the currently selected item.
func (m menuShell) SelectedItem() Widget {
	var arg0 *C.GtkMenuShell

	arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))

	var cret *C.GtkWidget

	cret = C.gtk_menu_shell_get_selected_item(arg0)

	var widget Widget

	widget = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return widget
}

// TakeFocus returns true if the menu shell will take the keyboard focus on
// popup.
func (m menuShell) TakeFocus() bool {
	var arg0 *C.GtkMenuShell

	arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))

	var cret C.gboolean

	cret = C.gtk_menu_shell_get_take_focus(arg0)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// Insert adds a new MenuItem to the menu shell’s item list at the position
// indicated by @position.
func (m menuShell) Insert(child Widget, position int) {
	var arg0 *C.GtkMenuShell
	var arg1 *C.GtkWidget
	var arg2 C.gint

	arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = C.gint(position)

	C.gtk_menu_shell_insert(arg0, arg1, arg2)
}

// Prepend adds a new MenuItem to the beginning of the menu shell's item
// list.
func (m menuShell) Prepend(child Widget) {
	var arg0 *C.GtkMenuShell
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_menu_shell_prepend(arg0, arg1)
}

// SelectFirst: select the first visible or selectable child of the menu
// shell; don’t select tearoff items unless the only item is a tearoff item.
func (m menuShell) SelectFirst(searchSensitive bool) {
	var arg0 *C.GtkMenuShell
	var arg1 C.gboolean

	arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))
	if searchSensitive {
		arg1 = C.gboolean(1)
	}

	C.gtk_menu_shell_select_first(arg0, arg1)
}

// SelectItem selects the menu item from the menu shell.
func (m menuShell) SelectItem(menuItem Widget) {
	var arg0 *C.GtkMenuShell
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(menuItem.Native()))

	C.gtk_menu_shell_select_item(arg0, arg1)
}

// SetTakeFocus: if @take_focus is true (the default) the menu shell will
// take the keyboard focus so that it will receive all keyboard events which
// is needed to enable keyboard navigation in menus.
//
// Setting @take_focus to false is useful only for special applications like
// virtual keyboard implementations which should not take keyboard focus.
//
// The @take_focus state of a menu or menu bar is automatically propagated
// to submenus whenever a submenu is popped up, so you don’t have to worry
// about recursively setting it for your entire menu hierarchy. Only when
// programmatically picking a submenu and popping it up manually, the
// @take_focus property of the submenu needs to be set explicitly.
//
// Note that setting it to false has side-effects:
//
// If the focus is in some other app, it keeps the focus and keynav in the
// menu doesn’t work. Consequently, keynav on the menu will only work if the
// focus is on some toplevel owned by the onscreen keyboard.
//
// To avoid confusing the user, menus with @take_focus set to false should
// not display mnemonics or accelerators, since it cannot be guaranteed that
// they will work.
//
// See also gdk_keyboard_grab()
func (m menuShell) SetTakeFocus(takeFocus bool) {
	var arg0 *C.GtkMenuShell
	var arg1 C.gboolean

	arg0 = (*C.GtkMenuShell)(unsafe.Pointer(m.Native()))
	if takeFocus {
		arg1 = C.gboolean(1)
	}

	C.gtk_menu_shell_set_take_focus(arg0, arg1)
}
