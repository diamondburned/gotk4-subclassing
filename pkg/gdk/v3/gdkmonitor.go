// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0 glib-2.0
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

// Geometry retrieves the size and position of an individual monitor within
// the display coordinate space. The returned geometry is in ”application
// pixels”, not in ”device pixels” (see gdk_monitor_get_scale_factor()).
func (m monitor) Geometry() Rectangle {
	var _arg0 *C.GdkMonitor // out

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var _geometry Rectangle

	C.gdk_monitor_get_geometry(_arg0, (*C.GdkRectangle)(unsafe.Pointer(&_geometry)))

	return _geometry
}

// HeightMm gets the height in millimeters of the monitor.
func (m monitor) HeightMm() int {
	var _arg0 *C.GdkMonitor // out

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var _cret C.int // in

	_cret = C.gdk_monitor_get_height_mm(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Manufacturer gets the name or PNP ID of the monitor's manufacturer, if
// available.
//
// Note that this value might also vary depending on actual display backend.
//
// PNP ID registry is located at https://uefi.org/pnp_id_list
func (m monitor) Manufacturer() string {
	var _arg0 *C.GdkMonitor // out

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var _cret *C.char // in

	_cret = C.gdk_monitor_get_manufacturer(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Model gets the a string identifying the monitor model, if available.
func (m monitor) Model() string {
	var _arg0 *C.GdkMonitor // out

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var _cret *C.char // in

	_cret = C.gdk_monitor_get_model(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// RefreshRate gets the refresh rate of the monitor, if available.
//
// The value is in milli-Hertz, so a refresh rate of 60Hz is returned as
// 60000.
func (m monitor) RefreshRate() int {
	var _arg0 *C.GdkMonitor // out

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var _cret C.int // in

	_cret = C.gdk_monitor_get_refresh_rate(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// ScaleFactor gets the internal scale factor that maps from monitor
// coordinates to the actual device pixels. On traditional systems this is
// 1, but on very high density outputs this can be a higher value (often 2).
//
// This can be used if you want to create pixel based data for a particular
// monitor, but most of the time you’re drawing to a window where it is
// better to use gdk_window_get_scale_factor() instead.
func (m monitor) ScaleFactor() int {
	var _arg0 *C.GdkMonitor // out

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var _cret C.int // in

	_cret = C.gdk_monitor_get_scale_factor(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// WidthMm gets the width in millimeters of the monitor.
func (m monitor) WidthMm() int {
	var _arg0 *C.GdkMonitor // out

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var _cret C.int // in

	_cret = C.gdk_monitor_get_width_mm(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
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
	var _arg0 *C.GdkMonitor // out

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var _workarea Rectangle

	C.gdk_monitor_get_workarea(_arg0, (*C.GdkRectangle)(unsafe.Pointer(&_workarea)))

	return _workarea
}

// IsPrimary gets whether this monitor should be considered primary (see
// gdk_display_get_primary_monitor()).
func (m monitor) IsPrimary() bool {
	var _arg0 *C.GdkMonitor // out

	_arg0 = (*C.GdkMonitor)(unsafe.Pointer(m.Native()))

	var _cret C.gboolean // in

	_cret = C.gdk_monitor_is_primary(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}
