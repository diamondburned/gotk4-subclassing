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
		{T: externglib.Type(C.g_inet_socket_address_get_type()), F: marshalInetSocketAddress},
	})
}

// InetSocketAddress: an IPv4 or IPv6 socket address; that is, the combination
// of a Address and a port number.
type InetSocketAddress interface {
	SocketAddress
	SocketConnectable

	// Address gets @address's Address.
	Address(a InetSocketAddress)
	// Flowinfo gets the `sin6_flowinfo` field from @address, which must be an
	// IPv6 address.
	Flowinfo(a InetSocketAddress)
	// Port gets @address's port.
	Port(a InetSocketAddress)
	// ScopeID gets the `sin6_scope_id` field from @address, which must be an
	// IPv6 address.
	ScopeID(a InetSocketAddress)
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

// NewInetSocketAddress constructs a class InetSocketAddress.
func NewInetSocketAddress(address InetAddress, port uint16) {
	var arg1 *C.GInetAddress
	var arg2 C.guint16

	arg1 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))
	arg2 = C.guint16(port)

	C.g_inet_socket_address_new(arg1, arg2)
}

// NewInetSocketAddressFromString constructs a class InetSocketAddress.
func NewInetSocketAddressFromString(address string, port uint) {
	var arg1 *C.char
	var arg2 C.guint

	arg1 = (*C.char)(C.CString(address))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.guint(port)

	C.g_inet_socket_address_new_from_string(arg1, arg2)
}

// Address gets @address's Address.
func (a inetSocketAddress) Address(a InetSocketAddress) {
	var arg0 *C.GInetSocketAddress

	arg0 = (*C.GInetSocketAddress)(unsafe.Pointer(a.Native()))

	C.g_inet_socket_address_get_address(arg0)
}

// Flowinfo gets the `sin6_flowinfo` field from @address, which must be an
// IPv6 address.
func (a inetSocketAddress) Flowinfo(a InetSocketAddress) {
	var arg0 *C.GInetSocketAddress

	arg0 = (*C.GInetSocketAddress)(unsafe.Pointer(a.Native()))

	C.g_inet_socket_address_get_flowinfo(arg0)
}

// Port gets @address's port.
func (a inetSocketAddress) Port(a InetSocketAddress) {
	var arg0 *C.GInetSocketAddress

	arg0 = (*C.GInetSocketAddress)(unsafe.Pointer(a.Native()))

	C.g_inet_socket_address_get_port(arg0)
}

// ScopeID gets the `sin6_scope_id` field from @address, which must be an
// IPv6 address.
func (a inetSocketAddress) ScopeID(a InetSocketAddress) {
	var arg0 *C.GInetSocketAddress

	arg0 = (*C.GInetSocketAddress)(unsafe.Pointer(a.Native()))

	C.g_inet_socket_address_get_scope_id(arg0)
}