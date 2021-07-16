// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// ConvertError: error codes returned by character set conversion routines.
type ConvertError int

const (
	// ConvertErrorNoConversion: conversion between the requested character sets
	// is not supported.
	ConvertErrorNoConversion ConvertError = iota
	// ConvertErrorIllegalSequence: invalid byte sequence in conversion input;
	// or the character sequence could not be represented in the target
	// character set.
	ConvertErrorIllegalSequence
	// ConvertErrorFailed: conversion failed for some reason.
	ConvertErrorFailed
	// ConvertErrorPartialInput: partial character sequence at end of input.
	ConvertErrorPartialInput
	// ConvertErrorBadURI: URI is invalid.
	ConvertErrorBadURI
	// ConvertErrorNotAbsolutePath: pathname is not an absolute path.
	ConvertErrorNotAbsolutePath
	// ConvertErrorNoMemory: no memory available. Since: 2.40
	ConvertErrorNoMemory
	// ConvertErrorEmbeddedNUL: embedded NUL character is present in conversion
	// output where a NUL-terminated string is expected. Since: 2.56
	ConvertErrorEmbeddedNUL
)

// String returns the name in string for ConvertError.
func (c ConvertError) String() string {
	switch c {
	case ConvertErrorNoConversion:
		return "NoConversion"
	case ConvertErrorIllegalSequence:
		return "IllegalSequence"
	case ConvertErrorFailed:
		return "Failed"
	case ConvertErrorPartialInput:
		return "PartialInput"
	case ConvertErrorBadURI:
		return "BadURI"
	case ConvertErrorNotAbsolutePath:
		return "NotAbsolutePath"
	case ConvertErrorNoMemory:
		return "NoMemory"
	case ConvertErrorEmbeddedNUL:
		return "EmbeddedNUL"
	default:
		return fmt.Sprintf("ConvertError(%d)", c)
	}
}

// FilenameDisplayBasename returns the display basename for the particular
// filename, guaranteed to be valid UTF-8. The display name might not be
// identical to the filename, for instance there might be problems converting it
// to UTF-8, and some files can be translated in the display.
//
// If GLib cannot make sense of the encoding of filename, as a last resort it
// replaces unknown characters with U+FFFD, the Unicode replacement character.
// You can search the result for the UTF-8 encoding of this character (which is
// "\357\277\275" in octal notation) to find out if filename was in an invalid
// encoding.
//
// You must pass the whole absolute pathname to this functions so that
// translation of well known locations can be done.
//
// This function is preferred over g_filename_display_name() if you know the
// whole path, as it allows translation.
func FilenameDisplayBasename(filename string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))

	_cret = C.g_filename_display_basename(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// FilenameDisplayName converts a filename into a valid UTF-8 string. The
// conversion is not necessarily reversible, so you should keep the original
// around and use the return value of this function only for display purposes.
// Unlike g_filename_to_utf8(), the result is guaranteed to be non-NULL even if
// the filename actually isn't in the GLib file name encoding.
//
// If GLib cannot make sense of the encoding of filename, as a last resort it
// replaces unknown characters with U+FFFD, the Unicode replacement character.
// You can search the result for the UTF-8 encoding of this character (which is
// "\357\277\275" in octal notation) to find out if filename was in an invalid
// encoding.
//
// If you know the whole pathname of the file you should use
// g_filename_display_basename(), since that allows location-based translation
// of filenames.
func FilenameDisplayName(filename string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))

	_cret = C.g_filename_display_name(_arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// FilenameFromURI converts an escaped ASCII-encoded URI to a local filename in
// the encoding used for filenames.
func FilenameFromURI(uri string) (hostname string, filename string, goerr error) {
	var _arg1 *C.gchar  // out
	var _arg2 *C.gchar  // in
	var _cret *C.gchar  // in
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))

	_cret = C.g_filename_from_uri(_arg1, &_arg2, &_cerr)

	var _hostname string // out
	var _filename string // out
	var _goerr error     // out

	_hostname = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
	defer C.free(unsafe.Pointer(_arg2))
	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _hostname, _filename, _goerr
}

// FilenameFromUTF8 converts a string from UTF-8 to the encoding GLib uses for
// filenames. Note that on Windows GLib uses UTF-8 for filenames; on other
// platforms, this function indirectly depends on the [current
// locale][setlocale].
//
// The input string shall not contain nul characters even if the len argument is
// positive. A nul character found inside the string will result in error
// G_CONVERT_ERROR_ILLEGAL_SEQUENCE. If the filename encoding is not UTF-8 and
// the conversion output contains a nul character, the error
// G_CONVERT_ERROR_EMBEDDED_NUL is set and the function returns NULL.
func FilenameFromUTF8(utf8String string, len int) (bytesRead uint, bytesWritten uint, filename string, goerr error) {
	var _arg1 *C.gchar  // out
	var _arg2 C.gssize  // out
	var _arg3 C.gsize   // in
	var _arg4 C.gsize   // in
	var _cret *C.gchar  // in
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(utf8String)))
	_arg2 = C.gssize(len)

	_cret = C.g_filename_from_utf8(_arg1, _arg2, &_arg3, &_arg4, &_cerr)

	var _bytesRead uint    // out
	var _bytesWritten uint // out
	var _filename string   // out
	var _goerr error       // out

	_bytesRead = uint(_arg3)
	_bytesWritten = uint(_arg4)
	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _bytesRead, _bytesWritten, _filename, _goerr
}

// FilenameToURI converts an absolute filename to an escaped ASCII-encoded URI,
// with the path component following Section 3.3. of RFC 2396.
func FilenameToURI(filename string, hostname string) (string, error) {
	var _arg1 *C.gchar  // out
	var _arg2 *C.gchar  // out
	var _cret *C.gchar  // in
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(filename)))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(hostname)))

	_cret = C.g_filename_to_uri(_arg1, _arg2, &_cerr)

	var _utf8 string // out
	var _goerr error // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _utf8, _goerr
}

// FilenameToUTF8 converts a string which is in the encoding used by GLib for
// filenames into a UTF-8 string. Note that on Windows GLib uses UTF-8 for
// filenames; on other platforms, this function indirectly depends on the
// [current locale][setlocale].
//
// The input string shall not contain nul characters even if the len argument is
// positive. A nul character found inside the string will result in error
// G_CONVERT_ERROR_ILLEGAL_SEQUENCE. If the source encoding is not UTF-8 and the
// conversion output contains a nul character, the error
// G_CONVERT_ERROR_EMBEDDED_NUL is set and the function returns NULL. Use
// g_convert() to produce output that may contain embedded nul characters.
func FilenameToUTF8(opsysstring string, len int) (bytesRead uint, bytesWritten uint, utf8 string, goerr error) {
	var _arg1 *C.gchar  // out
	var _arg2 C.gssize  // out
	var _arg3 C.gsize   // in
	var _arg4 C.gsize   // in
	var _cret *C.gchar  // in
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(opsysstring)))
	_arg2 = C.gssize(len)

	_cret = C.g_filename_to_utf8(_arg1, _arg2, &_arg3, &_arg4, &_cerr)

	var _bytesRead uint    // out
	var _bytesWritten uint // out
	var _utf8 string       // out
	var _goerr error       // out

	_bytesRead = uint(_arg3)
	_bytesWritten = uint(_arg4)
	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _bytesRead, _bytesWritten, _utf8, _goerr
}

// GetFilenameCharsets determines the preferred character sets used for
// filenames. The first character set from the charsets is the filename
// encoding, the subsequent character sets are used when trying to generate a
// displayable representation of a filename, see g_filename_display_name().
//
// On Unix, the character sets are determined by consulting the environment
// variables G_FILENAME_ENCODING and G_BROKEN_FILENAMES. On Windows, the
// character set used in the GLib API is always UTF-8 and said environment
// variables have no effect.
//
// G_FILENAME_ENCODING may be set to a comma-separated list of character set
// names. The special token "\locale" is taken to mean the character set for the
// [current locale][setlocale]. If G_FILENAME_ENCODING is not set, but
// G_BROKEN_FILENAMES is, the character set of the current locale is taken as
// the filename encoding. If neither environment variable is set, UTF-8 is taken
// as the filename encoding, but the character set of the current locale is also
// put in the list of encodings.
//
// The returned charsets belong to GLib and must not be freed.
//
// Note that on Unix, regardless of the locale character set or
// G_FILENAME_ENCODING value, the actual file names present on a system might be
// in any random encoding or just gibberish.
func GetFilenameCharsets() ([]string, bool) {
	var _arg1 **C.gchar
	var _cret C.gboolean // in

	_cret = C.g_get_filename_charsets(&_arg1)

	var _filenameCharsets []string
	var _ok bool // out

	{
		var i int
		var z *C.gchar
		for p := _arg1; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_arg1, i)
		_filenameCharsets = make([]string, i)
		for i := range src {
			_filenameCharsets[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}
	if _cret != 0 {
		_ok = true
	}

	return _filenameCharsets, _ok
}

// LocaleToUTF8 converts a string which is in the encoding used for strings by
// the C runtime (usually the same as that used by the operating system) in the
// [current locale][setlocale] into a UTF-8 string.
//
// If the source encoding is not UTF-8 and the conversion output contains a nul
// character, the error G_CONVERT_ERROR_EMBEDDED_NUL is set and the function
// returns NULL. If the source encoding is UTF-8, an embedded nul character is
// treated with the G_CONVERT_ERROR_ILLEGAL_SEQUENCE error for backward
// compatibility with earlier versions of this library. Use g_convert() to
// produce output that may contain embedded nul characters.
func LocaleToUTF8(opsysstring []byte) (bytesRead uint, bytesWritten uint, utf8 string, goerr error) {
	var _arg1 *C.gchar
	var _arg2 C.gssize
	var _arg3 C.gsize   // in
	var _arg4 C.gsize   // in
	var _cret *C.gchar  // in
	var _cerr *C.GError // in

	_arg2 = (C.gssize)(len(opsysstring))
	if len(opsysstring) > 0 {
		_arg1 = (*C.gchar)(unsafe.Pointer(&opsysstring[0]))
	}

	_cret = C.g_locale_to_utf8(_arg1, _arg2, &_arg3, &_arg4, &_cerr)

	var _bytesRead uint    // out
	var _bytesWritten uint // out
	var _utf8 string       // out
	var _goerr error       // out

	_bytesRead = uint(_arg3)
	_bytesWritten = uint(_arg4)
	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _bytesRead, _bytesWritten, _utf8, _goerr
}

// URIListExtractUris splits an URI list conforming to the text/uri-list mime
// type defined in RFC 2483 into individual URIs, discarding any comments. The
// URIs are not validated.
func UriListExtractUris(uriList string) []string {
	var _arg1 *C.gchar // out
	var _cret **C.gchar

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uriList)))

	_cret = C.g_uri_list_extract_uris(_arg1)

	var _utf8s []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}
