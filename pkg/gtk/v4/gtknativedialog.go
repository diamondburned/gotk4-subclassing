// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_native_dialog_get_type()), F: marshalNativeDialog},
	})
}

// NativeDialog: native dialogs are platform dialogs that don't use Dialog or
// Window. They are used in order to integrate better with a platform, by
// looking the same as other native applications and supporting platform
// specific features.
//
// The Dialog functions cannot be used on such objects, but we need a similar
// API in order to drive them. The NativeDialog object is an API that allows you
// to do this. It allows you to set various common properties on the dialog, as
// well as show and hide it and get a NativeDialog::response signal when the
// user finished with the dialog.
//
// Note that unlike Dialog, NativeDialog objects are not toplevel widgets, and
// GTK does not keep them alive. It is your responsibility to keep a reference
// until you are done with the object.
type NativeDialog interface {
	gextras.Objector

	// Destroy destroys a dialog.
	//
	// When a dialog is destroyed, it will break any references it holds to
	// other objects. If it is visible it will be hidden and any underlying
	// window system resources will be destroyed.
	//
	// Note that this does not release any reference to the object (as opposed
	// to destroying a GtkWindow) because there is no reference from the
	// windowing system to the NativeDialog.
	Destroy()
	// Modal returns whether the dialog is modal. See
	// gtk_native_dialog_set_modal().
	Modal() bool
	// Title gets the title of the NativeDialog.
	Title() string
	// TransientFor fetches the transient parent for this window. See
	// gtk_native_dialog_set_transient_for().
	TransientFor() Window
	// Visible determines whether the dialog is visible.
	Visible() bool
	// Hide hides the dialog if it is visilbe, aborting any interaction. Once
	// this is called the NativeDialog::response signal will not be emitted
	// until after the next call to gtk_native_dialog_show().
	//
	// If the dialog is not visible this does nothing.
	Hide()
	// SetModal sets a dialog modal or non-modal. Modal dialogs prevent
	// interaction with other windows in the same application. To keep modal
	// dialogs on top of main application windows, use
	// gtk_native_dialog_set_transient_for() to make the dialog transient for
	// the parent; most [window managers][gtk-X11-arch] will then disallow
	// lowering the dialog below the parent.
	SetModal(modal bool)
	// SetTitle sets the title of the NativeDialog.
	SetTitle(title string)
	// SetTransientFor: dialog windows should be set transient for the main
	// application window they were spawned from. This allows [window
	// managers][gtk-X11-arch] to e.g. keep the dialog on top of the main
	// window, or center the dialog over the main window.
	//
	// Passing nil for @parent unsets the current transient window.
	SetTransientFor(parent Window)
	// Show shows the dialog on the display, allowing the user to interact with
	// it. When the user accepts the state of the dialog the dialog will be
	// automatically hidden and the NativeDialog::response signal will be
	// emitted.
	//
	// Multiple calls while the dialog is visible will be ignored.
	Show()
}

// nativeDialog implements the NativeDialog interface.
type nativeDialog struct {
	gextras.Objector
}

var _ NativeDialog = (*nativeDialog)(nil)

// WrapNativeDialog wraps a GObject to the right type. It is
// primarily used internally.
func WrapNativeDialog(obj *externglib.Object) NativeDialog {
	return NativeDialog{
		Objector: obj,
	}
}

func marshalNativeDialog(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNativeDialog(obj), nil
}

// Destroy destroys a dialog.
//
// When a dialog is destroyed, it will break any references it holds to
// other objects. If it is visible it will be hidden and any underlying
// window system resources will be destroyed.
//
// Note that this does not release any reference to the object (as opposed
// to destroying a GtkWindow) because there is no reference from the
// windowing system to the NativeDialog.
func (s nativeDialog) Destroy() {
	var arg0 *C.GtkNativeDialog

	arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(s.Native()))

	C.gtk_native_dialog_destroy(arg0)
}

// Modal returns whether the dialog is modal. See
// gtk_native_dialog_set_modal().
func (s nativeDialog) Modal() bool {
	var arg0 *C.GtkNativeDialog

	arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(s.Native()))

	var cret C.gboolean

	cret = C.gtk_native_dialog_get_modal(arg0)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// Title gets the title of the NativeDialog.
func (s nativeDialog) Title() string {
	var arg0 *C.GtkNativeDialog

	arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(s.Native()))

	var cret *C.char

	cret = C.gtk_native_dialog_get_title(arg0)

	var utf8 string

	utf8 = C.GoString(cret)

	return utf8
}

// TransientFor fetches the transient parent for this window. See
// gtk_native_dialog_set_transient_for().
func (s nativeDialog) TransientFor() Window {
	var arg0 *C.GtkNativeDialog

	arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(s.Native()))

	var cret *C.GtkWindow

	cret = C.gtk_native_dialog_get_transient_for(arg0)

	var window Window

	window = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Window)

	return window
}

// Visible determines whether the dialog is visible.
func (s nativeDialog) Visible() bool {
	var arg0 *C.GtkNativeDialog

	arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(s.Native()))

	var cret C.gboolean

	cret = C.gtk_native_dialog_get_visible(arg0)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// Hide hides the dialog if it is visilbe, aborting any interaction. Once
// this is called the NativeDialog::response signal will not be emitted
// until after the next call to gtk_native_dialog_show().
//
// If the dialog is not visible this does nothing.
func (s nativeDialog) Hide() {
	var arg0 *C.GtkNativeDialog

	arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(s.Native()))

	C.gtk_native_dialog_hide(arg0)
}

// SetModal sets a dialog modal or non-modal. Modal dialogs prevent
// interaction with other windows in the same application. To keep modal
// dialogs on top of main application windows, use
// gtk_native_dialog_set_transient_for() to make the dialog transient for
// the parent; most [window managers][gtk-X11-arch] will then disallow
// lowering the dialog below the parent.
func (s nativeDialog) SetModal(modal bool) {
	var arg0 *C.GtkNativeDialog
	var arg1 C.gboolean

	arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(s.Native()))
	if modal {
		arg1 = C.gboolean(1)
	}

	C.gtk_native_dialog_set_modal(arg0, arg1)
}

// SetTitle sets the title of the NativeDialog.
func (s nativeDialog) SetTitle(title string) {
	var arg0 *C.GtkNativeDialog
	var arg1 *C.char

	arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(title))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_native_dialog_set_title(arg0, arg1)
}

// SetTransientFor: dialog windows should be set transient for the main
// application window they were spawned from. This allows [window
// managers][gtk-X11-arch] to e.g. keep the dialog on top of the main
// window, or center the dialog over the main window.
//
// Passing nil for @parent unsets the current transient window.
func (s nativeDialog) SetTransientFor(parent Window) {
	var arg0 *C.GtkNativeDialog
	var arg1 *C.GtkWindow

	arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))

	C.gtk_native_dialog_set_transient_for(arg0, arg1)
}

// Show shows the dialog on the display, allowing the user to interact with
// it. When the user accepts the state of the dialog the dialog will be
// automatically hidden and the NativeDialog::response signal will be
// emitted.
//
// Multiple calls while the dialog is visible will be ignored.
func (s nativeDialog) Show() {
	var arg0 *C.GtkNativeDialog

	arg0 = (*C.GtkNativeDialog)(unsafe.Pointer(s.Native()))

	C.gtk_native_dialog_show(arg0)
}
