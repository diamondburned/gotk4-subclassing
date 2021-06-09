// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

type Stock string

// TranslateFunc: the function used to translate messages in e.g. IconFactory
// and ActionGroup.
type TranslateFunc func() (utf8 string)

//export gotk4_TranslateFunc
func gotk4_TranslateFunc(arg0 *C.gchar, arg1 C.gpointer) *C.gchar {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(TranslateFunc)
	utf8 := fn()

	cret = (*C.gchar)(C.CString(utf8))
}

// StockAdd registers each of the stock items in @items. If an item already
// exists with the same stock ID as one of the @items, the old item gets
// replaced. The stock items are copied, so GTK+ does not hold any pointer into
// @items and @items can be freed. Use gtk_stock_add_static() if @items is
// persistent and GTK+ need not copy the array.
func StockAdd() {
	C.gtk_stock_add()
}

// StockAddStatic: same as gtk_stock_add(), but doesn’t copy @items, so @items
// must persist until application exit.
func StockAddStatic() {
	C.gtk_stock_add_static()
}

// StockListIds retrieves a list of all known stock IDs added to a IconFactory
// or registered with gtk_stock_add(). The list must be freed with
// g_slist_free(), and each string in the list must be freed with g_free().
func StockListIds() *glib.SList {
	var cret *C.GSList

	cret = C.gtk_stock_list_ids()

	var sList *glib.SList

	sList = glib.WrapSList(unsafe.Pointer(cret))
	runtime.SetFinalizer(sList, func(v *glib.SList) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return sList
}

// StockLookup fills @item with the registered values for @stock_id, returning
// true if @stock_id was known.
func StockLookup(stockId string) (item StockItem, ok bool) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(arg1))

	var item StockItem
	var cret C.gboolean

	cret = C.gtk_stock_lookup(arg1, (*C.GtkStockItem)(unsafe.Pointer(&item)))

	var ok bool

	if cret {
		ok = true
	}

	return item, ok
}

// StockSetTranslateFunc sets a function to be used for translating the @label
// of a stock item.
//
// If no function is registered for a translation domain, g_dgettext() is used.
//
// The function is used for all stock items whose @translation_domain matches
// @domain. Note that it is possible to use strings different from the actual
// gettext translation domain of your application for this, as long as your
// TranslateFunc uses the correct domain when calling dgettext(). This can be
// useful, e.g. when dealing with message contexts:
//
//    GtkStockItem items[] = {
//     { MY_ITEM1, NC_("odd items", "Item 1"), 0, 0, "odd-item-domain" },
//     { MY_ITEM2, NC_("even items", "Item 2"), 0, 0, "even-item-domain" },
//    };
//
//    gchar *
//    my_translate_func (const gchar *msgid,
//                       gpointer     data)
//    {
//      gchar *msgctxt = data;
//
//      return (gchar*)g_dpgettext2 (GETTEXT_PACKAGE, msgctxt, msgid);
//    }
//
//    ...
//
//    gtk_stock_add (items, G_N_ELEMENTS (items));
//    gtk_stock_set_translate_func ("odd-item-domain", my_translate_func, "odd items");
//    gtk_stock_set_translate_func ("even-item-domain", my_translate_func, "even items");
func StockSetTranslateFunc() {
	C.gtk_stock_set_translate_func()
}

type StockItem struct {
	native C.GtkStockItem
}

// WrapStockItem wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapStockItem(ptr unsafe.Pointer) *StockItem {
	if ptr == nil {
		return nil
	}

	return (*StockItem)(ptr)
}

func marshalStockItem(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapStockItem(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (s *StockItem) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// StockID gets the field inside the struct.
func (s *StockItem) StockID() string {
	var v string
	v = C.GoString(s.native.stock_id)
	return v
}

// Label gets the field inside the struct.
func (s *StockItem) Label() string {
	var v string
	v = C.GoString(s.native.label)
	return v
}

// Modifier gets the field inside the struct.
func (s *StockItem) Modifier() gdk.ModifierType {
	var v gdk.ModifierType
	v = gdk.ModifierType(s.native.modifier)
	return v
}

// Keyval gets the field inside the struct.
func (s *StockItem) Keyval() uint {
	var v uint
	v = (uint)(s.native.keyval)
	return v
}

// TranslationDomain gets the field inside the struct.
func (s *StockItem) TranslationDomain() string {
	var v string
	v = C.GoString(s.native.translation_domain)
	return v
}

// Copy copies a stock item, mostly useful for language bindings and not in
// applications.
func (i *StockItem) Copy() *StockItem {
	var arg0 *C.GtkStockItem

	arg0 = (*C.GtkStockItem)(unsafe.Pointer(i.Native()))

	var cret *C.GtkStockItem

	cret = C.gtk_stock_item_copy(arg0)

	var stockItem *StockItem

	stockItem = WrapStockItem(unsafe.Pointer(cret))

	return stockItem
}

// Free frees a stock item allocated on the heap, such as one returned by
// gtk_stock_item_copy(). Also frees the fields inside the stock item, if they
// are not nil.
func (i *StockItem) Free() {
	var arg0 *C.GtkStockItem

	arg0 = (*C.GtkStockItem)(unsafe.Pointer(i.Native()))

	C.gtk_stock_item_free(arg0)
}
