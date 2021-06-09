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
		{T: externglib.Type(C.g_tcp_wrapper_connection_get_type()), F: marshalTcpWrapperConnection},
	})
}

// TcpWrapperConnection: a WrapperConnection can be used to wrap a OStream that
// is based on a #GSocket, but which is not actually a Connection. This is used
// by Client so that it can always return a Connection, even when the connection
// it has actually created is not directly a Connection.
type TcpWrapperConnection interface {
	TcpConnection

	// BaseIOStream gets @conn's base OStream
	BaseIOStream() IOStream
}

// tcpWrapperConnection implements the TcpWrapperConnection interface.
type tcpWrapperConnection struct {
	TcpConnection
}

var _ TcpWrapperConnection = (*tcpWrapperConnection)(nil)

// WrapTcpWrapperConnection wraps a GObject to the right type. It is
// primarily used internally.
func WrapTcpWrapperConnection(obj *externglib.Object) TcpWrapperConnection {
	return TcpWrapperConnection{
		TcpConnection: WrapTcpConnection(obj),
	}
}

func marshalTcpWrapperConnection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTcpWrapperConnection(obj), nil
}

// NewTcpWrapperConnection constructs a class TcpWrapperConnection.
func NewTcpWrapperConnection(baseIoStream IOStream, socket Socket) TcpWrapperConnection {
	var arg1 *C.GIOStream
	var arg2 *C.GSocket

	arg1 = (*C.GIOStream)(unsafe.Pointer(baseIoStream.Native()))
	arg2 = (*C.GSocket)(unsafe.Pointer(socket.Native()))

	var cret C.GTcpWrapperConnection

	cret = C.g_tcp_wrapper_connection_new(arg1, arg2)

	var tcpWrapperConnection TcpWrapperConnection

	tcpWrapperConnection = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(TcpWrapperConnection)

	return tcpWrapperConnection
}

// BaseIOStream gets @conn's base OStream
func (c tcpWrapperConnection) BaseIOStream() IOStream {
	var arg0 *C.GTcpWrapperConnection

	arg0 = (*C.GTcpWrapperConnection)(unsafe.Pointer(c.Native()))

	var cret *C.GIOStream

	cret = C.g_tcp_wrapper_connection_get_base_io_stream(arg0)

	var ioStream IOStream

	ioStream = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(IOStream)

	return ioStream
}
