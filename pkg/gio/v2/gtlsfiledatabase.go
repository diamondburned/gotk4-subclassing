// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
		{T: externglib.Type(C.g_tls_file_database_get_type()), F: marshalTLSFileDatabaser},
	})
}

// TLSFileDatabase is implemented by Database objects which load their
// certificate information from a file. It is an interface which TLS library
// specific subtypes implement.
type TLSFileDatabase struct {
	TLSDatabase
}

// TLSFileDatabaser describes TLSFileDatabase's abstract methods.
type TLSFileDatabaser interface {
	externglib.Objector

	privateTLSFileDatabase()
}

var _ TLSFileDatabaser = (*TLSFileDatabase)(nil)

func wrapTLSFileDatabase(obj *externglib.Object) *TLSFileDatabase {
	return &TLSFileDatabase{
		TLSDatabase: TLSDatabase{
			Object: obj,
		},
	}
}

func marshalTLSFileDatabaser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTLSFileDatabase(obj), nil
}

func (*TLSFileDatabase) privateTLSFileDatabase() {}

// NewTLSFileDatabase creates a new FileDatabase which uses anchor certificate
// authorities in anchors to verify certificate chains.
//
// The certificates in anchors must be PEM encoded.
//
// The function takes the following parameters:
//
//    - anchors: filename of anchor certificate authorities.
//
func NewTLSFileDatabase(anchors string) (TLSFileDatabaser, error) {
	var _arg1 *C.gchar        // out
	var _cret *C.GTlsDatabase // in
	var _cerr *C.GError       // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(anchors)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_tls_file_database_new(_arg1, &_cerr)
	runtime.KeepAlive(anchors)

	var _tlsFileDatabase TLSFileDatabaser // out
	var _goerr error                      // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.TLSFileDatabaser is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(TLSFileDatabaser)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.TLSFileDatabaser")
		}
		_tlsFileDatabase = rv
	}
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsFileDatabase, _goerr
}
