// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_device_manager_core_get_type()), F: marshalX11DeviceManagerCore},
	})
}

// X11DeviceManagerCore:
type X11DeviceManagerCore interface {
	gdk.DeviceManager
}

// x11DeviceManagerCore implements the X11DeviceManagerCore class.
type x11DeviceManagerCore struct {
	gdk.DeviceManager
}

// WrapX11DeviceManagerCore wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11DeviceManagerCore(obj *externglib.Object) X11DeviceManagerCore {
	return x11DeviceManagerCore{
		DeviceManager: gdk.WrapDeviceManager(obj),
	}
}

func marshalX11DeviceManagerCore(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11DeviceManagerCore(obj), nil
}