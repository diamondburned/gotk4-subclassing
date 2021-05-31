// Code generated by girgen. DO NOT EDIT.

package xlib

import (
	"unsafe"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
//
import "C"

type Atom uint32

type Colormap uint32

type Cursor uint32

type Drawable uint32

type GC interface{}

type KeyCode byte

type KeySym uint32

type Picture uint32

type Time uint32

type VisualID uint32

type Window uint32

type XID uint32

type Pixmap uint32

func OpenDisplay() {

	C.XOpenDisplay()
}

type Display struct {
	native C.Display
}

// WrapDisplay wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDisplay(ptr unsafe.Pointer) *Display {
	if ptr == nil {
		return nil
	}

	return (*Display)(ptr)
}

func marshalDisplay(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapDisplay(unsafe.Pointer(b))
}

// Native returns the underlying C source pointer.
func (d *Display) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

type Screen struct {
	native C.Screen
}

// WrapScreen wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapScreen(ptr unsafe.Pointer) *Screen {
	if ptr == nil {
		return nil
	}

	return (*Screen)(ptr)
}

func marshalScreen(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapScreen(unsafe.Pointer(b))
}

// Native returns the underlying C source pointer.
func (s *Screen) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

type Visual struct {
	native C.Visual
}

// WrapVisual wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapVisual(ptr unsafe.Pointer) *Visual {
	if ptr == nil {
		return nil
	}

	return (*Visual)(ptr)
}

func marshalVisual(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapVisual(unsafe.Pointer(b))
}

// Native returns the underlying C source pointer.
func (v *Visual) Native() unsafe.Pointer {
	return unsafe.Pointer(&v.native)
}

type XConfigureEvent struct {
	native C.XConfigureEvent
}

// WrapXConfigureEvent wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapXConfigureEvent(ptr unsafe.Pointer) *XConfigureEvent {
	if ptr == nil {
		return nil
	}

	return (*XConfigureEvent)(ptr)
}

func marshalXConfigureEvent(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapXConfigureEvent(unsafe.Pointer(b))
}

// Native returns the underlying C source pointer.
func (x *XConfigureEvent) Native() unsafe.Pointer {
	return unsafe.Pointer(&x.native)
}

type XImage struct {
	native C.XImage
}

// WrapXImage wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapXImage(ptr unsafe.Pointer) *XImage {
	if ptr == nil {
		return nil
	}

	return (*XImage)(ptr)
}

func marshalXImage(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapXImage(unsafe.Pointer(b))
}

// Native returns the underlying C source pointer.
func (x *XImage) Native() unsafe.Pointer {
	return unsafe.Pointer(&x.native)
}

type XFontStruct struct {
	native C.XFontStruct
}

// WrapXFontStruct wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapXFontStruct(ptr unsafe.Pointer) *XFontStruct {
	if ptr == nil {
		return nil
	}

	return (*XFontStruct)(ptr)
}

func marshalXFontStruct(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapXFontStruct(unsafe.Pointer(b))
}

// Native returns the underlying C source pointer.
func (x *XFontStruct) Native() unsafe.Pointer {
	return unsafe.Pointer(&x.native)
}

type XTrapezoid struct {
	native C.XTrapezoid
}

// WrapXTrapezoid wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapXTrapezoid(ptr unsafe.Pointer) *XTrapezoid {
	if ptr == nil {
		return nil
	}

	return (*XTrapezoid)(ptr)
}

func marshalXTrapezoid(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapXTrapezoid(unsafe.Pointer(b))
}

// Native returns the underlying C source pointer.
func (x *XTrapezoid) Native() unsafe.Pointer {
	return unsafe.Pointer(&x.native)
}

type XVisualInfo struct {
	native C.XVisualInfo
}

// WrapXVisualInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapXVisualInfo(ptr unsafe.Pointer) *XVisualInfo {
	if ptr == nil {
		return nil
	}

	return (*XVisualInfo)(ptr)
}

func marshalXVisualInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapXVisualInfo(unsafe.Pointer(b))
}

// Native returns the underlying C source pointer.
func (x *XVisualInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&x.native)
}

type XWindowAttributes struct {
	native C.XWindowAttributes
}

// WrapXWindowAttributes wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapXWindowAttributes(ptr unsafe.Pointer) *XWindowAttributes {
	if ptr == nil {
		return nil
	}

	return (*XWindowAttributes)(ptr)
}

func marshalXWindowAttributes(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapXWindowAttributes(unsafe.Pointer(b))
}

// Native returns the underlying C source pointer.
func (x *XWindowAttributes) Native() unsafe.Pointer {
	return unsafe.Pointer(&x.native)
}
