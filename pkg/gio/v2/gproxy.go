// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_proxy_get_type()), F: marshalProxier},
	})
}

// PROXY_EXTENSION_POINT_NAME: extension point for proxy functionality. See
// [Extending GIO][extending-gio].
const PROXY_EXTENSION_POINT_NAME = "gio-proxy"

// ProxyOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ProxyOverrider interface {
	// ConnectProxy: given connection to communicate with a proxy (eg, a
	// Connection that is connected to the proxy server), this does the
	// necessary handshake to connect to proxy_address, and if required, wraps
	// the OStream to handle proxy payload.
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional): #GCancellable.
	//    - connection: OStream.
	//    - proxyAddress: Address.
	//
	// The function returns the following values:
	//
	//    - ioStream that will replace connection. This might be the same as
	//      connection, in which case a reference will be added.
	//
	ConnectProxy(ctx context.Context, connection IOStreamer, proxyAddress *ProxyAddress) (IOStreamer, error)
	// ConnectAsync asynchronous version of g_proxy_connect().
	//
	// The function takes the following parameters:
	//
	//    - ctx (optional): #GCancellable.
	//    - connection: OStream.
	//    - proxyAddress: Address.
	//    - callback (optional): ReadyCallback.
	//
	ConnectAsync(ctx context.Context, connection IOStreamer, proxyAddress *ProxyAddress, callback AsyncReadyCallback)
	// ConnectFinish: see g_proxy_connect().
	//
	// The function takes the following parameters:
	//
	//    - result: Result.
	//
	// The function returns the following values:
	//
	//    - ioStream: OStream.
	//
	ConnectFinish(result AsyncResulter) (IOStreamer, error)
	// SupportsHostname: some proxy protocols expect to be passed a hostname,
	// which they will resolve to an IP address themselves. Others, like SOCKS4,
	// do not allow this. This function will return FALSE if proxy is
	// implementing such a protocol. When FALSE is returned, the caller should
	// resolve the destination hostname first, and then pass a Address
	// containing the stringified IP address to g_proxy_connect() or
	// g_proxy_connect_async().
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if hostname resolution is supported.
	//
	SupportsHostname() bool
}

// Proxy handles connecting to a remote host via a given type of proxy server.
// It is implemented by the 'gio-proxy' extension point. The extensions are
// named after their proxy protocol name. As an example, a SOCKS5 proxy
// implementation can be retrieved with the name 'socks5' using the function
// g_io_extension_point_get_extension_by_name().
type Proxy struct {
	_ [0]func() // equal guard
	*externglib.Object
}

var (
	_ externglib.Objector = (*Proxy)(nil)
)

// Proxier describes Proxy's interface methods.
type Proxier interface {
	externglib.Objector

	// ConnectProxy: given connection to communicate with a proxy (eg, a
	// Connection that is connected to the proxy server), this does the
	// necessary handshake to connect to proxy_address, and if required, wraps
	// the OStream to handle proxy payload.
	ConnectProxy(ctx context.Context, connection IOStreamer, proxyAddress *ProxyAddress) (IOStreamer, error)
	// ConnectAsync asynchronous version of g_proxy_connect().
	ConnectAsync(ctx context.Context, connection IOStreamer, proxyAddress *ProxyAddress, callback AsyncReadyCallback)
	// ConnectFinish: see g_proxy_connect().
	ConnectFinish(result AsyncResulter) (IOStreamer, error)
	// SupportsHostname: some proxy protocols expect to be passed a hostname,
	// which they will resolve to an IP address themselves.
	SupportsHostname() bool
}

var _ Proxier = (*Proxy)(nil)

func wrapProxy(obj *externglib.Object) *Proxy {
	return &Proxy{
		Object: obj,
	}
}

func marshalProxier(p uintptr) (interface{}, error) {
	return wrapProxy(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectProxy: given connection to communicate with a proxy (eg, a Connection
// that is connected to the proxy server), this does the necessary handshake to
// connect to proxy_address, and if required, wraps the OStream to handle proxy
// payload.
//
// The function takes the following parameters:
//
//    - ctx (optional): #GCancellable.
//    - connection: OStream.
//    - proxyAddress: Address.
//
// The function returns the following values:
//
//    - ioStream that will replace connection. This might be the same as
//      connection, in which case a reference will be added.
//
func (proxy *Proxy) ConnectProxy(ctx context.Context, connection IOStreamer, proxyAddress *ProxyAddress) (IOStreamer, error) {
	var _arg0 *C.GProxy        // out
	var _arg3 *C.GCancellable  // out
	var _arg1 *C.GIOStream     // out
	var _arg2 *C.GProxyAddress // out
	var _cret *C.GIOStream     // in
	var _cerr *C.GError        // in

	_arg0 = (*C.GProxy)(unsafe.Pointer(proxy.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GIOStream)(unsafe.Pointer(connection.Native()))
	_arg2 = (*C.GProxyAddress)(unsafe.Pointer(proxyAddress.Native()))

	_cret = C.g_proxy_connect(_arg0, _arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(proxy)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(proxyAddress)

	var _ioStream IOStreamer // out
	var _goerr error         // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.IOStreamer is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(IOStreamer)
			return ok
		})
		rv, ok := casted.(IOStreamer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.IOStreamer")
		}
		_ioStream = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _ioStream, _goerr
}

// ConnectAsync asynchronous version of g_proxy_connect().
//
// The function takes the following parameters:
//
//    - ctx (optional): #GCancellable.
//    - connection: OStream.
//    - proxyAddress: Address.
//    - callback (optional): ReadyCallback.
//
func (proxy *Proxy) ConnectAsync(ctx context.Context, connection IOStreamer, proxyAddress *ProxyAddress, callback AsyncReadyCallback) {
	var _arg0 *C.GProxy             // out
	var _arg3 *C.GCancellable       // out
	var _arg1 *C.GIOStream          // out
	var _arg2 *C.GProxyAddress      // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GProxy)(unsafe.Pointer(proxy.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GIOStream)(unsafe.Pointer(connection.Native()))
	_arg2 = (*C.GProxyAddress)(unsafe.Pointer(proxyAddress.Native()))
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_proxy_connect_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(proxy)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(connection)
	runtime.KeepAlive(proxyAddress)
	runtime.KeepAlive(callback)
}

// ConnectFinish: see g_proxy_connect().
//
// The function takes the following parameters:
//
//    - result: Result.
//
// The function returns the following values:
//
//    - ioStream: OStream.
//
func (proxy *Proxy) ConnectFinish(result AsyncResulter) (IOStreamer, error) {
	var _arg0 *C.GProxy       // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GIOStream    // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GProxy)(unsafe.Pointer(proxy.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_proxy_connect_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(proxy)
	runtime.KeepAlive(result)

	var _ioStream IOStreamer // out
	var _goerr error         // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.IOStreamer is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(IOStreamer)
			return ok
		})
		rv, ok := casted.(IOStreamer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.IOStreamer")
		}
		_ioStream = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _ioStream, _goerr
}

// SupportsHostname: some proxy protocols expect to be passed a hostname, which
// they will resolve to an IP address themselves. Others, like SOCKS4, do not
// allow this. This function will return FALSE if proxy is implementing such a
// protocol. When FALSE is returned, the caller should resolve the destination
// hostname first, and then pass a Address containing the stringified IP address
// to g_proxy_connect() or g_proxy_connect_async().
//
// The function returns the following values:
//
//    - ok: TRUE if hostname resolution is supported.
//
func (proxy *Proxy) SupportsHostname() bool {
	var _arg0 *C.GProxy  // out
	var _cret C.gboolean // in

	_arg0 = (*C.GProxy)(unsafe.Pointer(proxy.Native()))

	_cret = C.g_proxy_supports_hostname(_arg0)
	runtime.KeepAlive(proxy)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ProxyGetDefaultForProtocol: find the gio-proxy extension point for a proxy
// implementation that supports the specified protocol.
//
// The function takes the following parameters:
//
//    - protocol: proxy protocol name (e.g. http, socks, etc).
//
// The function returns the following values:
//
//    - proxy (optional): return a #GProxy or NULL if protocol is not supported.
//
func ProxyGetDefaultForProtocol(protocol string) Proxier {
	var _arg1 *C.gchar  // out
	var _cret *C.GProxy // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(protocol)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_proxy_get_default_for_protocol(_arg1)
	runtime.KeepAlive(protocol)

	var _proxy Proxier // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(Proxier)
				return ok
			})
			rv, ok := casted.(Proxier)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.Proxier")
			}
			_proxy = rv
		}
	}

	return _proxy
}
