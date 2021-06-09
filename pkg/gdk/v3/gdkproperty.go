// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// PropertyChange changes the contents of a property on a window.
func PropertyChange(window Window, property Atom, typ Atom, format int, mode PropMode, data *byte, nelements int) {
	var arg1 *C.GdkWindow
	var arg2 C.GdkAtom
	var arg3 C.GdkAtom
	var arg4 C.gint
	var arg5 C.GdkPropMode
	var arg6 *C.guchar
	var arg7 C.gint

	arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	arg2 = (C.GdkAtom)(unsafe.Pointer(property.Native()))
	arg3 = (C.GdkAtom)(unsafe.Pointer(typ.Native()))
	arg4 = C.gint(format)
	arg5 = (C.GdkPropMode)(mode)
	arg6 = *C.guchar(data)
	arg7 = C.gint(nelements)

	C.gdk_property_change(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// PropertyDelete deletes a property from a window.
func PropertyDelete(window Window, property Atom) {
	var arg1 *C.GdkWindow
	var arg2 C.GdkAtom

	arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	arg2 = (C.GdkAtom)(unsafe.Pointer(property.Native()))

	C.gdk_property_delete(arg1, arg2)
}

// UTF8ToStringTarget converts an UTF-8 string into the best possible
// representation as a STRING. The representation of characters not in STRING is
// not specified; it may be as pseudo-escape sequences \x{ABCD}, or it may be in
// some other form of approximation.
func UTF8ToStringTarget(str string) string {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.gchar

	cret = C.gdk_utf8_to_string_target(arg1)

	var utf8 string

	utf8 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return utf8
}
