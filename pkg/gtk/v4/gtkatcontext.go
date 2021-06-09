// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_at_context_get_type()), F: marshalATContext},
	})
}

// ATContext: gtkATContext is an abstract class provided by GTK to communicate
// to platform-specific assistive technologies API.
//
// Each platform supported by GTK implements a ATContext subclass, and is
// responsible for updating the accessible state in response to state changes in
// Accessible.
type ATContext interface {
	gextras.Objector

	// Accessible retrieves the Accessible using this context.
	Accessible() Accessible
	// AccessibleRole retrieves the accessible role of this context.
	AccessibleRole() AccessibleRole
}

// atContext implements the ATContext interface.
type atContext struct {
	gextras.Objector
}

var _ ATContext = (*atContext)(nil)

// WrapATContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapATContext(obj *externglib.Object) ATContext {
	return ATContext{
		Objector: obj,
	}
}

func marshalATContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapATContext(obj), nil
}

// NewATContextCreate constructs a class ATContext.
func NewATContextCreate(accessibleRole AccessibleRole, accessible Accessible, display gdk.Display) ATContext {
	var arg1 C.GtkAccessibleRole
	var arg2 *C.GtkAccessible
	var arg3 *C.GdkDisplay

	arg1 = (C.GtkAccessibleRole)(accessibleRole)
	arg2 = (*C.GtkAccessible)(unsafe.Pointer(accessible.Native()))
	arg3 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	var cret C.GtkATContext

	cret = C.gtk_at_context_create(arg1, arg2, arg3)

	var atContext ATContext

	atContext = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ATContext)

	return atContext
}

// Accessible retrieves the Accessible using this context.
func (s atContext) Accessible() Accessible {
	var arg0 *C.GtkATContext

	arg0 = (*C.GtkATContext)(unsafe.Pointer(s.Native()))

	var cret *C.GtkAccessible

	cret = C.gtk_at_context_get_accessible(arg0)

	var accessible Accessible

	accessible = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Accessible)

	return accessible
}

// AccessibleRole retrieves the accessible role of this context.
func (s atContext) AccessibleRole() AccessibleRole {
	var arg0 *C.GtkATContext

	arg0 = (*C.GtkATContext)(unsafe.Pointer(s.Native()))

	var cret C.GtkAccessibleRole

	cret = C.gtk_at_context_get_accessible_role(arg0)

	var accessibleRole AccessibleRole

	accessibleRole = AccessibleRole(cret)

	return accessibleRole
}
