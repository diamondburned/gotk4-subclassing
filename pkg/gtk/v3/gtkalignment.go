// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
		{T: externglib.Type(C.gtk_alignment_get_type()), F: marshalAlignment},
	})
}

// Alignment: the Alignment widget controls the alignment and size of its child
// widget. It has four settings: xscale, yscale, xalign, and yalign.
//
// The scale settings are used to specify how much the child widget should
// expand to fill the space allocated to the Alignment. The values can range
// from 0 (meaning the child doesn’t expand at all) to 1 (meaning the child
// expands to fill all of the available space).
//
// The align settings are used to place the child widget within the available
// area. The values range from 0 (top or left) to 1 (bottom or right). Of
// course, if the scale settings are both set to 1, the alignment settings have
// no effect.
//
// GtkAlignment has been deprecated in 3.14 and should not be used in
// newly-written code. The desired effect can be achieved by using the
// Widget:halign, Widget:valign and Widget:margin properties on the child
// widget.
type Alignment interface {
	Bin
	Buildable

	// Padding gets the padding on the different sides of the widget. See
	// gtk_alignment_set_padding ().
	Padding() (paddingTop uint, paddingBottom uint, paddingLeft uint, paddingRight uint)
	// Set sets the Alignment values.
	Set(xalign float32, yalign float32, xscale float32, yscale float32)
	// SetPadding sets the padding on the different sides of the widget. The
	// padding adds blank space to the sides of the widget. For instance, this
	// can be used to indent the child widget towards the right by adding
	// padding on the left.
	SetPadding(paddingTop uint, paddingBottom uint, paddingLeft uint, paddingRight uint)
}

// alignment implements the Alignment interface.
type alignment struct {
	Bin
	Buildable
}

var _ Alignment = (*alignment)(nil)

// WrapAlignment wraps a GObject to the right type. It is
// primarily used internally.
func WrapAlignment(obj *externglib.Object) Alignment {
	return Alignment{
		Bin:       WrapBin(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalAlignment(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAlignment(obj), nil
}

// Padding gets the padding on the different sides of the widget. See
// gtk_alignment_set_padding ().
func (a alignment) Padding() (paddingTop uint, paddingBottom uint, paddingLeft uint, paddingRight uint) {
	var _arg0 *C.GtkAlignment // out

	_arg0 = (*C.GtkAlignment)(unsafe.Pointer(a.Native()))

	var _arg1 C.guint // in
	var _arg2 C.guint // in
	var _arg3 C.guint // in
	var _arg4 C.guint // in

	C.gtk_alignment_get_padding(_arg0, &_arg1, &_arg2, &_arg3, &_arg4)

	var _paddingTop uint    // out
	var _paddingBottom uint // out
	var _paddingLeft uint   // out
	var _paddingRight uint  // out

	_paddingTop = (uint)(_arg1)
	_paddingBottom = (uint)(_arg2)
	_paddingLeft = (uint)(_arg3)
	_paddingRight = (uint)(_arg4)

	return _paddingTop, _paddingBottom, _paddingLeft, _paddingRight
}

// Set sets the Alignment values.
func (a alignment) Set(xalign float32, yalign float32, xscale float32, yscale float32) {
	var _arg0 *C.GtkAlignment // out
	var _arg1 C.gfloat        // out
	var _arg2 C.gfloat        // out
	var _arg3 C.gfloat        // out
	var _arg4 C.gfloat        // out

	_arg0 = (*C.GtkAlignment)(unsafe.Pointer(a.Native()))
	_arg1 = C.gfloat(xalign)
	_arg2 = C.gfloat(yalign)
	_arg3 = C.gfloat(xscale)
	_arg4 = C.gfloat(yscale)

	C.gtk_alignment_set(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// SetPadding sets the padding on the different sides of the widget. The
// padding adds blank space to the sides of the widget. For instance, this
// can be used to indent the child widget towards the right by adding
// padding on the left.
func (a alignment) SetPadding(paddingTop uint, paddingBottom uint, paddingLeft uint, paddingRight uint) {
	var _arg0 *C.GtkAlignment // out
	var _arg1 C.guint         // out
	var _arg2 C.guint         // out
	var _arg3 C.guint         // out
	var _arg4 C.guint         // out

	_arg0 = (*C.GtkAlignment)(unsafe.Pointer(a.Native()))
	_arg1 = C.guint(paddingTop)
	_arg2 = C.guint(paddingBottom)
	_arg3 = C.guint(paddingLeft)
	_arg4 = C.guint(paddingRight)

	C.gtk_alignment_set_padding(_arg0, _arg1, _arg2, _arg3, _arg4)
}
