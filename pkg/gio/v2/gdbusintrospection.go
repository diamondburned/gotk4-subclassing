// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gio-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
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
//
// An instance of this type is always passed by reference.
type DBusAnnotationInfo struct {
	*dBusAnnotationInfo
}

// dBusAnnotationInfo is the struct that's finalized.
type dBusAnnotationInfo struct {
	native *C.GDBusAnnotationInfo
}

func marshalDBusAnnotationInfo(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &DBusAnnotationInfo{&dBusAnnotationInfo{(*C.GDBusAnnotationInfo)(b)}}, nil
}

// RefCount: reference count or -1 if statically allocated.
func (d *DBusAnnotationInfo) RefCount() int {
	var v int // out
	v = int(d.native.ref_count)
	return v
}

// Key: name of the annotation, e.g. "org.freedesktop.DBus.Deprecated".
func (d *DBusAnnotationInfo) Key() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(d.native.key)))
	return v
}

// Value: value of the annotation.
func (d *DBusAnnotationInfo) Value() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(d.native.value)))
	return v
}

// Annotations: pointer to a NULL-terminated array of pointers to
// BusAnnotationInfo structures or NULL if there are no annotations.
func (d *DBusAnnotationInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo // out
	{
		var i int
		var z *C.GDBusAnnotationInfo
		for p := d.native.annotations; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(d.native.annotations, i)
		v = make([]*DBusAnnotationInfo, i)
		for i := range src {
			v[i] = (*DBusAnnotationInfo)(gextras.NewStructNative(unsafe.Pointer(src[i])))
			C.g_dbus_annotation_info_ref(src[i])
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.g_dbus_annotation_info_unref((*C.GDBusAnnotationInfo)(intern.C))
				},
			)
		}
	}
	return v
}

// DBusAnnotationInfoLookup looks up the value of an annotation.
//
// The cost of this function is O(n) in number of annotations.
//
// The function takes the following parameters:
//
//    - annotations: NULL-terminated array of annotations or NULL.
//    - name of the annotation to look up.
//
func DBusAnnotationInfoLookup(annotations []*DBusAnnotationInfo, name string) string {
	var _arg1 **C.GDBusAnnotationInfo // out
	var _arg2 *C.gchar                // out
	var _cret *C.gchar                // in

	if annotations != nil {
		{
			_arg1 = (**C.GDBusAnnotationInfo)(C.calloc(C.size_t((len(annotations) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
			defer C.free(unsafe.Pointer(_arg1))
			{
				out := unsafe.Slice(_arg1, len(annotations)+1)
				var zero *C.GDBusAnnotationInfo
				out[len(annotations)] = zero
				for i := range annotations {
					out[i] = (*C.GDBusAnnotationInfo)(gextras.StructNative(unsafe.Pointer(annotations[i])))
				}
			}
		}
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_dbus_annotation_info_lookup(_arg1, _arg2)
	runtime.KeepAlive(annotations)
	runtime.KeepAlive(name)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// DBusArgInfo: information about an argument for a method or a signal.
//
// An instance of this type is always passed by reference.
type DBusArgInfo struct {
	*dBusArgInfo
}

// dBusArgInfo is the struct that's finalized.
type dBusArgInfo struct {
	native *C.GDBusArgInfo
}

func marshalDBusArgInfo(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &DBusArgInfo{&dBusArgInfo{(*C.GDBusArgInfo)(b)}}, nil
}

// RefCount: reference count or -1 if statically allocated.
func (d *DBusArgInfo) RefCount() int {
	var v int // out
	v = int(d.native.ref_count)
	return v
}

// Name of the argument, e.g. unix_user_id.
func (d *DBusArgInfo) Name() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(d.native.name)))
	return v
}

// Signature d-Bus signature of the argument (a single complete type).
func (d *DBusArgInfo) Signature() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(d.native.signature)))
	return v
}

// Annotations: pointer to a NULL-terminated array of pointers to
// BusAnnotationInfo structures or NULL if there are no annotations.
func (d *DBusArgInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo // out
	{
		var i int
		var z *C.GDBusAnnotationInfo
		for p := d.native.annotations; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(d.native.annotations, i)
		v = make([]*DBusAnnotationInfo, i)
		for i := range src {
			v[i] = (*DBusAnnotationInfo)(gextras.NewStructNative(unsafe.Pointer(src[i])))
			C.g_dbus_annotation_info_ref(src[i])
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.g_dbus_annotation_info_unref((*C.GDBusAnnotationInfo)(intern.C))
				},
			)
		}
	}
	return v
}

// DBusInterfaceInfo: information about a D-Bus interface.
//
// An instance of this type is always passed by reference.
type DBusInterfaceInfo struct {
	*dBusInterfaceInfo
}

// dBusInterfaceInfo is the struct that's finalized.
type dBusInterfaceInfo struct {
	native *C.GDBusInterfaceInfo
}

func marshalDBusInterfaceInfo(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &DBusInterfaceInfo{&dBusInterfaceInfo{(*C.GDBusInterfaceInfo)(b)}}, nil
}

// RefCount: reference count or -1 if statically allocated.
func (d *DBusInterfaceInfo) RefCount() int {
	var v int // out
	v = int(d.native.ref_count)
	return v
}

// Name: name of the D-Bus interface, e.g. "org.freedesktop.DBus.Properties".
func (d *DBusInterfaceInfo) Name() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(d.native.name)))
	return v
}

// Methods: pointer to a NULL-terminated array of pointers to BusMethodInfo
// structures or NULL if there are no methods.
func (d *DBusInterfaceInfo) Methods() []*DBusMethodInfo {
	var v []*DBusMethodInfo // out
	{
		var i int
		var z *C.GDBusMethodInfo
		for p := d.native.methods; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(d.native.methods, i)
		v = make([]*DBusMethodInfo, i)
		for i := range src {
			v[i] = (*DBusMethodInfo)(gextras.NewStructNative(unsafe.Pointer(src[i])))
			C.g_dbus_method_info_ref(src[i])
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.g_dbus_method_info_unref((*C.GDBusMethodInfo)(intern.C))
				},
			)
		}
	}
	return v
}

// Signals: pointer to a NULL-terminated array of pointers to BusSignalInfo
// structures or NULL if there are no signals.
func (d *DBusInterfaceInfo) Signals() []*DBusSignalInfo {
	var v []*DBusSignalInfo // out
	{
		var i int
		var z *C.GDBusSignalInfo
		for p := d.native.signals; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(d.native.signals, i)
		v = make([]*DBusSignalInfo, i)
		for i := range src {
			v[i] = (*DBusSignalInfo)(gextras.NewStructNative(unsafe.Pointer(src[i])))
			C.g_dbus_signal_info_ref(src[i])
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.g_dbus_signal_info_unref((*C.GDBusSignalInfo)(intern.C))
				},
			)
		}
	}
	return v
}

// Properties: pointer to a NULL-terminated array of pointers to BusPropertyInfo
// structures or NULL if there are no properties.
func (d *DBusInterfaceInfo) Properties() []*DBusPropertyInfo {
	var v []*DBusPropertyInfo // out
	{
		var i int
		var z *C.GDBusPropertyInfo
		for p := d.native.properties; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(d.native.properties, i)
		v = make([]*DBusPropertyInfo, i)
		for i := range src {
			v[i] = (*DBusPropertyInfo)(gextras.NewStructNative(unsafe.Pointer(src[i])))
			C.g_dbus_property_info_ref(src[i])
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.g_dbus_property_info_unref((*C.GDBusPropertyInfo)(intern.C))
				},
			)
		}
	}
	return v
}

// Annotations: pointer to a NULL-terminated array of pointers to
// BusAnnotationInfo structures or NULL if there are no annotations.
func (d *DBusInterfaceInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo // out
	{
		var i int
		var z *C.GDBusAnnotationInfo
		for p := d.native.annotations; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(d.native.annotations, i)
		v = make([]*DBusAnnotationInfo, i)
		for i := range src {
			v[i] = (*DBusAnnotationInfo)(gextras.NewStructNative(unsafe.Pointer(src[i])))
			C.g_dbus_annotation_info_ref(src[i])
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.g_dbus_annotation_info_unref((*C.GDBusAnnotationInfo)(intern.C))
				},
			)
		}
	}
	return v
}

// CacheBuild builds a lookup-cache to speed up
// g_dbus_interface_info_lookup_method(), g_dbus_interface_info_lookup_signal()
// and g_dbus_interface_info_lookup_property().
//
// If this has already been called with info, the existing cache is used and its
// use count is increased.
//
// Note that info cannot be modified until g_dbus_interface_info_cache_release()
// is called.
func (info *DBusInterfaceInfo) CacheBuild() {
	var _arg0 *C.GDBusInterfaceInfo // out

	_arg0 = (*C.GDBusInterfaceInfo)(gextras.StructNative(unsafe.Pointer(info)))

	C.g_dbus_interface_info_cache_build(_arg0)
	runtime.KeepAlive(info)
}

// CacheRelease decrements the usage count for the cache for info built by
// g_dbus_interface_info_cache_build() (if any) and frees the resources used by
// the cache if the usage count drops to zero.
func (info *DBusInterfaceInfo) CacheRelease() {
	var _arg0 *C.GDBusInterfaceInfo // out

	_arg0 = (*C.GDBusInterfaceInfo)(gextras.StructNative(unsafe.Pointer(info)))

	C.g_dbus_interface_info_cache_release(_arg0)
	runtime.KeepAlive(info)
}

// LookupMethod looks up information about a method.
//
// The cost of this function is O(n) in number of methods unless
// g_dbus_interface_info_cache_build() has been used on info.
func (info *DBusInterfaceInfo) LookupMethod(name string) *DBusMethodInfo {
	var _arg0 *C.GDBusInterfaceInfo // out
	var _arg1 *C.gchar              // out
	var _cret *C.GDBusMethodInfo    // in

	_arg0 = (*C.GDBusInterfaceInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_interface_info_lookup_method(_arg0, _arg1)
	runtime.KeepAlive(info)
	runtime.KeepAlive(name)

	var _dBusMethodInfo *DBusMethodInfo // out

	if _cret != nil {
		_dBusMethodInfo = (*DBusMethodInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.g_dbus_method_info_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_dBusMethodInfo)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_dbus_method_info_unref((*C.GDBusMethodInfo)(intern.C))
			},
		)
	}

	return _dBusMethodInfo
}

// LookupProperty looks up information about a property.
//
// The cost of this function is O(n) in number of properties unless
// g_dbus_interface_info_cache_build() has been used on info.
func (info *DBusInterfaceInfo) LookupProperty(name string) *DBusPropertyInfo {
	var _arg0 *C.GDBusInterfaceInfo // out
	var _arg1 *C.gchar              // out
	var _cret *C.GDBusPropertyInfo  // in

	_arg0 = (*C.GDBusInterfaceInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_interface_info_lookup_property(_arg0, _arg1)
	runtime.KeepAlive(info)
	runtime.KeepAlive(name)

	var _dBusPropertyInfo *DBusPropertyInfo // out

	if _cret != nil {
		_dBusPropertyInfo = (*DBusPropertyInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.g_dbus_property_info_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_dBusPropertyInfo)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_dbus_property_info_unref((*C.GDBusPropertyInfo)(intern.C))
			},
		)
	}

	return _dBusPropertyInfo
}

// LookupSignal looks up information about a signal.
//
// The cost of this function is O(n) in number of signals unless
// g_dbus_interface_info_cache_build() has been used on info.
func (info *DBusInterfaceInfo) LookupSignal(name string) *DBusSignalInfo {
	var _arg0 *C.GDBusInterfaceInfo // out
	var _arg1 *C.gchar              // out
	var _cret *C.GDBusSignalInfo    // in

	_arg0 = (*C.GDBusInterfaceInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_interface_info_lookup_signal(_arg0, _arg1)
	runtime.KeepAlive(info)
	runtime.KeepAlive(name)

	var _dBusSignalInfo *DBusSignalInfo // out

	if _cret != nil {
		_dBusSignalInfo = (*DBusSignalInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.g_dbus_signal_info_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_dBusSignalInfo)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_dbus_signal_info_unref((*C.GDBusSignalInfo)(intern.C))
			},
		)
	}

	return _dBusSignalInfo
}

// DBusMethodInfo: information about a method on an D-Bus interface.
//
// An instance of this type is always passed by reference.
type DBusMethodInfo struct {
	*dBusMethodInfo
}

// dBusMethodInfo is the struct that's finalized.
type dBusMethodInfo struct {
	native *C.GDBusMethodInfo
}

func marshalDBusMethodInfo(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &DBusMethodInfo{&dBusMethodInfo{(*C.GDBusMethodInfo)(b)}}, nil
}

// RefCount: reference count or -1 if statically allocated.
func (d *DBusMethodInfo) RefCount() int {
	var v int // out
	v = int(d.native.ref_count)
	return v
}

// Name: name of the D-Bus method, e.g. RequestName.
func (d *DBusMethodInfo) Name() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(d.native.name)))
	return v
}

// InArgs: pointer to a NULL-terminated array of pointers to BusArgInfo
// structures or NULL if there are no in arguments.
func (d *DBusMethodInfo) InArgs() []*DBusArgInfo {
	var v []*DBusArgInfo // out
	{
		var i int
		var z *C.GDBusArgInfo
		for p := d.native.in_args; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(d.native.in_args, i)
		v = make([]*DBusArgInfo, i)
		for i := range src {
			v[i] = (*DBusArgInfo)(gextras.NewStructNative(unsafe.Pointer(src[i])))
			C.g_dbus_arg_info_ref(src[i])
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.g_dbus_arg_info_unref((*C.GDBusArgInfo)(intern.C))
				},
			)
		}
	}
	return v
}

// OutArgs: pointer to a NULL-terminated array of pointers to BusArgInfo
// structures or NULL if there are no out arguments.
func (d *DBusMethodInfo) OutArgs() []*DBusArgInfo {
	var v []*DBusArgInfo // out
	{
		var i int
		var z *C.GDBusArgInfo
		for p := d.native.out_args; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(d.native.out_args, i)
		v = make([]*DBusArgInfo, i)
		for i := range src {
			v[i] = (*DBusArgInfo)(gextras.NewStructNative(unsafe.Pointer(src[i])))
			C.g_dbus_arg_info_ref(src[i])
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.g_dbus_arg_info_unref((*C.GDBusArgInfo)(intern.C))
				},
			)
		}
	}
	return v
}

// Annotations: pointer to a NULL-terminated array of pointers to
// BusAnnotationInfo structures or NULL if there are no annotations.
func (d *DBusMethodInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo // out
	{
		var i int
		var z *C.GDBusAnnotationInfo
		for p := d.native.annotations; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(d.native.annotations, i)
		v = make([]*DBusAnnotationInfo, i)
		for i := range src {
			v[i] = (*DBusAnnotationInfo)(gextras.NewStructNative(unsafe.Pointer(src[i])))
			C.g_dbus_annotation_info_ref(src[i])
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.g_dbus_annotation_info_unref((*C.GDBusAnnotationInfo)(intern.C))
				},
			)
		}
	}
	return v
}

// DBusNodeInfo: information about nodes in a remote object hierarchy.
//
// An instance of this type is always passed by reference.
type DBusNodeInfo struct {
	*dBusNodeInfo
}

// dBusNodeInfo is the struct that's finalized.
type dBusNodeInfo struct {
	native *C.GDBusNodeInfo
}

func marshalDBusNodeInfo(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &DBusNodeInfo{&dBusNodeInfo{(*C.GDBusNodeInfo)(b)}}, nil
}

// NewDBusNodeInfoForXML constructs a struct DBusNodeInfo.
func NewDBusNodeInfoForXML(xmlData string) (*DBusNodeInfo, error) {
	var _arg1 *C.gchar         // out
	var _cret *C.GDBusNodeInfo // in
	var _cerr *C.GError        // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(xmlData)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_node_info_new_for_xml(_arg1, &_cerr)
	runtime.KeepAlive(xmlData)

	var _dBusNodeInfo *DBusNodeInfo // out
	var _goerr error                // out

	_dBusNodeInfo = (*DBusNodeInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_dBusNodeInfo)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_dbus_node_info_unref((*C.GDBusNodeInfo)(intern.C))
		},
	)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _dBusNodeInfo, _goerr
}

// RefCount: reference count or -1 if statically allocated.
func (d *DBusNodeInfo) RefCount() int {
	var v int // out
	v = int(d.native.ref_count)
	return v
}

// Path: path of the node or NULL if omitted. Note that this may be a relative
// path. See the D-Bus specification for more details.
func (d *DBusNodeInfo) Path() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(d.native.path)))
	return v
}

// Interfaces: pointer to a NULL-terminated array of pointers to
// BusInterfaceInfo structures or NULL if there are no interfaces.
func (d *DBusNodeInfo) Interfaces() []*DBusInterfaceInfo {
	var v []*DBusInterfaceInfo // out
	{
		var i int
		var z *C.GDBusInterfaceInfo
		for p := d.native.interfaces; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(d.native.interfaces, i)
		v = make([]*DBusInterfaceInfo, i)
		for i := range src {
			v[i] = (*DBusInterfaceInfo)(gextras.NewStructNative(unsafe.Pointer(src[i])))
			C.g_dbus_interface_info_ref(src[i])
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.g_dbus_interface_info_unref((*C.GDBusInterfaceInfo)(intern.C))
				},
			)
		}
	}
	return v
}

// Nodes: pointer to a NULL-terminated array of pointers to BusNodeInfo
// structures or NULL if there are no nodes.
func (d *DBusNodeInfo) Nodes() []*DBusNodeInfo {
	var v []*DBusNodeInfo // out
	{
		var i int
		var z *C.GDBusNodeInfo
		for p := d.native.nodes; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(d.native.nodes, i)
		v = make([]*DBusNodeInfo, i)
		for i := range src {
			v[i] = (*DBusNodeInfo)(gextras.NewStructNative(unsafe.Pointer(src[i])))
			C.g_dbus_node_info_ref(src[i])
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.g_dbus_node_info_unref((*C.GDBusNodeInfo)(intern.C))
				},
			)
		}
	}
	return v
}

// Annotations: pointer to a NULL-terminated array of pointers to
// BusAnnotationInfo structures or NULL if there are no annotations.
func (d *DBusNodeInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo // out
	{
		var i int
		var z *C.GDBusAnnotationInfo
		for p := d.native.annotations; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(d.native.annotations, i)
		v = make([]*DBusAnnotationInfo, i)
		for i := range src {
			v[i] = (*DBusAnnotationInfo)(gextras.NewStructNative(unsafe.Pointer(src[i])))
			C.g_dbus_annotation_info_ref(src[i])
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.g_dbus_annotation_info_unref((*C.GDBusAnnotationInfo)(intern.C))
				},
			)
		}
	}
	return v
}

// LookupInterface looks up information about an interface.
//
// The cost of this function is O(n) in number of interfaces.
func (info *DBusNodeInfo) LookupInterface(name string) *DBusInterfaceInfo {
	var _arg0 *C.GDBusNodeInfo      // out
	var _arg1 *C.gchar              // out
	var _cret *C.GDBusInterfaceInfo // in

	_arg0 = (*C.GDBusNodeInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_dbus_node_info_lookup_interface(_arg0, _arg1)
	runtime.KeepAlive(info)
	runtime.KeepAlive(name)

	var _dBusInterfaceInfo *DBusInterfaceInfo // out

	if _cret != nil {
		_dBusInterfaceInfo = (*DBusInterfaceInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		C.g_dbus_interface_info_ref(_cret)
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_dBusInterfaceInfo)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.g_dbus_interface_info_unref((*C.GDBusInterfaceInfo)(intern.C))
			},
		)
	}

	return _dBusInterfaceInfo
}

// DBusPropertyInfo: information about a D-Bus property on a D-Bus interface.
//
// An instance of this type is always passed by reference.
type DBusPropertyInfo struct {
	*dBusPropertyInfo
}

// dBusPropertyInfo is the struct that's finalized.
type dBusPropertyInfo struct {
	native *C.GDBusPropertyInfo
}

func marshalDBusPropertyInfo(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &DBusPropertyInfo{&dBusPropertyInfo{(*C.GDBusPropertyInfo)(b)}}, nil
}

// RefCount: reference count or -1 if statically allocated.
func (d *DBusPropertyInfo) RefCount() int {
	var v int // out
	v = int(d.native.ref_count)
	return v
}

// Name: name of the D-Bus property, e.g. "SupportedFilesystems".
func (d *DBusPropertyInfo) Name() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(d.native.name)))
	return v
}

// Signature d-Bus signature of the property (a single complete type).
func (d *DBusPropertyInfo) Signature() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(d.native.signature)))
	return v
}

// Flags access control flags for the property.
func (d *DBusPropertyInfo) Flags() DBusPropertyInfoFlags {
	var v DBusPropertyInfoFlags // out
	v = DBusPropertyInfoFlags(d.native.flags)
	return v
}

// Annotations: pointer to a NULL-terminated array of pointers to
// BusAnnotationInfo structures or NULL if there are no annotations.
func (d *DBusPropertyInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo // out
	{
		var i int
		var z *C.GDBusAnnotationInfo
		for p := d.native.annotations; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(d.native.annotations, i)
		v = make([]*DBusAnnotationInfo, i)
		for i := range src {
			v[i] = (*DBusAnnotationInfo)(gextras.NewStructNative(unsafe.Pointer(src[i])))
			C.g_dbus_annotation_info_ref(src[i])
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.g_dbus_annotation_info_unref((*C.GDBusAnnotationInfo)(intern.C))
				},
			)
		}
	}
	return v
}

// DBusSignalInfo: information about a signal on a D-Bus interface.
//
// An instance of this type is always passed by reference.
type DBusSignalInfo struct {
	*dBusSignalInfo
}

// dBusSignalInfo is the struct that's finalized.
type dBusSignalInfo struct {
	native *C.GDBusSignalInfo
}

func marshalDBusSignalInfo(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &DBusSignalInfo{&dBusSignalInfo{(*C.GDBusSignalInfo)(b)}}, nil
}

// RefCount: reference count or -1 if statically allocated.
func (d *DBusSignalInfo) RefCount() int {
	var v int // out
	v = int(d.native.ref_count)
	return v
}

// Name: name of the D-Bus signal, e.g. "NameOwnerChanged".
func (d *DBusSignalInfo) Name() string {
	var v string // out
	v = C.GoString((*C.gchar)(unsafe.Pointer(d.native.name)))
	return v
}

// Args: pointer to a NULL-terminated array of pointers to BusArgInfo structures
// or NULL if there are no arguments.
func (d *DBusSignalInfo) Args() []*DBusArgInfo {
	var v []*DBusArgInfo // out
	{
		var i int
		var z *C.GDBusArgInfo
		for p := d.native.args; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(d.native.args, i)
		v = make([]*DBusArgInfo, i)
		for i := range src {
			v[i] = (*DBusArgInfo)(gextras.NewStructNative(unsafe.Pointer(src[i])))
			C.g_dbus_arg_info_ref(src[i])
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.g_dbus_arg_info_unref((*C.GDBusArgInfo)(intern.C))
				},
			)
		}
	}
	return v
}

// Annotations: pointer to a NULL-terminated array of pointers to
// BusAnnotationInfo structures or NULL if there are no annotations.
func (d *DBusSignalInfo) Annotations() []*DBusAnnotationInfo {
	var v []*DBusAnnotationInfo // out
	{
		var i int
		var z *C.GDBusAnnotationInfo
		for p := d.native.annotations; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(d.native.annotations, i)
		v = make([]*DBusAnnotationInfo, i)
		for i := range src {
			v[i] = (*DBusAnnotationInfo)(gextras.NewStructNative(unsafe.Pointer(src[i])))
			C.g_dbus_annotation_info_ref(src[i])
			runtime.SetFinalizer(
				gextras.StructIntern(unsafe.Pointer(v[i])),
				func(intern *struct{ C unsafe.Pointer }) {
					C.g_dbus_annotation_info_unref((*C.GDBusAnnotationInfo)(intern.C))
				},
			)
		}
	}
	return v
}
