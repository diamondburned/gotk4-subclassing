// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

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
		{T: externglib.Type(C.gdk_device_tool_type_get_type()), F: marshalDeviceToolType},
	})
}

// DeviceToolType indicates the specific type of tool being used being a tablet.
// Such as an airbrush, pencil, etc.
type DeviceToolType int

const (
	// unknown: tool is of an unknown type.
	DeviceToolTypeUnknown DeviceToolType = 0
	// pen: tool is a standard tablet stylus.
	DeviceToolTypePen DeviceToolType = 1
	// eraser: tool is standard tablet eraser.
	DeviceToolTypeEraser DeviceToolType = 2
	// brush: tool is a brush stylus.
	DeviceToolTypeBrush DeviceToolType = 3
	// pencil: tool is a pencil stylus.
	DeviceToolTypePencil DeviceToolType = 4
	// airbrush: tool is an airbrush stylus.
	DeviceToolTypeAirbrush DeviceToolType = 5
	// mouse: tool is a mouse.
	DeviceToolTypeMouse DeviceToolType = 6
	// lens: tool is a lens cursor.
	DeviceToolTypeLens DeviceToolType = 7
)

func marshalDeviceToolType(p uintptr) (interface{}, error) {
	return DeviceToolType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}