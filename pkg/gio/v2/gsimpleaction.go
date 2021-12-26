// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// NewSimpleAction creates a new action.
//
// The created action is stateless. See g_simple_action_new_stateful() to create
// an action that has state.
//
// The function takes the following parameters:
//
//    - name of the action.
//    - parameterType (optional): type of parameter that will be passed to
//      handlers for the Action::activate signal, or NULL for no parameter.
//
// The function returns the following values:
//
//    - simpleAction: new Action.
//
func NewSimpleAction(name string, parameterType *glib.VariantType) *SimpleAction {
	var _arg1 *C.gchar         // out
	var _arg2 *C.GVariantType  // out
	var _cret *C.GSimpleAction // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	if parameterType != nil {
		_arg2 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(parameterType)))
	}

	_cret = C.g_simple_action_new(_arg1, _arg2)
	runtime.KeepAlive(name)
	runtime.KeepAlive(parameterType)

	var _simpleAction *SimpleAction // out

	_simpleAction = wrapSimpleAction(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _simpleAction
}

// NewSimpleActionStateful creates a new stateful action.
//
// All future state values must have the same Type as the initial state.
//
// If the state #GVariant is floating, it is consumed.
//
// The function takes the following parameters:
//
//    - name of the action.
//    - parameterType (optional): type of the parameter that will be passed to
//      handlers for the Action::activate signal, or NULL for no parameter.
//    - state: initial state of the action.
//
// The function returns the following values:
//
//    - simpleAction: new Action.
//
func NewSimpleActionStateful(name string, parameterType *glib.VariantType, state *glib.Variant) *SimpleAction {
	var _arg1 *C.gchar         // out
	var _arg2 *C.GVariantType  // out
	var _arg3 *C.GVariant      // out
	var _cret *C.GSimpleAction // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	if parameterType != nil {
		_arg2 = (*C.GVariantType)(gextras.StructNative(unsafe.Pointer(parameterType)))
	}
	_arg3 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(state)))

	_cret = C.g_simple_action_new_stateful(_arg1, _arg2, _arg3)
	runtime.KeepAlive(name)
	runtime.KeepAlive(parameterType)
	runtime.KeepAlive(state)

	var _simpleAction *SimpleAction // out

	_simpleAction = wrapSimpleAction(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _simpleAction
}

// SetEnabled sets the action as enabled or not.
//
// An action must be enabled in order to be activated or in order to have its
// state changed from outside callers.
//
// This should only be called by the implementor of the action. Users of the
// action should not attempt to modify its enabled flag.
//
// The function takes the following parameters:
//
//    - enabled: whether the action is enabled.
//
func (simple *SimpleAction) SetEnabled(enabled bool) {
	var _arg0 *C.GSimpleAction // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GSimpleAction)(unsafe.Pointer(simple.Native()))
	if enabled {
		_arg1 = C.TRUE
	}

	C.g_simple_action_set_enabled(_arg0, _arg1)
	runtime.KeepAlive(simple)
	runtime.KeepAlive(enabled)
}

// SetState sets the state of the action.
//
// This directly updates the 'state' property to the given value.
//
// This should only be called by the implementor of the action. Users of the
// action should not attempt to directly modify the 'state' property. Instead,
// they should call g_action_change_state() to request the change.
//
// If the value GVariant is floating, it is consumed.
//
// The function takes the following parameters:
//
//    - value: new #GVariant for the state.
//
func (simple *SimpleAction) SetState(value *glib.Variant) {
	var _arg0 *C.GSimpleAction // out
	var _arg1 *C.GVariant      // out

	_arg0 = (*C.GSimpleAction)(unsafe.Pointer(simple.Native()))
	_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(value)))

	C.g_simple_action_set_state(_arg0, _arg1)
	runtime.KeepAlive(simple)
	runtime.KeepAlive(value)
}

// SetStateHint sets the state hint for the action.
//
// See g_action_get_state_hint() for more information about action state hints.
//
// The function takes the following parameters:
//
//    - stateHint (optional) representing the state hint.
//
func (simple *SimpleAction) SetStateHint(stateHint *glib.Variant) {
	var _arg0 *C.GSimpleAction // out
	var _arg1 *C.GVariant      // out

	_arg0 = (*C.GSimpleAction)(unsafe.Pointer(simple.Native()))
	if stateHint != nil {
		_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(stateHint)))
	}

	C.g_simple_action_set_state_hint(_arg0, _arg1)
	runtime.KeepAlive(simple)
	runtime.KeepAlive(stateHint)
}
