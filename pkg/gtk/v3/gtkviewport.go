// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
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
		{T: externglib.Type(C.gtk_viewport_get_type()), F: marshalViewport},
	})
}

// Viewport: the Viewport widget acts as an adaptor class, implementing
// scrollability for child widgets that lack their own scrolling capabilities.
// Use GtkViewport to scroll child widgets such as Grid, Box, and so on.
//
// If a widget has native scrolling abilities, such as TextView, TreeView or
// IconView, it can be added to a ScrolledWindow with gtk_container_add(). If a
// widget does not, you must first add the widget to a Viewport, then add the
// viewport to the scrolled window. gtk_container_add() does this automatically
// if a child that does not implement Scrollable is added to a ScrolledWindow,
// so you can ignore the presence of the viewport.
//
// The GtkViewport will start scrolling content only if allocated less than the
// child widget’s minimum size in a given orientation.
//
//
// CSS nodes
//
// GtkViewport has a single CSS node with name viewport.
type Viewport interface {
	gextras.Objector

	// BinWindow gets the bin window of the Viewport.
	BinWindow() *gdk.WindowClass
	// HAdjustment returns the horizontal adjustment of the viewport.
	//
	// Deprecated: since version 3.0.
	HAdjustment() *AdjustmentClass
	// ShadowType gets the shadow type of the Viewport. See
	// gtk_viewport_set_shadow_type().
	ShadowType() ShadowType
	// VAdjustment returns the vertical adjustment of the viewport.
	//
	// Deprecated: since version 3.0.
	VAdjustment() *AdjustmentClass
	// ViewWindow gets the view window of the Viewport.
	ViewWindow() *gdk.WindowClass
	// SetHAdjustment sets the horizontal adjustment of the viewport.
	//
	// Deprecated: since version 3.0.
	SetHAdjustment(adjustment Adjustment)
	// SetVAdjustment sets the vertical adjustment of the viewport.
	//
	// Deprecated: since version 3.0.
	SetVAdjustment(adjustment Adjustment)
}

// ViewportClass implements the Viewport interface.
type ViewportClass struct {
	*externglib.Object
	BinClass
	BuildableInterface
	ScrollableInterface
}

var _ Viewport = (*ViewportClass)(nil)

func wrapViewport(obj *externglib.Object) Viewport {
	return &ViewportClass{
		Object: obj,
		BinClass: BinClass{
			Object: obj,
			ContainerClass: ContainerClass{
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
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		ScrollableInterface: ScrollableInterface{
			Object: obj,
		},
	}
}

func marshalViewport(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapViewport(obj), nil
}

// NewViewport creates a new Viewport with the given adjustments, or with
// default adjustments if none are given.
func NewViewport(hadjustment Adjustment, vadjustment Adjustment) *ViewportClass {
	var _arg1 *C.GtkAdjustment // out
	var _arg2 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hadjustment.Native()))
	_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(vadjustment.Native()))

	_cret = C.gtk_viewport_new(_arg1, _arg2)

	var _viewport *ViewportClass // out

	_viewport = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ViewportClass)

	return _viewport
}

// BinWindow gets the bin window of the Viewport.
func (v *ViewportClass) BinWindow() *gdk.WindowClass {
	var _arg0 *C.GtkViewport // out
	var _cret *C.GdkWindow   // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	_cret = C.gtk_viewport_get_bin_window(_arg0)

	var _window *gdk.WindowClass // out

	_window = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gdk.WindowClass)

	return _window
}

// HAdjustment returns the horizontal adjustment of the viewport.
//
// Deprecated: since version 3.0.
func (v *ViewportClass) HAdjustment() *AdjustmentClass {
	var _arg0 *C.GtkViewport   // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	_cret = C.gtk_viewport_get_hadjustment(_arg0)

	var _adjustment *AdjustmentClass // out

	_adjustment = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*AdjustmentClass)

	return _adjustment
}

// ShadowType gets the shadow type of the Viewport. See
// gtk_viewport_set_shadow_type().
func (v *ViewportClass) ShadowType() ShadowType {
	var _arg0 *C.GtkViewport  // out
	var _cret C.GtkShadowType // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	_cret = C.gtk_viewport_get_shadow_type(_arg0)

	var _shadowType ShadowType // out

	_shadowType = (ShadowType)(_cret)

	return _shadowType
}

// VAdjustment returns the vertical adjustment of the viewport.
//
// Deprecated: since version 3.0.
func (v *ViewportClass) VAdjustment() *AdjustmentClass {
	var _arg0 *C.GtkViewport   // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	_cret = C.gtk_viewport_get_vadjustment(_arg0)

	var _adjustment *AdjustmentClass // out

	_adjustment = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*AdjustmentClass)

	return _adjustment
}

// ViewWindow gets the view window of the Viewport.
func (v *ViewportClass) ViewWindow() *gdk.WindowClass {
	var _arg0 *C.GtkViewport // out
	var _cret *C.GdkWindow   // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))

	_cret = C.gtk_viewport_get_view_window(_arg0)

	var _window *gdk.WindowClass // out

	_window = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gdk.WindowClass)

	return _window
}

// SetHAdjustment sets the horizontal adjustment of the viewport.
//
// Deprecated: since version 3.0.
func (v *ViewportClass) SetHAdjustment(adjustment Adjustment) {
	var _arg0 *C.GtkViewport   // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_viewport_set_hadjustment(_arg0, _arg1)
}

// SetVAdjustment sets the vertical adjustment of the viewport.
//
// Deprecated: since version 3.0.
func (v *ViewportClass) SetVAdjustment(adjustment Adjustment) {
	var _arg0 *C.GtkViewport   // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(v.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_viewport_set_vadjustment(_arg0, _arg1)
}
