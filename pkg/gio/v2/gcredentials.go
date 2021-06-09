// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
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
		{T: externglib.Type(C.g_credentials_get_type()), F: marshalCredentials},
	})
}

// Credentials: the #GCredentials type is a reference-counted wrapper for native
// credentials. This information is typically used for identifying,
// authenticating and authorizing other processes.
//
// Some operating systems supports looking up the credentials of the remote peer
// of a communication endpoint - see e.g. g_socket_get_credentials().
//
// Some operating systems supports securely sending and receiving credentials
// over a Unix Domain Socket, see CredentialsMessage,
// g_unix_connection_send_credentials() and
// g_unix_connection_receive_credentials() for details.
//
// On Linux, the native credential type is a `struct ucred` - see the unix(7)
// man page for details. This corresponds to G_CREDENTIALS_TYPE_LINUX_UCRED.
//
// On Apple operating systems (including iOS, tvOS, and macOS), the native
// credential type is a `struct xucred`. This corresponds to
// G_CREDENTIALS_TYPE_APPLE_XUCRED.
//
// On FreeBSD, Debian GNU/kFreeBSD, and GNU/Hurd, the native credential type is
// a `struct cmsgcred`. This corresponds to G_CREDENTIALS_TYPE_FREEBSD_CMSGCRED.
//
// On NetBSD, the native credential type is a `struct unpcbid`. This corresponds
// to G_CREDENTIALS_TYPE_NETBSD_UNPCBID.
//
// On OpenBSD, the native credential type is a `struct sockpeercred`. This
// corresponds to G_CREDENTIALS_TYPE_OPENBSD_SOCKPEERCRED.
//
// On Solaris (including OpenSolaris and its derivatives), the native credential
// type is a `ucred_t`. This corresponds to G_CREDENTIALS_TYPE_SOLARIS_UCRED.
type Credentials interface {
	gextras.Objector

	// Native gets a pointer to native credentials of type @native_type from
	// @credentials.
	//
	// It is a programming error (which will cause a warning to be logged) to
	// use this method if there is no #GCredentials support for the OS or if
	// @native_type isn't supported by the OS.
	Native(nativeType CredentialsType) interface{}
	// UnixPid tries to get the UNIX process identifier from @credentials. This
	// method is only available on UNIX platforms.
	//
	// This operation can fail if #GCredentials is not supported on the OS or if
	// the native credentials type does not contain information about the UNIX
	// process ID (for example this is the case for
	// G_CREDENTIALS_TYPE_APPLE_XUCRED).
	UnixPid() (gint int, goerr error)
	// UnixUser tries to get the UNIX user identifier from @credentials. This
	// method is only available on UNIX platforms.
	//
	// This operation can fail if #GCredentials is not supported on the OS or if
	// the native credentials type does not contain information about the UNIX
	// user.
	UnixUser() (guint uint, goerr error)
	// IsSameUser checks if @credentials and @other_credentials is the same
	// user.
	//
	// This operation can fail if #GCredentials is not supported on the the OS.
	IsSameUser(otherCredentials Credentials) error
	// SetNative copies the native credentials of type @native_type from @native
	// into @credentials.
	//
	// It is a programming error (which will cause a warning to be logged) to
	// use this method if there is no #GCredentials support for the OS or if
	// @native_type isn't supported by the OS.
	SetNative(nativeType CredentialsType, native interface{})
	// SetUnixUser tries to set the UNIX user identifier on @credentials. This
	// method is only available on UNIX platforms.
	//
	// This operation can fail if #GCredentials is not supported on the OS or if
	// the native credentials type does not contain information about the UNIX
	// user. It can also fail if the OS does not allow the use of "spoofed"
	// credentials.
	SetUnixUser(uid uint) error
	// String creates a human-readable textual representation of @credentials
	// that can be used in logging and debug messages. The format of the
	// returned string may change in future GLib release.
	String() string
}

// credentials implements the Credentials interface.
type credentials struct {
	gextras.Objector
}

var _ Credentials = (*credentials)(nil)

// WrapCredentials wraps a GObject to the right type. It is
// primarily used internally.
func WrapCredentials(obj *externglib.Object) Credentials {
	return Credentials{
		Objector: obj,
	}
}

func marshalCredentials(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCredentials(obj), nil
}

// NewCredentials constructs a class Credentials.
func NewCredentials() Credentials {
	var cret C.GCredentials

	cret = C.g_credentials_new()

	var credentials Credentials

	credentials = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Credentials)

	return credentials
}

// Native gets a pointer to native credentials of type @native_type from
// @credentials.
//
// It is a programming error (which will cause a warning to be logged) to
// use this method if there is no #GCredentials support for the OS or if
// @native_type isn't supported by the OS.
func (c credentials) Native(nativeType CredentialsType) interface{} {
	var arg0 *C.GCredentials
	var arg1 C.GCredentialsType

	arg0 = (*C.GCredentials)(unsafe.Pointer(c.Native()))
	arg1 = (C.GCredentialsType)(nativeType)

	var cret C.gpointer

	cret = C.g_credentials_get_native(arg0, arg1)

	var gpointer interface{}

	gpointer = (interface{})(cret)

	return gpointer
}

// UnixPid tries to get the UNIX process identifier from @credentials. This
// method is only available on UNIX platforms.
//
// This operation can fail if #GCredentials is not supported on the OS or if
// the native credentials type does not contain information about the UNIX
// process ID (for example this is the case for
// G_CREDENTIALS_TYPE_APPLE_XUCRED).
func (c credentials) UnixPid() (gint int, goerr error) {
	var arg0 *C.GCredentials

	arg0 = (*C.GCredentials)(unsafe.Pointer(c.Native()))

	var cret C.pid_t
	var cerr *C.GError

	cret = C.g_credentials_get_unix_pid(arg0, cerr)

	var gint int
	var goerr error

	gint = (int)(cret)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return gint, goerr
}

// UnixUser tries to get the UNIX user identifier from @credentials. This
// method is only available on UNIX platforms.
//
// This operation can fail if #GCredentials is not supported on the OS or if
// the native credentials type does not contain information about the UNIX
// user.
func (c credentials) UnixUser() (guint uint, goerr error) {
	var arg0 *C.GCredentials

	arg0 = (*C.GCredentials)(unsafe.Pointer(c.Native()))

	var cret C.uid_t
	var cerr *C.GError

	cret = C.g_credentials_get_unix_user(arg0, cerr)

	var guint uint
	var goerr error

	guint = (uint)(cret)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return guint, goerr
}

// IsSameUser checks if @credentials and @other_credentials is the same
// user.
//
// This operation can fail if #GCredentials is not supported on the the OS.
func (c credentials) IsSameUser(otherCredentials Credentials) error {
	var arg0 *C.GCredentials
	var arg1 *C.GCredentials

	arg0 = (*C.GCredentials)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GCredentials)(unsafe.Pointer(otherCredentials.Native()))

	var cerr *C.GError

	C.g_credentials_is_same_user(arg0, arg1, cerr)

	var goerr error

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// SetNative copies the native credentials of type @native_type from @native
// into @credentials.
//
// It is a programming error (which will cause a warning to be logged) to
// use this method if there is no #GCredentials support for the OS or if
// @native_type isn't supported by the OS.
func (c credentials) SetNative(nativeType CredentialsType, native interface{}) {
	var arg0 *C.GCredentials
	var arg1 C.GCredentialsType
	var arg2 C.gpointer

	arg0 = (*C.GCredentials)(unsafe.Pointer(c.Native()))
	arg1 = (C.GCredentialsType)(nativeType)
	arg2 = C.gpointer(native)

	C.g_credentials_set_native(arg0, arg1, arg2)
}

// SetUnixUser tries to set the UNIX user identifier on @credentials. This
// method is only available on UNIX platforms.
//
// This operation can fail if #GCredentials is not supported on the OS or if
// the native credentials type does not contain information about the UNIX
// user. It can also fail if the OS does not allow the use of "spoofed"
// credentials.
func (c credentials) SetUnixUser(uid uint) error {
	var arg0 *C.GCredentials
	var arg1 C.uid_t

	arg0 = (*C.GCredentials)(unsafe.Pointer(c.Native()))
	arg1 = C.uid_t(uid)

	var cerr *C.GError

	C.g_credentials_set_unix_user(arg0, arg1, cerr)

	var goerr error

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// String creates a human-readable textual representation of @credentials
// that can be used in logging and debug messages. The format of the
// returned string may change in future GLib release.
func (c credentials) String() string {
	var arg0 *C.GCredentials

	arg0 = (*C.GCredentials)(unsafe.Pointer(c.Native()))

	var cret *C.gchar

	cret = C.g_credentials_to_string(arg0)

	var utf8 string

	utf8 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return utf8
}
