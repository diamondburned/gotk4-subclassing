// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/graphene"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gsk/gsk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gsk_render_node_get_type()), F: marshalRenderNode},
	})
}

// ParseErrorFunc: type of callback that is called when an error occurs during
// node deserialization.
type ParseErrorFunc func(start *ParseLocation, end *ParseLocation, err error, userData interface{})

//export gotk4_ParseErrorFunc
func gotk4_ParseErrorFunc(arg0 *C.GskParseLocation, arg1 *C.GskParseLocation, arg2 *C.GError, arg3 C.gpointer) {
	v := box.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var start *ParseLocation // out
	var end *ParseLocation   // out
	var err error            // out
	var userData interface{} // out

	start = (*ParseLocation)(unsafe.Pointer(arg0))
	end = (*ParseLocation)(unsafe.Pointer(arg1))
	err = gerror.Take(unsafe.Pointer(arg2))
	userData = box.Get(uintptr(arg3))

	fn := v.(ParseErrorFunc)
	fn(start, end, err, userData)
}

// RenderNode: `GskRenderNode` is the basic block in a scene graph to be
// rendered using `GskRenderer`.
//
// Each node has a parent, except the top-level node; each node may have
// children nodes.
//
// Each node has an associated drawing surface, which has the size of the
// rectangle set when creating it.
//
// Render nodes are meant to be transient; once they have been associated to a
// [class@Gsk.Renderer] it's safe to release any reference you have on them. All
// [class@Gsk.RenderNode]s are immutable, you can only specify their properties
// during construction.
type RenderNode interface {
	gextras.Objector

	// Draw the contents of @node to the given cairo context.
	//
	// Typically, you'll use this function to implement fallback rendering of
	// `GskRenderNode`s on an intermediate Cairo context, instead of using the
	// drawing context associated to a `GdkSurface`'s rendering buffer.
	//
	// For advanced nodes that cannot be supported using Cairo, in particular
	// for nodes doing 3D operations, this function may fail.
	Draw(cr *cairo.Context)
	// Bounds retrieves the boundaries of the @node.
	//
	// The node will not draw outside of its boundaries.
	Bounds() graphene.Rect
	// NodeType returns the type of the @node.
	NodeType() RenderNodeType
	// Ref acquires a reference on the given `GskRenderNode`.
	ref() *RenderNodeClass
	// Unref releases a reference on the given `GskRenderNode`.
	//
	// If the reference was the last, the resources associated to the @node are
	// freed.
	unref()
	// WriteToFile: this function is equivalent to calling
	// gsk_render_node_serialize() followed by g_file_set_contents().
	//
	// See those two functions for details on the arguments.
	//
	// It is mostly intended for use inside a debugger to quickly dump a render
	// node to a file for later inspection.
	WriteToFile(filename string) error
}

// RenderNodeClass implements the RenderNode interface.
type RenderNodeClass struct {
	*externglib.Object
}

var _ RenderNode = (*RenderNodeClass)(nil)

func wrapRenderNode(obj *externglib.Object) RenderNode {
	return &RenderNodeClass{
		Object: obj,
	}
}

func marshalRenderNode(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRenderNode(obj), nil
}

// Draw the contents of @node to the given cairo context.
//
// Typically, you'll use this function to implement fallback rendering of
// `GskRenderNode`s on an intermediate Cairo context, instead of using the
// drawing context associated to a `GdkSurface`'s rendering buffer.
//
// For advanced nodes that cannot be supported using Cairo, in particular for
// nodes doing 3D operations, this function may fail.
func (n *RenderNodeClass) Draw(cr *cairo.Context) {
	var _arg0 *C.GskRenderNode // out
	var _arg1 *C.cairo_t       // out

	_arg0 = (*C.GskRenderNode)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr))

	C.gsk_render_node_draw(_arg0, _arg1)
}

// Bounds retrieves the boundaries of the @node.
//
// The node will not draw outside of its boundaries.
func (n *RenderNodeClass) Bounds() graphene.Rect {
	var _arg0 *C.GskRenderNode  // out
	var _arg1 C.graphene_rect_t // in

	_arg0 = (*C.GskRenderNode)(unsafe.Pointer(n.Native()))

	C.gsk_render_node_get_bounds(_arg0, &_arg1)

	var _bounds graphene.Rect // out

	_bounds = *(*graphene.Rect)(unsafe.Pointer((&_arg1)))

	return _bounds
}

// NodeType returns the type of the @node.
func (n *RenderNodeClass) NodeType() RenderNodeType {
	var _arg0 *C.GskRenderNode    // out
	var _cret C.GskRenderNodeType // in

	_arg0 = (*C.GskRenderNode)(unsafe.Pointer(n.Native()))

	_cret = C.gsk_render_node_get_node_type(_arg0)

	var _renderNodeType RenderNodeType // out

	_renderNodeType = (RenderNodeType)(_cret)

	return _renderNodeType
}

// Ref acquires a reference on the given `GskRenderNode`.
func (n *RenderNodeClass) ref() *RenderNodeClass {
	var _arg0 *C.GskRenderNode // out
	var _cret *C.GskRenderNode // in

	_arg0 = (*C.GskRenderNode)(unsafe.Pointer(n.Native()))

	_cret = C.gsk_render_node_ref(_arg0)

	var _renderNode *RenderNodeClass // out

	_renderNode = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*RenderNodeClass)

	return _renderNode
}

// Unref releases a reference on the given `GskRenderNode`.
//
// If the reference was the last, the resources associated to the @node are
// freed.
func (n *RenderNodeClass) unref() {
	var _arg0 *C.GskRenderNode // out

	_arg0 = (*C.GskRenderNode)(unsafe.Pointer(n.Native()))

	C.gsk_render_node_unref(_arg0)
}

// WriteToFile: this function is equivalent to calling
// gsk_render_node_serialize() followed by g_file_set_contents().
//
// See those two functions for details on the arguments.
//
// It is mostly intended for use inside a debugger to quickly dump a render node
// to a file for later inspection.
func (n *RenderNodeClass) WriteToFile(filename string) error {
	var _arg0 *C.GskRenderNode // out
	var _arg1 *C.char          // out
	var _cerr *C.GError        // in

	_arg0 = (*C.GskRenderNode)(unsafe.Pointer(n.Native()))
	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	C.gsk_render_node_write_to_file(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// ColorStop: color stop in a gradient node.
type ColorStop struct {
	native C.GskColorStop
}

// WrapColorStop wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapColorStop(ptr unsafe.Pointer) *ColorStop {
	return (*ColorStop)(ptr)
}

// Native returns the underlying C source pointer.
func (c *ColorStop) Native() unsafe.Pointer {
	return unsafe.Pointer(&c.native)
}

// Offset: the offset of the color stop
func (c *ColorStop) Offset() float32 {
	var v float32 // out
	v = float32(c.native.offset)
	return v
}

// Color: the color at the given offset
func (c *ColorStop) Color() gdk.RGBA {
	var v gdk.RGBA // out
	v = *(*gdk.RGBA)(unsafe.Pointer((&c.native.color)))
	return v
}

// ParseLocation: location in a parse buffer.
type ParseLocation struct {
	native C.GskParseLocation
}

// WrapParseLocation wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapParseLocation(ptr unsafe.Pointer) *ParseLocation {
	return (*ParseLocation)(ptr)
}

// Native returns the underlying C source pointer.
func (p *ParseLocation) Native() unsafe.Pointer {
	return unsafe.Pointer(&p.native)
}

// Bytes: the offset of the location in the parse buffer, as bytes
func (p *ParseLocation) Bytes() uint {
	var v uint // out
	v = uint(p.native.bytes)
	return v
}

// Chars: the offset of the location in the parse buffer, as characters
func (p *ParseLocation) Chars() uint {
	var v uint // out
	v = uint(p.native.chars)
	return v
}

// Lines: the line of the location in the parse buffer
func (p *ParseLocation) Lines() uint {
	var v uint // out
	v = uint(p.native.lines)
	return v
}

// LineBytes: the position in the line, as bytes
func (p *ParseLocation) LineBytes() uint {
	var v uint // out
	v = uint(p.native.line_bytes)
	return v
}

// LineChars: the position in the line, as characters
func (p *ParseLocation) LineChars() uint {
	var v uint // out
	v = uint(p.native.line_chars)
	return v
}

// Shadow: the shadow parameters in a shadow node.
type Shadow struct {
	native C.GskShadow
}

// WrapShadow wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapShadow(ptr unsafe.Pointer) *Shadow {
	return (*Shadow)(ptr)
}

// Native returns the underlying C source pointer.
func (s *Shadow) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// Color: the color of the shadow
func (s *Shadow) Color() gdk.RGBA {
	var v gdk.RGBA // out
	v = *(*gdk.RGBA)(unsafe.Pointer((&s.native.color)))
	return v
}

// Dx: the horizontal offset of the shadow
func (s *Shadow) Dx() float32 {
	var v float32 // out
	v = float32(s.native.dx)
	return v
}

// Dy: the vertical offset of the shadow
func (s *Shadow) Dy() float32 {
	var v float32 // out
	v = float32(s.native.dy)
	return v
}

// Radius: the radius of the shadow
func (s *Shadow) Radius() float32 {
	var v float32 // out
	v = float32(s.native.radius)
	return v
}
