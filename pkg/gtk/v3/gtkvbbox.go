// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_vbutton_box_get_type()), F: marshalVButtonBox},
	})
}

type VButtonBox interface {
	gextras.Objector

	privateVButtonBoxClass()
}

// VButtonBoxClass implements the VButtonBox interface.
type VButtonBoxClass struct {
	*externglib.Object
	ButtonBoxClass
	BuildableInterface
	OrientableInterface
}

var _ VButtonBox = (*VButtonBoxClass)(nil)

func wrapVButtonBox(obj *externglib.Object) VButtonBox {
	return &VButtonBoxClass{
		Object: obj,
		ButtonBoxClass: ButtonBoxClass{
			Object: obj,
			BoxClass: BoxClass{
				Object: obj,
				ContainerClass: ContainerClass{
					Object: obj,
					WidgetClass: WidgetClass{
						InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
						BuildableInterface: BuildableInterface{
							Object: obj,
						},
					},
					BuildableInterface: BuildableInterface{
						Object: obj,
					},
				},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
				OrientableInterface: OrientableInterface{
					Object: obj,
				},
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
			OrientableInterface: OrientableInterface{
				Object: obj,
			},
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		OrientableInterface: OrientableInterface{
			Object: obj,
		},
	}
}

func marshalVButtonBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapVButtonBox(obj), nil
}

// NewVButtonBox creates a new vertical button box.
//
// Deprecated: since version 3.2.
func NewVButtonBox() *VButtonBoxClass {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_vbutton_box_new()

	var _vButtonBox *VButtonBoxClass // out

	_vButtonBox = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*VButtonBoxClass)

	return _vButtonBox
}

func (*VButtonBoxClass) privateVButtonBoxClass() {}
