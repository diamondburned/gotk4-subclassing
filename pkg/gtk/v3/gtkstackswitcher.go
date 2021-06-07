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
		{T: externglib.Type(C.gtk_stack_switcher_get_type()), F: marshalStackSwitcher},
	})
}

// StackSwitcher: the GtkStackSwitcher widget acts as a controller for a Stack;
// it shows a row of buttons to switch between the various pages of the
// associated stack widget.
//
// All the content for the buttons comes from the child properties of the Stack;
// the button visibility in a StackSwitcher widget is controlled by the
// visibility of the child in the Stack.
//
// It is possible to associate multiple StackSwitcher widgets with the same
// Stack widget.
//
// The GtkStackSwitcher widget was added in 3.10.
//
//
// CSS nodes
//
// GtkStackSwitcher has a single CSS node named stackswitcher and style class
// .stack-switcher.
//
// When circumstances require it, GtkStackSwitcher adds the .needs-attention
// style class to the widgets representing the stack pages.
type StackSwitcher interface {
	Box
	Buildable
	Orientable

	// Stack retrieves the stack. See gtk_stack_switcher_set_stack().
	Stack(s StackSwitcher)
	// SetStack sets the stack to control.
	SetStack(s StackSwitcher, stack Stack)
}

// stackSwitcher implements the StackSwitcher interface.
type stackSwitcher struct {
	Box
	Buildable
	Orientable
}

var _ StackSwitcher = (*stackSwitcher)(nil)

// WrapStackSwitcher wraps a GObject to the right type. It is
// primarily used internally.
func WrapStackSwitcher(obj *externglib.Object) StackSwitcher {
	return StackSwitcher{
		Box:        WrapBox(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalStackSwitcher(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStackSwitcher(obj), nil
}

// NewStackSwitcher constructs a class StackSwitcher.
func NewStackSwitcher() {
	C.gtk_stack_switcher_new()
}

// Stack retrieves the stack. See gtk_stack_switcher_set_stack().
func (s stackSwitcher) Stack(s StackSwitcher) {
	var arg0 *C.GtkStackSwitcher

	arg0 = (*C.GtkStackSwitcher)(unsafe.Pointer(s.Native()))

	C.gtk_stack_switcher_get_stack(arg0)
}

// SetStack sets the stack to control.
func (s stackSwitcher) SetStack(s StackSwitcher, stack Stack) {
	var arg0 *C.GtkStackSwitcher
	var arg1 *C.GtkStack

	arg0 = (*C.GtkStackSwitcher)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	C.gtk_stack_switcher_set_stack(arg0, arg1)
}
