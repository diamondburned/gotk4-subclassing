// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_level_bar_get_type()), F: marshalLevelBar},
	})
}

// LevelBarOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type LevelBarOverrider interface {
	OffsetChanged(name string)
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
	gextras.Objector

	// AddOffsetValue adds a new offset marker on @self at the position
	// specified by @value. When the bar value is in the interval topped by
	// @value (or between @value and LevelBar:max-value in case the offset is
	// the last one on the bar) a style class named `level-`@name will be
	// applied when rendering the level bar fill. If another offset marker named
	// @name exists, its value will be replaced by @value.
	AddOffsetValue(name string, value float64)
	// Inverted: return the value of the LevelBar:inverted property.
	Inverted() bool
	// MaxValue returns the value of the LevelBar:max-value property.
	MaxValue() float64
	// MinValue returns the value of the LevelBar:min-value property.
	MinValue() float64
	// Mode returns the value of the LevelBar:mode property.
	Mode() LevelBarMode
	// OffsetValue fetches the value specified for the offset marker @name in
	// @self, returning true in case an offset named @name was found.
	OffsetValue(name string) (float64, bool)
	// Value returns the value of the LevelBar:value property.
	Value() float64
	// RemoveOffsetValue removes an offset marker previously added with
	// gtk_level_bar_add_offset_value().
	RemoveOffsetValue(name string)
	// SetInverted sets the value of the LevelBar:inverted property.
	SetInverted(inverted bool)
	// SetMaxValue sets the value of the LevelBar:max-value property.
	//
	// You probably want to update preexisting level offsets after calling this
	// function.
	SetMaxValue(value float64)
	// SetMinValue sets the value of the LevelBar:min-value property.
	//
	// You probably want to update preexisting level offsets after calling this
	// function.
	SetMinValue(value float64)
	// SetValue sets the value of the LevelBar:value property.
	SetValue(value float64)
}

// LevelBarClass implements the LevelBar interface.
type LevelBarClass struct {
	*externglib.Object
	WidgetClass
	BuildableInterface
	OrientableInterface
}

var _ LevelBar = (*LevelBarClass)(nil)

func wrapLevelBar(obj *externglib.Object) LevelBar {
	return &LevelBarClass{
		Object: obj,
		WidgetClass: WidgetClass{
			InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		OrientableInterface: OrientableInterface{
			Object: obj,
		},
	}
}

func marshalLevelBar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapLevelBar(obj), nil
}

// NewLevelBar creates a new LevelBar.
func NewLevelBar() *LevelBarClass {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_level_bar_new()

	var _levelBar *LevelBarClass // out

	_levelBar = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*LevelBarClass)

	return _levelBar
}

// NewLevelBarForInterval: utility constructor that creates a new LevelBar for
// the specified interval.
func NewLevelBarForInterval(minValue float64, maxValue float64) *LevelBarClass {
	var _arg1 C.gdouble    // out
	var _arg2 C.gdouble    // out
	var _cret *C.GtkWidget // in

	_arg1 = C.gdouble(minValue)
	_arg2 = C.gdouble(maxValue)

	_cret = C.gtk_level_bar_new_for_interval(_arg1, _arg2)

	var _levelBar *LevelBarClass // out

	_levelBar = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*LevelBarClass)

	return _levelBar
}

// AddOffsetValue adds a new offset marker on @self at the position specified by
// @value. When the bar value is in the interval topped by @value (or between
// @value and LevelBar:max-value in case the offset is the last one on the bar)
// a style class named `level-`@name will be applied when rendering the level
// bar fill. If another offset marker named @name exists, its value will be
// replaced by @value.
func (s *LevelBarClass) AddOffsetValue(name string, value float64) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 *C.gchar       // out
	var _arg2 C.gdouble      // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gdouble(value)

	C.gtk_level_bar_add_offset_value(_arg0, _arg1, _arg2)
}

// Inverted: return the value of the LevelBar:inverted property.
func (s *LevelBarClass) Inverted() bool {
	var _arg0 *C.GtkLevelBar // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_level_bar_get_inverted(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MaxValue returns the value of the LevelBar:max-value property.
func (s *LevelBarClass) MaxValue() float64 {
	var _arg0 *C.GtkLevelBar // out
	var _cret C.gdouble      // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_level_bar_get_max_value(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// MinValue returns the value of the LevelBar:min-value property.
func (s *LevelBarClass) MinValue() float64 {
	var _arg0 *C.GtkLevelBar // out
	var _cret C.gdouble      // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_level_bar_get_min_value(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Mode returns the value of the LevelBar:mode property.
func (s *LevelBarClass) Mode() LevelBarMode {
	var _arg0 *C.GtkLevelBar    // out
	var _cret C.GtkLevelBarMode // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_level_bar_get_mode(_arg0)

	var _levelBarMode LevelBarMode // out

	_levelBarMode = (LevelBarMode)(_cret)

	return _levelBarMode
}

// OffsetValue fetches the value specified for the offset marker @name in @self,
// returning true in case an offset named @name was found.
func (s *LevelBarClass) OffsetValue(name string) (float64, bool) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 *C.gchar       // out
	var _arg2 C.gdouble      // in
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_level_bar_get_offset_value(_arg0, _arg1, &_arg2)

	var _value float64 // out
	var _ok bool       // out

	_value = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

// Value returns the value of the LevelBar:value property.
func (s *LevelBarClass) Value() float64 {
	var _arg0 *C.GtkLevelBar // out
	var _cret C.gdouble      // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_level_bar_get_value(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// RemoveOffsetValue removes an offset marker previously added with
// gtk_level_bar_add_offset_value().
func (s *LevelBarClass) RemoveOffsetValue(name string) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_level_bar_remove_offset_value(_arg0, _arg1)
}

// SetInverted sets the value of the LevelBar:inverted property.
func (s *LevelBarClass) SetInverted(inverted bool) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	if inverted {
		_arg1 = C.TRUE
	}

	C.gtk_level_bar_set_inverted(_arg0, _arg1)
}

// SetMaxValue sets the value of the LevelBar:max-value property.
//
// You probably want to update preexisting level offsets after calling this
// function.
func (s *LevelBarClass) SetMaxValue(value float64) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 C.gdouble      // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	_arg1 = C.gdouble(value)

	C.gtk_level_bar_set_max_value(_arg0, _arg1)
}

// SetMinValue sets the value of the LevelBar:min-value property.
//
// You probably want to update preexisting level offsets after calling this
// function.
func (s *LevelBarClass) SetMinValue(value float64) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 C.gdouble      // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	_arg1 = C.gdouble(value)

	C.gtk_level_bar_set_min_value(_arg0, _arg1)
}

// SetValue sets the value of the LevelBar:value property.
func (s *LevelBarClass) SetValue(value float64) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 C.gdouble      // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(s.Native()))
	_arg1 = C.gdouble(value)

	C.gtk_level_bar_set_value(_arg0, _arg1)
}
