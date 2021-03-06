// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
// extern void _gotk4_gtk3_CellAccessibleClass_update_cache(GtkCellAccessible*, gboolean);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_accessible_get_type()), F: marshalCellAccessibler},
	})
}

// CellAccessibleOverrider contains methods that are overridable.
type CellAccessibleOverrider interface {
	// The function takes the following parameters:
	//
	UpdateCache(emitSignal bool)
}

type CellAccessible struct {
	_ [0]func() // equal guard
	Accessible

	*externglib.Object
	atk.Action
	atk.Component
	atk.ObjectClass
	atk.TableCell
}

var (
	_ externglib.Objector = (*CellAccessible)(nil)
)

func classInitCellAccessibler(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkCellAccessibleClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkCellAccessibleClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ UpdateCache(emitSignal bool) }); ok {
		pclass.update_cache = (*[0]byte)(C._gotk4_gtk3_CellAccessibleClass_update_cache)
	}
}

//export _gotk4_gtk3_CellAccessibleClass_update_cache
func _gotk4_gtk3_CellAccessibleClass_update_cache(arg0 *C.GtkCellAccessible, arg1 C.gboolean) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ UpdateCache(emitSignal bool) })

	var _emitSignal bool // out

	if arg1 != 0 {
		_emitSignal = true
	}

	iface.UpdateCache(_emitSignal)
}

func wrapCellAccessible(obj *externglib.Object) *CellAccessible {
	return &CellAccessible{
		Accessible: Accessible{
			ObjectClass: atk.ObjectClass{
				Object: obj,
			},
		},
		Object: obj,
		Action: atk.Action{
			Object: obj,
		},
		Component: atk.Component{
			Object: obj,
		},
		ObjectClass: atk.ObjectClass{
			Object: obj,
		},
		TableCell: atk.TableCell{
			ObjectClass: atk.ObjectClass{
				Object: obj,
			},
		},
	}
}

func marshalCellAccessibler(p uintptr) (interface{}, error) {
	return wrapCellAccessible(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
