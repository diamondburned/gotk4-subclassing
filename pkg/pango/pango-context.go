// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_context_get_type()), F: marshalContext},
	})
}

// Itemize breaks a piece of text into segments with consistent directional
// level and font.
//
// Each byte of @text will be contained in exactly one of the items in the
// returned list; the generated list of items will be in logical order (the
// start offsets of the items are ascending).
//
// @cached_iter should be an iterator over @attrs currently positioned at a
// range before or containing @start_index; @cached_iter will be advanced to the
// range covering the position just after @start_index + @length. (i.e. if
// itemizing in a loop, just keep passing in the same @cached_iter).
func Itemize(context Context, text string, startIndex int, length int, attrs *AttrList, cachedIter *AttrIterator) *glib.List {
	var arg1 *C.PangoContext
	var arg2 *C.char
	var arg3 C.int
	var arg4 C.int
	var arg5 *C.PangoAttrList
	var arg6 *C.PangoAttrIterator

	arg1 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	arg2 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = C.int(startIndex)
	arg4 = C.int(length)
	arg5 = (*C.PangoAttrList)(unsafe.Pointer(attrs.Native()))
	arg6 = (*C.PangoAttrIterator)(unsafe.Pointer(cachedIter.Native()))

	var cret *C.GList

	cret = C.pango_itemize(arg1, arg2, arg3, arg4, arg5, arg6)

	var list *glib.List

	list = glib.WrapList(unsafe.Pointer(cret))
	runtime.SetFinalizer(list, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return list
}

// ItemizeWithBaseDir: like `pango_itemize()`, but with an explicitly specified
// base direction.
//
// The base direction is used when computing bidirectional levels. (see
// [method@Pango.Context.set_base_dir]). [func@itemize] gets the base direction
// from the `PangoContext`.
func ItemizeWithBaseDir(context Context, baseDir Direction, text string, startIndex int, length int, attrs *AttrList, cachedIter *AttrIterator) *glib.List {
	var arg1 *C.PangoContext
	var arg2 C.PangoDirection
	var arg3 *C.char
	var arg4 C.int
	var arg5 C.int
	var arg6 *C.PangoAttrList
	var arg7 *C.PangoAttrIterator

	arg1 = (*C.PangoContext)(unsafe.Pointer(context.Native()))
	arg2 = (C.PangoDirection)(baseDir)
	arg3 = (*C.char)(C.CString(text))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.int(startIndex)
	arg5 = C.int(length)
	arg6 = (*C.PangoAttrList)(unsafe.Pointer(attrs.Native()))
	arg7 = (*C.PangoAttrIterator)(unsafe.Pointer(cachedIter.Native()))

	var cret *C.GList

	cret = C.pango_itemize_with_base_dir(arg1, arg2, arg3, arg4, arg5, arg6, arg7)

	var list *glib.List

	list = glib.WrapList(unsafe.Pointer(cret))
	runtime.SetFinalizer(list, func(v *glib.List) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return list
}

// Context: a `PangoContext` stores global information used to control the
// itemization process.
//
// The information stored by `PangoContext includes the fontmap used to look up
// fonts, and default values such as the default language, default gravity, or
// default font.
//
// To obtain a `PangoContext`, use [method@Pango.FontMap.create_context].
type Context interface {
	gextras.Objector

	// Changed forces a change in the context, which will cause any
	// `PangoLayout` using this context to re-layout.
	//
	// This function is only useful when implementing a new backend for Pango,
	// something applications won't do. Backends should call this function if
	// they have attached extra data to the context and such data is changed.
	Changed()
	// BaseDir retrieves the base direction for the context.
	//
	// See [method@Pango.Context.set_base_dir].
	BaseDir() Direction
	// BaseGravity retrieves the base gravity for the context.
	//
	// See [method@Pango.Context.set_base_gravity].
	BaseGravity() Gravity
	// FontDescription: retrieve the default font description for the context.
	FontDescription() *FontDescription
	// FontMap gets the `PangoFontMap` used to look up fonts for this context.
	FontMap() FontMap
	// Gravity retrieves the gravity for the context.
	//
	// This is similar to [method@Pango.Context.get_base_gravity], except for
	// when the base gravity is PANGO_GRAVITY_AUTO for which
	// [type_func@Pango.Gravity.get_for_matrix] is used to return the gravity
	// from the current context matrix.
	Gravity() Gravity
	// GravityHint retrieves the gravity hint for the context.
	//
	// See [method@Pango.Context.set_gravity_hint] for details.
	GravityHint() GravityHint
	// Language retrieves the global language tag for the context.
	Language() *Language
	// Matrix gets the transformation matrix that will be applied when rendering
	// with this context.
	//
	// See [method@Pango.Context.set_matrix].
	Matrix() *Matrix
	// Metrics: get overall metric information for a particular font
	// description.
	//
	// Since the metrics may be substantially different for different scripts, a
	// language tag can be provided to indicate that the metrics should be
	// retrieved that correspond to the script(s) used by that language.
	//
	// The `PangoFontDescription` is interpreted in the same way as by
	// [func@itemize], and the family name may be a comma separated list of
	// names. If characters from multiple of these families would be used to
	// render the string, then the returned fonts would be a composite of the
	// metrics for the fonts loaded for the individual families.
	Metrics(desc *FontDescription, language *Language) *FontMetrics
	// RoundGlyphPositions returns whether font rendering with this context
	// should round glyph positions and widths.
	RoundGlyphPositions() bool
	// Serial returns the current serial number of @context.
	//
	// The serial number is initialized to an small number larger than zero when
	// a new context is created and is increased whenever the context is changed
	// using any of the setter functions, or the `PangoFontMap` it uses to find
	// fonts has changed. The serial may wrap, but will never have the value 0.
	// Since it can wrap, never compare it with "less than", always use "not
	// equals".
	//
	// This can be used to automatically detect changes to a `PangoContext`, and
	// is only useful when implementing objects that need update when their
	// `PangoContext` changes, like `PangoLayout`.
	Serial() uint
	// ListFamilies: list all families for a context.
	ListFamilies()
	// LoadFont loads the font in one of the fontmaps in the context that is the
	// closest match for @desc.
	LoadFont(desc *FontDescription) Font
	// LoadFontset: load a set of fonts in the context that can be used to
	// render a font matching @desc.
	LoadFontset(desc *FontDescription, language *Language) Fontset
	// SetBaseDir sets the base direction for the context.
	//
	// The base direction is used in applying the Unicode bidirectional
	// algorithm; if the @direction is PANGO_DIRECTION_LTR or
	// PANGO_DIRECTION_RTL, then the value will be used as the paragraph
	// direction in the Unicode bidirectional algorithm. A value of
	// PANGO_DIRECTION_WEAK_LTR or PANGO_DIRECTION_WEAK_RTL is used only for
	// paragraphs that do not contain any strong characters themselves.
	SetBaseDir(direction Direction)
	// SetBaseGravity sets the base gravity for the context.
	//
	// The base gravity is used in laying vertical text out.
	SetBaseGravity(gravity Gravity)
	// SetFontDescription: set the default font description for the context
	SetFontDescription(desc *FontDescription)
	// SetFontMap sets the font map to be searched when fonts are looked-up in
	// this context.
	//
	// This is only for internal use by Pango backends, a `PangoContext`
	// obtained via one of the recommended methods should already have a
	// suitable font map.
	SetFontMap(fontMap FontMap)
	// SetGravityHint sets the gravity hint for the context.
	//
	// The gravity hint is used in laying vertical text out, and is only
	// relevant if gravity of the context as returned by
	// [method@Pango.Context.get_gravity] is set to PANGO_GRAVITY_EAST or
	// PANGO_GRAVITY_WEST.
	SetGravityHint(hint GravityHint)
	// SetLanguage sets the global language tag for the context.
	//
	// The default language for the locale of the running process can be found
	// using [type_func@Pango.Language.get_default].
	SetLanguage(language *Language)
	// SetMatrix sets the transformation matrix that will be applied when
	// rendering with this context.
	//
	// Note that reported metrics are in the user space coordinates before the
	// application of the matrix, not device-space coordinates after the
	// application of the matrix. So, they don't scale with the matrix, though
	// they may change slightly for different matrices, depending on how the
	// text is fit to the pixel grid.
	SetMatrix(matrix *Matrix)
	// SetRoundGlyphPositions sets whether font rendering with this context
	// should round glyph positions and widths to integral positions, in device
	// units.
	//
	// This is useful when the renderer can't handle subpixel positioning of
	// glyphs.
	//
	// The default value is to round glyph positions, to remain compatible with
	// previous Pango behavior.
	SetRoundGlyphPositions(roundPositions bool)
}

// context implements the Context interface.
type context struct {
	gextras.Objector
}

var _ Context = (*context)(nil)

// WrapContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapContext(obj *externglib.Object) Context {
	return Context{
		Objector: obj,
	}
}

func marshalContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapContext(obj), nil
}

// NewContext constructs a class Context.
func NewContext() Context {
	var cret C.PangoContext

	cret = C.pango_context_new()

	var context Context

	context = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Context)

	return context
}

// Changed forces a change in the context, which will cause any
// `PangoLayout` using this context to re-layout.
//
// This function is only useful when implementing a new backend for Pango,
// something applications won't do. Backends should call this function if
// they have attached extra data to the context and such data is changed.
func (c context) Changed() {
	var arg0 *C.PangoContext

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	C.pango_context_changed(arg0)
}

// BaseDir retrieves the base direction for the context.
//
// See [method@Pango.Context.set_base_dir].
func (c context) BaseDir() Direction {
	var arg0 *C.PangoContext

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	var cret C.PangoDirection

	cret = C.pango_context_get_base_dir(arg0)

	var direction Direction

	direction = Direction(cret)

	return direction
}

// BaseGravity retrieves the base gravity for the context.
//
// See [method@Pango.Context.set_base_gravity].
func (c context) BaseGravity() Gravity {
	var arg0 *C.PangoContext

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	var cret C.PangoGravity

	cret = C.pango_context_get_base_gravity(arg0)

	var gravity Gravity

	gravity = Gravity(cret)

	return gravity
}

// FontDescription: retrieve the default font description for the context.
func (c context) FontDescription() *FontDescription {
	var arg0 *C.PangoContext

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	var cret *C.PangoFontDescription

	cret = C.pango_context_get_font_description(arg0)

	var fontDescription *FontDescription

	fontDescription = WrapFontDescription(unsafe.Pointer(cret))

	return fontDescription
}

// FontMap gets the `PangoFontMap` used to look up fonts for this context.
func (c context) FontMap() FontMap {
	var arg0 *C.PangoContext

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	var cret *C.PangoFontMap

	cret = C.pango_context_get_font_map(arg0)

	var fontMap FontMap

	fontMap = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(FontMap)

	return fontMap
}

// Gravity retrieves the gravity for the context.
//
// This is similar to [method@Pango.Context.get_base_gravity], except for
// when the base gravity is PANGO_GRAVITY_AUTO for which
// [type_func@Pango.Gravity.get_for_matrix] is used to return the gravity
// from the current context matrix.
func (c context) Gravity() Gravity {
	var arg0 *C.PangoContext

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	var cret C.PangoGravity

	cret = C.pango_context_get_gravity(arg0)

	var gravity Gravity

	gravity = Gravity(cret)

	return gravity
}

// GravityHint retrieves the gravity hint for the context.
//
// See [method@Pango.Context.set_gravity_hint] for details.
func (c context) GravityHint() GravityHint {
	var arg0 *C.PangoContext

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	var cret C.PangoGravityHint

	cret = C.pango_context_get_gravity_hint(arg0)

	var gravityHint GravityHint

	gravityHint = GravityHint(cret)

	return gravityHint
}

// Language retrieves the global language tag for the context.
func (c context) Language() *Language {
	var arg0 *C.PangoContext

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	var cret *C.PangoLanguage

	cret = C.pango_context_get_language(arg0)

	var language *Language

	language = WrapLanguage(unsafe.Pointer(cret))
	runtime.SetFinalizer(language, func(v *Language) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return language
}

// Matrix gets the transformation matrix that will be applied when rendering
// with this context.
//
// See [method@Pango.Context.set_matrix].
func (c context) Matrix() *Matrix {
	var arg0 *C.PangoContext

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	var cret *C.PangoMatrix

	cret = C.pango_context_get_matrix(arg0)

	var matrix *Matrix

	matrix = WrapMatrix(unsafe.Pointer(cret))

	return matrix
}

// Metrics: get overall metric information for a particular font
// description.
//
// Since the metrics may be substantially different for different scripts, a
// language tag can be provided to indicate that the metrics should be
// retrieved that correspond to the script(s) used by that language.
//
// The `PangoFontDescription` is interpreted in the same way as by
// [func@itemize], and the family name may be a comma separated list of
// names. If characters from multiple of these families would be used to
// render the string, then the returned fonts would be a composite of the
// metrics for the fonts loaded for the individual families.
func (c context) Metrics(desc *FontDescription, language *Language) *FontMetrics {
	var arg0 *C.PangoContext
	var arg1 *C.PangoFontDescription
	var arg2 *C.PangoLanguage

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.PangoFontDescription)(unsafe.Pointer(desc.Native()))
	arg2 = (*C.PangoLanguage)(unsafe.Pointer(language.Native()))

	var cret *C.PangoFontMetrics

	cret = C.pango_context_get_metrics(arg0, arg1, arg2)

	var fontMetrics *FontMetrics

	fontMetrics = WrapFontMetrics(unsafe.Pointer(cret))
	runtime.SetFinalizer(fontMetrics, func(v *FontMetrics) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return fontMetrics
}

// RoundGlyphPositions returns whether font rendering with this context
// should round glyph positions and widths.
func (c context) RoundGlyphPositions() bool {
	var arg0 *C.PangoContext

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	var cret C.gboolean

	cret = C.pango_context_get_round_glyph_positions(arg0)

	var ok bool

	if cret {
		ok = true
	}

	return ok
}

// Serial returns the current serial number of @context.
//
// The serial number is initialized to an small number larger than zero when
// a new context is created and is increased whenever the context is changed
// using any of the setter functions, or the `PangoFontMap` it uses to find
// fonts has changed. The serial may wrap, but will never have the value 0.
// Since it can wrap, never compare it with "less than", always use "not
// equals".
//
// This can be used to automatically detect changes to a `PangoContext`, and
// is only useful when implementing objects that need update when their
// `PangoContext` changes, like `PangoLayout`.
func (c context) Serial() uint {
	var arg0 *C.PangoContext

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	var cret C.guint

	cret = C.pango_context_get_serial(arg0)

	var guint uint

	guint = (uint)(cret)

	return guint
}

// ListFamilies: list all families for a context.
func (c context) ListFamilies() {
	var arg0 *C.PangoContext

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))

	C.pango_context_list_families(arg0)

	return
}

// LoadFont loads the font in one of the fontmaps in the context that is the
// closest match for @desc.
func (c context) LoadFont(desc *FontDescription) Font {
	var arg0 *C.PangoContext
	var arg1 *C.PangoFontDescription

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.PangoFontDescription)(unsafe.Pointer(desc.Native()))

	var cret *C.PangoFont

	cret = C.pango_context_load_font(arg0, arg1)

	var font Font

	font = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Font)

	return font
}

// LoadFontset: load a set of fonts in the context that can be used to
// render a font matching @desc.
func (c context) LoadFontset(desc *FontDescription, language *Language) Fontset {
	var arg0 *C.PangoContext
	var arg1 *C.PangoFontDescription
	var arg2 *C.PangoLanguage

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.PangoFontDescription)(unsafe.Pointer(desc.Native()))
	arg2 = (*C.PangoLanguage)(unsafe.Pointer(language.Native()))

	var cret *C.PangoFontset

	cret = C.pango_context_load_fontset(arg0, arg1, arg2)

	var fontset Fontset

	fontset = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Fontset)

	return fontset
}

// SetBaseDir sets the base direction for the context.
//
// The base direction is used in applying the Unicode bidirectional
// algorithm; if the @direction is PANGO_DIRECTION_LTR or
// PANGO_DIRECTION_RTL, then the value will be used as the paragraph
// direction in the Unicode bidirectional algorithm. A value of
// PANGO_DIRECTION_WEAK_LTR or PANGO_DIRECTION_WEAK_RTL is used only for
// paragraphs that do not contain any strong characters themselves.
func (c context) SetBaseDir(direction Direction) {
	var arg0 *C.PangoContext
	var arg1 C.PangoDirection

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	arg1 = (C.PangoDirection)(direction)

	C.pango_context_set_base_dir(arg0, arg1)
}

// SetBaseGravity sets the base gravity for the context.
//
// The base gravity is used in laying vertical text out.
func (c context) SetBaseGravity(gravity Gravity) {
	var arg0 *C.PangoContext
	var arg1 C.PangoGravity

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	arg1 = (C.PangoGravity)(gravity)

	C.pango_context_set_base_gravity(arg0, arg1)
}

// SetFontDescription: set the default font description for the context
func (c context) SetFontDescription(desc *FontDescription) {
	var arg0 *C.PangoContext
	var arg1 *C.PangoFontDescription

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.PangoFontDescription)(unsafe.Pointer(desc.Native()))

	C.pango_context_set_font_description(arg0, arg1)
}

// SetFontMap sets the font map to be searched when fonts are looked-up in
// this context.
//
// This is only for internal use by Pango backends, a `PangoContext`
// obtained via one of the recommended methods should already have a
// suitable font map.
func (c context) SetFontMap(fontMap FontMap) {
	var arg0 *C.PangoContext
	var arg1 *C.PangoFontMap

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.PangoFontMap)(unsafe.Pointer(fontMap.Native()))

	C.pango_context_set_font_map(arg0, arg1)
}

// SetGravityHint sets the gravity hint for the context.
//
// The gravity hint is used in laying vertical text out, and is only
// relevant if gravity of the context as returned by
// [method@Pango.Context.get_gravity] is set to PANGO_GRAVITY_EAST or
// PANGO_GRAVITY_WEST.
func (c context) SetGravityHint(hint GravityHint) {
	var arg0 *C.PangoContext
	var arg1 C.PangoGravityHint

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	arg1 = (C.PangoGravityHint)(hint)

	C.pango_context_set_gravity_hint(arg0, arg1)
}

// SetLanguage sets the global language tag for the context.
//
// The default language for the locale of the running process can be found
// using [type_func@Pango.Language.get_default].
func (c context) SetLanguage(language *Language) {
	var arg0 *C.PangoContext
	var arg1 *C.PangoLanguage

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.PangoLanguage)(unsafe.Pointer(language.Native()))

	C.pango_context_set_language(arg0, arg1)
}

// SetMatrix sets the transformation matrix that will be applied when
// rendering with this context.
//
// Note that reported metrics are in the user space coordinates before the
// application of the matrix, not device-space coordinates after the
// application of the matrix. So, they don't scale with the matrix, though
// they may change slightly for different matrices, depending on how the
// text is fit to the pixel grid.
func (c context) SetMatrix(matrix *Matrix) {
	var arg0 *C.PangoContext
	var arg1 *C.PangoMatrix

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.PangoMatrix)(unsafe.Pointer(matrix.Native()))

	C.pango_context_set_matrix(arg0, arg1)
}

// SetRoundGlyphPositions sets whether font rendering with this context
// should round glyph positions and widths to integral positions, in device
// units.
//
// This is useful when the renderer can't handle subpixel positioning of
// glyphs.
//
// The default value is to round glyph positions, to remain compatible with
// previous Pango behavior.
func (c context) SetRoundGlyphPositions(roundPositions bool) {
	var arg0 *C.PangoContext
	var arg1 C.gboolean

	arg0 = (*C.PangoContext)(unsafe.Pointer(c.Native()))
	if roundPositions {
		arg1 = C.gboolean(1)
	}

	C.pango_context_set_round_glyph_positions(arg0, arg1)
}
