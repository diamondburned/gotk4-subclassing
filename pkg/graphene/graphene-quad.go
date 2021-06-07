// Code generated by girgen. DO NOT EDIT.

package graphene

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: graphene-gobject-1.0 graphene-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <graphene-gobject.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.graphene_quad_get_type()), F: marshalQuad},
	})
}

// Quad: a 4 vertex quadrilateral, as represented by four #graphene_point_t.
//
// The contents of a #graphene_quad_t are private and should never be accessed
// directly.
type Quad struct {
	native C.graphene_quad_t
}

// WrapQuad wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapQuad(ptr unsafe.Pointer) *Quad {
	if ptr == nil {
		return nil
	}

	return (*Quad)(ptr)
}

func marshalQuad(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return WrapQuad(unsafe.Pointer(b)), nil
}

// NewQuadAlloc constructs a struct Quad.
func NewQuadAlloc() {
	C.graphene_quad_alloc()
}

// Native returns the underlying C source pointer.
func (q *Quad) Native() unsafe.Pointer {
	return unsafe.Pointer(&q.native)
}

// Bounds computes the bounding rectangle of @q and places it into @r.
func (q *Quad) Bounds(q *Quad) *Rect {
	var arg0 *C.graphene_quad_t

	arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q.Native()))

	var arg1 C.graphene_rect_t
	var r *Rect

	C.graphene_quad_bounds(arg0, &arg1)

	r = WrapRect(unsafe.Pointer(&arg1))

	return r
}

// Contains checks if the given #graphene_quad_t contains the given
// #graphene_point_t.
func (q *Quad) Contains(q *Quad, p *Point) bool {
	var arg0 *C.graphene_quad_t
	var arg1 *C.graphene_point_t

	arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q.Native()))
	arg1 = (*C.graphene_point_t)(unsafe.Pointer(p.Native()))

	var cret C._Bool
	var ok bool

	cret = C.graphene_quad_contains(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// Free frees the resources allocated by graphene_quad_alloc()
func (q *Quad) Free(q *Quad) {
	var arg0 *C.graphene_quad_t

	arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q.Native()))

	C.graphene_quad_free(arg0)
}

// Point retrieves the point of a #graphene_quad_t at the given index.
func (q *Quad) Point(q *Quad, index_ uint) {
	var arg0 *C.graphene_quad_t
	var arg1 C.uint

	arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q.Native()))
	arg1 = C.uint(index_)

	C.graphene_quad_get_point(arg0, arg1)
}

// Init initializes a #graphene_quad_t with the given points.
func (q *Quad) Init(q *Quad, p1 *Point, p2 *Point, p3 *Point, p4 *Point) {
	var arg0 *C.graphene_quad_t
	var arg1 *C.graphene_point_t
	var arg2 *C.graphene_point_t
	var arg3 *C.graphene_point_t
	var arg4 *C.graphene_point_t

	arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q.Native()))
	arg1 = (*C.graphene_point_t)(unsafe.Pointer(p1.Native()))
	arg2 = (*C.graphene_point_t)(unsafe.Pointer(p2.Native()))
	arg3 = (*C.graphene_point_t)(unsafe.Pointer(p3.Native()))
	arg4 = (*C.graphene_point_t)(unsafe.Pointer(p4.Native()))

	C.graphene_quad_init(arg0, arg1, arg2, arg3, arg4)
}

// InitFromPoints initializes a #graphene_quad_t using an array of points.
func (q *Quad) InitFromPoints(q *Quad, points [4]Point) {
	var arg0 *C.graphene_quad_t
	var arg1 *C.graphene_point_t

	arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q.Native()))
	{
		out := (*[4]C.graphene_point_t)(unsafe.Pointer(arg1))
		for i := 0; i < 4; i++ {
			out[i] = (C.graphene_point_t)(unsafe.Pointer(points[i].Native()))
		}
	}

	C.graphene_quad_init_from_points(arg0, arg1)
}

// InitFromRect initializes a #graphene_quad_t using the four corners of the
// given #graphene_rect_t.
func (q *Quad) InitFromRect(q *Quad, r *Rect) {
	var arg0 *C.graphene_quad_t
	var arg1 *C.graphene_rect_t

	arg0 = (*C.graphene_quad_t)(unsafe.Pointer(q.Native()))
	arg1 = (*C.graphene_rect_t)(unsafe.Pointer(r.Native()))

	C.graphene_quad_init_from_rect(arg0, arg1)
}
