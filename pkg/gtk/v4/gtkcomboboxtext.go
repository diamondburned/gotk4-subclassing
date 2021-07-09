// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_combo_box_text_get_type()), F: marshalComboBoxText},
	})
}

// ComboBoxText: `GtkComboBoxText` is a simple variant of `GtkComboBox` for
// text-only use cases.
//
// !An example GtkComboBoxText (combo-box-text.png)
//
// `GtkComboBoxText` hides the model-view complexity of `GtkComboBox`.
//
// To create a `GtkComboBoxText`, use [ctor@Gtk.ComboBoxText.new] or
// [ctor@Gtk.ComboBoxText.new_with_entry].
//
// You can add items to a `GtkComboBoxText` with
// [method@Gtk.ComboBoxText.append_text], [method@Gtk.ComboBoxText.insert_text]
// or [method@Gtk.ComboBoxText.prepend_text] and remove options with
// [method@Gtk.ComboBoxText.remove].
//
// If the `GtkComboBoxText` contains an entry (via the
// [property@Gtk.ComboBox:has-entry] property), its contents can be retrieved
// using [method@Gtk.ComboBoxText.get_active_text].
//
// You should not call [method@Gtk.ComboBox.set_model] or attempt to pack more
// cells into this combo box via its [interface@Gtk.CellLayout] interface.
//
//
// GtkComboBoxText as GtkBuildable
//
// The `GtkComboBoxText` implementation of the `GtkBuildable` interface supports
// adding items directly using the <items> element and specifying <item>
// elements for each item. Each <item> element can specify the “id”
// corresponding to the appended text and also supports the regular translation
// attributes “translatable”, “context” and “comments”.
//
// Here is a UI definition fragment specifying `GtkComboBoxText` items: “`xml
// <object class="GtkComboBoxText"> <items> <item translatable="yes"
// id="factory">Factory</item> <item translatable="yes" id="home">Home</item>
// <item translatable="yes" id="subway">Subway</item> </items> </object> “`
//
//
// CSS nodes
//
// “` combobox ╰── box.linked ├── entry.combo ├── button.combo ╰── window.popup
// “`
//
// `GtkComboBoxText` has a single CSS node with name combobox. It adds the style
// class .combo to the main CSS nodes of its entry and button children, and the
// .linked class to the node of its internal box.
type ComboBoxText interface {
	gextras.Objector

	// Append appends @text to the list of strings stored in @combo_box.
	//
	// If @id is non-nil then it is used as the ID of the row.
	//
	// This is the same as calling [method@Gtk.ComboBoxText.insert] with a
	// position of -1.
	Append(id string, text string)
	// AppendText appends @text to the list of strings stored in @combo_box.
	//
	// This is the same as calling [method@Gtk.ComboBoxText.insert_text] with a
	// position of -1.
	AppendText(text string)
	// ActiveText returns the currently active string in @combo_box.
	//
	// If no row is currently selected, nil is returned. If @combo_box contains
	// an entry, this function will return its contents (which will not
	// necessarily be an item from the list).
	ActiveText() string
	// Insert inserts @text at @position in the list of strings stored in
	// @combo_box.
	//
	// If @id is non-nil then it is used as the ID of the row. See
	// [property@Gtk.ComboBox:id-column].
	//
	// If @position is negative then @text is appended.
	Insert(position int, id string, text string)
	// InsertText inserts @text at @position in the list of strings stored in
	// @combo_box.
	//
	// If @position is negative then @text is appended.
	//
	// This is the same as calling [method@Gtk.ComboBoxText.insert] with a nil
	// ID string.
	InsertText(position int, text string)
	// Prepend prepends @text to the list of strings stored in @combo_box.
	//
	// If @id is non-nil then it is used as the ID of the row.
	//
	// This is the same as calling [method@Gtk.ComboBoxText.insert] with a
	// position of 0.
	Prepend(id string, text string)
	// PrependText prepends @text to the list of strings stored in @combo_box.
	//
	// This is the same as calling [method@Gtk.ComboBoxText.insert_text] with a
	// position of 0.
	PrependText(text string)
	// Remove removes the string at @position from @combo_box.
	Remove(position int)
	// RemoveAll removes all the text entries from the combo box.
	RemoveAll()
}

// ComboBoxTextClass implements the ComboBoxText interface.
type ComboBoxTextClass struct {
	*externglib.Object
	ComboBoxClass
	AccessibleInterface
	BuildableInterface
	CellEditableInterface
	CellLayoutInterface
	ConstraintTargetInterface
}

var _ ComboBoxText = (*ComboBoxTextClass)(nil)

func wrapComboBoxText(obj *externglib.Object) ComboBoxText {
	return &ComboBoxTextClass{
		Object: obj,
		ComboBoxClass: ComboBoxClass{
			Object: obj,
			WidgetClass: WidgetClass{
				InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
				AccessibleInterface: AccessibleInterface{
					Object: obj,
				},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
				ConstraintTargetInterface: ConstraintTargetInterface{
					Object: obj,
				},
			},
			AccessibleInterface: AccessibleInterface{
				Object: obj,
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
			CellEditableInterface: CellEditableInterface{
				WidgetClass: WidgetClass{
					InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
					AccessibleInterface: AccessibleInterface{
						Object: obj,
					},
					BuildableInterface: BuildableInterface{
						Object: obj,
					},
					ConstraintTargetInterface: ConstraintTargetInterface{
						Object: obj,
					},
				},
			},
			CellLayoutInterface: CellLayoutInterface{
				Object: obj,
			},
			ConstraintTargetInterface: ConstraintTargetInterface{
				Object: obj,
			},
		},
		AccessibleInterface: AccessibleInterface{
			Object: obj,
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
		CellEditableInterface: CellEditableInterface{
			WidgetClass: WidgetClass{
				InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
				AccessibleInterface: AccessibleInterface{
					Object: obj,
				},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
				ConstraintTargetInterface: ConstraintTargetInterface{
					Object: obj,
				},
			},
		},
		CellLayoutInterface: CellLayoutInterface{
			Object: obj,
		},
		ConstraintTargetInterface: ConstraintTargetInterface{
			Object: obj,
		},
	}
}

func marshalComboBoxText(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapComboBoxText(obj), nil
}

// NewComboBoxText creates a new `GtkComboBoxText`.
func NewComboBoxText() *ComboBoxTextClass {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_combo_box_text_new()

	var _comboBoxText *ComboBoxTextClass // out

	_comboBoxText = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ComboBoxTextClass)

	return _comboBoxText
}

// NewComboBoxTextWithEntry creates a new `GtkComboBoxText` with an entry.
func NewComboBoxTextWithEntry() *ComboBoxTextClass {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_combo_box_text_new_with_entry()

	var _comboBoxText *ComboBoxTextClass // out

	_comboBoxText = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*ComboBoxTextClass)

	return _comboBoxText
}

// Append appends @text to the list of strings stored in @combo_box.
//
// If @id is non-nil then it is used as the ID of the row.
//
// This is the same as calling [method@Gtk.ComboBoxText.insert] with a position
// of -1.
func (c *ComboBoxTextClass) Append(id string, text string) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 *C.char            // out
	var _arg2 *C.char            // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(id))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_combo_box_text_append(_arg0, _arg1, _arg2)
}

// AppendText appends @text to the list of strings stored in @combo_box.
//
// This is the same as calling [method@Gtk.ComboBoxText.insert_text] with a
// position of -1.
func (c *ComboBoxTextClass) AppendText(text string) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_combo_box_text_append_text(_arg0, _arg1)
}

// ActiveText returns the currently active string in @combo_box.
//
// If no row is currently selected, nil is returned. If @combo_box contains an
// entry, this function will return its contents (which will not necessarily be
// an item from the list).
func (c *ComboBoxTextClass) ActiveText() string {
	var _arg0 *C.GtkComboBoxText // out
	var _cret *C.char            // in

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_combo_box_text_get_active_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Insert inserts @text at @position in the list of strings stored in
// @combo_box.
//
// If @id is non-nil then it is used as the ID of the row. See
// [property@Gtk.ComboBox:id-column].
//
// If @position is negative then @text is appended.
func (c *ComboBoxTextClass) Insert(position int, id string, text string) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 C.int              // out
	var _arg2 *C.char            // out
	var _arg3 *C.char            // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(c.Native()))
	_arg1 = C.int(position)
	_arg2 = (*C.char)(C.CString(id))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg3))

	C.gtk_combo_box_text_insert(_arg0, _arg1, _arg2, _arg3)
}

// InsertText inserts @text at @position in the list of strings stored in
// @combo_box.
//
// If @position is negative then @text is appended.
//
// This is the same as calling [method@Gtk.ComboBoxText.insert] with a nil ID
// string.
func (c *ComboBoxTextClass) InsertText(position int, text string) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 C.int              // out
	var _arg2 *C.char            // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(c.Native()))
	_arg1 = C.int(position)
	_arg2 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_combo_box_text_insert_text(_arg0, _arg1, _arg2)
}

// Prepend prepends @text to the list of strings stored in @combo_box.
//
// If @id is non-nil then it is used as the ID of the row.
//
// This is the same as calling [method@Gtk.ComboBoxText.insert] with a position
// of 0.
func (c *ComboBoxTextClass) Prepend(id string, text string) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 *C.char            // out
	var _arg2 *C.char            // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(id))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_combo_box_text_prepend(_arg0, _arg1, _arg2)
}

// PrependText prepends @text to the list of strings stored in @combo_box.
//
// This is the same as calling [method@Gtk.ComboBoxText.insert_text] with a
// position of 0.
func (c *ComboBoxTextClass) PrependText(text string) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_combo_box_text_prepend_text(_arg0, _arg1)
}

// Remove removes the string at @position from @combo_box.
func (c *ComboBoxTextClass) Remove(position int) {
	var _arg0 *C.GtkComboBoxText // out
	var _arg1 C.int              // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(c.Native()))
	_arg1 = C.int(position)

	C.gtk_combo_box_text_remove(_arg0, _arg1)
}

// RemoveAll removes all the text entries from the combo box.
func (c *ComboBoxTextClass) RemoveAll() {
	var _arg0 *C.GtkComboBoxText // out

	_arg0 = (*C.GtkComboBoxText)(unsafe.Pointer(c.Native()))

	C.gtk_combo_box_text_remove_all(_arg0)
}
