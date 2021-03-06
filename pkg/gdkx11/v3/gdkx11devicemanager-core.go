// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #include <stdlib.h>
// #include <gdk/gdkx.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_x11_device_manager_core_get_type()), F: marshalX11DeviceManagerCorer},
	})
}

// X11DeviceManagerCoreOverrider contains methods that are overridable.
type X11DeviceManagerCoreOverrider interface {
}

type X11DeviceManagerCore struct {
	_ [0]func() // equal guard
	gdk.DeviceManager
}

var (
	_ gdk.DeviceManagerer = (*X11DeviceManagerCore)(nil)
)

func classInitX11DeviceManagerCorer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapX11DeviceManagerCore(obj *externglib.Object) *X11DeviceManagerCore {
	return &X11DeviceManagerCore{
		DeviceManager: gdk.DeviceManager{
			Object: obj,
		},
	}
}

func marshalX11DeviceManagerCorer(p uintptr) (interface{}, error) {
	return wrapX11DeviceManagerCore(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
