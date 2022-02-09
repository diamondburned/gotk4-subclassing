// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcuts_section_get_type()), F: marshalShortcutsSectioner},
	})
}

// ShortcutsSection: GtkShortcutsSection collects all the keyboard shortcuts and
// gestures for a major application mode.
//
// If your application needs multiple sections, you should give each section a
// unique gtk.ShortcutsSection:section-name and a gtk.ShortcutsSection:title
// that can be shown in the section selector of the gtk.ShortcutsWindow.
//
// The gtk.ShortcutsSection:max-height property can be used to influence how the
// groups in the section are distributed over pages and columns.
//
// This widget is only meant to be used with gtk.ShortcutsWindow.
type ShortcutsSection struct {
	_ [0]func() // equal guard
	Box
}

var (
	_ Widgetter           = (*ShortcutsSection)(nil)
	_ externglib.Objector = (*ShortcutsSection)(nil)
)

func wrapShortcutsSection(obj *externglib.Object) *ShortcutsSection {
	return &ShortcutsSection{
		Box: Box{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Object: obj,
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
			},
			Object: obj,
			Orientable: Orientable{
				Object: obj,
			},
		},
	}
}

func marshalShortcutsSectioner(p uintptr) (interface{}, error) {
	return wrapShortcutsSection(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *ShortcutsSection) ConnectChangeCurrentPage(f func(object int) bool) externglib.SignalHandle {
	return v.Connect("change-current-page", externglib.GeneratedClosure{Func: f})
}
