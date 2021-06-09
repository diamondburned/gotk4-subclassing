// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"github.com/diamondburned/gotk4/internal/gerror"
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
		{T: externglib.Type(C.g_output_stream_get_type()), F: marshalOutputStream},
	})
}

// OutputStream has functions to write to a stream (g_output_stream_write()), to
// close a stream (g_output_stream_close()) and to flush pending writes
// (g_output_stream_flush()).
//
// To copy the content of an input stream to an output stream without manually
// handling the reads and writes, use g_output_stream_splice().
//
// See the documentation for OStream for details of thread safety of streaming
// APIs.
//
// All of these functions have async variants too.
type OutputStream interface {
	gextras.Objector

	// ClearPending clears the pending flag on @stream.
	ClearPending()
	// Close closes the stream, releasing resources related to it.
	//
	// Once the stream is closed, all other operations will return
	// G_IO_ERROR_CLOSED. Closing a stream multiple times will not return an
	// error.
	//
	// Closing a stream will automatically flush any outstanding buffers in the
	// stream.
	//
	// Streams will be automatically closed when the last reference is dropped,
	// but you might want to call this function to make sure resources are
	// released as early as possible.
	//
	// Some streams might keep the backing store of the stream (e.g. a file
	// descriptor) open after the stream is closed. See the documentation for
	// the individual stream for details.
	//
	// On failure the first error that happened will be reported, but the close
	// operation will finish as much as possible. A stream that failed to close
	// will still return G_IO_ERROR_CLOSED for all operations. Still, it is
	// important to check and report the error to the user, otherwise there
	// might be a loss of data as all data might not be written.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	// Cancelling a close will still leave the stream closed, but there some
	// streams can use a faster close that doesn't block to e.g. check errors.
	// On cancellation (as with any error) there is no guarantee that all
	// written data will reach the target.
	Close(cancellable Cancellable) error
	// CloseAsync requests an asynchronous close of the stream, releasing
	// resources related to it. When the operation is finished @callback will be
	// called. You can then call g_output_stream_close_finish() to get the
	// result of the operation.
	//
	// For behaviour details see g_output_stream_close().
	//
	// The asynchronous methods have a default fallback that uses threads to
	// implement asynchronicity, so they are optional for inheriting classes.
	// However, if you override one you must override all.
	CloseAsync()
	// CloseFinish closes an output stream.
	CloseFinish(result AsyncResult) error
	// Flush forces a write of all user-space buffered data for the given
	// @stream. Will block during the operation. Closing the stream will
	// implicitly cause a flush.
	//
	// This function is optional for inherited classes.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	Flush(cancellable Cancellable) error
	// FlushAsync forces an asynchronous write of all user-space buffered data
	// for the given @stream. For behaviour details see g_output_stream_flush().
	//
	// When the operation is finished @callback will be called. You can then
	// call g_output_stream_flush_finish() to get the result of the operation.
	FlushAsync()
	// FlushFinish finishes flushing an output stream.
	FlushFinish(result AsyncResult) error
	// HasPending checks if an output stream has pending actions.
	HasPending() bool
	// IsClosed checks if an output stream has already been closed.
	IsClosed() bool
	// IsClosing checks if an output stream is being closed. This can be used
	// inside e.g. a flush implementation to see if the flush (or other i/o
	// operation) is called from within the closing operation.
	IsClosing() bool
	// SetPending sets @stream to have actions pending. If the pending flag is
	// already set or @stream is closed, it will return false and set @error.
	SetPending() error
	// Splice splices an input stream into an output stream.
	Splice(source InputStream, flags OutputStreamSpliceFlags, cancellable Cancellable) (gssize int, goerr error)
	// SpliceAsync splices a stream asynchronously. When the operation is
	// finished @callback will be called. You can then call
	// g_output_stream_splice_finish() to get the result of the operation.
	//
	// For the synchronous, blocking version of this function, see
	// g_output_stream_splice().
	SpliceAsync()
	// SpliceFinish finishes an asynchronous stream splice operation.
	SpliceFinish(result AsyncResult) (gssize int, goerr error)
	// WriteAllAsync: request an asynchronous write of @count bytes from @buffer
	// into the stream. When the operation is finished @callback will be called.
	// You can then call g_output_stream_write_all_finish() to get the result of
	// the operation.
	//
	// This is the asynchronous version of g_output_stream_write_all().
	//
	// Call g_output_stream_write_all_finish() to collect the result.
	//
	// Any outstanding I/O request with higher priority (lower numerical value)
	// will be executed before an outstanding request with lower priority.
	// Default priority is G_PRIORITY_DEFAULT.
	//
	// Note that no copy of @buffer will be made, so it must stay valid until
	// @callback is called.
	WriteAllAsync()
	// WriteAllFinish finishes an asynchronous stream write operation started
	// with g_output_stream_write_all_async().
	//
	// As a special exception to the normal conventions for functions that use
	// #GError, if this function returns false (and sets @error) then
	// @bytes_written will be set to the number of bytes that were successfully
	// written before the error was encountered. This functionality is only
	// available from C. If you need it from another language then you must
	// write your own loop around g_output_stream_write_async().
	WriteAllFinish(result AsyncResult) (bytesWritten uint, goerr error)
	// WriteAsync: request an asynchronous write of @count bytes from @buffer
	// into the stream. When the operation is finished @callback will be called.
	// You can then call g_output_stream_write_finish() to get the result of the
	// operation.
	//
	// During an async request no other sync and async calls are allowed, and
	// will result in G_IO_ERROR_PENDING errors.
	//
	// A value of @count larger than G_MAXSSIZE will cause a
	// G_IO_ERROR_INVALID_ARGUMENT error.
	//
	// On success, the number of bytes written will be passed to the @callback.
	// It is not an error if this is not the same as the requested size, as it
	// can happen e.g. on a partial I/O error, but generally we try to write as
	// many bytes as requested.
	//
	// You are guaranteed that this method will never fail with
	// G_IO_ERROR_WOULD_BLOCK - if @stream can't accept more data, the method
	// will just wait until this changes.
	//
	// Any outstanding I/O request with higher priority (lower numerical value)
	// will be executed before an outstanding request with lower priority.
	// Default priority is G_PRIORITY_DEFAULT.
	//
	// The asynchronous methods have a default fallback that uses threads to
	// implement asynchronicity, so they are optional for inheriting classes.
	// However, if you override one you must override all.
	//
	// For the synchronous, blocking version of this function, see
	// g_output_stream_write().
	//
	// Note that no copy of @buffer will be made, so it must stay valid until
	// @callback is called. See g_output_stream_write_bytes_async() for a
	// #GBytes version that will automatically hold a reference to the contents
	// (without copying) for the duration of the call.
	WriteAsync()
	// WriteBytesAsync: this function is similar to
	// g_output_stream_write_async(), but takes a #GBytes as input. Due to the
	// refcounted nature of #GBytes, this allows the stream to avoid taking a
	// copy of the data.
	//
	// However, note that this function may still perform partial writes, just
	// like g_output_stream_write_async(). If that occurs, to continue writing,
	// you will need to create a new #GBytes containing just the remaining
	// bytes, using g_bytes_new_from_bytes(). Passing the same #GBytes instance
	// multiple times potentially can result in duplicated data in the output
	// stream.
	//
	// For the synchronous, blocking version of this function, see
	// g_output_stream_write_bytes().
	WriteBytesAsync()
	// WriteBytesFinish finishes a stream write-from-#GBytes operation.
	WriteBytesFinish(result AsyncResult) (gssize int, goerr error)
	// WriteFinish finishes a stream write operation.
	WriteFinish(result AsyncResult) (gssize int, goerr error)
	// WritevAllAsync: request an asynchronous write of the bytes contained in
	// the @n_vectors @vectors into the stream. When the operation is finished
	// @callback will be called. You can then call
	// g_output_stream_writev_all_finish() to get the result of the operation.
	//
	// This is the asynchronous version of g_output_stream_writev_all().
	//
	// Call g_output_stream_writev_all_finish() to collect the result.
	//
	// Any outstanding I/O request with higher priority (lower numerical value)
	// will be executed before an outstanding request with lower priority.
	// Default priority is G_PRIORITY_DEFAULT.
	//
	// Note that no copy of @vectors will be made, so it must stay valid until
	// @callback is called. The content of the individual elements of @vectors
	// might be changed by this function.
	WritevAllAsync()
	// WritevAllFinish finishes an asynchronous stream write operation started
	// with g_output_stream_writev_all_async().
	//
	// As a special exception to the normal conventions for functions that use
	// #GError, if this function returns false (and sets @error) then
	// @bytes_written will be set to the number of bytes that were successfully
	// written before the error was encountered. This functionality is only
	// available from C. If you need it from another language then you must
	// write your own loop around g_output_stream_writev_async().
	WritevAllFinish(result AsyncResult) (bytesWritten uint, goerr error)
	// WritevAsync: request an asynchronous write of the bytes contained in
	// @n_vectors @vectors into the stream. When the operation is finished
	// @callback will be called. You can then call
	// g_output_stream_writev_finish() to get the result of the operation.
	//
	// During an async request no other sync and async calls are allowed, and
	// will result in G_IO_ERROR_PENDING errors.
	//
	// On success, the number of bytes written will be passed to the @callback.
	// It is not an error if this is not the same as the requested size, as it
	// can happen e.g. on a partial I/O error, but generally we try to write as
	// many bytes as requested.
	//
	// You are guaranteed that this method will never fail with
	// G_IO_ERROR_WOULD_BLOCK — if @stream can't accept more data, the method
	// will just wait until this changes.
	//
	// Any outstanding I/O request with higher priority (lower numerical value)
	// will be executed before an outstanding request with lower priority.
	// Default priority is G_PRIORITY_DEFAULT.
	//
	// The asynchronous methods have a default fallback that uses threads to
	// implement asynchronicity, so they are optional for inheriting classes.
	// However, if you override one you must override all.
	//
	// For the synchronous, blocking version of this function, see
	// g_output_stream_writev().
	//
	// Note that no copy of @vectors will be made, so it must stay valid until
	// @callback is called.
	WritevAsync()
	// WritevFinish finishes a stream writev operation.
	WritevFinish(result AsyncResult) (bytesWritten uint, goerr error)
}

// outputStream implements the OutputStream interface.
type outputStream struct {
	gextras.Objector
}

var _ OutputStream = (*outputStream)(nil)

// WrapOutputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapOutputStream(obj *externglib.Object) OutputStream {
	return OutputStream{
		Objector: obj,
	}
}

func marshalOutputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapOutputStream(obj), nil
}

// ClearPending clears the pending flag on @stream.
func (s outputStream) ClearPending() {
	var arg0 *C.GOutputStream

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))

	C.g_output_stream_clear_pending(arg0)
}

// Close closes the stream, releasing resources related to it.
//
// Once the stream is closed, all other operations will return
// G_IO_ERROR_CLOSED. Closing a stream multiple times will not return an
// error.
//
// Closing a stream will automatically flush any outstanding buffers in the
// stream.
//
// Streams will be automatically closed when the last reference is dropped,
// but you might want to call this function to make sure resources are
// released as early as possible.
//
// Some streams might keep the backing store of the stream (e.g. a file
// descriptor) open after the stream is closed. See the documentation for
// the individual stream for details.
//
// On failure the first error that happened will be reported, but the close
// operation will finish as much as possible. A stream that failed to close
// will still return G_IO_ERROR_CLOSED for all operations. Still, it is
// important to check and report the error to the user, otherwise there
// might be a loss of data as all data might not be written.
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
// Cancelling a close will still leave the stream closed, but there some
// streams can use a faster close that doesn't block to e.g. check errors.
// On cancellation (as with any error) there is no guarantee that all
// written data will reach the target.
func (s outputStream) Close(cancellable Cancellable) error {
	var arg0 *C.GOutputStream
	var arg1 *C.GCancellable

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cerr *C.GError

	C.g_output_stream_close(arg0, arg1, cerr)

	var goerr error

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// CloseAsync requests an asynchronous close of the stream, releasing
// resources related to it. When the operation is finished @callback will be
// called. You can then call g_output_stream_close_finish() to get the
// result of the operation.
//
// For behaviour details see g_output_stream_close().
//
// The asynchronous methods have a default fallback that uses threads to
// implement asynchronicity, so they are optional for inheriting classes.
// However, if you override one you must override all.
func (s outputStream) CloseAsync() {
	var arg0 *C.GOutputStream

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))

	C.g_output_stream_close_async(arg0)
}

// CloseFinish closes an output stream.
func (s outputStream) CloseFinish(result AsyncResult) error {
	var arg0 *C.GOutputStream
	var arg1 *C.GAsyncResult

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var cerr *C.GError

	C.g_output_stream_close_finish(arg0, arg1, cerr)

	var goerr error

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// Flush forces a write of all user-space buffered data for the given
// @stream. Will block during the operation. Closing the stream will
// implicitly cause a flush.
//
// This function is optional for inherited classes.
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
func (s outputStream) Flush(cancellable Cancellable) error {
	var arg0 *C.GOutputStream
	var arg1 *C.GCancellable

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cerr *C.GError

	C.g_output_stream_flush(arg0, arg1, cerr)

	var goerr error

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// FlushAsync forces an asynchronous write of all user-space buffered data
// for the given @stream. For behaviour details see g_output_stream_flush().
//
// When the operation is finished @callback will be called. You can then
// call g_output_stream_flush_finish() to get the result of the operation.
func (s outputStream) FlushAsync() {
	var arg0 *C.GOutputStream

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))

	C.g_output_stream_flush_async(arg0)
}

// FlushFinish finishes flushing an output stream.
func (s outputStream) FlushFinish(result AsyncResult) error {
	var arg0 *C.GOutputStream
	var arg1 *C.GAsyncResult

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var cerr *C.GError

	C.g_output_stream_flush_finish(arg0, arg1, cerr)

	var goerr error

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// HasPending checks if an output stream has pending actions.
func (s outputStream) HasPending() bool {
	var arg0 *C.GOutputStream

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))

	var cret C.gboolean

	cret = C.g_output_stream_has_pending(arg0)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// IsClosed checks if an output stream has already been closed.
func (s outputStream) IsClosed() bool {
	var arg0 *C.GOutputStream

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))

	var cret C.gboolean

	cret = C.g_output_stream_is_closed(arg0)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// IsClosing checks if an output stream is being closed. This can be used
// inside e.g. a flush implementation to see if the flush (or other i/o
// operation) is called from within the closing operation.
func (s outputStream) IsClosing() bool {
	var arg0 *C.GOutputStream

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))

	var cret C.gboolean

	cret = C.g_output_stream_is_closing(arg0)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// SetPending sets @stream to have actions pending. If the pending flag is
// already set or @stream is closed, it will return false and set @error.
func (s outputStream) SetPending() error {
	var arg0 *C.GOutputStream

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))

	var cerr *C.GError

	C.g_output_stream_set_pending(arg0, cerr)

	var goerr error

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// Splice splices an input stream into an output stream.
func (s outputStream) Splice(source InputStream, flags OutputStreamSpliceFlags, cancellable Cancellable) (gssize int, goerr error) {
	var arg0 *C.GOutputStream
	var arg1 *C.GInputStream
	var arg2 C.GOutputStreamSpliceFlags
	var arg3 *C.GCancellable

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GInputStream)(unsafe.Pointer(source.Native()))
	arg2 = (C.GOutputStreamSpliceFlags)(flags)
	arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cret C.gssize
	var cerr *C.GError

	cret = C.g_output_stream_splice(arg0, arg1, arg2, arg3, cerr)

	var gssize int
	var goerr error

	gssize = (int)(cret)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return gssize, goerr
}

// SpliceAsync splices a stream asynchronously. When the operation is
// finished @callback will be called. You can then call
// g_output_stream_splice_finish() to get the result of the operation.
//
// For the synchronous, blocking version of this function, see
// g_output_stream_splice().
func (s outputStream) SpliceAsync() {
	var arg0 *C.GOutputStream

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))

	C.g_output_stream_splice_async(arg0)
}

// SpliceFinish finishes an asynchronous stream splice operation.
func (s outputStream) SpliceFinish(result AsyncResult) (gssize int, goerr error) {
	var arg0 *C.GOutputStream
	var arg1 *C.GAsyncResult

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var cret C.gssize
	var cerr *C.GError

	cret = C.g_output_stream_splice_finish(arg0, arg1, cerr)

	var gssize int
	var goerr error

	gssize = (int)(cret)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return gssize, goerr
}

// WriteAllAsync: request an asynchronous write of @count bytes from @buffer
// into the stream. When the operation is finished @callback will be called.
// You can then call g_output_stream_write_all_finish() to get the result of
// the operation.
//
// This is the asynchronous version of g_output_stream_write_all().
//
// Call g_output_stream_write_all_finish() to collect the result.
//
// Any outstanding I/O request with higher priority (lower numerical value)
// will be executed before an outstanding request with lower priority.
// Default priority is G_PRIORITY_DEFAULT.
//
// Note that no copy of @buffer will be made, so it must stay valid until
// @callback is called.
func (s outputStream) WriteAllAsync() {
	var arg0 *C.GOutputStream

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))

	C.g_output_stream_write_all_async(arg0)
}

// WriteAllFinish finishes an asynchronous stream write operation started
// with g_output_stream_write_all_async().
//
// As a special exception to the normal conventions for functions that use
// #GError, if this function returns false (and sets @error) then
// @bytes_written will be set to the number of bytes that were successfully
// written before the error was encountered. This functionality is only
// available from C. If you need it from another language then you must
// write your own loop around g_output_stream_write_async().
func (s outputStream) WriteAllFinish(result AsyncResult) (bytesWritten uint, goerr error) {
	var arg0 *C.GOutputStream
	var arg1 *C.GAsyncResult

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var arg2 C.gsize
	var cerr *C.GError

	C.g_output_stream_write_all_finish(arg0, arg1, &arg2, cerr)

	var bytesWritten uint
	var goerr error

	bytesWritten = (uint)(arg2)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return bytesWritten, goerr
}

// WriteAsync: request an asynchronous write of @count bytes from @buffer
// into the stream. When the operation is finished @callback will be called.
// You can then call g_output_stream_write_finish() to get the result of the
// operation.
//
// During an async request no other sync and async calls are allowed, and
// will result in G_IO_ERROR_PENDING errors.
//
// A value of @count larger than G_MAXSSIZE will cause a
// G_IO_ERROR_INVALID_ARGUMENT error.
//
// On success, the number of bytes written will be passed to the @callback.
// It is not an error if this is not the same as the requested size, as it
// can happen e.g. on a partial I/O error, but generally we try to write as
// many bytes as requested.
//
// You are guaranteed that this method will never fail with
// G_IO_ERROR_WOULD_BLOCK - if @stream can't accept more data, the method
// will just wait until this changes.
//
// Any outstanding I/O request with higher priority (lower numerical value)
// will be executed before an outstanding request with lower priority.
// Default priority is G_PRIORITY_DEFAULT.
//
// The asynchronous methods have a default fallback that uses threads to
// implement asynchronicity, so they are optional for inheriting classes.
// However, if you override one you must override all.
//
// For the synchronous, blocking version of this function, see
// g_output_stream_write().
//
// Note that no copy of @buffer will be made, so it must stay valid until
// @callback is called. See g_output_stream_write_bytes_async() for a
// #GBytes version that will automatically hold a reference to the contents
// (without copying) for the duration of the call.
func (s outputStream) WriteAsync() {
	var arg0 *C.GOutputStream

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))

	C.g_output_stream_write_async(arg0)
}

// WriteBytesAsync: this function is similar to
// g_output_stream_write_async(), but takes a #GBytes as input. Due to the
// refcounted nature of #GBytes, this allows the stream to avoid taking a
// copy of the data.
//
// However, note that this function may still perform partial writes, just
// like g_output_stream_write_async(). If that occurs, to continue writing,
// you will need to create a new #GBytes containing just the remaining
// bytes, using g_bytes_new_from_bytes(). Passing the same #GBytes instance
// multiple times potentially can result in duplicated data in the output
// stream.
//
// For the synchronous, blocking version of this function, see
// g_output_stream_write_bytes().
func (s outputStream) WriteBytesAsync() {
	var arg0 *C.GOutputStream

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))

	C.g_output_stream_write_bytes_async(arg0)
}

// WriteBytesFinish finishes a stream write-from-#GBytes operation.
func (s outputStream) WriteBytesFinish(result AsyncResult) (gssize int, goerr error) {
	var arg0 *C.GOutputStream
	var arg1 *C.GAsyncResult

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var cret C.gssize
	var cerr *C.GError

	cret = C.g_output_stream_write_bytes_finish(arg0, arg1, cerr)

	var gssize int
	var goerr error

	gssize = (int)(cret)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return gssize, goerr
}

// WriteFinish finishes a stream write operation.
func (s outputStream) WriteFinish(result AsyncResult) (gssize int, goerr error) {
	var arg0 *C.GOutputStream
	var arg1 *C.GAsyncResult

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var cret C.gssize
	var cerr *C.GError

	cret = C.g_output_stream_write_finish(arg0, arg1, cerr)

	var gssize int
	var goerr error

	gssize = (int)(cret)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return gssize, goerr
}

// WritevAllAsync: request an asynchronous write of the bytes contained in
// the @n_vectors @vectors into the stream. When the operation is finished
// @callback will be called. You can then call
// g_output_stream_writev_all_finish() to get the result of the operation.
//
// This is the asynchronous version of g_output_stream_writev_all().
//
// Call g_output_stream_writev_all_finish() to collect the result.
//
// Any outstanding I/O request with higher priority (lower numerical value)
// will be executed before an outstanding request with lower priority.
// Default priority is G_PRIORITY_DEFAULT.
//
// Note that no copy of @vectors will be made, so it must stay valid until
// @callback is called. The content of the individual elements of @vectors
// might be changed by this function.
func (s outputStream) WritevAllAsync() {
	var arg0 *C.GOutputStream

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))

	C.g_output_stream_writev_all_async(arg0)
}

// WritevAllFinish finishes an asynchronous stream write operation started
// with g_output_stream_writev_all_async().
//
// As a special exception to the normal conventions for functions that use
// #GError, if this function returns false (and sets @error) then
// @bytes_written will be set to the number of bytes that were successfully
// written before the error was encountered. This functionality is only
// available from C. If you need it from another language then you must
// write your own loop around g_output_stream_writev_async().
func (s outputStream) WritevAllFinish(result AsyncResult) (bytesWritten uint, goerr error) {
	var arg0 *C.GOutputStream
	var arg1 *C.GAsyncResult

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var arg2 C.gsize
	var cerr *C.GError

	C.g_output_stream_writev_all_finish(arg0, arg1, &arg2, cerr)

	var bytesWritten uint
	var goerr error

	bytesWritten = (uint)(arg2)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return bytesWritten, goerr
}

// WritevAsync: request an asynchronous write of the bytes contained in
// @n_vectors @vectors into the stream. When the operation is finished
// @callback will be called. You can then call
// g_output_stream_writev_finish() to get the result of the operation.
//
// During an async request no other sync and async calls are allowed, and
// will result in G_IO_ERROR_PENDING errors.
//
// On success, the number of bytes written will be passed to the @callback.
// It is not an error if this is not the same as the requested size, as it
// can happen e.g. on a partial I/O error, but generally we try to write as
// many bytes as requested.
//
// You are guaranteed that this method will never fail with
// G_IO_ERROR_WOULD_BLOCK — if @stream can't accept more data, the method
// will just wait until this changes.
//
// Any outstanding I/O request with higher priority (lower numerical value)
// will be executed before an outstanding request with lower priority.
// Default priority is G_PRIORITY_DEFAULT.
//
// The asynchronous methods have a default fallback that uses threads to
// implement asynchronicity, so they are optional for inheriting classes.
// However, if you override one you must override all.
//
// For the synchronous, blocking version of this function, see
// g_output_stream_writev().
//
// Note that no copy of @vectors will be made, so it must stay valid until
// @callback is called.
func (s outputStream) WritevAsync() {
	var arg0 *C.GOutputStream

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))

	C.g_output_stream_writev_async(arg0)
}

// WritevFinish finishes a stream writev operation.
func (s outputStream) WritevFinish(result AsyncResult) (bytesWritten uint, goerr error) {
	var arg0 *C.GOutputStream
	var arg1 *C.GAsyncResult

	arg0 = (*C.GOutputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var arg2 C.gsize
	var cerr *C.GError

	C.g_output_stream_writev_finish(arg0, arg1, &arg2, cerr)

	var bytesWritten uint
	var goerr error

	bytesWritten = (uint)(arg2)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return bytesWritten, goerr
}
