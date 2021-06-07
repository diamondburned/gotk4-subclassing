// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_language_get_type()), F: marshalLanguage},
	})
}

// LanguageFromString: convert a language tag to a `PangoLanguage`.
//
// The language tag must be in a RFC-3066 format. `PangoLanguage` pointers can
// be efficiently copied (copy the pointer) and compared with other language
// tags (compare the pointer.)
//
// This function first canonicalizes the string by converting it to lowercase,
// mapping '_' to '-', and stripping all characters other than letters and '-'.
//
// Use [type_func@Pango.Language.get_default] if you want to get the
// `PangoLanguage` for the current locale of the process.
func LanguageFromString(language string) {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(language))
	defer C.free(unsafe.Pointer(arg1))

	C.pango_language_from_string(arg1)
}

// LanguageGetDefault returns the `PangoLanguage` for the current locale of the
// process.
//
// On Unix systems, this is the return value is derived from `setlocale
// (LC_CTYPE, NULL)`, and the user can affect this through the environment
// variables LC_ALL, LC_CTYPE or LANG (checked in that order). The locale string
// typically is in the form lang_COUNTRY, where lang is an ISO-639 language
// code, and COUNTRY is an ISO-3166 country code. For instance, sv_FI for
// Swedish as written in Finland or pt_BR for Portuguese as written in Brazil.
//
// On Windows, the C library does not use any such environment variables, and
// setting them won't affect the behavior of functions like ctime(). The user
// sets the locale through the Regional Options in the Control Panel. The C
// library (in the setlocale() function) does not use country and language
// codes, but country and language names spelled out in English. However, this
// function does check the above environment variables, and does return a
// Unix-style locale string based on either said environment variables or the
// thread's current locale.
//
// Your application should call `setlocale(LC_ALL, "")` for the user settings to
// take effect. GTK does this in its initialization functions automatically (by
// calling gtk_set_locale()). See the setlocale() manpage for more details.
//
// Note that the default language can change over the life of an application.
func LanguageGetDefault() {
	C.pango_language_get_default()
}

// LanguageGetPreferred returns the list of languages that the user prefers.
//
// The list is specified by the `PANGO_LANGUAGE` or `LANGUAGE` environment
// variables, in order of preference. Note that this list does not necessarily
// include the language returned by [type_func@Pango.Language.get_default].
//
// When choosing language-specific resources, such as the sample text returned
// by [method@Pango.Language.get_sample_string], you should first try the
// default language, followed by the languages returned by this function.
func LanguageGetPreferred() {
	C.pango_language_get_preferred()
}

// Language: the `PangoLanguage` structure is used to represent a language.
//
// `PangoLanguage` pointers can be efficiently copied and compared with each
// other.
type Language struct {
	native C.PangoLanguage
}

// WrapLanguage wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapLanguage(ptr unsafe.Pointer) *Language {
	if ptr == nil {
		return nil
	}

	return (*Language)(ptr)
}

func marshalLanguage(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapLanguage(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (l *Language) Native() unsafe.Pointer {
	return unsafe.Pointer(&l.native)
}

// SampleString: get a string that is representative of the characters needed to
// render a particular language.
//
// The sample text may be a pangram, but is not necessarily. It is chosen to be
// demonstrative of normal text in the language, as well as exposing font
// feature requirements unique to the language. It is suitable for use as sample
// text in a font selection dialog.
//
// If @language is nil, the default language as found by
// [type_func@Pango.Language.get_default] is used.
//
// If Pango does not have a sample string for @language, the classic "The quick
// brown fox..." is returned. This can be detected by comparing the returned
// pointer value to that returned for (non-existent) language code "xx". That
// is, compare to:
//
// “` pango_language_get_sample_string (pango_language_from_string ("xx")) “`
func (l *Language) SampleString(l *Language) {
	var arg0 *C.PangoLanguage

	arg0 = (*C.PangoLanguage)(unsafe.Pointer(l.Native()))

	C.pango_language_get_sample_string(arg0)
}

// Scripts determines the scripts used to to write @language.
//
// If nothing is known about the language tag @language, or if @language is nil,
// then nil is returned. The list of scripts returned starts with the script
// that the language uses most and continues to the one it uses least.
//
// The value @num_script points at will be set to the number of scripts in the
// returned array (or zero if nil is returned).
//
// Most languages use only one script for writing, but there are some that use
// two (Latin and Cyrillic for example), and a few use three (Japanese for
// example). Applications should not make any assumptions on the maximum number
// of scripts returned though, except that it is positive if the return value is
// not nil, and it is a small number.
//
// The [method@Pango.Language.includes_script] function uses this function
// internally.
//
// Note: while the return value is declared as `PangoScript`, the returned
// values are from the `GUnicodeScript` enumeration, which may have more values.
// Callers need to handle unknown values.
func (l *Language) Scripts(l *Language) int {
	var arg0 *C.PangoLanguage

	arg0 = (*C.PangoLanguage)(unsafe.Pointer(l.Native()))

	var arg1 C.int
	var numScripts int

	C.pango_language_get_scripts(arg0, &arg1)

	numScripts = int(&arg1)

	return numScripts
}

// IncludesScript determines if @script is one of the scripts used to write
// @language. The returned value is conservative; if nothing is known about the
// language tag @language, true will be returned, since, as far as Pango knows,
// @script might be used to write @language.
//
// This routine is used in Pango's itemization process when determining if a
// supplied language tag is relevant to a particular section of text. It
// probably is not useful for applications in most circumstances.
//
// This function uses [method@Pango.Language.get_scripts] internally.
func (l *Language) IncludesScript(l *Language, script Script) bool {
	var arg0 *C.PangoLanguage
	var arg1 C.PangoScript

	arg0 = (*C.PangoLanguage)(unsafe.Pointer(l.Native()))
	arg1 = (C.PangoScript)(script)

	var cret C.gboolean
	var ok bool

	cret = C.pango_language_includes_script(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// Matches checks if a language tag matches one of the elements in a list of
// language ranges.
//
// A language tag is considered to match a range in the list if the range is
// '*', the range is exactly the tag, or the range is a prefix of the tag, and
// the character after it in the tag is '-'.
func (l *Language) Matches(l *Language, rangeList string) bool {
	var arg0 *C.PangoLanguage
	var arg1 *C.char

	arg0 = (*C.PangoLanguage)(unsafe.Pointer(l.Native()))
	arg1 = (*C.char)(C.CString(rangeList))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ok bool

	cret = C.pango_language_matches(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// String gets the RFC-3066 format string representing the given language tag.
func (l *Language) String(l *Language) {
	var arg0 *C.PangoLanguage

	arg0 = (*C.PangoLanguage)(unsafe.Pointer(l.Native()))

	C.pango_language_to_string(arg0)
}
