// Code generated by girgen. DO NOT EDIT.

package pango

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pango glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_attr_iterator_get_type()), F: marshalAttrIterator},
		{T: externglib.Type(C.pango_attr_list_get_type()), F: marshalAttrList},
		{T: externglib.Type(C.pango_attribute_get_type()), F: marshalAttribute},
		{T: externglib.Type(C.pango_color_get_type()), F: marshalColor},
	})
}

// AttrDataCopyFunc: type of a function that can duplicate user data for an
// attribute.
type AttrDataCopyFunc func() (gpointer interface{})

//export gotk4_AttrDataCopyFunc
func gotk4_AttrDataCopyFunc(arg0 C.gpointer) C.gpointer {
	v := box.Get(uintptr(arg0))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(AttrDataCopyFunc)
	gpointer := fn()

	cret = C.gpointer(gpointer)

	return gpointer
}

// MarkupParserFinish finishes parsing markup.
//
// After feeding a Pango markup parser some data with
// g_markup_parse_context_parse(), use this function to get the list of
// attributes and text out of the markup. This function will not free @context,
// use g_markup_parse_context_free() to do so.
func MarkupParserFinish(context *glib.MarkupParseContext) (*AttrList, string, uint32, error) {
	var _arg1 *C.GMarkupParseContext // out

	_arg1 = (*C.GMarkupParseContext)(unsafe.Pointer(context.Native()))

	var _attrList *AttrList
	var _arg3 *C.char    // in
	var _arg4 C.gunichar // in
	var _cerr *C.GError  // in

	C.pango_markup_parser_finish(_arg1, (**C.PangoAttrList)(unsafe.Pointer(&_attrList)), &_arg3, &_arg4, &_cerr)

	var _text string      // out
	var _accelChar uint32 // out
	var _goerr error      // out

	_text = C.GoString(_arg3)
	defer C.free(unsafe.Pointer(_arg3))
	_accelChar = (uint32)(_arg4)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _attrList, _text, _accelChar, _goerr
}

// ParseMarkup parses marked-up text to create a plain-text string and an
// attribute list.
//
// See the Pango Markup (pango_markup.html) docs for details about the supported
// markup.
//
// If @accel_marker is nonzero, the given character will mark the character
// following it as an accelerator. For example, @accel_marker might be an
// ampersand or underscore. All characters marked as an accelerator will receive
// a PANGO_UNDERLINE_LOW attribute, and the first character so marked will be
// returned in @accel_char. Two @accel_marker characters following each other
// produce a single literal @accel_marker character.
//
// To parse a stream of pango markup incrementally, use
// [func@markup_parser_new].
//
// If any error happens, none of the output arguments are touched except for
// @error.
func ParseMarkup(markupText string, length int, accelMarker uint32) (*AttrList, string, uint32, error) {
	var _arg1 *C.char    // out
	var _arg2 C.int      // out
	var _arg3 C.gunichar // out

	_arg1 = (*C.char)(C.CString(markupText))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(length)
	_arg3 = C.gunichar(accelMarker)

	var _attrList *AttrList
	var _arg5 *C.char    // in
	var _arg6 C.gunichar // in
	var _cerr *C.GError  // in

	C.pango_parse_markup(_arg1, _arg2, _arg3, (**C.PangoAttrList)(unsafe.Pointer(&_attrList)), &_arg5, &_arg6, &_cerr)

	var _text string      // out
	var _accelChar uint32 // out
	var _goerr error      // out

	_text = C.GoString(_arg5)
	defer C.free(unsafe.Pointer(_arg5))
	_accelChar = (uint32)(_arg6)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _attrList, _text, _accelChar, _goerr
}

// AttrColor: the `PangoAttrColor` structure is used to represent attributes
// that are colors.
type AttrColor struct {
	native C.PangoAttrColor
}

// WrapAttrColor wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAttrColor(ptr unsafe.Pointer) *AttrColor {
	if ptr == nil {
		return nil
	}

	return (*AttrColor)(ptr)
}

func marshalAttrColor(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapAttrColor(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *AttrColor) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// AttrFloat: the `PangoAttrFloat` structure is used to represent attributes
// with a float or double value.
type AttrFloat struct {
	native C.PangoAttrFloat
}

// WrapAttrFloat wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAttrFloat(ptr unsafe.Pointer) *AttrFloat {
	if ptr == nil {
		return nil
	}

	return (*AttrFloat)(ptr)
}

func marshalAttrFloat(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapAttrFloat(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *AttrFloat) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// Value gets the field inside the struct.
func (a *AttrFloat) Value() float64 {
	var v float64 // out
	v = (float64)(a.native.value)
	return v
}

// AttrFontDesc: the `PangoAttrFontDesc` structure is used to store an attribute
// that sets all aspects of the font description at once.
type AttrFontDesc struct {
	native C.PangoAttrFontDesc
}

// WrapAttrFontDesc wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAttrFontDesc(ptr unsafe.Pointer) *AttrFontDesc {
	if ptr == nil {
		return nil
	}

	return (*AttrFontDesc)(ptr)
}

func marshalAttrFontDesc(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapAttrFontDesc(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *AttrFontDesc) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// AttrFontFeatures: the `PangoAttrFontFeatures` structure is used to represent
// OpenType font features as an attribute.
type AttrFontFeatures struct {
	native C.PangoAttrFontFeatures
}

// WrapAttrFontFeatures wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAttrFontFeatures(ptr unsafe.Pointer) *AttrFontFeatures {
	if ptr == nil {
		return nil
	}

	return (*AttrFontFeatures)(ptr)
}

func marshalAttrFontFeatures(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapAttrFontFeatures(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *AttrFontFeatures) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// Features gets the field inside the struct.
func (a *AttrFontFeatures) Features() string {
	var v string // out
	v = C.GoString(a.native.features)
	return v
}

// AttrInt: the `PangoAttrInt` structure is used to represent attributes with an
// integer or enumeration value.
type AttrInt struct {
	native C.PangoAttrInt
}

// WrapAttrInt wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAttrInt(ptr unsafe.Pointer) *AttrInt {
	if ptr == nil {
		return nil
	}

	return (*AttrInt)(ptr)
}

func marshalAttrInt(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapAttrInt(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *AttrInt) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// Value gets the field inside the struct.
func (a *AttrInt) Value() int {
	var v int // out
	v = (int)(a.native.value)
	return v
}

// AttrIterator: a `PangoAttrIterator` is used to iterate through a
// `PangoAttrList`.
//
// A new iterator is created with [method@Pango.AttrList.get_iterator]. Once the
// iterator is created, it can be advanced through the style changes in the text
// using [method@Pango.AttrIterator.next]. At each style change, the range of
// the current style segment and the attributes currently in effect can be
// queried.
type AttrIterator struct {
	native C.PangoAttrIterator
}

// WrapAttrIterator wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAttrIterator(ptr unsafe.Pointer) *AttrIterator {
	if ptr == nil {
		return nil
	}

	return (*AttrIterator)(ptr)
}

func marshalAttrIterator(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapAttrIterator(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *AttrIterator) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// Destroy: destroy a `PangoAttrIterator` and free all associated memory.
func (i *AttrIterator) Destroy() {
	var _arg0 *C.PangoAttrIterator // out

	_arg0 = (*C.PangoAttrIterator)(unsafe.Pointer(i.Native()))

	C.pango_attr_iterator_destroy(_arg0)
}

// Font: get the font and other attributes at the current iterator position.
func (i *AttrIterator) Font(desc *FontDescription, language **Language, extraAttrs **glib.SList) {
	var _arg0 *C.PangoAttrIterator    // out
	var _arg1 *C.PangoFontDescription // out
	var _arg2 **C.PangoLanguage       // out
	var _arg3 **C.GSList              // out

	_arg0 = (*C.PangoAttrIterator)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.PangoFontDescription)(unsafe.Pointer(desc.Native()))
	_arg2 = (**C.PangoLanguage)(unsafe.Pointer(language.Native()))
	_arg3 = (**C.GSList)(unsafe.Pointer(extraAttrs.Native()))

	C.pango_attr_iterator_get_font(_arg0, _arg1, _arg2, _arg3)
}

// Next: advance the iterator until the next change of style.
func (i *AttrIterator) Next() bool {
	var _arg0 *C.PangoAttrIterator // out

	_arg0 = (*C.PangoAttrIterator)(unsafe.Pointer(i.Native()))

	var _cret C.gboolean // in

	_cret = C.pango_attr_iterator_next(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Range: get the range of the current segment. Note that the stored return
// values are signed, not unsigned like the values in `PangoAttribute`. To deal
// with this API oversight, stored return values that wouldn't fit into a signed
// integer are clamped to G_MAXINT.
func (i *AttrIterator) Range() (start int, end int) {
	var _arg0 *C.PangoAttrIterator // out

	_arg0 = (*C.PangoAttrIterator)(unsafe.Pointer(i.Native()))

	var _arg1 C.gint // in
	var _arg2 C.gint // in

	C.pango_attr_iterator_range(_arg0, &_arg1, &_arg2)

	var _start int // out
	var _end int   // out

	_start = (int)(_arg1)
	_end = (int)(_arg2)

	return _start, _end
}

// AttrLanguage: the `PangoAttrLanguage` structure is used to represent
// attributes that are languages.
type AttrLanguage struct {
	native C.PangoAttrLanguage
}

// WrapAttrLanguage wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAttrLanguage(ptr unsafe.Pointer) *AttrLanguage {
	if ptr == nil {
		return nil
	}

	return (*AttrLanguage)(ptr)
}

func marshalAttrLanguage(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapAttrLanguage(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *AttrLanguage) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// AttrList: a `PangoAttrList` represents a list of attributes that apply to a
// section of text.
//
// The attributes in a `PangoAttrList` are, in general, allowed to overlap in an
// arbitrary fashion. However, if the attributes are manipulated only through
// [method@Pango.AttrList.change], the overlap between properties will meet
// stricter criteria.
//
// Since the `PangoAttrList` structure is stored as a linear list, it is not
// suitable for storing attributes for large amounts of text. In general, you
// should not use a single `PangoAttrList` for more than one paragraph of text.
type AttrList struct {
	native C.PangoAttrList
}

// WrapAttrList wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAttrList(ptr unsafe.Pointer) *AttrList {
	if ptr == nil {
		return nil
	}

	return (*AttrList)(ptr)
}

func marshalAttrList(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapAttrList(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *AttrList) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// Change: insert the given attribute into the `PangoAttrList`.
//
// It will replace any attributes of the same type on that segment and be merged
// with any adjoining attributes that are identical.
//
// This function is slower than [method@Pango.AttrList.insert] for creating an
// attribute list in order (potentially much slower for large lists). However,
// [method@Pango.AttrList.insert] is not suitable for continually changing a set
// of attributes since it never removes or combines existing attributes.
func (l *AttrList) Change(attr *Attribute) {
	var _arg0 *C.PangoAttrList  // out
	var _arg1 *C.PangoAttribute // out

	_arg0 = (*C.PangoAttrList)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.PangoAttribute)(unsafe.Pointer(attr.Native()))

	C.pango_attr_list_change(_arg0, _arg1)
}

// Equal checks whether @list and @other_list contain the same attributes and
// whether those attributes apply to the same ranges. Beware that this will
// return wrong values if any list contains duplicates.
func (l *AttrList) Equal(otherList *AttrList) bool {
	var _arg0 *C.PangoAttrList // out
	var _arg1 *C.PangoAttrList // out

	_arg0 = (*C.PangoAttrList)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.PangoAttrList)(unsafe.Pointer(otherList.Native()))

	var _cret C.gboolean // in

	_cret = C.pango_attr_list_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Insert: insert the given attribute into the `PangoAttrList`.
//
// It will be inserted after all other attributes with a matching @start_index.
func (l *AttrList) Insert(attr *Attribute) {
	var _arg0 *C.PangoAttrList  // out
	var _arg1 *C.PangoAttribute // out

	_arg0 = (*C.PangoAttrList)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.PangoAttribute)(unsafe.Pointer(attr.Native()))

	C.pango_attr_list_insert(_arg0, _arg1)
}

// InsertBefore: insert the given attribute into the `PangoAttrList`.
//
// It will be inserted before all other attributes with a matching @start_index.
func (l *AttrList) InsertBefore(attr *Attribute) {
	var _arg0 *C.PangoAttrList  // out
	var _arg1 *C.PangoAttribute // out

	_arg0 = (*C.PangoAttrList)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.PangoAttribute)(unsafe.Pointer(attr.Native()))

	C.pango_attr_list_insert_before(_arg0, _arg1)
}

// Splice: this function opens up a hole in @list, fills it in with attributes
// from the left, and then merges @other on top of the hole.
//
// This operation is equivalent to stretching every attribute that applies at
// position @pos in @list by an amount @len, and then calling
// [method@Pango.AttrList.change] with a copy of each attribute in @other in
// sequence (offset in position by @pos).
//
// This operation proves useful for, for instance, inserting a pre-edit string
// in the middle of an edit buffer.
func (l *AttrList) Splice(other *AttrList, pos int, len int) {
	var _arg0 *C.PangoAttrList // out
	var _arg1 *C.PangoAttrList // out
	var _arg2 C.gint           // out
	var _arg3 C.gint           // out

	_arg0 = (*C.PangoAttrList)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.PangoAttrList)(unsafe.Pointer(other.Native()))
	_arg2 = C.gint(pos)
	_arg3 = C.gint(len)

	C.pango_attr_list_splice(_arg0, _arg1, _arg2, _arg3)
}

// Unref: decrease the reference count of the given attribute list by one. If
// the result is zero, free the attribute list and the attributes it contains.
func (l *AttrList) Unref() {
	var _arg0 *C.PangoAttrList // out

	_arg0 = (*C.PangoAttrList)(unsafe.Pointer(l.Native()))

	C.pango_attr_list_unref(_arg0)
}

// Update: update indices of attributes in @list for a change in the text they
// refer to.
//
// The change that this function applies is removing @remove bytes at position
// @pos and inserting @add bytes instead.
//
// Attributes that fall entirely in the (@pos, @pos + @remove) range are
// removed.
//
// Attributes that start or end inside the (@pos, @pos + @remove) range are
// shortened to reflect the removal.
//
// Attributes start and end positions are updated if they are behind @pos +
// @remove.
func (l *AttrList) Update(pos int, remove int, add int) {
	var _arg0 *C.PangoAttrList // out
	var _arg1 C.int            // out
	var _arg2 C.int            // out
	var _arg3 C.int            // out

	_arg0 = (*C.PangoAttrList)(unsafe.Pointer(l.Native()))
	_arg1 = C.int(pos)
	_arg2 = C.int(remove)
	_arg3 = C.int(add)

	C.pango_attr_list_update(_arg0, _arg1, _arg2, _arg3)
}

// AttrShape: the `PangoAttrShape` structure is used to represent attributes
// which impose shape restrictions.
type AttrShape struct {
	native C.PangoAttrShape
}

// WrapAttrShape wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAttrShape(ptr unsafe.Pointer) *AttrShape {
	if ptr == nil {
		return nil
	}

	return (*AttrShape)(ptr)
}

func marshalAttrShape(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapAttrShape(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *AttrShape) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// Data gets the field inside the struct.
func (a *AttrShape) Data() interface{} {
	var v interface{} // out
	v = (interface{})(a.native.data)
	return v
}

// AttrSize: the `PangoAttrSize` structure is used to represent attributes which
// set font size.
type AttrSize struct {
	native C.PangoAttrSize
}

// WrapAttrSize wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAttrSize(ptr unsafe.Pointer) *AttrSize {
	if ptr == nil {
		return nil
	}

	return (*AttrSize)(ptr)
}

func marshalAttrSize(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapAttrSize(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *AttrSize) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// Size gets the field inside the struct.
func (a *AttrSize) Size() int {
	var v int // out
	v = (int)(a.native.size)
	return v
}

// AttrString: the `PangoAttrString` structure is used to represent attributes
// with a string value.
type AttrString struct {
	native C.PangoAttrString
}

// WrapAttrString wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAttrString(ptr unsafe.Pointer) *AttrString {
	if ptr == nil {
		return nil
	}

	return (*AttrString)(ptr)
}

func marshalAttrString(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapAttrString(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *AttrString) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// Value gets the field inside the struct.
func (a *AttrString) Value() string {
	var v string // out
	v = C.GoString(a.native.value)
	return v
}

// Attribute: the `PangoAttribute` structure represents the common portions of
// all attributes.
//
// Particular types of attributes include this structure as their initial
// portion. The common portion of the attribute holds the range to which the
// value in the type-specific part of the attribute applies and should be
// initialized using [method@Pango.Attribute.init]. By default, an attribute
// will have an all-inclusive range of [0,G_MAXUINT].
type Attribute struct {
	native C.PangoAttribute
}

// WrapAttribute wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapAttribute(ptr unsafe.Pointer) *Attribute {
	if ptr == nil {
		return nil
	}

	return (*Attribute)(ptr)
}

func marshalAttribute(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapAttribute(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (a *Attribute) Native() unsafe.Pointer {
	return unsafe.Pointer(&a.native)
}

// StartIndex gets the field inside the struct.
func (a *Attribute) StartIndex() uint {
	var v uint // out
	v = (uint)(a.native.start_index)
	return v
}

// EndIndex gets the field inside the struct.
func (a *Attribute) EndIndex() uint {
	var v uint // out
	v = (uint)(a.native.end_index)
	return v
}

// Destroy: destroy a `PangoAttribute` and free all associated memory.
func (a *Attribute) Destroy() {
	var _arg0 *C.PangoAttribute // out

	_arg0 = (*C.PangoAttribute)(unsafe.Pointer(a.Native()))

	C.pango_attribute_destroy(_arg0)
}

// Equal: compare two attributes for equality. This compares only the actual
// value of the two attributes and not the ranges that the attributes apply to.
func (a *Attribute) Equal(attr2 *Attribute) bool {
	var _arg0 *C.PangoAttribute // out
	var _arg1 *C.PangoAttribute // out

	_arg0 = (*C.PangoAttribute)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.PangoAttribute)(unsafe.Pointer(attr2.Native()))

	var _cret C.gboolean // in

	_cret = C.pango_attribute_equal(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Color: the `PangoColor` structure is used to represent a color in an
// uncalibrated RGB color-space.
type Color struct {
	native C.PangoColor
}

// WrapColor wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapColor(ptr unsafe.Pointer) *Color {
	if ptr == nil {
		return nil
	}

	return (*Color)(ptr)
}

func marshalColor(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapColor(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (c *Color) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}

// Red gets the field inside the struct.
func (c *Color) Red() uint16 {
	var v uint16 // out
	v = (uint16)(c.native.red)
	return v
}

// Green gets the field inside the struct.
func (c *Color) Green() uint16 {
	var v uint16 // out
	v = (uint16)(c.native.green)
	return v
}

// Blue gets the field inside the struct.
func (c *Color) Blue() uint16 {
	var v uint16 // out
	v = (uint16)(c.native.blue)
	return v
}

// Free frees a color allocated by pango_color_copy().
func (c *Color) Free() {
	var _arg0 *C.PangoColor // out

	_arg0 = (*C.PangoColor)(unsafe.Pointer(c.Native()))

	C.pango_color_free(_arg0)
}

// Parse: fill in the fields of a color from a string specification.
//
// The string can either one of a large set of standard names. (Taken from the
// CSS Color specification (https://www.w3.org/TR/css-color-4/#named-colors), or
// it can be a value in the form `#rgb`, `#rrggbb`, `#rrrgggbbb` or
// `#rrrrggggbbbb`, where `r`, `g` and `b` are hex digits of the red, green, and
// blue components of the color, respectively. (White in the four forms is
// `#fff`, `#ffffff`, `#fffffffff` and `#ffffffffffff`.)
func (c *Color) Parse(spec string) bool {
	var _arg0 *C.PangoColor // out
	var _arg1 *C.char       // out

	_arg0 = (*C.PangoColor)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.char)(C.CString(spec))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean // in

	_cret = C.pango_color_parse(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// ParseWithAlpha: fill in the fields of a color from a string specification.
//
// The string can either one of a large set of standard names. (Taken from the
// CSS Color specification (https://www.w3.org/TR/css-color-4/#named-colors), or
// it can be a hexadecimal value in the form `#rgb`, `#rrggbb`, `#rrrgggbbb` or
// `#rrrrggggbbbb` where `r`, `g` and `b` are hex digits of the red, green, and
// blue components of the color, respectively. (White in the four forms is
// `#fff`, `#ffffff`, `#fffffffff` and `#ffffffffffff`.)
//
// Additionally, parse strings of the form `#rgba`, `#rrggbbaa`,
// `#rrrrggggbbbbaaaa`, if @alpha is not nil, and set @alpha to the value
// specified by the hex digits for `a`. If no alpha component is found in @spec,
// @alpha is set to 0xffff (for a solid color).
func (c *Color) ParseWithAlpha(spec string) (uint16, bool) {
	var _arg0 *C.PangoColor // out
	var _arg2 *C.char       // out

	_arg0 = (*C.PangoColor)(unsafe.Pointer(c.Native()))
	_arg2 = (*C.char)(C.CString(spec))
	defer C.free(unsafe.Pointer(_arg2))

	var _arg1 C.guint16  // in
	var _cret C.gboolean // in

	_cret = C.pango_color_parse_with_alpha(_arg0, _arg2, &_arg1)

	var _alpha uint16 // out
	var _ok bool      // out

	_alpha = (uint16)(_arg1)
	if _cret {
		_ok = true
	}

	return _alpha, _ok
}

// String returns a textual specification of @color.
//
// The string is in the hexadecimal form `#rrrrggggbbbb`, where `r`, `g` and `b`
// are hex digits representing the red, green, and blue components respectively.
func (c *Color) String() string {
	var _arg0 *C.PangoColor // out

	_arg0 = (*C.PangoColor)(unsafe.Pointer(c.Native()))

	var _cret *C.gchar // in

	_cret = C.pango_color_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
