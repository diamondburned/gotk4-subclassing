// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

// ShowURI: convenience function for launching the default application to show
// the uri. Like gtk_show_uri_on_window(), but takes a screen as transient
// parent instead of a window.
//
// Note that this function is deprecated as it does not pass the necessary
// information for helpers to parent their dialog properly, when run from
// sandboxed applications for example.
//
// Deprecated: Use gtk_show_uri_on_window() instead.
//
// The function takes the following parameters:
//
//    - screen to show the uri on or NULL for the default screen.
//    - uri to show.
//    - timestamp to prevent focus stealing.
//
func ShowURI(screen *gdk.Screen, uri string, timestamp uint32) error {
	var _arg1 *C.GdkScreen // out
	var _arg2 *C.gchar     // out
	var _arg3 C.guint32    // out
	var _cerr *C.GError    // in

	if screen != nil {
		_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.guint32(timestamp)

	C.gtk_show_uri(_arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(screen)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(timestamp)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// ShowURIOnWindow: this is a convenience function for launching the default
// application to show the uri. The uri must be of a form understood by GIO
// (i.e. you need to install gvfs to get support for uri schemes such as http://
// or ftp://, as only local files are handled by GIO itself). Typical examples
// are
//
// - file:///home/gnome/pict.jpg
//
// - http://www.gnome.org
//
// - mailto:megnome.org
//
// Ideally the timestamp is taken from the event triggering the gtk_show_uri()
// call. If timestamp is not known you can take GDK_CURRENT_TIME.
//
// This is the recommended call to be used as it passes information necessary
// for sandbox helpers to parent their dialogs properly.
//
// The function takes the following parameters:
//
//    - parent window.
//    - uri to show.
//    - timestamp to prevent focus stealing.
//
func ShowURIOnWindow(parent *Window, uri string, timestamp uint32) error {
	var _arg1 *C.GtkWindow // out
	var _arg2 *C.char      // out
	var _arg3 C.guint32    // out
	var _cerr *C.GError    // in

	if parent != nil {
		_arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	}
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.guint32(timestamp)

	C.gtk_show_uri_on_window(_arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(timestamp)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}
