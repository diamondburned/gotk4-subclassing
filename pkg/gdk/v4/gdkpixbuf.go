// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// PixbufGetFromSurface transfers image data from a cairo_surface_t and converts
// it to a GdkPixbuf.
//
// This allows you to efficiently read individual pixels from cairo surfaces.
//
// This function will create an RGB pixbuf with 8 bits per channel. The pixbuf
// will contain an alpha channel if the surface contains one.
//
// The function takes the following parameters:
//
//    - surface to copy from.
//    - srcX: source X coordinate within surface.
//    - srcY: source Y coordinate within surface.
//    - width: width in pixels of region to get.
//    - height: height in pixels of region to get.
//
func PixbufGetFromSurface(surface *cairo.Surface, srcX, srcY, width, height int) *gdkpixbuf.Pixbuf {
	var _arg1 *C.cairo_surface_t // out
	var _arg2 C.int              // out
	var _arg3 C.int              // out
	var _arg4 C.int              // out
	var _arg5 C.int              // out
	var _cret *C.GdkPixbuf       // in

	_arg1 = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))
	_arg2 = C.int(srcX)
	_arg3 = C.int(srcY)
	_arg4 = C.int(width)
	_arg5 = C.int(height)

	_cret = C.gdk_pixbuf_get_from_surface(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(surface)
	runtime.KeepAlive(srcX)
	runtime.KeepAlive(srcY)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	if _cret != nil {
		{
			obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
			_pixbuf = &gdkpixbuf.Pixbuf{
				Object: obj,
				LoadableIcon: gio.LoadableIcon{
					Icon: gio.Icon{
						Object: obj,
					},
				},
			}
		}
	}

	return _pixbuf
}

// PixbufGetFromTexture creates a new pixbuf from texture.
//
// This should generally not be used in newly written code as later stages will
// almost certainly convert the pixbuf back into a texture to draw it on screen.
//
// The function takes the following parameters:
//
//    - texture: GdkTexture.
//
func PixbufGetFromTexture(texture Texturer) *gdkpixbuf.Pixbuf {
	var _arg1 *C.GdkTexture // out
	var _cret *C.GdkPixbuf  // in

	_arg1 = (*C.GdkTexture)(unsafe.Pointer(texture.Native()))

	_cret = C.gdk_pixbuf_get_from_texture(_arg1)
	runtime.KeepAlive(texture)

	var _pixbuf *gdkpixbuf.Pixbuf // out

	if _cret != nil {
		{
			obj := externglib.AssumeOwnership(unsafe.Pointer(_cret))
			_pixbuf = &gdkpixbuf.Pixbuf{
				Object: obj,
				LoadableIcon: gio.LoadableIcon{
					Icon: gio.Icon{
						Object: obj,
					},
				},
			}
		}
	}

	return _pixbuf
}
