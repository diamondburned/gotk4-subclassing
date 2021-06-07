// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-wayland gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/wayland/gdkwayland.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_wayland_popup_get_type()), F: marshalWaylandPopup},
		{T: externglib.Type(C.gdk_wayland_surface_get_type()), F: marshalWaylandSurface},
		{T: externglib.Type(C.gdk_wayland_toplevel_get_type()), F: marshalWaylandToplevel},
	})
}

type WaylandPopup interface {
	WaylandSurface
}

// waylandPopup implements the WaylandPopup interface.
type waylandPopup struct {
	WaylandSurface
}

var _ WaylandPopup = (*waylandPopup)(nil)

// WrapWaylandPopup wraps a GObject to the right type. It is
// primarily used internally.
func WrapWaylandPopup(obj *externglib.Object) WaylandPopup {
	return WaylandPopup{
		WaylandSurface: WrapWaylandSurface(obj),
	}
}

func marshalWaylandPopup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandPopup(obj), nil
}

type WaylandSurface interface {
	gdk.Surface

	// WlSurface returns the Wayland surface of a Surface.
	WlSurface(s WaylandSurface)
}

// waylandSurface implements the WaylandSurface interface.
type waylandSurface struct {
	gdk.Surface
}

var _ WaylandSurface = (*waylandSurface)(nil)

// WrapWaylandSurface wraps a GObject to the right type. It is
// primarily used internally.
func WrapWaylandSurface(obj *externglib.Object) WaylandSurface {
	return WaylandSurface{
		gdk.Surface: gdk.WrapSurface(obj),
	}
}

func marshalWaylandSurface(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandSurface(obj), nil
}

// WlSurface returns the Wayland surface of a Surface.
func (s waylandSurface) WlSurface(s WaylandSurface) {
	var arg0 *C.GdkSurface

	arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	C.gdk_wayland_surface_get_wl_surface(arg0)
}

type WaylandToplevel interface {
	WaylandSurface

	// ExportHandle: asynchronously obtains a handle for a surface that can be
	// passed to other processes. When the handle has been obtained, @callback
	// will be called.
	//
	// It is an error to call this function on a surface that is already
	// exported.
	//
	// When the handle is no longer needed,
	// gdk_wayland_toplevel_unexport_handle() should be called to clean up
	// resources.
	//
	// The main purpose for obtaining a handle is to mark a surface from another
	// surface as transient for this one, see
	// gdk_wayland_toplevel_set_transient_for_exported().
	//
	// Note that this API depends on an unstable Wayland protocol, and thus may
	// require changes in the future.
	ExportHandle(t WaylandToplevel) bool
	// SetApplicationID sets the application id on a Toplevel.
	SetApplicationID(t WaylandToplevel, applicationID string)
	// SetTransientForExported marks @toplevel as transient for the surface to
	// which the given @parent_handle_str refers. Typically, the handle will
	// originate from a gdk_wayland_toplevel_export_handle() call in another
	// process.
	//
	// Note that this API depends on an unstable Wayland protocol, and thus may
	// require changes in the future.
	SetTransientForExported(t WaylandToplevel, parentHandleStr string) bool
	// UnexportHandle destroys the handle that was obtained with
	// gdk_wayland_toplevel_export_handle().
	//
	// It is an error to call this function on a surface that does not have a
	// handle.
	//
	// Note that this API depends on an unstable Wayland protocol, and thus may
	// require changes in the future.
	UnexportHandle(t WaylandToplevel)
}

// waylandToplevel implements the WaylandToplevel interface.
type waylandToplevel struct {
	WaylandSurface
}

var _ WaylandToplevel = (*waylandToplevel)(nil)

// WrapWaylandToplevel wraps a GObject to the right type. It is
// primarily used internally.
func WrapWaylandToplevel(obj *externglib.Object) WaylandToplevel {
	return WaylandToplevel{
		WaylandSurface: WrapWaylandSurface(obj),
	}
}

func marshalWaylandToplevel(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandToplevel(obj), nil
}

// ExportHandle: asynchronously obtains a handle for a surface that can be
// passed to other processes. When the handle has been obtained, @callback
// will be called.
//
// It is an error to call this function on a surface that is already
// exported.
//
// When the handle is no longer needed,
// gdk_wayland_toplevel_unexport_handle() should be called to clean up
// resources.
//
// The main purpose for obtaining a handle is to mark a surface from another
// surface as transient for this one, see
// gdk_wayland_toplevel_set_transient_for_exported().
//
// Note that this API depends on an unstable Wayland protocol, and thus may
// require changes in the future.
func (t waylandToplevel) ExportHandle(t WaylandToplevel) bool {
	var arg0 *C.GdkToplevel

	arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gdk_wayland_toplevel_export_handle(arg0, arg1, arg2, arg3)

	if cret {
		ok = true
	}

	return ok
}

// SetApplicationID sets the application id on a Toplevel.
func (t waylandToplevel) SetApplicationID(t WaylandToplevel, applicationID string) {
	var arg0 *C.GdkToplevel
	var arg1 *C.char

	arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))
	arg1 = (*C.char)(C.CString(applicationID))
	defer C.free(unsafe.Pointer(arg1))

	C.gdk_wayland_toplevel_set_application_id(arg0, arg1)
}

// SetTransientForExported marks @toplevel as transient for the surface to
// which the given @parent_handle_str refers. Typically, the handle will
// originate from a gdk_wayland_toplevel_export_handle() call in another
// process.
//
// Note that this API depends on an unstable Wayland protocol, and thus may
// require changes in the future.
func (t waylandToplevel) SetTransientForExported(t WaylandToplevel, parentHandleStr string) bool {
	var arg0 *C.GdkToplevel
	var arg1 *C.char

	arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))
	arg1 = (*C.char)(C.CString(parentHandleStr))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ok bool

	cret = C.gdk_wayland_toplevel_set_transient_for_exported(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// UnexportHandle destroys the handle that was obtained with
// gdk_wayland_toplevel_export_handle().
//
// It is an error to call this function on a surface that does not have a
// handle.
//
// Note that this API depends on an unstable Wayland protocol, and thus may
// require changes in the future.
func (t waylandToplevel) UnexportHandle(t WaylandToplevel) {
	var arg0 *C.GdkToplevel

	arg0 = (*C.GdkToplevel)(unsafe.Pointer(t.Native()))

	C.gdk_wayland_toplevel_unexport_handle(arg0)
}
