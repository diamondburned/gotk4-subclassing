// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"
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

// ResourceLoad loads a binary resource bundle and creates a #GResource
// representation of it, allowing you to query it for data.
//
// If you want to use this resource in the global resource namespace you need to
// register it with g_resources_register().
//
// If @filename is empty or the data in it is corrupt, G_RESOURCE_ERROR_INTERNAL
// will be returned. If @filename doesn’t exist, or there is an error in reading
// it, an error from g_mapped_file_new() will be returned.
func ResourceLoad(filename string) error {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))

	var errout *C.GError
	var err error

	C.g_resource_load(arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// ResourcesEnumerateChildren returns all the names of children at the specified
// @path in the set of globally registered resources. The return result is a nil
// terminated list of strings which should be released with g_strfreev().
//
// @lookup_flags controls the behaviour of the lookup.
func ResourcesEnumerateChildren(path string, lookupFlags ResourceLookupFlags) error {
	var arg1 *C.char
	var arg2 C.GResourceLookupFlags

	arg1 = (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (C.GResourceLookupFlags)(lookupFlags)

	var errout *C.GError
	var err error

	C.g_resources_enumerate_children(arg1, arg2, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// ResourcesGetInfo looks for a file at the specified @path in the set of
// globally registered resources and if found returns information about it.
//
// @lookup_flags controls the behaviour of the lookup.
func ResourcesGetInfo(path string, lookupFlags ResourceLookupFlags) (size uint, flags uint32, err error) {
	var arg1 *C.char
	var arg2 C.GResourceLookupFlags

	arg1 = (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (C.GResourceLookupFlags)(lookupFlags)

	var arg3 C.gsize
	var size uint
	var arg4 C.guint32
	var flags uint32
	var errout *C.GError
	var err error

	C.g_resources_get_info(arg1, arg2, &arg3, &arg4, &errout)

	size = uint(&arg3)
	flags = uint32(&arg4)
	err = gerror.Take(unsafe.Pointer(errout))

	return size, flags, err
}

// ResourcesLookupData looks for a file at the specified @path in the set of
// globally registered resources and returns a #GBytes that lets you directly
// access the data in memory.
//
// The data is always followed by a zero byte, so you can safely use the data as
// a C string. However, that byte is not included in the size of the GBytes.
//
// For uncompressed resource files this is a pointer directly into the resource
// bundle, which is typically in some readonly data section in the program
// binary. For compressed files we allocate memory on the heap and automatically
// uncompress the data.
//
// @lookup_flags controls the behaviour of the lookup.
func ResourcesLookupData(path string, lookupFlags ResourceLookupFlags) error {
	var arg1 *C.char
	var arg2 C.GResourceLookupFlags

	arg1 = (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (C.GResourceLookupFlags)(lookupFlags)

	var errout *C.GError
	var err error

	C.g_resources_lookup_data(arg1, arg2, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// ResourcesOpenStream looks for a file at the specified @path in the set of
// globally registered resources and returns a Stream that lets you read the
// data.
//
// @lookup_flags controls the behaviour of the lookup.
func ResourcesOpenStream(path string, lookupFlags ResourceLookupFlags) error {
	var arg1 *C.char
	var arg2 C.GResourceLookupFlags

	arg1 = (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (C.GResourceLookupFlags)(lookupFlags)

	var errout *C.GError
	var err error

	C.g_resources_open_stream(arg1, arg2, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// ResourcesRegister registers the resource with the process-global set of
// resources. Once a resource is registered the files in it can be accessed with
// the global resource lookup functions like g_resources_lookup_data().
func ResourcesRegister(resource *Resource) {
	var arg1 *C.GResource

	arg1 = (*C.GResource)(unsafe.Pointer(resource.Native()))

	C.g_resources_register(arg1)
}

// ResourcesUnregister unregisters the resource from the process-global set of
// resources.
func ResourcesUnregister(resource *Resource) {
	var arg1 *C.GResource

	arg1 = (*C.GResource)(unsafe.Pointer(resource.Native()))

	C.g_resources_unregister(arg1)
}

// StaticResource is an opaque data structure and can only be accessed using the
// following functions.
type StaticResource struct {
	native C.GStaticResource
}

// WrapStaticResource wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapStaticResource(ptr unsafe.Pointer) *StaticResource {
	if ptr == nil {
		return nil
	}

	return (*StaticResource)(ptr)
}

func marshalStaticResource(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapStaticResource(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *StaticResource) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// Fini: finalized a GResource initialized by g_static_resource_init().
//
// This is normally used by code generated by
// [glib-compile-resources][glib-compile-resources] and is not typically used by
// other code.
func (s *StaticResource) Fini(s *StaticResource) {
	var arg0 *C.GStaticResource

	arg0 = (*C.GStaticResource)(unsafe.Pointer(s.Native()))

	C.g_static_resource_fini(arg0)
}

// Resource gets the GResource that was registered by a call to
// g_static_resource_init().
//
// This is normally used by code generated by
// [glib-compile-resources][glib-compile-resources] and is not typically used by
// other code.
func (s *StaticResource) Resource(s *StaticResource) {
	var arg0 *C.GStaticResource

	arg0 = (*C.GStaticResource)(unsafe.Pointer(s.Native()))

	C.g_static_resource_get_resource(arg0)
}

// Init initializes a GResource from static data using a GStaticResource.
//
// This is normally used by code generated by
// [glib-compile-resources][glib-compile-resources] and is not typically used by
// other code.
func (s *StaticResource) Init(s *StaticResource) {
	var arg0 *C.GStaticResource

	arg0 = (*C.GStaticResource)(unsafe.Pointer(s.Native()))

	C.g_static_resource_init(arg0)
}