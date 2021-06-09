// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"

	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
import "C"

// NewPollableSource: utility method for InputStream and OutputStream
// implementations. Creates a new #GSource that expects a callback of type
// SourceFunc. The new source does not actually do anything on its own; use
// g_source_add_child_source() to add other sources to it to cause it to
// trigger.
func NewPollableSource(pollableStream gextras.Objector) *glib.Source {
	var arg1 *C.GObject

	arg1 = (*C.GObject)(unsafe.Pointer(pollableStream.Native()))

	var cret *C.GSource

	cret = C.g_pollable_source_new(arg1)

	var source *glib.Source

	source = glib.WrapSource(unsafe.Pointer(cret))
	runtime.SetFinalizer(source, func(v *glib.Source) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return source
}

// PollableSourceNewFull: utility method for InputStream and OutputStream
// implementations. Creates a new #GSource, as with g_pollable_source_new(), but
// also attaching @child_source (with a dummy callback), and @cancellable, if
// they are non-nil.
func PollableSourceNewFull(pollableStream gextras.Objector, childSource *glib.Source, cancellable Cancellable) *glib.Source {
	var arg1 C.gpointer
	var arg2 *C.GSource
	var arg3 *C.GCancellable

	arg1 = (*C.GObject)(unsafe.Pointer(pollableStream.Native()))
	arg2 = (*C.GSource)(unsafe.Pointer(childSource.Native()))
	arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cret *C.GSource

	cret = C.g_pollable_source_new_full(arg1, arg2, arg3)

	var source *glib.Source

	source = glib.WrapSource(unsafe.Pointer(cret))
	runtime.SetFinalizer(source, func(v *glib.Source) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return source
}
