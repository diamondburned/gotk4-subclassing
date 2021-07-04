// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gerror"
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
		{T: externglib.Type(C.g_converter_get_type()), F: marshalConverter},
	})
}

// Converter is implemented by objects that convert binary data in various ways.
// The conversion can be stateful and may fail at any place.
//
// Some example conversions are: character set conversion, compression,
// decompression and regular expression replace.
type Converter interface {
	gextras.Objector

	// Convert resets all internal state in the converter, making it behave as
	// if it was just created. If the converter has any internal state that
	// would produce output then that output is lost.
	Convert(inbuf []byte, outbuf []byte, flags ConverterFlags) (bytesRead uint, bytesWritten uint, converterResult ConverterResult, goerr error)
	// Reset resets all internal state in the converter, making it behave as if
	// it was just created. If the converter has any internal state that would
	// produce output then that output is lost.
	Reset()
}

// converter implements the Converter interface.
type converter struct {
	gextras.Objector
}

var _ Converter = (*converter)(nil)

// WrapConverter wraps a GObject to a type that implements
// interface Converter. It is primarily used internally.
func WrapConverter(obj *externglib.Object) Converter {
	return converter{
		Objector: obj,
	}
}

func marshalConverter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapConverter(obj), nil
}

func (c converter) Convert(inbuf []byte, outbuf []byte, flags ConverterFlags) (bytesRead uint, bytesWritten uint, converterResult ConverterResult, goerr error) {
	var _arg0 *C.GConverter // out
	var _arg1 *C.void
	var _arg2 C.gsize
	var _arg3 *C.void
	var _arg4 C.gsize
	var _arg5 C.GConverterFlags  // out
	var _arg6 C.gsize            // in
	var _arg7 C.gsize            // in
	var _cret C.GConverterResult // in
	var _cerr *C.GError          // in

	_arg0 = (*C.GConverter)(unsafe.Pointer(c.Native()))
	_arg2 = C.gsize(len(inbuf))
	_arg1 = (*C.void)(unsafe.Pointer(&inbuf[0]))
	_arg4 = C.gsize(len(outbuf))
	_arg3 = (*C.void)(unsafe.Pointer(&outbuf[0]))
	_arg5 = C.GConverterFlags(flags)

	_cret = C.g_converter_convert(_arg0, unsafe.Pointer(_arg1), _arg2, unsafe.Pointer(_arg3), _arg4, _arg5, &_arg6, &_arg7, &_cerr)

	var _bytesRead uint                  // out
	var _bytesWritten uint               // out
	var _converterResult ConverterResult // out
	var _goerr error                     // out

	_bytesRead = uint(_arg6)
	_bytesWritten = uint(_arg7)
	_converterResult = ConverterResult(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _bytesRead, _bytesWritten, _converterResult, _goerr
}

func (c converter) Reset() {
	var _arg0 *C.GConverter // out

	_arg0 = (*C.GConverter)(unsafe.Pointer(c.Native()))

	C.g_converter_reset(_arg0)
}