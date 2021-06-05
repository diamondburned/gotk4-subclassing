// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"github.com/diamondburned/gotk4/internal/box"
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

// BusNameAppearedCallback: invoked when the name being watched is known to have
// to have an owner.
type BusNameAppearedCallback func(connection DBusConnection, name string, nameOwner string)

//export gotk4_BusNameAppearedCallback
func gotk4_BusNameAppearedCallback(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 *C.gchar, arg3 C.gpointer) {
	v := box.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(BusNameAppearedCallback)
	fn(connection, name, nameOwner, userData)
}

// BusNameVanishedCallback: invoked when the name being watched is known not to
// have to have an owner.
//
// This is also invoked when the BusConnection on which the watch was
// established has been closed. In that case, @connection will be nil.
type BusNameVanishedCallback func(connection DBusConnection, name string)

//export gotk4_BusNameVanishedCallback
func gotk4_BusNameVanishedCallback(arg0 *C.GDBusConnection, arg1 *C.gchar, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(BusNameVanishedCallback)
	fn(connection, name, userData)
}

// BusUnwatchName stops watching a name.
//
// Note that there may still be D-Bus traffic to process (relating to watching
// and unwatching the name) in the current thread-default Context after this
// function has returned. You should continue to iterate the Context until the
// Notify function passed to g_bus_watch_name() is called, in order to avoid
// memory leaks through callbacks queued on the Context after it’s stopped being
// iterated.
func BusUnwatchName(watcherID uint) {
	var arg1 C.guint

	arg1 = C.guint(watcherID)

	C.g_bus_unwatch_name(watcherID)
}
