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
		{T: externglib.Type(C.g_proxy_address_get_type()), F: marshalProxyAddress},
	})
}

// ProxyAddress: support for proxied SocketAddress.
type ProxyAddress interface {
	InetSocketAddress
	SocketConnectable

	// DestinationHostname gets @proxy's destination hostname; that is, the name
	// of the host that will be connected to via the proxy, not the name of the
	// proxy itself.
	DestinationHostname(p ProxyAddress)
	// DestinationPort gets @proxy's destination port; that is, the port on the
	// destination host that will be connected to via the proxy, not the port
	// number of the proxy itself.
	DestinationPort(p ProxyAddress)
	// DestinationProtocol gets the protocol that is being spoken to the
	// destination server; eg, "http" or "ftp".
	DestinationProtocol(p ProxyAddress)
	// Password gets @proxy's password.
	Password(p ProxyAddress)
	// Protocol gets @proxy's protocol. eg, "socks" or "http"
	Protocol(p ProxyAddress)
	// URI gets the proxy URI that @proxy was constructed from.
	URI(p ProxyAddress)
	// Username gets @proxy's username.
	Username(p ProxyAddress)
}

// proxyAddress implements the ProxyAddress interface.
type proxyAddress struct {
	InetSocketAddress
	SocketConnectable
}

var _ ProxyAddress = (*proxyAddress)(nil)

// WrapProxyAddress wraps a GObject to the right type. It is
// primarily used internally.
func WrapProxyAddress(obj *externglib.Object) ProxyAddress {
	return ProxyAddress{
		InetSocketAddress: WrapInetSocketAddress(obj),
		SocketConnectable: WrapSocketConnectable(obj),
	}
}

func marshalProxyAddress(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapProxyAddress(obj), nil
}

// NewProxyAddress constructs a class ProxyAddress.
func NewProxyAddress(inetaddr InetAddress, port uint16, protocol string, destHostname string, destPort uint16, username string, password string) {
	var arg1 *C.GInetAddress
	var arg2 C.guint16
	var arg3 *C.gchar
	var arg4 *C.gchar
	var arg5 C.guint16
	var arg6 *C.gchar
	var arg7 *C.gchar

	arg1 = (*C.GInetAddress)(unsafe.Pointer(inetaddr.Native()))
	arg2 = C.guint16(port)
	arg3 = (*C.gchar)(C.CString(protocol))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = (*C.gchar)(C.CString(destHostname))
	defer C.free(unsafe.Pointer(arg4))
	arg5 = C.guint16(destPort)
	arg6 = (*C.gchar)(C.CString(username))
	defer C.free(unsafe.Pointer(arg6))
	arg7 = (*C.gchar)(C.CString(password))
	defer C.free(unsafe.Pointer(arg7))

	C.g_proxy_address_new(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// DestinationHostname gets @proxy's destination hostname; that is, the name
// of the host that will be connected to via the proxy, not the name of the
// proxy itself.
func (p proxyAddress) DestinationHostname(p ProxyAddress) {
	var arg0 *C.GProxyAddress

	arg0 = (*C.GProxyAddress)(unsafe.Pointer(p.Native()))

	C.g_proxy_address_get_destination_hostname(arg0)
}

// DestinationPort gets @proxy's destination port; that is, the port on the
// destination host that will be connected to via the proxy, not the port
// number of the proxy itself.
func (p proxyAddress) DestinationPort(p ProxyAddress) {
	var arg0 *C.GProxyAddress

	arg0 = (*C.GProxyAddress)(unsafe.Pointer(p.Native()))

	C.g_proxy_address_get_destination_port(arg0)
}

// DestinationProtocol gets the protocol that is being spoken to the
// destination server; eg, "http" or "ftp".
func (p proxyAddress) DestinationProtocol(p ProxyAddress) {
	var arg0 *C.GProxyAddress

	arg0 = (*C.GProxyAddress)(unsafe.Pointer(p.Native()))

	C.g_proxy_address_get_destination_protocol(arg0)
}

// Password gets @proxy's password.
func (p proxyAddress) Password(p ProxyAddress) {
	var arg0 *C.GProxyAddress

	arg0 = (*C.GProxyAddress)(unsafe.Pointer(p.Native()))

	C.g_proxy_address_get_password(arg0)
}

// Protocol gets @proxy's protocol. eg, "socks" or "http"
func (p proxyAddress) Protocol(p ProxyAddress) {
	var arg0 *C.GProxyAddress

	arg0 = (*C.GProxyAddress)(unsafe.Pointer(p.Native()))

	C.g_proxy_address_get_protocol(arg0)
}

// URI gets the proxy URI that @proxy was constructed from.
func (p proxyAddress) URI(p ProxyAddress) {
	var arg0 *C.GProxyAddress

	arg0 = (*C.GProxyAddress)(unsafe.Pointer(p.Native()))

	C.g_proxy_address_get_uri(arg0)
}

// Username gets @proxy's username.
func (p proxyAddress) Username(p ProxyAddress) {
	var arg0 *C.GProxyAddress

	arg0 = (*C.GProxyAddress)(unsafe.Pointer(p.Native()))

	C.g_proxy_address_get_username(arg0)
}