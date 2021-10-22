// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_editable_get_type()), F: marshalEditabler},
	})
}

// EditableOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type EditableOverrider interface {
	Changed()
	// DeleteText deletes a sequence of characters.
	//
	// The characters that are deleted are those characters at positions from
	// start_pos up to, but not including end_pos. If end_pos is negative, then
	// the characters deleted are those from start_pos to the end of the text.
	//
	// Note that the positions are specified in characters, not bytes.
	DeleteText(startPos, endPos int)
	// DoDeleteText deletes a sequence of characters.
	//
	// The characters that are deleted are those characters at positions from
	// start_pos up to, but not including end_pos. If end_pos is negative, then
	// the characters deleted are those from start_pos to the end of the text.
	//
	// Note that the positions are specified in characters, not bytes.
	DoDeleteText(startPos, endPos int)
	// Delegate gets the GtkEditable that editable is delegating its
	// implementation to.
	//
	// Typically, the delegate is a gtk.Text widget.
	Delegate() Editabler
	// SelectionBounds retrieves the selection bound of the editable.
	//
	// start_pos will be filled with the start of the selection and end_pos with
	// end. If no text was selected both will be identical and FALSE will be
	// returned.
	//
	// Note that positions are specified in characters, not bytes.
	SelectionBounds() (startPos int, endPos int, ok bool)
	// Text retrieves the contents of editable.
	//
	// The returned string is owned by GTK and must not be modified or freed.
	Text() string
	// SetSelectionBounds selects a region of text.
	//
	// The characters that are selected are those characters at positions from
	// start_pos up to, but not including end_pos. If end_pos is negative, then
	// the characters selected are those characters from start_pos to the end of
	// the text.
	//
	// Note that positions are specified in characters, not bytes.
	SetSelectionBounds(startPos, endPos int)
}

// Editable: GtkEditable is an interface for text editing widgets.
//
// Typical examples of editable widgets are gtk.Entry and gtk.SpinButton. It
// contains functions for generically manipulating an editable widget, a large
// number of action signals used for key bindings, and several signals that an
// application can connect to modify the behavior of a widget.
//
// As an example of the latter usage, by connecting the following handler to
// gtk.Editable::insert-text, an application can convert all entry into a widget
// into uppercase.
//
// Forcing entry to uppercase.
//
//    #include <ctype.h>
//
//    void
//    insert_text_handler (GtkEditable *editable,
//                         const char  *text,
//                         int          length,
//                         int         *position,
//                         gpointer     data)
//    {
//      char *result = g_utf8_strup (text, length);
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
//
//
//
// Implementing GtkEditable
//
// The most likely scenario for implementing GtkEditable on your own widget is
// that you will embed a Text inside a complex widget, and want to delegate the
// editable functionality to that text widget. GtkEditable provides some utility
// functions to make this easy.
//
// In your class_init function, call gtk.Editable().InstallProperties, passing
// the first available property ID:
//
//    static void
//    my_class_init (MyClass *class)
//    {
//      ...
//      g_object_class_install_properties (object_class, NUM_PROPERTIES, props);
//      gtk_editable_install_properties (object_clas, NUM_PROPERTIES);
//      ...
//    }
//
//
// In your interface_init function for the GtkEditable interface, provide an
// implementation for the get_delegate vfunc that returns your text widget:
//
//    GtkEditable *
//    get_editable_delegate (GtkEditable *editable)
//    {
//      return GTK_EDITABLE (MY_WIDGET (editable)->text_widget);
//    }
//
//    static void
//    my_editable_init (GtkEditableInterface *iface)
//    {
//      iface->get_delegate = get_editable_delegate;
//    }
//
//
// You don't need to provide any other vfuncs. The default implementations work
// by forwarding to the delegate that the GtkEditableInterface.get_delegate()
// vfunc returns.
//
// In your instance_init function, create your text widget, and then call
// gtk.Editable.InitDelegate():
//
//    static void
//    my_widget_init (MyWidget *self)
//    {
//      ...
//      self->text_widget = gtk_text_new ();
//      gtk_editable_init_delegate (GTK_EDITABLE (self));
//      ...
//    }
//
//
// In your dispose function, call gtk.Editable.FinishDelegate() before
// destroying your text widget:
//
//    static void
//    my_widget_dispose (GObject *object)
//    {
//      ...
//      gtk_editable_finish_delegate (GTK_EDITABLE (self));
//      g_clear_pointer (&self->text_widget, gtk_widget_unparent);
//      ...
//    }
//
//
// Finally, use gtk.Editable().DelegateSetProperty in your set_property function
// (and similar for get_property), to set the editable properties:
//
//      ...
//      if (gtk_editable_delegate_set_property (object, prop_id, value, pspec))
//        return;
//
//      switch (prop_id)
//      ...
//
//
// It is important to note that if you create a GtkEditable that uses a
// delegate, the low level gtk.Editable::insert-text and
// gtk.Editable::delete-text signals will be propagated from the "wrapper"
// editable to the delegate, but they will not be propagated from the delegate
// to the "wrapper" editable, as they would cause an infinite recursion. If you
// wish to connect to the gtk.Editable::insert-text and
// gtk.Editable::delete-text signals, you will need to connect to them on the
// delegate obtained via gtk.Editable.GetDelegate().
type Editable struct {
	Widget
}

// Editabler describes Editable's interface methods.
type Editabler interface {
	externglib.Objector

	// DeleteSelection deletes the currently selected text of the editable.
	DeleteSelection()
	// DeleteText deletes a sequence of characters.
	DeleteText(startPos, endPos int)
	// FinishDelegate undoes the setup done by gtk.Editable.InitDelegate().
	FinishDelegate()
	// Alignment gets the alignment of the editable.
	Alignment() float32
	// Chars retrieves a sequence of characters.
	Chars(startPos, endPos int) string
	// Delegate gets the GtkEditable that editable is delegating its
	// implementation to.
	Delegate() Editabler
	// Editable retrieves whether editable is editable.
	Editable() bool
	// EnableUndo gets if undo/redo actions are enabled for editable.
	EnableUndo() bool
	// MaxWidthChars retrieves the desired maximum width of editable, in
	// characters.
	MaxWidthChars() int
	// Position retrieves the current position of the cursor relative to the
	// start of the content of the editable.
	Position() int
	// SelectionBounds retrieves the selection bound of the editable.
	SelectionBounds() (startPos int, endPos int, ok bool)
	// Text retrieves the contents of editable.
	Text() string
	// WidthChars gets the number of characters of space reserved for the
	// contents of the editable.
	WidthChars() int
	// InitDelegate sets up a delegate for GtkEditable.
	InitDelegate()
	// SelectRegion selects a region of text.
	SelectRegion(startPos, endPos int)
	// SetAlignment sets the alignment for the contents of the editable.
	SetAlignment(xalign float32)
	// SetEditable determines if the user can edit the text in the editable
	// widget.
	SetEditable(isEditable bool)
	// SetEnableUndo: if enabled, changes to editable will be saved for
	// undo/redo actions.
	SetEnableUndo(enableUndo bool)
	// SetMaxWidthChars sets the desired maximum width in characters of
	// editable.
	SetMaxWidthChars(nChars int)
	// SetPosition sets the cursor position in the editable to the given value.
	SetPosition(position int)
	// SetText sets the text in the editable to the given value.
	SetText(text string)
	// SetWidthChars changes the size request of the editable to be about the
	// right size for n_chars characters.
	SetWidthChars(nChars int)
}

var _ Editabler = (*Editable)(nil)

func wrapEditable(obj *externglib.Object) *Editable {
	return &Editable{
		Widget: Widget{
			InitiallyUnowned: externglib.InitiallyUnowned{
				Object: obj,
			},
			Accessible: Accessible{
				Object: obj,
			},
			Buildable: Buildable{
				Object: obj,
			},
			ConstraintTarget: ConstraintTarget{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalEditabler(p uintptr) (interface{}, error) {
	return wrapEditable(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DeleteSelection deletes the currently selected text of the editable.
//
// This call doesn’t do anything if there is no selected text.
func (editable *Editable) DeleteSelection() {
	var _arg0 *C.GtkEditable // out

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))

	C.gtk_editable_delete_selection(_arg0)
	runtime.KeepAlive(editable)
}

// DeleteText deletes a sequence of characters.
//
// The characters that are deleted are those characters at positions from
// start_pos up to, but not including end_pos. If end_pos is negative, then the
// characters deleted are those from start_pos to the end of the text.
//
// Note that the positions are specified in characters, not bytes.
//
// The function takes the following parameters:
//
//    - startPos: start position.
//    - endPos: end position.
//
func (editable *Editable) DeleteText(startPos, endPos int) {
	var _arg0 *C.GtkEditable // out
	var _arg1 C.int          // out
	var _arg2 C.int          // out

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))
	_arg1 = C.int(startPos)
	_arg2 = C.int(endPos)

	C.gtk_editable_delete_text(_arg0, _arg1, _arg2)
	runtime.KeepAlive(editable)
	runtime.KeepAlive(startPos)
	runtime.KeepAlive(endPos)
}

// FinishDelegate undoes the setup done by gtk.Editable.InitDelegate().
//
// This is a helper function that should be called from dispose, before removing
// the delegate object.
func (editable *Editable) FinishDelegate() {
	var _arg0 *C.GtkEditable // out

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))

	C.gtk_editable_finish_delegate(_arg0)
	runtime.KeepAlive(editable)
}

// Alignment gets the alignment of the editable.
func (editable *Editable) Alignment() float32 {
	var _arg0 *C.GtkEditable // out
	var _cret C.float        // in

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))

	_cret = C.gtk_editable_get_alignment(_arg0)
	runtime.KeepAlive(editable)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// Chars retrieves a sequence of characters.
//
// The characters that are retrieved are those characters at positions from
// start_pos up to, but not including end_pos. If end_pos is negative, then the
// characters retrieved are those characters from start_pos to the end of the
// text.
//
// Note that positions are specified in characters, not bytes.
//
// The function takes the following parameters:
//
//    - startPos: start of text.
//    - endPos: end of text.
//
func (editable *Editable) Chars(startPos, endPos int) string {
	var _arg0 *C.GtkEditable // out
	var _arg1 C.int          // out
	var _arg2 C.int          // out
	var _cret *C.char        // in

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))
	_arg1 = C.int(startPos)
	_arg2 = C.int(endPos)

	_cret = C.gtk_editable_get_chars(_arg0, _arg1, _arg2)
	runtime.KeepAlive(editable)
	runtime.KeepAlive(startPos)
	runtime.KeepAlive(endPos)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Delegate gets the GtkEditable that editable is delegating its implementation
// to.
//
// Typically, the delegate is a gtk.Text widget.
func (editable *Editable) Delegate() Editabler {
	var _arg0 *C.GtkEditable // out
	var _cret *C.GtkEditable // in

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))

	_cret = C.gtk_editable_get_delegate(_arg0)
	runtime.KeepAlive(editable)

	var _ret Editabler // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(Editabler)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Editabler")
			}
			_ret = rv
		}
	}

	return _ret
}

// Editable retrieves whether editable is editable.
func (editable *Editable) Editable() bool {
	var _arg0 *C.GtkEditable // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))

	_cret = C.gtk_editable_get_editable(_arg0)
	runtime.KeepAlive(editable)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// EnableUndo gets if undo/redo actions are enabled for editable.
func (editable *Editable) EnableUndo() bool {
	var _arg0 *C.GtkEditable // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))

	_cret = C.gtk_editable_get_enable_undo(_arg0)
	runtime.KeepAlive(editable)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MaxWidthChars retrieves the desired maximum width of editable, in characters.
func (editable *Editable) MaxWidthChars() int {
	var _arg0 *C.GtkEditable // out
	var _cret C.int          // in

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))

	_cret = C.gtk_editable_get_max_width_chars(_arg0)
	runtime.KeepAlive(editable)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Position retrieves the current position of the cursor relative to the start
// of the content of the editable.
//
// Note that this position is in characters, not in bytes.
func (editable *Editable) Position() int {
	var _arg0 *C.GtkEditable // out
	var _cret C.int          // in

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))

	_cret = C.gtk_editable_get_position(_arg0)
	runtime.KeepAlive(editable)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// SelectionBounds retrieves the selection bound of the editable.
//
// start_pos will be filled with the start of the selection and end_pos with
// end. If no text was selected both will be identical and FALSE will be
// returned.
//
// Note that positions are specified in characters, not bytes.
func (editable *Editable) SelectionBounds() (startPos int, endPos int, ok bool) {
	var _arg0 *C.GtkEditable // out
	var _arg1 C.int          // in
	var _arg2 C.int          // in
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))

	_cret = C.gtk_editable_get_selection_bounds(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(editable)

	var _startPos int // out
	var _endPos int   // out
	var _ok bool      // out

	_startPos = int(_arg1)
	_endPos = int(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _startPos, _endPos, _ok
}

// Text retrieves the contents of editable.
//
// The returned string is owned by GTK and must not be modified or freed.
func (editable *Editable) Text() string {
	var _arg0 *C.GtkEditable // out
	var _cret *C.char        // in

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))

	_cret = C.gtk_editable_get_text(_arg0)
	runtime.KeepAlive(editable)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// WidthChars gets the number of characters of space reserved for the contents
// of the editable.
func (editable *Editable) WidthChars() int {
	var _arg0 *C.GtkEditable // out
	var _cret C.int          // in

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))

	_cret = C.gtk_editable_get_width_chars(_arg0)
	runtime.KeepAlive(editable)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// InitDelegate sets up a delegate for GtkEditable.
//
// This is assuming that the get_delegate vfunc in the GtkEditable interface has
// been set up for the editable's type.
//
// This is a helper function that should be called in instance init, after
// creating the delegate object.
func (editable *Editable) InitDelegate() {
	var _arg0 *C.GtkEditable // out

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))

	C.gtk_editable_init_delegate(_arg0)
	runtime.KeepAlive(editable)
}

// SelectRegion selects a region of text.
//
// The characters that are selected are those characters at positions from
// start_pos up to, but not including end_pos. If end_pos is negative, then the
// characters selected are those characters from start_pos to the end of the
// text.
//
// Note that positions are specified in characters, not bytes.
//
// The function takes the following parameters:
//
//    - startPos: start of region.
//    - endPos: end of region.
//
func (editable *Editable) SelectRegion(startPos, endPos int) {
	var _arg0 *C.GtkEditable // out
	var _arg1 C.int          // out
	var _arg2 C.int          // out

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))
	_arg1 = C.int(startPos)
	_arg2 = C.int(endPos)

	C.gtk_editable_select_region(_arg0, _arg1, _arg2)
	runtime.KeepAlive(editable)
	runtime.KeepAlive(startPos)
	runtime.KeepAlive(endPos)
}

// SetAlignment sets the alignment for the contents of the editable.
//
// This controls the horizontal positioning of the contents when the displayed
// text is shorter than the width of the editable.
//
// The function takes the following parameters:
//
//    - xalign: horizontal alignment, from 0 (left) to 1 (right). Reversed for
//    RTL layouts.
//
func (editable *Editable) SetAlignment(xalign float32) {
	var _arg0 *C.GtkEditable // out
	var _arg1 C.float        // out

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))
	_arg1 = C.float(xalign)

	C.gtk_editable_set_alignment(_arg0, _arg1)
	runtime.KeepAlive(editable)
	runtime.KeepAlive(xalign)
}

// SetEditable determines if the user can edit the text in the editable widget.
//
// The function takes the following parameters:
//
//    - isEditable: TRUE if the user is allowed to edit the text in the widget.
//
func (editable *Editable) SetEditable(isEditable bool) {
	var _arg0 *C.GtkEditable // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))
	if isEditable {
		_arg1 = C.TRUE
	}

	C.gtk_editable_set_editable(_arg0, _arg1)
	runtime.KeepAlive(editable)
	runtime.KeepAlive(isEditable)
}

// SetEnableUndo: if enabled, changes to editable will be saved for undo/redo
// actions.
//
// This results in an additional copy of text changes and are not stored in
// secure memory. As such, undo is forcefully disabled when gtk.Text:visibility
// is set to FALSE.
//
// The function takes the following parameters:
//
//    - enableUndo: if undo/redo should be enabled.
//
func (editable *Editable) SetEnableUndo(enableUndo bool) {
	var _arg0 *C.GtkEditable // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))
	if enableUndo {
		_arg1 = C.TRUE
	}

	C.gtk_editable_set_enable_undo(_arg0, _arg1)
	runtime.KeepAlive(editable)
	runtime.KeepAlive(enableUndo)
}

// SetMaxWidthChars sets the desired maximum width in characters of editable.
//
// The function takes the following parameters:
//
//    - nChars: new desired maximum width, in characters.
//
func (editable *Editable) SetMaxWidthChars(nChars int) {
	var _arg0 *C.GtkEditable // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))
	_arg1 = C.int(nChars)

	C.gtk_editable_set_max_width_chars(_arg0, _arg1)
	runtime.KeepAlive(editable)
	runtime.KeepAlive(nChars)
}

// SetPosition sets the cursor position in the editable to the given value.
//
// The cursor is displayed before the character with the given (base 0) index in
// the contents of the editable. The value must be less than or equal to the
// number of characters in the editable. A value of -1 indicates that the
// position should be set after the last character of the editable. Note that
// position is in characters, not in bytes.
//
// The function takes the following parameters:
//
//    - position of the cursor.
//
func (editable *Editable) SetPosition(position int) {
	var _arg0 *C.GtkEditable // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))
	_arg1 = C.int(position)

	C.gtk_editable_set_position(_arg0, _arg1)
	runtime.KeepAlive(editable)
	runtime.KeepAlive(position)
}

// SetText sets the text in the editable to the given value.
//
// This is replacing the current contents.
//
// The function takes the following parameters:
//
//    - text to set.
//
func (editable *Editable) SetText(text string) {
	var _arg0 *C.GtkEditable // out
	var _arg1 *C.char        // out

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_editable_set_text(_arg0, _arg1)
	runtime.KeepAlive(editable)
	runtime.KeepAlive(text)
}

// SetWidthChars changes the size request of the editable to be about the right
// size for n_chars characters.
//
// Note that it changes the size request, the size can still be affected by how
// you pack the widget into containers. If n_chars is -1, the size reverts to
// the default size.
//
// The function takes the following parameters:
//
//    - nChars: width in chars.
//
func (editable *Editable) SetWidthChars(nChars int) {
	var _arg0 *C.GtkEditable // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkEditable)(unsafe.Pointer(editable.Native()))
	_arg1 = C.int(nChars)

	C.gtk_editable_set_width_chars(_arg0, _arg1)
	runtime.KeepAlive(editable)
	runtime.KeepAlive(nChars)
}

// ConnectChanged: emitted at the end of a single user-visible operation on the
// contents.
//
// E.g., a paste operation that replaces the contents of the selection will
// cause only one signal emission (even though it is implemented by first
// deleting the selection, then inserting the new content, and may cause
// multiple ::notify::text signals to be emitted).
func (editable *Editable) ConnectChanged(f func()) externglib.SignalHandle {
	return editable.Connect("changed", f)
}

// ConnectDeleteText: emitted when text is deleted from the widget by the user.
//
// The default handler for this signal will normally be responsible for deleting
// the text, so by connecting to this signal and then stopping the signal with
// g_signal_stop_emission(), it is possible to modify the range of deleted text,
// or prevent it from being deleted entirely.
//
// The start_pos and end_pos parameters are interpreted as for
// gtk.Editable.DeleteText().
func (editable *Editable) ConnectDeleteText(f func(startPos, endPos int)) externglib.SignalHandle {
	return editable.Connect("delete-text", f)
}
