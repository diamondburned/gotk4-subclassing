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
		{T: externglib.Type(C.gtk_editable_get_type()), F: marshalEditable},
	})
}

// EditableOverrider contains methods that are overridable. This
// interface is a subset of the interface Editable.
type EditableOverrider interface {
	Changed(e Editable)
	// DeleteText deletes a sequence of characters. The characters that are
	// deleted are those characters at positions from @start_pos up to, but not
	// including @end_pos. If @end_pos is negative, then the characters deleted
	// are those from @start_pos to the end of the text.
	//
	// Note that the positions are specified in characters, not bytes.
	DeleteText(e Editable, startPos int, endPos int)
	// DoDeleteText deletes a sequence of characters. The characters that are
	// deleted are those characters at positions from @start_pos up to, but not
	// including @end_pos. If @end_pos is negative, then the characters deleted
	// are those from @start_pos to the end of the text.
	//
	// Note that the positions are specified in characters, not bytes.
	DoDeleteText(e Editable, startPos int, endPos int)
	// DoInsertText inserts @new_text_length bytes of @new_text into the
	// contents of the widget, at position @position.
	//
	// Note that the position is in characters, not in bytes. The function
	// updates @position to point after the newly inserted text.
	DoInsertText(e Editable, newText string, newTextLength int, position int)
	// Chars retrieves a sequence of characters. The characters that are
	// retrieved are those characters at positions from @start_pos up to, but
	// not including @end_pos. If @end_pos is negative, then the characters
	// retrieved are those characters from @start_pos to the end of the text.
	//
	// Note that positions are specified in characters, not bytes.
	Chars(e Editable, startPos int, endPos int)
	// Position retrieves the current position of the cursor relative to the
	// start of the content of the editable.
	//
	// Note that this position is in characters, not in bytes.
	Position(e Editable)
	// SelectionBounds retrieves the selection bound of the editable. start_pos
	// will be filled with the start of the selection and @end_pos with end. If
	// no text was selected both will be identical and false will be returned.
	//
	// Note that positions are specified in characters, not bytes.
	SelectionBounds(e Editable) (startPos int, endPos int, ok bool)
	// InsertText inserts @new_text_length bytes of @new_text into the contents
	// of the widget, at position @position.
	//
	// Note that the position is in characters, not in bytes. The function
	// updates @position to point after the newly inserted text.
	InsertText(e Editable, newText string, newTextLength int, position int)
	// SetPosition sets the cursor position in the editable to the given value.
	//
	// The cursor is displayed before the character with the given (base 0)
	// index in the contents of the editable. The value must be less than or
	// equal to the number of characters in the editable. A value of -1
	// indicates that the position should be set after the last character of the
	// editable. Note that @position is in characters, not in bytes.
	SetPosition(e Editable, position int)
	// SetSelectionBounds selects a region of text. The characters that are
	// selected are those characters at positions from @start_pos up to, but not
	// including @end_pos. If @end_pos is negative, then the characters selected
	// are those characters from @start_pos to the end of the text.
	//
	// Note that positions are specified in characters, not bytes.
	SetSelectionBounds(e Editable, startPos int, endPos int)
}

// Editable: the Editable interface is an interface which should be implemented
// by text editing widgets, such as Entry and SpinButton. It contains functions
// for generically manipulating an editable widget, a large number of action
// signals used for key bindings, and several signals that an application can
// connect to to modify the behavior of a widget.
//
// As an example of the latter usage, by connecting the following handler to
// Editable::insert-text, an application can convert all entry into a widget
// into uppercase.
//
// Forcing entry to uppercase.
//
//    #include <ctype.h>;
//
//    void
//    insert_text_handler (GtkEditable *editable,
//                         const gchar *text,
//                         gint         length,
//                         gint        *position,
//                         gpointer     data)
//    {
//      gchar *result = g_utf8_strup (text, length);
//
//      g_signal_handlers_block_by_func (editable,
//                                   (gpointer) insert_text_handler, data);
//      gtk_editable_insert_text (editable, result, length, position);
//      g_signal_handlers_unblock_by_func (editable,
//                                         (gpointer) insert_text_handler, data);
//
//      g_signal_stop_emission_by_name (editable, "insert_text");
//
//      g_free (result);
//    }
type Editable interface {
	gextras.Objector
	EditableOverrider

	// CopyClipboard copies the contents of the currently selected content in
	// the editable and puts it on the clipboard.
	CopyClipboard(e Editable)
	// CutClipboard removes the contents of the currently selected content in
	// the editable and puts it on the clipboard.
	CutClipboard(e Editable)
	// DeleteSelection deletes the currently selected text of the editable. This
	// call doesn’t do anything if there is no selected text.
	DeleteSelection(e Editable)
	// Editable retrieves whether @editable is editable. See
	// gtk_editable_set_editable().
	Editable(e Editable) bool
	// PasteClipboard pastes the content of the clipboard to the current
	// position of the cursor in the editable.
	PasteClipboard(e Editable)
	// SelectRegion selects a region of text. The characters that are selected
	// are those characters at positions from @start_pos up to, but not
	// including @end_pos. If @end_pos is negative, then the characters selected
	// are those characters from @start_pos to the end of the text.
	//
	// Note that positions are specified in characters, not bytes.
	SelectRegion(e Editable, startPos int, endPos int)
	// SetEditable determines if the user can edit the text in the editable
	// widget or not.
	SetEditable(e Editable, isEditable bool)
}

// editable implements the Editable interface.
type editable struct {
	gextras.Objector
}

var _ Editable = (*editable)(nil)

// WrapEditable wraps a GObject to a type that implements interface
// Editable. It is primarily used internally.
func WrapEditable(obj *externglib.Object) Editable {
	return Editable{
		Objector: obj,
	}
}

func marshalEditable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEditable(obj), nil
}

// CopyClipboard copies the contents of the currently selected content in
// the editable and puts it on the clipboard.
func (e editable) CopyClipboard(e Editable) {
	var arg0 *C.GtkEditable

	arg0 = (*C.GtkEditable)(unsafe.Pointer(e.Native()))

	C.gtk_editable_copy_clipboard(arg0)
}

// CutClipboard removes the contents of the currently selected content in
// the editable and puts it on the clipboard.
func (e editable) CutClipboard(e Editable) {
	var arg0 *C.GtkEditable

	arg0 = (*C.GtkEditable)(unsafe.Pointer(e.Native()))

	C.gtk_editable_cut_clipboard(arg0)
}

// DeleteSelection deletes the currently selected text of the editable. This
// call doesn’t do anything if there is no selected text.
func (e editable) DeleteSelection(e Editable) {
	var arg0 *C.GtkEditable

	arg0 = (*C.GtkEditable)(unsafe.Pointer(e.Native()))

	C.gtk_editable_delete_selection(arg0)
}

// DeleteText deletes a sequence of characters. The characters that are
// deleted are those characters at positions from @start_pos up to, but not
// including @end_pos. If @end_pos is negative, then the characters deleted
// are those from @start_pos to the end of the text.
//
// Note that the positions are specified in characters, not bytes.
func (e editable) DeleteText(e Editable, startPos int, endPos int) {
	var arg0 *C.GtkEditable
	var arg1 C.gint
	var arg2 C.gint

	arg0 = (*C.GtkEditable)(unsafe.Pointer(e.Native()))
	arg1 = C.gint(startPos)
	arg2 = C.gint(endPos)

	C.gtk_editable_delete_text(arg0, arg1, arg2)
}

// Chars retrieves a sequence of characters. The characters that are
// retrieved are those characters at positions from @start_pos up to, but
// not including @end_pos. If @end_pos is negative, then the characters
// retrieved are those characters from @start_pos to the end of the text.
//
// Note that positions are specified in characters, not bytes.
func (e editable) Chars(e Editable, startPos int, endPos int) {
	var arg0 *C.GtkEditable
	var arg1 C.gint
	var arg2 C.gint

	arg0 = (*C.GtkEditable)(unsafe.Pointer(e.Native()))
	arg1 = C.gint(startPos)
	arg2 = C.gint(endPos)

	C.gtk_editable_get_chars(arg0, arg1, arg2)
}

// Editable retrieves whether @editable is editable. See
// gtk_editable_set_editable().
func (e editable) Editable(e Editable) bool {
	var arg0 *C.GtkEditable

	arg0 = (*C.GtkEditable)(unsafe.Pointer(e.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_editable_get_editable(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Position retrieves the current position of the cursor relative to the
// start of the content of the editable.
//
// Note that this position is in characters, not in bytes.
func (e editable) Position(e Editable) {
	var arg0 *C.GtkEditable

	arg0 = (*C.GtkEditable)(unsafe.Pointer(e.Native()))

	C.gtk_editable_get_position(arg0)
}

// SelectionBounds retrieves the selection bound of the editable. start_pos
// will be filled with the start of the selection and @end_pos with end. If
// no text was selected both will be identical and false will be returned.
//
// Note that positions are specified in characters, not bytes.
func (e editable) SelectionBounds(e Editable) (startPos int, endPos int, ok bool) {
	var arg0 *C.GtkEditable

	arg0 = (*C.GtkEditable)(unsafe.Pointer(e.Native()))

	var arg1 C.gint
	var startPos int
	var arg2 C.gint
	var endPos int
	var cret C.gboolean
	var ok bool

	cret = C.gtk_editable_get_selection_bounds(arg0, &arg1, &arg2)

	startPos = int(&arg1)
	endPos = int(&arg2)
	if cret {
		ok = true
	}

	return startPos, endPos, ok
}

// InsertText inserts @new_text_length bytes of @new_text into the contents
// of the widget, at position @position.
//
// Note that the position is in characters, not in bytes. The function
// updates @position to point after the newly inserted text.
func (e editable) InsertText(e Editable, newText string, newTextLength int, position int) {
	var arg0 *C.GtkEditable
	var arg1 *C.gchar
	var arg2 C.gint
	var arg3 *C.gint

	arg0 = (*C.GtkEditable)(unsafe.Pointer(e.Native()))
	arg1 = (*C.gchar)(C.CString(newText))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gint(newTextLength)
	arg3 = *C.gint(position)

	C.gtk_editable_insert_text(arg0, arg1, arg2, arg3)
}

// PasteClipboard pastes the content of the clipboard to the current
// position of the cursor in the editable.
func (e editable) PasteClipboard(e Editable) {
	var arg0 *C.GtkEditable

	arg0 = (*C.GtkEditable)(unsafe.Pointer(e.Native()))

	C.gtk_editable_paste_clipboard(arg0)
}

// SelectRegion selects a region of text. The characters that are selected
// are those characters at positions from @start_pos up to, but not
// including @end_pos. If @end_pos is negative, then the characters selected
// are those characters from @start_pos to the end of the text.
//
// Note that positions are specified in characters, not bytes.
func (e editable) SelectRegion(e Editable, startPos int, endPos int) {
	var arg0 *C.GtkEditable
	var arg1 C.gint
	var arg2 C.gint

	arg0 = (*C.GtkEditable)(unsafe.Pointer(e.Native()))
	arg1 = C.gint(startPos)
	arg2 = C.gint(endPos)

	C.gtk_editable_select_region(arg0, arg1, arg2)
}

// SetEditable determines if the user can edit the text in the editable
// widget or not.
func (e editable) SetEditable(e Editable, isEditable bool) {
	var arg0 *C.GtkEditable
	var arg1 C.gboolean

	arg0 = (*C.GtkEditable)(unsafe.Pointer(e.Native()))
	if isEditable {
		arg1 = C.gboolean(1)
	}

	C.gtk_editable_set_editable(arg0, arg1)
}

// SetPosition sets the cursor position in the editable to the given value.
//
// The cursor is displayed before the character with the given (base 0)
// index in the contents of the editable. The value must be less than or
// equal to the number of characters in the editable. A value of -1
// indicates that the position should be set after the last character of the
// editable. Note that @position is in characters, not in bytes.
func (e editable) SetPosition(e Editable, position int) {
	var arg0 *C.GtkEditable
	var arg1 C.gint

	arg0 = (*C.GtkEditable)(unsafe.Pointer(e.Native()))
	arg1 = C.gint(position)

	C.gtk_editable_set_position(arg0, arg1)
}