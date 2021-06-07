// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_layout_manager_get_type()), F: marshalLayoutManager},
	})
}

// LayoutManager: layout managers are delegate classes that handle the preferred
// size and the allocation of a container widget.
//
// You typically subclass LayoutManager if you want to implement a layout policy
// for the children of a widget, or if you want to determine the size of a
// widget depending on its contents.
//
// Each Widget can only have a LayoutManager instance associated to it at any
// given time; it is possible, though, to replace the layout manager instance
// using gtk_widget_set_layout_manager().
//
//
// Layout properties
//
// A layout manager can expose properties for controlling the layout of each
// child, by creating an object type derived from LayoutChild and installing the
// properties on it as normal GObject properties.
//
// Each LayoutChild instance storing the layout properties for a specific child
// is created through the gtk_layout_manager_get_layout_child() method; a
// LayoutManager controls the creation of its LayoutChild instances by
// overriding the GtkLayoutManagerClass.create_layout_child() virtual function.
// The typical implementation should look like:
//
//    static GtkLayoutChild *
//    create_layout_child (GtkLayoutManager *manager,
//                         GtkWidget        *container,
//                         GtkWidget        *child)
//    {
//      return g_object_new (your_layout_child_get_type (),
//                           "layout-manager", manager,
//                           "child-widget", child,
//                           NULL);
//    }
//
// The LayoutChild:layout-manager and LayoutChild:child-widget properties on the
// newly created LayoutChild instance are mandatory. The LayoutManager will
// cache the newly created LayoutChild instance until the widget is removed from
// its parent, or the parent removes the layout manager.
//
// Each LayoutManager instance creating a LayoutChild should use
// gtk_layout_manager_get_layout_child() every time it needs to query the layout
// properties; each LayoutChild instance should call
// gtk_layout_manager_layout_changed() every time a property is updated, in
// order to queue a new size measuring and allocation.
type LayoutManager interface {
	gextras.Objector

	// Allocate: this function assigns the given @width, @height, and @baseline
	// to a @widget, and computes the position and sizes of the children of the
	// @widget using the layout management policy of @manager.
	Allocate(m LayoutManager, widget Widget, width int, height int, baseline int)
	// LayoutChild retrieves a LayoutChild instance for the LayoutManager,
	// creating one if necessary.
	//
	// The @child widget must be a child of the widget using @manager.
	//
	// The LayoutChild instance is owned by the LayoutManager, and is guaranteed
	// to exist as long as @child is a child of the Widget using the given
	// LayoutManager.
	LayoutChild(m LayoutManager, child Widget)
	// RequestMode retrieves the request mode of @manager.
	RequestMode(m LayoutManager)
	// Widget retrieves the Widget using the given LayoutManager.
	Widget(m LayoutManager)
	// LayoutChanged queues a resize on the Widget using @manager, if any.
	//
	// This function should be called by subclasses of LayoutManager in response
	// to changes to their layout management policies.
	LayoutChanged(m LayoutManager)
	// Measure measures the size of the @widget using @manager, for the given
	// @orientation and size.
	//
	// See [GtkWidget's geometry management section][geometry-management] for
	// more details.
	Measure(m LayoutManager, widget Widget, orientation Orientation, forSize int) (minimum int, natural int, minimumBaseline int, naturalBaseline int)
}

// layoutManager implements the LayoutManager interface.
type layoutManager struct {
	gextras.Objector
}

var _ LayoutManager = (*layoutManager)(nil)

// WrapLayoutManager wraps a GObject to the right type. It is
// primarily used internally.
func WrapLayoutManager(obj *externglib.Object) LayoutManager {
	return LayoutManager{
		Objector: obj,
	}
}

func marshalLayoutManager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapLayoutManager(obj), nil
}

// Allocate: this function assigns the given @width, @height, and @baseline
// to a @widget, and computes the position and sizes of the children of the
// @widget using the layout management policy of @manager.
func (m layoutManager) Allocate(m LayoutManager, widget Widget, width int, height int, baseline int) {
	var arg0 *C.GtkLayoutManager
	var arg1 *C.GtkWidget
	var arg2 C.int
	var arg3 C.int
	var arg4 C.int

	arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 = C.int(width)
	arg3 = C.int(height)
	arg4 = C.int(baseline)

	C.gtk_layout_manager_allocate(arg0, arg1, arg2, arg3, arg4)
}

// LayoutChild retrieves a LayoutChild instance for the LayoutManager,
// creating one if necessary.
//
// The @child widget must be a child of the widget using @manager.
//
// The LayoutChild instance is owned by the LayoutManager, and is guaranteed
// to exist as long as @child is a child of the Widget using the given
// LayoutManager.
func (m layoutManager) LayoutChild(m LayoutManager, child Widget) {
	var arg0 *C.GtkLayoutManager
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_layout_manager_get_layout_child(arg0, arg1)
}

// RequestMode retrieves the request mode of @manager.
func (m layoutManager) RequestMode(m LayoutManager) {
	var arg0 *C.GtkLayoutManager

	arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(m.Native()))

	C.gtk_layout_manager_get_request_mode(arg0)
}

// Widget retrieves the Widget using the given LayoutManager.
func (m layoutManager) Widget(m LayoutManager) {
	var arg0 *C.GtkLayoutManager

	arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(m.Native()))

	C.gtk_layout_manager_get_widget(arg0)
}

// LayoutChanged queues a resize on the Widget using @manager, if any.
//
// This function should be called by subclasses of LayoutManager in response
// to changes to their layout management policies.
func (m layoutManager) LayoutChanged(m LayoutManager) {
	var arg0 *C.GtkLayoutManager

	arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(m.Native()))

	C.gtk_layout_manager_layout_changed(arg0)
}

// Measure measures the size of the @widget using @manager, for the given
// @orientation and size.
//
// See [GtkWidget's geometry management section][geometry-management] for
// more details.
func (m layoutManager) Measure(m LayoutManager, widget Widget, orientation Orientation, forSize int) (minimum int, natural int, minimumBaseline int, naturalBaseline int) {
	var arg0 *C.GtkLayoutManager
	var arg1 *C.GtkWidget
	var arg2 C.GtkOrientation
	var arg3 C.int

	arg0 = (*C.GtkLayoutManager)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 = (C.GtkOrientation)(orientation)
	arg3 = C.int(forSize)

	var arg4 C.int
	var minimum int
	var arg5 C.int
	var natural int
	var arg6 C.int
	var minimumBaseline int
	var arg7 C.int
	var naturalBaseline int

	C.gtk_layout_manager_measure(arg0, arg1, arg2, arg3, &arg4, &arg5, &arg6, &arg7)

	minimum = int(&arg4)
	natural = int(&arg5)
	minimumBaseline = int(&arg6)
	naturalBaseline = int(&arg7)

	return minimum, natural, minimumBaseline, naturalBaseline
}
