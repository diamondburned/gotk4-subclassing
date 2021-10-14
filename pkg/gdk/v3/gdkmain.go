// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// AddOptionEntriesLibgtkOnly appends gdk option entries to the passed in option
// group. This is not public API and must not be used by applications.
//
// Deprecated: This symbol was never meant to be used outside of GTK+.
//
// The function takes the following parameters:
//
//    - group: option group.
//
func AddOptionEntriesLibgtkOnly(group *glib.OptionGroup) {
	var _arg1 *C.GOptionGroup // out

	_arg1 = (*C.GOptionGroup)(gextras.StructNative(unsafe.Pointer(group)))

	C.gdk_add_option_entries_libgtk_only(_arg1)
	runtime.KeepAlive(group)
}

// Beep emits a short beep on the default display.
func Beep() {
	C.gdk_beep()
}

// DisableMultidevice disables multidevice support in GDK. This call must happen
// prior to gdk_display_open(), gtk_init(), gtk_init_with_args() or
// gtk_init_check() in order to take effect.
//
// Most common GTK+ applications won’t ever need to call this. Only applications
// that do mixed GDK/Xlib calls could want to disable multidevice support if
// such Xlib code deals with input devices in any way and doesn’t observe the
// presence of XInput 2.
func DisableMultidevice() {
	C.gdk_disable_multidevice()
}

// ErrorTrapPop removes an error trap pushed with gdk_error_trap_push(). May
// block until an error has been definitively received or not received from the
// X server. gdk_error_trap_pop_ignored() is preferred if you don’t need to know
// whether an error occurred, because it never has to block. If you don't need
// the return value of gdk_error_trap_pop(), use gdk_error_trap_pop_ignored().
//
// Prior to GDK 3.0, this function would not automatically sync for you, so you
// had to gdk_flush() if your last call to Xlib was not a blocking round trip.
func ErrorTrapPop() int {
	var _cret C.gint // in

	_cret = C.gdk_error_trap_pop()

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ErrorTrapPopIgnored removes an error trap pushed with gdk_error_trap_push(),
// but without bothering to wait and see whether an error occurred. If an error
// arrives later asynchronously that was triggered while the trap was pushed,
// that error will be ignored.
func ErrorTrapPopIgnored() {
	C.gdk_error_trap_pop_ignored()
}

// ErrorTrapPush: this function allows X errors to be trapped instead of the
// normal behavior of exiting the application. It should only be used if it is
// not possible to avoid the X error in any other way. Errors are ignored on all
// Display currently known to the DisplayManager. If you don’t care which error
// happens and just want to ignore everything, pop with
// gdk_error_trap_pop_ignored(). If you need the error code, use
// gdk_error_trap_pop() which may have to block and wait for the error to arrive
// from the X server.
//
// This API exists on all platforms but only does anything on X.
//
// You can use gdk_x11_display_error_trap_push() to ignore errors on only a
// single display.
//
// Trapping an X error
//
//    gdk_error_trap_push ();
//
//     // ... Call the X function which may cause an error here ...
//
//
//    if (gdk_error_trap_pop ())
//     {
//       // ... Handle the error here ...
//     }.
func ErrorTrapPush() {
	C.gdk_error_trap_push()
}

// Flush flushes the output buffers of all display connections and waits until
// all requests have been processed. This is rarely needed by applications.
func Flush() {
	C.gdk_flush()
}

// GetDisplay gets the name of the display, which usually comes from the DISPLAY
// environment variable or the --display command line option.
//
// Deprecated: Call gdk_display_get_name (gdk_display_get_default ())) instead.
func GetDisplay() string {
	var _cret *C.gchar // in

	_cret = C.gdk_get_display()

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// GetDisplayArgName gets the display name specified in the command line
// arguments passed to gdk_init() or gdk_parse_args(), if any.
func GetDisplayArgName() string {
	var _cret *C.gchar // in

	_cret = C.gdk_get_display_arg_name()

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// GetProgramClass gets the program class. Unless the program class has
// explicitly been set with gdk_set_program_class() or with the --class
// commandline option, the default value is the program name (determined with
// g_get_prgname()) with the first character converted to uppercase.
func GetProgramClass() string {
	var _cret *C.gchar // in

	_cret = C.gdk_get_program_class()

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// KeyboardGrab grabs the keyboard so that all events are passed to this
// application until the keyboard is ungrabbed with gdk_keyboard_ungrab(). This
// overrides any previous keyboard grab by this client.
//
// If you set up anything at the time you take the grab that needs to be cleaned
// up when the grab ends, you should handle the EventGrabBroken events that are
// emitted when the grab ends unvoluntarily.
//
// Deprecated: Use gdk_device_grab() instead.
//
// The function takes the following parameters:
//
//    - window which will own the grab (the grab window).
//    - ownerEvents: if FALSE then all keyboard events are reported with
//    respect to window. If TRUE then keyboard events for this application are
//    reported as normal, but keyboard events outside this application are
//    reported with respect to window. Both key press and key release events
//    are always reported, independant of the event mask set by the
//    application.
//    - time_: timestamp from a Event, or GDK_CURRENT_TIME if no timestamp is
//    available.
//
func KeyboardGrab(window Windower, ownerEvents bool, time_ uint32) GrabStatus {
	var _arg1 *C.GdkWindow    // out
	var _arg2 C.gboolean      // out
	var _arg3 C.guint32       // out
	var _cret C.GdkGrabStatus // in

	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	if ownerEvents {
		_arg2 = C.TRUE
	}
	_arg3 = C.guint32(time_)

	_cret = C.gdk_keyboard_grab(_arg1, _arg2, _arg3)
	runtime.KeepAlive(window)
	runtime.KeepAlive(ownerEvents)
	runtime.KeepAlive(time_)

	var _grabStatus GrabStatus // out

	_grabStatus = GrabStatus(_cret)

	return _grabStatus
}

// KeyboardUngrab ungrabs the keyboard on the default display, if it is grabbed
// by this application.
//
// Deprecated: Use gdk_device_ungrab(), together with gdk_device_grab() instead.
//
// The function takes the following parameters:
//
//    - time_: timestamp from a Event, or GDK_CURRENT_TIME if no timestamp is
//    available.
//
func KeyboardUngrab(time_ uint32) {
	var _arg1 C.guint32 // out

	_arg1 = C.guint32(time_)

	C.gdk_keyboard_ungrab(_arg1)
	runtime.KeepAlive(time_)
}

// NotifyStartupComplete indicates to the GUI environment that the application
// has finished loading. If the applications opens windows, this function is
// normally called after opening the application’s initial set of windows.
//
// GTK+ will call this function automatically after opening the first Window
// unless gtk_window_set_auto_startup_notification() is called to disable that
// feature.
func NotifyStartupComplete() {
	C.gdk_notify_startup_complete()
}

// NotifyStartupCompleteWithID indicates to the GUI environment that the
// application has finished loading, using a given identifier.
//
// GTK+ will call this function automatically for Window with custom
// startup-notification identifier unless
// gtk_window_set_auto_startup_notification() is called to disable that feature.
//
// The function takes the following parameters:
//
//    - startupId: startup-notification identifier, for which notification
//    process should be completed.
//
func NotifyStartupCompleteWithID(startupId string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(startupId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_notify_startup_complete_with_id(_arg1)
	runtime.KeepAlive(startupId)
}

// PointerGrab grabs the pointer (usually a mouse) so that all events are passed
// to this application until the pointer is ungrabbed with gdk_pointer_ungrab(),
// or the grab window becomes unviewable. This overrides any previous pointer
// grab by this client.
//
// Pointer grabs are used for operations which need complete control over mouse
// events, even if the mouse leaves the application. For example in GTK+ it is
// used for Drag and Drop, for dragging the handle in the HPaned and VPaned
// widgets.
//
// Note that if the event mask of an X window has selected both button press and
// button release events, then a button press event will cause an automatic
// pointer grab until the button is released. X does this automatically since
// most applications expect to receive button press and release events in pairs.
// It is equivalent to a pointer grab on the window with owner_events set to
// TRUE.
//
// If you set up anything at the time you take the grab that needs to be cleaned
// up when the grab ends, you should handle the EventGrabBroken events that are
// emitted when the grab ends unvoluntarily.
//
// Deprecated: Use gdk_device_grab() instead.
//
// The function takes the following parameters:
//
//    - window which will own the grab (the grab window).
//    - ownerEvents: if FALSE then all pointer events are reported with respect
//    to window and are only reported if selected by event_mask. If TRUE then
//    pointer events for this application are reported as normal, but pointer
//    events outside this application are reported with respect to window and
//    only if selected by event_mask. In either mode, unreported events are
//    discarded.
//    - eventMask specifies the event mask, which is used in accordance with
//    owner_events. Note that only pointer events (i.e. button and motion
//    events) may be selected.
//    - confineTo: if non-NULL, the pointer will be confined to this window
//    during the grab. If the pointer is outside confine_to, it will
//    automatically be moved to the closest edge of confine_to and enter and
//    leave events will be generated as necessary.
//    - cursor to display while the grab is active. If this is NULL then the
//    normal cursors are used for window and its descendants, and the cursor
//    for window is used for all other windows.
//    - time_: timestamp of the event which led to this pointer grab. This
//    usually comes from a EventButton struct, though GDK_CURRENT_TIME can be
//    used if the time isn’t known.
//
func PointerGrab(window Windower, ownerEvents bool, eventMask EventMask, confineTo Windower, cursor Cursorrer, time_ uint32) GrabStatus {
	var _arg1 *C.GdkWindow    // out
	var _arg2 C.gboolean      // out
	var _arg3 C.GdkEventMask  // out
	var _arg4 *C.GdkWindow    // out
	var _arg5 *C.GdkCursor    // out
	var _arg6 C.guint32       // out
	var _cret C.GdkGrabStatus // in

	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	if ownerEvents {
		_arg2 = C.TRUE
	}
	_arg3 = C.GdkEventMask(eventMask)
	if confineTo != nil {
		_arg4 = (*C.GdkWindow)(unsafe.Pointer(confineTo.Native()))
	}
	if cursor != nil {
		_arg5 = (*C.GdkCursor)(unsafe.Pointer(cursor.Native()))
	}
	_arg6 = C.guint32(time_)

	_cret = C.gdk_pointer_grab(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(window)
	runtime.KeepAlive(ownerEvents)
	runtime.KeepAlive(eventMask)
	runtime.KeepAlive(confineTo)
	runtime.KeepAlive(cursor)
	runtime.KeepAlive(time_)

	var _grabStatus GrabStatus // out

	_grabStatus = GrabStatus(_cret)

	return _grabStatus
}

// PointerIsGrabbed returns TRUE if the pointer on the default display is
// currently grabbed by this application.
//
// Note that this does not take the inmplicit pointer grab on button presses
// into account.
//
// Deprecated: Use gdk_display_device_is_grabbed() instead.
func PointerIsGrabbed() bool {
	var _cret C.gboolean // in

	_cret = C.gdk_pointer_is_grabbed()

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PointerUngrab ungrabs the pointer on the default display, if it is grabbed by
// this application.
//
// Deprecated: Use gdk_device_ungrab(), together with gdk_device_grab() instead.
//
// The function takes the following parameters:
//
//    - time_: timestamp from a Event, or GDK_CURRENT_TIME if no timestamp is
//    available.
//
func PointerUngrab(time_ uint32) {
	var _arg1 C.guint32 // out

	_arg1 = C.guint32(time_)

	C.gdk_pointer_ungrab(_arg1)
	runtime.KeepAlive(time_)
}

// PreParseLibgtkOnly: prepare for parsing command line arguments for GDK. This
// is not public API and should not be used in application code.
//
// Deprecated: This symbol was never meant to be used outside of GTK+.
func PreParseLibgtkOnly() {
	C.gdk_pre_parse_libgtk_only()
}

// SetAllowedBackends sets a list of backends that GDK should try to use.
//
// This can be be useful if your application does not work with certain GDK
// backends.
//
// By default, GDK tries all included backends.
//
// For example,
//
//    gdk_set_allowed_backends ("wayland,quartz,*");
//
// instructs GDK to try the Wayland backend first, followed by the Quartz
// backend, and then all others.
//
// If the GDK_BACKEND environment variable is set, it determines what backends
// are tried in what order, while still respecting the set of allowed backends
// that are specified by this function.
//
// The possible backend names are x11, win32, quartz, broadway, wayland. You can
// also include a * in the list to try all remaining backends.
//
// This call must happen prior to gdk_display_open(), gtk_init(),
// gtk_init_with_args() or gtk_init_check() in order to take effect.
//
// The function takes the following parameters:
//
//    - backends: comma-separated list of backends.
//
func SetAllowedBackends(backends string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(backends)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_set_allowed_backends(_arg1)
	runtime.KeepAlive(backends)
}

// SetDoubleClickTime: set the double click time for the default display. See
// gdk_display_set_double_click_time(). See also
// gdk_display_set_double_click_distance(). Applications should not set this, it
// is a global user-configured setting.
//
// The function takes the following parameters:
//
//    - msec: double click time in milliseconds (thousandths of a second).
//
func SetDoubleClickTime(msec uint) {
	var _arg1 C.guint // out

	_arg1 = C.guint(msec)

	C.gdk_set_double_click_time(_arg1)
	runtime.KeepAlive(msec)
}

// SetProgramClass sets the program class. The X11 backend uses the program
// class to set the class name part of the WM_CLASS property on toplevel
// windows; see the ICCCM.
//
// The program class can still be overridden with the --class command line
// option.
//
// The function takes the following parameters:
//
//    - programClass: string.
//
func SetProgramClass(programClass string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(programClass)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_set_program_class(_arg1)
	runtime.KeepAlive(programClass)
}

// ScreenHeight gets the height of the default screen in pixels. The returned
// size is in ”application pixels”, not in ”device pixels” (see
// gdk_screen_get_monitor_scale_factor()).
//
// Deprecated: Use per-monitor information.
func ScreenHeight() int {
	var _cret C.gint // in

	_cret = C.gdk_screen_height()

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ScreenHeightMm returns the height of the default screen in millimeters. Note
// that on many X servers this value will not be correct.
//
// Deprecated: Use per-monitor information.
func ScreenHeightMm() int {
	var _cret C.gint // in

	_cret = C.gdk_screen_height_mm()

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ScreenWidth gets the width of the default screen in pixels. The returned size
// is in ”application pixels”, not in ”device pixels” (see
// gdk_screen_get_monitor_scale_factor()).
//
// Deprecated: Use per-monitor information.
func ScreenWidth() int {
	var _cret C.gint // in

	_cret = C.gdk_screen_width()

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ScreenWidthMm returns the width of the default screen in millimeters. Note
// that on many X servers this value will not be correct.
//
// Deprecated: Use per-monitor information.
func ScreenWidthMm() int {
	var _cret C.gint // in

	_cret = C.gdk_screen_width_mm()

	var _gint int // out

	_gint = int(_cret)

	return _gint
}
