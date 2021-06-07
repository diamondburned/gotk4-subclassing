// Code generated by girgen. DO NOT EDIT.

package pangoot

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: pangoot pango
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <pango/pango-ot.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.pango_ot_info_get_type()), F: marshalInfo},
		{T: externglib.Type(C.pango_ot_ruleset_get_type()), F: marshalRuleset},
	})
}

type Info interface {
	gextras.Objector

	// ListScripts obtains the list of available scripts.
	ListScripts(i Info, tableType TableType)
}

// info implements the Info interface.
type info struct {
	gextras.Objector
}

var _ Info = (*info)(nil)

// WrapInfo wraps a GObject to the right type. It is
// primarily used internally.
func WrapInfo(obj *externglib.Object) Info {
	return Info{
		Objector: obj,
	}
}

func marshalInfo(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapInfo(obj), nil
}

// ListScripts obtains the list of available scripts.
func (i info) ListScripts(i Info, tableType TableType) {
	var arg0 *C.PangoOTInfo
	var arg1 C.PangoOTTableType

	arg0 = (*C.PangoOTInfo)(unsafe.Pointer(i.Native()))
	arg1 = (C.PangoOTTableType)(tableType)

	C.pango_ot_info_list_scripts(arg0, arg1)
}

// Ruleset: the OTRuleset structure holds a set of features selected from the
// tables in an OpenType font. (A feature is an operation such as adjusting
// glyph positioning that should be applied to a text feature such as a certain
// type of accent.) A OTRuleset is created with pango_ot_ruleset_new(), features
// are added to it with pango_ot_ruleset_add_feature(), then it is applied to a
// GlyphString with pango_ot_ruleset_shape().
type Ruleset interface {
	gextras.Objector

	// AddFeature adds a feature to the ruleset.
	AddFeature(r Ruleset, tableType TableType, featureIndex uint, propertyBit uint32)
	// FeatureCount gets the number of GSUB and GPOS features in the ruleset.
	FeatureCount(r Ruleset) (nGsubFeatures uint, nGposFeatures uint)
	// MaybeAddFeatures: this is a convenience function that for each feature in
	// the feature map array @features converts the feature name to a OTTag
	// feature tag using PANGO_OT_TAG_MAKE() and calls
	// pango_ot_ruleset_maybe_add_feature() on it.
	MaybeAddFeatures(r Ruleset, tableType TableType, features *FeatureMap, nFeatures uint)
	// Position performs the OpenType GPOS positioning on @buffer using the
	// features in @ruleset
	Position(r Ruleset, buffer *Buffer)
	// Substitute performs the OpenType GSUB substitution on @buffer using the
	// features in @ruleset
	Substitute(r Ruleset, buffer *Buffer)
}

// ruleset implements the Ruleset interface.
type ruleset struct {
	gextras.Objector
}

var _ Ruleset = (*ruleset)(nil)

// WrapRuleset wraps a GObject to the right type. It is
// primarily used internally.
func WrapRuleset(obj *externglib.Object) Ruleset {
	return Ruleset{
		Objector: obj,
	}
}

func marshalRuleset(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRuleset(obj), nil
}

// NewRuleset constructs a class Ruleset.
func NewRuleset(info Info) {
	var arg1 *C.PangoOTInfo

	arg1 = (*C.PangoOTInfo)(unsafe.Pointer(info.Native()))

	C.pango_ot_ruleset_new(arg1)
}

// NewRulesetFor constructs a class Ruleset.
func NewRulesetFor(info Info, script pango.Script, language *pango.Language) {
	var arg1 *C.PangoOTInfo
	var arg2 C.PangoScript
	var arg3 *C.PangoLanguage

	arg1 = (*C.PangoOTInfo)(unsafe.Pointer(info.Native()))
	arg2 = (C.PangoScript)(script)
	arg3 = (*C.PangoLanguage)(unsafe.Pointer(language.Native()))

	C.pango_ot_ruleset_new_for(arg1, arg2, arg3)
}

// NewRulesetFromDescription constructs a class Ruleset.
func NewRulesetFromDescription(info Info, desc *RulesetDescription) {
	var arg1 *C.PangoOTInfo
	var arg2 *C.PangoOTRulesetDescription

	arg1 = (*C.PangoOTInfo)(unsafe.Pointer(info.Native()))
	arg2 = (*C.PangoOTRulesetDescription)(unsafe.Pointer(desc.Native()))

	C.pango_ot_ruleset_new_from_description(arg1, arg2)
}

// AddFeature adds a feature to the ruleset.
func (r ruleset) AddFeature(r Ruleset, tableType TableType, featureIndex uint, propertyBit uint32) {
	var arg0 *C.PangoOTRuleset
	var arg1 C.PangoOTTableType
	var arg2 C.guint
	var arg3 C.gulong

	arg0 = (*C.PangoOTRuleset)(unsafe.Pointer(r.Native()))
	arg1 = (C.PangoOTTableType)(tableType)
	arg2 = C.guint(featureIndex)
	arg3 = C.gulong(propertyBit)

	C.pango_ot_ruleset_add_feature(arg0, arg1, arg2, arg3)
}

// FeatureCount gets the number of GSUB and GPOS features in the ruleset.
func (r ruleset) FeatureCount(r Ruleset) (nGsubFeatures uint, nGposFeatures uint) {
	var arg0 *C.PangoOTRuleset

	arg0 = (*C.PangoOTRuleset)(unsafe.Pointer(r.Native()))

	var arg1 C.guint
	var nGsubFeatures uint
	var arg2 C.guint
	var nGposFeatures uint

	C.pango_ot_ruleset_get_feature_count(arg0, &arg1, &arg2)

	nGsubFeatures = uint(&arg1)
	nGposFeatures = uint(&arg2)

	return nGsubFeatures, nGposFeatures
}

// MaybeAddFeatures: this is a convenience function that for each feature in
// the feature map array @features converts the feature name to a OTTag
// feature tag using PANGO_OT_TAG_MAKE() and calls
// pango_ot_ruleset_maybe_add_feature() on it.
func (r ruleset) MaybeAddFeatures(r Ruleset, tableType TableType, features *FeatureMap, nFeatures uint) {
	var arg0 *C.PangoOTRuleset
	var arg1 C.PangoOTTableType
	var arg2 *C.PangoOTFeatureMap
	var arg3 C.guint

	arg0 = (*C.PangoOTRuleset)(unsafe.Pointer(r.Native()))
	arg1 = (C.PangoOTTableType)(tableType)
	arg2 = (*C.PangoOTFeatureMap)(unsafe.Pointer(features.Native()))
	arg3 = C.guint(nFeatures)

	C.pango_ot_ruleset_maybe_add_features(arg0, arg1, arg2, arg3)
}

// Position performs the OpenType GPOS positioning on @buffer using the
// features in @ruleset
func (r ruleset) Position(r Ruleset, buffer *Buffer) {
	var arg0 *C.PangoOTRuleset
	var arg1 *C.PangoOTBuffer

	arg0 = (*C.PangoOTRuleset)(unsafe.Pointer(r.Native()))
	arg1 = (*C.PangoOTBuffer)(unsafe.Pointer(buffer.Native()))

	C.pango_ot_ruleset_position(arg0, arg1)
}

// Substitute performs the OpenType GSUB substitution on @buffer using the
// features in @ruleset
func (r ruleset) Substitute(r Ruleset, buffer *Buffer) {
	var arg0 *C.PangoOTRuleset
	var arg1 *C.PangoOTBuffer

	arg0 = (*C.PangoOTRuleset)(unsafe.Pointer(r.Native()))
	arg1 = (*C.PangoOTBuffer)(unsafe.Pointer(buffer.Native()))

	C.pango_ot_ruleset_substitute(arg0, arg1)
}
