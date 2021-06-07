// Code generated by girgen. DO NOT EDIT.

package gdk

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// ListVisuals lists the available visuals for the default screen. (See
// gdk_screen_list_visuals()) A visual describes a hardware image data format.
// For example, a visual might support 24-bit color, or 8-bit color, and might
// expect pixels to be in a certain format.
//
// Call g_list_free() on the return value when you’re finished with it.
func ListVisuals() {
	C.gdk_list_visuals()
}

// QueryDepths: this function returns the available bit depths for the default
// screen. It’s equivalent to listing the visuals (gdk_list_visuals()) and then
// looking at the depth field in each visual, removing duplicates.
//
// The array returned by this function should not be freed.
func QueryDepths() {
	C.gdk_query_depths(&arg1, &arg2)

	return depths, count
}

// QueryVisualTypes: this function returns the available visual types for the
// default screen. It’s equivalent to listing the visuals (gdk_list_visuals())
// and then looking at the type field in each visual, removing duplicates.
//
// The array returned by this function should not be freed.
func QueryVisualTypes() {
	C.gdk_query_visual_types(&arg1, &arg2)

	return visualTypes, count
}