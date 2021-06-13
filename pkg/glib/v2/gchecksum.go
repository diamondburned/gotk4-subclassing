// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_checksum_get_type()), F: marshalChecksum},
	})
}

// ChecksumType: the hashing algorithm to be used by #GChecksum when performing
// the digest of some data.
//
// Note that the Type enumeration may be extended at a later date to include new
// hashing algorithm types.
type ChecksumType int

const (
	// ChecksumTypeMD5: use the MD5 hashing algorithm
	ChecksumTypeMD5 ChecksumType = 0
	// ChecksumTypeSHA1: use the SHA-1 hashing algorithm
	ChecksumTypeSHA1 ChecksumType = 1
	// ChecksumTypeSHA256: use the SHA-256 hashing algorithm
	ChecksumTypeSHA256 ChecksumType = 2
	// ChecksumTypeSHA512: use the SHA-512 hashing algorithm (Since: 2.36)
	ChecksumTypeSHA512 ChecksumType = 3
	// ChecksumTypeSHA384: use the SHA-384 hashing algorithm (Since: 2.51)
	ChecksumTypeSHA384 ChecksumType = 4
)

// ComputeChecksumForData computes the checksum for a binary @data of @length.
// This is a convenience wrapper for g_checksum_new(), g_checksum_get_string()
// and g_checksum_free().
//
// The hexadecimal string returned will be in lower case.
func ComputeChecksumForData(checksumType ChecksumType, data []byte) string {
	var _arg1 C.GChecksumType // out
	var _arg2 *C.guchar
	var _arg3 C.gsize

	_arg1 = (C.GChecksumType)(checksumType)
	_arg3 = C.gsize(len(data))
	_arg2 = (*C.guchar)(unsafe.Pointer(&data[0]))

	var _cret *C.gchar // in

	_cret = C.g_compute_checksum_for_data(_arg1, _arg2, _arg3)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ComputeChecksumForString computes the checksum of a string.
//
// The hexadecimal string returned will be in lower case.
func ComputeChecksumForString(checksumType ChecksumType, str string, length int) string {
	var _arg1 C.GChecksumType // out
	var _arg2 *C.gchar        // out
	var _arg3 C.gssize        // out

	_arg1 = (C.GChecksumType)(checksumType)
	_arg2 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gssize(length)

	var _cret *C.gchar // in

	_cret = C.g_compute_checksum_for_string(_arg1, _arg2, _arg3)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Checksum: an opaque structure representing a checksumming operation. To
// create a new GChecksum, use g_checksum_new(). To free a GChecksum, use
// g_checksum_free().
type Checksum struct {
	native C.GChecksum
}

// WrapChecksum wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapChecksum(ptr unsafe.Pointer) *Checksum {
	if ptr == nil {
		return nil
	}

	return (*Checksum)(ptr)
}

func marshalChecksum(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapChecksum(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *Checksum) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}

// Free frees the memory allocated for @checksum.
func (c *Checksum) Free() {
	var _arg0 *C.GChecksum // out

	_arg0 = (*C.GChecksum)(unsafe.Pointer(c.Native()))

	C.g_checksum_free(_arg0)
}

// String gets the digest as a hexadecimal string.
//
// Once this function has been called the #GChecksum can no longer be updated
// with g_checksum_update().
//
// The hexadecimal characters will be lower case.
func (c *Checksum) String() string {
	var _arg0 *C.GChecksum // out

	_arg0 = (*C.GChecksum)(unsafe.Pointer(c.Native()))

	var _cret *C.gchar // in

	_cret = C.g_checksum_get_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Reset resets the state of the @checksum back to its initial state.
func (c *Checksum) Reset() {
	var _arg0 *C.GChecksum // out

	_arg0 = (*C.GChecksum)(unsafe.Pointer(c.Native()))

	C.g_checksum_reset(_arg0)
}

// Update feeds @data into an existing #GChecksum. The checksum must still be
// open, that is g_checksum_get_string() or g_checksum_get_digest() must not
// have been called on @checksum.
func (c *Checksum) Update(data []byte) {
	var _arg0 *C.GChecksum // out
	var _arg1 *C.guchar
	var _arg2 C.gssize

	_arg0 = (*C.GChecksum)(unsafe.Pointer(c.Native()))
	_arg2 = C.gssize(len(data))
	_arg1 = (*C.guchar)(unsafe.Pointer(&data[0]))

	C.g_checksum_update(_arg0, _arg1, _arg2)
}
