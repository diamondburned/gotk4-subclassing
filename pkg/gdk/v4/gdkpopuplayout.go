// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_popup_layout_get_type()), F: marshalPopupLayout},
	})
}

// PopupLayout: the `GdkPopupLayout` struct contains information that is
// necessary position a [interface@Gdk.Popup] relative to its parent.
//
// The positioning requires a negotiation with the windowing system, since it
// depends on external constraints, such as the position of the parent surface,
// and the screen dimensions.
//
// The basic ingredients are a rectangle on the parent surface, and the anchor
// on both that rectangle and the popup. The anchors specify a side or corner to
// place next to each other.
//
// !Popup anchors (popup-anchors.png)
//
// For cases where placing the anchors next to each other would make the popup
// extend offscreen, the layout includes some hints for how to resolve this
// problem. The hints may suggest to flip the anchor position to the other side,
// or to 'slide' the popup along a side, or to resize it.
//
// !Flipping popups (popup-flip.png)
//
// !Sliding popups (popup-slide.png)
//
// These hints may be combined.
//
// Ultimatively, it is up to the windowing system to determine the position and
// size of the popup. You can learn about the result by calling
// [method@Gdk.Popup.get_position_x], [method@Gdk.Popup.get_position_y],
// [method@Gdk.Popup.get_rect_anchor] and [method@Gdk.Popup.get_surface_anchor]
// after the popup has been presented. This can be used to adjust the rendering.
// For example, [class@Gtk.Popover] changes its arrow position accordingly. But
// you have to be careful avoid changing the size of the popover, or it has to
// be presented again.
type PopupLayout struct {
	native C.GdkPopupLayout
}

// WrapPopupLayout wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapPopupLayout(ptr unsafe.Pointer) *PopupLayout {
	if ptr == nil {
		return nil
	}

	return (*PopupLayout)(ptr)
}

func marshalPopupLayout(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapPopupLayout(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (p *PopupLayout) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}

// Equal: check whether @layout and @other has identical layout properties.
func (l *PopupLayout) Equal(other *PopupLayout) bool {
	var _arg0 *C.GdkPopupLayout // out
	var _arg1 *C.GdkPopupLayout // out

	_arg0 = (*C.GdkPopupLayout)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.GdkPopupLayout)(unsafe.Pointer(other.Native()))

	var _cret C.gboolean // in

	_cret = C.gdk_popup_layout_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Offset retrieves the offset for the anchor rectangle.
func (l *PopupLayout) Offset() (dx int, dy int) {
	var _arg0 *C.GdkPopupLayout // out

	_arg0 = (*C.GdkPopupLayout)(unsafe.Pointer(l.Native()))

	var _arg1 C.int // in
	var _arg2 C.int // in

	C.gdk_popup_layout_get_offset(_arg0, &_arg1, &_arg2)

	var _dx int // out
	var _dy int // out

	_dx = (int)(_arg1)
	_dy = (int)(_arg2)

	return _dx, _dy
}

// ShadowWidth obtains the shadow widths of this layout.
func (l *PopupLayout) ShadowWidth() (left int, right int, top int, bottom int) {
	var _arg0 *C.GdkPopupLayout // out

	_arg0 = (*C.GdkPopupLayout)(unsafe.Pointer(l.Native()))

	var _arg1 C.int // in
	var _arg2 C.int // in
	var _arg3 C.int // in
	var _arg4 C.int // in

	C.gdk_popup_layout_get_shadow_width(_arg0, &_arg1, &_arg2, &_arg3, &_arg4)

	var _left int   // out
	var _right int  // out
	var _top int    // out
	var _bottom int // out

	_left = (int)(_arg1)
	_right = (int)(_arg2)
	_top = (int)(_arg3)
	_bottom = (int)(_arg4)

	return _left, _right, _top, _bottom
}

// SetAnchorHints: set new anchor hints.
//
// The set @anchor_hints determines how @surface will be moved if the anchor
// points cause it to move off-screen. For example, GDK_ANCHOR_FLIP_X will
// replace GDK_GRAVITY_NORTH_WEST with GDK_GRAVITY_NORTH_EAST and vice versa if
// @surface extends beyond the left or right edges of the monitor.
func (l *PopupLayout) SetAnchorHints(anchorHints AnchorHints) {
	var _arg0 *C.GdkPopupLayout // out
	var _arg1 C.GdkAnchorHints  // out

	_arg0 = (*C.GdkPopupLayout)(unsafe.Pointer(l.Native()))
	_arg1 = (C.GdkAnchorHints)(anchorHints)

	C.gdk_popup_layout_set_anchor_hints(_arg0, _arg1)
}

// SetAnchorRect: set the anchor rectangle.
func (l *PopupLayout) SetAnchorRect(anchorRect *Rectangle) {
	var _arg0 *C.GdkPopupLayout // out
	var _arg1 *C.GdkRectangle   // out

	_arg0 = (*C.GdkPopupLayout)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.GdkRectangle)(unsafe.Pointer(anchorRect.Native()))

	C.gdk_popup_layout_set_anchor_rect(_arg0, _arg1)
}

// SetOffset: offset the position of the anchor rectangle with the given delta.
func (l *PopupLayout) SetOffset(dx int, dy int) {
	var _arg0 *C.GdkPopupLayout // out
	var _arg1 C.int             // out
	var _arg2 C.int             // out

	_arg0 = (*C.GdkPopupLayout)(unsafe.Pointer(l.Native()))
	_arg1 = C.int(dx)
	_arg2 = C.int(dy)

	C.gdk_popup_layout_set_offset(_arg0, _arg1, _arg2)
}

// SetRectAnchor: set the anchor on the anchor rectangle.
func (l *PopupLayout) SetRectAnchor(anchor Gravity) {
	var _arg0 *C.GdkPopupLayout // out
	var _arg1 C.GdkGravity      // out

	_arg0 = (*C.GdkPopupLayout)(unsafe.Pointer(l.Native()))
	_arg1 = (C.GdkGravity)(anchor)

	C.gdk_popup_layout_set_rect_anchor(_arg0, _arg1)
}

// SetShadowWidth sets the shadow width of the popup.
//
// The shadow width corresponds to the part of the computed surface size that
// would consist of the shadow margin surrounding the window, would there be
// any.
func (l *PopupLayout) SetShadowWidth(left int, right int, top int, bottom int) {
	var _arg0 *C.GdkPopupLayout // out
	var _arg1 C.int             // out
	var _arg2 C.int             // out
	var _arg3 C.int             // out
	var _arg4 C.int             // out

	_arg0 = (*C.GdkPopupLayout)(unsafe.Pointer(l.Native()))
	_arg1 = C.int(left)
	_arg2 = C.int(right)
	_arg3 = C.int(top)
	_arg4 = C.int(bottom)

	C.gdk_popup_layout_set_shadow_width(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// SetSurfaceAnchor: set the anchor on the popup surface.
func (l *PopupLayout) SetSurfaceAnchor(anchor Gravity) {
	var _arg0 *C.GdkPopupLayout // out
	var _arg1 C.GdkGravity      // out

	_arg0 = (*C.GdkPopupLayout)(unsafe.Pointer(l.Native()))
	_arg1 = (C.GdkGravity)(anchor)

	C.gdk_popup_layout_set_surface_anchor(_arg0, _arg1)
}

// Unref decreases the reference count of @value.
func (l *PopupLayout) Unref() {
	var _arg0 *C.GdkPopupLayout // out

	_arg0 = (*C.GdkPopupLayout)(unsafe.Pointer(l.Native()))

	C.gdk_popup_layout_unref(_arg0)
}
