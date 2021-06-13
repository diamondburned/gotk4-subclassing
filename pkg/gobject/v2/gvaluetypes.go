// Code generated by girgen. DO NOT EDIT.

package gobject

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gobject-2.0 gobject-introspection-1.0 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib-object.h>
import "C"

func GTypeGetType() externglib.Type {
	var _cret C.GType // in

	_cret = C.g_gtype_get_type()

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// PointerTypeRegisterStatic creates a new G_TYPE_POINTER derived type id for a
// new pointer type with name @name.
func PointerTypeRegisterStatic(name string) externglib.Type {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GType // in

	_cret = C.g_pointer_type_register_static(_arg1)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// StrdupValueContents: return a newly allocated string, which describes the
// contents of a #GValue. The main purpose of this function is to describe
// #GValue contents for debugging output, the way in which the contents are
// described may change between different GLib versions.
func StrdupValueContents(value **externglib.Value) string {
	var _arg1 *C.GValue // out

	_arg1 = (*C.GValue)(value.GValue)

	var _cret *C.gchar // in

	_cret = C.g_strdup_value_contents(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
