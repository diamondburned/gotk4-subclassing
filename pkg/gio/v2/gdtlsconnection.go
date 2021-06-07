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
		{T: externglib.Type(C.g_dtls_connection_get_type()), F: marshalDTLSConnection},
	})
}

// DTLSConnectionOverrider contains methods that are overridable. This
// interface is a subset of the interface DTLSConnection.
type DTLSConnectionOverrider interface {
	AcceptCertificate(c DTLSConnection, peerCert TLSCertificate, errors TLSCertificateFlags) bool

	BindingData(c DTLSConnection, typ TLSChannelBindingType, data []byte) error
	// NegotiatedProtocol gets the name of the application-layer protocol
	// negotiated during the handshake.
	//
	// If the peer did not use the ALPN extension, or did not advertise a
	// protocol that matched one of @conn's protocols, or the TLS backend does
	// not support ALPN, then this will be nil. See
	// g_dtls_connection_set_advertised_protocols().
	NegotiatedProtocol(c DTLSConnection)
	// Handshake attempts a TLS handshake on @conn.
	//
	// On the client side, it is never necessary to call this method; although
	// the connection needs to perform a handshake after connecting, Connection
	// will handle this for you automatically when you try to send or receive
	// data on the connection. You can call g_dtls_connection_handshake()
	// manually if you want to know whether the initial handshake succeeded or
	// failed (as opposed to just immediately trying to use @conn to read or
	// write, in which case, if it fails, it may not be possible to tell if it
	// failed before or after completing the handshake), but beware that servers
	// may reject client authentication after the handshake has completed, so a
	// successful handshake does not indicate the connection will be usable.
	//
	// Likewise, on the server side, although a handshake is necessary at the
	// beginning of the communication, you do not need to call this function
	// explicitly unless you want clearer error reporting.
	//
	// Previously, calling g_dtls_connection_handshake() after the initial
	// handshake would trigger a rehandshake; however, this usage was deprecated
	// in GLib 2.60 because rehandshaking was removed from the TLS protocol in
	// TLS 1.3. Since GLib 2.64, calling this function after the initial
	// handshake will no longer do anything.
	//
	// Connection::accept_certificate may be emitted during the handshake.
	Handshake(c DTLSConnection, cancellable Cancellable) error
	// HandshakeAsync: asynchronously performs a TLS handshake on @conn. See
	// g_dtls_connection_handshake() for more information.
	HandshakeAsync(c DTLSConnection)
	// HandshakeFinish: finish an asynchronous TLS handshake operation. See
	// g_dtls_connection_handshake() for more information.
	HandshakeFinish(c DTLSConnection, result AsyncResult) error
	// SetAdvertisedProtocols sets the list of application-layer protocols to
	// advertise that the caller is willing to speak on this connection. The
	// Application-Layer Protocol Negotiation (ALPN) extension will be used to
	// negotiate a compatible protocol with the peer; use
	// g_dtls_connection_get_negotiated_protocol() to find the negotiated
	// protocol after the handshake. Specifying nil for the the value of
	// @protocols will disable ALPN negotiation.
	//
	// See IANA TLS ALPN Protocol IDs
	// (https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml#alpn-protocol-ids)
	// for a list of registered protocol IDs.
	SetAdvertisedProtocols(c DTLSConnection, protocols []string)
	// Shutdown: shut down part or all of a DTLS connection.
	//
	// If @shutdown_read is true then the receiving side of the connection is
	// shut down, and further reading is disallowed. Subsequent calls to
	// g_datagram_based_receive_messages() will return G_IO_ERROR_CLOSED.
	//
	// If @shutdown_write is true then the sending side of the connection is
	// shut down, and further writing is disallowed. Subsequent calls to
	// g_datagram_based_send_messages() will return G_IO_ERROR_CLOSED.
	//
	// It is allowed for both @shutdown_read and @shutdown_write to be TRUE —
	// this is equivalent to calling g_dtls_connection_close().
	//
	// If @cancellable is cancelled, the Connection may be left partially-closed
	// and any pending untransmitted data may be lost. Call
	// g_dtls_connection_shutdown() again to complete closing the Connection.
	Shutdown(c DTLSConnection, shutdownRead bool, shutdownWrite bool, cancellable Cancellable) error
	// ShutdownAsync: asynchronously shut down part or all of the DTLS
	// connection. See g_dtls_connection_shutdown() for more information.
	ShutdownAsync(c DTLSConnection)
	// ShutdownFinish: finish an asynchronous TLS shutdown operation. See
	// g_dtls_connection_shutdown() for more information.
	ShutdownFinish(c DTLSConnection, result AsyncResult) error
}

// DTLSConnection is the base DTLS connection class type, which wraps a Based
// and provides DTLS encryption on top of it. Its subclasses, ClientConnection
// and ServerConnection, implement client-side and server-side DTLS,
// respectively.
//
// For TLS support, see Connection.
//
// As DTLS is datagram based, Connection implements Based, presenting a
// datagram-socket-like API for the encrypted connection. This operates over a
// base datagram connection, which is also a Based (Connection:base-socket).
//
// To close a DTLS connection, use g_dtls_connection_close().
//
// Neither ServerConnection or ClientConnection set the peer address on their
// base Based if it is a #GSocket — it is up to the caller to do that if they
// wish. If they do not, and g_socket_close() is called on the base socket, the
// Connection will not raise a G_IO_ERROR_NOT_CONNECTED error on further I/O.
type DTLSConnection interface {
	DatagramBased
	DTLSConnectionOverrider

	// Close: close the DTLS connection. This is equivalent to calling
	// g_dtls_connection_shutdown() to shut down both sides of the connection.
	//
	// Closing a Connection waits for all buffered but untransmitted data to be
	// sent before it completes. It then sends a `close_notify` DTLS alert to
	// the peer and may wait for a `close_notify` to be received from the peer.
	// It does not close the underlying Connection:base-socket; that must be
	// closed separately.
	//
	// Once @conn is closed, all other operations will return G_IO_ERROR_CLOSED.
	// Closing a Connection multiple times will not return an error.
	//
	// Connections will be automatically closed when the last reference is
	// dropped, but you might want to call this function to make sure resources
	// are released as early as possible.
	//
	// If @cancellable is cancelled, the Connection may be left partially-closed
	// and any pending untransmitted data may be lost. Call
	// g_dtls_connection_close() again to complete closing the Connection.
	Close(c DTLSConnection, cancellable Cancellable) error
	// CloseAsync: asynchronously close the DTLS connection. See
	// g_dtls_connection_close() for more information.
	CloseAsync(c DTLSConnection)
	// CloseFinish: finish an asynchronous TLS close operation. See
	// g_dtls_connection_close() for more information.
	CloseFinish(c DTLSConnection, result AsyncResult) error
	// EmitAcceptCertificate: used by Connection implementations to emit the
	// Connection::accept-certificate signal.
	EmitAcceptCertificate(c DTLSConnection, peerCert TLSCertificate, errors TLSCertificateFlags) bool
	// Certificate gets @conn's certificate, as set by
	// g_dtls_connection_set_certificate().
	Certificate(c DTLSConnection)
	// ChannelBindingData: query the TLS backend for TLS channel binding data of
	// @type for @conn.
	//
	// This call retrieves TLS channel binding data as specified in RFC 5056
	// (https://tools.ietf.org/html/rfc5056), RFC 5929
	// (https://tools.ietf.org/html/rfc5929), and related RFCs. The binding data
	// is returned in @data. The @data is resized by the callee using Array
	// buffer management and will be freed when the @data is destroyed by
	// g_byte_array_unref(). If @data is nil, it will only check whether TLS
	// backend is able to fetch the data (e.g. whether @type is supported by the
	// TLS backend). It does not guarantee that the data will be available
	// though. That could happen if TLS connection does not support @type or the
	// binding data is not available yet due to additional negotiation or input
	// required.
	ChannelBindingData(c DTLSConnection, typ TLSChannelBindingType) (data []byte, err error)
	// Database gets the certificate database that @conn uses to verify peer
	// certificates. See g_dtls_connection_set_database().
	Database(c DTLSConnection)
	// Interaction: get the object that will be used to interact with the user.
	// It will be used for things like prompting the user for passwords. If nil
	// is returned, then no user interaction will occur for this connection.
	Interaction(c DTLSConnection)
	// PeerCertificate gets @conn's peer's certificate after the handshake has
	// completed or failed. (It is not set during the emission of
	// Connection::accept-certificate.)
	PeerCertificate(c DTLSConnection)
	// PeerCertificateErrors gets the errors associated with validating @conn's
	// peer's certificate, after the handshake has completed or failed. (It is
	// not set during the emission of Connection::accept-certificate.)
	PeerCertificateErrors(c DTLSConnection)
	// RehandshakeMode gets @conn rehandshaking mode. See
	// g_dtls_connection_set_rehandshake_mode() for details.
	RehandshakeMode(c DTLSConnection)
	// RequireCloseNotify tests whether or not @conn expects a proper TLS close
	// notification when the connection is closed. See
	// g_dtls_connection_set_require_close_notify() for details.
	RequireCloseNotify(c DTLSConnection) bool
	// SetCertificate: this sets the certificate that @conn will present to its
	// peer during the TLS handshake. For a ServerConnection, it is mandatory to
	// set this, and that will normally be done at construct time.
	//
	// For a ClientConnection, this is optional. If a handshake fails with
	// G_TLS_ERROR_CERTIFICATE_REQUIRED, that means that the server requires a
	// certificate, and if you try connecting again, you should call this method
	// first. You can call g_dtls_client_connection_get_accepted_cas() on the
	// failed connection to get a list of Certificate Authorities that the
	// server will accept certificates from.
	//
	// (It is also possible that a server will allow the connection with or
	// without a certificate; in that case, if you don't provide a certificate,
	// you can tell that the server requested one by the fact that
	// g_dtls_client_connection_get_accepted_cas() will return non-nil.)
	SetCertificate(c DTLSConnection, certificate TLSCertificate)
	// SetDatabase sets the certificate database that is used to verify peer
	// certificates. This is set to the default database by default. See
	// g_tls_backend_get_default_database(). If set to nil, then peer
	// certificate validation will always set the G_TLS_CERTIFICATE_UNKNOWN_CA
	// error (meaning Connection::accept-certificate will always be emitted on
	// client-side connections, unless that bit is not set in
	// ClientConnection:validation-flags).
	SetDatabase(c DTLSConnection, database TLSDatabase)
	// SetInteraction: set the object that will be used to interact with the
	// user. It will be used for things like prompting the user for passwords.
	//
	// The @interaction argument will normally be a derived subclass of
	// Interaction. nil can also be provided if no user interaction should occur
	// for this connection.
	SetInteraction(c DTLSConnection, interaction TLSInteraction)
	// SetRehandshakeMode: since GLib 2.64, changing the rehandshake mode is no
	// longer supported and will have no effect. With TLS 1.3, rehandshaking has
	// been removed from the TLS protocol, replaced by separate post-handshake
	// authentication and rekey operations.
	SetRehandshakeMode(c DTLSConnection, mode TLSRehandshakeMode)
	// SetRequireCloseNotify sets whether or not @conn expects a proper TLS
	// close notification before the connection is closed. If this is true (the
	// default), then @conn will expect to receive a TLS close notification from
	// its peer before the connection is closed, and will return a
	// G_TLS_ERROR_EOF error if the connection is closed without proper
	// notification (since this may indicate a network error, or
	// man-in-the-middle attack).
	//
	// In some protocols, the application will know whether or not the
	// connection was closed cleanly based on application-level data (because
	// the application-level data includes a length field, or is somehow
	// self-delimiting); in this case, the close notify is redundant and may be
	// omitted. You can use g_dtls_connection_set_require_close_notify() to tell
	// @conn to allow an "unannounced" connection close, in which case the close
	// will show up as a 0-length read, as in a non-TLS Based, and it is up to
	// the application to check that the data has been fully received.
	//
	// Note that this only affects the behavior when the peer closes the
	// connection; when the application calls g_dtls_connection_close_async() on
	// @conn itself, this will send a close notification regardless of the
	// setting of this property. If you explicitly want to do an unclean close,
	// you can close @conn's Connection:base-socket rather than closing @conn
	// itself.
	SetRequireCloseNotify(c DTLSConnection, requireCloseNotify bool)
}

// dtlsConnection implements the DTLSConnection interface.
type dtlsConnection struct {
	DatagramBased
}

var _ DTLSConnection = (*dtlsConnection)(nil)

// WrapDTLSConnection wraps a GObject to a type that implements interface
// DTLSConnection. It is primarily used internally.
func WrapDTLSConnection(obj *externglib.Object) DTLSConnection {
	return DTLSConnection{
		DatagramBased: WrapDatagramBased(obj),
	}
}

func marshalDTLSConnection(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDTLSConnection(obj), nil
}

// Close: close the DTLS connection. This is equivalent to calling
// g_dtls_connection_shutdown() to shut down both sides of the connection.
//
// Closing a Connection waits for all buffered but untransmitted data to be
// sent before it completes. It then sends a `close_notify` DTLS alert to
// the peer and may wait for a `close_notify` to be received from the peer.
// It does not close the underlying Connection:base-socket; that must be
// closed separately.
//
// Once @conn is closed, all other operations will return G_IO_ERROR_CLOSED.
// Closing a Connection multiple times will not return an error.
//
// Connections will be automatically closed when the last reference is
// dropped, but you might want to call this function to make sure resources
// are released as early as possible.
//
// If @cancellable is cancelled, the Connection may be left partially-closed
// and any pending untransmitted data may be lost. Call
// g_dtls_connection_close() again to complete closing the Connection.
func (c dtlsConnection) Close(c DTLSConnection, cancellable Cancellable) error {
	var arg0 *C.GDtlsConnection
	var arg1 *C.GCancellable

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var errout *C.GError
	var err error

	C.g_dtls_connection_close(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// CloseAsync: asynchronously close the DTLS connection. See
// g_dtls_connection_close() for more information.
func (c dtlsConnection) CloseAsync(c DTLSConnection) {
	var arg0 *C.GDtlsConnection

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))

	C.g_dtls_connection_close_async(arg0, arg1, arg2, arg3, arg4)
}

// CloseFinish: finish an asynchronous TLS close operation. See
// g_dtls_connection_close() for more information.
func (c dtlsConnection) CloseFinish(c DTLSConnection, result AsyncResult) error {
	var arg0 *C.GDtlsConnection
	var arg1 *C.GAsyncResult

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var errout *C.GError
	var err error

	C.g_dtls_connection_close_finish(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// EmitAcceptCertificate: used by Connection implementations to emit the
// Connection::accept-certificate signal.
func (c dtlsConnection) EmitAcceptCertificate(c DTLSConnection, peerCert TLSCertificate, errors TLSCertificateFlags) bool {
	var arg0 *C.GDtlsConnection
	var arg1 *C.GTlsCertificate
	var arg2 C.GTlsCertificateFlags

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GTlsCertificate)(unsafe.Pointer(peerCert.Native()))
	arg2 = (C.GTlsCertificateFlags)(errors)

	var cret C.gboolean
	var ok bool

	cret = C.g_dtls_connection_emit_accept_certificate(arg0, arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// Certificate gets @conn's certificate, as set by
// g_dtls_connection_set_certificate().
func (c dtlsConnection) Certificate(c DTLSConnection) {
	var arg0 *C.GDtlsConnection

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))

	C.g_dtls_connection_get_certificate(arg0)
}

// ChannelBindingData: query the TLS backend for TLS channel binding data of
// @type for @conn.
//
// This call retrieves TLS channel binding data as specified in RFC 5056
// (https://tools.ietf.org/html/rfc5056), RFC 5929
// (https://tools.ietf.org/html/rfc5929), and related RFCs. The binding data
// is returned in @data. The @data is resized by the callee using Array
// buffer management and will be freed when the @data is destroyed by
// g_byte_array_unref(). If @data is nil, it will only check whether TLS
// backend is able to fetch the data (e.g. whether @type is supported by the
// TLS backend). It does not guarantee that the data will be available
// though. That could happen if TLS connection does not support @type or the
// binding data is not available yet due to additional negotiation or input
// required.
func (c dtlsConnection) ChannelBindingData(c DTLSConnection, typ TLSChannelBindingType) (data []byte, err error) {
	var arg0 *C.GDtlsConnection
	var arg1 C.GTlsChannelBindingType

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = (C.GTlsChannelBindingType)(typ)

	var arg2 *C.GByteArray
	var data []byte
	var errout *C.GError
	var err error

	C.g_dtls_connection_get_channel_binding_data(arg0, arg1, &arg2, &errout)

	{
		var length int
		for p := arg2; *p != 0; p = (*C.GByteArray)(ptr.Add(unsafe.Pointer(p), C.sizeof_guint8)) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		data = make([]byte, length)
		for i := uintptr(0); i < uintptr(length); i += C.sizeof_guint8 {
			src := (C.guint8)(ptr.Add(unsafe.Pointer(arg2), i))
			data[i] = byte(src)
		}
	}
	err = gerror.Take(unsafe.Pointer(errout))

	return data, err
}

// Database gets the certificate database that @conn uses to verify peer
// certificates. See g_dtls_connection_set_database().
func (c dtlsConnection) Database(c DTLSConnection) {
	var arg0 *C.GDtlsConnection

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))

	C.g_dtls_connection_get_database(arg0)
}

// Interaction: get the object that will be used to interact with the user.
// It will be used for things like prompting the user for passwords. If nil
// is returned, then no user interaction will occur for this connection.
func (c dtlsConnection) Interaction(c DTLSConnection) {
	var arg0 *C.GDtlsConnection

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))

	C.g_dtls_connection_get_interaction(arg0)
}

// NegotiatedProtocol gets the name of the application-layer protocol
// negotiated during the handshake.
//
// If the peer did not use the ALPN extension, or did not advertise a
// protocol that matched one of @conn's protocols, or the TLS backend does
// not support ALPN, then this will be nil. See
// g_dtls_connection_set_advertised_protocols().
func (c dtlsConnection) NegotiatedProtocol(c DTLSConnection) {
	var arg0 *C.GDtlsConnection

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))

	C.g_dtls_connection_get_negotiated_protocol(arg0)
}

// PeerCertificate gets @conn's peer's certificate after the handshake has
// completed or failed. (It is not set during the emission of
// Connection::accept-certificate.)
func (c dtlsConnection) PeerCertificate(c DTLSConnection) {
	var arg0 *C.GDtlsConnection

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))

	C.g_dtls_connection_get_peer_certificate(arg0)
}

// PeerCertificateErrors gets the errors associated with validating @conn's
// peer's certificate, after the handshake has completed or failed. (It is
// not set during the emission of Connection::accept-certificate.)
func (c dtlsConnection) PeerCertificateErrors(c DTLSConnection) {
	var arg0 *C.GDtlsConnection

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))

	C.g_dtls_connection_get_peer_certificate_errors(arg0)
}

// RehandshakeMode gets @conn rehandshaking mode. See
// g_dtls_connection_set_rehandshake_mode() for details.
func (c dtlsConnection) RehandshakeMode(c DTLSConnection) {
	var arg0 *C.GDtlsConnection

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))

	C.g_dtls_connection_get_rehandshake_mode(arg0)
}

// RequireCloseNotify tests whether or not @conn expects a proper TLS close
// notification when the connection is closed. See
// g_dtls_connection_set_require_close_notify() for details.
func (c dtlsConnection) RequireCloseNotify(c DTLSConnection) bool {
	var arg0 *C.GDtlsConnection

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_dtls_connection_get_require_close_notify(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Handshake attempts a TLS handshake on @conn.
//
// On the client side, it is never necessary to call this method; although
// the connection needs to perform a handshake after connecting, Connection
// will handle this for you automatically when you try to send or receive
// data on the connection. You can call g_dtls_connection_handshake()
// manually if you want to know whether the initial handshake succeeded or
// failed (as opposed to just immediately trying to use @conn to read or
// write, in which case, if it fails, it may not be possible to tell if it
// failed before or after completing the handshake), but beware that servers
// may reject client authentication after the handshake has completed, so a
// successful handshake does not indicate the connection will be usable.
//
// Likewise, on the server side, although a handshake is necessary at the
// beginning of the communication, you do not need to call this function
// explicitly unless you want clearer error reporting.
//
// Previously, calling g_dtls_connection_handshake() after the initial
// handshake would trigger a rehandshake; however, this usage was deprecated
// in GLib 2.60 because rehandshaking was removed from the TLS protocol in
// TLS 1.3. Since GLib 2.64, calling this function after the initial
// handshake will no longer do anything.
//
// Connection::accept_certificate may be emitted during the handshake.
func (c dtlsConnection) Handshake(c DTLSConnection, cancellable Cancellable) error {
	var arg0 *C.GDtlsConnection
	var arg1 *C.GCancellable

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var errout *C.GError
	var err error

	C.g_dtls_connection_handshake(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// HandshakeAsync: asynchronously performs a TLS handshake on @conn. See
// g_dtls_connection_handshake() for more information.
func (c dtlsConnection) HandshakeAsync(c DTLSConnection) {
	var arg0 *C.GDtlsConnection

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))

	C.g_dtls_connection_handshake_async(arg0, arg1, arg2, arg3, arg4)
}

// HandshakeFinish: finish an asynchronous TLS handshake operation. See
// g_dtls_connection_handshake() for more information.
func (c dtlsConnection) HandshakeFinish(c DTLSConnection, result AsyncResult) error {
	var arg0 *C.GDtlsConnection
	var arg1 *C.GAsyncResult

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var errout *C.GError
	var err error

	C.g_dtls_connection_handshake_finish(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// SetAdvertisedProtocols sets the list of application-layer protocols to
// advertise that the caller is willing to speak on this connection. The
// Application-Layer Protocol Negotiation (ALPN) extension will be used to
// negotiate a compatible protocol with the peer; use
// g_dtls_connection_get_negotiated_protocol() to find the negotiated
// protocol after the handshake. Specifying nil for the the value of
// @protocols will disable ALPN negotiation.
//
// See IANA TLS ALPN Protocol IDs
// (https://www.iana.org/assignments/tls-extensiontype-values/tls-extensiontype-values.xhtml#alpn-protocol-ids)
// for a list of registered protocol IDs.
func (c dtlsConnection) SetAdvertisedProtocols(c DTLSConnection, protocols []string) {
	var arg0 *C.GDtlsConnection
	var arg1 **C.gchar

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = C.malloc(len(protocols) * (unsafe.Sizeof(int(0)) + 1))
	defer C.free(unsafe.Pointer(arg1))

	{
		var out []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(arg1), int(len(protocols)))

		for i := range protocols {
			out[i] = (*C.gchar)(C.CString(protocols[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.g_dtls_connection_set_advertised_protocols(arg0, arg1)
}

// SetCertificate: this sets the certificate that @conn will present to its
// peer during the TLS handshake. For a ServerConnection, it is mandatory to
// set this, and that will normally be done at construct time.
//
// For a ClientConnection, this is optional. If a handshake fails with
// G_TLS_ERROR_CERTIFICATE_REQUIRED, that means that the server requires a
// certificate, and if you try connecting again, you should call this method
// first. You can call g_dtls_client_connection_get_accepted_cas() on the
// failed connection to get a list of Certificate Authorities that the
// server will accept certificates from.
//
// (It is also possible that a server will allow the connection with or
// without a certificate; in that case, if you don't provide a certificate,
// you can tell that the server requested one by the fact that
// g_dtls_client_connection_get_accepted_cas() will return non-nil.)
func (c dtlsConnection) SetCertificate(c DTLSConnection, certificate TLSCertificate) {
	var arg0 *C.GDtlsConnection
	var arg1 *C.GTlsCertificate

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GTlsCertificate)(unsafe.Pointer(certificate.Native()))

	C.g_dtls_connection_set_certificate(arg0, arg1)
}

// SetDatabase sets the certificate database that is used to verify peer
// certificates. This is set to the default database by default. See
// g_tls_backend_get_default_database(). If set to nil, then peer
// certificate validation will always set the G_TLS_CERTIFICATE_UNKNOWN_CA
// error (meaning Connection::accept-certificate will always be emitted on
// client-side connections, unless that bit is not set in
// ClientConnection:validation-flags).
func (c dtlsConnection) SetDatabase(c DTLSConnection, database TLSDatabase) {
	var arg0 *C.GDtlsConnection
	var arg1 *C.GTlsDatabase

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GTlsDatabase)(unsafe.Pointer(database.Native()))

	C.g_dtls_connection_set_database(arg0, arg1)
}

// SetInteraction: set the object that will be used to interact with the
// user. It will be used for things like prompting the user for passwords.
//
// The @interaction argument will normally be a derived subclass of
// Interaction. nil can also be provided if no user interaction should occur
// for this connection.
func (c dtlsConnection) SetInteraction(c DTLSConnection, interaction TLSInteraction) {
	var arg0 *C.GDtlsConnection
	var arg1 *C.GTlsInteraction

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GTlsInteraction)(unsafe.Pointer(interaction.Native()))

	C.g_dtls_connection_set_interaction(arg0, arg1)
}

// SetRehandshakeMode: since GLib 2.64, changing the rehandshake mode is no
// longer supported and will have no effect. With TLS 1.3, rehandshaking has
// been removed from the TLS protocol, replaced by separate post-handshake
// authentication and rekey operations.
func (c dtlsConnection) SetRehandshakeMode(c DTLSConnection, mode TLSRehandshakeMode) {
	var arg0 *C.GDtlsConnection
	var arg1 C.GTlsRehandshakeMode

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = (C.GTlsRehandshakeMode)(mode)

	C.g_dtls_connection_set_rehandshake_mode(arg0, arg1)
}

// SetRequireCloseNotify sets whether or not @conn expects a proper TLS
// close notification before the connection is closed. If this is true (the
// default), then @conn will expect to receive a TLS close notification from
// its peer before the connection is closed, and will return a
// G_TLS_ERROR_EOF error if the connection is closed without proper
// notification (since this may indicate a network error, or
// man-in-the-middle attack).
//
// In some protocols, the application will know whether or not the
// connection was closed cleanly based on application-level data (because
// the application-level data includes a length field, or is somehow
// self-delimiting); in this case, the close notify is redundant and may be
// omitted. You can use g_dtls_connection_set_require_close_notify() to tell
// @conn to allow an "unannounced" connection close, in which case the close
// will show up as a 0-length read, as in a non-TLS Based, and it is up to
// the application to check that the data has been fully received.
//
// Note that this only affects the behavior when the peer closes the
// connection; when the application calls g_dtls_connection_close_async() on
// @conn itself, this will send a close notification regardless of the
// setting of this property. If you explicitly want to do an unclean close,
// you can close @conn's Connection:base-socket rather than closing @conn
// itself.
func (c dtlsConnection) SetRequireCloseNotify(c DTLSConnection, requireCloseNotify bool) {
	var arg0 *C.GDtlsConnection
	var arg1 C.gboolean

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))
	if requireCloseNotify {
		arg1 = C.gboolean(1)
	}

	C.g_dtls_connection_set_require_close_notify(arg0, arg1)
}

// Shutdown: shut down part or all of a DTLS connection.
//
// If @shutdown_read is true then the receiving side of the connection is
// shut down, and further reading is disallowed. Subsequent calls to
// g_datagram_based_receive_messages() will return G_IO_ERROR_CLOSED.
//
// If @shutdown_write is true then the sending side of the connection is
// shut down, and further writing is disallowed. Subsequent calls to
// g_datagram_based_send_messages() will return G_IO_ERROR_CLOSED.
//
// It is allowed for both @shutdown_read and @shutdown_write to be TRUE —
// this is equivalent to calling g_dtls_connection_close().
//
// If @cancellable is cancelled, the Connection may be left partially-closed
// and any pending untransmitted data may be lost. Call
// g_dtls_connection_shutdown() again to complete closing the Connection.
func (c dtlsConnection) Shutdown(c DTLSConnection, shutdownRead bool, shutdownWrite bool, cancellable Cancellable) error {
	var arg0 *C.GDtlsConnection
	var arg1 C.gboolean
	var arg2 C.gboolean
	var arg3 *C.GCancellable

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))
	if shutdownRead {
		arg1 = C.gboolean(1)
	}
	if shutdownWrite {
		arg2 = C.gboolean(1)
	}
	arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var errout *C.GError
	var err error

	C.g_dtls_connection_shutdown(arg0, arg1, arg2, arg3, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// ShutdownAsync: asynchronously shut down part or all of the DTLS
// connection. See g_dtls_connection_shutdown() for more information.
func (c dtlsConnection) ShutdownAsync(c DTLSConnection) {
	var arg0 *C.GDtlsConnection

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))

	C.g_dtls_connection_shutdown_async(arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// ShutdownFinish: finish an asynchronous TLS shutdown operation. See
// g_dtls_connection_shutdown() for more information.
func (c dtlsConnection) ShutdownFinish(c DTLSConnection, result AsyncResult) error {
	var arg0 *C.GDtlsConnection
	var arg1 *C.GAsyncResult

	arg0 = (*C.GDtlsConnection)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var errout *C.GError
	var err error

	C.g_dtls_connection_shutdown_finish(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}
