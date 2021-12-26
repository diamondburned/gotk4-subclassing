// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gdk_axis_use_get_type()), F: marshalAxisUse},
		{T: externglib.Type(C.gdk_gl_error_get_type()), F: marshalGLError},
		{T: externglib.Type(C.gdk_gravity_get_type()), F: marshalGravity},
		{T: externglib.Type(C.gdk_vulkan_error_get_type()), F: marshalVulkanError},
		{T: externglib.Type(C.gdk_axis_flags_get_type()), F: marshalAxisFlags},
		{T: externglib.Type(C.gdk_drag_action_get_type()), F: marshalDragAction},
		{T: externglib.Type(C.gdk_modifier_type_get_type()), F: marshalModifierType},
		{T: externglib.Type(C.gdk_content_formats_get_type()), F: marshalContentFormats},
		{T: externglib.Type(C.gdk_rectangle_get_type()), F: marshalRectangle},
	})
}

// ACTION_ALL defines all possible DND actions.
//
// This can be used in gdk.Drop.Status() messages when any drop can be accepted
// or a more specific drop method is not yet known.
const ACTION_ALL = 7

// CURRENT_TIME represents the current time, and can be used anywhere a time is
// expected.
const CURRENT_TIME = 0

// MODIFIER_MASK: mask covering all entries in GdkModifierType.
const MODIFIER_MASK = 469769999

// AxisUse defines how device axes are interpreted by GTK.
//
// Note that the X and Y axes are not really needed; pointer devices report
// their location via the x/y members of events regardless. Whether X and Y are
// present as axes depends on the GDK backend.
type AxisUse C.gint

const (
	// AxisIgnore axis is ignored.
	AxisIgnore AxisUse = iota
	// AxisX axis is used as the x axis.
	AxisX
	// AxisY axis is used as the y axis.
	AxisY
	// AxisDeltaX axis is used as the scroll x delta.
	AxisDeltaX
	// AxisDeltaY axis is used as the scroll y delta.
	AxisDeltaY
	// AxisPressure axis is used for pressure information.
	AxisPressure
	// AxisXtilt axis is used for x tilt information.
	AxisXtilt
	// AxisYtilt axis is used for y tilt information.
	AxisYtilt
	// AxisWheel axis is used for wheel information.
	AxisWheel
	// AxisDistance axis is used for pen/tablet distance information.
	AxisDistance
	// AxisRotation axis is used for pen rotation information.
	AxisRotation
	// AxisSlider axis is used for pen slider information.
	AxisSlider
	// AxisLast: constant equal to the numerically highest axis value.
	AxisLast
)

func marshalAxisUse(p uintptr) (interface{}, error) {
	return AxisUse(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for AxisUse.
func (a AxisUse) String() string {
	switch a {
	case AxisIgnore:
		return "Ignore"
	case AxisX:
		return "X"
	case AxisY:
		return "Y"
	case AxisDeltaX:
		return "DeltaX"
	case AxisDeltaY:
		return "DeltaY"
	case AxisPressure:
		return "Pressure"
	case AxisXtilt:
		return "Xtilt"
	case AxisYtilt:
		return "Ytilt"
	case AxisWheel:
		return "Wheel"
	case AxisDistance:
		return "Distance"
	case AxisRotation:
		return "Rotation"
	case AxisSlider:
		return "Slider"
	case AxisLast:
		return "Last"
	default:
		return fmt.Sprintf("AxisUse(%d)", a)
	}
}

// GLError: error enumeration for GdkGLContext.
type GLError C.gint

const (
	// GLErrorNotAvailable: openGL support is not available.
	GLErrorNotAvailable GLError = iota
	// GLErrorUnsupportedFormat: requested visual format is not supported.
	GLErrorUnsupportedFormat
	// GLErrorUnsupportedProfile: requested profile is not supported.
	GLErrorUnsupportedProfile
	// GLErrorCompilationFailed: shader compilation failed.
	GLErrorCompilationFailed
	// GLErrorLinkFailed: shader linking failed.
	GLErrorLinkFailed
)

func marshalGLError(p uintptr) (interface{}, error) {
	return GLError(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for GLError.
func (g GLError) String() string {
	switch g {
	case GLErrorNotAvailable:
		return "NotAvailable"
	case GLErrorUnsupportedFormat:
		return "UnsupportedFormat"
	case GLErrorUnsupportedProfile:
		return "UnsupportedProfile"
	case GLErrorCompilationFailed:
		return "CompilationFailed"
	case GLErrorLinkFailed:
		return "LinkFailed"
	default:
		return fmt.Sprintf("GLError(%d)", g)
	}
}

// Gravity defines the reference point of a surface and is used in PopupLayout.
type Gravity C.gint

const (
	// GravityNorthWest: reference point is at the top left corner.
	GravityNorthWest Gravity = 1
	// GravityNorth: reference point is in the middle of the top edge.
	GravityNorth Gravity = 2
	// GravityNorthEast: reference point is at the top right corner.
	GravityNorthEast Gravity = 3
	// GravityWest: reference point is at the middle of the left edge.
	GravityWest Gravity = 4
	// GravityCenter: reference point is at the center of the surface.
	GravityCenter Gravity = 5
	// GravityEast: reference point is at the middle of the right edge.
	GravityEast Gravity = 6
	// GravitySouthWest: reference point is at the lower left corner.
	GravitySouthWest Gravity = 7
	// GravitySouth: reference point is at the middle of the lower edge.
	GravitySouth Gravity = 8
	// GravitySouthEast: reference point is at the lower right corner.
	GravitySouthEast Gravity = 9
	// GravityStatic: reference point is at the top left corner of the surface
	// itself, ignoring window manager decorations.
	GravityStatic Gravity = 10
)

func marshalGravity(p uintptr) (interface{}, error) {
	return Gravity(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for Gravity.
func (g Gravity) String() string {
	switch g {
	case GravityNorthWest:
		return "NorthWest"
	case GravityNorth:
		return "North"
	case GravityNorthEast:
		return "NorthEast"
	case GravityWest:
		return "West"
	case GravityCenter:
		return "Center"
	case GravityEast:
		return "East"
	case GravitySouthWest:
		return "SouthWest"
	case GravitySouth:
		return "South"
	case GravitySouthEast:
		return "SouthEast"
	case GravityStatic:
		return "Static"
	default:
		return fmt.Sprintf("Gravity(%d)", g)
	}
}

// VulkanError: error enumeration for VulkanContext.
type VulkanError C.gint

const (
	// VulkanErrorUnsupported: vulkan is not supported on this backend or has
	// not been compiled in.
	VulkanErrorUnsupported VulkanError = iota
	// VulkanErrorNotAvailable: vulkan support is not available on this Surface.
	VulkanErrorNotAvailable
)

func marshalVulkanError(p uintptr) (interface{}, error) {
	return VulkanError(externglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for VulkanError.
func (v VulkanError) String() string {
	switch v {
	case VulkanErrorUnsupported:
		return "Unsupported"
	case VulkanErrorNotAvailable:
		return "NotAvailable"
	default:
		return fmt.Sprintf("VulkanError(%d)", v)
	}
}

// AxisFlags flags describing the current capabilities of a device/tool.
type AxisFlags C.guint

const (
	// AxisFlagX: x axis is present.
	AxisFlagX AxisFlags = 0b10
	// AxisFlagY: y axis is present.
	AxisFlagY AxisFlags = 0b100
	// AxisFlagDeltaX: scroll X delta axis is present.
	AxisFlagDeltaX AxisFlags = 0b1000
	// AxisFlagDeltaY: scroll Y delta axis is present.
	AxisFlagDeltaY AxisFlags = 0b10000
	// AxisFlagPressure: pressure axis is present.
	AxisFlagPressure AxisFlags = 0b100000
	// AxisFlagXtilt: x tilt axis is present.
	AxisFlagXtilt AxisFlags = 0b1000000
	// AxisFlagYtilt: y tilt axis is present.
	AxisFlagYtilt AxisFlags = 0b10000000
	// AxisFlagWheel: wheel axis is present.
	AxisFlagWheel AxisFlags = 0b100000000
	// AxisFlagDistance: distance axis is present.
	AxisFlagDistance AxisFlags = 0b1000000000
	// AxisFlagRotation z-axis rotation is present.
	AxisFlagRotation AxisFlags = 0b10000000000
	// AxisFlagSlider: slider axis is present.
	AxisFlagSlider AxisFlags = 0b100000000000
)

func marshalAxisFlags(p uintptr) (interface{}, error) {
	return AxisFlags(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for AxisFlags.
func (a AxisFlags) String() string {
	if a == 0 {
		return "AxisFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(157)

	for a != 0 {
		next := a & (a - 1)
		bit := a - next

		switch bit {
		case AxisFlagX:
			builder.WriteString("X|")
		case AxisFlagY:
			builder.WriteString("Y|")
		case AxisFlagDeltaX:
			builder.WriteString("DeltaX|")
		case AxisFlagDeltaY:
			builder.WriteString("DeltaY|")
		case AxisFlagPressure:
			builder.WriteString("Pressure|")
		case AxisFlagXtilt:
			builder.WriteString("Xtilt|")
		case AxisFlagYtilt:
			builder.WriteString("Ytilt|")
		case AxisFlagWheel:
			builder.WriteString("Wheel|")
		case AxisFlagDistance:
			builder.WriteString("Distance|")
		case AxisFlagRotation:
			builder.WriteString("Rotation|")
		case AxisFlagSlider:
			builder.WriteString("Slider|")
		default:
			builder.WriteString(fmt.Sprintf("AxisFlags(0b%b)|", bit))
		}

		a = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if a contains other.
func (a AxisFlags) Has(other AxisFlags) bool {
	return (a & other) == other
}

// DragAction: used in GdkDrop and GdkDrag to indicate the actions that the
// destination can and should do with the dropped data.
type DragAction C.guint

const (
	// ActionCopy: copy the data.
	ActionCopy DragAction = 0b1
	// ActionMove: move the data, i.e. first copy it, then delete it from the
	// source using the DELETE target of the X selection protocol.
	ActionMove DragAction = 0b10
	// ActionLink: add a link to the data. Note that this is only useful if
	// source and destination agree on what it means, and is not supported on
	// all platforms.
	ActionLink DragAction = 0b100
	// ActionAsk: ask the user what to do with the data.
	ActionAsk DragAction = 0b1000
)

func marshalDragAction(p uintptr) (interface{}, error) {
	return DragAction(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for DragAction.
func (d DragAction) String() string {
	if d == 0 {
		return "DragAction(0)"
	}

	var builder strings.Builder
	builder.Grow(42)

	for d != 0 {
		next := d & (d - 1)
		bit := d - next

		switch bit {
		case ActionCopy:
			builder.WriteString("Copy|")
		case ActionMove:
			builder.WriteString("Move|")
		case ActionLink:
			builder.WriteString("Link|")
		case ActionAsk:
			builder.WriteString("Ask|")
		default:
			builder.WriteString(fmt.Sprintf("DragAction(0b%b)|", bit))
		}

		d = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if d contains other.
func (d DragAction) Has(other DragAction) bool {
	return (d & other) == other
}

// ModifierType flags to indicate the state of modifier keys and mouse buttons
// in events.
//
// Typical modifier keys are Shift, Control, Meta, Super, Hyper, Alt, Compose,
// Apple, CapsLock or ShiftLock.
//
// Note that GDK may add internal values to events which include values outside
// of this enumeration. Your code should preserve and ignore them. You can use
// GDK_MODIFIER_MASK to remove all private values.
type ModifierType C.guint

const (
	// ShiftMask: shift key.
	ShiftMask ModifierType = 0b1
	// LockMask: lock key (depending on the modifier mapping of the X server
	// this may either be CapsLock or ShiftLock).
	LockMask ModifierType = 0b10
	// ControlMask: control key.
	ControlMask ModifierType = 0b100
	// AltMask: fourth modifier key (it depends on the modifier mapping of the X
	// server which key is interpreted as this modifier, but normally it is the
	// Alt key).
	AltMask ModifierType = 0b1000
	// Button1Mask: first mouse button.
	Button1Mask ModifierType = 0b100000000
	// Button2Mask: second mouse button.
	Button2Mask ModifierType = 0b1000000000
	// Button3Mask: third mouse button.
	Button3Mask ModifierType = 0b10000000000
	// Button4Mask: fourth mouse button.
	Button4Mask ModifierType = 0b100000000000
	// Button5Mask: fifth mouse button.
	Button5Mask ModifierType = 0b1000000000000
	// SuperMask: super modifier.
	SuperMask ModifierType = 0b100000000000000000000000000
	// HyperMask: hyper modifier.
	HyperMask ModifierType = 0b1000000000000000000000000000
	// MetaMask: meta modifier.
	MetaMask ModifierType = 0b10000000000000000000000000000
)

func marshalModifierType(p uintptr) (interface{}, error) {
	return ModifierType(externglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for ModifierType.
func (m ModifierType) String() string {
	if m == 0 {
		return "ModifierType(0)"
	}

	var builder strings.Builder
	builder.Grow(127)

	for m != 0 {
		next := m & (m - 1)
		bit := m - next

		switch bit {
		case ShiftMask:
			builder.WriteString("ShiftMask|")
		case LockMask:
			builder.WriteString("LockMask|")
		case ControlMask:
			builder.WriteString("ControlMask|")
		case AltMask:
			builder.WriteString("AltMask|")
		case Button1Mask:
			builder.WriteString("Button1Mask|")
		case Button2Mask:
			builder.WriteString("Button2Mask|")
		case Button3Mask:
			builder.WriteString("Button3Mask|")
		case Button4Mask:
			builder.WriteString("Button4Mask|")
		case Button5Mask:
			builder.WriteString("Button5Mask|")
		case SuperMask:
			builder.WriteString("SuperMask|")
		case HyperMask:
			builder.WriteString("HyperMask|")
		case MetaMask:
			builder.WriteString("MetaMask|")
		default:
			builder.WriteString(fmt.Sprintf("ModifierType(0b%b)|", bit))
		}

		m = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if m contains other.
func (m ModifierType) Has(other ModifierType) bool {
	return (m & other) == other
}

// ContentFormats: GdkContentFormats structure is used to advertise and
// negotiate the format of content.
//
// You will encounter GdkContentFormats when interacting with objects
// controlling operations that pass data between different widgets, window or
// application, like gdk.Drag, gdk.Drop, gdk.Clipboard or gdk.ContentProvider.
//
// GDK supports content in 2 forms: GType and mime type. Using GTypes is meant
// only for in-process content transfers. Mime types are meant to be used for
// data passing both in-process and out-of-process. The details of how data is
// passed is described in the documentation of the actual implementations. To
// transform between the two forms, gdk.ContentSerializer and
// gdk.ContentDeserializer are used.
//
// A GdkContentFormats describes a set of possible formats content can be
// exchanged in. It is assumed that this set is ordered. GTypes are more
// important than mime types. Order between different GTypes or mime types is
// the order they were added in, most important first. Functions that care about
// order, such as gdk.ContentFormats.Union(), will describe in their
// documentation how they interpret that order, though in general the order of
// the first argument is considered the primary order of the result, followed by
// the order of further arguments.
//
// For debugging purposes, the function gdk.ContentFormats.ToString() exists. It
// will print a comma-separated list of formats from most important to least
// important.
//
// GdkContentFormats is an immutable struct. After creation, you cannot change
// the types it represents. Instead, new GdkContentFormats have to be created.
// The gdk.ContentFormatsBuilder` structure is meant to help in this endeavor.
//
// An instance of this type is always passed by reference.
type ContentFormats struct {
	*contentFormats
}

// contentFormats is the struct that's finalized.
type contentFormats struct {
	native *C.GdkContentFormats
}

func marshalContentFormats(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &ContentFormats{&contentFormats{(*C.GdkContentFormats)(b)}}, nil
}

// NewContentFormats constructs a struct ContentFormats.
func NewContentFormats(mimeTypes []string) *ContentFormats {
	var _arg1 **C.char // out
	var _arg2 C.guint
	var _cret *C.GdkContentFormats // in

	_arg2 = (C.guint)(len(mimeTypes))
	_arg1 = (**C.char)(C.calloc(C.size_t(len(mimeTypes)), C.size_t(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg1))
	{
		out := unsafe.Slice((**C.char)(_arg1), len(mimeTypes))
		for i := range mimeTypes {
			out[i] = (*C.char)(unsafe.Pointer(C.CString(mimeTypes[i])))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	_cret = C.gdk_content_formats_new(_arg1, _arg2)
	runtime.KeepAlive(mimeTypes)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_contentFormats)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_content_formats_unref((*C.GdkContentFormats)(intern.C))
		},
	)

	return _contentFormats
}

// NewContentFormatsForGType constructs a struct ContentFormats.
func NewContentFormatsForGType(typ externglib.Type) *ContentFormats {
	var _arg1 C.GType              // out
	var _cret *C.GdkContentFormats // in

	_arg1 = C.GType(typ)

	_cret = C.gdk_content_formats_new_for_gtype(_arg1)
	runtime.KeepAlive(typ)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_contentFormats)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_content_formats_unref((*C.GdkContentFormats)(intern.C))
		},
	)

	return _contentFormats
}

// ContainGType checks if a given GType is part of the given formats.
//
// The function takes the following parameters:
//
//    - typ: GType to search for.
//
// The function returns the following values:
//
//    - ok: TRUE if the #GType was found.
//
func (formats *ContentFormats) ContainGType(typ externglib.Type) bool {
	var _arg0 *C.GdkContentFormats // out
	var _arg1 C.GType              // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))
	_arg1 = C.GType(typ)

	_cret = C.gdk_content_formats_contain_gtype(_arg0, _arg1)
	runtime.KeepAlive(formats)
	runtime.KeepAlive(typ)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ContainMIMEType checks if a given mime type is part of the given formats.
//
// The function takes the following parameters:
//
//    - mimeType: mime type to search for.
//
// The function returns the following values:
//
//    - ok: TRUE if the mime_type was found.
//
func (formats *ContentFormats) ContainMIMEType(mimeType string) bool {
	var _arg0 *C.GdkContentFormats // out
	var _arg1 *C.char              // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(mimeType)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gdk_content_formats_contain_mime_type(_arg0, _arg1)
	runtime.KeepAlive(formats)
	runtime.KeepAlive(mimeType)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// GTypes gets the GTypes included in formats.
//
// Note that formats may not contain any #GTypes, in particular when they are
// empty. In that case NULL will be returned.
//
// The function returns the following values:
//
//    - gTypes (optional): G_TYPE_INVALID-terminated array of types included in
//      formats or NULL if none.
//
func (formats *ContentFormats) GTypes() []externglib.Type {
	var _arg0 *C.GdkContentFormats // out
	var _cret *C.GType             // in
	var _arg1 C.gsize              // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))

	_cret = C.gdk_content_formats_get_gtypes(_arg0, &_arg1)
	runtime.KeepAlive(formats)

	var _gTypes []externglib.Type // out

	if _cret != nil {
		{
			src := unsafe.Slice(_cret, _arg1)
			_gTypes = make([]externglib.Type, _arg1)
			for i := 0; i < int(_arg1); i++ {
				_gTypes[i] = externglib.Type(src[i])
			}
		}
	}

	return _gTypes
}

// MIMETypes gets the mime types included in formats.
//
// Note that formats may not contain any mime types, in particular when they are
// empty. In that case NULL will be returned.
//
// The function returns the following values:
//
//    - utf8s (optional): NULL-terminated array of interned strings of mime types
//      included in formats or NULL if none.
//
func (formats *ContentFormats) MIMETypes() []string {
	var _arg0 *C.GdkContentFormats // out
	var _cret **C.char             // in
	var _arg1 C.gsize              // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))

	_cret = C.gdk_content_formats_get_mime_types(_arg0, &_arg1)
	runtime.KeepAlive(formats)

	var _utf8s []string // out

	if _cret != nil {
		{
			src := unsafe.Slice(_cret, _arg1)
			_utf8s = make([]string, _arg1)
			for i := 0; i < int(_arg1); i++ {
				_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			}
		}
	}

	return _utf8s
}

// Match checks if first and second have any matching formats.
//
// The function takes the following parameters:
//
//    - second: GdkContentFormats to intersect with.
//
// The function returns the following values:
//
//    - ok: TRUE if a matching format was found.
//
func (first *ContentFormats) Match(second *ContentFormats) bool {
	var _arg0 *C.GdkContentFormats // out
	var _arg1 *C.GdkContentFormats // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(first)))
	_arg1 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(second)))

	_cret = C.gdk_content_formats_match(_arg0, _arg1)
	runtime.KeepAlive(first)
	runtime.KeepAlive(second)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MatchGType finds the first GType from first that is also contained in second.
//
// If no matching GType is found, G_TYPE_INVALID is returned.
//
// The function takes the following parameters:
//
//    - second: GdkContentFormats to intersect with.
//
// The function returns the following values:
//
//    - gType: first common GType or G_TYPE_INVALID if none.
//
func (first *ContentFormats) MatchGType(second *ContentFormats) externglib.Type {
	var _arg0 *C.GdkContentFormats // out
	var _arg1 *C.GdkContentFormats // out
	var _cret C.GType              // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(first)))
	_arg1 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(second)))

	_cret = C.gdk_content_formats_match_gtype(_arg0, _arg1)
	runtime.KeepAlive(first)
	runtime.KeepAlive(second)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// MatchMIMEType finds the first mime type from first that is also contained in
// second.
//
// If no matching mime type is found, NULL is returned.
//
// The function takes the following parameters:
//
//    - second: GdkContentFormats to intersect with.
//
// The function returns the following values:
//
//    - utf8 (optional): first common mime type or NULL if none.
//
func (first *ContentFormats) MatchMIMEType(second *ContentFormats) string {
	var _arg0 *C.GdkContentFormats // out
	var _arg1 *C.GdkContentFormats // out
	var _cret *C.char              // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(first)))
	_arg1 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(second)))

	_cret = C.gdk_content_formats_match_mime_type(_arg0, _arg1)
	runtime.KeepAlive(first)
	runtime.KeepAlive(second)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// String prints the given formats into a human-readable string.
//
// This is a small wrapper around gdk.ContentFormats.Print() to help when
// debugging.
//
// The function returns the following values:
//
//    - utf8: new string.
//
func (formats *ContentFormats) String() string {
	var _arg0 *C.GdkContentFormats // out
	var _cret *C.char              // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))

	_cret = C.gdk_content_formats_to_string(_arg0)
	runtime.KeepAlive(formats)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Union: append all missing types from second to first, in the order they had
// in second.
//
// The function takes the following parameters:
//
//    - second: GdkContentFormats to merge from.
//
// The function returns the following values:
//
//    - contentFormats: new GdkContentFormats.
//
func (first *ContentFormats) Union(second *ContentFormats) *ContentFormats {
	var _arg0 *C.GdkContentFormats // out
	var _arg1 *C.GdkContentFormats // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(first)))
	_arg1 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(second)))

	_cret = C.gdk_content_formats_union(_arg0, _arg1)
	runtime.KeepAlive(first)
	runtime.KeepAlive(second)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_contentFormats)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_content_formats_unref((*C.GdkContentFormats)(intern.C))
		},
	)

	return _contentFormats
}

// UnionDeserializeGTypes: add GTypes for mime types in formats for which
// deserializers are registered.
//
// The function returns the following values:
//
//    - contentFormats: new GdkContentFormats.
//
func (formats *ContentFormats) UnionDeserializeGTypes() *ContentFormats {
	var _arg0 *C.GdkContentFormats // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))

	_cret = C.gdk_content_formats_union_deserialize_gtypes(_arg0)
	runtime.KeepAlive(formats)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_contentFormats)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_content_formats_unref((*C.GdkContentFormats)(intern.C))
		},
	)

	return _contentFormats
}

// UnionDeserializeMIMETypes: add mime types for GTypes in formats for which
// deserializers are registered.
//
// The function returns the following values:
//
//    - contentFormats: new GdkContentFormats.
//
func (formats *ContentFormats) UnionDeserializeMIMETypes() *ContentFormats {
	var _arg0 *C.GdkContentFormats // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))

	_cret = C.gdk_content_formats_union_deserialize_mime_types(_arg0)
	runtime.KeepAlive(formats)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_contentFormats)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_content_formats_unref((*C.GdkContentFormats)(intern.C))
		},
	)

	return _contentFormats
}

// UnionSerializeGTypes: add GTypes for the mime types in formats for which
// serializers are registered.
//
// The function returns the following values:
//
//    - contentFormats: new GdkContentFormats.
//
func (formats *ContentFormats) UnionSerializeGTypes() *ContentFormats {
	var _arg0 *C.GdkContentFormats // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))

	_cret = C.gdk_content_formats_union_serialize_gtypes(_arg0)
	runtime.KeepAlive(formats)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_contentFormats)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_content_formats_unref((*C.GdkContentFormats)(intern.C))
		},
	)

	return _contentFormats
}

// UnionSerializeMIMETypes: add mime types for GTypes in formats for which
// serializers are registered.
//
// The function returns the following values:
//
//    - contentFormats: new GdkContentFormats.
//
func (formats *ContentFormats) UnionSerializeMIMETypes() *ContentFormats {
	var _arg0 *C.GdkContentFormats // out
	var _cret *C.GdkContentFormats // in

	_arg0 = (*C.GdkContentFormats)(gextras.StructNative(unsafe.Pointer(formats)))

	_cret = C.gdk_content_formats_union_serialize_mime_types(_arg0)
	runtime.KeepAlive(formats)

	var _contentFormats *ContentFormats // out

	_contentFormats = (*ContentFormats)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_contentFormats)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gdk_content_formats_unref((*C.GdkContentFormats)(intern.C))
		},
	)

	return _contentFormats
}

// KeymapKey: GdkKeymapKey is a hardware key that can be mapped to a keyval.
//
// An instance of this type is always passed by reference.
type KeymapKey struct {
	*keymapKey
}

// keymapKey is the struct that's finalized.
type keymapKey struct {
	native *C.GdkKeymapKey
}

// NewKeymapKey creates a new KeymapKey instance from the given
// fields.
func NewKeymapKey(keycode uint, group, level int) KeymapKey {
	var f0 C.guint // out
	f0 = C.guint(keycode)
	var f1 C.int // out
	f1 = C.int(group)
	var f2 C.int // out
	f2 = C.int(level)

	v := C.GdkKeymapKey{
		keycode: f0,
		group:   f1,
		level:   f2,
	}

	return *(*KeymapKey)(gextras.NewStructNative(unsafe.Pointer(&v)))
}

// Keycode: hardware keycode. This is an identifying number for a physical key.
func (k *KeymapKey) Keycode() uint {
	var v uint // out
	v = uint(k.native.keycode)
	return v
}

// Group indicates movement in a horizontal direction. Usually groups are used
// for two different languages. In group 0, a key might have two English
// characters, and in group 1 it might have two Hebrew characters. The Hebrew
// characters will be printed on the key next to the English characters.
func (k *KeymapKey) Group() int {
	var v int // out
	v = int(k.native.group)
	return v
}

// Level indicates which symbol on the key will be used, in a vertical
// direction. So on a standard US keyboard, the key with the number “1” on it
// also has the exclamation point ("!") character on it. The level indicates
// whether to use the “1” or the “!” symbol. The letter keys are considered to
// have a lowercase letter at level 0, and an uppercase letter at level 1,
// though only the uppercase letter is printed.
func (k *KeymapKey) Level() int {
	var v int // out
	v = int(k.native.level)
	return v
}

// Rectangle: GdkRectangle data type for representing rectangles.
//
// GdkRectangle is identical to cairo_rectangle_t. Together with Cairo’s
// cairo_region_t data type, these are the central types for representing sets
// of pixels.
//
// The intersection of two rectangles can be computed with
// gdk.Rectangle.Intersect(); to find the union of two rectangles use
// gdk.Rectangle.Union().
//
// The cairo_region_t type provided by Cairo is usually used for managing
// non-rectangular clipping of graphical operations.
//
// The Graphene library has a number of other data types for regions and volumes
// in 2D and 3D.
//
// An instance of this type is always passed by reference.
type Rectangle struct {
	*rectangle
}

// rectangle is the struct that's finalized.
type rectangle struct {
	native *C.GdkRectangle
}

func marshalRectangle(p uintptr) (interface{}, error) {
	b := externglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Rectangle{&rectangle{(*C.GdkRectangle)(b)}}, nil
}

// NewRectangle creates a new Rectangle instance from the given
// fields.
func NewRectangle(x, y, width, height int) Rectangle {
	var f0 C.int // out
	f0 = C.int(x)
	var f1 C.int // out
	f1 = C.int(y)
	var f2 C.int // out
	f2 = C.int(width)
	var f3 C.int // out
	f3 = C.int(height)

	v := C.GdkRectangle{
		x:      f0,
		y:      f1,
		width:  f2,
		height: f3,
	}

	return *(*Rectangle)(gextras.NewStructNative(unsafe.Pointer(&v)))
}

// X: x coordinate of the top left corner.
func (r *Rectangle) X() int {
	var v int // out
	v = int(r.native.x)
	return v
}

// Y: y coordinate of the top left corner.
func (r *Rectangle) Y() int {
	var v int // out
	v = int(r.native.y)
	return v
}

// Width: width of the rectangle.
func (r *Rectangle) Width() int {
	var v int // out
	v = int(r.native.width)
	return v
}

// Height: height of the rectangle.
func (r *Rectangle) Height() int {
	var v int // out
	v = int(r.native.height)
	return v
}

// ContainsPoint returns UE if rect contains the point described by x and y.
//
// The function takes the following parameters:
//
//    - x: x coordinate.
//    - y: y coordinate.
//
// The function returns the following values:
//
//    - ok if rect contains the point.
//
func (rect *Rectangle) ContainsPoint(x int, y int) bool {
	var _arg0 *C.GdkRectangle // out
	var _arg1 C.int           // out
	var _arg2 C.int           // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(rect)))
	_arg1 = C.int(x)
	_arg2 = C.int(y)

	_cret = C.gdk_rectangle_contains_point(_arg0, _arg1, _arg2)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Equal checks if the two given rectangles are equal.
//
// The function takes the following parameters:
//
//    - rect2: GdkRectangle.
//
// The function returns the following values:
//
//    - ok: TRUE if the rectangles are equal.
//
func (rect1 *Rectangle) Equal(rect2 *Rectangle) bool {
	var _arg0 *C.GdkRectangle // out
	var _arg1 *C.GdkRectangle // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(rect1)))
	_arg1 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(rect2)))

	_cret = C.gdk_rectangle_equal(_arg0, _arg1)
	runtime.KeepAlive(rect1)
	runtime.KeepAlive(rect2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Intersect calculates the intersection of two rectangles.
//
// It is allowed for dest to be the same as either src1 or src2. If the
// rectangles do not intersect, dest’s width and height is set to 0 and its x
// and y values are undefined. If you are only interested in whether the
// rectangles intersect, but not in the intersecting area itself, pass NULL for
// dest.
//
// The function takes the following parameters:
//
//    - src2: GdkRectangle.
//
// The function returns the following values:
//
//    - dest (optional): return location for the intersection of src1 and src2,
//      or NULL.
//    - ok: TRUE if the rectangles intersect.
//
func (src1 *Rectangle) Intersect(src2 *Rectangle) (*Rectangle, bool) {
	var _arg0 *C.GdkRectangle // out
	var _arg1 *C.GdkRectangle // out
	var _arg2 C.GdkRectangle  // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(src1)))
	_arg1 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(src2)))

	_cret = C.gdk_rectangle_intersect(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(src1)
	runtime.KeepAlive(src2)

	var _dest *Rectangle // out
	var _ok bool         // out

	_dest = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))
	if _cret != 0 {
		_ok = true
	}

	return _dest, _ok
}

// Union calculates the union of two rectangles.
//
// The union of rectangles src1 and src2 is the smallest rectangle which
// includes both src1 and src2 within it. It is allowed for dest to be the same
// as either src1 or src2.
//
// Note that this function does not ignore 'empty' rectangles (ie. with zero
// width or height).
//
// The function takes the following parameters:
//
//    - src2: GdkRectangle.
//
// The function returns the following values:
//
//    - dest: return location for the union of src1 and src2.
//
func (src1 *Rectangle) Union(src2 *Rectangle) *Rectangle {
	var _arg0 *C.GdkRectangle // out
	var _arg1 *C.GdkRectangle // out
	var _arg2 C.GdkRectangle  // in

	_arg0 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(src1)))
	_arg1 = (*C.GdkRectangle)(gextras.StructNative(unsafe.Pointer(src2)))

	C.gdk_rectangle_union(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(src1)
	runtime.KeepAlive(src2)

	var _dest *Rectangle // out

	_dest = (*Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg2))))

	return _dest
}
