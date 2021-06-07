// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_settings_get_type()), F: marshalSettings},
	})
}

// RCPropertyParseBorder: a RcPropertyParser for use with
// gtk_settings_install_property_parser() or
// gtk_widget_class_install_style_property_parser() which parses borders in the
// form `"{ left, right, top, bottom }"` for integers left, right, top and
// bottom.
func RCPropertyParseBorder(pspec gobject.ParamSpec, gstring *glib.String, propertyValue *externglib.Value) bool {
	var arg1 *C.GParamSpec
	var arg2 *C.GString
	var arg3 *C.GValue

	arg1 = (*C.GParamSpec)(unsafe.Pointer(pspec.Native()))
	arg2 = (*C.GString)(unsafe.Pointer(gstring.Native()))
	arg3 = (*C.GValue)(propertyValue.GValue)

	var cret C.gboolean
	var ok bool

	cret = C.gtk_rc_property_parse_border(arg1, arg2, arg3)

	if cret {
		ok = true
	}

	return ok
}

// RCPropertyParseColor: a RcPropertyParser for use with
// gtk_settings_install_property_parser() or
// gtk_widget_class_install_style_property_parser() which parses a color given
// either by its name or in the form `{ red, green, blue }` where red, green and
// blue are integers between 0 and 65535 or floating-point numbers between 0 and
// 1.
func RCPropertyParseColor(pspec gobject.ParamSpec, gstring *glib.String, propertyValue *externglib.Value) bool {
	var arg1 *C.GParamSpec
	var arg2 *C.GString
	var arg3 *C.GValue

	arg1 = (*C.GParamSpec)(unsafe.Pointer(pspec.Native()))
	arg2 = (*C.GString)(unsafe.Pointer(gstring.Native()))
	arg3 = (*C.GValue)(propertyValue.GValue)

	var cret C.gboolean
	var ok bool

	cret = C.gtk_rc_property_parse_color(arg1, arg2, arg3)

	if cret {
		ok = true
	}

	return ok
}

// RCPropertyParseEnum: a RcPropertyParser for use with
// gtk_settings_install_property_parser() or
// gtk_widget_class_install_style_property_parser() which parses a single
// enumeration value.
//
// The enumeration value can be specified by its name, its nickname or its
// numeric value. For consistency with flags parsing, the value may be
// surrounded by parentheses.
func RCPropertyParseEnum(pspec gobject.ParamSpec, gstring *glib.String, propertyValue *externglib.Value) bool {
	var arg1 *C.GParamSpec
	var arg2 *C.GString
	var arg3 *C.GValue

	arg1 = (*C.GParamSpec)(unsafe.Pointer(pspec.Native()))
	arg2 = (*C.GString)(unsafe.Pointer(gstring.Native()))
	arg3 = (*C.GValue)(propertyValue.GValue)

	var cret C.gboolean
	var ok bool

	cret = C.gtk_rc_property_parse_enum(arg1, arg2, arg3)

	if cret {
		ok = true
	}

	return ok
}

// RCPropertyParseFlags: a RcPropertyParser for use with
// gtk_settings_install_property_parser() or
// gtk_widget_class_install_style_property_parser() which parses flags.
//
// Flags can be specified by their name, their nickname or numerically. Multiple
// flags can be specified in the form `"( flag1 | flag2 | ... )"`.
func RCPropertyParseFlags(pspec gobject.ParamSpec, gstring *glib.String, propertyValue *externglib.Value) bool {
	var arg1 *C.GParamSpec
	var arg2 *C.GString
	var arg3 *C.GValue

	arg1 = (*C.GParamSpec)(unsafe.Pointer(pspec.Native()))
	arg2 = (*C.GString)(unsafe.Pointer(gstring.Native()))
	arg3 = (*C.GValue)(propertyValue.GValue)

	var cret C.gboolean
	var ok bool

	cret = C.gtk_rc_property_parse_flags(arg1, arg2, arg3)

	if cret {
		ok = true
	}

	return ok
}

// RCPropertyParseRequisition: a RcPropertyParser for use with
// gtk_settings_install_property_parser() or
// gtk_widget_class_install_style_property_parser() which parses a requisition
// in the form `"{ width, height }"` for integers width and height.
func RCPropertyParseRequisition(pspec gobject.ParamSpec, gstring *glib.String, propertyValue *externglib.Value) bool {
	var arg1 *C.GParamSpec
	var arg2 *C.GString
	var arg3 *C.GValue

	arg1 = (*C.GParamSpec)(unsafe.Pointer(pspec.Native()))
	arg2 = (*C.GString)(unsafe.Pointer(gstring.Native()))
	arg3 = (*C.GValue)(propertyValue.GValue)

	var cret C.gboolean
	var ok bool

	cret = C.gtk_rc_property_parse_requisition(arg1, arg2, arg3)

	if cret {
		ok = true
	}

	return ok
}

// Settings gtkSettings provide a mechanism to share global settings between
// applications.
//
// On the X window system, this sharing is realized by an XSettings
// (http://www.freedesktop.org/wiki/Specifications/xsettings-spec) manager that
// is usually part of the desktop environment, along with utilities that let the
// user change these settings. In the absence of an Xsettings manager, GTK+
// reads default values for settings from `settings.ini` files in
// `/etc/gtk-3.0`, `$XDG_CONFIG_DIRS/gtk-3.0` and `$XDG_CONFIG_HOME/gtk-3.0`.
// These files must be valid key files (see File), and have a section called
// Settings. Themes can also provide default values for settings by installing a
// `settings.ini` file next to their `gtk.css` file.
//
// Applications can override system-wide settings by setting the property of the
// GtkSettings object with g_object_set(). This should be restricted to special
// cases though; GtkSettings are not meant as an application configuration
// facility. When doing so, you need to be aware that settings that are specific
// to individual widgets may not be available before the widget type has been
// realized at least once. The following example demonstrates a way to do this:
//
//      gtk_init (&argc, &argv);
//
//      // make sure the type is realized
//      g_type_class_unref (g_type_class_ref (GTK_TYPE_IMAGE_MENU_ITEM));
//
//      g_object_set (gtk_settings_get_default (), "gtk-enable-animations", FALSE, NULL);
//
// There is one GtkSettings instance per screen. It can be obtained with
// gtk_settings_get_for_screen(), but in many cases, it is more convenient to
// use gtk_widget_get_settings(). gtk_settings_get_default() returns the
// GtkSettings instance for the default screen.
type Settings interface {
	gextras.Objector
	StyleProvider

	// ResetProperty undoes the effect of calling g_object_set() to install an
	// application-specific value for a setting. After this call, the setting
	// will again follow the session-wide value for this setting.
	ResetProperty(s Settings, name string)

	SetDoubleProperty(s Settings, name string, vDouble float64, origin string)

	SetLongProperty(s Settings, name string, vLong int32, origin string)

	SetPropertyValue(s Settings, name string, svalue *SettingsValue)

	SetStringProperty(s Settings, name string, vString string, origin string)
}

// settings implements the Settings interface.
type settings struct {
	gextras.Objector
	StyleProvider
}

var _ Settings = (*settings)(nil)

// WrapSettings wraps a GObject to the right type. It is
// primarily used internally.
func WrapSettings(obj *externglib.Object) Settings {
	return Settings{
		Objector:      obj,
		StyleProvider: WrapStyleProvider(obj),
	}
}

func marshalSettings(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSettings(obj), nil
}

// ResetProperty undoes the effect of calling g_object_set() to install an
// application-specific value for a setting. After this call, the setting
// will again follow the session-wide value for this setting.
func (s settings) ResetProperty(s Settings, name string) {
	var arg0 *C.GtkSettings
	var arg1 *C.gchar

	arg0 = (*C.GtkSettings)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_settings_reset_property(arg0, arg1)
}

func (s settings) SetDoubleProperty(s Settings, name string, vDouble float64, origin string) {
	var arg0 *C.GtkSettings
	var arg1 *C.gchar
	var arg2 C.gdouble
	var arg3 *C.gchar

	arg0 = (*C.GtkSettings)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gdouble(vDouble)
	arg3 = (*C.gchar)(C.CString(origin))
	defer C.free(unsafe.Pointer(arg3))

	C.gtk_settings_set_double_property(arg0, arg1, arg2, arg3)
}

func (s settings) SetLongProperty(s Settings, name string, vLong int32, origin string) {
	var arg0 *C.GtkSettings
	var arg1 *C.gchar
	var arg2 C.glong
	var arg3 *C.gchar

	arg0 = (*C.GtkSettings)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.glong(vLong)
	arg3 = (*C.gchar)(C.CString(origin))
	defer C.free(unsafe.Pointer(arg3))

	C.gtk_settings_set_long_property(arg0, arg1, arg2, arg3)
}

func (s settings) SetPropertyValue(s Settings, name string, svalue *SettingsValue) {
	var arg0 *C.GtkSettings
	var arg1 *C.gchar
	var arg2 *C.GtkSettingsValue

	arg0 = (*C.GtkSettings)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GtkSettingsValue)(unsafe.Pointer(svalue.Native()))

	C.gtk_settings_set_property_value(arg0, arg1, arg2)
}

func (s settings) SetStringProperty(s Settings, name string, vString string, origin string) {
	var arg0 *C.GtkSettings
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar

	arg0 = (*C.GtkSettings)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(vString))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(origin))
	defer C.free(unsafe.Pointer(arg3))

	C.gtk_settings_set_string_property(arg0, arg1, arg2, arg3)
}

type SettingsValue struct {
	native C.GtkSettingsValue
}

// WrapSettingsValue wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapSettingsValue(ptr unsafe.Pointer) *SettingsValue {
	if ptr == nil {
		return nil
	}

	return (*SettingsValue)(ptr)
}

func marshalSettingsValue(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapSettingsValue(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *SettingsValue) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// Origin gets the field inside the struct.
func (s *SettingsValue) Origin() string {
	var v string
	v = C.GoString(s.native.origin)
	return v
}

// Value gets the field inside the struct.
func (s *SettingsValue) Value() *externglib.Value {
	var v *externglib.Value
	v = externglib.ValueFromNative(unsafe.Pointer(s.native.value))
	return v
}
