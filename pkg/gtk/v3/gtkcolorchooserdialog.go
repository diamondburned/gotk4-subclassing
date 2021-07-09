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
		{T: externglib.Type(C.gtk_color_chooser_dialog_get_type()), F: marshalColorChooserDialog},
	})
}

// ColorChooserDialog: the ColorChooserDialog widget is a dialog for choosing a
// color. It implements the ColorChooser interface.
type ColorChooserDialog interface {
	gextras.Objector

	privateColorChooserDialogClass()
}

// ColorChooserDialogClass implements the ColorChooserDialog interface.
type ColorChooserDialogClass struct {
	*externglib.Object
	DialogClass
	BuildableInterface
	ColorChooserInterface
}

var _ ColorChooserDialog = (*ColorChooserDialogClass)(nil)

func wrapColorChooserDialog(obj *externglib.Object) ColorChooserDialog {
	return &ColorChooserDialogClass{
		Object: obj,
		DialogClass: DialogClass{
			Object: obj,
			WindowClass: WindowClass{
				Object: obj,
				BinClass: BinClass{
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
				},
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
		ColorChooserInterface: ColorChooserInterface{
			Object: obj,
		},
	}
}

func marshalColorChooserDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapColorChooserDialog(obj), nil
}

// NewColorChooserDialog creates a new ColorChooserDialog.
func NewColorChooserDialog(title string, parent Window) *ColorChooserDialogClass {
	var _arg1 *C.gchar     // out
	var _arg2 *C.GtkWindow // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	_cret = C.gtk_color_chooser_dialog_new(_arg1, _arg2)

	var _colorChooserDialog *ColorChooserDialogClass // out

	_colorChooserDialog = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ColorChooserDialogClass)

	return _colorChooserDialog
}

func (*ColorChooserDialogClass) privateColorChooserDialogClass() {}
