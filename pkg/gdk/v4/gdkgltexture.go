// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_gl_texture_get_type()), F: marshalGLTexturer},
	})
}

// GLTextureOverrider contains methods that are overridable.
type GLTextureOverrider interface {
}

// GLTexture: gdkTexture representing a GL texture object.
type GLTexture struct {
	_ [0]func() // equal guard
	Texture
}

var (
	_ Texturer = (*GLTexture)(nil)
)

func classInitGLTexturer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapGLTexture(obj *externglib.Object) *GLTexture {
	return &GLTexture{
		Texture: Texture{
			Object: obj,
			Paintable: Paintable{
				Object: obj,
			},
		},
	}
}

func marshalGLTexturer(p uintptr) (interface{}, error) {
	return wrapGLTexture(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Release releases the GL resources held by a GdkGLTexture.
//
// The texture contents are still available via the gdk.Texture.Download()
// function, after this function has been called.
func (self *GLTexture) Release() {
	var _arg0 *C.GdkGLTexture // out

	_arg0 = (*C.GdkGLTexture)(unsafe.Pointer(self.Native()))

	C.gdk_gl_texture_release(_arg0)
	runtime.KeepAlive(self)
}
