// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_css_section_get_type()), F: marshalCSSSection},
	})
}

// CSSSection defines a part of a CSS document. Because sections are nested into
// one another, you can use gtk_css_section_get_parent() to get the containing
// region.
type CSSSection struct {
	native C.GtkCssSection
}

// WrapCSSSection wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapCSSSection(ptr unsafe.Pointer) *CSSSection {
	if ptr == nil {
		return nil
	}

	return (*CSSSection)(ptr)
}

func marshalCSSSection(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapCSSSection(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *CSSSection) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}

// EndLine returns the line in the CSS document where this section end. The line
// number is 0-indexed, so the first line of the document will return 0. This
// value may change in future invocations of this function if @section is not
// yet parsed completely. This will for example happen in the
// GtkCssProvider::parsing-error signal. The end position and line may be
// identical to the start position and line for sections which failed to parse
// anything successfully.
func (s *CSSSection) EndLine() uint {
	var _arg0 *C.GtkCssSection // out

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	var _cret C.guint // in

	_cret = C.gtk_css_section_get_end_line(_arg0)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// EndPosition returns the offset in bytes from the start of the current line
// returned via gtk_css_section_get_end_line(). This value may change in future
// invocations of this function if @section is not yet parsed completely. This
// will for example happen in the GtkCssProvider::parsing-error signal. The end
// position and line may be identical to the start position and line for
// sections which failed to parse anything successfully.
func (s *CSSSection) EndPosition() uint {
	var _arg0 *C.GtkCssSection // out

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	var _cret C.guint // in

	_cret = C.gtk_css_section_get_end_position(_arg0)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// StartLine returns the line in the CSS document where this section starts. The
// line number is 0-indexed, so the first line of the document will return 0.
func (s *CSSSection) StartLine() uint {
	var _arg0 *C.GtkCssSection // out

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	var _cret C.guint // in

	_cret = C.gtk_css_section_get_start_line(_arg0)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// StartPosition returns the offset in bytes from the start of the current line
// returned via gtk_css_section_get_start_line().
func (s *CSSSection) StartPosition() uint {
	var _arg0 *C.GtkCssSection // out

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	var _cret C.guint // in

	_cret = C.gtk_css_section_get_start_position(_arg0)

	var _guint uint // out

	_guint = (uint)(_cret)

	return _guint
}

// Unref decrements the reference count on @section, freeing the structure if
// the reference count reaches 0.
func (s *CSSSection) Unref() {
	var _arg0 *C.GtkCssSection // out

	_arg0 = (*C.GtkCssSection)(unsafe.Pointer(s.Native()))

	C.gtk_css_section_unref(_arg0)
}
