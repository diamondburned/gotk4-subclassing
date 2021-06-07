// Code generated by girgen. DO NOT EDIT.

package gdk

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// PixbufGetFromSurface transfers image data from a #cairo_surface_t and
// converts it to an RGB(A) representation inside a Pixbuf. This allows you to
// efficiently read individual pixels from cairo surfaces.
//
// This function will create an RGB pixbuf with 8 bits per channel. The pixbuf
// will contain an alpha channel if the @surface contains one.
func PixbufGetFromSurface(surface *cairo.Surface, srcX int, srcY int, width int, height int) {
	var arg1 *C.cairo_surface_t
	var arg2 C.int
	var arg3 C.int
	var arg4 C.int
	var arg5 C.int

	arg1 = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))
	arg2 = C.int(srcX)
	arg3 = C.int(srcY)
	arg4 = C.int(width)
	arg5 = C.int(height)

	C.gdk_pixbuf_get_from_surface(arg1, arg2, arg3, arg4, arg5)
}

// PixbufGetFromTexture creates a new pixbuf from @texture. This should
// generally not be used in newly written code as later stages will almost
// certainly convert the pixbuf back into a texture to draw it on screen.
func PixbufGetFromTexture(texture Texture) {
	var arg1 *C.GdkTexture

	arg1 = (*C.GdkTexture)(unsafe.Pointer(texture.Native()))

	C.gdk_pixbuf_get_from_texture(arg1)
}