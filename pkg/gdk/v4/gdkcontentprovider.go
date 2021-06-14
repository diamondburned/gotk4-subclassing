// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_content_provider_get_type()), F: marshalContentProvider},
	})
}

// ContentProvider: a `GdkContentProvider` is used to provide content for the
// clipboard or for drag-and-drop operations in a number of formats.
//
// To create a `GdkContentProvider`, use
// [ctor@Gdk.ContentProvider.new_for_value] or
// [ctor@Gdk.ContentProvider.new_for_bytes].
//
// GDK knows how to handle common text and image formats out-of-the-box. See
// [class@Gdk.ContentSerializer] and [class@Gdk.ContentDeserializer] if you want
// to add support for application-specific data formats.
type ContentProvider interface {
	gextras.Objector

	// ContentChanged emits the ::content-changed signal.
	ContentChanged()
	// Value gets the contents of @provider stored in @value.
	//
	// The @value will have been initialized to the `GType` the value should be
	// provided in. This given `GType` does not need to be listed in the formats
	// returned by [method@Gdk.ContentProvider.ref_formats]. However, if the
	// given `GType` is not supported, this operation can fail and
	// IO_ERROR_NOT_SUPPORTED will be reported.
	Value(value **externglib.Value) error
	// RefFormats gets the formats that the provider can provide its current
	// contents in.
	RefFormats() *ContentFormats
	// RefStorableFormats gets the formats that the provider suggests other
	// applications to store the data in.
	//
	// An example of such an application would be a clipboard manager.
	//
	// This can be assumed to be a subset of
	// [method@Gdk.ContentProvider.ref_formats].
	RefStorableFormats() *ContentFormats
	// WriteMIMETypeFinish finishes an asynchronous write operation.
	//
	// See [method@Gdk.ContentProvider.write_mime_type_async].
	WriteMIMETypeFinish(result gio.AsyncResult) error
}

// contentProvider implements the ContentProvider class.
type contentProvider struct {
	gextras.Objector
}

var _ ContentProvider = (*contentProvider)(nil)

// WrapContentProvider wraps a GObject to the right type. It is
// primarily used internally.
func WrapContentProvider(obj *externglib.Object) ContentProvider {
	return contentProvider{
		Objector: obj,
	}
}

func marshalContentProvider(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapContentProvider(obj), nil
}

// NewContentProviderForValue constructs a class ContentProvider.
func NewContentProviderForValue(value **externglib.Value) ContentProvider {
	var _arg1 *C.GValue // out

	_arg1 = (*C.GValue)(value.GValue)

	var _cret C.GdkContentProvider // in

	_cret = C.gdk_content_provider_new_for_value(_arg1)

	var _contentProvider ContentProvider // out

	_contentProvider = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ContentProvider)

	return _contentProvider
}

// NewContentProviderUnion constructs a class ContentProvider.
func NewContentProviderUnion(providers []ContentProvider) ContentProvider {
	var _arg1 **C.GdkContentProvider
	var _arg2 C.gsize

	_arg2 = C.gsize(len(providers))
	_arg1 = (**C.GdkContentProvider)(C.malloc(C.ulong(len(providers)) * C.ulong(unsafe.Sizeof(uint(0)))))
	{
		out := unsafe.Slice(_arg1, len(providers))
		for i := range providers {
			out[i] = (*C.GdkContentProvider)(unsafe.Pointer(providers[i].Native()))
		}
	}

	var _cret C.GdkContentProvider // in

	_cret = C.gdk_content_provider_new_union(_arg1, _arg2)

	var _contentProvider ContentProvider // out

	_contentProvider = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ContentProvider)

	return _contentProvider
}

// ContentChanged emits the ::content-changed signal.
func (p contentProvider) ContentChanged() {
	var _arg0 *C.GdkContentProvider // out

	_arg0 = (*C.GdkContentProvider)(unsafe.Pointer(p.Native()))

	C.gdk_content_provider_content_changed(_arg0)
}

// Value gets the contents of @provider stored in @value.
//
// The @value will have been initialized to the `GType` the value should be
// provided in. This given `GType` does not need to be listed in the formats
// returned by [method@Gdk.ContentProvider.ref_formats]. However, if the
// given `GType` is not supported, this operation can fail and
// IO_ERROR_NOT_SUPPORTED will be reported.
func (p contentProvider) Value(value **externglib.Value) error {
	var _arg0 *C.GdkContentProvider // out
	var _arg1 *C.GValue             // out

	_arg0 = (*C.GdkContentProvider)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GValue)(value.GValue)

	var _cerr *C.GError // in

	C.gdk_content_provider_get_value(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// RefFormats gets the formats that the provider can provide its current
// contents in.
func (p contentProvider) RefFormats() *ContentFormats {
	var _arg0 *C.GdkContentProvider // out

	_arg0 = (*C.GdkContentProvider)(unsafe.Pointer(p.Native()))

	var _cret *C.GdkContentFormats // in

	_cret = C.gdk_content_provider_ref_formats(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = WrapContentFormats(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_contentFormats, func(v *ContentFormats) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _contentFormats
}

// RefStorableFormats gets the formats that the provider suggests other
// applications to store the data in.
//
// An example of such an application would be a clipboard manager.
//
// This can be assumed to be a subset of
// [method@Gdk.ContentProvider.ref_formats].
func (p contentProvider) RefStorableFormats() *ContentFormats {
	var _arg0 *C.GdkContentProvider // out

	_arg0 = (*C.GdkContentProvider)(unsafe.Pointer(p.Native()))

	var _cret *C.GdkContentFormats // in

	_cret = C.gdk_content_provider_ref_storable_formats(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = WrapContentFormats(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_contentFormats, func(v *ContentFormats) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return _contentFormats
}

// WriteMIMETypeFinish finishes an asynchronous write operation.
//
// See [method@Gdk.ContentProvider.write_mime_type_async].
func (p contentProvider) WriteMIMETypeFinish(result gio.AsyncResult) error {
	var _arg0 *C.GdkContentProvider // out
	var _arg1 *C.GAsyncResult       // out

	_arg0 = (*C.GdkContentProvider)(unsafe.Pointer(p.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var _cerr *C.GError // in

	C.gdk_content_provider_write_mime_type_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}
