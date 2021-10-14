// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_drop_target_async_get_type()), F: marshalDropTargetAsyncer},
	})
}

// DropTargetAsync: GtkDropTargetAsync is an event controller to receive
// Drag-and-Drop operations, asynchronously.
//
// It is the more complete but also more complex method of handling drop
// operations compared to gtk.DropTarget, and you should only use it if
// GtkDropTarget doesn't provide all the features you need.
//
// To use a GtkDropTargetAsync to receive drops on a widget, you create a
// GtkDropTargetAsync object, configure which data formats and actions you
// support, connect to its signals, and then attach it to the widget with
// gtk.Widget.AddController().
//
// During a drag operation, the first signal that a GtkDropTargetAsync emits is
// gtk.DropTargetAsync::accept, which is meant to determine whether the target
// is a possible drop site for the ongoing drop. The default handler for the
// ::accept signal accepts the drop if it finds a compatible data format and an
// action that is supported on both sides.
//
// If it is, and the widget becomes a target, you will receive a
// gtk.DropTargetAsync::drag-enter signal, followed by
// gtk.DropTargetAsync::drag-motion signals as the pointer moves, optionally a
// gtk.DropTargetAsync::drop signal when a drop happens, and finally a
// gtk.DropTargetAsync::drag-leave signal when the pointer moves off the widget.
//
// The ::drag-enter and ::drag-motion handler return a GdkDragAction to update
// the status of the ongoing operation. The ::drop handler should decide if it
// ultimately accepts the drop and if it does, it should initiate the data
// transfer and finish the operation by calling gdk.Drop.Finish().
//
// Between the ::drag-enter and ::drag-leave signals the widget is a current
// drop target, and will receive the GTK_STATE_FLAG_DROP_ACTIVE state, which can
// be used by themes to style the widget as a drop target.
type DropTargetAsync struct {
	EventController
}

func wrapDropTargetAsync(obj *externglib.Object) *DropTargetAsync {
	return &DropTargetAsync{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalDropTargetAsyncer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapDropTargetAsync(obj), nil
}

// NewDropTargetAsync creates a new GtkDropTargetAsync object.
//
// The function takes the following parameters:
//
//    - formats: supported data formats.
//    - actions: supported actions.
//
func NewDropTargetAsync(formats *gdk.ContentFormats, actions gdk.DragAction) *DropTargetAsync {
	var _arg1 *C.GdkContentFormats  // out
	var _arg2 C.GdkDragAction       // out
	var _cret *C.GtkDropTargetAsync // in

	if formats != nil {
		_arg1 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))
	}
	_arg2 = C.GdkDragAction(actions)

	_cret = C.gtk_drop_target_async_new(_arg1, _arg2)
	runtime.KeepAlive(formats)
	runtime.KeepAlive(actions)

	var _dropTargetAsync *DropTargetAsync // out

	_dropTargetAsync = wrapDropTargetAsync(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _dropTargetAsync
}

// Actions gets the actions that this drop target supports.
func (self *DropTargetAsync) Actions() gdk.DragAction {
	var _arg0 *C.GtkDropTargetAsync // out
	var _cret C.GdkDragAction       // in

	_arg0 = (*C.GtkDropTargetAsync)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drop_target_async_get_actions(_arg0)
	runtime.KeepAlive(self)

	var _dragAction gdk.DragAction // out

	_dragAction = gdk.DragAction(_cret)

	return _dragAction
}

// Formats gets the data formats that this drop target accepts.
//
// If the result is NULL, all formats are expected to be supported.
func (self *DropTargetAsync) Formats() *gdk.ContentFormats {
	var _arg0 *C.GtkDropTargetAsync // out
	var _cret *C.GdkContentFormats  // in

	_arg0 = (*C.GtkDropTargetAsync)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_drop_target_async_get_formats(_arg0)
	runtime.KeepAlive(self)

	var _contentFormats *gdk.ContentFormats // out

	if _cret != nil {
		_contentFormats = (*gdk.ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_contentFormats)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gdk_content_formats_unref((*C.GdkContentFormats)(intern.C))
			},
		)
	}

	return _contentFormats
}

// RejectDrop sets the drop as not accepted on this drag site.
//
// This function should be used when delaying the decision on whether to accept
// a drag or not until after reading the data.
//
// The function takes the following parameters:
//
//    - drop of an ongoing drag operation.
//
func (self *DropTargetAsync) RejectDrop(drop gdk.Dropper) {
	var _arg0 *C.GtkDropTargetAsync // out
	var _arg1 *C.GdkDrop            // out

	_arg0 = (*C.GtkDropTargetAsync)(unsafe.Pointer(self.Native()))
	_arg1 = (*C.GdkDrop)(unsafe.Pointer(drop.Native()))

	C.gtk_drop_target_async_reject_drop(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(drop)
}

// SetActions sets the actions that this drop target supports.
//
// The function takes the following parameters:
//
//    - actions: supported actions.
//
func (self *DropTargetAsync) SetActions(actions gdk.DragAction) {
	var _arg0 *C.GtkDropTargetAsync // out
	var _arg1 C.GdkDragAction       // out

	_arg0 = (*C.GtkDropTargetAsync)(unsafe.Pointer(self.Native()))
	_arg1 = C.GdkDragAction(actions)

	C.gtk_drop_target_async_set_actions(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(actions)
}

// SetFormats sets the data formats that this drop target will accept.
//
// The function takes the following parameters:
//
//    - formats: supported data formats or NULL for any format.
//
func (self *DropTargetAsync) SetFormats(formats *gdk.ContentFormats) {
	var _arg0 *C.GtkDropTargetAsync // out
	var _arg1 *C.GdkContentFormats  // out

	_arg0 = (*C.GtkDropTargetAsync)(unsafe.Pointer(self.Native()))
	if formats != nil {
		_arg1 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))
	}

	C.gtk_drop_target_async_set_formats(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(formats)
}

// ConnectAccept: emitted on the drop site when a drop operation is about to
// begin.
//
// If the drop is not accepted, FALSE will be returned and the drop target will
// ignore the drop. If TRUE is returned, the drop is accepted for now but may be
// rejected later via a call to gtk.DropTargetAsync.RejectDrop() or ultimately
// by returning FALSE from a gtk.DropTargetAsync::drop handler.
//
// The default handler for this signal decides whether to accept the drop based
// on the formats provided by the drop.
//
// If the decision whether the drop will be accepted or rejected needs further
// processing, such as inspecting the data, this function should return TRUE and
// proceed as is drop was accepted and if it decides to reject the drop later,
// it should call gtk.DropTargetAsync.RejectDrop().
func (self *DropTargetAsync) ConnectAccept(f func(drop gdk.Dropper) bool) externglib.SignalHandle {
	return self.Connect("accept", f)
}

// ConnectDragEnter: emitted on the drop site when the pointer enters the
// widget.
//
// It can be used to set up custom highlighting.
func (self *DropTargetAsync) ConnectDragEnter(f func(drop gdk.Dropper, x, y float64) gdk.DragAction) externglib.SignalHandle {
	return self.Connect("drag-enter", f)
}

// ConnectDragLeave: emitted on the drop site when the pointer leaves the
// widget.
//
// Its main purpose it to undo things done in GtkDropTargetAsync::drag-enter.
func (self *DropTargetAsync) ConnectDragLeave(f func(drop gdk.Dropper)) externglib.SignalHandle {
	return self.Connect("drag-leave", f)
}

// ConnectDragMotion: emitted while the pointer is moving over the drop target.
func (self *DropTargetAsync) ConnectDragMotion(f func(drop gdk.Dropper, x, y float64) gdk.DragAction) externglib.SignalHandle {
	return self.Connect("drag-motion", f)
}

// ConnectDrop: emitted on the drop site when the user drops the data onto the
// widget.
//
// The signal handler must determine whether the pointer position is in a drop
// zone or not. If it is not in a drop zone, it returns FALSE and no further
// processing is necessary.
//
// Otherwise, the handler returns TRUE. In this case, this handler will accept
// the drop. The handler must ensure that gdk.Drop.Finish() is called to let the
// source know that the drop is done. The call to gdk.Drop.Finish() must only be
// done when all data has been received.
//
// To receive the data, use one of the read functions provided by gdk.Drop such
// as gdk.Drop.ReadAsync() or gdk.Drop.ReadValueAsync().
func (self *DropTargetAsync) ConnectDrop(f func(drop gdk.Dropper, x, y float64) bool) externglib.SignalHandle {
	return self.Connect("drop", f)
}
