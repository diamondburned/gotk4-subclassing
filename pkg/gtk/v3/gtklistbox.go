// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
type ListBoxCreateWidgetFunc func(item gextras.Objector) (widget Widget)

//export gotk4_ListBoxCreateWidgetFunc
func gotk4_ListBoxCreateWidgetFunc(arg0 C.gpointer, arg1 C.gpointer) *C.GtkWidget {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var item gextras.Objector // out

	item = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0.Native()))).(gextras.Objector)

	fn := v.(ListBoxCreateWidgetFunc)
	widget := fn(item)

	cret = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	return widget
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
	// SetPlaceholder sets the placeholder widget that is shown in the list when
	// it doesn't display any visible children.
	SetPlaceholder(placeholder Widget)
	// SetSelectionMode sets how selection works in the listbox. See
	// SelectionMode for details.
	SetSelectionMode(mode SelectionMode)
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

// DragHighlightRow: this is a helper function for implementing DnD onto a
// ListBox. The passed in @row will be highlighted via gtk_drag_highlight(),
// and any previously highlighted row will be unhighlighted.
//
// The row will also be unhighlighted when the widget gets a drag leave
// event.
func (b listBox) DragHighlightRow(row ListBoxRow) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	C.gtk_list_box_drag_highlight_row(_arg0, _arg1)
}

// DragUnhighlightRow: if a row has previously been highlighted via
// gtk_list_box_drag_highlight_row() it will have the highlight removed.
func (b listBox) DragUnhighlightRow() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_drag_unhighlight_row(_arg0)
}

// ActivateOnSingleClick returns whether rows activate on single clicks.
func (b listBox) ActivateOnSingleClick() bool {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_list_box_get_activate_on_single_click(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Insert: insert the @child into the @box at @position. If a sort function
// is set, the widget will actually be inserted at the calculated position
// and this function has the same effect of gtk_container_add().
//
// If @position is -1, or larger than the total number of items in the @box,
// then the @child will be appended to the end.
func (b listBox) Insert(child Widget, position int) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.gint        // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.gint(position)

	C.gtk_list_box_insert(_arg0, _arg1, _arg2)
}

// InvalidateFilter: update the filtering for all rows. Call this when
// result of the filter function on the @box is changed due to an external
// factor. For instance, this would be used if the filter function just
// looked for a specific search string and the entry with the search string
// has changed.
func (b listBox) InvalidateFilter() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_invalidate_filter(_arg0)
}

// InvalidateHeaders: update the separators for all rows. Call this when
// result of the header function on the @box is changed due to an external
// factor.
func (b listBox) InvalidateHeaders() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_invalidate_headers(_arg0)
}

// InvalidateSort: update the sorting for all rows. Call this when result of
// the sort function on the @box is changed due to an external factor.
func (b listBox) InvalidateSort() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_invalidate_sort(_arg0)
}

// Prepend: prepend a widget to the list. If a sort function is set, the
// widget will actually be inserted at the calculated position and this
// function has the same effect of gtk_container_add().
func (b listBox) Prepend(child Widget) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_list_box_prepend(_arg0, _arg1)
}

// SelectAll: select all children of @box, if the selection mode allows it.
func (b listBox) SelectAll() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_select_all(_arg0)
}

// SelectRow: make @row the currently selected row.
func (b listBox) SelectRow(row ListBoxRow) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	C.gtk_list_box_select_row(_arg0, _arg1)
}

// SelectedForeach calls a function for each selected child.
//
// Note that the selection cannot be modified from within this function.
func (b listBox) SelectedForeach(fn ListBoxForeachFunc) {
	var _arg0 *C.GtkListBox           // out
	var _arg1 C.GtkListBoxForeachFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*[0]byte)(C.gotk4_ListBoxForeachFunc)
	_arg2 = C.gpointer(box.Assign(fn))

	C.gtk_list_box_selected_foreach(_arg0, _arg1, _arg2)
}

// SetActivateOnSingleClick: if @single is true, rows will be activated when
// you click on them, otherwise you need to double-click.
func (b listBox) SetActivateOnSingleClick(single bool) {
	var _arg0 *C.GtkListBox // out
	var _arg1 C.gboolean    // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	if single {
		_arg1 = C.gboolean(1)
	}

	C.gtk_list_box_set_activate_on_single_click(_arg0, _arg1)
}

// SetAdjustment sets the adjustment (if any) that the widget uses to for
// vertical scrolling. For instance, this is used to get the page size for
// PageUp/Down key handling.
//
// In the normal case when the @box is packed inside a ScrolledWindow the
// adjustment from that will be picked up automatically, so there is no need
// to manually do that.
func (b listBox) SetAdjustment(adjustment Adjustment) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_list_box_set_adjustment(_arg0, _arg1)
}

// SetPlaceholder sets the placeholder widget that is shown in the list when
// it doesn't display any visible children.
func (b listBox) SetPlaceholder(placeholder Widget) {
	var _arg0 *C.GtkListBox // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(placeholder.Native()))

	C.gtk_list_box_set_placeholder(_arg0, _arg1)
}

// SetSelectionMode sets how selection works in the listbox. See
// SelectionMode for details.
func (b listBox) SetSelectionMode(mode SelectionMode) {
	var _arg0 *C.GtkListBox      // out
	var _arg1 C.GtkSelectionMode // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = (C.GtkSelectionMode)(mode)

	C.gtk_list_box_set_selection_mode(_arg0, _arg1)
}

// UnselectAll: unselect all children of @box, if the selection mode allows
// it.
func (b listBox) UnselectAll() {
	var _arg0 *C.GtkListBox // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))

	C.gtk_list_box_unselect_all(_arg0)
}

// UnselectRow unselects a single row of @box, if the selection mode allows
// it.
func (b listBox) UnselectRow(row ListBoxRow) {
	var _arg0 *C.GtkListBox    // out
	var _arg1 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBox)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkListBoxRow)(unsafe.Pointer(row.Native()))

	C.gtk_list_box_unselect_row(_arg0, _arg1)
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
	var _arg0 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	C.gtk_list_box_row_changed(_arg0)
}

// Activatable gets the value of the ListBoxRow:activatable property for
// this row.
func (r listBoxRow) Activatable() bool {
	var _arg0 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_list_box_row_get_activatable(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Index gets the current index of the @row in its ListBox container.
func (r listBoxRow) Index() int {
	var _arg0 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	var _cret C.gint // in

	_cret = C.gtk_list_box_row_get_index(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Selectable gets the value of the ListBoxRow:selectable property for this
// row.
func (r listBoxRow) Selectable() bool {
	var _arg0 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_list_box_row_get_selectable(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// IsSelected returns whether the child is currently selected in its ListBox
// container.
func (r listBoxRow) IsSelected() bool {
	var _arg0 *C.GtkListBoxRow // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_list_box_row_is_selected(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// SetActivatable: set the ListBoxRow:activatable property for this row.
func (r listBoxRow) SetActivatable(activatable bool) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))
	if activatable {
		_arg1 = C.gboolean(1)
	}

	C.gtk_list_box_row_set_activatable(_arg0, _arg1)
}

// SetHeader sets the current header of the @row. This is only allowed to be
// called from a ListBoxUpdateHeaderFunc. It will replace any existing
// header in the row, and be shown in front of the row in the listbox.
func (r listBoxRow) SetHeader(header Widget) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 *C.GtkWidget     // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(header.Native()))

	C.gtk_list_box_row_set_header(_arg0, _arg1)
}

// SetSelectable: set the ListBoxRow:selectable property for this row.
func (r listBoxRow) SetSelectable(selectable bool) {
	var _arg0 *C.GtkListBoxRow // out
	var _arg1 C.gboolean       // out

	_arg0 = (*C.GtkListBoxRow)(unsafe.Pointer(r.Native()))
	if selectable {
		_arg1 = C.gboolean(1)
	}

	C.gtk_list_box_row_set_selectable(_arg0, _arg1)
}
