// Code generated by girgen. DO NOT EDIT.

package gio

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
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

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_loadable_icon_get_type()), F: marshalLoadableIcon},
	})
}

// LoadableIconOverrider contains methods that are overridable. This
// interface is a subset of the interface LoadableIcon.
type LoadableIconOverrider interface {
	// Load loads a loadable icon. For the asynchronous version of this
	// function, see g_loadable_icon_load_async().
	Load(i LoadableIcon, size int, cancellable Cancellable) (typ string, err error)
	// LoadAsync loads an icon asynchronously. To finish this function, see
	// g_loadable_icon_load_finish(). For the synchronous, blocking version of
	// this function, see g_loadable_icon_load().
	LoadAsync(i LoadableIcon)
	// LoadFinish finishes an asynchronous icon load started in
	// g_loadable_icon_load_async().
	LoadFinish(i LoadableIcon, res AsyncResult) (typ string, err error)
}

// LoadableIcon extends the #GIcon interface and adds the ability to load icons
// from streams.
type LoadableIcon interface {
	Icon
	LoadableIconOverrider
}

// loadableIcon implements the LoadableIcon interface.
type loadableIcon struct {
	Icon
}

var _ LoadableIcon = (*loadableIcon)(nil)

// WrapLoadableIcon wraps a GObject to a type that implements interface
// LoadableIcon. It is primarily used internally.
func WrapLoadableIcon(obj *externglib.Object) LoadableIcon {
	return LoadableIcon{
		Icon: WrapIcon(obj),
	}
}

func marshalLoadableIcon(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapLoadableIcon(obj), nil
}

// Load loads a loadable icon. For the asynchronous version of this
// function, see g_loadable_icon_load_async().
func (i loadableIcon) Load(i LoadableIcon, size int, cancellable Cancellable) (typ string, err error) {
	var arg0 *C.GLoadableIcon
	var arg1 C.int
	var arg3 *C.GCancellable

	arg0 = (*C.GLoadableIcon)(unsafe.Pointer(i.Native()))
	arg1 = C.int(size)
	arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var arg2 *C.char
	var typ string
	var errout *C.GError
	var err error

	C.g_loadable_icon_load(arg0, arg1, &arg2, arg3, &errout)

	typ = C.GoString(&arg2)
	defer C.free(unsafe.Pointer(&arg2))
	err = gerror.Take(unsafe.Pointer(errout))

	return typ, err
}

// LoadAsync loads an icon asynchronously. To finish this function, see
// g_loadable_icon_load_finish(). For the synchronous, blocking version of
// this function, see g_loadable_icon_load().
func (i loadableIcon) LoadAsync(i LoadableIcon) {
	var arg0 *C.GLoadableIcon

	arg0 = (*C.GLoadableIcon)(unsafe.Pointer(i.Native()))

	C.g_loadable_icon_load_async(arg0, arg1, arg2, arg3, arg4)
}

// LoadFinish finishes an asynchronous icon load started in
// g_loadable_icon_load_async().
func (i loadableIcon) LoadFinish(i LoadableIcon, res AsyncResult) (typ string, err error) {
	var arg0 *C.GLoadableIcon
	var arg1 *C.GAsyncResult

	arg0 = (*C.GLoadableIcon)(unsafe.Pointer(i.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(res.Native()))

	var arg2 *C.char
	var typ string
	var errout *C.GError
	var err error

	C.g_loadable_icon_load_finish(arg0, arg1, &arg2, &errout)

	typ = C.GoString(&arg2)
	defer C.free(unsafe.Pointer(&arg2))
	err = gerror.Take(unsafe.Pointer(errout))

	return typ, err
}