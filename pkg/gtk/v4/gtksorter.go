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
		{T: externglib.Type(C.gtk_sorter_get_type()), F: marshalSorter},
	})
}

// Sorter is the way to describe sorting criteria. Its primary user is
// SortListModel.
//
// The model will use a sorter to determine the order in which its items should
// appear by calling gtk_sorter_compare() for pairs of items.
//
// Sorters may change their sorting behavior through their lifetime. In that
// case, they will emit the Sorter::changed signal to notify that the sort order
// is no longer valid and should be updated by calling gtk_sorter_compare()
// again.
//
// GTK provides various pre-made sorter implementations for common sorting
// operations. ColumnView has built-in support for sorting lists via the
// ColumnViewColumn:sorter property, where the user can change the sorting by
// clicking on list headers.
//
// Of course, in particular for large lists, it is also possible to subclass
// Sorter and provide one's own sorter.
type Sorter interface {
	gextras.Objector

	// Changed emits the Sorter::changed signal to notify all users of the
	// sorter that it has changed. Users of the sorter should then update the
	// sort order via gtk_sorter_compare().
	//
	// Depending on the @change parameter, it may be possible to update the sort
	// order without a full resorting. Refer to the SorterChange documentation
	// for details.
	//
	// This function is intended for implementors of Sorter subclasses and
	// should not be called from other functions.
	Changed(s Sorter, change SorterChange)
	// Compare compares two given items according to the sort order implemented
	// by the sorter.
	//
	// Sorters implement a partial order: * It is reflexive, ie a = a * It is
	// antisymmetric, ie if a < b and b < a, then a = b * It is transitive, ie
	// given any 3 items with a ≤ b and b ≤ c, then a ≤ c
	//
	// The sorter may signal it conforms to additional constraints via the
	// return value of gtk_sorter_get_order().
	Compare(s Sorter, item1 gextras.Objector, item2 gextras.Objector)
	// Order gets the order that @self conforms to. See SorterOrder for details
	// of the possible return values.
	//
	// This function is intended to allow optimizations.
	Order(s Sorter)
}

// sorter implements the Sorter interface.
type sorter struct {
	gextras.Objector
}

var _ Sorter = (*sorter)(nil)

// WrapSorter wraps a GObject to the right type. It is
// primarily used internally.
func WrapSorter(obj *externglib.Object) Sorter {
	return Sorter{
		Objector: obj,
	}
}

func marshalSorter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSorter(obj), nil
}

// Changed emits the Sorter::changed signal to notify all users of the
// sorter that it has changed. Users of the sorter should then update the
// sort order via gtk_sorter_compare().
//
// Depending on the @change parameter, it may be possible to update the sort
// order without a full resorting. Refer to the SorterChange documentation
// for details.
//
// This function is intended for implementors of Sorter subclasses and
// should not be called from other functions.
func (s sorter) Changed(s Sorter, change SorterChange) {
	var arg0 *C.GtkSorter
	var arg1 C.GtkSorterChange

	arg0 = (*C.GtkSorter)(unsafe.Pointer(s.Native()))
	arg1 = (C.GtkSorterChange)(change)

	C.gtk_sorter_changed(arg0, arg1)
}

// Compare compares two given items according to the sort order implemented
// by the sorter.
//
// Sorters implement a partial order: * It is reflexive, ie a = a * It is
// antisymmetric, ie if a < b and b < a, then a = b * It is transitive, ie
// given any 3 items with a ≤ b and b ≤ c, then a ≤ c
//
// The sorter may signal it conforms to additional constraints via the
// return value of gtk_sorter_get_order().
func (s sorter) Compare(s Sorter, item1 gextras.Objector, item2 gextras.Objector) {
	var arg0 *C.GtkSorter
	var arg1 C.gpointer
	var arg2 C.gpointer

	arg0 = (*C.GtkSorter)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GObject)(unsafe.Pointer(item1.Native()))
	arg2 = (*C.GObject)(unsafe.Pointer(item2.Native()))

	C.gtk_sorter_compare(arg0, arg1, arg2)
}

// Order gets the order that @self conforms to. See SorterOrder for details
// of the possible return values.
//
// This function is intended to allow optimizations.
func (s sorter) Order(s Sorter) {
	var arg0 *C.GtkSorter

	arg0 = (*C.GtkSorter)(unsafe.Pointer(s.Native()))

	C.gtk_sorter_get_order(arg0)
}
