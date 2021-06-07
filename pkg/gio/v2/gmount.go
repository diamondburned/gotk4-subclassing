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
		{T: externglib.Type(C.g_mount_get_type()), F: marshalMount},
	})
}

// MountOverrider contains methods that are overridable. This
// interface is a subset of the interface Mount.
type MountOverrider interface {
	// CanEject checks if @mount can be ejected.
	CanEject(m Mount) bool
	// CanUnmount checks if @mount can be unmounted.
	CanUnmount(m Mount) bool

	Changed(m Mount)
	// Eject ejects a mount. This is an asynchronous operation, and is finished
	// by calling g_mount_eject_finish() with the @mount and Result data
	// returned in the @callback.
	Eject(m Mount)
	// EjectFinish finishes ejecting a mount. If any errors occurred during the
	// operation, @error will be set to contain the errors and false will be
	// returned.
	EjectFinish(m Mount, result AsyncResult) error
	// EjectWithOperation ejects a mount. This is an asynchronous operation, and
	// is finished by calling g_mount_eject_with_operation_finish() with the
	// @mount and Result data returned in the @callback.
	EjectWithOperation(m Mount)
	// EjectWithOperationFinish finishes ejecting a mount. If any errors
	// occurred during the operation, @error will be set to contain the errors
	// and false will be returned.
	EjectWithOperationFinish(m Mount, result AsyncResult) error
	// DefaultLocation gets the default location of @mount. The default location
	// of the given @mount is a path that reflects the main entry point for the
	// user (e.g. the home directory, or the root of the volume).
	DefaultLocation(m Mount)
	// Drive gets the drive for the @mount.
	//
	// This is a convenience method for getting the #GVolume and then using that
	// object to get the #GDrive.
	Drive(m Mount)
	// Icon gets the icon for @mount.
	Icon(m Mount)
	// Name gets the name of @mount.
	Name(m Mount)
	// Root gets the root directory on @mount.
	Root(m Mount)
	// SortKey gets the sort key for @mount, if any.
	SortKey(m Mount)
	// SymbolicIcon gets the symbolic icon for @mount.
	SymbolicIcon(m Mount)
	// UUID gets the UUID for the @mount. The reference is typically based on
	// the file system UUID for the mount in question and should be considered
	// an opaque string. Returns nil if there is no UUID available.
	UUID(m Mount)
	// Volume gets the volume for the @mount.
	Volume(m Mount)
	// GuessContentType tries to guess the type of content stored on @mount.
	// Returns one or more textual identifiers of well-known content types
	// (typically prefixed with "x-content/"), e.g. x-content/image-dcf for
	// camera memory cards. See the shared-mime-info
	// (http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
	// specification for more on x-content types.
	//
	// This is an asynchronous operation (see g_mount_guess_content_type_sync()
	// for the synchronous version), and is finished by calling
	// g_mount_guess_content_type_finish() with the @mount and Result data
	// returned in the @callback.
	GuessContentType(m Mount)
	// GuessContentTypeFinish finishes guessing content types of @mount. If any
	// errors occurred during the operation, @error will be set to contain the
	// errors and false will be returned. In particular, you may get an
	// G_IO_ERROR_NOT_SUPPORTED if the mount does not support content guessing.
	GuessContentTypeFinish(m Mount, result AsyncResult) error
	// GuessContentTypeSync tries to guess the type of content stored on @mount.
	// Returns one or more textual identifiers of well-known content types
	// (typically prefixed with "x-content/"), e.g. x-content/image-dcf for
	// camera memory cards. See the shared-mime-info
	// (http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
	// specification for more on x-content types.
	//
	// This is a synchronous operation and as such may block doing IO; see
	// g_mount_guess_content_type() for the asynchronous version.
	GuessContentTypeSync(m Mount, forceRescan bool, cancellable Cancellable) error

	PreUnmount(m Mount)
	// Remount remounts a mount. This is an asynchronous operation, and is
	// finished by calling g_mount_remount_finish() with the @mount and Results
	// data returned in the @callback.
	//
	// Remounting is useful when some setting affecting the operation of the
	// volume has been changed, as these may need a remount to take affect.
	// While this is semantically equivalent with unmounting and then remounting
	// not all backends might need to actually be unmounted.
	Remount(m Mount)
	// RemountFinish finishes remounting a mount. If any errors occurred during
	// the operation, @error will be set to contain the errors and false will be
	// returned.
	RemountFinish(m Mount, result AsyncResult) error
	// Unmount unmounts a mount. This is an asynchronous operation, and is
	// finished by calling g_mount_unmount_finish() with the @mount and Result
	// data returned in the @callback.
	Unmount(m Mount)
	// UnmountFinish finishes unmounting a mount. If any errors occurred during
	// the operation, @error will be set to contain the errors and false will be
	// returned.
	UnmountFinish(m Mount, result AsyncResult) error
	// UnmountWithOperation unmounts a mount. This is an asynchronous operation,
	// and is finished by calling g_mount_unmount_with_operation_finish() with
	// the @mount and Result data returned in the @callback.
	UnmountWithOperation(m Mount)
	// UnmountWithOperationFinish finishes unmounting a mount. If any errors
	// occurred during the operation, @error will be set to contain the errors
	// and false will be returned.
	UnmountWithOperationFinish(m Mount, result AsyncResult) error

	Unmounted(m Mount)
}

// Mount: the #GMount interface represents user-visible mounts. Note, when
// porting from GnomeVFS, #GMount is the moral equivalent of VFSVolume.
//
// #GMount is a "mounted" filesystem that you can access. Mounted is in quotes
// because it's not the same as a unix mount, it might be a gvfs mount, but you
// can still access the files on it if you use GIO. Might or might not be
// related to a volume object.
//
// Unmounting a #GMount instance is an asynchronous operation. For more
// information about asynchronous operations, see Result and #GTask. To unmount
// a #GMount instance, first call g_mount_unmount_with_operation() with (at
// least) the #GMount instance and a ReadyCallback. The callback will be fired
// when the operation has resolved (either with success or failure), and a
// Result structure will be passed to the callback. That callback should then
// call g_mount_unmount_with_operation_finish() with the #GMount and the Result
// data to see if the operation was completed successfully. If an @error is
// present when g_mount_unmount_with_operation_finish() is called, then it will
// be filled with any error information.
type Mount interface {
	gextras.Objector
	MountOverrider

	// IsShadowed determines if @mount is shadowed. Applications or libraries
	// should avoid displaying @mount in the user interface if it is shadowed.
	//
	// A mount is said to be shadowed if there exists one or more user visible
	// objects (currently #GMount objects) with a root that is inside the root
	// of @mount.
	//
	// One application of shadow mounts is when exposing a single file system
	// that is used to address several logical volumes. In this situation, a
	// Monitor implementation would create two #GVolume objects (for example,
	// one for the camera functionality of the device and one for a SD card
	// reader on the device) with activation URIs
	// `gphoto2://[usb:001,002]/store1/` and `gphoto2://[usb:001,002]/store2/`.
	// When the underlying mount (with root `gphoto2://[usb:001,002]/`) is
	// mounted, said Monitor implementation would create two #GMount objects
	// (each with their root matching the corresponding volume activation root)
	// that would shadow the original mount.
	//
	// The proxy monitor in GVfs 2.26 and later, automatically creates and
	// manage shadow mounts (and shadows the underlying mount) if the activation
	// root on a #GVolume is set.
	IsShadowed(m Mount) bool
	// Shadow increments the shadow count on @mount. Usually used by Monitor
	// implementations when creating a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	Shadow(m Mount)
	// Unshadow decrements the shadow count on @mount. Usually used by Monitor
	// implementations when destroying a shadow mount for @mount, see
	// g_mount_is_shadowed() for more information. The caller will need to emit
	// the #GMount::changed signal on @mount manually.
	Unshadow(m Mount)
}

// mount implements the Mount interface.
type mount struct {
	gextras.Objector
}

var _ Mount = (*mount)(nil)

// WrapMount wraps a GObject to a type that implements interface
// Mount. It is primarily used internally.
func WrapMount(obj *externglib.Object) Mount {
	return Mount{
		Objector: obj,
	}
}

func marshalMount(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMount(obj), nil
}

// CanEject checks if @mount can be ejected.
func (m mount) CanEject(m Mount) bool {
	var arg0 *C.GMount

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_mount_can_eject(arg0)

	if cret {
		ok = true
	}

	return ok
}

// CanUnmount checks if @mount can be unmounted.
func (m mount) CanUnmount(m Mount) bool {
	var arg0 *C.GMount

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_mount_can_unmount(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Eject ejects a mount. This is an asynchronous operation, and is finished
// by calling g_mount_eject_finish() with the @mount and Result data
// returned in the @callback.
func (m mount) Eject(m Mount) {
	var arg0 *C.GMount

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_eject(arg0, arg1, arg2, arg3, arg4)
}

// EjectFinish finishes ejecting a mount. If any errors occurred during the
// operation, @error will be set to contain the errors and false will be
// returned.
func (m mount) EjectFinish(m Mount, result AsyncResult) error {
	var arg0 *C.GMount
	var arg1 *C.GAsyncResult

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var errout *C.GError
	var err error

	C.g_mount_eject_finish(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// EjectWithOperation ejects a mount. This is an asynchronous operation, and
// is finished by calling g_mount_eject_with_operation_finish() with the
// @mount and Result data returned in the @callback.
func (m mount) EjectWithOperation(m Mount) {
	var arg0 *C.GMount

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_eject_with_operation(arg0, arg1, arg2, arg3, arg4, arg5)
}

// EjectWithOperationFinish finishes ejecting a mount. If any errors
// occurred during the operation, @error will be set to contain the errors
// and false will be returned.
func (m mount) EjectWithOperationFinish(m Mount, result AsyncResult) error {
	var arg0 *C.GMount
	var arg1 *C.GAsyncResult

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var errout *C.GError
	var err error

	C.g_mount_eject_with_operation_finish(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// DefaultLocation gets the default location of @mount. The default location
// of the given @mount is a path that reflects the main entry point for the
// user (e.g. the home directory, or the root of the volume).
func (m mount) DefaultLocation(m Mount) {
	var arg0 *C.GMount

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_get_default_location(arg0)
}

// Drive gets the drive for the @mount.
//
// This is a convenience method for getting the #GVolume and then using that
// object to get the #GDrive.
func (m mount) Drive(m Mount) {
	var arg0 *C.GMount

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_get_drive(arg0)
}

// Icon gets the icon for @mount.
func (m mount) Icon(m Mount) {
	var arg0 *C.GMount

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_get_icon(arg0)
}

// Name gets the name of @mount.
func (m mount) Name(m Mount) {
	var arg0 *C.GMount

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_get_name(arg0)
}

// Root gets the root directory on @mount.
func (m mount) Root(m Mount) {
	var arg0 *C.GMount

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_get_root(arg0)
}

// SortKey gets the sort key for @mount, if any.
func (m mount) SortKey(m Mount) {
	var arg0 *C.GMount

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_get_sort_key(arg0)
}

// SymbolicIcon gets the symbolic icon for @mount.
func (m mount) SymbolicIcon(m Mount) {
	var arg0 *C.GMount

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_get_symbolic_icon(arg0)
}

// UUID gets the UUID for the @mount. The reference is typically based on
// the file system UUID for the mount in question and should be considered
// an opaque string. Returns nil if there is no UUID available.
func (m mount) UUID(m Mount) {
	var arg0 *C.GMount

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_get_uuid(arg0)
}

// Volume gets the volume for the @mount.
func (m mount) Volume(m Mount) {
	var arg0 *C.GMount

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_get_volume(arg0)
}

// GuessContentType tries to guess the type of content stored on @mount.
// Returns one or more textual identifiers of well-known content types
// (typically prefixed with "x-content/"), e.g. x-content/image-dcf for
// camera memory cards. See the shared-mime-info
// (http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
// specification for more on x-content types.
//
// This is an asynchronous operation (see g_mount_guess_content_type_sync()
// for the synchronous version), and is finished by calling
// g_mount_guess_content_type_finish() with the @mount and Result data
// returned in the @callback.
func (m mount) GuessContentType(m Mount) {
	var arg0 *C.GMount

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_guess_content_type(arg0, arg1, arg2, arg3, arg4)
}

// GuessContentTypeFinish finishes guessing content types of @mount. If any
// errors occurred during the operation, @error will be set to contain the
// errors and false will be returned. In particular, you may get an
// G_IO_ERROR_NOT_SUPPORTED if the mount does not support content guessing.
func (m mount) GuessContentTypeFinish(m Mount, result AsyncResult) error {
	var arg0 *C.GMount
	var arg1 *C.GAsyncResult

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var errout *C.GError
	var err error

	C.g_mount_guess_content_type_finish(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// GuessContentTypeSync tries to guess the type of content stored on @mount.
// Returns one or more textual identifiers of well-known content types
// (typically prefixed with "x-content/"), e.g. x-content/image-dcf for
// camera memory cards. See the shared-mime-info
// (http://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec)
// specification for more on x-content types.
//
// This is a synchronous operation and as such may block doing IO; see
// g_mount_guess_content_type() for the asynchronous version.
func (m mount) GuessContentTypeSync(m Mount, forceRescan bool, cancellable Cancellable) error {
	var arg0 *C.GMount
	var arg1 C.gboolean
	var arg2 *C.GCancellable

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	if forceRescan {
		arg1 = C.gboolean(1)
	}
	arg2 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))

	var errout *C.GError
	var err error

	C.g_mount_guess_content_type_sync(arg0, arg1, arg2, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// IsShadowed determines if @mount is shadowed. Applications or libraries
// should avoid displaying @mount in the user interface if it is shadowed.
//
// A mount is said to be shadowed if there exists one or more user visible
// objects (currently #GMount objects) with a root that is inside the root
// of @mount.
//
// One application of shadow mounts is when exposing a single file system
// that is used to address several logical volumes. In this situation, a
// Monitor implementation would create two #GVolume objects (for example,
// one for the camera functionality of the device and one for a SD card
// reader on the device) with activation URIs
// `gphoto2://[usb:001,002]/store1/` and `gphoto2://[usb:001,002]/store2/`.
// When the underlying mount (with root `gphoto2://[usb:001,002]/`) is
// mounted, said Monitor implementation would create two #GMount objects
// (each with their root matching the corresponding volume activation root)
// that would shadow the original mount.
//
// The proxy monitor in GVfs 2.26 and later, automatically creates and
// manage shadow mounts (and shadows the underlying mount) if the activation
// root on a #GVolume is set.
func (m mount) IsShadowed(m Mount) bool {
	var arg0 *C.GMount

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_mount_is_shadowed(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Remount remounts a mount. This is an asynchronous operation, and is
// finished by calling g_mount_remount_finish() with the @mount and Results
// data returned in the @callback.
//
// Remounting is useful when some setting affecting the operation of the
// volume has been changed, as these may need a remount to take affect.
// While this is semantically equivalent with unmounting and then remounting
// not all backends might need to actually be unmounted.
func (m mount) Remount(m Mount) {
	var arg0 *C.GMount

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_remount(arg0, arg1, arg2, arg3, arg4, arg5)
}

// RemountFinish finishes remounting a mount. If any errors occurred during
// the operation, @error will be set to contain the errors and false will be
// returned.
func (m mount) RemountFinish(m Mount, result AsyncResult) error {
	var arg0 *C.GMount
	var arg1 *C.GAsyncResult

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var errout *C.GError
	var err error

	C.g_mount_remount_finish(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// Shadow increments the shadow count on @mount. Usually used by Monitor
// implementations when creating a shadow mount for @mount, see
// g_mount_is_shadowed() for more information. The caller will need to emit
// the #GMount::changed signal on @mount manually.
func (m mount) Shadow(m Mount) {
	var arg0 *C.GMount

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_shadow(arg0)
}

// Unmount unmounts a mount. This is an asynchronous operation, and is
// finished by calling g_mount_unmount_finish() with the @mount and Result
// data returned in the @callback.
func (m mount) Unmount(m Mount) {
	var arg0 *C.GMount

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_unmount(arg0, arg1, arg2, arg3, arg4)
}

// UnmountFinish finishes unmounting a mount. If any errors occurred during
// the operation, @error will be set to contain the errors and false will be
// returned.
func (m mount) UnmountFinish(m Mount, result AsyncResult) error {
	var arg0 *C.GMount
	var arg1 *C.GAsyncResult

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var errout *C.GError
	var err error

	C.g_mount_unmount_finish(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// UnmountWithOperation unmounts a mount. This is an asynchronous operation,
// and is finished by calling g_mount_unmount_with_operation_finish() with
// the @mount and Result data returned in the @callback.
func (m mount) UnmountWithOperation(m Mount) {
	var arg0 *C.GMount

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_unmount_with_operation(arg0, arg1, arg2, arg3, arg4, arg5)
}

// UnmountWithOperationFinish finishes unmounting a mount. If any errors
// occurred during the operation, @error will be set to contain the errors
// and false will be returned.
func (m mount) UnmountWithOperationFinish(m Mount, result AsyncResult) error {
	var arg0 *C.GMount
	var arg1 *C.GAsyncResult

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var errout *C.GError
	var err error

	C.g_mount_unmount_with_operation_finish(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// Unshadow decrements the shadow count on @mount. Usually used by Monitor
// implementations when destroying a shadow mount for @mount, see
// g_mount_is_shadowed() for more information. The caller will need to emit
// the #GMount::changed signal on @mount manually.
func (m mount) Unshadow(m Mount) {
	var arg0 *C.GMount

	arg0 = (*C.GMount)(unsafe.Pointer(m.Native()))

	C.g_mount_unshadow(arg0)
}