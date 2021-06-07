// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_link_button_get_type()), F: marshalLinkButton},
	})
}

// LinkButton: a GtkLinkButton is a Button with a hyperlink, similar to the one
// used by web browsers, which triggers an action when clicked. It is useful to
// show quick links to resources.
//
// A link button is created by calling either gtk_link_button_new() or
// gtk_link_button_new_with_label(). If using the former, the URI you pass to
// the constructor is used as a label for the widget.
//
// The URI bound to a GtkLinkButton can be set specifically using
// gtk_link_button_set_uri(), and retrieved using gtk_link_button_get_uri().
//
// By default, GtkLinkButton calls gtk_show_uri_on_window() when the button is
// clicked. This behaviour can be overridden by connecting to the
// LinkButton::activate-link signal and returning true from the signal handler.
//
//
// CSS nodes
//
// GtkLinkButton has a single CSS node with name button. To differentiate it
// from a plain Button, it gets the .link style class.
type LinkButton interface {
	Button
	Actionable
	Activatable
	Buildable

	// URI retrieves the URI set using gtk_link_button_set_uri().
	URI(l LinkButton)
	// Visited retrieves the “visited” state of the URI where the LinkButton
	// points. The button becomes visited when it is clicked. If the URI is
	// changed on the button, the “visited” state is unset again.
	//
	// The state may also be changed using gtk_link_button_set_visited().
	Visited(l LinkButton) bool
	// SetURI sets @uri as the URI where the LinkButton points. As a side-effect
	// this unsets the “visited” state of the button.
	SetURI(l LinkButton, uri string)
	// SetVisited sets the “visited” state of the URI where the LinkButton
	// points. See gtk_link_button_get_visited() for more details.
	SetVisited(l LinkButton, visited bool)
}

// linkButton implements the LinkButton interface.
type linkButton struct {
	Button
	Actionable
	Activatable
	Buildable
}

var _ LinkButton = (*linkButton)(nil)

// WrapLinkButton wraps a GObject to the right type. It is
// primarily used internally.
func WrapLinkButton(obj *externglib.Object) LinkButton {
	return LinkButton{
		Button:      WrapButton(obj),
		Actionable:  WrapActionable(obj),
		Activatable: WrapActivatable(obj),
		Buildable:   WrapBuildable(obj),
	}
}

func marshalLinkButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapLinkButton(obj), nil
}

// NewLinkButton constructs a class LinkButton.
func NewLinkButton(uri string) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_link_button_new(arg1)
}

// NewLinkButtonWithLabel constructs a class LinkButton.
func NewLinkButtonWithLabel(uri string, label string) {
	var arg1 *C.gchar
	var arg2 *C.gchar

	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg2))

	C.gtk_link_button_new_with_label(arg1, arg2)
}

// URI retrieves the URI set using gtk_link_button_set_uri().
func (l linkButton) URI(l LinkButton) {
	var arg0 *C.GtkLinkButton

	arg0 = (*C.GtkLinkButton)(unsafe.Pointer(l.Native()))

	C.gtk_link_button_get_uri(arg0)
}

// Visited retrieves the “visited” state of the URI where the LinkButton
// points. The button becomes visited when it is clicked. If the URI is
// changed on the button, the “visited” state is unset again.
//
// The state may also be changed using gtk_link_button_set_visited().
func (l linkButton) Visited(l LinkButton) bool {
	var arg0 *C.GtkLinkButton

	arg0 = (*C.GtkLinkButton)(unsafe.Pointer(l.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_link_button_get_visited(arg0)

	if cret {
		ok = true
	}

	return ok
}

// SetURI sets @uri as the URI where the LinkButton points. As a side-effect
// this unsets the “visited” state of the button.
func (l linkButton) SetURI(l LinkButton, uri string) {
	var arg0 *C.GtkLinkButton
	var arg1 *C.gchar

	arg0 = (*C.GtkLinkButton)(unsafe.Pointer(l.Native()))
	arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_link_button_set_uri(arg0, arg1)
}

// SetVisited sets the “visited” state of the URI where the LinkButton
// points. See gtk_link_button_get_visited() for more details.
func (l linkButton) SetVisited(l LinkButton, visited bool) {
	var arg0 *C.GtkLinkButton
	var arg1 C.gboolean

	arg0 = (*C.GtkLinkButton)(unsafe.Pointer(l.Native()))
	if visited {
		arg1 = C.gboolean(1)
	}

	C.gtk_link_button_set_visited(arg0, arg1)
}
