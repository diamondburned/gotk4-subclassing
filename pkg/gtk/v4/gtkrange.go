// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_range_get_type()), F: marshalRanger},
	})
}

// RangeOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type RangeOverrider interface {
	AdjustBounds(newValue float64)
	ChangeValue(scroll ScrollType, newValue float64) bool
	RangeBorder(border_ *Border)
	MoveSlider(scroll ScrollType)
	ValueChanged()
}

// Range: GtkRange is the common base class for widgets which visualize an
// adjustment.
//
// Widgets that are derived from GtkRange include gtk.Scale and gtk.Scrollbar.
//
// Apart from signals for monitoring the parameters of the adjustment, GtkRange
// provides properties and methods for setting a “fill level” on range widgets.
// See gtk.Range.SetFillLevel().
type Range struct {
	Widget

	Orientable
	*externglib.Object
}

func wrapRange(obj *externglib.Object) *Range {
	return &Range{
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
		Orientable: Orientable{
			Object: obj,
		},
		Object: obj,
	}
}

func marshalRanger(p uintptr) (interface{}, error) {
	return wrapRange(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Adjustment: get the adjustment which is the “model” object for GtkRange.
func (_range *Range) Adjustment() *Adjustment {
	var _arg0 *C.GtkRange      // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_adjustment(_arg0)
	runtime.KeepAlive(_range)

	var _adjustment *Adjustment // out

	_adjustment = wrapAdjustment(externglib.Take(unsafe.Pointer(_cret)))

	return _adjustment
}

// FillLevel gets the current position of the fill level indicator.
func (_range *Range) FillLevel() float64 {
	var _arg0 *C.GtkRange // out
	var _cret C.double    // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_fill_level(_arg0)
	runtime.KeepAlive(_range)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Flippable gets whether the GtkRange respects text direction.
//
// See gtk.Range.SetFlippable().
func (_range *Range) Flippable() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_flippable(_arg0)
	runtime.KeepAlive(_range)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Inverted gets whether the range is inverted.
//
// See gtk.Range.SetInverted().
func (_range *Range) Inverted() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_inverted(_arg0)
	runtime.KeepAlive(_range)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RangeRect: this function returns the area that contains the range’s trough,
// in coordinates relative to range's origin.
//
// This function is useful mainly for GtkRange subclasses.
func (_range *Range) RangeRect() *gdk.Rectangle {
	var _arg0 *C.GtkRange    // out
	var _arg1 C.GdkRectangle // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	C.gtk_range_get_range_rect(_arg0, &_arg1)
	runtime.KeepAlive(_range)

	var _rangeRect *gdk.Rectangle // out

	_rangeRect = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _rangeRect
}

// RestrictToFillLevel gets whether the range is restricted to the fill level.
func (_range *Range) RestrictToFillLevel() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_restrict_to_fill_level(_arg0)
	runtime.KeepAlive(_range)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RoundDigits gets the number of digits to round the value to when it changes.
//
// See gtk.Range::change-value.
func (_range *Range) RoundDigits() int {
	var _arg0 *C.GtkRange // out
	var _cret C.int       // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_round_digits(_arg0)
	runtime.KeepAlive(_range)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ShowFillLevel gets whether the range displays the fill level graphically.
func (_range *Range) ShowFillLevel() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_show_fill_level(_arg0)
	runtime.KeepAlive(_range)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SliderRange: this function returns sliders range along the long dimension, in
// widget->window coordinates.
//
// This function is useful mainly for GtkRange subclasses.
func (_range *Range) SliderRange() (sliderStart int, sliderEnd int) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.int       // in
	var _arg2 C.int       // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	C.gtk_range_get_slider_range(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(_range)

	var _sliderStart int // out
	var _sliderEnd int   // out

	_sliderStart = int(_arg1)
	_sliderEnd = int(_arg2)

	return _sliderStart, _sliderEnd
}

// SliderSizeFixed: this function is useful mainly for GtkRange subclasses.
//
// See gtk.Range.SetSliderSizeFixed().
func (_range *Range) SliderSizeFixed() bool {
	var _arg0 *C.GtkRange // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_slider_size_fixed(_arg0)
	runtime.KeepAlive(_range)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Value gets the current value of the range.
func (_range *Range) Value() float64 {
	var _arg0 *C.GtkRange // out
	var _cret C.double    // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_value(_arg0)
	runtime.KeepAlive(_range)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetAdjustment sets the adjustment to be used as the “model” object for the
// GtkRange
//
// The adjustment indicates the current range value, the minimum and maximum
// range values, the step/page increments used for keybindings and scrolling,
// and the page size.
//
// The page size is normally 0 for GtkScale and nonzero for GtkScrollbar, and
// indicates the size of the visible area of the widget being scrolled. The page
// size affects the size of the scrollbar slider.
//
// The function takes the following parameters:
//
//    - adjustment: GtkAdjustment.
//
func (_range *Range) SetAdjustment(adjustment *Adjustment) {
	var _arg0 *C.GtkRange      // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_range_set_adjustment(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(adjustment)
}

// SetFillLevel: set the new position of the fill level indicator.
//
// The “fill level” is probably best described by its most prominent use case,
// which is an indicator for the amount of pre-buffering in a streaming media
// player. In that use case, the value of the range would indicate the current
// play position, and the fill level would be the position up to which the
// file/stream has been downloaded.
//
// This amount of prebuffering can be displayed on the range’s trough and is
// themeable separately from the trough. To enable fill level display, use
// gtk.Range.SetShowFillLevel(). The range defaults to not showing the fill
// level.
//
// Additionally, it’s possible to restrict the range’s slider position to values
// which are smaller than the fill level. This is controlled by
// gtk.Range.SetRestrictToFillLevel() and is by default enabled.
//
// The function takes the following parameters:
//
//    - fillLevel: new position of the fill level indicator.
//
func (_range *Range) SetFillLevel(fillLevel float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.double    // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	_arg1 = C.double(fillLevel)

	C.gtk_range_set_fill_level(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(fillLevel)
}

// SetFlippable sets whether the GtkRange respects text direction.
//
// If a range is flippable, it will switch its direction if it is horizontal and
// its direction is GTK_TEXT_DIR_RTL.
//
// See gtk.Widget.GetDirection().
//
// The function takes the following parameters:
//
//    - flippable: TRUE to make the range flippable.
//
func (_range *Range) SetFlippable(flippable bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	if flippable {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_flippable(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(flippable)
}

// SetIncrements sets the step and page sizes for the range.
//
// The step size is used when the user clicks the GtkScrollbar arrows or moves a
// GtkScale via arrow keys. The page size is used for example when moving via
// Page Up or Page Down keys.
//
// The function takes the following parameters:
//
//    - step size.
//    - page size.
//
func (_range *Range) SetIncrements(step, page float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.double    // out
	var _arg2 C.double    // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	_arg1 = C.double(step)
	_arg2 = C.double(page)

	C.gtk_range_set_increments(_arg0, _arg1, _arg2)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(step)
	runtime.KeepAlive(page)
}

// SetInverted sets whether to invert the range.
//
// Ranges normally move from lower to higher values as the slider moves from top
// to bottom or left to right. Inverted ranges have higher values at the top or
// on the right rather than on the bottom or left.
//
// The function takes the following parameters:
//
//    - setting: TRUE to invert the range.
//
func (_range *Range) SetInverted(setting bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_inverted(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(setting)
}

// SetRange sets the allowable values in the GtkRange.
//
// The range value is clamped to be between min and max. (If the range has a
// non-zero page size, it is clamped between min and max - page-size.).
//
// The function takes the following parameters:
//
//    - min: minimum range value.
//    - max: maximum range value.
//
func (_range *Range) SetRange(min, max float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.double    // out
	var _arg2 C.double    // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	_arg1 = C.double(min)
	_arg2 = C.double(max)

	C.gtk_range_set_range(_arg0, _arg1, _arg2)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
}

// SetRestrictToFillLevel sets whether the slider is restricted to the fill
// level.
//
// See gtk.Range.SetFillLevel() for a general description of the fill level
// concept.
//
// The function takes the following parameters:
//
//    - restrictToFillLevel: whether the fill level restricts slider movement.
//
func (_range *Range) SetRestrictToFillLevel(restrictToFillLevel bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	if restrictToFillLevel {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_restrict_to_fill_level(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(restrictToFillLevel)
}

// SetRoundDigits sets the number of digits to round the value to when it
// changes.
//
// See gtk.Range::change-value.
//
// The function takes the following parameters:
//
//    - roundDigits: precision in digits, or -1.
//
func (_range *Range) SetRoundDigits(roundDigits int) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.int       // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	_arg1 = C.int(roundDigits)

	C.gtk_range_set_round_digits(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(roundDigits)
}

// SetShowFillLevel sets whether a graphical fill level is show on the trough.
//
// See gtk.Range.SetFillLevel() for a general description of the fill level
// concept.
//
// The function takes the following parameters:
//
//    - showFillLevel: whether a fill level indicator graphics is shown.
//
func (_range *Range) SetShowFillLevel(showFillLevel bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	if showFillLevel {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_show_fill_level(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(showFillLevel)
}

// SetSliderSizeFixed sets whether the range’s slider has a fixed size, or a
// size that depends on its adjustment’s page size.
//
// This function is useful mainly for GtkRange subclasses.
//
// The function takes the following parameters:
//
//    - sizeFixed: TRUE to make the slider size constant.
//
func (_range *Range) SetSliderSizeFixed(sizeFixed bool) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	if sizeFixed {
		_arg1 = C.TRUE
	}

	C.gtk_range_set_slider_size_fixed(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(sizeFixed)
}

// SetValue sets the current value of the range.
//
// If the value is outside the minimum or maximum range values, it will be
// clamped to fit inside them. The range emits the gtk.Range::value-changed
// signal if the value changes.
//
// The function takes the following parameters:
//
//    - value: new value of the range.
//
func (_range *Range) SetValue(value float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.double    // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	_arg1 = C.double(value)

	C.gtk_range_set_value(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(value)
}

// ConnectAdjustBounds: emitted before clamping a value, to give the application
// a chance to adjust the bounds.
func (_range *Range) ConnectAdjustBounds(f func(value float64)) externglib.SignalHandle {
	return _range.Connect("adjust-bounds", f)
}

// ConnectChangeValue: emitted when a scroll action is performed on a range.
//
// It allows an application to determine the type of scroll event that occurred
// and the resultant new value. The application can handle the event itself and
// return TRUE to prevent further processing. Or, by returning FALSE, it can
// pass the event to other handlers until the default GTK handler is reached.
//
// The value parameter is unrounded. An application that overrides the
// ::change-value signal is responsible for clamping the value to the desired
// number of decimal digits; the default GTK handler clamps the value based on
// gtk.Range:round-digits.
func (_range *Range) ConnectChangeValue(f func(scroll ScrollType, value float64) bool) externglib.SignalHandle {
	return _range.Connect("change-value", f)
}

// ConnectMoveSlider: virtual function that moves the slider.
//
// Used for keybindings.
func (_range *Range) ConnectMoveSlider(f func(step ScrollType)) externglib.SignalHandle {
	return _range.Connect("move-slider", f)
}

// ConnectValueChanged: emitted when the range value changes.
func (_range *Range) ConnectValueChanged(f func()) externglib.SignalHandle {
	return _range.Connect("value-changed", f)
}
