// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
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
func ComputeChecksumForData() string {
	var cret *C.gchar

	cret = C.g_compute_checksum_for_data()

	var utf8 string

	utf8 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return utf8
}

// ComputeChecksumForString computes the checksum of a string.
//
// The hexadecimal string returned will be in lower case.
func ComputeChecksumForString(checksumType ChecksumType, str string, length int) string {
	var arg1 C.GChecksumType
	var arg2 *C.gchar
	var arg3 C.gssize

	arg1 = (C.GChecksumType)(checksumType)
	arg2 = (*C.gchar)(C.CString(str))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = C.gssize(length)

	var cret *C.gchar

	cret = C.g_compute_checksum_for_string(arg1, arg2, arg3)

	var utf8 string

	utf8 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return utf8
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

// NewChecksum constructs a struct Checksum.
func NewChecksum(checksumType ChecksumType) *Checksum {
	var arg1 C.GChecksumType

	arg1 = (C.GChecksumType)(checksumType)

	var cret *C.GChecksum

	cret = C.g_checksum_new(arg1)

	var checksum *Checksum

	checksum = WrapChecksum(unsafe.Pointer(cret))
	runtime.SetFinalizer(checksum, func(v *Checksum) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return checksum
}

// Native returns the underlying C source pointer.
func (c *Checksum) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}

// Copy copies a #GChecksum. If @checksum has been closed, by calling
// g_checksum_get_string() or g_checksum_get_digest(), the copied checksum will
// be closed as well.
func (c *Checksum) Copy() *Checksum {
	var arg0 *C.GChecksum

	arg0 = (*C.GChecksum)(unsafe.Pointer(c.Native()))

	var cret *C.GChecksum

	cret = C.g_checksum_copy(arg0)

	var ret *Checksum

	ret = WrapChecksum(unsafe.Pointer(cret))
	runtime.SetFinalizer(ret, func(v *Checksum) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return ret
}

// Free frees the memory allocated for @checksum.
func (c *Checksum) Free() {
	var arg0 *C.GChecksum

	arg0 = (*C.GChecksum)(unsafe.Pointer(c.Native()))

	C.g_checksum_free(arg0)
}

// Digest gets the digest from @checksum as a raw binary vector and places it
// into @buffer. The size of the digest depends on the type of checksum.
//
// Once this function has been called, the #GChecksum is closed and can no
// longer be updated with g_checksum_update().
func (c *Checksum) Digest() {
	var arg0 *C.GChecksum

	arg0 = (*C.GChecksum)(unsafe.Pointer(c.Native()))

	C.g_checksum_get_digest(arg0)
}

// String gets the digest as a hexadecimal string.
//
// Once this function has been called the #GChecksum can no longer be updated
// with g_checksum_update().
//
// The hexadecimal characters will be lower case.
func (c *Checksum) String() string {
	var arg0 *C.GChecksum

	arg0 = (*C.GChecksum)(unsafe.Pointer(c.Native()))

	var cret *C.gchar

	cret = C.g_checksum_get_string(arg0)

	var utf8 string

	utf8 = C.GoString(cret)

	return utf8
}

// Reset resets the state of the @checksum back to its initial state.
func (c *Checksum) Reset() {
	var arg0 *C.GChecksum

	arg0 = (*C.GChecksum)(unsafe.Pointer(c.Native()))

	C.g_checksum_reset(arg0)
}

// Update feeds @data into an existing #GChecksum. The checksum must still be
// open, that is g_checksum_get_string() or g_checksum_get_digest() must not
// have been called on @checksum.
func (c *Checksum) Update() {
	var arg0 *C.GChecksum

	arg0 = (*C.GChecksum)(unsafe.Pointer(c.Native()))

	C.g_checksum_update(arg0)
}
