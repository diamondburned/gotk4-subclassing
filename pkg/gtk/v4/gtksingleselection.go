// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_single_selection_get_type()), F: marshalSingleSelection},
	})
}

// SingleSelection: gtkSingleSelection is an implementation of the
// SelectionModel interface that allows selecting a single element. It is the
// default selection method used by list widgets in GTK.
//
// Note that the selection is *persistent* -- if the selected item is removed
// and re-added in the same Model::items-changed emission, it stays selected. In
// particular, this means that changing the sort order of an underlying sort
// model will preserve the selection.
type SingleSelection interface {
	gextras.Objector
	gio.ListModel
	SelectionModel

	// Autoselect checks if autoselect has been enabled or disabled via
	// gtk_single_selection_set_autoselect().
	Autoselect(s SingleSelection) bool
	// CanUnselect: if true, gtk_selection_model_unselect_item() is supported
	// and allows unselecting the selected item.
	CanUnselect(s SingleSelection) bool
	// Model gets the model that @self is wrapping.
	Model(s SingleSelection)
	// Selected gets the position of the selected item. If no item is selected,
	// K_INVALID_LIST_POSITION is returned.
	Selected(s SingleSelection)
	// SelectedItem gets the selected item.
	//
	// If no item is selected, nil is returned.
	SelectedItem(s SingleSelection)
	// SetAutoselect: if @autoselect is true, @self will enforce that an item is
	// always selected. It will select a new item when the currently selected
	// item is deleted and it will disallow unselecting the current item.
	SetAutoselect(s SingleSelection, autoselect bool)
	// SetCanUnselect: if true, unselecting the current item via
	// gtk_selection_model_unselect_item() is supported.
	//
	// Note that setting SingleSelection:autoselect will cause the unselecting
	// to not work, so it practically makes no sense to set both at the same
	// time the same time.
	SetCanUnselect(s SingleSelection, canUnselect bool)
	// SetModel sets the model that @self should wrap. If @model is nil, @self
	// will be empty.
	SetModel(s SingleSelection, model gio.ListModel)
	// SetSelected selects the item at the given position.
	//
	// If the list does not have an item at @position or K_INVALID_LIST_POSITION
	// is given, the behavior depends on the value of the
	// SingleSelection:autoselect property: If it is set, no change will occur
	// and the old item will stay selected. If it is unset, the selection will
	// be unset and no item will be selected.
	SetSelected(s SingleSelection, position uint)
}

// singleSelection implements the SingleSelection interface.
type singleSelection struct {
	gextras.Objector
	gio.ListModel
	SelectionModel
}

var _ SingleSelection = (*singleSelection)(nil)

// WrapSingleSelection wraps a GObject to the right type. It is
// primarily used internally.
func WrapSingleSelection(obj *externglib.Object) SingleSelection {
	return SingleSelection{
		Objector:       obj,
		gio.ListModel:  gio.WrapListModel(obj),
		SelectionModel: WrapSelectionModel(obj),
	}
}

func marshalSingleSelection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSingleSelection(obj), nil
}

// NewSingleSelection constructs a class SingleSelection.
func NewSingleSelection(model gio.ListModel) {
	var arg1 *C.GListModel

	arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	C.gtk_single_selection_new(arg1)
}

// Autoselect checks if autoselect has been enabled or disabled via
// gtk_single_selection_set_autoselect().
func (s singleSelection) Autoselect(s SingleSelection) bool {
	var arg0 *C.GtkSingleSelection

	arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_single_selection_get_autoselect(arg0)

	if cret {
		ok = true
	}

	return ok
}

// CanUnselect: if true, gtk_selection_model_unselect_item() is supported
// and allows unselecting the selected item.
func (s singleSelection) CanUnselect(s SingleSelection) bool {
	var arg0 *C.GtkSingleSelection

	arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_single_selection_get_can_unselect(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Model gets the model that @self is wrapping.
func (s singleSelection) Model(s SingleSelection) {
	var arg0 *C.GtkSingleSelection

	arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(s.Native()))

	C.gtk_single_selection_get_model(arg0)
}

// Selected gets the position of the selected item. If no item is selected,
// K_INVALID_LIST_POSITION is returned.
func (s singleSelection) Selected(s SingleSelection) {
	var arg0 *C.GtkSingleSelection

	arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(s.Native()))

	C.gtk_single_selection_get_selected(arg0)
}

// SelectedItem gets the selected item.
//
// If no item is selected, nil is returned.
func (s singleSelection) SelectedItem(s SingleSelection) {
	var arg0 *C.GtkSingleSelection

	arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(s.Native()))

	C.gtk_single_selection_get_selected_item(arg0)
}

// SetAutoselect: if @autoselect is true, @self will enforce that an item is
// always selected. It will select a new item when the currently selected
// item is deleted and it will disallow unselecting the current item.
func (s singleSelection) SetAutoselect(s SingleSelection, autoselect bool) {
	var arg0 *C.GtkSingleSelection
	var arg1 C.gboolean

	arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(s.Native()))
	if autoselect {
		arg1 = C.gboolean(1)
	}

	C.gtk_single_selection_set_autoselect(arg0, arg1)
}

// SetCanUnselect: if true, unselecting the current item via
// gtk_selection_model_unselect_item() is supported.
//
// Note that setting SingleSelection:autoselect will cause the unselecting
// to not work, so it practically makes no sense to set both at the same
// time the same time.
func (s singleSelection) SetCanUnselect(s SingleSelection, canUnselect bool) {
	var arg0 *C.GtkSingleSelection
	var arg1 C.gboolean

	arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(s.Native()))
	if canUnselect {
		arg1 = C.gboolean(1)
	}

	C.gtk_single_selection_set_can_unselect(arg0, arg1)
}

// SetModel sets the model that @self should wrap. If @model is nil, @self
// will be empty.
func (s singleSelection) SetModel(s SingleSelection, model gio.ListModel) {
	var arg0 *C.GtkSingleSelection
	var arg1 *C.GListModel

	arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	C.gtk_single_selection_set_model(arg0, arg1)
}

// SetSelected selects the item at the given position.
//
// If the list does not have an item at @position or K_INVALID_LIST_POSITION
// is given, the behavior depends on the value of the
// SingleSelection:autoselect property: If it is set, no change will occur
// and the old item will stay selected. If it is unset, the selection will
// be unset and no item will be selected.
func (s singleSelection) SetSelected(s SingleSelection, position uint) {
	var arg0 *C.GtkSingleSelection
	var arg1 C.guint

	arg0 = (*C.GtkSingleSelection)(unsafe.Pointer(s.Native()))
	arg1 = C.guint(position)

	C.gtk_single_selection_set_selected(arg0, arg1)
}
