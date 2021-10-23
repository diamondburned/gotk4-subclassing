// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_level_bar_get_type()), F: marshalLevelBarrer},
	})
}

// LEVEL_BAR_OFFSET_FULL: name used for the stock full offset included by
// LevelBar.
const LEVEL_BAR_OFFSET_FULL = "full"

// LEVEL_BAR_OFFSET_HIGH: name used for the stock high offset included by
// LevelBar.
const LEVEL_BAR_OFFSET_HIGH = "high"

// LEVEL_BAR_OFFSET_LOW: name used for the stock low offset included by
// LevelBar.
const LEVEL_BAR_OFFSET_LOW = "low"

// LevelBarOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type LevelBarOverrider interface {
	OffsetChanged(name string)
}

// LevelBar is a bar widget that can be used as a level indicator. Typical use
// cases are displaying the strength of a password, or showing the charge level
// of a battery.
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
type LevelBar struct {
	Widget

	Orientable
	*externglib.Object
}

func wrapLevelBar(obj *externglib.Object) *LevelBar {
	return &LevelBar{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			ImplementorIface: atk.ImplementorIface{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			Object: obj,
		},
		Orientable: Orientable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalLevelBarrer(p uintptr) (interface{}, error) {
	return wrapLevelBar(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewLevelBar creates a new LevelBar.
func NewLevelBar() *LevelBar {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_level_bar_new()

	var _levelBar *LevelBar // out

	_levelBar = wrapLevelBar(externglib.Take(unsafe.Pointer(_cret)))

	return _levelBar
}

// NewLevelBarForInterval: utility constructor that creates a new LevelBar for
// the specified interval.
//
// The function takes the following parameters:
//
//    - minValue: positive value.
//    - maxValue: positive value.
//
func NewLevelBarForInterval(minValue, maxValue float64) *LevelBar {
	var _arg1 C.gdouble    // out
	var _arg2 C.gdouble    // out
	var _cret *C.GtkWidget // in

	_arg1 = C.gdouble(minValue)
	_arg2 = C.gdouble(maxValue)

	_cret = C.gtk_level_bar_new_for_interval(_arg1, _arg2)
	runtime.KeepAlive(minValue)
	runtime.KeepAlive(maxValue)

	var _levelBar *LevelBar // out

	_levelBar = wrapLevelBar(externglib.Take(unsafe.Pointer(_cret)))

	return _levelBar
}

// AddOffsetValue adds a new offset marker on self at the position specified by
// value. When the bar value is in the interval topped by value (or between
// value and LevelBar:max-value in case the offset is the last one on the bar) a
// style class named level-name will be applied when rendering the level bar
// fill. If another offset marker named name exists, its value will be replaced
// by value.
//
// The function takes the following parameters:
//
//    - name of the new offset.
//    - value for the new offset.
//
func (self *LevelBar) AddOffsetValue(name string, value float64) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 *C.gchar       // out
	var _arg2 C.gdouble      // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gdouble(value)

	C.gtk_level_bar_add_offset_value(_arg0, _arg1, _arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
}

// Inverted: return the value of the LevelBar:inverted property.
func (self *LevelBar) Inverted() bool {
	var _arg0 *C.GtkLevelBar // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_level_bar_get_inverted(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MaxValue returns the value of the LevelBar:max-value property.
func (self *LevelBar) MaxValue() float64 {
	var _arg0 *C.GtkLevelBar // out
	var _cret C.gdouble      // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_level_bar_get_max_value(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// MinValue returns the value of the LevelBar:min-value property.
func (self *LevelBar) MinValue() float64 {
	var _arg0 *C.GtkLevelBar // out
	var _cret C.gdouble      // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_level_bar_get_min_value(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Mode returns the value of the LevelBar:mode property.
func (self *LevelBar) Mode() LevelBarMode {
	var _arg0 *C.GtkLevelBar    // out
	var _cret C.GtkLevelBarMode // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_level_bar_get_mode(_arg0)
	runtime.KeepAlive(self)

	var _levelBarMode LevelBarMode // out

	_levelBarMode = LevelBarMode(_cret)

	return _levelBarMode
}

// OffsetValue fetches the value specified for the offset marker name in self,
// returning TRUE in case an offset named name was found.
//
// The function takes the following parameters:
//
//    - name of an offset in the bar.
//
func (self *LevelBar) OffsetValue(name string) (float64, bool) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 *C.gchar       // out
	var _arg2 C.gdouble      // in
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))
	if name != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gtk_level_bar_get_offset_value(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(name)

	var _value float64 // out
	var _ok bool       // out

	_value = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

// Value returns the value of the LevelBar:value property.
func (self *LevelBar) Value() float64 {
	var _arg0 *C.GtkLevelBar // out
	var _cret C.gdouble      // in

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_level_bar_get_value(_arg0)
	runtime.KeepAlive(self)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// RemoveOffsetValue removes an offset marker previously added with
// gtk_level_bar_add_offset_value().
//
// The function takes the following parameters:
//
//    - name of an offset in the bar.
//
func (self *LevelBar) RemoveOffsetValue(name string) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 *C.gchar       // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))
	if name != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_level_bar_remove_offset_value(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(name)
}

// SetInverted sets the value of the LevelBar:inverted property.
//
// The function takes the following parameters:
//
//    - inverted: TRUE to invert the level bar.
//
func (self *LevelBar) SetInverted(inverted bool) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))
	if inverted {
		_arg1 = C.TRUE
	}

	C.gtk_level_bar_set_inverted(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(inverted)
}

// SetMaxValue sets the value of the LevelBar:max-value property.
//
// You probably want to update preexisting level offsets after calling this
// function.
//
// The function takes the following parameters:
//
//    - value: positive value.
//
func (self *LevelBar) SetMaxValue(value float64) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 C.gdouble      // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))
	_arg1 = C.gdouble(value)

	C.gtk_level_bar_set_max_value(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetMinValue sets the value of the LevelBar:min-value property.
//
// You probably want to update preexisting level offsets after calling this
// function.
//
// The function takes the following parameters:
//
//    - value: positive value.
//
func (self *LevelBar) SetMinValue(value float64) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 C.gdouble      // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))
	_arg1 = C.gdouble(value)

	C.gtk_level_bar_set_min_value(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// SetMode sets the value of the LevelBar:mode property.
//
// The function takes the following parameters:
//
//    - mode: LevelBarMode.
//
func (self *LevelBar) SetMode(mode LevelBarMode) {
	var _arg0 *C.GtkLevelBar    // out
	var _arg1 C.GtkLevelBarMode // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))
	_arg1 = C.GtkLevelBarMode(mode)

	C.gtk_level_bar_set_mode(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(mode)
}

// SetValue sets the value of the LevelBar:value property.
//
// The function takes the following parameters:
//
//    - value in the interval between LevelBar:min-value and
//    LevelBar:max-value.
//
func (self *LevelBar) SetValue(value float64) {
	var _arg0 *C.GtkLevelBar // out
	var _arg1 C.gdouble      // out

	_arg0 = (*C.GtkLevelBar)(unsafe.Pointer(self.Native()))
	_arg1 = C.gdouble(value)

	C.gtk_level_bar_set_value(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(value)
}

// ConnectOffsetChanged: emitted when an offset specified on the bar changes
// value as an effect to gtk_level_bar_add_offset_value() being called.
//
// The signal supports detailed connections; you can connect to the detailed
// signal "changed::x" in order to only receive callbacks when the value of
// offset "x" changes.
func (self *LevelBar) ConnectOffsetChanged(f func(name string)) externglib.SignalHandle {
	return self.Connect("offset-changed", f)
}
