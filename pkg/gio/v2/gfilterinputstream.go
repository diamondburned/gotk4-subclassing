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
		{T: externglib.Type(C.g_filter_input_stream_get_type()), F: marshalFilterInputStream},
	})
}

// FilterInputStream: base class for input stream implementations that perform
// some kind of filtering operation on a base stream. Typical examples of
// filtering operations are character set conversion, compression and byte order
// flipping.
type FilterInputStream interface {
	InputStream

	// BaseStream gets the base stream for the filter stream.
	BaseStream(s FilterInputStream)
	// CloseBaseStream returns whether the base stream will be closed when
	// @stream is closed.
	CloseBaseStream(s FilterInputStream) bool
	// SetCloseBaseStream sets whether the base stream will be closed when
	// @stream is closed.
	SetCloseBaseStream(s FilterInputStream, closeBase bool)
}

// filterInputStream implements the FilterInputStream interface.
type filterInputStream struct {
	InputStream
}

var _ FilterInputStream = (*filterInputStream)(nil)

// WrapFilterInputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapFilterInputStream(obj *externglib.Object) FilterInputStream {
	return FilterInputStream{
		InputStream: WrapInputStream(obj),
	}
}

func marshalFilterInputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFilterInputStream(obj), nil
}

// BaseStream gets the base stream for the filter stream.
func (s filterInputStream) BaseStream(s FilterInputStream) {
	var arg0 *C.GFilterInputStream

	arg0 = (*C.GFilterInputStream)(unsafe.Pointer(s.Native()))

	C.g_filter_input_stream_get_base_stream(arg0)
}

// CloseBaseStream returns whether the base stream will be closed when
// @stream is closed.
func (s filterInputStream) CloseBaseStream(s FilterInputStream) bool {
	var arg0 *C.GFilterInputStream

	arg0 = (*C.GFilterInputStream)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_filter_input_stream_get_close_base_stream(arg0)

	if cret {
		ok = true
	}

	return ok
}

// SetCloseBaseStream sets whether the base stream will be closed when
// @stream is closed.
func (s filterInputStream) SetCloseBaseStream(s FilterInputStream, closeBase bool) {
	var arg0 *C.GFilterInputStream
	var arg1 C.gboolean

	arg0 = (*C.GFilterInputStream)(unsafe.Pointer(s.Native()))
	if closeBase {
		arg1 = C.gboolean(1)
	}

	C.g_filter_input_stream_set_close_base_stream(arg0, arg1)
}
