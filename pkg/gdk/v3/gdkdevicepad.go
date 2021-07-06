// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_device_pad_feature_get_type()), F: marshalDevicePadFeature},
		{T: externglib.Type(C.gdk_device_pad_get_type()), F: marshalDevicePad},
	})
}

// DevicePadFeature: pad feature.
type DevicePadFeature int

const (
	// Button: button
	Button DevicePadFeature = iota
	// Ring: ring-shaped interactive area
	Ring
	// Strip: straight interactive area
	Strip
)

func marshalDevicePadFeature(p uintptr) (interface{}, error) {
	return DevicePadFeature(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// DevicePad is an interface implemented by devices of type
// GDK_SOURCE_TABLET_PAD, it allows querying the features provided by the pad
// device.
//
// Tablet pads may contain one or more groups, each containing a subset of the
// buttons/rings/strips available. gdk_device_pad_get_n_groups() can be used to
// obtain the number of groups, gdk_device_pad_get_n_features() and
// gdk_device_pad_get_feature_group() can be combined to find out the number of
// buttons/rings/strips the device has, and how are they grouped.
//
// Each of those groups have different modes, which may be used to map each
// individual pad feature to multiple actions. Only one mode is effective
// (current) for each given group, different groups may have different current
// modes. The number of available modes in a group can be found out through
// gdk_device_pad_get_group_n_modes(), and the current mode for a given group
// will be notified through the EventPadGroupMode event.
type DevicePad interface {
	gextras.Objector

	// AsDevice casts the class to the Device interface.
	AsDevice() Device

	// GetAssociatedDevice returns the associated device to @device, if @device
	// is of type GDK_DEVICE_TYPE_MASTER, it will return the paired pointer or
	// keyboard.
	//
	// If @device is of type GDK_DEVICE_TYPE_SLAVE, it will return the master
	// device to which @device is attached to.
	//
	// If @device is of type GDK_DEVICE_TYPE_FLOATING, nil will be returned, as
	// there is no associated device.
	//
	// This method is inherited from Device
	GetAssociatedDevice() Device
	// GetAxes returns the axes currently available on the device.
	//
	// This method is inherited from Device
	GetAxes() AxisFlags
	// GetAxisUse returns the axis use for @index_.
	//
	// This method is inherited from Device
	GetAxisUse(index_ uint) AxisUse
	// GetDeviceType returns the device type for @device.
	//
	// This method is inherited from Device
	GetDeviceType() DeviceType
	// GetDisplay returns the Display to which @device pertains.
	//
	// This method is inherited from Device
	GetDisplay() Display
	// GetHasCursor determines whether the pointer follows device motion. This
	// is not meaningful for keyboard devices, which don't have a pointer.
	//
	// This method is inherited from Device
	GetHasCursor() bool
	// GetKey: if @index_ has a valid keyval, this function will return true and
	// fill in @keyval and @modifiers with the keyval settings.
	//
	// This method is inherited from Device
	GetKey(index_ uint) (uint, ModifierType, bool)
	// GetLastEventWindow gets information about which window the given pointer
	// device is in, based on events that have been received so far from the
	// display server. If another application has a pointer grab, or this
	// application has a grab with owner_events = false, nil may be returned
	// even if the pointer is physically over one of this application's windows.
	//
	// This method is inherited from Device
	GetLastEventWindow() Window
	// GetMode determines the mode of the device.
	//
	// This method is inherited from Device
	GetMode() InputMode
	// GetNAxes returns the number of axes the device currently has.
	//
	// This method is inherited from Device
	GetNAxes() int
	// GetNKeys returns the number of keys the device currently has.
	//
	// This method is inherited from Device
	GetNKeys() int
	// GetName determines the name of the device.
	//
	// This method is inherited from Device
	GetName() string
	// GetPosition gets the current location of @device. As a slave device
	// coordinates are those of its master pointer, This function may not be
	// called on devices of type GDK_DEVICE_TYPE_SLAVE, unless there is an
	// ongoing grab on them, see gdk_device_grab().
	//
	// This method is inherited from Device
	GetPosition() (screen Screen, x int, y int)
	// GetPositionDouble gets the current location of @device in double
	// precision. As a slave device's coordinates are those of its master
	// pointer, this function may not be called on devices of type
	// GDK_DEVICE_TYPE_SLAVE, unless there is an ongoing grab on them. See
	// gdk_device_grab().
	//
	// This method is inherited from Device
	GetPositionDouble() (screen Screen, x float64, y float64)
	// GetProductID returns the product ID of this device, or nil if this
	// information couldn't be obtained. This ID is retrieved from the device,
	// and is thus constant for it. See gdk_device_get_vendor_id() for more
	// information.
	//
	// This method is inherited from Device
	GetProductID() string
	// GetSeat returns the Seat the device belongs to.
	//
	// This method is inherited from Device
	GetSeat() Seat
	// GetSource determines the type of the device.
	//
	// This method is inherited from Device
	GetSource() InputSource
	// GetVendorID returns the vendor ID of this device, or nil if this
	// information couldn't be obtained. This ID is retrieved from the device,
	// and is thus constant for it.
	//
	// This function, together with gdk_device_get_product_id(), can be used to
	// eg. compose #GSettings paths to store settings for this device.
	//
	//     static GSettings *
	//     get_device_settings (GdkDevice *device)
	//     {
	//       const gchar *vendor, *product;
	//       GSettings *settings;
	//       GdkDevice *device;
	//       gchar *path;
	//
	//       vendor = gdk_device_get_vendor_id (device);
	//       product = gdk_device_get_product_id (device);
	//
	//       path = g_strdup_printf ("/org/example/app/devices/s:s/", vendor, product);
	//       settings = g_settings_new_with_path (DEVICE_SCHEMA, path);
	//       g_free (path);
	//
	//       return settings;
	//     }
	//
	// This method is inherited from Device
	GetVendorID() string
	// GetWindowAtPosition obtains the window underneath @device, returning the
	// location of the device in @win_x and @win_y. Returns nil if the window
	// tree under @device is not known to GDK (for example, belongs to another
	// application).
	//
	// As a slave device coordinates are those of its master pointer, This
	// function may not be called on devices of type GDK_DEVICE_TYPE_SLAVE,
	// unless there is an ongoing grab on them, see gdk_device_grab().
	//
	// This method is inherited from Device
	GetWindowAtPosition() (winX int, winY int, window Window)
	// GetWindowAtPositionDouble obtains the window underneath @device,
	// returning the location of the device in @win_x and @win_y in double
	// precision. Returns nil if the window tree under @device is not known to
	// GDK (for example, belongs to another application).
	//
	// As a slave device coordinates are those of its master pointer, This
	// function may not be called on devices of type GDK_DEVICE_TYPE_SLAVE,
	// unless there is an ongoing grab on them, see gdk_device_grab().
	//
	// This method is inherited from Device
	GetWindowAtPositionDouble() (winX float64, winY float64, window Window)
	// Grab grabs the device so that all events coming from this device are
	// passed to this application until the device is ungrabbed with
	// gdk_device_ungrab(), or the window becomes unviewable. This overrides any
	// previous grab on the device by this client.
	//
	// Note that @device and @window need to be on the same display.
	//
	// Device grabs are used for operations which need complete control over the
	// given device events (either pointer or keyboard). For example in GTK+
	// this is used for Drag and Drop operations, popup menus and such.
	//
	// Note that if the event mask of an X window has selected both button press
	// and button release events, then a button press event will cause an
	// automatic pointer grab until the button is released. X does this
	// automatically since most applications expect to receive button press and
	// release events in pairs. It is equivalent to a pointer grab on the window
	// with @owner_events set to true.
	//
	// If you set up anything at the time you take the grab that needs to be
	// cleaned up when the grab ends, you should handle the EventGrabBroken
	// events that are emitted when the grab ends unvoluntarily.
	//
	// Deprecated: since version 3.20.
	//
	// This method is inherited from Device
	Grab(window Window, grabOwnership GrabOwnership, ownerEvents bool, eventMask EventMask, cursor Cursor, time_ uint32) GrabStatus
	// SetAxisUse specifies how an axis of a device is used.
	//
	// This method is inherited from Device
	SetAxisUse(index_ uint, use AxisUse)
	// SetKey specifies the X key event to generate when a macro button of a
	// device is pressed.
	//
	// This method is inherited from Device
	SetKey(index_ uint, keyval uint, modifiers ModifierType)
	// SetMode sets a the mode of an input device. The mode controls if the
	// device is active and whether the device’s range is mapped to the entire
	// screen or to a single window.
	//
	// Note: This is only meaningful for floating devices, master devices (and
	// slaves connected to these) drive the pointer cursor, which is not limited
	// by the input mode.
	//
	// This method is inherited from Device
	SetMode(mode InputMode) bool
	// Ungrab: release any grab on @device.
	//
	// Deprecated: since version 3.20.
	//
	// This method is inherited from Device
	Ungrab(time_ uint32)
	// Warp warps @device in @display to the point @x,@y on the screen @screen,
	// unless the device is confined to a window by a grab, in which case it
	// will be moved as far as allowed by the grab. Warping the pointer creates
	// events as if the user had moved the mouse instantaneously to the
	// destination.
	//
	// Note that the pointer should normally be under the control of the user.
	// This function was added to cover some rare use cases like keyboard
	// navigation support for the color picker in the ColorSelectionDialog.
	//
	// This method is inherited from Device
	Warp(screen Screen, x int, y int)

	// FeatureGroup returns the group the given @feature and @idx belong to, or
	// -1 if feature/index do not exist in @pad.
	FeatureGroup(feature DevicePadFeature, featureIdx int) int
	// GroupNModes returns the number of modes that @group may have.
	GroupNModes(groupIdx int) int
	// NFeatures returns the number of features a tablet pad has.
	NFeatures(feature DevicePadFeature) int
	// NGroups returns the number of groups this pad device has. Pads have at
	// least one group. A pad group is a subcollection of buttons/strip/rings
	// that is affected collectively by a same current mode.
	NGroups() int
}

// devicePad implements the DevicePad interface.
type devicePad struct {
	*externglib.Object
}

var _ DevicePad = (*devicePad)(nil)

// WrapDevicePad wraps a GObject to a type that implements
// interface DevicePad. It is primarily used internally.
func WrapDevicePad(obj *externglib.Object) DevicePad {
	return devicePad{obj}
}

func marshalDevicePad(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDevicePad(obj), nil
}

func (d devicePad) AsDevice() Device {
	return WrapDevice(gextras.InternObject(d))
}

func (d devicePad) GetAssociatedDevice() Device {
	return WrapDevice(gextras.InternObject(d)).GetAssociatedDevice()
}

func (d devicePad) GetAxes() AxisFlags {
	return WrapDevice(gextras.InternObject(d)).GetAxes()
}

func (d devicePad) GetAxisUse(index_ uint) AxisUse {
	return WrapDevice(gextras.InternObject(d)).GetAxisUse(index_)
}

func (d devicePad) GetDeviceType() DeviceType {
	return WrapDevice(gextras.InternObject(d)).GetDeviceType()
}

func (d devicePad) GetDisplay() Display {
	return WrapDevice(gextras.InternObject(d)).GetDisplay()
}

func (d devicePad) GetHasCursor() bool {
	return WrapDevice(gextras.InternObject(d)).GetHasCursor()
}

func (d devicePad) GetKey(index_ uint) (uint, ModifierType, bool) {
	return WrapDevice(gextras.InternObject(d)).GetKey(index_)
}

func (d devicePad) GetLastEventWindow() Window {
	return WrapDevice(gextras.InternObject(d)).GetLastEventWindow()
}

func (d devicePad) GetMode() InputMode {
	return WrapDevice(gextras.InternObject(d)).GetMode()
}

func (d devicePad) GetNAxes() int {
	return WrapDevice(gextras.InternObject(d)).GetNAxes()
}

func (d devicePad) GetNKeys() int {
	return WrapDevice(gextras.InternObject(d)).GetNKeys()
}

func (d devicePad) GetName() string {
	return WrapDevice(gextras.InternObject(d)).GetName()
}

func (d devicePad) GetPosition() (screen Screen, x int, y int) {
	return WrapDevice(gextras.InternObject(d)).GetPosition()
}

func (d devicePad) GetPositionDouble() (screen Screen, x float64, y float64) {
	return WrapDevice(gextras.InternObject(d)).GetPositionDouble()
}

func (d devicePad) GetProductID() string {
	return WrapDevice(gextras.InternObject(d)).GetProductID()
}

func (d devicePad) GetSeat() Seat {
	return WrapDevice(gextras.InternObject(d)).GetSeat()
}

func (d devicePad) GetSource() InputSource {
	return WrapDevice(gextras.InternObject(d)).GetSource()
}

func (d devicePad) GetVendorID() string {
	return WrapDevice(gextras.InternObject(d)).GetVendorID()
}

func (d devicePad) GetWindowAtPosition() (winX int, winY int, window Window) {
	return WrapDevice(gextras.InternObject(d)).GetWindowAtPosition()
}

func (d devicePad) GetWindowAtPositionDouble() (winX float64, winY float64, window Window) {
	return WrapDevice(gextras.InternObject(d)).GetWindowAtPositionDouble()
}

func (d devicePad) Grab(window Window, grabOwnership GrabOwnership, ownerEvents bool, eventMask EventMask, cursor Cursor, time_ uint32) GrabStatus {
	return WrapDevice(gextras.InternObject(d)).Grab(window, grabOwnership, ownerEvents, eventMask, cursor, time_)
}

func (d devicePad) SetAxisUse(index_ uint, use AxisUse) {
	WrapDevice(gextras.InternObject(d)).SetAxisUse(index_, use)
}

func (d devicePad) SetKey(index_ uint, keyval uint, modifiers ModifierType) {
	WrapDevice(gextras.InternObject(d)).SetKey(index_, keyval, modifiers)
}

func (d devicePad) SetMode(mode InputMode) bool {
	return WrapDevice(gextras.InternObject(d)).SetMode(mode)
}

func (d devicePad) Ungrab(time_ uint32) {
	WrapDevice(gextras.InternObject(d)).Ungrab(time_)
}

func (d devicePad) Warp(screen Screen, x int, y int) {
	WrapDevice(gextras.InternObject(d)).Warp(screen, x, y)
}

func (p devicePad) FeatureGroup(feature DevicePadFeature, featureIdx int) int {
	var _arg0 *C.GdkDevicePad       // out
	var _arg1 C.GdkDevicePadFeature // out
	var _arg2 C.gint                // out
	var _cret C.gint                // in

	_arg0 = (*C.GdkDevicePad)(unsafe.Pointer(p.Native()))
	_arg1 = C.GdkDevicePadFeature(feature)
	_arg2 = C.gint(featureIdx)

	_cret = C.gdk_device_pad_get_feature_group(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (p devicePad) GroupNModes(groupIdx int) int {
	var _arg0 *C.GdkDevicePad // out
	var _arg1 C.gint          // out
	var _cret C.gint          // in

	_arg0 = (*C.GdkDevicePad)(unsafe.Pointer(p.Native()))
	_arg1 = C.gint(groupIdx)

	_cret = C.gdk_device_pad_get_group_n_modes(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (p devicePad) NFeatures(feature DevicePadFeature) int {
	var _arg0 *C.GdkDevicePad       // out
	var _arg1 C.GdkDevicePadFeature // out
	var _cret C.gint                // in

	_arg0 = (*C.GdkDevicePad)(unsafe.Pointer(p.Native()))
	_arg1 = C.GdkDevicePadFeature(feature)

	_cret = C.gdk_device_pad_get_n_features(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (p devicePad) NGroups() int {
	var _arg0 *C.GdkDevicePad // out
	var _cret C.gint          // in

	_arg0 = (*C.GdkDevicePad)(unsafe.Pointer(p.Native()))

	_cret = C.gdk_device_pad_get_n_groups(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}
