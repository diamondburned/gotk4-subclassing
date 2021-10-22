// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_vpaned_get_type()), F: marshalVPanedder},
	})
}

// VPaned widget is a container widget with two children arranged vertically.
// The division between the two panes is adjustable by the user by dragging a
// handle. See Paned for details.
//
// GtkVPaned has been deprecated, use Paned instead.
type VPaned struct {
	Paned
}

func wrapVPaned(obj *externglib.Object) *VPaned {
	return &VPaned{
		Paned: Paned{
			Container: Container{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					ImplementorIface: atk.ImplementorIface{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
					Object: obj,
				},
			},
			Orientable: Orientable{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalVPanedder(p uintptr) (interface{}, error) {
	return wrapVPaned(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewVPaned: create a new VPaned
//
// Deprecated: Use gtk_paned_new() with GTK_ORIENTATION_VERTICAL instead.
func NewVPaned() *VPaned {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_vpaned_new()

	var _vPaned *VPaned // out

	_vPaned = wrapVPaned(externglib.Take(unsafe.Pointer(_cret)))

	return _vPaned
}
