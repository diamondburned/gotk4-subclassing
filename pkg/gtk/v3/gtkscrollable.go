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
		{T: externglib.Type(C.gtk_scrollable_get_type()), F: marshalScrollable},
	})
}

// ScrollableOverrider contains methods that are overridable .
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ScrollableOverrider interface {
	// Border returns the size of a non-scrolling border around the outside of
	// the scrollable. An example for this would be treeview headers. GTK+ can
	// use this information to display overlayed graphics, like the overshoot
	// indication, at the right position.
	Border() (Border, bool)
}

// Scrollable is an interface that is implemented by widgets with native
// scrolling ability.
//
// To implement this interface you should override the Scrollable:hadjustment
// and Scrollable:vadjustment properties.
//
//
// Creating a scrollable widget
//
// All scrollable widgets should do the following.
//
// - When a parent widget sets the scrollable child widget’s adjustments, the
// widget should populate the adjustments’ Adjustment:lower, Adjustment:upper,
// Adjustment:step-increment, Adjustment:page-increment and Adjustment:page-size
// properties and connect to the Adjustment::value-changed signal.
//
// - Because its preferred size is the size for a fully expanded widget, the
// scrollable widget must be able to cope with underallocations. This means that
// it must accept any value passed to its WidgetClass.size_allocate() function.
//
// - When the parent allocates space to the scrollable child widget, the widget
// should update the adjustments’ properties with new values.
//
// - When any of the adjustments emits the Adjustment::value-changed signal, the
// scrollable widget should scroll its contents.
type Scrollable interface {
	gextras.Objector

	// Border returns the size of a non-scrolling border around the outside of
	// the scrollable. An example for this would be treeview headers. GTK+ can
	// use this information to display overlayed graphics, like the overshoot
	// indication, at the right position.
	Border() (Border, bool)
	// HAdjustment retrieves the Adjustment used for horizontal scrolling.
	HAdjustment() Adjustment
	// HscrollPolicy gets the horizontal ScrollablePolicy.
	HscrollPolicy() ScrollablePolicy
	// VAdjustment retrieves the Adjustment used for vertical scrolling.
	VAdjustment() Adjustment
	// VscrollPolicy gets the vertical ScrollablePolicy.
	VscrollPolicy() ScrollablePolicy
	// SetHAdjustment sets the horizontal adjustment of the Scrollable.
	SetHAdjustment(hadjustment Adjustment)
	// SetHscrollPolicy sets the ScrollablePolicy to determine whether
	// horizontal scrolling should start below the minimum width or below the
	// natural width.
	SetHscrollPolicy(policy ScrollablePolicy)
	// SetVAdjustment sets the vertical adjustment of the Scrollable.
	SetVAdjustment(vadjustment Adjustment)
	// SetVscrollPolicy sets the ScrollablePolicy to determine whether vertical
	// scrolling should start below the minimum height or below the natural
	// height.
	SetVscrollPolicy(policy ScrollablePolicy)
}

// scrollable implements the Scrollable interface.
type scrollable struct {
	*externglib.Object
}

var _ Scrollable = (*scrollable)(nil)

// WrapScrollable wraps a GObject to a type that implements
// interface Scrollable. It is primarily used internally.
func WrapScrollable(obj *externglib.Object) Scrollable {
	return scrollable{obj}
}

func marshalScrollable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapScrollable(obj), nil
}

func (s scrollable) Border() (Border, bool) {
	var _arg0 *C.GtkScrollable // out
	var _arg1 C.GtkBorder      // in
	var _cret C.gboolean       // in

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrollable_get_border(_arg0, &_arg1)

	var _border Border // out
	var _ok bool       // out

	{
		var refTmpIn *C.GtkBorder
		var refTmpOut *Border

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*Border)(unsafe.Pointer(refTmpIn))

		_border = *refTmpOut
	}
	if _cret != 0 {
		_ok = true
	}

	return _border, _ok
}

func (s scrollable) HAdjustment() Adjustment {
	var _arg0 *C.GtkScrollable // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrollable_get_hadjustment(_arg0)

	var _adjustment Adjustment // out

	_adjustment = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Adjustment)

	return _adjustment
}

func (s scrollable) HscrollPolicy() ScrollablePolicy {
	var _arg0 *C.GtkScrollable      // out
	var _cret C.GtkScrollablePolicy // in

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrollable_get_hscroll_policy(_arg0)

	var _scrollablePolicy ScrollablePolicy // out

	_scrollablePolicy = ScrollablePolicy(_cret)

	return _scrollablePolicy
}

func (s scrollable) VAdjustment() Adjustment {
	var _arg0 *C.GtkScrollable // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrollable_get_vadjustment(_arg0)

	var _adjustment Adjustment // out

	_adjustment = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Adjustment)

	return _adjustment
}

func (s scrollable) VscrollPolicy() ScrollablePolicy {
	var _arg0 *C.GtkScrollable      // out
	var _cret C.GtkScrollablePolicy // in

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_scrollable_get_vscroll_policy(_arg0)

	var _scrollablePolicy ScrollablePolicy // out

	_scrollablePolicy = ScrollablePolicy(_cret)

	return _scrollablePolicy
}

func (s scrollable) SetHAdjustment(hadjustment Adjustment) {
	var _arg0 *C.GtkScrollable // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hadjustment.Native()))

	C.gtk_scrollable_set_hadjustment(_arg0, _arg1)
}

func (s scrollable) SetHscrollPolicy(policy ScrollablePolicy) {
	var _arg0 *C.GtkScrollable      // out
	var _arg1 C.GtkScrollablePolicy // out

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkScrollablePolicy(policy)

	C.gtk_scrollable_set_hscroll_policy(_arg0, _arg1)
}

func (s scrollable) SetVAdjustment(vadjustment Adjustment) {
	var _arg0 *C.GtkScrollable // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(vadjustment.Native()))

	C.gtk_scrollable_set_vadjustment(_arg0, _arg1)
}

func (s scrollable) SetVscrollPolicy(policy ScrollablePolicy) {
	var _arg0 *C.GtkScrollable      // out
	var _arg1 C.GtkScrollablePolicy // out

	_arg0 = (*C.GtkScrollable)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkScrollablePolicy(policy)

	C.gtk_scrollable_set_vscroll_policy(_arg0, _arg1)
}
