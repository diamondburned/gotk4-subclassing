// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gerror"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0 glib-2.0
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
		{T: externglib.Type(C.g_socket_listener_get_type()), F: marshalSocketListener},
	})
}

// SocketListener: a Listener is an object that keeps track of a set of server
// sockets and helps you accept sockets from any of the socket, either sync or
// async.
//
// Add addresses and ports to listen on using g_socket_listener_add_address()
// and g_socket_listener_add_inet_port(). These will be listened on until
// g_socket_listener_close() is called. Dropping your final reference to the
// Listener will not cause g_socket_listener_close() to be called implicitly, as
// some references to the Listener may be held internally.
//
// If you want to implement a network server, also look at Service and
// SocketService which are subclasses of Listener that make this even easier.
type SocketListener interface {
	gextras.Objector

	// AcceptAsync: this is the asynchronous version of
	// g_socket_listener_accept().
	//
	// When the operation is finished @callback will be called. You can then
	// call g_socket_listener_accept_socket() to get the result of the
	// operation.
	AcceptAsync(cancellable Cancellable, callback AsyncReadyCallback)
	// AcceptSocketAsync: this is the asynchronous version of
	// g_socket_listener_accept_socket().
	//
	// When the operation is finished @callback will be called. You can then
	// call g_socket_listener_accept_socket_finish() to get the result of the
	// operation.
	AcceptSocketAsync(cancellable Cancellable, callback AsyncReadyCallback)
	// AddAnyInetPort listens for TCP connections on any available port number
	// for both IPv6 and IPv4 (if each is available).
	//
	// This is useful if you need to have a socket for incoming connections but
	// don't care about the specific port number.
	//
	// @source_object will be passed out in the various calls to accept to
	// identify this particular source, which is useful if you're listening on
	// multiple addresses and do different things depending on what address is
	// connected to.
	AddAnyInetPort(sourceObject gextras.Objector) (uint16, error)
	// AddInetPort: helper function for g_socket_listener_add_address() that
	// creates a TCP/IP socket listening on IPv4 and IPv6 (if supported) on the
	// specified port on all interfaces.
	//
	// @source_object will be passed out in the various calls to accept to
	// identify this particular source, which is useful if you're listening on
	// multiple addresses and do different things depending on what address is
	// connected to.
	//
	// Call g_socket_listener_close() to stop listening on @port; this will not
	// be done automatically when you drop your final reference to @listener, as
	// references may be held internally.
	AddInetPort(port uint16, sourceObject gextras.Objector) error
	// AddSocket adds @socket to the set of sockets that we try to accept new
	// clients from. The socket must be bound to a local address and listened
	// to.
	//
	// @source_object will be passed out in the various calls to accept to
	// identify this particular source, which is useful if you're listening on
	// multiple addresses and do different things depending on what address is
	// connected to.
	//
	// The @socket will not be automatically closed when the @listener is
	// finalized unless the listener held the final reference to the socket.
	// Before GLib 2.42, the @socket was automatically closed on finalization of
	// the @listener, even if references to it were held elsewhere.
	AddSocket(socket Socket, sourceObject gextras.Objector) error
	// Close closes all the sockets in the listener.
	Close()
	// SetBacklog sets the listen backlog on the sockets in the listener. This
	// must be called before adding any sockets, addresses or ports to the
	// Listener (for example, by calling g_socket_listener_add_inet_port()) to
	// be effective.
	//
	// See g_socket_set_listen_backlog() for details
	SetBacklog(listenBacklog int)
}

// socketListener implements the SocketListener interface.
type socketListener struct {
	gextras.Objector
}

var _ SocketListener = (*socketListener)(nil)

// WrapSocketListener wraps a GObject to the right type. It is
// primarily used internally.
func WrapSocketListener(obj *externglib.Object) SocketListener {
	return SocketListener{
		Objector: obj,
	}
}

func marshalSocketListener(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSocketListener(obj), nil
}

// AcceptAsync: this is the asynchronous version of
// g_socket_listener_accept().
//
// When the operation is finished @callback will be called. You can then
// call g_socket_listener_accept_socket() to get the result of the
// operation.
func (l socketListener) AcceptAsync(cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GSocketListener    // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg2 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg3 = C.gpointer(box.Assign(callback))

	C.g_socket_listener_accept_async(_arg0, _arg1, _arg2, _arg3)
}

// AcceptSocketAsync: this is the asynchronous version of
// g_socket_listener_accept_socket().
//
// When the operation is finished @callback will be called. You can then
// call g_socket_listener_accept_socket_finish() to get the result of the
// operation.
func (l socketListener) AcceptSocketAsync(cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GSocketListener    // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg2 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg3 = C.gpointer(box.Assign(callback))

	C.g_socket_listener_accept_socket_async(_arg0, _arg1, _arg2, _arg3)
}

// AddAnyInetPort listens for TCP connections on any available port number
// for both IPv6 and IPv4 (if each is available).
//
// This is useful if you need to have a socket for incoming connections but
// don't care about the specific port number.
//
// @source_object will be passed out in the various calls to accept to
// identify this particular source, which is useful if you're listening on
// multiple addresses and do different things depending on what address is
// connected to.
func (l socketListener) AddAnyInetPort(sourceObject gextras.Objector) (uint16, error) {
	var _arg0 *C.GSocketListener // out
	var _arg1 *C.GObject         // out

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.GObject)(unsafe.Pointer(sourceObject.Native()))

	var _cret C.guint16 // in
	var _cerr *C.GError // in

	_cret = C.g_socket_listener_add_any_inet_port(_arg0, _arg1, &_cerr)

	var _guint16 uint16 // out
	var _goerr error    // out

	_guint16 = (uint16)(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _guint16, _goerr
}

// AddInetPort: helper function for g_socket_listener_add_address() that
// creates a TCP/IP socket listening on IPv4 and IPv6 (if supported) on the
// specified port on all interfaces.
//
// @source_object will be passed out in the various calls to accept to
// identify this particular source, which is useful if you're listening on
// multiple addresses and do different things depending on what address is
// connected to.
//
// Call g_socket_listener_close() to stop listening on @port; this will not
// be done automatically when you drop your final reference to @listener, as
// references may be held internally.
func (l socketListener) AddInetPort(port uint16, sourceObject gextras.Objector) error {
	var _arg0 *C.GSocketListener // out
	var _arg1 C.guint16          // out
	var _arg2 *C.GObject         // out

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(l.Native()))
	_arg1 = C.guint16(port)
	_arg2 = (*C.GObject)(unsafe.Pointer(sourceObject.Native()))

	var _cerr *C.GError // in

	C.g_socket_listener_add_inet_port(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// AddSocket adds @socket to the set of sockets that we try to accept new
// clients from. The socket must be bound to a local address and listened
// to.
//
// @source_object will be passed out in the various calls to accept to
// identify this particular source, which is useful if you're listening on
// multiple addresses and do different things depending on what address is
// connected to.
//
// The @socket will not be automatically closed when the @listener is
// finalized unless the listener held the final reference to the socket.
// Before GLib 2.42, the @socket was automatically closed on finalization of
// the @listener, even if references to it were held elsewhere.
func (l socketListener) AddSocket(socket Socket, sourceObject gextras.Objector) error {
	var _arg0 *C.GSocketListener // out
	var _arg1 *C.GSocket         // out
	var _arg2 *C.GObject         // out

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.GSocket)(unsafe.Pointer(socket.Native()))
	_arg2 = (*C.GObject)(unsafe.Pointer(sourceObject.Native()))

	var _cerr *C.GError // in

	C.g_socket_listener_add_socket(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// Close closes all the sockets in the listener.
func (l socketListener) Close() {
	var _arg0 *C.GSocketListener // out

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(l.Native()))

	C.g_socket_listener_close(_arg0)
}

// SetBacklog sets the listen backlog on the sockets in the listener. This
// must be called before adding any sockets, addresses or ports to the
// Listener (for example, by calling g_socket_listener_add_inet_port()) to
// be effective.
//
// See g_socket_set_listen_backlog() for details
func (l socketListener) SetBacklog(listenBacklog int) {
	var _arg0 *C.GSocketListener // out
	var _arg1 C.int              // out

	_arg0 = (*C.GSocketListener)(unsafe.Pointer(l.Native()))
	_arg1 = C.int(listenBacklog)

	C.g_socket_listener_set_backlog(_arg0, _arg1)
}
