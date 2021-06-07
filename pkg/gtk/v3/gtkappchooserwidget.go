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
		{T: externglib.Type(C.gtk_app_chooser_widget_get_type()), F: marshalAppChooserWidget},
	})
}

// AppChooserWidget is a widget for selecting applications. It is the main
// building block for AppChooserDialog. Most applications only need to use the
// latter; but you can use this widget as part of a larger widget if you have
// special needs.
//
// AppChooserWidget offers detailed control over what applications are shown,
// using the AppChooserWidget:show-default, AppChooserWidget:show-recommended,
// AppChooserWidget:show-fallback, AppChooserWidget:show-other and
// AppChooserWidget:show-all properties. See the AppChooser documentation for
// more information about these groups of applications.
//
// To keep track of the selected application, use the
// AppChooserWidget::application-selected and
// AppChooserWidget::application-activated signals.
//
//
// CSS nodes
//
// GtkAppChooserWidget has a single CSS node with name appchooser.
type AppChooserWidget interface {
	Box
	AppChooser
	Buildable
	Orientable

	// DefaultText returns the text that is shown if there are not applications
	// that can handle the content type.
	DefaultText(s AppChooserWidget)
	// ShowAll returns the current value of the AppChooserWidget:show-all
	// property.
	ShowAll(s AppChooserWidget) bool
	// ShowDefault returns the current value of the
	// AppChooserWidget:show-default property.
	ShowDefault(s AppChooserWidget) bool
	// ShowFallback returns the current value of the
	// AppChooserWidget:show-fallback property.
	ShowFallback(s AppChooserWidget) bool
	// ShowOther returns the current value of the AppChooserWidget:show-other
	// property.
	ShowOther(s AppChooserWidget) bool
	// ShowRecommended returns the current value of the
	// AppChooserWidget:show-recommended property.
	ShowRecommended(s AppChooserWidget) bool
	// SetDefaultText sets the text that is shown if there are not applications
	// that can handle the content type.
	SetDefaultText(s AppChooserWidget, text string)
	// SetShowAll sets whether the app chooser should show all applications in a
	// flat list.
	SetShowAll(s AppChooserWidget, setting bool)
	// SetShowDefault sets whether the app chooser should show the default
	// handler for the content type in a separate section.
	SetShowDefault(s AppChooserWidget, setting bool)
	// SetShowFallback sets whether the app chooser should show related
	// applications for the content type in a separate section.
	SetShowFallback(s AppChooserWidget, setting bool)
	// SetShowOther sets whether the app chooser should show applications which
	// are unrelated to the content type.
	SetShowOther(s AppChooserWidget, setting bool)
	// SetShowRecommended sets whether the app chooser should show recommended
	// applications for the content type in a separate section.
	SetShowRecommended(s AppChooserWidget, setting bool)
}

// appChooserWidget implements the AppChooserWidget interface.
type appChooserWidget struct {
	Box
	AppChooser
	Buildable
	Orientable
}

var _ AppChooserWidget = (*appChooserWidget)(nil)

// WrapAppChooserWidget wraps a GObject to the right type. It is
// primarily used internally.
func WrapAppChooserWidget(obj *externglib.Object) AppChooserWidget {
	return AppChooserWidget{
		Box:        WrapBox(obj),
		AppChooser: WrapAppChooser(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalAppChooserWidget(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAppChooserWidget(obj), nil
}

// NewAppChooserWidget constructs a class AppChooserWidget.
func NewAppChooserWidget(contentType string) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(contentType))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_app_chooser_widget_new(arg1)
}

// DefaultText returns the text that is shown if there are not applications
// that can handle the content type.
func (s appChooserWidget) DefaultText(s AppChooserWidget) {
	var arg0 *C.GtkAppChooserWidget

	arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))

	C.gtk_app_chooser_widget_get_default_text(arg0)
}

// ShowAll returns the current value of the AppChooserWidget:show-all
// property.
func (s appChooserWidget) ShowAll(s AppChooserWidget) bool {
	var arg0 *C.GtkAppChooserWidget

	arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_app_chooser_widget_get_show_all(arg0)

	if cret {
		ok = true
	}

	return ok
}

// ShowDefault returns the current value of the
// AppChooserWidget:show-default property.
func (s appChooserWidget) ShowDefault(s AppChooserWidget) bool {
	var arg0 *C.GtkAppChooserWidget

	arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_app_chooser_widget_get_show_default(arg0)

	if cret {
		ok = true
	}

	return ok
}

// ShowFallback returns the current value of the
// AppChooserWidget:show-fallback property.
func (s appChooserWidget) ShowFallback(s AppChooserWidget) bool {
	var arg0 *C.GtkAppChooserWidget

	arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_app_chooser_widget_get_show_fallback(arg0)

	if cret {
		ok = true
	}

	return ok
}

// ShowOther returns the current value of the AppChooserWidget:show-other
// property.
func (s appChooserWidget) ShowOther(s AppChooserWidget) bool {
	var arg0 *C.GtkAppChooserWidget

	arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_app_chooser_widget_get_show_other(arg0)

	if cret {
		ok = true
	}

	return ok
}

// ShowRecommended returns the current value of the
// AppChooserWidget:show-recommended property.
func (s appChooserWidget) ShowRecommended(s AppChooserWidget) bool {
	var arg0 *C.GtkAppChooserWidget

	arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_app_chooser_widget_get_show_recommended(arg0)

	if cret {
		ok = true
	}

	return ok
}

// SetDefaultText sets the text that is shown if there are not applications
// that can handle the content type.
func (s appChooserWidget) SetDefaultText(s AppChooserWidget, text string) {
	var arg0 *C.GtkAppChooserWidget
	var arg1 *C.gchar

	arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_app_chooser_widget_set_default_text(arg0, arg1)
}

// SetShowAll sets whether the app chooser should show all applications in a
// flat list.
func (s appChooserWidget) SetShowAll(s AppChooserWidget, setting bool) {
	var arg0 *C.GtkAppChooserWidget
	var arg1 C.gboolean

	arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))
	if setting {
		arg1 = C.gboolean(1)
	}

	C.gtk_app_chooser_widget_set_show_all(arg0, arg1)
}

// SetShowDefault sets whether the app chooser should show the default
// handler for the content type in a separate section.
func (s appChooserWidget) SetShowDefault(s AppChooserWidget, setting bool) {
	var arg0 *C.GtkAppChooserWidget
	var arg1 C.gboolean

	arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))
	if setting {
		arg1 = C.gboolean(1)
	}

	C.gtk_app_chooser_widget_set_show_default(arg0, arg1)
}

// SetShowFallback sets whether the app chooser should show related
// applications for the content type in a separate section.
func (s appChooserWidget) SetShowFallback(s AppChooserWidget, setting bool) {
	var arg0 *C.GtkAppChooserWidget
	var arg1 C.gboolean

	arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))
	if setting {
		arg1 = C.gboolean(1)
	}

	C.gtk_app_chooser_widget_set_show_fallback(arg0, arg1)
}

// SetShowOther sets whether the app chooser should show applications which
// are unrelated to the content type.
func (s appChooserWidget) SetShowOther(s AppChooserWidget, setting bool) {
	var arg0 *C.GtkAppChooserWidget
	var arg1 C.gboolean

	arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))
	if setting {
		arg1 = C.gboolean(1)
	}

	C.gtk_app_chooser_widget_set_show_other(arg0, arg1)
}

// SetShowRecommended sets whether the app chooser should show recommended
// applications for the content type in a separate section.
func (s appChooserWidget) SetShowRecommended(s AppChooserWidget, setting bool) {
	var arg0 *C.GtkAppChooserWidget
	var arg1 C.gboolean

	arg0 = (*C.GtkAppChooserWidget)(unsafe.Pointer(s.Native()))
	if setting {
		arg1 = C.gboolean(1)
	}

	C.gtk_app_chooser_widget_set_show_recommended(arg0, arg1)
}