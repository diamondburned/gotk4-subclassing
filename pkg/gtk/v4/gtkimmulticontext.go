// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_im_multicontext_get_type()), F: marshalIMMulticontext},
	})
}

type IMMulticontext interface {
	IMContext

	// ContextID gets the id of the currently active delegate of the @context.
	ContextID() string
	// SetContextID sets the context id for @context.
	//
	// This causes the currently active delegate of @context to be replaced by
	// the delegate corresponding to the new context id.
	SetContextID(contextId string)
}

// imMulticontext implements the IMMulticontext interface.
type imMulticontext struct {
	IMContext
}

var _ IMMulticontext = (*imMulticontext)(nil)

// WrapIMMulticontext wraps a GObject to the right type. It is
// primarily used internally.
func WrapIMMulticontext(obj *externglib.Object) IMMulticontext {
	return IMMulticontext{
		IMContext: WrapIMContext(obj),
	}
}

func marshalIMMulticontext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapIMMulticontext(obj), nil
}

// NewIMMulticontext constructs a class IMMulticontext.
func NewIMMulticontext() IMMulticontext {
	var cret C.GtkIMMulticontext

	cret = C.gtk_im_multicontext_new()

	var imMulticontext IMMulticontext

	imMulticontext = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(IMMulticontext)

	return imMulticontext
}

// ContextID gets the id of the currently active delegate of the @context.
func (c imMulticontext) ContextID() string {
	var arg0 *C.GtkIMMulticontext

	arg0 = (*C.GtkIMMulticontext)(unsafe.Pointer(c.Native()))

	var cret *C.char

	cret = C.gtk_im_multicontext_get_context_id(arg0)

	var utf8 string

	utf8 = C.GoString(cret)

	return utf8
}

// SetContextID sets the context id for @context.
//
// This causes the currently active delegate of @context to be replaced by
// the delegate corresponding to the new context id.
func (c imMulticontext) SetContextID(contextId string) {
	var arg0 *C.GtkIMMulticontext
	var arg1 *C.char

	arg0 = (*C.GtkIMMulticontext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.char)(C.CString(contextId))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_im_multicontext_set_context_id(arg0, arg1)
}
