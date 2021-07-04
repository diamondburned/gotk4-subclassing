// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_scrollbar_get_type()), F: marshalScrollbar},
	})
}

// Scrollbar: the Scrollbar widget is a horizontal or vertical scrollbar,
// depending on the value of the Orientable:orientation property.
//
// Its position and movement are controlled by the adjustment that is passed to
// or created by gtk_scrollbar_new(). See Adjustment for more details. The
// Adjustment:value field sets the position of the thumb and must be between
// Adjustment:lower and Adjustment:upper - Adjustment:page-size. The
// Adjustment:page-size represents the size of the visible scrollable area. The
// fields Adjustment:step-increment and Adjustment:page-increment fields are
// added to or subtracted from the Adjustment:value when the user asks to move
// by a step (using e.g. the cursor arrow keys or, if present, the stepper
// buttons) or by a page (using e.g. the Page Down/Up keys).
//
// CSS nodes
//
//    scrollbar[.fine-tune]
//    ╰── contents
//        ├── [button.up]
//        ├── [button.down]
//        ├── trough
//        │   ╰── slider
//        ├── [button.up]
//        ╰── [button.down]
//
// GtkScrollbar has a main CSS node with name scrollbar and a subnode for its
// contents, with subnodes named trough and slider.
//
// The main node gets the style class .fine-tune added when the scrollbar is in
// 'fine-tuning' mode.
//
// If steppers are enabled, they are represented by up to four additional
// subnodes with name button. These get the style classes .up and .down to
// indicate in which direction they are moving.
//
// Other style classes that may be added to scrollbars inside ScrolledWindow
// include the positional classes (.left, .right, .top, .bottom) and style
// classes related to overlay scrolling (.overlay-indicator, .dragging,
// .hovering).
type Scrollbar interface {
	Range
}

// scrollbar implements the Scrollbar class.
type scrollbar struct {
	Range
}

// WrapScrollbar wraps a GObject to the right type. It is
// primarily used internally.
func WrapScrollbar(obj *externglib.Object) Scrollbar {
	return scrollbar{
		Range: WrapRange(obj),
	}
}

func marshalScrollbar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapScrollbar(obj), nil
}

// NewScrollbar:
func NewScrollbar(orientation Orientation, adjustment Adjustment) Scrollbar {
	var _arg1 C.GtkOrientation // out
	var _arg2 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = C.GtkOrientation(orientation)
	_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	_cret = C.gtk_scrollbar_new(_arg1, _arg2)

	var _scrollbar Scrollbar // out

	_scrollbar = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Scrollbar)

	return _scrollbar
}

func (b scrollbar) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b scrollbar) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b scrollbar) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b scrollbar) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b scrollbar) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b scrollbar) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b scrollbar) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (o scrollbar) Orientation() Orientation {
	return WrapOrientable(gextras.InternObject(o)).Orientation()
}

func (o scrollbar) SetOrientation(orientation Orientation) {
	WrapOrientable(gextras.InternObject(o)).SetOrientation(orientation)
}