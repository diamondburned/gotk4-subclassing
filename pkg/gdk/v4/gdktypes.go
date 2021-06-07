// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gdk/gdk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_content_formats_get_type()), F: marshalContentFormats},
		{T: externglib.Type(C.gdk_rectangle_get_type()), F: marshalRectangle},
	})
}

// ContentFormats: this section describes the ContentFormats structure that is
// used to advertise and negotiate the format of content passed between
// different widgets, windows or applications using for example the clipboard or
// drag'n'drop.
//
// GDK supports content in 2 forms: #GType and mime type. Using #GTypes is meant
// only for in-process content transfers. Mime types are meant to be used for
// data passing both in-process and out-of-process. The details of how data is
// passed is described in the documentation of the actual implementations.
//
// A ContentFormats describes a set of possible formats content can be exchanged
// in. It is assumed that this set is ordered. #GTypes are more important than
// mime types. Order between different #GTypes or mime types is the order they
// were added in, most important first. Functions that care about order, such as
// gdk_content_formats_union() will describe in their documentation how they
// interpret that order, though in general the order of the first argument is
// considered the primary order of the result, followed by the order of further
// arguments.
//
// For debugging purposes, the function gdk_content_formats_to_string() exists.
// It will print a comma-seperated formats of formats from most important to
// least important.
//
// ContentFormats is an immutable struct. After creation, you cannot change the
// types it represents. Instead, new ContentFormats have to be created. The
// ContentFormatsBuilder structure is meant to help in this endeavor.
type ContentFormats struct {
	native C.GdkContentFormats
}

// WrapContentFormats wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapContentFormats(ptr unsafe.Pointer) *ContentFormats {
	if ptr == nil {
		return nil
	}

	return (*ContentFormats)(ptr)
}

func marshalContentFormats(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapContentFormats(unsafe.Pointer(b)), nil
}

// NewContentFormats constructs a struct ContentFormats.
func NewContentFormats() {
	C.gdk_content_formats_new(arg1, arg2)
}

// NewContentFormatsForGType constructs a struct ContentFormats.
func NewContentFormatsForGType(typ externglib.Type) {
	var arg1 C.GType

	arg1 := C.GType(typ)

	C.gdk_content_formats_new_for_gtype(arg1)
}

// Native returns the underlying C source pointer.
func (c *ContentFormats) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}

// ContainGType checks if a given #GType is part of the given @formats.
func (f *ContentFormats) ContainGType(f *ContentFormats, typ externglib.Type) bool {
	var arg0 *C.GdkContentFormats
	var arg1 C.GType

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))
	arg1 := C.GType(typ)

	var cret C.gboolean
	var ok bool

	cret = C.gdk_content_formats_contain_gtype(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// ContainMIMEType checks if a given mime type is part of the given @formats.
func (f *ContentFormats) ContainMIMEType(f *ContentFormats, mimeType string) bool {
	var arg0 *C.GdkContentFormats
	var arg1 *C.char

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))
	arg1 = (*C.char)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ok bool

	cret = C.gdk_content_formats_contain_mime_type(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// GTypes gets the #GTypes included in @formats. Note that @formats may not
// contain any #GTypes, in particular when they are empty. In that case nil will
// be returned.
func (f *ContentFormats) GTypes(f *ContentFormats) uint {
	var arg0 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	var arg1 C.gsize
	var nGTypes uint

	C.gdk_content_formats_get_gtypes(arg0, &arg1)

	nGTypes = uint(&arg1)

	return nGTypes
}

// MIMETypes gets the mime types included in @formats. Note that @formats may
// not contain any mime types, in particular when they are empty. In that case
// nil will be returned.
func (f *ContentFormats) MIMETypes(f *ContentFormats) uint {
	var arg0 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	var arg1 C.gsize
	var nMIMETypes uint

	C.gdk_content_formats_get_mime_types(arg0, &arg1)

	nMIMETypes = uint(&arg1)

	return nMIMETypes
}

// Match checks if @first and @second have any matching formats.
func (f *ContentFormats) Match(f *ContentFormats, second *ContentFormats) bool {
	var arg0 *C.GdkContentFormats
	var arg1 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))
	arg1 = (*C.GdkContentFormats)(unsafe.Pointer(second.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gdk_content_formats_match(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// MatchGType finds the first #GType from @first that is also contained in
// @second. If no matching #GType is found, G_TYPE_INVALID is returned.
func (f *ContentFormats) MatchGType(f *ContentFormats, second *ContentFormats) {
	var arg0 *C.GdkContentFormats
	var arg1 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))
	arg1 = (*C.GdkContentFormats)(unsafe.Pointer(second.Native()))

	C.gdk_content_formats_match_gtype(arg0, arg1)
}

// MatchMIMEType finds the first mime type from @first that is also contained in
// @second. If no matching mime type is found, nil is returned.
func (f *ContentFormats) MatchMIMEType(f *ContentFormats, second *ContentFormats) {
	var arg0 *C.GdkContentFormats
	var arg1 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))
	arg1 = (*C.GdkContentFormats)(unsafe.Pointer(second.Native()))

	C.gdk_content_formats_match_mime_type(arg0, arg1)
}

// Print prints the given @formats into a string for human consumption. This is
// meant for debugging and logging.
//
// The form of the representation may change at any time and is not guaranteed
// to stay identical.
func (f *ContentFormats) Print(f *ContentFormats, string *glib.String) {
	var arg0 *C.GdkContentFormats
	var arg1 *C.GString

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))
	arg1 = (*C.GString)(unsafe.Pointer(string.Native()))

	C.gdk_content_formats_print(arg0, arg1)
}

// Ref increases the reference count of a ContentFormats by one.
func (f *ContentFormats) Ref(f *ContentFormats) {
	var arg0 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	C.gdk_content_formats_ref(arg0)
}

// String prints the given @formats into a human-readable string. This is a
// small wrapper around gdk_content_formats_print() to help when debugging.
func (f *ContentFormats) String(f *ContentFormats) {
	var arg0 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	C.gdk_content_formats_to_string(arg0)
}

// Union: append all missing types from @second to @first, in the order they had
// in @second.
func (f *ContentFormats) Union(f *ContentFormats, second *ContentFormats) {
	var arg0 *C.GdkContentFormats
	var arg1 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))
	arg1 = (*C.GdkContentFormats)(unsafe.Pointer(second.Native()))

	C.gdk_content_formats_union(arg0, arg1)
}

// UnionDeserializeGTypes: add GTypes for mime types in @formats for which
// deserializers are registered.
func (f *ContentFormats) UnionDeserializeGTypes(f *ContentFormats) {
	var arg0 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	C.gdk_content_formats_union_deserialize_gtypes(arg0)
}

// UnionDeserializeMIMETypes: add mime types for GTypes in @formats for which
// deserializers are registered.
func (f *ContentFormats) UnionDeserializeMIMETypes(f *ContentFormats) {
	var arg0 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	C.gdk_content_formats_union_deserialize_mime_types(arg0)
}

// UnionSerializeGTypes: add GTypes for the mime types in @formats for which
// serializers are registered.
func (f *ContentFormats) UnionSerializeGTypes(f *ContentFormats) {
	var arg0 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	C.gdk_content_formats_union_serialize_gtypes(arg0)
}

// UnionSerializeMIMETypes: add mime types for GTypes in @formats for which
// serializers are registered.
func (f *ContentFormats) UnionSerializeMIMETypes(f *ContentFormats) {
	var arg0 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	C.gdk_content_formats_union_serialize_mime_types(arg0)
}

// Unref decreases the reference count of a ContentFormats by one. If the
// resulting reference count is zero, frees the formats.
func (f *ContentFormats) Unref(f *ContentFormats) {
	var arg0 *C.GdkContentFormats

	arg0 = (*C.GdkContentFormats)(unsafe.Pointer(f.Native()))

	C.gdk_content_formats_unref(arg0)
}

type DrawingContext struct {
	native C.GdkDrawingContext
}

// WrapDrawingContext wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapDrawingContext(ptr unsafe.Pointer) *DrawingContext {
	if ptr == nil {
		return nil
	}

	return (*DrawingContext)(ptr)
}

func marshalDrawingContext(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapDrawingContext(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (d *DrawingContext) Native() unsafe.Pointer {
	return unsafe.Pointer(&d.native)
}

// KeymapKey: a KeymapKey is a hardware key that can be mapped to a keyval.
type KeymapKey struct {
	native C.GdkKeymapKey
}

// WrapKeymapKey wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapKeymapKey(ptr unsafe.Pointer) *KeymapKey {
	if ptr == nil {
		return nil
	}

	return (*KeymapKey)(ptr)
}

func marshalKeymapKey(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapKeymapKey(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (k *KeymapKey) Native() unsafe.Pointer {
	return unsafe.Pointer(&k.native)
}

// Keycode gets the field inside the struct.
func (k *KeymapKey) Keycode() uint {
	var v uint
	v = uint(k.native.keycode)
	return v
}

// Group gets the field inside the struct.
func (k *KeymapKey) Group() int {
	var v int
	v = int(k.native.group)
	return v
}

// Level gets the field inside the struct.
func (k *KeymapKey) Level() int {
	var v int
	v = int(k.native.level)
	return v
}

// Rectangle defines the position and size of a rectangle. It is identical to
// #cairo_rectangle_int_t.
type Rectangle struct {
	native C.GdkRectangle
}

// WrapRectangle wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRectangle(ptr unsafe.Pointer) *Rectangle {
	if ptr == nil {
		return nil
	}

	return (*Rectangle)(ptr)
}

func marshalRectangle(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapRectangle(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *Rectangle) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// X gets the field inside the struct.
func (r *Rectangle) X() int {
	var v int
	v = int(r.native.x)
	return v
}

// Y gets the field inside the struct.
func (r *Rectangle) Y() int {
	var v int
	v = int(r.native.y)
	return v
}

// Width gets the field inside the struct.
func (r *Rectangle) Width() int {
	var v int
	v = int(r.native.width)
	return v
}

// Height gets the field inside the struct.
func (r *Rectangle) Height() int {
	var v int
	v = int(r.native.height)
	return v
}

// ContainsPoint returns UE if @rect contains the point described by @x and @y.
func (r *Rectangle) ContainsPoint(r *Rectangle, x int, y int) bool {
	var arg0 *C.GdkRectangle
	var arg1 C.int
	var arg2 C.int

	arg0 = (*C.GdkRectangle)(unsafe.Pointer(r.Native()))
	arg1 = C.int(x)
	arg2 = C.int(y)

	var cret C.gboolean
	var ok bool

	cret = C.gdk_rectangle_contains_point(arg0, arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// Equal checks if the two given rectangles are equal.
func (r *Rectangle) Equal(r *Rectangle, rect2 *Rectangle) bool {
	var arg0 *C.GdkRectangle
	var arg1 *C.GdkRectangle

	arg0 = (*C.GdkRectangle)(unsafe.Pointer(r.Native()))
	arg1 = (*C.GdkRectangle)(unsafe.Pointer(rect2.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gdk_rectangle_equal(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// Intersect calculates the intersection of two rectangles. It is allowed for
// @dest to be the same as either @src1 or @src2. If the rectangles do not
// intersect, @dest’s width and height is set to 0 and its x and y values are
// undefined. If you are only interested in whether the rectangles intersect,
// but not in the intersecting area itself, pass nil for @dest.
func (s *Rectangle) Intersect(s *Rectangle, src2 *Rectangle) (dest *Rectangle, ok bool) {
	var arg0 *C.GdkRectangle
	var arg1 *C.GdkRectangle

	arg0 = (*C.GdkRectangle)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GdkRectangle)(unsafe.Pointer(src2.Native()))

	var arg2 C.GdkRectangle
	var dest *Rectangle
	var cret C.gboolean
	var ok bool

	cret = C.gdk_rectangle_intersect(arg0, arg1, &arg2)

	dest = WrapRectangle(unsafe.Pointer(&arg2))
	if cret {
		ok = true
	}

	return dest, ok
}

// Union calculates the union of two rectangles. The union of rectangles @src1
// and @src2 is the smallest rectangle which includes both @src1 and @src2
// within it. It is allowed for @dest to be the same as either @src1 or @src2.
//
// Note that this function does not ignore 'empty' rectangles (ie. with zero
// width or height).
func (s *Rectangle) Union(s *Rectangle, src2 *Rectangle) *Rectangle {
	var arg0 *C.GdkRectangle
	var arg1 *C.GdkRectangle

	arg0 = (*C.GdkRectangle)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GdkRectangle)(unsafe.Pointer(src2.Native()))

	var arg2 C.GdkRectangle
	var dest *Rectangle

	C.gdk_rectangle_union(arg0, arg1, &arg2)

	dest = WrapRectangle(unsafe.Pointer(&arg2))

	return dest
}