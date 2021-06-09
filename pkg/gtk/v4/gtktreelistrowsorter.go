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
		{T: externglib.Type(C.gtk_tree_list_row_sorter_get_type()), F: marshalTreeListRowSorter},
	})
}

// TreeListRowSorter is a special-purpose sorter that will apply a given sorter
// to the levels in a tree, while respecting the tree structure.
//
// Here is an example for setting up a column view with a tree model and a
// GtkTreeListSorter:
//
//    column_sorter = gtk_column_view_get_sorter (view);
//    sorter = gtk_tree_list_row_sorter_new (g_object_ref (column_sorter));
//    sort_model = gtk_sort_list_model_new (tree_model, sorter);
//    selection = gtk_single_selection_new (sort_model);
//    gtk_column_view_set_model (view, G_LIST_MODEL (selection));
type TreeListRowSorter interface {
	Sorter

	// Sorter returns the sorter used by @self.
	Sorter() Sorter
	// SetSorter sets the sorter to use for items with the same parent.
	//
	// This sorter will be passed the TreeListRow:item of the tree list rows
	// passed to @self.
	SetSorter(sorter Sorter)
}

// treeListRowSorter implements the TreeListRowSorter interface.
type treeListRowSorter struct {
	Sorter
}

var _ TreeListRowSorter = (*treeListRowSorter)(nil)

// WrapTreeListRowSorter wraps a GObject to the right type. It is
// primarily used internally.
func WrapTreeListRowSorter(obj *externglib.Object) TreeListRowSorter {
	return TreeListRowSorter{
		Sorter: WrapSorter(obj),
	}
}

func marshalTreeListRowSorter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeListRowSorter(obj), nil
}

// NewTreeListRowSorter constructs a class TreeListRowSorter.
func NewTreeListRowSorter(sorter Sorter) TreeListRowSorter {
	var arg1 *C.GtkSorter

	arg1 = (*C.GtkSorter)(unsafe.Pointer(sorter.Native()))

	var cret C.GtkTreeListRowSorter

	cret = C.gtk_tree_list_row_sorter_new(arg1)

	var treeListRowSorter TreeListRowSorter

	treeListRowSorter = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(TreeListRowSorter)

	return treeListRowSorter
}

// Sorter returns the sorter used by @self.
func (s treeListRowSorter) Sorter() Sorter {
	var arg0 *C.GtkTreeListRowSorter

	arg0 = (*C.GtkTreeListRowSorter)(unsafe.Pointer(s.Native()))

	var cret *C.GtkSorter

	cret = C.gtk_tree_list_row_sorter_get_sorter(arg0)

	var sorter Sorter

	sorter = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Sorter)

	return sorter
}

// SetSorter sets the sorter to use for items with the same parent.
//
// This sorter will be passed the TreeListRow:item of the tree list rows
// passed to @self.
func (s treeListRowSorter) SetSorter(sorter Sorter) {
	var arg0 *C.GtkTreeListRowSorter
	var arg1 *C.GtkSorter

	arg0 = (*C.GtkTreeListRowSorter)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkSorter)(unsafe.Pointer(sorter.Native()))

	C.gtk_tree_list_row_sorter_set_sorter(arg0, arg1)
}
