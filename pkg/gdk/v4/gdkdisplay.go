// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_display_get_type()), F: marshalDisplay},
	})
}

// Display: `GdkDisplay` objects are the GDK representation of a workstation.
//
// Their purpose are two-fold:
//
// - To manage and provide information about input devices (pointers, keyboards,
// etc) - To manage and provide information about output devices (monitors,
// projectors, etc)
//
// Most of the input device handling has been factored out into separate
// [class@Gdk.Seat] objects. Every display has a one or more seats, which can be
// accessed with [method@Gdk.Display.get_default_seat] and
// [method@Gdk.Display.list_seats].
//
// Output devices are represented by [class@Gdk.Monitor] objects, which can be
// accessed with [method@Gdk.Display.get_monitor_at_surface] and similar APIs.
type Display interface {
	gextras.Objector

	// BeepDisplay:
	BeepDisplay()
	// CloseDisplay:
	CloseDisplay()
	// DeviceIsGrabbedDisplay:
	DeviceIsGrabbedDisplay(device Device) bool
	// FlushDisplay:
	FlushDisplay()
	// AppLaunchContext:
	AppLaunchContext() AppLaunchContext
	// Clipboard:
	Clipboard() Clipboard
	// DefaultSeat:
	DefaultSeat() Seat
	// MonitorAtSurface:
	MonitorAtSurface(surface Surface) Monitor
	// Name:
	Name() string
	// PrimaryClipboard:
	PrimaryClipboard() Clipboard
	// Setting:
	Setting(name string, value externglib.Value) bool
	// StartupNotificationID:
	StartupNotificationID() string
	// IsClosedDisplay:
	IsClosedDisplay() bool
	// IsCompositedDisplay:
	IsCompositedDisplay() bool
	// IsRGBADisplay:
	IsRGBADisplay() bool
	// MapKeycodeDisplay:
	MapKeycodeDisplay(keycode uint) ([]KeymapKey, []uint, bool)
	// MapKeyvalDisplay:
	MapKeyvalDisplay(keyval uint) ([]KeymapKey, bool)
	// NotifyStartupCompleteDisplay:
	NotifyStartupCompleteDisplay(startupId string)
	// PutEventDisplay:
	PutEventDisplay(event Event)
	// SupportsInputShapesDisplay:
	SupportsInputShapesDisplay() bool
	// SyncDisplay:
	SyncDisplay()
	// TranslateKeyDisplay:
	TranslateKeyDisplay(keycode uint, state ModifierType, group int) (keyval uint, effectiveGroup int, level int, consumed ModifierType, ok bool)
}

// display implements the Display class.
type display struct {
	gextras.Objector
}

// WrapDisplay wraps a GObject to the right type. It is
// primarily used internally.
func WrapDisplay(obj *externglib.Object) Display {
	return display{
		Objector: obj,
	}
}

func marshalDisplay(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDisplay(obj), nil
}

func (d display) BeepDisplay() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_display_beep(_arg0)
}

func (d display) CloseDisplay() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_display_close(_arg0)
}

func (d display) DeviceIsGrabbedDisplay(device Device) bool {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkDevice  // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))

	_cret = C.gdk_display_device_is_grabbed(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d display) FlushDisplay() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_display_flush(_arg0)
}

func (d display) AppLaunchContext() AppLaunchContext {
	var _arg0 *C.GdkDisplay          // out
	var _cret *C.GdkAppLaunchContext // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_app_launch_context(_arg0)

	var _appLaunchContext AppLaunchContext // out

	_appLaunchContext = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(AppLaunchContext)

	return _appLaunchContext
}

func (d display) Clipboard() Clipboard {
	var _arg0 *C.GdkDisplay   // out
	var _cret *C.GdkClipboard // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_clipboard(_arg0)

	var _clipboard Clipboard // out

	_clipboard = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Clipboard)

	return _clipboard
}

func (d display) DefaultSeat() Seat {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.GdkSeat    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_default_seat(_arg0)

	var _seat Seat // out

	_seat = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Seat)

	return _seat
}

func (d display) MonitorAtSurface(surface Surface) Monitor {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkSurface // out
	var _cret *C.GdkMonitor // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GdkSurface)(unsafe.Pointer(surface.Native()))

	_cret = C.gdk_display_get_monitor_at_surface(_arg0, _arg1)

	var _monitor Monitor // out

	_monitor = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Monitor)

	return _monitor
}

func (d display) Name() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (d display) PrimaryClipboard() Clipboard {
	var _arg0 *C.GdkDisplay   // out
	var _cret *C.GdkClipboard // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_primary_clipboard(_arg0)

	var _clipboard Clipboard // out

	_clipboard = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Clipboard)

	return _clipboard
}

func (d display) Setting(name string, value externglib.Value) bool {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out
	var _arg2 *C.GValue     // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	_cret = C.gdk_display_get_setting(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d display) StartupNotificationID() string {
	var _arg0 *C.GdkDisplay // out
	var _cret *C.char       // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_get_startup_notification_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (d display) IsClosedDisplay() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_is_closed(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d display) IsCompositedDisplay() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_is_composited(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d display) IsRGBADisplay() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_is_rgba(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d display) MapKeycodeDisplay(keycode uint) ([]KeymapKey, []uint, bool) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.guint       // out
	var _arg2 *C.GdkKeymapKey
	var _arg4 C.int // in
	var _arg3 *C.guint
	var _arg4 C.int      // in
	var _cret C.gboolean // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = C.guint(keycode)

	_cret = C.gdk_display_map_keycode(_arg0, _arg1, &_arg2, &_arg3, &_arg4)

	var _keys []KeymapKey
	var _keyvals []uint
	var _ok bool // out

	_keys = unsafe.Slice((*KeymapKey)(unsafe.Pointer(_arg2)), _arg4)
	runtime.SetFinalizer(&_keys, func(v *[]KeymapKey) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})
	_keyvals = unsafe.Slice((*uint)(unsafe.Pointer(_arg3)), _arg4)
	runtime.SetFinalizer(&_keyvals, func(v *[]uint) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})
	if _cret != 0 {
		_ok = true
	}

	return _keys, _keyvals, _ok
}

func (d display) MapKeyvalDisplay(keyval uint) ([]KeymapKey, bool) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 C.guint       // out
	var _arg2 *C.GdkKeymapKey
	var _arg3 C.int      // in
	var _cret C.gboolean // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = C.guint(keyval)

	_cret = C.gdk_display_map_keyval(_arg0, _arg1, &_arg2, &_arg3)

	var _keys []KeymapKey
	var _ok bool // out

	_keys = unsafe.Slice((*KeymapKey)(unsafe.Pointer(_arg2)), _arg3)
	runtime.SetFinalizer(&_keys, func(v *[]KeymapKey) {
		C.free(unsafe.Pointer(&(*v)[0]))
	})
	if _cret != 0 {
		_ok = true
	}

	return _keys, _ok
}

func (d display) NotifyStartupCompleteDisplay(startupId string) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.char)(C.CString(startupId))
	defer C.free(unsafe.Pointer(_arg1))

	C.gdk_display_notify_startup_complete(_arg0, _arg1)
}

func (d display) PutEventDisplay(event Event) {
	var _arg0 *C.GdkDisplay // out
	var _arg1 *C.GdkEvent   // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = (*C.GdkEvent)(unsafe.Pointer(event.Native()))

	C.gdk_display_put_event(_arg0, _arg1)
}

func (d display) SupportsInputShapesDisplay() bool {
	var _arg0 *C.GdkDisplay // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	_cret = C.gdk_display_supports_input_shapes(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (d display) SyncDisplay() {
	var _arg0 *C.GdkDisplay // out

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))

	C.gdk_display_sync(_arg0)
}

func (d display) TranslateKeyDisplay(keycode uint, state ModifierType, group int) (keyval uint, effectiveGroup int, level int, consumed ModifierType, ok bool) {
	var _arg0 *C.GdkDisplay     // out
	var _arg1 C.guint           // out
	var _arg2 C.GdkModifierType // out
	var _arg3 C.int             // out
	var _arg4 C.guint           // in
	var _arg5 C.int             // in
	var _arg6 C.int             // in
	var _arg7 C.GdkModifierType // in
	var _cret C.gboolean        // in

	_arg0 = (*C.GdkDisplay)(unsafe.Pointer(d.Native()))
	_arg1 = C.guint(keycode)
	_arg2 = C.GdkModifierType(state)
	_arg3 = C.int(group)

	_cret = C.gdk_display_translate_key(_arg0, _arg1, _arg2, _arg3, &_arg4, &_arg5, &_arg6, &_arg7)

	var _keyval uint           // out
	var _effectiveGroup int    // out
	var _level int             // out
	var _consumed ModifierType // out
	var _ok bool               // out

	_keyval = uint(_arg4)
	_effectiveGroup = int(_arg5)
	_level = int(_arg6)
	_consumed = ModifierType(_arg7)
	if _cret != 0 {
		_ok = true
	}

	return _keyval, _effectiveGroup, _level, _consumed, _ok
}