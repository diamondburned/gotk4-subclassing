// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_device_pad_get_type()), F: marshalDevicePad},
	})
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
// will be notified through events of type K_PAD_GROUP_MODE.
type DevicePad interface {
	Device

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
	Device
}

var _ DevicePad = (*devicePad)(nil)

// WrapDevicePad wraps a GObject to a type that implements interface
// DevicePad. It is primarily used internally.
func WrapDevicePad(obj *externglib.Object) DevicePad {
	return DevicePad{
		Device: WrapDevice(obj),
	}
}

func marshalDevicePad(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDevicePad(obj), nil
}

// FeatureGroup returns the group the given @feature and @idx belong to, or
// -1 if feature/index do not exist in @pad.
func (p devicePad) FeatureGroup(feature DevicePadFeature, featureIdx int) int {
	var arg0 *C.GdkDevicePad
	var arg1 C.GdkDevicePadFeature
	var arg2 C.int

	arg0 = (*C.GdkDevicePad)(unsafe.Pointer(p.Native()))
	arg1 = (C.GdkDevicePadFeature)(feature)
	arg2 = C.int(featureIdx)

	var cret C.int

	cret = C.gdk_device_pad_get_feature_group(arg0, arg1, arg2)

	var gint int

	gint = (int)(cret)

	return gint
}

// GroupNModes returns the number of modes that @group may have.
func (p devicePad) GroupNModes(groupIdx int) int {
	var arg0 *C.GdkDevicePad
	var arg1 C.int

	arg0 = (*C.GdkDevicePad)(unsafe.Pointer(p.Native()))
	arg1 = C.int(groupIdx)

	var cret C.int

	cret = C.gdk_device_pad_get_group_n_modes(arg0, arg1)

	var gint int

	gint = (int)(cret)

	return gint
}

// NFeatures returns the number of features a tablet pad has.
func (p devicePad) NFeatures(feature DevicePadFeature) int {
	var arg0 *C.GdkDevicePad
	var arg1 C.GdkDevicePadFeature

	arg0 = (*C.GdkDevicePad)(unsafe.Pointer(p.Native()))
	arg1 = (C.GdkDevicePadFeature)(feature)

	var cret C.int

	cret = C.gdk_device_pad_get_n_features(arg0, arg1)

	var gint int

	gint = (int)(cret)

	return gint
}

// NGroups returns the number of groups this pad device has. Pads have at
// least one group. A pad group is a subcollection of buttons/strip/rings
// that is affected collectively by a same current mode.
func (p devicePad) NGroups() int {
	var arg0 *C.GdkDevicePad

	arg0 = (*C.GdkDevicePad)(unsafe.Pointer(p.Native()))

	var cret C.int

	cret = C.gdk_device_pad_get_n_groups(arg0)

	var gint int

	gint = (int)(cret)

	return gint
}
