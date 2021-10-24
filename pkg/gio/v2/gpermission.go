// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// void _gotk4_gio2_AsyncReadyCallback(GObject*, GAsyncResult*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_permission_get_type()), F: marshalPermissioner},
	})
}

// PermissionOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type PermissionOverrider interface {
	// Acquire attempts to acquire the permission represented by permission.
	//
	// The precise method by which this happens depends on the permission and
	// the underlying authentication mechanism. A simple example is that a
	// dialog may appear asking the user to enter their password.
	//
	// You should check with g_permission_get_can_acquire() before calling this
	// function.
	//
	// If the permission is acquired then TRUE is returned. Otherwise, FALSE is
	// returned and error is set appropriately.
	//
	// This call is blocking, likely for a very long time (in the case that user
	// interaction is required). See g_permission_acquire_async() for the
	// non-blocking version.
	Acquire(ctx context.Context) error
	// AcquireAsync attempts to acquire the permission represented by
	// permission.
	//
	// This is the first half of the asynchronous version of
	// g_permission_acquire().
	AcquireAsync(ctx context.Context, callback AsyncReadyCallback)
	// AcquireFinish collects the result of attempting to acquire the permission
	// represented by permission.
	//
	// This is the second half of the asynchronous version of
	// g_permission_acquire().
	AcquireFinish(result AsyncResulter) error
	// Release attempts to release the permission represented by permission.
	//
	// The precise method by which this happens depends on the permission and
	// the underlying authentication mechanism. In most cases the permission
	// will be dropped immediately without further action.
	//
	// You should check with g_permission_get_can_release() before calling this
	// function.
	//
	// If the permission is released then TRUE is returned. Otherwise, FALSE is
	// returned and error is set appropriately.
	//
	// This call is blocking, likely for a very long time (in the case that user
	// interaction is required). See g_permission_release_async() for the
	// non-blocking version.
	Release(ctx context.Context) error
	// ReleaseAsync attempts to release the permission represented by
	// permission.
	//
	// This is the first half of the asynchronous version of
	// g_permission_release().
	ReleaseAsync(ctx context.Context, callback AsyncReadyCallback)
	// ReleaseFinish collects the result of attempting to release the permission
	// represented by permission.
	//
	// This is the second half of the asynchronous version of
	// g_permission_release().
	ReleaseFinish(result AsyncResulter) error
}

// Permission represents the status of the caller's permission to perform a
// certain action.
//
// You can query if the action is currently allowed and if it is possible to
// acquire the permission so that the action will be allowed in the future.
//
// There is also an API to actually acquire the permission and one to release
// it.
//
// As an example, a #GPermission might represent the ability for the user to
// write to a #GSettings object. This #GPermission object could then be used to
// decide if it is appropriate to show a "Click here to unlock" button in a
// dialog and to provide the mechanism to invoke when that button is clicked.
type Permission struct {
	*externglib.Object
}

// Permissioner describes types inherited from class Permission.
// To get the original type, the caller must assert this to an interface or
// another type.
type Permissioner interface {
	externglib.Objector

	// BasePermission returns the underlying base class.
	BasePermission() *Permission
}

var _ Permissioner = (*Permission)(nil)

func wrapPermission(obj *externglib.Object) *Permission {
	return &Permission{
		Object: obj,
	}
}

func marshalPermissioner(p uintptr) (interface{}, error) {
	return wrapPermission(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Acquire attempts to acquire the permission represented by permission.
//
// The precise method by which this happens depends on the permission and the
// underlying authentication mechanism. A simple example is that a dialog may
// appear asking the user to enter their password.
//
// You should check with g_permission_get_can_acquire() before calling this
// function.
//
// If the permission is acquired then TRUE is returned. Otherwise, FALSE is
// returned and error is set appropriately.
//
// This call is blocking, likely for a very long time (in the case that user
// interaction is required). See g_permission_acquire_async() for the
// non-blocking version.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//
func (permission *Permission) Acquire(ctx context.Context) error {
	var _arg0 *C.GPermission  // out
	var _arg1 *C.GCancellable // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(permission.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	C.g_permission_acquire(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(permission)
	runtime.KeepAlive(ctx)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// AcquireAsync attempts to acquire the permission represented by permission.
//
// This is the first half of the asynchronous version of g_permission_acquire().
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - callback to call when done.
//
func (permission *Permission) AcquireAsync(ctx context.Context, callback AsyncReadyCallback) {
	var _arg0 *C.GPermission        // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.GPermission)(unsafe.Pointer(permission.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if callback != nil {
		_arg2 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg3 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_permission_acquire_async(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(permission)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(callback)
}

// AcquireFinish collects the result of attempting to acquire the permission
// represented by permission.
//
// This is the second half of the asynchronous version of
// g_permission_acquire().
//
// The function takes the following parameters:
//
//    - result given to the ReadyCallback.
//
func (permission *Permission) AcquireFinish(result AsyncResulter) error {
	var _arg0 *C.GPermission  // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(permission.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_permission_acquire_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(permission)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Allowed gets the value of the 'allowed' property. This property is TRUE if
// the caller currently has permission to perform the action that permission
// represents the permission to perform.
func (permission *Permission) Allowed() bool {
	var _arg0 *C.GPermission // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(permission.Native()))

	_cret = C.g_permission_get_allowed(_arg0)
	runtime.KeepAlive(permission)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanAcquire gets the value of the 'can-acquire' property. This property is
// TRUE if it is generally possible to acquire the permission by calling
// g_permission_acquire().
func (permission *Permission) CanAcquire() bool {
	var _arg0 *C.GPermission // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(permission.Native()))

	_cret = C.g_permission_get_can_acquire(_arg0)
	runtime.KeepAlive(permission)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CanRelease gets the value of the 'can-release' property. This property is
// TRUE if it is generally possible to release the permission by calling
// g_permission_release().
func (permission *Permission) CanRelease() bool {
	var _arg0 *C.GPermission // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(permission.Native()))

	_cret = C.g_permission_get_can_release(_arg0)
	runtime.KeepAlive(permission)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ImplUpdate: this function is called by the #GPermission implementation to
// update the properties of the permission. You should never call this function
// except from a #GPermission implementation.
//
// GObject notify signals are generated, as appropriate.
//
// The function takes the following parameters:
//
//    - allowed: new value for the 'allowed' property.
//    - canAcquire: new value for the 'can-acquire' property.
//    - canRelease: new value for the 'can-release' property.
//
func (permission *Permission) ImplUpdate(allowed, canAcquire, canRelease bool) {
	var _arg0 *C.GPermission // out
	var _arg1 C.gboolean     // out
	var _arg2 C.gboolean     // out
	var _arg3 C.gboolean     // out

	_arg0 = (*C.GPermission)(unsafe.Pointer(permission.Native()))
	if allowed {
		_arg1 = C.TRUE
	}
	if canAcquire {
		_arg2 = C.TRUE
	}
	if canRelease {
		_arg3 = C.TRUE
	}

	C.g_permission_impl_update(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(permission)
	runtime.KeepAlive(allowed)
	runtime.KeepAlive(canAcquire)
	runtime.KeepAlive(canRelease)
}

// Release attempts to release the permission represented by permission.
//
// The precise method by which this happens depends on the permission and the
// underlying authentication mechanism. In most cases the permission will be
// dropped immediately without further action.
//
// You should check with g_permission_get_can_release() before calling this
// function.
//
// If the permission is released then TRUE is returned. Otherwise, FALSE is
// returned and error is set appropriately.
//
// This call is blocking, likely for a very long time (in the case that user
// interaction is required). See g_permission_release_async() for the
// non-blocking version.
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//
func (permission *Permission) Release(ctx context.Context) error {
	var _arg0 *C.GPermission  // out
	var _arg1 *C.GCancellable // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(permission.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	C.g_permission_release(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(permission)
	runtime.KeepAlive(ctx)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// ReleaseAsync attempts to release the permission represented by permission.
//
// This is the first half of the asynchronous version of g_permission_release().
//
// The function takes the following parameters:
//
//    - ctx or NULL.
//    - callback to call when done.
//
func (permission *Permission) ReleaseAsync(ctx context.Context, callback AsyncReadyCallback) {
	var _arg0 *C.GPermission        // out
	var _arg1 *C.GCancellable       // out
	var _arg2 C.GAsyncReadyCallback // out
	var _arg3 C.gpointer

	_arg0 = (*C.GPermission)(unsafe.Pointer(permission.Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	if callback != nil {
		_arg2 = (*[0]byte)(C._gotk4_gio2_AsyncReadyCallback)
		_arg3 = C.gpointer(gbox.AssignOnce(callback))
	}

	C.g_permission_release_async(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(permission)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(callback)
}

// ReleaseFinish collects the result of attempting to release the permission
// represented by permission.
//
// This is the second half of the asynchronous version of
// g_permission_release().
//
// The function takes the following parameters:
//
//    - result given to the ReadyCallback.
//
func (permission *Permission) ReleaseFinish(result AsyncResulter) error {
	var _arg0 *C.GPermission  // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GPermission)(unsafe.Pointer(permission.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.g_permission_release_finish(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(permission)
	runtime.KeepAlive(result)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// BasePermission returns permission.
func (permission *Permission) BasePermission() *Permission {
	return permission
}
