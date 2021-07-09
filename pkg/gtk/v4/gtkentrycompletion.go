// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_entry_completion_get_type()), F: marshalEntryCompletion},
	})
}

// EntryCompletionMatchFunc: function which decides whether the row indicated by
// @iter matches a given @key, and should be displayed as a possible completion
// for @key.
//
// Note that @key is normalized and case-folded (see g_utf8_normalize() and
// g_utf8_casefold()). If this is not appropriate, match functions have access
// to the unmodified key via `gtk_editable_get_text (GTK_EDITABLE
// (gtk_entry_completion_get_entry ()))`.
type EntryCompletionMatchFunc func(completion *EntryCompletionClass, key string, iter *TreeIter, userData interface{}) (ok bool)

//export gotk4_EntryCompletionMatchFunc
func gotk4_EntryCompletionMatchFunc(arg0 *C.GtkEntryCompletion, arg1 *C.char, arg2 *C.GtkTreeIter, arg3 C.gpointer) (cret C.gboolean) {
	v := box.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var completion *EntryCompletionClass // out
	var key string                       // out
	var iter *TreeIter                   // out
	var userData interface{}             // out

	completion = (gextras.CastObject(externglib.Take(unsafe.Pointer(arg0)))).(*EntryCompletionClass)
	key = C.GoString(arg1)
	iter = (*TreeIter)(unsafe.Pointer(arg2))
	userData = box.Get(uintptr(arg3))

	fn := v.(EntryCompletionMatchFunc)
	ok := fn(completion, key, iter, userData)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// EntryCompletion: `GtkEntryCompletion` is an auxiliary object to provide
// completion functionality for `GtkEntry`.
//
// It implements the [iface@Gtk.CellLayout] interface, to allow the user to add
// extra cells to the `GtkTreeView` with completion matches.
//
// “Completion functionality” means that when the user modifies the text in the
// entry, `GtkEntryCompletion` checks which rows in the model match the current
// content of the entry, and displays a list of matches. By default, the
// matching is done by comparing the entry text case-insensitively against the
// text column of the model (see [method@Gtk.EntryCompletion.set_text_column]),
// but this can be overridden with a custom match function (see
// [method@Gtk.EntryCompletion.set_match_func]).
//
// When the user selects a completion, the content of the entry is updated. By
// default, the content of the entry is replaced by the text column of the
// model, but this can be overridden by connecting to the
// [signal@Gtk.EntryCompletion::match-selected] signal and updating the entry in
// the signal handler. Note that you should return true from the signal handler
// to suppress the default behaviour.
//
// To add completion functionality to an entry, use
// [method@Gtk.Entry.set_completion].
//
// `GtkEntryCompletion` uses a [class@Gtk.TreeModelFilter] model to represent
// the subset of the entire model that is currently matching. While the
// `GtkEntryCompletion` signals [signal@Gtk.EntryCompletion::match-selected] and
// [signal@Gtk.EntryCompletion::cursor-on-match] take the original model and an
// iter pointing to that model as arguments, other callbacks and signals (such
// as `GtkCellLayoutDataFunc` or [signal@Gtk.CellArea::apply-attributes)] will
// generally take the filter model as argument. As long as you are only calling
// [method@Gtk.TreeModel.get], this will make no difference to you. If for some
// reason, you need the original model, use
// [method@Gtk.TreeModelFilter.get_model]. Don’t forget to use
// [method@Gtk.TreeModelFilter.convert_iter_to_child_iter] to obtain a matching
// iter.
type EntryCompletion interface {
	gextras.Objector

	// Complete requests a completion operation, or in other words a refiltering
	// of the current list with completions, using the current key.
	//
	// The completion list view will be updated accordingly.
	Complete()
	// ComputePrefix computes the common prefix that is shared by all rows in
	// @completion that start with @key.
	//
	// If no row matches @key, nil will be returned. Note that a text column
	// must have been set for this function to work, see
	// [method@Gtk.EntryCompletion.set_text_column] for details.
	ComputePrefix(key string) string
	// CompletionPrefix: get the original text entered by the user that
	// triggered the completion or nil if there’s no completion ongoing.
	CompletionPrefix() string
	// Entry gets the entry @completion has been attached to.
	Entry() *WidgetClass
	// InlineCompletion returns whether the common prefix of the possible
	// completions should be automatically inserted in the entry.
	InlineCompletion() bool
	// InlineSelection returns true if inline-selection mode is turned on.
	InlineSelection() bool
	// MinimumKeyLength returns the minimum key length as set for @completion.
	MinimumKeyLength() int
	// Model returns the model the `GtkEntryCompletion` is using as data source.
	//
	// Returns nil if the model is unset.
	Model() *TreeModelInterface
	// PopupCompletion returns whether the completions should be presented in a
	// popup window.
	PopupCompletion() bool
	// PopupSetWidth returns whether the completion popup window will be resized
	// to the width of the entry.
	PopupSetWidth() bool
	// PopupSingleMatch returns whether the completion popup window will appear
	// even if there is only a single match.
	PopupSingleMatch() bool
	// TextColumn returns the column in the model of @completion to get strings
	// from.
	TextColumn() int
	// InsertPrefix requests a prefix insertion.
	InsertPrefix()
	// SetInlineCompletion sets whether the common prefix of the possible
	// completions should be automatically inserted in the entry.
	SetInlineCompletion(inlineCompletion bool)
	// SetInlineSelection sets whether it is possible to cycle through the
	// possible completions inside the entry.
	SetInlineSelection(inlineSelection bool)
	// SetMinimumKeyLength requires the length of the search key for @completion
	// to be at least @length.
	//
	// This is useful for long lists, where completing using a small key takes a
	// lot of time and will come up with meaningless results anyway (ie, a too
	// large dataset).
	SetMinimumKeyLength(length int)
	// SetModel sets the model for a `GtkEntryCompletion`.
	//
	// If @completion already has a model set, it will remove it before setting
	// the new model. If model is nil, then it will unset the model.
	SetModel(model TreeModel)
	// SetPopupCompletion sets whether the completions should be presented in a
	// popup window.
	SetPopupCompletion(popupCompletion bool)
	// SetPopupSetWidth sets whether the completion popup window will be resized
	// to be the same width as the entry.
	SetPopupSetWidth(popupSetWidth bool)
	// SetPopupSingleMatch sets whether the completion popup window will appear
	// even if there is only a single match.
	//
	// You may want to set this to false if you are using
	// [property@Gtk.EntryCompletion:inline-completion].
	SetPopupSingleMatch(popupSingleMatch bool)
	// SetTextColumn: convenience function for setting up the most used case of
	// this code: a completion list with just strings.
	//
	// This function will set up @completion to have a list displaying all (and
	// just) strings in the completion list, and to get those strings from
	// @column in the model of @completion.
	//
	// This functions creates and adds a `GtkCellRendererText` for the selected
	// column. If you need to set the text column, but don't want the cell
	// renderer, use g_object_set() to set the
	// [property@Gtk.EntryCompletion:text-column] property directly.
	SetTextColumn(column int)
}

// EntryCompletionClass implements the EntryCompletion interface.
type EntryCompletionClass struct {
	*externglib.Object
	BuildableInterface
	CellLayoutInterface
}

var _ EntryCompletion = (*EntryCompletionClass)(nil)

func wrapEntryCompletion(obj *externglib.Object) EntryCompletion {
	return &EntryCompletionClass{
		Object: obj,
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		CellLayoutInterface: CellLayoutInterface{
			Object: obj,
		},
	}
}

func marshalEntryCompletion(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapEntryCompletion(obj), nil
}

// NewEntryCompletion creates a new `GtkEntryCompletion` object.
func NewEntryCompletion() *EntryCompletionClass {
	var _cret *C.GtkEntryCompletion // in

	_cret = C.gtk_entry_completion_new()

	var _entryCompletion *EntryCompletionClass // out

	_entryCompletion = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*EntryCompletionClass)

	return _entryCompletion
}

// NewEntryCompletionWithArea creates a new `GtkEntryCompletion` object using
// the specified @area.
//
// The `GtkCellArea` is used to layout cells in the underlying
// `GtkTreeViewColumn` for the drop-down menu.
func NewEntryCompletionWithArea(area CellArea) *EntryCompletionClass {
	var _arg1 *C.GtkCellArea        // out
	var _cret *C.GtkEntryCompletion // in

	_arg1 = (*C.GtkCellArea)(unsafe.Pointer(area.Native()))

	_cret = C.gtk_entry_completion_new_with_area(_arg1)

	var _entryCompletion *EntryCompletionClass // out

	_entryCompletion = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*EntryCompletionClass)

	return _entryCompletion
}

// Complete requests a completion operation, or in other words a refiltering of
// the current list with completions, using the current key.
//
// The completion list view will be updated accordingly.
func (c *EntryCompletionClass) Complete() {
	var _arg0 *C.GtkEntryCompletion // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	C.gtk_entry_completion_complete(_arg0)
}

// ComputePrefix computes the common prefix that is shared by all rows in
// @completion that start with @key.
//
// If no row matches @key, nil will be returned. Note that a text column must
// have been set for this function to work, see
// [method@Gtk.EntryCompletion.set_text_column] for details.
func (c *EntryCompletionClass) ComputePrefix(key string) string {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 *C.char               // out
	var _cret *C.char               // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_entry_completion_compute_prefix(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// CompletionPrefix: get the original text entered by the user that triggered
// the completion or nil if there’s no completion ongoing.
func (c *EntryCompletionClass) CompletionPrefix() string {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret *C.char               // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_completion_prefix(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Entry gets the entry @completion has been attached to.
func (c *EntryCompletionClass) Entry() *WidgetClass {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret *C.GtkWidget          // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_entry(_arg0)

	var _widget *WidgetClass // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*WidgetClass)

	return _widget
}

// InlineCompletion returns whether the common prefix of the possible
// completions should be automatically inserted in the entry.
func (c *EntryCompletionClass) InlineCompletion() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_inline_completion(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InlineSelection returns true if inline-selection mode is turned on.
func (c *EntryCompletionClass) InlineSelection() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_inline_selection(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MinimumKeyLength returns the minimum key length as set for @completion.
func (c *EntryCompletionClass) MinimumKeyLength() int {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.int                 // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_minimum_key_length(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Model returns the model the `GtkEntryCompletion` is using as data source.
//
// Returns nil if the model is unset.
func (c *EntryCompletionClass) Model() *TreeModelInterface {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret *C.GtkTreeModel       // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_model(_arg0)

	var _treeModel *TreeModelInterface // out

	_treeModel = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*TreeModelInterface)

	return _treeModel
}

// PopupCompletion returns whether the completions should be presented in a
// popup window.
func (c *EntryCompletionClass) PopupCompletion() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_popup_completion(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PopupSetWidth returns whether the completion popup window will be resized to
// the width of the entry.
func (c *EntryCompletionClass) PopupSetWidth() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_popup_set_width(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PopupSingleMatch returns whether the completion popup window will appear even
// if there is only a single match.
func (c *EntryCompletionClass) PopupSingleMatch() bool {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_popup_single_match(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TextColumn returns the column in the model of @completion to get strings
// from.
func (c *EntryCompletionClass) TextColumn() int {
	var _arg0 *C.GtkEntryCompletion // out
	var _cret C.int                 // in

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_entry_completion_get_text_column(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// InsertPrefix requests a prefix insertion.
func (c *EntryCompletionClass) InsertPrefix() {
	var _arg0 *C.GtkEntryCompletion // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))

	C.gtk_entry_completion_insert_prefix(_arg0)
}

// SetInlineCompletion sets whether the common prefix of the possible
// completions should be automatically inserted in the entry.
func (c *EntryCompletionClass) SetInlineCompletion(inlineCompletion bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	if inlineCompletion {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_inline_completion(_arg0, _arg1)
}

// SetInlineSelection sets whether it is possible to cycle through the possible
// completions inside the entry.
func (c *EntryCompletionClass) SetInlineSelection(inlineSelection bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	if inlineSelection {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_inline_selection(_arg0, _arg1)
}

// SetMinimumKeyLength requires the length of the search key for @completion to
// be at least @length.
//
// This is useful for long lists, where completing using a small key takes a lot
// of time and will come up with meaningless results anyway (ie, a too large
// dataset).
func (c *EntryCompletionClass) SetMinimumKeyLength(length int) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.int                 // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	_arg1 = C.int(length)

	C.gtk_entry_completion_set_minimum_key_length(_arg0, _arg1)
}

// SetModel sets the model for a `GtkEntryCompletion`.
//
// If @completion already has a model set, it will remove it before setting the
// new model. If model is nil, then it will unset the model.
func (c *EntryCompletionClass) SetModel(model TreeModel) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 *C.GtkTreeModel       // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(model.Native()))

	C.gtk_entry_completion_set_model(_arg0, _arg1)
}

// SetPopupCompletion sets whether the completions should be presented in a
// popup window.
func (c *EntryCompletionClass) SetPopupCompletion(popupCompletion bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	if popupCompletion {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_popup_completion(_arg0, _arg1)
}

// SetPopupSetWidth sets whether the completion popup window will be resized to
// be the same width as the entry.
func (c *EntryCompletionClass) SetPopupSetWidth(popupSetWidth bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	if popupSetWidth {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_popup_set_width(_arg0, _arg1)
}

// SetPopupSingleMatch sets whether the completion popup window will appear even
// if there is only a single match.
//
// You may want to set this to false if you are using
// [property@Gtk.EntryCompletion:inline-completion].
func (c *EntryCompletionClass) SetPopupSingleMatch(popupSingleMatch bool) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	if popupSingleMatch {
		_arg1 = C.TRUE
	}

	C.gtk_entry_completion_set_popup_single_match(_arg0, _arg1)
}

// SetTextColumn: convenience function for setting up the most used case of this
// code: a completion list with just strings.
//
// This function will set up @completion to have a list displaying all (and
// just) strings in the completion list, and to get those strings from @column
// in the model of @completion.
//
// This functions creates and adds a `GtkCellRendererText` for the selected
// column. If you need to set the text column, but don't want the cell renderer,
// use g_object_set() to set the [property@Gtk.EntryCompletion:text-column]
// property directly.
func (c *EntryCompletionClass) SetTextColumn(column int) {
	var _arg0 *C.GtkEntryCompletion // out
	var _arg1 C.int                 // out

	_arg0 = (*C.GtkEntryCompletion)(unsafe.Pointer(c.Native()))
	_arg1 = C.int(column)

	C.gtk_entry_completion_set_text_column(_arg0, _arg1)
}
