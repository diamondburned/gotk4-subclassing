// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_list_item_get_type()), F: marshalListItem},
	})
}

// ListItem: `GtkListItem` is used by list widgets to represent items in a
// `GListModel`.
//
// The `GtkListItem`s are managed by the list widget (with its factory) and
// cannot be created by applications, but they need to be populated by
// application code. This is done by calling [method@Gtk.ListItem.set_child].
//
// `GtkListItem`s exist in 2 stages:
//
// 1. The unbound stage where the listitem is not currently connected to an item
// in the list. In that case, the [property@Gtk.ListItem:item] property is set
// to nil.
//
// 2. The bound stage where the listitem references an item from the list. The
// [property@Gtk.ListItem:item] property is not nil.
type ListItem interface {
	gextras.Objector

	// Activatable checks if a list item has been set to be activatable via
	// gtk_list_item_set_activatable().
	Activatable() bool
	// Item gets the model item that associated with @self.
	//
	// If @self is unbound, this function returns nil.
	Item() gextras.Objector
	// Position gets the position in the model that @self currently displays.
	//
	// If @self is unbound, GTK_INVALID_LIST_POSITION is returned.
	Position() uint
	// Selectable checks if a list item has been set to be selectable via
	// gtk_list_item_set_selectable().
	//
	// Do not confuse this function with [method@Gtk.ListItem.get_selected].
	Selectable() bool
	// Selected checks if the item is displayed as selected.
	//
	// The selected state is maintained by the liste widget and its model and
	// cannot be set otherwise.
	Selected() bool
	// SetActivatable sets @self to be activatable.
	//
	// If an item is activatable, double-clicking on the item, using the Return
	// key or calling gtk_widget_activate() will activate the item. Activating
	// instructs the containing view to handle activation. `GtkListView` for
	// example will be emitting the [signal@Gtk.ListView::activate] signal.
	//
	// By default, list items are activatable.
	SetActivatable(activatable bool)
	// SetChild sets the child to be used for this listitem.
	//
	// This function is typically called by applications when setting up a
	// listitem so that the widget can be reused when binding it multiple times.
	SetChild(child Widget)
	// SetSelectable sets @self to be selectable.
	//
	// If an item is selectable, clicking on the item or using the keyboard will
	// try to select or unselect the item. If this succeeds is up to the model
	// to determine, as it is managing the selected state.
	//
	// Note that this means that making an item non-selectable has no influence
	// on the selected state at all. A non-selectable item may still be
	// selected.
	//
	// By default, list items are selectable. When rebinding them to a new item,
	// they will also be reset to be selectable by GTK.
	SetSelectable(selectable bool)
}

// listItem implements the ListItem interface.
type listItem struct {
	gextras.Objector
}

var _ ListItem = (*listItem)(nil)

// WrapListItem wraps a GObject to the right type. It is
// primarily used internally.
func WrapListItem(obj *externglib.Object) ListItem {
	return ListItem{
		Objector: obj,
	}
}

func marshalListItem(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapListItem(obj), nil
}

// Activatable checks if a list item has been set to be activatable via
// gtk_list_item_set_activatable().
func (s listItem) Activatable() bool {
	var _arg0 *C.GtkListItem // out

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_list_item_get_activatable(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Item gets the model item that associated with @self.
//
// If @self is unbound, this function returns nil.
func (s listItem) Item() gextras.Objector {
	var _arg0 *C.GtkListItem // out

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(s.Native()))

	var _cret C.gpointer // in

	_cret = C.gtk_list_item_get_item(_arg0)

	var _object gextras.Objector // out

	_object = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gextras.Objector)

	return _object
}

// Position gets the position in the model that @self currently displays.
//
// If @self is unbound, GTK_INVALID_LIST_POSITION is returned.
func (s listItem) Position() uint {
	var _arg0 *C.GtkListItem // out

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(s.Native()))

	var _cret C.guint // in

	_cret = C.gtk_list_item_get_position(_arg0)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// Selectable checks if a list item has been set to be selectable via
// gtk_list_item_set_selectable().
//
// Do not confuse this function with [method@Gtk.ListItem.get_selected].
func (s listItem) Selectable() bool {
	var _arg0 *C.GtkListItem // out

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_list_item_get_selectable(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Selected checks if the item is displayed as selected.
//
// The selected state is maintained by the liste widget and its model and
// cannot be set otherwise.
func (s listItem) Selected() bool {
	var _arg0 *C.GtkListItem // out

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_list_item_get_selected(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// SetActivatable sets @self to be activatable.
//
// If an item is activatable, double-clicking on the item, using the Return
// key or calling gtk_widget_activate() will activate the item. Activating
// instructs the containing view to handle activation. `GtkListView` for
// example will be emitting the [signal@Gtk.ListView::activate] signal.
//
// By default, list items are activatable.
func (s listItem) SetActivatable(activatable bool) {
	var _arg0 *C.GtkListItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(s.Native()))
	if activatable {
		_arg1 = C.gboolean(1)
	}

	C.gtk_list_item_set_activatable(_arg0, _arg1)
}

// SetChild sets the child to be used for this listitem.
//
// This function is typically called by applications when setting up a
// listitem so that the widget can be reused when binding it multiple times.
func (s listItem) SetChild(child Widget) {
	var _arg0 *C.GtkListItem // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_list_item_set_child(_arg0, _arg1)
}

// SetSelectable sets @self to be selectable.
//
// If an item is selectable, clicking on the item or using the keyboard will
// try to select or unselect the item. If this succeeds is up to the model
// to determine, as it is managing the selected state.
//
// Note that this means that making an item non-selectable has no influence
// on the selected state at all. A non-selectable item may still be
// selected.
//
// By default, list items are selectable. When rebinding them to a new item,
// they will also be reset to be selectable by GTK.
func (s listItem) SetSelectable(selectable bool) {
	var _arg0 *C.GtkListItem // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkListItem)(unsafe.Pointer(s.Native()))
	if selectable {
		_arg1 = C.gboolean(1)
	}

	C.gtk_list_item_set_selectable(_arg0, _arg1)
}
