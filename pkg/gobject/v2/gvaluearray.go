// Code generated by girgen. DO NOT EDIT.

package gobject

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gobject-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_value_array_get_type()), F: marshalValueArray},
	})
}

// ValueArray: a Array contains an array of #GValue elements.
type ValueArray struct {
	native C.GValueArray
}

// WrapValueArray wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapValueArray(ptr unsafe.Pointer) *ValueArray {
	if ptr == nil {
		return nil
	}

	return (*ValueArray)(ptr)
}

func marshalValueArray(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapValueArray(unsafe.Pointer(b)), nil
}

// NewValueArray constructs a struct ValueArray.
func NewValueArray(nPrealloced uint) *ValueArray {
	var arg1 C.guint

	arg1 = C.guint(nPrealloced)

	var cret *C.GValueArray

	cret = C.g_value_array_new(arg1)

	var valueArray *ValueArray

	valueArray = WrapValueArray(unsafe.Pointer(cret))
	runtime.SetFinalizer(valueArray, func(v *ValueArray) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return valueArray
}

// Native returns the underlying C source pointer.
func (v *ValueArray) Native() unsafe.Pointer {
	return unsafe.Pointer(&v.native)
}

// NValues gets the field inside the struct.
func (v *ValueArray) NValues() uint {
	var v uint
	v = (uint)(v.native.n_values)
	return v
}

// Values gets the field inside the struct.
func (v *ValueArray) Values() **externglib.Value {
	var v **externglib.Value
	v = externglib.ValueFromNative(unsafe.Pointer(v.native.values))
	return v
}

// Append: insert a copy of @value as last element of @value_array. If @value is
// nil, an uninitialized value is appended.
func (v *ValueArray) Append(value **externglib.Value) *ValueArray {
	var arg0 *C.GValueArray
	var arg1 *C.GValue

	arg0 = (*C.GValueArray)(unsafe.Pointer(v.Native()))
	arg1 = (*C.GValue)(value.GValue)

	var cret *C.GValueArray

	cret = C.g_value_array_append(arg0, arg1)

	var valueArray *ValueArray

	valueArray = WrapValueArray(unsafe.Pointer(cret))

	return valueArray
}

// Copy: construct an exact copy of a Array by duplicating all its contents.
func (v *ValueArray) Copy() *ValueArray {
	var arg0 *C.GValueArray

	arg0 = (*C.GValueArray)(unsafe.Pointer(v.Native()))

	var cret *C.GValueArray

	cret = C.g_value_array_copy(arg0)

	var valueArray *ValueArray

	valueArray = WrapValueArray(unsafe.Pointer(cret))
	runtime.SetFinalizer(valueArray, func(v *ValueArray) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return valueArray
}

// Free: free a Array including its contents.
func (v *ValueArray) Free() {
	var arg0 *C.GValueArray

	arg0 = (*C.GValueArray)(unsafe.Pointer(v.Native()))

	C.g_value_array_free(arg0)
}

// Nth: return a pointer to the value at @index_ containd in @value_array.
func (v *ValueArray) Nth(index_ uint) **externglib.Value {
	var arg0 *C.GValueArray
	var arg1 C.guint

	arg0 = (*C.GValueArray)(unsafe.Pointer(v.Native()))
	arg1 = C.guint(index_)

	var cret *C.GValue

	cret = C.g_value_array_get_nth(arg0, arg1)

	var value **externglib.Value

	value = externglib.ValueFromNative(unsafe.Pointer(cret))

	return value
}

// Insert: insert a copy of @value at specified position into @value_array. If
// @value is nil, an uninitialized value is inserted.
func (v *ValueArray) Insert(index_ uint, value **externglib.Value) *ValueArray {
	var arg0 *C.GValueArray
	var arg1 C.guint
	var arg2 *C.GValue

	arg0 = (*C.GValueArray)(unsafe.Pointer(v.Native()))
	arg1 = C.guint(index_)
	arg2 = (*C.GValue)(value.GValue)

	var cret *C.GValueArray

	cret = C.g_value_array_insert(arg0, arg1, arg2)

	var valueArray *ValueArray

	valueArray = WrapValueArray(unsafe.Pointer(cret))

	return valueArray
}

// Prepend: insert a copy of @value as first element of @value_array. If @value
// is nil, an uninitialized value is prepended.
func (v *ValueArray) Prepend(value **externglib.Value) *ValueArray {
	var arg0 *C.GValueArray
	var arg1 *C.GValue

	arg0 = (*C.GValueArray)(unsafe.Pointer(v.Native()))
	arg1 = (*C.GValue)(value.GValue)

	var cret *C.GValueArray

	cret = C.g_value_array_prepend(arg0, arg1)

	var valueArray *ValueArray

	valueArray = WrapValueArray(unsafe.Pointer(cret))

	return valueArray
}

// Remove: remove the value at position @index_ from @value_array.
func (v *ValueArray) Remove(index_ uint) *ValueArray {
	var arg0 *C.GValueArray
	var arg1 C.guint

	arg0 = (*C.GValueArray)(unsafe.Pointer(v.Native()))
	arg1 = C.guint(index_)

	var cret *C.GValueArray

	cret = C.g_value_array_remove(arg0, arg1)

	var valueArray *ValueArray

	valueArray = WrapValueArray(unsafe.Pointer(cret))

	return valueArray
}

// SortWithData: sort @value_array using @compare_func to compare the elements
// according to the semantics of DataFunc.
//
// The current implementation uses the same sorting algorithm as standard C
// qsort() function.
func (v *ValueArray) SortWithData() *ValueArray {
	var arg0 *C.GValueArray

	arg0 = (*C.GValueArray)(unsafe.Pointer(v.Native()))

	var cret *C.GValueArray

	cret = C.g_value_array_sort_with_data(arg0)

	var valueArray *ValueArray

	valueArray = WrapValueArray(unsafe.Pointer(cret))

	return valueArray
}
