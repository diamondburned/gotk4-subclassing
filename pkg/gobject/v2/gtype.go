// Code generated by girgen. DO NOT EDIT.

package gobject

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gobject-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib-object.h>
import "C"

// Type: a numerical value which represents the unique identifier of a
// registered type.
type Type uint

// TypeDebugFlags: these flags used to be passed to
// g_type_init_with_debug_flags() which is now deprecated.
//
// If you need to enable debugging features, use the GOBJECT_DEBUG environment
// variable.
type TypeDebugFlags int

const (
	// TypeDebugFlagsNone: print no messages
	TypeDebugFlagsNone TypeDebugFlags = 0
	// TypeDebugFlagsObjects: print messages about object bookkeeping
	TypeDebugFlagsObjects TypeDebugFlags = 1
	// TypeDebugFlagsSignals: print messages about signal emissions
	TypeDebugFlagsSignals TypeDebugFlags = 2
	// TypeDebugFlagsInstanceCount: keep a count of instances of each type
	TypeDebugFlagsInstanceCount TypeDebugFlags = 4
	// TypeDebugFlagsMask: mask covering all debug flags
	TypeDebugFlagsMask TypeDebugFlags = 7
)

// TypeFlags: bit masks used to check or determine characteristics of a type.
type TypeFlags int

const (
	// TypeFlagsAbstract indicates an abstract type. No instances can be created
	// for an abstract type
	TypeFlagsAbstract TypeFlags = 16
	// TypeFlagsValueAbstract indicates an abstract value type, i.e. a type that
	// introduces a value table, but can't be used for g_value_init()
	TypeFlagsValueAbstract TypeFlags = 32
)

// TypeFundamentalFlags: bit masks used to check or determine specific
// characteristics of a fundamental type.
type TypeFundamentalFlags int

const (
	// TypeFundamentalFlagsClassed indicates a classed type
	TypeFundamentalFlagsClassed TypeFundamentalFlags = 1
	// TypeFundamentalFlagsInstantiatable indicates an instantiable type
	// (implies classed)
	TypeFundamentalFlagsInstantiatable TypeFundamentalFlags = 2
	// TypeFundamentalFlagsDerivable indicates a flat derivable type
	TypeFundamentalFlagsDerivable TypeFundamentalFlags = 4
	// TypeFundamentalFlagsDeepDerivable indicates a deep derivable type
	// (implies derivable)
	TypeFundamentalFlagsDeepDerivable TypeFundamentalFlags = 8
)

// TypeAddClassPrivate registers a private class structure for a classed type;
// when the class is allocated, the private structures for the class and all of
// its parent types are allocated sequentially in the same memory block as the
// public structures, and are zero-filled.
//
// This function should be called in the type's get_type() function after the
// type is registered. The private structure can be retrieved using the
// G_TYPE_CLASS_GET_PRIVATE() macro.
func TypeAddClassPrivate(classType externglib.Type, privateSize uint) {
	var arg1 C.GType
	var arg2 C.gsize

	arg1 := C.GType(classType)
	arg2 = C.gsize(privateSize)

	C.g_type_add_class_private(arg1, arg2)
}

func TypeAddInstancePrivate(classType externglib.Type, privateSize uint) {
	var arg1 C.GType
	var arg2 C.gsize

	arg1 := C.GType(classType)
	arg2 = C.gsize(privateSize)

	C.g_type_add_instance_private(arg1, arg2)
}

// TypeAddInterfaceDynamic adds @interface_type to the dynamic
// @instantiable_type. The information contained in the Plugin structure pointed
// to by @plugin is used to manage the relationship.
func TypeAddInterfaceDynamic(instanceType externglib.Type, interfaceType externglib.Type, plugin TypePlugin) {
	var arg1 C.GType
	var arg2 C.GType
	var arg3 *C.GTypePlugin

	arg1 := C.GType(instanceType)
	arg2 := C.GType(interfaceType)
	arg3 = (*C.GTypePlugin)(unsafe.Pointer(plugin.Native()))

	C.g_type_add_interface_dynamic(arg1, arg2, arg3)
}

// TypeAddInterfaceStatic adds @interface_type to the static @instantiable_type.
// The information contained in the Info structure pointed to by @info is used
// to manage the relationship.
func TypeAddInterfaceStatic(instanceType externglib.Type, interfaceType externglib.Type, info *InterfaceInfo) {
	var arg1 C.GType
	var arg2 C.GType
	var arg3 *C.GInterfaceInfo

	arg1 := C.GType(instanceType)
	arg2 := C.GType(interfaceType)
	arg3 = (*C.GInterfaceInfo)(unsafe.Pointer(info.Native()))

	C.g_type_add_interface_static(arg1, arg2, arg3)
}

// TypeCheckInstance: private helper function to aid implementation of the
// G_TYPE_CHECK_INSTANCE() macro.
func TypeCheckInstance(instance *TypeInstance) bool {
	var arg1 *C.GTypeInstance

	arg1 = (*C.GTypeInstance)(unsafe.Pointer(instance.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_type_check_instance(arg1)

	if cret {
		ok = true
	}

	return ok
}

func TypeCheckInstanceCast(instance *TypeInstance, ifaceType externglib.Type) {
	var arg1 *C.GTypeInstance
	var arg2 C.GType

	arg1 = (*C.GTypeInstance)(unsafe.Pointer(instance.Native()))
	arg2 := C.GType(ifaceType)

	C.g_type_check_instance_cast(arg1, arg2)
}

func TypeCheckInstanceIsA(instance *TypeInstance, ifaceType externglib.Type) bool {
	var arg1 *C.GTypeInstance
	var arg2 C.GType

	arg1 = (*C.GTypeInstance)(unsafe.Pointer(instance.Native()))
	arg2 := C.GType(ifaceType)

	var cret C.gboolean
	var ok bool

	cret = C.g_type_check_instance_is_a(arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

func TypeCheckInstanceIsFundamentallyA(instance *TypeInstance, fundamentalType externglib.Type) bool {
	var arg1 *C.GTypeInstance
	var arg2 C.GType

	arg1 = (*C.GTypeInstance)(unsafe.Pointer(instance.Native()))
	arg2 := C.GType(fundamentalType)

	var cret C.gboolean
	var ok bool

	cret = C.g_type_check_instance_is_fundamentally_a(arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

func TypeCheckIsValueType(typ externglib.Type) bool {
	var arg1 C.GType

	arg1 := C.GType(typ)

	var cret C.gboolean
	var ok bool

	cret = C.g_type_check_is_value_type(arg1)

	if cret {
		ok = true
	}

	return ok
}

func TypeCheckValue(value *externglib.Value) bool {
	var arg1 *C.GValue

	arg1 = (*C.GValue)(value.GValue)

	var cret C.gboolean
	var ok bool

	cret = C.g_type_check_value(arg1)

	if cret {
		ok = true
	}

	return ok
}

func TypeCheckValueHolds(value *externglib.Value, typ externglib.Type) bool {
	var arg1 *C.GValue
	var arg2 C.GType

	arg1 = (*C.GValue)(value.GValue)
	arg2 := C.GType(typ)

	var cret C.gboolean
	var ok bool

	cret = C.g_type_check_value_holds(arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// TypeChildren: return a newly allocated and 0-terminated array of type IDs,
// listing the child types of @type.
func TypeChildren(typ externglib.Type) uint {
	var arg1 C.GType

	arg1 := C.GType(typ)

	var arg2 C.guint
	var nChildren uint

	C.g_type_children(arg1, &arg2)

	nChildren = uint(&arg2)

	return nChildren
}

func TypeClassAdjustPrivateOffset(gClass interface{}, privateSizeOrOffset int) {
	var arg1 C.gpointer
	var arg2 *C.gint

	arg1 = C.gpointer(gClass)
	arg2 = *C.gint(privateSizeOrOffset)

	C.g_type_class_adjust_private_offset(arg1, arg2)
}

// TypeClassPeek: this function is essentially the same as g_type_class_ref(),
// except that the classes reference count isn't incremented. As a consequence,
// this function may return nil if the class of the type passed in does not
// currently exist (hasn't been referenced before).
func TypeClassPeek(typ externglib.Type) {
	var arg1 C.GType

	arg1 := C.GType(typ)

	C.g_type_class_peek(arg1)
}

// TypeClassPeekStatic: a more efficient version of g_type_class_peek() which
// works only for static types.
func TypeClassPeekStatic(typ externglib.Type) {
	var arg1 C.GType

	arg1 := C.GType(typ)

	C.g_type_class_peek_static(arg1)
}

// TypeClassRef increments the reference count of the class structure belonging
// to @type. This function will demand-create the class if it doesn't exist
// already.
func TypeClassRef(typ externglib.Type) {
	var arg1 C.GType

	arg1 := C.GType(typ)

	C.g_type_class_ref(arg1)
}

// TypeCreateInstance creates and initializes an instance of @type if @type is
// valid and can be instantiated. The type system only performs basic allocation
// and structure setups for instances: actual instance creation should happen
// through functions supplied by the type's fundamental type implementation. So
// use of g_type_create_instance() is reserved for implementers of fundamental
// types only. E.g. instances of the #GObject hierarchy should be created via
// g_object_new() and never directly through g_type_create_instance() which
// doesn't handle things like singleton objects or object construction.
//
// The extended members of the returned instance are guaranteed to be filled
// with zeros.
//
// Note: Do not use this function, unless you're implementing a fundamental
// type. Also language bindings should not use this function, but g_object_new()
// instead.
func TypeCreateInstance(typ externglib.Type) {
	var arg1 C.GType

	arg1 := C.GType(typ)

	C.g_type_create_instance(arg1)
}

// TypeDefaultInterfacePeek: if the interface type @g_type is currently in use,
// returns its default interface vtable.
func TypeDefaultInterfacePeek(gType externglib.Type) {
	var arg1 C.GType

	arg1 := C.GType(gType)

	C.g_type_default_interface_peek(arg1)
}

// TypeDefaultInterfaceRef increments the reference count for the interface type
// @g_type, and returns the default interface vtable for the type.
//
// If the type is not currently in use, then the default vtable for the type
// will be created and initialized by calling the base interface init and
// default vtable init functions for the type (the @base_init and @class_init
// members of Info). Calling g_type_default_interface_ref() is useful when you
// want to make sure that signals and properties for an interface have been
// installed.
func TypeDefaultInterfaceRef(gType externglib.Type) {
	var arg1 C.GType

	arg1 := C.GType(gType)

	C.g_type_default_interface_ref(arg1)
}

// TypeDepth returns the length of the ancestry of the passed in type. This
// includes the type itself, so that e.g. a fundamental type has depth 1.
func TypeDepth(typ externglib.Type) {
	var arg1 C.GType

	arg1 := C.GType(typ)

	C.g_type_depth(arg1)
}

// TypeEnsure ensures that the indicated @type has been registered with the type
// system, and its _class_init() method has been run.
//
// In theory, simply calling the type's _get_type() method (or using the
// corresponding macro) is supposed take care of this. However, _get_type()
// methods are often marked G_GNUC_CONST for performance reasons, even though
// this is technically incorrect (since G_GNUC_CONST requires that the function
// not have side effects, which _get_type() methods do on the first call). As a
// result, if you write a bare call to a _get_type() macro, it may get optimized
// out by the compiler. Using g_type_ensure() guarantees that the type's
// _get_type() method is called.
func TypeEnsure(typ externglib.Type) {
	var arg1 C.GType

	arg1 := C.GType(typ)

	C.g_type_ensure(arg1)
}

// TypeFreeInstance frees an instance of a type, returning it to the instance
// pool for the type, if there is one.
//
// Like g_type_create_instance(), this function is reserved for implementors of
// fundamental types.
func TypeFreeInstance(instance *TypeInstance) {
	var arg1 *C.GTypeInstance

	arg1 = (*C.GTypeInstance)(unsafe.Pointer(instance.Native()))

	C.g_type_free_instance(arg1)
}

// TypeFromName: look up the type ID from a given type name, returning 0 if no
// type has been registered under this name (this is the preferred method to
// find out by name whether a specific type has been registered yet).
func TypeFromName(name string) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	C.g_type_from_name(arg1)
}

// TypeFundamental: internal function, used to extract the fundamental type ID
// portion. Use G_TYPE_FUNDAMENTAL() instead.
func TypeFundamental(typeID externglib.Type) {
	var arg1 C.GType

	arg1 := C.GType(typeID)

	C.g_type_fundamental(arg1)
}

// TypeFundamentalNext returns the next free fundamental type id which can be
// used to register a new fundamental type with g_type_register_fundamental().
// The returned type ID represents the highest currently registered fundamental
// type identifier.
func TypeFundamentalNext() {
	C.g_type_fundamental_next()
}

// TypeGetInstanceCount returns the number of instances allocated of the
// particular type; this is only available if GLib is built with debugging
// support and the instance_count debug flag is set (by setting the
// GOBJECT_DEBUG variable to include instance-count).
func TypeGetInstanceCount(typ externglib.Type) {
	var arg1 C.GType

	arg1 := C.GType(typ)

	C.g_type_get_instance_count(arg1)
}

// TypeGetPlugin returns the Plugin structure for @type.
func TypeGetPlugin(typ externglib.Type) {
	var arg1 C.GType

	arg1 := C.GType(typ)

	C.g_type_get_plugin(arg1)
}

// TypeGetTypeRegistrationSerial returns an opaque serial number that represents
// the state of the set of registered types. Any time a type is registered this
// serial changes, which means you can cache information based on type lookups
// (such as g_type_from_name()) and know if the cache is still valid at a later
// time by comparing the current serial with the one at the type lookup.
func TypeGetTypeRegistrationSerial() {
	C.g_type_get_type_registration_serial()
}

// TypeInit: this function used to initialise the type system. Since GLib 2.36,
// the type system is initialised automatically and this function does nothing.
func TypeInit() {
	C.g_type_init()
}

// TypeInitWithDebugFlags: this function used to initialise the type system with
// debugging flags. Since GLib 2.36, the type system is initialised
// automatically and this function does nothing.
//
// If you need to enable debugging features, use the GOBJECT_DEBUG environment
// variable.
func TypeInitWithDebugFlags(debugFlags TypeDebugFlags) {
	var arg1 C.GTypeDebugFlags

	arg1 = (C.GTypeDebugFlags)(debugFlags)

	C.g_type_init_with_debug_flags(arg1)
}

// TypeInterfaceAddPrerequisite adds @prerequisite_type to the list of
// prerequisites of @interface_type. This means that any type implementing
// @interface_type must also implement @prerequisite_type. Prerequisites can be
// thought of as an alternative to interface derivation (which GType doesn't
// support). An interface can have at most one instantiatable prerequisite type.
func TypeInterfaceAddPrerequisite(interfaceType externglib.Type, prerequisiteType externglib.Type) {
	var arg1 C.GType
	var arg2 C.GType

	arg1 := C.GType(interfaceType)
	arg2 := C.GType(prerequisiteType)

	C.g_type_interface_add_prerequisite(arg1, arg2)
}

// TypeInterfaceGetPlugin returns the Plugin structure for the dynamic interface
// @interface_type which has been added to @instance_type, or nil if
// @interface_type has not been added to @instance_type or does not have a
// Plugin structure. See g_type_add_interface_dynamic().
func TypeInterfaceGetPlugin(instanceType externglib.Type, interfaceType externglib.Type) {
	var arg1 C.GType
	var arg2 C.GType

	arg1 := C.GType(instanceType)
	arg2 := C.GType(interfaceType)

	C.g_type_interface_get_plugin(arg1, arg2)
}

// TypeInterfacePrerequisites returns the prerequisites of an interfaces type.
func TypeInterfacePrerequisites(interfaceType externglib.Type) uint {
	var arg1 C.GType

	arg1 := C.GType(interfaceType)

	var arg2 C.guint
	var nPrerequisites uint

	C.g_type_interface_prerequisites(arg1, &arg2)

	nPrerequisites = uint(&arg2)

	return nPrerequisites
}

// TypeInterfaces: return a newly allocated and 0-terminated array of type IDs,
// listing the interface types that @type conforms to.
func TypeInterfaces(typ externglib.Type) uint {
	var arg1 C.GType

	arg1 := C.GType(typ)

	var arg2 C.guint
	var nInterfaces uint

	C.g_type_interfaces(arg1, &arg2)

	nInterfaces = uint(&arg2)

	return nInterfaces
}

// TypeIsA: if @is_a_type is a derivable type, check whether @type is a
// descendant of @is_a_type. If @is_a_type is an interface, check whether @type
// conforms to it.
func TypeIsA(typ externglib.Type, isAType externglib.Type) bool {
	var arg1 C.GType
	var arg2 C.GType

	arg1 := C.GType(typ)
	arg2 := C.GType(isAType)

	var cret C.gboolean
	var ok bool

	cret = C.g_type_is_a(arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// TypeName: get the unique name that is assigned to a type ID. Note that this
// function (like all other GType API) cannot cope with invalid type IDs.
// G_TYPE_INVALID may be passed to this function, as may be any other validly
// registered type ID, but randomized type IDs should not be passed in and will
// most likely lead to a crash.
func TypeName(typ externglib.Type) {
	var arg1 C.GType

	arg1 := C.GType(typ)

	C.g_type_name(arg1)
}

func TypeNameFromInstance(instance *TypeInstance) {
	var arg1 *C.GTypeInstance

	arg1 = (*C.GTypeInstance)(unsafe.Pointer(instance.Native()))

	C.g_type_name_from_instance(arg1)
}

// TypeNextBase: given a @leaf_type and a @root_type which is contained in its
// anchestry, return the type that @root_type is the immediate parent of. In
// other words, this function determines the type that is derived directly from
// @root_type which is also a base class of @leaf_type. Given a root type and a
// leaf type, this function can be used to determine the types and order in
// which the leaf type is descended from the root type.
func TypeNextBase(leafType externglib.Type, rootType externglib.Type) {
	var arg1 C.GType
	var arg2 C.GType

	arg1 := C.GType(leafType)
	arg2 := C.GType(rootType)

	C.g_type_next_base(arg1, arg2)
}

// TypeParent: return the direct parent type of the passed in type. If the
// passed in type has no parent, i.e. is a fundamental type, 0 is returned.
func TypeParent(typ externglib.Type) {
	var arg1 C.GType

	arg1 := C.GType(typ)

	C.g_type_parent(arg1)
}

// TypeQname: get the corresponding quark of the type IDs name.
func TypeQname(typ externglib.Type) {
	var arg1 C.GType

	arg1 := C.GType(typ)

	C.g_type_qname(arg1)
}

// TypeQuery queries the type system for information about a specific type. This
// function will fill in a user-provided structure to hold type-specific
// information. If an invalid #GType is passed in, the @type member of the Query
// is 0. All members filled into the Query structure should be considered
// constant and have to be left untouched.
func TypeQuery(typ externglib.Type) *TypeQuery {
	var arg1 C.GType

	arg1 := C.GType(typ)

	var arg2 C.GTypeQuery
	var query *TypeQuery

	C.g_type_query(arg1, &arg2)

	query = WrapTypeQuery(unsafe.Pointer(&arg2))

	return query
}

// TypeRegisterDynamic registers @type_name as the name of a new dynamic type
// derived from @parent_type. The type system uses the information contained in
// the Plugin structure pointed to by @plugin to manage the type and its
// instances (if not abstract). The value of @flags determines the nature (e.g.
// abstract or not) of the type.
func TypeRegisterDynamic(parentType externglib.Type, typeName string, plugin TypePlugin, flags TypeFlags) {
	var arg1 C.GType
	var arg2 *C.gchar
	var arg3 *C.GTypePlugin
	var arg4 C.GTypeFlags

	arg1 := C.GType(parentType)
	arg2 = (*C.gchar)(C.CString(typeName))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.GTypePlugin)(unsafe.Pointer(plugin.Native()))
	arg4 = (C.GTypeFlags)(flags)

	C.g_type_register_dynamic(arg1, arg2, arg3, arg4)
}

// TypeRegisterFundamental registers @type_id as the predefined identifier and
// @type_name as the name of a fundamental type. If @type_id is already
// registered, or a type named @type_name is already registered, the behaviour
// is undefined. The type system uses the information contained in the Info
// structure pointed to by @info and the FundamentalInfo structure pointed to by
// @finfo to manage the type and its instances. The value of @flags determines
// additional characteristics of the fundamental type.
func TypeRegisterFundamental(typeID externglib.Type, typeName string, info *TypeInfo, finfo *TypeFundamentalInfo, flags TypeFlags) {
	var arg1 C.GType
	var arg2 *C.gchar
	var arg3 *C.GTypeInfo
	var arg4 *C.GTypeFundamentalInfo
	var arg5 C.GTypeFlags

	arg1 := C.GType(typeID)
	arg2 = (*C.gchar)(C.CString(typeName))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.GTypeInfo)(unsafe.Pointer(info.Native()))
	arg4 = (*C.GTypeFundamentalInfo)(unsafe.Pointer(finfo.Native()))
	arg5 = (C.GTypeFlags)(flags)

	C.g_type_register_fundamental(arg1, arg2, arg3, arg4, arg5)
}

// TypeRegisterStatic registers @type_name as the name of a new static type
// derived from @parent_type. The type system uses the information contained in
// the Info structure pointed to by @info to manage the type and its instances
// (if not abstract). The value of @flags determines the nature (e.g. abstract
// or not) of the type.
func TypeRegisterStatic(parentType externglib.Type, typeName string, info *TypeInfo, flags TypeFlags) {
	var arg1 C.GType
	var arg2 *C.gchar
	var arg3 *C.GTypeInfo
	var arg4 C.GTypeFlags

	arg1 := C.GType(parentType)
	arg2 = (*C.gchar)(C.CString(typeName))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.GTypeInfo)(unsafe.Pointer(info.Native()))
	arg4 = (C.GTypeFlags)(flags)

	C.g_type_register_static(arg1, arg2, arg3, arg4)
}

func TypeTestFlags(typ externglib.Type, flags uint) bool {
	var arg1 C.GType
	var arg2 C.guint

	arg1 := C.GType(typ)
	arg2 = C.guint(flags)

	var cret C.gboolean
	var ok bool

	cret = C.g_type_test_flags(arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// TypeValueTablePeek returns the location of the ValueTable associated with
// @type.
//
// Note that this function should only be used from source code that implements
// or has internal knowledge of the implementation of @type.
func TypeValueTablePeek(typ externglib.Type) {
	var arg1 C.GType

	arg1 := C.GType(typ)

	C.g_type_value_table_peek(arg1)
}

// InterfaceInfo: a structure that provides information to the type system which
// is used specifically for managing interface types.
type InterfaceInfo struct {
	native C.GInterfaceInfo
}

// WrapInterfaceInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapInterfaceInfo(ptr unsafe.Pointer) *InterfaceInfo {
	if ptr == nil {
		return nil
	}

	return (*InterfaceInfo)(ptr)
}

func marshalInterfaceInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapInterfaceInfo(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (i *InterfaceInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&i.native)
}

// InterfaceData gets the field inside the struct.
func (i *InterfaceInfo) InterfaceData() interface{} {
	var v interface{}
	v = interface{}(i.native.interface_data)
	return v
}

// TypeFundamentalInfo: a structure that provides information to the type system
// which is used specifically for managing fundamental types.
type TypeFundamentalInfo struct {
	native C.GTypeFundamentalInfo
}

// WrapTypeFundamentalInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTypeFundamentalInfo(ptr unsafe.Pointer) *TypeFundamentalInfo {
	if ptr == nil {
		return nil
	}

	return (*TypeFundamentalInfo)(ptr)
}

func marshalTypeFundamentalInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTypeFundamentalInfo(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TypeFundamentalInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// TypeFlags gets the field inside the struct.
func (t *TypeFundamentalInfo) TypeFlags() TypeFundamentalFlags {
	var v TypeFundamentalFlags
	v = TypeFundamentalFlags(t.native.type_flags)
	return v
}

// TypeInfo: this structure is used to provide the type system with the
// information required to initialize and destruct (finalize) a type's class and
// its instances.
//
// The initialized structure is passed to the g_type_register_static() function
// (or is copied into the provided Info structure in the
// g_type_plugin_complete_type_info()). The type system will perform a deep copy
// of this structure, so its memory does not need to be persistent across
// invocation of g_type_register_static().
type TypeInfo struct {
	native C.GTypeInfo
}

// WrapTypeInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTypeInfo(ptr unsafe.Pointer) *TypeInfo {
	if ptr == nil {
		return nil
	}

	return (*TypeInfo)(ptr)
}

func marshalTypeInfo(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTypeInfo(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TypeInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// ClassSize gets the field inside the struct.
func (t *TypeInfo) ClassSize() uint16 {
	var v uint16
	v = uint16(t.native.class_size)
	return v
}

// ClassData gets the field inside the struct.
func (t *TypeInfo) ClassData() interface{} {
	var v interface{}
	v = interface{}(t.native.class_data)
	return v
}

// InstanceSize gets the field inside the struct.
func (t *TypeInfo) InstanceSize() uint16 {
	var v uint16
	v = uint16(t.native.instance_size)
	return v
}

// NPreallocs gets the field inside the struct.
func (t *TypeInfo) NPreallocs() uint16 {
	var v uint16
	v = uint16(t.native.n_preallocs)
	return v
}

// TypeInstance: an opaque structure used as the base of all type instances.
type TypeInstance struct {
	native C.GTypeInstance
}

// WrapTypeInstance wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTypeInstance(ptr unsafe.Pointer) *TypeInstance {
	if ptr == nil {
		return nil
	}

	return (*TypeInstance)(ptr)
}

func marshalTypeInstance(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTypeInstance(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TypeInstance) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

func (i *TypeInstance) Private(i *TypeInstance, privateType externglib.Type) {
	var arg0 *C.GTypeInstance
	var arg1 C.GType

	arg0 = (*C.GTypeInstance)(unsafe.Pointer(i.Native()))
	arg1 := C.GType(privateType)

	C.g_type_instance_get_private(arg0, arg1)
}

// TypeQuery: a structure holding information for a specific type. It is filled
// in by the g_type_query() function.
type TypeQuery struct {
	native C.GTypeQuery
}

// WrapTypeQuery wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapTypeQuery(ptr unsafe.Pointer) *TypeQuery {
	if ptr == nil {
		return nil
	}

	return (*TypeQuery)(ptr)
}

func marshalTypeQuery(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapTypeQuery(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *TypeQuery) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// Type gets the field inside the struct.
func (t *TypeQuery) Type() externglib.Type {
	var v externglib.Type
	v = externglib.Type(t.native._type)
	return v
}

// TypeName gets the field inside the struct.
func (t *TypeQuery) TypeName() string {
	var v string
	v = C.GoString(t.native.type_name)
	return v
}

// ClassSize gets the field inside the struct.
func (t *TypeQuery) ClassSize() uint {
	var v uint
	v = uint(t.native.class_size)
	return v
}

// InstanceSize gets the field inside the struct.
func (t *TypeQuery) InstanceSize() uint {
	var v uint
	v = uint(t.native.instance_size)
	return v
}