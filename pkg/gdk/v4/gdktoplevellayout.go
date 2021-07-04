// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
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
type ToplevelLayout C.GdkToplevelLayout

// WrapToplevelLayout wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapToplevelLayout(ptr unsafe.Pointer) *ToplevelLayout {
	return (*ToplevelLayout)(ptr)
}

func marshalToplevelLayout(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*ToplevelLayout)(unsafe.Pointer(b)), nil
}

// NewToplevelLayout constructs a struct ToplevelLayout.
func NewToplevelLayout() *ToplevelLayout {
	var _cret *C.GdkToplevelLayout // in

	_cret = C.gdk_toplevel_layout_new()

	var _toplevelLayout *ToplevelLayout // out

	_toplevelLayout = (*ToplevelLayout)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_toplevelLayout, func(v **ToplevelLayout) {
		C.free(unsafe.Pointer(v))
	})

	return _toplevelLayout
}

// Native returns the underlying C source pointer.
func (t *ToplevelLayout) Native() unsafe.Pointer {
	return unsafe.Pointer(t)
}

// Copy decreases the reference count of @layout.
func (l *ToplevelLayout) Copy() *ToplevelLayout {
	var _arg0 *C.GdkToplevelLayout // out
	var _cret *C.GdkToplevelLayout // in

	_arg0 = (*C.GdkToplevelLayout)(unsafe.Pointer(l.Native()))

	_cret = C.gdk_toplevel_layout_copy(_arg0)

	var _toplevelLayout *ToplevelLayout // out

	_toplevelLayout = (*ToplevelLayout)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_toplevelLayout, func(v **ToplevelLayout) {
		C.free(unsafe.Pointer(v))
	})

	return _toplevelLayout
}

// Equal decreases the reference count of @layout.
func (l *ToplevelLayout) Equal(other *ToplevelLayout) bool {
	var _arg0 *C.GdkToplevelLayout // out
	var _arg1 *C.GdkToplevelLayout // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GdkToplevelLayout)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.GdkToplevelLayout)(unsafe.Pointer(other.Native()))

	_cret = C.gdk_toplevel_layout_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Fullscreen decreases the reference count of @layout.
func (l *ToplevelLayout) Fullscreen() (fullscreen bool, ok bool) {
	var _arg0 *C.GdkToplevelLayout // out
	var _arg1 C.gboolean           // in
	var _cret C.gboolean           // in

	_arg0 = (*C.GdkToplevelLayout)(unsafe.Pointer(l.Native()))

	_cret = C.gdk_toplevel_layout_get_fullscreen(_arg0, &_arg1)

	var _fullscreen bool // out
	var _ok bool         // out

	if _arg1 != 0 {
		_fullscreen = true
	}
	if _cret != 0 {
		_ok = true
	}

	return _fullscreen, _ok
}

// FullscreenMonitor decreases the reference count of @layout.
func (l *ToplevelLayout) FullscreenMonitor() Monitor {
	var _arg0 *C.GdkToplevelLayout // out
	var _cret *C.GdkMonitor        // in

	_arg0 = (*C.GdkToplevelLayout)(unsafe.Pointer(l.Native()))

	_cret = C.gdk_toplevel_layout_get_fullscreen_monitor(_arg0)

	var _monitor Monitor // out

	_monitor = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Monitor)

	return _monitor
}

// Maximized decreases the reference count of @layout.
func (l *ToplevelLayout) Maximized() (maximized bool, ok bool) {
	var _arg0 *C.GdkToplevelLayout // out
	var _arg1 C.gboolean           // in
	var _cret C.gboolean           // in

	_arg0 = (*C.GdkToplevelLayout)(unsafe.Pointer(l.Native()))

	_cret = C.gdk_toplevel_layout_get_maximized(_arg0, &_arg1)

	var _maximized bool // out
	var _ok bool        // out

	if _arg1 != 0 {
		_maximized = true
	}
	if _cret != 0 {
		_ok = true
	}

	return _maximized, _ok
}

// Resizable decreases the reference count of @layout.
func (l *ToplevelLayout) Resizable() bool {
	var _arg0 *C.GdkToplevelLayout // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GdkToplevelLayout)(unsafe.Pointer(l.Native()))

	_cret = C.gdk_toplevel_layout_get_resizable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Ref decreases the reference count of @layout.
func (l *ToplevelLayout) Ref() *ToplevelLayout {
	var _arg0 *C.GdkToplevelLayout // out
	var _cret *C.GdkToplevelLayout // in

	_arg0 = (*C.GdkToplevelLayout)(unsafe.Pointer(l.Native()))

	_cret = C.gdk_toplevel_layout_ref(_arg0)

	var _toplevelLayout *ToplevelLayout // out

	_toplevelLayout = (*ToplevelLayout)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(&_toplevelLayout, func(v **ToplevelLayout) {
		C.free(unsafe.Pointer(v))
	})

	return _toplevelLayout
}

// SetFullscreen decreases the reference count of @layout.
func (l *ToplevelLayout) SetFullscreen(fullscreen bool, monitor Monitor) {
	var _arg0 *C.GdkToplevelLayout // out
	var _arg1 C.gboolean           // out
	var _arg2 *C.GdkMonitor        // out

	_arg0 = (*C.GdkToplevelLayout)(unsafe.Pointer(l.Native()))
	if fullscreen {
		_arg1 = C.TRUE
	}
	_arg2 = (*C.GdkMonitor)(unsafe.Pointer(monitor.Native()))

	C.gdk_toplevel_layout_set_fullscreen(_arg0, _arg1, _arg2)
}

// SetMaximized decreases the reference count of @layout.
func (l *ToplevelLayout) SetMaximized(maximized bool) {
	var _arg0 *C.GdkToplevelLayout // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GdkToplevelLayout)(unsafe.Pointer(l.Native()))
	if maximized {
		_arg1 = C.TRUE
	}

	C.gdk_toplevel_layout_set_maximized(_arg0, _arg1)
}

// SetResizable decreases the reference count of @layout.
func (l *ToplevelLayout) SetResizable(resizable bool) {
	var _arg0 *C.GdkToplevelLayout // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GdkToplevelLayout)(unsafe.Pointer(l.Native()))
	if resizable {
		_arg1 = C.TRUE
	}

	C.gdk_toplevel_layout_set_resizable(_arg0, _arg1)
}

// Unref decreases the reference count of @layout.
func (l *ToplevelLayout) Unref() {
	var _arg0 *C.GdkToplevelLayout // out

	_arg0 = (*C.GdkToplevelLayout)(unsafe.Pointer(l.Native()))

	C.gdk_toplevel_layout_unref(_arg0)
}