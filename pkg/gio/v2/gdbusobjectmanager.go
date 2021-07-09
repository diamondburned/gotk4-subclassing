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
		{T: externglib.Type(C.g_dbus_object_manager_get_type()), F: marshalDBusObjectManager},
	})
}

// DBusObjectManagerOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type DBusObjectManagerOverrider interface {
	// Interface gets the interface proxy for @interface_name at @object_path,
	// if any.
	Interface(objectPath string, interfaceName string) *DBusInterfaceInterface
	// GetObject gets the BusObjectProxy at @object_path, if any.
	GetObject(objectPath string) *DBusObjectInterface
	// ObjectPath gets the object path that @manager is for.
	ObjectPath() string
	InterfaceAdded(object DBusObject, interface_ DBusInterface)
	InterfaceRemoved(object DBusObject, interface_ DBusInterface)
	ObjectAdded(object DBusObject)
	ObjectRemoved(object DBusObject)
}

// DBusObjectManager: the BusObjectManager type is the base type for service-
// and client-side implementations of the standardized
// org.freedesktop.DBus.ObjectManager
// (http://dbus.freedesktop.org/doc/dbus-specification.html#standard-interfaces-objectmanager)
// interface.
//
// See BusObjectManagerClient for the client-side implementation and
// BusObjectManagerServer for the service-side implementation.
type DBusObjectManager interface {
	gextras.Objector

	// Interface gets the interface proxy for @interface_name at @object_path,
	// if any.
	Interface(objectPath string, interfaceName string) *DBusInterfaceInterface
	// GetObject gets the BusObjectProxy at @object_path, if any.
	GetObject(objectPath string) *DBusObjectInterface
	// ObjectPath gets the object path that @manager is for.
	ObjectPath() string
}

// DBusObjectManagerInterface implements the DBusObjectManager interface.
type DBusObjectManagerInterface struct {
	*externglib.Object
}

var _ DBusObjectManager = (*DBusObjectManagerInterface)(nil)

func wrapDBusObjectManager(obj *externglib.Object) DBusObjectManager {
	return &DBusObjectManagerInterface{
		Object: obj,
	}
}

func marshalDBusObjectManager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDBusObjectManager(obj), nil
}

// Interface gets the interface proxy for @interface_name at @object_path, if
// any.
func (m *DBusObjectManagerInterface) Interface(objectPath string, interfaceName string) *DBusInterfaceInterface {
	var _arg0 *C.GDBusObjectManager // out
	var _arg1 *C.gchar              // out
	var _arg2 *C.gchar              // out
	var _cret *C.GDBusInterface     // in

	_arg0 = (*C.GDBusObjectManager)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(interfaceName))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_dbus_object_manager_get_interface(_arg0, _arg1, _arg2)

	var _dBusInterface *DBusInterfaceInterface // out

	_dBusInterface = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*DBusInterfaceInterface)

	return _dBusInterface
}

// GetObject gets the BusObjectProxy at @object_path, if any.
func (m *DBusObjectManagerInterface) GetObject(objectPath string) *DBusObjectInterface {
	var _arg0 *C.GDBusObjectManager // out
	var _arg1 *C.gchar              // out
	var _cret *C.GDBusObject        // in

	_arg0 = (*C.GDBusObjectManager)(unsafe.Pointer(m.Native()))
	_arg1 = (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_object_manager_get_object(_arg0, _arg1)

	var _dBusObject *DBusObjectInterface // out

	_dBusObject = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*DBusObjectInterface)

	return _dBusObject
}

// ObjectPath gets the object path that @manager is for.
func (m *DBusObjectManagerInterface) ObjectPath() string {
	var _arg0 *C.GDBusObjectManager // out
	var _cret *C.gchar              // in

	_arg0 = (*C.GDBusObjectManager)(unsafe.Pointer(m.Native()))

	_cret = C.g_dbus_object_manager_get_object_path(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}
