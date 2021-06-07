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
		{T: externglib.Type(C.g_file_output_stream_get_type()), F: marshalFileOutputStream},
	})
}

// FileOutputStream: GFileOutputStream provides output streams that write their
// content to a file.
//
// GFileOutputStream implements #GSeekable, which allows the output stream to
// jump to arbitrary positions in the file and to truncate the file, provided
// the filesystem of the file supports these operations.
//
// To find the position of a file output stream, use g_seekable_tell(). To find
// out if a file output stream supports seeking, use g_seekable_can_seek().To
// position a file output stream, use g_seekable_seek(). To find out if a file
// output stream supports truncating, use g_seekable_can_truncate(). To truncate
// a file output stream, use g_seekable_truncate().
type FileOutputStream interface {
	OutputStream
	Seekable

	// Etag gets the entity tag for the file when it has been written. This must
	// be called after the stream has been written and closed, as the etag can
	// change while writing.
	Etag(s FileOutputStream)
	// QueryInfo queries a file output stream for the given @attributes. This
	// function blocks while querying the stream. For the asynchronous version
	// of this function, see g_file_output_stream_query_info_async(). While the
	// stream is blocked, the stream will set the pending flag internally, and
	// any other operations on the stream will fail with G_IO_ERROR_PENDING.
	//
	// Can fail if the stream was already closed (with @error being set to
	// G_IO_ERROR_CLOSED), the stream has pending operations (with @error being
	// set to G_IO_ERROR_PENDING), or if querying info is not supported for the
	// stream's interface (with @error being set to G_IO_ERROR_NOT_SUPPORTED).
	// In all cases of failure, nil will be returned.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be set, and nil will
	// be returned.
	QueryInfo(s FileOutputStream, attributes string, cancellable Cancellable) error
	// QueryInfoAsync: asynchronously queries the @stream for a Info. When
	// completed, @callback will be called with a Result which can be used to
	// finish the operation with g_file_output_stream_query_info_finish().
	//
	// For the synchronous version of this function, see
	// g_file_output_stream_query_info().
	QueryInfoAsync(s FileOutputStream)
	// QueryInfoFinish finalizes the asynchronous query started by
	// g_file_output_stream_query_info_async().
	QueryInfoFinish(s FileOutputStream, result AsyncResult) error
}

// fileOutputStream implements the FileOutputStream interface.
type fileOutputStream struct {
	OutputStream
	Seekable
}

var _ FileOutputStream = (*fileOutputStream)(nil)

// WrapFileOutputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapFileOutputStream(obj *externglib.Object) FileOutputStream {
	return FileOutputStream{
		OutputStream: WrapOutputStream(obj),
		Seekable:     WrapSeekable(obj),
	}
}

func marshalFileOutputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileOutputStream(obj), nil
}

// Etag gets the entity tag for the file when it has been written. This must
// be called after the stream has been written and closed, as the etag can
// change while writing.
func (s fileOutputStream) Etag(s FileOutputStream) {
	var arg0 *C.GFileOutputStream

	arg0 = (*C.GFileOutputStream)(unsafe.Pointer(s.Native()))

	C.g_file_output_stream_get_etag(arg0)
}

// QueryInfo queries a file output stream for the given @attributes. This
// function blocks while querying the stream. For the asynchronous version
// of this function, see g_file_output_stream_query_info_async(). While the
// stream is blocked, the stream will set the pending flag internally, and
// any other operations on the stream will fail with G_IO_ERROR_PENDING.
//
// Can fail if the stream was already closed (with @error being set to
// G_IO_ERROR_CLOSED), the stream has pending operations (with @error being
// set to G_IO_ERROR_PENDING), or if querying info is not supported for the
// stream's interface (with @error being set to G_IO_ERROR_NOT_SUPPORTED).
// In all cases of failure, nil will be returned.
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be set, and nil will
// be returned.
func (s fileOutputStream) QueryInfo(s FileOutputStream, attributes string, cancellable Cancellable) error {
	var arg0 *C.GFileOutputStream
	var arg1 *C.char
	var arg2 *C.GCancellable

	arg0 = (*C.GFileOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var errout *C.GError
	var err error

	C.g_file_output_stream_query_info(arg0, arg1, arg2, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// QueryInfoAsync: asynchronously queries the @stream for a Info. When
// completed, @callback will be called with a Result which can be used to
// finish the operation with g_file_output_stream_query_info_finish().
//
// For the synchronous version of this function, see
// g_file_output_stream_query_info().
func (s fileOutputStream) QueryInfoAsync(s FileOutputStream) {
	var arg0 *C.GFileOutputStream

	arg0 = (*C.GFileOutputStream)(unsafe.Pointer(s.Native()))

	C.g_file_output_stream_query_info_async(arg0, arg1, arg2, arg3, arg4, arg5)
}

// QueryInfoFinish finalizes the asynchronous query started by
// g_file_output_stream_query_info_async().
func (s fileOutputStream) QueryInfoFinish(s FileOutputStream, result AsyncResult) error {
	var arg0 *C.GFileOutputStream
	var arg1 *C.GAsyncResult

	arg0 = (*C.GFileOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var errout *C.GError
	var err error

	C.g_file_output_stream_query_info_finish(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}
