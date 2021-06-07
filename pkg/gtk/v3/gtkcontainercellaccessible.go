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
		{T: externglib.Type(C.gtk_container_cell_accessible_get_type()), F: marshalContainerCellAccessible},
	})
}

type ContainerCellAccessible interface {
	CellAccessible

	AddChild(c ContainerCellAccessible, child CellAccessible)
	// Children: get a list of children.
	Children(c ContainerCellAccessible)

	RemoveChild(c ContainerCellAccessible, child CellAccessible)
}

// containerCellAccessible implements the ContainerCellAccessible interface.
type containerCellAccessible struct {
	CellAccessible
}

var _ ContainerCellAccessible = (*containerCellAccessible)(nil)

// WrapContainerCellAccessible wraps a GObject to the right type. It is
// primarily used internally.
func WrapContainerCellAccessible(obj *externglib.Object) ContainerCellAccessible {
	return ContainerCellAccessible{
		CellAccessible: WrapCellAccessible(obj),
	}
}

func marshalContainerCellAccessible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapContainerCellAccessible(obj), nil
}

// NewContainerCellAccessible constructs a class ContainerCellAccessible.
func NewContainerCellAccessible() {
	C.gtk_container_cell_accessible_new()
}

func (c containerCellAccessible) AddChild(c ContainerCellAccessible, child CellAccessible) {
	var arg0 *C.GtkContainerCellAccessible
	var arg1 *C.GtkCellAccessible

	arg0 = (*C.GtkContainerCellAccessible)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(child.Native()))

	C.gtk_container_cell_accessible_add_child(arg0, arg1)
}

// Children: get a list of children.
func (c containerCellAccessible) Children(c ContainerCellAccessible) {
	var arg0 *C.GtkContainerCellAccessible

	arg0 = (*C.GtkContainerCellAccessible)(unsafe.Pointer(c.Native()))

	C.gtk_container_cell_accessible_get_children(arg0)
}

func (c containerCellAccessible) RemoveChild(c ContainerCellAccessible, child CellAccessible) {
	var arg0 *C.GtkContainerCellAccessible
	var arg1 *C.GtkCellAccessible

	arg0 = (*C.GtkContainerCellAccessible)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkCellAccessible)(unsafe.Pointer(child.Native()))

	C.gtk_container_cell_accessible_remove_child(arg0, arg1)
}
