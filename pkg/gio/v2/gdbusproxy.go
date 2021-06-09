// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
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
		{T: externglib.Type(C.g_dbus_proxy_get_type()), F: marshalDBusProXY},
	})
}

// DBusProXY is a base class used for proxies to access a D-Bus interface on a
// remote object. A BusProxy can be constructed for both well-known and unique
// names.
//
// By default, BusProxy will cache all properties (and listen to changes) of the
// remote object, and proxy all signals that get emitted. This behaviour can be
// changed by passing suitable BusProxyFlags when the proxy is created. If the
// proxy is for a well-known name, the property cache is flushed when the name
// owner vanishes and reloaded when a name owner appears.
//
// The unique name owner of the proxy's name is tracked and can be read from
// BusProxy:g-name-owner. Connect to the #GObject::notify signal to get notified
// of changes. Additionally, only signals and property changes emitted from the
// current name owner are considered and calls are always sent to the current
// name owner. This avoids a number of race conditions when the name is lost by
// one owner and claimed by another. However, if no name owner currently exists,
// then calls will be sent to the well-known name which may result in the
// message bus launching an owner (unless G_DBUS_PROXY_FLAGS_DO_NOT_AUTO_START
// is set).
//
// The generic BusProxy::g-properties-changed and BusProxy::g-signal signals are
// not very convenient to work with. Therefore, the recommended way of working
// with proxies is to subclass BusProxy, and have more natural properties and
// signals in your derived class. This [example][gdbus-example-gdbus-codegen]
// shows how this can easily be done using the [gdbus-codegen][gdbus-codegen]
// tool.
//
// A BusProxy instance can be used from multiple threads but note that all
// signals (e.g. BusProxy::g-signal, BusProxy::g-properties-changed and
// #GObject::notify) are emitted in the [thread-default main
// context][g-main-context-push-thread-default] of the thread where the instance
// was constructed.
//
// An example using a proxy for a well-known name can be found in
// gdbus-example-watch-proxy.c
// (https://git.gnome.org/browse/glib/tree/gio/tests/gdbus-example-watch-proxy.c)
type DBusProXY interface {
	gextras.Objector
	AsyncInitable
	DBusInterface
	Initable

	// Call: asynchronously invokes the @method_name method on @proxy.
	//
	// If @method_name contains any dots, then @name is split into interface and
	// method name parts. This allows using @proxy for invoking methods on other
	// interfaces.
	//
	// If the BusConnection associated with @proxy is closed then the operation
	// will fail with G_IO_ERROR_CLOSED. If @cancellable is canceled, the
	// operation will fail with G_IO_ERROR_CANCELLED. If @parameters contains a
	// value not compatible with the D-Bus protocol, the operation fails with
	// G_IO_ERROR_INVALID_ARGUMENT.
	//
	// If the @parameters #GVariant is floating, it is consumed. This allows
	// convenient 'inline' use of g_variant_new(), e.g.:
	//
	//    g_dbus_proxy_call (proxy,
	//                       "TwoStrings",
	//                       g_variant_new ("(ss)",
	//                                      "Thing One",
	//                                      "Thing Two"),
	//                       G_DBUS_CALL_FLAGS_NONE,
	//                       -1,
	//                       NULL,
	//                       (GAsyncReadyCallback) two_strings_done,
	//                       &data);
	//
	// If @proxy has an expected interface (see BusProxy:g-interface-info) and
	// @method_name is referenced by it, then the return value is checked
	// against the return type.
	//
	// This is an asynchronous method. When the operation is finished, @callback
	// will be invoked in the [thread-default main
	// context][g-main-context-push-thread-default] of the thread you are
	// calling this method from. You can then call g_dbus_proxy_call_finish() to
	// get the result of the operation. See g_dbus_proxy_call_sync() for the
	// synchronous version of this method.
	//
	// If @callback is nil then the D-Bus method call message will be sent with
	// the G_DBUS_MESSAGE_FLAGS_NO_REPLY_EXPECTED flag set.
	Call()
	// CallFinish finishes an operation started with g_dbus_proxy_call().
	CallFinish(res AsyncResult) (variant *glib.Variant, goerr error)
	// CallSync: synchronously invokes the @method_name method on @proxy.
	//
	// If @method_name contains any dots, then @name is split into interface and
	// method name parts. This allows using @proxy for invoking methods on other
	// interfaces.
	//
	// If the BusConnection associated with @proxy is disconnected then the
	// operation will fail with G_IO_ERROR_CLOSED. If @cancellable is canceled,
	// the operation will fail with G_IO_ERROR_CANCELLED. If @parameters
	// contains a value not compatible with the D-Bus protocol, the operation
	// fails with G_IO_ERROR_INVALID_ARGUMENT.
	//
	// If the @parameters #GVariant is floating, it is consumed. This allows
	// convenient 'inline' use of g_variant_new(), e.g.:
	//
	//    g_dbus_proxy_call_sync (proxy,
	//                            "TwoStrings",
	//                            g_variant_new ("(ss)",
	//                                           "Thing One",
	//                                           "Thing Two"),
	//                            G_DBUS_CALL_FLAGS_NONE,
	//                            -1,
	//                            NULL,
	//                            &error);
	//
	// The calling thread is blocked until a reply is received. See
	// g_dbus_proxy_call() for the asynchronous version of this method.
	//
	// If @proxy has an expected interface (see BusProxy:g-interface-info) and
	// @method_name is referenced by it, then the return value is checked
	// against the return type.
	CallSync(methodName string, parameters *glib.Variant, flags DBusCallFlags, timeoutMsec int, cancellable Cancellable) (variant *glib.Variant, goerr error)
	// CallWithUnixFdList: like g_dbus_proxy_call() but also takes a FDList
	// object.
	//
	// This method is only available on UNIX.
	CallWithUnixFdList()
	// CallWithUnixFdListFinish finishes an operation started with
	// g_dbus_proxy_call_with_unix_fd_list().
	CallWithUnixFdListFinish(res AsyncResult) (outFdList UnixFDList, variant *glib.Variant, goerr error)
	// CallWithUnixFdListSync: like g_dbus_proxy_call_sync() but also takes and
	// returns FDList objects.
	//
	// This method is only available on UNIX.
	CallWithUnixFdListSync(methodName string, parameters *glib.Variant, flags DBusCallFlags, timeoutMsec int, fdList UnixFDList, cancellable Cancellable) (outFdList UnixFDList, variant *glib.Variant, goerr error)
	// CachedProperty looks up the value for a property from the cache. This
	// call does no blocking IO.
	//
	// If @proxy has an expected interface (see BusProxy:g-interface-info) and
	// @property_name is referenced by it, then @value is checked against the
	// type of the property.
	CachedProperty(propertyName string) *glib.Variant
	// CachedPropertyNames gets the names of all cached properties on @proxy.
	CachedPropertyNames() []string
	// Connection gets the connection @proxy is for.
	Connection() DBusConnection
	// DefaultTimeout gets the timeout to use if -1 (specifying default timeout)
	// is passed as @timeout_msec in the g_dbus_proxy_call() and
	// g_dbus_proxy_call_sync() functions.
	//
	// See the BusProxy:g-default-timeout property for more details.
	DefaultTimeout() int
	// Flags gets the flags that @proxy was constructed with.
	Flags() DBusProXYFlags
	// InterfaceInfo returns the BusInterfaceInfo, if any, specifying the
	// interface that @proxy conforms to. See the BusProxy:g-interface-info
	// property for more details.
	InterfaceInfo() *DBusInterfaceInfo
	// InterfaceName gets the D-Bus interface name @proxy is for.
	InterfaceName() string
	// Name gets the name that @proxy was constructed for.
	Name() string
	// NameOwner: the unique name that owns the name that @proxy is for or nil
	// if no-one currently owns that name. You may connect to the
	// #GObject::notify signal to track changes to the BusProxy:g-name-owner
	// property.
	NameOwner() string
	// ObjectPath gets the object path @proxy is for.
	ObjectPath() string
	// SetCachedProperty: if @value is not nil, sets the cached value for the
	// property with name @property_name to the value in @value.
	//
	// If @value is nil, then the cached value is removed from the property
	// cache.
	//
	// If @proxy has an expected interface (see BusProxy:g-interface-info) and
	// @property_name is referenced by it, then @value is checked against the
	// type of the property.
	//
	// If the @value #GVariant is floating, it is consumed. This allows
	// convenient 'inline' use of g_variant_new(), e.g.
	//
	//    g_dbus_proxy_set_cached_property (proxy,
	//                                      "SomeProperty",
	//                                      g_variant_new ("(si)",
	//                                                    "A String",
	//                                                    42));
	//
	// Normally you will not need to use this method since @proxy is tracking
	// changes using the `org.freedesktop.DBus.Properties.PropertiesChanged`
	// D-Bus signal. However, for performance reasons an object may decide to
	// not use this signal for some properties and instead use a proprietary
	// out-of-band mechanism to transmit changes.
	//
	// As a concrete example, consider an object with a property
	// `ChatroomParticipants` which is an array of strings. Instead of
	// transmitting the same (long) array every time the property changes, it is
	// more efficient to only transmit the delta using e.g. signals
	// `ChatroomParticipantJoined(String name)` and
	// `ChatroomParticipantParted(String name)`.
	SetCachedProperty(propertyName string, value *glib.Variant)
	// SetDefaultTimeout sets the timeout to use if -1 (specifying default
	// timeout) is passed as @timeout_msec in the g_dbus_proxy_call() and
	// g_dbus_proxy_call_sync() functions.
	//
	// See the BusProxy:g-default-timeout property for more details.
	SetDefaultTimeout(timeoutMsec int)
	// SetInterfaceInfo: ensure that interactions with @proxy conform to the
	// given interface. See the BusProxy:g-interface-info property for more
	// details.
	SetInterfaceInfo(info *DBusInterfaceInfo)
}

// dBusProXY implements the DBusProXY interface.
type dBusProXY struct {
	gextras.Objector
	AsyncInitable
	DBusInterface
	Initable
}

var _ DBusProXY = (*dBusProXY)(nil)

// WrapDBusProXY wraps a GObject to the right type. It is
// primarily used internally.
func WrapDBusProXY(obj *externglib.Object) DBusProXY {
	return DBusProXY{
		Objector:      obj,
		AsyncInitable: WrapAsyncInitable(obj),
		DBusInterface: WrapDBusInterface(obj),
		Initable:      WrapInitable(obj),
	}
}

func marshalDBusProXY(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDBusProXY(obj), nil
}

// NewDBusProXYFinish constructs a class DBusProXY.
func NewDBusProXYFinish(res AsyncResult) (dBusProxy DBusProXY, goerr error) {
	var arg1 *C.GAsyncResult

	arg1 = (*C.GAsyncResult)(unsafe.Pointer(res.Native()))

	var cret C.GDBusProxy
	var cerr *C.GError

	cret = C.g_dbus_proxy_new_finish(arg1, cerr)

	var dBusProxy DBusProXY
	var goerr error

	dBusProxy = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(DBusProXY)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return dBusProxy, goerr
}

// NewDBusProXYForBusFinish constructs a class DBusProXY.
func NewDBusProXYForBusFinish(res AsyncResult) (dBusProxy DBusProXY, goerr error) {
	var arg1 *C.GAsyncResult

	arg1 = (*C.GAsyncResult)(unsafe.Pointer(res.Native()))

	var cret C.GDBusProxy
	var cerr *C.GError

	cret = C.g_dbus_proxy_new_for_bus_finish(arg1, cerr)

	var dBusProxy DBusProXY
	var goerr error

	dBusProxy = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(DBusProXY)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return dBusProxy, goerr
}

// NewDBusProXYForBusSync constructs a class DBusProXY.
func NewDBusProXYForBusSync(busType BusType, flags DBusProXYFlags, info *DBusInterfaceInfo, name string, objectPath string, interfaceName string, cancellable Cancellable) (dBusProxy DBusProXY, goerr error) {
	var arg1 C.GBusType
	var arg2 C.GDBusProxyFlags
	var arg3 *C.GDBusInterfaceInfo
	var arg4 *C.gchar
	var arg5 *C.gchar
	var arg6 *C.gchar
	var arg7 *C.GCancellable

	arg1 = (C.GBusType)(busType)
	arg2 = (C.GDBusProxyFlags)(flags)
	arg3 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(info.Native()))
	arg4 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg4))
	arg5 = (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(arg5))
	arg6 = (*C.gchar)(C.CString(interfaceName))
	defer C.free(unsafe.Pointer(arg6))
	arg7 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cret C.GDBusProxy
	var cerr *C.GError

	cret = C.g_dbus_proxy_new_for_bus_sync(arg1, arg2, arg3, arg4, arg5, arg6, arg7, cerr)

	var dBusProxy DBusProXY
	var goerr error

	dBusProxy = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(DBusProXY)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return dBusProxy, goerr
}

// NewDBusProXYSync constructs a class DBusProXY.
func NewDBusProXYSync(connection DBusConnection, flags DBusProXYFlags, info *DBusInterfaceInfo, name string, objectPath string, interfaceName string, cancellable Cancellable) (dBusProxy DBusProXY, goerr error) {
	var arg1 *C.GDBusConnection
	var arg2 C.GDBusProxyFlags
	var arg3 *C.GDBusInterfaceInfo
	var arg4 *C.gchar
	var arg5 *C.gchar
	var arg6 *C.gchar
	var arg7 *C.GCancellable

	arg1 = (*C.GDBusConnection)(unsafe.Pointer(connection.Native()))
	arg2 = (C.GDBusProxyFlags)(flags)
	arg3 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(info.Native()))
	arg4 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg4))
	arg5 = (*C.gchar)(C.CString(objectPath))
	defer C.free(unsafe.Pointer(arg5))
	arg6 = (*C.gchar)(C.CString(interfaceName))
	defer C.free(unsafe.Pointer(arg6))
	arg7 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cret C.GDBusProxy
	var cerr *C.GError

	cret = C.g_dbus_proxy_new_sync(arg1, arg2, arg3, arg4, arg5, arg6, arg7, cerr)

	var dBusProxy DBusProXY
	var goerr error

	dBusProxy = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(DBusProXY)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return dBusProxy, goerr
}

// Call: asynchronously invokes the @method_name method on @proxy.
//
// If @method_name contains any dots, then @name is split into interface and
// method name parts. This allows using @proxy for invoking methods on other
// interfaces.
//
// If the BusConnection associated with @proxy is closed then the operation
// will fail with G_IO_ERROR_CLOSED. If @cancellable is canceled, the
// operation will fail with G_IO_ERROR_CANCELLED. If @parameters contains a
// value not compatible with the D-Bus protocol, the operation fails with
// G_IO_ERROR_INVALID_ARGUMENT.
//
// If the @parameters #GVariant is floating, it is consumed. This allows
// convenient 'inline' use of g_variant_new(), e.g.:
//
//    g_dbus_proxy_call (proxy,
//                       "TwoStrings",
//                       g_variant_new ("(ss)",
//                                      "Thing One",
//                                      "Thing Two"),
//                       G_DBUS_CALL_FLAGS_NONE,
//                       -1,
//                       NULL,
//                       (GAsyncReadyCallback) two_strings_done,
//                       &data);
//
// If @proxy has an expected interface (see BusProxy:g-interface-info) and
// @method_name is referenced by it, then the return value is checked
// against the return type.
//
// This is an asynchronous method. When the operation is finished, @callback
// will be invoked in the [thread-default main
// context][g-main-context-push-thread-default] of the thread you are
// calling this method from. You can then call g_dbus_proxy_call_finish() to
// get the result of the operation. See g_dbus_proxy_call_sync() for the
// synchronous version of this method.
//
// If @callback is nil then the D-Bus method call message will be sent with
// the G_DBUS_MESSAGE_FLAGS_NO_REPLY_EXPECTED flag set.
func (p dBusProXY) Call() {
	var arg0 *C.GDBusProxy

	arg0 = (*C.GDBusProxy)(unsafe.Pointer(p.Native()))

	C.g_dbus_proxy_call(arg0)
}

// CallFinish finishes an operation started with g_dbus_proxy_call().
func (p dBusProXY) CallFinish(res AsyncResult) (variant *glib.Variant, goerr error) {
	var arg0 *C.GDBusProxy
	var arg1 *C.GAsyncResult

	arg0 = (*C.GDBusProxy)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(res.Native()))

	var cret *C.GVariant
	var cerr *C.GError

	cret = C.g_dbus_proxy_call_finish(arg0, arg1, cerr)

	var variant *glib.Variant
	var goerr error

	variant = glib.WrapVariant(unsafe.Pointer(cret))
	runtime.SetFinalizer(variant, func(v *glib.Variant) {
		C.free(unsafe.Pointer(v.Native()))
	})
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return variant, goerr
}

// CallSync: synchronously invokes the @method_name method on @proxy.
//
// If @method_name contains any dots, then @name is split into interface and
// method name parts. This allows using @proxy for invoking methods on other
// interfaces.
//
// If the BusConnection associated with @proxy is disconnected then the
// operation will fail with G_IO_ERROR_CLOSED. If @cancellable is canceled,
// the operation will fail with G_IO_ERROR_CANCELLED. If @parameters
// contains a value not compatible with the D-Bus protocol, the operation
// fails with G_IO_ERROR_INVALID_ARGUMENT.
//
// If the @parameters #GVariant is floating, it is consumed. This allows
// convenient 'inline' use of g_variant_new(), e.g.:
//
//    g_dbus_proxy_call_sync (proxy,
//                            "TwoStrings",
//                            g_variant_new ("(ss)",
//                                           "Thing One",
//                                           "Thing Two"),
//                            G_DBUS_CALL_FLAGS_NONE,
//                            -1,
//                            NULL,
//                            &error);
//
// The calling thread is blocked until a reply is received. See
// g_dbus_proxy_call() for the asynchronous version of this method.
//
// If @proxy has an expected interface (see BusProxy:g-interface-info) and
// @method_name is referenced by it, then the return value is checked
// against the return type.
func (p dBusProXY) CallSync(methodName string, parameters *glib.Variant, flags DBusCallFlags, timeoutMsec int, cancellable Cancellable) (variant *glib.Variant, goerr error) {
	var arg0 *C.GDBusProxy
	var arg1 *C.gchar
	var arg2 *C.GVariant
	var arg3 C.GDBusCallFlags
	var arg4 C.gint
	var arg5 *C.GCancellable

	arg0 = (*C.GDBusProxy)(unsafe.Pointer(p.Native()))
	arg1 = (*C.gchar)(C.CString(methodName))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GVariant)(unsafe.Pointer(parameters.Native()))
	arg3 = (C.GDBusCallFlags)(flags)
	arg4 = C.gint(timeoutMsec)
	arg5 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cret *C.GVariant
	var cerr *C.GError

	cret = C.g_dbus_proxy_call_sync(arg0, arg1, arg2, arg3, arg4, arg5, cerr)

	var variant *glib.Variant
	var goerr error

	variant = glib.WrapVariant(unsafe.Pointer(cret))
	runtime.SetFinalizer(variant, func(v *glib.Variant) {
		C.free(unsafe.Pointer(v.Native()))
	})
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return variant, goerr
}

// CallWithUnixFdList: like g_dbus_proxy_call() but also takes a FDList
// object.
//
// This method is only available on UNIX.
func (p dBusProXY) CallWithUnixFdList() {
	var arg0 *C.GDBusProxy

	arg0 = (*C.GDBusProxy)(unsafe.Pointer(p.Native()))

	C.g_dbus_proxy_call_with_unix_fd_list(arg0)
}

// CallWithUnixFdListFinish finishes an operation started with
// g_dbus_proxy_call_with_unix_fd_list().
func (p dBusProXY) CallWithUnixFdListFinish(res AsyncResult) (outFdList UnixFDList, variant *glib.Variant, goerr error) {
	var arg0 *C.GDBusProxy
	var arg2 *C.GAsyncResult

	arg0 = (*C.GDBusProxy)(unsafe.Pointer(p.Native()))
	arg2 = (*C.GAsyncResult)(unsafe.Pointer(res.Native()))

	var arg1 *C.GUnixFDList
	var cret *C.GVariant
	var cerr *C.GError

	cret = C.g_dbus_proxy_call_with_unix_fd_list_finish(arg0, arg2, &arg1, cerr)

	var outFdList UnixFDList
	var variant *glib.Variant
	var goerr error

	outFdList = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(arg1.Native()))).(UnixFDList)
	variant = glib.WrapVariant(unsafe.Pointer(cret))
	runtime.SetFinalizer(variant, func(v *glib.Variant) {
		C.free(unsafe.Pointer(v.Native()))
	})
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return outFdList, variant, goerr
}

// CallWithUnixFdListSync: like g_dbus_proxy_call_sync() but also takes and
// returns FDList objects.
//
// This method is only available on UNIX.
func (p dBusProXY) CallWithUnixFdListSync(methodName string, parameters *glib.Variant, flags DBusCallFlags, timeoutMsec int, fdList UnixFDList, cancellable Cancellable) (outFdList UnixFDList, variant *glib.Variant, goerr error) {
	var arg0 *C.GDBusProxy
	var arg1 *C.gchar
	var arg2 *C.GVariant
	var arg3 C.GDBusCallFlags
	var arg4 C.gint
	var arg5 *C.GUnixFDList
	var arg7 *C.GCancellable

	arg0 = (*C.GDBusProxy)(unsafe.Pointer(p.Native()))
	arg1 = (*C.gchar)(C.CString(methodName))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GVariant)(unsafe.Pointer(parameters.Native()))
	arg3 = (C.GDBusCallFlags)(flags)
	arg4 = C.gint(timeoutMsec)
	arg5 = (*C.GUnixFDList)(unsafe.Pointer(fdList.Native()))
	arg7 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var arg6 *C.GUnixFDList
	var cret *C.GVariant
	var cerr *C.GError

	cret = C.g_dbus_proxy_call_with_unix_fd_list_sync(arg0, arg1, arg2, arg3, arg4, arg5, arg7, &arg6, cerr)

	var outFdList UnixFDList
	var variant *glib.Variant
	var goerr error

	outFdList = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(arg6.Native()))).(UnixFDList)
	variant = glib.WrapVariant(unsafe.Pointer(cret))
	runtime.SetFinalizer(variant, func(v *glib.Variant) {
		C.free(unsafe.Pointer(v.Native()))
	})
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return outFdList, variant, goerr
}

// CachedProperty looks up the value for a property from the cache. This
// call does no blocking IO.
//
// If @proxy has an expected interface (see BusProxy:g-interface-info) and
// @property_name is referenced by it, then @value is checked against the
// type of the property.
func (p dBusProXY) CachedProperty(propertyName string) *glib.Variant {
	var arg0 *C.GDBusProxy
	var arg1 *C.gchar

	arg0 = (*C.GDBusProxy)(unsafe.Pointer(p.Native()))
	arg1 = (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GVariant

	cret = C.g_dbus_proxy_get_cached_property(arg0, arg1)

	var variant *glib.Variant

	variant = glib.WrapVariant(unsafe.Pointer(cret))
	runtime.SetFinalizer(variant, func(v *glib.Variant) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return variant
}

// CachedPropertyNames gets the names of all cached properties on @proxy.
func (p dBusProXY) CachedPropertyNames() []string {
	var arg0 *C.GDBusProxy

	arg0 = (*C.GDBusProxy)(unsafe.Pointer(p.Native()))

	var cret **C.gchar

	cret = C.g_dbus_proxy_get_cached_property_names(arg0)

	var utf8s []string

	{
		var length int
		for p := cret; *p != 0; p = (**C.gchar)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(cret), int(length))

		utf8s = make([]string, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			utf8s = C.GoString(cret)
			defer C.free(unsafe.Pointer(cret))
		}
	}

	return utf8s
}

// Connection gets the connection @proxy is for.
func (p dBusProXY) Connection() DBusConnection {
	var arg0 *C.GDBusProxy

	arg0 = (*C.GDBusProxy)(unsafe.Pointer(p.Native()))

	var cret *C.GDBusConnection

	cret = C.g_dbus_proxy_get_connection(arg0)

	var dBusConnection DBusConnection

	dBusConnection = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(DBusConnection)

	return dBusConnection
}

// DefaultTimeout gets the timeout to use if -1 (specifying default timeout)
// is passed as @timeout_msec in the g_dbus_proxy_call() and
// g_dbus_proxy_call_sync() functions.
//
// See the BusProxy:g-default-timeout property for more details.
func (p dBusProXY) DefaultTimeout() int {
	var arg0 *C.GDBusProxy

	arg0 = (*C.GDBusProxy)(unsafe.Pointer(p.Native()))

	var cret C.gint

	cret = C.g_dbus_proxy_get_default_timeout(arg0)

	var gint int

	gint = (int)(cret)

	return gint
}

// Flags gets the flags that @proxy was constructed with.
func (p dBusProXY) Flags() DBusProXYFlags {
	var arg0 *C.GDBusProxy

	arg0 = (*C.GDBusProxy)(unsafe.Pointer(p.Native()))

	var cret C.GDBusProxyFlags

	cret = C.g_dbus_proxy_get_flags(arg0)

	var dBusProxyFlags DBusProXYFlags

	dBusProxyFlags = DBusProXYFlags(cret)

	return dBusProxyFlags
}

// InterfaceInfo returns the BusInterfaceInfo, if any, specifying the
// interface that @proxy conforms to. See the BusProxy:g-interface-info
// property for more details.
func (p dBusProXY) InterfaceInfo() *DBusInterfaceInfo {
	var arg0 *C.GDBusProxy

	arg0 = (*C.GDBusProxy)(unsafe.Pointer(p.Native()))

	var cret *C.GDBusInterfaceInfo

	cret = C.g_dbus_proxy_get_interface_info(arg0)

	var dBusInterfaceInfo *DBusInterfaceInfo

	dBusInterfaceInfo = WrapDBusInterfaceInfo(unsafe.Pointer(cret))

	return dBusInterfaceInfo
}

// InterfaceName gets the D-Bus interface name @proxy is for.
func (p dBusProXY) InterfaceName() string {
	var arg0 *C.GDBusProxy

	arg0 = (*C.GDBusProxy)(unsafe.Pointer(p.Native()))

	var cret *C.gchar

	cret = C.g_dbus_proxy_get_interface_name(arg0)

	var utf8 string

	utf8 = C.GoString(cret)

	return utf8
}

// Name gets the name that @proxy was constructed for.
func (p dBusProXY) Name() string {
	var arg0 *C.GDBusProxy

	arg0 = (*C.GDBusProxy)(unsafe.Pointer(p.Native()))

	var cret *C.gchar

	cret = C.g_dbus_proxy_get_name(arg0)

	var utf8 string

	utf8 = C.GoString(cret)

	return utf8
}

// NameOwner: the unique name that owns the name that @proxy is for or nil
// if no-one currently owns that name. You may connect to the
// #GObject::notify signal to track changes to the BusProxy:g-name-owner
// property.
func (p dBusProXY) NameOwner() string {
	var arg0 *C.GDBusProxy

	arg0 = (*C.GDBusProxy)(unsafe.Pointer(p.Native()))

	var cret *C.gchar

	cret = C.g_dbus_proxy_get_name_owner(arg0)

	var utf8 string

	utf8 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return utf8
}

// ObjectPath gets the object path @proxy is for.
func (p dBusProXY) ObjectPath() string {
	var arg0 *C.GDBusProxy

	arg0 = (*C.GDBusProxy)(unsafe.Pointer(p.Native()))

	var cret *C.gchar

	cret = C.g_dbus_proxy_get_object_path(arg0)

	var utf8 string

	utf8 = C.GoString(cret)

	return utf8
}

// SetCachedProperty: if @value is not nil, sets the cached value for the
// property with name @property_name to the value in @value.
//
// If @value is nil, then the cached value is removed from the property
// cache.
//
// If @proxy has an expected interface (see BusProxy:g-interface-info) and
// @property_name is referenced by it, then @value is checked against the
// type of the property.
//
// If the @value #GVariant is floating, it is consumed. This allows
// convenient 'inline' use of g_variant_new(), e.g.
//
//    g_dbus_proxy_set_cached_property (proxy,
//                                      "SomeProperty",
//                                      g_variant_new ("(si)",
//                                                    "A String",
//                                                    42));
//
// Normally you will not need to use this method since @proxy is tracking
// changes using the `org.freedesktop.DBus.Properties.PropertiesChanged`
// D-Bus signal. However, for performance reasons an object may decide to
// not use this signal for some properties and instead use a proprietary
// out-of-band mechanism to transmit changes.
//
// As a concrete example, consider an object with a property
// `ChatroomParticipants` which is an array of strings. Instead of
// transmitting the same (long) array every time the property changes, it is
// more efficient to only transmit the delta using e.g. signals
// `ChatroomParticipantJoined(String name)` and
// `ChatroomParticipantParted(String name)`.
func (p dBusProXY) SetCachedProperty(propertyName string, value *glib.Variant) {
	var arg0 *C.GDBusProxy
	var arg1 *C.gchar
	var arg2 *C.GVariant

	arg0 = (*C.GDBusProxy)(unsafe.Pointer(p.Native()))
	arg1 = (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GVariant)(unsafe.Pointer(value.Native()))

	C.g_dbus_proxy_set_cached_property(arg0, arg1, arg2)
}

// SetDefaultTimeout sets the timeout to use if -1 (specifying default
// timeout) is passed as @timeout_msec in the g_dbus_proxy_call() and
// g_dbus_proxy_call_sync() functions.
//
// See the BusProxy:g-default-timeout property for more details.
func (p dBusProXY) SetDefaultTimeout(timeoutMsec int) {
	var arg0 *C.GDBusProxy
	var arg1 C.gint

	arg0 = (*C.GDBusProxy)(unsafe.Pointer(p.Native()))
	arg1 = C.gint(timeoutMsec)

	C.g_dbus_proxy_set_default_timeout(arg0, arg1)
}

// SetInterfaceInfo: ensure that interactions with @proxy conform to the
// given interface. See the BusProxy:g-interface-info property for more
// details.
func (p dBusProXY) SetInterfaceInfo(info *DBusInterfaceInfo) {
	var arg0 *C.GDBusProxy
	var arg1 *C.GDBusInterfaceInfo

	arg0 = (*C.GDBusProxy)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(info.Native()))

	C.g_dbus_proxy_set_interface_info(arg0, arg1)
}
