// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_gl_context_get_type()), F: marshalX11GLContexter},
	})
}

// X11DisplayGetGLXVersion retrieves the version of the GLX implementation.
//
// The function takes the following parameters:
//
//    - display: Display.
//
func X11DisplayGetGLXVersion(display *gdk.Display) (major int, minor int, ok bool) {
	var _arg1 *C.GdkDisplay // out
	var _arg2 C.gint        // in
	var _arg3 C.gint        // in
	var _cret C.gboolean    // in

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	_cret = C.gdk_x11_display_get_glx_version(_arg1, &_arg2, &_arg3)
	runtime.KeepAlive(display)

	var _major int // out
	var _minor int // out
	var _ok bool   // out

	_major = int(_arg2)
	_minor = int(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _major, _minor, _ok
}

type X11GLContext struct {
	gdk.GLContext
}

func wrapX11GLContext(obj *externglib.Object) *X11GLContext {
	return &X11GLContext{
		GLContext: gdk.GLContext{
			Object: obj,
		},
	}
}

func marshalX11GLContexter(p uintptr) (interface{}, error) {
	return wrapX11GLContext(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
