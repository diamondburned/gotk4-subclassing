// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
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
		{T: externglib.Type(C.g_data_input_stream_get_type()), F: marshalDataInputStream},
	})
}

// DataInputStream: data input stream implements Stream and includes functions
// for reading structured data directly from a binary input stream.
type DataInputStream interface {
	BufferedInputStream
	Seekable

	// ByteOrder gets the byte order for the data input stream.
	ByteOrder() DataStreamByteOrder
	// NewlineType gets the current newline type for the @stream.
	NewlineType() DataStreamNewlineType
	// ReadByte reads an unsigned 8-bit/1-byte value from @stream.
	ReadByte(cancellable Cancellable) (guint8 byte, goerr error)
	// ReadInt16 reads a 16-bit/2-byte value from @stream.
	//
	// In order to get the correct byte order for this read operation, see
	// g_data_input_stream_get_byte_order() and
	// g_data_input_stream_set_byte_order().
	ReadInt16(cancellable Cancellable) (gint16 int16, goerr error)
	// ReadInt32 reads a signed 32-bit/4-byte value from @stream.
	//
	// In order to get the correct byte order for this read operation, see
	// g_data_input_stream_get_byte_order() and
	// g_data_input_stream_set_byte_order().
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	ReadInt32(cancellable Cancellable) (gint32 int32, goerr error)
	// ReadInt64 reads a 64-bit/8-byte value from @stream.
	//
	// In order to get the correct byte order for this read operation, see
	// g_data_input_stream_get_byte_order() and
	// g_data_input_stream_set_byte_order().
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	ReadInt64(cancellable Cancellable) (gint64 int64, goerr error)
	// ReadLine reads a line from the data input stream. Note that no encoding
	// checks or conversion is performed; the input is not guaranteed to be
	// UTF-8, and may in fact have embedded NUL characters.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	ReadLine(cancellable Cancellable) (length uint, guint8s []byte, goerr error)
	// ReadLineAsync: the asynchronous version of
	// g_data_input_stream_read_line(). It is an error to have two outstanding
	// calls to this function.
	//
	// When the operation is finished, @callback will be called. You can then
	// call g_data_input_stream_read_line_finish() to get the result of the
	// operation.
	ReadLineAsync()
	// ReadLineFinish: finish an asynchronous call started by
	// g_data_input_stream_read_line_async(). Note the warning about string
	// encoding in g_data_input_stream_read_line() applies here as well.
	ReadLineFinish(result AsyncResult) (length uint, guint8s []byte, goerr error)
	// ReadLineFinishUTF8: finish an asynchronous call started by
	// g_data_input_stream_read_line_async().
	ReadLineFinishUTF8(result AsyncResult) (length uint, utf8 string, goerr error)
	// ReadLineUTF8 reads a UTF-8 encoded line from the data input stream.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	ReadLineUTF8(cancellable Cancellable) (length uint, utf8 string, goerr error)
	// ReadUint16 reads an unsigned 16-bit/2-byte value from @stream.
	//
	// In order to get the correct byte order for this read operation, see
	// g_data_input_stream_get_byte_order() and
	// g_data_input_stream_set_byte_order().
	ReadUint16(cancellable Cancellable) (guint16 uint16, goerr error)
	// ReadUint32 reads an unsigned 32-bit/4-byte value from @stream.
	//
	// In order to get the correct byte order for this read operation, see
	// g_data_input_stream_get_byte_order() and
	// g_data_input_stream_set_byte_order().
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	ReadUint32(cancellable Cancellable) (guint32 uint32, goerr error)
	// ReadUint64 reads an unsigned 64-bit/8-byte value from @stream.
	//
	// In order to get the correct byte order for this read operation, see
	// g_data_input_stream_get_byte_order().
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
	ReadUint64(cancellable Cancellable) (guint64 uint64, goerr error)
	// ReadUntil reads a string from the data input stream, up to the first
	// occurrence of any of the stop characters.
	//
	// Note that, in contrast to g_data_input_stream_read_until_async(), this
	// function consumes the stop character that it finds.
	//
	// Don't use this function in new code. Its functionality is inconsistent
	// with g_data_input_stream_read_until_async(). Both functions will be
	// marked as deprecated in a future release. Use
	// g_data_input_stream_read_upto() instead, but note that that function does
	// not consume the stop character.
	ReadUntil(stopChars string, cancellable Cancellable) (length uint, utf8 string, goerr error)
	// ReadUntilAsync: the asynchronous version of
	// g_data_input_stream_read_until(). It is an error to have two outstanding
	// calls to this function.
	//
	// Note that, in contrast to g_data_input_stream_read_until(), this function
	// does not consume the stop character that it finds. You must read it for
	// yourself.
	//
	// When the operation is finished, @callback will be called. You can then
	// call g_data_input_stream_read_until_finish() to get the result of the
	// operation.
	//
	// Don't use this function in new code. Its functionality is inconsistent
	// with g_data_input_stream_read_until(). Both functions will be marked as
	// deprecated in a future release. Use g_data_input_stream_read_upto_async()
	// instead.
	ReadUntilAsync()
	// ReadUntilFinish: finish an asynchronous call started by
	// g_data_input_stream_read_until_async().
	ReadUntilFinish(result AsyncResult) (length uint, utf8 string, goerr error)
	// ReadUpto reads a string from the data input stream, up to the first
	// occurrence of any of the stop characters.
	//
	// In contrast to g_data_input_stream_read_until(), this function does not
	// consume the stop character. You have to use
	// g_data_input_stream_read_byte() to get it before calling
	// g_data_input_stream_read_upto() again.
	//
	// Note that @stop_chars may contain '\0' if @stop_chars_len is specified.
	//
	// The returned string will always be nul-terminated on success.
	ReadUpto(stopChars string, stopCharsLen int, cancellable Cancellable) (length uint, utf8 string, goerr error)
	// ReadUptoAsync: the asynchronous version of
	// g_data_input_stream_read_upto(). It is an error to have two outstanding
	// calls to this function.
	//
	// In contrast to g_data_input_stream_read_until(), this function does not
	// consume the stop character. You have to use
	// g_data_input_stream_read_byte() to get it before calling
	// g_data_input_stream_read_upto() again.
	//
	// Note that @stop_chars may contain '\0' if @stop_chars_len is specified.
	//
	// When the operation is finished, @callback will be called. You can then
	// call g_data_input_stream_read_upto_finish() to get the result of the
	// operation.
	ReadUptoAsync()
	// ReadUptoFinish: finish an asynchronous call started by
	// g_data_input_stream_read_upto_async().
	//
	// Note that this function does not consume the stop character. You have to
	// use g_data_input_stream_read_byte() to get it before calling
	// g_data_input_stream_read_upto_async() again.
	//
	// The returned string will always be nul-terminated on success.
	ReadUptoFinish(result AsyncResult) (length uint, utf8 string, goerr error)
	// SetByteOrder: this function sets the byte order for the given @stream.
	// All subsequent reads from the @stream will be read in the given @order.
	SetByteOrder(order DataStreamByteOrder)
	// SetNewlineType sets the newline type for the @stream.
	//
	// Note that using G_DATA_STREAM_NEWLINE_TYPE_ANY is slightly unsafe. If a
	// read chunk ends in "CR" we must read an additional byte to know if this
	// is "CR" or "CR LF", and this might block if there is no more data
	// available.
	SetNewlineType(typ DataStreamNewlineType)
}

// dataInputStream implements the DataInputStream interface.
type dataInputStream struct {
	BufferedInputStream
	Seekable
}

var _ DataInputStream = (*dataInputStream)(nil)

// WrapDataInputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapDataInputStream(obj *externglib.Object) DataInputStream {
	return DataInputStream{
		BufferedInputStream: WrapBufferedInputStream(obj),
		Seekable:            WrapSeekable(obj),
	}
}

func marshalDataInputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDataInputStream(obj), nil
}

// NewDataInputStream constructs a class DataInputStream.
func NewDataInputStream(baseStream InputStream) DataInputStream {
	var arg1 *C.GInputStream

	arg1 = (*C.GInputStream)(unsafe.Pointer(baseStream.Native()))

	var cret C.GDataInputStream

	cret = C.g_data_input_stream_new(arg1)

	var dataInputStream DataInputStream

	dataInputStream = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(DataInputStream)

	return dataInputStream
}

// ByteOrder gets the byte order for the data input stream.
func (s dataInputStream) ByteOrder() DataStreamByteOrder {
	var arg0 *C.GDataInputStream

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))

	var cret C.GDataStreamByteOrder

	cret = C.g_data_input_stream_get_byte_order(arg0)

	var dataStreamByteOrder DataStreamByteOrder

	dataStreamByteOrder = DataStreamByteOrder(cret)

	return dataStreamByteOrder
}

// NewlineType gets the current newline type for the @stream.
func (s dataInputStream) NewlineType() DataStreamNewlineType {
	var arg0 *C.GDataInputStream

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))

	var cret C.GDataStreamNewlineType

	cret = C.g_data_input_stream_get_newline_type(arg0)

	var dataStreamNewlineType DataStreamNewlineType

	dataStreamNewlineType = DataStreamNewlineType(cret)

	return dataStreamNewlineType
}

// ReadByte reads an unsigned 8-bit/1-byte value from @stream.
func (s dataInputStream) ReadByte(cancellable Cancellable) (guint8 byte, goerr error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GCancellable

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cret C.guchar
	var cerr *C.GError

	cret = C.g_data_input_stream_read_byte(arg0, arg1, cerr)

	var guint8 byte
	var goerr error

	guint8 = (byte)(cret)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return guint8, goerr
}

// ReadInt16 reads a 16-bit/2-byte value from @stream.
//
// In order to get the correct byte order for this read operation, see
// g_data_input_stream_get_byte_order() and
// g_data_input_stream_set_byte_order().
func (s dataInputStream) ReadInt16(cancellable Cancellable) (gint16 int16, goerr error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GCancellable

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cret C.gint16
	var cerr *C.GError

	cret = C.g_data_input_stream_read_int16(arg0, arg1, cerr)

	var gint16 int16
	var goerr error

	gint16 = (int16)(cret)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return gint16, goerr
}

// ReadInt32 reads a signed 32-bit/4-byte value from @stream.
//
// In order to get the correct byte order for this read operation, see
// g_data_input_stream_get_byte_order() and
// g_data_input_stream_set_byte_order().
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
func (s dataInputStream) ReadInt32(cancellable Cancellable) (gint32 int32, goerr error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GCancellable

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cret C.gint32
	var cerr *C.GError

	cret = C.g_data_input_stream_read_int32(arg0, arg1, cerr)

	var gint32 int32
	var goerr error

	gint32 = (int32)(cret)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return gint32, goerr
}

// ReadInt64 reads a 64-bit/8-byte value from @stream.
//
// In order to get the correct byte order for this read operation, see
// g_data_input_stream_get_byte_order() and
// g_data_input_stream_set_byte_order().
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
func (s dataInputStream) ReadInt64(cancellable Cancellable) (gint64 int64, goerr error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GCancellable

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cret C.gint64
	var cerr *C.GError

	cret = C.g_data_input_stream_read_int64(arg0, arg1, cerr)

	var gint64 int64
	var goerr error

	gint64 = (int64)(cret)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return gint64, goerr
}

// ReadLine reads a line from the data input stream. Note that no encoding
// checks or conversion is performed; the input is not guaranteed to be
// UTF-8, and may in fact have embedded NUL characters.
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
func (s dataInputStream) ReadLine(cancellable Cancellable) (length uint, guint8s []byte, goerr error) {
	var arg0 *C.GDataInputStream
	var arg2 *C.GCancellable

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var arg1 C.gsize
	var cret *C.char
	var cerr *C.GError

	cret = C.g_data_input_stream_read_line(arg0, arg2, &arg1, cerr)

	var length uint
	var guint8s []byte
	var goerr error

	length = (uint)(arg1)
	{
		var length int
		for p := cret; *p != 0; p = (*C.char)(ptr.Add(unsafe.Pointer(p), C.sizeof_guint8)) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []C.guint8
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(cret), int(length))

		guint8s = make([]byte, length)
		for i := uintptr(0); i < uintptr(length); i += C.sizeof_guint8 {
			guint8s = (byte)(cret)
		}
	}
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return length, guint8s, goerr
}

// ReadLineAsync: the asynchronous version of
// g_data_input_stream_read_line(). It is an error to have two outstanding
// calls to this function.
//
// When the operation is finished, @callback will be called. You can then
// call g_data_input_stream_read_line_finish() to get the result of the
// operation.
func (s dataInputStream) ReadLineAsync() {
	var arg0 *C.GDataInputStream

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))

	C.g_data_input_stream_read_line_async(arg0)
}

// ReadLineFinish: finish an asynchronous call started by
// g_data_input_stream_read_line_async(). Note the warning about string
// encoding in g_data_input_stream_read_line() applies here as well.
func (s dataInputStream) ReadLineFinish(result AsyncResult) (length uint, guint8s []byte, goerr error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GAsyncResult

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var arg2 C.gsize
	var cret *C.char
	var cerr *C.GError

	cret = C.g_data_input_stream_read_line_finish(arg0, arg1, &arg2, cerr)

	var length uint
	var guint8s []byte
	var goerr error

	length = (uint)(arg2)
	{
		var length int
		for p := cret; *p != 0; p = (*C.char)(ptr.Add(unsafe.Pointer(p), C.sizeof_guint8)) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []C.guint8
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(cret), int(length))

		guint8s = make([]byte, length)
		for i := uintptr(0); i < uintptr(length); i += C.sizeof_guint8 {
			guint8s = (byte)(cret)
		}
	}
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return length, guint8s, goerr
}

// ReadLineFinishUTF8: finish an asynchronous call started by
// g_data_input_stream_read_line_async().
func (s dataInputStream) ReadLineFinishUTF8(result AsyncResult) (length uint, utf8 string, goerr error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GAsyncResult

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var arg2 C.gsize
	var cret *C.char
	var cerr *C.GError

	cret = C.g_data_input_stream_read_line_finish_utf8(arg0, arg1, &arg2, cerr)

	var length uint
	var utf8 string
	var goerr error

	length = (uint)(arg2)
	utf8 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return length, utf8, goerr
}

// ReadLineUTF8 reads a UTF-8 encoded line from the data input stream.
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
func (s dataInputStream) ReadLineUTF8(cancellable Cancellable) (length uint, utf8 string, goerr error) {
	var arg0 *C.GDataInputStream
	var arg2 *C.GCancellable

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var arg1 C.gsize
	var cret *C.char
	var cerr *C.GError

	cret = C.g_data_input_stream_read_line_utf8(arg0, arg2, &arg1, cerr)

	var length uint
	var utf8 string
	var goerr error

	length = (uint)(arg1)
	utf8 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return length, utf8, goerr
}

// ReadUint16 reads an unsigned 16-bit/2-byte value from @stream.
//
// In order to get the correct byte order for this read operation, see
// g_data_input_stream_get_byte_order() and
// g_data_input_stream_set_byte_order().
func (s dataInputStream) ReadUint16(cancellable Cancellable) (guint16 uint16, goerr error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GCancellable

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cret C.guint16
	var cerr *C.GError

	cret = C.g_data_input_stream_read_uint16(arg0, arg1, cerr)

	var guint16 uint16
	var goerr error

	guint16 = (uint16)(cret)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return guint16, goerr
}

// ReadUint32 reads an unsigned 32-bit/4-byte value from @stream.
//
// In order to get the correct byte order for this read operation, see
// g_data_input_stream_get_byte_order() and
// g_data_input_stream_set_byte_order().
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
func (s dataInputStream) ReadUint32(cancellable Cancellable) (guint32 uint32, goerr error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GCancellable

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cret C.guint32
	var cerr *C.GError

	cret = C.g_data_input_stream_read_uint32(arg0, arg1, cerr)

	var guint32 uint32
	var goerr error

	guint32 = (uint32)(cret)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return guint32, goerr
}

// ReadUint64 reads an unsigned 64-bit/8-byte value from @stream.
//
// In order to get the correct byte order for this read operation, see
// g_data_input_stream_get_byte_order().
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned.
func (s dataInputStream) ReadUint64(cancellable Cancellable) (guint64 uint64, goerr error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GCancellable

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cret C.guint64
	var cerr *C.GError

	cret = C.g_data_input_stream_read_uint64(arg0, arg1, cerr)

	var guint64 uint64
	var goerr error

	guint64 = (uint64)(cret)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return guint64, goerr
}

// ReadUntil reads a string from the data input stream, up to the first
// occurrence of any of the stop characters.
//
// Note that, in contrast to g_data_input_stream_read_until_async(), this
// function consumes the stop character that it finds.
//
// Don't use this function in new code. Its functionality is inconsistent
// with g_data_input_stream_read_until_async(). Both functions will be
// marked as deprecated in a future release. Use
// g_data_input_stream_read_upto() instead, but note that that function does
// not consume the stop character.
func (s dataInputStream) ReadUntil(stopChars string, cancellable Cancellable) (length uint, utf8 string, goerr error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.gchar
	var arg3 *C.GCancellable

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(stopChars))
	defer C.free(unsafe.Pointer(arg1))
	arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var arg2 C.gsize
	var cret *C.char
	var cerr *C.GError

	cret = C.g_data_input_stream_read_until(arg0, arg1, arg3, &arg2, cerr)

	var length uint
	var utf8 string
	var goerr error

	length = (uint)(arg2)
	utf8 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return length, utf8, goerr
}

// ReadUntilAsync: the asynchronous version of
// g_data_input_stream_read_until(). It is an error to have two outstanding
// calls to this function.
//
// Note that, in contrast to g_data_input_stream_read_until(), this function
// does not consume the stop character that it finds. You must read it for
// yourself.
//
// When the operation is finished, @callback will be called. You can then
// call g_data_input_stream_read_until_finish() to get the result of the
// operation.
//
// Don't use this function in new code. Its functionality is inconsistent
// with g_data_input_stream_read_until(). Both functions will be marked as
// deprecated in a future release. Use g_data_input_stream_read_upto_async()
// instead.
func (s dataInputStream) ReadUntilAsync() {
	var arg0 *C.GDataInputStream

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))

	C.g_data_input_stream_read_until_async(arg0)
}

// ReadUntilFinish: finish an asynchronous call started by
// g_data_input_stream_read_until_async().
func (s dataInputStream) ReadUntilFinish(result AsyncResult) (length uint, utf8 string, goerr error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GAsyncResult

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var arg2 C.gsize
	var cret *C.char
	var cerr *C.GError

	cret = C.g_data_input_stream_read_until_finish(arg0, arg1, &arg2, cerr)

	var length uint
	var utf8 string
	var goerr error

	length = (uint)(arg2)
	utf8 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return length, utf8, goerr
}

// ReadUpto reads a string from the data input stream, up to the first
// occurrence of any of the stop characters.
//
// In contrast to g_data_input_stream_read_until(), this function does not
// consume the stop character. You have to use
// g_data_input_stream_read_byte() to get it before calling
// g_data_input_stream_read_upto() again.
//
// Note that @stop_chars may contain '\0' if @stop_chars_len is specified.
//
// The returned string will always be nul-terminated on success.
func (s dataInputStream) ReadUpto(stopChars string, stopCharsLen int, cancellable Cancellable) (length uint, utf8 string, goerr error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.gchar
	var arg2 C.gssize
	var arg4 *C.GCancellable

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(stopChars))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gssize(stopCharsLen)
	arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var arg3 C.gsize
	var cret *C.char
	var cerr *C.GError

	cret = C.g_data_input_stream_read_upto(arg0, arg1, arg2, arg4, &arg3, cerr)

	var length uint
	var utf8 string
	var goerr error

	length = (uint)(arg3)
	utf8 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return length, utf8, goerr
}

// ReadUptoAsync: the asynchronous version of
// g_data_input_stream_read_upto(). It is an error to have two outstanding
// calls to this function.
//
// In contrast to g_data_input_stream_read_until(), this function does not
// consume the stop character. You have to use
// g_data_input_stream_read_byte() to get it before calling
// g_data_input_stream_read_upto() again.
//
// Note that @stop_chars may contain '\0' if @stop_chars_len is specified.
//
// When the operation is finished, @callback will be called. You can then
// call g_data_input_stream_read_upto_finish() to get the result of the
// operation.
func (s dataInputStream) ReadUptoAsync() {
	var arg0 *C.GDataInputStream

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))

	C.g_data_input_stream_read_upto_async(arg0)
}

// ReadUptoFinish: finish an asynchronous call started by
// g_data_input_stream_read_upto_async().
//
// Note that this function does not consume the stop character. You have to
// use g_data_input_stream_read_byte() to get it before calling
// g_data_input_stream_read_upto_async() again.
//
// The returned string will always be nul-terminated on success.
func (s dataInputStream) ReadUptoFinish(result AsyncResult) (length uint, utf8 string, goerr error) {
	var arg0 *C.GDataInputStream
	var arg1 *C.GAsyncResult

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var arg2 C.gsize
	var cret *C.char
	var cerr *C.GError

	cret = C.g_data_input_stream_read_upto_finish(arg0, arg1, &arg2, cerr)

	var length uint
	var utf8 string
	var goerr error

	length = (uint)(arg2)
	utf8 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return length, utf8, goerr
}

// SetByteOrder: this function sets the byte order for the given @stream.
// All subsequent reads from the @stream will be read in the given @order.
func (s dataInputStream) SetByteOrder(order DataStreamByteOrder) {
	var arg0 *C.GDataInputStream
	var arg1 C.GDataStreamByteOrder

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (C.GDataStreamByteOrder)(order)

	C.g_data_input_stream_set_byte_order(arg0, arg1)
}

// SetNewlineType sets the newline type for the @stream.
//
// Note that using G_DATA_STREAM_NEWLINE_TYPE_ANY is slightly unsafe. If a
// read chunk ends in "CR" we must read an additional byte to know if this
// is "CR" or "CR LF", and this might block if there is no more data
// available.
func (s dataInputStream) SetNewlineType(typ DataStreamNewlineType) {
	var arg0 *C.GDataInputStream
	var arg1 C.GDataStreamNewlineType

	arg0 = (*C.GDataInputStream)(unsafe.Pointer(s.Native()))
	arg1 = (C.GDataStreamNewlineType)(typ)

	C.g_data_input_stream_set_newline_type(arg0, arg1)
}
