// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_tls_server_connection_get_type()), F: marshalTLSServerConnectioner},
	})
}

// TLSServerConnectionOverrider contains methods that are overridable.
type TLSServerConnectionOverrider interface {
}

// TLSServerConnection is the server-side subclass of Connection, representing a
// server-side TLS connection.
type TLSServerConnection struct {
	_ [0]func() // equal guard
	TLSConnection
}

var (
	_ TLSConnectioner = (*TLSServerConnection)(nil)
)

// TLSServerConnectioner describes TLSServerConnection's interface methods.
type TLSServerConnectioner interface {
	externglib.Objector

	baseTLSServerConnection() *TLSServerConnection
}

var _ TLSServerConnectioner = (*TLSServerConnection)(nil)

func ifaceInitTLSServerConnectioner(gifacePtr, data C.gpointer) {
}

func wrapTLSServerConnection(obj *externglib.Object) *TLSServerConnection {
	return &TLSServerConnection{
		TLSConnection: TLSConnection{
			IOStream: IOStream{
				Object: obj,
			},
		},
	}
}

func marshalTLSServerConnectioner(p uintptr) (interface{}, error) {
	return wrapTLSServerConnection(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *TLSServerConnection) baseTLSServerConnection() *TLSServerConnection {
	return v
}

// BaseTLSServerConnection returns the underlying base object.
func BaseTLSServerConnection(obj TLSServerConnectioner) *TLSServerConnection {
	return obj.baseTLSServerConnection()
}

// NewTLSServerConnection creates a new ServerConnection wrapping base_io_stream
// (which must have pollable input and output streams).
//
// See the documentation for Connection:base-io-stream for restrictions on when
// application code can run operations on the base_io_stream after this function
// has returned.
//
// The function takes the following parameters:
//
//    - baseIoStream to wrap.
//    - certificate (optional): default server certificate, or NULL.
//
// The function returns the following values:
//
//    - tlsServerConnection: new ServerConnection, or NULL on error.
//
func NewTLSServerConnection(baseIoStream IOStreamer, certificate TLSCertificater) (TLSServerConnectioner, error) {
	var _arg1 *C.GIOStream       // out
	var _arg2 *C.GTlsCertificate // out
	var _cret *C.GIOStream       // in
	var _cerr *C.GError          // in

	_arg1 = (*C.GIOStream)(unsafe.Pointer(baseIoStream.Native()))
	if certificate != nil {
		_arg2 = (*C.GTlsCertificate)(unsafe.Pointer(certificate.Native()))
	}

	_cret = C.g_tls_server_connection_new(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(baseIoStream)
	runtime.KeepAlive(certificate)

	var _tlsServerConnection TLSServerConnectioner // out
	var _goerr error                               // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.TLSServerConnectioner is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj externglib.Objector) bool {
			_, ok := obj.(TLSServerConnectioner)
			return ok
		})
		rv, ok := casted.(TLSServerConnectioner)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.TLSServerConnectioner")
		}
		_tlsServerConnection = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsServerConnection, _goerr
}
