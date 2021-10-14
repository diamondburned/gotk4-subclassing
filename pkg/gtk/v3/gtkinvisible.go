// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_invisible_get_type()), F: marshalInvisibler},
	})
}

// Invisible widget is used internally in GTK+, and is probably not very useful
// for application developers.
//
// It is used for reliable pointer grabs and selection handling in the code for
// drag-and-drop.
type Invisible struct {
	Widget
}

func wrapInvisible(obj *externglib.Object) *Invisible {
	return &Invisible{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			ImplementorIface: atk.ImplementorIface{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalInvisibler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapInvisible(obj), nil
}

// NewInvisible creates a new Invisible.
func NewInvisible() *Invisible {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_invisible_new()

	var _invisible *Invisible // out

	_invisible = wrapInvisible(externglib.Take(unsafe.Pointer(_cret)))

	return _invisible
}

// NewInvisibleForScreen creates a new Invisible object for a specified screen.
//
// The function takes the following parameters:
//
//    - screen which identifies on which the new Invisible will be created.
//
func NewInvisibleForScreen(screen *gdk.Screen) *Invisible {
	var _arg1 *C.GdkScreen // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	_cret = C.gtk_invisible_new_for_screen(_arg1)
	runtime.KeepAlive(screen)

	var _invisible *Invisible // out

	_invisible = wrapInvisible(externglib.Take(unsafe.Pointer(_cret)))

	return _invisible
}

// Screen returns the Screen object associated with invisible.
func (invisible *Invisible) Screen() *gdk.Screen {
	var _arg0 *C.GtkInvisible // out
	var _cret *C.GdkScreen    // in

	_arg0 = (*C.GtkInvisible)(unsafe.Pointer(invisible.Native()))

	_cret = C.gtk_invisible_get_screen(_arg0)
	runtime.KeepAlive(invisible)

	var _screen *gdk.Screen // out

	{
		obj := externglib.Take(unsafe.Pointer(_cret))
		_screen = &gdk.Screen{
			Object: obj,
		}
	}

	return _screen
}

// SetScreen sets the Screen where the Invisible object will be displayed.
//
// The function takes the following parameters:
//
//    - screen: Screen.
//
func (invisible *Invisible) SetScreen(screen *gdk.Screen) {
	var _arg0 *C.GtkInvisible // out
	var _arg1 *C.GdkScreen    // out

	_arg0 = (*C.GtkInvisible)(unsafe.Pointer(invisible.Native()))
	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	C.gtk_invisible_set_screen(_arg0, _arg1)
	runtime.KeepAlive(invisible)
	runtime.KeepAlive(screen)
}
