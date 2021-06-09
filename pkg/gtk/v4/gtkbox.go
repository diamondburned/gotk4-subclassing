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
		{T: externglib.Type(C.gtk_box_get_type()), F: marshalBox},
	})
}

// Box: the GtkBox widget arranges child widgets into a single row or column,
// depending upon the value of its Orientable:orientation property. Within the
// other dimension, all children are allocated the same size. Of course, the
// Widget:halign and Widget:valign properties can be used on the children to
// influence their allocation.
//
// Use repeated calls to gtk_box_append() to pack widgets into a GtkBox from
// start to end. Use gtk_box_remove() to remove widgets from the GtkBox.
// gtk_box_insert_child_after() can be used to add a child at a particular
// position.
//
// Use gtk_box_set_homogeneous() to specify whether or not all children of the
// GtkBox are forced to get the same amount of space.
//
// Use gtk_box_set_spacing() to determine how much space will be minimally
// placed between all children in the GtkBox. Note that spacing is added between
// the children.
//
// Use gtk_box_reorder_child_after() to move a child to a different place in the
// box.
//
//
// CSS nodes
//
// GtkBox uses a single CSS node with name box.
//
//
// Accessibility
//
// GtkBox uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type Box interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Orientable

	// Append adds @child as the last child to @box.
	Append(child Widget)
	// BaselinePosition gets the value set by gtk_box_set_baseline_position().
	BaselinePosition() BaselinePosition
	// Homogeneous returns whether the box is homogeneous (all children are the
	// same size). See gtk_box_set_homogeneous().
	Homogeneous() bool
	// Spacing gets the value set by gtk_box_set_spacing().
	Spacing() int
	// InsertChildAfter inserts @child in the position after @sibling in the
	// list of @box children. If @sibling is nil, insert @child at the first
	// position.
	InsertChildAfter(child Widget, sibling Widget)
	// Prepend adds @child as the first child to @box.
	Prepend(child Widget)
	// Remove removes a child widget from @box, after it has been added with
	// gtk_box_append(), gtk_box_prepend(), or gtk_box_insert_child_after().
	Remove(child Widget)
	// ReorderChildAfter moves @child to the position after @sibling in the list
	// of @box children. If @sibling is nil, move @child to the first position.
	ReorderChildAfter(child Widget, sibling Widget)
	// SetBaselinePosition sets the baseline position of a box. This affects
	// only horizontal boxes with at least one baseline aligned child. If there
	// is more vertical space available than requested, and the baseline is not
	// allocated by the parent then @position is used to allocate the baseline
	// wrt the extra space available.
	SetBaselinePosition(position BaselinePosition)
	// SetHomogeneous sets the Box:homogeneous property of @box, controlling
	// whether or not all children of @box are given equal space in the box.
	SetHomogeneous(homogeneous bool)
	// SetSpacing sets the Box:spacing property of @box, which is the number of
	// pixels to place between children of @box.
	SetSpacing(spacing int)
}

// box implements the Box interface.
type box struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Orientable
}

var _ Box = (*box)(nil)

// WrapBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapBox(obj *externglib.Object) Box {
	return Box{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		Orientable:       WrapOrientable(obj),
	}
}

func marshalBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBox(obj), nil
}

// NewBox constructs a class Box.
func NewBox(orientation Orientation, spacing int) Box {
	var arg1 C.GtkOrientation
	var arg2 C.int

	arg1 = (C.GtkOrientation)(orientation)
	arg2 = C.int(spacing)

	var cret C.GtkBox

	cret = C.gtk_box_new(arg1, arg2)

	var box Box

	box = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Box)

	return box
}

// Append adds @child as the last child to @box.
func (b box) Append(child Widget) {
	var arg0 *C.GtkBox
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_box_append(arg0, arg1)
}

// BaselinePosition gets the value set by gtk_box_set_baseline_position().
func (b box) BaselinePosition() BaselinePosition {
	var arg0 *C.GtkBox

	arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))

	var cret C.GtkBaselinePosition

	cret = C.gtk_box_get_baseline_position(arg0)

	var baselinePosition BaselinePosition

	baselinePosition = BaselinePosition(cret)

	return baselinePosition
}

// Homogeneous returns whether the box is homogeneous (all children are the
// same size). See gtk_box_set_homogeneous().
func (b box) Homogeneous() bool {
	var arg0 *C.GtkBox

	arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))

	var cret C.gboolean

	cret = C.gtk_box_get_homogeneous(arg0)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// Spacing gets the value set by gtk_box_set_spacing().
func (b box) Spacing() int {
	var arg0 *C.GtkBox

	arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))

	var cret C.int

	cret = C.gtk_box_get_spacing(arg0)

	var gint int

	gint = (int)(cret)

	return gint
}

// InsertChildAfter inserts @child in the position after @sibling in the
// list of @box children. If @sibling is nil, insert @child at the first
// position.
func (b box) InsertChildAfter(child Widget, sibling Widget) {
	var arg0 *C.GtkBox
	var arg1 *C.GtkWidget
	var arg2 *C.GtkWidget

	arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = (*C.GtkWidget)(unsafe.Pointer(sibling.Native()))

	C.gtk_box_insert_child_after(arg0, arg1, arg2)
}

// Prepend adds @child as the first child to @box.
func (b box) Prepend(child Widget) {
	var arg0 *C.GtkBox
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_box_prepend(arg0, arg1)
}

// Remove removes a child widget from @box, after it has been added with
// gtk_box_append(), gtk_box_prepend(), or gtk_box_insert_child_after().
func (b box) Remove(child Widget) {
	var arg0 *C.GtkBox
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_box_remove(arg0, arg1)
}

// ReorderChildAfter moves @child to the position after @sibling in the list
// of @box children. If @sibling is nil, move @child to the first position.
func (b box) ReorderChildAfter(child Widget, sibling Widget) {
	var arg0 *C.GtkBox
	var arg1 *C.GtkWidget
	var arg2 *C.GtkWidget

	arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = (*C.GtkWidget)(unsafe.Pointer(sibling.Native()))

	C.gtk_box_reorder_child_after(arg0, arg1, arg2)
}

// SetBaselinePosition sets the baseline position of a box. This affects
// only horizontal boxes with at least one baseline aligned child. If there
// is more vertical space available than requested, and the baseline is not
// allocated by the parent then @position is used to allocate the baseline
// wrt the extra space available.
func (b box) SetBaselinePosition(position BaselinePosition) {
	var arg0 *C.GtkBox
	var arg1 C.GtkBaselinePosition

	arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	arg1 = (C.GtkBaselinePosition)(position)

	C.gtk_box_set_baseline_position(arg0, arg1)
}

// SetHomogeneous sets the Box:homogeneous property of @box, controlling
// whether or not all children of @box are given equal space in the box.
func (b box) SetHomogeneous(homogeneous bool) {
	var arg0 *C.GtkBox
	var arg1 C.gboolean

	arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	if homogeneous {
		arg1 = C.gboolean(1)
	}

	C.gtk_box_set_homogeneous(arg0, arg1)
}

// SetSpacing sets the Box:spacing property of @box, which is the number of
// pixels to place between children of @box.
func (b box) SetSpacing(spacing int) {
	var arg0 *C.GtkBox
	var arg1 C.int

	arg0 = (*C.GtkBox)(unsafe.Pointer(b.Native()))
	arg1 = C.int(spacing)

	C.gtk_box_set_spacing(arg0, arg1)
}
