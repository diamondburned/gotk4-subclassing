// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk.h>
import "C"

// AssistantPageFunc: a function used by gtk_assistant_set_forward_page_func()
// to know which is the next page given a current one. It’s called both for
// computing the next page when the user presses the “forward” button and for
// handling the behavior of the “last” button.
type AssistantPageFunc func() (gint int)

//export gotk4_AssistantPageFunc
func gotk4_AssistantPageFunc(arg0 C.int, arg1 C.gpointer) C.int {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(AssistantPageFunc)
	gint := fn()

	cret = C.int(gint)
}
