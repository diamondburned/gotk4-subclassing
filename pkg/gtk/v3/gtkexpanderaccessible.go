// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_expander_accessible_get_type()), F: marshalExpanderAccessibler},
	})
}

// ExpanderAccessibleOverrider contains methods that are overridable.
type ExpanderAccessibleOverrider interface {
}

type ExpanderAccessible struct {
	_ [0]func() // equal guard
	ContainerAccessible

	atk.Action
}

var (
	_ externglib.Objector = (*ExpanderAccessible)(nil)
)

func classInitExpanderAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapExpanderAccessible(obj *externglib.Object) *ExpanderAccessible {
	return &ExpanderAccessible{
		ContainerAccessible: ContainerAccessible{
			WidgetAccessible: WidgetAccessible{
				Accessible: Accessible{
					ObjectClass: atk.ObjectClass{
						Object: obj,
					},
				},
				Component: atk.Component{
					Object: obj,
				},
			},
		},
		Action: atk.Action{
			Object: obj,
		},
	}
}

func marshalExpanderAccessibler(p uintptr) (interface{}, error) {
	return wrapExpanderAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
