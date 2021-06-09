// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_monitor_get_type()), F: marshalMonitor},
	})
}

// Monitor: gdkMonitor objects represent the individual outputs that are
// associated with a Display. GdkDisplay has APIs to enumerate monitors with
// gdk_display_get_n_monitors() and gdk_display_get_monitor(), and to find
// particular monitors with gdk_display_get_primary_monitor() or
// gdk_display_get_monitor_at_window().
//
// GdkMonitor was introduced in GTK+ 3.22 and supersedes earlier APIs in
// GdkScreen to obtain monitor-related information.
type Monitor interface {
	gextras.Objector

	// Display gets the display that this monitor belongs to.
	Display() Display
	// Geometry retrieves the size and position of an individual monitor within
	// the display coordinate space. The returned geometry is in ”application
	// pixels”, not in ”device pixels” (see gdk_monitor_get_scale_factor()).
	Geometry() Rectangle
	// HeightMm gets the height in millimeters of the monitor.
	HeightMm() int
	// Manufacturer gets the name or PNP ID of the monitor's manufacturer, if
	// available.
	//
	// Note that this value might also vary depending on actual display backend.
	//
	// PNP ID registry is located at https://uefi.org/pnp_id_list
	Manufacturer() string
	// Model gets the a string identifying the monitor model, if available.
	Model() string
	// RefreshRate gets the refresh rate of the monitor, if available.
	//
	// The value is in milli-Hertz, so a refresh rate of 60Hz is returned as
	// 60000.
	RefreshRate() int
	// ScaleFactor gets the internal scale factor that maps from monitor
	// coordinates to the actual device pixels. On traditional systems this is
	// 1, but on very high density outputs this can be a higher value (often 2).
	//
	// This can be used if you want to create pixel based data for a particular
	// monitor, but most of the time you’re drawing to a window where it is
	// better to use gdk_window_get_scale_factor() instead.
	ScaleFactor() int
	// SubpixelLayout gets information about the layout of red, green and blue
	// primaries for each pixel in this monitor, if available.
	SubpixelLayout() SubpixelLayout
	// WidthMm gets the width in millimeters of the monitor.
	WidthMm() int
	// Workarea retrieves the size and position of the “work area” on a monitor
	// within the display coordinate space. The returned geometry is in
	// ”application pixels”, not in ”device pixels” (see
	// gdk_monitor_get_scale_factor()).
	//
	// The work area should be considered when positioning menus and similar
	// popups, to avoid placing them below panels, docks or other desktop
	// components.
	//
	// Note that not all backends may have a concept of workarea. This function
	// will return the monitor geometry if a workarea is not available, or does
	// not apply.
	Workarea() Rectangle
	// IsPrimary gets whether this monitor should be considered primary (see
	// gdk_display_get_primary_monitor()).
	IsPrimary() bool
}

// monitor implements the Monitor interface.
type monitor struct {
	gextras.Objector
}

var _ Monitor = (*monitor)(nil)

// WrapMonitor wraps a GObject to the right type. It is
// primarily used internally.
func WrapMonitor(obj *externglib.Object) Monitor {
	return Monitor{
		Objector: obj,
	}
}

func marshalMonitor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMonitor(obj), nil
}

// Display gets the display that this monitor belongs to.
func (m monitor) Display() Display {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var cret *C.GdkDisplay

	cret = C.gdk_monitor_get_display(arg0)

	var display Display

	display = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Display)

	return display
}

// Geometry retrieves the size and position of an individual monitor within
// the display coordinate space. The returned geometry is in ”application
// pixels”, not in ”device pixels” (see gdk_monitor_get_scale_factor()).
func (m monitor) Geometry() Rectangle {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var geometry Rectangle

	C.gdk_monitor_get_geometry(arg0, (*C.GdkRectangle)(unsafe.Pointer(&geometry)))

	return geometry
}

// HeightMm gets the height in millimeters of the monitor.
func (m monitor) HeightMm() int {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var cret C.int

	cret = C.gdk_monitor_get_height_mm(arg0)

	var gint int

	gint = (int)(cret)

	return gint
}

// Manufacturer gets the name or PNP ID of the monitor's manufacturer, if
// available.
//
// Note that this value might also vary depending on actual display backend.
//
// PNP ID registry is located at https://uefi.org/pnp_id_list
func (m monitor) Manufacturer() string {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var cret *C.char

	cret = C.gdk_monitor_get_manufacturer(arg0)

	var utf8 string

	utf8 = C.GoString(cret)

	return utf8
}

// Model gets the a string identifying the monitor model, if available.
func (m monitor) Model() string {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var cret *C.char

	cret = C.gdk_monitor_get_model(arg0)

	var utf8 string

	utf8 = C.GoString(cret)

	return utf8
}

// RefreshRate gets the refresh rate of the monitor, if available.
//
// The value is in milli-Hertz, so a refresh rate of 60Hz is returned as
// 60000.
func (m monitor) RefreshRate() int {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var cret C.int

	cret = C.gdk_monitor_get_refresh_rate(arg0)

	var gint int

	gint = (int)(cret)

	return gint
}

// ScaleFactor gets the internal scale factor that maps from monitor
// coordinates to the actual device pixels. On traditional systems this is
// 1, but on very high density outputs this can be a higher value (often 2).
//
// This can be used if you want to create pixel based data for a particular
// monitor, but most of the time you’re drawing to a window where it is
// better to use gdk_window_get_scale_factor() instead.
func (m monitor) ScaleFactor() int {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var cret C.int

	cret = C.gdk_monitor_get_scale_factor(arg0)

	var gint int

	gint = (int)(cret)

	return gint
}

// SubpixelLayout gets information about the layout of red, green and blue
// primaries for each pixel in this monitor, if available.
func (m monitor) SubpixelLayout() SubpixelLayout {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var cret C.GdkSubpixelLayout

	cret = C.gdk_monitor_get_subpixel_layout(arg0)

	var subpixelLayout SubpixelLayout

	subpixelLayout = SubpixelLayout(cret)

	return subpixelLayout
}

// WidthMm gets the width in millimeters of the monitor.
func (m monitor) WidthMm() int {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var cret C.int

	cret = C.gdk_monitor_get_width_mm(arg0)

	var gint int

	gint = (int)(cret)

	return gint
}

// Workarea retrieves the size and position of the “work area” on a monitor
// within the display coordinate space. The returned geometry is in
// ”application pixels”, not in ”device pixels” (see
// gdk_monitor_get_scale_factor()).
//
// The work area should be considered when positioning menus and similar
// popups, to avoid placing them below panels, docks or other desktop
// components.
//
// Note that not all backends may have a concept of workarea. This function
// will return the monitor geometry if a workarea is not available, or does
// not apply.
func (m monitor) Workarea() Rectangle {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var workarea Rectangle

	C.gdk_monitor_get_workarea(arg0, (*C.GdkRectangle)(unsafe.Pointer(&workarea)))

	return workarea
}

// IsPrimary gets whether this monitor should be considered primary (see
// gdk_display_get_primary_monitor()).
func (m monitor) IsPrimary() bool {
	var arg0 *C.GdkMonitor

	arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var cret C.gboolean

	cret = C.gdk_monitor_is_primary(arg0)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}
