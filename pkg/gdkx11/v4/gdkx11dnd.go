// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4 gtk4-x11
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/x11/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_drag_get_type()), F: marshalX11Drag},
	})
}

type X11Drag interface {
	gdk.Drag
}

// x11Drag implements the X11Drag class.
type x11Drag struct {
	gdk.Drag
}

var _ X11Drag = (*x11Drag)(nil)

// WrapX11Drag wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Drag(obj *externglib.Object) X11Drag {
	return x11Drag{
		gdk.Drag: gdk.WrapDrag(obj),
	}
}

func marshalX11Drag(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Drag(obj), nil
}
