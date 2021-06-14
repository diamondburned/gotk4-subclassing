// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tree_view_column_sizing_get_type()), F: marshalTreeViewColumnSizing},
		{T: externglib.Type(C.gtk_tree_view_column_get_type()), F: marshalTreeViewColumn},
	})
}

// TreeViewColumnSizing: the sizing method the column uses to determine its
// width. Please note that @GTK_TREE_VIEW_COLUMN_AUTOSIZE are inefficient for
// large views, and can make columns appear choppy.
type TreeViewColumnSizing int

const (
	// TreeViewColumnSizingGrowOnly columns only get bigger in reaction to
	// changes in the model
	TreeViewColumnSizingGrowOnly TreeViewColumnSizing = 0
	// TreeViewColumnSizingAutosize columns resize to be the optimal size
	// everytime the model changes.
	TreeViewColumnSizingAutosize TreeViewColumnSizing = 1
	// TreeViewColumnSizingFixed columns are a fixed numbers of pixels wide.
	TreeViewColumnSizingFixed TreeViewColumnSizing = 2
)

func marshalTreeViewColumnSizing(p uintptr) (interface{}, error) {
	return TreeViewColumnSizing(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// TreeViewColumn: the GtkTreeViewColumn object represents a visible column in a
// TreeView widget. It allows to set properties of the column header, and
// functions as a holding pen for the cell renderers which determine how the
// data in the column is displayed.
//
// Please refer to the [tree widget conceptual overview][TreeWidget] for an
// overview of all the objects and data types related to the tree widget and how
// they work together.
type TreeViewColumn interface {
	gextras.Objector
	Buildable
	CellLayout

	// AddAttribute adds an attribute mapping to the list in @tree_column. The
	// @column is the column of the model to get a value from, and the
	// @attribute is the parameter on @cell_renderer to be set from the value.
	// So for example if column 2 of the model contains strings, you could have
	// the “text” attribute of a CellRendererText get its values from column 2.
	AddAttribute(cellRenderer CellRenderer, attribute string, column int)
	// CellGetPosition obtains the horizontal position and size of a cell in a
	// column. If the cell is not found in the column, @start_pos and @width are
	// not changed and false is returned.
	CellGetPosition(cellRenderer CellRenderer) (xOffset int, width int, ok bool)
	// CellGetSize obtains the width and height needed to render the column.
	// This is used primarily by the TreeView.
	CellGetSize(cellArea *gdk.Rectangle) (xOffset int, yOffset int, width int, height int)
	// CellIsVisible returns true if any of the cells packed into the
	// @tree_column are visible. For this to be meaningful, you must first
	// initialize the cells with gtk_tree_view_column_cell_set_cell_data()
	CellIsVisible() bool
	// CellSetCellData sets the cell renderer based on the @tree_model and
	// @iter. That is, for every attribute mapping in @tree_column, it will get
	// a value from the set column on the @iter, and use that value to set the
	// attribute on the cell renderer. This is used primarily by the TreeView.
	CellSetCellData(treeModel TreeModel, iter *TreeIter, isExpander bool, isExpanded bool)
	// Clear unsets all the mappings on all renderers on the @tree_column.
	Clear()
	// ClearAttributes clears all existing attributes previously set with
	// gtk_tree_view_column_set_attributes().
	ClearAttributes(cellRenderer CellRenderer)
	// Clicked emits the “clicked” signal on the column. This function will only
	// work if @tree_column is clickable.
	Clicked()
	// FocusCell sets the current keyboard focus to be at @cell, if the column
	// contains 2 or more editable and activatable cells.
	FocusCell(cell CellRenderer)
	// Alignment returns the current x alignment of @tree_column. This value can
	// range between 0.0 and 1.0.
	Alignment() float32
	// Button returns the button used in the treeview column header
	Button() Widget
	// Clickable returns true if the user can click on the header for the
	// column.
	Clickable() bool
	// Expand returns true if the column expands to fill available space.
	Expand() bool
	// FixedWidth gets the fixed width of the column. This may not be the actual
	// displayed width of the column; for that, use
	// gtk_tree_view_column_get_width().
	FixedWidth() int
	// MaxWidth returns the maximum width in pixels of the @tree_column, or -1
	// if no maximum width is set.
	MaxWidth() int
	// MinWidth returns the minimum width in pixels of the @tree_column, or -1
	// if no minimum width is set.
	MinWidth() int
	// Reorderable returns true if the @tree_column can be reordered by the
	// user.
	Reorderable() bool
	// Resizable returns true if the @tree_column can be resized by the end
	// user.
	Resizable() bool
	// Sizing returns the current type of @tree_column.
	Sizing() TreeViewColumnSizing
	// SortColumnID gets the logical @sort_column_id that the model sorts on
	// when this column is selected for sorting. See
	// gtk_tree_view_column_set_sort_column_id().
	SortColumnID() int
	// SortIndicator gets the value set by
	// gtk_tree_view_column_set_sort_indicator().
	SortIndicator() bool
	// SortOrder gets the value set by gtk_tree_view_column_set_sort_order().
	SortOrder() SortType
	// Spacing returns the spacing of @tree_column.
	Spacing() int
	// Title returns the title of the widget.
	Title() string
	// TreeView returns the TreeView wherein @tree_column has been inserted. If
	// @column is currently not inserted in any tree view, nil is returned.
	TreeView() Widget
	// Visible returns true if @tree_column is visible.
	Visible() bool
	// Widget returns the Widget in the button on the column header. If a custom
	// widget has not been set then nil is returned.
	Widget() Widget
	// Width returns the current size of @tree_column in pixels.
	Width() int
	// XOffset returns the current X offset of @tree_column in pixels.
	XOffset() int
	// PackEnd adds the @cell to end of the column. If @expand is false, then
	// the @cell is allocated no more space than it needs. Any unused space is
	// divided evenly between cells for which @expand is true.
	PackEnd(cell CellRenderer, expand bool)
	// PackStart packs the @cell into the beginning of the column. If @expand is
	// false, then the @cell is allocated no more space than it needs. Any
	// unused space is divided evenly between cells for which @expand is true.
	PackStart(cell CellRenderer, expand bool)
	// QueueResize flags the column, and the cell renderers added to this
	// column, to have their sizes renegotiated.
	QueueResize()
	// SetAlignment sets the alignment of the title or custom widget inside the
	// column header. The alignment determines its location inside the button --
	// 0.0 for left, 0.5 for center, 1.0 for right.
	SetAlignment(xalign float32)
	// SetClickable sets the header to be active if @clickable is true. When the
	// header is active, then it can take keyboard focus, and can be clicked.
	SetClickable(clickable bool)
	// SetExpand sets the column to take available extra space. This space is
	// shared equally amongst all columns that have the expand set to true. If
	// no column has this option set, then the last column gets all extra space.
	// By default, every column is created with this false.
	//
	// Along with “fixed-width”, the “expand” property changes when the column
	// is resized by the user.
	SetExpand(expand bool)
	// SetFixedWidth: if @fixed_width is not -1, sets the fixed width of
	// @tree_column; otherwise unsets it. The effective value of @fixed_width is
	// clamped between the minimum and maximum width of the column; however, the
	// value stored in the “fixed-width” property is not clamped. If the column
	// sizing is K_TREE_VIEW_COLUMN_GROW_ONLY or K_TREE_VIEW_COLUMN_AUTOSIZE,
	// setting a fixed width overrides the automatically calculated width. Note
	// that @fixed_width is only a hint to GTK+; the width actually allocated to
	// the column may be greater or less than requested.
	//
	// Along with “expand”, the “fixed-width” property changes when the column
	// is resized by the user.
	SetFixedWidth(fixedWidth int)
	// SetMaxWidth sets the maximum width of the @tree_column. If @max_width is
	// -1, then the maximum width is unset. Note, the column can actually be
	// wider than max width if it’s the last column in a view. In this case, the
	// column expands to fill any extra space.
	SetMaxWidth(maxWidth int)
	// SetMinWidth sets the minimum width of the @tree_column. If @min_width is
	// -1, then the minimum width is unset.
	SetMinWidth(minWidth int)
	// SetReorderable: if @reorderable is true, then the column can be reordered
	// by the end user dragging the header.
	SetReorderable(reorderable bool)
	// SetResizable: if @resizable is true, then the user can explicitly resize
	// the column by grabbing the outer edge of the column button. If resizable
	// is true and sizing mode of the column is K_TREE_VIEW_COLUMN_AUTOSIZE,
	// then the sizing mode is changed to K_TREE_VIEW_COLUMN_GROW_ONLY.
	SetResizable(resizable bool)
	// SetSizing sets the growth behavior of @tree_column to @type.
	SetSizing(typ TreeViewColumnSizing)
	// SetSortColumnID sets the logical @sort_column_id that this column sorts
	// on when this column is selected for sorting. Doing so makes the column
	// header clickable.
	SetSortColumnID(sortColumnId int)
	// SetSortIndicator: call this function with a @setting of true to display
	// an arrow in the header button indicating the column is sorted. Call
	// gtk_tree_view_column_set_sort_order() to change the direction of the
	// arrow.
	SetSortIndicator(setting bool)
	// SetSortOrder changes the appearance of the sort indicator.
	//
	// This does not actually sort the model. Use
	// gtk_tree_view_column_set_sort_column_id() if you want automatic sorting
	// support. This function is primarily for custom sorting behavior, and
	// should be used in conjunction with gtk_tree_sortable_set_sort_column_id()
	// to do that. For custom models, the mechanism will vary.
	//
	// The sort indicator changes direction to indicate normal sort or reverse
	// sort. Note that you must have the sort indicator enabled to see anything
	// when calling this function; see
	// gtk_tree_view_column_set_sort_indicator().
	SetSortOrder(order SortType)
	// SetSpacing sets the spacing field of @tree_column, which is the number of
	// pixels to place between cell renderers packed into it.
	SetSpacing(spacing int)
	// SetTitle sets the title of the @tree_column. If a custom widget has been
	// set, then this value is ignored.
	SetTitle(title string)
	// SetVisible sets the visibility of @tree_column.
	SetVisible(visible bool)
	// SetWidget sets the widget in the header to be @widget. If widget is nil,
	// then the header button is set with a Label set to the title of
	// @tree_column.
	SetWidget(widget Widget)
}

// treeViewColumn implements the TreeViewColumn class.
type treeViewColumn struct {
	gextras.Objector
	Buildable
	CellLayout
}

var _ TreeViewColumn = (*treeViewColumn)(nil)

// WrapTreeViewColumn wraps a GObject to the right type. It is
// primarily used internally.
func WrapTreeViewColumn(obj *externglib.Object) TreeViewColumn {
	return treeViewColumn{
		Objector:   obj,
		Buildable:  WrapBuildable(obj),
		CellLayout: WrapCellLayout(obj),
	}
}

func marshalTreeViewColumn(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTreeViewColumn(obj), nil
}

// NewTreeViewColumn constructs a class TreeViewColumn.
func NewTreeViewColumn() TreeViewColumn {
	var _cret C.GtkTreeViewColumn // in

	_cret = C.gtk_tree_view_column_new()

	var _treeViewColumn TreeViewColumn // out

	_treeViewColumn = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(TreeViewColumn)

	return _treeViewColumn
}

// NewTreeViewColumnWithArea constructs a class TreeViewColumn.
func NewTreeViewColumnWithArea(area CellArea) TreeViewColumn {
	var _arg1 *C.GtkCellArea // out

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer(area.Native()))

	var _cret C.GtkTreeViewColumn // in

	_cret = C.gtk_tree_view_column_new_with_area(_arg1)

	var _treeViewColumn TreeViewColumn // out

	_treeViewColumn = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(TreeViewColumn)

	return _treeViewColumn
}

// AddAttribute adds an attribute mapping to the list in @tree_column. The
// @column is the column of the model to get a value from, and the
// @attribute is the parameter on @cell_renderer to be set from the value.
// So for example if column 2 of the model contains strings, you could have
// the “text” attribute of a CellRendererText get its values from column 2.
func (t treeViewColumn) AddAttribute(cellRenderer CellRenderer, attribute string, column int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkCellRenderer   // out
	var _arg2 *C.gchar             // out
	var _arg3 C.gint               // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cellRenderer.Native()))
	_arg2 = (*C.gchar)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.gint(column)

	C.gtk_tree_view_column_add_attribute(_arg0, _arg1, _arg2, _arg3)
}

// CellGetPosition obtains the horizontal position and size of a cell in a
// column. If the cell is not found in the column, @start_pos and @width are
// not changed and false is returned.
func (t treeViewColumn) CellGetPosition(cellRenderer CellRenderer) (xOffset int, width int, ok bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkCellRenderer   // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cellRenderer.Native()))

	var _arg2 C.gint     // in
	var _arg3 C.gint     // in
	var _cret C.gboolean // in

	_cret = C.gtk_tree_view_column_cell_get_position(_arg0, _arg1, &_arg2, &_arg3)

	var _xOffset int // out
	var _width int   // out
	var _ok bool     // out

	_xOffset = (int)(_arg2)
	_width = (int)(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _xOffset, _width, _ok
}

// CellGetSize obtains the width and height needed to render the column.
// This is used primarily by the TreeView.
func (t treeViewColumn) CellGetSize(cellArea *gdk.Rectangle) (xOffset int, yOffset int, width int, height int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GdkRectangle      // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GdkRectangle)(unsafe.Pointer(cellArea.Native()))

	var _arg2 C.gint // in
	var _arg3 C.gint // in
	var _arg4 C.gint // in
	var _arg5 C.gint // in

	C.gtk_tree_view_column_cell_get_size(_arg0, _arg1, &_arg2, &_arg3, &_arg4, &_arg5)

	var _xOffset int // out
	var _yOffset int // out
	var _width int   // out
	var _height int  // out

	_xOffset = (int)(_arg2)
	_yOffset = (int)(_arg3)
	_width = (int)(_arg4)
	_height = (int)(_arg5)

	return _xOffset, _yOffset, _width, _height
}

// CellIsVisible returns true if any of the cells packed into the
// @tree_column are visible. For this to be meaningful, you must first
// initialize the cells with gtk_tree_view_column_cell_set_cell_data()
func (t treeViewColumn) CellIsVisible() bool {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_tree_view_column_cell_is_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CellSetCellData sets the cell renderer based on the @tree_model and
// @iter. That is, for every attribute mapping in @tree_column, it will get
// a value from the set column on the @iter, and use that value to set the
// attribute on the cell renderer. This is used primarily by the TreeView.
func (t treeViewColumn) CellSetCellData(treeModel TreeModel, iter *TreeIter, isExpander bool, isExpanded bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkTreeModel      // out
	var _arg2 *C.GtkTreeIter       // out
	var _arg3 C.gboolean           // out
	var _arg4 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(treeModel.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))
	if isExpander {
		_arg3 = C.TRUE
	}
	if isExpanded {
		_arg4 = C.TRUE
	}

	C.gtk_tree_view_column_cell_set_cell_data(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// Clear unsets all the mappings on all renderers on the @tree_column.
func (t treeViewColumn) Clear() {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	C.gtk_tree_view_column_clear(_arg0)
}

// ClearAttributes clears all existing attributes previously set with
// gtk_tree_view_column_set_attributes().
func (t treeViewColumn) ClearAttributes(cellRenderer CellRenderer) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkCellRenderer   // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cellRenderer.Native()))

	C.gtk_tree_view_column_clear_attributes(_arg0, _arg1)
}

// Clicked emits the “clicked” signal on the column. This function will only
// work if @tree_column is clickable.
func (t treeViewColumn) Clicked() {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	C.gtk_tree_view_column_clicked(_arg0)
}

// FocusCell sets the current keyboard focus to be at @cell, if the column
// contains 2 or more editable and activatable cells.
func (t treeViewColumn) FocusCell(cell CellRenderer) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkCellRenderer   // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))

	C.gtk_tree_view_column_focus_cell(_arg0, _arg1)
}

// Alignment returns the current x alignment of @tree_column. This value can
// range between 0.0 and 1.0.
func (t treeViewColumn) Alignment() float32 {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret C.gfloat // in

	_cret = C.gtk_tree_view_column_get_alignment(_arg0)

	var _gfloat float32 // out

	_gfloat = (float32)(_cret)

	return _gfloat
}

// Button returns the button used in the treeview column header
func (t treeViewColumn) Button() Widget {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret *C.GtkWidget // in

	_cret = C.gtk_tree_view_column_get_button(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// Clickable returns true if the user can click on the header for the
// column.
func (t treeViewColumn) Clickable() bool {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_tree_view_column_get_clickable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Expand returns true if the column expands to fill available space.
func (t treeViewColumn) Expand() bool {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_tree_view_column_get_expand(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// FixedWidth gets the fixed width of the column. This may not be the actual
// displayed width of the column; for that, use
// gtk_tree_view_column_get_width().
func (t treeViewColumn) FixedWidth() int {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret C.gint // in

	_cret = C.gtk_tree_view_column_get_fixed_width(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// MaxWidth returns the maximum width in pixels of the @tree_column, or -1
// if no maximum width is set.
func (t treeViewColumn) MaxWidth() int {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret C.gint // in

	_cret = C.gtk_tree_view_column_get_max_width(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// MinWidth returns the minimum width in pixels of the @tree_column, or -1
// if no minimum width is set.
func (t treeViewColumn) MinWidth() int {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret C.gint // in

	_cret = C.gtk_tree_view_column_get_min_width(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Reorderable returns true if the @tree_column can be reordered by the
// user.
func (t treeViewColumn) Reorderable() bool {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_tree_view_column_get_reorderable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Resizable returns true if the @tree_column can be resized by the end
// user.
func (t treeViewColumn) Resizable() bool {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_tree_view_column_get_resizable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Sizing returns the current type of @tree_column.
func (t treeViewColumn) Sizing() TreeViewColumnSizing {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret C.GtkTreeViewColumnSizing // in

	_cret = C.gtk_tree_view_column_get_sizing(_arg0)

	var _treeViewColumnSizing TreeViewColumnSizing // out

	_treeViewColumnSizing = TreeViewColumnSizing(_cret)

	return _treeViewColumnSizing
}

// SortColumnID gets the logical @sort_column_id that the model sorts on
// when this column is selected for sorting. See
// gtk_tree_view_column_set_sort_column_id().
func (t treeViewColumn) SortColumnID() int {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret C.gint // in

	_cret = C.gtk_tree_view_column_get_sort_column_id(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// SortIndicator gets the value set by
// gtk_tree_view_column_set_sort_indicator().
func (t treeViewColumn) SortIndicator() bool {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_tree_view_column_get_sort_indicator(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SortOrder gets the value set by gtk_tree_view_column_set_sort_order().
func (t treeViewColumn) SortOrder() SortType {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret C.GtkSortType // in

	_cret = C.gtk_tree_view_column_get_sort_order(_arg0)

	var _sortType SortType // out

	_sortType = SortType(_cret)

	return _sortType
}

// Spacing returns the spacing of @tree_column.
func (t treeViewColumn) Spacing() int {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret C.gint // in

	_cret = C.gtk_tree_view_column_get_spacing(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Title returns the title of the widget.
func (t treeViewColumn) Title() string {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret *C.gchar // in

	_cret = C.gtk_tree_view_column_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// TreeView returns the TreeView wherein @tree_column has been inserted. If
// @column is currently not inserted in any tree view, nil is returned.
func (t treeViewColumn) TreeView() Widget {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret *C.GtkWidget // in

	_cret = C.gtk_tree_view_column_get_tree_view(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// Visible returns true if @tree_column is visible.
func (t treeViewColumn) Visible() bool {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_tree_view_column_get_visible(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Widget returns the Widget in the button on the column header. If a custom
// widget has not been set then nil is returned.
func (t treeViewColumn) Widget() Widget {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret *C.GtkWidget // in

	_cret = C.gtk_tree_view_column_get_widget(_arg0)

	var _widget Widget // out

	_widget = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Widget)

	return _widget
}

// Width returns the current size of @tree_column in pixels.
func (t treeViewColumn) Width() int {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret C.gint // in

	_cret = C.gtk_tree_view_column_get_width(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// XOffset returns the current X offset of @tree_column in pixels.
func (t treeViewColumn) XOffset() int {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	var _cret C.gint // in

	_cret = C.gtk_tree_view_column_get_x_offset(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// PackEnd adds the @cell to end of the column. If @expand is false, then
// the @cell is allocated no more space than it needs. Any unused space is
// divided evenly between cells for which @expand is true.
func (t treeViewColumn) PackEnd(cell CellRenderer, expand bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkCellRenderer   // out
	var _arg2 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if expand {
		_arg2 = C.TRUE
	}

	C.gtk_tree_view_column_pack_end(_arg0, _arg1, _arg2)
}

// PackStart packs the @cell into the beginning of the column. If @expand is
// false, then the @cell is allocated no more space than it needs. Any
// unused space is divided evenly between cells for which @expand is true.
func (t treeViewColumn) PackStart(cell CellRenderer, expand bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkCellRenderer   // out
	var _arg2 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(cell.Native()))
	if expand {
		_arg2 = C.TRUE
	}

	C.gtk_tree_view_column_pack_start(_arg0, _arg1, _arg2)
}

// QueueResize flags the column, and the cell renderers added to this
// column, to have their sizes renegotiated.
func (t treeViewColumn) QueueResize() {
	var _arg0 *C.GtkTreeViewColumn // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))

	C.gtk_tree_view_column_queue_resize(_arg0)
}

// SetAlignment sets the alignment of the title or custom widget inside the
// column header. The alignment determines its location inside the button --
// 0.0 for left, 0.5 for center, 1.0 for right.
func (t treeViewColumn) SetAlignment(xalign float32) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gfloat             // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = C.gfloat(xalign)

	C.gtk_tree_view_column_set_alignment(_arg0, _arg1)
}

// SetClickable sets the header to be active if @clickable is true. When the
// header is active, then it can take keyboard focus, and can be clicked.
func (t treeViewColumn) SetClickable(clickable bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	if clickable {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_column_set_clickable(_arg0, _arg1)
}

// SetExpand sets the column to take available extra space. This space is
// shared equally amongst all columns that have the expand set to true. If
// no column has this option set, then the last column gets all extra space.
// By default, every column is created with this false.
//
// Along with “fixed-width”, the “expand” property changes when the column
// is resized by the user.
func (t treeViewColumn) SetExpand(expand bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	if expand {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_column_set_expand(_arg0, _arg1)
}

// SetFixedWidth: if @fixed_width is not -1, sets the fixed width of
// @tree_column; otherwise unsets it. The effective value of @fixed_width is
// clamped between the minimum and maximum width of the column; however, the
// value stored in the “fixed-width” property is not clamped. If the column
// sizing is K_TREE_VIEW_COLUMN_GROW_ONLY or K_TREE_VIEW_COLUMN_AUTOSIZE,
// setting a fixed width overrides the automatically calculated width. Note
// that @fixed_width is only a hint to GTK+; the width actually allocated to
// the column may be greater or less than requested.
//
// Along with “expand”, the “fixed-width” property changes when the column
// is resized by the user.
func (t treeViewColumn) SetFixedWidth(fixedWidth int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = C.gint(fixedWidth)

	C.gtk_tree_view_column_set_fixed_width(_arg0, _arg1)
}

// SetMaxWidth sets the maximum width of the @tree_column. If @max_width is
// -1, then the maximum width is unset. Note, the column can actually be
// wider than max width if it’s the last column in a view. In this case, the
// column expands to fill any extra space.
func (t treeViewColumn) SetMaxWidth(maxWidth int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = C.gint(maxWidth)

	C.gtk_tree_view_column_set_max_width(_arg0, _arg1)
}

// SetMinWidth sets the minimum width of the @tree_column. If @min_width is
// -1, then the minimum width is unset.
func (t treeViewColumn) SetMinWidth(minWidth int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = C.gint(minWidth)

	C.gtk_tree_view_column_set_min_width(_arg0, _arg1)
}

// SetReorderable: if @reorderable is true, then the column can be reordered
// by the end user dragging the header.
func (t treeViewColumn) SetReorderable(reorderable bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	if reorderable {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_column_set_reorderable(_arg0, _arg1)
}

// SetResizable: if @resizable is true, then the user can explicitly resize
// the column by grabbing the outer edge of the column button. If resizable
// is true and sizing mode of the column is K_TREE_VIEW_COLUMN_AUTOSIZE,
// then the sizing mode is changed to K_TREE_VIEW_COLUMN_GROW_ONLY.
func (t treeViewColumn) SetResizable(resizable bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	if resizable {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_column_set_resizable(_arg0, _arg1)
}

// SetSizing sets the growth behavior of @tree_column to @type.
func (t treeViewColumn) SetSizing(typ TreeViewColumnSizing) {
	var _arg0 *C.GtkTreeViewColumn      // out
	var _arg1 C.GtkTreeViewColumnSizing // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (C.GtkTreeViewColumnSizing)(typ)

	C.gtk_tree_view_column_set_sizing(_arg0, _arg1)
}

// SetSortColumnID sets the logical @sort_column_id that this column sorts
// on when this column is selected for sorting. Doing so makes the column
// header clickable.
func (t treeViewColumn) SetSortColumnID(sortColumnId int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = C.gint(sortColumnId)

	C.gtk_tree_view_column_set_sort_column_id(_arg0, _arg1)
}

// SetSortIndicator: call this function with a @setting of true to display
// an arrow in the header button indicating the column is sorted. Call
// gtk_tree_view_column_set_sort_order() to change the direction of the
// arrow.
func (t treeViewColumn) SetSortIndicator(setting bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_column_set_sort_indicator(_arg0, _arg1)
}

// SetSortOrder changes the appearance of the sort indicator.
//
// This does not actually sort the model. Use
// gtk_tree_view_column_set_sort_column_id() if you want automatic sorting
// support. This function is primarily for custom sorting behavior, and
// should be used in conjunction with gtk_tree_sortable_set_sort_column_id()
// to do that. For custom models, the mechanism will vary.
//
// The sort indicator changes direction to indicate normal sort or reverse
// sort. Note that you must have the sort indicator enabled to see anything
// when calling this function; see
// gtk_tree_view_column_set_sort_indicator().
func (t treeViewColumn) SetSortOrder(order SortType) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.GtkSortType        // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (C.GtkSortType)(order)

	C.gtk_tree_view_column_set_sort_order(_arg0, _arg1)
}

// SetSpacing sets the spacing field of @tree_column, which is the number of
// pixels to place between cell renderers packed into it.
func (t treeViewColumn) SetSpacing(spacing int) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gint               // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = C.gint(spacing)

	C.gtk_tree_view_column_set_spacing(_arg0, _arg1)
}

// SetTitle sets the title of the @tree_column. If a custom widget has been
// set, then this value is ignored.
func (t treeViewColumn) SetTitle(title string) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.gchar             // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_tree_view_column_set_title(_arg0, _arg1)
}

// SetVisible sets the visibility of @tree_column.
func (t treeViewColumn) SetVisible(visible bool) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 C.gboolean           // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_tree_view_column_set_visible(_arg0, _arg1)
}

// SetWidget sets the widget in the header to be @widget. If widget is nil,
// then the header button is set with a Label set to the title of
// @tree_column.
func (t treeViewColumn) SetWidget(widget Widget) {
	var _arg0 *C.GtkTreeViewColumn // out
	var _arg1 *C.GtkWidget         // out

	_arg0 = (*C.GtkTreeViewColumn)(unsafe.Pointer(t.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_tree_view_column_set_widget(_arg0, _arg1)
}
