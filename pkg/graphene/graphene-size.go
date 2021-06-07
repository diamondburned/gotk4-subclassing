// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: graphene-gobject-1.0 graphene-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.graphene_size_get_type()), F: marshalSize},
	})
}

// SizeZero: a constant pointer to a zero #graphene_size_t, useful for equality
// checks and interpolations.
func SizeZero() {
	C.graphene_size_zero()
}

// Size: a size.
type Size struct {
	native C.graphene_size_t
}

// WrapSize wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapSize(ptr unsafe.Pointer) *Size {
	if ptr == nil {
		return nil
	}

	return (*Size)(ptr)
}

func marshalSize(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapSize(unsafe.Pointer(b)), nil
}

// NewSizeAlloc constructs a struct Size.
func NewSizeAlloc() {
	C.graphene_size_alloc()
}

// Native returns the underlying C source pointer.
func (s *Size) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// Width gets the field inside the struct.
func (s *Size) Width() float32 {
	var v float32
	v = float32(s.native.width)
	return v
}

// Height gets the field inside the struct.
func (s *Size) Height() float32 {
	var v float32
	v = float32(s.native.height)
	return v
}

// Equal checks whether the two give #graphene_size_t are equal.
func (a *Size) Equal(a *Size, b *Size) bool {
	var arg0 *C.graphene_size_t
	var arg1 *C.graphene_size_t

	arg0 = (*C.graphene_size_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_size_t)(unsafe.Pointer(b.Native()))

	var cret C._Bool
	var ok bool

	cret = C.graphene_size_equal(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// Free frees the resources allocated by graphene_size_alloc().
func (s *Size) Free(s *Size) {
	var arg0 *C.graphene_size_t

	arg0 = (*C.graphene_size_t)(unsafe.Pointer(s.Native()))

	C.graphene_size_free(arg0)
}

// Init initializes a #graphene_size_t using the given @width and @height.
func (s *Size) Init(s *Size, width float32, height float32) {
	var arg0 *C.graphene_size_t
	var arg1 C.float
	var arg2 C.float

	arg0 = (*C.graphene_size_t)(unsafe.Pointer(s.Native()))
	arg1 = C.float(width)
	arg2 = C.float(height)

	C.graphene_size_init(arg0, arg1, arg2)
}

// InitFromSize initializes a #graphene_size_t using the width and height of the
// given @src.
func (s *Size) InitFromSize(s *Size, src *Size) {
	var arg0 *C.graphene_size_t
	var arg1 *C.graphene_size_t

	arg0 = (*C.graphene_size_t)(unsafe.Pointer(s.Native()))
	arg1 = (*C.graphene_size_t)(unsafe.Pointer(src.Native()))

	C.graphene_size_init_from_size(arg0, arg1)
}

// Interpolate: linearly interpolates the two given #graphene_size_t using the
// given interpolation @factor.
func (a *Size) Interpolate(a *Size, b *Size, factor float64) *Size {
	var arg0 *C.graphene_size_t
	var arg1 *C.graphene_size_t
	var arg2 C.double

	arg0 = (*C.graphene_size_t)(unsafe.Pointer(a.Native()))
	arg1 = (*C.graphene_size_t)(unsafe.Pointer(b.Native()))
	arg2 = C.double(factor)

	var arg3 C.graphene_size_t
	var res *Size

	C.graphene_size_interpolate(arg0, arg1, arg2, &arg3)

	res = WrapSize(unsafe.Pointer(&arg3))

	return res
}

// Scale scales the components of a #graphene_size_t using the given @factor.
func (s *Size) Scale(s *Size, factor float32) *Size {
	var arg0 *C.graphene_size_t
	var arg1 C.float

	arg0 = (*C.graphene_size_t)(unsafe.Pointer(s.Native()))
	arg1 = C.float(factor)

	var arg2 C.graphene_size_t
	var res *Size

	C.graphene_size_scale(arg0, arg1, &arg2)

	res = WrapSize(unsafe.Pointer(&arg2))

	return res
}
