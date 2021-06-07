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
		{T: externglib.Type(C.g_file_enumerator_get_type()), F: marshalFileEnumerator},
	})
}

// FileEnumerator allows you to operate on a set of #GFiles, returning a Info
// structure for each file enumerated (e.g. g_file_enumerate_children() will
// return a Enumerator for each of the children within a directory).
//
// To get the next file's information from a Enumerator, use
// g_file_enumerator_next_file() or its asynchronous version,
// g_file_enumerator_next_files_async(). Note that the asynchronous version will
// return a list of Infos, whereas the synchronous will only return the next
// file in the enumerator.
//
// The ordering of returned files is unspecified for non-Unix platforms; for
// more information, see g_dir_read_name(). On Unix, when operating on local
// files, returned files will be sorted by inode number. Effectively you can
// assume that the ordering of returned files will be stable between successive
// calls (and applications) assuming the directory is unchanged.
//
// If your application needs a specific ordering, such as by name or
// modification time, you will have to implement that in your application code.
//
// To close a Enumerator, use g_file_enumerator_close(), or its asynchronous
// version, g_file_enumerator_close_async(). Once a Enumerator is closed, no
// further actions may be performed on it, and it should be freed with
// g_object_unref().
type FileEnumerator interface {
	gextras.Objector

	// Close releases all resources used by this enumerator, making the
	// enumerator return G_IO_ERROR_CLOSED on all calls.
	//
	// This will be automatically called when the last reference is dropped, but
	// you might want to call this function to make sure resources are released
	// as early as possible.
	Close(e FileEnumerator, cancellable Cancellable) error
	// CloseAsync: asynchronously closes the file enumerator.
	//
	// If @cancellable is not nil, then the operation can be cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be returned in
	// g_file_enumerator_close_finish().
	CloseAsync(e FileEnumerator)
	// CloseFinish finishes closing a file enumerator, started from
	// g_file_enumerator_close_async().
	//
	// If the file enumerator was already closed when
	// g_file_enumerator_close_async() was called, then this function will
	// report G_IO_ERROR_CLOSED in @error, and return false. If the file
	// enumerator had pending operation when the close operation was started,
	// then this function will report G_IO_ERROR_PENDING, and return false. If
	// @cancellable was not nil, then the operation may have been cancelled by
	// triggering the cancellable object from another thread. If the operation
	// was cancelled, the error G_IO_ERROR_CANCELLED will be set, and false will
	// be returned.
	CloseFinish(e FileEnumerator, result AsyncResult) error
	// Child: return a new #GFile which refers to the file named by @info in the
	// source directory of @enumerator. This function is primarily intended to
	// be used inside loops with g_file_enumerator_next_file().
	//
	// This is a convenience method that's equivalent to:
	//
	//    gchar *name = g_file_info_get_name (info);
	//    GFile *child = g_file_get_child (g_file_enumerator_get_container (enumr),
	//                                     name);
	Child(e FileEnumerator, info FileInfo)
	// Container: get the #GFile container which is being enumerated.
	Container(e FileEnumerator)
	// HasPending checks if the file enumerator has pending operations.
	HasPending(e FileEnumerator) bool
	// IsClosed checks if the file enumerator has been closed.
	IsClosed(e FileEnumerator) bool
	// Iterate: this is a version of g_file_enumerator_next_file() that's easier
	// to use correctly from C programs. With g_file_enumerator_next_file(), the
	// gboolean return value signifies "end of iteration or error", which
	// requires allocation of a temporary #GError.
	//
	// In contrast, with this function, a false return from
	// g_file_enumerator_iterate() *always* means "error". End of iteration is
	// signaled by @out_info or @out_child being nil.
	//
	// Another crucial difference is that the references for @out_info and
	// @out_child are owned by @direnum (they are cached as hidden properties).
	// You must not unref them in your own code. This makes memory management
	// significantly easier for C code in combination with loops.
	//
	// Finally, this function optionally allows retrieving a #GFile as well.
	//
	// You must specify at least one of @out_info or @out_child.
	//
	// The code pattern for correctly using g_file_enumerator_iterate() from C
	// is:
	//
	//    direnum = g_file_enumerate_children (file, ...);
	//    while (TRUE)
	//      {
	//        GFileInfo *info;
	//        if (!g_file_enumerator_iterate (direnum, &info, NULL, cancellable, error))
	//          goto out;
	//        if (!info)
	//          break;
	//        ... do stuff with "info"; do not unref it! ...
	//      }
	//
	//    out:
	//      g_object_unref (direnum); // Note: frees the last @info
	Iterate(d FileEnumerator, cancellable Cancellable) (outInfo FileInfo, outChild *File, err error)
	// NextFile returns information for the next file in the enumerated object.
	// Will block until the information is available. The Info returned from
	// this function will contain attributes that match the attribute string
	// that was passed when the Enumerator was created.
	//
	// See the documentation of Enumerator for information about the order of
	// returned files.
	//
	// On error, returns nil and sets @error to the error. If the enumerator is
	// at the end, nil will be returned and @error will be unset.
	NextFile(e FileEnumerator, cancellable Cancellable) error
	// NextFilesAsync: request information for a number of files from the
	// enumerator asynchronously. When all i/o for the operation is finished the
	// @callback will be called with the requested information.
	//
	// See the documentation of Enumerator for information about the order of
	// returned files.
	//
	// The callback can be called with less than @num_files files in case of
	// error or at the end of the enumerator. In case of a partial error the
	// callback will be called with any succeeding items and no error, and on
	// the next request the error will be reported. If a request is cancelled
	// the callback will be called with G_IO_ERROR_CANCELLED.
	//
	// During an async request no other sync and async calls are allowed, and
	// will result in G_IO_ERROR_PENDING errors.
	//
	// Any outstanding i/o request with higher priority (lower numerical value)
	// will be executed before an outstanding request with lower priority.
	// Default priority is G_PRIORITY_DEFAULT.
	NextFilesAsync(e FileEnumerator)
	// NextFilesFinish finishes the asynchronous operation started with
	// g_file_enumerator_next_files_async().
	NextFilesFinish(e FileEnumerator, result AsyncResult) error
	// SetPending sets the file enumerator as having pending operations.
	SetPending(e FileEnumerator, pending bool)
}

// fileEnumerator implements the FileEnumerator interface.
type fileEnumerator struct {
	gextras.Objector
}

var _ FileEnumerator = (*fileEnumerator)(nil)

// WrapFileEnumerator wraps a GObject to the right type. It is
// primarily used internally.
func WrapFileEnumerator(obj *externglib.Object) FileEnumerator {
	return FileEnumerator{
		Objector: obj,
	}
}

func marshalFileEnumerator(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileEnumerator(obj), nil
}

// Close releases all resources used by this enumerator, making the
// enumerator return G_IO_ERROR_CLOSED on all calls.
//
// This will be automatically called when the last reference is dropped, but
// you might want to call this function to make sure resources are released
// as early as possible.
func (e fileEnumerator) Close(e FileEnumerator, cancellable Cancellable) error {
	var arg0 *C.GFileEnumerator
	var arg1 *C.GCancellable

	arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))
	arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var errout *C.GError
	var err error

	C.g_file_enumerator_close(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// CloseAsync: asynchronously closes the file enumerator.
//
// If @cancellable is not nil, then the operation can be cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be returned in
// g_file_enumerator_close_finish().
func (e fileEnumerator) CloseAsync(e FileEnumerator) {
	var arg0 *C.GFileEnumerator

	arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))

	C.g_file_enumerator_close_async(arg0, arg1, arg2, arg3, arg4)
}

// CloseFinish finishes closing a file enumerator, started from
// g_file_enumerator_close_async().
//
// If the file enumerator was already closed when
// g_file_enumerator_close_async() was called, then this function will
// report G_IO_ERROR_CLOSED in @error, and return false. If the file
// enumerator had pending operation when the close operation was started,
// then this function will report G_IO_ERROR_PENDING, and return false. If
// @cancellable was not nil, then the operation may have been cancelled by
// triggering the cancellable object from another thread. If the operation
// was cancelled, the error G_IO_ERROR_CANCELLED will be set, and false will
// be returned.
func (e fileEnumerator) CloseFinish(e FileEnumerator, result AsyncResult) error {
	var arg0 *C.GFileEnumerator
	var arg1 *C.GAsyncResult

	arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var errout *C.GError
	var err error

	C.g_file_enumerator_close_finish(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// Child: return a new #GFile which refers to the file named by @info in the
// source directory of @enumerator. This function is primarily intended to
// be used inside loops with g_file_enumerator_next_file().
//
// This is a convenience method that's equivalent to:
//
//    gchar *name = g_file_info_get_name (info);
//    GFile *child = g_file_get_child (g_file_enumerator_get_container (enumr),
//                                     name);
func (e fileEnumerator) Child(e FileEnumerator, info FileInfo) {
	var arg0 *C.GFileEnumerator
	var arg1 *C.GFileInfo

	arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))
	arg1 = (*C.GFileInfo)(unsafe.Pointer(info.Native()))

	C.g_file_enumerator_get_child(arg0, arg1)
}

// Container: get the #GFile container which is being enumerated.
func (e fileEnumerator) Container(e FileEnumerator) {
	var arg0 *C.GFileEnumerator

	arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))

	C.g_file_enumerator_get_container(arg0)
}

// HasPending checks if the file enumerator has pending operations.
func (e fileEnumerator) HasPending(e FileEnumerator) bool {
	var arg0 *C.GFileEnumerator

	arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_file_enumerator_has_pending(arg0)

	if cret {
		ok = true
	}

	return ok
}

// IsClosed checks if the file enumerator has been closed.
func (e fileEnumerator) IsClosed(e FileEnumerator) bool {
	var arg0 *C.GFileEnumerator

	arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_file_enumerator_is_closed(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Iterate: this is a version of g_file_enumerator_next_file() that's easier
// to use correctly from C programs. With g_file_enumerator_next_file(), the
// gboolean return value signifies "end of iteration or error", which
// requires allocation of a temporary #GError.
//
// In contrast, with this function, a false return from
// g_file_enumerator_iterate() *always* means "error". End of iteration is
// signaled by @out_info or @out_child being nil.
//
// Another crucial difference is that the references for @out_info and
// @out_child are owned by @direnum (they are cached as hidden properties).
// You must not unref them in your own code. This makes memory management
// significantly easier for C code in combination with loops.
//
// Finally, this function optionally allows retrieving a #GFile as well.
//
// You must specify at least one of @out_info or @out_child.
//
// The code pattern for correctly using g_file_enumerator_iterate() from C
// is:
//
//    direnum = g_file_enumerate_children (file, ...);
//    while (TRUE)
//      {
//        GFileInfo *info;
//        if (!g_file_enumerator_iterate (direnum, &info, NULL, cancellable, error))
//          goto out;
//        if (!info)
//          break;
//        ... do stuff with "info"; do not unref it! ...
//      }
//
//    out:
//      g_object_unref (direnum); // Note: frees the last @info
func (d fileEnumerator) Iterate(d FileEnumerator, cancellable Cancellable) (outInfo FileInfo, outChild *File, err error) {
	var arg0 *C.GFileEnumerator
	var arg3 *C.GCancellable

	arg0 = (*C.GFileEnumerator)(unsafe.Pointer(d.Native()))
	arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var arg1 *C.GFileInfo
	var outInfo FileInfo
	var arg2 *C.GFile
	var outChild *File
	var errout *C.GError
	var err error

	C.g_file_enumerator_iterate(arg0, &arg1, &arg2, arg3, &errout)

	outInfo = gextras.CastObject(externglib.Take(unsafe.Pointer(&arg1.Native()))).(FileInfo)
	outChild = gextras.CastObject(externglib.Take(unsafe.Pointer(&arg2.Native()))).(*File)
	err = gerror.Take(unsafe.Pointer(errout))

	return outInfo, outChild, err
}

// NextFile returns information for the next file in the enumerated object.
// Will block until the information is available. The Info returned from
// this function will contain attributes that match the attribute string
// that was passed when the Enumerator was created.
//
// See the documentation of Enumerator for information about the order of
// returned files.
//
// On error, returns nil and sets @error to the error. If the enumerator is
// at the end, nil will be returned and @error will be unset.
func (e fileEnumerator) NextFile(e FileEnumerator, cancellable Cancellable) error {
	var arg0 *C.GFileEnumerator
	var arg1 *C.GCancellable

	arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))
	arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var errout *C.GError
	var err error

	C.g_file_enumerator_next_file(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// NextFilesAsync: request information for a number of files from the
// enumerator asynchronously. When all i/o for the operation is finished the
// @callback will be called with the requested information.
//
// See the documentation of Enumerator for information about the order of
// returned files.
//
// The callback can be called with less than @num_files files in case of
// error or at the end of the enumerator. In case of a partial error the
// callback will be called with any succeeding items and no error, and on
// the next request the error will be reported. If a request is cancelled
// the callback will be called with G_IO_ERROR_CANCELLED.
//
// During an async request no other sync and async calls are allowed, and
// will result in G_IO_ERROR_PENDING errors.
//
// Any outstanding i/o request with higher priority (lower numerical value)
// will be executed before an outstanding request with lower priority.
// Default priority is G_PRIORITY_DEFAULT.
func (e fileEnumerator) NextFilesAsync(e FileEnumerator) {
	var arg0 *C.GFileEnumerator

	arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))

	C.g_file_enumerator_next_files_async(arg0, arg1, arg2, arg3, arg4, arg5)
}

// NextFilesFinish finishes the asynchronous operation started with
// g_file_enumerator_next_files_async().
func (e fileEnumerator) NextFilesFinish(e FileEnumerator, result AsyncResult) error {
	var arg0 *C.GFileEnumerator
	var arg1 *C.GAsyncResult

	arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var errout *C.GError
	var err error

	C.g_file_enumerator_next_files_finish(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// SetPending sets the file enumerator as having pending operations.
func (e fileEnumerator) SetPending(e FileEnumerator, pending bool) {
	var arg0 *C.GFileEnumerator
	var arg1 C.gboolean

	arg0 = (*C.GFileEnumerator)(unsafe.Pointer(e.Native()))
	if pending {
		arg1 = C.gboolean(1)
	}

	C.g_file_enumerator_set_pending(arg0, arg1)
}
