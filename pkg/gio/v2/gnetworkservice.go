// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_network_service_get_type()), F: marshalNetworkService},
	})
}

// NetworkService: like Address does with hostnames, Service provides an easy
// way to resolve a SRV record, and then attempt to connect to one of the hosts
// that implements that service, handling service priority/weighting, multiple
// IP addresses, and multiple address families.
//
// See Target for more information about SRV records, and see Connectable for an
// example of using the connectable interface.
type NetworkService interface {
	SocketConnectable

	// Domain:
	Domain() string
	// Protocol:
	Protocol() string
	// Scheme:
	Scheme() string
	// Service:
	Service() string
	// SetSchemeNetworkService:
	SetSchemeNetworkService(scheme string)
}

// networkService implements the NetworkService class.
type networkService struct {
	gextras.Objector
}

// WrapNetworkService wraps a GObject to the right type. It is
// primarily used internally.
func WrapNetworkService(obj *externglib.Object) NetworkService {
	return networkService{
		Objector: obj,
	}
}

func marshalNetworkService(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNetworkService(obj), nil
}

// NewNetworkService:
func NewNetworkService(service string, protocol string, domain string) NetworkService {
	var _arg1 *C.gchar              // out
	var _arg2 *C.gchar              // out
	var _arg3 *C.gchar              // out
	var _cret *C.GSocketConnectable // in

	_arg1 = (*C.gchar)(C.CString(service))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(protocol))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(domain))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.g_network_service_new(_arg1, _arg2, _arg3)

	var _networkService NetworkService // out

	_networkService = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(NetworkService)

	return _networkService
}

func (s networkService) Domain() string {
	var _arg0 *C.GNetworkService // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GNetworkService)(unsafe.Pointer(s.Native()))

	_cret = C.g_network_service_get_domain(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s networkService) Protocol() string {
	var _arg0 *C.GNetworkService // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GNetworkService)(unsafe.Pointer(s.Native()))

	_cret = C.g_network_service_get_protocol(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s networkService) Scheme() string {
	var _arg0 *C.GNetworkService // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GNetworkService)(unsafe.Pointer(s.Native()))

	_cret = C.g_network_service_get_scheme(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s networkService) Service() string {
	var _arg0 *C.GNetworkService // out
	var _cret *C.gchar           // in

	_arg0 = (*C.GNetworkService)(unsafe.Pointer(s.Native()))

	_cret = C.g_network_service_get_service(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s networkService) SetSchemeNetworkService(scheme string) {
	var _arg0 *C.GNetworkService // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GNetworkService)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(scheme))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_network_service_set_scheme(_arg0, _arg1)
}

func (c networkService) Enumerate() SocketAddressEnumerator {
	return WrapSocketConnectable(gextras.InternObject(c)).Enumerate()
}

func (c networkService) ProxyEnumerate() SocketAddressEnumerator {
	return WrapSocketConnectable(gextras.InternObject(c)).ProxyEnumerate()
}

func (c networkService) String() string {
	return WrapSocketConnectable(gextras.InternObject(c)).String()
}