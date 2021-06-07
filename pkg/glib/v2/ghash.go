// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_hash_table_get_type()), F: marshalHashTable},
	})
}

// HRFunc specifies the type of the function passed to
// g_hash_table_foreach_remove(). It is called with each key/value pair,
// together with the @user_data parameter passed to
// g_hash_table_foreach_remove(). It should return true if the key/value pair
// should be removed from the Table.
type HRFunc func(key interface{}, value interface{}) bool

//export gotk4_HRFunc
func gotk4_HRFunc(arg0 C.gpointer, arg1 C.gpointer, arg2 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(HRFunc)
	ret := fn(key, value, userData)

	if ret {
		cret = C.gboolean(1)
	}

	return cret
}

// DirectEqual compares two #gpointer arguments and returns true if they are
// equal. It can be passed to g_hash_table_new() as the @key_equal_func
// parameter, when using opaque pointers compared by pointer value as keys in a
// Table.
//
// This equality function is also appropriate for keys that are integers stored
// in pointers, such as `GINT_TO_POINTER (n)`.
func DirectEqual(v1 interface{}, v2 interface{}) bool {
	var arg1 C.gpointer
	var arg2 C.gpointer

	arg1 = C.gpointer(v1)
	arg2 = C.gpointer(v2)

	var cret C.gboolean
	var ok bool

	cret = C.g_direct_equal(arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// DirectHash converts a gpointer to a hash value. It can be passed to
// g_hash_table_new() as the @hash_func parameter, when using opaque pointers
// compared by pointer value as keys in a Table.
//
// This hash function is also appropriate for keys that are integers stored in
// pointers, such as `GINT_TO_POINTER (n)`.
func DirectHash(v interface{}) {
	var arg1 C.gpointer

	arg1 = C.gpointer(v)

	C.g_direct_hash(arg1)
}

// DoubleEqual compares the two #gdouble values being pointed to and returns
// true if they are equal. It can be passed to g_hash_table_new() as the
// @key_equal_func parameter, when using non-nil pointers to doubles as keys in
// a Table.
func DoubleEqual(v1 interface{}, v2 interface{}) bool {
	var arg1 C.gpointer
	var arg2 C.gpointer

	arg1 = C.gpointer(v1)
	arg2 = C.gpointer(v2)

	var cret C.gboolean
	var ok bool

	cret = C.g_double_equal(arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// DoubleHash converts a pointer to a #gdouble to a hash value. It can be passed
// to g_hash_table_new() as the @hash_func parameter, It can be passed to
// g_hash_table_new() as the @hash_func parameter, when using non-nil pointers
// to doubles as keys in a Table.
func DoubleHash(v interface{}) {
	var arg1 C.gpointer

	arg1 = C.gpointer(v)

	C.g_double_hash(arg1)
}

// HashTableAdd: this is a convenience function for using a Table as a set. It
// is equivalent to calling g_hash_table_replace() with @key as both the key and
// the value.
//
// In particular, this means that if @key already exists in the hash table, then
// the old copy of @key in the hash table is freed and @key replaces it in the
// table.
//
// When a hash table only ever contains keys that have themselves as the
// corresponding value it is able to be stored more efficiently. See the
// discussion in the section description.
//
// Starting from GLib 2.40, this function returns a boolean value to indicate
// whether the newly added value was already in the hash table or not.
func HashTableAdd(hashTable *HashTable, key interface{}) bool {
	var arg1 *C.GHashTable
	var arg2 C.gpointer

	arg1 = (*C.GHashTable)(unsafe.Pointer(hashTable.Native()))
	arg2 = C.gpointer(key)

	var cret C.gboolean
	var ok bool

	cret = C.g_hash_table_add(arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// HashTableContains checks if @key is in @hash_table.
func HashTableContains(hashTable *HashTable, key interface{}) bool {
	var arg1 *C.GHashTable
	var arg2 C.gpointer

	arg1 = (*C.GHashTable)(unsafe.Pointer(hashTable.Native()))
	arg2 = C.gpointer(key)

	var cret C.gboolean
	var ok bool

	cret = C.g_hash_table_contains(arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// HashTableDestroy destroys all keys and values in the Table and decrements its
// reference count by 1. If keys and/or values are dynamically allocated, you
// should either free them first or create the Table with destroy notifiers
// using g_hash_table_new_full(). In the latter case the destroy functions you
// supplied will be called on all keys and values during the destruction phase.
func HashTableDestroy(hashTable *HashTable) {
	var arg1 *C.GHashTable

	arg1 = (*C.GHashTable)(unsafe.Pointer(hashTable.Native()))

	C.g_hash_table_destroy(arg1)
}

// HashTableInsert inserts a new key and value into a Table.
//
// If the key already exists in the Table its current value is replaced with the
// new value. If you supplied a @value_destroy_func when creating the Table, the
// old value is freed using that function. If you supplied a @key_destroy_func
// when creating the Table, the passed key is freed using that function.
//
// Starting from GLib 2.40, this function returns a boolean value to indicate
// whether the newly added value was already in the hash table or not.
func HashTableInsert(hashTable *HashTable, key interface{}, value interface{}) bool {
	var arg1 *C.GHashTable
	var arg2 C.gpointer
	var arg3 C.gpointer

	arg1 = (*C.GHashTable)(unsafe.Pointer(hashTable.Native()))
	arg2 = C.gpointer(key)
	arg3 = C.gpointer(value)

	var cret C.gboolean
	var ok bool

	cret = C.g_hash_table_insert(arg1, arg2, arg3)

	if cret {
		ok = true
	}

	return ok
}

// HashTableLookup looks up a key in a Table. Note that this function cannot
// distinguish between a key that is not present and one which is present and
// has the value nil. If you need this distinction, use
// g_hash_table_lookup_extended().
func HashTableLookup(hashTable *HashTable, key interface{}) {
	var arg1 *C.GHashTable
	var arg2 C.gpointer

	arg1 = (*C.GHashTable)(unsafe.Pointer(hashTable.Native()))
	arg2 = C.gpointer(key)

	C.g_hash_table_lookup(arg1, arg2)
}

// HashTableLookupExtended looks up a key in the Table, returning the original
// key and the associated value and a #gboolean which is true if the key was
// found. This is useful if you need to free the memory allocated for the
// original key, for example before calling g_hash_table_remove().
//
// You can actually pass nil for @lookup_key to test whether the nil key exists,
// provided the hash and equal functions of @hash_table are nil-safe.
func HashTableLookupExtended(hashTable *HashTable, lookupKey interface{}) (origKey interface{}, value interface{}, ok bool) {
	var arg1 *C.GHashTable
	var arg2 C.gpointer

	arg1 = (*C.GHashTable)(unsafe.Pointer(hashTable.Native()))
	arg2 = C.gpointer(lookupKey)

	var arg3 C.gpointer
	var origKey interface{}
	var arg4 C.gpointer
	var value interface{}
	var cret C.gboolean
	var ok bool

	cret = C.g_hash_table_lookup_extended(arg1, arg2, &arg3, &arg4)

	origKey = interface{}(&arg3)
	value = interface{}(&arg4)
	if cret {
		ok = true
	}

	return origKey, value, ok
}

// HashTableRemove removes a key and its associated value from a Table.
//
// If the Table was created using g_hash_table_new_full(), the key and value are
// freed using the supplied destroy functions, otherwise you have to make sure
// that any dynamically allocated values are freed yourself.
func HashTableRemove(hashTable *HashTable, key interface{}) bool {
	var arg1 *C.GHashTable
	var arg2 C.gpointer

	arg1 = (*C.GHashTable)(unsafe.Pointer(hashTable.Native()))
	arg2 = C.gpointer(key)

	var cret C.gboolean
	var ok bool

	cret = C.g_hash_table_remove(arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// HashTableRemoveAll removes all keys and their associated values from a Table.
//
// If the Table was created using g_hash_table_new_full(), the keys and values
// are freed using the supplied destroy functions, otherwise you have to make
// sure that any dynamically allocated values are freed yourself.
func HashTableRemoveAll(hashTable *HashTable) {
	var arg1 *C.GHashTable

	arg1 = (*C.GHashTable)(unsafe.Pointer(hashTable.Native()))

	C.g_hash_table_remove_all(arg1)
}

// HashTableReplace inserts a new key and value into a Table similar to
// g_hash_table_insert(). The difference is that if the key already exists in
// the Table, it gets replaced by the new key. If you supplied a
// @value_destroy_func when creating the Table, the old value is freed using
// that function. If you supplied a @key_destroy_func when creating the Table,
// the old key is freed using that function.
//
// Starting from GLib 2.40, this function returns a boolean value to indicate
// whether the newly added value was already in the hash table or not.
func HashTableReplace(hashTable *HashTable, key interface{}, value interface{}) bool {
	var arg1 *C.GHashTable
	var arg2 C.gpointer
	var arg3 C.gpointer

	arg1 = (*C.GHashTable)(unsafe.Pointer(hashTable.Native()))
	arg2 = C.gpointer(key)
	arg3 = C.gpointer(value)

	var cret C.gboolean
	var ok bool

	cret = C.g_hash_table_replace(arg1, arg2, arg3)

	if cret {
		ok = true
	}

	return ok
}

// HashTableSize returns the number of elements contained in the Table.
func HashTableSize(hashTable *HashTable) {
	var arg1 *C.GHashTable

	arg1 = (*C.GHashTable)(unsafe.Pointer(hashTable.Native()))

	C.g_hash_table_size(arg1)
}

// HashTableSteal removes a key and its associated value from a Table without
// calling the key and value destroy functions.
func HashTableSteal(hashTable *HashTable, key interface{}) bool {
	var arg1 *C.GHashTable
	var arg2 C.gpointer

	arg1 = (*C.GHashTable)(unsafe.Pointer(hashTable.Native()))
	arg2 = C.gpointer(key)

	var cret C.gboolean
	var ok bool

	cret = C.g_hash_table_steal(arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// HashTableStealAll removes all keys and their associated values from a Table
// without calling the key and value destroy functions.
func HashTableStealAll(hashTable *HashTable) {
	var arg1 *C.GHashTable

	arg1 = (*C.GHashTable)(unsafe.Pointer(hashTable.Native()))

	C.g_hash_table_steal_all(arg1)
}

// HashTableStealExtended looks up a key in the Table, stealing the original key
// and the associated value and returning true if the key was found. If the key
// was not found, false is returned.
//
// If found, the stolen key and value are removed from the hash table without
// calling the key and value destroy functions, and ownership is transferred to
// the caller of this method; as with g_hash_table_steal().
//
// You can pass nil for @lookup_key, provided the hash and equal functions of
// @hash_table are nil-safe.
func HashTableStealExtended(hashTable *HashTable, lookupKey interface{}) (stolenKey interface{}, stolenValue interface{}, ok bool) {
	var arg1 *C.GHashTable
	var arg2 C.gpointer

	arg1 = (*C.GHashTable)(unsafe.Pointer(hashTable.Native()))
	arg2 = C.gpointer(lookupKey)

	var arg3 C.gpointer
	var stolenKey interface{}
	var arg4 C.gpointer
	var stolenValue interface{}
	var cret C.gboolean
	var ok bool

	cret = C.g_hash_table_steal_extended(arg1, arg2, &arg3, &arg4)

	stolenKey = interface{}(&arg3)
	stolenValue = interface{}(&arg4)
	if cret {
		ok = true
	}

	return stolenKey, stolenValue, ok
}

// HashTableUnref: atomically decrements the reference count of @hash_table by
// one. If the reference count drops to 0, all keys and values will be
// destroyed, and all memory allocated by the hash table is released. This
// function is MT-safe and may be called from any thread.
func HashTableUnref(hashTable *HashTable) {
	var arg1 *C.GHashTable

	arg1 = (*C.GHashTable)(unsafe.Pointer(hashTable.Native()))

	C.g_hash_table_unref(arg1)
}

// Int64Equal compares the two #gint64 values being pointed to and returns true
// if they are equal. It can be passed to g_hash_table_new() as the
// @key_equal_func parameter, when using non-nil pointers to 64-bit integers as
// keys in a Table.
func Int64Equal(v1 interface{}, v2 interface{}) bool {
	var arg1 C.gpointer
	var arg2 C.gpointer

	arg1 = C.gpointer(v1)
	arg2 = C.gpointer(v2)

	var cret C.gboolean
	var ok bool

	cret = C.g_int64_equal(arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// Int64Hash converts a pointer to a #gint64 to a hash value.
//
// It can be passed to g_hash_table_new() as the @hash_func parameter, when
// using non-nil pointers to 64-bit integer values as keys in a Table.
func Int64Hash(v interface{}) {
	var arg1 C.gpointer

	arg1 = C.gpointer(v)

	C.g_int64_hash(arg1)
}

// IntEqual compares the two #gint values being pointed to and returns true if
// they are equal. It can be passed to g_hash_table_new() as the @key_equal_func
// parameter, when using non-nil pointers to integers as keys in a Table.
//
// Note that this function acts on pointers to #gint, not on #gint directly: if
// your hash table's keys are of the form `GINT_TO_POINTER (n)`, use
// g_direct_equal() instead.
func IntEqual(v1 interface{}, v2 interface{}) bool {
	var arg1 C.gpointer
	var arg2 C.gpointer

	arg1 = C.gpointer(v1)
	arg2 = C.gpointer(v2)

	var cret C.gboolean
	var ok bool

	cret = C.g_int_equal(arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// IntHash converts a pointer to a #gint to a hash value. It can be passed to
// g_hash_table_new() as the @hash_func parameter, when using non-nil pointers
// to integer values as keys in a Table.
//
// Note that this function acts on pointers to #gint, not on #gint directly: if
// your hash table's keys are of the form `GINT_TO_POINTER (n)`, use
// g_direct_hash() instead.
func IntHash(v interface{}) {
	var arg1 C.gpointer

	arg1 = C.gpointer(v)

	C.g_int_hash(arg1)
}

// StrEqual compares two strings for byte-by-byte equality and returns true if
// they are equal. It can be passed to g_hash_table_new() as the @key_equal_func
// parameter, when using non-nil strings as keys in a Table.
//
// This function is typically used for hash table comparisons, but can be used
// for general purpose comparisons of non-nil strings. For a nil-safe string
// comparison function, see g_strcmp0().
func StrEqual(v1 interface{}, v2 interface{}) bool {
	var arg1 C.gpointer
	var arg2 C.gpointer

	arg1 = C.gpointer(v1)
	arg2 = C.gpointer(v2)

	var cret C.gboolean
	var ok bool

	cret = C.g_str_equal(arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// StrHash converts a string to a hash value.
//
// This function implements the widely used "djb" hash apparently posted by
// Daniel Bernstein to comp.lang.c some time ago. The 32 bit unsigned hash value
// starts at 5381 and for each byte 'c' in the string, is updated: `hash = hash
// * 33 + c`. This function uses the signed value of each byte.
//
// It can be passed to g_hash_table_new() as the @hash_func parameter, when
// using non-nil strings as keys in a Table.
//
// Note that this function may not be a perfect fit for all use cases. For
// example, it produces some hash collisions with strings as short as 2.
func StrHash(v interface{}) {
	var arg1 C.gpointer

	arg1 = C.gpointer(v)

	C.g_str_hash(arg1)
}

// HashTable: the Table struct is an opaque data structure to represent a [Hash
// Table][glib-Hash-Tables]. It should only be accessed via the following
// functions.
type HashTable struct {
	native C.GHashTable
}

// WrapHashTable wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapHashTable(ptr unsafe.Pointer) *HashTable {
	if ptr == nil {
		return nil
	}

	return (*HashTable)(ptr)
}

func marshalHashTable(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapHashTable(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (h *HashTable) Native() unsafe.Pointer {
	return unsafe.Pointer(&h.native)
}

// HashTableIter: a GHashTableIter structure represents an iterator that can be
// used to iterate over the elements of a Table. GHashTableIter structures are
// typically allocated on the stack and then initialized with
// g_hash_table_iter_init().
//
// The iteration order of a TableIter over the keys/values in a hash table is
// not defined.
type HashTableIter struct {
	native C.GHashTableIter
}

// WrapHashTableIter wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapHashTableIter(ptr unsafe.Pointer) *HashTableIter {
	if ptr == nil {
		return nil
	}

	return (*HashTableIter)(ptr)
}

func marshalHashTableIter(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapHashTableIter(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (h *HashTableIter) Native() unsafe.Pointer {
	return unsafe.Pointer(&h.native)
}

// HashTable returns the Table associated with @iter.
func (i *HashTableIter) HashTable(i *HashTableIter) {
	var arg0 *C.GHashTableIter

	arg0 = (*C.GHashTableIter)(unsafe.Pointer(i.Native()))

	C.g_hash_table_iter_get_hash_table(arg0)
}

// Init initializes a key/value pair iterator and associates it with
// @hash_table. Modifying the hash table after calling this function invalidates
// the returned iterator.
//
// The iteration order of a TableIter over the keys/values in a hash table is
// not defined.
//
//    GHashTableIter iter;
//    gpointer key, value;
//
//    g_hash_table_iter_init (&iter, hash_table);
//    while (g_hash_table_iter_next (&iter, &key, &value))
//      {
//        // do something with key and value
//      }
func (i *HashTableIter) Init(i *HashTableIter, hashTable *HashTable) {
	var arg0 *C.GHashTableIter
	var arg1 *C.GHashTable

	arg0 = (*C.GHashTableIter)(unsafe.Pointer(i.Native()))
	arg1 = (*C.GHashTable)(unsafe.Pointer(hashTable.Native()))

	C.g_hash_table_iter_init(arg0, arg1)
}

// Next advances @iter and retrieves the key and/or value that are now pointed
// to as a result of this advancement. If false is returned, @key and @value are
// not set, and the iterator becomes invalid.
func (i *HashTableIter) Next(i *HashTableIter) (key interface{}, value interface{}, ok bool) {
	var arg0 *C.GHashTableIter

	arg0 = (*C.GHashTableIter)(unsafe.Pointer(i.Native()))

	var arg1 C.gpointer
	var key interface{}
	var arg2 C.gpointer
	var value interface{}
	var cret C.gboolean
	var ok bool

	cret = C.g_hash_table_iter_next(arg0, &arg1, &arg2)

	key = interface{}(&arg1)
	value = interface{}(&arg2)
	if cret {
		ok = true
	}

	return key, value, ok
}

// Remove removes the key/value pair currently pointed to by the iterator from
// its associated Table. Can only be called after g_hash_table_iter_next()
// returned true, and cannot be called more than once for the same key/value
// pair.
//
// If the Table was created using g_hash_table_new_full(), the key and value are
// freed using the supplied destroy functions, otherwise you have to make sure
// that any dynamically allocated values are freed yourself.
//
// It is safe to continue iterating the Table afterward:
//
//    while (g_hash_table_iter_next (&iter, &key, &value))
//      {
//        if (condition)
//          g_hash_table_iter_remove (&iter);
//      }
func (i *HashTableIter) Remove(i *HashTableIter) {
	var arg0 *C.GHashTableIter

	arg0 = (*C.GHashTableIter)(unsafe.Pointer(i.Native()))

	C.g_hash_table_iter_remove(arg0)
}

// Replace replaces the value currently pointed to by the iterator from its
// associated Table. Can only be called after g_hash_table_iter_next() returned
// true.
//
// If you supplied a @value_destroy_func when creating the Table, the old value
// is freed using that function.
func (i *HashTableIter) Replace(i *HashTableIter, value interface{}) {
	var arg0 *C.GHashTableIter
	var arg1 C.gpointer

	arg0 = (*C.GHashTableIter)(unsafe.Pointer(i.Native()))
	arg1 = C.gpointer(value)

	C.g_hash_table_iter_replace(arg0, arg1)
}

// Steal removes the key/value pair currently pointed to by the iterator from
// its associated Table, without calling the key and value destroy functions.
// Can only be called after g_hash_table_iter_next() returned true, and cannot
// be called more than once for the same key/value pair.
func (i *HashTableIter) Steal(i *HashTableIter) {
	var arg0 *C.GHashTableIter

	arg0 = (*C.GHashTableIter)(unsafe.Pointer(i.Native()))

	C.g_hash_table_iter_steal(arg0)
}
