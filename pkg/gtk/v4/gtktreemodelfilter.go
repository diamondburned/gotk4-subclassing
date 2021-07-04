// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_model_filter_get_type()), F: marshalTreeModelFilter},
	})
}

// TreeModelFilter: a GtkTreeModel which hides parts of an underlying tree model
//
// A TreeModelFilter is a tree model which wraps another tree model, and can do
// the following things:
//
// - Filter specific rows, based on data from a “visible column”, a column
// storing booleans indicating whether the row should be filtered or not, or
// based on the return value of a “visible function”, which gets a model, iter
// and user_data and returns a boolean indicating whether the row should be
// filtered or not.
//
// - Modify the “appearance” of the model, using a modify function. This is
// extremely powerful and allows for just changing some values and also for
// creating a completely different model based on the given child model.
//
// - Set a different root node, also known as a “virtual root”. You can pass in
// a TreePath indicating the root node for the filter at construction time.
//
// The basic API is similar to TreeModelSort. For an example on its usage, see
// the section on TreeModelSort.
//
// When using TreeModelFilter, it is important to realize that TreeModelFilter
// maintains an internal cache of all nodes which are visible in its clients.
// The cache is likely to be a subtree of the tree exposed by the child model.
// TreeModelFilter will not cache the entire child model when unnecessary to not
// compromise the caching mechanism that is exposed by the reference counting
// scheme. If the child model implements reference counting, unnecessary signals
// may not be emitted because of reference counting rule 3, see the TreeModel
// documentation. (Note that e.g. TreeStore does not implement reference
// counting and will always emit all signals, even when the receiving node is
// not visible).
//
// Because of this, limitations for possible visible functions do apply. In
// general, visible functions should only use data or properties from the node
// for which the visibility state must be determined, its siblings or its
// parents. Usually, having a dependency on the state of any child node is not
// possible, unless references are taken on these explicitly. When no such
// reference exists, no signals may be received for these child nodes (see
// reference counting rule number 3 in the TreeModel section).
//
// Determining the visibility state of a given node based on the state of its
// child nodes is a frequently occurring use case. Therefore, TreeModelFilter
// explicitly supports this. For example, when a node does not have any
// children, you might not want the node to be visible. As soon as the first row
// is added to the node’s child level (or the last row removed), the node’s
// visibility should be updated.
//
// This introduces a dependency from the node on its child nodes. In order to
// accommodate this, TreeModelFilter must make sure the necessary signals are
// received from the child model. This is achieved by building, for all nodes
// which are exposed as visible nodes to TreeModelFilter's clients, the child
// level (if any) and take a reference on the first node in this level.
// Furthermore, for every row-inserted, row-changed or row-deleted signal (also
// these which were not handled because the node was not cached),
// TreeModelFilter will check if the visibility state of any parent node has
// changed.
//
// Beware, however, that this explicit support is limited to these two cases.
// For example, if you want a node to be visible only if two nodes in a child’s
// child level (2 levels deeper) are visible, you are on your own. In this case,
// either rely on TreeStore to emit all signals because it does not implement
// reference counting, or for models that do implement reference counting,
// obtain references on these child levels yourself.
type TreeModelFilter interface {
	TreeDragSource
	TreeModel

	// ClearCacheTreeModelFilter:
	ClearCacheTreeModelFilter()
	// ConvertChildIterToIterTreeModelFilter:
	ConvertChildIterToIterTreeModelFilter(childIter *TreeIter) (TreeIter, bool)
	// ConvertChildPathToPathTreeModelFilter:
	ConvertChildPathToPathTreeModelFilter(childPath *TreePath) *TreePath
	// ConvertIterToChildIterTreeModelFilter:
	ConvertIterToChildIterTreeModelFilter(filterIter *TreeIter) TreeIter
	// ConvertPathToChildPathTreeModelFilter:
	ConvertPathToChildPathTreeModelFilter(filterPath *TreePath) *TreePath
	// Model:
	Model() TreeModel
	// RefilterTreeModelFilter:
	RefilterTreeModelFilter()
	// SetVisibleColumnTreeModelFilter:
	SetVisibleColumnTreeModelFilter(column int)
}

// treeModelFilter implements the TreeModelFilter class.
type treeModelFilter struct {
	gextras.Objector
}

// WrapTreeModelFilter wraps a GObject to the right type. It is
// primarily used internally.
func WrapTreeModelFilter(obj *externglib.Object) TreeModelFilter {
	return treeModelFilter{
		Objector: obj,
	}
}

func marshalTreeModelFilter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeModelFilter(obj), nil
}

func (f treeModelFilter) ClearCacheTreeModelFilter() {
	var _arg0 *C.GtkTreeModelFilter // out

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(f.Native()))

	C.gtk_tree_model_filter_clear_cache(_arg0)
}

func (f treeModelFilter) ConvertChildIterToIterTreeModelFilter(childIter *TreeIter) (TreeIter, bool) {
	var _arg0 *C.GtkTreeModelFilter // out
	var _arg1 C.GtkTreeIter         // in
	var _arg2 *C.GtkTreeIter        // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(f.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(childIter.Native()))

	_cret = C.gtk_tree_model_filter_convert_child_iter_to_iter(_arg0, &_arg1, _arg2)

	var _filterIter TreeIter // out
	var _ok bool             // out

	{
		var refTmpIn *C.GtkTreeIter
		var refTmpOut *TreeIter

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*TreeIter)(unsafe.Pointer(refTmpIn))

		_filterIter = *refTmpOut
	}
	if _cret != 0 {
		_ok = true
	}

	return _filterIter, _ok
}

func (f treeModelFilter) ConvertChildPathToPathTreeModelFilter(childPath *TreePath) *TreePath {
	var _arg0 *C.GtkTreeModelFilter // out
	var _arg1 *C.GtkTreePath        // out
	var _cret *C.GtkTreePath        // in

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(childPath.Native()))

	_cret = C.gtk_tree_model_filter_convert_child_path_to_path(_arg0, _arg1)

	var _treePath *TreePath // out

	_treePath = (*TreePath)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_treePath, func(v **TreePath) {
		C.free(unsafe.Pointer(v))
	})

	return _treePath
}

func (f treeModelFilter) ConvertIterToChildIterTreeModelFilter(filterIter *TreeIter) TreeIter {
	var _arg0 *C.GtkTreeModelFilter // out
	var _arg1 C.GtkTreeIter         // in
	var _arg2 *C.GtkTreeIter        // out

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(f.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(filterIter.Native()))

	C.gtk_tree_model_filter_convert_iter_to_child_iter(_arg0, &_arg1, _arg2)

	var _childIter TreeIter // out

	{
		var refTmpIn *C.GtkTreeIter
		var refTmpOut *TreeIter

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*TreeIter)(unsafe.Pointer(refTmpIn))

		_childIter = *refTmpOut
	}

	return _childIter
}

func (f treeModelFilter) ConvertPathToChildPathTreeModelFilter(filterPath *TreePath) *TreePath {
	var _arg0 *C.GtkTreeModelFilter // out
	var _arg1 *C.GtkTreePath        // out
	var _cret *C.GtkTreePath        // in

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.GtkTreePath)(unsafe.Pointer(filterPath.Native()))

	_cret = C.gtk_tree_model_filter_convert_path_to_child_path(_arg0, _arg1)

	var _treePath *TreePath // out

	_treePath = (*TreePath)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_treePath, func(v **TreePath) {
		C.free(unsafe.Pointer(v))
	})

	return _treePath
}

func (f treeModelFilter) Model() TreeModel {
	var _arg0 *C.GtkTreeModelFilter // out
	var _cret *C.GtkTreeModel       // in

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(f.Native()))

	_cret = C.gtk_tree_model_filter_get_model(_arg0)

	var _treeModel TreeModel // out

	_treeModel = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(TreeModel)

	return _treeModel
}

func (f treeModelFilter) RefilterTreeModelFilter() {
	var _arg0 *C.GtkTreeModelFilter // out

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(f.Native()))

	C.gtk_tree_model_filter_refilter(_arg0)
}

func (f treeModelFilter) SetVisibleColumnTreeModelFilter(column int) {
	var _arg0 *C.GtkTreeModelFilter // out
	var _arg1 C.int                 // out

	_arg0 = (*C.GtkTreeModelFilter)(unsafe.Pointer(f.Native()))
	_arg1 = C.int(column)

	C.gtk_tree_model_filter_set_visible_column(_arg0, _arg1)
}

func (d treeModelFilter) DragDataDelete(path *TreePath) bool {
	return WrapTreeDragSource(gextras.InternObject(d)).DragDataDelete(path)
}

func (d treeModelFilter) DragDataGet(path *TreePath) gdk.ContentProvider {
	return WrapTreeDragSource(gextras.InternObject(d)).DragDataGet(path)
}

func (d treeModelFilter) RowDraggable(path *TreePath) bool {
	return WrapTreeDragSource(gextras.InternObject(d)).RowDraggable(path)
}

func (t treeModelFilter) NewFilter(root *TreePath) TreeModel {
	return WrapTreeModel(gextras.InternObject(t)).NewFilter(root)
}

func (t treeModelFilter) ColumnType(index_ int) externglib.Type {
	return WrapTreeModel(gextras.InternObject(t)).ColumnType(index_)
}

func (t treeModelFilter) Flags() TreeModelFlags {
	return WrapTreeModel(gextras.InternObject(t)).Flags()
}

func (t treeModelFilter) Iter(path *TreePath) (TreeIter, bool) {
	return WrapTreeModel(gextras.InternObject(t)).Iter(path)
}

func (t treeModelFilter) IterFirst() (TreeIter, bool) {
	return WrapTreeModel(gextras.InternObject(t)).IterFirst()
}

func (t treeModelFilter) IterFromString(pathString string) (TreeIter, bool) {
	return WrapTreeModel(gextras.InternObject(t)).IterFromString(pathString)
}

func (t treeModelFilter) NColumns() int {
	return WrapTreeModel(gextras.InternObject(t)).NColumns()
}

func (t treeModelFilter) Path(iter *TreeIter) *TreePath {
	return WrapTreeModel(gextras.InternObject(t)).Path(iter)
}

func (t treeModelFilter) StringFromIter(iter *TreeIter) string {
	return WrapTreeModel(gextras.InternObject(t)).StringFromIter(iter)
}

func (t treeModelFilter) Value(iter *TreeIter, column int) externglib.Value {
	return WrapTreeModel(gextras.InternObject(t)).Value(iter, column)
}

func (t treeModelFilter) IterChildren(parent *TreeIter) (TreeIter, bool) {
	return WrapTreeModel(gextras.InternObject(t)).IterChildren(parent)
}

func (t treeModelFilter) IterHasChild(iter *TreeIter) bool {
	return WrapTreeModel(gextras.InternObject(t)).IterHasChild(iter)
}

func (t treeModelFilter) IterNChildren(iter *TreeIter) int {
	return WrapTreeModel(gextras.InternObject(t)).IterNChildren(iter)
}

func (t treeModelFilter) IterNext(iter *TreeIter) bool {
	return WrapTreeModel(gextras.InternObject(t)).IterNext(iter)
}

func (t treeModelFilter) IterNthChild(parent *TreeIter, n int) (TreeIter, bool) {
	return WrapTreeModel(gextras.InternObject(t)).IterNthChild(parent, n)
}

func (t treeModelFilter) IterParent(child *TreeIter) (TreeIter, bool) {
	return WrapTreeModel(gextras.InternObject(t)).IterParent(child)
}

func (t treeModelFilter) IterPrevious(iter *TreeIter) bool {
	return WrapTreeModel(gextras.InternObject(t)).IterPrevious(iter)
}

func (t treeModelFilter) RefNode(iter *TreeIter) {
	WrapTreeModel(gextras.InternObject(t)).RefNode(iter)
}

func (t treeModelFilter) RowChanged(path *TreePath, iter *TreeIter) {
	WrapTreeModel(gextras.InternObject(t)).RowChanged(path, iter)
}

func (t treeModelFilter) RowDeleted(path *TreePath) {
	WrapTreeModel(gextras.InternObject(t)).RowDeleted(path)
}

func (t treeModelFilter) RowHasChildToggled(path *TreePath, iter *TreeIter) {
	WrapTreeModel(gextras.InternObject(t)).RowHasChildToggled(path, iter)
}

func (t treeModelFilter) RowInserted(path *TreePath, iter *TreeIter) {
	WrapTreeModel(gextras.InternObject(t)).RowInserted(path, iter)
}

func (t treeModelFilter) RowsReorderedWithLength(path *TreePath, iter *TreeIter, newOrder []int) {
	WrapTreeModel(gextras.InternObject(t)).RowsReorderedWithLength(path, iter, newOrder)
}

func (t treeModelFilter) UnrefNode(iter *TreeIter) {
	WrapTreeModel(gextras.InternObject(t)).UnrefNode(iter)
}