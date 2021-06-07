// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/internal/box"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// TextBufferDeserializeFunc: a function that is called to deserialize rich text
// that has been serialized with gtk_text_buffer_serialize(), and insert it at
// @iter.
type TextBufferDeserializeFunc func(registerBuffer TextBuffer, contentBuffer TextBuffer, iter *TextIter, data []byte, createTags bool) error

//export gotk4_TextBufferDeserializeFunc
func gotk4_TextBufferDeserializeFunc(arg0 *C.GtkTextBuffer, arg1 *C.GtkTextBuffer, arg2 *C.GtkTextIter, arg3 *C.guint8, arg4 C.gsize, arg5 C.gboolean, arg6 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg6))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(TextBufferDeserializeFunc)
	ret := fn(registerBuffer, contentBuffer, iter, data, length, createTags, userData)

	if ret {
		cret = C.gboolean(1)
	}

	return cret
}

// TextBufferSerializeFunc: a function that is called to serialize the content
// of a text buffer. It must return the serialized form of the content.
type TextBufferSerializeFunc func(registerBuffer TextBuffer, contentBuffer TextBuffer, start *TextIter, end *TextIter, length uint) byte

//export gotk4_TextBufferSerializeFunc
func gotk4_TextBufferSerializeFunc(arg0 *C.GtkTextBuffer, arg1 *C.GtkTextBuffer, arg2 *C.GtkTextIter, arg3 *C.GtkTextIter, arg4 *C.gsize, arg5 C.gpointer) *C.guint8 {
	v := box.Get(uintptr(arg5))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(TextBufferSerializeFunc)
	ret := fn(registerBuffer, contentBuffer, start, end, length, userData)

	cret = *C.guint8(ret)

	return cret
}
