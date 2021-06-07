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
		{T: externglib.Type(C.gtk_drop_down_get_type()), F: marshalDropDown},
	})
}

// DropDown: gtkDropDown is a widget that allows the user to choose an item from
// a list of options. The GtkDropDown displays the selected choice.
//
// The options are given to GtkDropDown in the form of Model, and how the
// individual options are represented is determined by a ListItemFactory. The
// default factory displays simple strings.
//
// GtkDropDown knows how to obtain strings from the items in a StringList; for
// other models, you have to provide an expression to find the strings via
// gtk_drop_down_set_expression().
//
// GtkDropDown can optionally allow search in the popup, which is useful if the
// list of options is long. To enable the search entry, use
// gtk_drop_down_set_enable_search().
//
//
// CSS nodes
//
// GtkDropDown has a single CSS node with name dropdown, with the button and
// popover nodes as children.
//
//
// Accessibility
//
// GtkDropDown uses the K_ACCESSIBLE_ROLE_COMBO_BOX role.
type DropDown interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget

	// EnableSearch returns whether search is enabled.
	EnableSearch(s DropDown) bool
	// Expression gets the expression set with gtk_drop_down_set_expression().
	Expression(s DropDown)
	// Factory gets the factory that's currently used to populate list items.
	//
	// The factory returned by this function is always used for the item in the
	// button. It is also used for items in the popup if DropDown:list-factory
	// is not set.
	Factory(s DropDown)
	// ListFactory gets the factory that's currently used to populate list items
	// in the popup.
	ListFactory(s DropDown)
	// Model gets the model that provides the displayed items.
	Model(s DropDown)
	// Selected gets the position of the selected item.
	Selected(s DropDown)
	// SelectedItem gets the selected item. If no item is selected, nil is
	// returned.
	SelectedItem(s DropDown)
	// SetEnableSearch sets whether a search entry will be shown in the popup
	// that allows to search for items in the list.
	//
	// Note that DropDown:expression must be set for search to work.
	SetEnableSearch(s DropDown, enableSearch bool)
	// SetExpression sets the expression that gets evaluated to obtain strings
	// from items when searching in the popup. The expression must have a value
	// type of TYPE_STRING.
	SetExpression(s DropDown, expression Expression)
	// SetFactory sets the ListItemFactory to use for populating list items.
	SetFactory(s DropDown, factory ListItemFactory)
	// SetListFactory sets the ListItemFactory to use for populating list items
	// in the popup.
	SetListFactory(s DropDown, factory ListItemFactory)
	// SetModel sets the Model to use.
	SetModel(s DropDown, model gio.ListModel)
	// SetSelected selects the item at the given position.
	SetSelected(s DropDown, position uint)
}

// dropDown implements the DropDown interface.
type dropDown struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
}

var _ DropDown = (*dropDown)(nil)

// WrapDropDown wraps a GObject to the right type. It is
// primarily used internally.
func WrapDropDown(obj *externglib.Object) DropDown {
	return DropDown{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalDropDown(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDropDown(obj), nil
}

// NewDropDown constructs a class DropDown.
func NewDropDown(model gio.ListModel, expression Expression) {
	var arg1 *C.GListModel
	var arg2 *C.GtkExpression

	arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
	arg2 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))

	C.gtk_drop_down_new(arg1, arg2)
}

// NewDropDownFromStrings constructs a class DropDown.
func NewDropDownFromStrings(strings []string) {
	var arg1 **C.char

	arg1 = C.malloc(len(strings) * (unsafe.Sizeof(int(0)) + 1))
	defer C.free(unsafe.Pointer(arg1))

	{
		var out []*C.char
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(arg1), int(len(strings)))

		for i := range strings {
			out[i] = (*C.char)(C.CString(strings[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.gtk_drop_down_new_from_strings(arg1)
}

// EnableSearch returns whether search is enabled.
func (s dropDown) EnableSearch(s DropDown) bool {
	var arg0 *C.GtkDropDown

	arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_drop_down_get_enable_search(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Expression gets the expression set with gtk_drop_down_set_expression().
func (s dropDown) Expression(s DropDown) {
	var arg0 *C.GtkDropDown

	arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))

	C.gtk_drop_down_get_expression(arg0)
}

// Factory gets the factory that's currently used to populate list items.
//
// The factory returned by this function is always used for the item in the
// button. It is also used for items in the popup if DropDown:list-factory
// is not set.
func (s dropDown) Factory(s DropDown) {
	var arg0 *C.GtkDropDown

	arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))

	C.gtk_drop_down_get_factory(arg0)
}

// ListFactory gets the factory that's currently used to populate list items
// in the popup.
func (s dropDown) ListFactory(s DropDown) {
	var arg0 *C.GtkDropDown

	arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))

	C.gtk_drop_down_get_list_factory(arg0)
}

// Model gets the model that provides the displayed items.
func (s dropDown) Model(s DropDown) {
	var arg0 *C.GtkDropDown

	arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))

	C.gtk_drop_down_get_model(arg0)
}

// Selected gets the position of the selected item.
func (s dropDown) Selected(s DropDown) {
	var arg0 *C.GtkDropDown

	arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))

	C.gtk_drop_down_get_selected(arg0)
}

// SelectedItem gets the selected item. If no item is selected, nil is
// returned.
func (s dropDown) SelectedItem(s DropDown) {
	var arg0 *C.GtkDropDown

	arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))

	C.gtk_drop_down_get_selected_item(arg0)
}

// SetEnableSearch sets whether a search entry will be shown in the popup
// that allows to search for items in the list.
//
// Note that DropDown:expression must be set for search to work.
func (s dropDown) SetEnableSearch(s DropDown, enableSearch bool) {
	var arg0 *C.GtkDropDown
	var arg1 C.gboolean

	arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))
	if enableSearch {
		arg1 = C.gboolean(1)
	}

	C.gtk_drop_down_set_enable_search(arg0, arg1)
}

// SetExpression sets the expression that gets evaluated to obtain strings
// from items when searching in the popup. The expression must have a value
// type of TYPE_STRING.
func (s dropDown) SetExpression(s DropDown, expression Expression) {
	var arg0 *C.GtkDropDown
	var arg1 *C.GtkExpression

	arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))

	C.gtk_drop_down_set_expression(arg0, arg1)
}

// SetFactory sets the ListItemFactory to use for populating list items.
func (s dropDown) SetFactory(s DropDown, factory ListItemFactory) {
	var arg0 *C.GtkDropDown
	var arg1 *C.GtkListItemFactory

	arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))

	C.gtk_drop_down_set_factory(arg0, arg1)
}

// SetListFactory sets the ListItemFactory to use for populating list items
// in the popup.
func (s dropDown) SetListFactory(s DropDown, factory ListItemFactory) {
	var arg0 *C.GtkDropDown
	var arg1 *C.GtkListItemFactory

	arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))

	C.gtk_drop_down_set_list_factory(arg0, arg1)
}

// SetModel sets the Model to use.
func (s dropDown) SetModel(s DropDown, model gio.ListModel) {
	var arg0 *C.GtkDropDown
	var arg1 *C.GListModel

	arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	C.gtk_drop_down_set_model(arg0, arg1)
}

// SetSelected selects the item at the given position.
func (s dropDown) SetSelected(s DropDown, position uint) {
	var arg0 *C.GtkDropDown
	var arg1 C.guint

	arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))
	arg1 = C.guint(position)

	C.gtk_drop_down_set_selected(arg0, arg1)
}
