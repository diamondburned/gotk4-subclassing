// Code generated by girgen. DO NOT EDIT.

package gtk

import (
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
		{T: externglib.Type(C.gtk_text_tag_get_type()), F: marshalTextTag},
	})
}

// TextTag: you may wish to begin by reading the [text widget conceptual
// overview][TextWidget] which gives an overview of all the objects and data
// types related to the text widget and how they work together.
//
// Tags should be in the TextTagTable for a given TextBuffer before using them
// with that buffer.
//
// gtk_text_buffer_create_tag() is the best way to create tags. See “gtk3-demo”
// for numerous examples.
//
// For each property of TextTag, there is a “set” property, e.g. “font-set”
// corresponds to “font”. These “set” properties reflect whether a property has
// been set or not. They are maintained by GTK+ and you should not set them
// independently.
type TextTag interface {
	gextras.Objector

	// Changed emits the TextTagTable::tag-changed signal on the TextTagTable
	// where the tag is included.
	//
	// The signal is already emitted when setting a TextTag property. This
	// function is useful for a TextTag subclass.
	Changed(t TextTag, sizeChanged bool)
	// Priority: get the tag priority.
	Priority(t TextTag)
	// SetPriority sets the priority of a TextTag. Valid priorities start at 0
	// and go to one less than gtk_text_tag_table_get_size(). Each tag in a
	// table has a unique priority; setting the priority of one tag shifts the
	// priorities of all the other tags in the table to maintain a unique
	// priority for each tag. Higher priority tags “win” if two tags both set
	// the same text attribute. When adding a tag to a tag table, it will be
	// assigned the highest priority in the table by default; so normally the
	// precedence of a set of tags is the order in which they were added to the
	// table, or created with gtk_text_buffer_create_tag(), which adds the tag
	// to the buffer’s table automatically.
	SetPriority(t TextTag, priority int)
}

// textTag implements the TextTag interface.
type textTag struct {
	gextras.Objector
}

var _ TextTag = (*textTag)(nil)

// WrapTextTag wraps a GObject to the right type. It is
// primarily used internally.
func WrapTextTag(obj *externglib.Object) TextTag {
	return TextTag{
		Objector: obj,
	}
}

func marshalTextTag(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapTextTag(obj), nil
}

// NewTextTag constructs a class TextTag.
func NewTextTag(name string) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_text_tag_new(arg1)
}

// Changed emits the TextTagTable::tag-changed signal on the TextTagTable
// where the tag is included.
//
// The signal is already emitted when setting a TextTag property. This
// function is useful for a TextTag subclass.
func (t textTag) Changed(t TextTag, sizeChanged bool) {
	var arg0 *C.GtkTextTag
	var arg1 C.gboolean

	arg0 = (*C.GtkTextTag)(unsafe.Pointer(t.Native()))
	if sizeChanged {
		arg1 = C.gboolean(1)
	}

	C.gtk_text_tag_changed(arg0, arg1)
}

// Priority: get the tag priority.
func (t textTag) Priority(t TextTag) {
	var arg0 *C.GtkTextTag

	arg0 = (*C.GtkTextTag)(unsafe.Pointer(t.Native()))

	C.gtk_text_tag_get_priority(arg0)
}

// SetPriority sets the priority of a TextTag. Valid priorities start at 0
// and go to one less than gtk_text_tag_table_get_size(). Each tag in a
// table has a unique priority; setting the priority of one tag shifts the
// priorities of all the other tags in the table to maintain a unique
// priority for each tag. Higher priority tags “win” if two tags both set
// the same text attribute. When adding a tag to a tag table, it will be
// assigned the highest priority in the table by default; so normally the
// precedence of a set of tags is the order in which they were added to the
// table, or created with gtk_text_buffer_create_tag(), which adds the tag
// to the buffer’s table automatically.
func (t textTag) SetPriority(t TextTag, priority int) {
	var arg0 *C.GtkTextTag
	var arg1 C.gint

	arg0 = (*C.GtkTextTag)(unsafe.Pointer(t.Native()))
	arg1 = C.gint(priority)

	C.gtk_text_tag_set_priority(arg0, arg1)
}
