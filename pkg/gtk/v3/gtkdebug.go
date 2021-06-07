// Code generated by girgen. DO NOT EDIT.

package gtk

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// GetDebugFlags returns the GTK+ debug flags.
//
// This function is intended for GTK+ modules that want to adjust their debug
// output based on GTK+ debug flags.
func GetDebugFlags() {
	C.gtk_get_debug_flags()
}

// SetDebugFlags sets the GTK+ debug flags.
func SetDebugFlags(flags uint) {
	var arg1 C.guint

	arg1 = C.guint(flags)

	C.gtk_set_debug_flags(arg1)
}