// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_socket_get_type()), F: marshalSocket},
	})
}

// Socket: together with Plug, Socket provides the ability to embed widgets from
// one process into another process in a fashion that is transparent to the
// user. One process creates a Socket widget and passes that widget’s window ID
// to the other process, which then creates a Plug with that window ID. Any
// widgets contained in the Plug then will appear inside the first application’s
// window.
//
// The socket’s window ID is obtained by using gtk_socket_get_id(). Before using
// this function, the socket must have been realized, and for hence, have been
// added to its parent.
//
// Obtaining the window ID of a socket.
//
//    GtkWidget *socket = gtk_socket_new ();
//    gtk_widget_show (socket);
//    gtk_container_add (GTK_CONTAINER (parent), socket);
//
//    // The following call is only necessary if one of
//    // the ancestors of the socket is not yet visible.
//    gtk_widget_realize (socket);
//    g_print ("The ID of the sockets window is %#x\n",
//             gtk_socket_get_id (socket));
//
// Note that if you pass the window ID of the socket to another process that
// will create a plug in the socket, you must make sure that the socket widget
// is not destroyed until that plug is created. Violating this rule will cause
// unpredictable consequences, the most likely consequence being that the plug
// will appear as a separate toplevel window. You can check if the plug has been
// created by using gtk_socket_get_plug_window(). If it returns a non-nil value,
// then the plug has been successfully created inside of the socket.
//
// When GTK+ is notified that the embedded window has been destroyed, then it
// will destroy the socket as well. You should always, therefore, be prepared
// for your sockets to be destroyed at any time when the main event loop is
// running. To prevent this from happening, you can connect to the
// Socket::plug-removed signal.
//
// The communication between a Socket and a Plug follows the XEmbed Protocol
// (http://www.freedesktop.org/Standards/xembed-spec). This protocol has also
// been implemented in other toolkits, e.g. Qt, allowing the same level of
// integration when embedding a Qt widget in GTK or vice versa.
//
// The Plug and Socket widgets are only available when GTK+ is compiled for the
// X11 platform and GDK_WINDOWING_X11 is defined. They can only be used on a
// X11Display. To use Plug and Socket, you need to include the `gtk/gtkx.h`
// header.
type Socket interface {
	Container
	Buildable

	// PlugWindow retrieves the window of the plug. Use this to check if the
	// plug has been created inside of the socket.
	PlugWindow() gdk.Window
}

// socket implements the Socket class.
type socket struct {
	Container
	Buildable
}

var _ Socket = (*socket)(nil)

// WrapSocket wraps a GObject to the right type. It is
// primarily used internally.
func WrapSocket(obj *externglib.Object) Socket {
	return socket{
		Container: WrapContainer(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalSocket(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapSocket(obj), nil
}

// NewSocket constructs a class Socket.
func NewSocket() Socket {
	var _cret C.GtkSocket // in

	_cret = C.gtk_socket_new()

	var _socket Socket // out

	_socket = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Socket)

	return _socket
}

// PlugWindow retrieves the window of the plug. Use this to check if the
// plug has been created inside of the socket.
func (s socket) PlugWindow() gdk.Window {
	var _arg0 *C.GtkSocket // out

	_arg0 = (*C.GtkSocket)(unsafe.Pointer(s.Native()))

	var _cret *C.GdkWindow // in

	_cret = C.gtk_socket_get_plug_window(_arg0)

	var _window gdk.Window // out

	_window = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gdk.Window)

	return _window
}
