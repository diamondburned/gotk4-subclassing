// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_error_get_type()), F: marshalError},
	})
}

// ClearError: if @err or *@err is nil, does nothing. Otherwise, calls
// g_error_free() on *@err and sets *@err to nil.
func ClearError() error {
	var _cerr *C.GError // in

	C.g_clear_error(&_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// PropagateError: if @dest is nil, free @src; otherwise, moves @src into
// *@dest. The error variable @dest points to must be nil.
//
// @src must be non-nil.
//
// Note that @src is no longer valid after this call. If you want to keep using
// the same GError*, you need to set it to nil after calling this function on
// it.
func PropagateError(src error) error {
	var _arg2 *C.GError // out

	_arg2 = (*C.GError)(gerror.New(unsafe.Pointer(src)))

	var _arg1 *C.GError // in

	C.g_propagate_error(_arg2, &_arg1)

	var _dest error // out

	_dest = gerror.Take(unsafe.Pointer(_arg1))

	return _dest
}

// Error: the `GError` structure contains information about an error that has
// occurred.
type Error struct {
	native C.GError
}

// WrapError wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapError(ptr unsafe.Pointer) *Error {
	if ptr == nil {
		return nil
	}

	return (*Error)(ptr)
}

func marshalError(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapError(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (e *Error) Native() unsafe.Pointer {
	return unsafe.Pointer(&e.native)
}

// Code gets the field inside the struct.
func (e *Error) Code() int {
	var v int // out
	v = (int)(e.native.code)
	return v
}

// Message gets the field inside the struct.
func (e *Error) Message() string {
	var v string // out
	v = C.GoString(e.native.message)
	return v
}

// Copy makes a copy of @error.
func (e *Error) Copy() error {
	var _arg0 *C.GError // out

	_arg0 = (*C.GError)(gerror.New(unsafe.Pointer(e)))
	defer C.g_error_free(_arg0)

	var _cret *C.GError // in

	_cret = C.g_error_copy(_arg0)

	var _err error // out

	_err = gerror.Take(unsafe.Pointer(_cret))

	return _err
}

// Free frees a #GError and associated resources.
func (e *Error) Free() {
	var _arg0 *C.GError // out

	_arg0 = (*C.GError)(gerror.New(unsafe.Pointer(e)))
	defer C.g_error_free(_arg0)

	C.g_error_free(_arg0)
}
