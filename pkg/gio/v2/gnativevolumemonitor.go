// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_native_volume_monitor_get_type()), F: marshalNativeVolumeMonitor},
	})
}

// NativeVolumeMonitor:
type NativeVolumeMonitor interface {
	VolumeMonitor
}

// nativeVolumeMonitor implements the NativeVolumeMonitor class.
type nativeVolumeMonitor struct {
	VolumeMonitor
}

// WrapNativeVolumeMonitor wraps a GObject to the right type. It is
// primarily used internally.
func WrapNativeVolumeMonitor(obj *externglib.Object) NativeVolumeMonitor {
	return nativeVolumeMonitor{
		VolumeMonitor: WrapVolumeMonitor(obj),
	}
}

func marshalNativeVolumeMonitor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNativeVolumeMonitor(obj), nil
}