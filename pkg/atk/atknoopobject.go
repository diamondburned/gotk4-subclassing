// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_no_op_object_get_type()), F: marshalNoOpObject},
	})
}

// NoOpObject is an AtkObject which purports to implement all ATK interfaces. It
// is the type of AtkObject which is created if an accessible object is
// requested for an object type for which no factory type is specified.
type NoOpObject interface {
	gextras.Objector

	privateNoOpObjectClass()
}

// NoOpObjectClass implements the NoOpObject interface.
type NoOpObjectClass struct {
	*externglib.Object
	ObjectClass
	ActionInterface
	ComponentInterface
	DocumentInterface
	EditableTextInterface
	HypertextInterface
	ImageInterface
	SelectionInterface
	TableInterface
	TableCellInterface
	TextInterface
	ValueInterface
	WindowInterface
}

var _ NoOpObject = (*NoOpObjectClass)(nil)

func wrapNoOpObject(obj *externglib.Object) NoOpObject {
	return &NoOpObjectClass{
		Object: obj,
		ObjectClass: ObjectClass{
			Object: obj,
		},
		ActionInterface: ActionInterface{
			Object: obj,
		},
		ComponentInterface: ComponentInterface{
			Object: obj,
		},
		DocumentInterface: DocumentInterface{
			Object: obj,
		},
		EditableTextInterface: EditableTextInterface{
			Object: obj,
		},
		HypertextInterface: HypertextInterface{
			Object: obj,
		},
		ImageInterface: ImageInterface{
			Object: obj,
		},
		SelectionInterface: SelectionInterface{
			Object: obj,
		},
		TableInterface: TableInterface{
			Object: obj,
		},
		TableCellInterface: TableCellInterface{
			ObjectClass: ObjectClass{
				Object: obj,
			},
		},
		TextInterface: TextInterface{
			Object: obj,
		},
		ValueInterface: ValueInterface{
			Object: obj,
		},
		WindowInterface: WindowInterface{
			ObjectClass: ObjectClass{
				Object: obj,
			},
		},
	}
}

func marshalNoOpObject(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapNoOpObject(obj), nil
}

// NewNoOpObject provides a default (non-functioning stub) Object. Application
// maintainers should not use this method.
func NewNoOpObject(obj gextras.Objector) *NoOpObjectClass {
	var _arg1 *C.GObject   // out
	var _cret *C.AtkObject // in

	_arg1 = (*C.GObject)(unsafe.Pointer(obj.Native()))

	_cret = C.atk_no_op_object_new(_arg1)

	var _noOpObject *NoOpObjectClass // out

	_noOpObject = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*NoOpObjectClass)

	return _noOpObject
}

func (*NoOpObjectClass) privateNoOpObjectClass() {}
