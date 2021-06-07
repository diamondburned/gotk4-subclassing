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
		{T: externglib.Type(C.g_converter_output_stream_get_type()), F: marshalConverterOutputStream},
	})
}

// ConverterOutputStream: converter output stream implements Stream and allows
// conversion of data of various types during reading.
//
// As of GLib 2.34, OutputStream implements OutputStream.
type ConverterOutputStream interface {
	FilterOutputStream
	PollableOutputStream

	// Converter gets the #GConverter that is used by @converter_stream.
	Converter(c ConverterOutputStream)
}

// converterOutputStream implements the ConverterOutputStream interface.
type converterOutputStream struct {
	FilterOutputStream
	PollableOutputStream
}

var _ ConverterOutputStream = (*converterOutputStream)(nil)

// WrapConverterOutputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapConverterOutputStream(obj *externglib.Object) ConverterOutputStream {
	return ConverterOutputStream{
		FilterOutputStream:   WrapFilterOutputStream(obj),
		PollableOutputStream: WrapPollableOutputStream(obj),
	}
}

func marshalConverterOutputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapConverterOutputStream(obj), nil
}

// NewConverterOutputStream constructs a class ConverterOutputStream.
func NewConverterOutputStream(baseStream OutputStream, converter Converter) {
	var arg1 *C.GOutputStream
	var arg2 *C.GConverter

	arg1 = (*C.GOutputStream)(unsafe.Pointer(baseStream.Native()))
	arg2 = (*C.GConverter)(unsafe.Pointer(converter.Native()))

	C.g_converter_output_stream_new(arg1, arg2)
}

// Converter gets the #GConverter that is used by @converter_stream.
func (c converterOutputStream) Converter(c ConverterOutputStream) {
	var arg0 *C.GConverterOutputStream

	arg0 = (*C.GConverterOutputStream)(unsafe.Pointer(c.Native()))

	C.g_converter_output_stream_get_converter(arg0)
}
