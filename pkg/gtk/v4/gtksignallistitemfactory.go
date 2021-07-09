// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_signal_list_item_factory_get_type()), F: marshalSignalListItemFactory},
	})
}

// SignalListItemFactory: `GtkSignalListItemFactory` is a `GtkListItemFactory`
// that emits signals to to manage listitems.
//
// Signals are emitted for every listitem in the same order:
//
//    1. [signal@Gtk.SignalListItemFactory::setup] is emitted to set up permanent
//    things on the listitem. This usually means constructing the widgets used in
//    the row and adding them to the listitem.
//
//    2. [signal@Gtk.SignalListItemFactory::bind] is emitted to bind the item passed
//    via [property@Gtk.ListItem:item] to the widgets that have been created in
//    step 1 or to add item-specific widgets. Signals are connected to listen to
//    changes - both to changes in the item to update the widgets or to changes
//    in the widgets to update the item. After this signal has been called, the
//    listitem may be shown in a list widget.
//
//    3. [signal@Gtk.SignalListItemFactory::unbind] is emitted to undo everything
//    done in step 2. Usually this means disconnecting signal handlers. Once this
//    signal has been called, the listitem will no longer be used in a list
//    widget.
//
//    4. [signal@Gtk.SignalListItemFactory::bind] and
//    [signal@Gtk.SignalListItemFactory::unbind] may be emitted multiple times
//    again to bind the listitem for use with new items. By reusing listitems,
//    potentially costly setup can be avoided. However, it means code needs to
//    make sure to properly clean up the listitem in step 3 so that no information
//    from the previous use leaks into the next use.
//
// 5. [signal@Gtk.SignalListItemFactory::teardown] is emitted to allow undoing
// the effects of [signal@Gtk.SignalListItemFactory::setup]. After this signal
// was emitted on a listitem, the listitem will be destroyed and not be used
// again.
//
// Note that during the signal emissions, changing properties on the ListItems
// passed will not trigger notify signals as the listitem's notifications are
// frozen. See g_object_freeze_notify() for details.
//
// For tracking changes in other properties in the `GtkListItem`, the ::notify
// signal is recommended. The signal can be connected in the
// [signal@Gtk.SignalListItemFactory::setup] signal and removed again during
// [signal@Gtk.SignalListItemFactory::teardown].
type SignalListItemFactory interface {
	gextras.Objector

	privateSignalListItemFactoryClass()
}

// SignalListItemFactoryClass implements the SignalListItemFactory interface.
type SignalListItemFactoryClass struct {
	ListItemFactoryClass
}

var _ SignalListItemFactory = (*SignalListItemFactoryClass)(nil)

func wrapSignalListItemFactory(obj *externglib.Object) SignalListItemFactory {
	return &SignalListItemFactoryClass{
		ListItemFactoryClass: ListItemFactoryClass{
			Object: obj,
		},
	}
}

func marshalSignalListItemFactory(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSignalListItemFactory(obj), nil
}

// NewSignalListItemFactory creates a new `GtkSignalListItemFactory`.
//
// You need to connect signal handlers before you use it.
func NewSignalListItemFactory() *SignalListItemFactoryClass {
	var _cret *C.GtkListItemFactory // in

	_cret = C.gtk_signal_list_item_factory_new()

	var _signalListItemFactory *SignalListItemFactoryClass // out

	_signalListItemFactory = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*SignalListItemFactoryClass)

	return _signalListItemFactory
}

func (*SignalListItemFactoryClass) privateSignalListItemFactoryClass() {}
