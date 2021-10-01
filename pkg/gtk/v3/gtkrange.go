// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
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
	RangeSizeRequest(orientation Orientation, minimum *int, natural *int)
	MoveSlider(scroll ScrollType)
	ValueChanged()
}

// Range is the common base class for widgets which visualize an adjustment, e.g
// Scale or Scrollbar.
//
// Apart from signals for monitoring the parameters of the adjustment, Range
// provides properties and methods for influencing the sensitivity of the
// “steppers”. It also provides properties and methods for setting a “fill
// level” on range widgets. See gtk_range_set_fill_level().
type Range struct {
	Widget

	Orientable
	*externglib.Object
}

// Ranger describes Range's abstract methods.
type Ranger interface {
	externglib.Objector

	// Adjustment: get the Adjustment which is the “model” object for Range.
	Adjustment() *Adjustment
	// FillLevel gets the current position of the fill level indicator.
	FillLevel() float64
	// Flippable gets the value set by gtk_range_set_flippable().
	Flippable() bool
	// Inverted gets the value set by gtk_range_set_inverted().
	Inverted() bool
	// LowerStepperSensitivity gets the sensitivity policy for the stepper that
	// points to the 'lower' end of the GtkRange’s adjustment.
	LowerStepperSensitivity() SensitivityType
	// MinSliderSize: this function is useful mainly for Range subclasses.
	MinSliderSize() int
	// RangeRect: this function returns the area that contains the range’s
	// trough and its steppers, in widget->window coordinates.
	RangeRect() gdk.Rectangle
	// RestrictToFillLevel gets whether the range is restricted to the fill
	// level.
	RestrictToFillLevel() bool
	// RoundDigits gets the number of digits to round the value to when it
	// changes.
	RoundDigits() int
	// ShowFillLevel gets whether the range displays the fill level graphically.
	ShowFillLevel() bool
	// SliderRange: this function returns sliders range along the long
	// dimension, in widget->window coordinates.
	SliderRange() (sliderStart int, sliderEnd int)
	// SliderSizeFixed: this function is useful mainly for Range subclasses.
	SliderSizeFixed() bool
	// UpperStepperSensitivity gets the sensitivity policy for the stepper that
	// points to the 'upper' end of the GtkRange’s adjustment.
	UpperStepperSensitivity() SensitivityType
	// Value gets the current value of the range.
	Value() float64
	// SetAdjustment sets the adjustment to be used as the “model” object for
	// this range widget.
	SetAdjustment(adjustment *Adjustment)
	// SetFillLevel: set the new position of the fill level indicator.
	SetFillLevel(fillLevel float64)
	// SetFlippable: if a range is flippable, it will switch its direction if it
	// is horizontal and its direction is GTK_TEXT_DIR_RTL.
	SetFlippable(flippable bool)
	// SetIncrements sets the step and page sizes for the range.
	SetIncrements(step float64, page float64)
	// SetInverted ranges normally move from lower to higher values as the
	// slider moves from top to bottom or left to right.
	SetInverted(setting bool)
	// SetLowerStepperSensitivity sets the sensitivity policy for the stepper
	// that points to the 'lower' end of the GtkRange’s adjustment.
	SetLowerStepperSensitivity(sensitivity SensitivityType)
	// SetMinSliderSize sets the minimum size of the range’s slider.
	SetMinSliderSize(minSize int)
	// SetRange sets the allowable values in the Range, and clamps the range
	// value to be between min and max.
	SetRange(min float64, max float64)
	// SetRestrictToFillLevel sets whether the slider is restricted to the fill
	// level.
	SetRestrictToFillLevel(restrictToFillLevel bool)
	// SetRoundDigits sets the number of digits to round the value to when it
	// changes.
	SetRoundDigits(roundDigits int)
	// SetShowFillLevel sets whether a graphical fill level is show on the
	// trough.
	SetShowFillLevel(showFillLevel bool)
	// SetSliderSizeFixed sets whether the range’s slider has a fixed size, or a
	// size that depends on its adjustment’s page size.
	SetSliderSizeFixed(sizeFixed bool)
	// SetUpperStepperSensitivity sets the sensitivity policy for the stepper
	// that points to the 'upper' end of the GtkRange’s adjustment.
	SetUpperStepperSensitivity(sensitivity SensitivityType)
	// SetValue sets the current value of the range; if the value is outside the
	// minimum or maximum range values, it will be clamped to fit inside them.
	SetValue(value float64)
}

var _ Ranger = (*Range)(nil)

func wrapRange(obj *externglib.Object) *Range {
	return &Range{
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

func marshalRanger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRange(obj), nil
}

// Adjustment: get the Adjustment which is the “model” object for Range. See
// gtk_range_set_adjustment() for details. The return value does not have a
// reference added, so should not be unreferenced.
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
	var _cret C.gdouble   // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_fill_level(_arg0)
	runtime.KeepAlive(_range)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Flippable gets the value set by gtk_range_set_flippable().
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

// Inverted gets the value set by gtk_range_set_inverted().
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

// LowerStepperSensitivity gets the sensitivity policy for the stepper that
// points to the 'lower' end of the GtkRange’s adjustment.
func (_range *Range) LowerStepperSensitivity() SensitivityType {
	var _arg0 *C.GtkRange          // out
	var _cret C.GtkSensitivityType // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_lower_stepper_sensitivity(_arg0)
	runtime.KeepAlive(_range)

	var _sensitivityType SensitivityType // out

	_sensitivityType = SensitivityType(_cret)

	return _sensitivityType
}

// MinSliderSize: this function is useful mainly for Range subclasses.
//
// See gtk_range_set_min_slider_size().
//
// Deprecated: Use the min-height/min-width CSS properties on the slider node.
func (_range *Range) MinSliderSize() int {
	var _arg0 *C.GtkRange // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_min_slider_size(_arg0)
	runtime.KeepAlive(_range)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// RangeRect: this function returns the area that contains the range’s trough
// and its steppers, in widget->window coordinates.
//
// This function is useful mainly for Range subclasses.
func (_range *Range) RangeRect() gdk.Rectangle {
	var _arg0 *C.GtkRange    // out
	var _arg1 C.GdkRectangle // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	C.gtk_range_get_range_rect(_arg0, &_arg1)
	runtime.KeepAlive(_range)

	var _rangeRect gdk.Rectangle // out

	_rangeRect = *(*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

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
// See Range::change-value.
func (_range *Range) RoundDigits() int {
	var _arg0 *C.GtkRange // out
	var _cret C.gint      // in

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
// This function is useful mainly for Range subclasses.
func (_range *Range) SliderRange() (sliderStart int, sliderEnd int) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gint      // in
	var _arg2 C.gint      // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	C.gtk_range_get_slider_range(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(_range)

	var _sliderStart int // out
	var _sliderEnd int   // out

	_sliderStart = int(_arg1)
	_sliderEnd = int(_arg2)

	return _sliderStart, _sliderEnd
}

// SliderSizeFixed: this function is useful mainly for Range subclasses.
//
// See gtk_range_set_slider_size_fixed().
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

// UpperStepperSensitivity gets the sensitivity policy for the stepper that
// points to the 'upper' end of the GtkRange’s adjustment.
func (_range *Range) UpperStepperSensitivity() SensitivityType {
	var _arg0 *C.GtkRange          // out
	var _cret C.GtkSensitivityType // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_upper_stepper_sensitivity(_arg0)
	runtime.KeepAlive(_range)

	var _sensitivityType SensitivityType // out

	_sensitivityType = SensitivityType(_cret)

	return _sensitivityType
}

// Value gets the current value of the range.
func (_range *Range) Value() float64 {
	var _arg0 *C.GtkRange // out
	var _cret C.gdouble   // in

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))

	_cret = C.gtk_range_get_value(_arg0)
	runtime.KeepAlive(_range)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// SetAdjustment sets the adjustment to be used as the “model” object for this
// range widget. The adjustment indicates the current range value, the minimum
// and maximum range values, the step/page increments used for keybindings and
// scrolling, and the page size. The page size is normally 0 for Scale and
// nonzero for Scrollbar, and indicates the size of the visible area of the
// widget being scrolled. The page size affects the size of the scrollbar
// slider.
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
// gtk_range_set_show_fill_level(). The range defaults to not showing the fill
// level.
//
// Additionally, it’s possible to restrict the range’s slider position to values
// which are smaller than the fill level. This is controller by
// gtk_range_set_restrict_to_fill_level() and is by default enabled.
func (_range *Range) SetFillLevel(fillLevel float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gdouble   // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	_arg1 = C.gdouble(fillLevel)

	C.gtk_range_set_fill_level(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(fillLevel)
}

// SetFlippable: if a range is flippable, it will switch its direction if it is
// horizontal and its direction is GTK_TEXT_DIR_RTL.
//
// See gtk_widget_get_direction().
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

// SetIncrements sets the step and page sizes for the range. The step size is
// used when the user clicks the Scrollbar arrows or moves Scale via arrow keys.
// The page size is used for example when moving via Page Up or Page Down keys.
func (_range *Range) SetIncrements(step float64, page float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gdouble   // out
	var _arg2 C.gdouble   // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	_arg1 = C.gdouble(step)
	_arg2 = C.gdouble(page)

	C.gtk_range_set_increments(_arg0, _arg1, _arg2)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(step)
	runtime.KeepAlive(page)
}

// SetInverted ranges normally move from lower to higher values as the slider
// moves from top to bottom or left to right. Inverted ranges have higher values
// at the top or on the right rather than on the bottom or left.
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

// SetLowerStepperSensitivity sets the sensitivity policy for the stepper that
// points to the 'lower' end of the GtkRange’s adjustment.
func (_range *Range) SetLowerStepperSensitivity(sensitivity SensitivityType) {
	var _arg0 *C.GtkRange          // out
	var _arg1 C.GtkSensitivityType // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	_arg1 = C.GtkSensitivityType(sensitivity)

	C.gtk_range_set_lower_stepper_sensitivity(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(sensitivity)
}

// SetMinSliderSize sets the minimum size of the range’s slider.
//
// This function is useful mainly for Range subclasses.
//
// Deprecated: Use the min-height/min-width CSS properties on the slider node.
func (_range *Range) SetMinSliderSize(minSize int) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	_arg1 = C.gint(minSize)

	C.gtk_range_set_min_slider_size(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(minSize)
}

// SetRange sets the allowable values in the Range, and clamps the range value
// to be between min and max. (If the range has a non-zero page size, it is
// clamped between min and max - page-size.).
func (_range *Range) SetRange(min float64, max float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gdouble   // out
	var _arg2 C.gdouble   // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	_arg1 = C.gdouble(min)
	_arg2 = C.gdouble(max)

	C.gtk_range_set_range(_arg0, _arg1, _arg2)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
}

// SetRestrictToFillLevel sets whether the slider is restricted to the fill
// level. See gtk_range_set_fill_level() for a general description of the fill
// level concept.
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
// changes. See Range::change-value.
func (_range *Range) SetRoundDigits(roundDigits int) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	_arg1 = C.gint(roundDigits)

	C.gtk_range_set_round_digits(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(roundDigits)
}

// SetShowFillLevel sets whether a graphical fill level is show on the trough.
// See gtk_range_set_fill_level() for a general description of the fill level
// concept.
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
// This function is useful mainly for Range subclasses.
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

// SetUpperStepperSensitivity sets the sensitivity policy for the stepper that
// points to the 'upper' end of the GtkRange’s adjustment.
func (_range *Range) SetUpperStepperSensitivity(sensitivity SensitivityType) {
	var _arg0 *C.GtkRange          // out
	var _arg1 C.GtkSensitivityType // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	_arg1 = C.GtkSensitivityType(sensitivity)

	C.gtk_range_set_upper_stepper_sensitivity(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(sensitivity)
}

// SetValue sets the current value of the range; if the value is outside the
// minimum or maximum range values, it will be clamped to fit inside them. The
// range emits the Range::value-changed signal if the value changes.
func (_range *Range) SetValue(value float64) {
	var _arg0 *C.GtkRange // out
	var _arg1 C.gdouble   // out

	_arg0 = (*C.GtkRange)(unsafe.Pointer(_range.Native()))
	_arg1 = C.gdouble(value)

	C.gtk_range_set_value(_arg0, _arg1)
	runtime.KeepAlive(_range)
	runtime.KeepAlive(value)
}
