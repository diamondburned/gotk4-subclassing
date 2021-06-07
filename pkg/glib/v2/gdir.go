// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

// Dir: an opaque structure representing an opened directory.
type Dir struct {
	native C.GDir
}

// WrapDir wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDir(ptr unsafe.Pointer) *Dir {
	if ptr == nil {
		return nil
	}

	return (*Dir)(ptr)
}

func marshalDir(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapDir(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *Dir) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// Close closes the directory and deallocates all related resources.
func (d *Dir) Close(d *Dir) {
	var arg0 *C.GDir

	arg0 = (*C.GDir)(unsafe.Pointer(d.Native()))

	C.g_dir_close(arg0)
}

// ReadName retrieves the name of another entry in the directory, or nil. The
// order of entries returned from this function is not defined, and may vary by
// file system or other operating-system dependent factors.
//
// nil may also be returned in case of errors. On Unix, you can check `errno` to
// find out if nil was returned because of an error.
//
// On Unix, the '.' and '..' entries are omitted, and the returned name is in
// the on-disk encoding.
//
// On Windows, as is true of all GLib functions which operate on filenames, the
// returned name is in UTF-8.
func (d *Dir) ReadName(d *Dir) {
	var arg0 *C.GDir

	arg0 = (*C.GDir)(unsafe.Pointer(d.Native()))

	C.g_dir_read_name(arg0)
}

// Rewind resets the given directory. The next call to g_dir_read_name() will
// return the first entry again.
func (d *Dir) Rewind(d *Dir) {
	var arg0 *C.GDir

	arg0 = (*C.GDir)(unsafe.Pointer(d.Native()))

	C.g_dir_rewind(arg0)
}
