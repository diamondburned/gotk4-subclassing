// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/internal/box"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_sortable_get_type()), F: marshalTreeSortable},
	})
}

// TreeIterCompareFunc: a GtkTreeIterCompareFunc should return a negative
// integer, zero, or a positive integer if @a sorts before @b, @a sorts with @b,
// or @a sorts after @b respectively. If two iters compare as equal, their order
// in the sorted model is undefined. In order to ensure that the TreeSortable
// behaves as expected, the GtkTreeIterCompareFunc must define a partial order
// on the model, i.e. it must be reflexive, antisymmetric and transitive.
//
// For example, if @model is a product catalogue, then a compare function for
// the “price” column could be one which returns `price_of(@a) - price_of(@b)`.
type TreeIterCompareFunc func(model TreeModel, a *TreeIter, b *TreeIter) int

//export gotk4_TreeIterCompareFunc
func gotk4_TreeIterCompareFunc(arg0 *C.GtkTreeModel, arg1 *C.GtkTreeIter, arg2 *C.GtkTreeIter, arg3 C.gpointer) C.gint {
	v := box.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(TreeIterCompareFunc)
	ret := fn(model, a, b, userData)

	cret = C.gint(ret)

	return cret
}

// TreeSortableOverrider contains methods that are overridable. This
// interface is a subset of the interface TreeSortable.
type TreeSortableOverrider interface {
	// SortColumnID fills in @sort_column_id and @order with the current sort
	// column and the order. It returns true unless the @sort_column_id is
	// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID or
	// GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID.
	SortColumnID(s TreeSortable) (sortColumnID int, order *SortType, ok bool)
	// HasDefaultSortFunc returns true if the model has a default sort function.
	// This is used primarily by GtkTreeViewColumns in order to determine if a
	// model can go back to the default state, or not.
	HasDefaultSortFunc(s TreeSortable) bool
	// SetDefaultSortFunc sets the default comparison function used when sorting
	// to be @sort_func. If the current sort column id of @sortable is
	// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID, then the model will sort using
	// this function.
	//
	// If @sort_func is nil, then there will be no default comparison function.
	// This means that once the model has been sorted, it can’t go back to the
	// default state. In this case, when the current sort column id of @sortable
	// is GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID, the model will be unsorted.
	SetDefaultSortFunc(s TreeSortable)
	// SetSortColumnID sets the current sort column to be @sort_column_id. The
	// @sortable will resort itself to reflect this change, after emitting a
	// TreeSortable::sort-column-changed signal. @sort_column_id may either be a
	// regular column id, or one of the following special values:
	//
	// - GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID: the default sort function
	// will be used, if it is set
	//
	// - GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID: no sorting will occur
	SetSortColumnID(s TreeSortable, sortColumnID int, order SortType)
	// SetSortFunc sets the comparison function used when sorting to be
	// @sort_func. If the current sort column id of @sortable is the same as
	// @sort_column_id, then the model will sort using this function.
	SetSortFunc(s TreeSortable)
	// SortColumnChanged emits a TreeSortable::sort-column-changed signal on
	// @sortable.
	SortColumnChanged(s TreeSortable)
}

// TreeSortable is an interface to be implemented by tree models which support
// sorting. The TreeView uses the methods provided by this interface to sort the
// model.
type TreeSortable interface {
	TreeModel
	TreeSortableOverrider
}

// treeSortable implements the TreeSortable interface.
type treeSortable struct {
	TreeModel
}

var _ TreeSortable = (*treeSortable)(nil)

// WrapTreeSortable wraps a GObject to a type that implements interface
// TreeSortable. It is primarily used internally.
func WrapTreeSortable(obj *externglib.Object) TreeSortable {
	return TreeSortable{
		TreeModel: WrapTreeModel(obj),
	}
}

func marshalTreeSortable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeSortable(obj), nil
}

// SortColumnID fills in @sort_column_id and @order with the current sort
// column and the order. It returns true unless the @sort_column_id is
// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID or
// GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID.
func (s treeSortable) SortColumnID(s TreeSortable) (sortColumnID int, order *SortType, ok bool) {
	var arg0 *C.GtkTreeSortable

	arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(s.Native()))

	var arg1 C.gint
	var sortColumnID int
	var arg2 C.GtkSortType
	var order *SortType
	var cret C.gboolean
	var ok bool

	cret = C.gtk_tree_sortable_get_sort_column_id(arg0, &arg1, &arg2)

	sortColumnID = int(&arg1)
	order = *SortType(&arg2)
	if cret {
		ok = true
	}

	return sortColumnID, order, ok
}

// HasDefaultSortFunc returns true if the model has a default sort function.
// This is used primarily by GtkTreeViewColumns in order to determine if a
// model can go back to the default state, or not.
func (s treeSortable) HasDefaultSortFunc(s TreeSortable) bool {
	var arg0 *C.GtkTreeSortable

	arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_tree_sortable_has_default_sort_func(arg0)

	if cret {
		ok = true
	}

	return ok
}

// SetDefaultSortFunc sets the default comparison function used when sorting
// to be @sort_func. If the current sort column id of @sortable is
// GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID, then the model will sort using
// this function.
//
// If @sort_func is nil, then there will be no default comparison function.
// This means that once the model has been sorted, it can’t go back to the
// default state. In this case, when the current sort column id of @sortable
// is GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID, the model will be unsorted.
func (s treeSortable) SetDefaultSortFunc(s TreeSortable) {
	var arg0 *C.GtkTreeSortable

	arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(s.Native()))

	C.gtk_tree_sortable_set_default_sort_func(arg0, arg1, arg2, arg3)
}

// SetSortColumnID sets the current sort column to be @sort_column_id. The
// @sortable will resort itself to reflect this change, after emitting a
// TreeSortable::sort-column-changed signal. @sort_column_id may either be a
// regular column id, or one of the following special values:
//
// - GTK_TREE_SORTABLE_DEFAULT_SORT_COLUMN_ID: the default sort function
// will be used, if it is set
//
// - GTK_TREE_SORTABLE_UNSORTED_SORT_COLUMN_ID: no sorting will occur
func (s treeSortable) SetSortColumnID(s TreeSortable, sortColumnID int, order SortType) {
	var arg0 *C.GtkTreeSortable
	var arg1 C.gint
	var arg2 C.GtkSortType

	arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(s.Native()))
	arg1 = C.gint(sortColumnID)
	arg2 = (C.GtkSortType)(order)

	C.gtk_tree_sortable_set_sort_column_id(arg0, arg1, arg2)
}

// SetSortFunc sets the comparison function used when sorting to be
// @sort_func. If the current sort column id of @sortable is the same as
// @sort_column_id, then the model will sort using this function.
func (s treeSortable) SetSortFunc(s TreeSortable) {
	var arg0 *C.GtkTreeSortable

	arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(s.Native()))

	C.gtk_tree_sortable_set_sort_func(arg0, arg1, arg2, arg3, arg4)
}

// SortColumnChanged emits a TreeSortable::sort-column-changed signal on
// @sortable.
func (s treeSortable) SortColumnChanged(s TreeSortable) {
	var arg0 *C.GtkTreeSortable

	arg0 = (*C.GtkTreeSortable)(unsafe.Pointer(s.Native()))

	C.gtk_tree_sortable_sort_column_changed(arg0)
}