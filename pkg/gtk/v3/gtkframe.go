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
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_frame_get_type()), F: marshalFrame},
	})
}

// Frame: the frame widget is a bin that surrounds its child with a decorative
// frame and an optional label. If present, the label is drawn in a gap in the
// top side of the frame. The position of the label can be controlled with
// gtk_frame_set_label_align().
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
//        <object class="GtkLabel" id="frame-label"/>
//      </child>
//      <child>
//        <object class="GtkEntry" id="frame-content"/>
//      </child>
//    </object>
//
// CSS nodes
//
//    frame
//    ├── border[.flat]
//    ├── <label widget>
//    ╰── <child>
//
// GtkFrame has a main CSS node named “frame” and a subnode named “border”. The
// “border” node is used to draw the visible border. You can set the appearance
// of the border using CSS properties like “border-style” on the “border” node.
//
// The border node can be given the style class “.flat”, which is used by themes
// to disable drawing of the border. To do this from code, call
// gtk_frame_set_shadow_type() with GTK_SHADOW_NONE to add the “.flat” class or
// any other shadow type to remove it.
type Frame interface {
	Bin
	Buildable

	// Label: if the frame’s label widget is a Label, returns the text in the
	// label widget. (The frame will have a Label for the label widget if a
	// non-nil argument was passed to gtk_frame_new().)
	Label() string
	// LabelAlign retrieves the X and Y alignment of the frame’s label. See
	// gtk_frame_set_label_align().
	LabelAlign() (xalign float32, yalign float32)
	// LabelWidget retrieves the label widget for the frame. See
	// gtk_frame_set_label_widget().
	LabelWidget() Widget
	// ShadowType retrieves the shadow type of the frame. See
	// gtk_frame_set_shadow_type().
	ShadowType() ShadowType
	// SetLabel removes the current Frame:label-widget. If @label is not nil,
	// creates a new Label with that text and adds it as the Frame:label-widget.
	SetLabel(label string)
	// SetLabelAlign sets the alignment of the frame widget’s label. The default
	// values for a newly created frame are 0.0 and 0.5.
	SetLabelAlign(xalign float32, yalign float32)
	// SetLabelWidget sets the Frame:label-widget for the frame. This is the
	// widget that will appear embedded in the top edge of the frame as a title.
	SetLabelWidget(labelWidget Widget)
	// SetShadowType sets the Frame:shadow-type for @frame, i.e. whether it is
	// drawn without (GTK_SHADOW_NONE) or with (other values) a visible border.
	// Values other than GTK_SHADOW_NONE are treated identically by GtkFrame.
	// The chosen type is applied by removing or adding the .flat class to the
	// CSS node named border.
	SetShadowType(typ ShadowType)
}

// frame implements the Frame interface.
type frame struct {
	Bin
	Buildable
}

var _ Frame = (*frame)(nil)

// WrapFrame wraps a GObject to the right type. It is
// primarily used internally.
func WrapFrame(obj *externglib.Object) Frame {
	return Frame{
		Bin:       WrapBin(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalFrame(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFrame(obj), nil
}

// NewFrame constructs a class Frame.
func NewFrame(label string) Frame {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkFrame

	cret = C.gtk_frame_new(arg1)

	var frame Frame

	frame = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(Frame)

	return frame
}

// Label: if the frame’s label widget is a Label, returns the text in the
// label widget. (The frame will have a Label for the label widget if a
// non-nil argument was passed to gtk_frame_new().)
func (f frame) Label() string {
	var arg0 *C.GtkFrame

	arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))

	var cret *C.gchar

	cret = C.gtk_frame_get_label(arg0)

	var utf8 string

	utf8 = C.GoString(cret)

	return utf8
}

// LabelAlign retrieves the X and Y alignment of the frame’s label. See
// gtk_frame_set_label_align().
func (f frame) LabelAlign() (xalign float32, yalign float32) {
	var arg0 *C.GtkFrame

	arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))

	var arg1 C.gfloat
	var arg2 C.gfloat

	C.gtk_frame_get_label_align(arg0, &arg1, &arg2)

	var xalign float32
	var yalign float32

	xalign = (float32)(arg1)
	yalign = (float32)(arg2)

	return xalign, yalign
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

// ShadowType retrieves the shadow type of the frame. See
// gtk_frame_set_shadow_type().
func (f frame) ShadowType() ShadowType {
	var arg0 *C.GtkFrame

	arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))

	var cret C.GtkShadowType

	cret = C.gtk_frame_get_shadow_type(arg0)

	var shadowType ShadowType

	shadowType = ShadowType(cret)

	return shadowType
}

// SetLabel removes the current Frame:label-widget. If @label is not nil,
// creates a new Label with that text and adds it as the Frame:label-widget.
func (f frame) SetLabel(label string) {
	var arg0 *C.GtkFrame
	var arg1 *C.gchar

	arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))
	arg1 = (*C.gchar)(C.CString(label))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_frame_set_label(arg0, arg1)
}

// SetLabelAlign sets the alignment of the frame widget’s label. The default
// values for a newly created frame are 0.0 and 0.5.
func (f frame) SetLabelAlign(xalign float32, yalign float32) {
	var arg0 *C.GtkFrame
	var arg1 C.gfloat
	var arg2 C.gfloat

	arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))
	arg1 = C.gfloat(xalign)
	arg2 = C.gfloat(yalign)

	C.gtk_frame_set_label_align(arg0, arg1, arg2)
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

// SetShadowType sets the Frame:shadow-type for @frame, i.e. whether it is
// drawn without (GTK_SHADOW_NONE) or with (other values) a visible border.
// Values other than GTK_SHADOW_NONE are treated identically by GtkFrame.
// The chosen type is applied by removing or adding the .flat class to the
// CSS node named border.
func (f frame) SetShadowType(typ ShadowType) {
	var arg0 *C.GtkFrame
	var arg1 C.GtkShadowType

	arg0 = (*C.GtkFrame)(unsafe.Pointer(f.Native()))
	arg1 = (C.GtkShadowType)(typ)

	C.gtk_frame_set_shadow_type(arg0, arg1)
}
