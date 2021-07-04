// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_window_controls_get_type()), F: marshalWindowControls},
	})
}

// WindowControls: `GtkWindowControls` shows window frame controls.
//
// Typical window frame controls are minimize, maximize and close buttons, and
// the window icon.
//
// !An example GtkWindowControls (windowcontrols.png)
//
// `GtkWindowControls` only displays start or end side of the controls (see
// [property@Gtk.WindowControls:side]), so it's intended to be always used in
// pair with another `GtkWindowControls` for the opposite side, for example:
//
// “`xml <object class="GtkBox"> <child> <object class="GtkWindowControls">
// <property name="side">start</property> </object> </child>
//
//    ...
//
//    <child>
//      <object class="GtkWindowControls">
//        <property name="side">end</property>
//      </object>
//    </child>
//
// </object> “`
//
//
// CSS nodes
//
// “` windowcontrols ├── [image.icon] ├── [button.minimize] ├──
// [button.maximize] ╰── [button.close] “`
//
// A `GtkWindowControls`' CSS node is called windowcontrols. It contains
// subnodes corresponding to each title button. Which of the title buttons exist
// and where they are placed exactly depends on the desktop environment and
// [property@Gtk.WindowControls:decoration-layout] value.
//
// When [property@Gtk.WindowControls:empty] is true, it gets the .empty style
// class.
//
//
// Accessibility
//
// `GtkWindowControls` uses the GTK_ACCESSIBLE_ROLE_GROUP role.
type WindowControls interface {
	Widget

	// DecorationLayout:
	DecorationLayout() string
	// Empty:
	Empty() bool
	// Side:
	Side() PackType
	// SetDecorationLayoutWindowControls:
	SetDecorationLayoutWindowControls(layout string)
	// SetSideWindowControls:
	SetSideWindowControls(side PackType)
}

// windowControls implements the WindowControls class.
type windowControls struct {
	Widget
}

// WrapWindowControls wraps a GObject to the right type. It is
// primarily used internally.
func WrapWindowControls(obj *externglib.Object) WindowControls {
	return windowControls{
		Widget: WrapWidget(obj),
	}
}

func marshalWindowControls(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapWindowControls(obj), nil
}

// NewWindowControls:
func NewWindowControls(side PackType) WindowControls {
	var _arg1 C.GtkPackType // out
	var _cret *C.GtkWidget  // in

	_arg1 = C.GtkPackType(side)

	_cret = C.gtk_window_controls_new(_arg1)

	var _windowControls WindowControls // out

	_windowControls = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(WindowControls)

	return _windowControls
}

func (s windowControls) DecorationLayout() string {
	var _arg0 *C.GtkWindowControls // out
	var _cret *C.char              // in

	_arg0 = (*C.GtkWindowControls)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_window_controls_get_decoration_layout(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (s windowControls) Empty() bool {
	var _arg0 *C.GtkWindowControls // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GtkWindowControls)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_window_controls_get_empty(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s windowControls) Side() PackType {
	var _arg0 *C.GtkWindowControls // out
	var _cret C.GtkPackType        // in

	_arg0 = (*C.GtkWindowControls)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_window_controls_get_side(_arg0)

	var _packType PackType // out

	_packType = PackType(_cret)

	return _packType
}

func (s windowControls) SetDecorationLayoutWindowControls(layout string) {
	var _arg0 *C.GtkWindowControls // out
	var _arg1 *C.char              // out

	_arg0 = (*C.GtkWindowControls)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(layout))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_window_controls_set_decoration_layout(_arg0, _arg1)
}

func (s windowControls) SetSideWindowControls(side PackType) {
	var _arg0 *C.GtkWindowControls // out
	var _arg1 C.GtkPackType        // out

	_arg0 = (*C.GtkWindowControls)(unsafe.Pointer(s.Native()))
	_arg1 = C.GtkPackType(side)

	C.gtk_window_controls_set_side(_arg0, _arg1)
}

func (s windowControls) AccessibleRole() AccessibleRole {
	return WrapAccessible(gextras.InternObject(s)).AccessibleRole()
}

func (s windowControls) ResetProperty(property AccessibleProperty) {
	WrapAccessible(gextras.InternObject(s)).ResetProperty(property)
}

func (s windowControls) ResetRelation(relation AccessibleRelation) {
	WrapAccessible(gextras.InternObject(s)).ResetRelation(relation)
}

func (s windowControls) ResetState(state AccessibleState) {
	WrapAccessible(gextras.InternObject(s)).ResetState(state)
}

func (s windowControls) UpdatePropertyValue(properties []AccessibleProperty, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdatePropertyValue(properties, values)
}

func (s windowControls) UpdateRelationValue(relations []AccessibleRelation, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateRelationValue(relations, values)
}

func (s windowControls) UpdateStateValue(states []AccessibleState, values []externglib.Value) {
	WrapAccessible(gextras.InternObject(s)).UpdateStateValue(states, values)
}

func (b windowControls) BuildableID() string {
	return WrapBuildable(gextras.InternObject(b)).BuildableID()
}