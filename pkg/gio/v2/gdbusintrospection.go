// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
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
		{T: externglib.Type(C.g_dbus_annotation_info_get_type()), F: marshalDBusAnnotationInfo},
		{T: externglib.Type(C.g_dbus_arg_info_get_type()), F: marshalDBusArgInfo},
		{T: externglib.Type(C.g_dbus_interface_info_get_type()), F: marshalDBusInterfaceInfo},
		{T: externglib.Type(C.g_dbus_method_info_get_type()), F: marshalDBusMethodInfo},
		{T: externglib.Type(C.g_dbus_node_info_get_type()), F: marshalDBusNodeInfo},
		{T: externglib.Type(C.g_dbus_property_info_get_type()), F: marshalDBusPropertyInfo},
		{T: externglib.Type(C.g_dbus_signal_info_get_type()), F: marshalDBusSignalInfo},
	})
}

// DBusAnnotationInfo: information about an annotation.
type DBusAnnotationInfo struct {
	native C.GDBusAnnotationInfo
}

// WrapDBusAnnotationInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusAnnotationInfo(ptr unsafe.Pointer) *DBusAnnotationInfo {
	if ptr == nil {
		return nil
	}

	return (*DBusAnnotationInfo)(ptr)
}

func marshalDBusAnnotationInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapDBusAnnotationInfo(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusAnnotationInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// RefCount gets the field inside the struct.
func (d *DBusAnnotationInfo) RefCount() int {
	var v int
	v = (int)(d.native.ref_count)
	return v
}

// Key gets the field inside the struct.
func (d *DBusAnnotationInfo) Key() string {
	var v string
	v = C.GoString(d.native.key)
	return v
}

// Value gets the field inside the struct.
func (d *DBusAnnotationInfo) Value() string {
	var v string
	v = C.GoString(d.native.value)
	return v
}

// Annotations gets the field inside the struct.
func (d *DBusAnnotationInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo
	{
		var length int
		for p := d.native.annotations; *p != 0; p = (**C.GDBusAnnotationInfo)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.GDBusAnnotationInfo
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(d.native.annotations), int(length))

		v = make([]*DBusAnnotationInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			v = WrapDBusAnnotationInfo(unsafe.Pointer(d.native.annotations))
		}
	}
	return v
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusAnnotationInfo) Ref() *DBusAnnotationInfo {
	var arg0 *C.GDBusAnnotationInfo

	arg0 = (*C.GDBusAnnotationInfo)(unsafe.Pointer(i.Native()))

	var cret *C.GDBusAnnotationInfo

	cret = C.g_dbus_annotation_info_ref(arg0)

	var dBusAnnotationInfo *DBusAnnotationInfo

	dBusAnnotationInfo = WrapDBusAnnotationInfo(unsafe.Pointer(cret))
	runtime.SetFinalizer(dBusAnnotationInfo, func(v *DBusAnnotationInfo) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return dBusAnnotationInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusAnnotationInfo) Unref() {
	var arg0 *C.GDBusAnnotationInfo

	arg0 = (*C.GDBusAnnotationInfo)(unsafe.Pointer(i.Native()))

	C.g_dbus_annotation_info_unref(arg0)
}

// DBusArgInfo: information about an argument for a method or a signal.
type DBusArgInfo struct {
	native C.GDBusArgInfo
}

// WrapDBusArgInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusArgInfo(ptr unsafe.Pointer) *DBusArgInfo {
	if ptr == nil {
		return nil
	}

	return (*DBusArgInfo)(ptr)
}

func marshalDBusArgInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapDBusArgInfo(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusArgInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// RefCount gets the field inside the struct.
func (d *DBusArgInfo) RefCount() int {
	var v int
	v = (int)(d.native.ref_count)
	return v
}

// Name gets the field inside the struct.
func (d *DBusArgInfo) Name() string {
	var v string
	v = C.GoString(d.native.name)
	return v
}

// Signature gets the field inside the struct.
func (d *DBusArgInfo) Signature() string {
	var v string
	v = C.GoString(d.native.signature)
	return v
}

// Annotations gets the field inside the struct.
func (d *DBusArgInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo
	{
		var length int
		for p := d.native.annotations; *p != 0; p = (**C.GDBusAnnotationInfo)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.GDBusAnnotationInfo
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(d.native.annotations), int(length))

		v = make([]*DBusAnnotationInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			v = WrapDBusAnnotationInfo(unsafe.Pointer(d.native.annotations))
		}
	}
	return v
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusArgInfo) Ref() *DBusArgInfo {
	var arg0 *C.GDBusArgInfo

	arg0 = (*C.GDBusArgInfo)(unsafe.Pointer(i.Native()))

	var cret *C.GDBusArgInfo

	cret = C.g_dbus_arg_info_ref(arg0)

	var dBusArgInfo *DBusArgInfo

	dBusArgInfo = WrapDBusArgInfo(unsafe.Pointer(cret))
	runtime.SetFinalizer(dBusArgInfo, func(v *DBusArgInfo) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return dBusArgInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusArgInfo) Unref() {
	var arg0 *C.GDBusArgInfo

	arg0 = (*C.GDBusArgInfo)(unsafe.Pointer(i.Native()))

	C.g_dbus_arg_info_unref(arg0)
}

// DBusInterfaceInfo: information about a D-Bus interface.
type DBusInterfaceInfo struct {
	native C.GDBusInterfaceInfo
}

// WrapDBusInterfaceInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusInterfaceInfo(ptr unsafe.Pointer) *DBusInterfaceInfo {
	if ptr == nil {
		return nil
	}

	return (*DBusInterfaceInfo)(ptr)
}

func marshalDBusInterfaceInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapDBusInterfaceInfo(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusInterfaceInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// RefCount gets the field inside the struct.
func (d *DBusInterfaceInfo) RefCount() int {
	var v int
	v = (int)(d.native.ref_count)
	return v
}

// Name gets the field inside the struct.
func (d *DBusInterfaceInfo) Name() string {
	var v string
	v = C.GoString(d.native.name)
	return v
}

// Methods gets the field inside the struct.
func (d *DBusInterfaceInfo) Methods() []*DBusMethodInfo {
	var v []*DBusMethodInfo
	{
		var length int
		for p := d.native.methods; *p != 0; p = (**C.GDBusMethodInfo)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.GDBusMethodInfo
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(d.native.methods), int(length))

		v = make([]*DBusMethodInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			v = WrapDBusMethodInfo(unsafe.Pointer(d.native.methods))
		}
	}
	return v
}

// Signals gets the field inside the struct.
func (d *DBusInterfaceInfo) Signals() []*DBusSignalInfo {
	var v []*DBusSignalInfo
	{
		var length int
		for p := d.native.signals; *p != 0; p = (**C.GDBusSignalInfo)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.GDBusSignalInfo
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(d.native.signals), int(length))

		v = make([]*DBusSignalInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			v = WrapDBusSignalInfo(unsafe.Pointer(d.native.signals))
		}
	}
	return v
}

// Properties gets the field inside the struct.
func (d *DBusInterfaceInfo) Properties() []*DBusPropertyInfo {
	var v []*DBusPropertyInfo
	{
		var length int
		for p := d.native.properties; *p != 0; p = (**C.GDBusPropertyInfo)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.GDBusPropertyInfo
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(d.native.properties), int(length))

		v = make([]*DBusPropertyInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			v = WrapDBusPropertyInfo(unsafe.Pointer(d.native.properties))
		}
	}
	return v
}

// Annotations gets the field inside the struct.
func (d *DBusInterfaceInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo
	{
		var length int
		for p := d.native.annotations; *p != 0; p = (**C.GDBusAnnotationInfo)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.GDBusAnnotationInfo
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(d.native.annotations), int(length))

		v = make([]*DBusAnnotationInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			v = WrapDBusAnnotationInfo(unsafe.Pointer(d.native.annotations))
		}
	}
	return v
}

// CacheBuild builds a lookup-cache to speed up
// g_dbus_interface_info_lookup_method(), g_dbus_interface_info_lookup_signal()
// and g_dbus_interface_info_lookup_property().
//
// If this has already been called with @info, the existing cache is used and
// its use count is increased.
//
// Note that @info cannot be modified until
// g_dbus_interface_info_cache_release() is called.
func (i *DBusInterfaceInfo) CacheBuild() {
	var arg0 *C.GDBusInterfaceInfo

	arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i.Native()))

	C.g_dbus_interface_info_cache_build(arg0)
}

// CacheRelease decrements the usage count for the cache for @info built by
// g_dbus_interface_info_cache_build() (if any) and frees the resources used by
// the cache if the usage count drops to zero.
func (i *DBusInterfaceInfo) CacheRelease() {
	var arg0 *C.GDBusInterfaceInfo

	arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i.Native()))

	C.g_dbus_interface_info_cache_release(arg0)
}

// GenerateXML appends an XML representation of @info (and its children) to
// @string_builder.
//
// This function is typically used for generating introspection XML documents at
// run-time for handling the `org.freedesktop.DBus.Introspectable.Introspect`
// method.
func (i *DBusInterfaceInfo) GenerateXML(indent uint, stringBuilder *glib.String) {
	var arg0 *C.GDBusInterfaceInfo
	var arg1 C.guint
	var arg2 *C.GString

	arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i.Native()))
	arg1 = C.guint(indent)
	arg2 = (*C.GString)(unsafe.Pointer(stringBuilder.Native()))

	C.g_dbus_interface_info_generate_xml(arg0, arg1, arg2)
}

// LookupMethod looks up information about a method.
//
// The cost of this function is O(n) in number of methods unless
// g_dbus_interface_info_cache_build() has been used on @info.
func (i *DBusInterfaceInfo) LookupMethod(name string) *DBusMethodInfo {
	var arg0 *C.GDBusInterfaceInfo
	var arg1 *C.gchar

	arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GDBusMethodInfo

	cret = C.g_dbus_interface_info_lookup_method(arg0, arg1)

	var dBusMethodInfo *DBusMethodInfo

	dBusMethodInfo = WrapDBusMethodInfo(unsafe.Pointer(cret))

	return dBusMethodInfo
}

// LookupProperty looks up information about a property.
//
// The cost of this function is O(n) in number of properties unless
// g_dbus_interface_info_cache_build() has been used on @info.
func (i *DBusInterfaceInfo) LookupProperty(name string) *DBusPropertyInfo {
	var arg0 *C.GDBusInterfaceInfo
	var arg1 *C.gchar

	arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GDBusPropertyInfo

	cret = C.g_dbus_interface_info_lookup_property(arg0, arg1)

	var dBusPropertyInfo *DBusPropertyInfo

	dBusPropertyInfo = WrapDBusPropertyInfo(unsafe.Pointer(cret))

	return dBusPropertyInfo
}

// LookupSignal looks up information about a signal.
//
// The cost of this function is O(n) in number of signals unless
// g_dbus_interface_info_cache_build() has been used on @info.
func (i *DBusInterfaceInfo) LookupSignal(name string) *DBusSignalInfo {
	var arg0 *C.GDBusInterfaceInfo
	var arg1 *C.gchar

	arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GDBusSignalInfo

	cret = C.g_dbus_interface_info_lookup_signal(arg0, arg1)

	var dBusSignalInfo *DBusSignalInfo

	dBusSignalInfo = WrapDBusSignalInfo(unsafe.Pointer(cret))

	return dBusSignalInfo
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusInterfaceInfo) Ref() *DBusInterfaceInfo {
	var arg0 *C.GDBusInterfaceInfo

	arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i.Native()))

	var cret *C.GDBusInterfaceInfo

	cret = C.g_dbus_interface_info_ref(arg0)

	var dBusInterfaceInfo *DBusInterfaceInfo

	dBusInterfaceInfo = WrapDBusInterfaceInfo(unsafe.Pointer(cret))
	runtime.SetFinalizer(dBusInterfaceInfo, func(v *DBusInterfaceInfo) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return dBusInterfaceInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusInterfaceInfo) Unref() {
	var arg0 *C.GDBusInterfaceInfo

	arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i.Native()))

	C.g_dbus_interface_info_unref(arg0)
}

// DBusMethodInfo: information about a method on an D-Bus interface.
type DBusMethodInfo struct {
	native C.GDBusMethodInfo
}

// WrapDBusMethodInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusMethodInfo(ptr unsafe.Pointer) *DBusMethodInfo {
	if ptr == nil {
		return nil
	}

	return (*DBusMethodInfo)(ptr)
}

func marshalDBusMethodInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapDBusMethodInfo(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusMethodInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// RefCount gets the field inside the struct.
func (d *DBusMethodInfo) RefCount() int {
	var v int
	v = (int)(d.native.ref_count)
	return v
}

// Name gets the field inside the struct.
func (d *DBusMethodInfo) Name() string {
	var v string
	v = C.GoString(d.native.name)
	return v
}

// InArgs gets the field inside the struct.
func (d *DBusMethodInfo) InArgs() []*DBusArgInfo {
	var v []*DBusArgInfo
	{
		var length int
		for p := d.native.in_args; *p != 0; p = (**C.GDBusArgInfo)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.GDBusArgInfo
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(d.native.in_args), int(length))

		v = make([]*DBusArgInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			v = WrapDBusArgInfo(unsafe.Pointer(d.native.in_args))
		}
	}
	return v
}

// OutArgs gets the field inside the struct.
func (d *DBusMethodInfo) OutArgs() []*DBusArgInfo {
	var v []*DBusArgInfo
	{
		var length int
		for p := d.native.out_args; *p != 0; p = (**C.GDBusArgInfo)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.GDBusArgInfo
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(d.native.out_args), int(length))

		v = make([]*DBusArgInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			v = WrapDBusArgInfo(unsafe.Pointer(d.native.out_args))
		}
	}
	return v
}

// Annotations gets the field inside the struct.
func (d *DBusMethodInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo
	{
		var length int
		for p := d.native.annotations; *p != 0; p = (**C.GDBusAnnotationInfo)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.GDBusAnnotationInfo
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(d.native.annotations), int(length))

		v = make([]*DBusAnnotationInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			v = WrapDBusAnnotationInfo(unsafe.Pointer(d.native.annotations))
		}
	}
	return v
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusMethodInfo) Ref() *DBusMethodInfo {
	var arg0 *C.GDBusMethodInfo

	arg0 = (*C.GDBusMethodInfo)(unsafe.Pointer(i.Native()))

	var cret *C.GDBusMethodInfo

	cret = C.g_dbus_method_info_ref(arg0)

	var dBusMethodInfo *DBusMethodInfo

	dBusMethodInfo = WrapDBusMethodInfo(unsafe.Pointer(cret))
	runtime.SetFinalizer(dBusMethodInfo, func(v *DBusMethodInfo) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return dBusMethodInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusMethodInfo) Unref() {
	var arg0 *C.GDBusMethodInfo

	arg0 = (*C.GDBusMethodInfo)(unsafe.Pointer(i.Native()))

	C.g_dbus_method_info_unref(arg0)
}

// DBusNodeInfo: information about nodes in a remote object hierarchy.
type DBusNodeInfo struct {
	native C.GDBusNodeInfo
}

// WrapDBusNodeInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusNodeInfo(ptr unsafe.Pointer) *DBusNodeInfo {
	if ptr == nil {
		return nil
	}

	return (*DBusNodeInfo)(ptr)
}

func marshalDBusNodeInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapDBusNodeInfo(unsafe.Pointer(b)), nil
}

// NewDBusNodeInfoForXML constructs a struct DBusNodeInfo.
func NewDBusNodeInfoForXML(xmlData string) (dBusNodeInfo *DBusNodeInfo, goerr error) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(xmlData))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GDBusNodeInfo
	var cerr *C.GError

	cret = C.g_dbus_node_info_new_for_xml(arg1, cerr)

	var dBusNodeInfo *DBusNodeInfo
	var goerr error

	dBusNodeInfo = WrapDBusNodeInfo(unsafe.Pointer(cret))
	runtime.SetFinalizer(dBusNodeInfo, func(v *DBusNodeInfo) {
		C.free(unsafe.Pointer(v.Native()))
	})
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return dBusNodeInfo, goerr
}

// Native returns the underlying C source pointer.
func (d *DBusNodeInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// RefCount gets the field inside the struct.
func (d *DBusNodeInfo) RefCount() int {
	var v int
	v = (int)(d.native.ref_count)
	return v
}

// Path gets the field inside the struct.
func (d *DBusNodeInfo) Path() string {
	var v string
	v = C.GoString(d.native.path)
	return v
}

// Interfaces gets the field inside the struct.
func (d *DBusNodeInfo) Interfaces() []*DBusInterfaceInfo {
	var v []*DBusInterfaceInfo
	{
		var length int
		for p := d.native.interfaces; *p != 0; p = (**C.GDBusInterfaceInfo)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.GDBusInterfaceInfo
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(d.native.interfaces), int(length))

		v = make([]*DBusInterfaceInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			v = WrapDBusInterfaceInfo(unsafe.Pointer(d.native.interfaces))
		}
	}
	return v
}

// Nodes gets the field inside the struct.
func (d *DBusNodeInfo) Nodes() []*DBusNodeInfo {
	var v []*DBusNodeInfo
	{
		var length int
		for p := d.native.nodes; *p != 0; p = (**C.GDBusNodeInfo)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.GDBusNodeInfo
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(d.native.nodes), int(length))

		v = make([]*DBusNodeInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			v = WrapDBusNodeInfo(unsafe.Pointer(d.native.nodes))
		}
	}
	return v
}

// Annotations gets the field inside the struct.
func (d *DBusNodeInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo
	{
		var length int
		for p := d.native.annotations; *p != 0; p = (**C.GDBusAnnotationInfo)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.GDBusAnnotationInfo
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(d.native.annotations), int(length))

		v = make([]*DBusAnnotationInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			v = WrapDBusAnnotationInfo(unsafe.Pointer(d.native.annotations))
		}
	}
	return v
}

// GenerateXML appends an XML representation of @info (and its children) to
// @string_builder.
//
// This function is typically used for generating introspection XML documents at
// run-time for handling the `org.freedesktop.DBus.Introspectable.Introspect`
// method.
func (i *DBusNodeInfo) GenerateXML(indent uint, stringBuilder *glib.String) {
	var arg0 *C.GDBusNodeInfo
	var arg1 C.guint
	var arg2 *C.GString

	arg0 = (*C.GDBusNodeInfo)(unsafe.Pointer(i.Native()))
	arg1 = C.guint(indent)
	arg2 = (*C.GString)(unsafe.Pointer(stringBuilder.Native()))

	C.g_dbus_node_info_generate_xml(arg0, arg1, arg2)
}

// LookupInterface looks up information about an interface.
//
// The cost of this function is O(n) in number of interfaces.
func (i *DBusNodeInfo) LookupInterface(name string) *DBusInterfaceInfo {
	var arg0 *C.GDBusNodeInfo
	var arg1 *C.gchar

	arg0 = (*C.GDBusNodeInfo)(unsafe.Pointer(i.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GDBusInterfaceInfo

	cret = C.g_dbus_node_info_lookup_interface(arg0, arg1)

	var dBusInterfaceInfo *DBusInterfaceInfo

	dBusInterfaceInfo = WrapDBusInterfaceInfo(unsafe.Pointer(cret))

	return dBusInterfaceInfo
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusNodeInfo) Ref() *DBusNodeInfo {
	var arg0 *C.GDBusNodeInfo

	arg0 = (*C.GDBusNodeInfo)(unsafe.Pointer(i.Native()))

	var cret *C.GDBusNodeInfo

	cret = C.g_dbus_node_info_ref(arg0)

	var dBusNodeInfo *DBusNodeInfo

	dBusNodeInfo = WrapDBusNodeInfo(unsafe.Pointer(cret))
	runtime.SetFinalizer(dBusNodeInfo, func(v *DBusNodeInfo) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return dBusNodeInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusNodeInfo) Unref() {
	var arg0 *C.GDBusNodeInfo

	arg0 = (*C.GDBusNodeInfo)(unsafe.Pointer(i.Native()))

	C.g_dbus_node_info_unref(arg0)
}

// DBusPropertyInfo: information about a D-Bus property on a D-Bus interface.
type DBusPropertyInfo struct {
	native C.GDBusPropertyInfo
}

// WrapDBusPropertyInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusPropertyInfo(ptr unsafe.Pointer) *DBusPropertyInfo {
	if ptr == nil {
		return nil
	}

	return (*DBusPropertyInfo)(ptr)
}

func marshalDBusPropertyInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapDBusPropertyInfo(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusPropertyInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// RefCount gets the field inside the struct.
func (d *DBusPropertyInfo) RefCount() int {
	var v int
	v = (int)(d.native.ref_count)
	return v
}

// Name gets the field inside the struct.
func (d *DBusPropertyInfo) Name() string {
	var v string
	v = C.GoString(d.native.name)
	return v
}

// Signature gets the field inside the struct.
func (d *DBusPropertyInfo) Signature() string {
	var v string
	v = C.GoString(d.native.signature)
	return v
}

// Flags gets the field inside the struct.
func (d *DBusPropertyInfo) Flags() DBusPropertyInfoFlags {
	var v DBusPropertyInfoFlags
	v = DBusPropertyInfoFlags(d.native.flags)
	return v
}

// Annotations gets the field inside the struct.
func (d *DBusPropertyInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo
	{
		var length int
		for p := d.native.annotations; *p != 0; p = (**C.GDBusAnnotationInfo)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.GDBusAnnotationInfo
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(d.native.annotations), int(length))

		v = make([]*DBusAnnotationInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			v = WrapDBusAnnotationInfo(unsafe.Pointer(d.native.annotations))
		}
	}
	return v
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusPropertyInfo) Ref() *DBusPropertyInfo {
	var arg0 *C.GDBusPropertyInfo

	arg0 = (*C.GDBusPropertyInfo)(unsafe.Pointer(i.Native()))

	var cret *C.GDBusPropertyInfo

	cret = C.g_dbus_property_info_ref(arg0)

	var dBusPropertyInfo *DBusPropertyInfo

	dBusPropertyInfo = WrapDBusPropertyInfo(unsafe.Pointer(cret))
	runtime.SetFinalizer(dBusPropertyInfo, func(v *DBusPropertyInfo) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return dBusPropertyInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusPropertyInfo) Unref() {
	var arg0 *C.GDBusPropertyInfo

	arg0 = (*C.GDBusPropertyInfo)(unsafe.Pointer(i.Native()))

	C.g_dbus_property_info_unref(arg0)
}

// DBusSignalInfo: information about a signal on a D-Bus interface.
type DBusSignalInfo struct {
	native C.GDBusSignalInfo
}

// WrapDBusSignalInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDBusSignalInfo(ptr unsafe.Pointer) *DBusSignalInfo {
	if ptr == nil {
		return nil
	}

	return (*DBusSignalInfo)(ptr)
}

func marshalDBusSignalInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapDBusSignalInfo(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DBusSignalInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// RefCount gets the field inside the struct.
func (d *DBusSignalInfo) RefCount() int {
	var v int
	v = (int)(d.native.ref_count)
	return v
}

// Name gets the field inside the struct.
func (d *DBusSignalInfo) Name() string {
	var v string
	v = C.GoString(d.native.name)
	return v
}

// Args gets the field inside the struct.
func (d *DBusSignalInfo) Args() []*DBusArgInfo {
	var v []*DBusArgInfo
	{
		var length int
		for p := d.native.args; *p != 0; p = (**C.GDBusArgInfo)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.GDBusArgInfo
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(d.native.args), int(length))

		v = make([]*DBusArgInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			v = WrapDBusArgInfo(unsafe.Pointer(d.native.args))
		}
	}
	return v
}

// Annotations gets the field inside the struct.
func (d *DBusSignalInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo
	{
		var length int
		for p := d.native.annotations; *p != 0; p = (**C.GDBusAnnotationInfo)(ptr.Add(unsafe.Pointer(p), unsafe.Sizeof(int(0)))) {
			length++
			if length < 0 {
				panic(`length overflow`)
			}
		}

		var src []*C.GDBusAnnotationInfo
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(d.native.annotations), int(length))

		v = make([]*DBusAnnotationInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			v = WrapDBusAnnotationInfo(unsafe.Pointer(d.native.annotations))
		}
	}
	return v
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusSignalInfo) Ref() *DBusSignalInfo {
	var arg0 *C.GDBusSignalInfo

	arg0 = (*C.GDBusSignalInfo)(unsafe.Pointer(i.Native()))

	var cret *C.GDBusSignalInfo

	cret = C.g_dbus_signal_info_ref(arg0)

	var dBusSignalInfo *DBusSignalInfo

	dBusSignalInfo = WrapDBusSignalInfo(unsafe.Pointer(cret))
	runtime.SetFinalizer(dBusSignalInfo, func(v *DBusSignalInfo) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return dBusSignalInfo
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusSignalInfo) Unref() {
	var arg0 *C.GDBusSignalInfo

	arg0 = (*C.GDBusSignalInfo)(unsafe.Pointer(i.Native()))

	C.g_dbus_signal_info_unref(arg0)
}
