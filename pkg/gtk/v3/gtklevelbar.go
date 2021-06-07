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
		{T: externglib.Type(C.gtk_level_bar_get_type()), F: marshalLevelBar},
	})
}

// LevelBar: the LevelBar is a bar widget that can be used as a level indicator.
// Typical use cases are displaying the strength of a password, or showing the
// charge level of a battery.
//
// Use gtk_level_bar_set_value() to set the current value, and
// gtk_level_bar_add_offset_value() to set the value offsets at which the bar
// will be considered in a different state. GTK will add a few offsets by
// default on the level bar: K_LEVEL_BAR_OFFSET_LOW, K_LEVEL_BAR_OFFSET_HIGH and
// K_LEVEL_BAR_OFFSET_FULL, with values 0.25, 0.75 and 1.0 respectively.
//
// Note that it is your responsibility to update preexisting offsets when
// changing the minimum or maximum value. GTK+ will simply clamp them to the new
// range.
//
// Adding a custom offset on the bar
//
//    levelbar[.discrete]
//    ╰── trough
//        ├── block.filled.level-name
//        ┊
//        ├── block.empty
//        ┊
//
// GtkLevelBar has a main CSS node with name levelbar and one of the style
// classes .discrete or .continuous and a subnode with name trough. Below the
// trough node are a number of nodes with name block and style class .filled or
// .empty. In continuous mode, there is exactly one node of each, in discrete
// mode, the number of filled and unfilled nodes corresponds to blocks that are
// drawn. The block.filled nodes also get a style class .level-name
// corresponding to the level for the current value.
//
// In horizontal orientation, the nodes are always arranged from left to right,
// regardless of text direction.
type LevelBar interface {
	Widget
	Buildable
	Orientable

	// AddOffsetValue adds a new offset marker on @self at the position
	// specified by @value. When the bar value is in the interval topped by
	// @value (or between @value and LevelBar:max-value in case the offset is
	// the last one on the bar) a style class named `level-`@name will be
	// applied when rendering the level bar fill. If another offset marker named
	// @name exists, its value will be replaced by @value.
	AddOffsetValue(s LevelBar, name string, value float64)
	// Inverted: return the value of the LevelBar:inverted property.
	Inverted(s LevelBar) bool
	// MaxValue returns the value of the LevelBar:max-value property.
	MaxValue(s LevelBar)
	// MinValue returns the value of the LevelBar:min-value property.
	MinValue(s LevelBar)
	// Mode returns the value of the LevelBar:mode property.
	Mode(s LevelBar)
	// OffsetValue fetches the value specified for the offset marker @name in
	// @self, returning true in case an offset named @name was found.
	OffsetValue(s LevelBar, name string) (value float64, ok bool)
	// Value returns the value of the LevelBar:value property.
	Value(s LevelBar)
	// RemoveOffsetValue removes an offset marker previously added with
	// gtk_level_bar_add_offset_value().
	RemoveOffsetValue(s LevelBar, name string)
	// SetInverted sets the value of the LevelBar:inverted property.
	SetInverted(s LevelBar, inverted bool)
	// SetMaxValue sets the value of the LevelBar:max-value property.
	//
	// You probably want to update preexisting level offsets after calling this
	// function.
	SetMaxValue(s LevelBar, value float64)
	// SetMinValue sets the value of the LevelBar:min-value property.
	//
	// You probably want to update preexisting level offsets after calling this
	// function.
	SetMinValue(s LevelBar, value float64)
	// SetMode sets the value of the LevelBar:mode property.
	SetMode(s LevelBar, mode LevelBarMode)
	// SetValue sets the value of the LevelBar:value property.
	SetValue(s LevelBar, value float64)
}

// levelBar implements the LevelBar interface.
type levelBar struct {
	Widget
	Buildable
	Orientable
}

var _ LevelBar = (*levelBar)(nil)

// WrapLevelBar wraps a GObject to the right type. It is
// primarily used internally.
func WrapLevelBar(obj *externglib.Object) LevelBar {
	return LevelBar{
		Widget:     WrapWidget(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
	}
}

func marshalLevelBar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapLevelBar(obj), nil
}

// NewLevelBar constructs a class LevelBar.
func NewLevelBar() {
	C.gtk_level_bar_new()
}

// NewLevelBarForInterval constructs a class LevelBar.
func NewLevelBarForInterval(minValue float64, maxValue float64) {
	var arg1 C.gdouble
	var arg2 C.gdouble

	arg1 = C.gdouble(minValue)
	arg2 = C.gdouble(maxValue)

	C.gtk_level_bar_new_for_interval(arg1, arg2)
}

// AddOffsetValue adds a new offset marker on @self at the position
// specified by @value. When the bar value is in the interval topped by
// @value (or between @value and LevelBar:max-value in case the offset is
// the last one on the bar) a style class named `level-`@name will be
// applied when rendering the level bar fill. If another offset marker named
// @name exists, its value will be replaced by @value.
func (s levelBar) AddOffsetValue(s LevelBar, name string, value float64) {
	var arg0 *C.GtkLevelBar
	var arg1 *C.gchar
	var arg2 C.gdouble

	arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gdouble(value)

	C.gtk_level_bar_add_offset_value(arg0, arg1, arg2)
}

// Inverted: return the value of the LevelBar:inverted property.
func (s levelBar) Inverted(s LevelBar) bool {
	var arg0 *C.GtkLevelBar

	arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_level_bar_get_inverted(arg0)

	if cret {
		ok = true
	}

	return ok
}

// MaxValue returns the value of the LevelBar:max-value property.
func (s levelBar) MaxValue(s LevelBar) {
	var arg0 *C.GtkLevelBar

	arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))

	C.gtk_level_bar_get_max_value(arg0)
}

// MinValue returns the value of the LevelBar:min-value property.
func (s levelBar) MinValue(s LevelBar) {
	var arg0 *C.GtkLevelBar

	arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))

	C.gtk_level_bar_get_min_value(arg0)
}

// Mode returns the value of the LevelBar:mode property.
func (s levelBar) Mode(s LevelBar) {
	var arg0 *C.GtkLevelBar

	arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))

	C.gtk_level_bar_get_mode(arg0)
}

// OffsetValue fetches the value specified for the offset marker @name in
// @self, returning true in case an offset named @name was found.
func (s levelBar) OffsetValue(s LevelBar, name string) (value float64, ok bool) {
	var arg0 *C.GtkLevelBar
	var arg1 *C.gchar

	arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	var arg2 C.gdouble
	var value float64
	var cret C.gboolean
	var ok bool

	cret = C.gtk_level_bar_get_offset_value(arg0, arg1, &arg2)

	value = float64(&arg2)
	if cret {
		ok = true
	}

	return value, ok
}

// Value returns the value of the LevelBar:value property.
func (s levelBar) Value(s LevelBar) {
	var arg0 *C.GtkLevelBar

	arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))

	C.gtk_level_bar_get_value(arg0)
}

// RemoveOffsetValue removes an offset marker previously added with
// gtk_level_bar_add_offset_value().
func (s levelBar) RemoveOffsetValue(s LevelBar, name string) {
	var arg0 *C.GtkLevelBar
	var arg1 *C.gchar

	arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_level_bar_remove_offset_value(arg0, arg1)
}

// SetInverted sets the value of the LevelBar:inverted property.
func (s levelBar) SetInverted(s LevelBar, inverted bool) {
	var arg0 *C.GtkLevelBar
	var arg1 C.gboolean

	arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	if inverted {
		arg1 = C.gboolean(1)
	}

	C.gtk_level_bar_set_inverted(arg0, arg1)
}

// SetMaxValue sets the value of the LevelBar:max-value property.
//
// You probably want to update preexisting level offsets after calling this
// function.
func (s levelBar) SetMaxValue(s LevelBar, value float64) {
	var arg0 *C.GtkLevelBar
	var arg1 C.gdouble

	arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	arg1 = C.gdouble(value)

	C.gtk_level_bar_set_max_value(arg0, arg1)
}

// SetMinValue sets the value of the LevelBar:min-value property.
//
// You probably want to update preexisting level offsets after calling this
// function.
func (s levelBar) SetMinValue(s LevelBar, value float64) {
	var arg0 *C.GtkLevelBar
	var arg1 C.gdouble

	arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	arg1 = C.gdouble(value)

	C.gtk_level_bar_set_min_value(arg0, arg1)
}

// SetMode sets the value of the LevelBar:mode property.
func (s levelBar) SetMode(s LevelBar, mode LevelBarMode) {
	var arg0 *C.GtkLevelBar
	var arg1 C.GtkLevelBarMode

	arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	arg1 = (C.GtkLevelBarMode)(mode)

	C.gtk_level_bar_set_mode(arg0, arg1)
}

// SetValue sets the value of the LevelBar:value property.
func (s levelBar) SetValue(s LevelBar, value float64) {
	var arg0 *C.GtkLevelBar
	var arg1 C.gdouble

	arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	arg1 = C.gdouble(value)

	C.gtk_level_bar_set_value(arg0, arg1)
}