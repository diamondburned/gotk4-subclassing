// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.g_dtls_client_connection_get_type()), F: marshalDTLSClientConnection},
	})
}

// DTLSClientConnection is the client-side subclass of Connection, representing
// a client-side DTLS connection.
type DTLSClientConnection interface {
	DatagramBasedDTLSConnection

	// AcceptedCAS gets the list of distinguished names of the Certificate
	// Authorities that the server will accept certificates from. This will be
	// set during the TLS handshake if the server requests a certificate.
	// Otherwise, it will be nil.
	//
	// Each item in the list is a Array which contains the complete subject DN
	// of the certificate authority.
	AcceptedCAS() *glib.List
	// ServerIdentity gets @conn's expected server identity
	ServerIdentity() SocketConnectable
	// ValidationFlags gets @conn's validation flags
	ValidationFlags() TLSCertificateFlags
	// SetServerIdentity sets @conn's expected server identity, which is used
	// both to tell servers on virtual hosts which certificate to present, and
	// also to let @conn know what name to look for in the certificate when
	// performing G_TLS_CERTIFICATE_BAD_IDENTITY validation, if enabled.
	SetServerIdentity(identity SocketConnectable)
	// SetValidationFlags sets @conn's validation flags, to override the default
	// set of checks performed when validating a server certificate. By default,
	// G_TLS_CERTIFICATE_VALIDATE_ALL is used.
	SetValidationFlags(flags TLSCertificateFlags)
}

// dtlsClientConnection implements the DTLSClientConnection interface.
type dtlsClientConnection struct {
	DatagramBased
	DTLSConnection
}

var _ DTLSClientConnection = (*dtlsClientConnection)(nil)

// WrapDTLSClientConnection wraps a GObject to a type that implements interface
// DTLSClientConnection. It is primarily used internally.
func WrapDTLSClientConnection(obj *externglib.Object) DTLSClientConnection {
	return DTLSClientConnection{
		DatagramBased:  WrapDatagramBased(obj),
		DTLSConnection: WrapDTLSConnection(obj),
	}
}

func marshalDTLSClientConnection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDTLSClientConnection(obj), nil
}

// AcceptedCAS gets the list of distinguished names of the Certificate
// Authorities that the server will accept certificates from. This will be
// set during the TLS handshake if the server requests a certificate.
// Otherwise, it will be nil.
//
// Each item in the list is a Array which contains the complete subject DN
// of the certificate authority.
func (c dtlsClientConnection) AcceptedCAS() *glib.List {
	var arg0 *C.GDtlsClientConnection

	arg0 = (*C.GDtlsClientConnection)(unsafe.Pointer(c.Native()))

	var cret *C.GList

	cret = C.g_dtls_client_connection_get_accepted_cas(arg0)

	var list *glib.List

	list = glib.WrapList(unsafe.Pointer(cret))
	runtime.SetFinalizer(list, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return list
}

// ServerIdentity gets @conn's expected server identity
func (c dtlsClientConnection) ServerIdentity() SocketConnectable {
	var arg0 *C.GDtlsClientConnection

	arg0 = (*C.GDtlsClientConnection)(unsafe.Pointer(c.Native()))

	var cret *C.GSocketConnectable

	cret = C.g_dtls_client_connection_get_server_identity(arg0)

	var socketConnectable SocketConnectable

	socketConnectable = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(SocketConnectable)

	return socketConnectable
}

// ValidationFlags gets @conn's validation flags
func (c dtlsClientConnection) ValidationFlags() TLSCertificateFlags {
	var arg0 *C.GDtlsClientConnection

	arg0 = (*C.GDtlsClientConnection)(unsafe.Pointer(c.Native()))

	var cret C.GTlsCertificateFlags

	cret = C.g_dtls_client_connection_get_validation_flags(arg0)

	var tlsCertificateFlags TLSCertificateFlags

	tlsCertificateFlags = TLSCertificateFlags(cret)

	return tlsCertificateFlags
}

// SetServerIdentity sets @conn's expected server identity, which is used
// both to tell servers on virtual hosts which certificate to present, and
// also to let @conn know what name to look for in the certificate when
// performing G_TLS_CERTIFICATE_BAD_IDENTITY validation, if enabled.
func (c dtlsClientConnection) SetServerIdentity(identity SocketConnectable) {
	var arg0 *C.GDtlsClientConnection
	var arg1 *C.GSocketConnectable

	arg0 = (*C.GDtlsClientConnection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GSocketConnectable)(unsafe.Pointer(identity.Native()))

	C.g_dtls_client_connection_set_server_identity(arg0, arg1)
}

// SetValidationFlags sets @conn's validation flags, to override the default
// set of checks performed when validating a server certificate. By default,
// G_TLS_CERTIFICATE_VALIDATE_ALL is used.
func (c dtlsClientConnection) SetValidationFlags(flags TLSCertificateFlags) {
	var arg0 *C.GDtlsClientConnection
	var arg1 C.GTlsCertificateFlags

	arg0 = (*C.GDtlsClientConnection)(unsafe.Pointer(c.Native()))
	arg1 = (C.GTlsCertificateFlags)(flags)

	C.g_dtls_client_connection_set_validation_flags(arg0, arg1)
}
