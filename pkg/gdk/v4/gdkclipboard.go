// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gerror"
	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_clipboard_get_type()), F: marshalClipboard},
	})
}

// Clipboard: the `GdkClipboard` object represents data shared between
// applications or inside an application.
//
// To get a `GdkClipboard` object, use [method@Gdk.Display.get_clipboard] or
// [method@Gdk.Display.get_primary_clipboard]. You can find out about the data
// that is currently available in a clipboard using
// [method@Gdk.Clipboard.get_formats].
//
// To make text or image data available in a clipboard, use
// [method@Gdk.Clipboard.set_text] or [method@Gdk.Clipboard.set_texture]. For
// other data, you can use [method@Gdk.Clipboard.set_content], which takes a
// [class@Gdk.ContentProvider] object.
//
// To read textual or image data from a clipboard, use
// [method@Gdk.Clipboard.read_text_async] or
// [method@Gdk.Clipboard.read_texture_async]. For other data, use
// [method@Gdk.Clipboard.read_async], which provides a `GInputStream` object.
type Clipboard interface {
	gextras.Objector

	// Content:
	Content() ContentProvider
	// Display:
	Display() Display
	// Formats:
	Formats() *ContentFormats
	// IsLocalClipboard:
	IsLocalClipboard() bool
	// ReadFinishClipboard:
	ReadFinishClipboard(result gio.AsyncResult) (string, gio.InputStream, error)
	// ReadTextFinishClipboard:
	ReadTextFinishClipboard(result gio.AsyncResult) (string, error)
	// ReadTextureFinishClipboard:
	ReadTextureFinishClipboard(result gio.AsyncResult) (Texture, error)
	// ReadValueFinishClipboard:
	ReadValueFinishClipboard(result gio.AsyncResult) (externglib.Value, error)
	// SetContentClipboard:
	SetContentClipboard(provider ContentProvider) bool
	// SetValueClipboard:
	SetValueClipboard(value externglib.Value)
	// StoreFinishClipboard:
	StoreFinishClipboard(result gio.AsyncResult) error
}

// clipboard implements the Clipboard class.
type clipboard struct {
	gextras.Objector
}

// WrapClipboard wraps a GObject to the right type. It is
// primarily used internally.
func WrapClipboard(obj *externglib.Object) Clipboard {
	return clipboard{
		Objector: obj,
	}
}

func marshalClipboard(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapClipboard(obj), nil
}

func (c clipboard) Content() ContentProvider {
	var _arg0 *C.GdkClipboard       // out
	var _cret *C.GdkContentProvider // in

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_clipboard_get_content(_arg0)

	var _contentProvider ContentProvider // out

	_contentProvider = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(ContentProvider)

	return _contentProvider
}

func (c clipboard) Display() Display {
	var _arg0 *C.GdkClipboard // out
	var _cret *C.GdkDisplay   // in

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_clipboard_get_display(_arg0)

	var _display Display // out

	_display = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Display)

	return _display
}

func (c clipboard) Formats() *ContentFormats {
	var _arg0 *C.GdkClipboard      // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_clipboard_get_formats(_arg0)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(unsafe.Pointer(_cret))

	return _contentFormats
}

func (c clipboard) IsLocalClipboard() bool {
	var _arg0 *C.GdkClipboard // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(c.Native()))

	_cret = C.gdk_clipboard_is_local(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c clipboard) ReadFinishClipboard(result gio.AsyncResult) (string, gio.InputStream, error) {
	var _arg0 *C.GdkClipboard // out
	var _arg1 *C.GAsyncResult // out
	var _arg2 *C.char         // in
	var _cret *C.GInputStream // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.gdk_clipboard_read_finish(_arg0, _arg1, &_arg2, &_cerr)

	var _outMimeType string          // out
	var _inputStream gio.InputStream // out
	var _goerr error                 // out

	_outMimeType = C.GoString(_arg2)
	_inputStream = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(gio.InputStream)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _outMimeType, _inputStream, _goerr
}

func (c clipboard) ReadTextFinishClipboard(result gio.AsyncResult) (string, error) {
	var _arg0 *C.GdkClipboard // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.char         // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.gdk_clipboard_read_text_finish(_arg0, _arg1, &_cerr)

	var _utf8 string // out
	var _goerr error // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _utf8, _goerr
}

func (c clipboard) ReadTextureFinishClipboard(result gio.AsyncResult) (Texture, error) {
	var _arg0 *C.GdkClipboard // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GdkTexture   // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.gdk_clipboard_read_texture_finish(_arg0, _arg1, &_cerr)

	var _texture Texture // out
	var _goerr error     // out

	_texture = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(Texture)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _texture, _goerr
}

func (c clipboard) ReadValueFinishClipboard(result gio.AsyncResult) (externglib.Value, error) {
	var _arg0 *C.GdkClipboard // out
	var _arg1 *C.GAsyncResult // out
	var _cret *C.GValue       // in
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	_cret = C.gdk_clipboard_read_value_finish(_arg0, _arg1, &_cerr)

	var _value externglib.Value // out
	var _goerr error            // out

	_value = externglib.ValueFromNative(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _value, _goerr
}

func (c clipboard) SetContentClipboard(provider ContentProvider) bool {
	var _arg0 *C.GdkClipboard       // out
	var _arg1 *C.GdkContentProvider // out
	var _cret C.gboolean            // in

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkContentProvider)(unsafe.Pointer(provider.Native()))

	_cret = C.gdk_clipboard_set_content(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c clipboard) SetValueClipboard(value externglib.Value) {
	var _arg0 *C.GdkClipboard // out
	var _arg1 *C.GValue       // out

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	C.gdk_clipboard_set_value(_arg0, _arg1)
}

func (c clipboard) StoreFinishClipboard(result gio.AsyncResult) error {
	var _arg0 *C.GdkClipboard // out
	var _arg1 *C.GAsyncResult // out
	var _cerr *C.GError       // in

	_arg0 = (*C.GdkClipboard)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	C.gdk_clipboard_store_finish(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}