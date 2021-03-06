// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gsk/v4"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_native_get_type()), F: marshalNativeSurfacer},
	})
}

// NativeSurfaceOverrider contains methods that are overridable.
type NativeSurfaceOverrider interface {
}

// NativeSurface: GtkNative is the interface implemented by all widgets that
// have their own GdkSurface.
//
// The obvious example of a GtkNative is GtkWindow.
//
// Every widget that is not itself a GtkNative is contained in one, and you can
// get it with gtk.Widget.GetNative().
//
// To get the surface of a GtkNative, use gtk.Native.GetSurface(). It is also
// possible to find the GtkNative to which a surface belongs, with
// gtk.Native().GetForSurface.
//
// In addition to a gdk.Surface, a GtkNative also provides a gsk.Renderer for
// rendering on that surface. To get the renderer, use gtk.Native.GetRenderer().
//
// This type has been renamed from Native.
type NativeSurface struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*NativeSurface)(nil)
)

// NativeSurfacer describes NativeSurface's interface methods.
type NativeSurfacer interface {
	externglib.Objector

	// Renderer returns the renderer that is used for this GtkNative.
	Renderer() gsk.Rendererer
	// Surface returns the surface of this GtkNative.
	Surface() gdk.Surfacer
	// SurfaceTransform retrieves the surface transform of self.
	SurfaceTransform() (x float64, y float64)
	// Realize realizes a GtkNative.
	Realize()
	// Unrealize unrealizes a GtkNative.
	Unrealize()
}

var _ NativeSurfacer = (*NativeSurface)(nil)

func ifaceInitNativeSurfacer(gifacePtr, data C.gpointer) {
}

func wrapNativeSurface(obj *externglib.Object) *NativeSurface {
	return &NativeSurface{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Object: obj,
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
		},
	}
}

func marshalNativeSurfacer(p uintptr) (interface{}, error) {
	return wrapNativeSurface(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Renderer returns the renderer that is used for this GtkNative.
//
// The function returns the following values:
//
//    - renderer for self.
//
func (self *NativeSurface) Renderer() gsk.Rendererer {
	var _arg0 *C.GtkNative   // out
	var _cret *C.GskRenderer // in

	_arg0 = (*C.GtkNative)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_native_get_renderer(_arg0)
	runtime.KeepAlive(self)

	var _renderer gsk.Rendererer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gsk.Rendererer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(gsk.Rendererer)
			return ok
		})
		rv, ok := casted.(gsk.Rendererer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gsk.Rendererer")
		}
		_renderer = rv
	}

	return _renderer
}

// Surface returns the surface of this GtkNative.
//
// The function returns the following values:
//
//    - surface of self.
//
func (self *NativeSurface) Surface() gdk.Surfacer {
	var _arg0 *C.GtkNative  // out
	var _cret *C.GdkSurface // in

	_arg0 = (*C.GtkNative)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_native_get_surface(_arg0)
	runtime.KeepAlive(self)

	var _surface gdk.Surfacer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Surfacer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(gdk.Surfacer)
			return ok
		})
		rv, ok := casted.(gdk.Surfacer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Surfacer")
		}
		_surface = rv
	}

	return _surface
}

// SurfaceTransform retrieves the surface transform of self.
//
// This is the translation from self's surface coordinates into self's widget
// coordinates.
//
// The function returns the following values:
//
//    - x: return location for the x coordinate.
//    - y: return location for the y coordinate.
//
func (self *NativeSurface) SurfaceTransform() (x float64, y float64) {
	var _arg0 *C.GtkNative // out
	var _arg1 C.double     // in
	var _arg2 C.double     // in

	_arg0 = (*C.GtkNative)(unsafe.Pointer(self.Native()))

	C.gtk_native_get_surface_transform(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(self)

	var _x float64 // out
	var _y float64 // out

	_x = float64(_arg1)
	_y = float64(_arg2)

	return _x, _y
}

// Realize realizes a GtkNative.
//
// This should only be used by subclasses.
func (self *NativeSurface) Realize() {
	var _arg0 *C.GtkNative // out

	_arg0 = (*C.GtkNative)(unsafe.Pointer(self.Native()))

	C.gtk_native_realize(_arg0)
	runtime.KeepAlive(self)
}

// Unrealize unrealizes a GtkNative.
//
// This should only be used by subclasses.
func (self *NativeSurface) Unrealize() {
	var _arg0 *C.GtkNative // out

	_arg0 = (*C.GtkNative)(unsafe.Pointer(self.Native()))

	C.gtk_native_unrealize(_arg0)
	runtime.KeepAlive(self)
}

// NativeSurfaceGetForSurface finds the GtkNative associated with the surface.
//
// The function takes the following parameters:
//
//    - surface: GdkSurface.
//
// The function returns the following values:
//
//    - native: GtkNative that is associated with surface.
//
func NativeSurfaceGetForSurface(surface gdk.Surfacer) NativeSurfacer {
	var _arg1 *C.GdkSurface // out
	var _cret *C.GtkNative  // in

	_arg1 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	_cret = C.gtk_native_get_for_surface(_arg1)
	runtime.KeepAlive(surface)

	var _native NativeSurfacer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gtk.NativeSurfacer is nil")
		}

		object := externglib.Take(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(NativeSurfacer)
			return ok
		})
		rv, ok := casted.(NativeSurfacer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gtk.NativeSurfacer")
		}
		_native = rv
	}

	return _native
}
