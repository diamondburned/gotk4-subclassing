// Code generated by girgen. DO NOT EDIT.

package gdkx11

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdkx.h>
import "C"

// X11DeviceGetID returns the device ID as seen by XInput2.
//
// > If gdk_disable_multidevice() has been called, this function > will
// respectively return 2/3 for the core pointer and keyboard, > (matching the
// IDs for the Virtual Core Pointer and Keyboard in > XInput 2), but calling
// this function on any slave devices (i.e. > those managed via XInput 1.x),
// will return 0.
func X11DeviceGetID(device X11DeviceCore) int {
	var _arg1 *C.GdkDevice // out

	_arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	var _cret C.gint // in

	_cret = C.gdk_x11_device_get_id(_arg1)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}
