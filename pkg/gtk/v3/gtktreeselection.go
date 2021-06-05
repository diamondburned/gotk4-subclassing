// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_selection_get_type()), F: marshalTreeSelection},
	})
}

// TreeSelectionForeachFunc: a function used by
// gtk_tree_selection_selected_foreach() to map all selected rows. It will be
// called on every selected row in the view.
type TreeSelectionForeachFunc func(model TreeModel, path *TreePath, iter *TreeIter)

//export gotk4_TreeSelectionForeachFunc
func gotk4_TreeSelectionForeachFunc(arg0 *C.GtkTreeModel, arg1 *C.GtkTreePath, arg2 *C.GtkTreeIter, arg3 C.gpointer) {
	v := box.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(TreeSelectionForeachFunc)
	fn(model, path, iter, data)
}

// TreeSelectionFunc: a function used by
// gtk_tree_selection_set_select_function() to filter whether or not a row may
// be selected. It is called whenever a row's state might change. A return value
// of true indicates to @selection that it is okay to change the selection.
type TreeSelectionFunc func(selection TreeSelection, model TreeModel, path *TreePath, pathCurrentlySelected bool) bool

//export gotk4_TreeSelectionFunc
func gotk4_TreeSelectionFunc(arg0 *C.GtkTreeSelection, arg1 *C.GtkTreeModel, arg2 *C.GtkTreePath, arg3 C.gboolean, arg4 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(TreeSelectionFunc)
	ret := fn(selection, model, path, pathCurrentlySelected, data)

	if ret {
		cret = C.gboolean(1)
	}

	return cret
}

// TreeSelection: the TreeSelection object is a helper object to manage the
// selection for a TreeView widget. The TreeSelection object is automatically
// created when a new TreeView widget is created, and cannot exist independently
// of this widget. The primary reason the TreeSelection objects exists is for
// cleanliness of code and API. That is, there is no conceptual reason all these
// functions could not be methods on the TreeView widget instead of a separate
// function.
//
// The TreeSelection object is gotten from a TreeView by calling
// gtk_tree_view_get_selection(). It can be manipulated to check the selection
// status of the tree, as well as select and deselect individual rows. Selection
// is done completely view side. As a result, multiple views of the same model
// can have completely different selections. Additionally, you cannot change the
// selection of a row on the model that is not currently displayed by the view
// without expanding its parents first.
//
// One of the important things to remember when monitoring the selection of a
// view is that the TreeSelection::changed signal is mostly a hint. That is, it
// may only emit one signal when a range of rows is selected. Additionally, it
// may on occasion emit a TreeSelection::changed signal when nothing has
// happened (mostly as a result of programmers calling select_row on an already
// selected row).
type TreeSelection interface {
	gextras.Objector

	// CountSelectedRows returns the number of rows that have been selected in
	// @tree.
	CountSelectedRows() int
	// Mode gets the selection mode for @selection. See
	// gtk_tree_selection_set_mode().
	Mode() SelectionMode
	// Selected sets @iter to the currently selected node if @selection is set
	// to K_SELECTION_SINGLE or K_SELECTION_BROWSE. @iter may be NULL if you
	// just want to test if @selection has any selected nodes. @model is filled
	// with the current model as a convenience. This function will not work if
	// you use @selection is K_SELECTION_MULTIPLE.
	Selected() (model TreeModel, iter TreeIter, ok bool)
	// SelectedRows creates a list of path of all selected rows. Additionally,
	// if you are planning on modifying the model after calling this function,
	// you may want to convert the returned list into a list of
	// TreeRowReferences. To do this, you can use gtk_tree_row_reference_new().
	//
	// To free the return value, use:
	//
	//    g_list_free_full (list, (GDestroyNotify) gtk_tree_path_free);
	SelectedRows() (model TreeModel, list *glib.List)
	// TreeView returns the tree view associated with @selection.
	TreeView() TreeView
	// UserData returns the user data for the selection function.
	UserData() interface{}
	// IterIsSelected returns true if the row at @iter is currently selected.
	IterIsSelected(iter *TreeIter) bool
	// PathIsSelected returns true if the row pointed to by @path is currently
	// selected. If @path does not point to a valid location, false is returned
	PathIsSelected(path *TreePath) bool
	// SelectAll selects all the nodes. @selection must be set to
	// K_SELECTION_MULTIPLE mode.
	SelectAll()
	// SelectIter selects the specified iterator.
	SelectIter(iter *TreeIter)
	// SelectPath: select the row at @path.
	SelectPath(path *TreePath)
	// SelectRange selects a range of nodes, determined by @start_path and
	// @end_path inclusive. @selection must be set to K_SELECTION_MULTIPLE mode.
	SelectRange(startPath *TreePath, endPath *TreePath)
	// SelectedForeach calls a function for each selected node. Note that you
	// cannot modify the tree or selection from within this function. As a
	// result, gtk_tree_selection_get_selected_rows() might be more useful.
	SelectedForeach(fn TreeSelectionForeachFunc)
	// SetMode sets the selection mode of the @selection. If the previous type
	// was K_SELECTION_MULTIPLE, then the anchor is kept selected, if it was
	// previously selected.
	SetMode(typ SelectionMode)
	// SetSelectFunction sets the selection function.
	//
	// If set, this function is called before any node is selected or
	// unselected, giving some control over which nodes are selected. The select
	// function should return true if the state of the node may be toggled, and
	// false if the state of the node should be left unchanged.
	SetSelectFunction(fn TreeSelectionFunc)
	// UnselectAll unselects all the nodes.
	UnselectAll()
	// UnselectIter unselects the specified iterator.
	UnselectIter(iter *TreeIter)
	// UnselectPath unselects the row at @path.
	UnselectPath(path *TreePath)
	// UnselectRange unselects a range of nodes, determined by @start_path and
	// @end_path inclusive.
	UnselectRange(startPath *TreePath, endPath *TreePath)
}

// treeSelection implements the TreeSelection interface.
type treeSelection struct {
	gextras.Objector
}

var _ TreeSelection = (*treeSelection)(nil)

// WrapTreeSelection wraps a GObject to the right type. It is
// primarily used internally.
func WrapTreeSelection(obj *externglib.Object) TreeSelection {
	return TreeSelection{
		Objector: obj,
	}
}

func marshalTreeSelection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeSelection(obj), nil
}

// CountSelectedRows returns the number of rows that have been selected in
// @tree.
func (s treeSelection) CountSelectedRows() int {
	var arg0 *C.GtkTreeSelection

	arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))

	var cret C.gint
	var goret1 int

	cret = C.gtk_tree_selection_count_selected_rows(arg0)

	goret1 = C.gint(cret)

	return goret1
}

// Mode gets the selection mode for @selection. See
// gtk_tree_selection_set_mode().
func (s treeSelection) Mode() SelectionMode {
	var arg0 *C.GtkTreeSelection

	arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))

	var cret C.GtkSelectionMode
	var goret1 SelectionMode

	cret = C.gtk_tree_selection_get_mode(arg0)

	goret1 = SelectionMode(cret)

	return goret1
}

// Selected sets @iter to the currently selected node if @selection is set
// to K_SELECTION_SINGLE or K_SELECTION_BROWSE. @iter may be NULL if you
// just want to test if @selection has any selected nodes. @model is filled
// with the current model as a convenience. This function will not work if
// you use @selection is K_SELECTION_MULTIPLE.
func (s treeSelection) Selected() (model TreeModel, iter TreeIter, ok bool) {
	var arg0 *C.GtkTreeSelection

	arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))

	var arg1 **C.GtkTreeModel
	var ret1 *TreeModel
	var arg2 *C.GtkTreeIter
	var ret2 *TreeIter
	var cret C.gboolean
	var goret3 bool

	cret = C.gtk_tree_selection_get_selected(arg0, &arg1, &arg2)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(arg1.Native()))).(*TreeModel)
	ret2 = WrapTreeIter(unsafe.Pointer(arg2))
	goret3 = C.bool(cret) != C.false

	return ret1, ret2, goret3
}

// SelectedRows creates a list of path of all selected rows. Additionally,
// if you are planning on modifying the model after calling this function,
// you may want to convert the returned list into a list of
// TreeRowReferences. To do this, you can use gtk_tree_row_reference_new().
//
// To free the return value, use:
//
//    g_list_free_full (list, (GDestroyNotify) gtk_tree_path_free);
func (s treeSelection) SelectedRows() (model TreeModel, list *glib.List) {
	var arg0 *C.GtkTreeSelection

	arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))

	var arg1 **C.GtkTreeModel
	var ret1 *TreeModel
	var cret *C.GList
	var goret2 *glib.List

	cret = C.gtk_tree_selection_get_selected_rows(arg0, &arg1)

	ret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(arg1.Native()))).(*TreeModel)
	goret2 = glib.WrapList(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret2, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret1, goret2
}

// TreeView returns the tree view associated with @selection.
func (s treeSelection) TreeView() TreeView {
	var arg0 *C.GtkTreeSelection

	arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))

	var cret *C.GtkTreeView
	var goret1 TreeView

	cret = C.gtk_tree_selection_get_tree_view(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(TreeView)

	return goret1
}

// UserData returns the user data for the selection function.
func (s treeSelection) UserData() interface{} {
	var arg0 *C.GtkTreeSelection

	arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))

	var cret C.gpointer
	var goret1 interface{}

	cret = C.gtk_tree_selection_get_user_data(arg0)

	goret1 = C.gpointer(cret)

	return goret1
}

// IterIsSelected returns true if the row at @iter is currently selected.
func (s treeSelection) IterIsSelected(iter *TreeIter) bool {
	var arg0 *C.GtkTreeSelection
	var arg1 *C.GtkTreeIter

	arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_tree_selection_iter_is_selected(arg0, iter)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// PathIsSelected returns true if the row pointed to by @path is currently
// selected. If @path does not point to a valid location, false is returned
func (s treeSelection) PathIsSelected(path *TreePath) bool {
	var arg0 *C.GtkTreeSelection
	var arg1 *C.GtkTreePath

	arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_tree_selection_path_is_selected(arg0, path)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// SelectAll selects all the nodes. @selection must be set to
// K_SELECTION_MULTIPLE mode.
func (s treeSelection) SelectAll() {
	var arg0 *C.GtkTreeSelection

	arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))

	C.gtk_tree_selection_select_all(arg0)
}

// SelectIter selects the specified iterator.
func (s treeSelection) SelectIter(iter *TreeIter) {
	var arg0 *C.GtkTreeSelection
	var arg1 *C.GtkTreeIter

	arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))

	C.gtk_tree_selection_select_iter(arg0, iter)
}

// SelectPath: select the row at @path.
func (s treeSelection) SelectPath(path *TreePath) {
	var arg0 *C.GtkTreeSelection
	var arg1 *C.GtkTreePath

	arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	C.gtk_tree_selection_select_path(arg0, path)
}

// SelectRange selects a range of nodes, determined by @start_path and
// @end_path inclusive. @selection must be set to K_SELECTION_MULTIPLE mode.
func (s treeSelection) SelectRange(startPath *TreePath, endPath *TreePath) {
	var arg0 *C.GtkTreeSelection
	var arg1 *C.GtkTreePath
	var arg2 *C.GtkTreePath

	arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkTreePath)(unsafe.Pointer(startPath.Native()))
	arg2 = (*C.GtkTreePath)(unsafe.Pointer(endPath.Native()))

	C.gtk_tree_selection_select_range(arg0, startPath, endPath)
}

// SelectedForeach calls a function for each selected node. Note that you
// cannot modify the tree or selection from within this function. As a
// result, gtk_tree_selection_get_selected_rows() might be more useful.
func (s treeSelection) SelectedForeach(fn TreeSelectionForeachFunc) {
	var arg0 *C.GtkTreeSelection

	arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))

	C.gtk_tree_selection_selected_foreach(arg0, fn, data)
}

// SetMode sets the selection mode of the @selection. If the previous type
// was K_SELECTION_MULTIPLE, then the anchor is kept selected, if it was
// previously selected.
func (s treeSelection) SetMode(typ SelectionMode) {
	var arg0 *C.GtkTreeSelection
	var arg1 C.GtkSelectionMode

	arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))
	arg1 = (C.GtkSelectionMode)(typ)

	C.gtk_tree_selection_set_mode(arg0, typ)
}

// SetSelectFunction sets the selection function.
//
// If set, this function is called before any node is selected or
// unselected, giving some control over which nodes are selected. The select
// function should return true if the state of the node may be toggled, and
// false if the state of the node should be left unchanged.
func (s treeSelection) SetSelectFunction(fn TreeSelectionFunc) {
	var arg0 *C.GtkTreeSelection

	arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))

	C.gtk_tree_selection_set_select_function(arg0, fn, data, destroy)
}

// UnselectAll unselects all the nodes.
func (s treeSelection) UnselectAll() {
	var arg0 *C.GtkTreeSelection

	arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))

	C.gtk_tree_selection_unselect_all(arg0)
}

// UnselectIter unselects the specified iterator.
func (s treeSelection) UnselectIter(iter *TreeIter) {
	var arg0 *C.GtkTreeSelection
	var arg1 *C.GtkTreeIter

	arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))

	C.gtk_tree_selection_unselect_iter(arg0, iter)
}

// UnselectPath unselects the row at @path.
func (s treeSelection) UnselectPath(path *TreePath) {
	var arg0 *C.GtkTreeSelection
	var arg1 *C.GtkTreePath

	arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkTreePath)(unsafe.Pointer(path.Native()))

	C.gtk_tree_selection_unselect_path(arg0, path)
}

// UnselectRange unselects a range of nodes, determined by @start_path and
// @end_path inclusive.
func (s treeSelection) UnselectRange(startPath *TreePath, endPath *TreePath) {
	var arg0 *C.GtkTreeSelection
	var arg1 *C.GtkTreePath
	var arg2 *C.GtkTreePath

	arg0 = (*C.GtkTreeSelection)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkTreePath)(unsafe.Pointer(startPath.Native()))
	arg2 = (*C.GtkTreePath)(unsafe.Pointer(endPath.Native()))

	C.gtk_tree_selection_unselect_range(arg0, startPath, endPath)
}

type TreeSelectionPrivate struct {
	native C.GtkTreeSelectionPrivate
}

// WrapTreeSelectionPrivate wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTreeSelectionPrivate(ptr unsafe.Pointer) *TreeSelectionPrivate {
	if ptr == nil {
		return nil
	}

	return (*TreeSelectionPrivate)(ptr)
}

func marshalTreeSelectionPrivate(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTreeSelectionPrivate(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TreeSelectionPrivate) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}
