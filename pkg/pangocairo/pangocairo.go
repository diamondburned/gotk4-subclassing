// Code generated by girgen. DO NOT EDIT.

package pangocairo

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/pango"
	"github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pangocairo
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <pango/pangocairo.h>
import "C"

func init() {
	glib.RegisterGValueMarshalers([]glib.TypeMarshaler{

		// Objects/Classes
	})
}

// ContextGetFontOptions: retrieves any font rendering options previously set
// with [func@PangoCairo.context_set_font_options].
//
// This function does not report options that are derived from the target
// surface by [func@update_context].
func ContextGetFontOptions(context *pango.Context) *cairo.FontOptions

// ContextGetResolution: gets the resolution for the context. See
// [func@PangoCairo.context_set_resolution]
func ContextGetResolution(context *pango.Context) float64

// ContextGetShapeRenderer: sets callback function for context to use for
// rendering attributes of type PANGO_ATTR_SHAPE.
//
// See `PangoCairoShapeRendererFunc` for details.
//
// Retrieves callback function and associated user data for rendering attributes
// of type PANGO_ATTR_SHAPE as set by
// [func@PangoCairo.context_set_shape_renderer], if any.
func ContextGetShapeRenderer(context *pango.Context, data unsafe.Pointer) ShapeRendererFunc

// ContextSetFontOptions: sets the font options used when rendering text with
// this context.
//
// These options override any options that [func@update_context] derives from
// the target surface.
func ContextSetFontOptions(context *pango.Context, options *cairo.FontOptions)

// ContextSetResolution: sets the resolution for the context.
//
// This is a scale factor between points specified in a `PangoFontDescription`
// and Cairo units. The default value is 96, meaning that a 10 point font will
// be 13 units high. (10 * 96. / 72. = 13.3).
func ContextSetResolution(context *pango.Context, dpi float64)

// ContextSetShapeRenderer: sets callback function for context to use for
// rendering attributes of type PANGO_ATTR_SHAPE.
//
// See `PangoCairoShapeRendererFunc` for details.
func ContextSetShapeRenderer(context *pango.Context, _func ShapeRendererFunc, data unsafe.Pointer, dnotify unsafe.Pointer)

// CreateContext: creates a context object set up to match the current
// transformation and target surface of the Cairo context.
//
// This context can then be used to create a layout using
// [ctor@Pango.Layout.new].
//
// This function is a convenience function that creates a context using the
// default font map, then updates it to @cr. If you just need to create a layout
// for use with @cr and do not need to access `PangoContext` directly, you can
// use [func@create_layout] instead.
func CreateContext(cr *cairo.Context) *pango.Context

// CreateLayout: creates a layout object set up to match the current
// transformation and target surface of the Cairo context.
//
// This layout can then be used for text measurement with functions like
// [method@Pango.Layout.get_size] or drawing with functions like
// [func@show_layout]. If you change the transformation or target surface for
// @cr, you need to call [func@update_layout].
//
// This function is the most convenient way to use Cairo with Pango, however it
// is slightly inefficient since it creates a separate `PangoContext` object for
// each layout. This might matter in an application that was laying out large
// amounts of text.
func CreateLayout(cr *cairo.Context) *pango.Layout

// ErrorUnderlinePath: add a squiggly line to the current path in the specified
// cairo context that approximately covers the given rectangle in the style of
// an underline used to indicate a spelling error.
//
// The width of the underline is rounded to an integer number of up/down
// segments and the resulting rectangle is centered in the original rectangle.
func ErrorUnderlinePath(cr *cairo.Context, x float64, y float64, width float64, height float64)

// FontMapGetDefault: gets a default `PangoCairoFontMap` to use with Cairo.
//
// Note that the type of the returned object will depend on the particular font
// backend Cairo was compiled to use; you generally should only use the
// `PangoFontMap` and `PangoCairoFontMap` interfaces on the returned object.
//
// The default Cairo fontmap can be changed by using
// [method@PangoCairo.FontMap.set_default]. This can be used to change the Cairo
// font backend that the default fontmap uses for example.
//
// Note that since Pango 1.32.6, the default fontmap is per-thread. Each thread
// gets its own default fontmap. In this way, PangoCairo can be used safely from
// multiple threads.
func FontMapGetDefault() *pango.FontMap

// NewFontMap: creates a new `PangoCairoFontMap` object.
//
// A fontmap is used to cache information about available fonts, and holds
// certain global parameters such as the resolution. In most cases, you can use
// `func@PangoCairo.font_map_get_default] instead.
//
// Note that the type of the returned object will depend on the particular font
// backend Cairo was compiled to use; You generally should only use the
// `PangoFontMap` and `PangoCairoFontMap` interfaces on the returned object.
//
// You can override the type of backend returned by using an environment
// variable PANGOCAIRO_BACKEND. Supported types, based on your build, are fc
// (fontconfig), win32, and coretext. If requested type is not available, NULL
// is returned. Ie. this is only useful for testing, when at least two backends
// are compiled in.
func NewFontMap() *pango.FontMap

// FontMapNewForFontType: creates a new `PangoCairoFontMap` object of the type
// suitable to be used with cairo font backend of type @fonttype.
//
// In most cases one should simply use [type_func@PangoCairo.FontMap.new], or in
// fact in most of those cases, just use [func@PangoCairo.FontMap.get_default].
func FontMapNewForFontType(fonttype cairo.FontType) *pango.FontMap

// GlyphStringPath: adds the glyphs in @glyphs to the current path in the
// specified cairo context.
//
// The origin of the glyphs (the left edge of the baseline) will be at the
// current point of the cairo context.
func GlyphStringPath(cr *cairo.Context, font *pango.Font, glyphs *pango.GlyphString)

// LayoutLinePath: adds the text in `PangoLayoutLine` to the current path in the
// specified cairo context.
//
// The origin of the glyphs (the left edge of the line) will be at the current
// point of the cairo context.
func LayoutLinePath(cr *cairo.Context, line *pango.LayoutLine)

// LayoutPath: adds the text in a `PangoLayout` to the current path in the
// specified cairo context.
//
// The top-left corner of the `PangoLayout` will be at the current point of the
// cairo context.
func LayoutPath(cr *cairo.Context, layout *pango.Layout)

// ShowErrorUnderline: draw a squiggly line in the specified cairo context that
// approximately covers the given rectangle in the style of an underline used to
// indicate a spelling error.
//
// The width of the underline is rounded to an integer number of up/down
// segments and the resulting rectangle is centered in the original rectangle.
func ShowErrorUnderline(cr *cairo.Context, x float64, y float64, width float64, height float64)

// ShowGlyphItem: draws the glyphs in @glyph_item in the specified cairo
// context,
//
// embedding the text associated with the glyphs in the output if the output
// format supports it (PDF for example), otherwise it acts similar to
// [func@show_glyph_string].
//
// The origin of the glyphs (the left edge of the baseline) will be drawn at the
// current point of the cairo context.
//
// Note that @text is the start of the text for layout, which is then indexed by
// `glyph_item->item->offset`.
func ShowGlyphItem(cr *cairo.Context, text string, glyphItem *pango.GlyphItem)

// ShowGlyphString: draws the glyphs in @glyphs in the specified cairo context.
//
// The origin of the glyphs (the left edge of the baseline) will be drawn at the
// current point of the cairo context.
func ShowGlyphString(cr *cairo.Context, font *pango.Font, glyphs *pango.GlyphString)

// ShowLayout: draws a `PangoLayout` in the specified cairo context.
//
// The top-left corner of the `PangoLayout` will be drawn at the current point
// of the cairo context.
func ShowLayout(cr *cairo.Context, layout *pango.Layout)

// ShowLayoutLine: draws a `PangoLayoutLine` in the specified cairo context.
//
// The origin of the glyphs (the left edge of the line) will be drawn at the
// current point of the cairo context.
func ShowLayoutLine(cr *cairo.Context, line *pango.LayoutLine)

// UpdateContext: updates a `PangoContext` previously created for use with Cairo
// to match the current transformation and target surface of a Cairo context.
//
// If any layouts have been created for the context, it's necessary to call
// [method@Pango.Layout.context_changed] on those layouts.
func UpdateContext(cr *cairo.Context, context *pango.Context)

// UpdateLayout: updates the private `PangoContext` of a `PangoLayout` created
// with [func@create_layout] to match the current transformation and target
// surface of a Cairo context.
func UpdateLayout(cr *cairo.Context, layout *pango.Layout)
