// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk.h>
import "C"

// CSSLocation represents a location in a file or other source of data parsed by
// the CSS engine.
//
// The bytes and line_bytes offsets are meant to be used to programmatically
// match data. The lines and line_chars offsets can be used for printing the
// location in a file.
//
// Note that the lines parameter starts from 0 and is increased whenever a CSS
// line break is encountered. (CSS defines the C character sequences "\r\n",
// "\r", "\n" and "\f" as newlines.) If your document uses different rules for
// line breaking, you might want run into problems here.
//
// An instance of this type is always passed by reference.
type CSSLocation struct {
	*cssLocation
}

// cssLocation is the struct that's finalized.
type cssLocation struct {
	native *C.GtkCssLocation
}

// NewCSSLocation creates a new CSSLocation instance from the given
// fields.
func NewCSSLocation(bytes, chars, lines, lineBytes, lineChars uint) CSSLocation {
	var f0 C.gsize // out
	f0 = C.gsize(bytes)
	var f1 C.gsize // out
	f1 = C.gsize(chars)
	var f2 C.gsize // out
	f2 = C.gsize(lines)
	var f3 C.gsize // out
	f3 = C.gsize(lineBytes)
	var f4 C.gsize // out
	f4 = C.gsize(lineChars)

	v := C.GtkCssLocation{
		bytes:      f0,
		chars:      f1,
		lines:      f2,
		line_bytes: f3,
		line_chars: f4,
	}

	return *(*CSSLocation)(gextras.NewStructNative(unsafe.Pointer(&v)))
}

// Bytes: number of bytes parsed since the beginning.
func (c *CSSLocation) Bytes() uint {
	var v uint // out
	v = uint(c.native.bytes)
	return v
}

// Chars: number of characters parsed since the beginning.
func (c *CSSLocation) Chars() uint {
	var v uint // out
	v = uint(c.native.chars)
	return v
}

// Lines: number of full lines that have been parsed If you want to display this
// as a line number, you need to add 1 to this.
func (c *CSSLocation) Lines() uint {
	var v uint // out
	v = uint(c.native.lines)
	return v
}

// LineBytes: number of bytes parsed since the last line break.
func (c *CSSLocation) LineBytes() uint {
	var v uint // out
	v = uint(c.native.line_bytes)
	return v
}

// LineChars: number of characters parsed since the last line break.
func (c *CSSLocation) LineChars() uint {
	var v uint // out
	v = uint(c.native.line_chars)
	return v
}
