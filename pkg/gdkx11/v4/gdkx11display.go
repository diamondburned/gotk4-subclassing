// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/x11/gdkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_display_get_type()), F: marshalX11Display},
	})
}

// X11SetSmClientID sets the `SM_CLIENT_ID` property on the application’s leader
// window so that the window manager can save the application’s state using the
// X11R6 ICCCM session management protocol.
//
// See the X Session Management Library documentation for more information on
// session management and the Inter-Client Communication Conventions Manual
func X11SetSmClientID(smClientID string) {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(smClientID))
	defer C.free(unsafe.Pointer(arg1))

	C.gdk_x11_set_sm_client_id(arg1)
}

type X11Display interface {
	gdk.Display

	// ErrorTrapPop pops the error trap pushed by
	// gdk_x11_display_error_trap_push(). Will XSync() if necessary and will
	// always block until the error is known to have occurred or not occurred,
	// so the error code can be returned.
	//
	// If you don’t need to use the return value,
	// gdk_x11_display_error_trap_pop_ignored() would be more efficient.
	ErrorTrapPop(d X11Display)
	// ErrorTrapPopIgnored pops the error trap pushed by
	// gdk_x11_display_error_trap_push(). Does not block to see if an error
	// occurred; merely records the range of requests to ignore errors for, and
	// ignores those errors if they arrive asynchronously.
	ErrorTrapPopIgnored(d X11Display)
	// ErrorTrapPush begins a range of X requests on @display for which X error
	// events will be ignored. Unignored errors (when no trap is pushed) will
	// abort the application. Use gdk_x11_display_error_trap_pop() or
	// gdk_x11_display_error_trap_pop_ignored()to lift a trap pushed with this
	// function.
	ErrorTrapPush(d X11Display)
	// DefaultGroup returns the default group leader surface for all toplevel
	// surfaces on @display. This surface is implicitly created by GDK. See
	// gdk_x11_surface_set_group().
	DefaultGroup(d X11Display)
	// GlxVersion retrieves the version of the GLX implementation.
	GlxVersion(d X11Display) (major int, minor int, ok bool)
	// PrimaryMonitor gets the primary monitor for the display.
	//
	// The primary monitor is considered the monitor where the “main desktop”
	// lives. While normal application surfaces typically allow the window
	// manager to place the surfaces, specialized desktop applications such as
	// panels should place themselves on the primary monitor.
	//
	// If no monitor is the designated primary monitor, any monitor (usually the
	// first) may be returned.
	PrimaryMonitor(d X11Display)
	// Screen retrieves the X11Screen of the @display.
	Screen(d X11Display)
	// StartupNotificationID gets the startup notification ID for a display.
	StartupNotificationID(d X11Display)
	// UserTime returns the timestamp of the last user interaction on @display.
	// The timestamp is taken from events caused by user interaction such as key
	// presses or pointer movements. See gdk_x11_surface_set_user_time().
	UserTime(d X11Display)
	// Xcursor returns the X cursor belonging to a Cursor, potentially creating
	// the cursor.
	//
	// Be aware that the returned cursor may not be unique to @cursor. It may
	// for example be shared with its fallback cursor. On old X servers that
	// don't support the XCursor extension, all cursors may even fall back to a
	// few default cursors.
	Xcursor(d X11Display, cursor gdk.Cursor)
	// Xdisplay returns the X display of a Display.
	Xdisplay(d X11Display)
	// Xrootwindow returns the root X window used by Display.
	Xrootwindow(d X11Display)
	// Xscreen returns the X Screen used by Display.
	Xscreen(d X11Display)
	// Grab: call XGrabServer() on @display. To ungrab the display again, use
	// gdk_x11_display_ungrab().
	//
	// gdk_x11_display_grab()/gdk_x11_display_ungrab() calls can be nested.
	Grab(d X11Display)
	// SetCursorTheme sets the cursor theme from which the images for cursor
	// should be taken.
	//
	// If the windowing system supports it, existing cursors created with
	// gdk_cursor_new_from_name() are updated to reflect the theme change.
	// Custom cursors constructed with gdk_cursor_new_from_texture() will have
	// to be handled by the application (GTK applications can learn about cursor
	// theme changes by listening for change notification for the corresponding
	// Setting).
	SetCursorTheme(d X11Display, theme string, size int)
	// SetStartupNotificationID sets the startup notification ID for a display.
	//
	// This is usually taken from the value of the DESKTOP_STARTUP_ID
	// environment variable, but in some cases (such as the application not
	// being launched using exec()) it can come from other sources.
	//
	// If the ID contains the string "_TIME" then the portion following that
	// string is taken to be the X11 timestamp of the event that triggered the
	// application to be launched and the GDK current event time is set
	// accordingly.
	//
	// The startup ID is also what is used to signal that the startup is
	// complete (for example, when opening a window or when calling
	// gdk_display_notify_startup_complete()).
	SetStartupNotificationID(d X11Display, startupID string)
	// SetSurfaceScale forces a specific window scale for all windows on this
	// display, instead of using the default or user configured scale. This is
	// can be used to disable scaling support by setting @scale to 1, or to
	// programmatically set the window scale.
	//
	// Once the scale is set by this call it will not change in response to
	// later user configuration changes.
	SetSurfaceScale(d X11Display, scale int)
	// StringToCompoundText: convert a string from the encoding of the current
	// locale into a form suitable for storing in a window property.
	StringToCompoundText(d X11Display, str string)
	// TextPropertyToTextList: convert a text string from the encoding as it is
	// stored in a property into an array of strings in the encoding of the
	// current locale. (The elements of the array represent the nul-separated
	// elements of the original text string.)
	TextPropertyToTextList(d X11Display, encoding string, format int, text byte, length int, list string)
	// Ungrab: ungrab @display after it has been grabbed with
	// gdk_x11_display_grab().
	Ungrab(d X11Display)
	// UTF8ToCompoundText converts from UTF-8 to compound text.
	UTF8ToCompoundText(d X11Display, str string) bool
}

// x11Display implements the X11Display interface.
type x11Display struct {
	gdk.Display
}

var _ X11Display = (*x11Display)(nil)

// WrapX11Display wraps a GObject to the right type. It is
// primarily used internally.
func WrapX11Display(obj *externglib.Object) X11Display {
	return X11Display{
		gdk.Display: gdk.WrapDisplay(obj),
	}
}

func marshalX11Display(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapX11Display(obj), nil
}

// ErrorTrapPop pops the error trap pushed by
// gdk_x11_display_error_trap_push(). Will XSync() if necessary and will
// always block until the error is known to have occurred or not occurred,
// so the error code can be returned.
//
// If you don’t need to use the return value,
// gdk_x11_display_error_trap_pop_ignored() would be more efficient.
func (d x11Display) ErrorTrapPop(d X11Display) {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_error_trap_pop(arg0)
}

// ErrorTrapPopIgnored pops the error trap pushed by
// gdk_x11_display_error_trap_push(). Does not block to see if an error
// occurred; merely records the range of requests to ignore errors for, and
// ignores those errors if they arrive asynchronously.
func (d x11Display) ErrorTrapPopIgnored(d X11Display) {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_error_trap_pop_ignored(arg0)
}

// ErrorTrapPush begins a range of X requests on @display for which X error
// events will be ignored. Unignored errors (when no trap is pushed) will
// abort the application. Use gdk_x11_display_error_trap_pop() or
// gdk_x11_display_error_trap_pop_ignored()to lift a trap pushed with this
// function.
func (d x11Display) ErrorTrapPush(d X11Display) {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_error_trap_push(arg0)
}

// DefaultGroup returns the default group leader surface for all toplevel
// surfaces on @display. This surface is implicitly created by GDK. See
// gdk_x11_surface_set_group().
func (d x11Display) DefaultGroup(d X11Display) {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_get_default_group(arg0)
}

// GlxVersion retrieves the version of the GLX implementation.
func (d x11Display) GlxVersion(d X11Display) (major int, minor int, ok bool) {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	var arg1 C.int
	var major int
	var arg2 C.int
	var minor int
	var cret C.gboolean
	var ok bool

	cret = C.gdk_x11_display_get_glx_version(arg0, &arg1, &arg2)

	major = int(&arg1)
	minor = int(&arg2)
	if cret {
		ok = true
	}

	return major, minor, ok
}

// PrimaryMonitor gets the primary monitor for the display.
//
// The primary monitor is considered the monitor where the “main desktop”
// lives. While normal application surfaces typically allow the window
// manager to place the surfaces, specialized desktop applications such as
// panels should place themselves on the primary monitor.
//
// If no monitor is the designated primary monitor, any monitor (usually the
// first) may be returned.
func (d x11Display) PrimaryMonitor(d X11Display) {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_get_primary_monitor(arg0)
}

// Screen retrieves the X11Screen of the @display.
func (d x11Display) Screen(d X11Display) {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_get_screen(arg0)
}

// StartupNotificationID gets the startup notification ID for a display.
func (d x11Display) StartupNotificationID(d X11Display) {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_get_startup_notification_id(arg0)
}

// UserTime returns the timestamp of the last user interaction on @display.
// The timestamp is taken from events caused by user interaction such as key
// presses or pointer movements. See gdk_x11_surface_set_user_time().
func (d x11Display) UserTime(d X11Display) {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_get_user_time(arg0)
}

// Xcursor returns the X cursor belonging to a Cursor, potentially creating
// the cursor.
//
// Be aware that the returned cursor may not be unique to @cursor. It may
// for example be shared with its fallback cursor. On old X servers that
// don't support the XCursor extension, all cursors may even fall back to a
// few default cursors.
func (d x11Display) Xcursor(d X11Display, cursor gdk.Cursor) {
	var arg0 *C.GdkDisplay
	var arg1 *C.GdkCursor

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	arg1 = (*C.GdkCursor)(unsafe.Pointer(cursor.Native()))

	C.gdk_x11_display_get_xcursor(arg0, arg1)
}

// Xdisplay returns the X display of a Display.
func (d x11Display) Xdisplay(d X11Display) {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_get_xdisplay(arg0)
}

// Xrootwindow returns the root X window used by Display.
func (d x11Display) Xrootwindow(d X11Display) {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_get_xrootwindow(arg0)
}

// Xscreen returns the X Screen used by Display.
func (d x11Display) Xscreen(d X11Display) {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_get_xscreen(arg0)
}

// Grab: call XGrabServer() on @display. To ungrab the display again, use
// gdk_x11_display_ungrab().
//
// gdk_x11_display_grab()/gdk_x11_display_ungrab() calls can be nested.
func (d x11Display) Grab(d X11Display) {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_grab(arg0)
}

// SetCursorTheme sets the cursor theme from which the images for cursor
// should be taken.
//
// If the windowing system supports it, existing cursors created with
// gdk_cursor_new_from_name() are updated to reflect the theme change.
// Custom cursors constructed with gdk_cursor_new_from_texture() will have
// to be handled by the application (GTK applications can learn about cursor
// theme changes by listening for change notification for the corresponding
// Setting).
func (d x11Display) SetCursorTheme(d X11Display, theme string, size int) {
	var arg0 *C.GdkDisplay
	var arg1 *C.char
	var arg2 C.int

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	arg1 = (*C.char)(C.CString(theme))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.int(size)

	C.gdk_x11_display_set_cursor_theme(arg0, arg1, arg2)
}

// SetStartupNotificationID sets the startup notification ID for a display.
//
// This is usually taken from the value of the DESKTOP_STARTUP_ID
// environment variable, but in some cases (such as the application not
// being launched using exec()) it can come from other sources.
//
// If the ID contains the string "_TIME" then the portion following that
// string is taken to be the X11 timestamp of the event that triggered the
// application to be launched and the GDK current event time is set
// accordingly.
//
// The startup ID is also what is used to signal that the startup is
// complete (for example, when opening a window or when calling
// gdk_display_notify_startup_complete()).
func (d x11Display) SetStartupNotificationID(d X11Display, startupID string) {
	var arg0 *C.GdkDisplay
	var arg1 *C.char

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	arg1 = (*C.char)(C.CString(startupID))
	defer C.free(unsafe.Pointer(arg1))

	C.gdk_x11_display_set_startup_notification_id(arg0, arg1)
}

// SetSurfaceScale forces a specific window scale for all windows on this
// display, instead of using the default or user configured scale. This is
// can be used to disable scaling support by setting @scale to 1, or to
// programmatically set the window scale.
//
// Once the scale is set by this call it will not change in response to
// later user configuration changes.
func (d x11Display) SetSurfaceScale(d X11Display, scale int) {
	var arg0 *C.GdkDisplay
	var arg1 C.int

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	arg1 = C.int(scale)

	C.gdk_x11_display_set_surface_scale(arg0, arg1)
}

// StringToCompoundText: convert a string from the encoding of the current
// locale into a form suitable for storing in a window property.
func (d x11Display) StringToCompoundText(d X11Display, str string) {
	var arg0 *C.GdkDisplay
	var arg1 *C.char

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))

	C.gdk_x11_display_string_to_compound_text(arg0, arg1, &arg2, &arg3, &arg4, &arg5)

	return encoding, format, ctext, length
}

// TextPropertyToTextList: convert a text string from the encoding as it is
// stored in a property into an array of strings in the encoding of the
// current locale. (The elements of the array represent the nul-separated
// elements of the original text string.)
func (d x11Display) TextPropertyToTextList(d X11Display, encoding string, format int, text byte, length int, list string) {
	var arg0 *C.GdkDisplay
	var arg1 *C.char
	var arg2 C.int
	var arg3 *C.guchar
	var arg4 C.int
	var arg5 ***C.char

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	arg1 = (*C.char)(C.CString(encoding))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.int(format)
	arg3 = *C.guchar(text)
	arg4 = C.int(length)
	arg5 = (***C.char)(C.CString(list))
	defer C.free(unsafe.Pointer(arg5))

	C.gdk_x11_display_text_property_to_text_list(arg0, arg1, arg2, arg3, arg4, arg5)
}

// Ungrab: ungrab @display after it has been grabbed with
// gdk_x11_display_grab().
func (d x11Display) Ungrab(d X11Display) {
	var arg0 *C.GdkDisplay

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_x11_display_ungrab(arg0)
}

// UTF8ToCompoundText converts from UTF-8 to compound text.
func (d x11Display) UTF8ToCompoundText(d X11Display, str string) bool {
	var arg0 *C.GdkDisplay
	var arg1 *C.char

	arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ok bool

	cret = C.gdk_x11_display_utf8_to_compound_text(arg0, arg1, &arg2, &arg3, &arg4, &arg5)

	if cret {
		ok = true
	}

	return encoding, format, ctext, length, ok
}