// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
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
		{T: externglib.Type(C.g_threaded_socket_service_get_type()), F: marshalThreadedSocketService},
	})
}

// ThreadedSocketService: a SocketService is a simple subclass of Service that
// handles incoming connections by creating a worker thread and dispatching the
// connection to it by emitting the SocketService::run signal in the new thread.
//
// The signal handler may perform blocking IO and need not return until the
// connection is closed.
//
// The service is implemented using a thread pool, so there is a limited amount
// of threads available to serve incoming requests. The service automatically
// stops the Service from accepting new connections when all threads are busy.
//
// As with Service, you may connect to SocketService::run, or subclass and
// override the default handler.
type ThreadedSocketService interface {
	SocketService
}

// threadedSocketService implements the ThreadedSocketService interface.
type threadedSocketService struct {
	SocketService
}

var _ ThreadedSocketService = (*threadedSocketService)(nil)

// WrapThreadedSocketService wraps a GObject to the right type. It is
// primarily used internally.
func WrapThreadedSocketService(obj *externglib.Object) ThreadedSocketService {
	return ThreadedSocketService{
		SocketService: WrapSocketService(obj),
	}
}

func marshalThreadedSocketService(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapThreadedSocketService(obj), nil
}

// NewThreadedSocketService constructs a class ThreadedSocketService.
func NewThreadedSocketService(maxThreads int) ThreadedSocketService {
	var arg1 C.int

	arg1 = C.int(maxThreads)

	var cret C.GThreadedSocketService

	cret = C.g_threaded_socket_service_new(arg1)

	var threadedSocketService ThreadedSocketService

	threadedSocketService = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(ThreadedSocketService)

	return threadedSocketService
}
