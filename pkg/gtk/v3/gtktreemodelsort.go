// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
		{T: externglib.Type(C.gtk_tree_model_sort_get_type()), F: marshalTreeModelSort},
	})
}

// TreeModelSort: the TreeModelSort is a model which implements the TreeSortable
// interface. It does not hold any data itself, but rather is created with a
// child model and proxies its data. It has identical column types to this child
// model, and the changes in the child are propagated. The primary purpose of
// this model is to provide a way to sort a different model without modifying
// it. Note that the sort function used by TreeModelSort is not guaranteed to be
// stable.
//
// The use of this is best demonstrated through an example. In the following
// sample code we create two TreeView widgets each with a view of the same data.
// As the model is wrapped here by a TreeModelSort, the two TreeViews can each
// sort their view of the data without affecting the other. By contrast, if we
// simply put the same model in each widget, then sorting the first would sort
// the second.
//
// Using a TreeModelSort
//
//    void
//    selection_changed (GtkTreeSelection *selection, gpointer data)
//    {
//      GtkTreeModel *sort_model = NULL;
//      GtkTreeModel *child_model;
//      GtkTreeIter sort_iter;
//      GtkTreeIter child_iter;
//      char *some_data = NULL;
//      char *modified_data;
//
//      // Get the current selected row and the model.
//      if (! gtk_tree_selection_get_selected (selection,
//                                             &sort_model,
//                                             &sort_iter))
//        return;
//
//      // Look up the current value on the selected row and get
//      // a new value to change it to.
//      gtk_tree_model_get (GTK_TREE_MODEL (sort_model), &sort_iter,
//                          COLUMN_1, &some_data,
//                          -1);
//
//      modified_data = change_the_data (some_data);
//      g_free (some_data);
//
//      // Get an iterator on the child model, instead of the sort model.
//      gtk_tree_model_sort_convert_iter_to_child_iter (GTK_TREE_MODEL_SORT (sort_model),
//                                                      &child_iter,
//                                                      &sort_iter);
//
//      // Get the child model and change the value of the row. In this
//      // example, the child model is a GtkListStore. It could be any other
//      // type of model, though.
//      child_model = gtk_tree_model_sort_get_model (GTK_TREE_MODEL_SORT (sort_model));
//      gtk_list_store_set (GTK_LIST_STORE (child_model), &child_iter,
//                          COLUMN_1, &modified_data,
//                          -1);
//      g_free (modified_data);
//    }
type TreeModelSort interface {
	gextras.Objector
	TreeDragSource
	TreeModel
	TreeSortable

	// ClearCache: this function should almost never be called. It clears the
	// @tree_model_sort of any cached iterators that haven’t been reffed with
	// gtk_tree_model_ref_node(). This might be useful if the child model being
	// sorted is static (and doesn’t change often) and there has been a lot of
	// unreffed access to nodes. As a side effect of this function, all unreffed
	// iters will be invalid.
	ClearCache(t TreeModelSort)
	// ConvertChildIterToIter sets @sort_iter to point to the row in
	// @tree_model_sort that corresponds to the row pointed at by @child_iter.
	// If @sort_iter was not set, false is returned. Note: a boolean is only
	// returned since 2.14.
	ConvertChildIterToIter(t TreeModelSort, childIter *TreeIter) (sortIter *TreeIter, ok bool)
	// ConvertChildPathToPath converts @child_path to a path relative to
	// @tree_model_sort. That is, @child_path points to a path in the child
	// model. The returned path will point to the same row in the sorted model.
	// If @child_path isn’t a valid path on the child model, then nil is
	// returned.
	ConvertChildPathToPath(t TreeModelSort, childPath *TreePath)
	// ConvertIterToChildIter sets @child_iter to point to the row pointed to by
	// @sorted_iter.
	ConvertIterToChildIter(t TreeModelSort, sortedIter *TreeIter) *TreeIter
	// ConvertPathToChildPath converts @sorted_path to a path on the child model
	// of @tree_model_sort. That is, @sorted_path points to a location in
	// @tree_model_sort. The returned path will point to the same location in
	// the model not being sorted. If @sorted_path does not point to a location
	// in the child model, nil is returned.
	ConvertPathToChildPath(t TreeModelSort, sortedPath *TreePath)
	// Model returns the model the TreeModelSort is sorting.
	Model(t TreeModelSort)
	// IterIsValid: > This function is slow. Only use it for debugging and/or
	// testing > purposes.
	//
	// Checks if the given iter is a valid iter for this TreeModelSort.
	IterIsValid(t TreeModelSort, iter *TreeIter) bool
	// ResetDefaultSortFunc: this resets the default sort function to be in the
	// “unsorted” state. That is, it is in the same order as the child model. It
	// will re-sort the model to be in the same order as the child model only if
	// the TreeModelSort is in “unsorted” state.
	ResetDefaultSortFunc(t TreeModelSort)
}

// treeModelSort implements the TreeModelSort interface.
type treeModelSort struct {
	gextras.Objector
	TreeDragSource
	TreeModel
	TreeSortable
}

var _ TreeModelSort = (*treeModelSort)(nil)

// WrapTreeModelSort wraps a GObject to the right type. It is
// primarily used internally.
func WrapTreeModelSort(obj *externglib.Object) TreeModelSort {
	return TreeModelSort{
		Objector:       obj,
		TreeDragSource: WrapTreeDragSource(obj),
		TreeModel:      WrapTreeModel(obj),
		TreeSortable:   WrapTreeSortable(obj),
	}
}

func marshalTreeModelSort(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeModelSort(obj), nil
}

// NewTreeModelSortWithModel constructs a class TreeModelSort.
func NewTreeModelSortWithModel(childModel TreeModel) {
	var arg1 *C.GtkTreeModel

	arg1 = (*C.GtkTreeModel)(unsafe.Pointer(childModel.Native()))

	C.gtk_tree_model_sort_new_with_model(arg1)
}

// ClearCache: this function should almost never be called. It clears the
// @tree_model_sort of any cached iterators that haven’t been reffed with
// gtk_tree_model_ref_node(). This might be useful if the child model being
// sorted is static (and doesn’t change often) and there has been a lot of
// unreffed access to nodes. As a side effect of this function, all unreffed
// iters will be invalid.
func (t treeModelSort) ClearCache(t TreeModelSort) {
	var arg0 *C.GtkTreeModelSort

	arg0 = (*C.GtkTreeModelSort)(unsafe.Pointer(t.Native()))

	C.gtk_tree_model_sort_clear_cache(arg0)
}

// ConvertChildIterToIter sets @sort_iter to point to the row in
// @tree_model_sort that corresponds to the row pointed at by @child_iter.
// If @sort_iter was not set, false is returned. Note: a boolean is only
// returned since 2.14.
func (t treeModelSort) ConvertChildIterToIter(t TreeModelSort, childIter *TreeIter) (sortIter *TreeIter, ok bool) {
	var arg0 *C.GtkTreeModelSort
	var arg2 *C.GtkTreeIter

	arg0 = (*C.GtkTreeModelSort)(unsafe.Pointer(t.Native()))
	arg2 = (*C.GtkTreeIter)(unsafe.Pointer(childIter.Native()))

	var arg1 C.GtkTreeIter
	var sortIter *TreeIter
	var cret C.gboolean
	var ok bool

	cret = C.gtk_tree_model_sort_convert_child_iter_to_iter(arg0, &arg1, arg2)

	sortIter = WrapTreeIter(unsafe.Pointer(&arg1))
	if cret {
		ok = true
	}

	return sortIter, ok
}

// ConvertChildPathToPath converts @child_path to a path relative to
// @tree_model_sort. That is, @child_path points to a path in the child
// model. The returned path will point to the same row in the sorted model.
// If @child_path isn’t a valid path on the child model, then nil is
// returned.
func (t treeModelSort) ConvertChildPathToPath(t TreeModelSort, childPath *TreePath) {
	var arg0 *C.GtkTreeModelSort
	var arg1 *C.GtkTreePath

	arg0 = (*C.GtkTreeModelSort)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkTreePath)(unsafe.Pointer(childPath.Native()))

	C.gtk_tree_model_sort_convert_child_path_to_path(arg0, arg1)
}

// ConvertIterToChildIter sets @child_iter to point to the row pointed to by
// @sorted_iter.
func (t treeModelSort) ConvertIterToChildIter(t TreeModelSort, sortedIter *TreeIter) *TreeIter {
	var arg0 *C.GtkTreeModelSort
	var arg2 *C.GtkTreeIter

	arg0 = (*C.GtkTreeModelSort)(unsafe.Pointer(t.Native()))
	arg2 = (*C.GtkTreeIter)(unsafe.Pointer(sortedIter.Native()))

	var arg1 C.GtkTreeIter
	var childIter *TreeIter

	C.gtk_tree_model_sort_convert_iter_to_child_iter(arg0, &arg1, arg2)

	childIter = WrapTreeIter(unsafe.Pointer(&arg1))

	return childIter
}

// ConvertPathToChildPath converts @sorted_path to a path on the child model
// of @tree_model_sort. That is, @sorted_path points to a location in
// @tree_model_sort. The returned path will point to the same location in
// the model not being sorted. If @sorted_path does not point to a location
// in the child model, nil is returned.
func (t treeModelSort) ConvertPathToChildPath(t TreeModelSort, sortedPath *TreePath) {
	var arg0 *C.GtkTreeModelSort
	var arg1 *C.GtkTreePath

	arg0 = (*C.GtkTreeModelSort)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkTreePath)(unsafe.Pointer(sortedPath.Native()))

	C.gtk_tree_model_sort_convert_path_to_child_path(arg0, arg1)
}

// Model returns the model the TreeModelSort is sorting.
func (t treeModelSort) Model(t TreeModelSort) {
	var arg0 *C.GtkTreeModelSort

	arg0 = (*C.GtkTreeModelSort)(unsafe.Pointer(t.Native()))

	C.gtk_tree_model_sort_get_model(arg0)
}

// IterIsValid: > This function is slow. Only use it for debugging and/or
// testing > purposes.
//
// Checks if the given iter is a valid iter for this TreeModelSort.
func (t treeModelSort) IterIsValid(t TreeModelSort, iter *TreeIter) bool {
	var arg0 *C.GtkTreeModelSort
	var arg1 *C.GtkTreeIter

	arg0 = (*C.GtkTreeModelSort)(unsafe.Pointer(t.Native()))
	arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_tree_model_sort_iter_is_valid(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// ResetDefaultSortFunc: this resets the default sort function to be in the
// “unsorted” state. That is, it is in the same order as the child model. It
// will re-sort the model to be in the same order as the child model only if
// the TreeModelSort is in “unsorted” state.
func (t treeModelSort) ResetDefaultSortFunc(t TreeModelSort) {
	var arg0 *C.GtkTreeModelSort

	arg0 = (*C.GtkTreeModelSort)(unsafe.Pointer(t.Native()))

	C.gtk_tree_model_sort_reset_default_sort_func(arg0)
}
