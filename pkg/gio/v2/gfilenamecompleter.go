// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
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
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_filename_completer_get_type()), F: marshalFilenameCompleter},
	})
}

// FilenameCompleter completes partial file and directory names given a partial
// string by looking in the file system for clues. Can return a list of possible
// completion strings for widget implementations.
type FilenameCompleter interface {
	gextras.Objector

	// CompletionSuffix obtains a completion for @initial_text from @completer.
	CompletionSuffix(initialText string) string
	// Completions gets an array of completion strings for a given initial text.
	Completions(initialText string) []string
	// SetDirsOnly: if @dirs_only is true, @completer will only complete
	// directory names, and not file names.
	SetDirsOnly(dirsOnly bool)
}

// filenameCompleter implements the FilenameCompleter class.
type filenameCompleter struct {
	gextras.Objector
}

var _ FilenameCompleter = (*filenameCompleter)(nil)

// WrapFilenameCompleter wraps a GObject to the right type. It is
// primarily used internally.
func WrapFilenameCompleter(obj *externglib.Object) FilenameCompleter {
	return filenameCompleter{
		Objector: obj,
	}
}

func marshalFilenameCompleter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFilenameCompleter(obj), nil
}

// NewFilenameCompleter constructs a class FilenameCompleter.
func NewFilenameCompleter() FilenameCompleter {
	var _cret C.GFilenameCompleter // in

	_cret = C.g_filename_completer_new()

	var _filenameCompleter FilenameCompleter // out

	_filenameCompleter = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(FilenameCompleter)

	return _filenameCompleter
}

// CompletionSuffix obtains a completion for @initial_text from @completer.
func (c filenameCompleter) CompletionSuffix(initialText string) string {
	var _arg0 *C.GFilenameCompleter // out
	var _arg1 *C.char               // out

	_arg0 = (*C.GFilenameCompleter)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(initialText))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.char // in

	_cret = C.g_filename_completer_get_completion_suffix(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Completions gets an array of completion strings for a given initial text.
func (c filenameCompleter) Completions(initialText string) []string {
	var _arg0 *C.GFilenameCompleter // out
	var _arg1 *C.char               // out

	_arg0 = (*C.GFilenameCompleter)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(initialText))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret **C.char

	_cret = C.g_filename_completer_get_completions(_arg0, _arg1)

	var _utf8s []string

	{
		var i int
		for p := _cret; *p != nil; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// SetDirsOnly: if @dirs_only is true, @completer will only complete
// directory names, and not file names.
func (c filenameCompleter) SetDirsOnly(dirsOnly bool) {
	var _arg0 *C.GFilenameCompleter // out
	var _arg1 C.gboolean            // out

	_arg0 = (*C.GFilenameCompleter)(unsafe.Pointer(c.Native()))
	if dirsOnly {
		_arg1 = C.TRUE
	}

	C.g_filename_completer_set_dirs_only(_arg0, _arg1)
}
