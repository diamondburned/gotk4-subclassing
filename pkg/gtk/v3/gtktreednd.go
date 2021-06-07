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
		{T: externglib.Type(C.gtk_tree_drag_dest_get_type()), F: marshalTreeDragDest},
		{T: externglib.Type(C.gtk_tree_drag_source_get_type()), F: marshalTreeDragSource},
	})
}

// TreeGetRowDragData obtains a @tree_model and @path from selection data of
// target type GTK_TREE_MODEL_ROW. Normally called from a drag_data_received
// handler. This function can only be used if @selection_data originates from
// the same process that’s calling this function, because a pointer to the tree
// model is being passed around. If you aren’t in the same process, then you'll
// get memory corruption. In the TreeDragDest drag_data_received handler, you
// can assume that selection data of type GTK_TREE_MODEL_ROW is in from the
// current process. The returned path must be freed with gtk_tree_path_free().
func TreeGetRowDragData(selectionData *SelectionData) (treeModel *TreeModel, path **TreePath, ok bool) {
	var arg1 *C.GtkSelectionData

	arg1 = (*C.GtkSelectionData)(unsafe.Pointer(selectionData.Native()))

	var arg2 *C.GtkTreeModel
	var treeModel *TreeModel
	var arg3 *C.GtkTreePath
	var path **TreePath
	var cret C.gboolean
	var ok bool

	cret = C.gtk_tree_get_row_drag_data(arg1, &arg2, &arg3)

	treeModel = gextras.CastObject(externglib.Take(unsafe.Pointer(&arg2.Native()))).(*TreeModel)
	path = WrapTreePath(unsafe.Pointer(&arg3))
	runtime.SetFinalizer(path, func(v **TreePath) {
		C.free(unsafe.Pointer(v.Native()))
	})
	if cret {
		ok = true
	}

	return treeModel, path, ok
}

// TreeSetRowDragData sets selection data of target type GTK_TREE_MODEL_ROW.
// Normally used in a drag_data_get handler.
func TreeSetRowDragData(selectionData *SelectionData, treeModel TreeModel, path *TreePath) bool {
	var arg1 *C.GtkSelectionData
	var arg2 *C.GtkTreeModel
	var arg3 *C.GtkTreePath

	arg1 = (*C.GtkSelectionData)(unsafe.Pointer(selectionData.Native()))
	arg2 = (*C.GtkTreeModel)(unsafe.Pointer(treeModel.Native()))
	arg3 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_tree_set_row_drag_data(arg1, arg2, arg3)

	if cret {
		ok = true
	}

	return ok
}

// TreeDragDestOverrider contains methods that are overridable. This
// interface is a subset of the interface TreeDragDest.
type TreeDragDestOverrider interface {
	// DragDataReceived asks the TreeDragDest to insert a row before the path
	// @dest, deriving the contents of the row from @selection_data. If @dest is
	// outside the tree so that inserting before it is impossible, false will be
	// returned. Also, false may be returned if the new row is not created for
	// some model-specific reason. Should robustly handle a @dest no longer
	// found in the model!
	DragDataReceived(d TreeDragDest, dest *TreePath, selectionData *SelectionData) bool
	// RowDropPossible determines whether a drop is possible before the given
	// @dest_path, at the same depth as @dest_path. i.e., can we drop the data
	// in @selection_data at that location. @dest_path does not have to exist;
	// the return value will almost certainly be false if the parent of
	// @dest_path doesn’t exist, though.
	RowDropPossible(d TreeDragDest, destPath *TreePath, selectionData *SelectionData) bool
}

type TreeDragDest interface {
	gextras.Objector
	TreeDragDestOverrider
}

// treeDragDest implements the TreeDragDest interface.
type treeDragDest struct {
	gextras.Objector
}

var _ TreeDragDest = (*treeDragDest)(nil)

// WrapTreeDragDest wraps a GObject to a type that implements interface
// TreeDragDest. It is primarily used internally.
func WrapTreeDragDest(obj *externglib.Object) TreeDragDest {
	return TreeDragDest{
		Objector: obj,
	}
}

func marshalTreeDragDest(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeDragDest(obj), nil
}

// DragDataReceived asks the TreeDragDest to insert a row before the path
// @dest, deriving the contents of the row from @selection_data. If @dest is
// outside the tree so that inserting before it is impossible, false will be
// returned. Also, false may be returned if the new row is not created for
// some model-specific reason. Should robustly handle a @dest no longer
// found in the model!
func (d treeDragDest) DragDataReceived(d TreeDragDest, dest *TreePath, selectionData *SelectionData) bool {
	var arg0 *C.GtkTreeDragDest
	var arg1 *C.GtkTreePath
	var arg2 *C.GtkSelectionData

	arg0 = (*C.GtkTreeDragDest)(unsafe.Pointer(d.Native()))
	arg1 = (*C.GtkTreePath)(unsafe.Pointer(dest.Native()))
	arg2 = (*C.GtkSelectionData)(unsafe.Pointer(selectionData.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_tree_drag_dest_drag_data_received(arg0, arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// RowDropPossible determines whether a drop is possible before the given
// @dest_path, at the same depth as @dest_path. i.e., can we drop the data
// in @selection_data at that location. @dest_path does not have to exist;
// the return value will almost certainly be false if the parent of
// @dest_path doesn’t exist, though.
func (d treeDragDest) RowDropPossible(d TreeDragDest, destPath *TreePath, selectionData *SelectionData) bool {
	var arg0 *C.GtkTreeDragDest
	var arg1 *C.GtkTreePath
	var arg2 *C.GtkSelectionData

	arg0 = (*C.GtkTreeDragDest)(unsafe.Pointer(d.Native()))
	arg1 = (*C.GtkTreePath)(unsafe.Pointer(destPath.Native()))
	arg2 = (*C.GtkSelectionData)(unsafe.Pointer(selectionData.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_tree_drag_dest_row_drop_possible(arg0, arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// TreeDragSourceOverrider contains methods that are overridable. This
// interface is a subset of the interface TreeDragSource.
type TreeDragSourceOverrider interface {
	// DragDataDelete asks the TreeDragSource to delete the row at @path,
	// because it was moved somewhere else via drag-and-drop. Returns false if
	// the deletion fails because @path no longer exists, or for some
	// model-specific reason. Should robustly handle a @path no longer found in
	// the model!
	DragDataDelete(d TreeDragSource, path *TreePath) bool
	// DragDataGet asks the TreeDragSource to fill in @selection_data with a
	// representation of the row at @path. @selection_data->target gives the
	// required type of the data. Should robustly handle a @path no longer found
	// in the model!
	DragDataGet(d TreeDragSource, path *TreePath, selectionData *SelectionData) bool
	// RowDraggable asks the TreeDragSource whether a particular row can be used
	// as the source of a DND operation. If the source doesn’t implement this
	// interface, the row is assumed draggable.
	RowDraggable(d TreeDragSource, path *TreePath) bool
}

type TreeDragSource interface {
	gextras.Objector
	TreeDragSourceOverrider
}

// treeDragSource implements the TreeDragSource interface.
type treeDragSource struct {
	gextras.Objector
}

var _ TreeDragSource = (*treeDragSource)(nil)

// WrapTreeDragSource wraps a GObject to a type that implements interface
// TreeDragSource. It is primarily used internally.
func WrapTreeDragSource(obj *externglib.Object) TreeDragSource {
	return TreeDragSource{
		Objector: obj,
	}
}

func marshalTreeDragSource(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeDragSource(obj), nil
}

// DragDataDelete asks the TreeDragSource to delete the row at @path,
// because it was moved somewhere else via drag-and-drop. Returns false if
// the deletion fails because @path no longer exists, or for some
// model-specific reason. Should robustly handle a @path no longer found in
// the model!
func (d treeDragSource) DragDataDelete(d TreeDragSource, path *TreePath) bool {
	var arg0 *C.GtkTreeDragSource
	var arg1 *C.GtkTreePath

	arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(d.Native()))
	arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_tree_drag_source_drag_data_delete(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// DragDataGet asks the TreeDragSource to fill in @selection_data with a
// representation of the row at @path. @selection_data->target gives the
// required type of the data. Should robustly handle a @path no longer found
// in the model!
func (d treeDragSource) DragDataGet(d TreeDragSource, path *TreePath, selectionData *SelectionData) bool {
	var arg0 *C.GtkTreeDragSource
	var arg1 *C.GtkTreePath
	var arg2 *C.GtkSelectionData

	arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(d.Native()))
	arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))
	arg2 = (*C.GtkSelectionData)(unsafe.Pointer(selectionData.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_tree_drag_source_drag_data_get(arg0, arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// RowDraggable asks the TreeDragSource whether a particular row can be used
// as the source of a DND operation. If the source doesn’t implement this
// interface, the row is assumed draggable.
func (d treeDragSource) RowDraggable(d TreeDragSource, path *TreePath) bool {
	var arg0 *C.GtkTreeDragSource
	var arg1 *C.GtkTreePath

	arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(d.Native()))
	arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_tree_drag_source_row_draggable(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}
