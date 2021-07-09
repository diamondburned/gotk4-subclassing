// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_volume_monitor_get_type()), F: marshalVolumeMonitor},
	})
}

// VolumeMonitorOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type VolumeMonitorOverrider interface {
	DriveChanged(drive Drive)
	DriveConnected(drive Drive)
	DriveDisconnected(drive Drive)
	DriveEjectButton(drive Drive)
	DriveStopButton(drive Drive)
	// MountForUUID finds a #GMount object by its UUID (see g_mount_get_uuid())
	MountForUUID(uuid string) *MountInterface
	// VolumeForUUID finds a #GVolume object by its UUID (see
	// g_volume_get_uuid())
	VolumeForUUID(uuid string) *VolumeInterface
	MountAdded(mount Mount)
	MountChanged(mount Mount)
	MountPreUnmount(mount Mount)
	MountRemoved(mount Mount)
	VolumeAdded(volume Volume)
	VolumeChanged(volume Volume)
	VolumeRemoved(volume Volume)
}

// VolumeMonitor is for listing the user interesting devices and volumes on the
// computer. In other words, what a file selector or file manager would show in
// a sidebar.
//
// Monitor is not [thread-default-context
// aware][g-main-context-push-thread-default], and so should not be used other
// than from the main thread, with no thread-default-context active.
//
// In order to receive updates about volumes and mounts monitored through GVFS,
// a main loop must be running.
type VolumeMonitor interface {
	gextras.Objector

	// MountForUUID finds a #GMount object by its UUID (see g_mount_get_uuid())
	MountForUUID(uuid string) *MountInterface
	// VolumeForUUID finds a #GVolume object by its UUID (see
	// g_volume_get_uuid())
	VolumeForUUID(uuid string) *VolumeInterface
}

// VolumeMonitorClass implements the VolumeMonitor interface.
type VolumeMonitorClass struct {
	*externglib.Object
}

var _ VolumeMonitor = (*VolumeMonitorClass)(nil)

func wrapVolumeMonitor(obj *externglib.Object) VolumeMonitor {
	return &VolumeMonitorClass{
		Object: obj,
	}
}

func marshalVolumeMonitor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapVolumeMonitor(obj), nil
}

// MountForUUID finds a #GMount object by its UUID (see g_mount_get_uuid())
func (v *VolumeMonitorClass) MountForUUID(uuid string) *MountInterface {
	var _arg0 *C.GVolumeMonitor // out
	var _arg1 *C.char           // out
	var _cret *C.GMount         // in

	_arg0 = (*C.GVolumeMonitor)(unsafe.Pointer(v.Native()))
	_arg1 = (*C.char)(C.CString(uuid))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_volume_monitor_get_mount_for_uuid(_arg0, _arg1)

	var _mount *MountInterface // out

	_mount = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*MountInterface)

	return _mount
}

// VolumeForUUID finds a #GVolume object by its UUID (see g_volume_get_uuid())
func (v *VolumeMonitorClass) VolumeForUUID(uuid string) *VolumeInterface {
	var _arg0 *C.GVolumeMonitor // out
	var _arg1 *C.char           // out
	var _cret *C.GVolume        // in

	_arg0 = (*C.GVolumeMonitor)(unsafe.Pointer(v.Native()))
	_arg1 = (*C.char)(C.CString(uuid))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_volume_monitor_get_volume_for_uuid(_arg0, _arg1)

	var _volume *VolumeInterface // out

	_volume = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*VolumeInterface)

	return _volume
}
