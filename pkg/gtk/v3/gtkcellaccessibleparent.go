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
		{T: externglib.Type(C.gtk_cell_accessible_parent_get_type()), F: marshalCellAccessibleParent},
	})
}

// CellAccessibleParentOverrider contains methods that are overridable. This
// interface is a subset of the interface CellAccessibleParent.
type CellAccessibleParentOverrider interface {
	Activate(p CellAccessibleParent, cell CellAccessible)

	Edit(p CellAccessibleParent, cell CellAccessible)

	ExpandCollapse(p CellAccessibleParent, cell CellAccessible)

	CellArea(p CellAccessibleParent, cell CellAccessible) *gdk.Rectangle

	CellPosition(p CellAccessibleParent, cell CellAccessible) (row int, column int)

	ChildIndex(p CellAccessibleParent, cell CellAccessible)

	ColumnHeaderCells(p CellAccessibleParent, cell CellAccessible)

	RendererState(p CellAccessibleParent, cell CellAccessible)

	RowHeaderCells(p CellAccessibleParent, cell CellAccessible)

	GrabFocus(p CellAccessibleParent, cell CellAccessible) bool
}

type CellAccessibleParent interface {
	gextras.Objector
	CellAccessibleParentOverrider
}

// cellAccessibleParent implements the CellAccessibleParent interface.
type cellAccessibleParent struct {
	gextras.Objector
}

var _ CellAccessibleParent = (*cellAccessibleParent)(nil)

// WrapCellAccessibleParent wraps a GObject to a type that implements interface
// CellAccessibleParent. It is primarily used internally.
func WrapCellAccessibleParent(obj *externglib.Object) CellAccessibleParent {
	return CellAccessibleParent{
		Objector: obj,
	}
}

func marshalCellAccessibleParent(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellAccessibleParent(obj), nil
}

func (p cellAccessibleParent) Activate(p CellAccessibleParent, cell CellAccessible) {
	var arg0 *C.GtkCellAccessibleParent
	var arg1 *C.GtkCellAccessible

	arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_accessible_parent_activate(arg0, arg1)
}

func (p cellAccessibleParent) Edit(p CellAccessibleParent, cell CellAccessible) {
	var arg0 *C.GtkCellAccessibleParent
	var arg1 *C.GtkCellAccessible

	arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_accessible_parent_edit(arg0, arg1)
}

func (p cellAccessibleParent) ExpandCollapse(p CellAccessibleParent, cell CellAccessible) {
	var arg0 *C.GtkCellAccessibleParent
	var arg1 *C.GtkCellAccessible

	arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_accessible_parent_expand_collapse(arg0, arg1)
}

func (p cellAccessibleParent) CellArea(p CellAccessibleParent, cell CellAccessible) *gdk.Rectangle {
	var arg0 *C.GtkCellAccessibleParent
	var arg1 *C.GtkCellAccessible

	arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	var arg2 C.GdkRectangle
	var cellRect *gdk.Rectangle

	C.gtk_cell_accessible_parent_get_cell_area(arg0, arg1, &arg2)

	cellRect = gdk.WrapRectangle(unsafe.Pointer(&arg2))

	return cellRect
}

func (p cellAccessibleParent) CellPosition(p CellAccessibleParent, cell CellAccessible) (row int, column int) {
	var arg0 *C.GtkCellAccessibleParent
	var arg1 *C.GtkCellAccessible

	arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	var arg2 C.gint
	var row int
	var arg3 C.gint
	var column int

	C.gtk_cell_accessible_parent_get_cell_position(arg0, arg1, &arg2, &arg3)

	row = int(&arg2)
	column = int(&arg3)

	return row, column
}

func (p cellAccessibleParent) ChildIndex(p CellAccessibleParent, cell CellAccessible) {
	var arg0 *C.GtkCellAccessibleParent
	var arg1 *C.GtkCellAccessible

	arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_accessible_parent_get_child_index(arg0, arg1)
}

func (p cellAccessibleParent) ColumnHeaderCells(p CellAccessibleParent, cell CellAccessible) {
	var arg0 *C.GtkCellAccessibleParent
	var arg1 *C.GtkCellAccessible

	arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_accessible_parent_get_column_header_cells(arg0, arg1)
}

func (p cellAccessibleParent) RendererState(p CellAccessibleParent, cell CellAccessible) {
	var arg0 *C.GtkCellAccessibleParent
	var arg1 *C.GtkCellAccessible

	arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_accessible_parent_get_renderer_state(arg0, arg1)
}

func (p cellAccessibleParent) RowHeaderCells(p CellAccessibleParent, cell CellAccessible) {
	var arg0 *C.GtkCellAccessibleParent
	var arg1 *C.GtkCellAccessible

	arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	C.gtk_cell_accessible_parent_get_row_header_cells(arg0, arg1)
}

func (p cellAccessibleParent) GrabFocus(p CellAccessibleParent, cell CellAccessible) bool {
	var arg0 *C.GtkCellAccessibleParent
	var arg1 *C.GtkCellAccessible

	arg0 = (*C.GtkCellAccessibleParent)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(cell.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_cell_accessible_parent_grab_focus(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}