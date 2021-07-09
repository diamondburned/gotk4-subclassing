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
		{T: externglib.Type(C.gtk_hbox_get_type()), F: marshalHBox},
	})
}

// HBox is a container that organizes child widgets into a single row.
//
// Use the Box packing interface to determine the arrangement, spacing, width,
// and alignment of HBox children.
//
// All children are allocated the same height.
//
// GtkHBox has been deprecated. You can use Box instead, which is a very quick
// and easy change. If you have derived your own classes from GtkHBox, you can
// simply change the inheritance to derive directly from Box. No further changes
// are needed, since the default value of the Orientable:orientation property is
// GTK_ORIENTATION_HORIZONTAL.
//
// If you have a grid-like layout composed of nested boxes, and you don’t need
// first-child or last-child styling, the recommendation is to switch to Grid.
// For more information about migrating to Grid, see [Migrating from other
// containers to GtkGrid][gtk-migrating-GtkGrid].
type HBox interface {
	gextras.Objector

	privateHBoxClass()
}

// HBoxClass implements the HBox interface.
type HBoxClass struct {
	*externglib.Object
	BoxClass
	BuildableInterface
	OrientableInterface
}

var _ HBox = (*HBoxClass)(nil)

func wrapHBox(obj *externglib.Object) HBox {
	return &HBoxClass{
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
	}
}

func marshalHBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapHBox(obj), nil
}

// NewHBox creates a new HBox.
//
// Deprecated: since version 3.2.
func NewHBox(homogeneous bool, spacing int) *HBoxClass {
	var _arg1 C.gboolean   // out
	var _arg2 C.gint       // out
	var _cret *C.GtkWidget // in

	if homogeneous {
		_arg1 = C.TRUE
	}
	_arg2 = C.gint(spacing)

	_cret = C.gtk_hbox_new(_arg1, _arg2)

	var _hBox *HBoxClass // out

	_hBox = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*HBoxClass)

	return _hBox
}

func (*HBoxClass) privateHBoxClass() {}
