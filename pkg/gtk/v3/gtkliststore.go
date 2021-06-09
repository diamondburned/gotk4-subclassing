// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/ptr"
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
		{T: externglib.Type(C.gtk_list_store_get_type()), F: marshalListStore},
	})
}

// ListStore: the ListStore object is a list model for use with a TreeView
// widget. It implements the TreeModel interface, and consequentialy, can use
// all of the methods available there. It also implements the TreeSortable
// interface so it can be sorted by the view. Finally, it also implements the
// tree [drag and drop][gtk3-GtkTreeView-drag-and-drop] interfaces.
//
// The ListStore can accept most GObject types as a column type, though it can’t
// accept all custom types. Internally, it will keep a copy of data passed in
// (such as a string or a boxed pointer). Columns that accept #GObjects are
// handled a little differently. The ListStore will keep a reference to the
// object instead of copying the value. As a result, if the object is modified,
// it is up to the application writer to call gtk_tree_model_row_changed() to
// emit the TreeModel::row_changed signal. This most commonly affects lists with
// Pixbufs stored.
//
// An example for creating a simple list store:
//
//    <object class="GtkListStore">
//      <columns>
//        <column type="gchararray"/>
//        <column type="gchararray"/>
//        <column type="gint"/>
//      </columns>
//      <data>
//        <row>
//          <col id="0">John</col>
//          <col id="1">Doe</col>
//          <col id="2">25</col>
//        </row>
//        <row>
//          <col id="0">Johan</col>
//          <col id="1">Dahlin</col>
//          <col id="2">50</col>
//        </row>
//      </data>
//    </object>
type ListStore interface {
	gextras.Objector
	Buildable
	TreeDragDest
	TreeDragSource
	TreeModel
	TreeSortable

	// Append appends a new row to @list_store. @iter will be changed to point
	// to this new row. The row will be empty after this function is called. To
	// fill in values, you need to call gtk_list_store_set() or
	// gtk_list_store_set_value().
	Append() TreeIter
	// Clear removes all rows from the list store.
	Clear()
	// Insert creates a new row at @position. @iter will be changed to point to
	// this new row. If @position is -1 or is larger than the number of rows on
	// the list, then the new row will be appended to the list. The row will be
	// empty after this function is called. To fill in values, you need to call
	// gtk_list_store_set() or gtk_list_store_set_value().
	Insert(position int) TreeIter
	// InsertAfter inserts a new row after @sibling. If @sibling is nil, then
	// the row will be prepended to the beginning of the list. @iter will be
	// changed to point to this new row. The row will be empty after this
	// function is called. To fill in values, you need to call
	// gtk_list_store_set() or gtk_list_store_set_value().
	InsertAfter(sibling *TreeIter) TreeIter
	// InsertBefore inserts a new row before @sibling. If @sibling is nil, then
	// the row will be appended to the end of the list. @iter will be changed to
	// point to this new row. The row will be empty after this function is
	// called. To fill in values, you need to call gtk_list_store_set() or
	// gtk_list_store_set_value().
	InsertBefore(sibling *TreeIter) TreeIter
	// InsertWithValuesv: a variant of gtk_list_store_insert_with_values() which
	// takes the columns and values as two arrays, instead of varargs. This
	// function is mainly intended for language-bindings.
	InsertWithValuesv() TreeIter
	// IterIsValid: > This function is slow. Only use it for debugging and/or
	// testing > purposes.
	//
	// Checks if the given iter is a valid iter for this ListStore.
	IterIsValid(iter *TreeIter) bool
	// MoveAfter moves @iter in @store to the position after @position. Note
	// that this function only works with unsorted stores. If @position is nil,
	// @iter will be moved to the start of the list.
	MoveAfter(iter *TreeIter, position *TreeIter)
	// MoveBefore moves @iter in @store to the position before @position. Note
	// that this function only works with unsorted stores. If @position is nil,
	// @iter will be moved to the end of the list.
	MoveBefore(iter *TreeIter, position *TreeIter)
	// Prepend prepends a new row to @list_store. @iter will be changed to point
	// to this new row. The row will be empty after this function is called. To
	// fill in values, you need to call gtk_list_store_set() or
	// gtk_list_store_set_value().
	Prepend() TreeIter
	// Remove removes the given row from the list store. After being removed,
	// @iter is set to be the next valid row, or invalidated if it pointed to
	// the last row in @list_store.
	Remove(iter *TreeIter) bool
	// Reorder reorders @store to follow the order indicated by @new_order. Note
	// that this function only works with unsorted stores.
	Reorder(newOrder []int)
	// SetValue sets the data in the cell specified by @iter and @column. The
	// type of @value must be convertible to the type of the column.
	SetValue(iter *TreeIter, column int, value **externglib.Value)
	// SetValuesv: a variant of gtk_list_store_set_valist() which takes the
	// columns and values as two arrays, instead of varargs. This function is
	// mainly intended for language-bindings and in case the number of columns
	// to change is not known until run-time.
	SetValuesv()
	// Swap swaps @a and @b in @store. Note that this function only works with
	// unsorted stores.
	Swap(a *TreeIter, b *TreeIter)
}

// listStore implements the ListStore interface.
type listStore struct {
	gextras.Objector
	Buildable
	TreeDragDest
	TreeDragSource
	TreeModel
	TreeSortable
}

var _ ListStore = (*listStore)(nil)

// WrapListStore wraps a GObject to the right type. It is
// primarily used internally.
func WrapListStore(obj *externglib.Object) ListStore {
	return ListStore{
		Objector:       obj,
		Buildable:      WrapBuildable(obj),
		TreeDragDest:   WrapTreeDragDest(obj),
		TreeDragSource: WrapTreeDragSource(obj),
		TreeModel:      WrapTreeModel(obj),
		TreeSortable:   WrapTreeSortable(obj),
	}
}

func marshalListStore(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapListStore(obj), nil
}

// Append appends a new row to @list_store. @iter will be changed to point
// to this new row. The row will be empty after this function is called. To
// fill in values, you need to call gtk_list_store_set() or
// gtk_list_store_set_value().
func (l listStore) Append() TreeIter {
	var arg0 *C.GtkListStore

	arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))

	var iter TreeIter

	C.gtk_list_store_append(arg0, (*C.GtkTreeIter)(unsafe.Pointer(&iter)))

	return iter
}

// Clear removes all rows from the list store.
func (l listStore) Clear() {
	var arg0 *C.GtkListStore

	arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))

	C.gtk_list_store_clear(arg0)
}

// Insert creates a new row at @position. @iter will be changed to point to
// this new row. If @position is -1 or is larger than the number of rows on
// the list, then the new row will be appended to the list. The row will be
// empty after this function is called. To fill in values, you need to call
// gtk_list_store_set() or gtk_list_store_set_value().
func (l listStore) Insert(position int) TreeIter {
	var arg0 *C.GtkListStore
	var arg2 C.gint

	arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))
	arg2 = C.gint(position)

	var iter TreeIter

	C.gtk_list_store_insert(arg0, arg2, (*C.GtkTreeIter)(unsafe.Pointer(&iter)))

	return iter
}

// InsertAfter inserts a new row after @sibling. If @sibling is nil, then
// the row will be prepended to the beginning of the list. @iter will be
// changed to point to this new row. The row will be empty after this
// function is called. To fill in values, you need to call
// gtk_list_store_set() or gtk_list_store_set_value().
func (l listStore) InsertAfter(sibling *TreeIter) TreeIter {
	var arg0 *C.GtkListStore
	var arg2 *C.GtkTreeIter

	arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))
	arg2 = (*C.GtkTreeIter)(unsafe.Pointer(sibling.Native()))

	var iter TreeIter

	C.gtk_list_store_insert_after(arg0, arg2, (*C.GtkTreeIter)(unsafe.Pointer(&iter)))

	return iter
}

// InsertBefore inserts a new row before @sibling. If @sibling is nil, then
// the row will be appended to the end of the list. @iter will be changed to
// point to this new row. The row will be empty after this function is
// called. To fill in values, you need to call gtk_list_store_set() or
// gtk_list_store_set_value().
func (l listStore) InsertBefore(sibling *TreeIter) TreeIter {
	var arg0 *C.GtkListStore
	var arg2 *C.GtkTreeIter

	arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))
	arg2 = (*C.GtkTreeIter)(unsafe.Pointer(sibling.Native()))

	var iter TreeIter

	C.gtk_list_store_insert_before(arg0, arg2, (*C.GtkTreeIter)(unsafe.Pointer(&iter)))

	return iter
}

// InsertWithValuesv: a variant of gtk_list_store_insert_with_values() which
// takes the columns and values as two arrays, instead of varargs. This
// function is mainly intended for language-bindings.
func (l listStore) InsertWithValuesv() TreeIter {
	var arg0 *C.GtkListStore

	arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))

	var iter TreeIter

	C.gtk_list_store_insert_with_valuesv(arg0, (*C.GtkTreeIter)(unsafe.Pointer(&iter)))

	return iter
}

// IterIsValid: > This function is slow. Only use it for debugging and/or
// testing > purposes.
//
// Checks if the given iter is a valid iter for this ListStore.
func (l listStore) IterIsValid(iter *TreeIter) bool {
	var arg0 *C.GtkListStore
	var arg1 *C.GtkTreeIter

	arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))
	arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))

	var cret C.gboolean

	cret = C.gtk_list_store_iter_is_valid(arg0, arg1)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// MoveAfter moves @iter in @store to the position after @position. Note
// that this function only works with unsorted stores. If @position is nil,
// @iter will be moved to the start of the list.
func (s listStore) MoveAfter(iter *TreeIter, position *TreeIter) {
	var arg0 *C.GtkListStore
	var arg1 *C.GtkTreeIter
	var arg2 *C.GtkTreeIter

	arg0 = (*C.GtkListStore)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))
	arg2 = (*C.GtkTreeIter)(unsafe.Pointer(position.Native()))

	C.gtk_list_store_move_after(arg0, arg1, arg2)
}

// MoveBefore moves @iter in @store to the position before @position. Note
// that this function only works with unsorted stores. If @position is nil,
// @iter will be moved to the end of the list.
func (s listStore) MoveBefore(iter *TreeIter, position *TreeIter) {
	var arg0 *C.GtkListStore
	var arg1 *C.GtkTreeIter
	var arg2 *C.GtkTreeIter

	arg0 = (*C.GtkListStore)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))
	arg2 = (*C.GtkTreeIter)(unsafe.Pointer(position.Native()))

	C.gtk_list_store_move_before(arg0, arg1, arg2)
}

// Prepend prepends a new row to @list_store. @iter will be changed to point
// to this new row. The row will be empty after this function is called. To
// fill in values, you need to call gtk_list_store_set() or
// gtk_list_store_set_value().
func (l listStore) Prepend() TreeIter {
	var arg0 *C.GtkListStore

	arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))

	var iter TreeIter

	C.gtk_list_store_prepend(arg0, (*C.GtkTreeIter)(unsafe.Pointer(&iter)))

	return iter
}

// Remove removes the given row from the list store. After being removed,
// @iter is set to be the next valid row, or invalidated if it pointed to
// the last row in @list_store.
func (l listStore) Remove(iter *TreeIter) bool {
	var arg0 *C.GtkListStore
	var arg1 *C.GtkTreeIter

	arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))
	arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))

	var cret C.gboolean

	cret = C.gtk_list_store_remove(arg0, arg1)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// Reorder reorders @store to follow the order indicated by @new_order. Note
// that this function only works with unsorted stores.
func (s listStore) Reorder(newOrder []int) {
	var arg0 *C.GtkListStore
	var arg1 *C.gint

	arg0 = (*C.GtkListStore)(unsafe.Pointer(s.Native()))
	arg1 = (*C.gint)(C.malloc((len(newOrder) + 1) * C.sizeof_gint))
	defer C.free(unsafe.Pointer(arg1))

	{
		var out []C.gint
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(arg1), int(len(newOrder)))

		for i := range newOrder {
			arg1 = C.gint(newOrder)
		}
	}

	C.gtk_list_store_reorder(arg0, arg1)
}

// SetValue sets the data in the cell specified by @iter and @column. The
// type of @value must be convertible to the type of the column.
func (l listStore) SetValue(iter *TreeIter, column int, value **externglib.Value) {
	var arg0 *C.GtkListStore
	var arg1 *C.GtkTreeIter
	var arg2 C.gint
	var arg3 *C.GValue

	arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))
	arg1 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))
	arg2 = C.gint(column)
	arg3 = (*C.GValue)(value.GValue)

	C.gtk_list_store_set_value(arg0, arg1, arg2, arg3)
}

// SetValuesv: a variant of gtk_list_store_set_valist() which takes the
// columns and values as two arrays, instead of varargs. This function is
// mainly intended for language-bindings and in case the number of columns
// to change is not known until run-time.
func (l listStore) SetValuesv() {
	var arg0 *C.GtkListStore

	arg0 = (*C.GtkListStore)(unsafe.Pointer(l.Native()))

	C.gtk_list_store_set_valuesv(arg0)
}

// Swap swaps @a and @b in @store. Note that this function only works with
// unsorted stores.
func (s listStore) Swap(a *TreeIter, b *TreeIter) {
	var arg0 *C.GtkListStore
	var arg1 *C.GtkTreeIter
	var arg2 *C.GtkTreeIter

	arg0 = (*C.GtkListStore)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkTreeIter)(unsafe.Pointer(a.Native()))
	arg2 = (*C.GtkTreeIter)(unsafe.Pointer(b.Native()))

	C.gtk_list_store_swap(arg0, arg1, arg2)
}
