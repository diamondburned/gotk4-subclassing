// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_column_view_get_type()), F: marshalColumnView},
	})
}

// ColumnView: gtkColumnView is a widget to present a view into a large dynamic
// list of items using multiple columns with headers.
//
// GtkColumnView uses the factories of its columns to generate a cell widget for
// each column, for each visible item and displays them together as the row for
// this item. The ColumnView:show-row-separators and
// ColumnView:show-column-separators properties offer a simple way to display
// separators between the rows or columns.
//
// GtkColumnView allows the user to select items according to the selection
// characteristics of the model. For models that allow multiple selected items,
// it is possible to turn on *rubberband selection*, using
// ColumnView:enable-rubberband.
//
// The column view supports sorting that can be customized by the user by
// clicking on column headers. To set this up, the Sorter returned by
// gtk_column_view_get_sorter() must be attached to a sort model for the data
// that the view is showing, and the columns must have sorters attached to them
// by calling gtk_column_view_column_set_sorter(). The initial sort order can be
// set with gtk_column_view_sort_by_column().
//
// The column view also supports interactive resizing and reordering of columns,
// via Drag-and-Drop of the column headers. This can be enabled or disabled with
// the ColumnView:reorderable and ColumnViewColumn:resizable properties.
//
// To learn more about the list widget framework, see the overview (Widget).
//
// CSS nodes
//
//    columnview[.column-separators][.rich-list][.navigation-sidebar][.data-table]
//    ├── header
//    │   ├── <column header>
//    ┊   ┊
//    │   ╰── <column header>
//    │
//    ├── listview
//    │
//    ┊
//    ╰── [rubberband]
//
//
// GtkColumnView uses a single CSS node named columnview. It may carry the
// .column-separators style class, when ColumnView:show-column-separators
// property is set. Header widets appear below a node with name header. The rows
// are contained in a GtkListView widget, so there is a listview node with the
// same structure as for a standalone GtkListView widget. If
// ColumnView:show-row-separators is set, it will be passed on to the list view,
// causing its CSS node to carry the .separators style class. For rubberband
// selection, a node with name rubberband is used.
//
// The main columnview node may also carry style classes to select the style of
// list presentation (ListContainers.html#list-styles): .rich-list,
// .navigation-sidebar or .data-table.
//
//
// Accessibility
//
// GtkColumnView uses the K_ACCESSIBLE_ROLE_TREE_GRID role, header title widgets
// are using the K_ACCESSIBLE_ROLE_COLUMN_HEADER role. The row widgets are using
// the K_ACCESSIBLE_ROLE_ROW role, and individual cells are using the
// K_ACCESSIBLE_ROLE_GRID_CELL role
type ColumnView interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Scrollable

	// AppendColumn appends the @column to the end of the columns in @self.
	AppendColumn(s ColumnView, column ColumnViewColumn)
	// Columns gets the list of columns in this column view. This list is
	// constant over the lifetime of @self and can be used to monitor changes to
	// the columns of @self by connecting to the Model:items-changed signal.
	Columns(s ColumnView)
	// EnableRubberband returns whether rows can be selected by dragging with
	// the mouse.
	EnableRubberband(s ColumnView) bool
	// Model gets the model that's currently used to read the items displayed.
	Model(s ColumnView)
	// Reorderable returns whether columns are reorderable.
	Reorderable(s ColumnView) bool
	// ShowColumnSeparators returns whether the list should show separators
	// between columns.
	ShowColumnSeparators(s ColumnView) bool
	// ShowRowSeparators returns whether the list should show separators between
	// rows.
	ShowRowSeparators(s ColumnView) bool
	// SingleClickActivate returns whether rows will be activated on single
	// click and selected on hover.
	SingleClickActivate(s ColumnView) bool
	// Sorter returns a special sorter that reflects the users sorting choices
	// in the column view.
	//
	// To allow users to customizable sorting by clicking on column headers,
	// this sorter needs to be set on the sort model underneath the model that
	// is displayed by the view.
	//
	// See gtk_column_view_column_set_sorter() for setting up per-column
	// sorting.
	//
	// Here is an example:
	//
	//    gtk_column_view_column_set_sorter (column, sorter);
	//    gtk_column_view_append_column (view, column);
	//    sorter = g_object_ref (gtk_column_view_get_sorter (view)));
	//    model = gtk_sort_list_model_new (store, sorter);
	//    selection = gtk_no_selection_new (model);
	//    gtk_column_view_set_model (view, selection);
	Sorter(s ColumnView)
	// InsertColumn inserts a column at the given position in the columns of
	// @self.
	//
	// If @column is already a column of @self, it will be repositioned.
	InsertColumn(s ColumnView, position uint, column ColumnViewColumn)
	// RemoveColumn removes the @column from the list of columns of @self.
	RemoveColumn(s ColumnView, column ColumnViewColumn)
	// SetEnableRubberband sets whether selections can be changed by dragging
	// with the mouse.
	SetEnableRubberband(s ColumnView, enableRubberband bool)
	// SetModel sets the SelectionModel to use.
	SetModel(s ColumnView, model SelectionModel)
	// SetReorderable sets whether columns should be reorderable by dragging.
	SetReorderable(s ColumnView, reorderable bool)
	// SetShowColumnSeparators sets whether the list should show separators
	// between columns.
	SetShowColumnSeparators(s ColumnView, showColumnSeparators bool)
	// SetShowRowSeparators sets whether the list should show separators between
	// rows.
	SetShowRowSeparators(s ColumnView, showRowSeparators bool)
	// SetSingleClickActivate sets whether rows should be activated on single
	// click and selected on hover.
	SetSingleClickActivate(s ColumnView, singleClickActivate bool)
	// SortByColumn sets the sorting of the view.
	//
	// This function should be used to set up the initial sorting. At runtime,
	// users can change the sorting of a column view by clicking on the list
	// headers.
	//
	// This call only has an effect if the sorter returned by
	// gtk_column_view_get_sorter() is set on a sort model, and
	// gtk_column_view_column_set_sorter() has been called on @column to
	// associate a sorter with the column.
	//
	// If @column is nil, the view will be unsorted.
	SortByColumn(s ColumnView, column ColumnViewColumn, direction SortType)
}

// columnView implements the ColumnView interface.
type columnView struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
	Scrollable
}

var _ ColumnView = (*columnView)(nil)

// WrapColumnView wraps a GObject to the right type. It is
// primarily used internally.
func WrapColumnView(obj *externglib.Object) ColumnView {
	return ColumnView{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		Scrollable:       WrapScrollable(obj),
	}
}

func marshalColumnView(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapColumnView(obj), nil
}

// NewColumnView constructs a class ColumnView.
func NewColumnView(model SelectionModel) {
	var arg1 *C.GtkSelectionModel

	arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))

	C.gtk_column_view_new(arg1)
}

// AppendColumn appends the @column to the end of the columns in @self.
func (s columnView) AppendColumn(s ColumnView, column ColumnViewColumn) {
	var arg0 *C.GtkColumnView
	var arg1 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkColumnViewColumn)(unsafe.Pointer(column.Native()))

	C.gtk_column_view_append_column(arg0, arg1)
}

// Columns gets the list of columns in this column view. This list is
// constant over the lifetime of @self and can be used to monitor changes to
// the columns of @self by connecting to the Model:items-changed signal.
func (s columnView) Columns(s ColumnView) {
	var arg0 *C.GtkColumnView

	arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))

	C.gtk_column_view_get_columns(arg0)
}

// EnableRubberband returns whether rows can be selected by dragging with
// the mouse.
func (s columnView) EnableRubberband(s ColumnView) bool {
	var arg0 *C.GtkColumnView

	arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_column_view_get_enable_rubberband(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Model gets the model that's currently used to read the items displayed.
func (s columnView) Model(s ColumnView) {
	var arg0 *C.GtkColumnView

	arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))

	C.gtk_column_view_get_model(arg0)
}

// Reorderable returns whether columns are reorderable.
func (s columnView) Reorderable(s ColumnView) bool {
	var arg0 *C.GtkColumnView

	arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_column_view_get_reorderable(arg0)

	if cret {
		ok = true
	}

	return ok
}

// ShowColumnSeparators returns whether the list should show separators
// between columns.
func (s columnView) ShowColumnSeparators(s ColumnView) bool {
	var arg0 *C.GtkColumnView

	arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_column_view_get_show_column_separators(arg0)

	if cret {
		ok = true
	}

	return ok
}

// ShowRowSeparators returns whether the list should show separators between
// rows.
func (s columnView) ShowRowSeparators(s ColumnView) bool {
	var arg0 *C.GtkColumnView

	arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_column_view_get_show_row_separators(arg0)

	if cret {
		ok = true
	}

	return ok
}

// SingleClickActivate returns whether rows will be activated on single
// click and selected on hover.
func (s columnView) SingleClickActivate(s ColumnView) bool {
	var arg0 *C.GtkColumnView

	arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_column_view_get_single_click_activate(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Sorter returns a special sorter that reflects the users sorting choices
// in the column view.
//
// To allow users to customizable sorting by clicking on column headers,
// this sorter needs to be set on the sort model underneath the model that
// is displayed by the view.
//
// See gtk_column_view_column_set_sorter() for setting up per-column
// sorting.
//
// Here is an example:
//
//    gtk_column_view_column_set_sorter (column, sorter);
//    gtk_column_view_append_column (view, column);
//    sorter = g_object_ref (gtk_column_view_get_sorter (view)));
//    model = gtk_sort_list_model_new (store, sorter);
//    selection = gtk_no_selection_new (model);
//    gtk_column_view_set_model (view, selection);
func (s columnView) Sorter(s ColumnView) {
	var arg0 *C.GtkColumnView

	arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))

	C.gtk_column_view_get_sorter(arg0)
}

// InsertColumn inserts a column at the given position in the columns of
// @self.
//
// If @column is already a column of @self, it will be repositioned.
func (s columnView) InsertColumn(s ColumnView, position uint, column ColumnViewColumn) {
	var arg0 *C.GtkColumnView
	var arg1 C.guint
	var arg2 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	arg1 = C.guint(position)
	arg2 = (*C.GtkColumnViewColumn)(unsafe.Pointer(column.Native()))

	C.gtk_column_view_insert_column(arg0, arg1, arg2)
}

// RemoveColumn removes the @column from the list of columns of @self.
func (s columnView) RemoveColumn(s ColumnView, column ColumnViewColumn) {
	var arg0 *C.GtkColumnView
	var arg1 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkColumnViewColumn)(unsafe.Pointer(column.Native()))

	C.gtk_column_view_remove_column(arg0, arg1)
}

// SetEnableRubberband sets whether selections can be changed by dragging
// with the mouse.
func (s columnView) SetEnableRubberband(s ColumnView, enableRubberband bool) {
	var arg0 *C.GtkColumnView
	var arg1 C.gboolean

	arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	if enableRubberband {
		arg1 = C.gboolean(1)
	}

	C.gtk_column_view_set_enable_rubberband(arg0, arg1)
}

// SetModel sets the SelectionModel to use.
func (s columnView) SetModel(s ColumnView, model SelectionModel) {
	var arg0 *C.GtkColumnView
	var arg1 *C.GtkSelectionModel

	arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))

	C.gtk_column_view_set_model(arg0, arg1)
}

// SetReorderable sets whether columns should be reorderable by dragging.
func (s columnView) SetReorderable(s ColumnView, reorderable bool) {
	var arg0 *C.GtkColumnView
	var arg1 C.gboolean

	arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	if reorderable {
		arg1 = C.gboolean(1)
	}

	C.gtk_column_view_set_reorderable(arg0, arg1)
}

// SetShowColumnSeparators sets whether the list should show separators
// between columns.
func (s columnView) SetShowColumnSeparators(s ColumnView, showColumnSeparators bool) {
	var arg0 *C.GtkColumnView
	var arg1 C.gboolean

	arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	if showColumnSeparators {
		arg1 = C.gboolean(1)
	}

	C.gtk_column_view_set_show_column_separators(arg0, arg1)
}

// SetShowRowSeparators sets whether the list should show separators between
// rows.
func (s columnView) SetShowRowSeparators(s ColumnView, showRowSeparators bool) {
	var arg0 *C.GtkColumnView
	var arg1 C.gboolean

	arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	if showRowSeparators {
		arg1 = C.gboolean(1)
	}

	C.gtk_column_view_set_show_row_separators(arg0, arg1)
}

// SetSingleClickActivate sets whether rows should be activated on single
// click and selected on hover.
func (s columnView) SetSingleClickActivate(s ColumnView, singleClickActivate bool) {
	var arg0 *C.GtkColumnView
	var arg1 C.gboolean

	arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	if singleClickActivate {
		arg1 = C.gboolean(1)
	}

	C.gtk_column_view_set_single_click_activate(arg0, arg1)
}

// SortByColumn sets the sorting of the view.
//
// This function should be used to set up the initial sorting. At runtime,
// users can change the sorting of a column view by clicking on the list
// headers.
//
// This call only has an effect if the sorter returned by
// gtk_column_view_get_sorter() is set on a sort model, and
// gtk_column_view_column_set_sorter() has been called on @column to
// associate a sorter with the column.
//
// If @column is nil, the view will be unsorted.
func (s columnView) SortByColumn(s ColumnView, column ColumnViewColumn, direction SortType) {
	var arg0 *C.GtkColumnView
	var arg1 *C.GtkColumnViewColumn
	var arg2 C.GtkSortType

	arg0 = (*C.GtkColumnView)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkColumnViewColumn)(unsafe.Pointer(column.Native()))
	arg2 = (C.GtkSortType)(direction)

	C.gtk_column_view_sort_by_column(arg0, arg1, arg2)
}
