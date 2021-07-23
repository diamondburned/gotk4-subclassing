// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
		{T: externglib.Type(C.g_data_output_stream_get_type()), F: marshalDataOutputStreamer},
	})
}

// DataOutputStream: data output stream implements Stream and includes functions
// for writing data directly to an output stream.
type DataOutputStream struct {
	FilterOutputStream

	Seekable
}

func wrapDataOutputStream(obj *externglib.Object) *DataOutputStream {
	return &DataOutputStream{
		FilterOutputStream: FilterOutputStream{
			OutputStream: OutputStream{
				Object: obj,
			},
		},
		Seekable: Seekable{
			Object: obj,
		},
	}
}

func marshalDataOutputStreamer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDataOutputStream(obj), nil
}

// NewDataOutputStream creates a new data output stream for base_stream.
func NewDataOutputStream(baseStream OutputStreamer) *DataOutputStream {
	var _arg1 *C.GOutputStream     // out
	var _cret *C.GDataOutputStream // in

	_arg1 = (*C.GOutputStream)(unsafe.Pointer(baseStream.Native()))

	_cret = C.g_data_output_stream_new(_arg1)

	var _dataOutputStream *DataOutputStream // out

	_dataOutputStream = wrapDataOutputStream(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dataOutputStream
}

// ByteOrder gets the byte order for the stream.
func (stream *DataOutputStream) ByteOrder() DataStreamByteOrder {
	var _arg0 *C.GDataOutputStream   // out
	var _cret C.GDataStreamByteOrder // in

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(stream.Native()))

	_cret = C.g_data_output_stream_get_byte_order(_arg0)

	var _dataStreamByteOrder DataStreamByteOrder // out

	_dataStreamByteOrder = DataStreamByteOrder(_cret)

	return _dataStreamByteOrder
}

// PutByte puts a byte into the output stream.
func (stream *DataOutputStream) PutByte(ctx context.Context, data byte) error {
	var _arg0 *C.GDataOutputStream // out
	var _arg2 *C.GCancellable      // out
	var _arg1 C.guchar             // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(stream.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.guchar(data)

	C.g_data_output_stream_put_byte(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PutInt16 puts a signed 16-bit integer into the output stream.
func (stream *DataOutputStream) PutInt16(ctx context.Context, data int16) error {
	var _arg0 *C.GDataOutputStream // out
	var _arg2 *C.GCancellable      // out
	var _arg1 C.gint16             // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(stream.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.gint16(data)

	C.g_data_output_stream_put_int16(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PutInt32 puts a signed 32-bit integer into the output stream.
func (stream *DataOutputStream) PutInt32(ctx context.Context, data int32) error {
	var _arg0 *C.GDataOutputStream // out
	var _arg2 *C.GCancellable      // out
	var _arg1 C.gint32             // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(stream.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.gint32(data)

	C.g_data_output_stream_put_int32(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PutInt64 puts a signed 64-bit integer into the stream.
func (stream *DataOutputStream) PutInt64(ctx context.Context, data int64) error {
	var _arg0 *C.GDataOutputStream // out
	var _arg2 *C.GCancellable      // out
	var _arg1 C.gint64             // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(stream.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.gint64(data)

	C.g_data_output_stream_put_int64(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PutString puts a string into the output stream.
func (stream *DataOutputStream) PutString(ctx context.Context, str string) error {
	var _arg0 *C.GDataOutputStream // out
	var _arg2 *C.GCancellable      // out
	var _arg1 *C.char              // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(stream.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_data_output_stream_put_string(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PutUint16 puts an unsigned 16-bit integer into the output stream.
func (stream *DataOutputStream) PutUint16(ctx context.Context, data uint16) error {
	var _arg0 *C.GDataOutputStream // out
	var _arg2 *C.GCancellable      // out
	var _arg1 C.guint16            // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(stream.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.guint16(data)

	C.g_data_output_stream_put_uint16(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PutUint32 puts an unsigned 32-bit integer into the stream.
func (stream *DataOutputStream) PutUint32(ctx context.Context, data uint32) error {
	var _arg0 *C.GDataOutputStream // out
	var _arg2 *C.GCancellable      // out
	var _arg1 C.guint32            // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(stream.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.guint32(data)

	C.g_data_output_stream_put_uint32(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// PutUint64 puts an unsigned 64-bit integer into the stream.
func (stream *DataOutputStream) PutUint64(ctx context.Context, data uint64) error {
	var _arg0 *C.GDataOutputStream // out
	var _arg2 *C.GCancellable      // out
	var _arg1 C.guint64            // out
	var _cerr *C.GError            // in

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(stream.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = C.guint64(data)

	C.g_data_output_stream_put_uint64(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// SetByteOrder sets the byte order of the data output stream to order.
func (stream *DataOutputStream) SetByteOrder(order DataStreamByteOrder) {
	var _arg0 *C.GDataOutputStream   // out
	var _arg1 C.GDataStreamByteOrder // out

	_arg0 = (*C.GDataOutputStream)(unsafe.Pointer(stream.Native()))
	_arg1 = C.GDataStreamByteOrder(order)

	C.g_data_output_stream_set_byte_order(_arg0, _arg1)
}
