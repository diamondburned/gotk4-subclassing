// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_script_get_type()), F: marshalScript},
		{T: externglib.Type(C.pango_script_iter_get_type()), F: marshalScriptIter},
	})
}

// Script: PangoScript enumeration identifies different writing systems.
//
// The values correspond to the names as defined in the Unicode standard. See
// Unicode Standard Annex 24: Script names
// (http://www.unicode.org/reports/tr24/)
//
// Note that this enumeration is deprecated and will not be updated to include
// values in newer versions of the Unicode standard. Applications should use the
// GUnicodeScript enumeration instead, whose values are interchangeable with
// PangoScript.
type Script int

const (
	// ScriptInvalidCode: value never returned from pango_script_for_unichar()
	ScriptInvalidCode Script = -1
	// ScriptCommon: character used by multiple different scripts
	ScriptCommon Script = 0
	// ScriptInherited: mark glyph that takes its script from the base glyph to
	// which it is attached
	ScriptInherited Script = 1
	// ScriptArabic: arabic
	ScriptArabic Script = 2
	// ScriptArmenian: armenian
	ScriptArmenian Script = 3
	// ScriptBengali: bengali
	ScriptBengali Script = 4
	// ScriptBopomofo: bopomofo
	ScriptBopomofo Script = 5
	// ScriptCherokee: cherokee
	ScriptCherokee Script = 6
	// ScriptCoptic: coptic
	ScriptCoptic Script = 7
	// ScriptCyrillic: cyrillic
	ScriptCyrillic Script = 8
	// ScriptDeseret: deseret
	ScriptDeseret Script = 9
	// ScriptDevanagari: devanagari
	ScriptDevanagari Script = 10
	// ScriptEthiopic: ethiopic
	ScriptEthiopic Script = 11
	// ScriptGeorgian: georgian
	ScriptGeorgian Script = 12
	// ScriptGothic: gothic
	ScriptGothic Script = 13
	// ScriptGreek: greek
	ScriptGreek Script = 14
	// ScriptGujarati: gujarati
	ScriptGujarati Script = 15
	// ScriptGurmukhi: gurmukhi
	ScriptGurmukhi Script = 16
	// ScriptHan: han
	ScriptHan Script = 17
	// ScriptHangul: hangul
	ScriptHangul Script = 18
	// ScriptHebrew: hebrew
	ScriptHebrew Script = 19
	// ScriptHiragana: hiragana
	ScriptHiragana Script = 20
	// ScriptKannada: kannada
	ScriptKannada Script = 21
	// ScriptKatakana: katakana
	ScriptKatakana Script = 22
	// ScriptKhmer: khmer
	ScriptKhmer Script = 23
	// ScriptLao: lao
	ScriptLao Script = 24
	// ScriptLatin: latin
	ScriptLatin Script = 25
	// ScriptMalayalam: malayalam
	ScriptMalayalam Script = 26
	// ScriptMongolian: mongolian
	ScriptMongolian Script = 27
	// ScriptMyanmar: myanmar
	ScriptMyanmar Script = 28
	// ScriptOgham: ogham
	ScriptOgham Script = 29
	// ScriptOldItalic: old Italic
	ScriptOldItalic Script = 30
	// ScriptOriya: oriya
	ScriptOriya Script = 31
	// ScriptRunic: runic
	ScriptRunic Script = 32
	// ScriptSinhala: sinhala
	ScriptSinhala Script = 33
	// ScriptSyriac: syriac
	ScriptSyriac Script = 34
	// ScriptTamil: tamil
	ScriptTamil Script = 35
	// ScriptTelugu: telugu
	ScriptTelugu Script = 36
	// ScriptThaana: thaana
	ScriptThaana Script = 37
	// ScriptThai: thai
	ScriptThai Script = 38
	// ScriptTibetan: tibetan
	ScriptTibetan Script = 39
	// ScriptCanadianAboriginal: canadian Aboriginal
	ScriptCanadianAboriginal Script = 40
	// ScriptYi: yi
	ScriptYi Script = 41
	// ScriptTagalog: tagalog
	ScriptTagalog Script = 42
	// ScriptHanunoo: hanunoo
	ScriptHanunoo Script = 43
	// ScriptBuhid: buhid
	ScriptBuhid Script = 44
	// ScriptTagbanwa: tagbanwa
	ScriptTagbanwa Script = 45
	// ScriptBraille: braille
	ScriptBraille Script = 46
	// ScriptCypriot: cypriot
	ScriptCypriot Script = 47
	// ScriptLimbu: limbu
	ScriptLimbu Script = 48
	// ScriptOsmanya: osmanya
	ScriptOsmanya Script = 49
	// ScriptShavian: shavian
	ScriptShavian Script = 50
	// ScriptLinearB: linear B
	ScriptLinearB Script = 51
	// ScriptTaiLe: tai Le
	ScriptTaiLe Script = 52
	// ScriptUgaritic: ugaritic
	ScriptUgaritic Script = 53
	// ScriptNewTaiLue: new Tai Lue. Since 1.10
	ScriptNewTaiLue Script = 54
	// ScriptBuginese: buginese. Since 1.10
	ScriptBuginese Script = 55
	// ScriptGlagolitic: glagolitic. Since 1.10
	ScriptGlagolitic Script = 56
	// ScriptTifinagh: tifinagh. Since 1.10
	ScriptTifinagh Script = 57
	// ScriptSylotiNagri: syloti Nagri. Since 1.10
	ScriptSylotiNagri Script = 58
	// ScriptOldPersian: old Persian. Since 1.10
	ScriptOldPersian Script = 59
	// ScriptKharoshthi: kharoshthi. Since 1.10
	ScriptKharoshthi Script = 60
	// ScriptUnknown: unassigned code point. Since 1.14
	ScriptUnknown Script = 61
	// ScriptBalinese: balinese. Since 1.14
	ScriptBalinese Script = 62
	// ScriptCuneiform: cuneiform. Since 1.14
	ScriptCuneiform Script = 63
	// ScriptPhoenician: phoenician. Since 1.14
	ScriptPhoenician Script = 64
	// ScriptPhagsPa: phags-pa. Since 1.14
	ScriptPhagsPa Script = 65
	// ScriptNko: n'Ko. Since 1.14
	ScriptNko Script = 66
	// ScriptKayahLi: kayah Li. Since 1.20.1
	ScriptKayahLi Script = 67
	// ScriptLepcha: lepcha. Since 1.20.1
	ScriptLepcha Script = 68
	// ScriptRejang: rejang. Since 1.20.1
	ScriptRejang Script = 69
	// ScriptSundanese: sundanese. Since 1.20.1
	ScriptSundanese Script = 70
	// ScriptSaurashtra: saurashtra. Since 1.20.1
	ScriptSaurashtra Script = 71
	// ScriptCham: cham. Since 1.20.1
	ScriptCham Script = 72
	// ScriptOlChiki: ol Chiki. Since 1.20.1
	ScriptOlChiki Script = 73
	// ScriptVai: vai. Since 1.20.1
	ScriptVai Script = 74
	// ScriptCarian: carian. Since 1.20.1
	ScriptCarian Script = 75
	// ScriptLycian: lycian. Since 1.20.1
	ScriptLycian Script = 76
	// ScriptLydian: lydian. Since 1.20.1
	ScriptLydian Script = 77
	// ScriptBatak: batak. Since 1.32
	ScriptBatak Script = 78
	// ScriptBrahmi: brahmi. Since 1.32
	ScriptBrahmi Script = 79
	// ScriptMandaic: mandaic. Since 1.32
	ScriptMandaic Script = 80
	// ScriptChakma: chakma. Since: 1.32
	ScriptChakma Script = 81
	// ScriptMeroiticCursive: meroitic Cursive. Since: 1.32
	ScriptMeroiticCursive Script = 82
	// ScriptMeroiticHieroglyphs: meroitic Hieroglyphs. Since: 1.32
	ScriptMeroiticHieroglyphs Script = 83
	// ScriptMiao: miao. Since: 1.32
	ScriptMiao Script = 84
	// ScriptSharada: sharada. Since: 1.32
	ScriptSharada Script = 85
	// ScriptSoraSompeng: sora Sompeng. Since: 1.32
	ScriptSoraSompeng Script = 86
	// ScriptTakri: takri. Since: 1.32
	ScriptTakri Script = 87
	// ScriptBassaVah: bassa. Since: 1.40
	ScriptBassaVah Script = 88
	// ScriptCaucasianAlbanian: caucasian Albanian. Since: 1.40
	ScriptCaucasianAlbanian Script = 89
	// ScriptDuployan: duployan. Since: 1.40
	ScriptDuployan Script = 90
	// ScriptElbasan: elbasan. Since: 1.40
	ScriptElbasan Script = 91
	// ScriptGrantha: grantha. Since: 1.40
	ScriptGrantha Script = 92
	// ScriptKhojki: kjohki. Since: 1.40
	ScriptKhojki Script = 93
	// ScriptKhudawadi: khudawadi, Sindhi. Since: 1.40
	ScriptKhudawadi Script = 94
	// ScriptLinearA: linear A. Since: 1.40
	ScriptLinearA Script = 95
	// ScriptMahajani: mahajani. Since: 1.40
	ScriptMahajani Script = 96
	// ScriptManichaean: manichaean. Since: 1.40
	ScriptManichaean Script = 97
	// ScriptMendeKikakui: mende Kikakui. Since: 1.40
	ScriptMendeKikakui Script = 98
	// ScriptModi: modi. Since: 1.40
	ScriptModi Script = 99
	// ScriptMro: mro. Since: 1.40
	ScriptMro Script = 100
	// ScriptNabataean: nabataean. Since: 1.40
	ScriptNabataean Script = 101
	// ScriptOldNorthArabian: old North Arabian. Since: 1.40
	ScriptOldNorthArabian Script = 102
	// ScriptOldPermic: old Permic. Since: 1.40
	ScriptOldPermic Script = 103
	// ScriptPahawhHmong: pahawh Hmong. Since: 1.40
	ScriptPahawhHmong Script = 104
	// ScriptPalmyrene: palmyrene. Since: 1.40
	ScriptPalmyrene Script = 105
	// ScriptPauCinHau: pau Cin Hau. Since: 1.40
	ScriptPauCinHau Script = 106
	// ScriptPsalterPahlavi: psalter Pahlavi. Since: 1.40
	ScriptPsalterPahlavi Script = 107
	// ScriptSiddham: siddham. Since: 1.40
	ScriptSiddham Script = 108
	// ScriptTirhuta: tirhuta. Since: 1.40
	ScriptTirhuta Script = 109
	// ScriptWarangCiti: warang Citi. Since: 1.40
	ScriptWarangCiti Script = 110
	// ScriptAhom: ahom. Since: 1.40
	ScriptAhom Script = 111
	// ScriptAnatolianHieroglyphs: anatolian Hieroglyphs. Since: 1.40
	ScriptAnatolianHieroglyphs Script = 112
	// ScriptHatran: hatran. Since: 1.40
	ScriptHatran Script = 113
	// ScriptMultani: multani. Since: 1.40
	ScriptMultani Script = 114
	// ScriptOldHungarian: old Hungarian. Since: 1.40
	ScriptOldHungarian Script = 115
	// ScriptSignwriting: signwriting. Since: 1.40
	ScriptSignwriting Script = 116
)

func marshalScript(p uintptr) (interface{}, error) {
	return Script(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for Script.
func (s Script) String() string {
	switch s {
	case ScriptInvalidCode:
		return "InvalidCode"
	case ScriptCommon:
		return "Common"
	case ScriptInherited:
		return "Inherited"
	case ScriptArabic:
		return "Arabic"
	case ScriptArmenian:
		return "Armenian"
	case ScriptBengali:
		return "Bengali"
	case ScriptBopomofo:
		return "Bopomofo"
	case ScriptCherokee:
		return "Cherokee"
	case ScriptCoptic:
		return "Coptic"
	case ScriptCyrillic:
		return "Cyrillic"
	case ScriptDeseret:
		return "Deseret"
	case ScriptDevanagari:
		return "Devanagari"
	case ScriptEthiopic:
		return "Ethiopic"
	case ScriptGeorgian:
		return "Georgian"
	case ScriptGothic:
		return "Gothic"
	case ScriptGreek:
		return "Greek"
	case ScriptGujarati:
		return "Gujarati"
	case ScriptGurmukhi:
		return "Gurmukhi"
	case ScriptHan:
		return "Han"
	case ScriptHangul:
		return "Hangul"
	case ScriptHebrew:
		return "Hebrew"
	case ScriptHiragana:
		return "Hiragana"
	case ScriptKannada:
		return "Kannada"
	case ScriptKatakana:
		return "Katakana"
	case ScriptKhmer:
		return "Khmer"
	case ScriptLao:
		return "Lao"
	case ScriptLatin:
		return "Latin"
	case ScriptMalayalam:
		return "Malayalam"
	case ScriptMongolian:
		return "Mongolian"
	case ScriptMyanmar:
		return "Myanmar"
	case ScriptOgham:
		return "Ogham"
	case ScriptOldItalic:
		return "OldItalic"
	case ScriptOriya:
		return "Oriya"
	case ScriptRunic:
		return "Runic"
	case ScriptSinhala:
		return "Sinhala"
	case ScriptSyriac:
		return "Syriac"
	case ScriptTamil:
		return "Tamil"
	case ScriptTelugu:
		return "Telugu"
	case ScriptThaana:
		return "Thaana"
	case ScriptThai:
		return "Thai"
	case ScriptTibetan:
		return "Tibetan"
	case ScriptCanadianAboriginal:
		return "CanadianAboriginal"
	case ScriptYi:
		return "Yi"
	case ScriptTagalog:
		return "Tagalog"
	case ScriptHanunoo:
		return "Hanunoo"
	case ScriptBuhid:
		return "Buhid"
	case ScriptTagbanwa:
		return "Tagbanwa"
	case ScriptBraille:
		return "Braille"
	case ScriptCypriot:
		return "Cypriot"
	case ScriptLimbu:
		return "Limbu"
	case ScriptOsmanya:
		return "Osmanya"
	case ScriptShavian:
		return "Shavian"
	case ScriptLinearB:
		return "LinearB"
	case ScriptTaiLe:
		return "TaiLe"
	case ScriptUgaritic:
		return "Ugaritic"
	case ScriptNewTaiLue:
		return "NewTaiLue"
	case ScriptBuginese:
		return "Buginese"
	case ScriptGlagolitic:
		return "Glagolitic"
	case ScriptTifinagh:
		return "Tifinagh"
	case ScriptSylotiNagri:
		return "SylotiNagri"
	case ScriptOldPersian:
		return "OldPersian"
	case ScriptKharoshthi:
		return "Kharoshthi"
	case ScriptUnknown:
		return "Unknown"
	case ScriptBalinese:
		return "Balinese"
	case ScriptCuneiform:
		return "Cuneiform"
	case ScriptPhoenician:
		return "Phoenician"
	case ScriptPhagsPa:
		return "PhagsPa"
	case ScriptNko:
		return "Nko"
	case ScriptKayahLi:
		return "KayahLi"
	case ScriptLepcha:
		return "Lepcha"
	case ScriptRejang:
		return "Rejang"
	case ScriptSundanese:
		return "Sundanese"
	case ScriptSaurashtra:
		return "Saurashtra"
	case ScriptCham:
		return "Cham"
	case ScriptOlChiki:
		return "OlChiki"
	case ScriptVai:
		return "Vai"
	case ScriptCarian:
		return "Carian"
	case ScriptLycian:
		return "Lycian"
	case ScriptLydian:
		return "Lydian"
	case ScriptBatak:
		return "Batak"
	case ScriptBrahmi:
		return "Brahmi"
	case ScriptMandaic:
		return "Mandaic"
	case ScriptChakma:
		return "Chakma"
	case ScriptMeroiticCursive:
		return "MeroiticCursive"
	case ScriptMeroiticHieroglyphs:
		return "MeroiticHieroglyphs"
	case ScriptMiao:
		return "Miao"
	case ScriptSharada:
		return "Sharada"
	case ScriptSoraSompeng:
		return "SoraSompeng"
	case ScriptTakri:
		return "Takri"
	case ScriptBassaVah:
		return "BassaVah"
	case ScriptCaucasianAlbanian:
		return "CaucasianAlbanian"
	case ScriptDuployan:
		return "Duployan"
	case ScriptElbasan:
		return "Elbasan"
	case ScriptGrantha:
		return "Grantha"
	case ScriptKhojki:
		return "Khojki"
	case ScriptKhudawadi:
		return "Khudawadi"
	case ScriptLinearA:
		return "LinearA"
	case ScriptMahajani:
		return "Mahajani"
	case ScriptManichaean:
		return "Manichaean"
	case ScriptMendeKikakui:
		return "MendeKikakui"
	case ScriptModi:
		return "Modi"
	case ScriptMro:
		return "Mro"
	case ScriptNabataean:
		return "Nabataean"
	case ScriptOldNorthArabian:
		return "OldNorthArabian"
	case ScriptOldPermic:
		return "OldPermic"
	case ScriptPahawhHmong:
		return "PahawhHmong"
	case ScriptPalmyrene:
		return "Palmyrene"
	case ScriptPauCinHau:
		return "PauCinHau"
	case ScriptPsalterPahlavi:
		return "PsalterPahlavi"
	case ScriptSiddham:
		return "Siddham"
	case ScriptTirhuta:
		return "Tirhuta"
	case ScriptWarangCiti:
		return "WarangCiti"
	case ScriptAhom:
		return "Ahom"
	case ScriptAnatolianHieroglyphs:
		return "AnatolianHieroglyphs"
	case ScriptHatran:
		return "Hatran"
	case ScriptMultani:
		return "Multani"
	case ScriptOldHungarian:
		return "OldHungarian"
	case ScriptSignwriting:
		return "Signwriting"
	default:
		return fmt.Sprintf("Script(%d)", s)
	}
}

// ScriptForUnichar looks up the script for a particular character.
//
// The script of a character is defined by Unicode Standard Annex \#24. No check
// is made for ch being a valid Unicode character; if you pass in invalid
// character, the result is undefined.
//
// Note that while the return type of this function is declared as PangoScript,
// as of Pango 1.18, this function simply returns the return value of
// g_unichar_get_script(). Callers must be prepared to handle unknown values.
//
// Deprecated: Use g_unichar_get_script().
func ScriptForUnichar(ch uint32) Script {
	var _arg1 C.gunichar    // out
	var _cret C.PangoScript // in

	_arg1 = C.gunichar(ch)

	_cret = C.pango_script_for_unichar(_arg1)

	var _script Script // out

	_script = Script(_cret)

	return _script
}

// ScriptGetSampleLanguage finds a language tag that is reasonably
// representative of script.
//
// The language will usually be the most widely spoken or used language written
// in that script: for instance, the sample language for PANGO_SCRIPT_CYRILLIC
// is ru (Russian), the sample language for PANGO_SCRIPT_ARABIC is ar.
//
// For some scripts, no sample language will be returned because there is no
// language that is sufficiently representative. The best example of this is
// PANGO_SCRIPT_HAN, where various different variants of written Chinese,
// Japanese, and Korean all use significantly different sets of Han characters
// and forms of shared characters. No sample language can be provided for many
// historical scripts as well.
//
// As of 1.18, this function checks the environment variables PANGO_LANGUAGE and
// LANGUAGE (checked in that order) first. If one of them is set, it is parsed
// as a list of language tags separated by colons or other separators. This
// function will return the first language in the parsed list that Pango
// believes may use script for writing. This last predicate is tested using
// pango.Language.IncludesScript(). This can be used to control Pango's font
// selection for non-primary languages. For example, a PANGO_LANGUAGE enviroment
// variable set to "en:fa" makes Pango choose fonts suitable for Persian (fa)
// instead of Arabic (ar) when a segment of Arabic text is found in an otherwise
// non-Arabic text. The same trick can be used to choose a default language for
// PANGO_SCRIPT_HAN when setting context language is not feasible.
func ScriptGetSampleLanguage(script Script) *Language {
	var _arg1 C.PangoScript    // out
	var _cret *C.PangoLanguage // in

	_arg1 = C.PangoScript(script)

	_cret = C.pango_script_get_sample_language(_arg1)

	var _language *Language // out

	_language = (*Language)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_language, func(v *Language) {
		C.free(gextras.StructNative(unsafe.Pointer(v)))
	})

	return _language
}

// ScriptIter: PangoScriptIter is used to iterate through a string and identify
// ranges in different scripts.
type ScriptIter struct {
	nocopy gextras.NoCopy
	native *C.PangoScriptIter
}

func marshalScriptIter(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &ScriptIter{native: (*C.PangoScriptIter)(unsafe.Pointer(b))}, nil
}

// NewScriptIter constructs a struct ScriptIter.
func NewScriptIter(text string, length int) *ScriptIter {
	var _arg1 *C.char            // out
	var _arg2 C.int              // out
	var _cret *C.PangoScriptIter // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(length)

	_cret = C.pango_script_iter_new(_arg1, _arg2)

	var _scriptIter *ScriptIter // out

	_scriptIter = (*ScriptIter)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_scriptIter, func(v *ScriptIter) {
		C.pango_script_iter_free((*C.PangoScriptIter)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _scriptIter
}

// Range gets information about the range to which iter currently points. The
// range is the set of locations p where *start <= p < *end. (That is, it
// doesn't include the character stored at *end)
//
// Note that while the type of the script argument is declared as PangoScript,
// as of Pango 1.18, this function simply returns GUnicodeScript values. Callers
// must be prepared to handle unknown values.
func (iter *ScriptIter) Range() (start string, end string, script Script) {
	var _arg0 *C.PangoScriptIter // out
	var _arg1 *C.char            // in
	var _arg2 *C.char            // in
	var _arg3 C.PangoScript      // in

	_arg0 = (*C.PangoScriptIter)(gextras.StructNative(unsafe.Pointer(iter)))

	C.pango_script_iter_get_range(_arg0, &_arg1, &_arg2, &_arg3)

	var _start string  // out
	var _end string    // out
	var _script Script // out

	if _arg1 != nil {
		_start = C.GoString((*C.gchar)(unsafe.Pointer(_arg1)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	if _arg2 != nil {
		_end = C.GoString((*C.gchar)(unsafe.Pointer(_arg2)))
		defer C.free(unsafe.Pointer(_arg2))
	}
	_script = Script(_arg3)

	return _start, _end, _script
}

// Next advances a ScriptIter to the next range. If iter is already at the end,
// it is left unchanged and FALSE is returned.
func (iter *ScriptIter) Next() bool {
	var _arg0 *C.PangoScriptIter // out
	var _cret C.gboolean         // in

	_arg0 = (*C.PangoScriptIter)(gextras.StructNative(unsafe.Pointer(iter)))

	_cret = C.pango_script_iter_next(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
