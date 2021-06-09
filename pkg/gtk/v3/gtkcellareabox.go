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
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_area_box_get_type()), F: marshalCellAreaBox},
	})
}

// CellAreaBox: the CellAreaBox renders cell renderers into a row or a column
// depending on its Orientation.
//
// GtkCellAreaBox uses a notion of packing. Packing refers to adding cell
// renderers with reference to a particular position in a CellAreaBox. There are
// two reference positions: the start and the end of the box. When the
// CellAreaBox is oriented in the GTK_ORIENTATION_VERTICAL orientation, the
// start is defined as the top of the box and the end is defined as the bottom.
// In the GTK_ORIENTATION_HORIZONTAL orientation start is defined as the left
// side and the end is defined as the right side.
//
// Alignments of CellRenderers rendered in adjacent rows can be configured by
// configuring the CellAreaBox align child cell property with
// gtk_cell_area_cell_set_property() or by specifying the "align" argument to
// gtk_cell_area_box_pack_start() and gtk_cell_area_box_pack_end().
type CellAreaBox interface {
	CellArea
	Buildable
	CellLayout
	Orientable

	// Spacing gets the spacing added between cell renderers.
	Spacing() int
	// PackEnd adds @renderer to @box, packed with reference to the end of @box.
	//
	// The @renderer is packed after (away from end of) any other CellRenderer
	// packed with reference to the end of @box.
	PackEnd(renderer CellRenderer, expand bool, align bool, fixed bool)
	// PackStart adds @renderer to @box, packed with reference to the start of
	// @box.
	//
	// The @renderer is packed after any other CellRenderer packed with
	// reference to the start of @box.
	PackStart(renderer CellRenderer, expand bool, align bool, fixed bool)
	// SetSpacing sets the spacing to add between cell renderers in @box.
	SetSpacing(spacing int)
}

// cellAreaBox implements the CellAreaBox interface.
type cellAreaBox struct {
	CellArea
	Buildable
	CellLayout
	Orientable
}

var _ CellAreaBox = (*cellAreaBox)(nil)

// WrapCellAreaBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapCellAreaBox(obj *externglib.Object) CellAreaBox {
	return CellAreaBox{
		CellArea:   WrapCellArea(obj),
		Buildable:  WrapBuildable(obj),
		CellLayout: WrapCellLayout(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalCellAreaBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellAreaBox(obj), nil
}

// NewCellAreaBox constructs a class CellAreaBox.
func NewCellAreaBox() CellAreaBox {
	var cret C.GtkCellAreaBox

	cret = C.gtk_cell_area_box_new()

	var cellAreaBox CellAreaBox

	cellAreaBox = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(CellAreaBox)

	return cellAreaBox
}

// Spacing gets the spacing added between cell renderers.
func (b cellAreaBox) Spacing() int {
	var arg0 *C.GtkCellAreaBox

	arg0 = (*C.GtkCellAreaBox)(unsafe.Pointer(b.Native()))

	var cret C.gint

	cret = C.gtk_cell_area_box_get_spacing(arg0)

	var gint int

	gint = (int)(cret)

	return gint
}

// PackEnd adds @renderer to @box, packed with reference to the end of @box.
//
// The @renderer is packed after (away from end of) any other CellRenderer
// packed with reference to the end of @box.
func (b cellAreaBox) PackEnd(renderer CellRenderer, expand bool, align bool, fixed bool) {
	var arg0 *C.GtkCellAreaBox
	var arg1 *C.GtkCellRenderer
	var arg2 C.gboolean
	var arg3 C.gboolean
	var arg4 C.gboolean

	arg0 = (*C.GtkCellAreaBox)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	if expand {
		arg2 = C.gboolean(1)
	}
	if align {
		arg3 = C.gboolean(1)
	}
	if fixed {
		arg4 = C.gboolean(1)
	}

	C.gtk_cell_area_box_pack_end(arg0, arg1, arg2, arg3, arg4)
}

// PackStart adds @renderer to @box, packed with reference to the start of
// @box.
//
// The @renderer is packed after any other CellRenderer packed with
// reference to the start of @box.
func (b cellAreaBox) PackStart(renderer CellRenderer, expand bool, align bool, fixed bool) {
	var arg0 *C.GtkCellAreaBox
	var arg1 *C.GtkCellRenderer
	var arg2 C.gboolean
	var arg3 C.gboolean
	var arg4 C.gboolean

	arg0 = (*C.GtkCellAreaBox)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	if expand {
		arg2 = C.gboolean(1)
	}
	if align {
		arg3 = C.gboolean(1)
	}
	if fixed {
		arg4 = C.gboolean(1)
	}

	C.gtk_cell_area_box_pack_start(arg0, arg1, arg2, arg3, arg4)
}

// SetSpacing sets the spacing to add between cell renderers in @box.
func (b cellAreaBox) SetSpacing(spacing int) {
	var arg0 *C.GtkCellAreaBox
	var arg1 C.gint

	arg0 = (*C.GtkCellAreaBox)(unsafe.Pointer(b.Native()))
	arg1 = C.gint(spacing)

	C.gtk_cell_area_box_set_spacing(arg0, arg1)
}
