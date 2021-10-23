// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
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
