// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_constraint_layout_get_type()), F: marshalConstraintLayout},
		{T: externglib.Type(C.gtk_constraint_layout_child_get_type()), F: marshalConstraintLayoutChild},
	})
}

// ConstraintLayout: gtkConstraintLayout is a layout manager that uses relations
// between widget attributes, expressed via Constraint instances, to measure and
// allocate widgets.
//
//
// How do constraints work
//
// Constraints are objects defining the relationship between attributes of a
// widget; you can read the description of the Constraint class to have a more
// in depth definition.
//
// By taking multiple constraints and applying them to the children of a widget
// using ConstraintLayout, it's possible to describe complex layout policies;
// each constraint applied to a child or to the parent widgets contributes to
// the full description of the layout, in terms of parameters for resolving the
// value of each attribute.
//
// It is important to note that a layout is defined by the totality of
// constraints; removing a child, or a constraint, from an existing layout
// without changing the remaining constraints may result in an unstable or
// unsolvable layout.
//
// Constraints have an implicit "reading order"; you should start describing
// each edge of each child, as well as their relationship with the parent
// container, from the top left (or top right, in RTL languages), horizontally
// first, and then vertically.
//
// A constraint-based layout with too few constraints can become "unstable",
// that is: have more than one solution. The behavior of an unstable layout is
// undefined.
//
// A constraint-based layout with conflicting constraints may be unsolvable, and
// lead to an unstable layout. You can use the Constraint:strength property of
// Constraint to "nudge" the layout towards a solution.
//
//
// GtkConstraintLayout as GtkBuildable
//
// GtkConstraintLayout implements the Buildable interface and has a custom
// "constraints" element which allows describing constraints in a GtkBuilder UI
// file.
//
// An example of a UI definition fragment specifying a constraint:
//
//    <object class="GtkConstraintLayout">
//      <constraints>
//        <constraint target="button" target-attribute="start"
//                    relation="eq"
//                    source="super" source-attribute="start"
//                    constant="12"
//                    strength="required" />
//        <constraint target="button" target-attribute="width"
//                    relation="ge"
//                    constant="250"
//                    strength="strong" />
//      </constraints>
//    </object>
//
// The definition above will add two constraints to the GtkConstraintLayout:
//
//    - a required constraint between the leading edge of "button" and
//      the leading edge of the widget using the constraint layout, plus
//      12 pixels
//    - a strong, constant constraint making the width of "button" greater
//      than, or equal to 250 pixels
//
// The "target" and "target-attribute" attributes are required.
//
// The "source" and "source-attribute" attributes of the "constraint" element
// are optional; if they are not specified, the constraint is assumed to be a
// constant.
//
// The "relation" attribute is optional; if not specified, the constraint is
// assumed to be an equality.
//
// The "strength" attribute is optional; if not specified, the constraint is
// assumed to be required.
//
// The "source" and "target" attributes can be set to "super" to indicate that
// the constraint target is the widget using the GtkConstraintLayout.
//
// There can be "constant" and "multiplier" attributes.
//
// Additionally, the "constraints" element can also contain a description of the
// ConstraintGuides used by the layout:
//
//    <constraints>
//      <guide min-width="100" max-width="500" name="hspace"/>
//      <guide min-height="64" nat-height="128" name="vspace" strength="strong"/>
//    </constraints>
//
// The "guide" element has the following optional attributes:
//
//    - "min-width", "nat-width", and "max-width", describe the minimum,
//      natural, and maximum width of the guide, respectively
//    - "min-height", "nat-height", and "max-height", describe the minimum,
//      natural, and maximum height of the guide, respectively
//    - "strength" describes the strength of the constraint on the natural
//      size of the guide; if not specified, the constraint is assumed to
//      have a medium strength
//    - "name" describes a name for the guide, useful when debugging
//
//
// Using the Visual Format Language
//
// Complex constraints can be described using a compact syntax called VFL, or
// *Visual Format Language*.
//
// The Visual Format Language describes all the constraints on a row or column,
// typically starting from the leading edge towards the trailing one. Each
// element of the layout is composed by "views", which identify a
// ConstraintTarget.
//
// For instance:
//
//    [button]-[textField]
//
// Describes a constraint that binds the trailing edge of "button" to the
// leading edge of "textField", leaving a default space between the two.
//
// Using VFL is also possible to specify predicates that describe constraints on
// attributes like width and height:
//
//      // Width must be greater than, or equal to 50
//      [button(>=50)]
//
//      // Width of button1 must be equal to width of button2
//      [button1(==button2)]
//
// The default orientation for a VFL description is horizontal, unless otherwise
// specified:
//
//      // horizontal orientation, default attribute: width
//      H:[button(>=150)]
//
//      // vertical orientation, default attribute: height
//      V:[button1(==button2)]
//
// It's also possible to specify multiple predicates, as well as their strength:
//
//    // minimum width of button must be 150
//    // natural width of button can be 250
//    [button(>=150@required, ==250@medium)]
//
// Finally, it's also possible to use simple arithmetic operators:
//
//    // width of button1 must be equal to width of button2
//    // divided by 2 plus 12
//    [button1(button2 / 2 + 12)]
type ConstraintLayout interface {
	LayoutManager
	Buildable

	// AddConstraint adds a Constraint to the layout manager.
	//
	// The Constraint:source and Constraint:target properties of @constraint can
	// be:
	//
	//    - set to nil to indicate that the constraint refers to the
	//      widget using @layout
	//    - set to the Widget using @layout
	//    - set to a child of the Widget using @layout
	//    - set to a guide that is part of @layout
	//
	// The @layout acquires the ownership of @constraint after calling this
	// function.
	AddConstraint(l ConstraintLayout, constraint Constraint)
	// AddGuide adds a guide to @layout. A guide can be used as the source or
	// target of constraints, like a widget, but it is not visible.
	//
	// The @layout acquires the ownership of @guide after calling this function.
	AddGuide(l ConstraintLayout, guide ConstraintGuide)
	// ObserveConstraints returns a Model to track the constraints that are part
	// of @layout.
	//
	// Calling this function will enable extra internal bookkeeping to track
	// constraints and emit signals on the returned listmodel. It may slow down
	// operations a lot.
	//
	// Applications should try hard to avoid calling this function because of
	// the slowdowns.
	ObserveConstraints(l ConstraintLayout)
	// ObserveGuides returns a Model to track the guides that are part of
	// @layout.
	//
	// Calling this function will enable extra internal bookkeeping to track
	// guides and emit signals on the returned listmodel. It may slow down
	// operations a lot.
	//
	// Applications should try hard to avoid calling this function because of
	// the slowdowns.
	ObserveGuides(l ConstraintLayout)
	// RemoveAllConstraints removes all constraints from the layout manager.
	RemoveAllConstraints(l ConstraintLayout)
	// RemoveConstraint removes @constraint from the layout manager, so that it
	// no longer influences the layout.
	RemoveConstraint(l ConstraintLayout, constraint Constraint)
	// RemoveGuide removes @guide from the layout manager, so that it no longer
	// influences the layout.
	RemoveGuide(l ConstraintLayout, guide ConstraintGuide)
}

// constraintLayout implements the ConstraintLayout interface.
type constraintLayout struct {
	LayoutManager
	Buildable
}

var _ ConstraintLayout = (*constraintLayout)(nil)

// WrapConstraintLayout wraps a GObject to the right type. It is
// primarily used internally.
func WrapConstraintLayout(obj *externglib.Object) ConstraintLayout {
	return ConstraintLayout{
		LayoutManager: WrapLayoutManager(obj),
		Buildable:     WrapBuildable(obj),
	}
}

func marshalConstraintLayout(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapConstraintLayout(obj), nil
}

// NewConstraintLayout constructs a class ConstraintLayout.
func NewConstraintLayout() {
	C.gtk_constraint_layout_new()
}

// AddConstraint adds a Constraint to the layout manager.
//
// The Constraint:source and Constraint:target properties of @constraint can
// be:
//
//    - set to nil to indicate that the constraint refers to the
//      widget using @layout
//    - set to the Widget using @layout
//    - set to a child of the Widget using @layout
//    - set to a guide that is part of @layout
//
// The @layout acquires the ownership of @constraint after calling this
// function.
func (l constraintLayout) AddConstraint(l ConstraintLayout, constraint Constraint) {
	var arg0 *C.GtkConstraintLayout
	var arg1 *C.GtkConstraint

	arg0 = (*C.GtkConstraintLayout)(unsafe.Pointer(l.Native()))
	arg1 = (*C.GtkConstraint)(unsafe.Pointer(constraint.Native()))

	C.gtk_constraint_layout_add_constraint(arg0, arg1)
}

// AddGuide adds a guide to @layout. A guide can be used as the source or
// target of constraints, like a widget, but it is not visible.
//
// The @layout acquires the ownership of @guide after calling this function.
func (l constraintLayout) AddGuide(l ConstraintLayout, guide ConstraintGuide) {
	var arg0 *C.GtkConstraintLayout
	var arg1 *C.GtkConstraintGuide

	arg0 = (*C.GtkConstraintLayout)(unsafe.Pointer(l.Native()))
	arg1 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))

	C.gtk_constraint_layout_add_guide(arg0, arg1)
}

// ObserveConstraints returns a Model to track the constraints that are part
// of @layout.
//
// Calling this function will enable extra internal bookkeeping to track
// constraints and emit signals on the returned listmodel. It may slow down
// operations a lot.
//
// Applications should try hard to avoid calling this function because of
// the slowdowns.
func (l constraintLayout) ObserveConstraints(l ConstraintLayout) {
	var arg0 *C.GtkConstraintLayout

	arg0 = (*C.GtkConstraintLayout)(unsafe.Pointer(l.Native()))

	C.gtk_constraint_layout_observe_constraints(arg0)
}

// ObserveGuides returns a Model to track the guides that are part of
// @layout.
//
// Calling this function will enable extra internal bookkeeping to track
// guides and emit signals on the returned listmodel. It may slow down
// operations a lot.
//
// Applications should try hard to avoid calling this function because of
// the slowdowns.
func (l constraintLayout) ObserveGuides(l ConstraintLayout) {
	var arg0 *C.GtkConstraintLayout

	arg0 = (*C.GtkConstraintLayout)(unsafe.Pointer(l.Native()))

	C.gtk_constraint_layout_observe_guides(arg0)
}

// RemoveAllConstraints removes all constraints from the layout manager.
func (l constraintLayout) RemoveAllConstraints(l ConstraintLayout) {
	var arg0 *C.GtkConstraintLayout

	arg0 = (*C.GtkConstraintLayout)(unsafe.Pointer(l.Native()))

	C.gtk_constraint_layout_remove_all_constraints(arg0)
}

// RemoveConstraint removes @constraint from the layout manager, so that it
// no longer influences the layout.
func (l constraintLayout) RemoveConstraint(l ConstraintLayout, constraint Constraint) {
	var arg0 *C.GtkConstraintLayout
	var arg1 *C.GtkConstraint

	arg0 = (*C.GtkConstraintLayout)(unsafe.Pointer(l.Native()))
	arg1 = (*C.GtkConstraint)(unsafe.Pointer(constraint.Native()))

	C.gtk_constraint_layout_remove_constraint(arg0, arg1)
}

// RemoveGuide removes @guide from the layout manager, so that it no longer
// influences the layout.
func (l constraintLayout) RemoveGuide(l ConstraintLayout, guide ConstraintGuide) {
	var arg0 *C.GtkConstraintLayout
	var arg1 *C.GtkConstraintGuide

	arg0 = (*C.GtkConstraintLayout)(unsafe.Pointer(l.Native()))
	arg1 = (*C.GtkConstraintGuide)(unsafe.Pointer(guide.Native()))

	C.gtk_constraint_layout_remove_guide(arg0, arg1)
}

// ConstraintLayoutChild: a LayoutChild in a ConstraintLayout.
type ConstraintLayoutChild interface {
	LayoutChild
}

// constraintLayoutChild implements the ConstraintLayoutChild interface.
type constraintLayoutChild struct {
	LayoutChild
}

var _ ConstraintLayoutChild = (*constraintLayoutChild)(nil)

// WrapConstraintLayoutChild wraps a GObject to the right type. It is
// primarily used internally.
func WrapConstraintLayoutChild(obj *externglib.Object) ConstraintLayoutChild {
	return ConstraintLayoutChild{
		LayoutChild: WrapLayoutChild(obj),
	}
}

func marshalConstraintLayoutChild(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapConstraintLayoutChild(obj), nil
}
