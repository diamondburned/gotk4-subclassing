// Code generated by girgen. DO NOT EDIT.

package gdkpixbuf

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #cgo pkg-config: gdk-pixbuf-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk-pixbuf/gdk-pixbuf.h>
import "C"

// PixbufSaveFunc: save functions used by
// [method@GdkPixbuf.Pixbuf.save_to_callback].
//
// This function is called once for each block of bytes that is "written" by
// `gdk_pixbuf_save_to_callback()`.
//
// If successful it should return `TRUE`; if an error occurs it should set
// `error` and return `FALSE`, in which case `gdk_pixbuf_save_to_callback()`
// will fail with the same error.
type PixbufSaveFunc func(buf []byte) (err *error, ok bool)

//export gotk4_PixbufSaveFunc
func gotk4_PixbufSaveFunc(arg0 *C.gchar, arg1 C.gsize, arg2 **C.GError, arg3 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var buf []byte

	{
		var src []C.guint8
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(arg0), int(arg1))

		buf = make([]byte, arg1)
		for i := 0; i < uintptr(arg1); i++ {
			buf = (byte)(arg0)
		}
	}

	fn := v.(PixbufSaveFunc)
	err, ok := fn(buf)

	arg2 = (*C.GError)(gerror.New(unsafe.Pointer(err)))
	if ok {
		cret = C.gboolean(1)
	}

	return ok
}
