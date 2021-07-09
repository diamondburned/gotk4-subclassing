// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_aspect_frame_get_type()), F: marshalAspectFrame},
	})
}

// AspectFrame: `GtkAspectFrame` preserves the aspect ratio of its child.
//
// The frame can respect the aspect ratio of the child widget, or use its own
// aspect ratio.
//
//
// CSS nodes
//
// `GtkAspectFrame` uses a CSS node with name `frame`.
type AspectFrame interface {
	gextras.Objector

	// Child gets the child widget of @self.
	Child() *WidgetClass
	// ObeyChild returns whether the child's size request should override the
	// set aspect ratio of the `GtkAspectFrame`.
	ObeyChild() bool
	// Ratio returns the desired aspect ratio of the child.
	Ratio() float32
	// Xalign returns the horizontal alignment of the child within the
	// allocation of the `GtkAspectFrame`.
	Xalign() float32
	// Yalign returns the vertical alignment of the child within the allocation
	// of the `GtkAspectFrame`.
	Yalign() float32
	// SetChild sets the child widget of @self.
	SetChild(child Widget)
	// SetObeyChild sets whether the aspect ratio of the child's size request
	// should override the set aspect ratio of the `GtkAspectFrame`.
	SetObeyChild(obeyChild bool)
	// SetRatio sets the desired aspect ratio of the child.
	SetRatio(ratio float32)
	// SetXalign sets the horizontal alignment of the child within the
	// allocation of the `GtkAspectFrame`.
	SetXalign(xalign float32)
	// SetYalign sets the vertical alignment of the child within the allocation
	// of the `GtkAspectFrame`.
	SetYalign(yalign float32)
}

// AspectFrameClass implements the AspectFrame interface.
type AspectFrameClass struct {
	*externglib.Object
	WidgetClass
	AccessibleInterface
	BuildableInterface
	ConstraintTargetInterface
}

var _ AspectFrame = (*AspectFrameClass)(nil)

func wrapAspectFrame(obj *externglib.Object) AspectFrame {
	return &AspectFrameClass{
		Object: obj,
		WidgetClass: WidgetClass{
			InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
			AccessibleInterface: AccessibleInterface{
				Object: obj,
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
			ConstraintTargetInterface: ConstraintTargetInterface{
				Object: obj,
			},
		},
		AccessibleInterface: AccessibleInterface{
			Object: obj,
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		ConstraintTargetInterface: ConstraintTargetInterface{
			Object: obj,
		},
	}
}

func marshalAspectFrame(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapAspectFrame(obj), nil
}

// NewAspectFrame: create a new `GtkAspectFrame`.
func NewAspectFrame(xalign float32, yalign float32, ratio float32, obeyChild bool) *AspectFrameClass {
	var _arg1 C.float      // out
	var _arg2 C.float      // out
	var _arg3 C.float      // out
	var _arg4 C.gboolean   // out
	var _cret *C.GtkWidget // in

	_arg1 = C.float(xalign)
	_arg2 = C.float(yalign)
	_arg3 = C.float(ratio)
	if obeyChild {
		_arg4 = C.TRUE
	}

	_cret = C.gtk_aspect_frame_new(_arg1, _arg2, _arg3, _arg4)

	var _aspectFrame *AspectFrameClass // out

	_aspectFrame = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*AspectFrameClass)

	return _aspectFrame
}

// Child gets the child widget of @self.
func (s *AspectFrameClass) Child() *WidgetClass {
	var _arg0 *C.GtkAspectFrame // out
	var _cret *C.GtkWidget      // in

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_aspect_frame_get_child(_arg0)

	var _widget *WidgetClass // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*WidgetClass)

	return _widget
}

// ObeyChild returns whether the child's size request should override the set
// aspect ratio of the `GtkAspectFrame`.
func (s *AspectFrameClass) ObeyChild() bool {
	var _arg0 *C.GtkAspectFrame // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_aspect_frame_get_obey_child(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Ratio returns the desired aspect ratio of the child.
func (s *AspectFrameClass) Ratio() float32 {
	var _arg0 *C.GtkAspectFrame // out
	var _cret C.float           // in

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_aspect_frame_get_ratio(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Xalign returns the horizontal alignment of the child within the allocation of
// the `GtkAspectFrame`.
func (s *AspectFrameClass) Xalign() float32 {
	var _arg0 *C.GtkAspectFrame // out
	var _cret C.float           // in

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_aspect_frame_get_xalign(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Yalign returns the vertical alignment of the child within the allocation of
// the `GtkAspectFrame`.
func (s *AspectFrameClass) Yalign() float32 {
	var _arg0 *C.GtkAspectFrame // out
	var _cret C.float           // in

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(s.Native()))

	_cret = C.gtk_aspect_frame_get_yalign(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// SetChild sets the child widget of @self.
func (s *AspectFrameClass) SetChild(child Widget) {
	var _arg0 *C.GtkAspectFrame // out
	var _arg1 *C.GtkWidget      // out

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_aspect_frame_set_child(_arg0, _arg1)
}

// SetObeyChild sets whether the aspect ratio of the child's size request should
// override the set aspect ratio of the `GtkAspectFrame`.
func (s *AspectFrameClass) SetObeyChild(obeyChild bool) {
	var _arg0 *C.GtkAspectFrame // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(s.Native()))
	if obeyChild {
		_arg1 = C.TRUE
	}

	C.gtk_aspect_frame_set_obey_child(_arg0, _arg1)
}

// SetRatio sets the desired aspect ratio of the child.
func (s *AspectFrameClass) SetRatio(ratio float32) {
	var _arg0 *C.GtkAspectFrame // out
	var _arg1 C.float           // out

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(s.Native()))
	_arg1 = C.float(ratio)

	C.gtk_aspect_frame_set_ratio(_arg0, _arg1)
}

// SetXalign sets the horizontal alignment of the child within the allocation of
// the `GtkAspectFrame`.
func (s *AspectFrameClass) SetXalign(xalign float32) {
	var _arg0 *C.GtkAspectFrame // out
	var _arg1 C.float           // out

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(s.Native()))
	_arg1 = C.float(xalign)

	C.gtk_aspect_frame_set_xalign(_arg0, _arg1)
}

// SetYalign sets the vertical alignment of the child within the allocation of
// the `GtkAspectFrame`.
func (s *AspectFrameClass) SetYalign(yalign float32) {
	var _arg0 *C.GtkAspectFrame // out
	var _arg1 C.float           // out

	_arg0 = (*C.GtkAspectFrame)(unsafe.Pointer(s.Native()))
	_arg1 = C.float(yalign)

	C.gtk_aspect_frame_set_yalign(_arg0, _arg1)
}
