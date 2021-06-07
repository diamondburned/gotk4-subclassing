// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

// KeyvalConvertCase obtains the upper- and lower-case versions of the keyval
// @symbol. Examples of keyvals are K_KEY_a, K_KEY_Enter, K_KEY_F1, etc.
func KeyvalConvertCase(symbol uint) (lower uint, upper uint) {
	var arg1 C.guint

	arg1 = C.guint(symbol)

	var arg2 C.guint
	var lower uint
	var arg3 C.guint
	var upper uint

	C.gdk_keyval_convert_case(arg1, &arg2, &arg3)

	lower = uint(&arg2)
	upper = uint(&arg3)

	return lower, upper
}

// KeyvalFromName converts a key name to a key value.
//
// The names are the same as those in the `gdk/gdkkeysyms.h` header file but
// without the leading “GDK_KEY_”.
func KeyvalFromName(keyvalName string) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(keyvalName))
	defer C.free(unsafe.Pointer(arg1))

	C.gdk_keyval_from_name(arg1)
}

// KeyvalIsLower returns true if the given key value is in lower case.
func KeyvalIsLower(keyval uint) bool {
	var arg1 C.guint

	arg1 = C.guint(keyval)

	var cret C.gboolean
	var ok bool

	cret = C.gdk_keyval_is_lower(arg1)

	if cret {
		ok = true
	}

	return ok
}

// KeyvalIsUpper returns true if the given key value is in upper case.
func KeyvalIsUpper(keyval uint) bool {
	var arg1 C.guint

	arg1 = C.guint(keyval)

	var cret C.gboolean
	var ok bool

	cret = C.gdk_keyval_is_upper(arg1)

	if cret {
		ok = true
	}

	return ok
}

// KeyvalName converts a key value into a symbolic name.
//
// The names are the same as those in the `gdk/gdkkeysyms.h` header file but
// without the leading “GDK_KEY_”.
func KeyvalName(keyval uint) {
	var arg1 C.guint

	arg1 = C.guint(keyval)

	C.gdk_keyval_name(arg1)
}

// KeyvalToLower converts a key value to lower case, if applicable.
func KeyvalToLower(keyval uint) {
	var arg1 C.guint

	arg1 = C.guint(keyval)

	C.gdk_keyval_to_lower(arg1)
}

// KeyvalToUnicode: convert from a GDK key symbol to the corresponding ISO10646
// (Unicode) character.
func KeyvalToUnicode(keyval uint) {
	var arg1 C.guint

	arg1 = C.guint(keyval)

	C.gdk_keyval_to_unicode(arg1)
}

// KeyvalToUpper converts a key value to upper case, if applicable.
func KeyvalToUpper(keyval uint) {
	var arg1 C.guint

	arg1 = C.guint(keyval)

	C.gdk_keyval_to_upper(arg1)
}

// UnicodeToKeyval: convert from a ISO10646 character to a key symbol.
func UnicodeToKeyval(wc uint32) {
	var arg1 C.guint32

	arg1 = C.guint32(wc)

	C.gdk_unicode_to_keyval(arg1)
}

// KeymapKey: a KeymapKey is a hardware key that can be mapped to a keyval.
type KeymapKey struct {
	native C.GdkKeymapKey
}

// WrapKeymapKey wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapKeymapKey(ptr unsafe.Pointer) *KeymapKey {
	if ptr == nil {
		return nil
	}

	return (*KeymapKey)(ptr)
}

func marshalKeymapKey(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapKeymapKey(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (k *KeymapKey) Native() unsafe.Pointer {
	return unsafe.Pointer(&k.native)
}

// Keycode gets the field inside the struct.
func (k *KeymapKey) Keycode() uint {
	var v uint
	v = uint(k.native.keycode)
	return v
}

// Group gets the field inside the struct.
func (k *KeymapKey) Group() int {
	var v int
	v = int(k.native.group)
	return v
}

// Level gets the field inside the struct.
func (k *KeymapKey) Level() int {
	var v int
	v = int(k.native.level)
	return v
}
