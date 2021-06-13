// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcut_get_type()), F: marshalShortcut},
	})
}

// Shortcut: a `GtkShortcut` describes a keyboard shortcut.
//
// It contains a description of how to trigger the shortcut via a
// [class@Gtk.ShortcutTrigger] and a way to activate the shortcut on a widget
// via a [class@Gtk.ShortcutAction].
//
// The actual work is usually done via [class@Gtk.ShortcutController], which
// decides if and when to activate a shortcut. Using that controller directly
// however is rarely necessary as various higher level convenience APIs exist on
// Widgets that make it easier to use shortcuts in GTK.
//
// `GtkShortcut` does provide functionality to make it easy for users to work
// with shortcuts, either by providing informational strings for display
// purposes or by allowing shortcuts to be configured.
type Shortcut interface {
	gextras.Objector

	// SetAction sets the new action for @self to be @action.
	SetAction(action ShortcutAction)
	// SetArguments sets the arguments to pass when activating the shortcut.
	SetArguments(args *glib.Variant)
	// SetTrigger sets the new trigger for @self to be @trigger.
	SetTrigger(trigger ShortcutTrigger)
}

// shortcut implements the Shortcut interface.
type shortcut struct {
	gextras.Objector
}

var _ Shortcut = (*shortcut)(nil)

// WrapShortcut wraps a GObject to the right type. It is
// primarily used internally.
func WrapShortcut(obj *externglib.Object) Shortcut {
	return Shortcut{
		Objector: obj,
	}
}

func marshalShortcut(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapShortcut(obj), nil
}

// SetAction sets the new action for @self to be @action.
func (s shortcut) SetAction(action ShortcutAction) {
	var _arg0 *C.GtkShortcut       // out
	var _arg1 *C.GtkShortcutAction // out

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkShortcutAction)(unsafe.Pointer(action.Native()))

	C.gtk_shortcut_set_action(_arg0, _arg1)
}

// SetArguments sets the arguments to pass when activating the shortcut.
func (s shortcut) SetArguments(args *glib.Variant) {
	var _arg0 *C.GtkShortcut // out
	var _arg1 *C.GVariant    // out

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GVariant)(unsafe.Pointer(args.Native()))

	C.gtk_shortcut_set_arguments(_arg0, _arg1)
}

// SetTrigger sets the new trigger for @self to be @trigger.
func (s shortcut) SetTrigger(trigger ShortcutTrigger) {
	var _arg0 *C.GtkShortcut        // out
	var _arg1 *C.GtkShortcutTrigger // out

	_arg0 = (*C.GtkShortcut)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkShortcutTrigger)(unsafe.Pointer(trigger.Native()))

	C.gtk_shortcut_set_trigger(_arg0, _arg1)
}
