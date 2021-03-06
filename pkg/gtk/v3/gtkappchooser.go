// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_app_chooser_get_type()), F: marshalAppChooserer},
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
type AppChooser struct {
	_ [0]func() // equal guard
	Widget
}

var (
	_ Widgetter = (*AppChooser)(nil)
)

// AppChooserer describes AppChooser's interface methods.
type AppChooserer interface {
	externglib.Objector

	// AppInfo returns the currently selected application.
	AppInfo() gio.AppInfor
	// ContentType returns the current value of the AppChooser:content-type
	// property.
	ContentType() string
	// Refresh reloads the list of applications.
	Refresh()
}

var _ AppChooserer = (*AppChooser)(nil)

func wrapAppChooser(obj *externglib.Object) *AppChooser {
	return &AppChooser{
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
	}
}

func marshalAppChooserer(p uintptr) (interface{}, error) {
	return wrapAppChooser(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AppInfo returns the currently selected application.
//
// The function returns the following values:
//
//    - appInfo (optional) for the currently selected application, or NULL if
//      none is selected. Free with g_object_unref().
//
func (self *AppChooser) AppInfo() gio.AppInfor {
	var _arg0 *C.GtkAppChooser // out
	var _cret *C.GAppInfo      // in

	_arg0 = (*C.GtkAppChooser)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_get_app_info(_arg0)
	runtime.KeepAlive(self)

	var _appInfo gio.AppInfor // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj externglib.Objector) bool {
				_, ok := obj.(gio.AppInfor)
				return ok
			})
			rv, ok := casted.(gio.AppInfor)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.AppInfor")
			}
			_appInfo = rv
		}
	}

	return _appInfo
}

// ContentType returns the current value of the AppChooser:content-type
// property.
//
// The function returns the following values:
//
//    - utf8: content type of self. Free with g_free().
//
func (self *AppChooser) ContentType() string {
	var _arg0 *C.GtkAppChooser // out
	var _cret *C.gchar         // in

	_arg0 = (*C.GtkAppChooser)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_get_content_type(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Refresh reloads the list of applications.
func (self *AppChooser) Refresh() {
	var _arg0 *C.GtkAppChooser // out

	_arg0 = (*C.GtkAppChooser)(unsafe.Pointer(self.Native()))

	C.gtk_app_chooser_refresh(_arg0)
	runtime.KeepAlive(self)
}
