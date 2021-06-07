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
		{T: externglib.Type(C.pango_font_get_type()), F: marshalFont},
		{T: externglib.Type(C.pango_font_face_get_type()), F: marshalFontFace},
		{T: externglib.Type(C.pango_font_family_get_type()), F: marshalFontFamily},
		{T: externglib.Type(C.pango_font_description_get_type()), F: marshalFontDescription},
		{T: externglib.Type(C.pango_font_metrics_get_type()), F: marshalFontMetrics},
	})
}

// FontDescriptionFromString creates a new font description from a string
// representation.
//
// The string must have the form
//
//    "\[FAMILY-LIST] \[STYLE-OPTIONS] \[SIZE] \[VARIATIONS]",
//
// where FAMILY-LIST is a comma-separated list of families optionally terminated
// by a comma, STYLE_OPTIONS is a whitespace-separated list of words where each
// word describes one of style, variant, weight, stretch, or gravity, and SIZE
// is a decimal number (size in points) or optionally followed by the unit
// modifier "px" for absolute size. VARIATIONS is a comma-separated list of font
// variation specifications of the form "\@axis=value" (the = sign is optional).
//
// The following words are understood as styles: "Normal", "Roman", "Oblique",
// "Italic".
//
// The following words are understood as variants: "Small-Caps".
//
// The following words are understood as weights: "Thin", "Ultra-Light",
// "Extra-Light", "Light", "Semi-Light", "Demi-Light", "Book", "Regular",
// "Medium", "Semi-Bold", "Demi-Bold", "Bold", "Ultra-Bold", "Extra-Bold",
// "Heavy", "Black", "Ultra-Black", "Extra-Black".
//
// The following words are understood as stretch values: "Ultra-Condensed",
// "Extra-Condensed", "Condensed", "Semi-Condensed", "Semi-Expanded",
// "Expanded", "Extra-Expanded", "Ultra-Expanded".
//
// The following words are understood as gravity values: "Not-Rotated", "South",
// "Upside-Down", "North", "Rotated-Left", "East", "Rotated-Right", "West".
//
// Any one of the options may be absent. If FAMILY-LIST is absent, then the
// family_name field of the resulting font description will be initialized to
// nil. If STYLE-OPTIONS is missing, then all style options will be set to the
// default values. If SIZE is missing, the size in the resulting font
// description will be set to 0.
//
// A typical example:
//
//    "Cantarell Italic Light 15 \@wght=200"
func FontDescriptionFromString(str string) {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(str))
	defer C.free(unsafe.Pointer(arg1))

	C.pango_font_description_from_string(arg1)
}

// Font: a `PangoFont` is used to represent a font in a
// rendering-system-independent manner.
type Font interface {
	gextras.Objector

	// Describe returns a description of the font, with font size set in points.
	//
	// Use [method@Pango.Font.describe_with_absolute_size] if you want the font
	// size in device units.
	Describe(f Font)
	// DescribeWithAbsoluteSize returns a description of the font, with absolute
	// font size set in device units.
	//
	// Use [method@Pango.Font.describe] if you want the font size in points.
	DescribeWithAbsoluteSize(f Font)
	// Coverage computes the coverage map for a given font and language tag.
	Coverage(f Font, language *Language)
	// Face gets the `PangoFontFace` to which @font belongs.
	Face(f Font)
	// FontMap gets the font map for which the font was created.
	//
	// Note that the font maintains a *weak* reference to the font map, so if
	// all references to font map are dropped, the font map will be finalized
	// even if there are fonts created with the font map that are still alive.
	// In that case this function will return nil.
	//
	// It is the responsibility of the user to ensure that the font map is kept
	// alive. In most uses this is not an issue as a Context holds a reference
	// to the font map.
	FontMap(f Font)
	// HbFont: get a `hb_font_t` object backing this font.
	//
	// Note that the objects returned by this function are cached and immutable.
	// If you need to make changes to the `hb_font_t`, use
	// hb_font_create_sub_font().
	HbFont(f Font)
	// Metrics gets overall metric information for a font.
	//
	// Since the metrics may be substantially different for different scripts, a
	// language tag can be provided to indicate that the metrics should be
	// retrieved that correspond to the script(s) used by that language.
	//
	// If @font is nil, this function gracefully sets some sane values in the
	// output variables and returns.
	Metrics(f Font, language *Language)
	// HasChar returns whether the font provides a glyph for this character.
	//
	// Returns true if @font can render @wc
	HasChar(f Font, wc uint32) bool
}

// font implements the Font interface.
type font struct {
	gextras.Objector
}

var _ Font = (*font)(nil)

// WrapFont wraps a GObject to the right type. It is
// primarily used internally.
func WrapFont(obj *externglib.Object) Font {
	return Font{
		Objector: obj,
	}
}

func marshalFont(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFont(obj), nil
}

// Describe returns a description of the font, with font size set in points.
//
// Use [method@Pango.Font.describe_with_absolute_size] if you want the font
// size in device units.
func (f font) Describe(f Font) {
	var arg0 *C.PangoFont

	arg0 = (*C.PangoFont)(unsafe.Pointer(f.Native()))

	C.pango_font_describe(arg0)
}

// DescribeWithAbsoluteSize returns a description of the font, with absolute
// font size set in device units.
//
// Use [method@Pango.Font.describe] if you want the font size in points.
func (f font) DescribeWithAbsoluteSize(f Font) {
	var arg0 *C.PangoFont

	arg0 = (*C.PangoFont)(unsafe.Pointer(f.Native()))

	C.pango_font_describe_with_absolute_size(arg0)
}

// Coverage computes the coverage map for a given font and language tag.
func (f font) Coverage(f Font, language *Language) {
	var arg0 *C.PangoFont
	var arg1 *C.PangoLanguage

	arg0 = (*C.PangoFont)(unsafe.Pointer(f.Native()))
	arg1 = (*C.PangoLanguage)(unsafe.Pointer(language.Native()))

	C.pango_font_get_coverage(arg0, arg1)
}

// Face gets the `PangoFontFace` to which @font belongs.
func (f font) Face(f Font) {
	var arg0 *C.PangoFont

	arg0 = (*C.PangoFont)(unsafe.Pointer(f.Native()))

	C.pango_font_get_face(arg0)
}

// FontMap gets the font map for which the font was created.
//
// Note that the font maintains a *weak* reference to the font map, so if
// all references to font map are dropped, the font map will be finalized
// even if there are fonts created with the font map that are still alive.
// In that case this function will return nil.
//
// It is the responsibility of the user to ensure that the font map is kept
// alive. In most uses this is not an issue as a Context holds a reference
// to the font map.
func (f font) FontMap(f Font) {
	var arg0 *C.PangoFont

	arg0 = (*C.PangoFont)(unsafe.Pointer(f.Native()))

	C.pango_font_get_font_map(arg0)
}

// HbFont: get a `hb_font_t` object backing this font.
//
// Note that the objects returned by this function are cached and immutable.
// If you need to make changes to the `hb_font_t`, use
// hb_font_create_sub_font().
func (f font) HbFont(f Font) {
	var arg0 *C.PangoFont

	arg0 = (*C.PangoFont)(unsafe.Pointer(f.Native()))

	C.pango_font_get_hb_font(arg0)
}

// Metrics gets overall metric information for a font.
//
// Since the metrics may be substantially different for different scripts, a
// language tag can be provided to indicate that the metrics should be
// retrieved that correspond to the script(s) used by that language.
//
// If @font is nil, this function gracefully sets some sane values in the
// output variables and returns.
func (f font) Metrics(f Font, language *Language) {
	var arg0 *C.PangoFont
	var arg1 *C.PangoLanguage

	arg0 = (*C.PangoFont)(unsafe.Pointer(f.Native()))
	arg1 = (*C.PangoLanguage)(unsafe.Pointer(language.Native()))

	C.pango_font_get_metrics(arg0, arg1)
}

// HasChar returns whether the font provides a glyph for this character.
//
// Returns true if @font can render @wc
func (f font) HasChar(f Font, wc uint32) bool {
	var arg0 *C.PangoFont
	var arg1 C.gunichar

	arg0 = (*C.PangoFont)(unsafe.Pointer(f.Native()))
	arg1 = C.gunichar(wc)

	var cret C.gboolean
	var ok bool

	cret = C.pango_font_has_char(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// FontFace: a `PangoFontFace` is used to represent a group of fonts with the
// same family, slant, weight, and width, but varying sizes.
type FontFace interface {
	gextras.Objector

	// Describe returns the family, style, variant, weight and stretch of a
	// `PangoFontFace`. The size field of the resulting font description will be
	// unset.
	Describe(f FontFace)
	// FaceName gets a name representing the style of this face among the
	// different faces in the `PangoFontFamily` for the face. The name is
	// suitable for displaying to users.
	FaceName(f FontFace)
	// Family gets the `PangoFontFamily` that @face belongs to.
	Family(f FontFace)
	// IsSynthesized returns whether a `PangoFontFace` is synthesized by the
	// underlying font rendering engine from another face, perhaps by shearing,
	// emboldening, or lightening it.
	IsSynthesized(f FontFace) bool
	// ListSizes: list the available sizes for a font.
	//
	// This is only applicable to bitmap fonts. For scalable fonts, stores nil
	// at the location pointed to by @sizes and 0 at the location pointed to by
	// @n_sizes. The sizes returned are in Pango units and are sorted in
	// ascending order.
	ListSizes(f FontFace)
}

// fontFace implements the FontFace interface.
type fontFace struct {
	gextras.Objector
}

var _ FontFace = (*fontFace)(nil)

// WrapFontFace wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontFace(obj *externglib.Object) FontFace {
	return FontFace{
		Objector: obj,
	}
}

func marshalFontFace(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontFace(obj), nil
}

// Describe returns the family, style, variant, weight and stretch of a
// `PangoFontFace`. The size field of the resulting font description will be
// unset.
func (f fontFace) Describe(f FontFace) {
	var arg0 *C.PangoFontFace

	arg0 = (*C.PangoFontFace)(unsafe.Pointer(f.Native()))

	C.pango_font_face_describe(arg0)
}

// FaceName gets a name representing the style of this face among the
// different faces in the `PangoFontFamily` for the face. The name is
// suitable for displaying to users.
func (f fontFace) FaceName(f FontFace) {
	var arg0 *C.PangoFontFace

	arg0 = (*C.PangoFontFace)(unsafe.Pointer(f.Native()))

	C.pango_font_face_get_face_name(arg0)
}

// Family gets the `PangoFontFamily` that @face belongs to.
func (f fontFace) Family(f FontFace) {
	var arg0 *C.PangoFontFace

	arg0 = (*C.PangoFontFace)(unsafe.Pointer(f.Native()))

	C.pango_font_face_get_family(arg0)
}

// IsSynthesized returns whether a `PangoFontFace` is synthesized by the
// underlying font rendering engine from another face, perhaps by shearing,
// emboldening, or lightening it.
func (f fontFace) IsSynthesized(f FontFace) bool {
	var arg0 *C.PangoFontFace

	arg0 = (*C.PangoFontFace)(unsafe.Pointer(f.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.pango_font_face_is_synthesized(arg0)

	if cret {
		ok = true
	}

	return ok
}

// ListSizes: list the available sizes for a font.
//
// This is only applicable to bitmap fonts. For scalable fonts, stores nil
// at the location pointed to by @sizes and 0 at the location pointed to by
// @n_sizes. The sizes returned are in Pango units and are sorted in
// ascending order.
func (f fontFace) ListSizes(f FontFace) {
	var arg0 *C.PangoFontFace

	arg0 = (*C.PangoFontFace)(unsafe.Pointer(f.Native()))

	C.pango_font_face_list_sizes(arg0, &arg1, &arg2)

	return sizes, nSizes
}

// FontFamily: a `PangoFontFamily` is used to represent a family of related font
// faces.
//
// The font faces in a family share a common design, but differ in slant,
// weight, width or other aspects.
type FontFamily interface {
	gextras.Objector

	// Face gets the `PangoFontFace` of @family with the given name.
	Face(f FontFamily, name string)
	// Name gets the name of the family.
	//
	// The name is unique among all fonts for the font backend and can be used
	// in a `PangoFontDescription` to specify that a face from this family is
	// desired.
	Name(f FontFamily)
	// IsMonospace: a monospace font is a font designed for text display where
	// the the characters form a regular grid.
	//
	// For Western languages this would mean that the advance width of all
	// characters are the same, but this categorization also includes Asian
	// fonts which include double-width characters: characters that occupy two
	// grid cells. g_unichar_iswide() returns a result that indicates whether a
	// character is typically double-width in a monospace font.
	//
	// The best way to find out the grid-cell size is to call
	// [method@Pango.FontMetrics.get_approximate_digit_width], since the results
	// of [method@Pango.FontMetrics.get_approximate_char_width] may be affected
	// by double-width characters.
	IsMonospace(f FontFamily) bool
	// IsVariable: a variable font is a font which has axes that can be modified
	// to produce different faces.
	IsVariable(f FontFamily) bool
	// ListFaces lists the different font faces that make up @family.
	//
	// The faces in a family share a common design, but differ in slant, weight,
	// width and other aspects.
	ListFaces(f FontFamily)
}

// fontFamily implements the FontFamily interface.
type fontFamily struct {
	gextras.Objector
}

var _ FontFamily = (*fontFamily)(nil)

// WrapFontFamily wraps a GObject to the right type. It is
// primarily used internally.
func WrapFontFamily(obj *externglib.Object) FontFamily {
	return FontFamily{
		Objector: obj,
	}
}

func marshalFontFamily(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFontFamily(obj), nil
}

// Face gets the `PangoFontFace` of @family with the given name.
func (f fontFamily) Face(f FontFamily, name string) {
	var arg0 *C.PangoFontFamily
	var arg1 *C.char

	arg0 = (*C.PangoFontFamily)(unsafe.Pointer(f.Native()))
	arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	C.pango_font_family_get_face(arg0, arg1)
}

// Name gets the name of the family.
//
// The name is unique among all fonts for the font backend and can be used
// in a `PangoFontDescription` to specify that a face from this family is
// desired.
func (f fontFamily) Name(f FontFamily) {
	var arg0 *C.PangoFontFamily

	arg0 = (*C.PangoFontFamily)(unsafe.Pointer(f.Native()))

	C.pango_font_family_get_name(arg0)
}

// IsMonospace: a monospace font is a font designed for text display where
// the the characters form a regular grid.
//
// For Western languages this would mean that the advance width of all
// characters are the same, but this categorization also includes Asian
// fonts which include double-width characters: characters that occupy two
// grid cells. g_unichar_iswide() returns a result that indicates whether a
// character is typically double-width in a monospace font.
//
// The best way to find out the grid-cell size is to call
// [method@Pango.FontMetrics.get_approximate_digit_width], since the results
// of [method@Pango.FontMetrics.get_approximate_char_width] may be affected
// by double-width characters.
func (f fontFamily) IsMonospace(f FontFamily) bool {
	var arg0 *C.PangoFontFamily

	arg0 = (*C.PangoFontFamily)(unsafe.Pointer(f.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.pango_font_family_is_monospace(arg0)

	if cret {
		ok = true
	}

	return ok
}

// IsVariable: a variable font is a font which has axes that can be modified
// to produce different faces.
func (f fontFamily) IsVariable(f FontFamily) bool {
	var arg0 *C.PangoFontFamily

	arg0 = (*C.PangoFontFamily)(unsafe.Pointer(f.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.pango_font_family_is_variable(arg0)

	if cret {
		ok = true
	}

	return ok
}

// ListFaces lists the different font faces that make up @family.
//
// The faces in a family share a common design, but differ in slant, weight,
// width and other aspects.
func (f fontFamily) ListFaces(f FontFamily) {
	var arg0 *C.PangoFontFamily

	arg0 = (*C.PangoFontFamily)(unsafe.Pointer(f.Native()))

	C.pango_font_family_list_faces(arg0, &arg1, &arg2)

	return faces, nFaces
}

// FontDescription: a `PangoFontDescription` describes a font in an
// implementation-independent manner.
//
// `PangoFontDescription` structures are used both to list what fonts are
// available on the system and also for specifying the characteristics of a font
// to load.
type FontDescription struct {
	native C.PangoFontDescription
}

// WrapFontDescription wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFontDescription(ptr unsafe.Pointer) *FontDescription {
	if ptr == nil {
		return nil
	}

	return (*FontDescription)(ptr)
}

func marshalFontDescription(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapFontDescription(unsafe.Pointer(b)), nil
}

// NewFontDescription constructs a struct FontDescription.
func NewFontDescription() {
	C.pango_font_description_new()
}

// Native returns the underlying C source pointer.
func (f *FontDescription) Native() unsafe.Pointer {
	return unsafe.Pointer(&f.native)
}

// BetterMatch determines if the style attributes of @new_match are a closer
// match for @desc than those of @old_match are, or if @old_match is nil,
// determines if @new_match is a match at all.
//
// Approximate matching is done for weight and style; other style attributes
// must match exactly. Style attributes are all attributes other than family and
// size-related attributes. Approximate matching for style considers
// PANGO_STYLE_OBLIQUE and PANGO_STYLE_ITALIC as matches, but not as good a
// match as when the styles are equal.
//
// Note that @old_match must match @desc.
func (d *FontDescription) BetterMatch(d *FontDescription, oldMatch *FontDescription, newMatch *FontDescription) bool {
	var arg0 *C.PangoFontDescription
	var arg1 *C.PangoFontDescription
	var arg2 *C.PangoFontDescription

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	arg1 = (*C.PangoFontDescription)(unsafe.Pointer(oldMatch.Native()))
	arg2 = (*C.PangoFontDescription)(unsafe.Pointer(newMatch.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.pango_font_description_better_match(arg0, arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// Copy: make a copy of a `PangoFontDescription`.
func (d *FontDescription) Copy(d *FontDescription) {
	var arg0 *C.PangoFontDescription

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	C.pango_font_description_copy(arg0)
}

// CopyStatic: make a copy of a `PangoFontDescription`, but don't duplicate
// allocated fields.
//
// This is like [method@Pango.FontDescription.copy], but only a shallow copy is
// made of the family name and other allocated fields. The result can only be
// used until @desc is modified or freed. This is meant to be used when the copy
// is only needed temporarily.
func (d *FontDescription) CopyStatic(d *FontDescription) {
	var arg0 *C.PangoFontDescription

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	C.pango_font_description_copy_static(arg0)
}

// Equal compares two font descriptions for equality.
//
// Two font descriptions are considered equal if the fonts they describe are
// provably identical. This means that their masks do not have to match, as long
// as other fields are all the same. (Two font descriptions may result in
// identical fonts being loaded, but still compare false.)
func (d *FontDescription) Equal(d *FontDescription, desc2 *FontDescription) bool {
	var arg0 *C.PangoFontDescription
	var arg1 *C.PangoFontDescription

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	arg1 = (*C.PangoFontDescription)(unsafe.Pointer(desc2.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.pango_font_description_equal(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// Free frees a font description.
func (d *FontDescription) Free(d *FontDescription) {
	var arg0 *C.PangoFontDescription

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	C.pango_font_description_free(arg0)
}

// Family gets the family name field of a font description.
//
// See [method@Pango.FontDescription.set_family].
func (d *FontDescription) Family(d *FontDescription) {
	var arg0 *C.PangoFontDescription

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	C.pango_font_description_get_family(arg0)
}

// Gravity gets the gravity field of a font description.
//
// See [method@Pango.FontDescription.set_gravity].
func (d *FontDescription) Gravity(d *FontDescription) {
	var arg0 *C.PangoFontDescription

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	C.pango_font_description_get_gravity(arg0)
}

// SetFields determines which fields in a font description have been set.
func (d *FontDescription) SetFields(d *FontDescription) {
	var arg0 *C.PangoFontDescription

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	C.pango_font_description_get_set_fields(arg0)
}

// Size gets the size field of a font description.
//
// See [method@Pango.FontDescription.set_size].
func (d *FontDescription) Size(d *FontDescription) {
	var arg0 *C.PangoFontDescription

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	C.pango_font_description_get_size(arg0)
}

// SizeIsAbsolute determines whether the size of the font is in points (not
// absolute) or device units (absolute).
//
// See [method@Pango.FontDescription.set_size] and
// [method@Pango.FontDescription.set_absolute_size].
func (d *FontDescription) SizeIsAbsolute(d *FontDescription) bool {
	var arg0 *C.PangoFontDescription

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.pango_font_description_get_size_is_absolute(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Stretch gets the stretch field of a font description.
//
// See [method@Pango.FontDescription.set_stretch].
func (d *FontDescription) Stretch(d *FontDescription) {
	var arg0 *C.PangoFontDescription

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	C.pango_font_description_get_stretch(arg0)
}

// Style gets the style field of a `PangoFontDescription`.
//
// See [method@Pango.FontDescription.set_style].
func (d *FontDescription) Style(d *FontDescription) {
	var arg0 *C.PangoFontDescription

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	C.pango_font_description_get_style(arg0)
}

// Variant gets the variant field of a `PangoFontDescription`.
//
// See [method@Pango.FontDescription.set_variant].
func (d *FontDescription) Variant(d *FontDescription) {
	var arg0 *C.PangoFontDescription

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	C.pango_font_description_get_variant(arg0)
}

// Variations gets the variations field of a font description.
//
// See [method@Pango.FontDescription.set_variations].
func (d *FontDescription) Variations(d *FontDescription) {
	var arg0 *C.PangoFontDescription

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	C.pango_font_description_get_variations(arg0)
}

// Weight gets the weight field of a font description.
//
// See [method@Pango.FontDescription.set_weight].
func (d *FontDescription) Weight(d *FontDescription) {
	var arg0 *C.PangoFontDescription

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	C.pango_font_description_get_weight(arg0)
}

// Hash computes a hash of a `PangoFontDescription` structure.
//
// This is suitable to be used, for example, as an argument to
// g_hash_table_new(). The hash value is independent of @desc->mask.
func (d *FontDescription) Hash(d *FontDescription) {
	var arg0 *C.PangoFontDescription

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	C.pango_font_description_hash(arg0)
}

// Merge merges the fields that are set in @desc_to_merge into the fields in
// @desc.
//
// If @replace_existing is false, only fields in @desc that are not already set
// are affected. If true, then fields that are already set will be replaced as
// well.
//
// If @desc_to_merge is nil, this function performs nothing.
func (d *FontDescription) Merge(d *FontDescription, descToMerge *FontDescription, replaceExisting bool) {
	var arg0 *C.PangoFontDescription
	var arg1 *C.PangoFontDescription
	var arg2 C.gboolean

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	arg1 = (*C.PangoFontDescription)(unsafe.Pointer(descToMerge.Native()))
	if replaceExisting {
		arg2 = C.gboolean(1)
	}

	C.pango_font_description_merge(arg0, arg1, arg2)
}

// MergeStatic merges the fields that are set in @desc_to_merge into the fields
// in @desc, without copying allocated fields.
//
// This is like [method@Pango.FontDescription.merge], but only a shallow copy is
// made of the family name and other allocated fields. @desc can only be used
// until @desc_to_merge is modified or freed. This is meant to be used when the
// merged font description is only needed temporarily.
func (d *FontDescription) MergeStatic(d *FontDescription, descToMerge *FontDescription, replaceExisting bool) {
	var arg0 *C.PangoFontDescription
	var arg1 *C.PangoFontDescription
	var arg2 C.gboolean

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	arg1 = (*C.PangoFontDescription)(unsafe.Pointer(descToMerge.Native()))
	if replaceExisting {
		arg2 = C.gboolean(1)
	}

	C.pango_font_description_merge_static(arg0, arg1, arg2)
}

// SetAbsoluteSize sets the size field of a font description, in device units.
//
// This is mutually exclusive with [method@Pango.FontDescription.set_size] which
// sets the font size in points.
func (d *FontDescription) SetAbsoluteSize(d *FontDescription, size float64) {
	var arg0 *C.PangoFontDescription
	var arg1 C.double

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	arg1 = C.double(size)

	C.pango_font_description_set_absolute_size(arg0, arg1)
}

// SetFamily sets the family name field of a font description.
//
// The family name represents a family of related font styles, and will resolve
// to a particular `PangoFontFamily`. In some uses of `PangoFontDescription`, it
// is also possible to use a comma separated list of family names for this
// field.
func (d *FontDescription) SetFamily(d *FontDescription, family string) {
	var arg0 *C.PangoFontDescription
	var arg1 *C.char

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	arg1 = (*C.char)(C.CString(family))
	defer C.free(unsafe.Pointer(arg1))

	C.pango_font_description_set_family(arg0, arg1)
}

// SetFamilyStatic sets the family name field of a font description, without
// copying the string.
//
// This is like [method@Pango.FontDescription.set_family], except that no copy
// of @family is made. The caller must make sure that the string passed in stays
// around until @desc has been freed or the name is set again. This function can
// be used if @family is a static string such as a C string literal, or if @desc
// is only needed temporarily.
func (d *FontDescription) SetFamilyStatic(d *FontDescription, family string) {
	var arg0 *C.PangoFontDescription
	var arg1 *C.char

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	arg1 = (*C.char)(C.CString(family))
	defer C.free(unsafe.Pointer(arg1))

	C.pango_font_description_set_family_static(arg0, arg1)
}

// SetGravity sets the gravity field of a font description.
//
// The gravity field specifies how the glyphs should be rotated. If @gravity is
// PANGO_GRAVITY_AUTO, this actually unsets the gravity mask on the font
// description.
//
// This function is seldom useful to the user. Gravity should normally be set on
// a `PangoContext`.
func (d *FontDescription) SetGravity(d *FontDescription, gravity Gravity) {
	var arg0 *C.PangoFontDescription
	var arg1 C.PangoGravity

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	arg1 = (C.PangoGravity)(gravity)

	C.pango_font_description_set_gravity(arg0, arg1)
}

// SetSize sets the size field of a font description in fractional points.
//
// This is mutually exclusive with
// [method@Pango.FontDescription.set_absolute_size].
func (d *FontDescription) SetSize(d *FontDescription, size int) {
	var arg0 *C.PangoFontDescription
	var arg1 C.gint

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	arg1 = C.gint(size)

	C.pango_font_description_set_size(arg0, arg1)
}

// SetStretch sets the stretch field of a font description.
//
// The [enum@Pango.Stretch] field specifies how narrow or wide the font should
// be.
func (d *FontDescription) SetStretch(d *FontDescription, stretch Stretch) {
	var arg0 *C.PangoFontDescription
	var arg1 C.PangoStretch

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	arg1 = (C.PangoStretch)(stretch)

	C.pango_font_description_set_stretch(arg0, arg1)
}

// SetStyle sets the style field of a `PangoFontDescription`.
//
// The [enum@Pango.Style] enumeration describes whether the font is slanted and
// the manner in which it is slanted; it can be either NGO_STYLE_NORMAL,
// NGO_STYLE_ITALIC, or NGO_STYLE_OBLIQUE.
//
// Most fonts will either have a italic style or an oblique style, but not both,
// and font matching in Pango will match italic specifications with oblique
// fonts and vice-versa if an exact match is not found.
func (d *FontDescription) SetStyle(d *FontDescription, style Style) {
	var arg0 *C.PangoFontDescription
	var arg1 C.PangoStyle

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	arg1 = (C.PangoStyle)(style)

	C.pango_font_description_set_style(arg0, arg1)
}

// SetVariant sets the variant field of a font description.
//
// The [enum@Pango.Variant] can either be PANGO_VARIANT_NORMAL or
// PANGO_VARIANT_SMALL_CAPS.
func (d *FontDescription) SetVariant(d *FontDescription, variant Variant) {
	var arg0 *C.PangoFontDescription
	var arg1 C.PangoVariant

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	arg1 = (C.PangoVariant)(variant)

	C.pango_font_description_set_variant(arg0, arg1)
}

// SetVariations sets the variations field of a font description.
//
// OpenType font variations allow to select a font instance by specifying values
// for a number of axes, such as width or weight.
//
// The format of the variations string is
//
//    AXIS1=VALUE,AXIS2=VALUE...
//
// with each AXIS a 4 character tag that identifies a font axis, and each VALUE
// a floating point number. Unknown axes are ignored, and values are clamped to
// their allowed range.
//
// Pango does not currently have a way to find supported axes of a font. Both
// harfbuzz or freetype have API for this.
func (d *FontDescription) SetVariations(d *FontDescription, variations string) {
	var arg0 *C.PangoFontDescription
	var arg1 *C.char

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	arg1 = (*C.char)(C.CString(variations))
	defer C.free(unsafe.Pointer(arg1))

	C.pango_font_description_set_variations(arg0, arg1)
}

// SetVariationsStatic sets the variations field of a font description.
//
// This is like [method@Pango.FontDescription.set_variations], except that no
// copy of @variations is made. The caller must make sure that the string passed
// in stays around until @desc has been freed or the name is set again. This
// function can be used if @variations is a static string such as a C string
// literal, or if @desc is only needed temporarily.
func (d *FontDescription) SetVariationsStatic(d *FontDescription, variations string) {
	var arg0 *C.PangoFontDescription
	var arg1 *C.char

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	arg1 = (*C.char)(C.CString(variations))
	defer C.free(unsafe.Pointer(arg1))

	C.pango_font_description_set_variations_static(arg0, arg1)
}

// SetWeight sets the weight field of a font description.
//
// The weight field specifies how bold or light the font should be. In addition
// to the values of the [enum@Pango.Weight] enumeration, other intermediate
// numeric values are possible.
func (d *FontDescription) SetWeight(d *FontDescription, weight Weight) {
	var arg0 *C.PangoFontDescription
	var arg1 C.PangoWeight

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	arg1 = (C.PangoWeight)(weight)

	C.pango_font_description_set_weight(arg0, arg1)
}

// ToFilename creates a filename representation of a font description.
//
// The filename is identical to the result from calling
// [method@Pango.FontDescription.to_string], but with underscores instead of
// characters that are untypical in filenames, and in lower case only.
func (d *FontDescription) ToFilename(d *FontDescription) {
	var arg0 *C.PangoFontDescription

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	C.pango_font_description_to_filename(arg0)
}

// String creates a string representation of a font description.
//
// See [type_func@Pango.FontDescription.from_string] for a description of the
// format of the string representation. The family list in the string
// description will only have a terminating comma if the last word of the list
// is a valid style option.
func (d *FontDescription) String(d *FontDescription) {
	var arg0 *C.PangoFontDescription

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))

	C.pango_font_description_to_string(arg0)
}

// UnsetFields unsets some of the fields in a `PangoFontDescription`.
//
// The unset fields will get back to their default values.
func (d *FontDescription) UnsetFields(d *FontDescription, toUnset FontMask) {
	var arg0 *C.PangoFontDescription
	var arg1 C.PangoFontMask

	arg0 = (*C.PangoFontDescription)(unsafe.Pointer(d.Native()))
	arg1 = (C.PangoFontMask)(toUnset)

	C.pango_font_description_unset_fields(arg0, arg1)
}

// FontMetrics: a `PangoFontMetrics` structure holds the overall metric
// information for a font.
//
// The information in a `PangoFontMetrics` structure may be restricted to a
// script. The fields of this structure are private to implementations of a font
// backend. See the documentation of the corresponding getters for documentation
// of their meaning.
type FontMetrics struct {
	native C.PangoFontMetrics
}

// WrapFontMetrics wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFontMetrics(ptr unsafe.Pointer) *FontMetrics {
	if ptr == nil {
		return nil
	}

	return (*FontMetrics)(ptr)
}

func marshalFontMetrics(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapFontMetrics(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (f *FontMetrics) Native() unsafe.Pointer {
	return unsafe.Pointer(&f.native)
}

// ApproximateCharWidth gets the approximate character width for a font metrics
// structure.
//
// This is merely a representative value useful, for example, for determining
// the initial size for a window. Actual characters in text will be wider and
// narrower than this.
func (m *FontMetrics) ApproximateCharWidth(m *FontMetrics) {
	var arg0 *C.PangoFontMetrics

	arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	C.pango_font_metrics_get_approximate_char_width(arg0)
}

// ApproximateDigitWidth gets the approximate digit width for a font metrics
// structure.
//
// This is merely a representative value useful, for example, for determining
// the initial size for a window. Actual digits in text can be wider or narrower
// than this, though this value is generally somewhat more accurate than the
// result of pango_font_metrics_get_approximate_char_width() for digits.
func (m *FontMetrics) ApproximateDigitWidth(m *FontMetrics) {
	var arg0 *C.PangoFontMetrics

	arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	C.pango_font_metrics_get_approximate_digit_width(arg0)
}

// Ascent gets the ascent from a font metrics structure.
//
// The ascent is the distance from the baseline to the logical top of a line of
// text. (The logical top may be above or below the top of the actual drawn ink.
// It is necessary to lay out the text to figure where the ink will be.)
func (m *FontMetrics) Ascent(m *FontMetrics) {
	var arg0 *C.PangoFontMetrics

	arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	C.pango_font_metrics_get_ascent(arg0)
}

// Descent gets the descent from a font metrics structure.
//
// The descent is the distance from the baseline to the logical bottom of a line
// of text. (The logical bottom may be above or below the bottom of the actual
// drawn ink. It is necessary to lay out the text to figure where the ink will
// be.)
func (m *FontMetrics) Descent(m *FontMetrics) {
	var arg0 *C.PangoFontMetrics

	arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	C.pango_font_metrics_get_descent(arg0)
}

// Height gets the line height from a font metrics structure.
//
// The line height is the distance between successive baselines in wrapped text.
//
// If the line height is not available, 0 is returned.
func (m *FontMetrics) Height(m *FontMetrics) {
	var arg0 *C.PangoFontMetrics

	arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	C.pango_font_metrics_get_height(arg0)
}

// StrikethroughPosition gets the suggested position to draw the strikethrough.
//
// The value returned is the distance *above* the baseline of the top of the
// strikethrough.
func (m *FontMetrics) StrikethroughPosition(m *FontMetrics) {
	var arg0 *C.PangoFontMetrics

	arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	C.pango_font_metrics_get_strikethrough_position(arg0)
}

// StrikethroughThickness gets the suggested thickness to draw for the
// strikethrough.
func (m *FontMetrics) StrikethroughThickness(m *FontMetrics) {
	var arg0 *C.PangoFontMetrics

	arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	C.pango_font_metrics_get_strikethrough_thickness(arg0)
}

// UnderlinePosition gets the suggested position to draw the underline.
//
// The value returned is the distance *above* the baseline of the top of the
// underline. Since most fonts have underline positions beneath the baseline,
// this value is typically negative.
func (m *FontMetrics) UnderlinePosition(m *FontMetrics) {
	var arg0 *C.PangoFontMetrics

	arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	C.pango_font_metrics_get_underline_position(arg0)
}

// UnderlineThickness gets the suggested thickness to draw for the underline.
func (m *FontMetrics) UnderlineThickness(m *FontMetrics) {
	var arg0 *C.PangoFontMetrics

	arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	C.pango_font_metrics_get_underline_thickness(arg0)
}

// Ref: increase the reference count of a font metrics structure by one.
func (m *FontMetrics) Ref(m *FontMetrics) {
	var arg0 *C.PangoFontMetrics

	arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	C.pango_font_metrics_ref(arg0)
}

// Unref: decrease the reference count of a font metrics structure by one. If
// the result is zero, frees the structure and any associated memory.
func (m *FontMetrics) Unref(m *FontMetrics) {
	var arg0 *C.PangoFontMetrics

	arg0 = (*C.PangoFontMetrics)(unsafe.Pointer(m.Native()))

	C.pango_font_metrics_unref(arg0)
}