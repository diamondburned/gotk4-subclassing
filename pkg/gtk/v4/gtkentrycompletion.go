// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/internal/box"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk.h>
import "C"

// EntryCompletionMatchFunc: a function which decides whether the row indicated
// by @iter matches a given @key, and should be displayed as a possible
// completion for @key. Note that @key is normalized and case-folded (see
// g_utf8_normalize() and g_utf8_casefold()). If this is not appropriate, match
// functions have access to the unmodified key via `gtk_editable_get_text
// (GTK_EDITABLE (gtk_entry_completion_get_entry ()))`.
type EntryCompletionMatchFunc func(completion EntryCompletion, key string, iter *TreeIter) bool

//export gotk4_EntryCompletionMatchFunc
func gotk4_EntryCompletionMatchFunc(arg0 *C.GtkEntryCompletion, arg1 *C.char, arg2 *C.GtkTreeIter, arg3 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(EntryCompletionMatchFunc)
	ret := fn(completion, key, iter, userData)

	if ret {
		cret = C.gboolean(1)
	}

	return cret
}
