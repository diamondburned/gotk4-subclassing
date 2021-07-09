// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_simple_action_group_get_type()), F: marshalSimpleActionGroup},
	})
}

// SimpleActionGroup is a hash table filled with #GAction objects, implementing
// the Group and Map interfaces.
type SimpleActionGroup interface {
	gextras.Objector

	// AddEntries: convenience function for creating multiple Action instances
	// and adding them to the action group.
	//
	// Deprecated: since version 2.38.
	AddEntries(entries []ActionEntry, userData interface{})
	// Insert adds an action to the action group.
	//
	// If the action group already contains an action with the same name as
	// @action then the old action is dropped from the group.
	//
	// The action group takes its own reference on @action.
	//
	// Deprecated: since version 2.38.
	Insert(action Action)
	// Lookup looks up the action with the name @action_name in the group.
	//
	// If no such action exists, returns nil.
	//
	// Deprecated: since version 2.38.
	Lookup(actionName string) *ActionInterface
	// Remove removes the named action from the action group.
	//
	// If no action of this name is in the group then nothing happens.
	//
	// Deprecated: since version 2.38.
	Remove(actionName string)
}

// SimpleActionGroupClass implements the SimpleActionGroup interface.
type SimpleActionGroupClass struct {
	*externglib.Object
	ActionGroupInterface
	ActionMapInterface
}

var _ SimpleActionGroup = (*SimpleActionGroupClass)(nil)

func wrapSimpleActionGroup(obj *externglib.Object) SimpleActionGroup {
	return &SimpleActionGroupClass{
		Object: obj,
		ActionGroupInterface: ActionGroupInterface{
			Object: obj,
		},
		ActionMapInterface: ActionMapInterface{
			Object: obj,
		},
	}
}

func marshalSimpleActionGroup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapSimpleActionGroup(obj), nil
}

// NewSimpleActionGroup creates a new, empty, ActionGroup.
func NewSimpleActionGroup() *SimpleActionGroupClass {
	var _cret *C.GSimpleActionGroup // in

	_cret = C.g_simple_action_group_new()

	var _simpleActionGroup *SimpleActionGroupClass // out

	_simpleActionGroup = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*SimpleActionGroupClass)

	return _simpleActionGroup
}

// AddEntries: convenience function for creating multiple Action instances and
// adding them to the action group.
//
// Deprecated: since version 2.38.
func (s *SimpleActionGroupClass) AddEntries(entries []ActionEntry, userData interface{}) {
	var _arg0 *C.GSimpleActionGroup // out
	var _arg1 *C.GActionEntry
	var _arg2 C.gint
	var _arg3 C.gpointer // out

	_arg0 = (*C.GSimpleActionGroup)(unsafe.Pointer(s.Native()))
	_arg2 = C.gint(len(entries))
	_arg1 = (*C.GActionEntry)(unsafe.Pointer(&entries[0]))
	_arg3 = (C.gpointer)(box.Assign(userData))

	C.g_simple_action_group_add_entries(_arg0, _arg1, _arg2, _arg3)
}

// Insert adds an action to the action group.
//
// If the action group already contains an action with the same name as @action
// then the old action is dropped from the group.
//
// The action group takes its own reference on @action.
//
// Deprecated: since version 2.38.
func (s *SimpleActionGroupClass) Insert(action Action) {
	var _arg0 *C.GSimpleActionGroup // out
	var _arg1 *C.GAction            // out

	_arg0 = (*C.GSimpleActionGroup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GAction)(unsafe.Pointer(action.Native()))

	C.g_simple_action_group_insert(_arg0, _arg1)
}

// Lookup looks up the action with the name @action_name in the group.
//
// If no such action exists, returns nil.
//
// Deprecated: since version 2.38.
func (s *SimpleActionGroupClass) Lookup(actionName string) *ActionInterface {
	var _arg0 *C.GSimpleActionGroup // out
	var _arg1 *C.gchar              // out
	var _cret *C.GAction            // in

	_arg0 = (*C.GSimpleActionGroup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_simple_action_group_lookup(_arg0, _arg1)

	var _action *ActionInterface // out

	_action = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ActionInterface)

	return _action
}

// Remove removes the named action from the action group.
//
// If no action of this name is in the group then nothing happens.
//
// Deprecated: since version 2.38.
func (s *SimpleActionGroupClass) Remove(actionName string) {
	var _arg0 *C.GSimpleActionGroup // out
	var _arg1 *C.gchar              // out

	_arg0 = (*C.GSimpleActionGroup)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_simple_action_group_remove(_arg0, _arg1)
}
