// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_drop_down_get_type()), F: marshalDropDown},
	})
}

// DropDown: `GtkDropDown` is a widget that allows the user to choose an item
// from a list of options.
//
// !An example GtkDropDown (drop-down.png)
//
// The `GtkDropDown` displays the selected choice.
//
// The options are given to `GtkDropDown` in the form of `GListModel` and how
// the individual options are represented is determined by a
// [class@Gtk.ListItemFactory]. The default factory displays simple strings.
//
// `GtkDropDown` knows how to obtain strings from the items in a
// [class@Gtk.StringList]; for other models, you have to provide an expression
// to find the strings via [method@Gtk.DropDown.set_expression].
//
// `GtkDropDown` can optionally allow search in the popup, which is useful if
// the list of options is long. To enable the search entry, use
// [method@Gtk.DropDown.set_enable_search].
//
//
// CSS nodes
//
// `GtkDropDown` has a single CSS node with name dropdown, with the button and
// popover nodes as children.
//
//
// Accessibility
//
// `GtkDropDown` uses the GTK_ACCESSIBLE_ROLE_COMBO_BOX role.
type DropDown interface {
	Widget

	// EnableSearch:
	EnableSearch() bool
	// Expression:
	Expression() Expression
	// Factory:
	Factory() ListItemFactory
	// ListFactory:
	ListFactory() ListItemFactory
	// Model:
	Model() gio.ListModel
	// Selected:
	Selected() uint
	// SelectedItem:
	SelectedItem() gextras.Objector
	// SetEnableSearchDropDown:
	SetEnableSearchDropDown(enableSearch bool)
	// SetExpressionDropDown:
	SetExpressionDropDown(expression Expression)
	// SetFactoryDropDown:
	SetFactoryDropDown(factory ListItemFactory)
	// SetListFactoryDropDown:
	SetListFactoryDropDown(factory ListItemFactory)
	// SetModelDropDown:
	SetModelDropDown(model gio.ListModel)
	// SetSelectedDropDown:
	SetSelectedDropDown(position uint)
}

// dropDown implements the DropDown class.
type dropDown struct {
	Widget
}

// WrapDropDown wraps a GObject to the right type. It is
// primarily used internally.
func WrapDropDown(obj *externglib.Object) DropDown {
	return dropDown{
		Widget: WrapWidget(obj),
	}
}

func marshalDropDown(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDropDown(obj), nil
}

// NewDropDown:
func NewDropDown(model gio.ListModel, expression Expression) DropDown {
	var _arg1 *C.GListModel    // out
	var _arg2 *C.GtkExpression // out
	var _cret *C.GtkWidget     // in

	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
	_arg2 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))

	_cret = C.gtk_drop_down_new(_arg1, _arg2)

	var _dropDown DropDown // out

	_dropDown = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(DropDown)

	return _dropDown
}

// NewDropDownFromStrings:
func NewDropDownFromStrings(strings []string) DropDown {
	var _arg1 **C.char
	var _cret *C.GtkWidget // in

	_arg1 = (**C.char)(C.malloc(C.ulong(len(strings)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice(_arg1, len(strings))
		for i := range strings {
			out[i] = (*C.char)(C.CString(strings[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	_cret = C.gtk_drop_down_new_from_strings(_arg1)

	var _dropDown DropDown // out

	_dropDown = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(DropDown)

	return _dropDown
}

func (s dropDown) EnableSearch() bool {
	var _arg0 *C.GtkDropDown // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_drop_down_get_enable_search(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s dropDown) Expression() Expression {
	var _arg0 *C.GtkDropDown   // out
	var _cret *C.GtkExpression // in

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_drop_down_get_expression(_arg0)

	var _expression Expression // out

	_expression = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Expression)

	return _expression
}

func (s dropDown) Factory() ListItemFactory {
	var _arg0 *C.GtkDropDown        // out
	var _cret *C.GtkListItemFactory // in

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_drop_down_get_factory(_arg0)

	var _listItemFactory ListItemFactory // out

	_listItemFactory = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ListItemFactory)

	return _listItemFactory
}

func (s dropDown) ListFactory() ListItemFactory {
	var _arg0 *C.GtkDropDown        // out
	var _cret *C.GtkListItemFactory // in

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_drop_down_get_list_factory(_arg0)

	var _listItemFactory ListItemFactory // out

	_listItemFactory = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ListItemFactory)

	return _listItemFactory
}

func (s dropDown) Model() gio.ListModel {
	var _arg0 *C.GtkDropDown // out
	var _cret *C.GListModel  // in

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_drop_down_get_model(_arg0)

	var _listModel gio.ListModel // out

	_listModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gio.ListModel)

	return _listModel
}

func (s dropDown) Selected() uint {
	var _arg0 *C.GtkDropDown // out
	var _cret C.guint        // in

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_drop_down_get_selected(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (s dropDown) SelectedItem() gextras.Objector {
	var _arg0 *C.GtkDropDown // out
	var _cret C.gpointer     // in

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_drop_down_get_selected_item(_arg0)

	var _object gextras.Objector // out

	_object = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gextras.Objector)

	return _object
}

func (s dropDown) SetEnableSearchDropDown(enableSearch bool) {
	var _arg0 *C.GtkDropDown // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))
	if enableSearch {
		_arg1 = C.TRUE
	}

	C.gtk_drop_down_set_enable_search(_arg0, _arg1)
}

func (s dropDown) SetExpressionDropDown(expression Expression) {
	var _arg0 *C.GtkDropDown   // out
	var _arg1 *C.GtkExpression // out

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))

	C.gtk_drop_down_set_expression(_arg0, _arg1)
}

func (s dropDown) SetFactoryDropDown(factory ListItemFactory) {
	var _arg0 *C.GtkDropDown        // out
	var _arg1 *C.GtkListItemFactory // out

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))

	C.gtk_drop_down_set_factory(_arg0, _arg1)
}

func (s dropDown) SetListFactoryDropDown(factory ListItemFactory) {
	var _arg0 *C.GtkDropDown        // out
	var _arg1 *C.GtkListItemFactory // out

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))

	C.gtk_drop_down_set_list_factory(_arg0, _arg1)
}

func (s dropDown) SetModelDropDown(model gio.ListModel) {
	var _arg0 *C.GtkDropDown // out
	var _arg1 *C.GListModel  // out

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	C.gtk_drop_down_set_model(_arg0, _arg1)
}

func (s dropDown) SetSelectedDropDown(position uint) {
	var _arg0 *C.GtkDropDown // out
	var _arg1 C.guint        // out

	_arg0 = (*C.GtkDropDown)(unsafe.Pointer(s.Native()))
	_arg1 = C.guint(position)

	C.gtk_drop_down_set_selected(_arg0, _arg1)
}

func (s dropDown) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s dropDown) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s dropDown) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s dropDown) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s dropDown) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s dropDown) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s dropDown) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b dropDown) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}