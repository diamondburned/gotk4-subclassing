// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_plug_get_type()), F: marshalPlug},
	})
}

// PlugOverrider contains methods that are overridable .
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type PlugOverrider interface {
	ObjectID() string
}

// Plug: see Socket
type Plug interface {
	gextras.Objector

	// AsObject casts the class to the Object interface.
	AsObject() Object
	// AsComponent casts the class to the Component interface.
	AsComponent() Component

	// AddRelationship adds a relationship of the specified type with the
	// specified target.
	//
	// This method is inherited from Object
	AddRelationship(relationship RelationType, target Object) bool
	// GetAccessibleID gets the accessible id of the accessible.
	//
	// This method is inherited from Object
	GetAccessibleID() string
	// GetDescription gets the accessible description of the accessible.
	//
	// This method is inherited from Object
	GetDescription() string
	// GetIndexInParent gets the 0-based index of this accessible in its parent;
	// returns -1 if the accessible does not have an accessible parent.
	//
	// This method is inherited from Object
	GetIndexInParent() int
	// GetLayer gets the layer of the accessible.
	//
	// Deprecated.
	//
	// This method is inherited from Object
	GetLayer() Layer
	// GetMDIZOrder gets the zorder of the accessible. The value G_MININT will
	// be returned if the layer of the accessible is not ATK_LAYER_MDI.
	//
	// Deprecated.
	//
	// This method is inherited from Object
	GetMDIZOrder() int
	// GetNAccessibleChildren gets the number of accessible children of the
	// accessible.
	//
	// This method is inherited from Object
	GetNAccessibleChildren() int
	// GetName gets the accessible name of the accessible.
	//
	// This method is inherited from Object
	GetName() string
	// GetObjectLocale gets a UTF-8 string indicating the POSIX-style
	// LC_MESSAGES locale of @accessible.
	//
	// This method is inherited from Object
	GetObjectLocale() string
	// GetParent gets the accessible parent of the accessible. By default this
	// is the one assigned with atk_object_set_parent(), but it is assumed that
	// ATK implementors have ways to get the parent of the object without the
	// need of assigning it manually with atk_object_set_parent(), and will
	// return it with this method.
	//
	// If you are only interested on the parent assigned with
	// atk_object_set_parent(), use atk_object_peek_parent().
	//
	// This method is inherited from Object
	GetParent() Object
	// GetRole gets the role of the accessible.
	//
	// This method is inherited from Object
	GetRole() Role
	// Initialize: this function is called when implementing subclasses of
	// Object. It does initialization required for the new object. It is
	// intended that this function should called only in the ..._new() functions
	// used to create an instance of a subclass of Object
	//
	// This method is inherited from Object
	Initialize(data interface{})
	// PeekParent gets the accessible parent of the accessible, if it has been
	// manually assigned with atk_object_set_parent. Otherwise, this function
	// returns nil.
	//
	// This method is intended as an utility for ATK implementors, and not to be
	// exposed to accessible tools. See atk_object_get_parent() for further
	// reference.
	//
	// This method is inherited from Object
	PeekParent() Object
	// RefAccessibleChild gets a reference to the specified accessible child of
	// the object. The accessible children are 0-based so the first accessible
	// child is at index 0, the second at index 1 and so on.
	//
	// This method is inherited from Object
	RefAccessibleChild(i int) Object
	// RefRelationSet gets the RelationSet associated with the object.
	//
	// This method is inherited from Object
	RefRelationSet() RelationSet
	// RefStateSet gets a reference to the state set of the accessible; the
	// caller must unreference it when it is no longer needed.
	//
	// This method is inherited from Object
	RefStateSet() StateSet
	// RemovePropertyChangeHandler removes a property change handler.
	//
	// Deprecated: since version 2.12.
	//
	// This method is inherited from Object
	RemovePropertyChangeHandler(handlerId uint)
	// RemoveRelationship removes a relationship of the specified type with the
	// specified target.
	//
	// This method is inherited from Object
	RemoveRelationship(relationship RelationType, target Object) bool
	// SetAccessibleID sets the accessible ID of the accessible. This is not
	// meant to be presented to the user, but to be an ID which is stable over
	// application development. Typically, this is the gtkbuilder ID. Such an ID
	// will be available for instance to identify a given well-known accessible
	// object for tailored screen reading, or for automatic regression testing.
	//
	// This method is inherited from Object
	SetAccessibleID(name string)
	// SetDescription sets the accessible description of the accessible. You
	// can't set the description to NULL. This is reserved for the initial
	// value. In this aspect NULL is similar to ATK_ROLE_UNKNOWN. If you want to
	// set the name to a empty value you can use "".
	//
	// This method is inherited from Object
	SetDescription(description string)
	// SetName sets the accessible name of the accessible. You can't set the
	// name to NULL. This is reserved for the initial value. In this aspect NULL
	// is similar to ATK_ROLE_UNKNOWN. If you want to set the name to a empty
	// value you can use "".
	//
	// This method is inherited from Object
	SetName(name string)
	// SetParent sets the accessible parent of the accessible. @parent can be
	// NULL.
	//
	// This method is inherited from Object
	SetParent(parent Object)
	// SetRole sets the role of the accessible.
	//
	// This method is inherited from Object
	SetRole(role Role)
	// Contains checks whether the specified point is within the extent of the
	// @component.
	//
	// Toolkit implementor note: ATK provides a default implementation for this
	// virtual method. In general there are little reason to re-implement it.
	//
	// This method is inherited from Component
	Contains(x int, y int, coordType CoordType) bool
	// GetAlpha returns the alpha value (i.e. the opacity) for this @component,
	// on a scale from 0 (fully transparent) to 1.0 (fully opaque).
	//
	// This method is inherited from Component
	GetAlpha() float64
	// GetExtents gets the rectangle which gives the extent of the @component.
	//
	// If the extent can not be obtained (e.g. a non-embedded plug or missing
	// support), all of x, y, width, height are set to -1.
	//
	// This method is inherited from Component
	GetExtents(coordType CoordType) (x int, y int, width int, height int)
	// GetLayer gets the layer of the component.
	//
	// This method is inherited from Component
	GetLayer() Layer
	// GetMDIZOrder gets the zorder of the component. The value G_MININT will be
	// returned if the layer of the component is not ATK_LAYER_MDI or
	// ATK_LAYER_WINDOW.
	//
	// This method is inherited from Component
	GetMDIZOrder() int
	// GetPosition gets the position of @component in the form of a point
	// specifying @component's top-left corner.
	//
	// If the position can not be obtained (e.g. a non-embedded plug or missing
	// support), x and y are set to -1.
	//
	// Deprecated.
	//
	// This method is inherited from Component
	GetPosition(coordType CoordType) (x int, y int)
	// GetSize gets the size of the @component in terms of width and height.
	//
	// If the size can not be obtained (e.g. a non-embedded plug or missing
	// support), width and height are set to -1.
	//
	// Deprecated.
	//
	// This method is inherited from Component
	GetSize() (width int, height int)
	// GrabFocus grabs focus for this @component.
	//
	// This method is inherited from Component
	GrabFocus() bool
	// RefAccessibleAtPoint gets a reference to the accessible child, if one
	// exists, at the coordinate point specified by @x and @y.
	//
	// This method is inherited from Component
	RefAccessibleAtPoint(x int, y int, coordType CoordType) Object
	// RemoveFocusHandler: remove the handler specified by @handler_id from the
	// list of functions to be executed when this object receives focus events
	// (in or out).
	//
	// Deprecated: since version 2.9.4.
	//
	// This method is inherited from Component
	RemoveFocusHandler(handlerId uint)
	// ScrollTo makes @component visible on the screen by scrolling all
	// necessary parents.
	//
	// Contrary to atk_component_set_position, this does not actually move
	// @component in its parent, this only makes the parents scroll so that the
	// object shows up on the screen, given its current position within the
	// parents.
	//
	// This method is inherited from Component
	ScrollTo(typ ScrollType) bool
	// ScrollToPoint: move the top-left of @component to a given position of the
	// screen by scrolling all necessary parents.
	//
	// This method is inherited from Component
	ScrollToPoint(coords CoordType, x int, y int) bool
	// SetExtents sets the extents of @component.
	//
	// This method is inherited from Component
	SetExtents(x int, y int, width int, height int, coordType CoordType) bool
	// SetPosition sets the position of @component.
	//
	// Contrary to atk_component_scroll_to, this does not trigger any scrolling,
	// this just moves @component in its parent.
	//
	// This method is inherited from Component
	SetPosition(x int, y int, coordType CoordType) bool
	// SetSize: set the size of the @component in terms of width and height.
	//
	// This method is inherited from Component
	SetSize(width int, height int) bool

	// ID gets the unique ID of an Plug object, which can be used to embed
	// inside of an Socket using atk_socket_embed().
	//
	// Internally, this calls a class function that should be registered by the
	// IPC layer (usually at-spi2-atk). The implementor of an Plug object should
	// call this function (after atk-bridge is loaded) and pass the value to the
	// process implementing the Socket, so it could embed the plug.
	ID() string
	// SetChild sets @child as accessible child of @plug and @plug as accessible
	// parent of @child. @child can be NULL.
	//
	// In some cases, one can not use the AtkPlug type directly as accessible
	// object for the toplevel widget of the application. For instance in the
	// gtk case, GtkPlugAccessible can not inherit both from GtkWindowAccessible
	// and from AtkPlug. In such a case, one can create, in addition to the
	// standard accessible object for the toplevel widget, an AtkPlug object,
	// and make the former the child of the latter by calling
	// atk_plug_set_child().
	SetChild(child Object)
}

// plug implements the Plug interface.
type plug struct {
	*externglib.Object
}

var _ Plug = (*plug)(nil)

// WrapPlug wraps a GObject to a type that implements
// interface Plug. It is primarily used internally.
func WrapPlug(obj *externglib.Object) Plug {
	return plug{obj}
}

func marshalPlug(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPlug(obj), nil
}

// NewPlug creates a new Plug instance.
func NewPlug() Plug {
	var _cret *C.AtkObject // in

	_cret = C.atk_plug_new()

	var _plug Plug // out

	_plug = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Plug)

	return _plug
}

func (p plug) AsObject() Object {
	return WrapObject(gextras.InternObject(p))
}

func (p plug) AsComponent() Component {
	return WrapComponent(gextras.InternObject(p))
}

func (o plug) AddRelationship(relationship RelationType, target Object) bool {
	return WrapObject(gextras.InternObject(o)).AddRelationship(relationship, target)
}

func (a plug) GetAccessibleID() string {
	return WrapObject(gextras.InternObject(a)).GetAccessibleID()
}

func (a plug) GetDescription() string {
	return WrapObject(gextras.InternObject(a)).GetDescription()
}

func (a plug) GetIndexInParent() int {
	return WrapObject(gextras.InternObject(a)).GetIndexInParent()
}

func (a plug) GetLayer() Layer {
	return WrapObject(gextras.InternObject(a)).GetLayer()
}

func (a plug) GetMDIZOrder() int {
	return WrapObject(gextras.InternObject(a)).GetMDIZOrder()
}

func (a plug) GetNAccessibleChildren() int {
	return WrapObject(gextras.InternObject(a)).GetNAccessibleChildren()
}

func (a plug) GetName() string {
	return WrapObject(gextras.InternObject(a)).GetName()
}

func (a plug) GetObjectLocale() string {
	return WrapObject(gextras.InternObject(a)).GetObjectLocale()
}

func (a plug) GetParent() Object {
	return WrapObject(gextras.InternObject(a)).GetParent()
}

func (a plug) GetRole() Role {
	return WrapObject(gextras.InternObject(a)).GetRole()
}

func (a plug) Initialize(data interface{}) {
	WrapObject(gextras.InternObject(a)).Initialize(data)
}

func (a plug) PeekParent() Object {
	return WrapObject(gextras.InternObject(a)).PeekParent()
}

func (a plug) RefAccessibleChild(i int) Object {
	return WrapObject(gextras.InternObject(a)).RefAccessibleChild(i)
}

func (a plug) RefRelationSet() RelationSet {
	return WrapObject(gextras.InternObject(a)).RefRelationSet()
}

func (a plug) RefStateSet() StateSet {
	return WrapObject(gextras.InternObject(a)).RefStateSet()
}

func (a plug) RemovePropertyChangeHandler(handlerId uint) {
	WrapObject(gextras.InternObject(a)).RemovePropertyChangeHandler(handlerId)
}

func (o plug) RemoveRelationship(relationship RelationType, target Object) bool {
	return WrapObject(gextras.InternObject(o)).RemoveRelationship(relationship, target)
}

func (a plug) SetAccessibleID(name string) {
	WrapObject(gextras.InternObject(a)).SetAccessibleID(name)
}

func (a plug) SetDescription(description string) {
	WrapObject(gextras.InternObject(a)).SetDescription(description)
}

func (a plug) SetName(name string) {
	WrapObject(gextras.InternObject(a)).SetName(name)
}

func (a plug) SetParent(parent Object) {
	WrapObject(gextras.InternObject(a)).SetParent(parent)
}

func (a plug) SetRole(role Role) {
	WrapObject(gextras.InternObject(a)).SetRole(role)
}

func (c plug) Contains(x int, y int, coordType CoordType) bool {
	return WrapComponent(gextras.InternObject(c)).Contains(x, y, coordType)
}

func (c plug) GetAlpha() float64 {
	return WrapComponent(gextras.InternObject(c)).GetAlpha()
}

func (c plug) GetExtents(coordType CoordType) (x int, y int, width int, height int) {
	return WrapComponent(gextras.InternObject(c)).GetExtents(coordType)
}

func (c plug) GetLayer() Layer {
	return WrapComponent(gextras.InternObject(c)).GetLayer()
}

func (c plug) GetMDIZOrder() int {
	return WrapComponent(gextras.InternObject(c)).GetMDIZOrder()
}

func (c plug) GetPosition(coordType CoordType) (x int, y int) {
	return WrapComponent(gextras.InternObject(c)).GetPosition(coordType)
}

func (c plug) GetSize() (width int, height int) {
	return WrapComponent(gextras.InternObject(c)).GetSize()
}

func (c plug) GrabFocus() bool {
	return WrapComponent(gextras.InternObject(c)).GrabFocus()
}

func (c plug) RefAccessibleAtPoint(x int, y int, coordType CoordType) Object {
	return WrapComponent(gextras.InternObject(c)).RefAccessibleAtPoint(x, y, coordType)
}

func (c plug) RemoveFocusHandler(handlerId uint) {
	WrapComponent(gextras.InternObject(c)).RemoveFocusHandler(handlerId)
}

func (c plug) ScrollTo(typ ScrollType) bool {
	return WrapComponent(gextras.InternObject(c)).ScrollTo(typ)
}

func (c plug) ScrollToPoint(coords CoordType, x int, y int) bool {
	return WrapComponent(gextras.InternObject(c)).ScrollToPoint(coords, x, y)
}

func (c plug) SetExtents(x int, y int, width int, height int, coordType CoordType) bool {
	return WrapComponent(gextras.InternObject(c)).SetExtents(x, y, width, height, coordType)
}

func (c plug) SetPosition(x int, y int, coordType CoordType) bool {
	return WrapComponent(gextras.InternObject(c)).SetPosition(x, y, coordType)
}

func (c plug) SetSize(width int, height int) bool {
	return WrapComponent(gextras.InternObject(c)).SetSize(width, height)
}

func (p plug) ID() string {
	var _arg0 *C.AtkPlug // out
	var _cret *C.gchar   // in

	_arg0 = (*C.AtkPlug)(unsafe.Pointer(p.Native()))

	_cret = C.atk_plug_get_id(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (p plug) SetChild(child Object) {
	var _arg0 *C.AtkPlug   // out
	var _arg1 *C.AtkObject // out

	_arg0 = (*C.AtkPlug)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.AtkObject)(unsafe.Pointer(child.Native()))

	C.atk_plug_set_child(_arg0, _arg1)
}
