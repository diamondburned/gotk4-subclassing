// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-x11-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_cursor_get_type()), F: marshalX11Cursor},
	})
}

type X11Cursor interface {
	gdk.Cursor

	// Xcursor returns the X cursor belonging to a Cursor.
	Xcursor(c X11Cursor)
	// Xdisplay returns the display of a Cursor.
	Xdisplay(c X11Cursor)
}

// x11Cursor implements the X11Cursor interface.
type x11Cursor struct {
	gdk.Cursor
}

var _ X11Cursor = (*x11Cursor)(nil)

// WrapX11Cursor wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Cursor(obj *externglib.Object) X11Cursor {
	return X11Cursor{
		gdk.Cursor: gdk.WrapCursor(obj),
	}
}

func marshalX11Cursor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Cursor(obj), nil
}

// Xcursor returns the X cursor belonging to a Cursor.
func (c x11Cursor) Xcursor(c X11Cursor) {
	var arg0 *C.GdkCursor

	arg0 = (*C.GdkCursor)(unsafe.Pointer(c.Native()))

	C.gdk_x11_cursor_get_xcursor(arg0)
}

// Xdisplay returns the display of a Cursor.
func (c x11Cursor) Xdisplay(c X11Cursor) {
	var arg0 *C.GdkCursor

	arg0 = (*C.GdkCursor)(unsafe.Pointer(c.Native()))

	C.gdk_x11_cursor_get_xdisplay(arg0)
}
