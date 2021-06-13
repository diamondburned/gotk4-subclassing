// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gobject/v2"
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
		{T: externglib.Type(C.gtk_style_provider_get_type()), F: marshalStyleProvider},
	})
}

// StyleProviderOverrider contains methods that are overridable. This
// interface is a subset of the interface StyleProvider.
type StyleProviderOverrider interface {
	// StyleProperty looks up a widget style property as defined by @provider
	// for the widget represented by @path.
	StyleProperty(path *WidgetPath, state StateFlags, pspec gobject.ParamSpec) (*externglib.Value, bool)
}

// StyleProvider: gtkStyleProvider is an interface used to provide style
// information to a StyleContext. See gtk_style_context_add_provider() and
// gtk_style_context_add_provider_for_screen().
type StyleProvider interface {
	gextras.Objector
	StyleProviderOverrider
}

// styleProvider implements the StyleProvider interface.
type styleProvider struct {
	gextras.Objector
}

var _ StyleProvider = (*styleProvider)(nil)

// WrapStyleProvider wraps a GObject to a type that implements interface
// StyleProvider. It is primarily used internally.
func WrapStyleProvider(obj *externglib.Object) StyleProvider {
	return StyleProvider{
		Objector: obj,
	}
}

func marshalStyleProvider(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStyleProvider(obj), nil
}

// StyleProperty looks up a widget style property as defined by @provider
// for the widget represented by @path.
func (p styleProvider) StyleProperty(path *WidgetPath, state StateFlags, pspec gobject.ParamSpec) (*externglib.Value, bool) {
	var _arg0 *C.GtkStyleProvider // out
	var _arg1 *C.GtkWidgetPath    // out
	var _arg2 C.GtkStateFlags     // out
	var _arg3 *C.GParamSpec       // out

	_arg0 = (*C.GtkStyleProvider)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidgetPath)(unsafe.Pointer(path.Native()))
	_arg2 = (C.GtkStateFlags)(state)
	_arg3 = (*C.GParamSpec)(unsafe.Pointer(pspec.Native()))

	var _arg4 C.GValue   // in
	var _cret C.gboolean // in

	_cret = C.gtk_style_provider_get_style_property(_arg0, _arg1, _arg2, _arg3, &_arg4)

	var _value *externglib.Value // out
	var _ok bool                 // out

	_value = externglib.ValueFromNative(unsafe.Pointer(_arg4))
	if _cret {
		_ok = true
	}

	return _value, _ok
}
