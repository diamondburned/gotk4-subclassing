// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern gboolean _gotk4_gtk3_TreeDragDestIface_drag_data_received(GtkTreeDragDest*, GtkTreePath*, GtkSelectionData*);
// extern gboolean _gotk4_gtk3_TreeDragDestIface_row_drop_possible(GtkTreeDragDest*, GtkTreePath*, GtkSelectionData*);
// extern gboolean _gotk4_gtk3_TreeDragSourceIface_drag_data_delete(GtkTreeDragSource*, GtkTreePath*);
// extern gboolean _gotk4_gtk3_TreeDragSourceIface_drag_data_get(GtkTreeDragSource*, GtkTreePath*, GtkSelectionData*);
// extern gboolean _gotk4_gtk3_TreeDragSourceIface_row_draggable(GtkTreeDragSource*, GtkTreePath*);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_drag_dest_get_type()), F: marshalTreeDragDester},
		{T: externglib.Type(C.gtk_tree_drag_source_get_type()), F: marshalTreeDragSourcer},
	})
}

// TreeSetRowDragData sets selection data of target type GTK_TREE_MODEL_ROW.
// Normally used in a drag_data_get handler.
//
// The function takes the following parameters:
//
//    - selectionData: some SelectionData.
//    - treeModel: TreeModel.
//    - path: row in tree_model.
//
// The function returns the following values:
//
//    - ok: TRUE if the SelectionData had the proper target type to allow us to
//      set a tree row.
//
func TreeSetRowDragData(selectionData *SelectionData, treeModel TreeModeller, path *TreePath) bool {
	var _arg1 *C.GtkSelectionData // out
	var _arg2 *C.GtkTreeModel     // out
	var _arg3 *C.GtkTreePath      // out
	var _cret C.gboolean          // in

	_arg1 = (*C.GtkSelectionData)(gextras.StructNative(unsafe.Pointer(selectionData)))
	_arg2 = (*C.GtkTreeModel)(unsafe.Pointer(treeModel.Native()))
	_arg3 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))

	_cret = C.gtk_tree_set_row_drag_data(_arg1, _arg2, _arg3)
	runtime.KeepAlive(selectionData)
	runtime.KeepAlive(treeModel)
	runtime.KeepAlive(path)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TreeDragDestOverrider contains methods that are overridable.
type TreeDragDestOverrider interface {
	// DragDataReceived asks the TreeDragDest to insert a row before the path
	// dest, deriving the contents of the row from selection_data. If dest is
	// outside the tree so that inserting before it is impossible, FALSE will be
	// returned. Also, FALSE may be returned if the new row is not created for
	// some model-specific reason. Should robustly handle a dest no longer found
	// in the model!.
	//
	// The function takes the following parameters:
	//
	//    - dest: row to drop in front of.
	//    - selectionData: data to drop.
	//
	// The function returns the following values:
	//
	//    - ok: whether a new row was created before position dest.
	//
	DragDataReceived(dest *TreePath, selectionData *SelectionData) bool
	// RowDropPossible determines whether a drop is possible before the given
	// dest_path, at the same depth as dest_path. i.e., can we drop the data in
	// selection_data at that location. dest_path does not have to exist; the
	// return value will almost certainly be FALSE if the parent of dest_path
	// doesn???t exist, though.
	//
	// The function takes the following parameters:
	//
	//    - destPath: destination row.
	//    - selectionData: data being dragged.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if a drop is possible before dest_path.
	//
	RowDropPossible(destPath *TreePath, selectionData *SelectionData) bool
}

type TreeDragDest struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*TreeDragDest)(nil)
)

// TreeDragDester describes TreeDragDest's interface methods.
type TreeDragDester interface {
	externglib.Objector

	// DragDataReceived asks the TreeDragDest to insert a row before the path
	// dest, deriving the contents of the row from selection_data.
	DragDataReceived(dest *TreePath, selectionData *SelectionData) bool
	// RowDropPossible determines whether a drop is possible before the given
	// dest_path, at the same depth as dest_path.
	RowDropPossible(destPath *TreePath, selectionData *SelectionData) bool
}

var _ TreeDragDester = (*TreeDragDest)(nil)

func ifaceInitTreeDragDester(gifacePtr, data C.gpointer) {
	iface := (*C.GtkTreeDragDestIface)(unsafe.Pointer(gifacePtr))
	iface.drag_data_received = (*[0]byte)(C._gotk4_gtk3_TreeDragDestIface_drag_data_received)
	iface.row_drop_possible = (*[0]byte)(C._gotk4_gtk3_TreeDragDestIface_row_drop_possible)
}

//export _gotk4_gtk3_TreeDragDestIface_drag_data_received
func _gotk4_gtk3_TreeDragDestIface_drag_data_received(arg0 *C.GtkTreeDragDest, arg1 *C.GtkTreePath, arg2 *C.GtkSelectionData) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TreeDragDestOverrider)

	var _dest *TreePath               // out
	var _selectionData *SelectionData // out

	_dest = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_selectionData = (*SelectionData)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := iface.DragDataReceived(_dest, _selectionData)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_TreeDragDestIface_row_drop_possible
func _gotk4_gtk3_TreeDragDestIface_row_drop_possible(arg0 *C.GtkTreeDragDest, arg1 *C.GtkTreePath, arg2 *C.GtkSelectionData) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TreeDragDestOverrider)

	var _destPath *TreePath           // out
	var _selectionData *SelectionData // out

	_destPath = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_selectionData = (*SelectionData)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := iface.RowDropPossible(_destPath, _selectionData)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapTreeDragDest(obj *externglib.Object) *TreeDragDest {
	return &TreeDragDest{
		Object: obj,
	}
}

func marshalTreeDragDester(p uintptr) (interface{}, error) {
	return wrapTreeDragDest(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DragDataReceived asks the TreeDragDest to insert a row before the path dest,
// deriving the contents of the row from selection_data. If dest is outside the
// tree so that inserting before it is impossible, FALSE will be returned. Also,
// FALSE may be returned if the new row is not created for some model-specific
// reason. Should robustly handle a dest no longer found in the model!.
//
// The function takes the following parameters:
//
//    - dest: row to drop in front of.
//    - selectionData: data to drop.
//
// The function returns the following values:
//
//    - ok: whether a new row was created before position dest.
//
func (dragDest *TreeDragDest) DragDataReceived(dest *TreePath, selectionData *SelectionData) bool {
	var _arg0 *C.GtkTreeDragDest  // out
	var _arg1 *C.GtkTreePath      // out
	var _arg2 *C.GtkSelectionData // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkTreeDragDest)(unsafe.Pointer(dragDest.Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(dest)))
	_arg2 = (*C.GtkSelectionData)(gextras.StructNative(unsafe.Pointer(selectionData)))

	_cret = C.gtk_tree_drag_dest_drag_data_received(_arg0, _arg1, _arg2)
	runtime.KeepAlive(dragDest)
	runtime.KeepAlive(dest)
	runtime.KeepAlive(selectionData)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RowDropPossible determines whether a drop is possible before the given
// dest_path, at the same depth as dest_path. i.e., can we drop the data in
// selection_data at that location. dest_path does not have to exist; the return
// value will almost certainly be FALSE if the parent of dest_path doesn???t
// exist, though.
//
// The function takes the following parameters:
//
//    - destPath: destination row.
//    - selectionData: data being dragged.
//
// The function returns the following values:
//
//    - ok: TRUE if a drop is possible before dest_path.
//
func (dragDest *TreeDragDest) RowDropPossible(destPath *TreePath, selectionData *SelectionData) bool {
	var _arg0 *C.GtkTreeDragDest  // out
	var _arg1 *C.GtkTreePath      // out
	var _arg2 *C.GtkSelectionData // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkTreeDragDest)(unsafe.Pointer(dragDest.Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(destPath)))
	_arg2 = (*C.GtkSelectionData)(gextras.StructNative(unsafe.Pointer(selectionData)))

	_cret = C.gtk_tree_drag_dest_row_drop_possible(_arg0, _arg1, _arg2)
	runtime.KeepAlive(dragDest)
	runtime.KeepAlive(destPath)
	runtime.KeepAlive(selectionData)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TreeDragSourceOverrider contains methods that are overridable.
type TreeDragSourceOverrider interface {
	// DragDataDelete asks the TreeDragSource to delete the row at path, because
	// it was moved somewhere else via drag-and-drop. Returns FALSE if the
	// deletion fails because path no longer exists, or for some model-specific
	// reason. Should robustly handle a path no longer found in the model!.
	//
	// The function takes the following parameters:
	//
	//    - path: row that was being dragged.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if the row was successfully deleted.
	//
	DragDataDelete(path *TreePath) bool
	// DragDataGet asks the TreeDragSource to fill in selection_data with a
	// representation of the row at path. selection_data->target gives the
	// required type of the data. Should robustly handle a path no longer found
	// in the model!.
	//
	// The function takes the following parameters:
	//
	//    - path: row that was dragged.
	//    - selectionData to fill with data from the dragged row.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if data of the required type was provided.
	//
	DragDataGet(path *TreePath, selectionData *SelectionData) bool
	// RowDraggable asks the TreeDragSource whether a particular row can be used
	// as the source of a DND operation. If the source doesn???t implement this
	// interface, the row is assumed draggable.
	//
	// The function takes the following parameters:
	//
	//    - path: row on which user is initiating a drag.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if the row can be dragged.
	//
	RowDraggable(path *TreePath) bool
}

type TreeDragSource struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*TreeDragSource)(nil)
)

// TreeDragSourcer describes TreeDragSource's interface methods.
type TreeDragSourcer interface {
	externglib.Objector

	// DragDataDelete asks the TreeDragSource to delete the row at path, because
	// it was moved somewhere else via drag-and-drop.
	DragDataDelete(path *TreePath) bool
	// DragDataGet asks the TreeDragSource to fill in selection_data with a
	// representation of the row at path.
	DragDataGet(path *TreePath, selectionData *SelectionData) bool
	// RowDraggable asks the TreeDragSource whether a particular row can be used
	// as the source of a DND operation.
	RowDraggable(path *TreePath) bool
}

var _ TreeDragSourcer = (*TreeDragSource)(nil)

func ifaceInitTreeDragSourcer(gifacePtr, data C.gpointer) {
	iface := (*C.GtkTreeDragSourceIface)(unsafe.Pointer(gifacePtr))
	iface.drag_data_delete = (*[0]byte)(C._gotk4_gtk3_TreeDragSourceIface_drag_data_delete)
	iface.drag_data_get = (*[0]byte)(C._gotk4_gtk3_TreeDragSourceIface_drag_data_get)
	iface.row_draggable = (*[0]byte)(C._gotk4_gtk3_TreeDragSourceIface_row_draggable)
}

//export _gotk4_gtk3_TreeDragSourceIface_drag_data_delete
func _gotk4_gtk3_TreeDragSourceIface_drag_data_delete(arg0 *C.GtkTreeDragSource, arg1 *C.GtkTreePath) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TreeDragSourceOverrider)

	var _path *TreePath // out

	_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := iface.DragDataDelete(_path)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_TreeDragSourceIface_drag_data_get
func _gotk4_gtk3_TreeDragSourceIface_drag_data_get(arg0 *C.GtkTreeDragSource, arg1 *C.GtkTreePath, arg2 *C.GtkSelectionData) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TreeDragSourceOverrider)

	var _path *TreePath               // out
	var _selectionData *SelectionData // out

	_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_selectionData = (*SelectionData)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := iface.DragDataGet(_path, _selectionData)

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gtk3_TreeDragSourceIface_row_draggable
func _gotk4_gtk3_TreeDragSourceIface_row_draggable(arg0 *C.GtkTreeDragSource, arg1 *C.GtkTreePath) (cret C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(TreeDragSourceOverrider)

	var _path *TreePath // out

	_path = (*TreePath)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := iface.RowDraggable(_path)

	if ok {
		cret = C.TRUE
	}

	return cret
}

func wrapTreeDragSource(obj *externglib.Object) *TreeDragSource {
	return &TreeDragSource{
		Object: obj,
	}
}

func marshalTreeDragSourcer(p uintptr) (interface{}, error) {
	return wrapTreeDragSource(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DragDataDelete asks the TreeDragSource to delete the row at path, because it
// was moved somewhere else via drag-and-drop. Returns FALSE if the deletion
// fails because path no longer exists, or for some model-specific reason.
// Should robustly handle a path no longer found in the model!.
//
// The function takes the following parameters:
//
//    - path: row that was being dragged.
//
// The function returns the following values:
//
//    - ok: TRUE if the row was successfully deleted.
//
func (dragSource *TreeDragSource) DragDataDelete(path *TreePath) bool {
	var _arg0 *C.GtkTreeDragSource // out
	var _arg1 *C.GtkTreePath       // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(dragSource.Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))

	_cret = C.gtk_tree_drag_source_drag_data_delete(_arg0, _arg1)
	runtime.KeepAlive(dragSource)
	runtime.KeepAlive(path)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DragDataGet asks the TreeDragSource to fill in selection_data with a
// representation of the row at path. selection_data->target gives the required
// type of the data. Should robustly handle a path no longer found in the
// model!.
//
// The function takes the following parameters:
//
//    - path: row that was dragged.
//    - selectionData to fill with data from the dragged row.
//
// The function returns the following values:
//
//    - ok: TRUE if data of the required type was provided.
//
func (dragSource *TreeDragSource) DragDataGet(path *TreePath, selectionData *SelectionData) bool {
	var _arg0 *C.GtkTreeDragSource // out
	var _arg1 *C.GtkTreePath       // out
	var _arg2 *C.GtkSelectionData  // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(dragSource.Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))
	_arg2 = (*C.GtkSelectionData)(gextras.StructNative(unsafe.Pointer(selectionData)))

	_cret = C.gtk_tree_drag_source_drag_data_get(_arg0, _arg1, _arg2)
	runtime.KeepAlive(dragSource)
	runtime.KeepAlive(path)
	runtime.KeepAlive(selectionData)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RowDraggable asks the TreeDragSource whether a particular row can be used as
// the source of a DND operation. If the source doesn???t implement this
// interface, the row is assumed draggable.
//
// The function takes the following parameters:
//
//    - path: row on which user is initiating a drag.
//
// The function returns the following values:
//
//    - ok: TRUE if the row can be dragged.
//
func (dragSource *TreeDragSource) RowDraggable(path *TreePath) bool {
	var _arg0 *C.GtkTreeDragSource // out
	var _arg1 *C.GtkTreePath       // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkTreeDragSource)(unsafe.Pointer(dragSource.Native()))
	_arg1 = (*C.GtkTreePath)(gextras.StructNative(unsafe.Pointer(path)))

	_cret = C.gtk_tree_drag_source_row_draggable(_arg0, _arg1)
	runtime.KeepAlive(dragSource)
	runtime.KeepAlive(path)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
