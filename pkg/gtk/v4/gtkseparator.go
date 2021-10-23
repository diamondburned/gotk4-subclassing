// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_separator_get_type()), F: marshalSeparatorrer},
	})
}

// Separator: GtkSeparator is a horizontal or vertical separator widget.
//
// !An example GtkSeparator (separators.png)
//
// A GtkSeparator can be used to group the widgets within a window. It displays
// a line with a shadow to make it appear sunken into the interface.
//
//
// CSS nodes
//
// GtkSeparator has a single CSS node with name separator. The node gets one of
// the .horizontal or .vertical style classes.
//
//
// Accessibility
//
// GtkSeparator uses the K_ACCESSIBLE_ROLE_SEPARATOR role.
type Separator struct {
	Widget

	Orientable
	*externglib.Object
}

func wrapSeparator(obj *externglib.Object) *Separator {
	return &Separator{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
		Orientable: Orientable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalSeparatorrer(p uintptr) (interface{}, error) {
	return wrapSeparator(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSeparator creates a new GtkSeparator with the given orientation.
//
// The function takes the following parameters:
//
//    - orientation separator’s orientation.
//
func NewSeparator(orientation Orientation) *Separator {
	var _arg1 C.GtkOrientation // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)

	_cret = C.gtk_separator_new(_arg1)
	runtime.KeepAlive(orientation)

	var _separator *Separator // out

	_separator = wrapSeparator(externglib.Take(unsafe.Pointer(_cret)))

	return _separator
}
