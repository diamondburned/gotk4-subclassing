// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
)

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

// DBusErrorEncodeGError creates a D-Bus error name to use for error. If error
// matches a registered error (cf. g_dbus_error_register_error()), the
// corresponding D-Bus error name will be returned.
//
// Otherwise the a name of the form
// org.gtk.GDBus.UnmappedGError.Quark._ESCAPED_QUARK_NAME.Code_ERROR_CODE will
// be used. This allows other GDBus applications to map the error on the wire
// back to a #GError using g_dbus_error_new_for_dbus_error().
//
// This function is typically only used in object mappings to put a #GError on
// the wire. Regular applications should not use it.
func DBusErrorEncodeGError(err error) string {
	var _arg1 *C.GError // out
	var _cret *C.gchar  // in

	_arg1 = (*C.GError)(gerror.New(err))

	_cret = C.g_dbus_error_encode_gerror(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// DBusErrorGetRemoteError gets the D-Bus error name used for error, if any.
//
// This function is guaranteed to return a D-Bus error name for all #GErrors
// returned from functions handling remote method calls (e.g.
// g_dbus_connection_call_finish()) unless g_dbus_error_strip_remote_error() has
// been used on error.
func DBusErrorGetRemoteError(err error) string {
	var _arg1 *C.GError // out
	var _cret *C.gchar  // in

	_arg1 = (*C.GError)(gerror.New(err))

	_cret = C.g_dbus_error_get_remote_error(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// DBusErrorIsRemoteError checks if error represents an error received via D-Bus
// from a remote peer. If so, use g_dbus_error_get_remote_error() to get the
// name of the error.
func DBusErrorIsRemoteError(err error) bool {
	var _arg1 *C.GError  // out
	var _cret C.gboolean // in

	_arg1 = (*C.GError)(gerror.New(err))

	_cret = C.g_dbus_error_is_remote_error(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DBusErrorNewForDBusError creates a #GError based on the contents of
// dbus_error_name and dbus_error_message.
//
// Errors registered with g_dbus_error_register_error() will be looked up using
// dbus_error_name and if a match is found, the error domain and code is used.
// Applications can use g_dbus_error_get_remote_error() to recover
// dbus_error_name.
//
// If a match against a registered error is not found and the D-Bus error name
// is in a form as returned by g_dbus_error_encode_gerror() the error domain and
// code encoded in the name is used to create the #GError. Also, dbus_error_name
// is added to the error message such that it can be recovered with
// g_dbus_error_get_remote_error().
//
// Otherwise, a #GError with the error code G_IO_ERROR_DBUS_ERROR in the
// IO_ERROR error domain is returned. Also, dbus_error_name is added to the
// error message such that it can be recovered with
// g_dbus_error_get_remote_error().
//
// In all three cases, dbus_error_name can always be recovered from the returned
// #GError using the g_dbus_error_get_remote_error() function (unless
// g_dbus_error_strip_remote_error() hasn't been used on the returned error).
//
// This function is typically only used in object mappings to prepare #GError
// instances for applications. Regular applications should not use it.
func DBusErrorNewForDBusError(dbusErrorName string, dbusErrorMessage string) error {
	var _arg1 *C.gchar  // out
	var _arg2 *C.gchar  // out
	var _cret *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(dbusErrorName)))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(dbusErrorMessage)))

	_cret = C.g_dbus_error_new_for_dbus_error(_arg1, _arg2)

	var _err error // out

	_err = gerror.Take(unsafe.Pointer(_cret))

	return _err
}

// DBusErrorRegisterErrorDomain: helper function for associating a #GError error
// domain with D-Bus error names.
//
// While quark_volatile has a volatile qualifier, this is a historical artifact
// and the argument passed to it should not be volatile.
func DBusErrorRegisterErrorDomain(errorDomainQuarkName string, quarkVolatile *uint, entries []DBusErrorEntry) {
	var _arg1 *C.gchar // out
	var _arg2 *C.gsize // out
	var _arg3 *C.GDBusErrorEntry
	var _arg4 C.guint

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(errorDomainQuarkName)))
	_arg2 = (*C.gsize)(unsafe.Pointer(quarkVolatile))
	_arg4 = (C.guint)(len(entries))
	if len(entries) > 0 {
		_arg3 = (*C.GDBusErrorEntry)(unsafe.Pointer(&entries[0]))
	}

	C.g_dbus_error_register_error_domain(_arg1, _arg2, _arg3, _arg4)
}

// DBusErrorStripRemoteError looks for extra information in the error message
// used to recover the D-Bus error name and strips it if found. If stripped, the
// message field in error will correspond exactly to what was received on the
// wire.
//
// This is typically used when presenting errors to the end user.
func DBusErrorStripRemoteError(err error) bool {
	var _arg1 *C.GError  // out
	var _cret C.gboolean // in

	_arg1 = (*C.GError)(gerror.New(err))

	_cret = C.g_dbus_error_strip_remote_error(_arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DBusErrorEntry: struct used in g_dbus_error_register_error_domain().
type DBusErrorEntry struct {
	native C.GDBusErrorEntry
}

// Native returns the underlying C source pointer.
func (d *DBusErrorEntry) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}
