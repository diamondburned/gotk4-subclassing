// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_orientable_get_type()), F: marshalOrientabler},
	})
}

// Orientable: GtkOrientable interface is implemented by all widgets that can be
// oriented horizontally or vertically.
//
// GtkOrientable is more flexible in that it allows the orientation to be
// changed at runtime, allowing the widgets to “flip”.
type Orientable struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Orientable)(nil)
)

// Orientabler describes Orientable's interface methods.
type Orientabler interface {
	externglib.Objector

	// Orientation retrieves the orientation of the orientable.
	Orientation() Orientation
	// SetOrientation sets the orientation of the orientable.
	SetOrientation(orientation Orientation)
}

var _ Orientabler = (*Orientable)(nil)

func wrapOrientable(obj *externglib.Object) *Orientable {
	return &Orientable{
		Object: obj,
	}
}

func marshalOrientabler(p uintptr) (interface{}, error) {
	return wrapOrientable(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Orientation retrieves the orientation of the orientable.
//
// The function returns the following values:
//
//    - orientation of the orientable.
//
func (orientable *Orientable) Orientation() Orientation {
	var _arg0 *C.GtkOrientable // out
	var _cret C.GtkOrientation // in

	_arg0 = (*C.GtkOrientable)(unsafe.Pointer(orientable.Native()))

	_cret = C.gtk_orientable_get_orientation(_arg0)
	runtime.KeepAlive(orientable)

	var _orientation Orientation // out

	_orientation = Orientation(_cret)

	return _orientation
}

// SetOrientation sets the orientation of the orientable.
//
// The function takes the following parameters:
//
//    - orientation orientable’s new orientation.
//
func (orientable *Orientable) SetOrientation(orientation Orientation) {
	var _arg0 *C.GtkOrientable // out
	var _arg1 C.GtkOrientation // out

	_arg0 = (*C.GtkOrientable)(unsafe.Pointer(orientable.Native()))
	_arg1 = C.GtkOrientation(orientation)

	C.gtk_orientable_set_orientation(_arg0, _arg1)
	runtime.KeepAlive(orientable)
	runtime.KeepAlive(orientation)
}
