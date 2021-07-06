// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

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
		{T: externglib.Type(C.g_mount_operation_get_type()), F: marshalMountOperation},
	})
}

// MountOperationOverrider contains methods that are overridable .
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type MountOperationOverrider interface {
	Aborted()
	AskPassword(message string, defaultUser string, defaultDomain string, flags AskPasswordFlags)
	// AskQuestion: virtual implementation of Operation::ask-question.
	AskQuestion(message string, choices []string)
	// Reply emits the Operation::reply signal.
	Reply(result MountOperationResult)
	ShowUnmountProgress(message string, timeLeft int64, bytesLeft int64)
}

// MountOperation provides a mechanism for interacting with the user. It can be
// used for authenticating mountable operations, such as loop mounting files,
// hard drive partitions or server locations. It can also be used to ask the
// user questions or show a list of applications preventing unmount or eject
// operations from completing.
//
// Note that Operation is used for more than just #GMount objects – for example
// it is also used in g_drive_start() and g_drive_stop().
//
// Users should instantiate a subclass of this that implements all the various
// callbacks to show the required dialogs, such as MountOperation. If no user
// interaction is desired (for example when automounting filesystems at login
// time), usually nil can be passed, see each method taking a Operation for
// details.
//
// The term ‘TCRYPT’ is used to mean ‘compatible with TrueCrypt and VeraCrypt’.
// TrueCrypt (https://en.wikipedia.org/wiki/TrueCrypt) is a discontinued system
// for encrypting file containers, partitions or whole disks, typically used
// with Windows. VeraCrypt (https://www.veracrypt.fr/) is a maintained fork of
// TrueCrypt with various improvements and auditing fixes.
type MountOperation interface {
	gextras.Objector

	// Anonymous: check to see whether the mount operation is being used for an
	// anonymous user.
	Anonymous() bool
	// Choice gets a choice from the mount operation.
	Choice() int
	// Domain gets the domain of the mount operation.
	Domain() string
	// IsTcryptHiddenVolume: check to see whether the mount operation is being
	// used for a TCRYPT hidden volume.
	IsTcryptHiddenVolume() bool
	// IsTcryptSystemVolume: check to see whether the mount operation is being
	// used for a TCRYPT system volume.
	IsTcryptSystemVolume() bool
	// Password gets a password from the mount operation.
	Password() string
	// PasswordSave gets the state of saving passwords for the mount operation.
	PasswordSave() PasswordSave
	// Pim gets a PIM from the mount operation.
	Pim() uint
	// Username: get the user name from the mount operation.
	Username() string
	// Reply emits the Operation::reply signal.
	Reply(result MountOperationResult)
	// SetAnonymous sets the mount operation to use an anonymous user if
	// @anonymous is true.
	SetAnonymous(anonymous bool)
	// SetChoice sets a default choice for the mount operation.
	SetChoice(choice int)
	// SetDomain sets the mount operation's domain.
	SetDomain(domain string)
	// SetIsTcryptHiddenVolume sets the mount operation to use a hidden volume
	// if @hidden_volume is true.
	SetIsTcryptHiddenVolume(hiddenVolume bool)
	// SetIsTcryptSystemVolume sets the mount operation to use a system volume
	// if @system_volume is true.
	SetIsTcryptSystemVolume(systemVolume bool)
	// SetPassword sets the mount operation's password to @password.
	SetPassword(password string)
	// SetPasswordSave sets the state of saving passwords for the mount
	// operation.
	SetPasswordSave(save PasswordSave)
	// SetPim sets the mount operation's PIM to @pim.
	SetPim(pim uint)
	// SetUsername sets the user name within @op to @username.
	SetUsername(username string)
}

// mountOperation implements the MountOperation interface.
type mountOperation struct {
	*externglib.Object
}

var _ MountOperation = (*mountOperation)(nil)

// WrapMountOperation wraps a GObject to a type that implements
// interface MountOperation. It is primarily used internally.
func WrapMountOperation(obj *externglib.Object) MountOperation {
	return mountOperation{obj}
}

func marshalMountOperation(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMountOperation(obj), nil
}

// NewMountOperation creates a new mount operation.
func NewMountOperation() MountOperation {
	var _cret *C.GMountOperation // in

	_cret = C.g_mount_operation_new()

	var _mountOperation MountOperation // out

	_mountOperation = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(MountOperation)

	return _mountOperation
}

func (o mountOperation) Anonymous() bool {
	var _arg0 *C.GMountOperation // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(o.Native()))

	_cret = C.g_mount_operation_get_anonymous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (o mountOperation) Choice() int {
	var _arg0 *C.GMountOperation // out
	var _cret C.int              // in

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(o.Native()))

	_cret = C.g_mount_operation_get_choice(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (o mountOperation) Domain() string {
	var _arg0 *C.GMountOperation // out
	var _cret *C.char            // in

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(o.Native()))

	_cret = C.g_mount_operation_get_domain(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (o mountOperation) IsTcryptHiddenVolume() bool {
	var _arg0 *C.GMountOperation // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(o.Native()))

	_cret = C.g_mount_operation_get_is_tcrypt_hidden_volume(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (o mountOperation) IsTcryptSystemVolume() bool {
	var _arg0 *C.GMountOperation // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(o.Native()))

	_cret = C.g_mount_operation_get_is_tcrypt_system_volume(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (o mountOperation) Password() string {
	var _arg0 *C.GMountOperation // out
	var _cret *C.char            // in

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(o.Native()))

	_cret = C.g_mount_operation_get_password(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (o mountOperation) PasswordSave() PasswordSave {
	var _arg0 *C.GMountOperation // out
	var _cret C.GPasswordSave    // in

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(o.Native()))

	_cret = C.g_mount_operation_get_password_save(_arg0)

	var _passwordSave PasswordSave // out

	_passwordSave = PasswordSave(_cret)

	return _passwordSave
}

func (o mountOperation) Pim() uint {
	var _arg0 *C.GMountOperation // out
	var _cret C.guint            // in

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(o.Native()))

	_cret = C.g_mount_operation_get_pim(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

func (o mountOperation) Username() string {
	var _arg0 *C.GMountOperation // out
	var _cret *C.char            // in

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(o.Native()))

	_cret = C.g_mount_operation_get_username(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (o mountOperation) Reply(result MountOperationResult) {
	var _arg0 *C.GMountOperation      // out
	var _arg1 C.GMountOperationResult // out

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(o.Native()))
	_arg1 = C.GMountOperationResult(result)

	C.g_mount_operation_reply(_arg0, _arg1)
}

func (o mountOperation) SetAnonymous(anonymous bool) {
	var _arg0 *C.GMountOperation // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(o.Native()))
	if anonymous {
		_arg1 = C.TRUE
	}

	C.g_mount_operation_set_anonymous(_arg0, _arg1)
}

func (o mountOperation) SetChoice(choice int) {
	var _arg0 *C.GMountOperation // out
	var _arg1 C.int              // out

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(o.Native()))
	_arg1 = C.int(choice)

	C.g_mount_operation_set_choice(_arg0, _arg1)
}

func (o mountOperation) SetDomain(domain string) {
	var _arg0 *C.GMountOperation // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.char)(C.CString(domain))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_mount_operation_set_domain(_arg0, _arg1)
}

func (o mountOperation) SetIsTcryptHiddenVolume(hiddenVolume bool) {
	var _arg0 *C.GMountOperation // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(o.Native()))
	if hiddenVolume {
		_arg1 = C.TRUE
	}

	C.g_mount_operation_set_is_tcrypt_hidden_volume(_arg0, _arg1)
}

func (o mountOperation) SetIsTcryptSystemVolume(systemVolume bool) {
	var _arg0 *C.GMountOperation // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(o.Native()))
	if systemVolume {
		_arg1 = C.TRUE
	}

	C.g_mount_operation_set_is_tcrypt_system_volume(_arg0, _arg1)
}

func (o mountOperation) SetPassword(password string) {
	var _arg0 *C.GMountOperation // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.char)(C.CString(password))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_mount_operation_set_password(_arg0, _arg1)
}

func (o mountOperation) SetPasswordSave(save PasswordSave) {
	var _arg0 *C.GMountOperation // out
	var _arg1 C.GPasswordSave    // out

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(o.Native()))
	_arg1 = C.GPasswordSave(save)

	C.g_mount_operation_set_password_save(_arg0, _arg1)
}

func (o mountOperation) SetPim(pim uint) {
	var _arg0 *C.GMountOperation // out
	var _arg1 C.guint            // out

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(o.Native()))
	_arg1 = C.guint(pim)

	C.g_mount_operation_set_pim(_arg0, _arg1)
}

func (o mountOperation) SetUsername(username string) {
	var _arg0 *C.GMountOperation // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GMountOperation)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.char)(C.CString(username))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_mount_operation_set_username(_arg0, _arg1)
}
