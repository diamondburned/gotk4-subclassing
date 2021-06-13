// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"github.com/diamondburned/gotk4/pkg/cairo"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_surface_get_type()), F: marshalSurface},
	})
}

// Surface: a `GdkSurface` is a rectangular region on the screen.
//
// It’s a low-level object, used to implement high-level objects such as
// [class@Gtk.Window] or [class@Gtk.Dialog] in GTK.
//
// The surfaces you see in practice are either [class@Gdk.Toplevel] or
// [class@Gdk.Popup], and those interfaces provide much of the required API to
// interact with these surfaces. Other, more specialized surface types exist,
// but you will rarely interact with them directly.
type Surface interface {
	gextras.Objector

	// Beep emits a short beep associated to @surface.
	//
	// If the display of @surface does not support per-surface beeps, emits a
	// short beep on the display just as [method@Gdk.Display.beep].
	Beep()
	// Destroy destroys the window system resources associated with @surface and
	// decrements @surface's reference count.
	//
	// The window system resources for all children of @surface are also
	// destroyed, but the children’s reference counts are not decremented.
	//
	// Note that a surface will not be destroyed automatically when its
	// reference count reaches zero. You must call this function yourself before
	// that happens.
	Destroy()
	// Height returns the height of the given @surface.
	//
	// Surface size is reported in ”application pixels”, not ”device pixels”
	// (see [method@Gdk.Surface.get_scale_factor]).
	Height() int
	// Mapped checks whether the surface has been mapped.
	//
	// A surface is mapped with [method@Gdk.Toplevel.present] or
	// [method@Gdk.Popup.present].
	Mapped() bool
	// ScaleFactor returns the internal scale factor that maps from surface
	// coordinates to the actual device pixels.
	//
	// On traditional systems this is 1, but on very high density outputs this
	// can be a higher value (often 2). A higher value means that drawing is
	// automatically scaled up to a higher resolution, so any code doing drawing
	// will automatically look nicer. However, if you are supplying pixel-based
	// data the scale value can be used to determine whether to use a pixel
	// resource with higher resolution data.
	//
	// The scale of a surface may change during runtime.
	ScaleFactor() int
	// Width returns the width of the given @surface.
	//
	// Surface size is reported in ”application pixels”, not ”device pixels”
	// (see [method@Gdk.Surface.get_scale_factor]).
	Width() int
	// Hide: hide the surface.
	//
	// For toplevel surfaces, withdraws them, so they will no longer be known to
	// the window manager; for all surfaces, unmaps them, so they won’t be
	// displayed. Normally done automatically as part of
	// [method@Gtk.Widget.hide].
	Hide()
	// IsDestroyed: check to see if a surface is destroyed.
	IsDestroyed() bool
	// QueueRender forces a [signal@Gdk.Surface::render] signal emission for
	// @surface to be scheduled.
	//
	// This function is useful for implementations that track invalid regions on
	// their own.
	QueueRender()
	// RequestLayout: request a layout phase from the surface's frame clock.
	//
	// See [method@Gdk.FrameClock.request_phase].
	RequestLayout()
	// SetCursor sets the default mouse pointer for a `GdkSurface`.
	//
	// Passing nil for the @cursor argument means that @surface will use the
	// cursor of its parent surface. Most surfaces should use this default. Note
	// that @cursor must be for the same display as @surface.
	//
	// Use [ctor@Gdk.Cursor.new_from_name] or [ctor@Gdk.Cursor.new_from_texture]
	// to create the cursor. To make the cursor invisible, use GDK_BLANK_CURSOR.
	SetCursor(cursor Cursor)
	// SetDeviceCursor sets a specific `GdkCursor` for a given device when it
	// gets inside @surface.
	//
	// Passing nil for the @cursor argument means that @surface will use the
	// cursor of its parent surface. Most surfaces should use this default.
	//
	// Use [ctor@Gdk.Cursor.new_from_name] or [ctor@Gdk.Cursor.new_from_texture]
	// to create the cursor. To make the cursor invisible, use GDK_BLANK_CURSOR.
	SetDeviceCursor(device Device, cursor Cursor)
	// SetInputRegion: apply the region to the surface for the purpose of event
	// handling.
	//
	// Mouse events which happen while the pointer position corresponds to an
	// unset bit in the mask will be passed on the surface below @surface.
	//
	// An input region is typically used with RGBA surfaces. The alpha channel
	// of the surface defines which pixels are invisible and allows for nicely
	// antialiased borders, and the input region controls where the surface is
	// “clickable”.
	//
	// Use [method@Gdk.Display.supports_input_shapes] to find out if a
	// particular backend supports input regions.
	SetInputRegion(region *cairo.Region)
	// SetOpaqueRegion marks a region of the `GdkSurface` as opaque.
	//
	// For optimisation purposes, compositing window managers may like to not
	// draw obscured regions of surfaces, or turn off blending during for these
	// regions. With RGB windows with no transparency, this is just the shape of
	// the window, but with ARGB32 windows, the compositor does not know what
	// regions of the window are transparent or not.
	//
	// This function only works for toplevel surfaces.
	//
	// GTK will update this property automatically if the @surface background is
	// opaque, as we know where the opaque regions are. If your surface
	// background is not opaque, please update this property in your
	// WidgetClass.css_changed() handler.
	SetOpaqueRegion(region *cairo.Region)
}

// surface implements the Surface interface.
type surface struct {
	gextras.Objector
}

var _ Surface = (*surface)(nil)

// WrapSurface wraps a GObject to the right type. It is
// primarily used internally.
func WrapSurface(obj *externglib.Object) Surface {
	return Surface{
		Objector: obj,
	}
}

func marshalSurface(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSurface(obj), nil
}

// Beep emits a short beep associated to @surface.
//
// If the display of @surface does not support per-surface beeps, emits a
// short beep on the display just as [method@Gdk.Display.beep].
func (s surface) Beep() {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	C.gdk_surface_beep(_arg0)
}

// Destroy destroys the window system resources associated with @surface and
// decrements @surface's reference count.
//
// The window system resources for all children of @surface are also
// destroyed, but the children’s reference counts are not decremented.
//
// Note that a surface will not be destroyed automatically when its
// reference count reaches zero. You must call this function yourself before
// that happens.
func (s surface) Destroy() {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	C.gdk_surface_destroy(_arg0)
}

// Height returns the height of the given @surface.
//
// Surface size is reported in ”application pixels”, not ”device pixels”
// (see [method@Gdk.Surface.get_scale_factor]).
func (s surface) Height() int {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	var _cret C.int // in

	_cret = C.gdk_surface_get_height(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Mapped checks whether the surface has been mapped.
//
// A surface is mapped with [method@Gdk.Toplevel.present] or
// [method@Gdk.Popup.present].
func (s surface) Mapped() bool {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gdk_surface_get_mapped(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// ScaleFactor returns the internal scale factor that maps from surface
// coordinates to the actual device pixels.
//
// On traditional systems this is 1, but on very high density outputs this
// can be a higher value (often 2). A higher value means that drawing is
// automatically scaled up to a higher resolution, so any code doing drawing
// will automatically look nicer. However, if you are supplying pixel-based
// data the scale value can be used to determine whether to use a pixel
// resource with higher resolution data.
//
// The scale of a surface may change during runtime.
func (s surface) ScaleFactor() int {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	var _cret C.int // in

	_cret = C.gdk_surface_get_scale_factor(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Width returns the width of the given @surface.
//
// Surface size is reported in ”application pixels”, not ”device pixels”
// (see [method@Gdk.Surface.get_scale_factor]).
func (s surface) Width() int {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	var _cret C.int // in

	_cret = C.gdk_surface_get_width(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Hide: hide the surface.
//
// For toplevel surfaces, withdraws them, so they will no longer be known to
// the window manager; for all surfaces, unmaps them, so they won’t be
// displayed. Normally done automatically as part of
// [method@Gtk.Widget.hide].
func (s surface) Hide() {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	C.gdk_surface_hide(_arg0)
}

// IsDestroyed: check to see if a surface is destroyed.
func (s surface) IsDestroyed() bool {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gdk_surface_is_destroyed(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// QueueRender forces a [signal@Gdk.Surface::render] signal emission for
// @surface to be scheduled.
//
// This function is useful for implementations that track invalid regions on
// their own.
func (s surface) QueueRender() {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	C.gdk_surface_queue_render(_arg0)
}

// RequestLayout: request a layout phase from the surface's frame clock.
//
// See [method@Gdk.FrameClock.request_phase].
func (s surface) RequestLayout() {
	var _arg0 *C.GdkSurface // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))

	C.gdk_surface_request_layout(_arg0)
}

// SetCursor sets the default mouse pointer for a `GdkSurface`.
//
// Passing nil for the @cursor argument means that @surface will use the
// cursor of its parent surface. Most surfaces should use this default. Note
// that @cursor must be for the same display as @surface.
//
// Use [ctor@Gdk.Cursor.new_from_name] or [ctor@Gdk.Cursor.new_from_texture]
// to create the cursor. To make the cursor invisible, use GDK_BLANK_CURSOR.
func (s surface) SetCursor(cursor Cursor) {
	var _arg0 *C.GdkSurface // out
	var _arg1 *C.GdkCursor  // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkCursor)(unsafe.Pointer(cursor.Native()))

	C.gdk_surface_set_cursor(_arg0, _arg1)
}

// SetDeviceCursor sets a specific `GdkCursor` for a given device when it
// gets inside @surface.
//
// Passing nil for the @cursor argument means that @surface will use the
// cursor of its parent surface. Most surfaces should use this default.
//
// Use [ctor@Gdk.Cursor.new_from_name] or [ctor@Gdk.Cursor.new_from_texture]
// to create the cursor. To make the cursor invisible, use GDK_BLANK_CURSOR.
func (s surface) SetDeviceCursor(device Device, cursor Cursor) {
	var _arg0 *C.GdkSurface // out
	var _arg1 *C.GdkDevice  // out
	var _arg2 *C.GdkCursor  // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	_arg2 = (*C.GdkCursor)(unsafe.Pointer(cursor.Native()))

	C.gdk_surface_set_device_cursor(_arg0, _arg1, _arg2)
}

// SetInputRegion: apply the region to the surface for the purpose of event
// handling.
//
// Mouse events which happen while the pointer position corresponds to an
// unset bit in the mask will be passed on the surface below @surface.
//
// An input region is typically used with RGBA surfaces. The alpha channel
// of the surface defines which pixels are invisible and allows for nicely
// antialiased borders, and the input region controls where the surface is
// “clickable”.
//
// Use [method@Gdk.Display.supports_input_shapes] to find out if a
// particular backend supports input regions.
func (s surface) SetInputRegion(region *cairo.Region) {
	var _arg0 *C.GdkSurface     // out
	var _arg1 *C.cairo_region_t // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.cairo_region_t)(unsafe.Pointer(region.Native()))

	C.gdk_surface_set_input_region(_arg0, _arg1)
}

// SetOpaqueRegion marks a region of the `GdkSurface` as opaque.
//
// For optimisation purposes, compositing window managers may like to not
// draw obscured regions of surfaces, or turn off blending during for these
// regions. With RGB windows with no transparency, this is just the shape of
// the window, but with ARGB32 windows, the compositor does not know what
// regions of the window are transparent or not.
//
// This function only works for toplevel surfaces.
//
// GTK will update this property automatically if the @surface background is
// opaque, as we know where the opaque regions are. If your surface
// background is not opaque, please update this property in your
// WidgetClass.css_changed() handler.
func (s surface) SetOpaqueRegion(region *cairo.Region) {
	var _arg0 *C.GdkSurface     // out
	var _arg1 *C.cairo_region_t // out

	_arg0 = (*C.GdkSurface)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.cairo_region_t)(unsafe.Pointer(region.Native()))

	C.gdk_surface_set_opaque_region(_arg0, _arg1)
}
