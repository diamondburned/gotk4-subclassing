// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

// RandomDouble returns a random #gdouble equally distributed over the range
// [0..1).
func RandomDouble() float64 {
	var _cret C.gdouble // in

	_cret = C.g_random_double()

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// RandomDoubleRange returns a random #gdouble equally distributed over the
// range [@begin..@end).
func RandomDoubleRange(begin float64, end float64) float64 {
	var _arg1 C.gdouble // out
	var _arg2 C.gdouble // out

	_arg1 = C.gdouble(begin)
	_arg2 = C.gdouble(end)

	var _cret C.gdouble // in

	_cret = C.g_random_double_range(_arg1, _arg2)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// RandomInt: return a random #guint32 equally distributed over the range
// [0..2^32-1].
func RandomInt() uint32 {
	var _cret C.guint32 // in

	_cret = C.g_random_int()

	var _guint32 uint32 // out

	_guint32 = (uint32)(_cret)

	return _guint32
}

// RandomIntRange returns a random #gint32 equally distributed over the range
// [@begin..@end-1].
func RandomIntRange(begin int32, end int32) int32 {
	var _arg1 C.gint32 // out
	var _arg2 C.gint32 // out

	_arg1 = C.gint32(begin)
	_arg2 = C.gint32(end)

	var _cret C.gint32 // in

	_cret = C.g_random_int_range(_arg1, _arg2)

	var _gint32 int32 // out

	_gint32 = (int32)(_cret)

	return _gint32
}

// RandomSetSeed sets the seed for the global random number generator, which is
// used by the g_random_* functions, to @seed.
func RandomSetSeed(seed uint32) {
	var _arg1 C.guint32 // out

	_arg1 = C.guint32(seed)

	C.g_random_set_seed(_arg1)
}

// Rand: the GRand struct is an opaque data structure. It should only be
// accessed through the g_rand_* functions.
type Rand struct {
	native C.GRand
}

// WrapRand wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRand(ptr unsafe.Pointer) *Rand {
	if ptr == nil {
		return nil
	}

	return (*Rand)(ptr)
}

func marshalRand(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRand(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *Rand) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Double returns the next random #gdouble from @rand_ equally distributed over
// the range [0..1).
func (r *Rand) Double() float64 {
	var _arg0 *C.GRand // out

	_arg0 = (*C.GRand)(unsafe.Pointer(r.Native()))

	var _cret C.gdouble // in

	_cret = C.g_rand_double(_arg0)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// DoubleRange returns the next random #gdouble from @rand_ equally distributed
// over the range [@begin..@end).
func (r *Rand) DoubleRange(begin float64, end float64) float64 {
	var _arg0 *C.GRand  // out
	var _arg1 C.gdouble // out
	var _arg2 C.gdouble // out

	_arg0 = (*C.GRand)(unsafe.Pointer(r.Native()))
	_arg1 = C.gdouble(begin)
	_arg2 = C.gdouble(end)

	var _cret C.gdouble // in

	_cret = C.g_rand_double_range(_arg0, _arg1, _arg2)

	var _gdouble float64 // out

	_gdouble = (float64)(_cret)

	return _gdouble
}

// Free frees the memory allocated for the #GRand.
func (r *Rand) Free() {
	var _arg0 *C.GRand // out

	_arg0 = (*C.GRand)(unsafe.Pointer(r.Native()))

	C.g_rand_free(_arg0)
}

// Int returns the next random #guint32 from @rand_ equally distributed over the
// range [0..2^32-1].
func (r *Rand) Int() uint32 {
	var _arg0 *C.GRand // out

	_arg0 = (*C.GRand)(unsafe.Pointer(r.Native()))

	var _cret C.guint32 // in

	_cret = C.g_rand_int(_arg0)

	var _guint32 uint32 // out

	_guint32 = (uint32)(_cret)

	return _guint32
}

// IntRange returns the next random #gint32 from @rand_ equally distributed over
// the range [@begin..@end-1].
func (r *Rand) IntRange(begin int32, end int32) int32 {
	var _arg0 *C.GRand // out
	var _arg1 C.gint32 // out
	var _arg2 C.gint32 // out

	_arg0 = (*C.GRand)(unsafe.Pointer(r.Native()))
	_arg1 = C.gint32(begin)
	_arg2 = C.gint32(end)

	var _cret C.gint32 // in

	_cret = C.g_rand_int_range(_arg0, _arg1, _arg2)

	var _gint32 int32 // out

	_gint32 = (int32)(_cret)

	return _gint32
}

// SetSeed sets the seed for the random number generator #GRand to @seed.
func (r *Rand) SetSeed(seed uint32) {
	var _arg0 *C.GRand  // out
	var _arg1 C.guint32 // out

	_arg0 = (*C.GRand)(unsafe.Pointer(r.Native()))
	_arg1 = C.guint32(seed)

	C.g_rand_set_seed(_arg0, _arg1)
}

// SetSeedArray initializes the random number generator by an array of longs.
// Array can be of arbitrary size, though only the first 624 values are taken.
// This function is useful if you have many low entropy seeds, or if you require
// more then 32 bits of actual entropy for your application.
func (r *Rand) SetSeedArray(seed *uint32, seedLength uint) {
	var _arg0 *C.GRand   // out
	var _arg1 *C.guint32 // out
	var _arg2 C.guint    // out

	_arg0 = (*C.GRand)(unsafe.Pointer(r.Native()))
	_arg1 = *C.guint32(seed)
	_arg2 = C.guint(seedLength)

	C.g_rand_set_seed_array(_arg0, _arg1, _arg2)
}
