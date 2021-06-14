// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_popover_get_type()), F: marshalPopover},
	})
}

// Popover: `GtkPopover` is a bubble-like context popup.
//
// !An example GtkPopover (popover.png)
//
// It is primarily meant to provide context-dependent information or options.
// Popovers are attached to a parent widget. By default, they point to the whole
// widget area, although this behavior can be changed with
// [method@Gtk.Popover.set_pointing_to].
//
// The position of a popover relative to the widget it is attached to can also
// be changed with [method@Gtk.Popover.set_position]
//
// By default, `GtkPopover` performs a grab, in order to ensure input events get
// redirected to it while it is shown, and also so the popover is dismissed in
// the expected situations (clicks outside the popover, or the Escape key being
// pressed). If no such modal behavior is desired on a popover,
// [method@Gtk.Popover.set_autohide] may be called on it to tweak its behavior.
//
//
// GtkPopover as menu replacement
//
// `GtkPopover` is often used to replace menus. The best was to do this is to
// use the [class@Gtk.PopoverMenu] subclass which supports being populated from
// a `GMenuModel` with [ctor@Gtk.PopoverMenu.new_from_model].
//
// “`xml <section> <attribute name="display-hint">horizontal-buttons</attribute>
// <item> <attribute name="label">Cut</attribute> <attribute
// name="action">app.cut</attribute> <attribute
// name="verb-icon">edit-cut-symbolic</attribute> </item> <item> <attribute
// name="label">Copy</attribute> <attribute name="action">app.copy</attribute>
// <attribute name="verb-icon">edit-copy-symbolic</attribute> </item> <item>
// <attribute name="label">Paste</attribute> <attribute
// name="action">app.paste</attribute> <attribute
// name="verb-icon">edit-paste-symbolic</attribute> </item> </section> “`
//
//
// CSS nodes
//
// “` popover[.menu] ├── arrow ╰── contents.background ╰── <child> “`
//
// The contents child node always gets the .background style class and the
// popover itself gets the .menu style class if the popover is menu-like (i.e.
// `GtkPopoverMenu`).
//
// Particular uses of `GtkPopover`, such as touch selection popups or magnifiers
// in `GtkEntry` or `GtkTextView` get style classes like .touch-selection or
// .magnifier to differentiate from plain popovers.
//
// When styling a popover directly, the popover node should usually not have any
// background. The visible part of the popover can have a shadow. To specify it
// in CSS, set the box-shadow of the contents node.
//
// Note that, in order to accomplish appropriate arrow visuals, `GtkPopover`
// uses custom drawing for the arrow node. This makes it possible for the arrow
// to change its shape dynamically, but it also limits the possibilities of
// styling it using CSS. In particular, the arrow gets drawn over the content
// node's border and shadow, so they look like one shape, which means that the
// border width of the content node and the arrow node should be the same. The
// arrow also does not support any border shape other than solid, no
// border-radius, only one border width (border-bottom-width is used) and no
// box-shadow.
type Popover interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Native
	ShortcutManager

	// Autohide returns whether the popover is modal.
	//
	// See [method@Gtk.Popover.set_autohide] for the implications of this.
	Autohide() bool
	// CascadePopdown returns whether the popover will close after a modal child
	// is closed.
	CascadePopdown() bool
	// Child gets the child widget of @popover.
	Child() Widget
	// HasArrow gets whether this popover is showing an arrow pointing at the
	// widget that it is relative to.
	HasArrow() bool
	// MnemonicsVisible gets whether mnemonics are visible.
	MnemonicsVisible() bool
	// Offset gets the offset previous set with gtk_popover_set_offset().
	Offset() (xOffset int, yOffset int)
	// PointingTo gets the rectangle that the popover points to.
	//
	// If a rectangle to point to has been set, this function will return true
	// and fill in @rect with such rectangle, otherwise it will return false and
	// fill in @rect with the parent widget coordinates.
	PointingTo() (gdk.Rectangle, bool)
	// Position returns the preferred position of @popover.
	Position() PositionType
	// Popdown pops @popover down.
	//
	// This is different from a [method@Gtk.Widget.hide] call in that it may
	// show the popover with a transition. If you want to hide the popover
	// without a transition, just use [method@Gtk.Widget.hide].
	Popdown()
	// Popup pops @popover up.
	//
	// This is different from a [method@Gtk.Widget.show() call in that it may
	// show the popover with a transition. If you want to show the popover
	// without a transition, just use [method@Gtk.Widget.show].
	Popup()
	// Present presents the popover to the user.
	Present()
	// SetAutohide sets whether @popover is modal.
	//
	// A modal popover will grab the keyboard focus on it when being displayed.
	// Clicking outside the popover area or pressing Esc will dismiss the
	// popover.
	//
	// Called this function on an already showing popup with a new autohide
	// value different from the current one, will cause the popup to be hidden.
	SetAutohide(autohide bool)
	// SetCascadePopdown: if @cascade_popdown is true, the popover will be
	// closed when a child modal popover is closed.
	//
	// If false, @popover will stay visible.
	SetCascadePopdown(cascadePopdown bool)
	// SetChild sets the child widget of @popover.
	SetChild(child Widget)
	// SetDefaultWidget sets the default widget of a `GtkPopover`.
	//
	// The default widget is the widget that’s activated when the user presses
	// Enter in a dialog (for example). This function sets or unsets the default
	// widget for a `GtkPopover`.
	SetDefaultWidget(widget Widget)
	// SetHasArrow sets whether this popover should draw an arrow pointing at
	// the widget it is relative to.
	SetHasArrow(hasArrow bool)
	// SetMnemonicsVisible sets whether mnemonics should be visible.
	SetMnemonicsVisible(mnemonicsVisible bool)
	// SetOffset sets the offset to use when calculating the position of the
	// popover.
	//
	// These values are used when preparing the [struct@Gdk.PopupLayout] for
	// positioning the popover.
	SetOffset(xOffset int, yOffset int)
	// SetPointingTo sets the rectangle that @popover points to.
	//
	// This is in the coordinate space of the @popover parent.
	SetPointingTo(rect *gdk.Rectangle)
	// SetPosition sets the preferred position for @popover to appear.
	//
	// If the @popover is currently visible, it will be immediately updated.
	//
	// This preference will be respected where possible, although on lack of
	// space (eg. if close to the window edges), the `GtkPopover` may choose to
	// appear on the opposite side.
	SetPosition(position PositionType)
}

// popover implements the Popover class.
type popover struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Native
	ShortcutManager
}

var _ Popover = (*popover)(nil)

// WrapPopover wraps a GObject to the right type. It is
// primarily used internally.
func WrapPopover(obj *externglib.Object) Popover {
	return popover{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		Native:           WrapNative(obj),
		ShortcutManager:  WrapShortcutManager(obj),
	}
}

func marshalPopover(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPopover(obj), nil
}

// NewPopover constructs a class Popover.
func NewPopover() Popover {
	var _cret C.GtkPopover // in

	_cret = C.gtk_popover_new()

	var _popover Popover // out

	_popover = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Popover)

	return _popover
}

// Autohide returns whether the popover is modal.
//
// See [method@Gtk.Popover.set_autohide] for the implications of this.
func (p popover) Autohide() bool {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_popover_get_autohide(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CascadePopdown returns whether the popover will close after a modal child
// is closed.
func (p popover) CascadePopdown() bool {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_popover_get_cascade_popdown(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Child gets the child widget of @popover.
func (p popover) Child() Widget {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var _cret *C.GtkWidget // in

	_cret = C.gtk_popover_get_child(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// HasArrow gets whether this popover is showing an arrow pointing at the
// widget that it is relative to.
func (p popover) HasArrow() bool {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_popover_get_has_arrow(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MnemonicsVisible gets whether mnemonics are visible.
func (p popover) MnemonicsVisible() bool {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_popover_get_mnemonics_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Offset gets the offset previous set with gtk_popover_set_offset().
func (p popover) Offset() (xOffset int, yOffset int) {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var _arg1 C.int // in
	var _arg2 C.int // in

	C.gtk_popover_get_offset(_arg0, &_arg1, &_arg2)

	var _xOffset int // out
	var _yOffset int // out

	_xOffset = (int)(_arg1)
	_yOffset = (int)(_arg2)

	return _xOffset, _yOffset
}

// PointingTo gets the rectangle that the popover points to.
//
// If a rectangle to point to has been set, this function will return true
// and fill in @rect with such rectangle, otherwise it will return false and
// fill in @rect with the parent widget coordinates.
func (p popover) PointingTo() (gdk.Rectangle, bool) {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var _rect gdk.Rectangle
	var _cret C.gboolean // in

	_cret = C.gtk_popover_get_pointing_to(_arg0, (*C.GdkRectangle)(unsafe.Pointer(&_rect)))

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _rect, _ok
}

// Position returns the preferred position of @popover.
func (p popover) Position() PositionType {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	var _cret C.GtkPositionType // in

	_cret = C.gtk_popover_get_position(_arg0)

	var _positionType PositionType // out

	_positionType = PositionType(_cret)

	return _positionType
}

// Popdown pops @popover down.
//
// This is different from a [method@Gtk.Widget.hide] call in that it may
// show the popover with a transition. If you want to hide the popover
// without a transition, just use [method@Gtk.Widget.hide].
func (p popover) Popdown() {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	C.gtk_popover_popdown(_arg0)
}

// Popup pops @popover up.
//
// This is different from a [method@Gtk.Widget.show() call in that it may
// show the popover with a transition. If you want to show the popover
// without a transition, just use [method@Gtk.Widget.show].
func (p popover) Popup() {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	C.gtk_popover_popup(_arg0)
}

// Present presents the popover to the user.
func (p popover) Present() {
	var _arg0 *C.GtkPopover // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))

	C.gtk_popover_present(_arg0)
}

// SetAutohide sets whether @popover is modal.
//
// A modal popover will grab the keyboard focus on it when being displayed.
// Clicking outside the popover area or pressing Esc will dismiss the
// popover.
//
// Called this function on an already showing popup with a new autohide
// value different from the current one, will cause the popup to be hidden.
func (p popover) SetAutohide(autohide bool) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	if autohide {
		_arg1 = C.TRUE
	}

	C.gtk_popover_set_autohide(_arg0, _arg1)
}

// SetCascadePopdown: if @cascade_popdown is true, the popover will be
// closed when a child modal popover is closed.
//
// If false, @popover will stay visible.
func (p popover) SetCascadePopdown(cascadePopdown bool) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	if cascadePopdown {
		_arg1 = C.TRUE
	}

	C.gtk_popover_set_cascade_popdown(_arg0, _arg1)
}

// SetChild sets the child widget of @popover.
func (p popover) SetChild(child Widget) {
	var _arg0 *C.GtkPopover // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_popover_set_child(_arg0, _arg1)
}

// SetDefaultWidget sets the default widget of a `GtkPopover`.
//
// The default widget is the widget that’s activated when the user presses
// Enter in a dialog (for example). This function sets or unsets the default
// widget for a `GtkPopover`.
func (p popover) SetDefaultWidget(widget Widget) {
	var _arg0 *C.GtkPopover // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_popover_set_default_widget(_arg0, _arg1)
}

// SetHasArrow sets whether this popover should draw an arrow pointing at
// the widget it is relative to.
func (p popover) SetHasArrow(hasArrow bool) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	if hasArrow {
		_arg1 = C.TRUE
	}

	C.gtk_popover_set_has_arrow(_arg0, _arg1)
}

// SetMnemonicsVisible sets whether mnemonics should be visible.
func (p popover) SetMnemonicsVisible(mnemonicsVisible bool) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	if mnemonicsVisible {
		_arg1 = C.TRUE
	}

	C.gtk_popover_set_mnemonics_visible(_arg0, _arg1)
}

// SetOffset sets the offset to use when calculating the position of the
// popover.
//
// These values are used when preparing the [struct@Gdk.PopupLayout] for
// positioning the popover.
func (p popover) SetOffset(xOffset int, yOffset int) {
	var _arg0 *C.GtkPopover // out
	var _arg1 C.int         // out
	var _arg2 C.int         // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = C.int(xOffset)
	_arg2 = C.int(yOffset)

	C.gtk_popover_set_offset(_arg0, _arg1, _arg2)
}

// SetPointingTo sets the rectangle that @popover points to.
//
// This is in the coordinate space of the @popover parent.
func (p popover) SetPointingTo(rect *gdk.Rectangle) {
	var _arg0 *C.GtkPopover   // out
	var _arg1 *C.GdkRectangle // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GdkRectangle)(unsafe.Pointer(rect.Native()))

	C.gtk_popover_set_pointing_to(_arg0, _arg1)
}

// SetPosition sets the preferred position for @popover to appear.
//
// If the @popover is currently visible, it will be immediately updated.
//
// This preference will be respected where possible, although on lack of
// space (eg. if close to the window edges), the `GtkPopover` may choose to
// appear on the opposite side.
func (p popover) SetPosition(position PositionType) {
	var _arg0 *C.GtkPopover     // out
	var _arg1 C.GtkPositionType // out

	_arg0 = (*C.GtkPopover)(unsafe.Pointer(p.Native()))
	_arg1 = (C.GtkPositionType)(position)

	C.gtk_popover_set_position(_arg0, _arg1)
}
