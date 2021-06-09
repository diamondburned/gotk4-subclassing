// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_frame_get_type()), F: marshalFrame},
	})
}

// Frame: the frame widget is a widget that surrounds its child with a
// decorative frame and an optional label. If present, the label is drawn inside
// the top edge of the frame. The horizontal position of the label can be
// controlled with gtk_frame_set_label_align().
//
// GtkFrame clips its child. You can use this to add rounded corners to widgets,
// but be aware that it also cuts off shadows.
//
//
// GtkFrame as GtkBuildable
//
// The GtkFrame implementation of the GtkBuildable interface supports placing a
// child in the label position by specifying “label” as the “type” attribute of
// a <child> element. A normal content child can be specified without specifying
// a <child> type attribute.
//
// An example of a UI definition fragment with GtkFrame:
//
//    <object class="GtkFrame">
//      <child type="label">
//        <object class="GtkLabel" id="frame_label"/>
//      </child>
//      <child>
//        <object class="GtkEntry" id="frame_content"/>
//      </child>
//    </object>
//
// CSS nodes
//
//    frame
//    ├── <label widget>
//    ╰── <child>
//
// GtkFrame has a main CSS node with name “frame”, which is used to draw the
// visible border. You can set the appearance of the border using CSS properties
// like “border-style” on this node.
type Frame interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget

	// Child gets the child widget of @frame.
	Child() Widget
	// Label: if the frame’s label widget is a Label, returns the text in the
	// label widget. (The frame will have a Label for the label widget if a
	// non-nil argument was passed to gtk_frame_new().)
	Label() string
	// LabelAlign retrieves the X alignment of the frame’s label. See
	// gtk_frame_set_label_align().
	LabelAlign() float32
	// LabelWidget retrieves the label widget for the frame. See
	// gtk_frame_set_label_widget().
	LabelWidget() Widget
	// SetChild sets the child widget of @frame.
	SetChild(child Widget)
	// SetLabel removes the current Frame:label-widget. If @label is not nil,
	// creates a new Label with that text and adds it as the Frame:label-widget.
	SetLabel(label string)
	// SetLabelAlign sets the X alignment of the frame widget’s label. The
	// default value for a newly created frame is 0.0.
	SetLabelAlign(xalign float32)
	// SetLabelWidget sets the Frame:label-widget for the frame. This is the
	// widget that will appear embedded in the top edge of the frame as a title.
	SetLabelWidget(labelWidget Widget)
}

// frame implements the Frame interface.
type frame struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
}

var _ Frame = (*frame)(nil)

// WrapFrame wraps a GObject to the right type. It is
// primarily used internally.
func WrapFrame(obj *externglib.Object) Frame {
	return Frame{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalFrame(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFrame(obj), nil
}

// NewFrame constructs a class Frame.
func NewFrame(label string) Frame {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkFrame

	cret = C.gtk_frame_new(arg1)

	var frame Frame

	frame = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Frame)

	return frame
}

// Child gets the child widget of @frame.
func (f frame) Child() Widget {
	var arg0 *C.GtkFrame

	arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))

	var cret *C.GtkWidget

	cret = C.gtk_frame_get_child(arg0)

	var widget Widget

	widget = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return widget
}

// Label: if the frame’s label widget is a Label, returns the text in the
// label widget. (The frame will have a Label for the label widget if a
// non-nil argument was passed to gtk_frame_new().)
func (f frame) Label() string {
	var arg0 *C.GtkFrame

	arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))

	var cret *C.char

	cret = C.gtk_frame_get_label(arg0)

	var utf8 string

	utf8 = C.GoString(cret)

	return utf8
}

// LabelAlign retrieves the X alignment of the frame’s label. See
// gtk_frame_set_label_align().
func (f frame) LabelAlign() float32 {
	var arg0 *C.GtkFrame

	arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))

	var cret C.float

	cret = C.gtk_frame_get_label_align(arg0)

	var gfloat float32

	gfloat = (float32)(cret)

	return gfloat
}

// LabelWidget retrieves the label widget for the frame. See
// gtk_frame_set_label_widget().
func (f frame) LabelWidget() Widget {
	var arg0 *C.GtkFrame

	arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))

	var cret *C.GtkWidget

	cret = C.gtk_frame_get_label_widget(arg0)

	var widget Widget

	widget = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Widget)

	return widget
}

// SetChild sets the child widget of @frame.
func (f frame) SetChild(child Widget) {
	var arg0 *C.GtkFrame
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_frame_set_child(arg0, arg1)
}

// SetLabel removes the current Frame:label-widget. If @label is not nil,
// creates a new Label with that text and adds it as the Frame:label-widget.
func (f frame) SetLabel(label string) {
	var arg0 *C.GtkFrame
	var arg1 *C.char

	arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))
	arg1 = (*C.char)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_frame_set_label(arg0, arg1)
}

// SetLabelAlign sets the X alignment of the frame widget’s label. The
// default value for a newly created frame is 0.0.
func (f frame) SetLabelAlign(xalign float32) {
	var arg0 *C.GtkFrame
	var arg1 C.float

	arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))
	arg1 = C.float(xalign)

	C.gtk_frame_set_label_align(arg0, arg1)
}

// SetLabelWidget sets the Frame:label-widget for the frame. This is the
// widget that will appear embedded in the top edge of the frame as a title.
func (f frame) SetLabelWidget(labelWidget Widget) {
	var arg0 *C.GtkFrame
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(labelWidget.Native()))

	C.gtk_frame_set_label_widget(arg0, arg1)
}
