// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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
	gextras.Objector

	// Alignment gets the X and Y alignment of the widget within its allocation.
	// See gtk_misc_set_alignment().
	//
	// Deprecated: since version 3.14.
	Alignment() (xalign float32, yalign float32)
	// Padding gets the padding in the X and Y directions of the widget. See
	// gtk_misc_set_padding().
	//
	// Deprecated: since version 3.14.
	Padding() (xpad int, ypad int)
	// SetAlignment sets the alignment of the widget.
	//
	// Deprecated: since version 3.14.
	SetAlignment(xalign float32, yalign float32)
	// SetPadding sets the amount of space to add around the widget.
	//
	// Deprecated: since version 3.14.
	SetPadding(xpad int, ypad int)
}

// MiscClass implements the Misc interface.
type MiscClass struct {
	*externglib.Object
	WidgetClass
	BuildableInterface
}

var _ Misc = (*MiscClass)(nil)

func wrapMisc(obj *externglib.Object) Misc {
	return &MiscClass{
		Object: obj,
		WidgetClass: WidgetClass{
			InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
	}
}

func marshalMisc(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMisc(obj), nil
}

// Alignment gets the X and Y alignment of the widget within its allocation. See
// gtk_misc_set_alignment().
//
// Deprecated: since version 3.14.
func (m *MiscClass) Alignment() (xalign float32, yalign float32) {
	var _arg0 *C.GtkMisc // out
	var _arg1 C.gfloat   // in
	var _arg2 C.gfloat   // in

	_arg0 = (*C.GtkMisc)(unsafe.Pointer(m.Native()))

	C.gtk_misc_get_alignment(_arg0, &_arg1, &_arg2)

	var _xalign float32 // out
	var _yalign float32 // out

	_xalign = float32(_arg1)
	_yalign = float32(_arg2)

	return _xalign, _yalign
}

// Padding gets the padding in the X and Y directions of the widget. See
// gtk_misc_set_padding().
//
// Deprecated: since version 3.14.
func (m *MiscClass) Padding() (xpad int, ypad int) {
	var _arg0 *C.GtkMisc // out
	var _arg1 C.gint     // in
	var _arg2 C.gint     // in

	_arg0 = (*C.GtkMisc)(unsafe.Pointer(m.Native()))

	C.gtk_misc_get_padding(_arg0, &_arg1, &_arg2)

	var _xpad int // out
	var _ypad int // out

	_xpad = int(_arg1)
	_ypad = int(_arg2)

	return _xpad, _ypad
}

// SetAlignment sets the alignment of the widget.
//
// Deprecated: since version 3.14.
func (m *MiscClass) SetAlignment(xalign float32, yalign float32) {
	var _arg0 *C.GtkMisc // out
	var _arg1 C.gfloat   // out
	var _arg2 C.gfloat   // out

	_arg0 = (*C.GtkMisc)(unsafe.Pointer(m.Native()))
	_arg1 = C.gfloat(xalign)
	_arg2 = C.gfloat(yalign)

	C.gtk_misc_set_alignment(_arg0, _arg1, _arg2)
}

// SetPadding sets the amount of space to add around the widget.
//
// Deprecated: since version 3.14.
func (m *MiscClass) SetPadding(xpad int, ypad int) {
	var _arg0 *C.GtkMisc // out
	var _arg1 C.gint     // out
	var _arg2 C.gint     // out

	_arg0 = (*C.GtkMisc)(unsafe.Pointer(m.Native()))
	_arg1 = C.gint(xpad)
	_arg2 = C.gint(ypad)

	C.gtk_misc_set_padding(_arg0, _arg1, _arg2)
}
