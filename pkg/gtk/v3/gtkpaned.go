// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
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
		{T: externglib.Type(C.gtk_paned_get_type()), F: marshalPaned},
	})
}

// Paned has two panes, arranged either horizontally or vertically. The division
// between the two panes is adjustable by the user by dragging a handle.
//
// Child widgets are added to the panes of the widget with gtk_paned_pack1() and
// gtk_paned_pack2(). The division between the two children is set by default
// from the size requests of the children, but it can be adjusted by the user.
//
// A paned widget draws a separator between the two child widgets and a small
// handle that the user can drag to adjust the division. It does not draw any
// relief around the children or around the separator. (The space in which the
// separator is called the gutter.) Often, it is useful to put each child inside
// a Frame with the shadow type set to GTK_SHADOW_IN so that the gutter appears
// as a ridge. No separator is drawn if one of the children is missing.
//
// Each child has two options that can be set, @resize and @shrink. If @resize
// is true, then when the Paned is resized, that child will expand or shrink
// along with the paned widget. If @shrink is true, then that child can be made
// smaller than its requisition by the user. Setting @shrink to false allows the
// application to set a minimum size. If @resize is false for both children,
// then this is treated as if @resize is true for both children.
//
// The application can set the position of the slider as if it were set by the
// user, by calling gtk_paned_set_position().
//
// CSS nodes
//
//    GtkWidget *hpaned = gtk_paned_new (GTK_ORIENTATION_HORIZONTAL);
//    GtkWidget *frame1 = gtk_frame_new (NULL);
//    GtkWidget *frame2 = gtk_frame_new (NULL);
//    gtk_frame_set_shadow_type (GTK_FRAME (frame1), GTK_SHADOW_IN);
//    gtk_frame_set_shadow_type (GTK_FRAME (frame2), GTK_SHADOW_IN);
//
//    gtk_widget_set_size_request (hpaned, 200, -1);
//
//    gtk_paned_pack1 (GTK_PANED (hpaned), frame1, TRUE, FALSE);
//    gtk_widget_set_size_request (frame1, 50, -1);
//
//    gtk_paned_pack2 (GTK_PANED (hpaned), frame2, FALSE, FALSE);
//    gtk_widget_set_size_request (frame2, 50, -1);
type Paned interface {
	Container
	Buildable
	Orientable

	// Add1 adds a child to the top or left pane with default parameters. This
	// is equivalent to `gtk_paned_pack1 (paned, child, FALSE, TRUE)`.
	Add1(child Widget)
	// Add2 adds a child to the bottom or right pane with default parameters.
	// This is equivalent to `gtk_paned_pack2 (paned, child, TRUE, TRUE)`.
	Add2(child Widget)
	// Child1 obtains the first child of the paned widget.
	Child1() Widget
	// Child2 obtains the second child of the paned widget.
	Child2() Widget
	// HandleWindow returns the Window of the handle. This function is useful
	// when handling button or motion events because it enables the callback to
	// distinguish between the window of the paned, a child and the handle.
	HandleWindow() gdk.Window
	// Position obtains the position of the divider between the two panes.
	Position() int
	// WideHandle gets the Paned:wide-handle property.
	WideHandle() bool
	// Pack1 adds a child to the top or left pane.
	Pack1(child Widget, resize bool, shrink bool)
	// Pack2 adds a child to the bottom or right pane.
	Pack2(child Widget, resize bool, shrink bool)
	// SetPosition sets the position of the divider between the two panes.
	SetPosition(position int)
	// SetWideHandle sets the Paned:wide-handle property.
	SetWideHandle(wide bool)
}

// paned implements the Paned interface.
type paned struct {
	Container
	Buildable
	Orientable
}

var _ Paned = (*paned)(nil)

// WrapPaned wraps a GObject to the right type. It is
// primarily used internally.
func WrapPaned(obj *externglib.Object) Paned {
	return Paned{
		Container:  WrapContainer(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalPaned(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPaned(obj), nil
}

// NewPaned constructs a class Paned.
func NewPaned(orientation Orientation) Paned {
	var arg1 C.GtkOrientation

	arg1 = (C.GtkOrientation)(orientation)

	var cret C.GtkPaned

	cret = C.gtk_paned_new(arg1)

	var paned Paned

	paned = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Paned)

	return paned
}

// Add1 adds a child to the top or left pane with default parameters. This
// is equivalent to `gtk_paned_pack1 (paned, child, FALSE, TRUE)`.
func (p paned) Add1(child Widget) {
	var arg0 *C.GtkPaned
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_paned_add1(arg0, arg1)
}

// Add2 adds a child to the bottom or right pane with default parameters.
// This is equivalent to `gtk_paned_pack2 (paned, child, TRUE, TRUE)`.
func (p paned) Add2(child Widget) {
	var arg0 *C.GtkPaned
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_paned_add2(arg0, arg1)
}

// Child1 obtains the first child of the paned widget.
func (p paned) Child1() Widget {
	var arg0 *C.GtkPaned

	arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))

	var cret *C.GtkWidget

	cret = C.gtk_paned_get_child1(arg0)

	var widget Widget

	widget = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return widget
}

// Child2 obtains the second child of the paned widget.
func (p paned) Child2() Widget {
	var arg0 *C.GtkPaned

	arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))

	var cret *C.GtkWidget

	cret = C.gtk_paned_get_child2(arg0)

	var widget Widget

	widget = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return widget
}

// HandleWindow returns the Window of the handle. This function is useful
// when handling button or motion events because it enables the callback to
// distinguish between the window of the paned, a child and the handle.
func (p paned) HandleWindow() gdk.Window {
	var arg0 *C.GtkPaned

	arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))

	var cret *C.GdkWindow

	cret = C.gtk_paned_get_handle_window(arg0)

	var window gdk.Window

	window = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdk.Window)

	return window
}

// Position obtains the position of the divider between the two panes.
func (p paned) Position() int {
	var arg0 *C.GtkPaned

	arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))

	var cret C.gint

	cret = C.gtk_paned_get_position(arg0)

	var gint int

	gint = (int)(cret)

	return gint
}

// WideHandle gets the Paned:wide-handle property.
func (p paned) WideHandle() bool {
	var arg0 *C.GtkPaned

	arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))

	var cret C.gboolean

	cret = C.gtk_paned_get_wide_handle(arg0)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// Pack1 adds a child to the top or left pane.
func (p paned) Pack1(child Widget, resize bool, shrink bool) {
	var arg0 *C.GtkPaned
	var arg1 *C.GtkWidget
	var arg2 C.gboolean
	var arg3 C.gboolean

	arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if resize {
		arg2 = C.gboolean(1)
	}
	if shrink {
		arg3 = C.gboolean(1)
	}

	C.gtk_paned_pack1(arg0, arg1, arg2, arg3)
}

// Pack2 adds a child to the bottom or right pane.
func (p paned) Pack2(child Widget, resize bool, shrink bool) {
	var arg0 *C.GtkPaned
	var arg1 *C.GtkWidget
	var arg2 C.gboolean
	var arg3 C.gboolean

	arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if resize {
		arg2 = C.gboolean(1)
	}
	if shrink {
		arg3 = C.gboolean(1)
	}

	C.gtk_paned_pack2(arg0, arg1, arg2, arg3)
}

// SetPosition sets the position of the divider between the two panes.
func (p paned) SetPosition(position int) {
	var arg0 *C.GtkPaned
	var arg1 C.gint

	arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	arg1 = C.gint(position)

	C.gtk_paned_set_position(arg0, arg1)
}

// SetWideHandle sets the Paned:wide-handle property.
func (p paned) SetWideHandle(wide bool) {
	var arg0 *C.GtkPaned
	var arg1 C.gboolean

	arg0 = (*C.GtkPaned)(unsafe.Pointer(p.Native()))
	if wide {
		arg1 = C.gboolean(1)
	}

	C.gtk_paned_set_wide_handle(arg0, arg1)
}
