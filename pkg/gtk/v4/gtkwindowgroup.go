// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_window_group_get_type()), F: marshalWindowGroup},
	})
}

// WindowGroup: `GtkWindowGroup` makes group of windows behave like separate
// applications.
//
// It achieves this by limiting the effect of GTK grabs and modality to windows
// in the same group.
//
// A window can be a member in at most one window group at a time. Windows that
// have not been explicitly assigned to a group are implicitly treated like
// windows of the default window group.
//
// `GtkWindowGroup` objects are referenced by each window in the group, so once
// you have added all windows to a `GtkWindowGroup`, you can drop the initial
// reference to the window group with g_object_unref(). If the windows in the
// window group are subsequently destroyed, then they will be removed from the
// window group and drop their references on the window group; when all window
// have been removed, the window group will be freed.
type WindowGroup interface {
	gextras.Objector

	// AddWindow adds a window to a `GtkWindowGroup`.
	AddWindow(window Window)
	// RemoveWindow removes a window from a `GtkWindowGroup`.
	RemoveWindow(window Window)
}

// windowGroup implements the WindowGroup interface.
type windowGroup struct {
	gextras.Objector
}

var _ WindowGroup = (*windowGroup)(nil)

// WrapWindowGroup wraps a GObject to the right type. It is
// primarily used internally.
func WrapWindowGroup(obj *externglib.Object) WindowGroup {
	return WindowGroup{
		Objector: obj,
	}
}

func marshalWindowGroup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWindowGroup(obj), nil
}

// AddWindow adds a window to a `GtkWindowGroup`.
func (w windowGroup) AddWindow(window Window) {
	var _arg0 *C.GtkWindowGroup // out
	var _arg1 *C.GtkWindow      // out

	_arg0 = (*C.GtkWindowGroup)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	C.gtk_window_group_add_window(_arg0, _arg1)
}

// RemoveWindow removes a window from a `GtkWindowGroup`.
func (w windowGroup) RemoveWindow(window Window) {
	var _arg0 *C.GtkWindowGroup // out
	var _arg1 *C.GtkWindow      // out

	_arg0 = (*C.GtkWindowGroup)(unsafe.Pointer(w.Native()))
	_arg1 = (*C.GtkWindow)(unsafe.Pointer(window.Native()))

	C.gtk_window_group_remove_window(_arg0, _arg1)
}
