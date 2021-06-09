// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_mount_operation_get_type()), F: marshalMountOperation},
	})
}

// MountOperation: this should not be accessed directly. Use the accessor
// functions below.
type MountOperation interface {
	gio.MountOperation

	// Parent gets the transient parent used by the MountOperation
	Parent() Window
	// Screen gets the screen on which windows of the MountOperation will be
	// shown.
	Screen() gdk.Screen
	// IsShowing returns whether the MountOperation is currently displaying a
	// window.
	IsShowing() bool
	// SetParent sets the transient parent for windows shown by the
	// MountOperation.
	SetParent(parent Window)
	// SetScreen sets the screen to show windows of the MountOperation on.
	SetScreen(screen gdk.Screen)
}

// mountOperation implements the MountOperation interface.
type mountOperation struct {
	gio.MountOperation
}

var _ MountOperation = (*mountOperation)(nil)

// WrapMountOperation wraps a GObject to the right type. It is
// primarily used internally.
func WrapMountOperation(obj *externglib.Object) MountOperation {
	return MountOperation{
		gio.MountOperation: gio.WrapMountOperation(obj),
	}
}

func marshalMountOperation(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMountOperation(obj), nil
}

// NewMountOperation constructs a class MountOperation.
func NewMountOperation(parent Window) MountOperation {
	var arg1 *C.GtkWindow

	arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	var cret C.GtkMountOperation

	cret = C.gtk_mount_operation_new(arg1)

	var mountOperation MountOperation

	mountOperation = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(MountOperation)

	return mountOperation
}

// Parent gets the transient parent used by the MountOperation
func (o mountOperation) Parent() Window {
	var arg0 *C.GtkMountOperation

	arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))

	var cret *C.GtkWindow

	cret = C.gtk_mount_operation_get_parent(arg0)

	var window Window

	window = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Window)

	return window
}

// Screen gets the screen on which windows of the MountOperation will be
// shown.
func (o mountOperation) Screen() gdk.Screen {
	var arg0 *C.GtkMountOperation

	arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))

	var cret *C.GdkScreen

	cret = C.gtk_mount_operation_get_screen(arg0)

	var screen gdk.Screen

	screen = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdk.Screen)

	return screen
}

// IsShowing returns whether the MountOperation is currently displaying a
// window.
func (o mountOperation) IsShowing() bool {
	var arg0 *C.GtkMountOperation

	arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))

	var cret C.gboolean

	cret = C.gtk_mount_operation_is_showing(arg0)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// SetParent sets the transient parent for windows shown by the
// MountOperation.
func (o mountOperation) SetParent(parent Window) {
	var arg0 *C.GtkMountOperation
	var arg1 *C.GtkWindow

	arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))
	arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	C.gtk_mount_operation_set_parent(arg0, arg1)
}

// SetScreen sets the screen to show windows of the MountOperation on.
func (o mountOperation) SetScreen(screen gdk.Screen) {
	var arg0 *C.GtkMountOperation
	var arg1 *C.GdkScreen

	arg0 = (*C.GtkMountOperation)(unsafe.Pointer(o.Native()))
	arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	C.gtk_mount_operation_set_screen(arg0, arg1)
}
