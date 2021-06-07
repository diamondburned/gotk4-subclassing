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
		{T: externglib.Type(C.g_proxy_get_type()), F: marshalProxy},
	})
}

// ProxyGetDefaultForProtocol: find the `gio-proxy` extension point for a proxy
// implementation that supports the specified protocol.
func ProxyGetDefaultForProtocol(protocol string) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(protocol))
	defer C.free(unsafe.Pointer(arg1))

	C.g_proxy_get_default_for_protocol(arg1)
}

// ProxyOverrider contains methods that are overridable. This
// interface is a subset of the interface Proxy.
type ProxyOverrider interface {
	// Connect: given @connection to communicate with a proxy (eg, a Connection
	// that is connected to the proxy server), this does the necessary handshake
	// to connect to @proxy_address, and if required, wraps the OStream to
	// handle proxy payload.
	Connect(p Proxy, connection IOStream, proxyAddress ProxyAddress, cancellable Cancellable) error
	// ConnectAsync asynchronous version of g_proxy_connect().
	ConnectAsync(p Proxy)
	// ConnectFinish: see g_proxy_connect().
	ConnectFinish(p Proxy, result AsyncResult) error
	// SupportsHostname: some proxy protocols expect to be passed a hostname,
	// which they will resolve to an IP address themselves. Others, like SOCKS4,
	// do not allow this. This function will return false if @proxy is
	// implementing such a protocol. When false is returned, the caller should
	// resolve the destination hostname first, and then pass a Address
	// containing the stringified IP address to g_proxy_connect() or
	// g_proxy_connect_async().
	SupportsHostname(p Proxy) bool
}

// Proxy: a #GProxy handles connecting to a remote host via a given type of
// proxy server. It is implemented by the 'gio-proxy' extension point. The
// extensions are named after their proxy protocol name. As an example, a SOCKS5
// proxy implementation can be retrieved with the name 'socks5' using the
// function g_io_extension_point_get_extension_by_name().
type Proxy interface {
	gextras.Objector
	ProxyOverrider
}

// proxy implements the Proxy interface.
type proxy struct {
	gextras.Objector
}

var _ Proxy = (*proxy)(nil)

// WrapProxy wraps a GObject to a type that implements interface
// Proxy. It is primarily used internally.
func WrapProxy(obj *externglib.Object) Proxy {
	return Proxy{
		Objector: obj,
	}
}

func marshalProxy(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapProxy(obj), nil
}

// Connect: given @connection to communicate with a proxy (eg, a Connection
// that is connected to the proxy server), this does the necessary handshake
// to connect to @proxy_address, and if required, wraps the OStream to
// handle proxy payload.
func (p proxy) Connect(p Proxy, connection IOStream, proxyAddress ProxyAddress, cancellable Cancellable) error {
	var arg0 *C.GProxy
	var arg1 *C.GIOStream
	var arg2 *C.GProxyAddress
	var arg3 *C.GCancellable

	arg0 = (*C.GProxy)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GIOStream)(unsafe.Pointer(connection.Native()))
	arg2 = (*C.GProxyAddress)(unsafe.Pointer(proxyAddress.Native()))
	arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var errout *C.GError
	var err error

	C.g_proxy_connect(arg0, arg1, arg2, arg3, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// ConnectAsync asynchronous version of g_proxy_connect().
func (p proxy) ConnectAsync(p Proxy) {
	var arg0 *C.GProxy

	arg0 = (*C.GProxy)(unsafe.Pointer(p.Native()))

	C.g_proxy_connect_async(arg0, arg1, arg2, arg3, arg4, arg5)
}

// ConnectFinish: see g_proxy_connect().
func (p proxy) ConnectFinish(p Proxy, result AsyncResult) error {
	var arg0 *C.GProxy
	var arg1 *C.GAsyncResult

	arg0 = (*C.GProxy)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var errout *C.GError
	var err error

	C.g_proxy_connect_finish(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// SupportsHostname: some proxy protocols expect to be passed a hostname,
// which they will resolve to an IP address themselves. Others, like SOCKS4,
// do not allow this. This function will return false if @proxy is
// implementing such a protocol. When false is returned, the caller should
// resolve the destination hostname first, and then pass a Address
// containing the stringified IP address to g_proxy_connect() or
// g_proxy_connect_async().
func (p proxy) SupportsHostname(p Proxy) bool {
	var arg0 *C.GProxy

	arg0 = (*C.GProxy)(unsafe.Pointer(p.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_proxy_supports_hostname(arg0)

	if cret {
		ok = true
	}

	return ok
}