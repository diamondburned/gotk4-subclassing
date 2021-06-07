// Code generated by girgen. DO NOT EDIT.

package gio

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_simple_async_result_get_type()), F: marshalSimpleAsyncResult},
	})
}

// SimpleAsyncResult as of GLib 2.46, AsyncResult is deprecated in favor of
// #GTask, which provides a simpler API.
//
// AsyncResult implements Result.
//
// GSimpleAsyncResult handles ReadyCallbacks, error reporting, operation
// cancellation and the final state of an operation, completely transparent to
// the application. Results can be returned as a pointer e.g. for functions that
// return data that is collected asynchronously, a boolean value for checking
// the success or failure of an operation, or a #gssize for operations which
// return the number of bytes modified by the operation; all of the simple
// return cases are covered.
//
// Most of the time, an application will not need to know of the details of this
// API; it is handled transparently, and any necessary operations are handled by
// Result's interface. However, if implementing a new GIO module, for writing
// language bindings, or for complex applications that need better control of
// how asynchronous operations are completed, it is important to understand this
// functionality.
//
// GSimpleAsyncResults are tagged with the calling function to ensure that
// asynchronous functions and their finishing functions are used together
// correctly.
//
// To create a new AsyncResult, call g_simple_async_result_new(). If the result
// needs to be created for a #GError, use g_simple_async_result_new_from_error()
// or g_simple_async_result_new_take_error(). If a #GError is not available
// (e.g. the asynchronous operation's doesn't take a #GError argument), but the
// result still needs to be created for an error condition, use
// g_simple_async_result_new_error() (or g_simple_async_result_set_error_va() if
// your application or binding requires passing a variable argument list
// directly), and the error can then be propagated through the use of
// g_simple_async_result_propagate_error().
//
// An asynchronous operation can be made to ignore a cancellation event by
// calling g_simple_async_result_set_handle_cancellation() with a AsyncResult
// for the operation and false. This is useful for operations that are dangerous
// to cancel, such as close (which would cause a leak if cancelled before being
// run).
//
// GSimpleAsyncResult can integrate into GLib's event loop, Loop, or it can use
// #GThreads. g_simple_async_result_complete() will finish an I/O task directly
// from the point where it is called. g_simple_async_result_complete_in_idle()
// will finish it from an idle handler in the [thread-default main
// context][g-main-context-push-thread-default] where the AsyncResult was
// created. g_simple_async_result_run_in_thread() will run the job in a separate
// thread and then use g_simple_async_result_complete_in_idle() to deliver the
// result.
//
// To set the results of an asynchronous function,
// g_simple_async_result_set_op_res_gpointer(),
// g_simple_async_result_set_op_res_gboolean(), and
// g_simple_async_result_set_op_res_gssize() are provided, setting the
// operation's result to a gpointer, gboolean, or gssize, respectively.
//
// Likewise, to get the result of an asynchronous function,
// g_simple_async_result_get_op_res_gpointer(),
// g_simple_async_result_get_op_res_gboolean(), and
// g_simple_async_result_get_op_res_gssize() are provided, getting the
// operation's result as a gpointer, gboolean, and gssize, respectively.
//
// For the details of the requirements implementations must respect, see Result.
// A typical implementation of an asynchronous operation using
// GSimpleAsyncResult looks something like this:
//
//    static void
//    baked_cb (Cake    *cake,
//              gpointer user_data)
//    {
//      // In this example, this callback is not given a reference to the cake,
//      // so the GSimpleAsyncResult has to take a reference to it.
//      GSimpleAsyncResult *result = user_data;
//
//      if (cake == NULL)
//        g_simple_async_result_set_error (result,
//                                         BAKER_ERRORS,
//                                         BAKER_ERROR_NO_FLOUR,
//                                         "Go to the supermarket");
//      else
//        g_simple_async_result_set_op_res_gpointer (result,
//                                                   g_object_ref (cake),
//                                                   g_object_unref);
//
//
//      // In this example, we assume that baked_cb is called as a callback from
//      // the mainloop, so it's safe to complete the operation synchronously here.
//      // If, however, _baker_prepare_cake () might call its callback without
//      // first returning to the mainloop — inadvisable, but some APIs do so —
//      // we would need to use g_simple_async_result_complete_in_idle().
//      g_simple_async_result_complete (result);
//      g_object_unref (result);
//    }
//
//    void
//    baker_bake_cake_async (Baker              *self,
//                           guint               radius,
//                           GAsyncReadyCallback callback,
//                           gpointer            user_data)
//    {
//      GSimpleAsyncResult *simple;
//      Cake               *cake;
//
//      if (radius < 3)
//        {
//          g_simple_async_report_error_in_idle (G_OBJECT (self),
//                                               callback,
//                                               user_data,
//                                               BAKER_ERRORS,
//                                               BAKER_ERROR_TOO_SMALL,
//                                               "ucm radius cakes are silly",
//                                               radius);
//          return;
//        }
//
//      simple = g_simple_async_result_new (G_OBJECT (self),
//                                          callback,
//                                          user_data,
//                                          baker_bake_cake_async);
//      cake = _baker_get_cached_cake (self, radius);
//
//      if (cake != NULL)
//        {
//          g_simple_async_result_set_op_res_gpointer (simple,
//                                                     g_object_ref (cake),
//                                                     g_object_unref);
//          g_simple_async_result_complete_in_idle (simple);
//          g_object_unref (simple);
//          // Drop the reference returned by _baker_get_cached_cake();
//          // the GSimpleAsyncResult has taken its own reference.
//          g_object_unref (cake);
//          return;
//        }
//
//      _baker_prepare_cake (self, radius, baked_cb, simple);
//    }
//
//    Cake *
//    baker_bake_cake_finish (Baker        *self,
//                            GAsyncResult *result,
//                            GError      **error)
//    {
//      GSimpleAsyncResult *simple;
//      Cake               *cake;
//
//      g_return_val_if_fail (g_simple_async_result_is_valid (result,
//                                                            G_OBJECT (self),
//                                                            baker_bake_cake_async),
//                            NULL);
//
//      simple = (GSimpleAsyncResult *) result;
//
//      if (g_simple_async_result_propagate_error (simple, error))
//        return NULL;
//
//      cake = CAKE (g_simple_async_result_get_op_res_gpointer (simple));
//      return g_object_ref (cake);
//    }
type SimpleAsyncResult interface {
	gextras.Objector
	AsyncResult

	// Complete completes an asynchronous I/O job immediately. Must be called in
	// the thread where the asynchronous result was to be delivered, as it
	// invokes the callback directly. If you are in a different thread use
	// g_simple_async_result_complete_in_idle().
	//
	// Calling this function takes a reference to @simple for as long as is
	// needed to complete the call.
	Complete(s SimpleAsyncResult)
	// CompleteInIdle completes an asynchronous function in an idle handler in
	// the [thread-default main context][g-main-context-push-thread-default] of
	// the thread that @simple was initially created in (and re-pushes that
	// context around the invocation of the callback).
	//
	// Calling this function takes a reference to @simple for as long as is
	// needed to complete the call.
	CompleteInIdle(s SimpleAsyncResult)
	// OpResGboolean gets the operation result boolean from within the
	// asynchronous result.
	OpResGboolean(s SimpleAsyncResult) bool
	// OpResGpointer gets a pointer result as returned by the asynchronous
	// function.
	OpResGpointer(s SimpleAsyncResult)
	// OpResGssize gets a gssize from the asynchronous result.
	OpResGssize(s SimpleAsyncResult)
	// SourceTag gets the source tag for the AsyncResult.
	SourceTag(s SimpleAsyncResult)
	// PropagateError propagates an error from within the simple asynchronous
	// result to a given destination.
	//
	// If the #GCancellable given to a prior call to
	// g_simple_async_result_set_check_cancellable() is cancelled then this
	// function will return true with @dest set appropriately.
	PropagateError(s SimpleAsyncResult) error
	// SetCheckCancellable sets a #GCancellable to check before dispatching
	// results.
	//
	// This function has one very specific purpose: the provided cancellable is
	// checked at the time of g_simple_async_result_propagate_error() If it is
	// cancelled, these functions will return an "Operation was cancelled" error
	// (G_IO_ERROR_CANCELLED).
	//
	// Implementors of cancellable asynchronous functions should use this in
	// order to provide a guarantee to their callers that cancelling an async
	// operation will reliably result in an error being returned for that
	// operation (even if a positive result for the operation has already been
	// sent as an idle to the main context to be dispatched).
	//
	// The checking described above is done regardless of any call to the
	// unrelated g_simple_async_result_set_handle_cancellation() function.
	SetCheckCancellable(s SimpleAsyncResult, checkCancellable Cancellable)
	// SetFromError sets the result from a #GError.
	SetFromError(s SimpleAsyncResult, error error)
	// SetHandleCancellation sets whether to handle cancellation within the
	// asynchronous operation.
	//
	// This function has nothing to do with
	// g_simple_async_result_set_check_cancellable(). It only refers to the
	// #GCancellable passed to g_simple_async_result_run_in_thread().
	SetHandleCancellation(s SimpleAsyncResult, handleCancellation bool)
	// SetOpResGboolean sets the operation result to a boolean within the
	// asynchronous result.
	SetOpResGboolean(s SimpleAsyncResult, opRes bool)
	// SetOpResGssize sets the operation result within the asynchronous result
	// to the given @op_res.
	SetOpResGssize(s SimpleAsyncResult, opRes int)
	// TakeError sets the result from @error, and takes over the caller's
	// ownership of @error, so the caller does not need to free it any more.
	TakeError(s SimpleAsyncResult, error error)
}

// simpleAsyncResult implements the SimpleAsyncResult interface.
type simpleAsyncResult struct {
	gextras.Objector
	AsyncResult
}

var _ SimpleAsyncResult = (*simpleAsyncResult)(nil)

// WrapSimpleAsyncResult wraps a GObject to the right type. It is
// primarily used internally.
func WrapSimpleAsyncResult(obj *externglib.Object) SimpleAsyncResult {
	return SimpleAsyncResult{
		Objector:    obj,
		AsyncResult: WrapAsyncResult(obj),
	}
}

func marshalSimpleAsyncResult(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSimpleAsyncResult(obj), nil
}

// Complete completes an asynchronous I/O job immediately. Must be called in
// the thread where the asynchronous result was to be delivered, as it
// invokes the callback directly. If you are in a different thread use
// g_simple_async_result_complete_in_idle().
//
// Calling this function takes a reference to @simple for as long as is
// needed to complete the call.
func (s simpleAsyncResult) Complete(s SimpleAsyncResult) {
	var arg0 *C.GSimpleAsyncResult

	arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(s.Native()))

	C.g_simple_async_result_complete(arg0)
}

// CompleteInIdle completes an asynchronous function in an idle handler in
// the [thread-default main context][g-main-context-push-thread-default] of
// the thread that @simple was initially created in (and re-pushes that
// context around the invocation of the callback).
//
// Calling this function takes a reference to @simple for as long as is
// needed to complete the call.
func (s simpleAsyncResult) CompleteInIdle(s SimpleAsyncResult) {
	var arg0 *C.GSimpleAsyncResult

	arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(s.Native()))

	C.g_simple_async_result_complete_in_idle(arg0)
}

// OpResGboolean gets the operation result boolean from within the
// asynchronous result.
func (s simpleAsyncResult) OpResGboolean(s SimpleAsyncResult) bool {
	var arg0 *C.GSimpleAsyncResult

	arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_simple_async_result_get_op_res_gboolean(arg0)

	if cret {
		ok = true
	}

	return ok
}

// OpResGpointer gets a pointer result as returned by the asynchronous
// function.
func (s simpleAsyncResult) OpResGpointer(s SimpleAsyncResult) {
	var arg0 *C.GSimpleAsyncResult

	arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(s.Native()))

	C.g_simple_async_result_get_op_res_gpointer(arg0)
}

// OpResGssize gets a gssize from the asynchronous result.
func (s simpleAsyncResult) OpResGssize(s SimpleAsyncResult) {
	var arg0 *C.GSimpleAsyncResult

	arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(s.Native()))

	C.g_simple_async_result_get_op_res_gssize(arg0)
}

// SourceTag gets the source tag for the AsyncResult.
func (s simpleAsyncResult) SourceTag(s SimpleAsyncResult) {
	var arg0 *C.GSimpleAsyncResult

	arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(s.Native()))

	C.g_simple_async_result_get_source_tag(arg0)
}

// PropagateError propagates an error from within the simple asynchronous
// result to a given destination.
//
// If the #GCancellable given to a prior call to
// g_simple_async_result_set_check_cancellable() is cancelled then this
// function will return true with @dest set appropriately.
func (s simpleAsyncResult) PropagateError(s SimpleAsyncResult) error {
	var arg0 *C.GSimpleAsyncResult

	arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(s.Native()))

	var errout *C.GError
	var err error

	C.g_simple_async_result_propagate_error(arg0, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// SetCheckCancellable sets a #GCancellable to check before dispatching
// results.
//
// This function has one very specific purpose: the provided cancellable is
// checked at the time of g_simple_async_result_propagate_error() If it is
// cancelled, these functions will return an "Operation was cancelled" error
// (G_IO_ERROR_CANCELLED).
//
// Implementors of cancellable asynchronous functions should use this in
// order to provide a guarantee to their callers that cancelling an async
// operation will reliably result in an error being returned for that
// operation (even if a positive result for the operation has already been
// sent as an idle to the main context to be dispatched).
//
// The checking described above is done regardless of any call to the
// unrelated g_simple_async_result_set_handle_cancellation() function.
func (s simpleAsyncResult) SetCheckCancellable(s SimpleAsyncResult, checkCancellable Cancellable) {
	var arg0 *C.GSimpleAsyncResult
	var arg1 *C.GCancellable

	arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GCancellable)(unsafe.Pointer(checkCancellable.Native()))

	C.g_simple_async_result_set_check_cancellable(arg0, arg1)
}

// SetFromError sets the result from a #GError.
func (s simpleAsyncResult) SetFromError(s SimpleAsyncResult, error error) {
	var arg0 *C.GSimpleAsyncResult
	var arg1 *C.GError

	arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GError)(gerror.New(unsafe.Pointer(error)))
	defer C.g_error_free(arg1)

	C.g_simple_async_result_set_from_error(arg0, arg1)
}

// SetHandleCancellation sets whether to handle cancellation within the
// asynchronous operation.
//
// This function has nothing to do with
// g_simple_async_result_set_check_cancellable(). It only refers to the
// #GCancellable passed to g_simple_async_result_run_in_thread().
func (s simpleAsyncResult) SetHandleCancellation(s SimpleAsyncResult, handleCancellation bool) {
	var arg0 *C.GSimpleAsyncResult
	var arg1 C.gboolean

	arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(s.Native()))
	if handleCancellation {
		arg1 = C.gboolean(1)
	}

	C.g_simple_async_result_set_handle_cancellation(arg0, arg1)
}

// SetOpResGboolean sets the operation result to a boolean within the
// asynchronous result.
func (s simpleAsyncResult) SetOpResGboolean(s SimpleAsyncResult, opRes bool) {
	var arg0 *C.GSimpleAsyncResult
	var arg1 C.gboolean

	arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(s.Native()))
	if opRes {
		arg1 = C.gboolean(1)
	}

	C.g_simple_async_result_set_op_res_gboolean(arg0, arg1)
}

// SetOpResGssize sets the operation result within the asynchronous result
// to the given @op_res.
func (s simpleAsyncResult) SetOpResGssize(s SimpleAsyncResult, opRes int) {
	var arg0 *C.GSimpleAsyncResult
	var arg1 C.gssize

	arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(s.Native()))
	arg1 = C.gssize(opRes)

	C.g_simple_async_result_set_op_res_gssize(arg0, arg1)
}

// TakeError sets the result from @error, and takes over the caller's
// ownership of @error, so the caller does not need to free it any more.
func (s simpleAsyncResult) TakeError(s SimpleAsyncResult, error error) {
	var arg0 *C.GSimpleAsyncResult
	var arg1 *C.GError

	arg0 = (*C.GSimpleAsyncResult)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GError)(gerror.New(unsafe.Pointer(error)))
	defer C.g_error_free(arg1)

	C.g_simple_async_result_take_error(arg0, arg1)
}
