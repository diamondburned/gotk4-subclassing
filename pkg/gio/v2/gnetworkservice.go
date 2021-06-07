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
	gextras.Objector
	SocketConnectable

	// Domain gets the domain that @srv serves. This might be either UTF-8 or
	// ASCII-encoded, depending on what @srv was created with.
	Domain(s NetworkService)
	// Protocol gets @srv's protocol name (eg, "tcp").
	Protocol(s NetworkService)
	// Scheme gets the URI scheme used to resolve proxies. By default, the
	// service name is used as scheme.
	Scheme(s NetworkService)
	// Service gets @srv's service name (eg, "ldap").
	Service(s NetworkService)
	// SetScheme set's the URI scheme used to resolve proxies. By default, the
	// service name is used as scheme.
	SetScheme(s NetworkService, scheme string)
}

// networkService implements the NetworkService interface.
type networkService struct {
	gextras.Objector
	SocketConnectable
}

var _ NetworkService = (*networkService)(nil)

// WrapNetworkService wraps a GObject to the right type. It is
// primarily used internally.
func WrapNetworkService(obj *externglib.Object) NetworkService {
	return NetworkService{
		Objector:          obj,
		SocketConnectable: WrapSocketConnectable(obj),
	}
}

func marshalNetworkService(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNetworkService(obj), nil
}

// NewNetworkService constructs a class NetworkService.
func NewNetworkService(service string, protocol string, domain string) {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar

	arg1 = (*C.gchar)(C.CString(service))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(protocol))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(domain))
	defer C.free(unsafe.Pointer(arg3))

	C.g_network_service_new(arg1, arg2, arg3)
}

// Domain gets the domain that @srv serves. This might be either UTF-8 or
// ASCII-encoded, depending on what @srv was created with.
func (s networkService) Domain(s NetworkService) {
	var arg0 *C.GNetworkService

	arg0 = (*C.GNetworkService)(unsafe.Pointer(s.Native()))

	C.g_network_service_get_domain(arg0)
}

// Protocol gets @srv's protocol name (eg, "tcp").
func (s networkService) Protocol(s NetworkService) {
	var arg0 *C.GNetworkService

	arg0 = (*C.GNetworkService)(unsafe.Pointer(s.Native()))

	C.g_network_service_get_protocol(arg0)
}

// Scheme gets the URI scheme used to resolve proxies. By default, the
// service name is used as scheme.
func (s networkService) Scheme(s NetworkService) {
	var arg0 *C.GNetworkService

	arg0 = (*C.GNetworkService)(unsafe.Pointer(s.Native()))

	C.g_network_service_get_scheme(arg0)
}

// Service gets @srv's service name (eg, "ldap").
func (s networkService) Service(s NetworkService) {
	var arg0 *C.GNetworkService

	arg0 = (*C.GNetworkService)(unsafe.Pointer(s.Native()))

	C.g_network_service_get_service(arg0)
}

// SetScheme set's the URI scheme used to resolve proxies. By default, the
// service name is used as scheme.
func (s networkService) SetScheme(s NetworkService, scheme string) {
	var arg0 *C.GNetworkService
	var arg1 *C.gchar

	arg0 = (*C.GNetworkService)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(scheme))
	defer C.free(unsafe.Pointer(arg1))

	C.g_network_service_set_scheme(arg0, arg1)
}