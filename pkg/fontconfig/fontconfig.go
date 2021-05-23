// Code generated by girgen. DO NOT EDIT.

package fontconfig

import (
	"runtime"
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
//
//
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{

		// Records
		// Skipped Pattern.
		// Skipped CharSet.

	})
}

func Init() {
	C.FcInit()
}

type Pattern struct {
	native *C.FcPattern
}

func wrapPattern(p *C.FcPattern) *Pattern {
	v := Pattern{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*Pattern).free)

	return &v
}

func marshalPattern(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.FcPattern)(unsafe.Pointer(b))

	return wrapPattern(c)
}

func (p *Pattern) free() {}

// Native returns the pointer to *C.FcPattern. The caller is expected to
// cast.
func (p *Pattern) Native() unsafe.Pointer {
	return unsafe.Pointer(p.native)
}

type CharSet struct {
	native *C.FcCharSet
}

func wrapCharSet(p *C.FcCharSet) *CharSet {
	v := CharSet{native: p}

	runtime.SetFinalizer(&v, nil)
	runtime.SetFinalizer(&v, (*CharSet).free)

	return &v
}

func marshalCharSet(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	c := (*C.FcCharSet)(unsafe.Pointer(b))

	return wrapCharSet(c)
}

func (c *CharSet) free() {}

// Native returns the pointer to *C.FcCharSet. The caller is expected to
// cast.
func (c *CharSet) Native() unsafe.Pointer {
	return unsafe.Pointer(c.native)
}
