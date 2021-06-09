// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_drag_get_type()), F: marshalGestureDrag},
	})
}

// GestureDrag is a Gesture implementation that recognizes drag operations. The
// drag operation itself can be tracked throughout the GestureDrag::drag-begin,
// GestureDrag::drag-update and GestureDrag::drag-end signals, or the relevant
// coordinates be extracted through gtk_gesture_drag_get_offset() and
// gtk_gesture_drag_get_start_point().
type GestureDrag interface {
	GestureSingle

	// Offset: if the @gesture is active, this function returns true and fills
	// in @x and @y with the coordinates of the current point, as an offset to
	// the starting drag point.
	Offset() (x float64, y float64, ok bool)
	// StartPoint: if the @gesture is active, this function returns true and
	// fills in @x and @y with the drag start coordinates, in window-relative
	// coordinates.
	StartPoint() (x float64, y float64, ok bool)
}

// gestureDrag implements the GestureDrag interface.
type gestureDrag struct {
	GestureSingle
}

var _ GestureDrag = (*gestureDrag)(nil)

// WrapGestureDrag wraps a GObject to the right type. It is
// primarily used internally.
func WrapGestureDrag(obj *externglib.Object) GestureDrag {
	return GestureDrag{
		GestureSingle: WrapGestureSingle(obj),
	}
}

func marshalGestureDrag(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGestureDrag(obj), nil
}

// NewGestureDrag constructs a class GestureDrag.
func NewGestureDrag() GestureDrag {
	var cret C.GtkGestureDrag

	cret = C.gtk_gesture_drag_new()

	var gestureDrag GestureDrag

	gestureDrag = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(GestureDrag)

	return gestureDrag
}

// Offset: if the @gesture is active, this function returns true and fills
// in @x and @y with the coordinates of the current point, as an offset to
// the starting drag point.
func (g gestureDrag) Offset() (x float64, y float64, ok bool) {
	var arg0 *C.GtkGestureDrag

	arg0 = (*C.GtkGestureDrag)(unsafe.Pointer(g.Native()))

	var arg1 C.double
	var arg2 C.double
	var cret C.gboolean

	cret = C.gtk_gesture_drag_get_offset(arg0, &arg1, &arg2)

	var x float64
	var y float64
	var ok bool

	x = (float64)(arg1)
	y = (float64)(arg2)
	if cret {
		ok = true
	}

	return x, y, ok
}

// StartPoint: if the @gesture is active, this function returns true and
// fills in @x and @y with the drag start coordinates, in window-relative
// coordinates.
func (g gestureDrag) StartPoint() (x float64, y float64, ok bool) {
	var arg0 *C.GtkGestureDrag

	arg0 = (*C.GtkGestureDrag)(unsafe.Pointer(g.Native()))

	var arg1 C.double
	var arg2 C.double
	var cret C.gboolean

	cret = C.gtk_gesture_drag_get_start_point(arg0, &arg1, &arg2)

	var x float64
	var y float64
	var ok bool

	x = (float64)(arg1)
	y = (float64)(arg2)
	if cret {
		ok = true
	}

	return x, y, ok
}
