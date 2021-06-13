// Code generated by girgen. DO NOT EDIT.

package gio

import (
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
		{T: externglib.Type(C.g_inet_socket_address_get_type()), F: marshalInetSocketAddress},
	})
}

// InetSocketAddress: an IPv4 or IPv6 socket address; that is, the combination
// of a Address and a port number.
type InetSocketAddress interface {
	SocketAddress
	SocketConnectable

	// Flowinfo gets the `sin6_flowinfo` field from @address, which must be an
	// IPv6 address.
	Flowinfo() uint32
	// Port gets @address's port.
	Port() uint16
	// ScopeID gets the `sin6_scope_id` field from @address, which must be an
	// IPv6 address.
	ScopeID() uint32
}

// inetSocketAddress implements the InetSocketAddress interface.
type inetSocketAddress struct {
	SocketAddress
	SocketConnectable
}

var _ InetSocketAddress = (*inetSocketAddress)(nil)

// WrapInetSocketAddress wraps a GObject to the right type. It is
// primarily used internally.
func WrapInetSocketAddress(obj *externglib.Object) InetSocketAddress {
	return InetSocketAddress{
		SocketAddress:     WrapSocketAddress(obj),
		SocketConnectable: WrapSocketConnectable(obj),
	}
}

func marshalInetSocketAddress(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapInetSocketAddress(obj), nil
}

// Flowinfo gets the `sin6_flowinfo` field from @address, which must be an
// IPv6 address.
func (a inetSocketAddress) Flowinfo() uint32 {
	var _arg0 *C.GInetSocketAddress // out

	_arg0 = (*C.GInetSocketAddress)(unsafe.Pointer(a.Native()))

	var _cret C.guint32 // in

	_cret = C.g_inet_socket_address_get_flowinfo(_arg0)

	var _guint32 uint32 // out

	_guint32 = (uint32)(_cret)

	return _guint32
}

// Port gets @address's port.
func (a inetSocketAddress) Port() uint16 {
	var _arg0 *C.GInetSocketAddress // out

	_arg0 = (*C.GInetSocketAddress)(unsafe.Pointer(a.Native()))

	var _cret C.guint16 // in

	_cret = C.g_inet_socket_address_get_port(_arg0)

	var _guint16 uint16 // out

	_guint16 = (uint16)(_cret)

	return _guint16
}

// ScopeID gets the `sin6_scope_id` field from @address, which must be an
// IPv6 address.
func (a inetSocketAddress) ScopeID() uint32 {
	var _arg0 *C.GInetSocketAddress // out

	_arg0 = (*C.GInetSocketAddress)(unsafe.Pointer(a.Native()))

	var _cret C.guint32 // in

	_cret = C.g_inet_socket_address_get_scope_id(_arg0)

	var _guint32 uint32 // out

	_guint32 = (uint32)(_cret)

	return _guint32
}
