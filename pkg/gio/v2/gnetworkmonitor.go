// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"github.com/diamondburned/gotk4/internal/gerror"
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
		{T: externglib.Type(C.g_network_monitor_get_type()), F: marshalNetworkMonitor},
	})
}

// NetworkMonitorOverrider contains methods that are overridable. This
// interface is a subset of the interface NetworkMonitor.
type NetworkMonitorOverrider interface {
	// CanReach attempts to determine whether or not the host pointed to by
	// @connectable can be reached, without actually trying to connect to it.
	//
	// This may return true even when Monitor:network-available is false, if,
	// for example, @monitor can determine that @connectable refers to a host on
	// a local network.
	//
	// If @monitor believes that an attempt to connect to @connectable will
	// succeed, it will return true. Otherwise, it will return false and set
	// @error to an appropriate error (such as G_IO_ERROR_HOST_UNREACHABLE).
	//
	// Note that although this does not attempt to connect to @connectable, it
	// may still block for a brief period of time (eg, trying to do multicast
	// DNS on the local network), so if you do not want to block, you should use
	// g_network_monitor_can_reach_async().
	CanReach(connectable SocketConnectable, cancellable Cancellable) error
	// CanReachAsync: asynchronously attempts to determine whether or not the
	// host pointed to by @connectable can be reached, without actually trying
	// to connect to it.
	//
	// For more details, see g_network_monitor_can_reach().
	//
	// When the operation is finished, @callback will be called. You can then
	// call g_network_monitor_can_reach_finish() to get the result of the
	// operation.
	CanReachAsync()
	// CanReachFinish finishes an async network connectivity test. See
	// g_network_monitor_can_reach_async().
	CanReachFinish(result AsyncResult) error

	NetworkChanged(networkAvailable bool)
}

// NetworkMonitor provides an easy-to-use cross-platform API for monitoring
// network connectivity. On Linux, the available implementations are based on
// the kernel's netlink interface and on NetworkManager.
//
// There is also an implementation for use inside Flatpak sandboxes.
type NetworkMonitor interface {
	Initable
	NetworkMonitorOverrider

	// Connectivity gets a more detailed networking state than
	// g_network_monitor_get_network_available().
	//
	// If Monitor:network-available is false, then the connectivity state will
	// be G_NETWORK_CONNECTIVITY_LOCAL.
	//
	// If Monitor:network-available is true, then the connectivity state will be
	// G_NETWORK_CONNECTIVITY_FULL (if there is full Internet connectivity),
	// G_NETWORK_CONNECTIVITY_LIMITED (if the host has a default route, but
	// appears to be unable to actually reach the full Internet), or
	// G_NETWORK_CONNECTIVITY_PORTAL (if the host is trapped behind a "captive
	// portal" that requires some sort of login or acknowledgement before
	// allowing full Internet access).
	//
	// Note that in the case of G_NETWORK_CONNECTIVITY_LIMITED and
	// G_NETWORK_CONNECTIVITY_PORTAL, it is possible that some sites are
	// reachable but others are not. In this case, applications can attempt to
	// connect to remote servers, but should gracefully fall back to their
	// "offline" behavior if the connection attempt fails.
	Connectivity() NetworkConnectivity
	// NetworkAvailable checks if the network is available. "Available" here
	// means that the system has a default route available for at least one of
	// IPv4 or IPv6. It does not necessarily imply that the public Internet is
	// reachable. See Monitor:network-available for more details.
	NetworkAvailable() bool
	// NetworkMetered checks if the network is metered. See
	// Monitor:network-metered for more details.
	NetworkMetered() bool
}

// networkMonitor implements the NetworkMonitor interface.
type networkMonitor struct {
	Initable
}

var _ NetworkMonitor = (*networkMonitor)(nil)

// WrapNetworkMonitor wraps a GObject to a type that implements interface
// NetworkMonitor. It is primarily used internally.
func WrapNetworkMonitor(obj *externglib.Object) NetworkMonitor {
	return NetworkMonitor{
		Initable: WrapInitable(obj),
	}
}

func marshalNetworkMonitor(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNetworkMonitor(obj), nil
}

// CanReach attempts to determine whether or not the host pointed to by
// @connectable can be reached, without actually trying to connect to it.
//
// This may return true even when Monitor:network-available is false, if,
// for example, @monitor can determine that @connectable refers to a host on
// a local network.
//
// If @monitor believes that an attempt to connect to @connectable will
// succeed, it will return true. Otherwise, it will return false and set
// @error to an appropriate error (such as G_IO_ERROR_HOST_UNREACHABLE).
//
// Note that although this does not attempt to connect to @connectable, it
// may still block for a brief period of time (eg, trying to do multicast
// DNS on the local network), so if you do not want to block, you should use
// g_network_monitor_can_reach_async().
func (m networkMonitor) CanReach(connectable SocketConnectable, cancellable Cancellable) error {
	var arg0 *C.GNetworkMonitor
	var arg1 *C.GSocketConnectable
	var arg2 *C.GCancellable

	arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GSocketConnectable)(unsafe.Pointer(connectable.Native()))
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var cerr *C.GError

	C.g_network_monitor_can_reach(arg0, arg1, arg2, cerr)

	var goerr error

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// CanReachAsync: asynchronously attempts to determine whether or not the
// host pointed to by @connectable can be reached, without actually trying
// to connect to it.
//
// For more details, see g_network_monitor_can_reach().
//
// When the operation is finished, @callback will be called. You can then
// call g_network_monitor_can_reach_finish() to get the result of the
// operation.
func (m networkMonitor) CanReachAsync() {
	var arg0 *C.GNetworkMonitor

	arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(m.Native()))

	C.g_network_monitor_can_reach_async(arg0)
}

// CanReachFinish finishes an async network connectivity test. See
// g_network_monitor_can_reach_async().
func (m networkMonitor) CanReachFinish(result AsyncResult) error {
	var arg0 *C.GNetworkMonitor
	var arg1 *C.GAsyncResult

	arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var cerr *C.GError

	C.g_network_monitor_can_reach_finish(arg0, arg1, cerr)

	var goerr error

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// Connectivity gets a more detailed networking state than
// g_network_monitor_get_network_available().
//
// If Monitor:network-available is false, then the connectivity state will
// be G_NETWORK_CONNECTIVITY_LOCAL.
//
// If Monitor:network-available is true, then the connectivity state will be
// G_NETWORK_CONNECTIVITY_FULL (if there is full Internet connectivity),
// G_NETWORK_CONNECTIVITY_LIMITED (if the host has a default route, but
// appears to be unable to actually reach the full Internet), or
// G_NETWORK_CONNECTIVITY_PORTAL (if the host is trapped behind a "captive
// portal" that requires some sort of login or acknowledgement before
// allowing full Internet access).
//
// Note that in the case of G_NETWORK_CONNECTIVITY_LIMITED and
// G_NETWORK_CONNECTIVITY_PORTAL, it is possible that some sites are
// reachable but others are not. In this case, applications can attempt to
// connect to remote servers, but should gracefully fall back to their
// "offline" behavior if the connection attempt fails.
func (m networkMonitor) Connectivity() NetworkConnectivity {
	var arg0 *C.GNetworkMonitor

	arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(m.Native()))

	var cret C.GNetworkConnectivity

	cret = C.g_network_monitor_get_connectivity(arg0)

	var networkConnectivity NetworkConnectivity

	networkConnectivity = NetworkConnectivity(cret)

	return networkConnectivity
}

// NetworkAvailable checks if the network is available. "Available" here
// means that the system has a default route available for at least one of
// IPv4 or IPv6. It does not necessarily imply that the public Internet is
// reachable. See Monitor:network-available for more details.
func (m networkMonitor) NetworkAvailable() bool {
	var arg0 *C.GNetworkMonitor

	arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(m.Native()))

	var cret C.gboolean

	cret = C.g_network_monitor_get_network_available(arg0)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// NetworkMetered checks if the network is metered. See
// Monitor:network-metered for more details.
func (m networkMonitor) NetworkMetered() bool {
	var arg0 *C.GNetworkMonitor

	arg0 = (*C.GNetworkMonitor)(unsafe.Pointer(m.Native()))

	var cret C.gboolean

	cret = C.g_network_monitor_get_network_metered(arg0)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}
