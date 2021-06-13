// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-x11 gtk4 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/x11/gdkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_screen_get_type()), F: marshalX11Screen},
	})
}

type X11Screen interface {
	gextras.Objector

	// CurrentDesktop returns the current workspace for @screen when running
	// under a window manager that supports multiple workspaces, as described in
	// the Extended Window Manager Hints
	// (http://www.freedesktop.org/Standards/wm-spec) specification.
	CurrentDesktop() uint32
	// NumberOfDesktops returns the number of workspaces for @screen when
	// running under a window manager that supports multiple workspaces, as
	// described in the Extended Window Manager Hints
	// (http://www.freedesktop.org/Standards/wm-spec) specification.
	NumberOfDesktops() uint32
	// ScreenNumber returns the index of a X11Screen.
	ScreenNumber() int
	// WindowManagerName returns the name of the window manager for @screen.
	WindowManagerName() string
	// SupportsNetWmHint: this function is specific to the X11 backend of GDK,
	// and indicates whether the window manager supports a certain hint from the
	// Extended Window Manager Hints
	// (http://www.freedesktop.org/Standards/wm-spec) specification.
	//
	// When using this function, keep in mind that the window manager can change
	// over time; so you shouldn’t use this function in a way that impacts
	// persistent application state. A common bug is that your application can
	// start up before the window manager does when the user logs in, and before
	// the window manager starts gdk_x11_screen_supports_net_wm_hint() will
	// return false for every property. You can monitor the
	// window_manager_changed signal on X11Screen to detect a window manager
	// change.
	SupportsNetWmHint(propertyName string) bool
}

// x11Screen implements the X11Screen interface.
type x11Screen struct {
	gextras.Objector
}

var _ X11Screen = (*x11Screen)(nil)

// WrapX11Screen wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Screen(obj *externglib.Object) X11Screen {
	return X11Screen{
		Objector: obj,
	}
}

func marshalX11Screen(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Screen(obj), nil
}

// CurrentDesktop returns the current workspace for @screen when running
// under a window manager that supports multiple workspaces, as described in
// the Extended Window Manager Hints
// (http://www.freedesktop.org/Standards/wm-spec) specification.
func (s x11Screen) CurrentDesktop() uint32 {
	var _arg0 *C.GdkX11Screen // out

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(s.Native()))

	var _cret C.guint32 // in

	_cret = C.gdk_x11_screen_get_current_desktop(_arg0)

	var _guint32 uint32 // out

	_guint32 = (uint32)(_cret)

	return _guint32
}

// NumberOfDesktops returns the number of workspaces for @screen when
// running under a window manager that supports multiple workspaces, as
// described in the Extended Window Manager Hints
// (http://www.freedesktop.org/Standards/wm-spec) specification.
func (s x11Screen) NumberOfDesktops() uint32 {
	var _arg0 *C.GdkX11Screen // out

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(s.Native()))

	var _cret C.guint32 // in

	_cret = C.gdk_x11_screen_get_number_of_desktops(_arg0)

	var _guint32 uint32 // out

	_guint32 = (uint32)(_cret)

	return _guint32
}

// ScreenNumber returns the index of a X11Screen.
func (s x11Screen) ScreenNumber() int {
	var _arg0 *C.GdkX11Screen // out

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(s.Native()))

	var _cret C.int // in

	_cret = C.gdk_x11_screen_get_screen_number(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// WindowManagerName returns the name of the window manager for @screen.
func (s x11Screen) WindowManagerName() string {
	var _arg0 *C.GdkX11Screen // out

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(s.Native()))

	var _cret *C.char // in

	_cret = C.gdk_x11_screen_get_window_manager_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// SupportsNetWmHint: this function is specific to the X11 backend of GDK,
// and indicates whether the window manager supports a certain hint from the
// Extended Window Manager Hints
// (http://www.freedesktop.org/Standards/wm-spec) specification.
//
// When using this function, keep in mind that the window manager can change
// over time; so you shouldn’t use this function in a way that impacts
// persistent application state. A common bug is that your application can
// start up before the window manager does when the user logs in, and before
// the window manager starts gdk_x11_screen_supports_net_wm_hint() will
// return false for every property. You can monitor the
// window_manager_changed signal on X11Screen to detect a window manager
// change.
func (s x11Screen) SupportsNetWmHint(propertyName string) bool {
	var _arg0 *C.GdkX11Screen // out
	var _arg1 *C.char         // out

	_arg0 = (*C.GdkX11Screen)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean // in

	_cret = C.gdk_x11_screen_supports_net_wm_hint(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}
