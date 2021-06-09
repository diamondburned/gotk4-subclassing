// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// DragAbort aborts a drag without dropping.
//
// This function is called by the drag source.
//
// This function does not need to be called in managed drag and drop operations.
// See gdk_drag_context_manage_dnd() for more information.
func DragAbort(context DragContext, time_ uint32) {
	var arg1 *C.GdkDragContext
	var arg2 C.guint32

	arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	arg2 = C.guint32(time_)

	C.gdk_drag_abort(arg1, arg2)
}

// DragBegin starts a drag and creates a new drag context for it. This function
// assumes that the drag is controlled by the client pointer device, use
// gdk_drag_begin_for_device() to begin a drag with a different device.
//
// This function is called by the drag source.
func DragBegin(window Window, targets *glib.List) DragContext {
	var arg1 *C.GdkWindow
	var arg2 *C.GList

	arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	arg2 = (*C.GList)(unsafe.Pointer(targets.Native()))

	var cret *C.GdkDragContext

	cret = C.gdk_drag_begin(arg1, arg2)

	var dragContext DragContext

	dragContext = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(DragContext)

	return dragContext
}

// DragBeginForDevice starts a drag and creates a new drag context for it.
//
// This function is called by the drag source.
func DragBeginForDevice(window Window, device Device, targets *glib.List) DragContext {
	var arg1 *C.GdkWindow
	var arg2 *C.GdkDevice
	var arg3 *C.GList

	arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	arg2 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	arg3 = (*C.GList)(unsafe.Pointer(targets.Native()))

	var cret *C.GdkDragContext

	cret = C.gdk_drag_begin_for_device(arg1, arg2, arg3)

	var dragContext DragContext

	dragContext = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(DragContext)

	return dragContext
}

// DragBeginFromPoint starts a drag and creates a new drag context for it.
//
// This function is called by the drag source.
func DragBeginFromPoint(window Window, device Device, targets *glib.List, xRoot int, yRoot int) DragContext {
	var arg1 *C.GdkWindow
	var arg2 *C.GdkDevice
	var arg3 *C.GList
	var arg4 C.gint
	var arg5 C.gint

	arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	arg2 = (*C.GdkDevice)(unsafe.Pointer(device.Native()))
	arg3 = (*C.GList)(unsafe.Pointer(targets.Native()))
	arg4 = C.gint(xRoot)
	arg5 = C.gint(yRoot)

	var cret *C.GdkDragContext

	cret = C.gdk_drag_begin_from_point(arg1, arg2, arg3, arg4, arg5)

	var dragContext DragContext

	dragContext = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(DragContext)

	return dragContext
}

// DragDrop drops on the current destination.
//
// This function is called by the drag source.
//
// This function does not need to be called in managed drag and drop operations.
// See gdk_drag_context_manage_dnd() for more information.
func DragDrop(context DragContext, time_ uint32) {
	var arg1 *C.GdkDragContext
	var arg2 C.guint32

	arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	arg2 = C.guint32(time_)

	C.gdk_drag_drop(arg1, arg2)
}

// DragDropDone: inform GDK if the drop ended successfully. Passing false for
// @success may trigger a drag cancellation animation.
//
// This function is called by the drag source, and should be the last call
// before dropping the reference to the @context.
//
// The DragContext will only take the first gdk_drag_drop_done() call as
// effective, if this function is called multiple times, all subsequent calls
// will be ignored.
func DragDropDone(context DragContext, success bool) {
	var arg1 *C.GdkDragContext
	var arg2 C.gboolean

	arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	if success {
		arg2 = C.gboolean(1)
	}

	C.gdk_drag_drop_done(arg1, arg2)
}

// DragDropSucceeded returns whether the dropped data has been successfully
// transferred. This function is intended to be used while handling a
// GDK_DROP_FINISHED event, its return value is meaningless at other times.
func DragDropSucceeded(context DragContext) bool {
	var arg1 *C.GdkDragContext

	arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	var cret C.gboolean

	cret = C.gdk_drag_drop_succeeded(arg1)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// DragFindWindowForScreen finds the destination window and DND protocol to use
// at the given pointer position.
//
// This function is called by the drag source to obtain the @dest_window and
// @protocol parameters for gdk_drag_motion().
func DragFindWindowForScreen(context DragContext, dragWindow Window, screen Screen, xRoot int, yRoot int) (destWindow Window, protocol DragProtocol) {
	var arg1 *C.GdkDragContext
	var arg2 *C.GdkWindow
	var arg3 *C.GdkScreen
	var arg4 C.gint
	var arg5 C.gint

	arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	arg2 = (*C.GdkWindow)(unsafe.Pointer(dragWindow.Native()))
	arg3 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))
	arg4 = C.gint(xRoot)
	arg5 = C.gint(yRoot)

	var arg6 *C.GdkWindow
	var arg7 C.GdkDragProtocol

	C.gdk_drag_find_window_for_screen(arg1, arg2, arg3, arg4, arg5, &arg6, &arg7)

	var destWindow Window
	var protocol DragProtocol

	destWindow = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(arg6.Native()))).(Window)
	protocol = DragProtocol(arg7)

	return destWindow, protocol
}

// DragGetSelection returns the selection atom for the current source window.
func DragGetSelection(context DragContext) Atom {
	var arg1 *C.GdkDragContext

	arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))

	var cret C.GdkAtom

	cret = C.gdk_drag_get_selection(arg1)

	var atom Atom

	atom = *WrapAtom(unsafe.Pointer(&cret))

	return atom
}

// DragMotion updates the drag context when the pointer moves or the set of
// actions changes.
//
// This function is called by the drag source.
//
// This function does not need to be called in managed drag and drop operations.
// See gdk_drag_context_manage_dnd() for more information.
func DragMotion(context DragContext, destWindow Window, protocol DragProtocol, xRoot int, yRoot int, suggestedAction DragAction, possibleActions DragAction, time_ uint32) bool {
	var arg1 *C.GdkDragContext
	var arg2 *C.GdkWindow
	var arg3 C.GdkDragProtocol
	var arg4 C.gint
	var arg5 C.gint
	var arg6 C.GdkDragAction
	var arg7 C.GdkDragAction
	var arg8 C.guint32

	arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	arg2 = (*C.GdkWindow)(unsafe.Pointer(destWindow.Native()))
	arg3 = (C.GdkDragProtocol)(protocol)
	arg4 = C.gint(xRoot)
	arg5 = C.gint(yRoot)
	arg6 = (C.GdkDragAction)(suggestedAction)
	arg7 = (C.GdkDragAction)(possibleActions)
	arg8 = C.guint32(time_)

	var cret C.gboolean

	cret = C.gdk_drag_motion(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// DragStatus selects one of the actions offered by the drag source.
//
// This function is called by the drag destination in response to
// gdk_drag_motion() called by the drag source.
func DragStatus(context DragContext, action DragAction, time_ uint32) {
	var arg1 *C.GdkDragContext
	var arg2 C.GdkDragAction
	var arg3 C.guint32

	arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	arg2 = (C.GdkDragAction)(action)
	arg3 = C.guint32(time_)

	C.gdk_drag_status(arg1, arg2, arg3)
}

// DropFinish ends the drag operation after a drop.
//
// This function is called by the drag destination.
func DropFinish(context DragContext, success bool, time_ uint32) {
	var arg1 *C.GdkDragContext
	var arg2 C.gboolean
	var arg3 C.guint32

	arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	if success {
		arg2 = C.gboolean(1)
	}
	arg3 = C.guint32(time_)

	C.gdk_drop_finish(arg1, arg2, arg3)
}

// DropReply accepts or rejects a drop.
//
// This function is called by the drag destination in response to a drop
// initiated by the drag source.
func DropReply(context DragContext, accepted bool, time_ uint32) {
	var arg1 *C.GdkDragContext
	var arg2 C.gboolean
	var arg3 C.guint32

	arg1 = (*C.GdkDragContext)(unsafe.Pointer(context.Native()))
	if accepted {
		arg2 = C.gboolean(1)
	}
	arg3 = C.guint32(time_)

	C.gdk_drop_reply(arg1, arg2, arg3)
}
