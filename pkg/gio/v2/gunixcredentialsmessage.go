// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

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
		{T: externglib.Type(C.g_unix_credentials_message_get_type()), F: marshalUnixCredentialsMessage},
	})
}

// UnixCredentialsMessage: this ControlMessage contains a #GCredentials
// instance. It may be sent using g_socket_send_message() and received using
// g_socket_receive_message() over UNIX sockets (ie: sockets in the
// G_SOCKET_FAMILY_UNIX family).
//
// For an easier way to send and receive credentials over stream-oriented UNIX
// sockets, see g_unix_connection_send_credentials() and
// g_unix_connection_receive_credentials(). To receive credentials of a foreign
// process connected to a socket, use g_socket_get_credentials().
type UnixCredentialsMessage interface {
	SocketControlMessage

	// Credentials gets the credentials stored in @message.
	Credentials() Credentials
}

// unixCredentialsMessage implements the UnixCredentialsMessage interface.
type unixCredentialsMessage struct {
	SocketControlMessage
}

var _ UnixCredentialsMessage = (*unixCredentialsMessage)(nil)

// WrapUnixCredentialsMessage wraps a GObject to the right type. It is
// primarily used internally.
func WrapUnixCredentialsMessage(obj *externglib.Object) UnixCredentialsMessage {
	return UnixCredentialsMessage{
		SocketControlMessage: WrapSocketControlMessage(obj),
	}
}

func marshalUnixCredentialsMessage(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapUnixCredentialsMessage(obj), nil
}

// NewUnixCredentialsMessage constructs a class UnixCredentialsMessage.
func NewUnixCredentialsMessage() UnixCredentialsMessage {
	var cret C.GUnixCredentialsMessage

	cret = C.g_unix_credentials_message_new()

	var unixCredentialsMessage UnixCredentialsMessage

	unixCredentialsMessage = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(UnixCredentialsMessage)

	return unixCredentialsMessage
}

// NewUnixCredentialsMessageWithCredentials constructs a class UnixCredentialsMessage.
func NewUnixCredentialsMessageWithCredentials(credentials Credentials) UnixCredentialsMessage {
	var arg1 *C.GCredentials

	arg1 = (*C.GCredentials)(unsafe.Pointer(credentials.Native()))

	var cret C.GUnixCredentialsMessage

	cret = C.g_unix_credentials_message_new_with_credentials(arg1)

	var unixCredentialsMessage UnixCredentialsMessage

	unixCredentialsMessage = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(UnixCredentialsMessage)

	return unixCredentialsMessage
}

// Credentials gets the credentials stored in @message.
func (m unixCredentialsMessage) Credentials() Credentials {
	var arg0 *C.GUnixCredentialsMessage

	arg0 = (*C.GUnixCredentialsMessage)(unsafe.Pointer(m.Native()))

	var cret *C.GCredentials

	cret = C.g_unix_credentials_message_get_credentials(arg0)

	var credentials Credentials

	credentials = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Credentials)

	return credentials
}
