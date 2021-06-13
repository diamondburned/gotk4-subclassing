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
		{T: externglib.Type(C.gtk_event_box_get_type()), F: marshalEventBox},
	})
}

// EventBox: the EventBox widget is a subclass of Bin which also has its own
// window. It is useful since it allows you to catch events for widgets which do
// not have their own window.
type EventBox interface {
	Bin
	Buildable

	// AboveChild returns whether the event box window is above or below the
	// windows of its child. See gtk_event_box_set_above_child() for details.
	AboveChild() bool
	// VisibleWindow returns whether the event box has a visible window. See
	// gtk_event_box_set_visible_window() for details.
	VisibleWindow() bool
	// SetAboveChild: set whether the event box window is positioned above the
	// windows of its child, as opposed to below it. If the window is above, all
	// events inside the event box will go to the event box. If the window is
	// below, events in windows of child widgets will first got to that widget,
	// and then to its parents.
	//
	// The default is to keep the window below the child.
	SetAboveChild(aboveChild bool)
	// SetVisibleWindow: set whether the event box uses a visible or invisible
	// child window. The default is to use visible windows.
	//
	// In an invisible window event box, the window that the event box creates
	// is a GDK_INPUT_ONLY window, which means that it is invisible and only
	// serves to receive events.
	//
	// A visible window event box creates a visible (GDK_INPUT_OUTPUT) window
	// that acts as the parent window for all the widgets contained in the event
	// box.
	//
	// You should generally make your event box invisible if you just want to
	// trap events. Creating a visible window may cause artifacts that are
	// visible to the user, especially if the user is using a theme with
	// gradients or pixmaps.
	//
	// The main reason to create a non input-only event box is if you want to
	// set the background to a different color or draw on it.
	//
	// There is one unexpected issue for an invisible event box that has its
	// window below the child. (See gtk_event_box_set_above_child().) Since the
	// input-only window is not an ancestor window of any windows that
	// descendent widgets of the event box create, events on these windows
	// aren’t propagated up by the windowing system, but only by GTK+. The
	// practical effect of this is if an event isn’t in the event mask for the
	// descendant window (see gtk_widget_add_events()), it won’t be received by
	// the event box.
	//
	// This problem doesn’t occur for visible event boxes, because in that case,
	// the event box window is actually the ancestor of the descendant windows,
	// not just at the same place on the screen.
	SetVisibleWindow(visibleWindow bool)
}

// eventBox implements the EventBox interface.
type eventBox struct {
	Bin
	Buildable
}

var _ EventBox = (*eventBox)(nil)

// WrapEventBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapEventBox(obj *externglib.Object) EventBox {
	return EventBox{
		Bin:       WrapBin(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalEventBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEventBox(obj), nil
}

// AboveChild returns whether the event box window is above or below the
// windows of its child. See gtk_event_box_set_above_child() for details.
func (e eventBox) AboveChild() bool {
	var _arg0 *C.GtkEventBox // out

	_arg0 = (*C.GtkEventBox)(unsafe.Pointer(e.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_event_box_get_above_child(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// VisibleWindow returns whether the event box has a visible window. See
// gtk_event_box_set_visible_window() for details.
func (e eventBox) VisibleWindow() bool {
	var _arg0 *C.GtkEventBox // out

	_arg0 = (*C.GtkEventBox)(unsafe.Pointer(e.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_event_box_get_visible_window(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// SetAboveChild: set whether the event box window is positioned above the
// windows of its child, as opposed to below it. If the window is above, all
// events inside the event box will go to the event box. If the window is
// below, events in windows of child widgets will first got to that widget,
// and then to its parents.
//
// The default is to keep the window below the child.
func (e eventBox) SetAboveChild(aboveChild bool) {
	var _arg0 *C.GtkEventBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkEventBox)(unsafe.Pointer(e.Native()))
	if aboveChild {
		_arg1 = C.gboolean(1)
	}

	C.gtk_event_box_set_above_child(_arg0, _arg1)
}

// SetVisibleWindow: set whether the event box uses a visible or invisible
// child window. The default is to use visible windows.
//
// In an invisible window event box, the window that the event box creates
// is a GDK_INPUT_ONLY window, which means that it is invisible and only
// serves to receive events.
//
// A visible window event box creates a visible (GDK_INPUT_OUTPUT) window
// that acts as the parent window for all the widgets contained in the event
// box.
//
// You should generally make your event box invisible if you just want to
// trap events. Creating a visible window may cause artifacts that are
// visible to the user, especially if the user is using a theme with
// gradients or pixmaps.
//
// The main reason to create a non input-only event box is if you want to
// set the background to a different color or draw on it.
//
// There is one unexpected issue for an invisible event box that has its
// window below the child. (See gtk_event_box_set_above_child().) Since the
// input-only window is not an ancestor window of any windows that
// descendent widgets of the event box create, events on these windows
// aren’t propagated up by the windowing system, but only by GTK+. The
// practical effect of this is if an event isn’t in the event mask for the
// descendant window (see gtk_widget_add_events()), it won’t be received by
// the event box.
//
// This problem doesn’t occur for visible event boxes, because in that case,
// the event box window is actually the ancestor of the descendant windows,
// not just at the same place on the screen.
func (e eventBox) SetVisibleWindow(visibleWindow bool) {
	var _arg0 *C.GtkEventBox // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkEventBox)(unsafe.Pointer(e.Native()))
	if visibleWindow {
		_arg1 = C.gboolean(1)
	}

	C.gtk_event_box_set_visible_window(_arg0, _arg1)
}
