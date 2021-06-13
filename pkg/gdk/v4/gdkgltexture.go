// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_gl_texture_get_type()), F: marshalGLTexture},
	})
}

// GLTexture: a GdkTexture representing a GL texture object.
type GLTexture interface {
	Texture
	Paintable

	// Release releases the GL resources held by a `GdkGLTexture`.
	//
	// The texture contents are still available via the
	// [method@Gdk.Texture.download] function, after this function has been
	// called.
	Release()
}

// glTexture implements the GLTexture interface.
type glTexture struct {
	Texture
	Paintable
}

var _ GLTexture = (*glTexture)(nil)

// WrapGLTexture wraps a GObject to the right type. It is
// primarily used internally.
func WrapGLTexture(obj *externglib.Object) GLTexture {
	return GLTexture{
		Texture:   WrapTexture(obj),
		Paintable: WrapPaintable(obj),
	}
}

func marshalGLTexture(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapGLTexture(obj), nil
}

// Release releases the GL resources held by a `GdkGLTexture`.
//
// The texture contents are still available via the
// [method@Gdk.Texture.download] function, after this function has been
// called.
func (s glTexture) Release() {
	var _arg0 *C.GdkGLTexture // out

	_arg0 = (*C.GdkGLTexture)(unsafe.Pointer(s.Native()))

	C.gdk_gl_texture_release(_arg0)
}
