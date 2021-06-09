// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_drop_target_get_type()), F: marshalDropTarget},
	})
}

// DropTarget: gtkDropTarget is an event controller implementing a simple way to
// receive Drag-and-Drop operations.
//
// The most basic way to use a DropTarget to receive drops on a widget is to
// create it via gtk_drop_target_new() passing in the #GType of the data you
// want to receive and connect to the DropTarget::drop signal to receive the
// data:
//
//    static gboolean
//    on_drop (GtkDropTarget *target,
//             const GValue  *value,
//             double         x,
//             double         y,
//             gpointer       data)
//    {
//      MyWidget *self = data;
//
//      // Call the appropriate setter depending on the type of data
//      // that we received
//      if (G_VALUE_HOLDS (value, G_TYPE_FILE))
//        my_widget_set_file (self, g_value_get_object (value));
//      else if (G_VALUE_HOLDS (value, GDK_TYPE_PIXBUF))
//        my_widget_set_pixbuf (self, g_value_get_object (value));
//      else
//        return FALSE;
//
//      return TRUE;
//    }
//
//    static void
//    my_widget_init (MyWidget *self)
//    {
//      GtkDropTarget *target =
//        gtk_drop_target_new (G_TYPE_INVALID, GDK_ACTION_COPY);
//
//      // This widget accepts two types of drop types: GFile objects
//      // and GdkPixbuf objects
//      gtk_drop_target_set_gtypes (target, (GTypes [2]) {
//        G_TYPE_FILE,
//        GDK_TYPE_PIXBUF,
//      }, 2);
//
//      gtk_widget_add_controller (GTK_WIDGET (self), GTK_EVENT_CONTROLLER (target));
//    }
//
// DropTarget supports more options, such as:
//
//    * rejecting potential drops via the DropTarget::accept signal
//      and the gtk_drop_target_reject() function to let other drop
//      targets handle the drop
//    * tracking an ongoing drag operation before the drop via the
//      DropTarget::enter, DropTarget::motion and
//      DropTarget::leave signals
//    * configuring how to receive data by setting the
//      DropTarget:preload property and listening for its availability
//      via the DropTarget:value property
//
// However, DropTarget is ultimately modeled in a synchronous way and only
// supports data transferred via #GType. If you want full control over an
// ongoing drop, the DropTargetAsync object gives you this ability.
//
// While a pointer is dragged over the drop target's widget and the drop has not
// been rejected, that widget will receive the GTK_STATE_FLAG_DROP_ACTIVE state,
// which can be used to style the widget.
type DropTarget interface {
	EventController

	// Actions gets the actions that this drop target supports.
	Actions() gdk.DragAction
	// Drop gets the currently handled drop operation.
	//
	// If no drop operation is going on, nil is returned.
	Drop() gdk.Drop
	// Formats gets the data formats that this drop target accepts.
	//
	// If the result is nil, all formats are expected to be supported.
	Formats() *gdk.ContentFormats
	// GTypes gets the list of supported #GTypes for @self. If no type have been
	// set, nil will be returned.
	GTypes() []externglib.Type
	// Preload gets the value of the GtkDropTarget:preload property.
	Preload() bool
	// Value gets the value of the GtkDropTarget:value property.
	Value() **externglib.Value
	// Reject rejects the ongoing drop operation.
	//
	// If no drop operation is ongoing - when GdkDropTarget:drop returns nil -
	// this function does nothing.
	//
	// This function should be used when delaying the decision on whether to
	// accept a drag or not until after reading the data.
	Reject()
	// SetActions sets the actions that this drop target supports.
	SetActions(actions gdk.DragAction)
	// SetGTypes sets the supported #GTypes for this drop target.
	SetGTypes()
	// SetPreload sets the GtkDropTarget:preload property.
	SetPreload(preload bool)
}

// dropTarget implements the DropTarget interface.
type dropTarget struct {
	EventController
}

var _ DropTarget = (*dropTarget)(nil)

// WrapDropTarget wraps a GObject to the right type. It is
// primarily used internally.
func WrapDropTarget(obj *externglib.Object) DropTarget {
	return DropTarget{
		EventController: WrapEventController(obj),
	}
}

func marshalDropTarget(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDropTarget(obj), nil
}

// NewDropTarget constructs a class DropTarget.
func NewDropTarget(typ externglib.Type, actions gdk.DragAction) DropTarget {
	var arg1 C.GType
	var arg2 C.GdkDragAction

	arg1 = C.GType(typ)
	arg2 = (C.GdkDragAction)(actions)

	var cret C.GtkDropTarget

	cret = C.gtk_drop_target_new(arg1, arg2)

	var dropTarget DropTarget

	dropTarget = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(DropTarget)

	return dropTarget
}

// Actions gets the actions that this drop target supports.
func (s dropTarget) Actions() gdk.DragAction {
	var arg0 *C.GtkDropTarget

	arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))

	var cret C.GdkDragAction

	cret = C.gtk_drop_target_get_actions(arg0)

	var dragAction gdk.DragAction

	dragAction = gdk.DragAction(cret)

	return dragAction
}

// Drop gets the currently handled drop operation.
//
// If no drop operation is going on, nil is returned.
func (s dropTarget) Drop() gdk.Drop {
	var arg0 *C.GtkDropTarget

	arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))

	var cret *C.GdkDrop

	cret = C.gtk_drop_target_get_drop(arg0)

	var drop gdk.Drop

	drop = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gdk.Drop)

	return drop
}

// Formats gets the data formats that this drop target accepts.
//
// If the result is nil, all formats are expected to be supported.
func (s dropTarget) Formats() *gdk.ContentFormats {
	var arg0 *C.GtkDropTarget

	arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))

	var cret *C.GdkContentFormats

	cret = C.gtk_drop_target_get_formats(arg0)

	var contentFormats *gdk.ContentFormats

	contentFormats = gdk.WrapContentFormats(unsafe.Pointer(cret))
	runtime.SetFinalizer(contentFormats, func(v *gdk.ContentFormats) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return contentFormats
}

// GTypes gets the list of supported #GTypes for @self. If no type have been
// set, nil will be returned.
func (s dropTarget) GTypes() []externglib.Type {
	var arg0 *C.GtkDropTarget

	arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))

	var cret *C.GType
	var arg1 *C.gsize

	cret = C.gtk_drop_target_get_gtypes(arg0)

	var gTypes []externglib.Type

	{
		var src []C.GType
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(cret), int(arg1))

		gTypes = make([]externglib.Type, arg1)
		for i := 0; i < uintptr(arg1); i++ {
			gTypes = externglib.Type(cret)
		}
	}

	return gTypes
}

// Preload gets the value of the GtkDropTarget:preload property.
func (s dropTarget) Preload() bool {
	var arg0 *C.GtkDropTarget

	arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))

	var cret C.gboolean

	cret = C.gtk_drop_target_get_preload(arg0)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// Value gets the value of the GtkDropTarget:value property.
func (s dropTarget) Value() **externglib.Value {
	var arg0 *C.GtkDropTarget

	arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))

	var cret *C.GValue

	cret = C.gtk_drop_target_get_value(arg0)

	var value **externglib.Value

	value = externglib.ValueFromNative(unsafe.Pointer(cret))

	return value
}

// Reject rejects the ongoing drop operation.
//
// If no drop operation is ongoing - when GdkDropTarget:drop returns nil -
// this function does nothing.
//
// This function should be used when delaying the decision on whether to
// accept a drag or not until after reading the data.
func (s dropTarget) Reject() {
	var arg0 *C.GtkDropTarget

	arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))

	C.gtk_drop_target_reject(arg0)
}

// SetActions sets the actions that this drop target supports.
func (s dropTarget) SetActions(actions gdk.DragAction) {
	var arg0 *C.GtkDropTarget
	var arg1 C.GdkDragAction

	arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))
	arg1 = (C.GdkDragAction)(actions)

	C.gtk_drop_target_set_actions(arg0, arg1)
}

// SetGTypes sets the supported #GTypes for this drop target.
func (s dropTarget) SetGTypes() {
	var arg0 *C.GtkDropTarget

	arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))

	C.gtk_drop_target_set_gtypes(arg0)
}

// SetPreload sets the GtkDropTarget:preload property.
func (s dropTarget) SetPreload(preload bool) {
	var arg0 *C.GtkDropTarget
	var arg1 C.gboolean

	arg0 = (*C.GtkDropTarget)(unsafe.Pointer(s.Native()))
	if preload {
		arg1 = C.gboolean(1)
	}

	C.gtk_drop_target_set_preload(arg0, arg1)
}
