// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_viewport_get_type()), F: marshalViewporter},
	})
}

// Viewport: GtkViewport implements scrollability for widgets that lack their
// own scrolling capabilities.
//
// Use GtkViewport to scroll child widgets such as GtkGrid, GtkBox, and so on.
//
// The GtkViewport will start scrolling content only if allocated less than the
// child widget’s minimum size in a given orientation.
//
//
// CSS nodes
//
// GtkViewport has a single CSS node with name viewport.
//
//
// Accessibility
//
// GtkViewport uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type Viewport struct {
	Widget

	Scrollable
	*externglib.Object
}

func wrapViewport(obj *externglib.Object) *Viewport {
	return &Viewport{
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
		Scrollable: Scrollable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalViewporter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapViewport(obj), nil
}

// NewViewport creates a new GtkViewport.
//
// The new viewport uses the given adjustments, or default adjustments if none
// are given.
func NewViewport(hadjustment *Adjustment, vadjustment *Adjustment) *Viewport {
	var _arg1 *C.GtkAdjustment // out
	var _arg2 *C.GtkAdjustment // out
	var _cret *C.GtkWidget     // in

	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(hadjustment.Native()))
	_arg2 = (*C.GtkAdjustment)(unsafe.Pointer(vadjustment.Native()))

	_cret = C.gtk_viewport_new(_arg1, _arg2)

	var _viewport *Viewport // out

	_viewport = wrapViewport(externglib.Take(unsafe.Pointer(_cret)))

	return _viewport
}

// Native solves the ambiguous selector of this class or interface.
func (viewport *Viewport) Native() uintptr {
	return viewport.Object.Native()
}

// Child gets the child widget of viewport.
func (viewport *Viewport) Child() Widgetter {
	var _arg0 *C.GtkViewport // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(viewport.Native()))

	_cret = C.gtk_viewport_get_child(_arg0)

	var _widget Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// ScrollToFocus gets whether the viewport is scrolling to keep the focused
// child in view.
func (viewport *Viewport) ScrollToFocus() bool {
	var _arg0 *C.GtkViewport // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(viewport.Native()))

	_cret = C.gtk_viewport_get_scroll_to_focus(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetChild sets the child widget of viewport.
func (viewport *Viewport) SetChild(child Widgetter) {
	var _arg0 *C.GtkViewport // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(viewport.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_viewport_set_child(_arg0, _arg1)
}

// SetScrollToFocus sets whether the viewport should automatically scroll to
// keep the focused child in view.
func (viewport *Viewport) SetScrollToFocus(scrollToFocus bool) {
	var _arg0 *C.GtkViewport // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkViewport)(unsafe.Pointer(viewport.Native()))
	if scrollToFocus {
		_arg1 = C.TRUE
	}

	C.gtk_viewport_set_scroll_to_focus(_arg0, _arg1)
}
