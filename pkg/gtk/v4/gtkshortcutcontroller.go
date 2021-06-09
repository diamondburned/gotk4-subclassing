// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_shortcut_controller_get_type()), F: marshalShortcutController},
	})
}

// ShortcutController is an event controller that manages shortcuts.
//
// Most common shortcuts are using this controller implicitly, e.g. by adding a
// mnemonic underline to a Label, or by installing a key binding using
// gtk_widget_class_add_binding(), or by adding accelerators to global actions
// using gtk_application_set_accels_for_action().
//
// But it is possible to create your own shortcut controller, and add shortcuts
// to it.
//
// ShortcutController implements Model for querying the shortcuts that have been
// added to it.
//
//
// GtkShortcutController as a GtkBuildable
//
// GtkShortcutControllers can be creates in ui files to set up shortcuts in the
// same place as the widgets.
//
// An example of a UI definition fragment with GtkShortcutController:
//
//    <object class='GtkButton'>
//      <child>
//        <object class='GtkShortcutController'>
//          <property name='scope'>managed</property>
//          <child>
//            <object class='GtkShortcut'>
//              <property name='trigger'>&amp;lt;Control&amp;gt;k</property>
//              <property name='action'>activate</property>
//            </object>
//          </child>
//        </object>
//      </child>
//    </object>
//
// This example creates a ActivateAction for triggering the `activate` signal of
// the GtkButton. See gtk_shortcut_action_parse_string() for the syntax for
// other kinds of ShortcutAction. See gtk_shortcut_trigger_parse_string() to
// learn more about the syntax for triggers.
type ShortcutController interface {
	EventController
	gio.ListModel
	Buildable

	// AddShortcut adds @shortcut to the list of shortcuts handled by @self.
	//
	// If this controller uses an external shortcut list, this function does
	// nothing.
	AddShortcut(shortcut Shortcut)
	// MnemonicsModifiers gets the mnemonics modifiers for when this controller
	// activates its shortcuts. See
	// gtk_shortcut_controller_set_mnemonics_modifiers() for details.
	MnemonicsModifiers() gdk.ModifierType
	// Scope gets the scope for when this controller activates its shortcuts.
	// See gtk_shortcut_controller_set_scope() for details.
	Scope() ShortcutScope
	// RemoveShortcut removes @shortcut from the list of shortcuts handled by
	// @self.
	//
	// If @shortcut had not been added to @controller or this controller uses an
	// external shortcut list, this function does nothing.
	RemoveShortcut(shortcut Shortcut)
	// SetMnemonicsModifiers sets the controller to have the given
	// @mnemonics_modifiers.
	//
	// The mnemonics modifiers determines which modifiers need to be pressed to
	// allow activation of shortcuts with mnemonics triggers.
	//
	// GTK normally uses the Alt modifier for mnemonics, except in PopoverMenus,
	// where mnemonics can be triggered without any modifiers. It should be very
	// rarely necessary to change this, and doing so is likely to interfere with
	// other shortcuts.
	//
	// This value is only relevant for local shortcut controllers. Global and
	// managed shortcut controllers will have their shortcuts activated from
	// other places which have their own modifiers for activating mnemonics.
	SetMnemonicsModifiers(modifiers gdk.ModifierType)
	// SetScope sets the controller to have the given @scope.
	//
	// The scope allows shortcuts to be activated outside of the normal event
	// propagation. In particular, it allows installing global keyboard
	// shortcuts that can be activated even when a widget does not have focus.
	//
	// With GTK_SHORTCUT_SCOPE_LOCAL, shortcuts will only be activated when the
	// widget has focus.
	SetScope(scope ShortcutScope)
}

// shortcutController implements the ShortcutController interface.
type shortcutController struct {
	EventController
	gio.ListModel
	Buildable
}

var _ ShortcutController = (*shortcutController)(nil)

// WrapShortcutController wraps a GObject to the right type. It is
// primarily used internally.
func WrapShortcutController(obj *externglib.Object) ShortcutController {
	return ShortcutController{
		EventController: WrapEventController(obj),
		gio.ListModel:   gio.WrapListModel(obj),
		Buildable:       WrapBuildable(obj),
	}
}

func marshalShortcutController(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapShortcutController(obj), nil
}

// NewShortcutController constructs a class ShortcutController.
func NewShortcutController() ShortcutController {
	var cret C.GtkShortcutController

	cret = C.gtk_shortcut_controller_new()

	var shortcutController ShortcutController

	shortcutController = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ShortcutController)

	return shortcutController
}

// NewShortcutControllerForModel constructs a class ShortcutController.
func NewShortcutControllerForModel(model gio.ListModel) ShortcutController {
	var arg1 *C.GListModel

	arg1 = (*C.GListModel)(unsafe.Pointer(model.Native()))

	var cret C.GtkShortcutController

	cret = C.gtk_shortcut_controller_new_for_model(arg1)

	var shortcutController ShortcutController

	shortcutController = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ShortcutController)

	return shortcutController
}

// AddShortcut adds @shortcut to the list of shortcuts handled by @self.
//
// If this controller uses an external shortcut list, this function does
// nothing.
func (s shortcutController) AddShortcut(shortcut Shortcut) {
	var arg0 *C.GtkShortcutController
	var arg1 *C.GtkShortcut

	arg0 = (*C.GtkShortcutController)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkShortcut)(unsafe.Pointer(shortcut.Native()))

	C.gtk_shortcut_controller_add_shortcut(arg0, arg1)
}

// MnemonicsModifiers gets the mnemonics modifiers for when this controller
// activates its shortcuts. See
// gtk_shortcut_controller_set_mnemonics_modifiers() for details.
func (s shortcutController) MnemonicsModifiers() gdk.ModifierType {
	var arg0 *C.GtkShortcutController

	arg0 = (*C.GtkShortcutController)(unsafe.Pointer(s.Native()))

	var cret C.GdkModifierType

	cret = C.gtk_shortcut_controller_get_mnemonics_modifiers(arg0)

	var modifierType gdk.ModifierType

	modifierType = gdk.ModifierType(cret)

	return modifierType
}

// Scope gets the scope for when this controller activates its shortcuts.
// See gtk_shortcut_controller_set_scope() for details.
func (s shortcutController) Scope() ShortcutScope {
	var arg0 *C.GtkShortcutController

	arg0 = (*C.GtkShortcutController)(unsafe.Pointer(s.Native()))

	var cret C.GtkShortcutScope

	cret = C.gtk_shortcut_controller_get_scope(arg0)

	var shortcutScope ShortcutScope

	shortcutScope = ShortcutScope(cret)

	return shortcutScope
}

// RemoveShortcut removes @shortcut from the list of shortcuts handled by
// @self.
//
// If @shortcut had not been added to @controller or this controller uses an
// external shortcut list, this function does nothing.
func (s shortcutController) RemoveShortcut(shortcut Shortcut) {
	var arg0 *C.GtkShortcutController
	var arg1 *C.GtkShortcut

	arg0 = (*C.GtkShortcutController)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkShortcut)(unsafe.Pointer(shortcut.Native()))

	C.gtk_shortcut_controller_remove_shortcut(arg0, arg1)
}

// SetMnemonicsModifiers sets the controller to have the given
// @mnemonics_modifiers.
//
// The mnemonics modifiers determines which modifiers need to be pressed to
// allow activation of shortcuts with mnemonics triggers.
//
// GTK normally uses the Alt modifier for mnemonics, except in PopoverMenus,
// where mnemonics can be triggered without any modifiers. It should be very
// rarely necessary to change this, and doing so is likely to interfere with
// other shortcuts.
//
// This value is only relevant for local shortcut controllers. Global and
// managed shortcut controllers will have their shortcuts activated from
// other places which have their own modifiers for activating mnemonics.
func (s shortcutController) SetMnemonicsModifiers(modifiers gdk.ModifierType) {
	var arg0 *C.GtkShortcutController
	var arg1 C.GdkModifierType

	arg0 = (*C.GtkShortcutController)(unsafe.Pointer(s.Native()))
	arg1 = (C.GdkModifierType)(modifiers)

	C.gtk_shortcut_controller_set_mnemonics_modifiers(arg0, arg1)
}

// SetScope sets the controller to have the given @scope.
//
// The scope allows shortcuts to be activated outside of the normal event
// propagation. In particular, it allows installing global keyboard
// shortcuts that can be activated even when a widget does not have focus.
//
// With GTK_SHORTCUT_SCOPE_LOCAL, shortcuts will only be activated when the
// widget has focus.
func (s shortcutController) SetScope(scope ShortcutScope) {
	var arg0 *C.GtkShortcutController
	var arg1 C.GtkShortcutScope

	arg0 = (*C.GtkShortcutController)(unsafe.Pointer(s.Native()))
	arg1 = (C.GtkShortcutScope)(scope)

	C.gtk_shortcut_controller_set_scope(arg0, arg1)
}
