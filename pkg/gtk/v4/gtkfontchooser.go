// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/pango"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_font_chooser_get_type()), F: marshalFontChooser},
	})
}

// FontChooserOverrider contains methods that are overridable. This
// interface is a subset of the interface FontChooser.
type FontChooserOverrider interface {
	FontActivated(fontname string)
	// FontSize: the selected font size.
	FontSize() int
	// SetFontMap sets a custom font map to use for this font chooser widget.
	//
	// A custom font map can be used to present application-specific fonts
	// instead of or in addition to the normal system fonts.
	//
	// “`c FcConfig *config; PangoFontMap *fontmap;
	//
	// config = FcInitLoadConfigAndFonts (); FcConfigAppFontAddFile (config,
	// my_app_font_file);
	//
	// fontmap = pango_cairo_font_map_new_for_font_type (CAIRO_FONT_TYPE_FT);
	// pango_fc_font_map_set_config (PANGO_FC_FONT_MAP (fontmap), config);
	//
	// gtk_font_chooser_set_font_map (font_chooser, fontmap); “`
	//
	// Note that other GTK widgets will only be able to use the
	// application-specific font if it is present in the font map they use:
	//
	// “`c context = gtk_widget_get_pango_context (label);
	// pango_context_set_font_map (context, fontmap); “`
	SetFontMap(fontmap pango.FontMap)
}

// FontChooser: `GtkFontChooser` is an interface that can be implemented by
// widgets for choosing fonts.
//
// In GTK, the main objects that implement this interface are
// [class@Gtk.FontChooserWidget], [class@Gtk.FontChooserDialog] and
// [class@Gtk.FontButton].
type FontChooser interface {
	gextras.Objector
	FontChooserOverrider

	// Font gets the currently-selected font name.
	//
	// Note that this can be a different string than what you set with
	// [method@Gtk.FontChooser.set_font], as the font chooser widget may
	// normalize font names and thus return a string with a different structure.
	// For example, “Helvetica Italic Bold 12” could be normalized to “Helvetica
	// Bold Italic 12”.
	//
	// Use [method@Pango.FontDescription.equal] if you want to compare two font
	// descriptions.
	Font() string
	// FontFeatures gets the currently-selected font features.
	FontFeatures() string
	// Language gets the language that is used for font features.
	Language() string
	// PreviewText gets the text displayed in the preview area.
	PreviewText() string
	// ShowPreviewEntry returns whether the preview entry is shown or not.
	ShowPreviewEntry() bool
	// SetFont sets the currently-selected font.
	SetFont(fontname string)
	// SetFontDesc sets the currently-selected font from @font_desc.
	SetFontDesc(fontDesc *pango.FontDescription)
	// SetLanguage sets the language to use for font features.
	SetLanguage(language string)
	// SetLevel sets the desired level of granularity for selecting fonts.
	SetLevel(level FontChooserLevel)
	// SetPreviewText sets the text displayed in the preview area.
	//
	// The @text is used to show how the selected font looks.
	SetPreviewText(text string)
	// SetShowPreviewEntry shows or hides the editable preview entry.
	SetShowPreviewEntry(showPreviewEntry bool)
}

// fontChooser implements the FontChooser interface.
type fontChooser struct {
	gextras.Objector
}

var _ FontChooser = (*fontChooser)(nil)

// WrapFontChooser wraps a GObject to a type that implements interface
// FontChooser. It is primarily used internally.
func WrapFontChooser(obj *externglib.Object) FontChooser {
	return FontChooser{
		Objector: obj,
	}
}

func marshalFontChooser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontChooser(obj), nil
}

// Font gets the currently-selected font name.
//
// Note that this can be a different string than what you set with
// [method@Gtk.FontChooser.set_font], as the font chooser widget may
// normalize font names and thus return a string with a different structure.
// For example, “Helvetica Italic Bold 12” could be normalized to “Helvetica
// Bold Italic 12”.
//
// Use [method@Pango.FontDescription.equal] if you want to compare two font
// descriptions.
func (f fontChooser) Font() string {
	var _arg0 *C.GtkFontChooser // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	var _cret *C.char // in

	_cret = C.gtk_font_chooser_get_font(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// FontFeatures gets the currently-selected font features.
func (f fontChooser) FontFeatures() string {
	var _arg0 *C.GtkFontChooser // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	var _cret *C.char // in

	_cret = C.gtk_font_chooser_get_font_features(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// FontSize: the selected font size.
func (f fontChooser) FontSize() int {
	var _arg0 *C.GtkFontChooser // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	var _cret C.int // in

	_cret = C.gtk_font_chooser_get_font_size(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Language gets the language that is used for font features.
func (f fontChooser) Language() string {
	var _arg0 *C.GtkFontChooser // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	var _cret *C.char // in

	_cret = C.gtk_font_chooser_get_language(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// PreviewText gets the text displayed in the preview area.
func (f fontChooser) PreviewText() string {
	var _arg0 *C.GtkFontChooser // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	var _cret *C.char // in

	_cret = C.gtk_font_chooser_get_preview_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// ShowPreviewEntry returns whether the preview entry is shown or not.
func (f fontChooser) ShowPreviewEntry() bool {
	var _arg0 *C.GtkFontChooser // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_font_chooser_get_show_preview_entry(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// SetFont sets the currently-selected font.
func (f fontChooser) SetFont(fontname string) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.char)(C.CString(fontname))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_chooser_set_font(_arg0, _arg1)
}

// SetFontDesc sets the currently-selected font from @font_desc.
func (f fontChooser) SetFontDesc(fontDesc *pango.FontDescription) {
	var _arg0 *C.GtkFontChooser       // out
	var _arg1 *C.PangoFontDescription // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.PangoFontDescription)(unsafe.Pointer(fontDesc.Native()))

	C.gtk_font_chooser_set_font_desc(_arg0, _arg1)
}

// SetFontMap sets a custom font map to use for this font chooser widget.
//
// A custom font map can be used to present application-specific fonts
// instead of or in addition to the normal system fonts.
//
// “`c FcConfig *config; PangoFontMap *fontmap;
//
// config = FcInitLoadConfigAndFonts (); FcConfigAppFontAddFile (config,
// my_app_font_file);
//
// fontmap = pango_cairo_font_map_new_for_font_type (CAIRO_FONT_TYPE_FT);
// pango_fc_font_map_set_config (PANGO_FC_FONT_MAP (fontmap), config);
//
// gtk_font_chooser_set_font_map (font_chooser, fontmap); “`
//
// Note that other GTK widgets will only be able to use the
// application-specific font if it is present in the font map they use:
//
// “`c context = gtk_widget_get_pango_context (label);
// pango_context_set_font_map (context, fontmap); “`
func (f fontChooser) SetFontMap(fontmap pango.FontMap) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 *C.PangoFontMap   // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.PangoFontMap)(unsafe.Pointer(fontmap.Native()))

	C.gtk_font_chooser_set_font_map(_arg0, _arg1)
}

// SetLanguage sets the language to use for font features.
func (f fontChooser) SetLanguage(language string) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.char)(C.CString(language))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_chooser_set_language(_arg0, _arg1)
}

// SetLevel sets the desired level of granularity for selecting fonts.
func (f fontChooser) SetLevel(level FontChooserLevel) {
	var _arg0 *C.GtkFontChooser     // out
	var _arg1 C.GtkFontChooserLevel // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	_arg1 = (C.GtkFontChooserLevel)(level)

	C.gtk_font_chooser_set_level(_arg0, _arg1)
}

// SetPreviewText sets the text displayed in the preview area.
//
// The @text is used to show how the selected font looks.
func (f fontChooser) SetPreviewText(text string) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 *C.char           // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_font_chooser_set_preview_text(_arg0, _arg1)
}

// SetShowPreviewEntry shows or hides the editable preview entry.
func (f fontChooser) SetShowPreviewEntry(showPreviewEntry bool) {
	var _arg0 *C.GtkFontChooser // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkFontChooser)(unsafe.Pointer(f.Native()))
	if showPreviewEntry {
		_arg1 = C.gboolean(1)
	}

	C.gtk_font_chooser_set_show_preview_entry(_arg0, _arg1)
}
