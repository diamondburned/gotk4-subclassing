// Code generated by girgen. DO NOT EDIT.

package pangoft2

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/freetype2"
	"github.com/diamondburned/gotk4/pkg/pango"
	"github.com/diamondburned/gotk4/pkg/pangofc"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pangoft2
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pangoft2.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{

		// Classes
		{T: externglib.Type(C.pango_ft2_font_map_get_type()), F: marshalFontMap},
	})
}

// FontGetCoverage: gets the Coverage for a `PangoFT2Font`. Use
// pango_font_get_coverage() instead.
func FontGetCoverage(font pango.Font, language *pango.Language) pango.Coverage {
	var arg0 pango.Font
	arg0 = wrapFont(font)

	var arg1 *pango.Language
	arg1 = wrapLanguage(language)

	c0 := C.pango_ft2_font_get_coverage(arg0, arg1)

	var ret0 pango.Coverage
	ret0 = wrapCoverage(c0)

	return ret0
}

// FontGetFace: returns the native FreeType2 `FT_Face` structure used for this
// Font. This may be useful if you want to use FreeType2 functions directly.
//
// Use pango_fc_font_lock_face() instead; when you are done with a face from
// pango_fc_font_lock_face() you must call pango_fc_font_unlock_face().
func FontGetFace(font pango.Font) freetype2.Face {
	var arg0 pango.Font
	arg0 = wrapFont(font)

	c0 := C.pango_ft2_font_get_face(arg0)

	var ret0 freetype2.Face
	ret0 = wrapFace(c0)

	return ret0
}

// FontGetKerning: retrieves kerning information for a combination of two
// glyphs.
//
// Use pango_fc_font_kern_glyphs() instead.
func FontGetKerning(font pango.Font, left pango.Glyph, right pango.Glyph) int {
	var arg0 pango.Font
	arg0 = wrapFont(font)

	var arg1 pango.Glyph
	{
		tmp := uint32(left)
		arg1 = Glyph(tmp)
	}

	var arg2 pango.Glyph
	{
		tmp := uint32(right)
		arg2 = Glyph(tmp)
	}

	c0 := C.pango_ft2_font_get_kerning(arg0, arg1, arg2)

	var ret0 int
	ret0 = int(c0)

	return ret0
}

// GetContext: retrieves a `PangoContext` for the default PangoFT2 fontmap (see
// pango_ft2_font_map_for_display()) and sets the resolution for the default
// fontmap to @dpi_x by @dpi_y.
func GetContext(dpiX float64, dpiY float64) pango.Context {
	var arg0 float64
	arg0 = float64(dpiX)

	var arg1 float64
	arg1 = float64(dpiY)

	c0 := C.pango_ft2_get_context(arg0, arg1)

	var ret0 pango.Context
	ret0 = wrapContext(c0)

	return ret0
}

// GetUnknownGlyph: return the index of a glyph suitable for drawing unknown
// characters with @font, or PANGO_GLYPH_EMPTY if no suitable glyph found.
//
// If you want to draw an unknown-box for a character that is not covered by the
// font, use PANGO_GET_UNKNOWN_GLYPH() instead.
func GetUnknownGlyph(font pango.Font) pango.Glyph {
	var arg0 pango.Font
	arg0 = wrapFont(font)

	c0 := C.pango_ft2_get_unknown_glyph(arg0)

	var ret0 pango.Glyph
	{
		tmp := uint32(c0)
		ret0 = Glyph(tmp)
	}

	return ret0
}

// Render: renders a GlyphString onto a FreeType2 bitmap.
func Render(bitmap *freetype2.Bitmap, font pango.Font, glyphs *pango.GlyphString, x int, y int) {
	var arg0 *freetype2.Bitmap
	arg0 = wrapBitmap(bitmap)

	var arg1 pango.Font
	arg1 = wrapFont(font)

	var arg2 *pango.GlyphString
	arg2 = wrapGlyphString(glyphs)

	var arg3 int
	arg3 = int(x)

	var arg4 int
	arg4 = int(y)

	C.pango_ft2_render(arg0, arg1, arg2, arg3, arg4)
}

// RenderLayout: render a Layout onto a FreeType2 bitmap
func RenderLayout(bitmap *freetype2.Bitmap, layout pango.Layout, x int, y int) {
	var arg0 *freetype2.Bitmap
	arg0 = wrapBitmap(bitmap)

	var arg1 pango.Layout
	arg1 = wrapLayout(layout)

	var arg2 int
	arg2 = int(x)

	var arg3 int
	arg3 = int(y)

	C.pango_ft2_render_layout(arg0, arg1, arg2, arg3)
}

// RenderLayoutLine: render a LayoutLine onto a FreeType2 bitmap
func RenderLayoutLine(bitmap *freetype2.Bitmap, line *pango.LayoutLine, x int, y int) {
	var arg0 *freetype2.Bitmap
	arg0 = wrapBitmap(bitmap)

	var arg1 *pango.LayoutLine
	arg1 = wrapLayoutLine(line)

	var arg2 int
	arg2 = int(x)

	var arg3 int
	arg3 = int(y)

	C.pango_ft2_render_layout_line(arg0, arg1, arg2, arg3)
}

// RenderLayoutLineSubpixel: render a LayoutLine onto a FreeType2 bitmap, with
// he location specified in fixed-point Pango units rather than pixels. (Using
// this will avoid extra inaccuracies from rounding to integer pixels multiple
// times, even if the final glyph positions are integers.)
func RenderLayoutLineSubpixel(bitmap *freetype2.Bitmap, line *pango.LayoutLine, x int, y int) {
	var arg0 *freetype2.Bitmap
	arg0 = wrapBitmap(bitmap)

	var arg1 *pango.LayoutLine
	arg1 = wrapLayoutLine(line)

	var arg2 int
	arg2 = int(x)

	var arg3 int
	arg3 = int(y)

	C.pango_ft2_render_layout_line_subpixel(arg0, arg1, arg2, arg3)
}

// RenderLayoutSubpixel: render a Layout onto a FreeType2 bitmap, with he
// location specified in fixed-point Pango units rather than pixels. (Using this
// will avoid extra inaccuracies from rounding to integer pixels multiple times,
// even if the final glyph positions are integers.)
func RenderLayoutSubpixel(bitmap *freetype2.Bitmap, layout pango.Layout, x int, y int) {
	var arg0 *freetype2.Bitmap
	arg0 = wrapBitmap(bitmap)

	var arg1 pango.Layout
	arg1 = wrapLayout(layout)

	var arg2 int
	arg2 = int(x)

	var arg3 int
	arg3 = int(y)

	C.pango_ft2_render_layout_subpixel(arg0, arg1, arg2, arg3)
}

// RenderTransformed: renders a GlyphString onto a FreeType2 bitmap, possibly
// transforming the layed-out coordinates through a transformation matrix. Note
// that the transformation matrix for @font is not changed, so to produce
// correct rendering results, the @font must have been loaded using a Context
// with an identical transformation matrix to that passed in to this function.
func RenderTransformed(bitmap *freetype2.Bitmap, matrix *pango.Matrix, font pango.Font, glyphs *pango.GlyphString, x int, y int) {
	var arg0 *freetype2.Bitmap
	arg0 = wrapBitmap(bitmap)

	var arg1 *pango.Matrix
	arg1 = wrapMatrix(matrix)

	var arg2 pango.Font
	arg2 = wrapFont(font)

	var arg3 *pango.GlyphString
	arg3 = wrapGlyphString(glyphs)

	var arg4 int
	arg4 = int(x)

	var arg5 int
	arg5 = int(y)

	C.pango_ft2_render_transformed(arg0, arg1, arg2, arg3, arg4, arg5)
}

// ShutdownDisplay: free the global fontmap. (See
// pango_ft2_font_map_for_display()) Use of the global PangoFT2 fontmap is
// deprecated.
func ShutdownDisplay() {
	C.pango_ft2_shutdown_display()
}

// FontMap: the `PangoFT2FontMap` is the `PangoFontMap` implementation for
// FreeType fonts.
type FontMap interface {
	pangofc.FontMap

	// CreateContext: create a `PangoContext` for the given fontmap.
	CreateContext() pango.Context
	// SetDefaultSubstitute: sets a function that will be called to do final
	// configuration substitution on a `FcPattern` before it is used to load the
	// font.
	//
	// This function can be used to do things like set hinting and antialiasing
	// options.
	SetDefaultSubstitute(_func SubstituteFunc, data unsafe.Pointer, notify unsafe.Pointer)
	// SetResolution: sets the horizontal and vertical resolutions for the
	// fontmap.
	SetResolution(dpiX float64, dpiY float64)
	// SubstituteChanged: call this function any time the results of the default
	// substitution function set with
	// pango_ft2_font_map_set_default_substitute() change.
	//
	// That is, if your substitution function will return different results for
	// the same input pattern, you must call this function.
	SubstituteChanged()
}

type fontMap struct {
	pangofc.fontMap
}

func wrapFontMap(obj *externglib.Object) FontMap {
	return fontMap{pangofc.fontMap{pango.fontMap{*externglib.Object{obj}}}}
}

func marshalFontMap(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapWidget(obj), nil
}

func NewFontMap() FontMap

func (f fontMap) CreateContext() pango.Context

func (f fontMap) SetDefaultSubstitute(_func SubstituteFunc, data unsafe.Pointer, notify unsafe.Pointer)

func (f fontMap) SetResolution(dpiX float64, dpiY float64)

func (f fontMap) SubstituteChanged()
