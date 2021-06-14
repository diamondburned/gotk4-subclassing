// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4 gtk4-wayland
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/wayland/gdkwayland.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_wayland_device_get_type()), F: marshalWaylandDevice},
	})
}

// WaylandDevice: the Wayland implementation of `GdkDevice`.
//
// Beyond the regular [class@Gdk.Device] API, the Wayland implementation
// provides access to Wayland objects such as the `wl_seat` with
// [method@GdkWayland.WaylandDevice.get_wl_seat], the `wl_keyboard` with
// [method@GdkWayland.WaylandDevice.get_wl_keyboard] and the `wl_pointer` with
// [method@GdkWayland.WaylandDevice.get_wl_pointer].
type WaylandDevice interface {
	gdk.Device

	// NodePath returns the `/dev/input/event*` path of this device.
	//
	// For `GdkDevice`s that possibly coalesce multiple hardware devices (eg.
	// mouse, keyboard, touch,...), this function will return nil.
	//
	// This is most notably implemented for devices of type GDK_SOURCE_PEN,
	// GDK_SOURCE_TABLET_PAD.
	NodePath() string
}

// waylandDevice implements the WaylandDevice class.
type waylandDevice struct {
	gdk.Device
}

var _ WaylandDevice = (*waylandDevice)(nil)

// WrapWaylandDevice wraps a GObject to the right type. It is
// primarily used internally.
func WrapWaylandDevice(obj *externglib.Object) WaylandDevice {
	return waylandDevice{
		gdk.Device: gdk.WrapDevice(obj),
	}
}

func marshalWaylandDevice(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandDevice(obj), nil
}

// NodePath returns the `/dev/input/event*` path of this device.
//
// For `GdkDevice`s that possibly coalesce multiple hardware devices (eg.
// mouse, keyboard, touch,...), this function will return nil.
//
// This is most notably implemented for devices of type GDK_SOURCE_PEN,
// GDK_SOURCE_TABLET_PAD.
func (d waylandDevice) NodePath() string {
	var _arg0 *C.GdkDevice // out

	_arg0 = (*C.GdkDevice)(unsafe.Pointer(d.Native()))

	var _cret *C.char // in

	_cret = C.gdk_wayland_device_get_node_path(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}
