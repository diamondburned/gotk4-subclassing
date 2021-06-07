// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-wayland gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/wayland/gdkwayland.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_wayland_display_get_type()), F: marshalWaylandDisplay},
	})
}

type WaylandDisplay interface {
	gdk.Display

	// StartupNotificationID gets the startup notification ID for a Wayland
	// display, or nil if no ID has been defined.
	StartupNotificationID(d WaylandDisplay)
	// WlCompositor returns the Wayland global singleton compositor of a
	// Display.
	WlCompositor(d WaylandDisplay)
	// WlDisplay returns the Wayland wl_display of a Display.
	WlDisplay(d WaylandDisplay)
	// QueryRegistry returns true if the the interface was found in the display
	// `wl_registry.global` handler.
	QueryRegistry(d WaylandDisplay, global string) bool
	// SetCursorTheme sets the cursor theme for the given @display.
	SetCursorTheme(d WaylandDisplay, name string, size int)
	// SetStartupNotificationID sets the startup notification ID for a display.
	//
	// This is usually taken from the value of the DESKTOP_STARTUP_ID
	// environment variable, but in some cases (such as the application not
	// being launched using exec()) it can come from other sources.
	//
	// The startup ID is also what is used to signal that the startup is
	// complete (for example, when opening a window or when calling
	// gdk_display_notify_startup_complete()).
	SetStartupNotificationID(d WaylandDisplay, startupID string)
}

// waylandDisplay implements the WaylandDisplay interface.
type waylandDisplay struct {
	gdk.Display
}

var _ WaylandDisplay = (*waylandDisplay)(nil)

// WrapWaylandDisplay wraps a GObject to the right type. It is
// primarily used internally.
func WrapWaylandDisplay(obj *externglib.Object) WaylandDisplay {
	return WaylandDisplay{
		gdk.Display: gdk.WrapDisplay(obj),
	}
}

func marshalWaylandDisplay(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWaylandDisplay(obj), nil
}

// StartupNotificationID gets the startup notification ID for a Wayland
// display, or nil if no ID has been defined.
func (d waylandDisplay) StartupNotificationID(d WaylandDisplay) {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_wayland_display_get_startup_notification_id(arg0)
}

// WlCompositor returns the Wayland global singleton compositor of a
// Display.
func (d waylandDisplay) WlCompositor(d WaylandDisplay) {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_wayland_display_get_wl_compositor(arg0)
}

// WlDisplay returns the Wayland wl_display of a Display.
func (d waylandDisplay) WlDisplay(d WaylandDisplay) {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_wayland_display_get_wl_display(arg0)
}

// QueryRegistry returns true if the the interface was found in the display
// `wl_registry.global` handler.
func (d waylandDisplay) QueryRegistry(d WaylandDisplay, global string) bool {
	var arg0 *C.GdkDisplay
	var arg1 *C.char

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	arg1 = (*C.char)(C.CString(global))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ok bool

	cret = C.gdk_wayland_display_query_registry(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// SetCursorTheme sets the cursor theme for the given @display.
func (d waylandDisplay) SetCursorTheme(d WaylandDisplay, name string, size int) {
	var arg0 *C.GdkDisplay
	var arg1 *C.char
	var arg2 C.int

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.int(size)

	C.gdk_wayland_display_set_cursor_theme(arg0, arg1, arg2)
}

// SetStartupNotificationID sets the startup notification ID for a display.
//
// This is usually taken from the value of the DESKTOP_STARTUP_ID
// environment variable, but in some cases (such as the application not
// being launched using exec()) it can come from other sources.
//
// The startup ID is also what is used to signal that the startup is
// complete (for example, when opening a window or when calling
// gdk_display_notify_startup_complete()).
func (d waylandDisplay) SetStartupNotificationID(d WaylandDisplay, startupID string) {
	var arg0 *C.GdkDisplay
	var arg1 *C.char

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	arg1 = (*C.char)(C.CString(startupID))
	defer C.free(unsafe.Pointer(arg1))

	C.gdk_wayland_display_set_startup_notification_id(arg0, arg1)
}