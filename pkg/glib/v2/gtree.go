// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_tree_get_type()), F: marshalTree},
	})
}

// Tree: the GTree struct is an opaque data structure representing a [balanced
// binary tree][glib-Balanced-Binary-Trees]. It should be accessed only by using
// the following functions.
type Tree struct {
	native C.GTree
}

// WrapTree wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTree(ptr unsafe.Pointer) *Tree {
	if ptr == nil {
		return nil
	}

	return (*Tree)(ptr)
}

func marshalTree(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTree(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *Tree) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// Destroy removes all keys and values from the #GTree and decreases its
// reference count by one. If keys and/or values are dynamically allocated, you
// should either free them first or create the #GTree using g_tree_new_full().
// In the latter case the destroy functions you supplied will be called on all
// keys and values before destroying the #GTree.
func (t *Tree) Destroy() {
	var _arg0 *C.GTree // out

	_arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))

	C.g_tree_destroy(_arg0)
}

// Height gets the height of a #GTree.
//
// If the #GTree contains no nodes, the height is 0. If the #GTree contains only
// one root node the height is 1. If the root node has children the height is 2,
// etc.
func (t *Tree) Height() int {
	var _arg0 *C.GTree // out

	_arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))

	var _cret C.gint // in

	_cret = C.g_tree_height(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Nnodes gets the number of nodes in a #GTree.
func (t *Tree) Nnodes() int {
	var _arg0 *C.GTree // out

	_arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))

	var _cret C.gint // in

	_cret = C.g_tree_nnodes(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// NodeFirst returns the first in-order node of the tree, or nil for an empty
// tree.
func (t *Tree) NodeFirst() *TreeNode {
	var _arg0 *C.GTree // out

	_arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))

	var _cret *C.GTreeNode // in

	_cret = C.g_tree_node_first(_arg0)

	var _treeNode *TreeNode // out

	_treeNode = WrapTreeNode(unsafe.Pointer(_cret))

	return _treeNode
}

// NodeLast returns the last in-order node of the tree, or nil for an empty
// tree.
func (t *Tree) NodeLast() *TreeNode {
	var _arg0 *C.GTree // out

	_arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))

	var _cret *C.GTreeNode // in

	_cret = C.g_tree_node_last(_arg0)

	var _treeNode *TreeNode // out

	_treeNode = WrapTreeNode(unsafe.Pointer(_cret))

	return _treeNode
}

// Ref increments the reference count of @tree by one.
//
// It is safe to call this function from any thread.
func (t *Tree) Ref() *Tree {
	var _arg0 *C.GTree // out

	_arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))

	var _cret *C.GTree // in

	_cret = C.g_tree_ref(_arg0)

	var _ret *Tree // out

	_ret = WrapTree(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_ret, func(v *Tree) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _ret
}

// Unref decrements the reference count of @tree by one. If the reference count
// drops to 0, all keys and values will be destroyed (if destroy functions were
// specified) and all memory allocated by @tree will be released.
//
// It is safe to call this function from any thread.
func (t *Tree) Unref() {
	var _arg0 *C.GTree // out

	_arg0 = (*C.GTree)(unsafe.Pointer(t.Native()))

	C.g_tree_unref(_arg0)
}

// TreeNode: an opaque type which identifies a specific node in a #GTree.
type TreeNode struct {
	native C.GTreeNode
}

// WrapTreeNode wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTreeNode(ptr unsafe.Pointer) *TreeNode {
	if ptr == nil {
		return nil
	}

	return (*TreeNode)(ptr)
}

// Native returns the underlying C source pointer.
func (t *TreeNode) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// Next returns the next in-order node of the tree, or nil if the passed node
// was already the last one.
func (n *TreeNode) Next() *TreeNode {
	var _arg0 *C.GTreeNode // out

	_arg0 = (*C.GTreeNode)(unsafe.Pointer(n.Native()))

	var _cret *C.GTreeNode // in

	_cret = C.g_tree_node_next(_arg0)

	var _treeNode *TreeNode // out

	_treeNode = WrapTreeNode(unsafe.Pointer(_cret))

	return _treeNode
}

// Previous returns the previous in-order node of the tree, or nil if the passed
// node was already the first one.
func (n *TreeNode) Previous() *TreeNode {
	var _arg0 *C.GTreeNode // out

	_arg0 = (*C.GTreeNode)(unsafe.Pointer(n.Native()))

	var _cret *C.GTreeNode // in

	_cret = C.g_tree_node_previous(_arg0)

	var _treeNode *TreeNode // out

	_treeNode = WrapTreeNode(unsafe.Pointer(_cret))

	return _treeNode
}
