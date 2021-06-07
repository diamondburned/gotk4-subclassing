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
		{T: externglib.Type(C.g_charset_converter_get_type()), F: marshalCharsetConverter},
	})
}

// CharsetConverter is an implementation of #GConverter based on GIConv.
type CharsetConverter interface {
	gextras.Objector
	Converter
	Initable

	// NumFallbacks gets the number of fallbacks that @converter has applied so
	// far.
	NumFallbacks(c CharsetConverter)
	// UseFallback gets the Converter:use-fallback property.
	UseFallback(c CharsetConverter) bool
	// SetUseFallback sets the Converter:use-fallback property.
	SetUseFallback(c CharsetConverter, useFallback bool)
}

// charsetConverter implements the CharsetConverter interface.
type charsetConverter struct {
	gextras.Objector
	Converter
	Initable
}

var _ CharsetConverter = (*charsetConverter)(nil)

// WrapCharsetConverter wraps a GObject to the right type. It is
// primarily used internally.
func WrapCharsetConverter(obj *externglib.Object) CharsetConverter {
	return CharsetConverter{
		Objector:  obj,
		Converter: WrapConverter(obj),
		Initable:  WrapInitable(obj),
	}
}

func marshalCharsetConverter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCharsetConverter(obj), nil
}

// NewCharsetConverter constructs a class CharsetConverter.
func NewCharsetConverter(toCharset string, fromCharset string) error {
	var arg1 *C.gchar
	var arg2 *C.gchar

	arg1 = (*C.gchar)(C.CString(toCharset))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(fromCharset))
	defer C.free(unsafe.Pointer(arg2))

	var errout *C.GError
	var err error

	C.g_charset_converter_new(arg1, arg2, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// NumFallbacks gets the number of fallbacks that @converter has applied so
// far.
func (c charsetConverter) NumFallbacks(c CharsetConverter) {
	var arg0 *C.GCharsetConverter

	arg0 = (*C.GCharsetConverter)(unsafe.Pointer(c.Native()))

	C.g_charset_converter_get_num_fallbacks(arg0)
}

// UseFallback gets the Converter:use-fallback property.
func (c charsetConverter) UseFallback(c CharsetConverter) bool {
	var arg0 *C.GCharsetConverter

	arg0 = (*C.GCharsetConverter)(unsafe.Pointer(c.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_charset_converter_get_use_fallback(arg0)

	if cret {
		ok = true
	}

	return ok
}

// SetUseFallback sets the Converter:use-fallback property.
func (c charsetConverter) SetUseFallback(c CharsetConverter, useFallback bool) {
	var arg0 *C.GCharsetConverter
	var arg1 C.gboolean

	arg0 = (*C.GCharsetConverter)(unsafe.Pointer(c.Native()))
	if useFallback {
		arg1 = C.gboolean(1)
	}

	C.g_charset_converter_set_use_fallback(arg0, arg1)
}