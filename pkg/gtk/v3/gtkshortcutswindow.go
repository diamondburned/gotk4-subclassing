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
// extern void _gotk4_gtk3_ShortcutsWindowClass_close(GtkShortcutsWindow*);
// extern void _gotk4_gtk3_ShortcutsWindowClass_search(GtkShortcutsWindow*);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcuts_window_get_type()), F: marshalShortcutsWindower},
	})
}

// ShortcutsWindowOverrider contains methods that are overridable.
type ShortcutsWindowOverrider interface {
	Close()
	Search()
}

// ShortcutsWindow shows brief information about the keyboard shortcuts and
// gestures of an application. The shortcuts can be grouped, and you can have
// multiple sections in this window, corresponding to the major modes of your
// application.
//
// Additionally, the shortcuts can be filtered by the current view, to avoid
// showing information that is not relevant in the current application context.
//
// The recommended way to construct a GtkShortcutsWindow is with GtkBuilder, by
// populating a ShortcutsWindow with one or more ShortcutsSection objects, which
// contain ShortcutsGroups that in turn contain objects of class
// ShortcutsShortcut.
//
// A simple example:
//
// ! (gedit-shortcuts.png)
//
// This example has as single section. As you can see, the shortcut groups are
// arranged in columns, and spread across several pages if there are too many to
// find on a single page.
//
// The .ui file for this example can be found here
// (https://git.gnome.org/browse/gtk+/tree/demos/gtk-demo/shortcuts-gedit.ui).
//
// An example with multiple views:
//
// ! (clocks-shortcuts.png)
//
// This example shows a ShortcutsWindow that has been configured to show only
// the shortcuts relevant to the "stopwatch" view.
//
// The .ui file for this example can be found here
// (https://git.gnome.org/browse/gtk+/tree/demos/gtk-demo/shortcuts-clocks.ui).
//
// An example with multiple sections:
//
// ! (builder-shortcuts.png)
//
// This example shows a ShortcutsWindow with two sections, "Editor Shortcuts"
// and "Terminal Shortcuts".
//
// The .ui file for this example can be found here
// (https://git.gnome.org/browse/gtk+/tree/demos/gtk-demo/shortcuts-builder.ui).
type ShortcutsWindow struct {
	_ [0]func() // equal guard
	Window
}

var (
	_ Binner = (*ShortcutsWindow)(nil)
)

func classInitShortcutsWindower(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

	goval := gbox.Get(uintptr(data))
	pclass := (*C.GtkShortcutsWindowClass)(unsafe.Pointer(gclassPtr))
	// gclass := (*C.GTypeClass)(unsafe.Pointer(gclassPtr))
	// pclass := (*C.GtkShortcutsWindowClass)(unsafe.Pointer(C.g_type_class_peek_parent(gclass)))

	if _, ok := goval.(interface{ Close() }); ok {
		pclass.close = (*[0]byte)(C._gotk4_gtk3_ShortcutsWindowClass_close)
	}

	if _, ok := goval.(interface{ Search() }); ok {
		pclass.search = (*[0]byte)(C._gotk4_gtk3_ShortcutsWindowClass_search)
	}
}

//export _gotk4_gtk3_ShortcutsWindowClass_close
func _gotk4_gtk3_ShortcutsWindowClass_close(arg0 *C.GtkShortcutsWindow) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Close() })

	iface.Close()
}

//export _gotk4_gtk3_ShortcutsWindowClass_search
func _gotk4_gtk3_ShortcutsWindowClass_search(arg0 *C.GtkShortcutsWindow) {
	goval := externglib.GoPrivateFromObject(unsafe.Pointer(arg0))
	iface := goval.(interface{ Search() })

	iface.Search()
}

func wrapShortcutsWindow(obj *externglib.Object) *ShortcutsWindow {
	return &ShortcutsWindow{
		Window: Window{
			Bin: Bin{
				Container: Container{
					Widget: Widget{
						InitiallyUnowned: externglib.InitiallyUnowned{
							Object: obj,
						},
						Object: obj,
						ImplementorIface: atk.ImplementorIface{
							Object: obj,
						},
						Buildable: Buildable{
							Object: obj,
						},
					},
				},
			},
		},
	}
}

func marshalShortcutsWindower(p uintptr) (interface{}, error) {
	return wrapShortcutsWindow(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectClose signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted when the user uses a keybinding to close the window.
//
// The default binding for this signal is the Escape key.
func (self *ShortcutsWindow) ConnectClose(f func()) externglib.SignalHandle {
	return self.Connect("close", externglib.GeneratedClosure{Func: f})
}

// ConnectSearch signal is a [keybinding signal][GtkBindingSignal] which gets
// emitted when the user uses a keybinding to start a search.
//
// The default binding for this signal is Control-F.
func (self *ShortcutsWindow) ConnectSearch(f func()) externglib.SignalHandle {
	return self.Connect("search", externglib.GeneratedClosure{Func: f})
}
