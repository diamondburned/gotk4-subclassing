// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_device_tool_type_get_type()), F: marshalDeviceToolType},
		{T: externglib.Type(C.gdk_device_tool_get_type()), F: marshalDeviceTooler},
	})
}

// DeviceToolType indicates the specific type of tool being used being a tablet.
// Such as an airbrush, pencil, etc.
type DeviceToolType C.gint

const (
	// DeviceToolTypeUnknown: tool is of an unknown type.
	DeviceToolTypeUnknown DeviceToolType = iota
	// DeviceToolTypePen: tool is a standard tablet stylus.
	DeviceToolTypePen
	// DeviceToolTypeEraser: tool is standard tablet eraser.
	DeviceToolTypeEraser
	// DeviceToolTypeBrush: tool is a brush stylus.
	DeviceToolTypeBrush
	// DeviceToolTypePencil: tool is a pencil stylus.
	DeviceToolTypePencil
	// DeviceToolTypeAirbrush: tool is an airbrush stylus.
	DeviceToolTypeAirbrush
	// DeviceToolTypeMouse: tool is a mouse.
	DeviceToolTypeMouse
	// DeviceToolTypeLens: tool is a lens cursor.
	DeviceToolTypeLens
)

func marshalDeviceToolType(p uintptr) (interface{}, error) {
	return DeviceToolType(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for DeviceToolType.
func (d DeviceToolType) String() string {
	switch d {
	case DeviceToolTypeUnknown:
		return "Unknown"
	case DeviceToolTypePen:
		return "Pen"
	case DeviceToolTypeEraser:
		return "Eraser"
	case DeviceToolTypeBrush:
		return "Brush"
	case DeviceToolTypePencil:
		return "Pencil"
	case DeviceToolTypeAirbrush:
		return "Airbrush"
	case DeviceToolTypeMouse:
		return "Mouse"
	case DeviceToolTypeLens:
		return "Lens"
	default:
		return fmt.Sprintf("DeviceToolType(%d)", d)
	}
}

// DeviceTool: physical tool associated to a GdkDevice.
type DeviceTool struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*DeviceTool)(nil)
)

func wrapDeviceTool(obj *externglib.Object) *DeviceTool {
	return &DeviceTool{
		Object: obj,
	}
}

func marshalDeviceTooler(p uintptr) (interface{}, error) {
	return wrapDeviceTool(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Axes gets the axes of the tool.
//
// The function returns the following values:
//
//    - axisFlags axes of tool.
//
func (tool *DeviceTool) Axes() AxisFlags {
	var _arg0 *C.GdkDeviceTool // out
	var _cret C.GdkAxisFlags   // in

	_arg0 = (*C.GdkDeviceTool)(unsafe.Pointer(tool.Native()))

	_cret = C.gdk_device_tool_get_axes(_arg0)
	runtime.KeepAlive(tool)

	var _axisFlags AxisFlags // out

	_axisFlags = AxisFlags(_cret)

	return _axisFlags
}

// HardwareID gets the hardware ID of this tool, or 0 if it's not known.
//
// When non-zero, the identificator is unique for the given tool model, meaning
// that two identical tools will share the same hardware_id, but will have
// different serial numbers (see gdk.DeviceTool.GetSerial()).
//
// This is a more concrete (and device specific) method to identify a
// GdkDeviceTool than gdk.DeviceTool.GetToolType(), as a tablet may support
// multiple devices with the same GdkDeviceToolType, but different hardware
// identificators.
//
// The function returns the following values:
//
//    - guint64: hardware identificator of this tool.
//
func (tool *DeviceTool) HardwareID() uint64 {
	var _arg0 *C.GdkDeviceTool // out
	var _cret C.guint64        // in

	_arg0 = (*C.GdkDeviceTool)(unsafe.Pointer(tool.Native()))

	_cret = C.gdk_device_tool_get_hardware_id(_arg0)
	runtime.KeepAlive(tool)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

// Serial gets the serial number of this tool.
//
// This value can be used to identify a physical tool (eg. a tablet pen) across
// program executions.
//
// The function returns the following values:
//
//    - guint64: serial ID for this tool.
//
func (tool *DeviceTool) Serial() uint64 {
	var _arg0 *C.GdkDeviceTool // out
	var _cret C.guint64        // in

	_arg0 = (*C.GdkDeviceTool)(unsafe.Pointer(tool.Native()))

	_cret = C.gdk_device_tool_get_serial(_arg0)
	runtime.KeepAlive(tool)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

// ToolType gets the GdkDeviceToolType of the tool.
//
// The function returns the following values:
//
//    - deviceToolType: physical type for this tool. This can be used to figure
//      out what sort of pen is being used, such as an airbrush or a pencil.
//
func (tool *DeviceTool) ToolType() DeviceToolType {
	var _arg0 *C.GdkDeviceTool    // out
	var _cret C.GdkDeviceToolType // in

	_arg0 = (*C.GdkDeviceTool)(unsafe.Pointer(tool.Native()))

	_cret = C.gdk_device_tool_get_tool_type(_arg0)
	runtime.KeepAlive(tool)

	var _deviceToolType DeviceToolType // out

	_deviceToolType = DeviceToolType(_cret)

	return _deviceToolType
}
