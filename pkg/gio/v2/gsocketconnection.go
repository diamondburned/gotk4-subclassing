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
		{T: externglib.Type(C.g_socket_connection_get_type()), F: marshalSocketConnection},
	})
}

// SocketConnection is a OStream for a connected socket. They can be created
// either by Client when connecting to a host, or by Listener when accepting a
// new client.
//
// The type of the Connection object returned from these calls depends on the
// type of the underlying socket that is in use. For instance, for a TCP/IP
// connection it will be a Connection.
//
// Choosing what type of object to construct is done with the socket connection
// factory, and it is possible for 3rd parties to register custom socket
// connection types for specific combination of socket family/type/protocol
// using g_socket_connection_factory_register_type().
//
// To close a Connection, use g_io_stream_close(). Closing both substreams of
// the OStream separately will not close the underlying #GSocket.
type SocketConnection interface {
	IOStream

	// Connect: connect @connection to the specified remote address.
	Connect(address SocketAddress, cancellable Cancellable) error
	// ConnectAsync: asynchronously connect @connection to the specified remote
	// address.
	//
	// This clears the #GSocket:blocking flag on @connection's underlying socket
	// if it is currently set.
	//
	// Use g_socket_connection_connect_finish() to retrieve the result.
	ConnectAsync()
	// ConnectFinish gets the result of a g_socket_connection_connect_async()
	// call.
	ConnectFinish(result AsyncResult) error
	// LocalAddress: try to get the local address of a socket connection.
	LocalAddress() (socketAddress SocketAddress, goerr error)
	// RemoteAddress: try to get the remote address of a socket connection.
	//
	// Since GLib 2.40, when used with g_socket_client_connect() or
	// g_socket_client_connect_async(), during emission of
	// G_SOCKET_CLIENT_CONNECTING, this function will return the remote address
	// that will be used for the connection. This allows applications to print
	// e.g. "Connecting to example.com (10.42.77.3)...".
	RemoteAddress() (socketAddress SocketAddress, goerr error)
	// Socket gets the underlying #GSocket object of the connection. This can be
	// useful if you want to do something unusual on it not supported by the
	// Connection APIs.
	Socket() Socket
	// IsConnected checks if @connection is connected. This is equivalent to
	// calling g_socket_is_connected() on @connection's underlying #GSocket.
	IsConnected() bool
}

// socketConnection implements the SocketConnection interface.
type socketConnection struct {
	IOStream
}

var _ SocketConnection = (*socketConnection)(nil)

// WrapSocketConnection wraps a GObject to the right type. It is
// primarily used internally.
func WrapSocketConnection(obj *externglib.Object) SocketConnection {
	return SocketConnection{
		IOStream: WrapIOStream(obj),
	}
}

func marshalSocketConnection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSocketConnection(obj), nil
}

// Connect: connect @connection to the specified remote address.
func (c socketConnection) Connect(address SocketAddress, cancellable Cancellable) error {
	var arg0 *C.GSocketConnection
	var arg1 *C.GSocketAddress
	var arg2 *C.GCancellable

	arg0 = (*C.GSocketConnection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GSocketAddress)(unsafe.Pointer(address.Native()))
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cerr *C.GError

	C.g_socket_connection_connect(arg0, arg1, arg2, cerr)

	var goerr error

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// ConnectAsync: asynchronously connect @connection to the specified remote
// address.
//
// This clears the #GSocket:blocking flag on @connection's underlying socket
// if it is currently set.
//
// Use g_socket_connection_connect_finish() to retrieve the result.
func (c socketConnection) ConnectAsync() {
	var arg0 *C.GSocketConnection

	arg0 = (*C.GSocketConnection)(unsafe.Pointer(c.Native()))

	C.g_socket_connection_connect_async(arg0)
}

// ConnectFinish gets the result of a g_socket_connection_connect_async()
// call.
func (c socketConnection) ConnectFinish(result AsyncResult) error {
	var arg0 *C.GSocketConnection
	var arg1 *C.GAsyncResult

	arg0 = (*C.GSocketConnection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var cerr *C.GError

	C.g_socket_connection_connect_finish(arg0, arg1, cerr)

	var goerr error

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// LocalAddress: try to get the local address of a socket connection.
func (c socketConnection) LocalAddress() (socketAddress SocketAddress, goerr error) {
	var arg0 *C.GSocketConnection

	arg0 = (*C.GSocketConnection)(unsafe.Pointer(c.Native()))

	var cret *C.GSocketAddress
	var cerr *C.GError

	cret = C.g_socket_connection_get_local_address(arg0, cerr)

	var socketAddress SocketAddress
	var goerr error

	socketAddress = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(SocketAddress)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return socketAddress, goerr
}

// RemoteAddress: try to get the remote address of a socket connection.
//
// Since GLib 2.40, when used with g_socket_client_connect() or
// g_socket_client_connect_async(), during emission of
// G_SOCKET_CLIENT_CONNECTING, this function will return the remote address
// that will be used for the connection. This allows applications to print
// e.g. "Connecting to example.com (10.42.77.3)...".
func (c socketConnection) RemoteAddress() (socketAddress SocketAddress, goerr error) {
	var arg0 *C.GSocketConnection

	arg0 = (*C.GSocketConnection)(unsafe.Pointer(c.Native()))

	var cret *C.GSocketAddress
	var cerr *C.GError

	cret = C.g_socket_connection_get_remote_address(arg0, cerr)

	var socketAddress SocketAddress
	var goerr error

	socketAddress = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(SocketAddress)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return socketAddress, goerr
}

// Socket gets the underlying #GSocket object of the connection. This can be
// useful if you want to do something unusual on it not supported by the
// Connection APIs.
func (c socketConnection) Socket() Socket {
	var arg0 *C.GSocketConnection

	arg0 = (*C.GSocketConnection)(unsafe.Pointer(c.Native()))

	var cret *C.GSocket

	cret = C.g_socket_connection_get_socket(arg0)

	var socket Socket

	socket = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Socket)

	return socket
}

// IsConnected checks if @connection is connected. This is equivalent to
// calling g_socket_is_connected() on @connection's underlying #GSocket.
func (c socketConnection) IsConnected() bool {
	var arg0 *C.GSocketConnection

	arg0 = (*C.GSocketConnection)(unsafe.Pointer(c.Native()))

	var cret C.gboolean

	cret = C.g_socket_connection_is_connected(arg0)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}
