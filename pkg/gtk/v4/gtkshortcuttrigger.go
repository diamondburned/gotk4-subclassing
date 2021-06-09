// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_alternative_trigger_get_type()), F: marshalAlternativeTrigger},
		{T: externglib.Type(C.gtk_keyval_trigger_get_type()), F: marshalKeyvalTrigger},
		{T: externglib.Type(C.gtk_mnemonic_trigger_get_type()), F: marshalMnemonicTrigger},
		{T: externglib.Type(C.gtk_never_trigger_get_type()), F: marshalNeverTrigger},
		{T: externglib.Type(C.gtk_shortcut_trigger_get_type()), F: marshalShortcutTrigger},
	})
}

// AlternativeTrigger: a ShortcutTrigger that triggers when either of two
// ShortcutTriggers trigger.
type AlternativeTrigger interface {
	ShortcutTrigger

	// First gets the first of the two alternative triggers that may trigger
	// @self. gtk_alternative_trigger_get_second() will return the other one.
	First() ShortcutTrigger
	// Second gets the second of the two alternative triggers that may trigger
	// @self. gtk_alternative_trigger_get_first() will return the other one.
	Second() ShortcutTrigger
}

// alternativeTrigger implements the AlternativeTrigger interface.
type alternativeTrigger struct {
	ShortcutTrigger
}

var _ AlternativeTrigger = (*alternativeTrigger)(nil)

// WrapAlternativeTrigger wraps a GObject to the right type. It is
// primarily used internally.
func WrapAlternativeTrigger(obj *externglib.Object) AlternativeTrigger {
	return AlternativeTrigger{
		ShortcutTrigger: WrapShortcutTrigger(obj),
	}
}

func marshalAlternativeTrigger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAlternativeTrigger(obj), nil
}

// NewAlternativeTrigger constructs a class AlternativeTrigger.
func NewAlternativeTrigger(first ShortcutTrigger, second ShortcutTrigger) AlternativeTrigger {
	var arg1 *C.GtkShortcutTrigger
	var arg2 *C.GtkShortcutTrigger

	arg1 = (*C.GtkShortcutTrigger)(unsafe.Pointer(first.Native()))
	arg2 = (*C.GtkShortcutTrigger)(unsafe.Pointer(second.Native()))

	var cret C.GtkAlternativeTrigger

	cret = C.gtk_alternative_trigger_new(arg1, arg2)

	var alternativeTrigger AlternativeTrigger

	alternativeTrigger = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(AlternativeTrigger)

	return alternativeTrigger
}

// First gets the first of the two alternative triggers that may trigger
// @self. gtk_alternative_trigger_get_second() will return the other one.
func (s alternativeTrigger) First() ShortcutTrigger {
	var arg0 *C.GtkAlternativeTrigger

	arg0 = (*C.GtkAlternativeTrigger)(unsafe.Pointer(s.Native()))

	var cret *C.GtkShortcutTrigger

	cret = C.gtk_alternative_trigger_get_first(arg0)

	var shortcutTrigger ShortcutTrigger

	shortcutTrigger = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ShortcutTrigger)

	return shortcutTrigger
}

// Second gets the second of the two alternative triggers that may trigger
// @self. gtk_alternative_trigger_get_first() will return the other one.
func (s alternativeTrigger) Second() ShortcutTrigger {
	var arg0 *C.GtkAlternativeTrigger

	arg0 = (*C.GtkAlternativeTrigger)(unsafe.Pointer(s.Native()))

	var cret *C.GtkShortcutTrigger

	cret = C.gtk_alternative_trigger_get_second(arg0)

	var shortcutTrigger ShortcutTrigger

	shortcutTrigger = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ShortcutTrigger)

	return shortcutTrigger
}

// KeyvalTrigger: a ShortcutTrigger that triggers when a specific keyval and
// (optionally) modifiers are pressed.
type KeyvalTrigger interface {
	ShortcutTrigger

	// Keyval gets the keyval that must be pressed to succeed triggering @self.
	Keyval() uint
	// Modifiers gets the modifiers that must be present to succeed triggering
	// @self.
	Modifiers() gdk.ModifierType
}

// keyvalTrigger implements the KeyvalTrigger interface.
type keyvalTrigger struct {
	ShortcutTrigger
}

var _ KeyvalTrigger = (*keyvalTrigger)(nil)

// WrapKeyvalTrigger wraps a GObject to the right type. It is
// primarily used internally.
func WrapKeyvalTrigger(obj *externglib.Object) KeyvalTrigger {
	return KeyvalTrigger{
		ShortcutTrigger: WrapShortcutTrigger(obj),
	}
}

func marshalKeyvalTrigger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapKeyvalTrigger(obj), nil
}

// NewKeyvalTrigger constructs a class KeyvalTrigger.
func NewKeyvalTrigger(keyval uint, modifiers gdk.ModifierType) KeyvalTrigger {
	var arg1 C.guint
	var arg2 C.GdkModifierType

	arg1 = C.guint(keyval)
	arg2 = (C.GdkModifierType)(modifiers)

	var cret C.GtkKeyvalTrigger

	cret = C.gtk_keyval_trigger_new(arg1, arg2)

	var keyvalTrigger KeyvalTrigger

	keyvalTrigger = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(KeyvalTrigger)

	return keyvalTrigger
}

// Keyval gets the keyval that must be pressed to succeed triggering @self.
func (s keyvalTrigger) Keyval() uint {
	var arg0 *C.GtkKeyvalTrigger

	arg0 = (*C.GtkKeyvalTrigger)(unsafe.Pointer(s.Native()))

	var cret C.guint

	cret = C.gtk_keyval_trigger_get_keyval(arg0)

	var guint uint

	guint = (uint)(cret)

	return guint
}

// Modifiers gets the modifiers that must be present to succeed triggering
// @self.
func (s keyvalTrigger) Modifiers() gdk.ModifierType {
	var arg0 *C.GtkKeyvalTrigger

	arg0 = (*C.GtkKeyvalTrigger)(unsafe.Pointer(s.Native()))

	var cret C.GdkModifierType

	cret = C.gtk_keyval_trigger_get_modifiers(arg0)

	var modifierType gdk.ModifierType

	modifierType = gdk.ModifierType(cret)

	return modifierType
}

// MnemonicTrigger: a ShortcutTrigger that triggers when a specific mnemonic is
// pressed.
type MnemonicTrigger interface {
	ShortcutTrigger

	// Keyval gets the keyval that must be pressed to succeed triggering @self.
	Keyval() uint
}

// mnemonicTrigger implements the MnemonicTrigger interface.
type mnemonicTrigger struct {
	ShortcutTrigger
}

var _ MnemonicTrigger = (*mnemonicTrigger)(nil)

// WrapMnemonicTrigger wraps a GObject to the right type. It is
// primarily used internally.
func WrapMnemonicTrigger(obj *externglib.Object) MnemonicTrigger {
	return MnemonicTrigger{
		ShortcutTrigger: WrapShortcutTrigger(obj),
	}
}

func marshalMnemonicTrigger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMnemonicTrigger(obj), nil
}

// NewMnemonicTrigger constructs a class MnemonicTrigger.
func NewMnemonicTrigger(keyval uint) MnemonicTrigger {
	var arg1 C.guint

	arg1 = C.guint(keyval)

	var cret C.GtkMnemonicTrigger

	cret = C.gtk_mnemonic_trigger_new(arg1)

	var mnemonicTrigger MnemonicTrigger

	mnemonicTrigger = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(MnemonicTrigger)

	return mnemonicTrigger
}

// Keyval gets the keyval that must be pressed to succeed triggering @self.
func (s mnemonicTrigger) Keyval() uint {
	var arg0 *C.GtkMnemonicTrigger

	arg0 = (*C.GtkMnemonicTrigger)(unsafe.Pointer(s.Native()))

	var cret C.guint

	cret = C.gtk_mnemonic_trigger_get_keyval(arg0)

	var guint uint

	guint = (uint)(cret)

	return guint
}

// NeverTrigger: a ShortcutTrigger that never triggers.
type NeverTrigger interface {
	ShortcutTrigger
}

// neverTrigger implements the NeverTrigger interface.
type neverTrigger struct {
	ShortcutTrigger
}

var _ NeverTrigger = (*neverTrigger)(nil)

// WrapNeverTrigger wraps a GObject to the right type. It is
// primarily used internally.
func WrapNeverTrigger(obj *externglib.Object) NeverTrigger {
	return NeverTrigger{
		ShortcutTrigger: WrapShortcutTrigger(obj),
	}
}

func marshalNeverTrigger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapNeverTrigger(obj), nil
}

// ShortcutTrigger is the object used to track if a Shortcut should be
// activated. For this purpose, gtk_shortcut_trigger_trigger() can be called on
// a Event.
//
// ShortcutTriggers contain functions that allow easy presentation to end users
// as well as being printed for debugging.
//
// All ShortcutTriggers are immutable, you can only specify their properties
// during construction. If you want to change a trigger, you have to replace it
// with a new one.
type ShortcutTrigger interface {
	gextras.Objector

	// Compare: the types of @trigger1 and @trigger2 are #gconstpointer only to
	// allow use of this function as a Func. They must each be a
	// ShortcutTrigger.
	Compare(trigger2 ShortcutTrigger) int
	// Equal checks if @trigger1 and @trigger2 trigger under the same
	// conditions.
	//
	// The types of @one and @two are #gconstpointer only to allow use of this
	// function with Table. They must each be a ShortcutTrigger.
	Equal(trigger2 ShortcutTrigger) bool
	// Hash generates a hash value for a ShortcutTrigger.
	//
	// The output of this function is guaranteed to be the same for a given
	// value only per-process. It may change between different processor
	// architectures or even different versions of GTK. Do not use this function
	// as a basis for building protocols or file formats.
	//
	// The types of @trigger is #gconstpointer only to allow use of this
	// function with Table. They must each be a ShortcutTrigger.
	Hash() uint
	// Print prints the given trigger into a string for the developer. This is
	// meant for debugging and logging.
	//
	// The form of the representation may change at any time and is not
	// guaranteed to stay identical.
	Print(string *glib.String)
	// PrintLabel prints the given trigger into a string. This function is
	// returning a translated string for presentation to end users for example
	// in menu items or in help texts.
	//
	// The @display in use may influence the resulting string in various forms,
	// such as resolving hardware keycodes or by causing display-specific
	// modifier names.
	//
	// The form of the representation may change at any time and is not
	// guaranteed to stay identical.
	PrintLabel(display gdk.Display, string *glib.String) bool
	// ToLabel gets textual representation for the given trigger. This function
	// is returning a translated string for presentation to end users for
	// example in menu items or in help texts.
	//
	// The @display in use may influence the resulting string in various forms,
	// such as resolving hardware keycodes or by causing display-specific
	// modifier names.
	//
	// The form of the representation may change at any time and is not
	// guaranteed to stay identical.
	ToLabel(display gdk.Display) string
	// String prints the given trigger into a human-readable string. This is a
	// small wrapper around gtk_shortcut_trigger_print() to help when debugging.
	String() string
	// Trigger checks if the given @event triggers @self.
	Trigger(event gdk.Event, enableMnemonics bool) gdk.KeyMatch
}

// shortcutTrigger implements the ShortcutTrigger interface.
type shortcutTrigger struct {
	gextras.Objector
}

var _ ShortcutTrigger = (*shortcutTrigger)(nil)

// WrapShortcutTrigger wraps a GObject to the right type. It is
// primarily used internally.
func WrapShortcutTrigger(obj *externglib.Object) ShortcutTrigger {
	return ShortcutTrigger{
		Objector: obj,
	}
}

func marshalShortcutTrigger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapShortcutTrigger(obj), nil
}

// NewShortcutTriggerParseString constructs a class ShortcutTrigger.
func NewShortcutTriggerParseString(string string) ShortcutTrigger {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkShortcutTrigger

	cret = C.gtk_shortcut_trigger_parse_string(arg1)

	var shortcutTrigger ShortcutTrigger

	shortcutTrigger = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ShortcutTrigger)

	return shortcutTrigger
}

// Compare: the types of @trigger1 and @trigger2 are #gconstpointer only to
// allow use of this function as a Func. They must each be a
// ShortcutTrigger.
func (t shortcutTrigger) Compare(trigger2 ShortcutTrigger) int {
	var arg0 C.gpointer
	var arg1 C.gpointer

	arg0 = (C.gpointer)(unsafe.Pointer(t.Native()))
	arg1 = (C.gpointer)(unsafe.Pointer(trigger2.Native()))

	var cret C.int

	cret = C.gtk_shortcut_trigger_compare(arg0, arg1)

	var gint int

	gint = (int)(cret)

	return gint
}

// Equal checks if @trigger1 and @trigger2 trigger under the same
// conditions.
//
// The types of @one and @two are #gconstpointer only to allow use of this
// function with Table. They must each be a ShortcutTrigger.
func (t shortcutTrigger) Equal(trigger2 ShortcutTrigger) bool {
	var arg0 C.gpointer
	var arg1 C.gpointer

	arg0 = (C.gpointer)(unsafe.Pointer(t.Native()))
	arg1 = (C.gpointer)(unsafe.Pointer(trigger2.Native()))

	var cret C.gboolean

	cret = C.gtk_shortcut_trigger_equal(arg0, arg1)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// Hash generates a hash value for a ShortcutTrigger.
//
// The output of this function is guaranteed to be the same for a given
// value only per-process. It may change between different processor
// architectures or even different versions of GTK. Do not use this function
// as a basis for building protocols or file formats.
//
// The types of @trigger is #gconstpointer only to allow use of this
// function with Table. They must each be a ShortcutTrigger.
func (t shortcutTrigger) Hash() uint {
	var arg0 C.gpointer

	arg0 = (C.gpointer)(unsafe.Pointer(t.Native()))

	var cret C.guint

	cret = C.gtk_shortcut_trigger_hash(arg0)

	var guint uint

	guint = (uint)(cret)

	return guint
}

// Print prints the given trigger into a string for the developer. This is
// meant for debugging and logging.
//
// The form of the representation may change at any time and is not
// guaranteed to stay identical.
func (s shortcutTrigger) Print(string *glib.String) {
	var arg0 *C.GtkShortcutTrigger
	var arg1 *C.GString

	arg0 = (*C.GtkShortcutTrigger)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GString)(unsafe.Pointer(string.Native()))

	C.gtk_shortcut_trigger_print(arg0, arg1)
}

// PrintLabel prints the given trigger into a string. This function is
// returning a translated string for presentation to end users for example
// in menu items or in help texts.
//
// The @display in use may influence the resulting string in various forms,
// such as resolving hardware keycodes or by causing display-specific
// modifier names.
//
// The form of the representation may change at any time and is not
// guaranteed to stay identical.
func (s shortcutTrigger) PrintLabel(display gdk.Display, string *glib.String) bool {
	var arg0 *C.GtkShortcutTrigger
	var arg1 *C.GdkDisplay
	var arg2 *C.GString

	arg0 = (*C.GtkShortcutTrigger)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))
	arg2 = (*C.GString)(unsafe.Pointer(string.Native()))

	var cret C.gboolean

	cret = C.gtk_shortcut_trigger_print_label(arg0, arg1, arg2)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// ToLabel gets textual representation for the given trigger. This function
// is returning a translated string for presentation to end users for
// example in menu items or in help texts.
//
// The @display in use may influence the resulting string in various forms,
// such as resolving hardware keycodes or by causing display-specific
// modifier names.
//
// The form of the representation may change at any time and is not
// guaranteed to stay identical.
func (s shortcutTrigger) ToLabel(display gdk.Display) string {
	var arg0 *C.GtkShortcutTrigger
	var arg1 *C.GdkDisplay

	arg0 = (*C.GtkShortcutTrigger)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	var cret *C.char

	cret = C.gtk_shortcut_trigger_to_label(arg0, arg1)

	var utf8 string

	utf8 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return utf8
}

// String prints the given trigger into a human-readable string. This is a
// small wrapper around gtk_shortcut_trigger_print() to help when debugging.
func (s shortcutTrigger) String() string {
	var arg0 *C.GtkShortcutTrigger

	arg0 = (*C.GtkShortcutTrigger)(unsafe.Pointer(s.Native()))

	var cret *C.char

	cret = C.gtk_shortcut_trigger_to_string(arg0)

	var utf8 string

	utf8 = C.GoString(cret)
	defer C.free(unsafe.Pointer(cret))

	return utf8
}

// Trigger checks if the given @event triggers @self.
func (s shortcutTrigger) Trigger(event gdk.Event, enableMnemonics bool) gdk.KeyMatch {
	var arg0 *C.GtkShortcutTrigger
	var arg1 *C.GdkEvent
	var arg2 C.gboolean

	arg0 = (*C.GtkShortcutTrigger)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GdkEvent)(unsafe.Pointer(event.Native()))
	if enableMnemonics {
		arg2 = C.gboolean(1)
	}

	var cret C.GdkKeyMatch

	cret = C.gtk_shortcut_trigger_trigger(arg0, arg1, arg2)

	var keyMatch gdk.KeyMatch

	keyMatch = gdk.KeyMatch(cret)

	return keyMatch
}
