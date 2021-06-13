// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_toplevel_layout_get_type()), F: marshalToplevelLayout},
	})
}

// ToplevelLayout: the `GdkToplevelLayout` struct contains information that is
// necessary to present a sovereign window on screen.
//
// The `GdkToplevelLayout` struct is necessary for using
// [method@Gdk.Toplevel.present].
//
// Toplevel surfaces are sovereign windows that can be presented to the user in
// various states (maximized, on all workspaces, etc).
type ToplevelLayout struct {
	native C.GdkToplevelLayout
}

// WrapToplevelLayout wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapToplevelLayout(ptr unsafe.Pointer) *ToplevelLayout {
	if ptr == nil {
		return nil
	}

	return (*ToplevelLayout)(ptr)
}

func marshalToplevelLayout(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapToplevelLayout(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *ToplevelLayout) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// Equal: check whether @layout and @other has identical layout properties.
func (l *ToplevelLayout) Equal(other *ToplevelLayout) bool {
	var _arg0 *C.GdkToplevelLayout // out
	var _arg1 *C.GdkToplevelLayout // out

	_arg0 = (*C.GdkToplevelLayout)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.GdkToplevelLayout)(unsafe.Pointer(other.Native()))

	var _cret C.gboolean // in

	_cret = C.gdk_toplevel_layout_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Fullscreen: if the layout specifies whether to the toplevel should go
// fullscreen, the value pointed to by @fullscreen is set to true if it should
// go fullscreen, or false, if it should go unfullscreen.
func (l *ToplevelLayout) Fullscreen() (fullscreen bool, ok bool) {
	var _arg0 *C.GdkToplevelLayout // out

	_arg0 = (*C.GdkToplevelLayout)(unsafe.Pointer(l.Native()))

	var _arg1 C.gboolean // in
	var _cret C.gboolean // in

	_cret = C.gdk_toplevel_layout_get_fullscreen(_arg0, &_arg1)

	var _fullscreen bool // out
	var _ok bool         // out

	if _arg1 {
		_fullscreen = true
	}
	if _cret {
		_ok = true
	}

	return _fullscreen, _ok
}

// Maximized: if the layout specifies whether to the toplevel should go
// maximized, the value pointed to by @maximized is set to true if it should go
// fullscreen, or false, if it should go unmaximized.
func (l *ToplevelLayout) Maximized() (maximized bool, ok bool) {
	var _arg0 *C.GdkToplevelLayout // out

	_arg0 = (*C.GdkToplevelLayout)(unsafe.Pointer(l.Native()))

	var _arg1 C.gboolean // in
	var _cret C.gboolean // in

	_cret = C.gdk_toplevel_layout_get_maximized(_arg0, &_arg1)

	var _maximized bool // out
	var _ok bool        // out

	if _arg1 {
		_maximized = true
	}
	if _cret {
		_ok = true
	}

	return _maximized, _ok
}

// Resizable returns whether the layout should allow the user to resize the
// surface.
func (l *ToplevelLayout) Resizable() bool {
	var _arg0 *C.GdkToplevelLayout // out

	_arg0 = (*C.GdkToplevelLayout)(unsafe.Pointer(l.Native()))

	var _cret C.gboolean // in

	_cret = C.gdk_toplevel_layout_get_resizable(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// SetFullscreen sets whether the layout should cause the surface to be
// fullscreen when presented.
func (l *ToplevelLayout) SetFullscreen(fullscreen bool, monitor Monitor) {
	var _arg0 *C.GdkToplevelLayout // out
	var _arg1 C.gboolean           // out
	var _arg2 *C.GdkMonitor        // out

	_arg0 = (*C.GdkToplevelLayout)(unsafe.Pointer(l.Native()))
	if fullscreen {
		_arg1 = C.gboolean(1)
	}
	_arg2 = (*C.GdkMonitor)(unsafe.Pointer(monitor.Native()))

	C.gdk_toplevel_layout_set_fullscreen(_arg0, _arg1, _arg2)
}

// SetMaximized sets whether the layout should cause the surface to be maximized
// when presented.
func (l *ToplevelLayout) SetMaximized(maximized bool) {
	var _arg0 *C.GdkToplevelLayout // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GdkToplevelLayout)(unsafe.Pointer(l.Native()))
	if maximized {
		_arg1 = C.gboolean(1)
	}

	C.gdk_toplevel_layout_set_maximized(_arg0, _arg1)
}

// SetResizable sets whether the layout should allow the user to resize the
// surface after it has been presented.
func (l *ToplevelLayout) SetResizable(resizable bool) {
	var _arg0 *C.GdkToplevelLayout // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GdkToplevelLayout)(unsafe.Pointer(l.Native()))
	if resizable {
		_arg1 = C.gboolean(1)
	}

	C.gdk_toplevel_layout_set_resizable(_arg0, _arg1)
}

// Unref decreases the reference count of @layout.
func (l *ToplevelLayout) Unref() {
	var _arg0 *C.GdkToplevelLayout // out

	_arg0 = (*C.GdkToplevelLayout)(unsafe.Pointer(l.Native()))

	C.gdk_toplevel_layout_unref(_arg0)
}
