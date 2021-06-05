// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdbool.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_list_box_get_type()), F: marshalListBox},
		{T: externglib.Type(C.gtk_list_box_row_get_type()), F: marshalListBoxRow},
	})
}

// ListBoxCreateWidgetFunc: called for list boxes that are bound to a Model with
// gtk_list_box_bind_model() for each item that gets added to the model.
//
// Versions of GTK+ prior to 3.18 called gtk_widget_show_all() on the rows
// created by the GtkListBoxCreateWidgetFunc, but this forced all widgets inside
// the row to be shown, and is no longer the case. Applications should be
// updated to show the desired row widgets.
type ListBoxCreateWidgetFunc func(item gextras.Objector) Widget

//export gotk4_ListBoxCreateWidgetFunc
func gotk4_ListBoxCreateWidgetFunc(arg0 C.gpointer, arg1 C.gpointer) *C.GtkWidget {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(ListBoxCreateWidgetFunc)
	ret := fn(item, userData)

	cret = (*C.GtkWidget)(unsafe.Pointer(ret.Native()))

	return cret
}

// ListBoxFilterFunc: will be called whenever the row changes or is added and
// lets you control if the row should be visible or not.
type ListBoxFilterFunc func(row ListBoxRow) bool

//export gotk4_ListBoxFilterFunc
func gotk4_ListBoxFilterFunc(arg0 *C.GtkListBoxRow, arg1 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(ListBoxFilterFunc)
	ret := fn(row, userData)

	if ret {
		cret = C.gboolean(1)
	}

	return cret
}

// ListBoxForeachFunc: a function used by gtk_list_box_selected_foreach(). It
// will be called on every selected child of the @box.
type ListBoxForeachFunc func(box ListBox, row ListBoxRow)

//export gotk4_ListBoxForeachFunc
func gotk4_ListBoxForeachFunc(arg0 *C.GtkListBox, arg1 *C.GtkListBoxRow, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(ListBoxForeachFunc)
	fn(box, row, userData)
}

// ListBoxSortFunc: compare two rows to determine which should be first.
type ListBoxSortFunc func(row1 ListBoxRow, row2 ListBoxRow) int

//export gotk4_ListBoxSortFunc
func gotk4_ListBoxSortFunc(arg0 *C.GtkListBoxRow, arg1 *C.GtkListBoxRow, arg2 C.gpointer) C.gint {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(ListBoxSortFunc)
	ret := fn(row1, row2, userData)

	cret = C.gint(ret)

	return cret
}

// ListBoxUpdateHeaderFunc: whenever @row changes or which row is before @row
// changes this is called, which lets you update the header on @row. You may
// remove or set a new one via gtk_list_box_row_set_header() or just change the
// state of the current header widget.
type ListBoxUpdateHeaderFunc func(row ListBoxRow, before ListBoxRow)

//export gotk4_ListBoxUpdateHeaderFunc
func gotk4_ListBoxUpdateHeaderFunc(arg0 *C.GtkListBoxRow, arg1 *C.GtkListBoxRow, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(ListBoxUpdateHeaderFunc)
	fn(row, before, userData)
}

// ListBox: a GtkListBox is a vertical container that contains GtkListBoxRow
// children. These rows can by dynamically sorted and filtered, and headers can
// be added dynamically depending on the row content. It also allows keyboard
// and mouse navigation and selection like a typical list.
//
// Using GtkListBox is often an alternative to TreeView, especially when the
// list contents has a more complicated layout than what is allowed by a
// CellRenderer, or when the contents is interactive (i.e. has a button in it).
//
// Although a ListBox must have only ListBoxRow children you can add any kind of
// widget to it via gtk_container_add(), and a ListBoxRow widget will
// automatically be inserted between the list and the widget.
//
// ListBoxRows can be marked as activatable or selectable. If a row is
// activatable, ListBox::row-activated will be emitted for it when the user
// tries to activate it. If it is selectable, the row will be marked as selected
// when the user tries to select it.
//
// The GtkListBox widget was added in GTK+ 3.10.
//
//
// GtkListBox as GtkBuildable
//
// The GtkListBox implementation of the Buildable interface supports setting a
// child as the placeholder by specifying “placeholder” as the “type” attribute
// of a <child> element. See gtk_list_box_set_placeholder() for info.
//
// CSS nodes
//
//    list
//    ╰── row[.activatable]
//
// GtkListBox uses a single CSS node named list. Each GtkListBoxRow uses a
// single CSS node named row. The row nodes get the .activatable style class
// added when appropriate.
type ListBox interface {
	Container
	Buildable

	// BindModel binds @model to @box.
	//
	// If @box was already bound to a model, that previous binding is destroyed.
	//
	// The contents of @box are cleared and then filled with widgets that
	// represent items from @model. @box is updated whenever @model changes. If
	// @model is nil, @box is left empty.
	//
	// It is undefined to add or remove widgets directly (for example, with
	// gtk_list_box_insert() or gtk_container_add()) while @box is bound to a
	// model.
	//
	// Note that using a model is incompatible with the filtering and sorting
	// functionality in GtkListBox. When using a model, filtering and sorting
	// should be implemented by the model.
	BindModel(model gio.ListModel, createWidgetFunc ListBoxCreateWidgetFunc)
	// DragHighlightRow: this is a helper function for implementing DnD onto a
	// ListBox. The passed in @row will be highlighted via gtk_drag_highlight(),
	// and any previously highlighted row will be unhighlighted.
	//
	// The row will also be unhighlighted when the widget gets a drag leave
	// event.
	DragHighlightRow(row ListBoxRow)
	// DragUnhighlightRow: if a row has previously been highlighted via
	// gtk_list_box_drag_highlight_row() it will have the highlight removed.
	DragUnhighlightRow()
	// ActivateOnSingleClick returns whether rows activate on single clicks.
	ActivateOnSingleClick() bool
	// Adjustment gets the adjustment (if any) that the widget uses to for
	// vertical scrolling.
	Adjustment() Adjustment
	// RowAtIndex gets the n-th child in the list (not counting headers). If
	// @_index is negative or larger than the number of items in the list, nil
	// is returned.
	RowAtIndex(index_ int) ListBoxRow
	// RowAtY gets the row at the @y position.
	RowAtY(y int) ListBoxRow
	// SelectedRow gets the selected row.
	//
	// Note that the box may allow multiple selection, in which case you should
	// use gtk_list_box_selected_foreach() to find all selected rows.
	SelectedRow() ListBoxRow
	// SelectedRows creates a list of all selected children.
	SelectedRows() *glib.List
	// SelectionMode gets the selection mode of the listbox.
	SelectionMode() SelectionMode
	// Insert: insert the @child into the @box at @position. If a sort function
	// is set, the widget will actually be inserted at the calculated position
	// and this function has the same effect of gtk_container_add().
	//
	// If @position is -1, or larger than the total number of items in the @box,
	// then the @child will be appended to the end.
	Insert(child Widget, position int)
	// InvalidateFilter: update the filtering for all rows. Call this when
	// result of the filter function on the @box is changed due to an external
	// factor. For instance, this would be used if the filter function just
	// looked for a specific search string and the entry with the search string
	// has changed.
	InvalidateFilter()
	// InvalidateHeaders: update the separators for all rows. Call this when
	// result of the header function on the @box is changed due to an external
	// factor.
	InvalidateHeaders()
	// InvalidateSort: update the sorting for all rows. Call this when result of
	// the sort function on the @box is changed due to an external factor.
	InvalidateSort()
	// Prepend: prepend a widget to the list. If a sort function is set, the
	// widget will actually be inserted at the calculated position and this
	// function has the same effect of gtk_container_add().
	Prepend(child Widget)
	// SelectAll: select all children of @box, if the selection mode allows it.
	SelectAll()
	// SelectRow: make @row the currently selected row.
	SelectRow(row ListBoxRow)
	// SelectedForeach calls a function for each selected child.
	//
	// Note that the selection cannot be modified from within this function.
	SelectedForeach(fn ListBoxForeachFunc)
	// SetActivateOnSingleClick: if @single is true, rows will be activated when
	// you click on them, otherwise you need to double-click.
	SetActivateOnSingleClick(single bool)
	// SetAdjustment sets the adjustment (if any) that the widget uses to for
	// vertical scrolling. For instance, this is used to get the page size for
	// PageUp/Down key handling.
	//
	// In the normal case when the @box is packed inside a ScrolledWindow the
	// adjustment from that will be picked up automatically, so there is no need
	// to manually do that.
	SetAdjustment(adjustment Adjustment)
	// SetFilterFunc: by setting a filter function on the @box one can decide
	// dynamically which of the rows to show. For instance, to implement a
	// search function on a list that filters the original list to only show the
	// matching rows.
	//
	// The @filter_func will be called for each row after the call, and it will
	// continue to be called each time a row changes (via
	// gtk_list_box_row_changed()) or when gtk_list_box_invalidate_filter() is
	// called.
	//
	// Note that using a filter function is incompatible with using a model (see
	// gtk_list_box_bind_model()).
	SetFilterFunc(filterFunc ListBoxFilterFunc)
	// SetHeaderFunc: by setting a header function on the @box one can
	// dynamically add headers in front of rows, depending on the contents of
	// the row and its position in the list. For instance, one could use it to
	// add headers in front of the first item of a new kind, in a list sorted by
	// the kind.
	//
	// The @update_header can look at the current header widget using
	// gtk_list_box_row_get_header() and either update the state of the widget
	// as needed, or set a new one using gtk_list_box_row_set_header(). If no
	// header is needed, set the header to nil.
	//
	// Note that you may get many calls @update_header to this for a particular
	// row when e.g. changing things that don’t affect the header. In this case
	// it is important for performance to not blindly replace an existing header
	// with an identical one.
	//
	// The @update_header function will be called for each row after the call,
	// and it will continue to be called each time a row changes (via
	// gtk_list_box_row_changed()) and when the row before changes (either by
	// gtk_list_box_row_changed() on the previous row, or when the previous row
	// becomes a different row). It is also called for all rows when
	// gtk_list_box_invalidate_headers() is called.
	SetHeaderFunc(updateHeader ListBoxUpdateHeaderFunc)
	// SetPlaceholder sets the placeholder widget that is shown in the list when
	// it doesn't display any visible children.
	SetPlaceholder(placeholder Widget)
	// SetSelectionMode sets how selection works in the listbox. See
	// SelectionMode for details.
	SetSelectionMode(mode SelectionMode)
	// SetSortFunc: by setting a sort function on the @box one can dynamically
	// reorder the rows of the list, based on the contents of the rows.
	//
	// The @sort_func will be called for each row after the call, and will
	// continue to be called each time a row changes (via
	// gtk_list_box_row_changed()) and when gtk_list_box_invalidate_sort() is
	// called.
	//
	// Note that using a sort function is incompatible with using a model (see
	// gtk_list_box_bind_model()).
	SetSortFunc(sortFunc ListBoxSortFunc)
	// UnselectAll: unselect all children of @box, if the selection mode allows
	// it.
	UnselectAll()
	// UnselectRow unselects a single row of @box, if the selection mode allows
	// it.
	UnselectRow(row ListBoxRow)
}

// listBox implements the ListBox interface.
type listBox struct {
	Container
	Buildable
}

var _ ListBox = (*listBox)(nil)

// WrapListBox wraps a GObject to the right type. It is
// primarily used internally.
func WrapListBox(obj *externglib.Object) ListBox {
	return ListBox{
		Container: WrapContainer(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalListBox(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapListBox(obj), nil
}

// NewListBox constructs a class ListBox.
func NewListBox() ListBox {
	var cret C.GtkListBox
	var goret1 ListBox

	cret = C.gtk_list_box_new()

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ListBox)

	return goret1
}

// BindModel binds @model to @box.
//
// If @box was already bound to a model, that previous binding is destroyed.
//
// The contents of @box are cleared and then filled with widgets that
// represent items from @model. @box is updated whenever @model changes. If
// @model is nil, @box is left empty.
//
// It is undefined to add or remove widgets directly (for example, with
// gtk_list_box_insert() or gtk_container_add()) while @box is bound to a
// model.
//
// Note that using a model is incompatible with the filtering and sorting
// functionality in GtkListBox. When using a model, filtering and sorting
// should be implemented by the model.
func (b listBox) BindModel(model gio.ListModel, createWidgetFunc ListBoxCreateWidgetFunc) {
	var arg0 *C.GtkListBox

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_bind_model(arg0, model, createWidgetFunc, userData, userDataFreeFunc)
}

// DragHighlightRow: this is a helper function for implementing DnD onto a
// ListBox. The passed in @row will be highlighted via gtk_drag_highlight(),
// and any previously highlighted row will be unhighlighted.
//
// The row will also be unhighlighted when the widget gets a drag leave
// event.
func (b listBox) DragHighlightRow(row ListBoxRow) {
	var arg0 *C.GtkListBox
	var arg1 *C.GtkListBoxRow

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	C.gtk_list_box_drag_highlight_row(arg0, row)
}

// DragUnhighlightRow: if a row has previously been highlighted via
// gtk_list_box_drag_highlight_row() it will have the highlight removed.
func (b listBox) DragUnhighlightRow() {
	var arg0 *C.GtkListBox

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_drag_unhighlight_row(arg0)
}

// ActivateOnSingleClick returns whether rows activate on single clicks.
func (b listBox) ActivateOnSingleClick() bool {
	var arg0 *C.GtkListBox

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_list_box_get_activate_on_single_click(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// Adjustment gets the adjustment (if any) that the widget uses to for
// vertical scrolling.
func (b listBox) Adjustment() Adjustment {
	var arg0 *C.GtkListBox

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	var cret *C.GtkAdjustment
	var goret1 Adjustment

	cret = C.gtk_list_box_get_adjustment(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Adjustment)

	return goret1
}

// RowAtIndex gets the n-th child in the list (not counting headers). If
// @_index is negative or larger than the number of items in the list, nil
// is returned.
func (b listBox) RowAtIndex(index_ int) ListBoxRow {
	var arg0 *C.GtkListBox
	var arg1 C.gint

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	arg1 = C.gint(index_)

	var cret *C.GtkListBoxRow
	var goret1 ListBoxRow

	cret = C.gtk_list_box_get_row_at_index(arg0, index_)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ListBoxRow)

	return goret1
}

// RowAtY gets the row at the @y position.
func (b listBox) RowAtY(y int) ListBoxRow {
	var arg0 *C.GtkListBox
	var arg1 C.gint

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	arg1 = C.gint(y)

	var cret *C.GtkListBoxRow
	var goret1 ListBoxRow

	cret = C.gtk_list_box_get_row_at_y(arg0, y)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ListBoxRow)

	return goret1
}

// SelectedRow gets the selected row.
//
// Note that the box may allow multiple selection, in which case you should
// use gtk_list_box_selected_foreach() to find all selected rows.
func (b listBox) SelectedRow() ListBoxRow {
	var arg0 *C.GtkListBox

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	var cret *C.GtkListBoxRow
	var goret1 ListBoxRow

	cret = C.gtk_list_box_get_selected_row(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ListBoxRow)

	return goret1
}

// SelectedRows creates a list of all selected children.
func (b listBox) SelectedRows() *glib.List {
	var arg0 *C.GtkListBox

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	var cret *C.GList
	var goret1 *glib.List

	cret = C.gtk_list_box_get_selected_rows(arg0)

	goret1 = glib.WrapList(unsafe.Pointer(cret))
	runtime.SetFinalizer(goret1, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return goret1
}

// SelectionMode gets the selection mode of the listbox.
func (b listBox) SelectionMode() SelectionMode {
	var arg0 *C.GtkListBox

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	var cret C.GtkSelectionMode
	var goret1 SelectionMode

	cret = C.gtk_list_box_get_selection_mode(arg0)

	goret1 = SelectionMode(cret)

	return goret1
}

// Insert: insert the @child into the @box at @position. If a sort function
// is set, the widget will actually be inserted at the calculated position
// and this function has the same effect of gtk_container_add().
//
// If @position is -1, or larger than the total number of items in the @box,
// then the @child will be appended to the end.
func (b listBox) Insert(child Widget, position int) {
	var arg0 *C.GtkListBox
	var arg1 *C.GtkWidget
	var arg2 C.gint

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	arg2 = C.gint(position)

	C.gtk_list_box_insert(arg0, child, position)
}

// InvalidateFilter: update the filtering for all rows. Call this when
// result of the filter function on the @box is changed due to an external
// factor. For instance, this would be used if the filter function just
// looked for a specific search string and the entry with the search string
// has changed.
func (b listBox) InvalidateFilter() {
	var arg0 *C.GtkListBox

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_invalidate_filter(arg0)
}

// InvalidateHeaders: update the separators for all rows. Call this when
// result of the header function on the @box is changed due to an external
// factor.
func (b listBox) InvalidateHeaders() {
	var arg0 *C.GtkListBox

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_invalidate_headers(arg0)
}

// InvalidateSort: update the sorting for all rows. Call this when result of
// the sort function on the @box is changed due to an external factor.
func (b listBox) InvalidateSort() {
	var arg0 *C.GtkListBox

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_invalidate_sort(arg0)
}

// Prepend: prepend a widget to the list. If a sort function is set, the
// widget will actually be inserted at the calculated position and this
// function has the same effect of gtk_container_add().
func (b listBox) Prepend(child Widget) {
	var arg0 *C.GtkListBox
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_list_box_prepend(arg0, child)
}

// SelectAll: select all children of @box, if the selection mode allows it.
func (b listBox) SelectAll() {
	var arg0 *C.GtkListBox

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_select_all(arg0)
}

// SelectRow: make @row the currently selected row.
func (b listBox) SelectRow(row ListBoxRow) {
	var arg0 *C.GtkListBox
	var arg1 *C.GtkListBoxRow

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	C.gtk_list_box_select_row(arg0, row)
}

// SelectedForeach calls a function for each selected child.
//
// Note that the selection cannot be modified from within this function.
func (b listBox) SelectedForeach(fn ListBoxForeachFunc) {
	var arg0 *C.GtkListBox

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_selected_foreach(arg0, fn, data)
}

// SetActivateOnSingleClick: if @single is true, rows will be activated when
// you click on them, otherwise you need to double-click.
func (b listBox) SetActivateOnSingleClick(single bool) {
	var arg0 *C.GtkListBox
	var arg1 C.gboolean

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	if single {
		arg1 = C.gboolean(1)
	}

	C.gtk_list_box_set_activate_on_single_click(arg0, single)
}

// SetAdjustment sets the adjustment (if any) that the widget uses to for
// vertical scrolling. For instance, this is used to get the page size for
// PageUp/Down key handling.
//
// In the normal case when the @box is packed inside a ScrolledWindow the
// adjustment from that will be picked up automatically, so there is no need
// to manually do that.
func (b listBox) SetAdjustment(adjustment Adjustment) {
	var arg0 *C.GtkListBox
	var arg1 *C.GtkAdjustment

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_list_box_set_adjustment(arg0, adjustment)
}

// SetFilterFunc: by setting a filter function on the @box one can decide
// dynamically which of the rows to show. For instance, to implement a
// search function on a list that filters the original list to only show the
// matching rows.
//
// The @filter_func will be called for each row after the call, and it will
// continue to be called each time a row changes (via
// gtk_list_box_row_changed()) or when gtk_list_box_invalidate_filter() is
// called.
//
// Note that using a filter function is incompatible with using a model (see
// gtk_list_box_bind_model()).
func (b listBox) SetFilterFunc(filterFunc ListBoxFilterFunc) {
	var arg0 *C.GtkListBox

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_set_filter_func(arg0, filterFunc, userData, destroy)
}

// SetHeaderFunc: by setting a header function on the @box one can
// dynamically add headers in front of rows, depending on the contents of
// the row and its position in the list. For instance, one could use it to
// add headers in front of the first item of a new kind, in a list sorted by
// the kind.
//
// The @update_header can look at the current header widget using
// gtk_list_box_row_get_header() and either update the state of the widget
// as needed, or set a new one using gtk_list_box_row_set_header(). If no
// header is needed, set the header to nil.
//
// Note that you may get many calls @update_header to this for a particular
// row when e.g. changing things that don’t affect the header. In this case
// it is important for performance to not blindly replace an existing header
// with an identical one.
//
// The @update_header function will be called for each row after the call,
// and it will continue to be called each time a row changes (via
// gtk_list_box_row_changed()) and when the row before changes (either by
// gtk_list_box_row_changed() on the previous row, or when the previous row
// becomes a different row). It is also called for all rows when
// gtk_list_box_invalidate_headers() is called.
func (b listBox) SetHeaderFunc(updateHeader ListBoxUpdateHeaderFunc) {
	var arg0 *C.GtkListBox

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_set_header_func(arg0, updateHeader, userData, destroy)
}

// SetPlaceholder sets the placeholder widget that is shown in the list when
// it doesn't display any visible children.
func (b listBox) SetPlaceholder(placeholder Widget) {
	var arg0 *C.GtkListBox
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(placeholder.Native()))

	C.gtk_list_box_set_placeholder(arg0, placeholder)
}

// SetSelectionMode sets how selection works in the listbox. See
// SelectionMode for details.
func (b listBox) SetSelectionMode(mode SelectionMode) {
	var arg0 *C.GtkListBox
	var arg1 C.GtkSelectionMode

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	arg1 = (C.GtkSelectionMode)(mode)

	C.gtk_list_box_set_selection_mode(arg0, mode)
}

// SetSortFunc: by setting a sort function on the @box one can dynamically
// reorder the rows of the list, based on the contents of the rows.
//
// The @sort_func will be called for each row after the call, and will
// continue to be called each time a row changes (via
// gtk_list_box_row_changed()) and when gtk_list_box_invalidate_sort() is
// called.
//
// Note that using a sort function is incompatible with using a model (see
// gtk_list_box_bind_model()).
func (b listBox) SetSortFunc(sortFunc ListBoxSortFunc) {
	var arg0 *C.GtkListBox

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_set_sort_func(arg0, sortFunc, userData, destroy)
}

// UnselectAll: unselect all children of @box, if the selection mode allows
// it.
func (b listBox) UnselectAll() {
	var arg0 *C.GtkListBox

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_unselect_all(arg0)
}

// UnselectRow unselects a single row of @box, if the selection mode allows
// it.
func (b listBox) UnselectRow(row ListBoxRow) {
	var arg0 *C.GtkListBox
	var arg1 *C.GtkListBoxRow

	arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	C.gtk_list_box_unselect_row(arg0, row)
}

type ListBoxRow interface {
	Bin
	Actionable
	Buildable

	// Changed marks @row as changed, causing any state that depends on this to
	// be updated. This affects sorting, filtering and headers.
	//
	// Note that calls to this method must be in sync with the data used for the
	// row functions. For instance, if the list is mirroring some external data
	// set, and *two* rows changed in the external data set then when you call
	// gtk_list_box_row_changed() on the first row the sort function must only
	// read the new data for the first of the two changed rows, otherwise the
	// resorting of the rows will be wrong.
	//
	// This generally means that if you don’t fully control the data model you
	// have to duplicate the data that affects the listbox row functions into
	// the row widgets themselves. Another alternative is to call
	// gtk_list_box_invalidate_sort() on any model change, but that is more
	// expensive.
	Changed()
	// Activatable gets the value of the ListBoxRow:activatable property for
	// this row.
	Activatable() bool
	// Header returns the current header of the @row. This can be used in a
	// ListBoxUpdateHeaderFunc to see if there is a header set already, and if
	// so to update the state of it.
	Header() Widget
	// Index gets the current index of the @row in its ListBox container.
	Index() int
	// Selectable gets the value of the ListBoxRow:selectable property for this
	// row.
	Selectable() bool
	// IsSelected returns whether the child is currently selected in its ListBox
	// container.
	IsSelected() bool
	// SetActivatable: set the ListBoxRow:activatable property for this row.
	SetActivatable(activatable bool)
	// SetHeader sets the current header of the @row. This is only allowed to be
	// called from a ListBoxUpdateHeaderFunc. It will replace any existing
	// header in the row, and be shown in front of the row in the listbox.
	SetHeader(header Widget)
	// SetSelectable: set the ListBoxRow:selectable property for this row.
	SetSelectable(selectable bool)
}

// listBoxRow implements the ListBoxRow interface.
type listBoxRow struct {
	Bin
	Actionable
	Buildable
}

var _ ListBoxRow = (*listBoxRow)(nil)

// WrapListBoxRow wraps a GObject to the right type. It is
// primarily used internally.
func WrapListBoxRow(obj *externglib.Object) ListBoxRow {
	return ListBoxRow{
		Bin:        WrapBin(obj),
		Actionable: WrapActionable(obj),
		Buildable:  WrapBuildable(obj),
	}
}

func marshalListBoxRow(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapListBoxRow(obj), nil
}

// NewListBoxRow constructs a class ListBoxRow.
func NewListBoxRow() ListBoxRow {
	var cret C.GtkListBoxRow
	var goret1 ListBoxRow

	cret = C.gtk_list_box_row_new()

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(ListBoxRow)

	return goret1
}

// Changed marks @row as changed, causing any state that depends on this to
// be updated. This affects sorting, filtering and headers.
//
// Note that calls to this method must be in sync with the data used for the
// row functions. For instance, if the list is mirroring some external data
// set, and *two* rows changed in the external data set then when you call
// gtk_list_box_row_changed() on the first row the sort function must only
// read the new data for the first of the two changed rows, otherwise the
// resorting of the rows will be wrong.
//
// This generally means that if you don’t fully control the data model you
// have to duplicate the data that affects the listbox row functions into
// the row widgets themselves. Another alternative is to call
// gtk_list_box_invalidate_sort() on any model change, but that is more
// expensive.
func (r listBoxRow) Changed() {
	var arg0 *C.GtkListBoxRow

	arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	C.gtk_list_box_row_changed(arg0)
}

// Activatable gets the value of the ListBoxRow:activatable property for
// this row.
func (r listBoxRow) Activatable() bool {
	var arg0 *C.GtkListBoxRow

	arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_list_box_row_get_activatable(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// Header returns the current header of the @row. This can be used in a
// ListBoxUpdateHeaderFunc to see if there is a header set already, and if
// so to update the state of it.
func (r listBoxRow) Header() Widget {
	var arg0 *C.GtkListBoxRow

	arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	var cret *C.GtkWidget
	var goret1 Widget

	cret = C.gtk_list_box_row_get_header(arg0)

	goret1 = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return goret1
}

// Index gets the current index of the @row in its ListBox container.
func (r listBoxRow) Index() int {
	var arg0 *C.GtkListBoxRow

	arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	var cret C.gint
	var goret1 int

	cret = C.gtk_list_box_row_get_index(arg0)

	goret1 = C.gint(cret)

	return goret1
}

// Selectable gets the value of the ListBoxRow:selectable property for this
// row.
func (r listBoxRow) Selectable() bool {
	var arg0 *C.GtkListBoxRow

	arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_list_box_row_get_selectable(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// IsSelected returns whether the child is currently selected in its ListBox
// container.
func (r listBoxRow) IsSelected() bool {
	var arg0 *C.GtkListBoxRow

	arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	var cret C.gboolean
	var goret1 bool

	cret = C.gtk_list_box_row_is_selected(arg0)

	goret1 = C.bool(cret) != C.false

	return goret1
}

// SetActivatable: set the ListBoxRow:activatable property for this row.
func (r listBoxRow) SetActivatable(activatable bool) {
	var arg0 *C.GtkListBoxRow
	var arg1 C.gboolean

	arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))
	if activatable {
		arg1 = C.gboolean(1)
	}

	C.gtk_list_box_row_set_activatable(arg0, activatable)
}

// SetHeader sets the current header of the @row. This is only allowed to be
// called from a ListBoxUpdateHeaderFunc. It will replace any existing
// header in the row, and be shown in front of the row in the listbox.
func (r listBoxRow) SetHeader(header Widget) {
	var arg0 *C.GtkListBoxRow
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(header.Native()))

	C.gtk_list_box_row_set_header(arg0, header)
}

// SetSelectable: set the ListBoxRow:selectable property for this row.
func (r listBoxRow) SetSelectable(selectable bool) {
	var arg0 *C.GtkListBoxRow
	var arg1 C.gboolean

	arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))
	if selectable {
		arg1 = C.gboolean(1)
	}

	C.gtk_list_box_row_set_selectable(arg0, selectable)
}
