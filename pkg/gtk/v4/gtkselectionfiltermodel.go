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
		{T: externglib.Type(C.gtk_selection_filter_model_get_type()), F: marshalSelectionFilterModel},
	})
}

// SelectionFilterModel is a list model that presents the selected items in a
// SelectionModel as its own list model.
type SelectionFilterModel interface {
	gextras.Objector
	gio.ListModel

	// Model gets the model currently filtered or nil if none.
	Model(s SelectionFilterModel)
	// SetModel sets the model to be filtered.
	//
	// Note that GTK makes no effort to ensure that @model conforms to the item
	// type of @self. It assumes that the caller knows what they are doing and
	// have set up an appropriate filter to ensure that item types match.
	SetModel(s SelectionFilterModel, model SelectionModel)
}

// selectionFilterModel implements the SelectionFilterModel interface.
type selectionFilterModel struct {
	gextras.Objector
	gio.ListModel
}

var _ SelectionFilterModel = (*selectionFilterModel)(nil)

// WrapSelectionFilterModel wraps a GObject to the right type. It is
// primarily used internally.
func WrapSelectionFilterModel(obj *externglib.Object) SelectionFilterModel {
	return SelectionFilterModel{
		Objector:      obj,
		gio.ListModel: gio.WrapListModel(obj),
	}
}

func marshalSelectionFilterModel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSelectionFilterModel(obj), nil
}

// NewSelectionFilterModel constructs a class SelectionFilterModel.
func NewSelectionFilterModel(model SelectionModel) {
	var arg1 *C.GtkSelectionModel

	arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))

	C.gtk_selection_filter_model_new(arg1)
}

// Model gets the model currently filtered or nil if none.
func (s selectionFilterModel) Model(s SelectionFilterModel) {
	var arg0 *C.GtkSelectionFilterModel

	arg0 = (*C.GtkSelectionFilterModel)(unsafe.Pointer(s.Native()))

	C.gtk_selection_filter_model_get_model(arg0)
}

// SetModel sets the model to be filtered.
//
// Note that GTK makes no effort to ensure that @model conforms to the item
// type of @self. It assumes that the caller knows what they are doing and
// have set up an appropriate filter to ensure that item types match.
func (s selectionFilterModel) SetModel(s SelectionFilterModel, model SelectionModel) {
	var arg0 *C.GtkSelectionFilterModel
	var arg1 *C.GtkSelectionModel

	arg0 = (*C.GtkSelectionFilterModel)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))

	C.gtk_selection_filter_model_set_model(arg0, arg1)
}