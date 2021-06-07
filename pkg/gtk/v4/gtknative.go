// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_native_get_type()), F: marshalNative},
	})
}

// NativeGetForSurface finds the GtkNative associated with the surface.
func NativeGetForSurface(surface gdk.Surface) {
	var arg1 *C.GdkSurface

	arg1 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	C.gtk_native_get_for_surface(arg1)
}

// Native is the interface implemented by all widgets that can provide a
// GdkSurface for widgets to render on.
//
// The obvious example of a Native is Window.
type Native interface {
	Widget

	// Renderer returns the renderer that is used for this Native.
	Renderer(s Native)
	// Surface returns the surface of this Native.
	Surface(s Native)
	// SurfaceTransform retrieves the surface transform of @self. This is the
	// translation from @self's surface coordinates into @self's widget
	// coordinates.
	SurfaceTransform(s Native) (x float64, y float64)
	// Realize realizes a Native.
	Realize(s Native)
	// Unrealize unrealizes a Native.
	Unrealize(s Native)
}

// native implements the Native interface.
type native struct {
	Widget
}

var _ Native = (*native)(nil)

// WrapNative wraps a GObject to a type that implements interface
// Native. It is primarily used internally.
func WrapNative(obj *externglib.Object) Native {
	return Native{
		Widget: WrapWidget(obj),
	}
}

func marshalNative(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNative(obj), nil
}

// Renderer returns the renderer that is used for this Native.
func (s native) Renderer(s Native) {
	var arg0 *C.GtkNative

	arg0 = (*C.GtkNative)(unsafe.Pointer(s.Native()))

	C.gtk_native_get_renderer(arg0)
}

// Surface returns the surface of this Native.
func (s native) Surface(s Native) {
	var arg0 *C.GtkNative

	arg0 = (*C.GtkNative)(unsafe.Pointer(s.Native()))

	C.gtk_native_get_surface(arg0)
}

// SurfaceTransform retrieves the surface transform of @self. This is the
// translation from @self's surface coordinates into @self's widget
// coordinates.
func (s native) SurfaceTransform(s Native) (x float64, y float64) {
	var arg0 *C.GtkNative

	arg0 = (*C.GtkNative)(unsafe.Pointer(s.Native()))

	var arg1 C.double
	var x float64
	var arg2 C.double
	var y float64

	C.gtk_native_get_surface_transform(arg0, &arg1, &arg2)

	x = float64(&arg1)
	y = float64(&arg2)

	return x, y
}

// Realize realizes a Native.
func (s native) Realize(s Native) {
	var arg0 *C.GtkNative

	arg0 = (*C.GtkNative)(unsafe.Pointer(s.Native()))

	C.gtk_native_realize(arg0)
}

// Unrealize unrealizes a Native.
func (s native) Unrealize(s Native) {
	var arg0 *C.GtkNative

	arg0 = (*C.GtkNative)(unsafe.Pointer(s.Native()))

	C.gtk_native_unrealize(arg0)
}
