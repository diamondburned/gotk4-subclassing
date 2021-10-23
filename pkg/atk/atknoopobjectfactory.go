// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_no_op_object_factory_get_type()), F: marshalNoOpObjectFactorier},
	})
}

// NoOpObjectFactory: atkObjectFactory which creates an AtkNoOpObject. An
// instance of this is created by an AtkRegistry if no factory type has not been
// specified to create an accessible object of a particular type.
type NoOpObjectFactory struct {
	ObjectFactory
}

func wrapNoOpObjectFactory(obj *externglib.Object) *NoOpObjectFactory {
	return &NoOpObjectFactory{
		ObjectFactory: ObjectFactory{
			Object: obj,
		},
	}
}

func marshalNoOpObjectFactorier(p uintptr) (interface{}, error) {
	return wrapNoOpObjectFactory(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewNoOpObjectFactory creates an instance of an ObjectFactory which generates
// primitive (non-functioning) Objects.
func NewNoOpObjectFactory() *NoOpObjectFactory {
	var _cret *C.AtkObjectFactory // in

	_cret = C.atk_no_op_object_factory_new()

	var _noOpObjectFactory *NoOpObjectFactory // out

	_noOpObjectFactory = wrapNoOpObjectFactory(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _noOpObjectFactory
}
