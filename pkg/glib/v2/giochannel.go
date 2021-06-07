// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_io_channel_get_type()), F: marshalIOChannel},
	})
}

// IOChannelError: error codes returned by OChannel operations.
type IOChannelError int

const (
	// IOChannelErrorFbig: file too large.
	IOChannelErrorFbig IOChannelError = 0
	// IOChannelErrorInval: invalid argument.
	IOChannelErrorInval IOChannelError = 1
	// IOChannelErrorIO: IO error.
	IOChannelErrorIO IOChannelError = 2
	// IOChannelErrorIsdir: file is a directory.
	IOChannelErrorIsdir IOChannelError = 3
	// IOChannelErrorNospc: no space left on device.
	IOChannelErrorNospc IOChannelError = 4
	// IOChannelErrorNxio: no such device or address.
	IOChannelErrorNxio IOChannelError = 5
	// IOChannelErrorOverflow: value too large for defined datatype.
	IOChannelErrorOverflow IOChannelError = 6
	// IOChannelErrorPipe: broken pipe.
	IOChannelErrorPipe IOChannelError = 7
	// IOChannelErrorFailed: some other error.
	IOChannelErrorFailed IOChannelError = 8
)

// IOError is only used by the deprecated functions g_io_channel_read(),
// g_io_channel_write(), and g_io_channel_seek().
type IOError int

const (
	// IOErrorNone: no error
	IOErrorNone IOError = 0
	// IOErrorAgain: an EAGAIN error occurred
	IOErrorAgain IOError = 1
	// IOErrorInval: an EINVAL error occurred
	IOErrorInval IOError = 2
	// IOErrorUnknown: another error occurred
	IOErrorUnknown IOError = 3
)

// IOStatus statuses returned by most of the OFuncs functions.
type IOStatus int

const (
	// IOStatusError: an error occurred.
	IOStatusError IOStatus = 0
	// IOStatusNormal: success.
	IOStatusNormal IOStatus = 1
	// IOStatusEOF: end of file.
	IOStatusEOF IOStatus = 2
	// IOStatusAgain: resource temporarily unavailable.
	IOStatusAgain IOStatus = 3
)

// SeekType: an enumeration specifying the base position for a
// g_io_channel_seek_position() operation.
type SeekType int

const (
	// SeekTypeCur: the current position in the file.
	SeekTypeCur SeekType = 0
	// SeekTypeSet: the start of the file.
	SeekTypeSet SeekType = 1
	// SeekTypeEnd: the end of the file.
	SeekTypeEnd SeekType = 2
)

// IOFlags specifies properties of a OChannel. Some of the flags can only be
// read with g_io_channel_get_flags(), but not changed with
// g_io_channel_set_flags().
type IOFlags int

const (
	// IOFlagsAppend turns on append mode, corresponds to O_APPEND (see the
	// documentation of the UNIX open() syscall)
	IOFlagsAppend IOFlags = 1
	// IOFlagsNonblock turns on nonblocking mode, corresponds to
	// O_NONBLOCK/O_NDELAY (see the documentation of the UNIX open() syscall)
	IOFlagsNonblock IOFlags = 2
	// IOFlagsIsReadable indicates that the io channel is readable. This flag
	// cannot be changed.
	IOFlagsIsReadable IOFlags = 4
	// IOFlagsIsWritable indicates that the io channel is writable. This flag
	// cannot be changed.
	IOFlagsIsWritable IOFlags = 8
	// IOFlagsIsWriteable: a misspelled version of @G_IO_FLAG_IS_WRITABLE that
	// existed before the spelling was fixed in GLib 2.30. It is kept here for
	// compatibility reasons. Deprecated since 2.30
	IOFlagsIsWriteable IOFlags = 8
	// IOFlagsIsSeekable indicates that the io channel is seekable, i.e. that
	// g_io_channel_seek_position() can be used on it. This flag cannot be
	// changed.
	IOFlagsIsSeekable IOFlags = 16
	// IOFlagsMask: the mask that specifies all the valid flags.
	IOFlagsMask IOFlags = 31
	// IOFlagsGetMask: the mask of the flags that are returned from
	// g_io_channel_get_flags()
	IOFlagsGetMask IOFlags = 31
	// IOFlagsSetMask: the mask of the flags that the user can modify with
	// g_io_channel_set_flags()
	IOFlagsSetMask IOFlags = 3
)

// IOAddWatchFull adds the OChannel into the default main loop context with the
// given priority.
//
// This internally creates a main loop source using g_io_create_watch() and
// attaches it to the main loop context with g_source_attach(). You can do these
// steps manually if you need greater control.
func IOAddWatchFull() {
	C.g_io_add_watch_full(arg1, arg2, arg3, arg4, arg5, arg6)
}

// IOChannelErrorFromErrno converts an `errno` error number to a OChannelError.
func IOChannelErrorFromErrno(en int) {
	var arg1 C.gint

	arg1 = C.gint(en)

	C.g_io_channel_error_from_errno(arg1)
}

// IOCreateWatch creates a #GSource that's dispatched when @condition is met for
// the given @channel. For example, if condition is IO_IN, the source will be
// dispatched when there's data available for reading.
//
// The callback function invoked by the #GSource should be added with
// g_source_set_callback(), but it has type OFunc (not Func).
//
// g_io_add_watch() is a simpler interface to this same functionality, for the
// case where you want to add the source to the default main loop context at the
// default priority.
//
// On Windows, polling a #GSource created to watch a channel for a socket puts
// the socket in non-blocking mode. This is a side-effect of the implementation
// and unavoidable.
func IOCreateWatch(channel *IOChannel, condition IOCondition) {
	var arg1 *C.GIOChannel
	var arg2 C.GIOCondition

	arg1 = (*C.GIOChannel)(unsafe.Pointer(channel.Native()))
	arg2 = (C.GIOCondition)(condition)

	C.g_io_create_watch(arg1, arg2)
}

// IOChannel: a data structure representing an IO Channel. The fields should be
// considered private and should only be accessed with the following functions.
type IOChannel struct {
	native C.GIOChannel
}

// WrapIOChannel wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapIOChannel(ptr unsafe.Pointer) *IOChannel {
	if ptr == nil {
		return nil
	}

	return (*IOChannel)(ptr)
}

func marshalIOChannel(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapIOChannel(unsafe.Pointer(b)), nil
}

// NewIOChannelFile constructs a struct IOChannel.
func NewIOChannelFile(filename string, mode string) error {
	var arg1 *C.gchar
	var arg2 *C.gchar

	arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(mode))
	defer C.free(unsafe.Pointer(arg2))

	var errout *C.GError
	var err error

	C.g_io_channel_new_file(arg1, arg2, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// NewIOChannelUnix constructs a struct IOChannel.
func NewIOChannelUnix(fd int) {
	var arg1 C.int

	arg1 = C.int(fd)

	C.g_io_channel_unix_new(arg1)
}

// Native returns the underlying C source pointer.
func (i *IOChannel) Native() unsafe.Pointer {
	return unsafe.Pointer(&i.native)
}

// Close: close an IO channel. Any pending data to be written will be flushed,
// ignoring errors. The channel will not be freed until the last reference is
// dropped using g_io_channel_unref().
func (c *IOChannel) Close(c *IOChannel) {
	var arg0 *C.GIOChannel

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))

	C.g_io_channel_close(arg0)
}

// Flush flushes the write buffer for the GIOChannel.
func (c *IOChannel) Flush(c *IOChannel) error {
	var arg0 *C.GIOChannel

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))

	var errout *C.GError
	var err error

	C.g_io_channel_flush(arg0, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// BufferCondition: this function returns a OCondition depending on whether
// there is data to be read/space to write data in the internal buffers in the
// OChannel. Only the flags G_IO_IN and G_IO_OUT may be set.
func (c *IOChannel) BufferCondition(c *IOChannel) {
	var arg0 *C.GIOChannel

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))

	C.g_io_channel_get_buffer_condition(arg0)
}

// BufferSize gets the buffer size.
func (c *IOChannel) BufferSize(c *IOChannel) {
	var arg0 *C.GIOChannel

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))

	C.g_io_channel_get_buffer_size(arg0)
}

// Buffered returns whether @channel is buffered.
func (c *IOChannel) Buffered(c *IOChannel) bool {
	var arg0 *C.GIOChannel

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_io_channel_get_buffered(arg0)

	if cret {
		ok = true
	}

	return ok
}

// CloseOnUnref returns whether the file/socket/whatever associated with
// @channel will be closed when @channel receives its final unref and is
// destroyed. The default value of this is true for channels created by
// g_io_channel_new_file (), and false for all other channels.
func (c *IOChannel) CloseOnUnref(c *IOChannel) bool {
	var arg0 *C.GIOChannel

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_io_channel_get_close_on_unref(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Encoding gets the encoding for the input/output of the channel. The internal
// encoding is always UTF-8. The encoding nil makes the channel safe for binary
// data.
func (c *IOChannel) Encoding(c *IOChannel) {
	var arg0 *C.GIOChannel

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))

	C.g_io_channel_get_encoding(arg0)
}

// Flags gets the current flags for a OChannel, including read-only flags such
// as G_IO_FLAG_IS_READABLE.
//
// The values of the flags G_IO_FLAG_IS_READABLE and G_IO_FLAG_IS_WRITABLE are
// cached for internal use by the channel when it is created. If they should
// change at some later point (e.g. partial shutdown of a socket with the UNIX
// shutdown() function), the user should immediately call
// g_io_channel_get_flags() to update the internal values of these flags.
func (c *IOChannel) Flags(c *IOChannel) {
	var arg0 *C.GIOChannel

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))

	C.g_io_channel_get_flags(arg0)
}

// LineTerm: this returns the string that OChannel uses to determine where in
// the file a line break occurs. A value of nil indicates autodetection.
func (c *IOChannel) LineTerm(c *IOChannel, length int) {
	var arg0 *C.GIOChannel
	var arg1 *C.gint

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))
	arg1 = *C.gint(length)

	C.g_io_channel_get_line_term(arg0, arg1)
}

// Init initializes a OChannel struct.
//
// This is called by each of the above functions when creating a OChannel, and
// so is not often needed by the application programmer (unless you are creating
// a new type of OChannel).
func (c *IOChannel) Init(c *IOChannel) {
	var arg0 *C.GIOChannel

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))

	C.g_io_channel_init(arg0)
}

// Read reads data from a OChannel.
func (c *IOChannel) Read(c *IOChannel, buf string, count uint, bytesRead uint) {
	var arg0 *C.GIOChannel
	var arg1 *C.gchar
	var arg2 C.gsize
	var arg3 *C.gsize

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))
	arg1 = (*C.gchar)(C.CString(buf))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gsize(count)
	arg3 = *C.gsize(bytesRead)

	C.g_io_channel_read(arg0, arg1, arg2, arg3)
}

// ReadLine reads a line, including the terminating character(s), from a
// OChannel into a newly-allocated string. @str_return will contain allocated
// memory if the return is G_IO_STATUS_NORMAL.
func (c *IOChannel) ReadLine(c *IOChannel) (strReturn string, length uint, terminatorPos uint, err error) {
	var arg0 *C.GIOChannel

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))

	var arg1 *C.gchar
	var strReturn string
	var arg2 C.gsize
	var length uint
	var arg3 C.gsize
	var terminatorPos uint
	var errout *C.GError
	var err error

	C.g_io_channel_read_line(arg0, &arg1, &arg2, &arg3, &errout)

	strReturn = C.GoString(&arg1)
	defer C.free(unsafe.Pointer(&arg1))
	length = uint(&arg2)
	terminatorPos = uint(&arg3)
	err = gerror.Take(unsafe.Pointer(errout))

	return strReturn, length, terminatorPos, err
}

// ReadLineString reads a line from a OChannel, using a #GString as a buffer.
func (c *IOChannel) ReadLineString(c *IOChannel, buffer *String, terminatorPos uint) error {
	var arg0 *C.GIOChannel
	var arg1 *C.GString
	var arg2 *C.gsize

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GString)(unsafe.Pointer(buffer.Native()))
	arg2 = *C.gsize(terminatorPos)

	var errout *C.GError
	var err error

	C.g_io_channel_read_line_string(arg0, arg1, arg2, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// ReadToEnd reads all the remaining data from the file.
func (c *IOChannel) ReadToEnd(c *IOChannel) error {
	var arg0 *C.GIOChannel

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))

	var errout *C.GError
	var err error

	C.g_io_channel_read_to_end(arg0, &arg1, &arg2, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return strReturn, length, err
}

// ReadUnichar reads a Unicode character from @channel. This function cannot be
// called on a channel with nil encoding.
func (c *IOChannel) ReadUnichar(c *IOChannel) (thechar uint32, err error) {
	var arg0 *C.GIOChannel

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))

	var arg1 C.gunichar
	var thechar uint32
	var errout *C.GError
	var err error

	C.g_io_channel_read_unichar(arg0, &arg1, &errout)

	thechar = uint32(&arg1)
	err = gerror.Take(unsafe.Pointer(errout))

	return thechar, err
}

// Ref increments the reference count of a OChannel.
func (c *IOChannel) Ref(c *IOChannel) {
	var arg0 *C.GIOChannel

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))

	C.g_io_channel_ref(arg0)
}

// Seek sets the current position in the OChannel, similar to the standard
// library function fseek().
func (c *IOChannel) Seek(c *IOChannel, offset int64, typ SeekType) {
	var arg0 *C.GIOChannel
	var arg1 C.gint64
	var arg2 C.GSeekType

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))
	arg1 = C.gint64(offset)
	arg2 = (C.GSeekType)(typ)

	C.g_io_channel_seek(arg0, arg1, arg2)
}

// SeekPosition: replacement for g_io_channel_seek() with the new API.
func (c *IOChannel) SeekPosition(c *IOChannel, offset int64, typ SeekType) error {
	var arg0 *C.GIOChannel
	var arg1 C.gint64
	var arg2 C.GSeekType

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))
	arg1 = C.gint64(offset)
	arg2 = (C.GSeekType)(typ)

	var errout *C.GError
	var err error

	C.g_io_channel_seek_position(arg0, arg1, arg2, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// SetBufferSize sets the buffer size.
func (c *IOChannel) SetBufferSize(c *IOChannel, size uint) {
	var arg0 *C.GIOChannel
	var arg1 C.gsize

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))
	arg1 = C.gsize(size)

	C.g_io_channel_set_buffer_size(arg0, arg1)
}

// SetBuffered: the buffering state can only be set if the channel's encoding is
// nil. For any other encoding, the channel must be buffered.
//
// A buffered channel can only be set unbuffered if the channel's internal
// buffers have been flushed. Newly created channels or channels which have
// returned G_IO_STATUS_EOF not require such a flush. For write-only channels, a
// call to g_io_channel_flush () is sufficient. For all other channels, the
// buffers may be flushed by a call to g_io_channel_seek_position (). This
// includes the possibility of seeking with seek type G_SEEK_CUR and an offset
// of zero. Note that this means that socket-based channels cannot be set
// unbuffered once they have had data read from them.
//
// On unbuffered channels, it is safe to mix read and write calls from the new
// and old APIs, if this is necessary for maintaining old code.
//
// The default state of the channel is buffered.
func (c *IOChannel) SetBuffered(c *IOChannel, buffered bool) {
	var arg0 *C.GIOChannel
	var arg1 C.gboolean

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))
	if buffered {
		arg1 = C.gboolean(1)
	}

	C.g_io_channel_set_buffered(arg0, arg1)
}

// SetCloseOnUnref: whether to close the channel on the final unref of the
// OChannel data structure. The default value of this is true for channels
// created by g_io_channel_new_file (), and false for all other channels.
//
// Setting this flag to true for a channel you have already closed can cause
// problems when the final reference to the OChannel is dropped.
func (c *IOChannel) SetCloseOnUnref(c *IOChannel, doClose bool) {
	var arg0 *C.GIOChannel
	var arg1 C.gboolean

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))
	if doClose {
		arg1 = C.gboolean(1)
	}

	C.g_io_channel_set_close_on_unref(arg0, arg1)
}

// SetEncoding sets the encoding for the input/output of the channel. The
// internal encoding is always UTF-8. The default encoding for the external file
// is UTF-8.
//
// The encoding nil is safe to use with binary data.
//
// The encoding can only be set if one of the following conditions is true:
//
// - The channel was just created, and has not been written to or read from yet.
//
// - The channel is write-only.
//
// - The channel is a file, and the file pointer was just repositioned by a call
// to g_io_channel_seek_position(). (This flushes all the internal buffers.)
//
// - The current encoding is nil or UTF-8.
//
// - One of the (new API) read functions has just returned G_IO_STATUS_EOF (or,
// in the case of g_io_channel_read_to_end(), G_IO_STATUS_NORMAL).
//
// - One of the functions g_io_channel_read_chars() or
// g_io_channel_read_unichar() has returned G_IO_STATUS_AGAIN or
// G_IO_STATUS_ERROR. This may be useful in the case of
// G_CONVERT_ERROR_ILLEGAL_SEQUENCE. Returning one of these statuses from
// g_io_channel_read_line(), g_io_channel_read_line_string(), or
// g_io_channel_read_to_end() does not guarantee that the encoding can be
// changed.
//
// Channels which do not meet one of the above conditions cannot call
// g_io_channel_seek_position() with an offset of G_SEEK_CUR, and, if they are
// "seekable", cannot call g_io_channel_write_chars() after calling one of the
// API "read" functions.
func (c *IOChannel) SetEncoding(c *IOChannel, encoding string) error {
	var arg0 *C.GIOChannel
	var arg1 *C.gchar

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))
	arg1 = (*C.gchar)(C.CString(encoding))
	defer C.free(unsafe.Pointer(arg1))

	var errout *C.GError
	var err error

	C.g_io_channel_set_encoding(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// SetFlags sets the (writeable) flags in @channel to (@flags &
// G_IO_FLAG_SET_MASK).
func (c *IOChannel) SetFlags(c *IOChannel, flags IOFlags) error {
	var arg0 *C.GIOChannel
	var arg1 C.GIOFlags

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))
	arg1 = (C.GIOFlags)(flags)

	var errout *C.GError
	var err error

	C.g_io_channel_set_flags(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// SetLineTerm: this sets the string that OChannel uses to determine where in
// the file a line break occurs.
func (c *IOChannel) SetLineTerm(c *IOChannel, lineTerm string, length int) {
	var arg0 *C.GIOChannel
	var arg1 *C.gchar
	var arg2 C.gint

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))
	arg1 = (*C.gchar)(C.CString(lineTerm))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gint(length)

	C.g_io_channel_set_line_term(arg0, arg1, arg2)
}

// Shutdown: close an IO channel. Any pending data to be written will be flushed
// if @flush is true. The channel will not be freed until the last reference is
// dropped using g_io_channel_unref().
func (c *IOChannel) Shutdown(c *IOChannel, flush bool) error {
	var arg0 *C.GIOChannel
	var arg1 C.gboolean

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))
	if flush {
		arg1 = C.gboolean(1)
	}

	var errout *C.GError
	var err error

	C.g_io_channel_shutdown(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// UnixGetFd returns the file descriptor of the OChannel.
//
// On Windows this function returns the file descriptor or socket of the
// OChannel.
func (c *IOChannel) UnixGetFd(c *IOChannel) {
	var arg0 *C.GIOChannel

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))

	C.g_io_channel_unix_get_fd(arg0)
}

// Unref decrements the reference count of a OChannel.
func (c *IOChannel) Unref(c *IOChannel) {
	var arg0 *C.GIOChannel

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))

	C.g_io_channel_unref(arg0)
}

// Write writes data to a OChannel.
func (c *IOChannel) Write(c *IOChannel, buf string, count uint, bytesWritten uint) {
	var arg0 *C.GIOChannel
	var arg1 *C.gchar
	var arg2 C.gsize
	var arg3 *C.gsize

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))
	arg1 = (*C.gchar)(C.CString(buf))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gsize(count)
	arg3 = *C.gsize(bytesWritten)

	C.g_io_channel_write(arg0, arg1, arg2, arg3)
}

// WriteUnichar writes a Unicode character to @channel. This function cannot be
// called on a channel with nil encoding.
func (c *IOChannel) WriteUnichar(c *IOChannel, thechar uint32) error {
	var arg0 *C.GIOChannel
	var arg1 C.gunichar

	arg0 = (*C.GIOChannel)(unsafe.Pointer(c.Native()))
	arg1 = C.gunichar(thechar)

	var errout *C.GError
	var err error

	C.g_io_channel_write_unichar(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}