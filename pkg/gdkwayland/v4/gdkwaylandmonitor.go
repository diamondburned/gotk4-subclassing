// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-wayland gtk4 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/wayland/gdkwayland.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_wayland_monitor_get_type()), F: marshalWaylandMonitor},
	})
}

// WaylandMonitor: the Wayland implementation of `GdkMonitor`.
//
// Beyond the [class@Gdk.Monitor] API, the Wayland implementation offers access
// to the Wayland `wl_output` object with
// [method@GdkWayland.WaylandMonitor.get_wl_output].
type WaylandMonitor interface {
	gdk.Monitor

	// WlOutput returns the Wayland `wl_output` of a `GdkMonitor`.
	WlOutput() *interface{}
}

// waylandMonitor implements the WaylandMonitor interface.
type waylandMonitor struct {
	gdk.Monitor
}

var _ WaylandMonitor = (*waylandMonitor)(nil)

// WrapWaylandMonitor wraps a GObject to the right type. It is
// primarily used internally.
func WrapWaylandMonitor(obj *externglib.Object) WaylandMonitor {
	return WaylandMonitor{
		gdk.Monitor: gdk.WrapMonitor(obj),
	}
}

func marshalWaylandMonitor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandMonitor(obj), nil
}

// WlOutput returns the Wayland `wl_output` of a `GdkMonitor`.
func (m waylandMonitor) WlOutput() *interface{} {
	var _arg0 *C.GdkMonitor // out

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var _cret *C.wl_output // in

	_cret = C.gdk_wayland_monitor_get_wl_output(_arg0)

	var _gpointer *interface{} // out

	_gpointer = (*interface{})(_cret)

	return _gpointer
}
