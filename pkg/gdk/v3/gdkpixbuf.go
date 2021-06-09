// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// PixbufGetFromSurface transfers image data from a #cairo_surface_t and
// converts it to an RGB(A) representation inside a Pixbuf. This allows you to
// efficiently read individual pixels from cairo surfaces. For Windows, use
// gdk_pixbuf_get_from_window() instead.
//
// This function will create an RGB pixbuf with 8 bits per channel. The pixbuf
// will contain an alpha channel if the @surface contains one.
func PixbufGetFromSurface(surface *cairo.Surface, srcX int, srcY int, width int, height int) gdkpixbuf.Pixbuf {
	var arg1 *C.cairo_surface_t
	var arg2 C.gint
	var arg3 C.gint
	var arg4 C.gint
	var arg5 C.gint

	arg1 = (*C.cairo_surface_t)(unsafe.Pointer(surface.Native()))
	arg2 = C.gint(srcX)
	arg3 = C.gint(srcY)
	arg4 = C.gint(width)
	arg5 = C.gint(height)

	var cret *C.GdkPixbuf

	cret = C.gdk_pixbuf_get_from_surface(arg1, arg2, arg3, arg4, arg5)

	var pixbuf gdkpixbuf.Pixbuf

	pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(gdkpixbuf.Pixbuf)

	return pixbuf
}

// PixbufGetFromWindow transfers image data from a Window and converts it to an
// RGB(A) representation inside a Pixbuf. In other words, copies image data from
// a server-side drawable to a client-side RGB(A) buffer. This allows you to
// efficiently read individual pixels on the client side.
//
// This function will create an RGB pixbuf with 8 bits per channel with the size
// specified by the @width and @height arguments scaled by the scale factor of
// @window. The pixbuf will contain an alpha channel if the @window contains
// one.
//
// If the window is off the screen, then there is no image data in the
// obscured/offscreen regions to be placed in the pixbuf. The contents of
// portions of the pixbuf corresponding to the offscreen region are undefined.
//
// If the window you’re obtaining data from is partially obscured by other
// windows, then the contents of the pixbuf areas corresponding to the obscured
// regions are undefined.
//
// If the window is not mapped (typically because it’s iconified/minimized or
// not on the current workspace), then nil will be returned.
//
// If memory can’t be allocated for the return value, nil will be returned
// instead.
//
// (In short, there are several ways this function can fail, and if it fails it
// returns nil; so check the return value.)
func PixbufGetFromWindow(window Window, srcX int, srcY int, width int, height int) gdkpixbuf.Pixbuf {
	var arg1 *C.GdkWindow
	var arg2 C.gint
	var arg3 C.gint
	var arg4 C.gint
	var arg5 C.gint

	arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	arg2 = C.gint(srcX)
	arg3 = C.gint(srcY)
	arg4 = C.gint(width)
	arg5 = C.gint(height)

	var cret *C.GdkPixbuf

	cret = C.gdk_pixbuf_get_from_window(arg1, arg2, arg3, arg4, arg5)

	var pixbuf gdkpixbuf.Pixbuf

	pixbuf = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(gdkpixbuf.Pixbuf)

	return pixbuf
}
