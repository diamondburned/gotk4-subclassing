// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdkx.h>
import "C"

// X11DeviceManagerLookup returns the Device that wraps the given device ID.
func X11DeviceManagerLookup(deviceManager X11DeviceManagerCore, deviceId int) X11DeviceCore {
	var arg1 *C.GdkDeviceManager
	var arg2 C.gint

	arg1 = (*C.GdkDeviceManager)(unsafe.Pointer(deviceManager.Native()))
	arg2 = C.gint(deviceId)

	var cret *C.GdkDevice

	cret = C.gdk_x11_device_manager_lookup(arg1, arg2)

	var x11DeviceCore X11DeviceCore

	x11DeviceCore = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(X11DeviceCore)

	return x11DeviceCore
}
