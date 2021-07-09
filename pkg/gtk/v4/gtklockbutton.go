// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_lock_button_get_type()), F: marshalLockButton},
	})
}

// LockButton: `GtkLockButton` is a widget to obtain and revoke authorizations
// needed to operate the controls.
//
// !An example GtkLockButton (lock-button.png)
//
// It is typically used in preference dialogs or control panels.
//
// The required authorization is represented by a `GPermission` object. Concrete
// implementations of `GPermission` may use PolicyKit or some other
// authorization framework. To obtain a PolicyKit-based `GPermission`, use
// `polkit_permission_new()`.
//
// If the user is not currently allowed to perform the action, but can obtain
// the permission, the widget looks like this:
//
// ! (lockbutton-locked.png)
//
// and the user can click the button to request the permission. Depending on the
// platform, this may pop up an authentication dialog or ask the user to
// authenticate in some other way. Once the user has obtained the permission,
// the widget changes to this:
//
// ! (lockbutton-unlocked.png)
//
// and the permission can be dropped again by clicking the button. If the user
// is not able to obtain the permission at all, the widget looks like this:
//
// ! (lockbutton-sorry.png)
//
// If the user has the permission and cannot drop it, the button is hidden.
//
// The text (and tooltips) that are shown in the various cases can be adjusted
// with the [property@Gtk.LockButton:text-lock],
// [property@Gtk.LockButton:text-unlock],
// [property@Gtk.LockButton:tooltip-lock],
// [property@Gtk.LockButton:tooltip-unlock] and
// [property@Gtk.LockButton:tooltip-not-authorized] properties.
type LockButton interface {
	gextras.Objector

	// Permission obtains the `GPermission` object that controls @button.
	Permission() *gio.PermissionClass
	// SetPermission sets the `GPermission` object that controls @button.
	SetPermission(permission gio.Permission)
}

// LockButtonClass implements the LockButton interface.
type LockButtonClass struct {
	*externglib.Object
	ButtonClass
	AccessibleInterface
	ActionableInterface
	BuildableInterface
	ConstraintTargetInterface
}

var _ LockButton = (*LockButtonClass)(nil)

func wrapLockButton(obj *externglib.Object) LockButton {
	return &LockButtonClass{
		Object: obj,
		ButtonClass: ButtonClass{
			Object: obj,
			WidgetClass: WidgetClass{
				InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
				AccessibleInterface: AccessibleInterface{
					Object: obj,
				},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
				ConstraintTargetInterface: ConstraintTargetInterface{
					Object: obj,
				},
			},
			AccessibleInterface: AccessibleInterface{
				Object: obj,
			},
			ActionableInterface: ActionableInterface{
				WidgetClass: WidgetClass{
					InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
					AccessibleInterface: AccessibleInterface{
						Object: obj,
					},
					BuildableInterface: BuildableInterface{
						Object: obj,
					},
					ConstraintTargetInterface: ConstraintTargetInterface{
						Object: obj,
					},
				},
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
			ConstraintTargetInterface: ConstraintTargetInterface{
				Object: obj,
			},
		},
		AccessibleInterface: AccessibleInterface{
			Object: obj,
		},
		ActionableInterface: ActionableInterface{
			WidgetClass: WidgetClass{
				InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
				AccessibleInterface: AccessibleInterface{
					Object: obj,
				},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
				ConstraintTargetInterface: ConstraintTargetInterface{
					Object: obj,
				},
			},
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		ConstraintTargetInterface: ConstraintTargetInterface{
			Object: obj,
		},
	}
}

func marshalLockButton(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapLockButton(obj), nil
}

// NewLockButton creates a new lock button which reflects the @permission.
func NewLockButton(permission gio.Permission) *LockButtonClass {
	var _arg1 *C.GPermission // out
	var _cret *C.GtkWidget   // in

	_arg1 = (*C.GPermission)(unsafe.Pointer(permission.Native()))

	_cret = C.gtk_lock_button_new(_arg1)

	var _lockButton *LockButtonClass // out

	_lockButton = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*LockButtonClass)

	return _lockButton
}

// Permission obtains the `GPermission` object that controls @button.
func (b *LockButtonClass) Permission() *gio.PermissionClass {
	var _arg0 *C.GtkLockButton // out
	var _cret *C.GPermission   // in

	_arg0 = (*C.GtkLockButton)(unsafe.Pointer(b.Native()))

	_cret = C.gtk_lock_button_get_permission(_arg0)

	var _permission *gio.PermissionClass // out

	_permission = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gio.PermissionClass)

	return _permission
}

// SetPermission sets the `GPermission` object that controls @button.
func (b *LockButtonClass) SetPermission(permission gio.Permission) {
	var _arg0 *C.GtkLockButton // out
	var _arg1 *C.GPermission   // out

	_arg0 = (*C.GtkLockButton)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GPermission)(unsafe.Pointer(permission.Native()))

	C.gtk_lock_button_set_permission(_arg0, _arg1)
}
