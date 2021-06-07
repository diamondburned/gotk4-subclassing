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
		{T: externglib.Type(C.gtk_filter_get_type()), F: marshalFilter},
	})
}

// Filter: a Filter object describes the filtering to be performed by a
// FilterListModel.
//
// The model will use the filter to determine if it should include items or not
// by calling gtk_filter_match() for each item and only keeping the ones that
// the function returns true for.
//
// Filters may change what items they match through their lifetime. In that
// case, they will emit the Filter::changed signal to notify that previous
// filter results are no longer valid and that items should be checked again via
// gtk_filter_match().
//
// GTK provides various pre-made filter implementations for common filtering
// operations. These filters often include properties that can be linked to
// various widgets to easily allow searches.
//
// However, in particular for large lists or complex search methods, it is also
// possible to subclass Filter and provide one's own filter.
type Filter interface {
	gextras.Objector

	// Changed emits the Filter::changed signal to notify all users of the
	// filter that the filter changed. Users of the filter should then check
	// items again via gtk_filter_match().
	//
	// Depending on the @change parameter, not all items need to be changed, but
	// only some. Refer to the FilterChange documentation for details.
	//
	// This function is intended for implementors of Filter subclasses and
	// should not be called from other functions.
	Changed(s Filter, change FilterChange)
	// Strictness gets the known strictness of @filters. If the strictness is
	// not known, GTK_FILTER_MATCH_SOME is returned.
	//
	// This value may change after emission of the Filter::changed signal.
	//
	// This function is meant purely for optimization purposes, filters can
	// choose to omit implementing it, but FilterListModel uses it.
	Strictness(s Filter)
	// Match checks if the given @item is matched by the filter or not.
	Match(s Filter, item gextras.Objector) bool
}

// filter implements the Filter interface.
type filter struct {
	gextras.Objector
}

var _ Filter = (*filter)(nil)

// WrapFilter wraps a GObject to the right type. It is
// primarily used internally.
func WrapFilter(obj *externglib.Object) Filter {
	return Filter{
		Objector: obj,
	}
}

func marshalFilter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFilter(obj), nil
}

// Changed emits the Filter::changed signal to notify all users of the
// filter that the filter changed. Users of the filter should then check
// items again via gtk_filter_match().
//
// Depending on the @change parameter, not all items need to be changed, but
// only some. Refer to the FilterChange documentation for details.
//
// This function is intended for implementors of Filter subclasses and
// should not be called from other functions.
func (s filter) Changed(s Filter, change FilterChange) {
	var arg0 *C.GtkFilter
	var arg1 C.GtkFilterChange

	arg0 = (*C.GtkFilter)(unsafe.Pointer(s.Native()))
	arg1 = (C.GtkFilterChange)(change)

	C.gtk_filter_changed(arg0, arg1)
}

// Strictness gets the known strictness of @filters. If the strictness is
// not known, GTK_FILTER_MATCH_SOME is returned.
//
// This value may change after emission of the Filter::changed signal.
//
// This function is meant purely for optimization purposes, filters can
// choose to omit implementing it, but FilterListModel uses it.
func (s filter) Strictness(s Filter) {
	var arg0 *C.GtkFilter

	arg0 = (*C.GtkFilter)(unsafe.Pointer(s.Native()))

	C.gtk_filter_get_strictness(arg0)
}

// Match checks if the given @item is matched by the filter or not.
func (s filter) Match(s Filter, item gextras.Objector) bool {
	var arg0 *C.GtkFilter
	var arg1 C.gpointer

	arg0 = (*C.GtkFilter)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GObject)(unsafe.Pointer(item.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_filter_match(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}
