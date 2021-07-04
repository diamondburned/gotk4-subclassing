// Code generated by girgen. DO NOT EDIT.

package gobject

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gobject-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_binding_flags_get_type()), F: marshalBindingFlags},
	})
}

// BindingFlags flags to be passed to g_object_bind_property() or
// g_object_bind_property_full().
//
// This enumeration can be extended at later date.
type BindingFlags int

const (
	// BindingFlagsDefault: the default binding; if the source property changes,
	// the target property is updated with its value.
	BindingFlagsDefault BindingFlags = 0b0
	// BindingFlagsBidirectional: bidirectional binding; if either the property
	// of the source or the property of the target changes, the other is
	// updated.
	BindingFlagsBidirectional BindingFlags = 0b1
	// BindingFlagsSyncCreate: synchronize the values of the source and target
	// properties when creating the binding; the direction of the
	// synchronization is always from the source to the target.
	BindingFlagsSyncCreate BindingFlags = 0b10
	// BindingFlagsInvertBoolean: if the two properties being bound are
	// booleans, setting one to true will result in the other being set to false
	// and vice versa. This flag will only work for boolean properties, and
	// cannot be used when passing custom transformation functions to
	// g_object_bind_property_full().
	BindingFlagsInvertBoolean BindingFlags = 0b100
)

func marshalBindingFlags(p uintptr) (interface{}, error) {
	return BindingFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}