// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

// CompareDataFunc specifies the type of a comparison function used to compare
// two values. The function should return a negative integer if the first value
// comes before the second, 0 if they are equal, or a positive integer if the
// first value comes after the second.
type CompareDataFunc func(a interface{}, b interface{}) int

//export gotk4_CompareDataFunc
func gotk4_CompareDataFunc(arg0 C.gpointer, arg1 C.gpointer, arg2 C.gpointer) C.gint {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(CompareDataFunc)
	ret := fn(a, b, userData)

	cret = C.gint(ret)

	return cret
}

// Func specifies the type of functions passed to g_list_foreach() and
// g_slist_foreach().
type Func func(data interface{})

//export gotk4_Func
func gotk4_Func(arg0 C.gpointer, arg1 C.gpointer) {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(Func)
	fn(data, userData)
}

// HFunc specifies the type of the function passed to g_hash_table_foreach(). It
// is called with each key/value pair, together with the @user_data parameter
// which is passed to g_hash_table_foreach().
type HFunc func(key interface{}, value interface{})

//export gotk4_HFunc
func gotk4_HFunc(arg0 C.gpointer, arg1 C.gpointer, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(HFunc)
	fn(key, value, userData)
}

// TimeVal represents a precise time, with seconds and microseconds. Similar to
// the struct timeval returned by the gettimeofday() UNIX system call.
//
// GLib is attempting to unify around the use of 64-bit integers to represent
// microsecond-precision time. As such, this type will be removed from a future
// version of GLib. A consequence of using `glong` for `tv_sec` is that on
// 32-bit systems `GTimeVal` is subject to the year 2038 problem.
type TimeVal struct {
	native C.GTimeVal
}

// WrapTimeVal wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTimeVal(ptr unsafe.Pointer) *TimeVal {
	if ptr == nil {
		return nil
	}

	return (*TimeVal)(ptr)
}

func marshalTimeVal(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTimeVal(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TimeVal) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// TvSec gets the field inside the struct.
func (t *TimeVal) TvSec() int32 {
	var v int32
	v = int32(t.native.tv_sec)
	return v
}

// TvUsec gets the field inside the struct.
func (t *TimeVal) TvUsec() int32 {
	var v int32
	v = int32(t.native.tv_usec)
	return v
}

// Add adds the given number of microseconds to @time_. @microseconds can also
// be negative to decrease the value of @time_.
func (t *TimeVal) Add(t *TimeVal, microseconds int32) {
	var arg0 *C.GTimeVal
	var arg1 C.glong

	arg0 = (*C.GTimeVal)(unsafe.Pointer(t.Native()))
	arg1 = C.glong(microseconds)

	C.g_time_val_add(arg0, arg1)
}

// ToISO8601 converts @time_ into an RFC 3339 encoded string, relative to the
// Coordinated Universal Time (UTC). This is one of the many formats allowed by
// ISO 8601.
//
// ISO 8601 allows a large number of date/time formats, with or without
// punctuation and optional elements. The format returned by this function is a
// complete date and time, with optional punctuation included, the UTC time zone
// represented as "Z", and the @tv_usec part included if and only if it is
// nonzero, i.e. either "YYYY-MM-DDTHH:MM:SSZ" or "YYYY-MM-DDTHH:MM:SS.fffffZ".
//
// This corresponds to the Internet date/time format defined by RFC 3339
// (https://www.ietf.org/rfc/rfc3339.txt), and to either of the two most-precise
// formats defined by the W3C Note Date and Time Formats
// (http://www.w3.org/TR/NOTE-datetime-19980827). Both of these documents are
// profiles of ISO 8601.
//
// Use g_date_time_format() or g_strdup_printf() if a different variation of ISO
// 8601 format is required.
//
// If @time_ represents a date which is too large to fit into a `struct tm`, nil
// will be returned. This is platform dependent. Note also that since `GTimeVal`
// stores the number of seconds as a `glong`, on 32-bit systems it is subject to
// the year 2038 problem. Accordingly, since GLib 2.62, this function has been
// deprecated. Equivalent functionality is available using:
//
//    GDateTime *dt = g_date_time_new_from_unix_utc (time_val);
//    iso8601_string = g_date_time_format_iso8601 (dt);
//    g_date_time_unref (dt);
//
// The return value of g_time_val_to_iso8601() has been nullable since GLib
// 2.54; before then, GLib would crash under the same conditions.
func (t *TimeVal) ToISO8601(t *TimeVal) {
	var arg0 *C.GTimeVal

	arg0 = (*C.GTimeVal)(unsafe.Pointer(t.Native()))

	C.g_time_val_to_iso8601(arg0)
}