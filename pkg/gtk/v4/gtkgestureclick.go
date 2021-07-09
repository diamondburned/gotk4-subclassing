// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_click_get_type()), F: marshalGestureClick},
	})
}

// GestureClick: `GtkGestureClick` is a `GtkGesture` implementation for clicks.
//
// It is able to recognize multiple clicks on a nearby zone, which can be
// listened for through the [signal@Gtk.GestureClick::pressed] signal. Whenever
// time or distance between clicks exceed the GTK defaults,
// [signal@Gtk.GestureClick::stopped] is emitted, and the click counter is
// reset.
type GestureClick interface {
	gextras.Objector

	privateGestureClickClass()
}

// GestureClickClass implements the GestureClick interface.
type GestureClickClass struct {
	GestureSingleClass
}

var _ GestureClick = (*GestureClickClass)(nil)

func wrapGestureClick(obj *externglib.Object) GestureClick {
	return &GestureClickClass{
		GestureSingleClass: GestureSingleClass{
			GestureClass: GestureClass{
				EventControllerClass: EventControllerClass{
					Object: obj,
				},
			},
		},
	}
}

func marshalGestureClick(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapGestureClick(obj), nil
}

// NewGestureClick returns a newly created `GtkGesture` that recognizes single
// and multiple presses.
func NewGestureClick() *GestureClickClass {
	var _cret *C.GtkGesture // in

	_cret = C.gtk_gesture_click_new()

	var _gestureClick *GestureClickClass // out

	_gestureClick = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*GestureClickClass)

	return _gestureClick
}

func (*GestureClickClass) privateGestureClickClass() {}
