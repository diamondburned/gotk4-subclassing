// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_rgba_get_type()), F: marshalRGBA},
	})
}

// RGBA is used to represent a (possibly translucent) color, in a way that is
// compatible with cairo’s notion of color.
type RGBA struct {
	*rgbA
}

// rgbA is the struct that's finalized.
type rgbA struct {
	native *C.GdkRGBA
}

func marshalRGBA(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &RGBA{&rgbA{(*C.GdkRGBA)(unsafe.Pointer(b))}}, nil
}

// Red: intensity of the red channel from 0.0 to 1.0 inclusive
func (r *RGBA) Red() float64 {
	var v float64 // out
	v = float64(r.native.red)
	return v
}

// Green: intensity of the green channel from 0.0 to 1.0 inclusive
func (r *RGBA) Green() float64 {
	var v float64 // out
	v = float64(r.native.green)
	return v
}

// Blue: intensity of the blue channel from 0.0 to 1.0 inclusive
func (r *RGBA) Blue() float64 {
	var v float64 // out
	v = float64(r.native.blue)
	return v
}

// Alpha: opacity of the color from 0.0 for completely translucent to 1.0 for
// opaque
func (r *RGBA) Alpha() float64 {
	var v float64 // out
	v = float64(r.native.alpha)
	return v
}

// Copy makes a copy of a RGBA.
//
// The result must be freed through gdk_rgba_free().
func (rgba *RGBA) Copy() *RGBA {
	var _arg0 *C.GdkRGBA // out
	var _cret *C.GdkRGBA // in

	_arg0 = (*C.GdkRGBA)(gextras.StructNative(unsafe.Pointer(rgba)))

	_cret = C.gdk_rgba_copy(_arg0)
	runtime.KeepAlive(rgba)

	var _rgbA *RGBA // out

	_rgbA = (*RGBA)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_rgbA)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_rgba_free((*C.GdkRGBA)(intern.C))
		},
	)

	return _rgbA
}

// Equal compares two RGBA colors.
func (p1 *RGBA) Equal(p2 *RGBA) bool {
	var _arg0 C.gconstpointer // out
	var _arg1 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg0 = C.gconstpointer(gextras.StructNative(unsafe.Pointer(p1)))
	_arg1 = C.gconstpointer(gextras.StructNative(unsafe.Pointer(p2)))

	_cret = C.gdk_rgba_equal(_arg0, _arg1)
	runtime.KeepAlive(p1)
	runtime.KeepAlive(p2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Hash: hash function suitable for using for a hash table that stores RGBAs.
func (p *RGBA) Hash() uint {
	var _arg0 C.gconstpointer // out
	var _cret C.guint         // in

	_arg0 = C.gconstpointer(gextras.StructNative(unsafe.Pointer(p)))

	_cret = C.gdk_rgba_hash(_arg0)
	runtime.KeepAlive(p)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Parse parses a textual representation of a color, filling in the red, green,
// blue and alpha fields of the rgba RGBA.
//
// The string can be either one of:
//
// - A standard name (Taken from the X11 rgb.txt file).
//
// - A hexadecimal value in the form “\#rgb”, “\#rrggbb”, “\#rrrgggbbb” or
// ”\#rrrrggggbbbb”
//
// - A RGB color in the form “rgb(r,g,b)” (In this case the color will have full
// opacity)
//
// - A RGBA color in the form “rgba(r,g,b,a)”
//
// Where “r”, “g”, “b” and “a” are respectively the red, green, blue and alpha
// color values. In the last two cases, “r”, “g”, and “b” are either integers in
// the range 0 to 255 or percentage values in the range 0% to 100%, and a is a
// floating point value in the range 0 to 1.
func (rgba *RGBA) Parse(spec string) bool {
	var _arg0 *C.GdkRGBA // out
	var _arg1 *C.gchar   // out
	var _cret C.gboolean // in

	_arg0 = (*C.GdkRGBA)(gextras.StructNative(unsafe.Pointer(rgba)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(spec)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_rgba_parse(_arg0, _arg1)
	runtime.KeepAlive(rgba)
	runtime.KeepAlive(spec)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// String returns a textual specification of rgba in the form rgb(r,g,b) or
// rgba(r g,b,a), where “r”, “g”, “b” and “a” represent the red, green, blue and
// alpha values respectively. “r”, “g”, and “b” are represented as integers in
// the range 0 to 255, and “a” is represented as a floating point value in the
// range 0 to 1.
//
// These string forms are string forms that are supported by the CSS3 colors
// module, and can be parsed by gdk_rgba_parse().
//
// Note that this string representation may lose some precision, since “r”, “g”
// and “b” are represented as 8-bit integers. If this is a concern, you should
// use a different representation.
func (rgba *RGBA) String() string {
	var _arg0 *C.GdkRGBA // out
	var _cret *C.gchar   // in

	_arg0 = (*C.GdkRGBA)(gextras.StructNative(unsafe.Pointer(rgba)))

	_cret = C.gdk_rgba_to_string(_arg0)
	runtime.KeepAlive(rgba)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
