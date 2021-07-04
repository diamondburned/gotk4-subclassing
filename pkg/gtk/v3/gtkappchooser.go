// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
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
		{T: externglib.Type(C.gtk_app_chooser_get_type()), F: marshalAppChooser},
	})
}

// AppChooser is an interface that can be implemented by widgets which allow the
// user to choose an application (typically for the purpose of opening a file).
// The main objects that implement this interface are AppChooserWidget,
// AppChooserDialog and AppChooserButton.
//
// Applications are represented by GIO Info objects here. GIO has a concept of
// recommended and fallback applications for a given content type. Recommended
// applications are those that claim to handle the content type itself, while
// fallback also includes applications that handle a more generic content type.
// GIO also knows the default and last-used application for a given content
// type. The AppChooserWidget provides detailed control over whether the shown
// list of applications should include default, recommended or fallback
// applications.
//
// To obtain the application that has been selected in a AppChooser, use
// gtk_app_chooser_get_app_info().
type AppChooser interface {
	Widget

	// AppInfo reloads the list of applications.
	AppInfo() gio.AppInfo
	// ContentType reloads the list of applications.
	ContentType() string
	// Refresh reloads the list of applications.
	Refresh()
}

// appChooser implements the AppChooser interface.
type appChooser struct {
	Widget
}

var _ AppChooser = (*appChooser)(nil)

// WrapAppChooser wraps a GObject to a type that implements
// interface AppChooser. It is primarily used internally.
func WrapAppChooser(obj *externglib.Object) AppChooser {
	return appChooser{
		Widget: WrapWidget(obj),
	}
}

func marshalAppChooser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAppChooser(obj), nil
}

func (s appChooser) AppInfo() gio.AppInfo {
	var _arg0 *C.GtkAppChooser // out
	var _cret *C.GAppInfo      // in

	_arg0 = (*C.GtkAppChooser)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_app_chooser_get_app_info(_arg0)

	var _appInfo gio.AppInfo // out

	_appInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(gio.AppInfo)

	return _appInfo
}

func (s appChooser) ContentType() string {
	var _arg0 *C.GtkAppChooser // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkAppChooser)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_app_chooser_get_content_type(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (s appChooser) Refresh() {
	var _arg0 *C.GtkAppChooser // out

	_arg0 = (*C.GtkAppChooser)(unsafe.Pointer(s.Native()))

	C.gtk_app_chooser_refresh(_arg0)
}