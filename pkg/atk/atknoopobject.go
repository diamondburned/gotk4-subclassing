// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_no_op_object_get_type()), F: marshalNoOpObjector},
	})
}

// NoOpObjectOverrider contains methods that are overridable.
type NoOpObjectOverrider interface {
}

// NoOpObject is an AtkObject which purports to implement all ATK interfaces. It
// is the type of AtkObject which is created if an accessible object is
// requested for an object type for which no factory type is specified.
type NoOpObject struct {
	_ [0]func() // equal guard
	ObjectClass

	*externglib.Object
	Action
	Component
	Document
	EditableText
	Hypertext
	Image
	Selection
	Table
	TableCell
	Text
	Value
	Window
}

var (
	_ externglib.Objector = (*NoOpObject)(nil)
)

func classInitNoOpObjector(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapNoOpObject(obj *externglib.Object) *NoOpObject {
	return &NoOpObject{
		ObjectClass: ObjectClass{
			Object: obj,
		},
		Object: obj,
		Action: Action{
			Object: obj,
		},
		Component: Component{
			Object: obj,
		},
		Document: Document{
			Object: obj,
		},
		EditableText: EditableText{
			Object: obj,
		},
		Hypertext: Hypertext{
			Object: obj,
		},
		Image: Image{
			Object: obj,
		},
		Selection: Selection{
			Object: obj,
		},
		Table: Table{
			Object: obj,
		},
		TableCell: TableCell{
			ObjectClass: ObjectClass{
				Object: obj,
			},
		},
		Text: Text{
			Object: obj,
		},
		Value: Value{
			Object: obj,
		},
		Window: Window{
			ObjectClass: ObjectClass{
				Object: obj,
			},
		},
	}
}

func marshalNoOpObjector(p uintptr) (interface{}, error) {
	return wrapNoOpObject(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewNoOpObject provides a default (non-functioning stub) Object. Application
// maintainers should not use this method.
//
// The function takes the following parameters:
//
//    - obj: #GObject.
//
// The function returns the following values:
//
//    - noOpObject: default (non-functioning stub) Object.
//
func NewNoOpObject(obj *externglib.Object) *NoOpObject {
	var _arg1 *C.GObject   // out
	var _cret *C.AtkObject // in

	_arg1 = (*C.GObject)(unsafe.Pointer(obj.Native()))

	_cret = C.atk_no_op_object_new(_arg1)
	runtime.KeepAlive(obj)

	var _noOpObject *NoOpObject // out

	_noOpObject = wrapNoOpObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _noOpObject
}
