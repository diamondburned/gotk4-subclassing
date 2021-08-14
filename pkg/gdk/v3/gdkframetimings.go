// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_frame_timings_get_type()), F: marshalFrameTimings},
	})
}

// FrameTimings object holds timing information for a single frame of the
// application’s displays. To retrieve FrameTimings objects, use
// gdk_frame_clock_get_timings() or gdk_frame_clock_get_current_timings(). The
// information in FrameTimings is useful for precise synchronization of video
// with the event or audio streams, and for measuring quality metrics for the
// application’s display, such as latency and jitter.
type FrameTimings struct {
	*frameTimings
}

// frameTimings is the struct that's finalized.
type frameTimings struct {
	native *C.GdkFrameTimings
}

func marshalFrameTimings(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &FrameTimings{&frameTimings{(*C.GdkFrameTimings)(unsafe.Pointer(b))}}, nil
}

// Complete: timing information in a FrameTimings is filled in incrementally as
// the frame as drawn and passed off to the window system for processing and
// display to the user. The accessor functions for FrameTimings can return 0 to
// indicate an unavailable value for two reasons: either because the information
// is not yet available, or because it isn't available at all. Once
// gdk_frame_timings_get_complete() returns TRUE for a frame, you can be certain
// that no further values will become available and be stored in the
// FrameTimings.
func (timings *FrameTimings) Complete() bool {
	var _arg0 *C.GdkFrameTimings // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GdkFrameTimings)(gextras.StructNative(unsafe.Pointer(timings)))

	_cret = C.gdk_frame_timings_get_complete(_arg0)
	runtime.KeepAlive(timings)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// FrameCounter gets the frame counter value of the FrameClock when this this
// frame was drawn.
func (timings *FrameTimings) FrameCounter() int64 {
	var _arg0 *C.GdkFrameTimings // out
	var _cret C.gint64           // in

	_arg0 = (*C.GdkFrameTimings)(gextras.StructNative(unsafe.Pointer(timings)))

	_cret = C.gdk_frame_timings_get_frame_counter(_arg0)
	runtime.KeepAlive(timings)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// FrameTime returns the frame time for the frame. This is the time value that
// is typically used to time animations for the frame. See
// gdk_frame_clock_get_frame_time().
func (timings *FrameTimings) FrameTime() int64 {
	var _arg0 *C.GdkFrameTimings // out
	var _cret C.gint64           // in

	_arg0 = (*C.GdkFrameTimings)(gextras.StructNative(unsafe.Pointer(timings)))

	_cret = C.gdk_frame_timings_get_frame_time(_arg0)
	runtime.KeepAlive(timings)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// PredictedPresentationTime gets the predicted time at which this frame will be
// displayed. Although no predicted time may be available, if one is available,
// it will be available while the frame is being generated, in contrast to
// gdk_frame_timings_get_presentation_time(), which is only available after the
// frame has been presented. In general, if you are simply animating, you should
// use gdk_frame_clock_get_frame_time() rather than this function, but this
// function is useful for applications that want exact control over latency. For
// example, a movie player may want this information for Audio/Video
// synchronization.
func (timings *FrameTimings) PredictedPresentationTime() int64 {
	var _arg0 *C.GdkFrameTimings // out
	var _cret C.gint64           // in

	_arg0 = (*C.GdkFrameTimings)(gextras.StructNative(unsafe.Pointer(timings)))

	_cret = C.gdk_frame_timings_get_predicted_presentation_time(_arg0)
	runtime.KeepAlive(timings)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// PresentationTime reurns the presentation time. This is the time at which the
// frame became visible to the user.
func (timings *FrameTimings) PresentationTime() int64 {
	var _arg0 *C.GdkFrameTimings // out
	var _cret C.gint64           // in

	_arg0 = (*C.GdkFrameTimings)(gextras.StructNative(unsafe.Pointer(timings)))

	_cret = C.gdk_frame_timings_get_presentation_time(_arg0)
	runtime.KeepAlive(timings)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// RefreshInterval gets the natural interval between presentation times for the
// display that this frame was displayed on. Frame presentation usually happens
// during the “vertical blanking interval”.
func (timings *FrameTimings) RefreshInterval() int64 {
	var _arg0 *C.GdkFrameTimings // out
	var _cret C.gint64           // in

	_arg0 = (*C.GdkFrameTimings)(gextras.StructNative(unsafe.Pointer(timings)))

	_cret = C.gdk_frame_timings_get_refresh_interval(_arg0)
	runtime.KeepAlive(timings)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}
