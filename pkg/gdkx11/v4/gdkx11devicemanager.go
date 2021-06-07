// Code generated by girgen. DO NOT EDIT.

package gdkx11

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/x11/gdkx.h>
import "C"

// X11DeviceManagerLookup returns the Device that wraps the given device ID.
func X11DeviceManagerLookup(deviceManager X11DeviceManagerXI2, deviceID int) {
	var arg1 *C.GdkX11DeviceManagerXI2
	var arg2 C.int

	arg1 = (*C.GdkX11DeviceManagerXI2)(unsafe.Pointer(deviceManager.Native()))
	arg2 = C.int(deviceID)

	C.gdk_x11_device_manager_lookup(arg1, arg2)
}
