// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
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

// ResourcesEnumerateChildren returns all the names of children at the specified
// @path in the set of globally registered resources. The return result is a nil
// terminated list of strings which should be released with g_strfreev().
//
// @lookup_flags controls the behaviour of the lookup.
func ResourcesEnumerateChildren(path string, lookupFlags ResourceLookupFlags) (utf8s []string, goerr error) {
	var arg1 *C.char
	var arg2 C.GResourceLookupFlags

	arg1 = (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (C.GResourceLookupFlags)(lookupFlags)

	var cret **C.char
	var cerr *C.GError

	cret = C.g_resources_enumerate_children(arg1, arg2, cerr)

	var utf8s []string
	var goerr error

	{
		var length int
		for p := cret; *p != 0; p = (**C.char)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
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
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return utf8s, goerr
}

// ResourcesGetInfo looks for a file at the specified @path in the set of
// globally registered resources and if found returns information about it.
//
// @lookup_flags controls the behaviour of the lookup.
func ResourcesGetInfo(path string, lookupFlags ResourceLookupFlags) (size uint, flags uint32, goerr error) {
	var arg1 *C.char
	var arg2 C.GResourceLookupFlags

	arg1 = (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (C.GResourceLookupFlags)(lookupFlags)

	var arg3 C.gsize
	var arg4 C.guint32
	var cerr *C.GError

	C.g_resources_get_info(arg1, arg2, &arg3, &arg4, cerr)

	var size uint
	var flags uint32
	var goerr error

	size = (uint)(arg3)
	flags = (uint32)(arg4)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return size, flags, goerr
}

// ResourcesOpenStream looks for a file at the specified @path in the set of
// globally registered resources and returns a Stream that lets you read the
// data.
//
// @lookup_flags controls the behaviour of the lookup.
func ResourcesOpenStream(path string, lookupFlags ResourceLookupFlags) (inputStream InputStream, goerr error) {
	var arg1 *C.char
	var arg2 C.GResourceLookupFlags

	arg1 = (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (C.GResourceLookupFlags)(lookupFlags)

	var cret *C.GInputStream
	var cerr *C.GError

	cret = C.g_resources_open_stream(arg1, arg2, cerr)

	var inputStream InputStream
	var goerr error

	inputStream = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(InputStream)
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return inputStream, goerr
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
func (s *StaticResource) Fini() {
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
func (s *StaticResource) Resource() *Resource {
	var arg0 *C.GStaticResource

	arg0 = (*C.GStaticResource)(unsafe.Pointer(s.Native()))

	var cret *C.GResource

	cret = C.g_static_resource_get_resource(arg0)

	var resource *Resource

	resource = WrapResource(unsafe.Pointer(cret))

	return resource
}

// Init initializes a GResource from static data using a GStaticResource.
//
// This is normally used by code generated by
// [glib-compile-resources][glib-compile-resources] and is not typically used by
// other code.
func (s *StaticResource) Init() {
	var arg0 *C.GStaticResource

	arg0 = (*C.GStaticResource)(unsafe.Pointer(s.Native()))

	C.g_static_resource_init(arg0)
}
