// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"unsafe"

	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: atk
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.atk_relation_type_get_type()), F: marshalRelationType},
	})
}

// RelationType describes the type of the relation
type RelationType int

const (
	// Null: not used, represens "no relationship" or an error condition.
	Null RelationType = iota
	// ControlledBy indicates an object controlled by one or more target
	// objects.
	ControlledBy
	// ControllerFor indicates an object is an controller for one or more target
	// objects.
	ControllerFor
	// LabelFor indicates an object is a label for one or more target objects.
	LabelFor
	// LabelledBy indicates an object is labelled by one or more target objects.
	LabelledBy
	// MemberOf indicates an object is a member of a group of one or more target
	// objects.
	MemberOf
	// NodeChildOf indicates an object is a cell in a treetable which is
	// displayed because a cell in the same column is expanded and identifies
	// that cell.
	NodeChildOf
	// FlowsTo indicates that the object has content that flows logically to
	// another AtkObject in a sequential way, (for instance text-flow).
	FlowsTo
	// FlowsFrom indicates that the object has content that flows logically from
	// another AtkObject in a sequential way, (for instance text-flow).
	FlowsFrom
	// SubwindowOf indicates a subwindow attached to a component but otherwise
	// has no connection in the UI heirarchy to that component.
	SubwindowOf
	// Embeds indicates that the object visually embeds another object's
	// content, i.e. this object's content flows around another's content.
	Embeds
	// EmbeddedBy: reciprocal of ATK_RELATION_EMBEDS, indicates that this
	// object's content is visualy embedded in another object.
	EmbeddedBy
	// PopupFor indicates that an object is a popup for another object.
	PopupFor
	// ParentWindowOf indicates that an object is a parent window of another
	// object.
	ParentWindowOf
	// DescribedBy: reciprocal of ATK_RELATION_DESCRIPTION_FOR. Indicates that
	// one or more target objects provide descriptive information about this
	// object. This relation type is most appropriate for information that is
	// not essential as its presentation may be user-configurable and/or limited
	// to an on-demand mechanism such as an assistive technology command. For
	// brief, essential information such as can be found in a widget's on-screen
	// label, use ATK_RELATION_LABELLED_BY. For an on-screen error message, use
	// ATK_RELATION_ERROR_MESSAGE. For lengthy extended descriptive information
	// contained in an on-screen object, consider using ATK_RELATION_DETAILS as
	// assistive technologies may provide a means for the user to navigate to
	// objects containing detailed descriptions so that their content can be
	// more closely reviewed.
	DescribedBy
	// DescriptionFor: reciprocal of ATK_RELATION_DESCRIBED_BY. Indicates that
	// this object provides descriptive information about the target object(s).
	// See also ATK_RELATION_DETAILS_FOR and ATK_RELATION_ERROR_FOR.
	DescriptionFor
	// NodeParentOf indicates an object is a cell in a treetable and is expanded
	// to display other cells in the same column.
	NodeParentOf
	// Details: reciprocal of ATK_RELATION_DETAILS_FOR. Indicates that this
	// object has a detailed or extended description, the contents of which can
	// be found in the target object(s). This relation type is most appropriate
	// for information that is sufficiently lengthy as to make navigation to the
	// container of that information desirable. For less verbose information
	// suitable for announcement only, see ATK_RELATION_DESCRIBED_BY. If the
	// detailed information describes an error condition, ATK_RELATION_ERROR_FOR
	// should be used instead. @Since: ATK-2.26.
	Details
	// DetailsFor: reciprocal of ATK_RELATION_DETAILS. Indicates that this
	// object provides a detailed or extended description about the target
	// object(s). See also ATK_RELATION_DESCRIPTION_FOR and
	// ATK_RELATION_ERROR_FOR. @Since: ATK-2.26.
	DetailsFor
	// ErrorMessage: reciprocal of ATK_RELATION_ERROR_FOR. Indicates that this
	// object has one or more errors, the nature of which is described in the
	// contents of the target object(s). Objects that have this relation type
	// should also contain ATK_STATE_INVALID_ENTRY in their StateSet. @Since:
	// ATK-2.26.
	ErrorMessage
	// ErrorFor: reciprocal of ATK_RELATION_ERROR_MESSAGE. Indicates that this
	// object contains an error message describing an invalid condition in the
	// target object(s). @Since: ATK_2.26.
	ErrorFor
	// LastDefined: not used, this value indicates the end of the enumeration.
	LastDefined
)

func marshalRelationType(p uintptr) (interface{}, error) {
	return RelationType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}
