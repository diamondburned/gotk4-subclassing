// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
		{T: externglib.Type(C.gtk_tool_palette_get_type()), F: marshalToolPalette},
	})
}

// ToolPalette: a ToolPalette allows you to add ToolItems to a palette-like
// container with different categories and drag and drop support.
//
// A ToolPalette is created with a call to gtk_tool_palette_new().
//
// ToolItems cannot be added directly to a ToolPalette - instead they are added
// to a ToolItemGroup which can than be added to a ToolPalette. To add a
// ToolItemGroup to a ToolPalette, use gtk_container_add().
//
//    static void
//    passive_canvas_drag_data_received (GtkWidget        *widget,
//                                       GdkDragContext   *context,
//                                       gint              x,
//                                       gint              y,
//                                       GtkSelectionData *selection,
//                                       guint             info,
//                                       guint             time,
//                                       gpointer          data)
//    {
//      GtkWidget *palette;
//      GtkWidget *item;
//
//      // Get the dragged item
//      palette = gtk_widget_get_ancestor (gtk_drag_get_source_widget (context),
//                                         GTK_TYPE_TOOL_PALETTE);
//      if (palette != NULL)
//        item = gtk_tool_palette_get_drag_item (GTK_TOOL_PALETTE (palette),
//                                               selection);
//
//      // Do something with item
//    }
//
//    GtkWidget *target, palette;
//
//    palette = gtk_tool_palette_new ();
//    target = gtk_drawing_area_new ();
//
//    g_signal_connect (G_OBJECT (target), "drag-data-received",
//                      G_CALLBACK (passive_canvas_drag_data_received), NULL);
//    gtk_tool_palette_add_drag_dest (GTK_TOOL_PALETTE (palette), target,
//                                    GTK_DEST_DEFAULT_ALL,
//                                    GTK_TOOL_PALETTE_DRAG_ITEMS,
//                                    GDK_ACTION_COPY);
//
//
// CSS nodes
//
// GtkToolPalette has a single CSS node named toolpalette.
type ToolPalette interface {
	Container
	Buildable
	Orientable
	Scrollable

	// AddDragDest sets @palette as drag source (see
	// gtk_tool_palette_set_drag_source()) and sets @widget as a drag
	// destination for drags from @palette. See gtk_drag_dest_set().
	AddDragDest(p ToolPalette, widget Widget, flags DestDefaults, targets ToolPaletteDragTargets, actions gdk.DragAction)
	// DragItem: get the dragged item from the selection. This could be a
	// ToolItem or a ToolItemGroup.
	DragItem(p ToolPalette, selection *SelectionData)
	// DropGroup gets the group at position (x, y).
	DropGroup(p ToolPalette, x int, y int)
	// DropItem gets the item at position (x, y). See
	// gtk_tool_palette_get_drop_group().
	DropItem(p ToolPalette, x int, y int)
	// Exclusive gets whether @group is exclusive or not. See
	// gtk_tool_palette_set_exclusive().
	Exclusive(p ToolPalette, group ToolItemGroup) bool
	// Expand gets whether group should be given extra space. See
	// gtk_tool_palette_set_expand().
	Expand(p ToolPalette, group ToolItemGroup) bool
	// GroupPosition gets the position of @group in @palette as index. See
	// gtk_tool_palette_set_group_position().
	GroupPosition(p ToolPalette, group ToolItemGroup)
	// HAdjustment gets the horizontal adjustment of the tool palette.
	HAdjustment(p ToolPalette)
	// IconSize gets the size of icons in the tool palette. See
	// gtk_tool_palette_set_icon_size().
	IconSize(p ToolPalette)
	// Style gets the style (icons, text or both) of items in the tool palette.
	Style(p ToolPalette)
	// VAdjustment gets the vertical adjustment of the tool palette.
	VAdjustment(p ToolPalette)
	// SetDragSource sets the tool palette as a drag source. Enables all groups
	// and items in the tool palette as drag sources on button 1 and button 3
	// press with copy and move actions. See gtk_drag_source_set().
	SetDragSource(p ToolPalette, targets ToolPaletteDragTargets)
	// SetExclusive sets whether the group should be exclusive or not. If an
	// exclusive group is expanded all other groups are collapsed.
	SetExclusive(p ToolPalette, group ToolItemGroup, exclusive bool)
	// SetExpand sets whether the group should be given extra space.
	SetExpand(p ToolPalette, group ToolItemGroup, expand bool)
	// SetGroupPosition sets the position of the group as an index of the tool
	// palette. If position is 0 the group will become the first child, if
	// position is -1 it will become the last child.
	SetGroupPosition(p ToolPalette, group ToolItemGroup, position int)
	// SetIconSize sets the size of icons in the tool palette.
	SetIconSize(p ToolPalette, iconSize int)
	// SetStyle sets the style (text, icons or both) of items in the tool
	// palette.
	SetStyle(p ToolPalette, style ToolbarStyle)
	// UnsetIconSize unsets the tool palette icon size set with
	// gtk_tool_palette_set_icon_size(), so that user preferences will be used
	// to determine the icon size.
	UnsetIconSize(p ToolPalette)
	// UnsetStyle unsets a toolbar style set with gtk_tool_palette_set_style(),
	// so that user preferences will be used to determine the toolbar style.
	UnsetStyle(p ToolPalette)
}

// toolPalette implements the ToolPalette interface.
type toolPalette struct {
	Container
	Buildable
	Orientable
	Scrollable
}

var _ ToolPalette = (*toolPalette)(nil)

// WrapToolPalette wraps a GObject to the right type. It is
// primarily used internally.
func WrapToolPalette(obj *externglib.Object) ToolPalette {
	return ToolPalette{
		Container:  WrapContainer(obj),
		Buildable:  WrapBuildable(obj),
		Orientable: WrapOrientable(obj),
		Scrollable: WrapScrollable(obj),
	}
}

func marshalToolPalette(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapToolPalette(obj), nil
}

// NewToolPalette constructs a class ToolPalette.
func NewToolPalette() {
	C.gtk_tool_palette_new()
}

// AddDragDest sets @palette as drag source (see
// gtk_tool_palette_set_drag_source()) and sets @widget as a drag
// destination for drags from @palette. See gtk_drag_dest_set().
func (p toolPalette) AddDragDest(p ToolPalette, widget Widget, flags DestDefaults, targets ToolPaletteDragTargets, actions gdk.DragAction) {
	var arg0 *C.GtkToolPalette
	var arg1 *C.GtkWidget
	var arg2 C.GtkDestDefaults
	var arg3 C.GtkToolPaletteDragTargets
	var arg4 C.GdkDragAction

	arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	arg2 = (C.GtkDestDefaults)(flags)
	arg3 = (C.GtkToolPaletteDragTargets)(targets)
	arg4 = (C.GdkDragAction)(actions)

	C.gtk_tool_palette_add_drag_dest(arg0, arg1, arg2, arg3, arg4)
}

// DragItem: get the dragged item from the selection. This could be a
// ToolItem or a ToolItemGroup.
func (p toolPalette) DragItem(p ToolPalette, selection *SelectionData) {
	var arg0 *C.GtkToolPalette
	var arg1 *C.GtkSelectionData

	arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkSelectionData)(unsafe.Pointer(selection.Native()))

	C.gtk_tool_palette_get_drag_item(arg0, arg1)
}

// DropGroup gets the group at position (x, y).
func (p toolPalette) DropGroup(p ToolPalette, x int, y int) {
	var arg0 *C.GtkToolPalette
	var arg1 C.gint
	var arg2 C.gint

	arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	arg1 = C.gint(x)
	arg2 = C.gint(y)

	C.gtk_tool_palette_get_drop_group(arg0, arg1, arg2)
}

// DropItem gets the item at position (x, y). See
// gtk_tool_palette_get_drop_group().
func (p toolPalette) DropItem(p ToolPalette, x int, y int) {
	var arg0 *C.GtkToolPalette
	var arg1 C.gint
	var arg2 C.gint

	arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	arg1 = C.gint(x)
	arg2 = C.gint(y)

	C.gtk_tool_palette_get_drop_item(arg0, arg1, arg2)
}

// Exclusive gets whether @group is exclusive or not. See
// gtk_tool_palette_set_exclusive().
func (p toolPalette) Exclusive(p ToolPalette, group ToolItemGroup) bool {
	var arg0 *C.GtkToolPalette
	var arg1 *C.GtkToolItemGroup

	arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_tool_palette_get_exclusive(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// Expand gets whether group should be given extra space. See
// gtk_tool_palette_set_expand().
func (p toolPalette) Expand(p ToolPalette, group ToolItemGroup) bool {
	var arg0 *C.GtkToolPalette
	var arg1 *C.GtkToolItemGroup

	arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_tool_palette_get_expand(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// GroupPosition gets the position of @group in @palette as index. See
// gtk_tool_palette_set_group_position().
func (p toolPalette) GroupPosition(p ToolPalette, group ToolItemGroup) {
	var arg0 *C.GtkToolPalette
	var arg1 *C.GtkToolItemGroup

	arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))

	C.gtk_tool_palette_get_group_position(arg0, arg1)
}

// HAdjustment gets the horizontal adjustment of the tool palette.
func (p toolPalette) HAdjustment(p ToolPalette) {
	var arg0 *C.GtkToolPalette

	arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))

	C.gtk_tool_palette_get_hadjustment(arg0)
}

// IconSize gets the size of icons in the tool palette. See
// gtk_tool_palette_set_icon_size().
func (p toolPalette) IconSize(p ToolPalette) {
	var arg0 *C.GtkToolPalette

	arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))

	C.gtk_tool_palette_get_icon_size(arg0)
}

// Style gets the style (icons, text or both) of items in the tool palette.
func (p toolPalette) Style(p ToolPalette) {
	var arg0 *C.GtkToolPalette

	arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))

	C.gtk_tool_palette_get_style(arg0)
}

// VAdjustment gets the vertical adjustment of the tool palette.
func (p toolPalette) VAdjustment(p ToolPalette) {
	var arg0 *C.GtkToolPalette

	arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))

	C.gtk_tool_palette_get_vadjustment(arg0)
}

// SetDragSource sets the tool palette as a drag source. Enables all groups
// and items in the tool palette as drag sources on button 1 and button 3
// press with copy and move actions. See gtk_drag_source_set().
func (p toolPalette) SetDragSource(p ToolPalette, targets ToolPaletteDragTargets) {
	var arg0 *C.GtkToolPalette
	var arg1 C.GtkToolPaletteDragTargets

	arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	arg1 = (C.GtkToolPaletteDragTargets)(targets)

	C.gtk_tool_palette_set_drag_source(arg0, arg1)
}

// SetExclusive sets whether the group should be exclusive or not. If an
// exclusive group is expanded all other groups are collapsed.
func (p toolPalette) SetExclusive(p ToolPalette, group ToolItemGroup, exclusive bool) {
	var arg0 *C.GtkToolPalette
	var arg1 *C.GtkToolItemGroup
	var arg2 C.gboolean

	arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	if exclusive {
		arg2 = C.gboolean(1)
	}

	C.gtk_tool_palette_set_exclusive(arg0, arg1, arg2)
}

// SetExpand sets whether the group should be given extra space.
func (p toolPalette) SetExpand(p ToolPalette, group ToolItemGroup, expand bool) {
	var arg0 *C.GtkToolPalette
	var arg1 *C.GtkToolItemGroup
	var arg2 C.gboolean

	arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	if expand {
		arg2 = C.gboolean(1)
	}

	C.gtk_tool_palette_set_expand(arg0, arg1, arg2)
}

// SetGroupPosition sets the position of the group as an index of the tool
// palette. If position is 0 the group will become the first child, if
// position is -1 it will become the last child.
func (p toolPalette) SetGroupPosition(p ToolPalette, group ToolItemGroup, position int) {
	var arg0 *C.GtkToolPalette
	var arg1 *C.GtkToolItemGroup
	var arg2 C.gint

	arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	arg1 = (*C.GtkToolItemGroup)(unsafe.Pointer(group.Native()))
	arg2 = C.gint(position)

	C.gtk_tool_palette_set_group_position(arg0, arg1, arg2)
}

// SetIconSize sets the size of icons in the tool palette.
func (p toolPalette) SetIconSize(p ToolPalette, iconSize int) {
	var arg0 *C.GtkToolPalette
	var arg1 C.GtkIconSize

	arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	arg1 = C.GtkIconSize(iconSize)

	C.gtk_tool_palette_set_icon_size(arg0, arg1)
}

// SetStyle sets the style (text, icons or both) of items in the tool
// palette.
func (p toolPalette) SetStyle(p ToolPalette, style ToolbarStyle) {
	var arg0 *C.GtkToolPalette
	var arg1 C.GtkToolbarStyle

	arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))
	arg1 = (C.GtkToolbarStyle)(style)

	C.gtk_tool_palette_set_style(arg0, arg1)
}

// UnsetIconSize unsets the tool palette icon size set with
// gtk_tool_palette_set_icon_size(), so that user preferences will be used
// to determine the icon size.
func (p toolPalette) UnsetIconSize(p ToolPalette) {
	var arg0 *C.GtkToolPalette

	arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))

	C.gtk_tool_palette_unset_icon_size(arg0)
}

// UnsetStyle unsets a toolbar style set with gtk_tool_palette_set_style(),
// so that user preferences will be used to determine the toolbar style.
func (p toolPalette) UnsetStyle(p ToolPalette) {
	var arg0 *C.GtkToolPalette

	arg0 = (*C.GtkToolPalette)(unsafe.Pointer(p.Native()))

	C.gtk_tool_palette_unset_style(arg0)
}