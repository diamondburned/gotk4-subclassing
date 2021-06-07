// Code generated by girgen. DO NOT EDIT.

package gio

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
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

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_tls_password_get_type()), F: marshalTLSPassword},
	})
}

// TLSPassword holds a password used in TLS.
type TLSPassword interface {
	gextras.Objector

	// Description: get a description string about what the password will be
	// used for.
	Description(p TLSPassword)
	// Flags: get flags about the password.
	Flags(p TLSPassword)
	// Value: get the password value. If @length is not nil then it will be
	// filled in with the length of the password value. (Note that the password
	// value is not nul-terminated, so you can only pass nil for @length in
	// contexts where you know the password will have a certain fixed length.)
	Value(p TLSPassword, length uint)
	// Warning: get a user readable translated warning. Usually this warning is
	// a representation of the password flags returned from
	// g_tls_password_get_flags().
	Warning(p TLSPassword)
	// SetDescription: set a description string about what the password will be
	// used for.
	SetDescription(p TLSPassword, description string)
	// SetFlags: set flags about the password.
	SetFlags(p TLSPassword, flags TLSPasswordFlags)
	// SetValue: set the value for this password. The @value will be copied by
	// the password object.
	//
	// Specify the @length, for a non-nul-terminated password. Pass -1 as
	// @length if using a nul-terminated password, and @length will be
	// calculated automatically. (Note that the terminating nul is not
	// considered part of the password in this case.)
	SetValue(p TLSPassword)
	// SetWarning: set a user readable translated warning. Usually this warning
	// is a representation of the password flags returned from
	// g_tls_password_get_flags().
	SetWarning(p TLSPassword, warning string)
}

// tlsPassword implements the TLSPassword interface.
type tlsPassword struct {
	gextras.Objector
}

var _ TLSPassword = (*tlsPassword)(nil)

// WrapTLSPassword wraps a GObject to the right type. It is
// primarily used internally.
func WrapTLSPassword(obj *externglib.Object) TLSPassword {
	return TLSPassword{
		Objector: obj,
	}
}

func marshalTLSPassword(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTLSPassword(obj), nil
}

// NewTLSPassword constructs a class TLSPassword.
func NewTLSPassword(flags TLSPasswordFlags, description string) {
	var arg1 C.GTlsPasswordFlags
	var arg2 *C.gchar

	arg1 = (C.GTlsPasswordFlags)(flags)
	arg2 = (*C.gchar)(C.CString(description))
	defer C.free(unsafe.Pointer(arg2))

	C.g_tls_password_new(arg1, arg2)
}

// Description: get a description string about what the password will be
// used for.
func (p tlsPassword) Description(p TLSPassword) {
	var arg0 *C.GTlsPassword

	arg0 = (*C.GTlsPassword)(unsafe.Pointer(p.Native()))

	C.g_tls_password_get_description(arg0)
}

// Flags: get flags about the password.
func (p tlsPassword) Flags(p TLSPassword) {
	var arg0 *C.GTlsPassword

	arg0 = (*C.GTlsPassword)(unsafe.Pointer(p.Native()))

	C.g_tls_password_get_flags(arg0)
}

// Value: get the password value. If @length is not nil then it will be
// filled in with the length of the password value. (Note that the password
// value is not nul-terminated, so you can only pass nil for @length in
// contexts where you know the password will have a certain fixed length.)
func (p tlsPassword) Value(p TLSPassword, length uint) {
	var arg0 *C.GTlsPassword
	var arg1 *C.gsize

	arg0 = (*C.GTlsPassword)(unsafe.Pointer(p.Native()))
	arg1 = *C.gsize(length)

	C.g_tls_password_get_value(arg0, arg1)
}

// Warning: get a user readable translated warning. Usually this warning is
// a representation of the password flags returned from
// g_tls_password_get_flags().
func (p tlsPassword) Warning(p TLSPassword) {
	var arg0 *C.GTlsPassword

	arg0 = (*C.GTlsPassword)(unsafe.Pointer(p.Native()))

	C.g_tls_password_get_warning(arg0)
}

// SetDescription: set a description string about what the password will be
// used for.
func (p tlsPassword) SetDescription(p TLSPassword, description string) {
	var arg0 *C.GTlsPassword
	var arg1 *C.gchar

	arg0 = (*C.GTlsPassword)(unsafe.Pointer(p.Native()))
	arg1 = (*C.gchar)(C.CString(description))
	defer C.free(unsafe.Pointer(arg1))

	C.g_tls_password_set_description(arg0, arg1)
}

// SetFlags: set flags about the password.
func (p tlsPassword) SetFlags(p TLSPassword, flags TLSPasswordFlags) {
	var arg0 *C.GTlsPassword
	var arg1 C.GTlsPasswordFlags

	arg0 = (*C.GTlsPassword)(unsafe.Pointer(p.Native()))
	arg1 = (C.GTlsPasswordFlags)(flags)

	C.g_tls_password_set_flags(arg0, arg1)
}

// SetValue: set the value for this password. The @value will be copied by
// the password object.
//
// Specify the @length, for a non-nul-terminated password. Pass -1 as
// @length if using a nul-terminated password, and @length will be
// calculated automatically. (Note that the terminating nul is not
// considered part of the password in this case.)
func (p tlsPassword) SetValue(p TLSPassword) {
	var arg0 *C.GTlsPassword

	arg0 = (*C.GTlsPassword)(unsafe.Pointer(p.Native()))

	C.g_tls_password_set_value(arg0, arg1, arg2)
}

// SetWarning: set a user readable translated warning. Usually this warning
// is a representation of the password flags returned from
// g_tls_password_get_flags().
func (p tlsPassword) SetWarning(p TLSPassword, warning string) {
	var arg0 *C.GTlsPassword
	var arg1 *C.gchar

	arg0 = (*C.GTlsPassword)(unsafe.Pointer(p.Native()))
	arg1 = (*C.gchar)(C.CString(warning))
	defer C.free(unsafe.Pointer(arg1))

	C.g_tls_password_set_warning(arg0, arg1)
}
