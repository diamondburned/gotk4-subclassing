// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_misc_get_type()), F: marshalMisc},
	})
}

// Misc: the Misc widget is an abstract widget which is not useful itself, but
// is used to derive subclasses which have alignment and padding attributes.
//
// The horizontal and vertical padding attributes allows extra space to be added
// around the widget.
//
// The horizontal and vertical alignment attributes enable the widget to be
// positioned within its allocated area. Note that if the widget is added to a
// container in such a way that it expands automatically to fill its allocated
// area, the alignment settings will not alter the widget's position.
//
// Note that the desired effect can in most cases be achieved by using the
// Widget:halign, Widget:valign and Widget:margin properties on the child
// widget, so GtkMisc should not be used in new code. To reflect this fact, all
// Misc API has been deprecated.
type Misc interface {
	Widget
	Buildable

	// Alignment gets the X and Y alignment of the widget within its allocation.
	// See gtk_misc_set_alignment().
	Alignment() (xalign float32, yalign float32)
	// Padding gets the padding in the X and Y directions of the widget. See
	// gtk_misc_set_padding().
	Padding() (xpad int, ypad int)
	// SetAlignment sets the alignment of the widget.
	SetAlignment(xalign float32, yalign float32)
	// SetPadding sets the amount of space to add around the widget.
	SetPadding(xpad int, ypad int)
}

// misc implements the Misc interface.
type misc struct {
	Widget
	Buildable
}

var _ Misc = (*misc)(nil)

// WrapMisc wraps a GObject to the right type. It is
// primarily used internally.
func WrapMisc(obj *externglib.Object) Misc {
	return Misc{
		Widget:    WrapWidget(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalMisc(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMisc(obj), nil
}

// Alignment gets the X and Y alignment of the widget within its allocation.
// See gtk_misc_set_alignment().
func (m misc) Alignment() (xalign float32, yalign float32) {
	var arg0 *C.GtkMisc

	arg0 = (*C.GtkMisc)(unsafe.Pointer(m.Native()))

	var arg1 C.gfloat
	var arg2 C.gfloat

	C.gtk_misc_get_alignment(arg0, &arg1, &arg2)

	var xalign float32
	var yalign float32

	xalign = (float32)(arg1)
	yalign = (float32)(arg2)

	return xalign, yalign
}

// Padding gets the padding in the X and Y directions of the widget. See
// gtk_misc_set_padding().
func (m misc) Padding() (xpad int, ypad int) {
	var arg0 *C.GtkMisc

	arg0 = (*C.GtkMisc)(unsafe.Pointer(m.Native()))

	var arg1 C.gint
	var arg2 C.gint

	C.gtk_misc_get_padding(arg0, &arg1, &arg2)

	var xpad int
	var ypad int

	xpad = (int)(arg1)
	ypad = (int)(arg2)

	return xpad, ypad
}

// SetAlignment sets the alignment of the widget.
func (m misc) SetAlignment(xalign float32, yalign float32) {
	var arg0 *C.GtkMisc
	var arg1 C.gfloat
	var arg2 C.gfloat

	arg0 = (*C.GtkMisc)(unsafe.Pointer(m.Native()))
	arg1 = C.gfloat(xalign)
	arg2 = C.gfloat(yalign)

	C.gtk_misc_set_alignment(arg0, arg1, arg2)
}

// SetPadding sets the amount of space to add around the widget.
func (m misc) SetPadding(xpad int, ypad int) {
	var arg0 *C.GtkMisc
	var arg1 C.gint
	var arg2 C.gint

	arg0 = (*C.GtkMisc)(unsafe.Pointer(m.Native()))
	arg1 = C.gint(xpad)
	arg2 = C.gint(ypad)

	C.gtk_misc_set_padding(arg0, arg1, arg2)
}
