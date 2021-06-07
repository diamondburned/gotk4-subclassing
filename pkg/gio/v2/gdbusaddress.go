// Code generated by girgen. DO NOT EDIT.

package gio

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
import "C"

// DBusAddressEscapeValue: escape @string so it can appear in a D-Bus address as
// the value part of a key-value pair.
//
// For instance, if @string is `/run/bus-for-:0`, this function would return
// `/run/bus-for-3A0`, which could be used in a D-Bus address like
// `unix:nonce-tcp:host=127.0.0.1,port=42,noncefile=/run/bus-for-3A0`.
func DBusAddressEscapeValue(string string) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))

	C.g_dbus_address_escape_value(arg1)
}

// DBusAddressGetForBusSync: synchronously looks up the D-Bus address for the
// well-known message bus instance specified by @bus_type. This may involve
// using various platform specific mechanisms.
//
// The returned address will be in the D-Bus address format
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
func DBusAddressGetForBusSync(busType BusType, cancellable Cancellable) error {
	var arg1 C.GBusType
	var arg2 *C.GCancellable

	arg1 = (C.GBusType)(busType)
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var errout *C.GError
	var err error

	C.g_dbus_address_get_for_bus_sync(arg1, arg2, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// DBusAddressGetStream: asynchronously connects to an endpoint specified by
// @address and sets up the connection so it is in a state to run the
// client-side of the D-Bus authentication conversation. @address must be in the
// D-Bus address format
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
//
// When the operation is finished, @callback will be invoked. You can then call
// g_dbus_address_get_stream_finish() to get the result of the operation.
//
// This is an asynchronous failable function. See
// g_dbus_address_get_stream_sync() for the synchronous version.
func DBusAddressGetStream() {
	C.g_dbus_address_get_stream(arg1, arg2, arg3, arg4)
}

// DBusAddressGetStreamFinish finishes an operation started with
// g_dbus_address_get_stream().
func DBusAddressGetStreamFinish(res AsyncResult) (outGuid string, err error) {
	var arg1 *C.GAsyncResult

	arg1 = (*C.GAsyncResult)(unsafe.Pointer(res.Native()))

	var arg2 *C.gchar
	var outGuid string
	var errout *C.GError
	var err error

	C.g_dbus_address_get_stream_finish(arg1, &arg2, &errout)

	outGuid = C.GoString(&arg2)
	defer C.free(unsafe.Pointer(&arg2))
	err = gerror.Take(unsafe.Pointer(errout))

	return outGuid, err
}

// DBusAddressGetStreamSync: synchronously connects to an endpoint specified by
// @address and sets up the connection so it is in a state to run the
// client-side of the D-Bus authentication conversation. @address must be in the
// D-Bus address format
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
//
// This is a synchronous failable function. See g_dbus_address_get_stream() for
// the asynchronous version.
func DBusAddressGetStreamSync(address string, cancellable Cancellable) (outGuid string, err error) {
	var arg1 *C.gchar
	var arg3 *C.GCancellable

	arg1 = (*C.gchar)(C.CString(address))
	defer C.free(unsafe.Pointer(arg1))
	arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var arg2 *C.gchar
	var outGuid string
	var errout *C.GError
	var err error

	C.g_dbus_address_get_stream_sync(arg1, &arg2, arg3, &errout)

	outGuid = C.GoString(&arg2)
	defer C.free(unsafe.Pointer(&arg2))
	err = gerror.Take(unsafe.Pointer(errout))

	return outGuid, err
}

// DBusIsAddress checks if @string is a D-Bus address
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
//
// This doesn't check if @string is actually supported by BusServer or
// BusConnection - use g_dbus_is_supported_address() to do more checks.
func DBusIsAddress(string string) bool {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ok bool

	cret = C.g_dbus_is_address(arg1)

	if cret {
		ok = true
	}

	return ok
}

// DBusIsSupportedAddress: like g_dbus_is_address() but also checks if the
// library supports the transports in @string and that key/value pairs for each
// transport are valid. See the specification of the D-Bus address format
// (https://dbus.freedesktop.org/doc/dbus-specification.html#addresses).
func DBusIsSupportedAddress(string string) error {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))

	var errout *C.GError
	var err error

	C.g_dbus_is_supported_address(arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}
