// Code generated by girgen. DO NOT EDIT.

package gdkx11

// #cgo pkg-config: gtk4-x11 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/x11/gdkx.h>
import "C"

// X11FreeCompoundText frees the data returned from
// gdk_x11_display_string_to_compound_text().
func X11FreeCompoundText(ctext byte) {
	var arg1 *C.guchar

	arg1 = *C.guchar(ctext)

	C.gdk_x11_free_compound_text(arg1)
}

// X11FreeTextList frees the array of strings created by
// gdk_x11_display_text_property_to_text_list().
func X11FreeTextList(list string) {
	var arg1 **C.char

	arg1 = (**C.char)(C.CString(list))
	defer C.free(unsafe.Pointer(arg1))

	C.gdk_x11_free_text_list(arg1)
}
