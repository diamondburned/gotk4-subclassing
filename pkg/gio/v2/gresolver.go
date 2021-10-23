// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
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
// #include <glib-object.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_resolver_name_lookup_flags_get_type()), F: marshalResolverNameLookupFlags},
		{T: externglib.Type(C.g_resolver_get_type()), F: marshalResolverer},
	})
}

// ResolverNameLookupFlags flags to modify lookup behavior.
type ResolverNameLookupFlags C.guint

const (
	// ResolverNameLookupFlagsDefault: default behavior (same as
	// g_resolver_lookup_by_name()).
	ResolverNameLookupFlagsDefault ResolverNameLookupFlags = 0b0
	// ResolverNameLookupFlagsIPv4Only: only resolve ipv4 addresses.
	ResolverNameLookupFlagsIPv4Only ResolverNameLookupFlags = 0b1
	// ResolverNameLookupFlagsIPv6Only: only resolve ipv6 addresses.
	ResolverNameLookupFlagsIPv6Only ResolverNameLookupFlags = 0b10
)

func marshalResolverNameLookupFlags(p uintptr) (interface{}, error) {
	return ResolverNameLookupFlags(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for ResolverNameLookupFlags.
func (r ResolverNameLookupFlags) String() string {
	if r == 0 {
		return "ResolverNameLookupFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(94)

	for r != 0 {
		next := r & (r - 1)
		bit := r - next

		switch bit {
		case ResolverNameLookupFlagsDefault:
			builder.WriteString("Default|")
		case ResolverNameLookupFlagsIPv4Only:
			builder.WriteString("IPv4Only|")
		case ResolverNameLookupFlagsIPv6Only:
			builder.WriteString("IPv6Only|")
		default:
			builder.WriteString(fmt.Sprintf("ResolverNameLookupFlags(0b%b)|", bit))
		}

		r = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if r contains other.
func (r ResolverNameLookupFlags) Has(other ResolverNameLookupFlags) bool {
	return (r & other) == other
}

// ResolverOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ResolverOverrider interface {
	// LookupByAddress: synchronously reverse-resolves address to determine its
	// associated hostname.
	//
	// If the DNS resolution fails, error (if non-NULL) will be set to a value
	// from Error.
	//
	// If cancellable is non-NULL, it can be used to cancel the operation, in
	// which case error (if non-NULL) will be set to G_IO_ERROR_CANCELLED.
	LookupByAddress(ctx context.Context, address *InetAddress) (string, error)
	// LookupByAddressAsync begins asynchronously reverse-resolving address to
	// determine its associated hostname, and eventually calls callback, which
	// must call g_resolver_lookup_by_address_finish() to get the final result.
	LookupByAddressAsync(ctx context.Context, address *InetAddress, callback AsyncReadyCallback)
	// LookupByAddressFinish retrieves the result of a previous call to
	// g_resolver_lookup_by_address_async().
	//
	// If the DNS resolution failed, error (if non-NULL) will be set to a value
	// from Error. If the operation was cancelled, error will be set to
	// G_IO_ERROR_CANCELLED.
	LookupByAddressFinish(result AsyncResulter) (string, error)
	// LookupByName: synchronously resolves hostname to determine its associated
	// IP address(es). hostname may be an ASCII-only or UTF-8 hostname, or the
	// textual form of an IP address (in which case this just becomes a wrapper
	// around g_inet_address_new_from_string()).
	//
	// On success, g_resolver_lookup_by_name() will return a non-empty #GList of
	// Address, sorted in order of preference and guaranteed to not contain
	// duplicates. That is, if using the result to connect to hostname, you
	// should attempt to connect to the first address first, then the second if
	// the first fails, etc. If you are using the result to listen on a socket,
	// it is appropriate to add each result using e.g.
	// g_socket_listener_add_address().
	//
	// If the DNS resolution fails, error (if non-NULL) will be set to a value
	// from Error and NULL will be returned.
	//
	// If cancellable is non-NULL, it can be used to cancel the operation, in
	// which case error (if non-NULL) will be set to G_IO_ERROR_CANCELLED.
	//
	// If you are planning to connect to a socket on the resolved IP address, it
	// may be easier to create a Address and use its Connectable interface.
	LookupByName(ctx context.Context, hostname string) ([]InetAddress, error)
	// LookupByNameAsync begins asynchronously resolving hostname to determine
	// its associated IP address(es), and eventually calls callback, which must
	// call g_resolver_lookup_by_name_finish() to get the result. See
	// g_resolver_lookup_by_name() for more details.
	LookupByNameAsync(ctx context.Context, hostname string, callback AsyncReadyCallback)
	// LookupByNameFinish retrieves the result of a call to
	// g_resolver_lookup_by_name_async().
	//
	// If the DNS resolution failed, error (if non-NULL) will be set to a value
	// from Error. If the operation was cancelled, error will be set to
	// G_IO_ERROR_CANCELLED.
	LookupByNameFinish(result AsyncResulter) ([]InetAddress, error)
	// LookupByNameWithFlags: this differs from g_resolver_lookup_by_name() in
	// that you can modify the lookup behavior with flags. For example this can
	// be used to limit results with RESOLVER_NAME_LOOKUP_FLAGS_IPV4_ONLY.
	LookupByNameWithFlags(ctx context.Context, hostname string, flags ResolverNameLookupFlags) ([]InetAddress, error)
	// LookupByNameWithFlagsAsync begins asynchronously resolving hostname to
	// determine its associated IP address(es), and eventually calls callback,
	// which must call g_resolver_lookup_by_name_with_flags_finish() to get the
	// result. See g_resolver_lookup_by_name() for more details.
	LookupByNameWithFlagsAsync(ctx context.Context, hostname string, flags ResolverNameLookupFlags, callback AsyncReadyCallback)
	// LookupByNameWithFlagsFinish retrieves the result of a call to
	// g_resolver_lookup_by_name_with_flags_async().
	//
	// If the DNS resolution failed, error (if non-NULL) will be set to a value
	// from Error. If the operation was cancelled, error will be set to
	// G_IO_ERROR_CANCELLED.
	LookupByNameWithFlagsFinish(result AsyncResulter) ([]InetAddress, error)
	// LookupRecords: synchronously performs a DNS record lookup for the given
	// rrname and returns a list of records as #GVariant tuples. See RecordType
	// for information on what the records contain for each record_type.
	//
	// If the DNS resolution fails, error (if non-NULL) will be set to a value
	// from Error and NULL will be returned.
	//
	// If cancellable is non-NULL, it can be used to cancel the operation, in
	// which case error (if non-NULL) will be set to G_IO_ERROR_CANCELLED.
	LookupRecords(ctx context.Context, rrname string, recordType ResolverRecordType) ([]*glib.Variant, error)
	// LookupRecordsAsync begins asynchronously performing a DNS lookup for the
	// given rrname, and eventually calls callback, which must call
	// g_resolver_lookup_records_finish() to get the final result. See
	// g_resolver_lookup_records() for more details.
	LookupRecordsAsync(ctx context.Context, rrname string, recordType ResolverRecordType, callback AsyncReadyCallback)
	// LookupRecordsFinish retrieves the result of a previous call to
	// g_resolver_lookup_records_async(). Returns a non-empty list of records as
	// #GVariant tuples. See RecordType for information on what the records
	// contain.
	//
	// If the DNS resolution failed, error (if non-NULL) will be set to a value
	// from Error. If the operation was cancelled, error will be set to
	// G_IO_ERROR_CANCELLED.
	LookupRecordsFinish(result AsyncResulter) ([]*glib.Variant, error)
	LookupServiceAsync(ctx context.Context, rrname string, callback AsyncReadyCallback)
	// LookupServiceFinish retrieves the result of a previous call to
	// g_resolver_lookup_service_async().
	//
	// If the DNS resolution failed, error (if non-NULL) will be set to a value
	// from Error. If the operation was cancelled, error will be set to
	// G_IO_ERROR_CANCELLED.
	LookupServiceFinish(result AsyncResulter) ([]*SrvTarget, error)
	Reload()
}

// Resolver provides cancellable synchronous and asynchronous DNS resolution,
// for hostnames (g_resolver_lookup_by_address(), g_resolver_lookup_by_name()
// and their async variants) and SRV (service) records
// (g_resolver_lookup_service()).
//
// Address and Service provide wrappers around #GResolver functionality that
// also implement Connectable, making it easy to connect to a remote
// host/service.
type Resolver struct {
	*externglib.Object
}

// Resolverer describes types inherited from class Resolver.
// To get the original type, the caller must assert this to an interface or
// another type.
type Resolverer interface {
	externglib.Objector

	// BaseResolver returns the underlying base class.
	BaseResolver() *Resolver
}

var _ Resolverer = (*Resolver)(nil)

func wrapResolver(obj *externglib.Object) *Resolver {
	return &Resolver{
		Object: obj,
	}
}

func marshalResolverer(p uintptr) (interface{}, error) {
	return wrapResolver(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// LookupByAddress: synchronously reverse-resolves address to determine its
// associated hostname.
//
// If the DNS resolution fails, error (if non-NULL) will be set to a value from
// Error.
//
// If cancellable is non-NULL, it can be used to cancel the operation, in which
// case error (if non-NULL) will be set to G_IO_ERROR_CANCELLED.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - address to reverse-resolve.
//
func (resolver *Resolver) LookupByAddress(ctx context.Context, address *InetAddress) (string, error) {
	var _arg0 *C.GResolver    // out
	var _arg2 *C.GCancellable // out
	var _arg1 *C.GInetAddress // out
	var _cret *C.gchar        // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GResolver)(unsafe.Pointer(resolver.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))

	_cret = C.g_resolver_lookup_by_address(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(address)

	var _utf8 string // out
	var _goerr error // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _utf8, _goerr
}

// LookupByAddressAsync begins asynchronously reverse-resolving address to
// determine its associated hostname, and eventually calls callback, which must
// call g_resolver_lookup_by_address_finish() to get the final result.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - address to reverse-resolve.
//    - callback to call after resolution completes.
//
func (resolver *Resolver) LookupByAddressAsync(ctx context.Context, address *InetAddress, callback AsyncReadyCallback) {
	var _arg0 *C.GResolver          // out
	var _arg2 *C.GCancellable       // out
	var _arg1 *C.GInetAddress       // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GResolver)(unsafe.Pointer(resolver.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.GInetAddress)(unsafe.Pointer(address.Native()))
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_resolver_lookup_by_address_async(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(address)
	runtime.KeepAlive(callback)
}

// LookupByAddressFinish retrieves the result of a previous call to
// g_resolver_lookup_by_address_async().
//
// If the DNS resolution failed, error (if non-NULL) will be set to a value from
// Error. If the operation was cancelled, error will be set to
// G_IO_ERROR_CANCELLED.
//
// The function takes the following parameters:
//
//    - result passed to your ReadyCallback.
//
func (resolver *Resolver) LookupByAddressFinish(result AsyncResulter) (string, error) {
	var _arg0 *C.GResolver    // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.gchar        // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GResolver)(unsafe.Pointer(resolver.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_resolver_lookup_by_address_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(result)

	var _utf8 string // out
	var _goerr error // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _utf8, _goerr
}

// LookupByName: synchronously resolves hostname to determine its associated IP
// address(es). hostname may be an ASCII-only or UTF-8 hostname, or the textual
// form of an IP address (in which case this just becomes a wrapper around
// g_inet_address_new_from_string()).
//
// On success, g_resolver_lookup_by_name() will return a non-empty #GList of
// Address, sorted in order of preference and guaranteed to not contain
// duplicates. That is, if using the result to connect to hostname, you should
// attempt to connect to the first address first, then the second if the first
// fails, etc. If you are using the result to listen on a socket, it is
// appropriate to add each result using e.g. g_socket_listener_add_address().
//
// If the DNS resolution fails, error (if non-NULL) will be set to a value from
// Error and NULL will be returned.
//
// If cancellable is non-NULL, it can be used to cancel the operation, in which
// case error (if non-NULL) will be set to G_IO_ERROR_CANCELLED.
//
// If you are planning to connect to a socket on the resolved IP address, it may
// be easier to create a Address and use its Connectable interface.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - hostname to look up.
//
func (resolver *Resolver) LookupByName(ctx context.Context, hostname string) ([]InetAddress, error) {
	var _arg0 *C.GResolver    // out
	var _arg2 *C.GCancellable // out
	var _arg1 *C.gchar        // out
	var _cret *C.GList        // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GResolver)(unsafe.Pointer(resolver.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(hostname)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_resolver_lookup_by_name(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(hostname)

	var _list []InetAddress // out
	var _goerr error        // out

	_list = make([]InetAddress, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GInetAddress)(v)
		var dst InetAddress // out
		dst = *wrapInetAddress(externglib.AssumeOwnership(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _list, _goerr
}

// LookupByNameAsync begins asynchronously resolving hostname to determine its
// associated IP address(es), and eventually calls callback, which must call
// g_resolver_lookup_by_name_finish() to get the result. See
// g_resolver_lookup_by_name() for more details.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - hostname to look up the address of.
//    - callback to call after resolution completes.
//
func (resolver *Resolver) LookupByNameAsync(ctx context.Context, hostname string, callback AsyncReadyCallback) {
	var _arg0 *C.GResolver          // out
	var _arg2 *C.GCancellable       // out
	var _arg1 *C.gchar              // out
	var _arg3 C.GAsyncReadyCallback // out
	var _arg4 C.gpointer

	_arg0 = (*C.GResolver)(unsafe.Pointer(resolver.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(hostname)))
	defer C.free(unsafe.Pointer(_arg1))
	if callback != nil {
		_arg3 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg4 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_resolver_lookup_by_name_async(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(hostname)
	runtime.KeepAlive(callback)
}

// LookupByNameFinish retrieves the result of a call to
// g_resolver_lookup_by_name_async().
//
// If the DNS resolution failed, error (if non-NULL) will be set to a value from
// Error. If the operation was cancelled, error will be set to
// G_IO_ERROR_CANCELLED.
//
// The function takes the following parameters:
//
//    - result passed to your ReadyCallback.
//
func (resolver *Resolver) LookupByNameFinish(result AsyncResulter) ([]InetAddress, error) {
	var _arg0 *C.GResolver    // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GList        // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GResolver)(unsafe.Pointer(resolver.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_resolver_lookup_by_name_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(result)

	var _list []InetAddress // out
	var _goerr error        // out

	_list = make([]InetAddress, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GInetAddress)(v)
		var dst InetAddress // out
		dst = *wrapInetAddress(externglib.AssumeOwnership(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _list, _goerr
}

// LookupByNameWithFlags: this differs from g_resolver_lookup_by_name() in that
// you can modify the lookup behavior with flags. For example this can be used
// to limit results with RESOLVER_NAME_LOOKUP_FLAGS_IPV4_ONLY.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - hostname to look up.
//    - flags: extra NameLookupFlags for the lookup.
//
func (resolver *Resolver) LookupByNameWithFlags(ctx context.Context, hostname string, flags ResolverNameLookupFlags) ([]InetAddress, error) {
	var _arg0 *C.GResolver               // out
	var _arg3 *C.GCancellable            // out
	var _arg1 *C.gchar                   // out
	var _arg2 C.GResolverNameLookupFlags // out
	var _cret *C.GList                   // in
	var _cerr *C.GError                  // in

	_arg0 = (*C.GResolver)(unsafe.Pointer(resolver.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(hostname)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GResolverNameLookupFlags(flags)

	_cret = C.g_resolver_lookup_by_name_with_flags(_arg0, _arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(hostname)
	runtime.KeepAlive(flags)

	var _list []InetAddress // out
	var _goerr error        // out

	_list = make([]InetAddress, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GInetAddress)(v)
		var dst InetAddress // out
		dst = *wrapInetAddress(externglib.AssumeOwnership(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _list, _goerr
}

// LookupByNameWithFlagsAsync begins asynchronously resolving hostname to
// determine its associated IP address(es), and eventually calls callback, which
// must call g_resolver_lookup_by_name_with_flags_finish() to get the result.
// See g_resolver_lookup_by_name() for more details.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - hostname to look up the address of.
//    - flags: extra NameLookupFlags for the lookup.
//    - callback to call after resolution completes.
//
func (resolver *Resolver) LookupByNameWithFlagsAsync(ctx context.Context, hostname string, flags ResolverNameLookupFlags, callback AsyncReadyCallback) {
	var _arg0 *C.GResolver               // out
	var _arg3 *C.GCancellable            // out
	var _arg1 *C.gchar                   // out
	var _arg2 C.GResolverNameLookupFlags // out
	var _arg4 C.GAsyncReadyCallback      // out
	var _arg5 C.gpointer

	_arg0 = (*C.GResolver)(unsafe.Pointer(resolver.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(hostname)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GResolverNameLookupFlags(flags)
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_resolver_lookup_by_name_with_flags_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(hostname)
	runtime.KeepAlive(flags)
	runtime.KeepAlive(callback)
}

// LookupByNameWithFlagsFinish retrieves the result of a call to
// g_resolver_lookup_by_name_with_flags_async().
//
// If the DNS resolution failed, error (if non-NULL) will be set to a value from
// Error. If the operation was cancelled, error will be set to
// G_IO_ERROR_CANCELLED.
//
// The function takes the following parameters:
//
//    - result passed to your ReadyCallback.
//
func (resolver *Resolver) LookupByNameWithFlagsFinish(result AsyncResulter) ([]InetAddress, error) {
	var _arg0 *C.GResolver    // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GList        // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GResolver)(unsafe.Pointer(resolver.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_resolver_lookup_by_name_with_flags_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(result)

	var _list []InetAddress // out
	var _goerr error        // out

	_list = make([]InetAddress, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GInetAddress)(v)
		var dst InetAddress // out
		dst = *wrapInetAddress(externglib.AssumeOwnership(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _list, _goerr
}

// LookupRecords: synchronously performs a DNS record lookup for the given
// rrname and returns a list of records as #GVariant tuples. See RecordType for
// information on what the records contain for each record_type.
//
// If the DNS resolution fails, error (if non-NULL) will be set to a value from
// Error and NULL will be returned.
//
// If cancellable is non-NULL, it can be used to cancel the operation, in which
// case error (if non-NULL) will be set to G_IO_ERROR_CANCELLED.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - rrname: DNS name to look up the record for.
//    - recordType: type of DNS record to look up.
//
func (resolver *Resolver) LookupRecords(ctx context.Context, rrname string, recordType ResolverRecordType) ([]*glib.Variant, error) {
	var _arg0 *C.GResolver          // out
	var _arg3 *C.GCancellable       // out
	var _arg1 *C.gchar              // out
	var _arg2 C.GResolverRecordType // out
	var _cret *C.GList              // in
	var _cerr *C.GError             // in

	_arg0 = (*C.GResolver)(unsafe.Pointer(resolver.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(rrname)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GResolverRecordType(recordType)

	_cret = C.g_resolver_lookup_records(_arg0, _arg1, _arg2, _arg3, &_cerr)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(rrname)
	runtime.KeepAlive(recordType)

	var _list []*glib.Variant // out
	var _goerr error          // out

	_list = make([]*glib.Variant, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GVariant)(v)
		var dst *glib.Variant // out
		dst = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
		_list = append(_list, dst)
	})
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _list, _goerr
}

// LookupRecordsAsync begins asynchronously performing a DNS lookup for the
// given rrname, and eventually calls callback, which must call
// g_resolver_lookup_records_finish() to get the final result. See
// g_resolver_lookup_records() for more details.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - rrname: DNS name to look up the record for.
//    - recordType: type of DNS record to look up.
//    - callback to call after resolution completes.
//
func (resolver *Resolver) LookupRecordsAsync(ctx context.Context, rrname string, recordType ResolverRecordType, callback AsyncReadyCallback) {
	var _arg0 *C.GResolver          // out
	var _arg3 *C.GCancellable       // out
	var _arg1 *C.gchar              // out
	var _arg2 C.GResolverRecordType // out
	var _arg4 C.GAsyncReadyCallback // out
	var _arg5 C.gpointer

	_arg0 = (*C.GResolver)(unsafe.Pointer(resolver.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg3 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(rrname)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GResolverRecordType(recordType)
	if callback != nil {
		_arg4 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg5 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_resolver_lookup_records_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(rrname)
	runtime.KeepAlive(recordType)
	runtime.KeepAlive(callback)
}

// LookupRecordsFinish retrieves the result of a previous call to
// g_resolver_lookup_records_async(). Returns a non-empty list of records as
// #GVariant tuples. See RecordType for information on what the records contain.
//
// If the DNS resolution failed, error (if non-NULL) will be set to a value from
// Error. If the operation was cancelled, error will be set to
// G_IO_ERROR_CANCELLED.
//
// The function takes the following parameters:
//
//    - result passed to your ReadyCallback.
//
func (resolver *Resolver) LookupRecordsFinish(result AsyncResulter) ([]*glib.Variant, error) {
	var _arg0 *C.GResolver    // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GList        // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GResolver)(unsafe.Pointer(resolver.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_resolver_lookup_records_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(result)

	var _list []*glib.Variant // out
	var _goerr error          // out

	_list = make([]*glib.Variant, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GVariant)(v)
		var dst *glib.Variant // out
		dst = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_variant_unref((*C.GVariant)(intern.C))
			},
		)
		_list = append(_list, dst)
	})
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _list, _goerr
}

// LookupService: synchronously performs a DNS SRV lookup for the given service
// and protocol in the given domain and returns an array of Target. domain may
// be an ASCII-only or UTF-8 hostname. Note also that the service and protocol
// arguments do not include the leading underscore that appears in the actual
// DNS entry.
//
// On success, g_resolver_lookup_service() will return a non-empty #GList of
// Target, sorted in order of preference. (That is, you should attempt to
// connect to the first target first, then the second if the first fails, etc.)
//
// If the DNS resolution fails, error (if non-NULL) will be set to a value from
// Error and NULL will be returned.
//
// If cancellable is non-NULL, it can be used to cancel the operation, in which
// case error (if non-NULL) will be set to G_IO_ERROR_CANCELLED.
//
// If you are planning to connect to the service, it is usually easier to create
// a Service and use its Connectable interface.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - service type to look up (eg, "ldap").
//    - protocol: networking protocol to use for service (eg, "tcp").
//    - domain: DNS domain to look up the service in.
//
func (resolver *Resolver) LookupService(ctx context.Context, service, protocol, domain string) ([]*SrvTarget, error) {
	var _arg0 *C.GResolver    // out
	var _arg4 *C.GCancellable // out
	var _arg1 *C.gchar        // out
	var _arg2 *C.gchar        // out
	var _arg3 *C.gchar        // out
	var _cret *C.GList        // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GResolver)(unsafe.Pointer(resolver.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(service)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(protocol)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.g_resolver_lookup_service(_arg0, _arg1, _arg2, _arg3, _arg4, &_cerr)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(service)
	runtime.KeepAlive(protocol)
	runtime.KeepAlive(domain)

	var _list []*SrvTarget // out
	var _goerr error       // out

	_list = make([]*SrvTarget, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GSrvTarget)(v)
		var dst *SrvTarget // out
		dst = (*SrvTarget)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_srv_target_free((*C.GSrvTarget)(intern.C))
			},
		)
		_list = append(_list, dst)
	})
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _list, _goerr
}

// LookupServiceAsync begins asynchronously performing a DNS SRV lookup for the
// given service and protocol in the given domain, and eventually calls
// callback, which must call g_resolver_lookup_service_finish() to get the final
// result. See g_resolver_lookup_service() for more details.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - service type to look up (eg, "ldap").
//    - protocol: networking protocol to use for service (eg, "tcp").
//    - domain: DNS domain to look up the service in.
//    - callback to call after resolution completes.
//
func (resolver *Resolver) LookupServiceAsync(ctx context.Context, service, protocol, domain string, callback AsyncReadyCallback) {
	var _arg0 *C.GResolver          // out
	var _arg4 *C.GCancellable       // out
	var _arg1 *C.gchar              // out
	var _arg2 *C.gchar              // out
	var _arg3 *C.gchar              // out
	var _arg5 C.GAsyncReadyCallback // out
	var _arg6 C.gpointer

	_arg0 = (*C.GResolver)(unsafe.Pointer(resolver.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(service)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(protocol)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(domain)))
	defer C.free(unsafe.Pointer(_arg3))
	if callback != nil {
		_arg5 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg6 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_resolver_lookup_service_async(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(service)
	runtime.KeepAlive(protocol)
	runtime.KeepAlive(domain)
	runtime.KeepAlive(callback)
}

// LookupServiceFinish retrieves the result of a previous call to
// g_resolver_lookup_service_async().
//
// If the DNS resolution failed, error (if non-NULL) will be set to a value from
// Error. If the operation was cancelled, error will be set to
// G_IO_ERROR_CANCELLED.
//
// The function takes the following parameters:
//
//    - result passed to your ReadyCallback.
//
func (resolver *Resolver) LookupServiceFinish(result AsyncResulter) ([]*SrvTarget, error) {
	var _arg0 *C.GResolver    // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GList        // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GResolver)(unsafe.Pointer(resolver.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.g_resolver_lookup_service_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(resolver)
	runtime.KeepAlive(result)

	var _list []*SrvTarget // out
	var _goerr error       // out

	_list = make([]*SrvTarget, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GSrvTarget)(v)
		var dst *SrvTarget // out
		dst = (*SrvTarget)(gextras.NewStructNative(unsafe.Pointer(src)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(dst)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_srv_target_free((*C.GSrvTarget)(intern.C))
			},
		)
		_list = append(_list, dst)
	})
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _list, _goerr
}

// SetDefault sets resolver to be the application's default resolver (reffing
// resolver, and unreffing the previous default resolver, if any). Future calls
// to g_resolver_get_default() will return this resolver.
//
// This can be used if an application wants to perform any sort of DNS caching
// or "pinning"; it can implement its own #GResolver that calls the original
// default resolver for DNS operations, and implements its own cache policies on
// top of that, and then set itself as the default resolver for all later code
// to use.
func (resolver *Resolver) SetDefault() {
	var _arg0 *C.GResolver // out

	_arg0 = (*C.GResolver)(unsafe.Pointer(resolver.Native()))

	C.g_resolver_set_default(_arg0)
	runtime.KeepAlive(resolver)
}

// BaseResolver returns resolver.
func (resolver *Resolver) BaseResolver() *Resolver {
	return resolver
}

// ConnectReload: emitted when the resolver notices that the system resolver
// configuration has changed.
func (resolver *Resolver) ConnectReload(f func()) externglib.SignalHandle {
	return resolver.Connect("reload", f)
}

// ResolverGetDefault gets the default #GResolver. You should unref it when you
// are done with it. #GResolver may use its reference count as a hint about how
// many threads it should allocate for concurrent DNS resolutions.
func ResolverGetDefault() Resolverer {
	var _cret *C.GResolver // in

	_cret = C.g_resolver_get_default()

	var _resolver Resolverer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gio.Resolverer is nil")
		}

		object := externglib.AssumeOwnership(objptr)
		rv, ok := (externglib.CastObject(object)).(Resolverer)
		if !ok {
			panic("object of type " + object.TypeFromInstance().String() + " is not gio.Resolverer")
		}
		_resolver = rv
	}

	return _resolver
}
