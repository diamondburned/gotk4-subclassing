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
		{T: externglib.Type(C.gtk_expander_get_type()), F: marshalExpander},
	})
}

// Expander: a Expander allows the user to hide or show its child by clicking on
// an expander triangle similar to the triangles used in a TreeView.
//
// Normally you use an expander as you would use any other descendant of Bin;
// you create the child widget and use gtk_container_add() to add it to the
// expander. When the expander is toggled, it will take care of showing and
// hiding the child automatically.
//
//
// Special Usage
//
// There are situations in which you may prefer to show and hide the expanded
// widget yourself, such as when you want to actually create the widget at
// expansion time. In this case, create a Expander but do not add a child to it.
// The expander widget has an Expander:expanded property which can be used to
// monitor its expansion state. You should watch this property with a signal
// connection as follows:
//
//    expander
//    ├── title
//    │   ├── arrow
//    │   ╰── <label widget>
//    ╰── <child>
//
// GtkExpander has three CSS nodes, the main node with the name expander, a
// subnode with name title and node below it with name arrow. The arrow of an
// expander that is showing its child gets the :checked pseudoclass added to it.
type Expander interface {
	Bin
	Buildable

	// Expanded queries a Expander and returns its current state. Returns true
	// if the child widget is revealed.
	//
	// See gtk_expander_set_expanded().
	Expanded(e Expander) bool
	// Label fetches the text from a label widget including any embedded
	// underlines indicating mnemonics and Pango markup, as set by
	// gtk_expander_set_label(). If the label text has not been set the return
	// value will be nil. This will be the case if you create an empty button
	// with gtk_button_new() to use as a container.
	//
	// Note that this function behaved differently in versions prior to 2.14 and
	// used to return the label text stripped of embedded underlines indicating
	// mnemonics and Pango markup. This problem can be avoided by fetching the
	// label text directly from the label widget.
	Label(e Expander)
	// LabelFill returns whether the label widget will fill all available
	// horizontal space allocated to @expander.
	LabelFill(e Expander) bool
	// LabelWidget retrieves the label widget for the frame. See
	// gtk_expander_set_label_widget().
	LabelWidget(e Expander)
	// ResizeToplevel returns whether the expander will resize the toplevel
	// widget containing the expander upon resizing and collpasing.
	ResizeToplevel(e Expander) bool
	// Spacing gets the value set by gtk_expander_set_spacing().
	Spacing(e Expander)
	// UseMarkup returns whether the label’s text is interpreted as marked up
	// with the [Pango text markup language][PangoMarkupFormat]. See
	// gtk_expander_set_use_markup().
	UseMarkup(e Expander) bool
	// UseUnderline returns whether an embedded underline in the expander label
	// indicates a mnemonic. See gtk_expander_set_use_underline().
	UseUnderline(e Expander) bool
	// SetExpanded sets the state of the expander. Set to true, if you want the
	// child widget to be revealed, and false if you want the child widget to be
	// hidden.
	SetExpanded(e Expander, expanded bool)
	// SetLabel sets the text of the label of the expander to @label.
	//
	// This will also clear any previously set labels.
	SetLabel(e Expander, label string)
	// SetLabelFill sets whether the label widget should fill all available
	// horizontal space allocated to @expander.
	//
	// Note that this function has no effect since 3.20.
	SetLabelFill(e Expander, labelFill bool)
	// SetLabelWidget: set the label widget for the expander. This is the widget
	// that will appear embedded alongside the expander arrow.
	SetLabelWidget(e Expander, labelWidget Widget)
	// SetResizeToplevel sets whether the expander will resize the toplevel
	// widget containing the expander upon resizing and collpasing.
	SetResizeToplevel(e Expander, resizeToplevel bool)
	// SetSpacing sets the spacing field of @expander, which is the number of
	// pixels to place between expander and the child.
	SetSpacing(e Expander, spacing int)
	// SetUseMarkup sets whether the text of the label contains markup in
	// [Pango’s text markup language][PangoMarkupFormat]. See
	// gtk_label_set_markup().
	SetUseMarkup(e Expander, useMarkup bool)
	// SetUseUnderline: if true, an underline in the text of the expander label
	// indicates the next character should be used for the mnemonic accelerator
	// key.
	SetUseUnderline(e Expander, useUnderline bool)
}

// expander implements the Expander interface.
type expander struct {
	Bin
	Buildable
}

var _ Expander = (*expander)(nil)

// WrapExpander wraps a GObject to the right type. It is
// primarily used internally.
func WrapExpander(obj *externglib.Object) Expander {
	return Expander{
		Bin:       WrapBin(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalExpander(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapExpander(obj), nil
}

// NewExpander constructs a class Expander.
func NewExpander(label string) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_expander_new(arg1)
}

// NewExpanderWithMnemonic constructs a class Expander.
func NewExpanderWithMnemonic(label string) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_expander_new_with_mnemonic(arg1)
}

// Expanded queries a Expander and returns its current state. Returns true
// if the child widget is revealed.
//
// See gtk_expander_set_expanded().
func (e expander) Expanded(e Expander) bool {
	var arg0 *C.GtkExpander

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_expander_get_expanded(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Label fetches the text from a label widget including any embedded
// underlines indicating mnemonics and Pango markup, as set by
// gtk_expander_set_label(). If the label text has not been set the return
// value will be nil. This will be the case if you create an empty button
// with gtk_button_new() to use as a container.
//
// Note that this function behaved differently in versions prior to 2.14 and
// used to return the label text stripped of embedded underlines indicating
// mnemonics and Pango markup. This problem can be avoided by fetching the
// label text directly from the label widget.
func (e expander) Label(e Expander) {
	var arg0 *C.GtkExpander

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	C.gtk_expander_get_label(arg0)
}

// LabelFill returns whether the label widget will fill all available
// horizontal space allocated to @expander.
func (e expander) LabelFill(e Expander) bool {
	var arg0 *C.GtkExpander

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_expander_get_label_fill(arg0)

	if cret {
		ok = true
	}

	return ok
}

// LabelWidget retrieves the label widget for the frame. See
// gtk_expander_set_label_widget().
func (e expander) LabelWidget(e Expander) {
	var arg0 *C.GtkExpander

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	C.gtk_expander_get_label_widget(arg0)
}

// ResizeToplevel returns whether the expander will resize the toplevel
// widget containing the expander upon resizing and collpasing.
func (e expander) ResizeToplevel(e Expander) bool {
	var arg0 *C.GtkExpander

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_expander_get_resize_toplevel(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Spacing gets the value set by gtk_expander_set_spacing().
func (e expander) Spacing(e Expander) {
	var arg0 *C.GtkExpander

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	C.gtk_expander_get_spacing(arg0)
}

// UseMarkup returns whether the label’s text is interpreted as marked up
// with the [Pango text markup language][PangoMarkupFormat]. See
// gtk_expander_set_use_markup().
func (e expander) UseMarkup(e Expander) bool {
	var arg0 *C.GtkExpander

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_expander_get_use_markup(arg0)

	if cret {
		ok = true
	}

	return ok
}

// UseUnderline returns whether an embedded underline in the expander label
// indicates a mnemonic. See gtk_expander_set_use_underline().
func (e expander) UseUnderline(e Expander) bool {
	var arg0 *C.GtkExpander

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_expander_get_use_underline(arg0)

	if cret {
		ok = true
	}

	return ok
}

// SetExpanded sets the state of the expander. Set to true, if you want the
// child widget to be revealed, and false if you want the child widget to be
// hidden.
func (e expander) SetExpanded(e Expander, expanded bool) {
	var arg0 *C.GtkExpander
	var arg1 C.gboolean

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	if expanded {
		arg1 = C.gboolean(1)
	}

	C.gtk_expander_set_expanded(arg0, arg1)
}

// SetLabel sets the text of the label of the expander to @label.
//
// This will also clear any previously set labels.
func (e expander) SetLabel(e Expander, label string) {
	var arg0 *C.GtkExpander
	var arg1 *C.gchar

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_expander_set_label(arg0, arg1)
}

// SetLabelFill sets whether the label widget should fill all available
// horizontal space allocated to @expander.
//
// Note that this function has no effect since 3.20.
func (e expander) SetLabelFill(e Expander, labelFill bool) {
	var arg0 *C.GtkExpander
	var arg1 C.gboolean

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	if labelFill {
		arg1 = C.gboolean(1)
	}

	C.gtk_expander_set_label_fill(arg0, arg1)
}

// SetLabelWidget: set the label widget for the expander. This is the widget
// that will appear embedded alongside the expander arrow.
func (e expander) SetLabelWidget(e Expander, labelWidget Widget) {
	var arg0 *C.GtkExpander
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(labelWidget.Native()))

	C.gtk_expander_set_label_widget(arg0, arg1)
}

// SetResizeToplevel sets whether the expander will resize the toplevel
// widget containing the expander upon resizing and collpasing.
func (e expander) SetResizeToplevel(e Expander, resizeToplevel bool) {
	var arg0 *C.GtkExpander
	var arg1 C.gboolean

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	if resizeToplevel {
		arg1 = C.gboolean(1)
	}

	C.gtk_expander_set_resize_toplevel(arg0, arg1)
}

// SetSpacing sets the spacing field of @expander, which is the number of
// pixels to place between expander and the child.
func (e expander) SetSpacing(e Expander, spacing int) {
	var arg0 *C.GtkExpander
	var arg1 C.gint

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	arg1 = C.gint(spacing)

	C.gtk_expander_set_spacing(arg0, arg1)
}

// SetUseMarkup sets whether the text of the label contains markup in
// [Pango’s text markup language][PangoMarkupFormat]. See
// gtk_label_set_markup().
func (e expander) SetUseMarkup(e Expander, useMarkup bool) {
	var arg0 *C.GtkExpander
	var arg1 C.gboolean

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	if useMarkup {
		arg1 = C.gboolean(1)
	}

	C.gtk_expander_set_use_markup(arg0, arg1)
}

// SetUseUnderline: if true, an underline in the text of the expander label
// indicates the next character should be used for the mnemonic accelerator
// key.
func (e expander) SetUseUnderline(e Expander, useUnderline bool) {
	var arg0 *C.GtkExpander
	var arg1 C.gboolean

	arg0 = (*C.GtkExpander)(unsafe.Pointer(e.Native()))
	if useUnderline {
		arg1 = C.gboolean(1)
	}

	C.gtk_expander_set_use_underline(arg0, arg1)
}