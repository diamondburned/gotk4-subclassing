// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_invisible_get_type()), F: marshalInvisible},
	})
}

// Invisible: the Invisible widget is used internally in GTK+, and is probably
// not very useful for application developers.
//
// It is used for reliable pointer grabs and selection handling in the code for
// drag-and-drop.
type Invisible interface {
	Widget
	Buildable

	// SetScreen sets the Screen where the Invisible object will be displayed.
	SetScreen(screen gdk.Screen)
}

// invisible implements the Invisible interface.
type invisible struct {
	Widget
	Buildable
}

var _ Invisible = (*invisible)(nil)

// WrapInvisible wraps a GObject to the right type. It is
// primarily used internally.
func WrapInvisible(obj *externglib.Object) Invisible {
	return Invisible{
		Widget:    WrapWidget(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalInvisible(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapInvisible(obj), nil
}

// SetScreen sets the Screen where the Invisible object will be displayed.
func (i invisible) SetScreen(screen gdk.Screen) {
	var _arg0 *C.GtkInvisible // out
	var _arg1 *C.GdkScreen    // out

	_arg0 = (*C.GtkInvisible)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	C.gtk_invisible_set_screen(_arg0, _arg1)
}
