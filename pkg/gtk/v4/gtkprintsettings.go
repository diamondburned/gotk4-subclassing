// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
// void _gotk4_gtk4_PrintSettingsFunc(char*, char*, gpointer);
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_print_settings_get_type()), F: marshalPrintSettingser},
	})
}

type PrintSettingsFunc func(key string, value string)

//export _gotk4_gtk4_PrintSettingsFunc
func _gotk4_gtk4_PrintSettingsFunc(arg0 *C.char, arg1 *C.char, arg2 C.gpointer) {
	v := gbox.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var key string   // out
	var value string // out

	key = C.GoString((*C.gchar)(unsafe.Pointer(arg0)))
	defer C.free(unsafe.Pointer(arg0))
	value = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	defer C.free(unsafe.Pointer(arg1))

	fn := v.(PrintSettingsFunc)
	fn(key, value)
}

// PrintSettings: GtkPrintSettings object represents the settings of a print
// dialog in a system-independent way.
//
// The main use for this object is that once you’ve printed you can get a
// settings object that represents the settings the user chose, and the next
// time you print you can pass that object in so that the user doesn’t have to
// re-set all his settings.
//
// Its also possible to enumerate the settings so that you can easily save the
// settings for the next time your app runs, or even store them in a document.
// The predefined keys try to use shared values as much as possible so that
// moving such a document between systems still works.
type PrintSettings struct {
	*externglib.Object
}

func wrapPrintSettings(obj *externglib.Object) *PrintSettings {
	return &PrintSettings{
		Object: obj,
	}
}

func marshalPrintSettingser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapPrintSettings(obj), nil
}

// NewPrintSettings creates a new GtkPrintSettings object.
func NewPrintSettings() *PrintSettings {
	var _cret *C.GtkPrintSettings // in

	_cret = C.gtk_print_settings_new()

	var _printSettings *PrintSettings // out

	_printSettings = wrapPrintSettings(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _printSettings
}

// NewPrintSettingsFromFile reads the print settings from file_name.
//
// Returns a new GtkPrintSettings object with the restored settings, or NULL if
// an error occurred. If the file could not be loaded then error is set to
// either a GFileError or GKeyFileError.
//
// See gtk.PrintSettings.ToFile().
func NewPrintSettingsFromFile(fileName string) (*PrintSettings, error) {
	var _arg1 *C.char             // out
	var _cret *C.GtkPrintSettings // in
	var _cerr *C.GError           // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(fileName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_new_from_file(_arg1, &_cerr)

	var _printSettings *PrintSettings // out
	var _goerr error                  // out

	_printSettings = wrapPrintSettings(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _printSettings, _goerr
}

// NewPrintSettingsFromGVariant: deserialize print settings from an a{sv}
// variant.
//
// The variant must be in the format produced by gtk.PrintSettings.ToGVariant().
func NewPrintSettingsFromGVariant(variant *glib.Variant) *PrintSettings {
	var _arg1 *C.GVariant         // out
	var _cret *C.GtkPrintSettings // in

	_arg1 = (*C.GVariant)(gextras.StructNative(unsafe.Pointer(variant)))

	_cret = C.gtk_print_settings_new_from_gvariant(_arg1)

	var _printSettings *PrintSettings // out

	_printSettings = wrapPrintSettings(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _printSettings
}

// NewPrintSettingsFromKeyFile reads the print settings from the group
// group_name in key_file.
//
// Returns a new GtkPrintSettings object with the restored settings, or NULL if
// an error occurred. If the file could not be loaded then error is set to
// either GFileError or GKeyFileError.
func NewPrintSettingsFromKeyFile(keyFile *glib.KeyFile, groupName string) (*PrintSettings, error) {
	var _arg1 *C.GKeyFile         // out
	var _arg2 *C.char             // out
	var _cret *C.GtkPrintSettings // in
	var _cerr *C.GError           // in

	_arg1 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	if groupName != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(groupName)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.gtk_print_settings_new_from_key_file(_arg1, _arg2, &_cerr)

	var _printSettings *PrintSettings // out
	var _goerr error                  // out

	_printSettings = wrapPrintSettings(externglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _printSettings, _goerr
}

// Copy copies a GtkPrintSettings object.
func (other *PrintSettings) Copy() *PrintSettings {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.GtkPrintSettings // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(other.Native()))

	_cret = C.gtk_print_settings_copy(_arg0)

	var _printSettings *PrintSettings // out

	_printSettings = wrapPrintSettings(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _printSettings
}

// Foreach calls func for each key-value pair of settings.
func (settings *PrintSettings) Foreach(fn PrintSettingsFunc) {
	var _arg0 *C.GtkPrintSettings    // out
	var _arg1 C.GtkPrintSettingsFunc // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*[0]byte)(C._gotk4_gtk4_PrintSettingsFunc)
	_arg2 = C.gpointer(gbox.Assign(fn))
	defer gbox.Delete(uintptr(_arg2))

	C.gtk_print_settings_foreach(_arg0, _arg1, _arg2)
}

// Get looks up the string value associated with key.
func (settings *PrintSettings) Get(key string) string {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_get(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Bool returns the boolean represented by the value that is associated with
// key.
//
// The string “true” represents TRUE, any other string FALSE.
func (settings *PrintSettings) Bool(key string) bool {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_get_bool(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Collate gets the value of GTK_PRINT_SETTINGS_COLLATE.
func (settings *PrintSettings) Collate() bool {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_collate(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DefaultSource gets the value of GTK_PRINT_SETTINGS_DEFAULT_SOURCE.
func (settings *PrintSettings) DefaultSource() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_default_source(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Dither gets the value of GTK_PRINT_SETTINGS_DITHER.
func (settings *PrintSettings) Dither() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_dither(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Double returns the double value associated with key, or 0.
func (settings *PrintSettings) Double(key string) float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_get_double(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// DoubleWithDefault returns the floating point number represented by the value
// that is associated with key, or default_val if the value does not represent a
// floating point number.
//
// Floating point numbers are parsed with g_ascii_strtod().
func (settings *PrintSettings) DoubleWithDefault(key string, def float64) float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.double            // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.double(def)

	_cret = C.gtk_print_settings_get_double_with_default(_arg0, _arg1, _arg2)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Duplex gets the value of GTK_PRINT_SETTINGS_DUPLEX.
func (settings *PrintSettings) Duplex() PrintDuplex {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.GtkPrintDuplex    // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_duplex(_arg0)

	var _printDuplex PrintDuplex // out

	_printDuplex = PrintDuplex(_cret)

	return _printDuplex
}

// Finishings gets the value of GTK_PRINT_SETTINGS_FINISHINGS.
func (settings *PrintSettings) Finishings() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_finishings(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Int returns the integer value of key, or 0.
func (settings *PrintSettings) Int(key string) int {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_get_int(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// IntWithDefault returns the value of key, interpreted as an integer, or the
// default value.
func (settings *PrintSettings) IntWithDefault(key string, def int) int {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.int               // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(def)

	_cret = C.gtk_print_settings_get_int_with_default(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Length returns the value associated with key, interpreted as a length.
//
// The returned value is converted to units.
func (settings *PrintSettings) Length(key string, unit Unit) float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.GtkUnit           // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GtkUnit(unit)

	_cret = C.gtk_print_settings_get_length(_arg0, _arg1, _arg2)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// MediaType gets the value of GTK_PRINT_SETTINGS_MEDIA_TYPE.
//
// The set of media types is defined in PWG 5101.1-2002 PWG.
func (settings *PrintSettings) MediaType() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_media_type(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// NCopies gets the value of GTK_PRINT_SETTINGS_N_COPIES.
func (settings *PrintSettings) NCopies() int {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_n_copies(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NumberUp gets the value of GTK_PRINT_SETTINGS_NUMBER_UP.
func (settings *PrintSettings) NumberUp() int {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_number_up(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NumberUpLayout gets the value of GTK_PRINT_SETTINGS_NUMBER_UP_LAYOUT.
func (settings *PrintSettings) NumberUpLayout() NumberUpLayout {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.GtkNumberUpLayout // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_number_up_layout(_arg0)

	var _numberUpLayout NumberUpLayout // out

	_numberUpLayout = NumberUpLayout(_cret)

	return _numberUpLayout
}

// Orientation: get the value of GTK_PRINT_SETTINGS_ORIENTATION, converted to a
// GtkPageOrientation.
func (settings *PrintSettings) Orientation() PageOrientation {
	var _arg0 *C.GtkPrintSettings  // out
	var _cret C.GtkPageOrientation // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_orientation(_arg0)

	var _pageOrientation PageOrientation // out

	_pageOrientation = PageOrientation(_cret)

	return _pageOrientation
}

// OutputBin gets the value of GTK_PRINT_SETTINGS_OUTPUT_BIN.
func (settings *PrintSettings) OutputBin() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_output_bin(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// PageSet gets the value of GTK_PRINT_SETTINGS_PAGE_SET.
func (settings *PrintSettings) PageSet() PageSet {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.GtkPageSet        // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_page_set(_arg0)

	var _pageSet PageSet // out

	_pageSet = PageSet(_cret)

	return _pageSet
}

// PaperHeight gets the value of GTK_PRINT_SETTINGS_PAPER_HEIGHT, converted to
// unit.
func (settings *PrintSettings) PaperHeight(unit Unit) float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkUnit           // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_print_settings_get_paper_height(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// PaperSize gets the value of GTK_PRINT_SETTINGS_PAPER_FORMAT, converted to a
// GtkPaperSize.
func (settings *PrintSettings) PaperSize() *PaperSize {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.GtkPaperSize     // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_paper_size(_arg0)

	var _paperSize *PaperSize // out

	_paperSize = (*PaperSize)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_paperSize, func(v *PaperSize) {
		C.gtk_paper_size_free((*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _paperSize
}

// PaperWidth gets the value of GTK_PRINT_SETTINGS_PAPER_WIDTH, converted to
// unit.
func (settings *PrintSettings) PaperWidth(unit Unit) float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkUnit           // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = C.GtkUnit(unit)

	_cret = C.gtk_print_settings_get_paper_width(_arg0, _arg1)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// PrintPages gets the value of GTK_PRINT_SETTINGS_PRINT_PAGES.
func (settings *PrintSettings) PrintPages() PrintPages {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.GtkPrintPages     // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_print_pages(_arg0)

	var _printPages PrintPages // out

	_printPages = PrintPages(_cret)

	return _printPages
}

// Printer: convenience function to obtain the value of
// GTK_PRINT_SETTINGS_PRINTER.
func (settings *PrintSettings) Printer() string {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.char             // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_printer(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// PrinterLpi gets the value of GTK_PRINT_SETTINGS_PRINTER_LPI.
func (settings *PrintSettings) PrinterLpi() float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_printer_lpi(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// Quality gets the value of GTK_PRINT_SETTINGS_QUALITY.
func (settings *PrintSettings) Quality() PrintQuality {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.GtkPrintQuality   // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_quality(_arg0)

	var _printQuality PrintQuality // out

	_printQuality = PrintQuality(_cret)

	return _printQuality
}

// Resolution gets the value of GTK_PRINT_SETTINGS_RESOLUTION.
func (settings *PrintSettings) Resolution() int {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_resolution(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ResolutionX gets the value of GTK_PRINT_SETTINGS_RESOLUTION_X.
func (settings *PrintSettings) ResolutionX() int {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_resolution_x(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// ResolutionY gets the value of GTK_PRINT_SETTINGS_RESOLUTION_Y.
func (settings *PrintSettings) ResolutionY() int {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.int               // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_resolution_y(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Reverse gets the value of GTK_PRINT_SETTINGS_REVERSE.
func (settings *PrintSettings) Reverse() bool {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_reverse(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Scale gets the value of GTK_PRINT_SETTINGS_SCALE.
func (settings *PrintSettings) Scale() float64 {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.double            // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_scale(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// UseColor gets the value of GTK_PRINT_SETTINGS_USE_COLOR.
func (settings *PrintSettings) UseColor() bool {
	var _arg0 *C.GtkPrintSettings // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_get_use_color(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasKey returns TRUE, if a value is associated with key.
func (settings *PrintSettings) HasKey(key string) bool {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_print_settings_has_key(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LoadFile reads the print settings from file_name.
//
// If the file could not be loaded then error is set to either a GFileError or
// GKeyFileError.
//
// See gtk.PrintSettings.ToFile().
func (settings *PrintSettings) LoadFile(fileName string) error {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(fileName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_load_file(_arg0, _arg1, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// LoadKeyFile reads the print settings from the group group_name in key_file.
//
// If the file could not be loaded then error is set to either a GFileError or
// GKeyFileError.
func (settings *PrintSettings) LoadKeyFile(keyFile *glib.KeyFile, groupName string) error {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.GKeyFile         // out
	var _arg2 *C.char             // out
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	if groupName != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(groupName)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	C.gtk_print_settings_load_key_file(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Set associates value with key.
func (settings *PrintSettings) Set(key string, value string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))
	if value != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(value)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	C.gtk_print_settings_set(_arg0, _arg1, _arg2)
}

// SetBool sets key to a boolean value.
func (settings *PrintSettings) SetBool(key string, value bool) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.gboolean          // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))
	if value {
		_arg2 = C.TRUE
	}

	C.gtk_print_settings_set_bool(_arg0, _arg1, _arg2)
}

// SetCollate sets the value of GTK_PRINT_SETTINGS_COLLATE.
func (settings *PrintSettings) SetCollate(collate bool) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	if collate {
		_arg1 = C.TRUE
	}

	C.gtk_print_settings_set_collate(_arg0, _arg1)
}

// SetDefaultSource sets the value of GTK_PRINT_SETTINGS_DEFAULT_SOURCE.
func (settings *PrintSettings) SetDefaultSource(defaultSource string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(defaultSource)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_default_source(_arg0, _arg1)
}

// SetDither sets the value of GTK_PRINT_SETTINGS_DITHER.
func (settings *PrintSettings) SetDither(dither string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(dither)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_dither(_arg0, _arg1)
}

// SetDouble sets key to a double value.
func (settings *PrintSettings) SetDouble(key string, value float64) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.double            // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.double(value)

	C.gtk_print_settings_set_double(_arg0, _arg1, _arg2)
}

// SetDuplex sets the value of GTK_PRINT_SETTINGS_DUPLEX.
func (settings *PrintSettings) SetDuplex(duplex PrintDuplex) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkPrintDuplex    // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = C.GtkPrintDuplex(duplex)

	C.gtk_print_settings_set_duplex(_arg0, _arg1)
}

// SetFinishings sets the value of GTK_PRINT_SETTINGS_FINISHINGS.
func (settings *PrintSettings) SetFinishings(finishings string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(finishings)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_finishings(_arg0, _arg1)
}

// SetInt sets key to an integer value.
func (settings *PrintSettings) SetInt(key string, value int) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.int               // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(value)

	C.gtk_print_settings_set_int(_arg0, _arg1, _arg2)
}

// SetLength associates a length in units of unit with key.
func (settings *PrintSettings) SetLength(key string, value float64, unit Unit) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _arg2 C.double            // out
	var _arg3 C.GtkUnit           // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.double(value)
	_arg3 = C.GtkUnit(unit)

	C.gtk_print_settings_set_length(_arg0, _arg1, _arg2, _arg3)
}

// SetMediaType sets the value of GTK_PRINT_SETTINGS_MEDIA_TYPE.
//
// The set of media types is defined in PWG 5101.1-2002 PWG.
func (settings *PrintSettings) SetMediaType(mediaType string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(mediaType)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_media_type(_arg0, _arg1)
}

// SetNCopies sets the value of GTK_PRINT_SETTINGS_N_COPIES.
func (settings *PrintSettings) SetNCopies(numCopies int) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.int               // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = C.int(numCopies)

	C.gtk_print_settings_set_n_copies(_arg0, _arg1)
}

// SetNumberUp sets the value of GTK_PRINT_SETTINGS_NUMBER_UP.
func (settings *PrintSettings) SetNumberUp(numberUp int) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.int               // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = C.int(numberUp)

	C.gtk_print_settings_set_number_up(_arg0, _arg1)
}

// SetNumberUpLayout sets the value of GTK_PRINT_SETTINGS_NUMBER_UP_LAYOUT.
func (settings *PrintSettings) SetNumberUpLayout(numberUpLayout NumberUpLayout) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkNumberUpLayout // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = C.GtkNumberUpLayout(numberUpLayout)

	C.gtk_print_settings_set_number_up_layout(_arg0, _arg1)
}

// SetOrientation sets the value of GTK_PRINT_SETTINGS_ORIENTATION.
func (settings *PrintSettings) SetOrientation(orientation PageOrientation) {
	var _arg0 *C.GtkPrintSettings  // out
	var _arg1 C.GtkPageOrientation // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = C.GtkPageOrientation(orientation)

	C.gtk_print_settings_set_orientation(_arg0, _arg1)
}

// SetOutputBin sets the value of GTK_PRINT_SETTINGS_OUTPUT_BIN.
func (settings *PrintSettings) SetOutputBin(outputBin string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(outputBin)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_output_bin(_arg0, _arg1)
}

// SetPageRanges sets the value of GTK_PRINT_SETTINGS_PAGE_RANGES.
func (settings *PrintSettings) SetPageRanges(pageRanges []PageRange) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.GtkPageRange     // out
	var _arg2 C.int

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg2 = (C.int)(len(pageRanges))
	if len(pageRanges) > 0 {
		_arg1 = (*C.GtkPageRange)(unsafe.Pointer(&pageRanges[0]))
	}

	C.gtk_print_settings_set_page_ranges(_arg0, _arg1, _arg2)
}

// SetPageSet sets the value of GTK_PRINT_SETTINGS_PAGE_SET.
func (settings *PrintSettings) SetPageSet(pageSet PageSet) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkPageSet        // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = C.GtkPageSet(pageSet)

	C.gtk_print_settings_set_page_set(_arg0, _arg1)
}

// SetPaperHeight sets the value of GTK_PRINT_SETTINGS_PAPER_HEIGHT.
func (settings *PrintSettings) SetPaperHeight(height float64, unit Unit) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.double            // out
	var _arg2 C.GtkUnit           // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = C.double(height)
	_arg2 = C.GtkUnit(unit)

	C.gtk_print_settings_set_paper_height(_arg0, _arg1, _arg2)
}

// SetPaperSize sets the value of GTK_PRINT_SETTINGS_PAPER_FORMAT,
// GTK_PRINT_SETTINGS_PAPER_WIDTH and GTK_PRINT_SETTINGS_PAPER_HEIGHT.
func (settings *PrintSettings) SetPaperSize(paperSize *PaperSize) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.GtkPaperSize     // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.GtkPaperSize)(gextras.StructNative(unsafe.Pointer(paperSize)))

	C.gtk_print_settings_set_paper_size(_arg0, _arg1)
}

// SetPaperWidth sets the value of GTK_PRINT_SETTINGS_PAPER_WIDTH.
func (settings *PrintSettings) SetPaperWidth(width float64, unit Unit) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.double            // out
	var _arg2 C.GtkUnit           // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = C.double(width)
	_arg2 = C.GtkUnit(unit)

	C.gtk_print_settings_set_paper_width(_arg0, _arg1, _arg2)
}

// SetPrintPages sets the value of GTK_PRINT_SETTINGS_PRINT_PAGES.
func (settings *PrintSettings) SetPrintPages(pages PrintPages) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkPrintPages     // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = C.GtkPrintPages(pages)

	C.gtk_print_settings_set_print_pages(_arg0, _arg1)
}

// SetPrinter: convenience function to set GTK_PRINT_SETTINGS_PRINTER to
// printer.
func (settings *PrintSettings) SetPrinter(printer string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(printer)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_set_printer(_arg0, _arg1)
}

// SetPrinterLpi sets the value of GTK_PRINT_SETTINGS_PRINTER_LPI.
func (settings *PrintSettings) SetPrinterLpi(lpi float64) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.double            // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = C.double(lpi)

	C.gtk_print_settings_set_printer_lpi(_arg0, _arg1)
}

// SetQuality sets the value of GTK_PRINT_SETTINGS_QUALITY.
func (settings *PrintSettings) SetQuality(quality PrintQuality) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.GtkPrintQuality   // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = C.GtkPrintQuality(quality)

	C.gtk_print_settings_set_quality(_arg0, _arg1)
}

// SetResolution sets the values of GTK_PRINT_SETTINGS_RESOLUTION,
// GTK_PRINT_SETTINGS_RESOLUTION_X and GTK_PRINT_SETTINGS_RESOLUTION_Y.
func (settings *PrintSettings) SetResolution(resolution int) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.int               // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = C.int(resolution)

	C.gtk_print_settings_set_resolution(_arg0, _arg1)
}

// SetResolutionXY sets the values of GTK_PRINT_SETTINGS_RESOLUTION,
// GTK_PRINT_SETTINGS_RESOLUTION_X and GTK_PRINT_SETTINGS_RESOLUTION_Y.
func (settings *PrintSettings) SetResolutionXY(resolutionX int, resolutionY int) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.int               // out
	var _arg2 C.int               // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = C.int(resolutionX)
	_arg2 = C.int(resolutionY)

	C.gtk_print_settings_set_resolution_xy(_arg0, _arg1, _arg2)
}

// SetReverse sets the value of GTK_PRINT_SETTINGS_REVERSE.
func (settings *PrintSettings) SetReverse(reverse bool) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	if reverse {
		_arg1 = C.TRUE
	}

	C.gtk_print_settings_set_reverse(_arg0, _arg1)
}

// SetScale sets the value of GTK_PRINT_SETTINGS_SCALE.
func (settings *PrintSettings) SetScale(scale float64) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.double            // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = C.double(scale)

	C.gtk_print_settings_set_scale(_arg0, _arg1)
}

// SetUseColor sets the value of GTK_PRINT_SETTINGS_USE_COLOR.
func (settings *PrintSettings) SetUseColor(useColor bool) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	if useColor {
		_arg1 = C.TRUE
	}

	C.gtk_print_settings_set_use_color(_arg0, _arg1)
}

// ToFile: this function saves the print settings from settings to file_name.
//
// If the file could not be written then error is set to either a GFileError or
// GKeyFileError.
func (settings *PrintSettings) ToFile(fileName string) error {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out
	var _cerr *C.GError           // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(fileName)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_to_file(_arg0, _arg1, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// ToGVariant: serialize print settings to an a{sv} variant.
func (settings *PrintSettings) ToGVariant() *glib.Variant {
	var _arg0 *C.GtkPrintSettings // out
	var _cret *C.GVariant         // in

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))

	_cret = C.gtk_print_settings_to_gvariant(_arg0)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(gextras.StructNative(unsafe.Pointer(v))))
	})

	return _variant
}

// ToKeyFile: this function adds the print settings from settings to key_file.
func (settings *PrintSettings) ToKeyFile(keyFile *glib.KeyFile, groupName string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.GKeyFile         // out
	var _arg2 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.GKeyFile)(gextras.StructNative(unsafe.Pointer(keyFile)))
	if groupName != "" {
		_arg2 = (*C.char)(unsafe.Pointer(C.CString(groupName)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	C.gtk_print_settings_to_key_file(_arg0, _arg1, _arg2)
}

// Unset removes any value associated with key.
//
// This has the same effect as setting the value to NULL.
func (settings *PrintSettings) Unset(key string) {
	var _arg0 *C.GtkPrintSettings // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkPrintSettings)(unsafe.Pointer(settings.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_print_settings_unset(_arg0, _arg1)
}

// PageRange: range of pages to print.
//
// See also gtk.PrintSettings.SetPageRanges().
type PageRange struct {
	nocopy gextras.NoCopy
	native *C.GtkPageRange
}

// Start: start of page range.
func (p *PageRange) Start() int {
	var v int // out
	v = int(p.native.start)
	return v
}

// End: end of page range.
func (p *PageRange) End() int {
	var v int // out
	v = int(p.native.end)
	return v
}
