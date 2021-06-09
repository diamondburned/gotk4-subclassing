// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// PangoLayoutGetClipRegion obtains a clip region which contains the areas where
// the given ranges of text would be drawn. @x_origin and @y_origin are the top
// left point to center the layout. @index_ranges should contain ranges of bytes
// in the layout’s text.
//
// Note that the regions returned correspond to logical extents of the text
// ranges, not ink extents. So the drawn layout may in fact touch areas out of
// the clip region. The clip region is mainly useful for highlightling parts of
// text, such as when text is selected.
func PangoLayoutGetClipRegion(layout pango.Layout, xOrigin int, yOrigin int, indexRanges *int, nRanges int) *cairo.Region {
	var arg1 *C.PangoLayout
	var arg2 C.int
	var arg3 C.int
	var arg4 *C.int
	var arg5 C.int

	arg1 = (*C.PangoLayout)(unsafe.Pointer(layout.Native()))
	arg2 = C.int(xOrigin)
	arg3 = C.int(yOrigin)
	arg4 = *C.int(indexRanges)
	arg5 = C.int(nRanges)

	var cret *C.cairo_region_t

	cret = C.gdk_pango_layout_get_clip_region(arg1, arg2, arg3, arg4, arg5)

	var region *cairo.Region

	region = cairo.WrapRegion(unsafe.Pointer(cret))
	runtime.SetFinalizer(region, func(v *cairo.Region) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return region
}
