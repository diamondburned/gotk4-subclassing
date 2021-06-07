// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// ClipboardImageReceivedFunc: a function to be called when the results of
// gtk_clipboard_request_image() are received, or when the request fails.
type ClipboardImageReceivedFunc func(clipboard Clipboard, pixbuf gdkpixbuf.Pixbuf)

//export gotk4_ClipboardImageReceivedFunc
func gotk4_ClipboardImageReceivedFunc(arg0 *C.GtkClipboard, arg1 *C.GdkPixbuf, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(ClipboardImageReceivedFunc)
	fn(clipboard, pixbuf, data)
}

// ClipboardReceivedFunc: a function to be called when the results of
// gtk_clipboard_request_contents() are received, or when the request fails.
type ClipboardReceivedFunc func(clipboard Clipboard, selectionData *SelectionData)

//export gotk4_ClipboardReceivedFunc
func gotk4_ClipboardReceivedFunc(arg0 *C.GtkClipboard, arg1 *C.GtkSelectionData, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(ClipboardReceivedFunc)
	fn(clipboard, selectionData, data)
}

// ClipboardRichTextReceivedFunc: a function to be called when the results of
// gtk_clipboard_request_rich_text() are received, or when the request fails.
type ClipboardRichTextReceivedFunc func(clipboard Clipboard, format gdk.Atom, text string, length uint)

//export gotk4_ClipboardRichTextReceivedFunc
func gotk4_ClipboardRichTextReceivedFunc(arg0 *C.GtkClipboard, arg1 C.GdkAtom, arg2 *C.guint8, arg3 C.gsize, arg4 C.gpointer) {
	v := box.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(ClipboardRichTextReceivedFunc)
	fn(clipboard, format, text, length, data)
}

// ClipboardTargetsReceivedFunc: a function to be called when the results of
// gtk_clipboard_request_targets() are received, or when the request fails.
type ClipboardTargetsReceivedFunc func(clipboard Clipboard, atoms []gdk.Atom)

//export gotk4_ClipboardTargetsReceivedFunc
func gotk4_ClipboardTargetsReceivedFunc(arg0 *C.GtkClipboard, arg1 *C.GdkAtom, arg2 C.gint, arg3 C.gpointer) {
	v := box.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(ClipboardTargetsReceivedFunc)
	fn(clipboard, atoms, nAtoms, data)
}

// ClipboardTextReceivedFunc: a function to be called when the results of
// gtk_clipboard_request_text() are received, or when the request fails.
type ClipboardTextReceivedFunc func(clipboard Clipboard, text string)

//export gotk4_ClipboardTextReceivedFunc
func gotk4_ClipboardTextReceivedFunc(arg0 *C.GtkClipboard, arg1 *C.gchar, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(ClipboardTextReceivedFunc)
	fn(clipboard, text, data)
}

// ClipboardURIReceivedFunc: a function to be called when the results of
// gtk_clipboard_request_uris() are received, or when the request fails.
type ClipboardURIReceivedFunc func(clipboard Clipboard, uris []string)

//export gotk4_ClipboardURIReceivedFunc
func gotk4_ClipboardURIReceivedFunc(arg0 *C.GtkClipboard, arg1 **C.gchar, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(ClipboardURIReceivedFunc)
	fn(clipboard, uris, data)
}