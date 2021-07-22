// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_app_chooser_widget_get_type()), F: marshalAppChooserWidgetter},
	})
}

// AppChooserWidget: GtkAppChooserWidget is a widget for selecting applications.
//
// It is the main building block for gtk.AppChooserDialog. Most applications
// only need to use the latter; but you can use this widget as part of a larger
// widget if you have special needs.
//
// GtkAppChooserWidget offers detailed control over what applications are shown,
// using the gtk.AppChooserWidget:show-default,
// gtk.AppChooserWidget:show-recommended, gtk.AppChooserWidget:show-fallback,
// gtk.AppChooserWidget:show-other and gtk.AppChooserWidget:show-all properties.
// See the gtk.AppChooser documentation for more information about these groups
// of applications.
//
// To keep track of the selected application, use the
// gtk.AppChooserWidget::application-selected and
// gtk.AppChooserWidget::application-activated signals.
//
//
// CSS nodes
//
// GtkAppChooserWidget has a single CSS node with name appchooser.
type AppChooserWidget struct {
	Widget

	AppChooser
	*externglib.Object
}

func wrapAppChooserWidget(obj *externglib.Object) *AppChooserWidget {
	return &AppChooserWidget{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
		AppChooser: AppChooser{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				Accessible: Accessible{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
				ConstraintTarget: ConstraintTarget{
					Object: obj,
				},
				Object: obj,
			},
		},
		Object: obj,
	}
}

func marshalAppChooserWidgetter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAppChooserWidget(obj), nil
}

// NewAppChooserWidget creates a new GtkAppChooserWidget for applications that
// can handle content of the given type.
func NewAppChooserWidget(contentType string) *AppChooserWidget {
	var _arg1 *C.char      // out
	var _cret *C.GtkWidget // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(contentType)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_app_chooser_widget_new(_arg1)

	var _appChooserWidget *AppChooserWidget // out

	_appChooserWidget = wrapAppChooserWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _appChooserWidget
}

// Native solves the ambiguous selector of this class or interface.
func (self *AppChooserWidget) Native() uintptr {
	return self.Object.Native()
}

// DefaultText returns the text that is shown if there are not applications that
// can handle the content type.
func (self *AppChooserWidget) DefaultText() string {
	var _arg0 *C.GtkAppChooserWidget // out
	var _cret *C.char                // in

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_widget_get_default_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// ShowAll gets whether the app chooser should show all applications in a flat
// list.
func (self *AppChooserWidget) ShowAll() bool {
	var _arg0 *C.GtkAppChooserWidget // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_widget_get_show_all(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowDefault gets whether the app chooser should show the default handler for
// the content type in a separate section.
func (self *AppChooserWidget) ShowDefault() bool {
	var _arg0 *C.GtkAppChooserWidget // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_widget_get_show_default(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowFallback gets whether the app chooser should show related applications
// for the content type in a separate section.
func (self *AppChooserWidget) ShowFallback() bool {
	var _arg0 *C.GtkAppChooserWidget // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_widget_get_show_fallback(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowOther gets whether the app chooser should show applications which are
// unrelated to the content type.
func (self *AppChooserWidget) ShowOther() bool {
	var _arg0 *C.GtkAppChooserWidget // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_widget_get_show_other(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowRecommended gets whether the app chooser should show recommended
// applications for the content type in a separate section.
func (self *AppChooserWidget) ShowRecommended() bool {
	var _arg0 *C.GtkAppChooserWidget // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_app_chooser_widget_get_show_recommended(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetDefaultText sets the text that is shown if there are not applications that
// can handle the content type.
func (self *AppChooserWidget) SetDefaultText(text string) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 *C.char                // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_app_chooser_widget_set_default_text(_arg0, _arg1)
}

// SetShowAll sets whether the app chooser should show all applications in a
// flat list.
func (self *AppChooserWidget) SetShowAll(setting bool) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_widget_set_show_all(_arg0, _arg1)
}

// SetShowDefault sets whether the app chooser should show the default handler
// for the content type in a separate section.
func (self *AppChooserWidget) SetShowDefault(setting bool) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_widget_set_show_default(_arg0, _arg1)
}

// SetShowFallback sets whether the app chooser should show related applications
// for the content type in a separate section.
func (self *AppChooserWidget) SetShowFallback(setting bool) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_widget_set_show_fallback(_arg0, _arg1)
}

// SetShowOther sets whether the app chooser should show applications which are
// unrelated to the content type.
func (self *AppChooserWidget) SetShowOther(setting bool) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_widget_set_show_other(_arg0, _arg1)
}

// SetShowRecommended sets whether the app chooser should show recommended
// applications for the content type in a separate section.
func (self *AppChooserWidget) SetShowRecommended(setting bool) {
	var _arg0 *C.GtkAppChooserWidget // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(self.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_app_chooser_widget_set_show_recommended(_arg0, _arg1)
}
