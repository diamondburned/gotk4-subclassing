// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_recent_chooser_widget_get_type()), F: marshalRecentChooserWidget},
	})
}

// RecentChooserWidget is a widget suitable for selecting recently used files.
// It is the main building block of a RecentChooserDialog. Most applications
// will only need to use the latter; you can use RecentChooserWidget as part of
// a larger window if you have special needs.
//
// Note that RecentChooserWidget does not have any methods of its own. Instead,
// you should use the functions that work on a RecentChooser.
//
// Recently used files are supported since GTK+ 2.10.
type RecentChooserWidget interface {
	Box
	Buildable
	Orientable
	RecentChooser
}

// recentChooserWidget implements the RecentChooserWidget interface.
type recentChooserWidget struct {
	Box
	Buildable
	Orientable
	RecentChooser
}

var _ RecentChooserWidget = (*recentChooserWidget)(nil)

// WrapRecentChooserWidget wraps a GObject to the right type. It is
// primarily used internally.
func WrapRecentChooserWidget(obj *externglib.Object) RecentChooserWidget {
	return RecentChooserWidget{
		Box:           WrapBox(obj),
		Buildable:     WrapBuildable(obj),
		Orientable:    WrapOrientable(obj),
		RecentChooser: WrapRecentChooser(obj),
	}
}

func marshalRecentChooserWidget(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRecentChooserWidget(obj), nil
}
