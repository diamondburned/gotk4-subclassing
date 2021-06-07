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
		{T: externglib.Type(C.gtk_column_view_column_get_type()), F: marshalColumnViewColumn},
	})
}

// ColumnViewColumn: gtkColumnViewColumn represents the columns being added to
// ColumnView.
//
// Columns have a title, and can optionally have a header menu set with
// gtk_column_view_column_set_header_menu().
//
// A sorter can be associated with a column using
// gtk_column_view_column_set_sorter(), to let users influence sorting by
// clicking on the column header.
type ColumnViewColumn interface {
	gextras.Objector

	// ColumnView gets the column view that's currently displaying this column.
	//
	// If @self has not been added to a column view yet, nil is returned.
	ColumnView(s ColumnViewColumn)
	// Expand returns whether this column should expand.
	Expand(s ColumnViewColumn) bool
	// Factory gets the factory that's currently used to populate list items for
	// this column.
	Factory(s ColumnViewColumn)
	// FixedWidth gets the fixed width of the column.
	FixedWidth(s ColumnViewColumn)
	// HeaderMenu gets the menu model that is used to create the context menu
	// for the column header.
	HeaderMenu(s ColumnViewColumn)
	// Resizable returns whether this column is resizable.
	Resizable(s ColumnViewColumn) bool
	// Sorter returns the sorter that is associated with the column.
	Sorter(s ColumnViewColumn)
	// Title returns the title set with gtk_column_view_column_set_title().
	Title(s ColumnViewColumn)
	// Visible returns whether this column is visible.
	Visible(s ColumnViewColumn) bool
	// SetExpand sets the column to take available extra space.
	//
	// The extra space is shared equally amongst all columns that have the
	// expand set to true.
	SetExpand(s ColumnViewColumn, expand bool)
	// SetFactory sets the ListItemFactory to use for populating list items for
	// this column.
	SetFactory(s ColumnViewColumn, factory ListItemFactory)
	// SetFixedWidth: if @fixed_width is not -1, sets the fixed width of
	// @column; otherwise unsets it.
	//
	// Setting a fixed width overrides the automatically calculated width.
	// Interactive resizing also sets the “fixed-width” property.
	SetFixedWidth(s ColumnViewColumn, fixedWidth int)
	// SetHeaderMenu sets the menu model that is used to create the context menu
	// for the column header.
	SetHeaderMenu(s ColumnViewColumn, menu gio.MenuModel)
	// SetResizable sets whether this column should be resizable by dragging.
	SetResizable(s ColumnViewColumn, resizable bool)
	// SetSorter associates a sorter with the column.
	//
	// If @sorter is nil, the column will not let users change the sorting by
	// clicking on its header.
	//
	// This sorter can be made active by clicking on the column header, or by
	// calling gtk_column_view_sort_by_column().
	//
	// See gtk_column_view_get_sorter() for the necessary steps for setting up
	// customizable sorting for ColumnView.
	SetSorter(s ColumnViewColumn, sorter Sorter)
	// SetTitle sets the title of this column. The title is displayed in the
	// header of a ColumnView for this column and is therefore user-facing text
	// that should be translated.
	SetTitle(s ColumnViewColumn, title string)
	// SetVisible sets whether this column should be visible in views.
	SetVisible(s ColumnViewColumn, visible bool)
}

// columnViewColumn implements the ColumnViewColumn interface.
type columnViewColumn struct {
	gextras.Objector
}

var _ ColumnViewColumn = (*columnViewColumn)(nil)

// WrapColumnViewColumn wraps a GObject to the right type. It is
// primarily used internally.
func WrapColumnViewColumn(obj *externglib.Object) ColumnViewColumn {
	return ColumnViewColumn{
		Objector: obj,
	}
}

func marshalColumnViewColumn(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapColumnViewColumn(obj), nil
}

// NewColumnViewColumn constructs a class ColumnViewColumn.
func NewColumnViewColumn(title string, factory ListItemFactory) {
	var arg1 *C.char
	var arg2 *C.GtkListItemFactory

	arg1 = (*C.char)(C.CString(title))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))

	C.gtk_column_view_column_new(arg1, arg2)
}

// ColumnView gets the column view that's currently displaying this column.
//
// If @self has not been added to a column view yet, nil is returned.
func (s columnViewColumn) ColumnView(s ColumnViewColumn) {
	var arg0 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	C.gtk_column_view_column_get_column_view(arg0)
}

// Expand returns whether this column should expand.
func (s columnViewColumn) Expand(s ColumnViewColumn) bool {
	var arg0 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_column_view_column_get_expand(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Factory gets the factory that's currently used to populate list items for
// this column.
func (s columnViewColumn) Factory(s ColumnViewColumn) {
	var arg0 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	C.gtk_column_view_column_get_factory(arg0)
}

// FixedWidth gets the fixed width of the column.
func (s columnViewColumn) FixedWidth(s ColumnViewColumn) {
	var arg0 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	C.gtk_column_view_column_get_fixed_width(arg0)
}

// HeaderMenu gets the menu model that is used to create the context menu
// for the column header.
func (s columnViewColumn) HeaderMenu(s ColumnViewColumn) {
	var arg0 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	C.gtk_column_view_column_get_header_menu(arg0)
}

// Resizable returns whether this column is resizable.
func (s columnViewColumn) Resizable(s ColumnViewColumn) bool {
	var arg0 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_column_view_column_get_resizable(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Sorter returns the sorter that is associated with the column.
func (s columnViewColumn) Sorter(s ColumnViewColumn) {
	var arg0 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	C.gtk_column_view_column_get_sorter(arg0)
}

// Title returns the title set with gtk_column_view_column_set_title().
func (s columnViewColumn) Title(s ColumnViewColumn) {
	var arg0 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	C.gtk_column_view_column_get_title(arg0)
}

// Visible returns whether this column is visible.
func (s columnViewColumn) Visible(s ColumnViewColumn) bool {
	var arg0 *C.GtkColumnViewColumn

	arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_column_view_column_get_visible(arg0)

	if cret {
		ok = true
	}

	return ok
}

// SetExpand sets the column to take available extra space.
//
// The extra space is shared equally amongst all columns that have the
// expand set to true.
func (s columnViewColumn) SetExpand(s ColumnViewColumn, expand bool) {
	var arg0 *C.GtkColumnViewColumn
	var arg1 C.gboolean

	arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	if expand {
		arg1 = C.gboolean(1)
	}

	C.gtk_column_view_column_set_expand(arg0, arg1)
}

// SetFactory sets the ListItemFactory to use for populating list items for
// this column.
func (s columnViewColumn) SetFactory(s ColumnViewColumn, factory ListItemFactory) {
	var arg0 *C.GtkColumnViewColumn
	var arg1 *C.GtkListItemFactory

	arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))

	C.gtk_column_view_column_set_factory(arg0, arg1)
}

// SetFixedWidth: if @fixed_width is not -1, sets the fixed width of
// @column; otherwise unsets it.
//
// Setting a fixed width overrides the automatically calculated width.
// Interactive resizing also sets the “fixed-width” property.
func (s columnViewColumn) SetFixedWidth(s ColumnViewColumn, fixedWidth int) {
	var arg0 *C.GtkColumnViewColumn
	var arg1 C.int

	arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	arg1 = C.int(fixedWidth)

	C.gtk_column_view_column_set_fixed_width(arg0, arg1)
}

// SetHeaderMenu sets the menu model that is used to create the context menu
// for the column header.
func (s columnViewColumn) SetHeaderMenu(s ColumnViewColumn, menu gio.MenuModel) {
	var arg0 *C.GtkColumnViewColumn
	var arg1 *C.GMenuModel

	arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GMenuModel)(unsafe.Pointer(menu.Native()))

	C.gtk_column_view_column_set_header_menu(arg0, arg1)
}

// SetResizable sets whether this column should be resizable by dragging.
func (s columnViewColumn) SetResizable(s ColumnViewColumn, resizable bool) {
	var arg0 *C.GtkColumnViewColumn
	var arg1 C.gboolean

	arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	if resizable {
		arg1 = C.gboolean(1)
	}

	C.gtk_column_view_column_set_resizable(arg0, arg1)
}

// SetSorter associates a sorter with the column.
//
// If @sorter is nil, the column will not let users change the sorting by
// clicking on its header.
//
// This sorter can be made active by clicking on the column header, or by
// calling gtk_column_view_sort_by_column().
//
// See gtk_column_view_get_sorter() for the necessary steps for setting up
// customizable sorting for ColumnView.
func (s columnViewColumn) SetSorter(s ColumnViewColumn, sorter Sorter) {
	var arg0 *C.GtkColumnViewColumn
	var arg1 *C.GtkSorter

	arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkSorter)(unsafe.Pointer(sorter.Native()))

	C.gtk_column_view_column_set_sorter(arg0, arg1)
}

// SetTitle sets the title of this column. The title is displayed in the
// header of a ColumnView for this column and is therefore user-facing text
// that should be translated.
func (s columnViewColumn) SetTitle(s ColumnViewColumn, title string) {
	var arg0 *C.GtkColumnViewColumn
	var arg1 *C.char

	arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(title))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_column_view_column_set_title(arg0, arg1)
}

// SetVisible sets whether this column should be visible in views.
func (s columnViewColumn) SetVisible(s ColumnViewColumn, visible bool) {
	var arg0 *C.GtkColumnViewColumn
	var arg1 C.gboolean

	arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	if visible {
		arg1 = C.gboolean(1)
	}

	C.gtk_column_view_column_set_visible(arg0, arg1)
}
