// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

// PatternMatch matches a string against a compiled pattern. Passing the correct
// length of the string given is mandatory. The reversed string can be omitted
// by passing nil, this is more efficient if the reversed version of the string
// to be matched is not at hand, as g_pattern_match() will only construct it if
// the compiled pattern requires reverse matches.
//
// Note that, if the user code will (possibly) match a string against a
// multitude of patterns containing wildcards, chances are high that some
// patterns will require a reversed string. In this case, it's more efficient to
// provide the reversed string to avoid multiple constructions thereof in the
// various calls to g_pattern_match().
//
// Note also that the reverse of a UTF-8 encoded string can in general not be
// obtained by g_strreverse(). This works only if the string does not contain
// any multibyte characters. GLib offers the g_utf8_strreverse() function to
// reverse UTF-8 encoded strings.
func PatternMatch(pspec *PatternSpec, stringLength uint, string string, stringReversed string) bool {
	var arg1 *C.GPatternSpec
	var arg2 C.guint
	var arg3 *C.gchar
	var arg4 *C.gchar

	arg1 = (*C.GPatternSpec)(unsafe.Pointer(pspec.Native()))
	arg2 = C.guint(stringLength)
	arg3 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = (*C.gchar)(C.CString(stringReversed))
	defer C.free(unsafe.Pointer(arg4))

	var cret C.gboolean

	cret = C.g_pattern_match(arg1, arg2, arg3, arg4)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// PatternMatchSimple matches a string against a pattern given as a string. If
// this function is to be called in a loop, it's more efficient to compile the
// pattern once with g_pattern_spec_new() and call g_pattern_match_string()
// repeatedly.
func PatternMatchSimple(pattern string, string string) bool {
	var arg1 *C.gchar
	var arg2 *C.gchar

	arg1 = (*C.gchar)(C.CString(pattern))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg2))

	var cret C.gboolean

	cret = C.g_pattern_match_simple(arg1, arg2)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// PatternMatchString matches a string against a compiled pattern. If the string
// is to be matched against more than one pattern, consider using
// g_pattern_match() instead while supplying the reversed string.
func PatternMatchString(pspec *PatternSpec, string string) bool {
	var arg1 *C.GPatternSpec
	var arg2 *C.gchar

	arg1 = (*C.GPatternSpec)(unsafe.Pointer(pspec.Native()))
	arg2 = (*C.gchar)(C.CString(string))
	defer C.free(unsafe.Pointer(arg2))

	var cret C.gboolean

	cret = C.g_pattern_match_string(arg1, arg2)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// PatternSpec: a GPatternSpec struct is the 'compiled' form of a pattern. This
// structure is opaque and its fields cannot be accessed directly.
type PatternSpec struct {
	native C.GPatternSpec
}

// WrapPatternSpec wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPatternSpec(ptr unsafe.Pointer) *PatternSpec {
	if ptr == nil {
		return nil
	}

	return (*PatternSpec)(ptr)
}

func marshalPatternSpec(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapPatternSpec(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (p *PatternSpec) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}

// Equal compares two compiled pattern specs and returns whether they will match
// the same set of strings.
func (p *PatternSpec) Equal(pspec2 *PatternSpec) bool {
	var arg0 *C.GPatternSpec
	var arg1 *C.GPatternSpec

	arg0 = (*C.GPatternSpec)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GPatternSpec)(unsafe.Pointer(pspec2.Native()))

	var cret C.gboolean

	cret = C.g_pattern_spec_equal(arg0, arg1)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// Free frees the memory allocated for the Spec.
func (p *PatternSpec) Free() {
	var arg0 *C.GPatternSpec

	arg0 = (*C.GPatternSpec)(unsafe.Pointer(p.Native()))

	C.g_pattern_spec_free(arg0)
}
