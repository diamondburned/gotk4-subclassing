// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/ptr"
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

// DBusAnnotationInfoLookup looks up the value of an annotation.
//
// The cost of this function is O(n) in number of annotations.
func DBusAnnotationInfoLookup(annotations []*DBusAnnotationInfo, name string) {
	var arg1 **C.GDBusAnnotationInfo
	var arg2 *C.gchar

	arg1 = C.malloc(len(annotations) * (unsafe.Sizeof(int(0)) + 1))
	defer C.free(unsafe.Pointer(arg1))

	{
		var out []*C.GDBusAnnotationInfo
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(arg1), int(len(annotations)))

		for i := range annotations {
			out[i] = (*C.GDBusAnnotationInfo)(unsafe.Pointer(annotations[i].Native()))
		}
	}
	arg2 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg2))

	C.g_dbus_annotation_info_lookup(arg1, arg2)
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
	v = int(d.native.ref_count)
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

		v = make([]*DBusAnnotationInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.GDBusAnnotationInfo)(ptr.Add(unsafe.Pointer(d.native.annotations), i))
			v[i] = WrapDBusAnnotationInfo(unsafe.Pointer(src))
		}
	}
	return v
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusAnnotationInfo) Ref(i *DBusAnnotationInfo) {
	var arg0 *C.GDBusAnnotationInfo

	arg0 = (*C.GDBusAnnotationInfo)(unsafe.Pointer(i.Native()))

	C.g_dbus_annotation_info_ref(arg0)
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusAnnotationInfo) Unref(i *DBusAnnotationInfo) {
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
	v = int(d.native.ref_count)
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

		v = make([]*DBusAnnotationInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.GDBusAnnotationInfo)(ptr.Add(unsafe.Pointer(d.native.annotations), i))
			v[i] = WrapDBusAnnotationInfo(unsafe.Pointer(src))
		}
	}
	return v
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusArgInfo) Ref(i *DBusArgInfo) {
	var arg0 *C.GDBusArgInfo

	arg0 = (*C.GDBusArgInfo)(unsafe.Pointer(i.Native()))

	C.g_dbus_arg_info_ref(arg0)
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusArgInfo) Unref(i *DBusArgInfo) {
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
	v = int(d.native.ref_count)
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

		v = make([]*DBusMethodInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.GDBusMethodInfo)(ptr.Add(unsafe.Pointer(d.native.methods), i))
			v[i] = WrapDBusMethodInfo(unsafe.Pointer(src))
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

		v = make([]*DBusSignalInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.GDBusSignalInfo)(ptr.Add(unsafe.Pointer(d.native.signals), i))
			v[i] = WrapDBusSignalInfo(unsafe.Pointer(src))
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

		v = make([]*DBusPropertyInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.GDBusPropertyInfo)(ptr.Add(unsafe.Pointer(d.native.properties), i))
			v[i] = WrapDBusPropertyInfo(unsafe.Pointer(src))
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

		v = make([]*DBusAnnotationInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.GDBusAnnotationInfo)(ptr.Add(unsafe.Pointer(d.native.annotations), i))
			v[i] = WrapDBusAnnotationInfo(unsafe.Pointer(src))
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
func (i *DBusInterfaceInfo) CacheBuild(i *DBusInterfaceInfo) {
	var arg0 *C.GDBusInterfaceInfo

	arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i.Native()))

	C.g_dbus_interface_info_cache_build(arg0)
}

// CacheRelease decrements the usage count for the cache for @info built by
// g_dbus_interface_info_cache_build() (if any) and frees the resources used by
// the cache if the usage count drops to zero.
func (i *DBusInterfaceInfo) CacheRelease(i *DBusInterfaceInfo) {
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
func (i *DBusInterfaceInfo) GenerateXML(i *DBusInterfaceInfo, indent uint, stringBuilder *glib.String) {
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
func (i *DBusInterfaceInfo) LookupMethod(i *DBusInterfaceInfo, name string) {
	var arg0 *C.GDBusInterfaceInfo
	var arg1 *C.gchar

	arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	C.g_dbus_interface_info_lookup_method(arg0, arg1)
}

// LookupProperty looks up information about a property.
//
// The cost of this function is O(n) in number of properties unless
// g_dbus_interface_info_cache_build() has been used on @info.
func (i *DBusInterfaceInfo) LookupProperty(i *DBusInterfaceInfo, name string) {
	var arg0 *C.GDBusInterfaceInfo
	var arg1 *C.gchar

	arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	C.g_dbus_interface_info_lookup_property(arg0, arg1)
}

// LookupSignal looks up information about a signal.
//
// The cost of this function is O(n) in number of signals unless
// g_dbus_interface_info_cache_build() has been used on @info.
func (i *DBusInterfaceInfo) LookupSignal(i *DBusInterfaceInfo, name string) {
	var arg0 *C.GDBusInterfaceInfo
	var arg1 *C.gchar

	arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	C.g_dbus_interface_info_lookup_signal(arg0, arg1)
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusInterfaceInfo) Ref(i *DBusInterfaceInfo) {
	var arg0 *C.GDBusInterfaceInfo

	arg0 = (*C.GDBusInterfaceInfo)(unsafe.Pointer(i.Native()))

	C.g_dbus_interface_info_ref(arg0)
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusInterfaceInfo) Unref(i *DBusInterfaceInfo) {
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
	v = int(d.native.ref_count)
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

		v = make([]*DBusArgInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.GDBusArgInfo)(ptr.Add(unsafe.Pointer(d.native.in_args), i))
			v[i] = WrapDBusArgInfo(unsafe.Pointer(src))
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

		v = make([]*DBusArgInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.GDBusArgInfo)(ptr.Add(unsafe.Pointer(d.native.out_args), i))
			v[i] = WrapDBusArgInfo(unsafe.Pointer(src))
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

		v = make([]*DBusAnnotationInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.GDBusAnnotationInfo)(ptr.Add(unsafe.Pointer(d.native.annotations), i))
			v[i] = WrapDBusAnnotationInfo(unsafe.Pointer(src))
		}
	}
	return v
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusMethodInfo) Ref(i *DBusMethodInfo) {
	var arg0 *C.GDBusMethodInfo

	arg0 = (*C.GDBusMethodInfo)(unsafe.Pointer(i.Native()))

	C.g_dbus_method_info_ref(arg0)
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusMethodInfo) Unref(i *DBusMethodInfo) {
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
func NewDBusNodeInfoForXML(xmlData string) error {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(xmlData))
	defer C.free(unsafe.Pointer(arg1))

	var errout *C.GError
	var err error

	C.g_dbus_node_info_new_for_xml(arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// Native returns the underlying C source pointer.
func (d *DBusNodeInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// RefCount gets the field inside the struct.
func (d *DBusNodeInfo) RefCount() int {
	var v int
	v = int(d.native.ref_count)
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

		v = make([]*DBusInterfaceInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.GDBusInterfaceInfo)(ptr.Add(unsafe.Pointer(d.native.interfaces), i))
			v[i] = WrapDBusInterfaceInfo(unsafe.Pointer(src))
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

		v = make([]*DBusNodeInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.GDBusNodeInfo)(ptr.Add(unsafe.Pointer(d.native.nodes), i))
			v[i] = WrapDBusNodeInfo(unsafe.Pointer(src))
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

		v = make([]*DBusAnnotationInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.GDBusAnnotationInfo)(ptr.Add(unsafe.Pointer(d.native.annotations), i))
			v[i] = WrapDBusAnnotationInfo(unsafe.Pointer(src))
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
func (i *DBusNodeInfo) GenerateXML(i *DBusNodeInfo, indent uint, stringBuilder *glib.String) {
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
func (i *DBusNodeInfo) LookupInterface(i *DBusNodeInfo, name string) {
	var arg0 *C.GDBusNodeInfo
	var arg1 *C.gchar

	arg0 = (*C.GDBusNodeInfo)(unsafe.Pointer(i.Native()))
	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	C.g_dbus_node_info_lookup_interface(arg0, arg1)
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusNodeInfo) Ref(i *DBusNodeInfo) {
	var arg0 *C.GDBusNodeInfo

	arg0 = (*C.GDBusNodeInfo)(unsafe.Pointer(i.Native()))

	C.g_dbus_node_info_ref(arg0)
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusNodeInfo) Unref(i *DBusNodeInfo) {
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
	v = int(d.native.ref_count)
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

		v = make([]*DBusAnnotationInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.GDBusAnnotationInfo)(ptr.Add(unsafe.Pointer(d.native.annotations), i))
			v[i] = WrapDBusAnnotationInfo(unsafe.Pointer(src))
		}
	}
	return v
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusPropertyInfo) Ref(i *DBusPropertyInfo) {
	var arg0 *C.GDBusPropertyInfo

	arg0 = (*C.GDBusPropertyInfo)(unsafe.Pointer(i.Native()))

	C.g_dbus_property_info_ref(arg0)
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusPropertyInfo) Unref(i *DBusPropertyInfo) {
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
	v = int(d.native.ref_count)
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

		v = make([]*DBusArgInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.GDBusArgInfo)(ptr.Add(unsafe.Pointer(d.native.args), i))
			v[i] = WrapDBusArgInfo(unsafe.Pointer(src))
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

		v = make([]*DBusAnnotationInfo, length)
		for i := uintptr(0); i < uintptr(length); i += unsafe.Sizeof(int(0)) {
			src := (*C.GDBusAnnotationInfo)(ptr.Add(unsafe.Pointer(d.native.annotations), i))
			v[i] = WrapDBusAnnotationInfo(unsafe.Pointer(src))
		}
	}
	return v
}

// Ref: if @info is statically allocated does nothing. Otherwise increases the
// reference count.
func (i *DBusSignalInfo) Ref(i *DBusSignalInfo) {
	var arg0 *C.GDBusSignalInfo

	arg0 = (*C.GDBusSignalInfo)(unsafe.Pointer(i.Native()))

	C.g_dbus_signal_info_ref(arg0)
}

// Unref: if @info is statically allocated, does nothing. Otherwise decreases
// the reference count of @info. When its reference count drops to 0, the memory
// used is freed.
func (i *DBusSignalInfo) Unref(i *DBusSignalInfo) {
	var arg0 *C.GDBusSignalInfo

	arg0 = (*C.GDBusSignalInfo)(unsafe.Pointer(i.Native()))

	C.g_dbus_signal_info_unref(arg0)
}