// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_slice_list_model_get_type()), F: marshalSliceListModeller},
	})
}

// SliceListModel: GtkSliceListModel is a list model that presents a slice of
// another model.
//
// This is useful when implementing paging by setting the size to the number of
// elements per page and updating the offset whenever a different page is
// opened.
type SliceListModel struct {
	_ [0]func() // equal guard
	*externglib.Object

	gio.ListModel
}

var (
	_ externglib.Objector = (*SliceListModel)(nil)
)

func wrapSliceListModel(obj *externglib.Object) *SliceListModel {
	return &SliceListModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalSliceListModeller(p uintptr) (interface{}, error) {
	return wrapSliceListModel(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSliceListModel creates a new slice model.
//
// It presents the slice from offset to offset + size of the given model.
//
// The function takes the following parameters:
//
//    - model (optional) to use, or NULL.
//    - offset of the slice.
//    - size: maximum size of the slice.
//
// The function returns the following values:
//
//    - sliceListModel: new GtkSliceListModel.
//
func NewSliceListModel(model gio.ListModeller, offset, size uint) *SliceListModel {
	var _arg1 *C.GListModel        // out
	var _arg2 C.guint              // out
	var _arg3 C.guint              // out
	var _cret *C.GtkSliceListModel // in

	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
		C.g_object_ref(C.gpointer(model.Native()))
	}
	_arg2 = C.guint(offset)
	_arg3 = C.guint(size)

	_cret = C.gtk_slice_list_model_new(_arg1, _arg2, _arg3)
	runtime.KeepAlive(model)
	runtime.KeepAlive(offset)
	runtime.KeepAlive(size)

	var _sliceListModel *SliceListModel // out

	_sliceListModel = wrapSliceListModel(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _sliceListModel
}

// Model gets the model that is currently being used or NULL if none.
//
// The function returns the following values:
//
//    - listModel (optional): model in use.
//
func (self *SliceListModel) Model() gio.ListModeller {
	var _arg0 *C.GtkSliceListModel // out
	var _cret *C.GListModel        // in

	_arg0 = (*C.GtkSliceListModel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_slice_list_model_get_model(_arg0)
	runtime.KeepAlive(self)

	var _listModel gio.ListModeller // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(gio.ListModeller)
				return ok
			})
			rv, ok := casted.(gio.ListModeller)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.ListModeller")
			}
			_listModel = rv
		}
	}

	return _listModel
}

// Offset gets the offset set via gtk_slice_list_model_set_offset().
//
// The function returns the following values:
//
//    - guint: offset.
//
func (self *SliceListModel) Offset() uint {
	var _arg0 *C.GtkSliceListModel // out
	var _cret C.guint              // in

	_arg0 = (*C.GtkSliceListModel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_slice_list_model_get_offset(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Size gets the size set via gtk_slice_list_model_set_size().
//
// The function returns the following values:
//
//    - guint: size.
//
func (self *SliceListModel) Size() uint {
	var _arg0 *C.GtkSliceListModel // out
	var _cret C.guint              // in

	_arg0 = (*C.GtkSliceListModel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_slice_list_model_get_size(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// SetModel sets the model to show a slice of.
//
// The model's item type must conform to self's item type.
//
// The function takes the following parameters:
//
//    - model (optional) to be sliced.
//
func (self *SliceListModel) SetModel(model gio.ListModeller) {
	var _arg0 *C.GtkSliceListModel // out
	var _arg1 *C.GListModel        // out

	_arg0 = (*C.GtkSliceListModel)(unsafe.Pointer(self.Native()))
	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
	}

	C.gtk_slice_list_model_set_model(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}

// SetOffset sets the offset into the original model for this slice.
//
// If the offset is too large for the sliced model, self will end up empty.
//
// The function takes the following parameters:
//
//    - offset: new offset to use.
//
func (self *SliceListModel) SetOffset(offset uint) {
	var _arg0 *C.GtkSliceListModel // out
	var _arg1 C.guint              // out

	_arg0 = (*C.GtkSliceListModel)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(offset)

	C.gtk_slice_list_model_set_offset(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(offset)
}

// SetSize sets the maximum size. self will never have more items than size.
//
// It can however have fewer items if the offset is too large or the model
// sliced from doesn't have enough items.
//
// The function takes the following parameters:
//
//    - size: maximum size.
//
func (self *SliceListModel) SetSize(size uint) {
	var _arg0 *C.GtkSliceListModel // out
	var _arg1 C.guint              // out

	_arg0 = (*C.GtkSliceListModel)(unsafe.Pointer(self.Native()))
	_arg1 = C.guint(size)

	C.gtk_slice_list_model_set_size(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(size)
}
