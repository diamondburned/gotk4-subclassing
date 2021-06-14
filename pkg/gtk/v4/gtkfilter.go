// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_filter_change_get_type()), F: marshalFilterChange},
		{T: externglib.Type(C.gtk_filter_match_get_type()), F: marshalFilterMatch},
		{T: externglib.Type(C.gtk_filter_get_type()), F: marshalFilter},
	})
}

// FilterChange describes changes in a filter in more detail and allows objects
// using the filter to optimize refiltering items.
//
// If you are writing an implementation and are not sure which value to pass,
// GTK_FILTER_CHANGE_DIFFERENT is always a correct choice.
type FilterChange int

const (
	// FilterChangeDifferent: the filter change cannot be described with any of
	// the other enumeration values.
	FilterChangeDifferent FilterChange = 0
	// FilterChangeLessStrict: the filter is less strict than it was before: All
	// items that it used to return true for still return true, others now may,
	// too.
	FilterChangeLessStrict FilterChange = 1
	// FilterChangeMoreStrict: the filter is more strict than it was before: All
	// items that it used to return false for still return false, others now
	// may, too.
	FilterChangeMoreStrict FilterChange = 2
)

func marshalFilterChange(p uintptr) (interface{}, error) {
	return FilterChange(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// FilterMatch describes the known strictness of a filter.
//
// Note that for filters where the strictness is not known,
// GTK_FILTER_MATCH_SOME is always an acceptable value, even if a filter does
// match all or no items.
type FilterMatch int

const (
	// FilterMatchSome: the filter matches some items, gtk_filter_match() may
	// return true or false
	FilterMatchSome FilterMatch = 0
	// FilterMatchNone: the filter does not match any item, gtk_filter_match()
	// will always return false.
	FilterMatchNone FilterMatch = 1
	// FilterMatchAll: the filter matches all items, gtk_filter_match() will
	// alays return true.
	FilterMatchAll FilterMatch = 2
)

func marshalFilterMatch(p uintptr) (interface{}, error) {
	return FilterMatch(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Filter: a `GtkFilter` object describes the filtering to be performed by a
// `GtkFilterListModel`.
//
// The model will use the filter to determine if it should include items or not
// by calling [method@Gtk.Filter.match] for each item and only keeping the ones
// that the function returns true for.
//
// Filters may change what items they match through their lifetime. In that
// case, they will emit the [signal@Gtk.Filter::changed] signal to notify that
// previous filter results are no longer valid and that items should be checked
// again via [method@Gtk.Filter.match].
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
	Changed(change FilterChange)
	// Strictness gets the known strictness of @filters. If the strictness is
	// not known, GTK_FILTER_MATCH_SOME is returned.
	//
	// This value may change after emission of the Filter::changed signal.
	//
	// This function is meant purely for optimization purposes, filters can
	// choose to omit implementing it, but FilterListModel uses it.
	Strictness() FilterMatch
	// Match checks if the given @item is matched by the filter or not.
	Match(item gextras.Objector) bool
}

// filter implements the Filter class.
type filter struct {
	gextras.Objector
}

var _ Filter = (*filter)(nil)

// WrapFilter wraps a GObject to the right type. It is
// primarily used internally.
func WrapFilter(obj *externglib.Object) Filter {
	return filter{
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
func (s filter) Changed(change FilterChange) {
	var _arg0 *C.GtkFilter      // out
	var _arg1 C.GtkFilterChange // out

	_arg0 = (*C.GtkFilter)(unsafe.Pointer(s.Native()))
	_arg1 = (C.GtkFilterChange)(change)

	C.gtk_filter_changed(_arg0, _arg1)
}

// Strictness gets the known strictness of @filters. If the strictness is
// not known, GTK_FILTER_MATCH_SOME is returned.
//
// This value may change after emission of the Filter::changed signal.
//
// This function is meant purely for optimization purposes, filters can
// choose to omit implementing it, but FilterListModel uses it.
func (s filter) Strictness() FilterMatch {
	var _arg0 *C.GtkFilter // out

	_arg0 = (*C.GtkFilter)(unsafe.Pointer(s.Native()))

	var _cret C.GtkFilterMatch // in

	_cret = C.gtk_filter_get_strictness(_arg0)

	var _filterMatch FilterMatch // out

	_filterMatch = FilterMatch(_cret)

	return _filterMatch
}

// Match checks if the given @item is matched by the filter or not.
func (s filter) Match(item gextras.Objector) bool {
	var _arg0 *C.GtkFilter // out
	var _arg1 C.gpointer   // out

	_arg0 = (*C.GtkFilter)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GObject)(unsafe.Pointer(item.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_filter_match(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
