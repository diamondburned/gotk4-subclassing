// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/internal/box"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_flow_box_child_get_type()), F: marshalFlowBoxChild},
	})
}

// FlowBoxCreateWidgetFunc: called for flow boxes that are bound to a Model with
// gtk_flow_box_bind_model() for each item that gets added to the model.
type FlowBoxCreateWidgetFunc func(item gextras.Objector) Widget

//export gotk4_FlowBoxCreateWidgetFunc
func gotk4_FlowBoxCreateWidgetFunc(arg0 C.gpointer, arg1 C.gpointer) *C.GtkWidget {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(FlowBoxCreateWidgetFunc)
	ret := fn(item, userData)

	cret = (*C.GtkWidget)(unsafe.Pointer(ret.Native()))

	return cret
}

// FlowBoxFilterFunc: a function that will be called whenever a child changes or
// is added. It lets you control if the child should be visible or not.
type FlowBoxFilterFunc func(child FlowBoxChild) bool

//export gotk4_FlowBoxFilterFunc
func gotk4_FlowBoxFilterFunc(arg0 *C.GtkFlowBoxChild, arg1 C.gpointer) C.gboolean {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(FlowBoxFilterFunc)
	ret := fn(child, userData)

	if ret {
		cret = C.gboolean(1)
	}

	return cret
}

// FlowBoxForeachFunc: a function used by gtk_flow_box_selected_foreach(). It
// will be called on every selected child of the @box.
type FlowBoxForeachFunc func(box FlowBox, child FlowBoxChild)

//export gotk4_FlowBoxForeachFunc
func gotk4_FlowBoxForeachFunc(arg0 *C.GtkFlowBox, arg1 *C.GtkFlowBoxChild, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(FlowBoxForeachFunc)
	fn(box, child, userData)
}

// FlowBoxSortFunc: a function to compare two children to determine which should
// come first.
type FlowBoxSortFunc func(child1 FlowBoxChild, child2 FlowBoxChild) int

//export gotk4_FlowBoxSortFunc
func gotk4_FlowBoxSortFunc(arg0 *C.GtkFlowBoxChild, arg1 *C.GtkFlowBoxChild, arg2 C.gpointer) C.int {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(FlowBoxSortFunc)
	ret := fn(child1, child2, userData)

	cret = C.int(ret)

	return cret
}

type FlowBoxChild interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget

	// Changed marks @child as changed, causing any state that depends on this
	// to be updated. This affects sorting and filtering.
	//
	// Note that calls to this method must be in sync with the data used for the
	// sorting and filtering functions. For instance, if the list is mirroring
	// some external data set, and *two* children changed in the external data
	// set when you call gtk_flow_box_child_changed() on the first child, the
	// sort function must only read the new data for the first of the two
	// changed children, otherwise the resorting of the children will be wrong.
	//
	// This generally means that if you don’t fully control the data model, you
	// have to duplicate the data that affects the sorting and filtering
	// functions into the widgets themselves. Another alternative is to call
	// gtk_flow_box_invalidate_sort() on any model change, but that is more
	// expensive.
	Changed(c FlowBoxChild)
	// Child gets the child widget of @self.
	Child(s FlowBoxChild)
	// Index gets the current index of the @child in its FlowBox container.
	Index(c FlowBoxChild)
	// IsSelected returns whether the @child is currently selected in its
	// FlowBox container.
	IsSelected(c FlowBoxChild) bool
	// SetChild sets the child widget of @self.
	SetChild(s FlowBoxChild, child Widget)
}

// flowBoxChild implements the FlowBoxChild interface.
type flowBoxChild struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
}

var _ FlowBoxChild = (*flowBoxChild)(nil)

// WrapFlowBoxChild wraps a GObject to the right type. It is
// primarily used internally.
func WrapFlowBoxChild(obj *externglib.Object) FlowBoxChild {
	return FlowBoxChild{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalFlowBoxChild(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFlowBoxChild(obj), nil
}

// NewFlowBoxChild constructs a class FlowBoxChild.
func NewFlowBoxChild() {
	C.gtk_flow_box_child_new()
}

// Changed marks @child as changed, causing any state that depends on this
// to be updated. This affects sorting and filtering.
//
// Note that calls to this method must be in sync with the data used for the
// sorting and filtering functions. For instance, if the list is mirroring
// some external data set, and *two* children changed in the external data
// set when you call gtk_flow_box_child_changed() on the first child, the
// sort function must only read the new data for the first of the two
// changed children, otherwise the resorting of the children will be wrong.
//
// This generally means that if you don’t fully control the data model, you
// have to duplicate the data that affects the sorting and filtering
// functions into the widgets themselves. Another alternative is to call
// gtk_flow_box_invalidate_sort() on any model change, but that is more
// expensive.
func (c flowBoxChild) Changed(c FlowBoxChild) {
	var arg0 *C.GtkFlowBoxChild

	arg0 = (*C.GtkFlowBoxChild)(unsafe.Pointer(c.Native()))

	C.gtk_flow_box_child_changed(arg0)
}

// Child gets the child widget of @self.
func (s flowBoxChild) Child(s FlowBoxChild) {
	var arg0 *C.GtkFlowBoxChild

	arg0 = (*C.GtkFlowBoxChild)(unsafe.Pointer(s.Native()))

	C.gtk_flow_box_child_get_child(arg0)
}

// Index gets the current index of the @child in its FlowBox container.
func (c flowBoxChild) Index(c FlowBoxChild) {
	var arg0 *C.GtkFlowBoxChild

	arg0 = (*C.GtkFlowBoxChild)(unsafe.Pointer(c.Native()))

	C.gtk_flow_box_child_get_index(arg0)
}

// IsSelected returns whether the @child is currently selected in its
// FlowBox container.
func (c flowBoxChild) IsSelected(c FlowBoxChild) bool {
	var arg0 *C.GtkFlowBoxChild

	arg0 = (*C.GtkFlowBoxChild)(unsafe.Pointer(c.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_flow_box_child_is_selected(arg0)

	if cret {
		ok = true
	}

	return ok
}

// SetChild sets the child widget of @self.
func (s flowBoxChild) SetChild(s FlowBoxChild, child Widget) {
	var arg0 *C.GtkFlowBoxChild
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkFlowBoxChild)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_flow_box_child_set_child(arg0, arg1)
}
