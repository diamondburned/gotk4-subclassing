// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_radio_action_get_type()), F: marshalRadioAction},
	})
}

// RadioActionOverrider contains methods that are overridable .
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type RadioActionOverrider interface {
	Changed(current RadioAction)
}

// RadioAction is similar to RadioMenuItem. A number of radio actions can be
// linked together so that only one may be active at any one time.
type RadioAction interface {
	gextras.Objector

	// AsToggleAction casts the class to the ToggleAction interface.
	AsToggleAction() ToggleAction
	// AsBuildable casts the class to the Buildable interface.
	AsBuildable() Buildable

	// GetActive returns the checked state of the toggle action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from ToggleAction
	GetActive() bool
	// GetDrawAsRadio returns whether the action should have proxies like a
	// radio action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from ToggleAction
	GetDrawAsRadio() bool
	// SetActive sets the checked state on the toggle action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from ToggleAction
	SetActive(isActive bool)
	// SetDrawAsRadio sets whether the action should have proxies like a radio
	// action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from ToggleAction
	SetDrawAsRadio(drawAsRadio bool)
	// Toggled emits the “toggled” signal on the toggle action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from ToggleAction
	Toggled()
	// Activate emits the “activate” signal on the specified action, if it isn't
	// insensitive. This gets called by the proxy widgets when they get
	// activated.
	//
	// It can also be used to manually activate an action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	Activate()
	// BlockActivate: disable activation signals from the action
	//
	// This is needed when updating the state of your proxy Activatable widget
	// could result in calling gtk_action_activate(), this is a convenience
	// function to avoid recursing in those cases (updating toggle state for
	// instance).
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	BlockActivate()
	// ConnectAccelerator installs the accelerator for @action if @action has an
	// accel path and group. See gtk_action_set_accel_path() and
	// gtk_action_set_accel_group()
	//
	// Since multiple proxies may independently trigger the installation of the
	// accelerator, the @action counts the number of times this function has
	// been called and doesn’t remove the accelerator until
	// gtk_action_disconnect_accelerator() has been called as many times.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	ConnectAccelerator()
	// CreateIcon: this function is intended for use by action implementations
	// to create icons displayed in the proxy widgets.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	CreateIcon(iconSize int) Widget
	// CreateMenu: if @action provides a Menu widget as a submenu for the menu
	// item or the toolbar item it creates, this function returns an instance of
	// that menu.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	CreateMenu() Widget
	// CreateMenuItem creates a menu item widget that proxies for the given
	// action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	CreateMenuItem() Widget
	// CreateToolItem creates a toolbar item widget that proxies for the given
	// action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	CreateToolItem() Widget
	// DisconnectAccelerator undoes the effect of one call to
	// gtk_action_connect_accelerator().
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	DisconnectAccelerator()
	// GetAccelPath returns the accel path for this action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetAccelPath() string
	// GetAlwaysShowImage returns whether @action's menu item proxies will
	// always show their image, if available.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetAlwaysShowImage() bool
	// GetIconName gets the icon name of @action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetIconName() string
	// GetIsImportant checks whether @action is important or not
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetIsImportant() bool
	// GetLabel gets the label text of @action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetLabel() string
	// GetName returns the name of the action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetName() string
	// GetSensitive returns whether the action itself is sensitive. Note that
	// this doesn’t necessarily mean effective sensitivity. See
	// gtk_action_is_sensitive() for that.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetSensitive() bool
	// GetShortLabel gets the short label text of @action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetShortLabel() string
	// GetStockID gets the stock id of @action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetStockID() string
	// GetTooltip gets the tooltip text of @action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetTooltip() string
	// GetVisible returns whether the action itself is visible. Note that this
	// doesn’t necessarily mean effective visibility. See
	// gtk_action_is_sensitive() for that.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetVisible() bool
	// GetVisibleHorizontal checks whether @action is visible when horizontal
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetVisibleHorizontal() bool
	// GetVisibleVertical checks whether @action is visible when horizontal
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	GetVisibleVertical() bool
	// IsSensitive returns whether the action is effectively sensitive.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	IsSensitive() bool
	// IsVisible returns whether the action is effectively visible.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	IsVisible() bool
	// SetAccelGroup sets the AccelGroup in which the accelerator for this
	// action will be installed.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetAccelGroup(accelGroup AccelGroup)
	// SetAccelPath sets the accel path for this action. All proxy widgets
	// associated with the action will have this accel path, so that their
	// accelerators are consistent.
	//
	// Note that @accel_path string will be stored in a #GQuark. Therefore, if
	// you pass a static string, you can save some memory by interning it first
	// with g_intern_static_string().
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetAccelPath(accelPath string)
	// SetAlwaysShowImage sets whether @action's menu item proxies will ignore
	// the Settings:gtk-menu-images setting and always show their image, if
	// available.
	//
	// Use this if the menu item would be useless or hard to use without their
	// image.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetAlwaysShowImage(alwaysShow bool)
	// SetIconName sets the icon name on @action
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetIconName(iconName string)
	// SetIsImportant sets whether the action is important, this attribute is
	// used primarily by toolbar items to decide whether to show a label or not.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetIsImportant(isImportant bool)
	// SetLabel sets the label of @action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetLabel(label string)
	// SetSensitive sets the :sensitive property of the action to @sensitive.
	// Note that this doesn’t necessarily mean effective sensitivity. See
	// gtk_action_is_sensitive() for that.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetSensitive(sensitive bool)
	// SetShortLabel sets a shorter label text on @action.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetShortLabel(shortLabel string)
	// SetStockID sets the stock id on @action
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetStockID(stockId string)
	// SetTooltip sets the tooltip text on @action
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetTooltip(tooltip string)
	// SetVisible sets the :visible property of the action to @visible. Note
	// that this doesn’t necessarily mean effective visibility. See
	// gtk_action_is_visible() for that.
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetVisible(visible bool)
	// SetVisibleHorizontal sets whether @action is visible when horizontal
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetVisibleHorizontal(visibleHorizontal bool)
	// SetVisibleVertical sets whether @action is visible when vertical
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	SetVisibleVertical(visibleVertical bool)
	// UnblockActivate: reenable activation signals from the action
	//
	// Deprecated: since version 3.10.
	//
	// This method is inherited from Action
	UnblockActivate()
	// AddChild adds a child to @buildable. @type is an optional string
	// describing how the child should be added.
	//
	// This method is inherited from Buildable
	AddChild(builder Builder, child gextras.Objector, typ string)
	// ConstructChild constructs a child of @buildable with the name @name.
	//
	// Builder calls this function if a “constructor” has been specified in the
	// UI definition.
	//
	// This method is inherited from Buildable
	ConstructChild(builder Builder, name string) gextras.Objector
	// CustomFinished: this is similar to gtk_buildable_parser_finished() but is
	// called once for each custom tag handled by the @buildable.
	//
	// This method is inherited from Buildable
	CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{})
	// CustomTagEnd: this is called at the end of each custom element handled by
	// the buildable.
	//
	// This method is inherited from Buildable
	CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data interface{})
	// CustomTagStart: this is called for each unknown element under <child>.
	//
	// This method is inherited from Buildable
	CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool)
	// GetInternalChild: get the internal child called @childname of the
	// @buildable object.
	//
	// This method is inherited from Buildable
	GetInternalChild(builder Builder, childname string) gextras.Objector
	// GetName gets the name of the @buildable object.
	//
	// Builder sets the name based on the [GtkBuilder UI definition][BUILDER-UI]
	// used to construct the @buildable.
	//
	// This method is inherited from Buildable
	GetName() string
	// ParserFinished: called when the builder finishes the parsing of a
	// [GtkBuilder UI definition][BUILDER-UI]. Note that this will be called
	// once for each time gtk_builder_add_from_file() or
	// gtk_builder_add_from_string() is called on a builder.
	//
	// This method is inherited from Buildable
	ParserFinished(builder Builder)
	// SetBuildableProperty sets the property name @name to @value on the
	// @buildable object.
	//
	// This method is inherited from Buildable
	SetBuildableProperty(builder Builder, name string, value externglib.Value)
	// SetName sets the name of the @buildable object.
	//
	// This method is inherited from Buildable
	SetName(name string)
	// AddChild adds a child to @buildable. @type is an optional string
	// describing how the child should be added.
	//
	// This method is inherited from Buildable
	AddChild(builder Builder, child gextras.Objector, typ string)
	// ConstructChild constructs a child of @buildable with the name @name.
	//
	// Builder calls this function if a “constructor” has been specified in the
	// UI definition.
	//
	// This method is inherited from Buildable
	ConstructChild(builder Builder, name string) gextras.Objector
	// CustomFinished: this is similar to gtk_buildable_parser_finished() but is
	// called once for each custom tag handled by the @buildable.
	//
	// This method is inherited from Buildable
	CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{})
	// CustomTagEnd: this is called at the end of each custom element handled by
	// the buildable.
	//
	// This method is inherited from Buildable
	CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data interface{})
	// CustomTagStart: this is called for each unknown element under <child>.
	//
	// This method is inherited from Buildable
	CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool)
	// GetInternalChild: get the internal child called @childname of the
	// @buildable object.
	//
	// This method is inherited from Buildable
	GetInternalChild(builder Builder, childname string) gextras.Objector
	// GetName gets the name of the @buildable object.
	//
	// Builder sets the name based on the [GtkBuilder UI definition][BUILDER-UI]
	// used to construct the @buildable.
	//
	// This method is inherited from Buildable
	GetName() string
	// ParserFinished: called when the builder finishes the parsing of a
	// [GtkBuilder UI definition][BUILDER-UI]. Note that this will be called
	// once for each time gtk_builder_add_from_file() or
	// gtk_builder_add_from_string() is called on a builder.
	//
	// This method is inherited from Buildable
	ParserFinished(builder Builder)
	// SetBuildableProperty sets the property name @name to @value on the
	// @buildable object.
	//
	// This method is inherited from Buildable
	SetBuildableProperty(builder Builder, name string, value externglib.Value)
	// SetName sets the name of the @buildable object.
	//
	// This method is inherited from Buildable
	SetName(name string)
	// AddChild adds a child to @buildable. @type is an optional string
	// describing how the child should be added.
	//
	// This method is inherited from Buildable
	AddChild(builder Builder, child gextras.Objector, typ string)
	// ConstructChild constructs a child of @buildable with the name @name.
	//
	// Builder calls this function if a “constructor” has been specified in the
	// UI definition.
	//
	// This method is inherited from Buildable
	ConstructChild(builder Builder, name string) gextras.Objector
	// CustomFinished: this is similar to gtk_buildable_parser_finished() but is
	// called once for each custom tag handled by the @buildable.
	//
	// This method is inherited from Buildable
	CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{})
	// CustomTagEnd: this is called at the end of each custom element handled by
	// the buildable.
	//
	// This method is inherited from Buildable
	CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data interface{})
	// CustomTagStart: this is called for each unknown element under <child>.
	//
	// This method is inherited from Buildable
	CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool)
	// GetInternalChild: get the internal child called @childname of the
	// @buildable object.
	//
	// This method is inherited from Buildable
	GetInternalChild(builder Builder, childname string) gextras.Objector
	// GetName gets the name of the @buildable object.
	//
	// Builder sets the name based on the [GtkBuilder UI definition][BUILDER-UI]
	// used to construct the @buildable.
	//
	// This method is inherited from Buildable
	GetName() string
	// ParserFinished: called when the builder finishes the parsing of a
	// [GtkBuilder UI definition][BUILDER-UI]. Note that this will be called
	// once for each time gtk_builder_add_from_file() or
	// gtk_builder_add_from_string() is called on a builder.
	//
	// This method is inherited from Buildable
	ParserFinished(builder Builder)
	// SetBuildableProperty sets the property name @name to @value on the
	// @buildable object.
	//
	// This method is inherited from Buildable
	SetBuildableProperty(builder Builder, name string, value externglib.Value)
	// SetName sets the name of the @buildable object.
	//
	// This method is inherited from Buildable
	SetName(name string)

	// CurrentValue obtains the value property of the currently active member of
	// the group to which @action belongs.
	//
	// Deprecated: since version 3.10.
	CurrentValue() int
	// JoinGroup joins a radio action object to the group of another radio
	// action object.
	//
	// Use this in language bindings instead of the gtk_radio_action_get_group()
	// and gtk_radio_action_set_group() methods
	//
	// A common way to set up a group of radio actions is the following:
	//
	//     GtkRadioAction *action;
	//     GtkRadioAction *last_action;
	//
	//     while ( ...more actions to add... /)
	//       {
	//          action = gtk_radio_action_new (...);
	//
	//          gtk_radio_action_join_group (action, last_action);
	//          last_action = action;
	//       }
	//
	// Deprecated: since version 3.10.
	JoinGroup(groupSource RadioAction)
	// SetCurrentValue sets the currently active group member to the member with
	// value property @current_value.
	//
	// Deprecated: since version 3.10.
	SetCurrentValue(currentValue int)
}

// radioAction implements the RadioAction interface.
type radioAction struct {
	*externglib.Object
}

var _ RadioAction = (*radioAction)(nil)

// WrapRadioAction wraps a GObject to a type that implements
// interface RadioAction. It is primarily used internally.
func WrapRadioAction(obj *externglib.Object) RadioAction {
	return radioAction{obj}
}

func marshalRadioAction(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRadioAction(obj), nil
}

// NewRadioAction creates a new RadioAction object. To add the action to a
// ActionGroup and set the accelerator for the action, call
// gtk_action_group_add_action_with_accel().
//
// Deprecated: since version 3.10.
func NewRadioAction(name string, label string, tooltip string, stockId string, value int) RadioAction {
	var _arg1 *C.gchar          // out
	var _arg2 *C.gchar          // out
	var _arg3 *C.gchar          // out
	var _arg4 *C.gchar          // out
	var _arg5 C.gint            // out
	var _cret *C.GtkRadioAction // in

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(tooltip))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = C.gint(value)

	_cret = C.gtk_radio_action_new(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _radioAction RadioAction // out

	_radioAction = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(RadioAction)

	return _radioAction
}

func (r radioAction) AsToggleAction() ToggleAction {
	return WrapToggleAction(gextras.InternObject(r))
}

func (r radioAction) AsBuildable() Buildable {
	return WrapBuildable(gextras.InternObject(r))
}

func (a radioAction) GetActive() bool {
	return WrapToggleAction(gextras.InternObject(a)).GetActive()
}

func (a radioAction) GetDrawAsRadio() bool {
	return WrapToggleAction(gextras.InternObject(a)).GetDrawAsRadio()
}

func (a radioAction) SetActive(isActive bool) {
	WrapToggleAction(gextras.InternObject(a)).SetActive(isActive)
}

func (a radioAction) SetDrawAsRadio(drawAsRadio bool) {
	WrapToggleAction(gextras.InternObject(a)).SetDrawAsRadio(drawAsRadio)
}

func (a radioAction) Toggled() {
	WrapToggleAction(gextras.InternObject(a)).Toggled()
}

func (a radioAction) Activate() {
	WrapAction(gextras.InternObject(a)).Activate()
}

func (a radioAction) BlockActivate() {
	WrapAction(gextras.InternObject(a)).BlockActivate()
}

func (a radioAction) ConnectAccelerator() {
	WrapAction(gextras.InternObject(a)).ConnectAccelerator()
}

func (a radioAction) CreateIcon(iconSize int) Widget {
	return WrapAction(gextras.InternObject(a)).CreateIcon(iconSize)
}

func (a radioAction) CreateMenu() Widget {
	return WrapAction(gextras.InternObject(a)).CreateMenu()
}

func (a radioAction) CreateMenuItem() Widget {
	return WrapAction(gextras.InternObject(a)).CreateMenuItem()
}

func (a radioAction) CreateToolItem() Widget {
	return WrapAction(gextras.InternObject(a)).CreateToolItem()
}

func (a radioAction) DisconnectAccelerator() {
	WrapAction(gextras.InternObject(a)).DisconnectAccelerator()
}

func (a radioAction) GetAccelPath() string {
	return WrapAction(gextras.InternObject(a)).GetAccelPath()
}

func (a radioAction) GetAlwaysShowImage() bool {
	return WrapAction(gextras.InternObject(a)).GetAlwaysShowImage()
}

func (a radioAction) GetIconName() string {
	return WrapAction(gextras.InternObject(a)).GetIconName()
}

func (a radioAction) GetIsImportant() bool {
	return WrapAction(gextras.InternObject(a)).GetIsImportant()
}

func (a radioAction) GetLabel() string {
	return WrapAction(gextras.InternObject(a)).GetLabel()
}

func (a radioAction) GetName() string {
	return WrapAction(gextras.InternObject(a)).GetName()
}

func (a radioAction) GetSensitive() bool {
	return WrapAction(gextras.InternObject(a)).GetSensitive()
}

func (a radioAction) GetShortLabel() string {
	return WrapAction(gextras.InternObject(a)).GetShortLabel()
}

func (a radioAction) GetStockID() string {
	return WrapAction(gextras.InternObject(a)).GetStockID()
}

func (a radioAction) GetTooltip() string {
	return WrapAction(gextras.InternObject(a)).GetTooltip()
}

func (a radioAction) GetVisible() bool {
	return WrapAction(gextras.InternObject(a)).GetVisible()
}

func (a radioAction) GetVisibleHorizontal() bool {
	return WrapAction(gextras.InternObject(a)).GetVisibleHorizontal()
}

func (a radioAction) GetVisibleVertical() bool {
	return WrapAction(gextras.InternObject(a)).GetVisibleVertical()
}

func (a radioAction) IsSensitive() bool {
	return WrapAction(gextras.InternObject(a)).IsSensitive()
}

func (a radioAction) IsVisible() bool {
	return WrapAction(gextras.InternObject(a)).IsVisible()
}

func (a radioAction) SetAccelGroup(accelGroup AccelGroup) {
	WrapAction(gextras.InternObject(a)).SetAccelGroup(accelGroup)
}

func (a radioAction) SetAccelPath(accelPath string) {
	WrapAction(gextras.InternObject(a)).SetAccelPath(accelPath)
}

func (a radioAction) SetAlwaysShowImage(alwaysShow bool) {
	WrapAction(gextras.InternObject(a)).SetAlwaysShowImage(alwaysShow)
}

func (a radioAction) SetIconName(iconName string) {
	WrapAction(gextras.InternObject(a)).SetIconName(iconName)
}

func (a radioAction) SetIsImportant(isImportant bool) {
	WrapAction(gextras.InternObject(a)).SetIsImportant(isImportant)
}

func (a radioAction) SetLabel(label string) {
	WrapAction(gextras.InternObject(a)).SetLabel(label)
}

func (a radioAction) SetSensitive(sensitive bool) {
	WrapAction(gextras.InternObject(a)).SetSensitive(sensitive)
}

func (a radioAction) SetShortLabel(shortLabel string) {
	WrapAction(gextras.InternObject(a)).SetShortLabel(shortLabel)
}

func (a radioAction) SetStockID(stockId string) {
	WrapAction(gextras.InternObject(a)).SetStockID(stockId)
}

func (a radioAction) SetTooltip(tooltip string) {
	WrapAction(gextras.InternObject(a)).SetTooltip(tooltip)
}

func (a radioAction) SetVisible(visible bool) {
	WrapAction(gextras.InternObject(a)).SetVisible(visible)
}

func (a radioAction) SetVisibleHorizontal(visibleHorizontal bool) {
	WrapAction(gextras.InternObject(a)).SetVisibleHorizontal(visibleHorizontal)
}

func (a radioAction) SetVisibleVertical(visibleVertical bool) {
	WrapAction(gextras.InternObject(a)).SetVisibleVertical(visibleVertical)
}

func (a radioAction) UnblockActivate() {
	WrapAction(gextras.InternObject(a)).UnblockActivate()
}

func (b radioAction) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b radioAction) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b radioAction) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b radioAction) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b radioAction) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b radioAction) GetInternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).GetInternalChild(builder, childname)
}

func (b radioAction) GetName() string {
	return WrapBuildable(gextras.InternObject(b)).GetName()
}

func (b radioAction) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b radioAction) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b radioAction) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (b radioAction) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b radioAction) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b radioAction) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b radioAction) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b radioAction) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b radioAction) GetInternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).GetInternalChild(builder, childname)
}

func (b radioAction) GetName() string {
	return WrapBuildable(gextras.InternObject(b)).GetName()
}

func (b radioAction) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b radioAction) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b radioAction) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (b radioAction) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b radioAction) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b radioAction) CustomFinished(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomFinished(builder, child, tagname, data)
}

func (b radioAction) CustomTagEnd(builder Builder, child gextras.Objector, tagname string, data interface{}) {
	WrapBuildable(gextras.InternObject(b)).CustomTagEnd(builder, child, tagname, data)
}

func (b radioAction) CustomTagStart(builder Builder, child gextras.Objector, tagname string) (glib.MarkupParser, interface{}, bool) {
	return WrapBuildable(gextras.InternObject(b)).CustomTagStart(builder, child, tagname)
}

func (b radioAction) GetInternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).GetInternalChild(builder, childname)
}

func (b radioAction) GetName() string {
	return WrapBuildable(gextras.InternObject(b)).GetName()
}

func (b radioAction) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b radioAction) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b radioAction) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (a radioAction) CurrentValue() int {
	var _arg0 *C.GtkRadioAction // out
	var _cret C.gint            // in

	_arg0 = (*C.GtkRadioAction)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_radio_action_get_current_value(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (a radioAction) JoinGroup(groupSource RadioAction) {
	var _arg0 *C.GtkRadioAction // out
	var _arg1 *C.GtkRadioAction // out

	_arg0 = (*C.GtkRadioAction)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkRadioAction)(unsafe.Pointer(groupSource.Native()))

	C.gtk_radio_action_join_group(_arg0, _arg1)
}

func (a radioAction) SetCurrentValue(currentValue int) {
	var _arg0 *C.GtkRadioAction // out
	var _arg1 C.gint            // out

	_arg0 = (*C.GtkRadioAction)(unsafe.Pointer(a.Native()))
	_arg1 = C.gint(currentValue)

	C.gtk_radio_action_set_current_value(_arg0, _arg1)
}
