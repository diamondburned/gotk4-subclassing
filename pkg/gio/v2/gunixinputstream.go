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
		{T: externglib.Type(C.g_unix_input_stream_get_type()), F: marshalUnixInputStream},
	})
}

// UnixInputStream implements Stream for reading from a UNIX file descriptor,
// including asynchronous operations. (If the file descriptor refers to a socket
// or pipe, this will use poll() to do asynchronous I/O. If it refers to a
// regular file, it will fall back to doing asynchronous I/O in another thread.)
//
// Note that `<gio/gunixinputstream.h>` belongs to the UNIX-specific GIO
// interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config file when
// using it.
type UnixInputStream interface {
	InputStream
	FileDescriptorBased
	PollableInputStream

	// CloseFd returns whether the file descriptor of @stream will be closed
	// when the stream is closed.
	CloseFd(s UnixInputStream) bool
	// Fd: return the UNIX file descriptor that the stream reads from.
	Fd(s UnixInputStream)
	// SetCloseFd sets whether the file descriptor of @stream shall be closed
	// when the stream is closed.
	SetCloseFd(s UnixInputStream, closeFd bool)
}

// unixInputStream implements the UnixInputStream interface.
type unixInputStream struct {
	InputStream
	FileDescriptorBased
	PollableInputStream
}

var _ UnixInputStream = (*unixInputStream)(nil)

// WrapUnixInputStream wraps a GObject to the right type. It is
// primarily used internally.
func WrapUnixInputStream(obj *externglib.Object) UnixInputStream {
	return UnixInputStream{
		InputStream:         WrapInputStream(obj),
		FileDescriptorBased: WrapFileDescriptorBased(obj),
		PollableInputStream: WrapPollableInputStream(obj),
	}
}

func marshalUnixInputStream(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapUnixInputStream(obj), nil
}

// NewUnixInputStream constructs a class UnixInputStream.
func NewUnixInputStream(fd int, closeFd bool) {
	var arg1 C.gint
	var arg2 C.gboolean

	arg1 = C.gint(fd)
	if closeFd {
		arg2 = C.gboolean(1)
	}

	C.g_unix_input_stream_new(arg1, arg2)
}

// CloseFd returns whether the file descriptor of @stream will be closed
// when the stream is closed.
func (s unixInputStream) CloseFd(s UnixInputStream) bool {
	var arg0 *C.GUnixInputStream

	arg0 = (*C.GUnixInputStream)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_unix_input_stream_get_close_fd(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Fd: return the UNIX file descriptor that the stream reads from.
func (s unixInputStream) Fd(s UnixInputStream) {
	var arg0 *C.GUnixInputStream

	arg0 = (*C.GUnixInputStream)(unsafe.Pointer(s.Native()))

	C.g_unix_input_stream_get_fd(arg0)
}

// SetCloseFd sets whether the file descriptor of @stream shall be closed
// when the stream is closed.
func (s unixInputStream) SetCloseFd(s UnixInputStream, closeFd bool) {
	var arg0 *C.GUnixInputStream
	var arg1 C.gboolean

	arg0 = (*C.GUnixInputStream)(unsafe.Pointer(s.Native()))
	if closeFd {
		arg1 = C.gboolean(1)
	}

	C.g_unix_input_stream_set_close_fd(arg0, arg1)
}
