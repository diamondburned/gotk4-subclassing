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
		{T: externglib.Type(C.gtk_tool_item_group_get_type()), F: marshalToolItemGroup},
	})
}

// ToolItemGroup: a ToolItemGroup is used together with ToolPalette to add
// ToolItems to a palette like container with different categories and drag and
// drop support.
//
//
// CSS nodes
//
// GtkToolItemGroup has a single CSS node named toolitemgroup.
type ToolItemGroup interface {
	Container
	Buildable
	ToolShell

	// Collapsed gets whether @group is collapsed or expanded.
	Collapsed(g ToolItemGroup) bool
	// DropItem gets the tool item at position (x, y).
	DropItem(g ToolItemGroup, x int, y int)
	// Ellipsize gets the ellipsization mode of @group.
	Ellipsize(g ToolItemGroup)
	// HeaderRelief gets the relief mode of the header button of @group.
	HeaderRelief(g ToolItemGroup)
	// ItemPosition gets the position of @item in @group as index.
	ItemPosition(g ToolItemGroup, item ToolItem)
	// Label gets the label of @group.
	Label(g ToolItemGroup)
	// LabelWidget gets the label widget of @group. See
	// gtk_tool_item_group_set_label_widget().
	LabelWidget(g ToolItemGroup)
	// NItems gets the number of tool items in @group.
	NItems(g ToolItemGroup)
	// NthItem gets the tool item at @index in group.
	NthItem(g ToolItemGroup, index uint)
	// Insert inserts @item at @position in the list of children of @group.
	Insert(g ToolItemGroup, item ToolItem, position int)
	// SetCollapsed sets whether the @group should be collapsed or expanded.
	SetCollapsed(g ToolItemGroup, collapsed bool)
	// SetEllipsize sets the ellipsization mode which should be used by labels
	// in @group.
	SetEllipsize(g ToolItemGroup, ellipsize pango.EllipsizeMode)
	// SetHeaderRelief: set the button relief of the group header. See
	// gtk_button_set_relief() for details.
	SetHeaderRelief(g ToolItemGroup, style ReliefStyle)
	// SetItemPosition sets the position of @item in the list of children of
	// @group.
	SetItemPosition(g ToolItemGroup, item ToolItem, position int)
	// SetLabel sets the label of the tool item group. The label is displayed in
	// the header of the group.
	SetLabel(g ToolItemGroup, label string)
	// SetLabelWidget sets the label of the tool item group. The label widget is
	// displayed in the header of the group, in place of the usual label.
	SetLabelWidget(g ToolItemGroup, labelWidget Widget)
}

// toolItemGroup implements the ToolItemGroup interface.
type toolItemGroup struct {
	Container
	Buildable
	ToolShell
}

var _ ToolItemGroup = (*toolItemGroup)(nil)

// WrapToolItemGroup wraps a GObject to the right type. It is
// primarily used internally.
func WrapToolItemGroup(obj *externglib.Object) ToolItemGroup {
	return ToolItemGroup{
		Container: WrapContainer(obj),
		Buildable: WrapBuildable(obj),
		ToolShell: WrapToolShell(obj),
	}
}

func marshalToolItemGroup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapToolItemGroup(obj), nil
}

// NewToolItemGroup constructs a class ToolItemGroup.
func NewToolItemGroup(label string) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_tool_item_group_new(arg1)
}

// Collapsed gets whether @group is collapsed or expanded.
func (g toolItemGroup) Collapsed(g ToolItemGroup) bool {
	var arg0 *C.GtkToolItemGroup

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_tool_item_group_get_collapsed(arg0)

	if cret {
		ok = true
	}

	return ok
}

// DropItem gets the tool item at position (x, y).
func (g toolItemGroup) DropItem(g ToolItemGroup, x int, y int) {
	var arg0 *C.GtkToolItemGroup
	var arg1 C.gint
	var arg2 C.gint

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	arg1 = C.gint(x)
	arg2 = C.gint(y)

	C.gtk_tool_item_group_get_drop_item(arg0, arg1, arg2)
}

// Ellipsize gets the ellipsization mode of @group.
func (g toolItemGroup) Ellipsize(g ToolItemGroup) {
	var arg0 *C.GtkToolItemGroup

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))

	C.gtk_tool_item_group_get_ellipsize(arg0)
}

// HeaderRelief gets the relief mode of the header button of @group.
func (g toolItemGroup) HeaderRelief(g ToolItemGroup) {
	var arg0 *C.GtkToolItemGroup

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))

	C.gtk_tool_item_group_get_header_relief(arg0)
}

// ItemPosition gets the position of @item in @group as index.
func (g toolItemGroup) ItemPosition(g ToolItemGroup, item ToolItem) {
	var arg0 *C.GtkToolItemGroup
	var arg1 *C.GtkToolItem

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	arg1 = (*C.GtkToolItem)(unsafe.Pointer(item.Native()))

	C.gtk_tool_item_group_get_item_position(arg0, arg1)
}

// Label gets the label of @group.
func (g toolItemGroup) Label(g ToolItemGroup) {
	var arg0 *C.GtkToolItemGroup

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))

	C.gtk_tool_item_group_get_label(arg0)
}

// LabelWidget gets the label widget of @group. See
// gtk_tool_item_group_set_label_widget().
func (g toolItemGroup) LabelWidget(g ToolItemGroup) {
	var arg0 *C.GtkToolItemGroup

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))

	C.gtk_tool_item_group_get_label_widget(arg0)
}

// NItems gets the number of tool items in @group.
func (g toolItemGroup) NItems(g ToolItemGroup) {
	var arg0 *C.GtkToolItemGroup

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))

	C.gtk_tool_item_group_get_n_items(arg0)
}

// NthItem gets the tool item at @index in group.
func (g toolItemGroup) NthItem(g ToolItemGroup, index uint) {
	var arg0 *C.GtkToolItemGroup
	var arg1 C.guint

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	arg1 = C.guint(index)

	C.gtk_tool_item_group_get_nth_item(arg0, arg1)
}

// Insert inserts @item at @position in the list of children of @group.
func (g toolItemGroup) Insert(g ToolItemGroup, item ToolItem, position int) {
	var arg0 *C.GtkToolItemGroup
	var arg1 *C.GtkToolItem
	var arg2 C.gint

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	arg1 = (*C.GtkToolItem)(unsafe.Pointer(item.Native()))
	arg2 = C.gint(position)

	C.gtk_tool_item_group_insert(arg0, arg1, arg2)
}

// SetCollapsed sets whether the @group should be collapsed or expanded.
func (g toolItemGroup) SetCollapsed(g ToolItemGroup, collapsed bool) {
	var arg0 *C.GtkToolItemGroup
	var arg1 C.gboolean

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	if collapsed {
		arg1 = C.gboolean(1)
	}

	C.gtk_tool_item_group_set_collapsed(arg0, arg1)
}

// SetEllipsize sets the ellipsization mode which should be used by labels
// in @group.
func (g toolItemGroup) SetEllipsize(g ToolItemGroup, ellipsize pango.EllipsizeMode) {
	var arg0 *C.GtkToolItemGroup
	var arg1 C.PangoEllipsizeMode

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	arg1 = (C.PangoEllipsizeMode)(ellipsize)

	C.gtk_tool_item_group_set_ellipsize(arg0, arg1)
}

// SetHeaderRelief: set the button relief of the group header. See
// gtk_button_set_relief() for details.
func (g toolItemGroup) SetHeaderRelief(g ToolItemGroup, style ReliefStyle) {
	var arg0 *C.GtkToolItemGroup
	var arg1 C.GtkReliefStyle

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	arg1 = (C.GtkReliefStyle)(style)

	C.gtk_tool_item_group_set_header_relief(arg0, arg1)
}

// SetItemPosition sets the position of @item in the list of children of
// @group.
func (g toolItemGroup) SetItemPosition(g ToolItemGroup, item ToolItem, position int) {
	var arg0 *C.GtkToolItemGroup
	var arg1 *C.GtkToolItem
	var arg2 C.gint

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	arg1 = (*C.GtkToolItem)(unsafe.Pointer(item.Native()))
	arg2 = C.gint(position)

	C.gtk_tool_item_group_set_item_position(arg0, arg1, arg2)
}

// SetLabel sets the label of the tool item group. The label is displayed in
// the header of the group.
func (g toolItemGroup) SetLabel(g ToolItemGroup, label string) {
	var arg0 *C.GtkToolItemGroup
	var arg1 *C.gchar

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_tool_item_group_set_label(arg0, arg1)
}

// SetLabelWidget sets the label of the tool item group. The label widget is
// displayed in the header of the group, in place of the usual label.
func (g toolItemGroup) SetLabelWidget(g ToolItemGroup, labelWidget Widget) {
	var arg0 *C.GtkToolItemGroup
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkToolItemGroup)(unsafe.Pointer(g.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(labelWidget.Native()))

	C.gtk_tool_item_group_set_label_widget(arg0, arg1)
}