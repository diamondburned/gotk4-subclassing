// Code generated by girgen. DO NOT EDIT.

package gdkwayland

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4-wayland gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/wayland/gdkwayland.h>
import "C"

// WaylandToplevelExported: callback that gets called when the handle for a
// surface has been obtained from the Wayland compositor.
//
// This callback is used in [method@GdkWayland.WaylandToplevel.export_handle].
//
// The @handle can be passed to other processes, for the purpose of marking
// surfaces as transient for out-of-process surfaces.
type WaylandToplevelExported func(toplevel *WaylandToplevelClass, handle string, userData interface{})

//export gotk4_WaylandToplevelExported
func gotk4_WaylandToplevelExported(arg0 *C.GdkToplevel, arg1 *C.char, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var toplevel *WaylandToplevelClass // out
	var handle string                  // out
	var userData interface{}           // out

	toplevel = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*WaylandToplevelClass)
	handle = C.GoString(arg1)
	userData = box.Get(uintptr(arg2))

	fn := v.(WaylandToplevelExported)
	fn(toplevel, handle, userData)
}
