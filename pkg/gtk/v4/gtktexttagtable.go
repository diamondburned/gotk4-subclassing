// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
// void _gotk4_gtk4_TextTagTableForeach(GtkTextTag*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_text_tag_table_get_type()), F: marshalTextTagTabler},
	})
}

// TextTagTableForeach: function used with gtk_text_tag_table_foreach(), to
// iterate over every GtkTextTag inside a GtkTextTagTable.
type TextTagTableForeach func(tag *TextTag)

//export _gotk4_gtk4_TextTagTableForeach
func _gotk4_gtk4_TextTagTableForeach(arg0 *C.GtkTextTag, arg1 C.gpointer) {
	v := gbox.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var tag *TextTag // out

	tag = wrapTextTag(externglib.Take(unsafe.Pointer(arg0)))

	fn := v.(TextTagTableForeach)
	fn(tag)
}

// TextTagTable: collection of tags in a GtkTextBuffer
//
// You may wish to begin by reading the text widget conceptual overview
// (section-text-widget.html), which gives an overview of all the objects and
// data types related to the text widget and how they work together.
//
//
// GtkTextTagTables as GtkBuildable
//
// The GtkTextTagTable implementation of the GtkBuildable interface supports
// adding tags by specifying “tag” as the “type” attribute of a <child> element.
//
// An example of a UI definition fragment specifying tags:
//
//    <object class="GtkTextTagTable">
//     <child type="tag">
//       <object class="GtkTextTag"/>
//     </child>
//    </object>.
type TextTagTable struct {
	*externglib.Object

	Buildable
}

func wrapTextTagTable(obj *externglib.Object) *TextTagTable {
	return &TextTagTable{
		Object: obj,
		Buildable: Buildable{
			Object: obj,
		},
	}
}

func marshalTextTagTabler(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTextTagTable(obj), nil
}

// NewTextTagTable creates a new GtkTextTagTable.
//
// The table contains no tags by default.
func NewTextTagTable() *TextTagTable {
	var _cret *C.GtkTextTagTable // in

	_cret = C.gtk_text_tag_table_new()

	var _textTagTable *TextTagTable // out

	_textTagTable = wrapTextTagTable(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _textTagTable
}

// Add a tag to the table.
//
// The tag is assigned the highest priority in the table.
//
// tag must not be in a tag table already, and may not have the same name as an
// already-added tag.
func (table *TextTagTable) Add(tag *TextTag) bool {
	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.GtkTextTag      // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(table.Native()))
	_arg1 = (*C.GtkTextTag)(unsafe.Pointer(tag.Native()))

	_cret = C.gtk_text_tag_table_add(_arg0, _arg1)
	runtime.KeepAlive(table)
	runtime.KeepAlive(tag)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Foreach calls func on each tag in table, with user data data.
//
// Note that the table may not be modified while iterating over it (you can’t
// add/remove tags).
func (table *TextTagTable) Foreach(fn TextTagTableForeach) {
	var _arg0 *C.GtkTextTagTable       // out
	var _arg1 C.GtkTextTagTableForeach // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(table.Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk4_TextTagTableForeach)
	_arg2 = C.gpointer(gbox.Assign(fn))
	defer gbox.Delete(uintptr(_arg2))

	C.gtk_text_tag_table_foreach(_arg0, _arg1, _arg2)
	runtime.KeepAlive(table)
	runtime.KeepAlive(fn)
}

// Size returns the size of the table (number of tags).
func (table *TextTagTable) Size() int {
	var _arg0 *C.GtkTextTagTable // out
	var _cret C.int              // in

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(table.Native()))

	_cret = C.gtk_text_tag_table_get_size(_arg0)
	runtime.KeepAlive(table)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Lookup: look up a named tag.
func (table *TextTagTable) Lookup(name string) *TextTag {
	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.char            // out
	var _cret *C.GtkTextTag      // in

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(table.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_text_tag_table_lookup(_arg0, _arg1)
	runtime.KeepAlive(table)
	runtime.KeepAlive(name)

	var _textTag *TextTag // out

	if _cret != nil {
		_textTag = wrapTextTag(externglib.Take(unsafe.Pointer(_cret)))
	}

	return _textTag
}

// Remove a tag from the table.
//
// If a GtkTextBuffer has table as its tag table, the tag is removed from the
// buffer. The table’s reference to the tag is removed, so the tag will end up
// destroyed if you don’t have a reference to it.
func (table *TextTagTable) Remove(tag *TextTag) {
	var _arg0 *C.GtkTextTagTable // out
	var _arg1 *C.GtkTextTag      // out

	_arg0 = (*C.GtkTextTagTable)(unsafe.Pointer(table.Native()))
	_arg1 = (*C.GtkTextTag)(unsafe.Pointer(tag.Native()))

	C.gtk_text_tag_table_remove(_arg0, _arg1)
	runtime.KeepAlive(table)
	runtime.KeepAlive(tag)
}
