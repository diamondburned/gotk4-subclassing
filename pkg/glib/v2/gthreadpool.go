// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

// ThreadPoolGetMaxIdleTime: this function will return the maximum @interval
// that a thread will wait in the thread pool for new tasks before being
// stopped.
//
// If this function returns 0, threads waiting in the thread pool for new work
// are not stopped.
func ThreadPoolGetMaxIdleTime() {
	C.g_thread_pool_get_max_idle_time()
}

// ThreadPoolGetMaxUnusedThreads returns the maximal allowed number of unused
// threads.
func ThreadPoolGetMaxUnusedThreads() {
	C.g_thread_pool_get_max_unused_threads()
}

// ThreadPoolGetNumUnusedThreads returns the number of currently unused threads.
func ThreadPoolGetNumUnusedThreads() {
	C.g_thread_pool_get_num_unused_threads()
}

// ThreadPoolSetMaxIdleTime: this function will set the maximum @interval that a
// thread waiting in the pool for new tasks can be idle for before being
// stopped. This function is similar to calling
// g_thread_pool_stop_unused_threads() on a regular timeout, except this is done
// on a per thread basis.
//
// By setting @interval to 0, idle threads will not be stopped.
//
// The default value is 15000 (15 seconds).
func ThreadPoolSetMaxIdleTime(interval uint) {
	var arg1 C.guint

	arg1 = C.guint(interval)

	C.g_thread_pool_set_max_idle_time(arg1)
}

// ThreadPoolSetMaxUnusedThreads sets the maximal number of unused threads to
// @max_threads. If @max_threads is -1, no limit is imposed on the number of
// unused threads.
//
// The default value is 2.
func ThreadPoolSetMaxUnusedThreads(maxThreads int) {
	var arg1 C.gint

	arg1 = C.gint(maxThreads)

	C.g_thread_pool_set_max_unused_threads(arg1)
}

// ThreadPoolStopUnusedThreads stops all currently unused threads. This does not
// change the maximal number of unused threads. This function can be used to
// regularly stop all unused threads e.g. from g_timeout_add().
func ThreadPoolStopUnusedThreads() {
	C.g_thread_pool_stop_unused_threads()
}

// ThreadPool: the Pool struct represents a thread pool. It has three public
// read-only members, but the underlying struct is bigger, so you must not copy
// this struct.
type ThreadPool struct {
	native C.GThreadPool
}

// WrapThreadPool wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapThreadPool(ptr unsafe.Pointer) *ThreadPool {
	if ptr == nil {
		return nil
	}

	return (*ThreadPool)(ptr)
}

func marshalThreadPool(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapThreadPool(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (t *ThreadPool) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// UserData gets the field inside the struct.
func (t *ThreadPool) UserData() interface{} {
	var v interface{}
	v = interface{}(t.native.user_data)
	return v
}

// Exclusive gets the field inside the struct.
func (t *ThreadPool) Exclusive() bool {
	var v bool
	if t.native.exclusive {
		v = true
	}
	return v
}

// Free frees all resources allocated for @pool.
//
// If @immediate is true, no new task is processed for @pool. Otherwise @pool is
// not freed before the last task is processed. Note however, that no thread of
// this pool is interrupted while processing a task. Instead at least all still
// running threads can finish their tasks before the @pool is freed.
//
// If @wait_ is true, this function does not return before all tasks to be
// processed (dependent on @immediate, whether all or only the currently
// running) are ready. Otherwise this function returns immediately.
//
// After calling this function @pool must not be used anymore.
func (p *ThreadPool) Free(p *ThreadPool, immediate bool, wait_ bool) {
	var arg0 *C.GThreadPool
	var arg1 C.gboolean
	var arg2 C.gboolean

	arg0 = (*C.GThreadPool)(unsafe.Pointer(p.Native()))
	if immediate {
		arg1 = C.gboolean(1)
	}
	if wait_ {
		arg2 = C.gboolean(1)
	}

	C.g_thread_pool_free(arg0, arg1, arg2)
}

// MaxThreads returns the maximal number of threads for @pool.
func (p *ThreadPool) MaxThreads(p *ThreadPool) {
	var arg0 *C.GThreadPool

	arg0 = (*C.GThreadPool)(unsafe.Pointer(p.Native()))

	C.g_thread_pool_get_max_threads(arg0)
}

// NumThreads returns the number of threads currently running in @pool.
func (p *ThreadPool) NumThreads(p *ThreadPool) {
	var arg0 *C.GThreadPool

	arg0 = (*C.GThreadPool)(unsafe.Pointer(p.Native()))

	C.g_thread_pool_get_num_threads(arg0)
}

// MoveToFront moves the item to the front of the queue of unprocessed items, so
// that it will be processed next.
func (p *ThreadPool) MoveToFront(p *ThreadPool, data interface{}) bool {
	var arg0 *C.GThreadPool
	var arg1 C.gpointer

	arg0 = (*C.GThreadPool)(unsafe.Pointer(p.Native()))
	arg1 = C.gpointer(data)

	var cret C.gboolean
	var ok bool

	cret = C.g_thread_pool_move_to_front(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// Push inserts @data into the list of tasks to be executed by @pool.
//
// When the number of currently running threads is lower than the maximal
// allowed number of threads, a new thread is started (or reused) with the
// properties given to g_thread_pool_new(). Otherwise, @data stays in the queue
// until a thread in this pool finishes its previous task and processes @data.
//
// @error can be nil to ignore errors, or non-nil to report errors. An error can
// only occur when a new thread couldn't be created. In that case @data is
// simply appended to the queue of work to do.
//
// Before version 2.32, this function did not return a success status.
func (p *ThreadPool) Push(p *ThreadPool, data interface{}) error {
	var arg0 *C.GThreadPool
	var arg1 C.gpointer

	arg0 = (*C.GThreadPool)(unsafe.Pointer(p.Native()))
	arg1 = C.gpointer(data)

	var errout *C.GError
	var err error

	C.g_thread_pool_push(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// SetMaxThreads sets the maximal allowed number of threads for @pool. A value
// of -1 means that the maximal number of threads is unlimited. If @pool is an
// exclusive thread pool, setting the maximal number of threads to -1 is not
// allowed.
//
// Setting @max_threads to 0 means stopping all work for @pool. It is
// effectively frozen until @max_threads is set to a non-zero value again.
//
// A thread is never terminated while calling @func, as supplied by
// g_thread_pool_new(). Instead the maximal number of threads only has effect
// for the allocation of new threads in g_thread_pool_push(). A new thread is
// allocated, whenever the number of currently running threads in @pool is
// smaller than the maximal number.
//
// @error can be nil to ignore errors, or non-nil to report errors. An error can
// only occur when a new thread couldn't be created.
//
// Before version 2.32, this function did not return a success status.
func (p *ThreadPool) SetMaxThreads(p *ThreadPool, maxThreads int) error {
	var arg0 *C.GThreadPool
	var arg1 C.gint

	arg0 = (*C.GThreadPool)(unsafe.Pointer(p.Native()))
	arg1 = C.gint(maxThreads)

	var errout *C.GError
	var err error

	C.g_thread_pool_set_max_threads(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// SetSortFunction sets the function used to sort the list of tasks. This allows
// the tasks to be processed by a priority determined by @func, and not just in
// the order in which they were added to the pool.
//
// Note, if the maximum number of threads is more than 1, the order that threads
// are executed cannot be guaranteed 100%. Threads are scheduled by the
// operating system and are executed at random. It cannot be assumed that
// threads are executed in the order they are created.
func (p *ThreadPool) SetSortFunction(p *ThreadPool) {
	var arg0 *C.GThreadPool

	arg0 = (*C.GThreadPool)(unsafe.Pointer(p.Native()))

	C.g_thread_pool_set_sort_function(arg0, arg1, arg2)
}

// Unprocessed returns the number of tasks still unprocessed in @pool.
func (p *ThreadPool) Unprocessed(p *ThreadPool) {
	var arg0 *C.GThreadPool

	arg0 = (*C.GThreadPool)(unsafe.Pointer(p.Native()))

	C.g_thread_pool_unprocessed(arg0)
}