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
		{T: externglib.Type(C.gtk_multi_selection_get_type()), F: marshalMultiSelection},
	})
}

// MultiSelection: `GtkMultiSelection` is a `GtkSelectionModel` that allows
// selecting multiple elements.
type MultiSelection interface {
	gextras.Objector
	gio.ListModel
	SelectionModel

	// Model returns the underlying model of @self.
	Model() gio.ListModel
	// SetModel sets the model that @self should wrap.
	//
	// If @model is nil, @self will be empty.
	SetModel(model gio.ListModel)
}

// multiSelection implements the MultiSelection class.
type multiSelection struct {
	gextras.Objector
	gio.ListModel
	SelectionModel
}

var _ MultiSelection = (*multiSelection)(nil)

// WrapMultiSelection wraps a GObject to the right type. It is
// primarily used internally.
func WrapMultiSelection(obj *externglib.Object) MultiSelection {
	return multiSelection{
		Objector:       obj,
		gio.ListModel:  gio.WrapListModel(obj),
		SelectionModel: WrapSelectionModel(obj),
	}
}

func marshalMultiSelection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMultiSelection(obj), nil
}

// NewMultiSelection constructs a class MultiSelection.
func NewMultiSelection(model gio.ListModel) MultiSelection {
	var _arg1 *C.GListModel // out

	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	var _cret C.GtkMultiSelection // in

	_cret = C.gtk_multi_selection_new(_arg1)

	var _multiSelection MultiSelection // out

	_multiSelection = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(MultiSelection)

	return _multiSelection
}

// Model returns the underlying model of @self.
func (s multiSelection) Model() gio.ListModel {
	var _arg0 *C.GtkMultiSelection // out

	_arg0 = (*C.GtkMultiSelection)(unsafe.Pointer(s.Native()))

	var _cret *C.GListModel // in

	_cret = C.gtk_multi_selection_get_model(_arg0)

	var _listModel gio.ListModel // out

	_listModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gio.ListModel)

	return _listModel
}

// SetModel sets the model that @self should wrap.
//
// If @model is nil, @self will be empty.
func (s multiSelection) SetModel(model gio.ListModel) {
	var _arg0 *C.GtkMultiSelection // out
	var _arg1 *C.GListModel        // out

	_arg0 = (*C.GtkMultiSelection)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	C.gtk_multi_selection_set_model(_arg0, _arg1)
}
