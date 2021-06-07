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
		{T: externglib.Type(C.g_settings_backend_get_type()), F: marshalSettingsBackend},
	})
}

// NewKeyfileSettingsBackend creates a keyfile-backed Backend.
//
// The filename of the keyfile to use is given by @filename.
//
// All settings read to or written from the backend must fall under the path
// given in @root_path (which must start and end with a slash and not contain
// two consecutive slashes). @root_path may be "/".
//
// If @root_group is non-nil then it specifies the name of the keyfile group
// used for keys that are written directly below @root_path. For example, if
// @root_path is "/apps/example/" and @root_group is "toplevel", then settings
// the key "/apps/example/enabled" to a value of true will cause the following
// to appear in the keyfile:
//
//    [toplevel]
//    enabled=true
//
// If @root_group is nil then it is not permitted to store keys directly below
// the @root_path.
//
// For keys not stored directly below @root_path (ie: in a sub-path), the name
// of the subpath (with the final slash stripped) is used as the name of the
// keyfile group. To continue the example, if
// "/apps/example/profiles/default/font-size" were set to 12 then the following
// would appear in the keyfile:
//
//    [profiles/default]
//    font-size=12
//
// The backend will refuse writes (and return writability as being false) for
// keys outside of @root_path and, in the event that @root_group is nil, also
// for keys directly under @root_path. Writes will also be refused if the
// backend detects that it has the inability to rewrite the keyfile (ie: the
// containing directory is not writable).
//
// There is no checking done for your key namespace clashing with the syntax of
// the key file format. For example, if you have '[' or ']' characters in your
// path names or '=' in your key names you may be in trouble.
//
// The backend reads default values from a keyfile called `defaults` in the
// directory specified by the SettingsBackend:defaults-dir property, and a list
// of locked keys from a text file with the name `locks` in the same location.
func NewKeyfileSettingsBackend(filename string, rootPath string, rootGroup string) {
	var arg1 *C.gchar
	var arg2 *C.gchar
	var arg3 *C.gchar

	arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.gchar)(C.CString(rootPath))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.gchar)(C.CString(rootGroup))
	defer C.free(unsafe.Pointer(arg3))

	C.g_keyfile_settings_backend_new(arg1, arg2, arg3)
}

// NewMemorySettingsBackend creates a memory-backed Backend.
//
// This backend allows changes to settings, but does not write them to any
// backing storage, so the next time you run your application, the memory
// backend will start out with the default values again.
func NewMemorySettingsBackend() {
	C.g_memory_settings_backend_new()
}

// NewNullSettingsBackend creates a readonly Backend.
//
// This backend does not allow changes to settings, so all settings will always
// have their default values.
func NewNullSettingsBackend() {
	C.g_null_settings_backend_new()
}

// SettingsBackend: the Backend interface defines a generic interface for
// non-strictly-typed data that is stored in a hierarchy. To implement an
// alternative storage backend for #GSettings, you need to implement the Backend
// interface and then make it implement the extension point
// SETTINGS_BACKEND_EXTENSION_POINT_NAME.
//
// The interface defines methods for reading and writing values, a method for
// determining if writing of certain values will fail (lockdown) and a change
// notification mechanism.
//
// The semantics of the interface are very precisely defined and implementations
// must carefully adhere to the expectations of callers that are documented on
// each of the interface methods.
//
// Some of the Backend functions accept or return a #GTree. These trees always
// have strings as keys and #GVariant as values.
// g_settings_backend_create_tree() is a convenience function to create suitable
// trees.
//
// The Backend API is exported to allow third-party implementations, but does
// not carry the same stability guarantees as the public GIO API. For this
// reason, you have to define the C preprocessor symbol
// G_SETTINGS_ENABLE_BACKEND before including `gio/gsettingsbackend.h`.
type SettingsBackend interface {
	gextras.Objector

	// Changed signals that a single key has possibly changed. Backend
	// implementations should call this if a key has possibly changed its value.
	//
	// @key must be a valid key (ie starting with a slash, not containing '//',
	// and not ending with a slash).
	//
	// The implementation must call this function during any call to
	// g_settings_backend_write(), before the call returns (except in the case
	// that no keys are actually changed and it cares to detect this fact). It
	// may not rely on the existence of a mainloop for dispatching the signal
	// later.
	//
	// The implementation may call this function at any other time it likes in
	// response to other events (such as changes occurring outside of the
	// program). These calls may originate from a mainloop or may originate in
	// response to any other action (including from calls to
	// g_settings_backend_write()).
	//
	// In the case that this call is in response to a call to
	// g_settings_backend_write() then @origin_tag must be set to the same value
	// that was passed to that call.
	Changed(b SettingsBackend, key string, originTag interface{})
	// ChangedTree: this call is a convenience wrapper. It gets the list of
	// changes from @tree, computes the longest common prefix and calls
	// g_settings_backend_changed().
	ChangedTree(b SettingsBackend, tree *glib.Tree, originTag interface{})
	// KeysChanged signals that a list of keys have possibly changed. Backend
	// implementations should call this if keys have possibly changed their
	// values.
	//
	// @path must be a valid path (ie starting and ending with a slash and not
	// containing '//'). Each string in @items must form a valid key name when
	// @path is prefixed to it (ie: each item must not start or end with '/' and
	// must not contain '//').
	//
	// The meaning of this signal is that any of the key names resulting from
	// the contatenation of @path with each item in @items may have changed.
	//
	// The same rules for when notifications must occur apply as per
	// g_settings_backend_changed(). These two calls can be used interchangeably
	// if exactly one item has changed (although in that case
	// g_settings_backend_changed() is definitely preferred).
	//
	// For efficiency reasons, the implementation should strive for @path to be
	// as long as possible (ie: the longest common prefix of all of the keys
	// that were changed) but this is not strictly required.
	KeysChanged(b SettingsBackend, path string, items []string, originTag interface{})
	// PathChanged signals that all keys below a given path may have possibly
	// changed. Backend implementations should call this if an entire path of
	// keys have possibly changed their values.
	//
	// @path must be a valid path (ie starting and ending with a slash and not
	// containing '//').
	//
	// The meaning of this signal is that any of the key which has a name
	// starting with @path may have changed.
	//
	// The same rules for when notifications must occur apply as per
	// g_settings_backend_changed(). This call might be an appropriate reasponse
	// to a 'reset' call but implementations are also free to explicitly list
	// the keys that were affected by that call if they can easily do so.
	//
	// For efficiency reasons, the implementation should strive for @path to be
	// as long as possible (ie: the longest common prefix of all of the keys
	// that were changed) but this is not strictly required. As an example, if
	// this function is called with the path of "/" then every single key in the
	// application will be notified of a possible change.
	PathChanged(b SettingsBackend, path string, originTag interface{})
	// PathWritableChanged signals that the writability of all keys below a
	// given path may have changed.
	//
	// Since GSettings performs no locking operations for itself, this call will
	// always be made in response to external events.
	PathWritableChanged(b SettingsBackend, path string)
	// WritableChanged signals that the writability of a single key has possibly
	// changed.
	//
	// Since GSettings performs no locking operations for itself, this call will
	// always be made in response to external events.
	WritableChanged(b SettingsBackend, key string)
}

// settingsBackend implements the SettingsBackend interface.
type settingsBackend struct {
	gextras.Objector
}

var _ SettingsBackend = (*settingsBackend)(nil)

// WrapSettingsBackend wraps a GObject to the right type. It is
// primarily used internally.
func WrapSettingsBackend(obj *externglib.Object) SettingsBackend {
	return SettingsBackend{
		Objector: obj,
	}
}

func marshalSettingsBackend(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSettingsBackend(obj), nil
}

// Changed signals that a single key has possibly changed. Backend
// implementations should call this if a key has possibly changed its value.
//
// @key must be a valid key (ie starting with a slash, not containing '//',
// and not ending with a slash).
//
// The implementation must call this function during any call to
// g_settings_backend_write(), before the call returns (except in the case
// that no keys are actually changed and it cares to detect this fact). It
// may not rely on the existence of a mainloop for dispatching the signal
// later.
//
// The implementation may call this function at any other time it likes in
// response to other events (such as changes occurring outside of the
// program). These calls may originate from a mainloop or may originate in
// response to any other action (including from calls to
// g_settings_backend_write()).
//
// In the case that this call is in response to a call to
// g_settings_backend_write() then @origin_tag must be set to the same value
// that was passed to that call.
func (b settingsBackend) Changed(b SettingsBackend, key string, originTag interface{}) {
	var arg0 *C.GSettingsBackend
	var arg1 *C.gchar
	var arg2 C.gpointer

	arg0 = (*C.GSettingsBackend)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gpointer(originTag)

	C.g_settings_backend_changed(arg0, arg1, arg2)
}

// ChangedTree: this call is a convenience wrapper. It gets the list of
// changes from @tree, computes the longest common prefix and calls
// g_settings_backend_changed().
func (b settingsBackend) ChangedTree(b SettingsBackend, tree *glib.Tree, originTag interface{}) {
	var arg0 *C.GSettingsBackend
	var arg1 *C.GTree
	var arg2 C.gpointer

	arg0 = (*C.GSettingsBackend)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GTree)(unsafe.Pointer(tree.Native()))
	arg2 = C.gpointer(originTag)

	C.g_settings_backend_changed_tree(arg0, arg1, arg2)
}

// KeysChanged signals that a list of keys have possibly changed. Backend
// implementations should call this if keys have possibly changed their
// values.
//
// @path must be a valid path (ie starting and ending with a slash and not
// containing '//'). Each string in @items must form a valid key name when
// @path is prefixed to it (ie: each item must not start or end with '/' and
// must not contain '//').
//
// The meaning of this signal is that any of the key names resulting from
// the contatenation of @path with each item in @items may have changed.
//
// The same rules for when notifications must occur apply as per
// g_settings_backend_changed(). These two calls can be used interchangeably
// if exactly one item has changed (although in that case
// g_settings_backend_changed() is definitely preferred).
//
// For efficiency reasons, the implementation should strive for @path to be
// as long as possible (ie: the longest common prefix of all of the keys
// that were changed) but this is not strictly required.
func (b settingsBackend) KeysChanged(b SettingsBackend, path string, items []string, originTag interface{}) {
	var arg0 *C.GSettingsBackend
	var arg1 *C.gchar
	var arg2 **C.gchar
	var arg3 C.gpointer

	arg0 = (*C.GSettingsBackend)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.malloc(len(items) * (unsafe.Sizeof(int(0)) + 1))
	defer C.free(unsafe.Pointer(arg2))

	{
		var out []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(arg2), int(len(items)))

		for i := range items {
			out[i] = (*C.gchar)(C.CString(items[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}
	arg3 = C.gpointer(originTag)

	C.g_settings_backend_keys_changed(arg0, arg1, arg2, arg3)
}

// PathChanged signals that all keys below a given path may have possibly
// changed. Backend implementations should call this if an entire path of
// keys have possibly changed their values.
//
// @path must be a valid path (ie starting and ending with a slash and not
// containing '//').
//
// The meaning of this signal is that any of the key which has a name
// starting with @path may have changed.
//
// The same rules for when notifications must occur apply as per
// g_settings_backend_changed(). This call might be an appropriate reasponse
// to a 'reset' call but implementations are also free to explicitly list
// the keys that were affected by that call if they can easily do so.
//
// For efficiency reasons, the implementation should strive for @path to be
// as long as possible (ie: the longest common prefix of all of the keys
// that were changed) but this is not strictly required. As an example, if
// this function is called with the path of "/" then every single key in the
// application will be notified of a possible change.
func (b settingsBackend) PathChanged(b SettingsBackend, path string, originTag interface{}) {
	var arg0 *C.GSettingsBackend
	var arg1 *C.gchar
	var arg2 C.gpointer

	arg0 = (*C.GSettingsBackend)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gpointer(originTag)

	C.g_settings_backend_path_changed(arg0, arg1, arg2)
}

// PathWritableChanged signals that the writability of all keys below a
// given path may have changed.
//
// Since GSettings performs no locking operations for itself, this call will
// always be made in response to external events.
func (b settingsBackend) PathWritableChanged(b SettingsBackend, path string) {
	var arg0 *C.GSettingsBackend
	var arg1 *C.gchar

	arg0 = (*C.GSettingsBackend)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(path))
	defer C.free(unsafe.Pointer(arg1))

	C.g_settings_backend_path_writable_changed(arg0, arg1)
}

// WritableChanged signals that the writability of a single key has possibly
// changed.
//
// Since GSettings performs no locking operations for itself, this call will
// always be made in response to external events.
func (b settingsBackend) WritableChanged(b SettingsBackend, key string) {
	var arg0 *C.GSettingsBackend
	var arg1 *C.gchar

	arg0 = (*C.GSettingsBackend)(unsafe.Pointer(b.Native()))
	arg1 = (*C.gchar)(C.CString(key))
	defer C.free(unsafe.Pointer(arg1))

	C.g_settings_backend_writable_changed(arg0, arg1)
}
