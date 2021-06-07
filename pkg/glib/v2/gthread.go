// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_thread_get_type()), F: marshalThread},
	})
}

// OnceStatus: the possible statuses of a one-time initialization function
// controlled by a #GOnce struct.
type OnceStatus int

const (
	// OnceStatusNotcalled: the function has not been called yet.
	OnceStatusNotcalled OnceStatus = 0
	// OnceStatusProgress: the function call is currently in progress.
	OnceStatusProgress OnceStatus = 1
	// OnceStatusReady: the function has been called.
	OnceStatusReady OnceStatus = 2
)

// ThreadError: possible errors of thread related functions.
type ThreadError int

const (
	// ThreadErrorThreadErrorAgain: a thread couldn't be created due to resource
	// shortage. Try again later.
	ThreadErrorThreadErrorAgain ThreadError = 0
)

// GetNumProcessors: determine the approximate number of threads that the system
// will schedule simultaneously for this process. This is intended to be used as
// a parameter to g_thread_pool_new() for CPU bound tasks and similar cases.
func GetNumProcessors() {
	C.g_get_num_processors()
}

// OnceInitEnter: function to be called when starting a critical initialization
// section. The argument @location must point to a static 0-initialized variable
// that will be set to a value other than 0 at the end of the initialization
// section. In combination with g_once_init_leave() and the unique address
// @value_location, it can be ensured that an initialization section will be
// executed only once during a program's life time, and that concurrent threads
// are blocked until initialization completed. To be used in constructs like
// this:
//
//      static gsize initialization_value = 0;
//
//      if (g_once_init_enter (&initialization_value))
//        {
//          gsize setup_value = 42; // initialization code here
//
//          g_once_init_leave (&initialization_value, setup_value);
//        }
//
//      // use initialization_value here
func OnceInitEnter(location interface{}) bool {
	var arg1 *C.void

	arg1 = *C.void(location)

	var cret C.gboolean
	var ok bool

	cret = C.g_once_init_enter(arg1)

	if cret {
		ok = true
	}

	return ok
}

// OnceInitLeave: counterpart to g_once_init_enter(). Expects a location of a
// static 0-initialized initialization variable, and an initialization value
// other than 0. Sets the variable to the initialization value, and releases
// concurrent threads blocking in g_once_init_enter() on this initialization
// variable.
func OnceInitLeave(location interface{}, result uint) {
	var arg1 *C.void
	var arg2 C.gsize

	arg1 = *C.void(location)
	arg2 = C.gsize(result)

	C.g_once_init_leave(arg1, arg2)
}

// ThreadExit terminates the current thread.
//
// If another thread is waiting for us using g_thread_join() then the waiting
// thread will be woken up and get @retval as the return value of
// g_thread_join().
//
// Calling g_thread_exit() with a parameter @retval is equivalent to returning
// @retval from the function @func, as given to g_thread_new().
//
// You must only call g_thread_exit() from a thread that you created yourself
// with g_thread_new() or related APIs. You must not call this function from a
// thread created with another threading library or or from within a Pool.
func ThreadExit(retval interface{}) {
	var arg1 C.gpointer

	arg1 = C.gpointer(retval)

	C.g_thread_exit(arg1)
}

// ThreadSelf: this function returns the #GThread corresponding to the current
// thread. Note that this function does not increase the reference count of the
// returned struct.
//
// This function will return a #GThread even for threads that were not created
// by GLib (i.e. those created by other threading APIs). This may be useful for
// thread identification purposes (i.e. comparisons) but you must not use GLib
// functions (such as g_thread_join()) on these threads.
func ThreadSelf() {
	C.g_thread_self()
}

// ThreadYield causes the calling thread to voluntarily relinquish the CPU, so
// that other threads can run.
//
// This function is often used as a method to make busy wait less evil.
func ThreadYield() {
	C.g_thread_yield()
}

// Cond: the #GCond struct is an opaque data structure that represents a
// condition. Threads can block on a #GCond if they find a certain condition to
// be false. If other threads change the state of this condition they signal the
// #GCond, and that causes the waiting threads to be woken up.
//
// Consider the following example of a shared variable. One or more threads can
// wait for data to be published to the variable and when another thread
// publishes the data, it can signal one of the waiting threads to wake up to
// collect the data.
//
// Here is an example for using GCond to block a thread until a condition is
// satisfied:
//
//      gpointer current_data = NULL;
//      GMutex data_mutex;
//      GCond data_cond;
//
//      void
//      push_data (gpointer data)
//      {
//        g_mutex_lock (&data_mutex);
//        current_data = data;
//        g_cond_signal (&data_cond);
//        g_mutex_unlock (&data_mutex);
//      }
//
//      gpointer
//      pop_data (void)
//      {
//        gpointer data;
//
//        g_mutex_lock (&data_mutex);
//        while (!current_data)
//          g_cond_wait (&data_cond, &data_mutex);
//        data = current_data;
//        current_data = NULL;
//        g_mutex_unlock (&data_mutex);
//
//        return data;
//      }
//
// Whenever a thread calls pop_data() now, it will wait until current_data is
// non-nil, i.e. until some other thread has called push_data().
//
// The example shows that use of a condition variable must always be paired with
// a mutex. Without the use of a mutex, there would be a race between the check
// of @current_data by the while loop in pop_data() and waiting. Specifically,
// another thread could set @current_data after the check, and signal the cond
// (with nobody waiting on it) before the first thread goes to sleep. #GCond is
// specifically useful for its ability to release the mutex and go to sleep
// atomically.
//
// It is also important to use the g_cond_wait() and g_cond_wait_until()
// functions only inside a loop which checks for the condition to be true. See
// g_cond_wait() for an explanation of why the condition may not be true even
// after it returns.
//
// If a #GCond is allocated in static storage then it can be used without
// initialisation. Otherwise, you should call g_cond_init() on it and
// g_cond_clear() when done.
//
// A #GCond should only be accessed via the g_cond_ functions.
type Cond struct {
	native C.GCond
}

// WrapCond wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapCond(ptr unsafe.Pointer) *Cond {
	if ptr == nil {
		return nil
	}

	return (*Cond)(ptr)
}

func marshalCond(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapCond(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *Cond) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}

// Broadcast: if threads are waiting for @cond, all of them are unblocked. If no
// threads are waiting for @cond, this function has no effect. It is good
// practice to lock the same mutex as the waiting threads while calling this
// function, though not required.
func (c *Cond) Broadcast(c *Cond) {
	var arg0 *C.GCond

	arg0 = (*C.GCond)(unsafe.Pointer(c.Native()))

	C.g_cond_broadcast(arg0)
}

// Clear frees the resources allocated to a #GCond with g_cond_init().
//
// This function should not be used with a #GCond that has been statically
// allocated.
//
// Calling g_cond_clear() for a #GCond on which threads are blocking leads to
// undefined behaviour.
func (c *Cond) Clear(c *Cond) {
	var arg0 *C.GCond

	arg0 = (*C.GCond)(unsafe.Pointer(c.Native()))

	C.g_cond_clear(arg0)
}

// Init initialises a #GCond so that it can be used.
//
// This function is useful to initialise a #GCond that has been allocated as
// part of a larger structure. It is not necessary to initialise a #GCond that
// has been statically allocated.
//
// To undo the effect of g_cond_init() when a #GCond is no longer needed, use
// g_cond_clear().
//
// Calling g_cond_init() on an already-initialised #GCond leads to undefined
// behaviour.
func (c *Cond) Init(c *Cond) {
	var arg0 *C.GCond

	arg0 = (*C.GCond)(unsafe.Pointer(c.Native()))

	C.g_cond_init(arg0)
}

// Signal: if threads are waiting for @cond, at least one of them is unblocked.
// If no threads are waiting for @cond, this function has no effect. It is good
// practice to hold the same lock as the waiting thread while calling this
// function, though not required.
func (c *Cond) Signal(c *Cond) {
	var arg0 *C.GCond

	arg0 = (*C.GCond)(unsafe.Pointer(c.Native()))

	C.g_cond_signal(arg0)
}

// Once: a #GOnce struct controls a one-time initialization function. Any
// one-time initialization function must have its own unique #GOnce struct.
type Once struct {
	native C.GOnce
}

// WrapOnce wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapOnce(ptr unsafe.Pointer) *Once {
	if ptr == nil {
		return nil
	}

	return (*Once)(ptr)
}

func marshalOnce(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapOnce(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (o *Once) Native() unsafe.Pointer {
	return unsafe.Pointer(&o.native)
}

// Status gets the field inside the struct.
func (o *Once) Status() OnceStatus {
	var v OnceStatus
	v = OnceStatus(o.native.status)
	return v
}

// Retval gets the field inside the struct.
func (o *Once) Retval() interface{} {
	var v interface{}
	v = interface{}(o.native.retval)
	return v
}

// RWLock: the GRWLock struct is an opaque data structure to represent a
// reader-writer lock. It is similar to a #GMutex in that it allows multiple
// threads to coordinate access to a shared resource.
//
// The difference to a mutex is that a reader-writer lock discriminates between
// read-only ('reader') and full ('writer') access. While only one thread at a
// time is allowed write access (by holding the 'writer' lock via
// g_rw_lock_writer_lock()), multiple threads can gain simultaneous read-only
// access (by holding the 'reader' lock via g_rw_lock_reader_lock()).
//
// It is unspecified whether readers or writers have priority in acquiring the
// lock when a reader already holds the lock and a writer is queued to acquire
// it.
//
// Here is an example for an array with access functions: |[<!-- language="C"
// --> GRWLock lock; GPtrArray *array;
//
//     gpointer
//     my_array_get (guint index)
//     {
//       gpointer retval = NULL;
//
//       if (!array)
//         return NULL;
//
//       g_rw_lock_reader_lock (&lock);
//       if (index < array->len)
//         retval = g_ptr_array_index (array, index);
//       g_rw_lock_reader_unlock (&lock);
//
//       return retval;
//     }
//
//     void
//     my_array_set (guint index, gpointer data)
//     {
//       g_rw_lock_writer_lock (&lock);
//
//       if (!array)
//         array = g_ptr_array_new ();
//
//       if (index >= array->len)
//         g_ptr_array_set_size (array, index+1);
//       g_ptr_array_index (array, index) = data;
//
//       g_rw_lock_writer_unlock (&lock);
//     }
//    ]|
//
// This example shows an array which can be accessed by many readers (the
// my_array_get() function) simultaneously, whereas the writers (the
// my_array_set() function) will only be allowed one at a time and only if no
// readers currently access the array. This is because of the potentially
// dangerous resizing of the array. Using these functions is fully multi-thread
// safe now.
//
// If a WLock is allocated in static storage then it can be used without
// initialisation. Otherwise, you should call g_rw_lock_init() on it and
// g_rw_lock_clear() when done.
//
// A GRWLock should only be accessed with the g_rw_lock_ functions.
type RWLock struct {
	native C.GRWLock
}

// WrapRWLock wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRWLock(ptr unsafe.Pointer) *RWLock {
	if ptr == nil {
		return nil
	}

	return (*RWLock)(ptr)
}

func marshalRWLock(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRWLock(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RWLock) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Clear frees the resources allocated to a lock with g_rw_lock_init().
//
// This function should not be used with a WLock that has been statically
// allocated.
//
// Calling g_rw_lock_clear() when any thread holds the lock leads to undefined
// behaviour.
//
// Sine: 2.32
func (r *RWLock) Clear(r *RWLock) {
	var arg0 *C.GRWLock

	arg0 = (*C.GRWLock)(unsafe.Pointer(r.Native()))

	C.g_rw_lock_clear(arg0)
}

// Init initializes a WLock so that it can be used.
//
// This function is useful to initialize a lock that has been allocated on the
// stack, or as part of a larger structure. It is not necessary to initialise a
// reader-writer lock that has been statically allocated.
//
//      typedef struct {
//        GRWLock l;
//        ...
//      } Blob;
//
//    Blob *b;
//
//    b = g_new (Blob, 1);
//    g_rw_lock_init (&b->l);
//
// To undo the effect of g_rw_lock_init() when a lock is no longer needed, use
// g_rw_lock_clear().
//
// Calling g_rw_lock_init() on an already initialized WLock leads to undefined
// behaviour.
func (r *RWLock) Init(r *RWLock) {
	var arg0 *C.GRWLock

	arg0 = (*C.GRWLock)(unsafe.Pointer(r.Native()))

	C.g_rw_lock_init(arg0)
}

// ReaderLock: obtain a read lock on @rw_lock. If another thread currently holds
// the write lock on @rw_lock, the current thread will block. If another thread
// does not hold the write lock, but is waiting for it, it is implementation
// defined whether the reader or writer will block. Read locks can be taken
// recursively.
//
// It is implementation-defined how many threads are allowed to hold read locks
// on the same lock simultaneously. If the limit is hit, or if a deadlock is
// detected, a critical warning will be emitted.
func (r *RWLock) ReaderLock(r *RWLock) {
	var arg0 *C.GRWLock

	arg0 = (*C.GRWLock)(unsafe.Pointer(r.Native()))

	C.g_rw_lock_reader_lock(arg0)
}

// ReaderTrylock tries to obtain a read lock on @rw_lock and returns true if the
// read lock was successfully obtained. Otherwise it returns false.
func (r *RWLock) ReaderTrylock(r *RWLock) bool {
	var arg0 *C.GRWLock

	arg0 = (*C.GRWLock)(unsafe.Pointer(r.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_rw_lock_reader_trylock(arg0)

	if cret {
		ok = true
	}

	return ok
}

// ReaderUnlock: release a read lock on @rw_lock.
//
// Calling g_rw_lock_reader_unlock() on a lock that is not held by the current
// thread leads to undefined behaviour.
func (r *RWLock) ReaderUnlock(r *RWLock) {
	var arg0 *C.GRWLock

	arg0 = (*C.GRWLock)(unsafe.Pointer(r.Native()))

	C.g_rw_lock_reader_unlock(arg0)
}

// WriterLock: obtain a write lock on @rw_lock. If any thread already holds a
// read or write lock on @rw_lock, the current thread will block until all other
// threads have dropped their locks on @rw_lock.
func (r *RWLock) WriterLock(r *RWLock) {
	var arg0 *C.GRWLock

	arg0 = (*C.GRWLock)(unsafe.Pointer(r.Native()))

	C.g_rw_lock_writer_lock(arg0)
}

// WriterTrylock tries to obtain a write lock on @rw_lock. If any other thread
// holds a read or write lock on @rw_lock, it immediately returns false.
// Otherwise it locks @rw_lock and returns true.
func (r *RWLock) WriterTrylock(r *RWLock) bool {
	var arg0 *C.GRWLock

	arg0 = (*C.GRWLock)(unsafe.Pointer(r.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_rw_lock_writer_trylock(arg0)

	if cret {
		ok = true
	}

	return ok
}

// WriterUnlock: release a write lock on @rw_lock.
//
// Calling g_rw_lock_writer_unlock() on a lock that is not held by the current
// thread leads to undefined behaviour.
func (r *RWLock) WriterUnlock(r *RWLock) {
	var arg0 *C.GRWLock

	arg0 = (*C.GRWLock)(unsafe.Pointer(r.Native()))

	C.g_rw_lock_writer_unlock(arg0)
}

// RecMutex: the GRecMutex struct is an opaque data structure to represent a
// recursive mutex. It is similar to a #GMutex with the difference that it is
// possible to lock a GRecMutex multiple times in the same thread without
// deadlock. When doing so, care has to be taken to unlock the recursive mutex
// as often as it has been locked.
//
// If a Mutex is allocated in static storage then it can be used without
// initialisation. Otherwise, you should call g_rec_mutex_init() on it and
// g_rec_mutex_clear() when done.
//
// A GRecMutex should only be accessed with the g_rec_mutex_ functions.
type RecMutex struct {
	native C.GRecMutex
}

// WrapRecMutex wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRecMutex(ptr unsafe.Pointer) *RecMutex {
	if ptr == nil {
		return nil
	}

	return (*RecMutex)(ptr)
}

func marshalRecMutex(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRecMutex(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *RecMutex) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Clear frees the resources allocated to a recursive mutex with
// g_rec_mutex_init().
//
// This function should not be used with a Mutex that has been statically
// allocated.
//
// Calling g_rec_mutex_clear() on a locked recursive mutex leads to undefined
// behaviour.
//
// Sine: 2.32
func (r *RecMutex) Clear(r *RecMutex) {
	var arg0 *C.GRecMutex

	arg0 = (*C.GRecMutex)(unsafe.Pointer(r.Native()))

	C.g_rec_mutex_clear(arg0)
}

// Init initializes a Mutex so that it can be used.
//
// This function is useful to initialize a recursive mutex that has been
// allocated on the stack, or as part of a larger structure.
//
// It is not necessary to initialise a recursive mutex that has been statically
// allocated.
//
//      typedef struct {
//        GRecMutex m;
//        ...
//      } Blob;
//
//    Blob *b;
//
//    b = g_new (Blob, 1);
//    g_rec_mutex_init (&b->m);
//
// Calling g_rec_mutex_init() on an already initialized Mutex leads to undefined
// behaviour.
//
// To undo the effect of g_rec_mutex_init() when a recursive mutex is no longer
// needed, use g_rec_mutex_clear().
func (r *RecMutex) Init(r *RecMutex) {
	var arg0 *C.GRecMutex

	arg0 = (*C.GRecMutex)(unsafe.Pointer(r.Native()))

	C.g_rec_mutex_init(arg0)
}

// Lock locks @rec_mutex. If @rec_mutex is already locked by another thread, the
// current thread will block until @rec_mutex is unlocked by the other thread.
// If @rec_mutex is already locked by the current thread, the 'lock count' of
// @rec_mutex is increased. The mutex will only become available again when it
// is unlocked as many times as it has been locked.
func (r *RecMutex) Lock(r *RecMutex) {
	var arg0 *C.GRecMutex

	arg0 = (*C.GRecMutex)(unsafe.Pointer(r.Native()))

	C.g_rec_mutex_lock(arg0)
}

// Trylock tries to lock @rec_mutex. If @rec_mutex is already locked by another
// thread, it immediately returns false. Otherwise it locks @rec_mutex and
// returns true.
func (r *RecMutex) Trylock(r *RecMutex) bool {
	var arg0 *C.GRecMutex

	arg0 = (*C.GRecMutex)(unsafe.Pointer(r.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_rec_mutex_trylock(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Unlock unlocks @rec_mutex. If another thread is blocked in a
// g_rec_mutex_lock() call for @rec_mutex, it will become unblocked and can lock
// @rec_mutex itself.
//
// Calling g_rec_mutex_unlock() on a recursive mutex that is not locked by the
// current thread leads to undefined behaviour.
func (r *RecMutex) Unlock(r *RecMutex) {
	var arg0 *C.GRecMutex

	arg0 = (*C.GRecMutex)(unsafe.Pointer(r.Native()))

	C.g_rec_mutex_unlock(arg0)
}

// Thread: the #GThread struct represents a running thread. This struct is
// returned by g_thread_new() or g_thread_try_new(). You can obtain the #GThread
// struct representing the current thread by calling g_thread_self().
//
// GThread is refcounted, see g_thread_ref() and g_thread_unref(). The thread
// represented by it holds a reference while it is running, and g_thread_join()
// consumes the reference that it is given, so it is normally not necessary to
// manage GThread references explicitly.
//
// The structure is opaque -- none of its fields may be directly accessed.
type Thread struct {
	native C.GThread
}

// WrapThread wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapThread(ptr unsafe.Pointer) *Thread {
	if ptr == nil {
		return nil
	}

	return (*Thread)(ptr)
}

func marshalThread(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapThread(unsafe.Pointer(b)), nil
}

// NewThread constructs a struct Thread.
func NewThread() {
	C.g_thread_new(arg1, arg2, arg3)
}

// NewThreadTry constructs a struct Thread.
func NewThreadTry() error {
	var errout *C.GError
	var err error

	C.g_thread_try_new(arg1, arg2, arg3, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// Native returns the underlying C source pointer.
func (t *Thread) Native() unsafe.Pointer {
	return unsafe.Pointer(&t.native)
}

// Join waits until @thread finishes, i.e. the function @func, as given to
// g_thread_new(), returns or g_thread_exit() is called. If @thread has already
// terminated, then g_thread_join() returns immediately.
//
// Any thread can wait for any other thread by calling g_thread_join(), not just
// its 'creator'. Calling g_thread_join() from multiple threads for the same
// @thread leads to undefined behaviour.
//
// The value returned by @func or given to g_thread_exit() is returned by this
// function.
//
// g_thread_join() consumes the reference to the passed-in @thread. This will
// usually cause the #GThread struct and associated resources to be freed. Use
// g_thread_ref() to obtain an extra reference if you want to keep the GThread
// alive beyond the g_thread_join() call.
func (t *Thread) Join(t *Thread) {
	var arg0 *C.GThread

	arg0 = (*C.GThread)(unsafe.Pointer(t.Native()))

	C.g_thread_join(arg0)
}

// Ref: increase the reference count on @thread.
func (t *Thread) Ref(t *Thread) {
	var arg0 *C.GThread

	arg0 = (*C.GThread)(unsafe.Pointer(t.Native()))

	C.g_thread_ref(arg0)
}

// Unref: decrease the reference count on @thread, possibly freeing all
// resources associated with it.
//
// Note that each thread holds a reference to its #GThread while it is running,
// so it is safe to drop your own reference to it if you don't need it anymore.
func (t *Thread) Unref(t *Thread) {
	var arg0 *C.GThread

	arg0 = (*C.GThread)(unsafe.Pointer(t.Native()))

	C.g_thread_unref(arg0)
}
