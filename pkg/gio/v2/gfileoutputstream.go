// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0 glib-2.0
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
	Etag() string
	// QueryInfoAsync: asynchronously queries the @stream for a Info. When
	// completed, @callback will be called with a Result which can be used to
	// finish the operation with g_file_output_stream_query_info_finish().
	//
	// For the synchronous version of this function, see
	// g_file_output_stream_query_info().
	QueryInfoAsync(attributes string, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback)
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
func (s fileOutputStream) Etag() string {
	var _arg0 *C.GFileOutputStream // out

	_arg0 = (*C.GFileOutputStream)(unsafe.Pointer(s.Native()))

	var _cret *C.char // in

	_cret = C.g_file_output_stream_get_etag(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// QueryInfoAsync: asynchronously queries the @stream for a Info. When
// completed, @callback will be called with a Result which can be used to
// finish the operation with g_file_output_stream_query_info_finish().
//
// For the synchronous version of this function, see
// g_file_output_stream_query_info().
func (s fileOutputStream) QueryInfoAsync(attributes string, ioPriority int, cancellable Cancellable, callback AsyncReadyCallback) {
	var _arg0 *C.GFileOutputStream  // out
	var _arg1 *C.char               // out
	var _arg2 C.int                 // out
	var _arg3 *C.GCancellable       // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GFileOutputStream)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(ioPriority)
	_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	_arg4 = (*[0]byte)(C.gotk4_AsyncReadyCallback)
	_arg5 = C.gpointer(box.Assign(callback))

	C.g_file_output_stream_query_info_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
}
