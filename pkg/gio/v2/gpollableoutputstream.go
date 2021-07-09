// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.g_pollable_output_stream_get_type()), F: marshalPollableOutputStream},
	})
}

// PollableOutputStreamOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type PollableOutputStreamOverrider interface {
	// CanPoll checks if @stream is actually pollable. Some classes may
	// implement OutputStream but have only certain instances of that class be
	// pollable. If this method returns false, then the behavior of other
	// OutputStream methods is undefined.
	//
	// For any given stream, the value returned by this method is constant; a
	// stream cannot switch from pollable to non-pollable or vice versa.
	CanPoll() bool
	// CreateSource creates a #GSource that triggers when @stream can be
	// written, or @cancellable is triggered or an error occurs. The callback on
	// the source is of the SourceFunc type.
	//
	// As with g_pollable_output_stream_is_writable(), it is possible that the
	// stream may not actually be writable even after the source triggers, so
	// you should use g_pollable_output_stream_write_nonblocking() rather than
	// g_output_stream_write() from the callback.
	CreateSource(cancellable Cancellable) *glib.Source
	// IsWritable checks if @stream can be written.
	//
	// Note that some stream types may not be able to implement this 100%
	// reliably, and it is possible that a call to g_output_stream_write() after
	// this returns true would still block. To guarantee non-blocking behavior,
	// you should always use g_pollable_output_stream_write_nonblocking(), which
	// will return a G_IO_ERROR_WOULD_BLOCK error rather than blocking.
	IsWritable() bool
	// WriteNonblocking attempts to write up to @count bytes from @buffer to
	// @stream, as with g_output_stream_write(). If @stream is not currently
	// writable, this will immediately return G_IO_ERROR_WOULD_BLOCK, and you
	// can use g_pollable_output_stream_create_source() to create a #GSource
	// that will be triggered when @stream is writable.
	//
	// Note that since this method never blocks, you cannot actually use
	// @cancellable to cancel it. However, it will return an error if
	// @cancellable has already been cancelled when you call, which may happen
	// if you call this method after a source triggers due to having been
	// cancelled.
	//
	// Also note that if G_IO_ERROR_WOULD_BLOCK is returned some underlying
	// transports like D/TLS require that you re-send the same @buffer and
	// @count in the next write call.
	WriteNonblocking(buffer []byte) (int, error)
	// WritevNonblocking attempts to write the bytes contained in the @n_vectors
	// @vectors to @stream, as with g_output_stream_writev(). If @stream is not
	// currently writable, this will immediately return
	// %@G_POLLABLE_RETURN_WOULD_BLOCK, and you can use
	// g_pollable_output_stream_create_source() to create a #GSource that will
	// be triggered when @stream is writable. @error will *not* be set in that
	// case.
	//
	// Note that since this method never blocks, you cannot actually use
	// @cancellable to cancel it. However, it will return an error if
	// @cancellable has already been cancelled when you call, which may happen
	// if you call this method after a source triggers due to having been
	// cancelled.
	//
	// Also note that if G_POLLABLE_RETURN_WOULD_BLOCK is returned some
	// underlying transports like D/TLS require that you re-send the same
	// @vectors and @n_vectors in the next write call.
	WritevNonblocking(vectors []OutputVector) (uint, PollableReturn, error)
}

// PollableOutputStream is implemented by Streams that can be polled for
// readiness to write. This can be used when interfacing with a non-GIO API that
// expects UNIX-file-descriptor-style asynchronous I/O rather than GIO-style.
type PollableOutputStream interface {
	gextras.Objector

	// CanPoll checks if @stream is actually pollable. Some classes may
	// implement OutputStream but have only certain instances of that class be
	// pollable. If this method returns false, then the behavior of other
	// OutputStream methods is undefined.
	//
	// For any given stream, the value returned by this method is constant; a
	// stream cannot switch from pollable to non-pollable or vice versa.
	CanPoll() bool
	// CreateSource creates a #GSource that triggers when @stream can be
	// written, or @cancellable is triggered or an error occurs. The callback on
	// the source is of the SourceFunc type.
	//
	// As with g_pollable_output_stream_is_writable(), it is possible that the
	// stream may not actually be writable even after the source triggers, so
	// you should use g_pollable_output_stream_write_nonblocking() rather than
	// g_output_stream_write() from the callback.
	CreateSource(cancellable Cancellable) *glib.Source
	// IsWritable checks if @stream can be written.
	//
	// Note that some stream types may not be able to implement this 100%
	// reliably, and it is possible that a call to g_output_stream_write() after
	// this returns true would still block. To guarantee non-blocking behavior,
	// you should always use g_pollable_output_stream_write_nonblocking(), which
	// will return a G_IO_ERROR_WOULD_BLOCK error rather than blocking.
	IsWritable() bool
	// WriteNonblocking attempts to write up to @count bytes from @buffer to
	// @stream, as with g_output_stream_write(). If @stream is not currently
	// writable, this will immediately return G_IO_ERROR_WOULD_BLOCK, and you
	// can use g_pollable_output_stream_create_source() to create a #GSource
	// that will be triggered when @stream is writable.
	//
	// Note that since this method never blocks, you cannot actually use
	// @cancellable to cancel it. However, it will return an error if
	// @cancellable has already been cancelled when you call, which may happen
	// if you call this method after a source triggers due to having been
	// cancelled.
	//
	// Also note that if G_IO_ERROR_WOULD_BLOCK is returned some underlying
	// transports like D/TLS require that you re-send the same @buffer and
	// @count in the next write call.
	WriteNonblocking(buffer []byte, cancellable Cancellable) (int, error)
	// WritevNonblocking attempts to write the bytes contained in the @n_vectors
	// @vectors to @stream, as with g_output_stream_writev(). If @stream is not
	// currently writable, this will immediately return
	// %@G_POLLABLE_RETURN_WOULD_BLOCK, and you can use
	// g_pollable_output_stream_create_source() to create a #GSource that will
	// be triggered when @stream is writable. @error will *not* be set in that
	// case.
	//
	// Note that since this method never blocks, you cannot actually use
	// @cancellable to cancel it. However, it will return an error if
	// @cancellable has already been cancelled when you call, which may happen
	// if you call this method after a source triggers due to having been
	// cancelled.
	//
	// Also note that if G_POLLABLE_RETURN_WOULD_BLOCK is returned some
	// underlying transports like D/TLS require that you re-send the same
	// @vectors and @n_vectors in the next write call.
	WritevNonblocking(vectors []OutputVector, cancellable Cancellable) (uint, PollableReturn, error)
}

// PollableOutputStreamInterface implements the PollableOutputStream interface.
type PollableOutputStreamInterface struct {
	OutputStreamClass
}

var _ PollableOutputStream = (*PollableOutputStreamInterface)(nil)

func wrapPollableOutputStream(obj *externglib.Object) PollableOutputStream {
	return &PollableOutputStreamInterface{
		OutputStreamClass: OutputStreamClass{
			Object: obj,
		},
	}
}

func marshalPollableOutputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPollableOutputStream(obj), nil
}

// CanPoll checks if @stream is actually pollable. Some classes may implement
// OutputStream but have only certain instances of that class be pollable. If
// this method returns false, then the behavior of other OutputStream methods is
// undefined.
//
// For any given stream, the value returned by this method is constant; a stream
// cannot switch from pollable to non-pollable or vice versa.
func (s *PollableOutputStreamInterface) CanPoll() bool {
	var _arg0 *C.GPollableOutputStream // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(s.Native()))

	_cret = C.g_pollable_output_stream_can_poll(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CreateSource creates a #GSource that triggers when @stream can be written, or
// @cancellable is triggered or an error occurs. The callback on the source is
// of the SourceFunc type.
//
// As with g_pollable_output_stream_is_writable(), it is possible that the
// stream may not actually be writable even after the source triggers, so you
// should use g_pollable_output_stream_write_nonblocking() rather than
// g_output_stream_write() from the callback.
func (s *PollableOutputStreamInterface) CreateSource(cancellable Cancellable) *glib.Source {
	var _arg0 *C.GPollableOutputStream // out
	var _arg1 *C.GCancellable          // out
	var _cret *C.GSource               // in

	_arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_pollable_output_stream_create_source(_arg0, _arg1)

	var _source *glib.Source // out

	_source = (*glib.Source)(unsafe.Pointer(_cret))
	C.g_source_ref(_cret)
	runtime.SetFinalizer(_source, func(v *glib.Source) {
		C.g_source_unref((*C.GSource)(unsafe.Pointer(v)))
	})

	return _source
}

// IsWritable checks if @stream can be written.
//
// Note that some stream types may not be able to implement this 100% reliably,
// and it is possible that a call to g_output_stream_write() after this returns
// true would still block. To guarantee non-blocking behavior, you should always
// use g_pollable_output_stream_write_nonblocking(), which will return a
// G_IO_ERROR_WOULD_BLOCK error rather than blocking.
func (s *PollableOutputStreamInterface) IsWritable() bool {
	var _arg0 *C.GPollableOutputStream // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(s.Native()))

	_cret = C.g_pollable_output_stream_is_writable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// WriteNonblocking attempts to write up to @count bytes from @buffer to
// @stream, as with g_output_stream_write(). If @stream is not currently
// writable, this will immediately return G_IO_ERROR_WOULD_BLOCK, and you can
// use g_pollable_output_stream_create_source() to create a #GSource that will
// be triggered when @stream is writable.
//
// Note that since this method never blocks, you cannot actually use
// @cancellable to cancel it. However, it will return an error if @cancellable
// has already been cancelled when you call, which may happen if you call this
// method after a source triggers due to having been cancelled.
//
// Also note that if G_IO_ERROR_WOULD_BLOCK is returned some underlying
// transports like D/TLS require that you re-send the same @buffer and @count in
// the next write call.
func (s *PollableOutputStreamInterface) WriteNonblocking(buffer []byte, cancellable Cancellable) (int, error) {
	var _arg0 *C.GPollableOutputStream // out
	var _arg1 *C.void
	var _arg2 C.gsize
	var _arg3 *C.GCancellable // out
	var _cret C.gssize        // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(s.Native()))
	_arg2 = C.gsize(len(buffer))
	_arg1 = (*C.void)(unsafe.Pointer(&buffer[0]))
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_pollable_output_stream_write_nonblocking(_arg0, unsafe.Pointer(_arg1), _arg2, _arg3, &_cerr)

	var _gssize int  // out
	var _goerr error // out

	_gssize = int(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _gssize, _goerr
}

// WritevNonblocking attempts to write the bytes contained in the @n_vectors
// @vectors to @stream, as with g_output_stream_writev(). If @stream is not
// currently writable, this will immediately return
// %@G_POLLABLE_RETURN_WOULD_BLOCK, and you can use
// g_pollable_output_stream_create_source() to create a #GSource that will be
// triggered when @stream is writable. @error will *not* be set in that case.
//
// Note that since this method never blocks, you cannot actually use
// @cancellable to cancel it. However, it will return an error if @cancellable
// has already been cancelled when you call, which may happen if you call this
// method after a source triggers due to having been cancelled.
//
// Also note that if G_POLLABLE_RETURN_WOULD_BLOCK is returned some underlying
// transports like D/TLS require that you re-send the same @vectors and
// @n_vectors in the next write call.
func (s *PollableOutputStreamInterface) WritevNonblocking(vectors []OutputVector, cancellable Cancellable) (uint, PollableReturn, error) {
	var _arg0 *C.GPollableOutputStream // out
	var _arg1 *C.GOutputVector
	var _arg2 C.gsize
	var _arg3 C.gsize           // in
	var _arg4 *C.GCancellable   // out
	var _cret C.GPollableReturn // in
	var _cerr *C.GError         // in

	_arg0 = (*C.GPollableOutputStream)(unsafe.Pointer(s.Native()))
	_arg2 = C.gsize(len(vectors))
	_arg1 = (*C.GOutputVector)(unsafe.Pointer(&vectors[0]))
	_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	_cret = C.g_pollable_output_stream_writev_nonblocking(_arg0, _arg1, _arg2, &_arg3, _arg4, &_cerr)

	var _bytesWritten uint             // out
	var _pollableReturn PollableReturn // out
	var _goerr error                   // out

	_bytesWritten = uint(_arg3)
	_pollableReturn = (PollableReturn)(_cret)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _bytesWritten, _pollableReturn, _goerr
}
