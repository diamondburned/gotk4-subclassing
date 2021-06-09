// Code generated by girgen. DO NOT EDIT.

package gtk

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk.h>
import "C"

// GetDebugFlags returns the GTK debug flags that are currently active.
//
// This function is intended for GTK modules that want to adjust their debug
// output based on GTK debug flags.
func GetDebugFlags() DebugFlags {
	var cret C.GtkDebugFlags

	cret = C.gtk_get_debug_flags()

	var debugFlags DebugFlags

	debugFlags = DebugFlags(cret)

	return debugFlags
}

// SetDebugFlags sets the GTK debug flags.
func SetDebugFlags(flags DebugFlags) {
	var arg1 C.GtkDebugFlags

	arg1 = (C.GtkDebugFlags)(flags)

	C.gtk_set_debug_flags(arg1)
}
