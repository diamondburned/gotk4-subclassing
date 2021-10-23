// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/pango"
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
		{T: externglib.Type(C.gtk_text_attributes_get_type()), F: marshalTextAttributes},
	})
}

// TextAppearance: instance of this type is always passed by reference.
type TextAppearance struct {
	*textAppearance
}

// textAppearance is the struct that's finalized.
type textAppearance struct {
	native *C.GtkTextAppearance
}

// BgColor: background Color.
func (t *TextAppearance) BgColor() *gdk.Color {
	var v *gdk.Color // out
	v = (*gdk.Color)(gextras.NewStructNative(unsafe.Pointer((&t.native.bg_color))))
	return v
}

// FgColor: foreground Color.
func (t *TextAppearance) FgColor() *gdk.Color {
	var v *gdk.Color // out
	v = (*gdk.Color)(gextras.NewStructNative(unsafe.Pointer((&t.native.fg_color))))
	return v
}

// Rise: super/subscript rise, can be negative.
func (t *TextAppearance) Rise() int {
	var v int // out
	v = int(t.native.rise)
	return v
}

// TextAttributes: using TextAttributes directly should rarely be necessary.
// It’s primarily useful with gtk_text_iter_get_attributes(). As with most GTK+
// structs, the fields in this struct should only be read, never modified
// directly.
//
// An instance of this type is always passed by reference.
type TextAttributes struct {
	*textAttributes
}

// textAttributes is the struct that's finalized.
type textAttributes struct {
	native *C.GtkTextAttributes
}

func marshalTextAttributes(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &TextAttributes{&textAttributes{(*C.GtkTextAttributes)(b)}}, nil
}

// NewTextAttributes constructs a struct TextAttributes.
func NewTextAttributes() *TextAttributes {
	var _cret *C.GtkTextAttributes // in

	_cret = C.gtk_text_attributes_new()

	var _textAttributes *TextAttributes // out

	_textAttributes = (*TextAttributes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_textAttributes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_text_attributes_unref((*C.GtkTextAttributes)(intern.C))
		},
	)

	return _textAttributes
}

// Appearance for text.
func (t *TextAttributes) Appearance() *TextAppearance {
	var v *TextAppearance // out
	v = (*TextAppearance)(gextras.NewStructNative(unsafe.Pointer((&t.native.appearance))))
	return v
}

// Justification for text.
func (t *TextAttributes) Justification() Justification {
	var v Justification // out
	v = Justification(t.native.justification)
	return v
}

// Direction for text.
func (t *TextAttributes) Direction() TextDirection {
	var v TextDirection // out
	v = TextDirection(t.native.direction)
	return v
}

// Font for text.
func (t *TextAttributes) Font() *pango.FontDescription {
	var v *pango.FontDescription // out
	v = (*pango.FontDescription)(gextras.NewStructNative(unsafe.Pointer(t.native.font)))
	return v
}

// FontScale: font scale factor.
func (t *TextAttributes) FontScale() float64 {
	var v float64 // out
	v = float64(t.native.font_scale)
	return v
}

// LeftMargin: width of the left margin in pixels.
func (t *TextAttributes) LeftMargin() int {
	var v int // out
	v = int(t.native.left_margin)
	return v
}

// RightMargin: width of the right margin in pixels.
func (t *TextAttributes) RightMargin() int {
	var v int // out
	v = int(t.native.right_margin)
	return v
}

// Indent: amount to indent the paragraph, in pixels.
func (t *TextAttributes) Indent() int {
	var v int // out
	v = int(t.native.indent)
	return v
}

// PixelsAboveLines pixels of blank space above paragraphs.
func (t *TextAttributes) PixelsAboveLines() int {
	var v int // out
	v = int(t.native.pixels_above_lines)
	return v
}

// PixelsBelowLines pixels of blank space below paragraphs.
func (t *TextAttributes) PixelsBelowLines() int {
	var v int // out
	v = int(t.native.pixels_below_lines)
	return v
}

// PixelsInsideWrap pixels of blank space between wrapped lines in a paragraph.
func (t *TextAttributes) PixelsInsideWrap() int {
	var v int // out
	v = int(t.native.pixels_inside_wrap)
	return v
}

// Tabs: custom TabArray for this text.
func (t *TextAttributes) Tabs() *pango.TabArray {
	var v *pango.TabArray // out
	v = (*pango.TabArray)(gextras.NewStructNative(unsafe.Pointer(t.native.tabs)))
	return v
}

// WrapMode for text.
func (t *TextAttributes) WrapMode() WrapMode {
	var v WrapMode // out
	v = WrapMode(t.native.wrap_mode)
	return v
}

// Language for text.
func (t *TextAttributes) Language() *pango.Language {
	var v *pango.Language // out
	v = (*pango.Language)(gextras.NewStructNative(unsafe.Pointer(t.native.language)))
	return v
}

// LetterSpacing: extra space to insert between graphemes, in Pango units.
func (t *TextAttributes) LetterSpacing() int {
	var v int // out
	v = int(t.native.letter_spacing)
	return v
}

// Copy copies src and returns a new TextAttributes.
func (src *TextAttributes) Copy() *TextAttributes {
	var _arg0 *C.GtkTextAttributes // out
	var _cret *C.GtkTextAttributes // in

	_arg0 = (*C.GtkTextAttributes)(gextras.StructNative(unsafe.Pointer(src)))

	_cret = C.gtk_text_attributes_copy(_arg0)
	runtime.KeepAlive(src)

	var _textAttributes *TextAttributes // out

	_textAttributes = (*TextAttributes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_textAttributes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gtk_text_attributes_unref((*C.GtkTextAttributes)(intern.C))
		},
	)

	return _textAttributes
}

// CopyValues copies the values from src to dest so that dest has the same
// values as src. Frees existing values in dest.
func (src *TextAttributes) CopyValues(dest *TextAttributes) {
	var _arg0 *C.GtkTextAttributes // out
	var _arg1 *C.GtkTextAttributes // out

	_arg0 = (*C.GtkTextAttributes)(gextras.StructNative(unsafe.Pointer(src)))
	_arg1 = (*C.GtkTextAttributes)(gextras.StructNative(unsafe.Pointer(dest)))

	C.gtk_text_attributes_copy_values(_arg0, _arg1)
	runtime.KeepAlive(src)
	runtime.KeepAlive(dest)
}
