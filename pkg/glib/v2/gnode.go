// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// TraverseType specifies the type of traversal performed by g_tree_traverse(),
// g_node_traverse() and g_node_find(). The different orders are illustrated
// here: - In order: A, B, C, D, E, F, G, H, I !
// (Sorted_binary_tree_inorder.svg) - Pre order: F, B, A, D, C, E, G, I, H !
// (Sorted_binary_tree_preorder.svg) - Post order: A, C, E, D, B, H, I, G, F !
// (Sorted_binary_tree_postorder.svg) - Level order: F, B, G, A, D, I, C, E, H !
// (Sorted_binary_tree_breadth-first_traversal.svg)
type TraverseType int

const (
	// TraverseTypeInOrder vists a node's left child first, then the node
	// itself, then its right child. This is the one to use if you want the
	// output sorted according to the compare function.
	TraverseTypeInOrder TraverseType = 0
	// TraverseTypePreOrder visits a node, then its children.
	TraverseTypePreOrder TraverseType = 1
	// TraverseTypePostOrder visits the node's children, then the node itself.
	TraverseTypePostOrder TraverseType = 2
	// TraverseTypeLevelOrder is not implemented for [balanced binary
	// trees][glib-Balanced-Binary-Trees]. For [n-ary trees][glib-N-ary-Trees],
	// it vists the root node first, then its children, then its grandchildren,
	// and so on. Note that this is less efficient than the other orders.
	TraverseTypeLevelOrder TraverseType = 3
)

// TraverseFlags specifies which nodes are visited during several of the tree
// functions, including g_node_traverse() and g_node_find().
type TraverseFlags int

const (
	// TraverseFlagsLeaves: only leaf nodes should be visited. This name has
	// been introduced in 2.6, for older version use G_TRAVERSE_LEAFS.
	TraverseFlagsLeaves TraverseFlags = 1
	// TraverseFlagsNonLeaves: only non-leaf nodes should be visited. This name
	// has been introduced in 2.6, for older version use G_TRAVERSE_NON_LEAFS.
	TraverseFlagsNonLeaves TraverseFlags = 2
	// TraverseFlagsAll: all nodes should be visited.
	TraverseFlagsAll TraverseFlags = 3
	// TraverseFlagsMask: a mask of all traverse flags.
	TraverseFlagsMask TraverseFlags = 3
	// TraverseFlagsLeafs: identical to G_TRAVERSE_LEAVES.
	TraverseFlagsLeafs TraverseFlags = 1
	// TraverseFlagsNonLeafs: identical to G_TRAVERSE_NON_LEAVES.
	TraverseFlagsNonLeafs TraverseFlags = 2
)

// Node: the #GNode struct represents one node in a [n-ary
// tree][glib-N-ary-Trees].
type Node struct {
	native C.GNode
}

// WrapNode wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapNode(ptr unsafe.Pointer) *Node {
	if ptr == nil {
		return nil
	}

	return (*Node)(ptr)
}

// Native returns the underlying C source pointer.
func (n *Node) Native() unsafe.Pointer {
	return unsafe.Pointer(&n.native)
}

// Next gets the field inside the struct.
func (n *Node) Next() *Node {
	var v *Node // out
	v = WrapNode(unsafe.Pointer(n.native.next))
	return v
}

// Prev gets the field inside the struct.
func (n *Node) Prev() *Node {
	var v *Node // out
	v = WrapNode(unsafe.Pointer(n.native.prev))
	return v
}

// Parent gets the field inside the struct.
func (n *Node) Parent() *Node {
	var v *Node // out
	v = WrapNode(unsafe.Pointer(n.native.parent))
	return v
}

// Children gets the field inside the struct.
func (n *Node) Children() *Node {
	var v *Node // out
	v = WrapNode(unsafe.Pointer(n.native.children))
	return v
}

// ChildPosition gets the position of a #GNode with respect to its siblings.
// @child must be a child of @node. The first child is numbered 0, the second 1,
// and so on.
func (n *Node) ChildPosition(child *Node) int {
	var _arg0 *C.GNode // out
	var _arg1 *C.GNode // out

	_arg0 = (*C.GNode)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GNode)(unsafe.Pointer(child.Native()))

	var _cret C.gint // in

	_cret = C.g_node_child_position(_arg0, _arg1)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Depth gets the depth of a #GNode.
//
// If @node is nil the depth is 0. The root node has a depth of 1. For the
// children of the root node the depth is 2. And so on.
func (n *Node) Depth() uint {
	var _arg0 *C.GNode // out

	_arg0 = (*C.GNode)(unsafe.Pointer(n.Native()))

	var _cret C.guint // in

	_cret = C.g_node_depth(_arg0)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// Destroy removes @root and its children from the tree, freeing any memory
// allocated.
func (r *Node) Destroy() {
	var _arg0 *C.GNode // out

	_arg0 = (*C.GNode)(unsafe.Pointer(r.Native()))

	C.g_node_destroy(_arg0)
}

// IsAncestor returns true if @node is an ancestor of @descendant. This is true
// if node is the parent of @descendant, or if node is the grandparent of
// @descendant etc.
func (n *Node) IsAncestor(descendant *Node) bool {
	var _arg0 *C.GNode // out
	var _arg1 *C.GNode // out

	_arg0 = (*C.GNode)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.GNode)(unsafe.Pointer(descendant.Native()))

	var _cret C.gboolean // in

	_cret = C.g_node_is_ancestor(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MaxHeight gets the maximum height of all branches beneath a #GNode. This is
// the maximum distance from the #GNode to all leaf nodes.
//
// If @root is nil, 0 is returned. If @root has no children, 1 is returned. If
// @root has children, 2 is returned. And so on.
func (r *Node) MaxHeight() uint {
	var _arg0 *C.GNode // out

	_arg0 = (*C.GNode)(unsafe.Pointer(r.Native()))

	var _cret C.guint // in

	_cret = C.g_node_max_height(_arg0)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// NChildren gets the number of children of a #GNode.
func (n *Node) NChildren() uint {
	var _arg0 *C.GNode // out

	_arg0 = (*C.GNode)(unsafe.Pointer(n.Native()))

	var _cret C.guint // in

	_cret = C.g_node_n_children(_arg0)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// NNodes gets the number of nodes in a tree.
func (r *Node) NNodes(flags TraverseFlags) uint {
	var _arg0 *C.GNode         // out
	var _arg1 C.GTraverseFlags // out

	_arg0 = (*C.GNode)(unsafe.Pointer(r.Native()))
	_arg1 = (C.GTraverseFlags)(flags)

	var _cret C.guint // in

	_cret = C.g_node_n_nodes(_arg0, _arg1)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// ReverseChildren reverses the order of the children of a #GNode. (It doesn't
// change the order of the grandchildren.)
func (n *Node) ReverseChildren() {
	var _arg0 *C.GNode // out

	_arg0 = (*C.GNode)(unsafe.Pointer(n.Native()))

	C.g_node_reverse_children(_arg0)
}

// Unlink unlinks a #GNode from a tree, resulting in two separate trees.
func (n *Node) Unlink() {
	var _arg0 *C.GNode // out

	_arg0 = (*C.GNode)(unsafe.Pointer(n.Native()))

	C.g_node_unlink(_arg0)
}
