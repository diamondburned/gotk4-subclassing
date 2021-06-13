// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_column_view_column_get_type()), F: marshalColumnViewColumn},
	})
}

// ColumnViewColumn: `GtkColumnViewColumn` represents the columns being added to
// `GtkColumnView`.
//
// The main ingredient for a `GtkColumnViewColumn` is the `GtkListItemFactory`
// that tells the columnview how to create cells for this column from items in
// the model.
//
// Columns have a title, and can optionally have a header menu set with
// [method@Gtk.ColumnViewColumn.set_header_menu].
//
// A sorter can be associated with a column using
// [method@Gtk.ColumnViewColumn.set_sorter], to let users influence sorting by
// clicking on the column header.
type ColumnViewColumn interface {
	gextras.Objector

	// Expand returns whether this column should expand.
	Expand() bool
	// FixedWidth gets the fixed width of the column.
	FixedWidth() int
	// Resizable returns whether this column is resizable.
	Resizable() bool
	// Title returns the title set with gtk_column_view_column_set_title().
	Title() string
	// Visible returns whether this column is visible.
	Visible() bool
	// SetExpand sets the column to take available extra space.
	//
	// The extra space is shared equally amongst all columns that have the
	// expand set to true.
	SetExpand(expand bool)
	// SetFactory sets the `GtkListItemFactory` to use for populating list items
	// for this column.
	SetFactory(factory ListItemFactory)
	// SetFixedWidth: if @fixed_width is not -1, sets the fixed width of
	// @column; otherwise unsets it.
	//
	// Setting a fixed width overrides the automatically calculated width.
	// Interactive resizing also sets the “fixed-width” property.
	SetFixedWidth(fixedWidth int)
	// SetHeaderMenu sets the menu model that is used to create the context menu
	// for the column header.
	SetHeaderMenu(menu gio.MenuModel)
	// SetResizable sets whether this column should be resizable by dragging.
	SetResizable(resizable bool)
	// SetSorter associates a sorter with the column.
	//
	// If @sorter is nil, the column will not let users change the sorting by
	// clicking on its header.
	//
	// This sorter can be made active by clicking on the column header, or by
	// calling [method@Gtk.ColumnView.sort_by_column].
	//
	// See [method@Gtk.ColumnView.get_sorter] for the necessary steps for
	// setting up customizable sorting for [class@Gtk.ColumnView].
	SetSorter(sorter Sorter)
	// SetTitle sets the title of this column.
	//
	// The title is displayed in the header of a `GtkColumnView` for this column
	// and is therefore user-facing text that should be translated.
	SetTitle(title string)
	// SetVisible sets whether this column should be visible in views.
	SetVisible(visible bool)
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

// Expand returns whether this column should expand.
func (s columnViewColumn) Expand() bool {
	var _arg0 *C.GtkColumnViewColumn // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_column_view_column_get_expand(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// FixedWidth gets the fixed width of the column.
func (s columnViewColumn) FixedWidth() int {
	var _arg0 *C.GtkColumnViewColumn // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	var _cret C.int // in

	_cret = C.gtk_column_view_column_get_fixed_width(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Resizable returns whether this column is resizable.
func (s columnViewColumn) Resizable() bool {
	var _arg0 *C.GtkColumnViewColumn // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_column_view_column_get_resizable(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Title returns the title set with gtk_column_view_column_set_title().
func (s columnViewColumn) Title() string {
	var _arg0 *C.GtkColumnViewColumn // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	var _cret *C.char // in

	_cret = C.gtk_column_view_column_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Visible returns whether this column is visible.
func (s columnViewColumn) Visible() bool {
	var _arg0 *C.GtkColumnViewColumn // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_column_view_column_get_visible(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// SetExpand sets the column to take available extra space.
//
// The extra space is shared equally amongst all columns that have the
// expand set to true.
func (s columnViewColumn) SetExpand(expand bool) {
	var _arg0 *C.GtkColumnViewColumn // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	if expand {
		_arg1 = C.gboolean(1)
	}

	C.gtk_column_view_column_set_expand(_arg0, _arg1)
}

// SetFactory sets the `GtkListItemFactory` to use for populating list items
// for this column.
func (s columnViewColumn) SetFactory(factory ListItemFactory) {
	var _arg0 *C.GtkColumnViewColumn // out
	var _arg1 *C.GtkListItemFactory  // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))

	C.gtk_column_view_column_set_factory(_arg0, _arg1)
}

// SetFixedWidth: if @fixed_width is not -1, sets the fixed width of
// @column; otherwise unsets it.
//
// Setting a fixed width overrides the automatically calculated width.
// Interactive resizing also sets the “fixed-width” property.
func (s columnViewColumn) SetFixedWidth(fixedWidth int) {
	var _arg0 *C.GtkColumnViewColumn // out
	var _arg1 C.int                  // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	_arg1 = C.int(fixedWidth)

	C.gtk_column_view_column_set_fixed_width(_arg0, _arg1)
}

// SetHeaderMenu sets the menu model that is used to create the context menu
// for the column header.
func (s columnViewColumn) SetHeaderMenu(menu gio.MenuModel) {
	var _arg0 *C.GtkColumnViewColumn // out
	var _arg1 *C.GMenuModel          // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GMenuModel)(unsafe.Pointer(menu.Native()))

	C.gtk_column_view_column_set_header_menu(_arg0, _arg1)
}

// SetResizable sets whether this column should be resizable by dragging.
func (s columnViewColumn) SetResizable(resizable bool) {
	var _arg0 *C.GtkColumnViewColumn // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	if resizable {
		_arg1 = C.gboolean(1)
	}

	C.gtk_column_view_column_set_resizable(_arg0, _arg1)
}

// SetSorter associates a sorter with the column.
//
// If @sorter is nil, the column will not let users change the sorting by
// clicking on its header.
//
// This sorter can be made active by clicking on the column header, or by
// calling [method@Gtk.ColumnView.sort_by_column].
//
// See [method@Gtk.ColumnView.get_sorter] for the necessary steps for
// setting up customizable sorting for [class@Gtk.ColumnView].
func (s columnViewColumn) SetSorter(sorter Sorter) {
	var _arg0 *C.GtkColumnViewColumn // out
	var _arg1 *C.GtkSorter           // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GtkSorter)(unsafe.Pointer(sorter.Native()))

	C.gtk_column_view_column_set_sorter(_arg0, _arg1)
}

// SetTitle sets the title of this column.
//
// The title is displayed in the header of a `GtkColumnView` for this column
// and is therefore user-facing text that should be translated.
func (s columnViewColumn) SetTitle(title string) {
	var _arg0 *C.GtkColumnViewColumn // out
	var _arg1 *C.char                // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.char)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_column_view_column_set_title(_arg0, _arg1)
}

// SetVisible sets whether this column should be visible in views.
func (s columnViewColumn) SetVisible(visible bool) {
	var _arg0 *C.GtkColumnViewColumn // out
	var _arg1 C.gboolean             // out

	_arg0 = (*C.GtkColumnViewColumn)(unsafe.Pointer(s.Native()))
	if visible {
		_arg1 = C.gboolean(1)
	}

	C.gtk_column_view_column_set_visible(_arg0, _arg1)
}
