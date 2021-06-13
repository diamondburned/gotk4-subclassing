// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-x11 gtk4 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/x11/gdkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_app_launch_context_get_type()), F: marshalX11AppLaunchContext},
	})
}

type X11AppLaunchContext interface {
	gdk.AppLaunchContext
}

// x11AppLaunchContext implements the X11AppLaunchContext interface.
type x11AppLaunchContext struct {
	gdk.AppLaunchContext
}

var _ X11AppLaunchContext = (*x11AppLaunchContext)(nil)

// WrapX11AppLaunchContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11AppLaunchContext(obj *externglib.Object) X11AppLaunchContext {
	return X11AppLaunchContext{
		gdk.AppLaunchContext: gdk.WrapAppLaunchContext(obj),
	}
}

func marshalX11AppLaunchContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11AppLaunchContext(obj), nil
}
