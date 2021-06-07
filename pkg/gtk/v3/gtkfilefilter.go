// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// FileFilterFunc: the type of function that is used with custom filters, see
// gtk_file_filter_add_custom().
type FileFilterFunc func(filterInfo *FileFilterInfo) bool

//export gotk4_FileFilterFunc
func gotk4_FileFilterFunc(arg0 *C.GtkFileFilterInfo, arg1 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(FileFilterFunc)
	ret := fn(filterInfo, data)

	if ret {
		cret = C.gboolean(1)
	}

	return cret
}

// FileFilterInfo: a FileFilterInfo-struct is used to pass information about the
// tested file to gtk_file_filter_filter().
type FileFilterInfo struct {
	native C.GtkFileFilterInfo
}

// WrapFileFilterInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFileFilterInfo(ptr unsafe.Pointer) *FileFilterInfo {
	if ptr == nil {
		return nil
	}

	return (*FileFilterInfo)(ptr)
}

func marshalFileFilterInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapFileFilterInfo(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (f *FileFilterInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&f.native)
}

// Contains gets the field inside the struct.
func (f *FileFilterInfo) Contains() FileFilterFlags {
	var v FileFilterFlags
	v = FileFilterFlags(f.native.contains)
	return v
}

// Filename gets the field inside the struct.
func (f *FileFilterInfo) Filename() string {
	var v string
	v = C.GoString(f.native.filename)
	return v
}

// URI gets the field inside the struct.
func (f *FileFilterInfo) URI() string {
	var v string
	v = C.GoString(f.native.uri)
	return v
}

// DisplayName gets the field inside the struct.
func (f *FileFilterInfo) DisplayName() string {
	var v string
	v = C.GoString(f.native.display_name)
	return v
}

// MIMEType gets the field inside the struct.
func (f *FileFilterInfo) MIMEType() string {
	var v string
	v = C.GoString(f.native.mime_type)
	return v
}
