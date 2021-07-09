// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-pixbuf-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk-pixbuf/gdk-pixbuf.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_pixbuf_simple_anim_get_type()), F: marshalPixbufSimpleAnim},
	})
}

// PixbufSimpleAnim: opaque struct representing a simple animation.
type PixbufSimpleAnim interface {
	gextras.Objector

	// AddFrame adds a new frame to @animation. The @pixbuf must have the
	// dimensions specified when the animation was constructed.
	AddFrame(pixbuf Pixbuf)
	// Loop gets whether @animation should loop indefinitely when it reaches the
	// end.
	Loop() bool
	// SetLoop sets whether @animation should loop indefinitely when it reaches
	// the end.
	SetLoop(loop bool)
}

// PixbufSimpleAnimClass implements the PixbufSimpleAnim interface.
type PixbufSimpleAnimClass struct {
	PixbufAnimationClass
}

var _ PixbufSimpleAnim = (*PixbufSimpleAnimClass)(nil)

func wrapPixbufSimpleAnim(obj *externglib.Object) PixbufSimpleAnim {
	return &PixbufSimpleAnimClass{
		PixbufAnimationClass: PixbufAnimationClass{
			Object: obj,
		},
	}
}

func marshalPixbufSimpleAnim(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPixbufSimpleAnim(obj), nil
}

// NewPixbufSimpleAnim creates a new, empty animation.
func NewPixbufSimpleAnim(width int, height int, rate float32) *PixbufSimpleAnimClass {
	var _arg1 C.gint                 // out
	var _arg2 C.gint                 // out
	var _arg3 C.gfloat               // out
	var _cret *C.GdkPixbufSimpleAnim // in

	_arg1 = C.gint(width)
	_arg2 = C.gint(height)
	_arg3 = C.gfloat(rate)

	_cret = C.gdk_pixbuf_simple_anim_new(_arg1, _arg2, _arg3)

	var _pixbufSimpleAnim *PixbufSimpleAnimClass // out

	_pixbufSimpleAnim = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*PixbufSimpleAnimClass)

	return _pixbufSimpleAnim
}

// AddFrame adds a new frame to @animation. The @pixbuf must have the dimensions
// specified when the animation was constructed.
func (a *PixbufSimpleAnimClass) AddFrame(pixbuf Pixbuf) {
	var _arg0 *C.GdkPixbufSimpleAnim // out
	var _arg1 *C.GdkPixbuf           // out

	_arg0 = (*C.GdkPixbufSimpleAnim)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gdk_pixbuf_simple_anim_add_frame(_arg0, _arg1)
}

// Loop gets whether @animation should loop indefinitely when it reaches the
// end.
func (a *PixbufSimpleAnimClass) Loop() bool {
	var _arg0 *C.GdkPixbufSimpleAnim // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GdkPixbufSimpleAnim)(unsafe.Pointer(a.Native()))

	_cret = C.gdk_pixbuf_simple_anim_get_loop(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetLoop sets whether @animation should loop indefinitely when it reaches the
// end.
func (a *PixbufSimpleAnimClass) SetLoop(loop bool) {
	var _arg0 *C.GdkPixbufSimpleAnim // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GdkPixbufSimpleAnim)(unsafe.Pointer(a.Native()))
	if loop {
		_arg1 = C.TRUE
	}

	C.gdk_pixbuf_simple_anim_set_loop(_arg0, _arg1)
}
