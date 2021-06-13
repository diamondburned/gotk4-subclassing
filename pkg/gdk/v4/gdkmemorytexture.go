// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_memory_texture_get_type()), F: marshalMemoryTexture},
	})
}

// MemoryTexture: a `GdkTexture` representing image data in memory.
type MemoryTexture interface {
	Texture
	Paintable
}

// memoryTexture implements the MemoryTexture interface.
type memoryTexture struct {
	Texture
	Paintable
}

var _ MemoryTexture = (*memoryTexture)(nil)

// WrapMemoryTexture wraps a GObject to the right type. It is
// primarily used internally.
func WrapMemoryTexture(obj *externglib.Object) MemoryTexture {
	return MemoryTexture{
		Texture:   WrapTexture(obj),
		Paintable: WrapPaintable(obj),
	}
}

func marshalMemoryTexture(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapMemoryTexture(obj), nil
}
