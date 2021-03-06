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
		{T: externglib.Type(C.gtk_sort_list_model_get_type()), F: marshalSortListModeller},
	})
}

// SortListModelOverrider contains methods that are overridable.
type SortListModelOverrider interface {
}

// SortListModel: GtkSortListModel is a list model that sorts the elements of
// the underlying model according to a GtkSorter.
//
// The model can be set up to do incremental sorting, so that sorting long lists
// doesn't block the UI. See gtk.SortListModel.SetIncremental() for details.
//
// GtkSortListModel is a generic model and because of that it cannot take
// advantage of any external knowledge when sorting. If you run into performance
// issues with GtkSortListModel, it is strongly recommended that you write your
// own sorting list model.
type SortListModel struct {
	_ [0]func() // equal guard
	*externglib.Object

	gio.ListModel
}

var (
	_ externglib.Objector = (*SortListModel)(nil)
)

func classInitSortListModeller(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapSortListModel(obj *externglib.Object) *SortListModel {
	return &SortListModel{
		Object: obj,
		ListModel: gio.ListModel{
			Object: obj,
		},
	}
}

func marshalSortListModeller(p uintptr) (interface{}, error) {
	return wrapSortListModel(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSortListModel creates a new sort list model that uses the sorter to sort
// model.
//
// The function takes the following parameters:
//
//    - model (optional) to sort, or NULL.
//    - sorter (optional): GtkSorter to sort model with, or NULL.
//
// The function returns the following values:
//
//    - sortListModel: new GtkSortListModel.
//
func NewSortListModel(model gio.ListModeller, sorter *Sorter) *SortListModel {
	var _arg1 *C.GListModel       // out
	var _arg2 *C.GtkSorter        // out
	var _cret *C.GtkSortListModel // in

	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
		C.g_object_ref(C.gpointer(model.Native()))
	}
	if sorter != nil {
		_arg2 = (*C.GtkSorter)(unsafe.Pointer(sorter.Native()))
		C.g_object_ref(C.gpointer(sorter.Native()))
	}

	_cret = C.gtk_sort_list_model_new(_arg1, _arg2)
	runtime.KeepAlive(model)
	runtime.KeepAlive(sorter)

	var _sortListModel *SortListModel // out

	_sortListModel = wrapSortListModel(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _sortListModel
}

// Incremental returns whether incremental sorting is enabled.
//
// See gtk.SortListModel.SetIncremental().
//
// The function returns the following values:
//
//    - ok: TRUE if incremental sorting is enabled.
//
func (self *SortListModel) Incremental() bool {
	var _arg0 *C.GtkSortListModel // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkSortListModel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_sort_list_model_get_incremental(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Model gets the model currently sorted or NULL if none.
//
// The function returns the following values:
//
//    - listModel (optional): model that gets sorted.
//
func (self *SortListModel) Model() gio.ListModeller {
	var _arg0 *C.GtkSortListModel // out
	var _cret *C.GListModel       // in

	_arg0 = (*C.GtkSortListModel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_sort_list_model_get_model(_arg0)
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

// Pending estimates progress of an ongoing sorting operation.
//
// The estimate is the number of items that would still need to be sorted to
// finish the sorting operation if this was a linear algorithm. So this number
// is not related to how many items are already correctly sorted.
//
// If you want to estimate the progress, you can use code like this:
//
//    pending = gtk_sort_list_model_get_pending (self);
//    model = gtk_sort_list_model_get_model (self);
//    progress = 1.0 - pending / (double) MAX (1, g_list_model_get_n_items (model));
//
//
// If no sort operation is ongoing - in particular when
// gtk.SortListModel:incremental is FALSE - this function returns 0.
//
// The function returns the following values:
//
//    - guint progress estimate of remaining items to sort.
//
func (self *SortListModel) Pending() uint {
	var _arg0 *C.GtkSortListModel // out
	var _cret C.guint             // in

	_arg0 = (*C.GtkSortListModel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_sort_list_model_get_pending(_arg0)
	runtime.KeepAlive(self)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Sorter gets the sorter that is used to sort self.
//
// The function returns the following values:
//
//    - sorter (optional) of #self.
//
func (self *SortListModel) Sorter() *Sorter {
	var _arg0 *C.GtkSortListModel // out
	var _cret *C.GtkSorter        // in

	_arg0 = (*C.GtkSortListModel)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_sort_list_model_get_sorter(_arg0)
	runtime.KeepAlive(self)

	var _sorter *Sorter // out

	if _cret != nil {
		_sorter = wrapSorter(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _sorter
}

// SetIncremental sets the sort model to do an incremental sort.
//
// When incremental sorting is enabled, the GtkSortListModel will not do a
// complete sort immediately, but will instead queue an idle handler that
// incrementally sorts the items towards their correct position. This of course
// means that items do not instantly appear in the right place. It also means
// that the total sorting time is a lot slower.
//
// When your filter blocks the UI while sorting, you might consider turning this
// on. Depending on your model and sorters, this may become interesting around
// 10,000 to 100,000 items.
//
// By default, incremental sorting is disabled.
//
// See gtk.SortListModel.GetPending() for progress information about an ongoing
// incremental sorting operation.
//
// The function takes the following parameters:
//
//    - incremental: TRUE to sort incrementally.
//
func (self *SortListModel) SetIncremental(incremental bool) {
	var _arg0 *C.GtkSortListModel // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkSortListModel)(unsafe.Pointer(self.Native()))
	if incremental {
		_arg1 = C.TRUE
	}

	C.gtk_sort_list_model_set_incremental(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(incremental)
}

// SetModel sets the model to be sorted.
//
// The model's item type must conform to the item type of self.
//
// The function takes the following parameters:
//
//    - model (optional) to be sorted.
//
func (self *SortListModel) SetModel(model gio.ListModeller) {
	var _arg0 *C.GtkSortListModel // out
	var _arg1 *C.GListModel       // out

	_arg0 = (*C.GtkSortListModel)(unsafe.Pointer(self.Native()))
	if model != nil {
		_arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))
	}

	C.gtk_sort_list_model_set_model(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(model)
}

// SetSorter sets a new sorter on self.
//
// The function takes the following parameters:
//
//    - sorter (optional): GtkSorter to sort model with.
//
func (self *SortListModel) SetSorter(sorter *Sorter) {
	var _arg0 *C.GtkSortListModel // out
	var _arg1 *C.GtkSorter        // out

	_arg0 = (*C.GtkSortListModel)(unsafe.Pointer(self.Native()))
	if sorter != nil {
		_arg1 = (*C.GtkSorter)(unsafe.Pointer(sorter.Native()))
	}

	C.gtk_sort_list_model_set_sorter(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(sorter)
}
