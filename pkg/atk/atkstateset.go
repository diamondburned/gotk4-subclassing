// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_state_set_get_type()), F: marshalStateSet},
	})
}

// StateSet is a read-only representation of the full set of States that apply
// to an object at a given time. This set is not meant to be modified, but
// rather created when #atk_object_ref_state_set() is called.
type StateSet interface {
	gextras.Objector

	// AddStateStateSet adds the state of the specified type to the state set if
	// it is not already present.
	//
	// Note that because an StateSet is a read-only object, this method should
	// be used to add a state to a newly-created set which will then be returned
	// by #atk_object_ref_state_set. It should not be used to modify the
	// existing state of an object. See also #atk_object_notify_state_change.
	AddStateStateSet(typ StateType) bool
	// AddStatesStateSet adds the states of the specified types to the state
	// set.
	//
	// Note that because an StateSet is a read-only object, this method should
	// be used to add states to a newly-created set which will then be returned
	// by #atk_object_ref_state_set. It should not be used to modify the
	// existing state of an object. See also #atk_object_notify_state_change.
	AddStatesStateSet(types []StateType)
	// AndSetsStateSet constructs the intersection of the two sets, returning
	// nil if the intersection is empty.
	AndSetsStateSet(compareSet StateSet) StateSet
	// ClearStatesStateSet removes all states from the state set.
	ClearStatesStateSet()
	// ContainsStateStateSet checks whether the state for the specified type is
	// in the specified set.
	ContainsStateStateSet(typ StateType) bool
	// ContainsStatesStateSet checks whether the states for all the specified
	// types are in the specified set.
	ContainsStatesStateSet(types []StateType) bool
	// IsEmptyStateSet checks whether the state set is empty, i.e. has no states
	// set.
	IsEmptyStateSet() bool
	// OrSetsStateSet constructs the union of the two sets.
	OrSetsStateSet(compareSet StateSet) StateSet
	// RemoveStateStateSet removes the state for the specified type from the
	// state set.
	//
	// Note that because an StateSet is a read-only object, this method should
	// be used to remove a state to a newly-created set which will then be
	// returned by #atk_object_ref_state_set. It should not be used to modify
	// the existing state of an object. See also
	// #atk_object_notify_state_change.
	RemoveStateStateSet(typ StateType) bool
	// XorSetsStateSet constructs the exclusive-or of the two sets, returning
	// nil is empty. The set returned by this operation contains the states in
	// exactly one of the two sets.
	XorSetsStateSet(compareSet StateSet) StateSet
}

// stateSet implements the StateSet class.
type stateSet struct {
	gextras.Objector
}

// WrapStateSet wraps a GObject to the right type. It is
// primarily used internally.
func WrapStateSet(obj *externglib.Object) StateSet {
	return stateSet{
		Objector: obj,
	}
}

func marshalStateSet(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStateSet(obj), nil
}

// NewStateSet creates a new empty state set.
func NewStateSet() StateSet {
	var _cret *C.AtkStateSet // in

	_cret = C.atk_state_set_new()

	var _stateSet StateSet // out

	_stateSet = WrapStateSet(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _stateSet
}

func (s stateSet) AddStateStateSet(typ StateType) bool {
	var _arg0 *C.AtkStateSet // out
	var _arg1 C.AtkStateType // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(s.Native()))
	_arg1 = C.AtkStateType(typ)

	_cret = C.atk_state_set_add_state(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s stateSet) AddStatesStateSet(types []StateType) {
	var _arg0 *C.AtkStateSet // out
	var _arg1 *C.AtkStateType
	var _arg2 C.gint

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(s.Native()))
	_arg2 = C.gint(len(types))
	_arg1 = (*C.AtkStateType)(C.malloc(C.ulong(len(types)) * C.ulong(C.sizeof_AtkStateType)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice(_arg1, len(types))
		for i := range types {
			out[i] = C.AtkStateType(types[i])
		}
	}

	C.atk_state_set_add_states(_arg0, _arg1, _arg2)
}

func (s stateSet) AndSetsStateSet(compareSet StateSet) StateSet {
	var _arg0 *C.AtkStateSet // out
	var _arg1 *C.AtkStateSet // out
	var _cret *C.AtkStateSet // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.AtkStateSet)(unsafe.Pointer(compareSet.Native()))

	_cret = C.atk_state_set_and_sets(_arg0, _arg1)

	var _stateSet StateSet // out

	_stateSet = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(StateSet)

	return _stateSet
}

func (s stateSet) ClearStatesStateSet() {
	var _arg0 *C.AtkStateSet // out

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(s.Native()))

	C.atk_state_set_clear_states(_arg0)
}

func (s stateSet) ContainsStateStateSet(typ StateType) bool {
	var _arg0 *C.AtkStateSet // out
	var _arg1 C.AtkStateType // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(s.Native()))
	_arg1 = C.AtkStateType(typ)

	_cret = C.atk_state_set_contains_state(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s stateSet) ContainsStatesStateSet(types []StateType) bool {
	var _arg0 *C.AtkStateSet // out
	var _arg1 *C.AtkStateType
	var _arg2 C.gint
	var _cret C.gboolean // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(s.Native()))
	_arg2 = C.gint(len(types))
	_arg1 = (*C.AtkStateType)(C.malloc(C.ulong(len(types)) * C.ulong(C.sizeof_AtkStateType)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice(_arg1, len(types))
		for i := range types {
			out[i] = C.AtkStateType(types[i])
		}
	}

	_cret = C.atk_state_set_contains_states(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s stateSet) IsEmptyStateSet() bool {
	var _arg0 *C.AtkStateSet // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(s.Native()))

	_cret = C.atk_state_set_is_empty(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s stateSet) OrSetsStateSet(compareSet StateSet) StateSet {
	var _arg0 *C.AtkStateSet // out
	var _arg1 *C.AtkStateSet // out
	var _cret *C.AtkStateSet // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.AtkStateSet)(unsafe.Pointer(compareSet.Native()))

	_cret = C.atk_state_set_or_sets(_arg0, _arg1)

	var _stateSet StateSet // out

	_stateSet = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(StateSet)

	return _stateSet
}

func (s stateSet) RemoveStateStateSet(typ StateType) bool {
	var _arg0 *C.AtkStateSet // out
	var _arg1 C.AtkStateType // out
	var _cret C.gboolean     // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(s.Native()))
	_arg1 = C.AtkStateType(typ)

	_cret = C.atk_state_set_remove_state(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (s stateSet) XorSetsStateSet(compareSet StateSet) StateSet {
	var _arg0 *C.AtkStateSet // out
	var _arg1 *C.AtkStateSet // out
	var _cret *C.AtkStateSet // in

	_arg0 = (*C.AtkStateSet)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.AtkStateSet)(unsafe.Pointer(compareSet.Native()))

	_cret = C.atk_state_set_xor_sets(_arg0, _arg1)

	var _stateSet StateSet // out

	_stateSet = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(StateSet)

	return _stateSet
}