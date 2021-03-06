// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_fixed_get_type()), F: marshalFixedder},
	})
}

// FixedOverrider contains methods that are overridable.
type FixedOverrider interface {
}

// Fixed widget is a container which can place child widgets at fixed positions
// and with fixed sizes, given in pixels. Fixed performs no automatic layout
// management.
//
// For most applications, you should not use this container! It keeps you from
// having to learn about the other GTK+ containers, but it results in broken
// applications. With Fixed, the following things will result in truncated text,
// overlapping widgets, and other display bugs:
//
// - Themes, which may change widget sizes.
//
// - Fonts other than the one you used to write the app will of course change
// the size of widgets containing text; keep in mind that users may use a larger
// font because of difficulty reading the default, or they may be using a
// different OS that provides different fonts.
//
// - Translation of text into other languages changes its size. Also, display of
// non-English text will use a different font in many cases.
//
// In addition, Fixed does not pay attention to text direction and thus may
// produce unwanted results if your app is run under right-to-left languages
// such as Hebrew or Arabic. That is: normally GTK+ will order containers
// appropriately for the text direction, e.g. to put labels to the right of the
// thing they label when using an RTL language, but it can’t do that with Fixed.
// So if you need to reorder widgets depending on the text direction, you would
// need to manually detect it and adjust child positions accordingly.
//
// Finally, fixed positioning makes it kind of annoying to add/remove GUI
// elements, since you have to reposition all the other elements. This is a
// long-term maintenance problem for your application.
//
// If you know none of these things are an issue for your application, and
// prefer the simplicity of Fixed, by all means use the widget. But you should
// be aware of the tradeoffs.
//
// See also Layout, which shares the ability to perform fixed positioning of
// child widgets and additionally adds custom drawing and scrollability.
type Fixed struct {
	_ [0]func() // equal guard
	Container
}

var (
	_ Containerer = (*Fixed)(nil)
)

func classInitFixedder(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapFixed(obj *externglib.Object) *Fixed {
	return &Fixed{
		Container: Container{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
		},
	}
}

func marshalFixedder(p uintptr) (interface{}, error) {
	return wrapFixed(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewFixed creates a new Fixed.
//
// The function returns the following values:
//
//    - fixed: new Fixed.
//
func NewFixed() *Fixed {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_fixed_new()

	var _fixed *Fixed // out

	_fixed = wrapFixed(externglib.Take(unsafe.Pointer(_cret)))

	return _fixed
}

// Move moves a child of a Fixed container to the given position.
//
// The function takes the following parameters:
//
//    - widget: child widget.
//    - x: horizontal position to move the widget to.
//    - y: vertical position to move the widget to.
//
func (fixed *Fixed) Move(widget Widgetter, x, y int) {
	var _arg0 *C.GtkFixed  // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gint       // out
	var _arg3 C.gint       // out

	_arg0 = (*C.GtkFixed)(unsafe.Pointer(fixed.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)

	C.gtk_fixed_move(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(fixed)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// Put adds a widget to a Fixed container at the given position.
//
// The function takes the following parameters:
//
//    - widget to add.
//    - x: horizontal position to place the widget at.
//    - y: vertical position to place the widget at.
//
func (fixed *Fixed) Put(widget Widgetter, x, y int) {
	var _arg0 *C.GtkFixed  // out
	var _arg1 *C.GtkWidget // out
	var _arg2 C.gint       // out
	var _arg3 C.gint       // out

	_arg0 = (*C.GtkFixed)(unsafe.Pointer(fixed.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.gint(x)
	_arg3 = C.gint(y)

	C.gtk_fixed_put(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(fixed)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
}

// FixedChild: instance of this type is always passed by reference.
type FixedChild struct {
	*fixedChild
}

// fixedChild is the struct that's finalized.
type fixedChild struct {
	native *C.GtkFixedChild
}

func (f *FixedChild) Widget() Widgetter {
	var v Widgetter // out
	{
		objptr := unsafe.Pointer(f.native.widget)
		if objptr == nil {
			panic("object of type gtk.Widgetter is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(Widgetter)
			return ok
		})
		rv, ok := casted.(Widgetter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.Widgetter")
		}
		v = rv
	}
	return v
}

func (f *FixedChild) X() int {
	var v int // out
	v = int(f.native.x)
	return v
}

func (f *FixedChild) Y() int {
	var v int // out
	v = int(f.native.y)
	return v
}
