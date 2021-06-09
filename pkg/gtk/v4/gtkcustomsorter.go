// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_custom_sorter_get_type()), F: marshalCustomSorter},
	})
}

// CustomSorter: gtkCustomSorter is a Sorter implementation that sorts via a
// traditional DataFunc callback.
type CustomSorter interface {
	Sorter

	// SetSortFunc sets (or unsets) the function used for sorting items.
	//
	// If @sort_func is nil, all items are considered equal.
	//
	// If the sort func changes its sorting behavior, gtk_sorter_changed() needs
	// to be called.
	//
	// If a previous function was set, its @user_destroy will be called now.
	SetSortFunc()
}

// customSorter implements the CustomSorter interface.
type customSorter struct {
	Sorter
}

var _ CustomSorter = (*customSorter)(nil)

// WrapCustomSorter wraps a GObject to the right type. It is
// primarily used internally.
func WrapCustomSorter(obj *externglib.Object) CustomSorter {
	return CustomSorter{
		Sorter: WrapSorter(obj),
	}
}

func marshalCustomSorter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCustomSorter(obj), nil
}

// NewCustomSorter constructs a class CustomSorter.
func NewCustomSorter() CustomSorter {
	var cret C.GtkCustomSorter

	cret = C.gtk_custom_sorter_new()

	var customSorter CustomSorter

	customSorter = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(CustomSorter)

	return customSorter
}

// SetSortFunc sets (or unsets) the function used for sorting items.
//
// If @sort_func is nil, all items are considered equal.
//
// If the sort func changes its sorting behavior, gtk_sorter_changed() needs
// to be called.
//
// If a previous function was set, its @user_destroy will be called now.
func (s customSorter) SetSortFunc() {
	var arg0 *C.GtkCustomSorter

	arg0 = (*C.GtkCustomSorter)(unsafe.Pointer(s.Native()))

	C.gtk_custom_sorter_set_sort_func(arg0)
}
