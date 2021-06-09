// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gobject/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_container_get_type()), F: marshalContainer},
	})
}

// Container: a GTK+ user interface is constructed by nesting widgets inside
// widgets. Container widgets are the inner nodes in the resulting tree of
// widgets: they contain other widgets. So, for example, you might have a Window
// containing a Frame containing a Label. If you wanted an image instead of a
// textual label inside the frame, you might replace the Label widget with a
// Image widget.
//
// There are two major kinds of container widgets in GTK+. Both are subclasses
// of the abstract GtkContainer base class.
//
// The first type of container widget has a single child widget and derives from
// Bin. These containers are decorators, which add some kind of functionality to
// the child. For example, a Button makes its child into a clickable button; a
// Frame draws a frame around its child and a Window places its child widget
// inside a top-level window.
//
// The second type of container can have more than one child; its purpose is to
// manage layout. This means that these containers assign sizes and positions to
// their children. For example, a HBox arranges its children in a horizontal
// row, and a Grid arranges the widgets it contains in a two-dimensional grid.
//
// For implementations of Container the virtual method ContainerClass.forall()
// is always required, since it's used for drawing and other internal operations
// on the children. If the Container implementation expect to have non internal
// children it's needed to implement both ContainerClass.add() and
// ContainerClass.remove(). If the GtkContainer implementation has internal
// children, they should be added with gtk_widget_set_parent() on init() and
// removed with gtk_widget_unparent() in the WidgetClass.destroy()
// implementation. See more about implementing custom widgets at
// https://wiki.gnome.org/HowDoI/CustomWidgets
//
//
// Height for width geometry management
//
// GTK+ uses a height-for-width (and width-for-height) geometry management
// system. Height-for-width means that a widget can change how much vertical
// space it needs, depending on the amount of horizontal space that it is given
// (and similar for width-for-height).
//
// There are some things to keep in mind when implementing container widgets
// that make use of GTK+’s height for width geometry management system. First,
// it’s important to note that a container must prioritize one of its
// dimensions, that is to say that a widget or container can only have a
// SizeRequestMode that is GTK_SIZE_REQUEST_HEIGHT_FOR_WIDTH or
// GTK_SIZE_REQUEST_WIDTH_FOR_HEIGHT. However, every widget and container must
// be able to respond to the APIs for both dimensions, i.e. even if a widget has
// a request mode that is height-for-width, it is possible that its parent will
// request its sizes using the width-for-height APIs.
//
// To ensure that everything works properly, here are some guidelines to follow
// when implementing height-for-width (or width-for-height) containers.
//
// Each request mode involves 2 virtual methods. Height-for-width apis run
// through gtk_widget_get_preferred_width() and then through
// gtk_widget_get_preferred_height_for_width(). When handling requests in the
// opposite SizeRequestMode it is important that every widget request at least
// enough space to display all of its content at all times.
//
// When gtk_widget_get_preferred_height() is called on a container that is
// height-for-width, the container must return the height for its minimum width.
// This is easily achieved by simply calling the reverse apis implemented for
// itself as follows:
//
//    static void
//    foo_container_get_preferred_width_for_height (GtkWidget *widget,
//                                                  gint for_height,
//                                                  gint *min_width,
//                                                  gint *nat_width)
//    {
//       if (i_am_in_height_for_width_mode)
//         {
//           GTK_WIDGET_GET_CLASS (widget)->get_preferred_width (widget,
//                                                               min_width,
//                                                               nat_width);
//         }
//       else
//         {
//           ... execute the real width-for-height request here based on
//           the required width of the children collectively if the
//           container were to be allocated the said height ...
//         }
//    }
//
// Height for width requests are generally implemented in terms of a virtual
// allocation of widgets in the input orientation. Assuming an height-for-width
// request mode, a container would implement the
// get_preferred_height_for_width() virtual function by first calling
// gtk_widget_get_preferred_width() for each of its children.
//
// For each potential group of children that are lined up horizontally, the
// values returned by gtk_widget_get_preferred_width() should be collected in an
// array of RequestedSize structures. Any child spacing should be removed from
// the input @for_width and then the collective size should be allocated using
// the gtk_distribute_natural_allocation() convenience function.
//
// The container will then move on to request the preferred height for each
// child by using gtk_widget_get_preferred_height_for_width() and using the
// sizes stored in the RequestedSize array.
//
// To allocate a height-for-width container, it’s again important to consider
// that a container must prioritize one dimension over the other. So if a
// container is a height-for-width container it must first allocate all widgets
// horizontally using a RequestedSize array and
// gtk_distribute_natural_allocation() and then add any extra space (if and
// where appropriate) for the widget to expand.
//
// After adding all the expand space, the container assumes it was allocated
// sufficient height to fit all of its content. At this time, the container must
// use the total horizontal sizes of each widget to request the height-for-width
// of each of its children and store the requests in a RequestedSize array for
// any widgets that stack vertically (for tabular containers this can be
// generalized into the heights and widths of rows and columns). The vertical
// space must then again be distributed using
// gtk_distribute_natural_allocation() while this time considering the allocated
// height of the widget minus any vertical spacing that the container adds. Then
// vertical expand space should be added where appropriate and available and the
// container should go on to actually allocating the child widgets.
//
// See [GtkWidget’s geometry management section][geometry-management] to learn
// more about implementing height-for-width geometry management for widgets.
//
//
// Child properties
//
// GtkContainer introduces child properties. These are object properties that
// are not specific to either the container or the contained widget, but rather
// to their relation. Typical examples of child properties are the position or
// pack-type of a widget which is contained in a Box.
//
// Use gtk_container_class_install_child_property() to install child properties
// for a container class and gtk_container_class_find_child_property() or
// gtk_container_class_list_child_properties() to get information about existing
// child properties.
//
// To set the value of a child property, use gtk_container_child_set_property(),
// gtk_container_child_set() or gtk_container_child_set_valist(). To obtain the
// value of a child property, use gtk_container_child_get_property(),
// gtk_container_child_get() or gtk_container_child_get_valist(). To emit
// notification about child property changes, use gtk_widget_child_notify().
//
//
// GtkContainer as GtkBuildable
//
// The GtkContainer implementation of the GtkBuildable interface supports a
// <packing> element for children, which can contain multiple <property>
// elements that specify child properties for the child.
//
// Since 2.16, child properties can also be marked as translatable using the
// same “translatable”, “comments” and “context” attributes that are used for
// regular properties.
//
// Since 3.16, containers can have a <focus-chain> element containing multiple
// <widget> elements, one for each child that should be added to the focus
// chain. The ”name” attribute gives the id of the widget.
//
// An example of these properties in UI definitions:
//
//    <object class="GtkBox">
//      <child>
//        <object class="GtkEntry" id="entry1"/>
//        <packing>
//          <property name="pack-type">start</property>
//        </packing>
//      </child>
//      <child>
//        <object class="GtkEntry" id="entry2"/>
//      </child>
//      <focus-chain>
//        <widget name="entry1"/>
//        <widget name="entry2"/>
//      </focus-chain>
//    </object>
type Container interface {
	Widget
	Buildable

	// Add adds @widget to @container. Typically used for simple containers such
	// as Window, Frame, or Button; for more complicated layout containers such
	// as Box or Grid, this function will pick default packing parameters that
	// may not be correct. So consider functions such as gtk_box_pack_start()
	// and gtk_grid_attach() as an alternative to gtk_container_add() in those
	// cases. A widget may be added to only one container at a time; you can’t
	// place the same widget inside two different containers.
	//
	// Note that some containers, such as ScrolledWindow or ListBox, may add
	// intermediate children between the added widget and the container.
	Add(widget Widget)

	CheckResize()
	// ChildGetProperty gets the value of a child property for @child and
	// @container.
	ChildGetProperty(child Widget, propertyName string, value **externglib.Value)
	// ChildNotify emits a Widget::child-notify signal for the [child
	// property][child-properties] @child_property on the child.
	//
	// This is an analogue of g_object_notify() for child properties.
	//
	// Also see gtk_widget_child_notify().
	ChildNotify(child Widget, childProperty string)
	// ChildNotifyByPspec emits a Widget::child-notify signal for the [child
	// property][child-properties] specified by @pspec on the child.
	//
	// This is an analogue of g_object_notify_by_pspec() for child properties.
	ChildNotifyByPspec(child Widget, pspec gobject.ParamSpec)
	// ChildSetProperty sets a child property for @child and @container.
	ChildSetProperty(child Widget, propertyName string, value **externglib.Value)
	// ChildType returns the type of the children supported by the container.
	//
	// Note that this may return G_TYPE_NONE to indicate that no more children
	// can be added, e.g. for a Paned which already has two children.
	ChildType() externglib.Type
	// Forall invokes @callback on each direct child of @container, including
	// children that are considered “internal” (implementation details of the
	// container). “Internal” children generally weren’t added by the user of
	// the container, but were added by the container implementation itself.
	//
	// Most applications should use gtk_container_foreach(), rather than
	// gtk_container_forall().
	Forall()
	// Foreach invokes @callback on each non-internal child of @container. See
	// gtk_container_forall() for details on what constitutes an “internal”
	// child. For all practical purposes, this function should iterate over
	// precisely those child widgets that were added to the container by the
	// application with explicit add() calls.
	//
	// It is permissible to remove the child from the @callback handler.
	//
	// Most applications should use gtk_container_foreach(), rather than
	// gtk_container_forall().
	Foreach()
	// BorderWidth retrieves the border width of the container. See
	// gtk_container_set_border_width().
	BorderWidth() uint
	// Children returns the container’s non-internal children. See
	// gtk_container_forall() for details on what constitutes an "internal"
	// child.
	Children() *glib.List
	// FocusChain retrieves the focus chain of the container, if one has been
	// set explicitly. If no focus chain has been explicitly set, GTK+ computes
	// the focus chain based on the positions of the children. In that case,
	// GTK+ stores nil in @focusable_widgets and returns false.
	FocusChain() (focusableWidgets *glib.List, ok bool)
	// FocusChild returns the current focus child widget inside @container. This
	// is not the currently focused widget. That can be obtained by calling
	// gtk_window_get_focus().
	FocusChild() Widget
	// FocusHAdjustment retrieves the horizontal focus adjustment for the
	// container. See gtk_container_set_focus_hadjustment ().
	FocusHAdjustment() Adjustment
	// FocusVAdjustment retrieves the vertical focus adjustment for the
	// container. See gtk_container_set_focus_vadjustment().
	FocusVAdjustment() Adjustment
	// PathForChild returns a newly created widget path representing all the
	// widget hierarchy from the toplevel down to and including @child.
	PathForChild(child Widget) *WidgetPath
	// ResizeMode returns the resize mode for the container. See
	// gtk_container_set_resize_mode ().
	ResizeMode() ResizeMode
	// PropagateDraw: when a container receives a call to the draw function, it
	// must send synthetic Widget::draw calls to all children that don’t have
	// their own Windows. This function provides a convenient way of doing this.
	// A container, when it receives a call to its Widget::draw function, calls
	// gtk_container_propagate_draw() once for each child, passing in the @cr
	// the container received.
	//
	// gtk_container_propagate_draw() takes care of translating the origin of
	// @cr, and deciding whether the draw needs to be sent to the child. It is a
	// convenient and optimized way of getting the same effect as calling
	// gtk_widget_draw() on the child directly.
	//
	// In most cases, a container can simply either inherit the Widget::draw
	// implementation from Container, or do some drawing and then chain to the
	// ::draw implementation from Container.
	PropagateDraw(child Widget, cr *cairo.Context)
	// Remove removes @widget from @container. @widget must be inside
	// @container. Note that @container will own a reference to @widget, and
	// that this may be the last reference held; so removing a widget from its
	// container can destroy that widget. If you want to use @widget again, you
	// need to add a reference to it before removing it from a container, using
	// g_object_ref(). If you don’t want to use @widget again it’s usually more
	// efficient to simply destroy it directly using gtk_widget_destroy() since
	// this will remove it from the container and help break any circular
	// reference count cycles.
	Remove(widget Widget)

	ResizeChildren()
	// SetBorderWidth sets the border width of the container.
	//
	// The border width of a container is the amount of space to leave around
	// the outside of the container. The only exception to this is Window;
	// because toplevel windows can’t leave space outside, they leave the space
	// inside. The border is added on all sides of the container. To add space
	// to only one side, use a specific Widget:margin property on the child
	// widget, for example Widget:margin-top.
	SetBorderWidth(borderWidth uint)
	// SetFocusChain sets a focus chain, overriding the one computed
	// automatically by GTK+.
	//
	// In principle each widget in the chain should be a descendant of the
	// container, but this is not enforced by this method, since it’s allowed to
	// set the focus chain before you pack the widgets, or have a widget in the
	// chain that isn’t always packed. The necessary checks are done when the
	// focus chain is actually traversed.
	SetFocusChain(focusableWidgets *glib.List)
	// SetFocusChild: sets, or unsets if @child is nil, the focused child of
	// @container.
	//
	// This function emits the GtkContainer::set_focus_child signal of
	// @container. Implementations of Container can override the default
	// behaviour by overriding the class closure of this signal.
	//
	// This is function is mostly meant to be used by widgets. Applications can
	// use gtk_widget_grab_focus() to manually set the focus to a specific
	// widget.
	SetFocusChild(child Widget)
	// SetFocusHAdjustment hooks up an adjustment to focus handling in a
	// container, so when a child of the container is focused, the adjustment is
	// scrolled to show that widget. This function sets the horizontal
	// alignment. See gtk_scrolled_window_get_hadjustment() for a typical way of
	// obtaining the adjustment and gtk_container_set_focus_vadjustment() for
	// setting the vertical adjustment.
	//
	// The adjustments have to be in pixel units and in the same coordinate
	// system as the allocation for immediate children of the container.
	SetFocusHAdjustment(adjustment Adjustment)
	// SetFocusVAdjustment hooks up an adjustment to focus handling in a
	// container, so when a child of the container is focused, the adjustment is
	// scrolled to show that widget. This function sets the vertical alignment.
	// See gtk_scrolled_window_get_vadjustment() for a typical way of obtaining
	// the adjustment and gtk_container_set_focus_hadjustment() for setting the
	// horizontal adjustment.
	//
	// The adjustments have to be in pixel units and in the same coordinate
	// system as the allocation for immediate children of the container.
	SetFocusVAdjustment(adjustment Adjustment)
	// SetReallocateRedraws sets the @reallocate_redraws flag of the container
	// to the given value.
	//
	// Containers requesting reallocation redraws get automatically redrawn if
	// any of their children changed allocation.
	SetReallocateRedraws(needsRedraws bool)
	// SetResizeMode sets the resize mode for the container.
	//
	// The resize mode of a container determines whether a resize request will
	// be passed to the container’s parent, queued for later execution or
	// executed immediately.
	SetResizeMode(resizeMode ResizeMode)
	// UnsetFocusChain removes a focus chain explicitly set with
	// gtk_container_set_focus_chain().
	UnsetFocusChain()
}

// container implements the Container interface.
type container struct {
	Widget
	Buildable
}

var _ Container = (*container)(nil)

// WrapContainer wraps a GObject to the right type. It is
// primarily used internally.
func WrapContainer(obj *externglib.Object) Container {
	return Container{
		Widget:    WrapWidget(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalContainer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapContainer(obj), nil
}

// Add adds @widget to @container. Typically used for simple containers such
// as Window, Frame, or Button; for more complicated layout containers such
// as Box or Grid, this function will pick default packing parameters that
// may not be correct. So consider functions such as gtk_box_pack_start()
// and gtk_grid_attach() as an alternative to gtk_container_add() in those
// cases. A widget may be added to only one container at a time; you can’t
// place the same widget inside two different containers.
//
// Note that some containers, such as ScrolledWindow or ListBox, may add
// intermediate children between the added widget and the container.
func (c container) Add(widget Widget) {
	var arg0 *C.GtkContainer
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_container_add(arg0, arg1)
}

func (c container) CheckResize() {
	var arg0 *C.GtkContainer

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))

	C.gtk_container_check_resize(arg0)
}

// ChildGetProperty gets the value of a child property for @child and
// @container.
func (c container) ChildGetProperty(child Widget, propertyName string, value **externglib.Value) {
	var arg0 *C.GtkContainer
	var arg1 *C.GtkWidget
	var arg2 *C.gchar
	var arg3 *C.GValue

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.GValue)(value.GValue)

	C.gtk_container_child_get_property(arg0, arg1, arg2, arg3)
}

// ChildNotify emits a Widget::child-notify signal for the [child
// property][child-properties] @child_property on the child.
//
// This is an analogue of g_object_notify() for child properties.
//
// Also see gtk_widget_child_notify().
func (c container) ChildNotify(child Widget, childProperty string) {
	var arg0 *C.GtkContainer
	var arg1 *C.GtkWidget
	var arg2 *C.gchar

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = (*C.gchar)(C.CString(childProperty))
	defer C.free(unsafe.Pointer(arg2))

	C.gtk_container_child_notify(arg0, arg1, arg2)
}

// ChildNotifyByPspec emits a Widget::child-notify signal for the [child
// property][child-properties] specified by @pspec on the child.
//
// This is an analogue of g_object_notify_by_pspec() for child properties.
func (c container) ChildNotifyByPspec(child Widget, pspec gobject.ParamSpec) {
	var arg0 *C.GtkContainer
	var arg1 *C.GtkWidget
	var arg2 *C.GParamSpec

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = (*C.GParamSpec)(unsafe.Pointer(pspec.Native()))

	C.gtk_container_child_notify_by_pspec(arg0, arg1, arg2)
}

// ChildSetProperty sets a child property for @child and @container.
func (c container) ChildSetProperty(child Widget, propertyName string, value **externglib.Value) {
	var arg0 *C.GtkContainer
	var arg1 *C.GtkWidget
	var arg2 *C.gchar
	var arg3 *C.GValue

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (*C.GValue)(value.GValue)

	C.gtk_container_child_set_property(arg0, arg1, arg2, arg3)
}

// ChildType returns the type of the children supported by the container.
//
// Note that this may return G_TYPE_NONE to indicate that no more children
// can be added, e.g. for a Paned which already has two children.
func (c container) ChildType() externglib.Type {
	var arg0 *C.GtkContainer

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))

	var cret C.GType

	cret = C.gtk_container_child_type(arg0)

	var gType externglib.Type

	gType = externglib.Type(cret)

	return gType
}

// Forall invokes @callback on each direct child of @container, including
// children that are considered “internal” (implementation details of the
// container). “Internal” children generally weren’t added by the user of
// the container, but were added by the container implementation itself.
//
// Most applications should use gtk_container_foreach(), rather than
// gtk_container_forall().
func (c container) Forall() {
	var arg0 *C.GtkContainer

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))

	C.gtk_container_forall(arg0)
}

// Foreach invokes @callback on each non-internal child of @container. See
// gtk_container_forall() for details on what constitutes an “internal”
// child. For all practical purposes, this function should iterate over
// precisely those child widgets that were added to the container by the
// application with explicit add() calls.
//
// It is permissible to remove the child from the @callback handler.
//
// Most applications should use gtk_container_foreach(), rather than
// gtk_container_forall().
func (c container) Foreach() {
	var arg0 *C.GtkContainer

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))

	C.gtk_container_foreach(arg0)
}

// BorderWidth retrieves the border width of the container. See
// gtk_container_set_border_width().
func (c container) BorderWidth() uint {
	var arg0 *C.GtkContainer

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))

	var cret C.guint

	cret = C.gtk_container_get_border_width(arg0)

	var guint uint

	guint = (uint)(cret)

	return guint
}

// Children returns the container’s non-internal children. See
// gtk_container_forall() for details on what constitutes an "internal"
// child.
func (c container) Children() *glib.List {
	var arg0 *C.GtkContainer

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))

	var cret *C.GList

	cret = C.gtk_container_get_children(arg0)

	var list *glib.List

	list = glib.WrapList(unsafe.Pointer(cret))
	runtime.SetFinalizer(list, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return list
}

// FocusChain retrieves the focus chain of the container, if one has been
// set explicitly. If no focus chain has been explicitly set, GTK+ computes
// the focus chain based on the positions of the children. In that case,
// GTK+ stores nil in @focusable_widgets and returns false.
func (c container) FocusChain() (focusableWidgets *glib.List, ok bool) {
	var arg0 *C.GtkContainer

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))

	var focusableWidgets *glib.List
	var cret C.gboolean

	cret = C.gtk_container_get_focus_chain(arg0, (**C.GList)(unsafe.Pointer(&focusableWidgets)))

	var ok bool

	if cret {
		ok = true
	}

	return focusableWidgets, ok
}

// FocusChild returns the current focus child widget inside @container. This
// is not the currently focused widget. That can be obtained by calling
// gtk_window_get_focus().
func (c container) FocusChild() Widget {
	var arg0 *C.GtkContainer

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))

	var cret *C.GtkWidget

	cret = C.gtk_container_get_focus_child(arg0)

	var widget Widget

	widget = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return widget
}

// FocusHAdjustment retrieves the horizontal focus adjustment for the
// container. See gtk_container_set_focus_hadjustment ().
func (c container) FocusHAdjustment() Adjustment {
	var arg0 *C.GtkContainer

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))

	var cret *C.GtkAdjustment

	cret = C.gtk_container_get_focus_hadjustment(arg0)

	var adjustment Adjustment

	adjustment = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Adjustment)

	return adjustment
}

// FocusVAdjustment retrieves the vertical focus adjustment for the
// container. See gtk_container_set_focus_vadjustment().
func (c container) FocusVAdjustment() Adjustment {
	var arg0 *C.GtkContainer

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))

	var cret *C.GtkAdjustment

	cret = C.gtk_container_get_focus_vadjustment(arg0)

	var adjustment Adjustment

	adjustment = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Adjustment)

	return adjustment
}

// PathForChild returns a newly created widget path representing all the
// widget hierarchy from the toplevel down to and including @child.
func (c container) PathForChild(child Widget) *WidgetPath {
	var arg0 *C.GtkContainer
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	var cret *C.GtkWidgetPath

	cret = C.gtk_container_get_path_for_child(arg0, arg1)

	var widgetPath *WidgetPath

	widgetPath = WrapWidgetPath(unsafe.Pointer(cret))
	runtime.SetFinalizer(widgetPath, func(v *WidgetPath) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return widgetPath
}

// ResizeMode returns the resize mode for the container. See
// gtk_container_set_resize_mode ().
func (c container) ResizeMode() ResizeMode {
	var arg0 *C.GtkContainer

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))

	var cret C.GtkResizeMode

	cret = C.gtk_container_get_resize_mode(arg0)

	var resizeMode ResizeMode

	resizeMode = ResizeMode(cret)

	return resizeMode
}

// PropagateDraw: when a container receives a call to the draw function, it
// must send synthetic Widget::draw calls to all children that don’t have
// their own Windows. This function provides a convenient way of doing this.
// A container, when it receives a call to its Widget::draw function, calls
// gtk_container_propagate_draw() once for each child, passing in the @cr
// the container received.
//
// gtk_container_propagate_draw() takes care of translating the origin of
// @cr, and deciding whether the draw needs to be sent to the child. It is a
// convenient and optimized way of getting the same effect as calling
// gtk_widget_draw() on the child directly.
//
// In most cases, a container can simply either inherit the Widget::draw
// implementation from Container, or do some drawing and then chain to the
// ::draw implementation from Container.
func (c container) PropagateDraw(child Widget, cr *cairo.Context) {
	var arg0 *C.GtkContainer
	var arg1 *C.GtkWidget
	var arg2 *C.cairo_t

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))

	C.gtk_container_propagate_draw(arg0, arg1, arg2)
}

// Remove removes @widget from @container. @widget must be inside
// @container. Note that @container will own a reference to @widget, and
// that this may be the last reference held; so removing a widget from its
// container can destroy that widget. If you want to use @widget again, you
// need to add a reference to it before removing it from a container, using
// g_object_ref(). If you don’t want to use @widget again it’s usually more
// efficient to simply destroy it directly using gtk_widget_destroy() since
// this will remove it from the container and help break any circular
// reference count cycles.
func (c container) Remove(widget Widget) {
	var arg0 *C.GtkContainer
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_container_remove(arg0, arg1)
}

func (c container) ResizeChildren() {
	var arg0 *C.GtkContainer

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))

	C.gtk_container_resize_children(arg0)
}

// SetBorderWidth sets the border width of the container.
//
// The border width of a container is the amount of space to leave around
// the outside of the container. The only exception to this is Window;
// because toplevel windows can’t leave space outside, they leave the space
// inside. The border is added on all sides of the container. To add space
// to only one side, use a specific Widget:margin property on the child
// widget, for example Widget:margin-top.
func (c container) SetBorderWidth(borderWidth uint) {
	var arg0 *C.GtkContainer
	var arg1 C.guint

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))
	arg1 = C.guint(borderWidth)

	C.gtk_container_set_border_width(arg0, arg1)
}

// SetFocusChain sets a focus chain, overriding the one computed
// automatically by GTK+.
//
// In principle each widget in the chain should be a descendant of the
// container, but this is not enforced by this method, since it’s allowed to
// set the focus chain before you pack the widgets, or have a widget in the
// chain that isn’t always packed. The necessary checks are done when the
// focus chain is actually traversed.
func (c container) SetFocusChain(focusableWidgets *glib.List) {
	var arg0 *C.GtkContainer
	var arg1 *C.GList

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GList)(unsafe.Pointer(focusableWidgets.Native()))

	C.gtk_container_set_focus_chain(arg0, arg1)
}

// SetFocusChild: sets, or unsets if @child is nil, the focused child of
// @container.
//
// This function emits the GtkContainer::set_focus_child signal of
// @container. Implementations of Container can override the default
// behaviour by overriding the class closure of this signal.
//
// This is function is mostly meant to be used by widgets. Applications can
// use gtk_widget_grab_focus() to manually set the focus to a specific
// widget.
func (c container) SetFocusChild(child Widget) {
	var arg0 *C.GtkContainer
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_container_set_focus_child(arg0, arg1)
}

// SetFocusHAdjustment hooks up an adjustment to focus handling in a
// container, so when a child of the container is focused, the adjustment is
// scrolled to show that widget. This function sets the horizontal
// alignment. See gtk_scrolled_window_get_hadjustment() for a typical way of
// obtaining the adjustment and gtk_container_set_focus_vadjustment() for
// setting the vertical adjustment.
//
// The adjustments have to be in pixel units and in the same coordinate
// system as the allocation for immediate children of the container.
func (c container) SetFocusHAdjustment(adjustment Adjustment) {
	var arg0 *C.GtkContainer
	var arg1 *C.GtkAdjustment

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_container_set_focus_hadjustment(arg0, arg1)
}

// SetFocusVAdjustment hooks up an adjustment to focus handling in a
// container, so when a child of the container is focused, the adjustment is
// scrolled to show that widget. This function sets the vertical alignment.
// See gtk_scrolled_window_get_vadjustment() for a typical way of obtaining
// the adjustment and gtk_container_set_focus_hadjustment() for setting the
// horizontal adjustment.
//
// The adjustments have to be in pixel units and in the same coordinate
// system as the allocation for immediate children of the container.
func (c container) SetFocusVAdjustment(adjustment Adjustment) {
	var arg0 *C.GtkContainer
	var arg1 *C.GtkAdjustment

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_container_set_focus_vadjustment(arg0, arg1)
}

// SetReallocateRedraws sets the @reallocate_redraws flag of the container
// to the given value.
//
// Containers requesting reallocation redraws get automatically redrawn if
// any of their children changed allocation.
func (c container) SetReallocateRedraws(needsRedraws bool) {
	var arg0 *C.GtkContainer
	var arg1 C.gboolean

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))
	if needsRedraws {
		arg1 = C.gboolean(1)
	}

	C.gtk_container_set_reallocate_redraws(arg0, arg1)
}

// SetResizeMode sets the resize mode for the container.
//
// The resize mode of a container determines whether a resize request will
// be passed to the container’s parent, queued for later execution or
// executed immediately.
func (c container) SetResizeMode(resizeMode ResizeMode) {
	var arg0 *C.GtkContainer
	var arg1 C.GtkResizeMode

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))
	arg1 = (C.GtkResizeMode)(resizeMode)

	C.gtk_container_set_resize_mode(arg0, arg1)
}

// UnsetFocusChain removes a focus chain explicitly set with
// gtk_container_set_focus_chain().
func (c container) UnsetFocusChain() {
	var arg0 *C.GtkContainer

	arg0 = (*C.GtkContainer)(unsafe.Pointer(c.Native()))

	C.gtk_container_unset_focus_chain(arg0)
}
