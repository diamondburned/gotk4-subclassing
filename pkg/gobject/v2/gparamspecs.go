// Code generated by girgen. DO NOT EDIT.

package gobject

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gobject-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <glib-object.h>
import "C"

// ParamSpecBoolean creates a new SpecBoolean instance specifying a
// G_TYPE_BOOLEAN property. In many cases, it may be more appropriate to use an
// enum with g_param_spec_enum(), both to improve code clarity by using
// explicitly named values, and to allow for more values to be added in future
// without breaking API.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecBoolean(name string, nick string, blurb string, defaultValue bool, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.gchar      // out
	var _arg4 C.gboolean    // out
	var _arg5 C.GParamFlags // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	if defaultValue {
		_arg4 = C.TRUE
	}
	_arg5 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_boolean(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// ParamSpecBoxed creates a new SpecBoxed instance specifying a G_TYPE_BOXED
// derived property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecBoxed(name string, nick string, blurb string, boxedType externglib.Type, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.gchar      // out
	var _arg4 C.GType       // out
	var _arg5 C.GParamFlags // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.GType(boxedType)
	_arg5 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_boxed(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// ParamSpecChar creates a new SpecChar instance specifying a G_TYPE_CHAR
// property.
func ParamSpecChar(name string, nick string, blurb string, minimum int8, maximum int8, defaultValue int8, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.gchar      // out
	var _arg4 C.gint8       // out
	var _arg5 C.gint8       // out
	var _arg6 C.gint8       // out
	var _arg7 C.GParamFlags // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.gint8(minimum)
	_arg5 = C.gint8(maximum)
	_arg6 = C.gint8(defaultValue)
	_arg7 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_char(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// ParamSpecDouble creates a new SpecDouble instance specifying a G_TYPE_DOUBLE
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecDouble(name string, nick string, blurb string, minimum float64, maximum float64, defaultValue float64, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.gchar      // out
	var _arg4 C.gdouble     // out
	var _arg5 C.gdouble     // out
	var _arg6 C.gdouble     // out
	var _arg7 C.GParamFlags // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.gdouble(minimum)
	_arg5 = C.gdouble(maximum)
	_arg6 = C.gdouble(defaultValue)
	_arg7 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_double(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// ParamSpecEnum creates a new SpecEnum instance specifying a G_TYPE_ENUM
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecEnum(name string, nick string, blurb string, enumType externglib.Type, defaultValue int, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.gchar      // out
	var _arg4 C.GType       // out
	var _arg5 C.gint        // out
	var _arg6 C.GParamFlags // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.GType(enumType)
	_arg5 = C.gint(defaultValue)
	_arg6 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_enum(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// ParamSpecFlags creates a new SpecFlags instance specifying a G_TYPE_FLAGS
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecFlags(name string, nick string, blurb string, flagsType externglib.Type, defaultValue uint, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.gchar      // out
	var _arg4 C.GType       // out
	var _arg5 C.guint       // out
	var _arg6 C.GParamFlags // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.GType(flagsType)
	_arg5 = C.guint(defaultValue)
	_arg6 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_flags(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// ParamSpecFloat creates a new SpecFloat instance specifying a G_TYPE_FLOAT
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecFloat(name string, nick string, blurb string, minimum float32, maximum float32, defaultValue float32, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.gchar      // out
	var _arg4 C.gfloat      // out
	var _arg5 C.gfloat      // out
	var _arg6 C.gfloat      // out
	var _arg7 C.GParamFlags // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.gfloat(minimum)
	_arg5 = C.gfloat(maximum)
	_arg6 = C.gfloat(defaultValue)
	_arg7 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_float(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// ParamSpecGType creates a new SpecGType instance specifying a G_TYPE_GTYPE
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecGType(name string, nick string, blurb string, isAType externglib.Type, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.gchar      // out
	var _arg4 C.GType       // out
	var _arg5 C.GParamFlags // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.GType(isAType)
	_arg5 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_gtype(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// ParamSpecInt creates a new SpecInt instance specifying a G_TYPE_INT property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecInt(name string, nick string, blurb string, minimum int, maximum int, defaultValue int, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.gchar      // out
	var _arg4 C.gint        // out
	var _arg5 C.gint        // out
	var _arg6 C.gint        // out
	var _arg7 C.GParamFlags // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.gint(minimum)
	_arg5 = C.gint(maximum)
	_arg6 = C.gint(defaultValue)
	_arg7 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_int(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// ParamSpecInt64 creates a new SpecInt64 instance specifying a G_TYPE_INT64
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecInt64(name string, nick string, blurb string, minimum int64, maximum int64, defaultValue int64, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.gchar      // out
	var _arg4 C.gint64      // out
	var _arg5 C.gint64      // out
	var _arg6 C.gint64      // out
	var _arg7 C.GParamFlags // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.gint64(minimum)
	_arg5 = C.gint64(maximum)
	_arg6 = C.gint64(defaultValue)
	_arg7 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_int64(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// ParamSpecLong creates a new SpecLong instance specifying a G_TYPE_LONG
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecLong(name string, nick string, blurb string, minimum int32, maximum int32, defaultValue int32, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.gchar      // out
	var _arg4 C.glong       // out
	var _arg5 C.glong       // out
	var _arg6 C.glong       // out
	var _arg7 C.GParamFlags // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.glong(minimum)
	_arg5 = C.glong(maximum)
	_arg6 = C.glong(defaultValue)
	_arg7 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_long(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// ParamSpecObject creates a new SpecBoxed instance specifying a G_TYPE_OBJECT
// derived property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecObject(name string, nick string, blurb string, objectType externglib.Type, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.gchar      // out
	var _arg4 C.GType       // out
	var _arg5 C.GParamFlags // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.GType(objectType)
	_arg5 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_object(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// ParamSpecParam creates a new SpecParam instance specifying a G_TYPE_PARAM
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecParam(name string, nick string, blurb string, paramType externglib.Type, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.gchar      // out
	var _arg4 C.GType       // out
	var _arg5 C.GParamFlags // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.GType(paramType)
	_arg5 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_param(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// ParamSpecPointer creates a new SpecPointer instance specifying a pointer
// property. Where possible, it is better to use g_param_spec_object() or
// g_param_spec_boxed() to expose memory management information.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecPointer(name string, nick string, blurb string, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.gchar      // out
	var _arg4 C.GParamFlags // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_pointer(_arg1, _arg2, _arg3, _arg4)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// ParamSpecString creates a new SpecString instance.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecString(name string, nick string, blurb string, defaultValue string, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.gchar      // out
	var _arg4 *C.gchar      // out
	var _arg5 C.GParamFlags // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(C.CString(defaultValue))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_string(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// ParamSpecUchar creates a new SpecUChar instance specifying a G_TYPE_UCHAR
// property.
func ParamSpecUchar(name string, nick string, blurb string, minimum byte, maximum byte, defaultValue byte, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.gchar      // out
	var _arg4 C.guint8      // out
	var _arg5 C.guint8      // out
	var _arg6 C.guint8      // out
	var _arg7 C.GParamFlags // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.guint8(minimum)
	_arg5 = C.guint8(maximum)
	_arg6 = C.guint8(defaultValue)
	_arg7 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_uchar(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// ParamSpecUint creates a new SpecUInt instance specifying a G_TYPE_UINT
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecUint(name string, nick string, blurb string, minimum uint, maximum uint, defaultValue uint, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.gchar      // out
	var _arg4 C.guint       // out
	var _arg5 C.guint       // out
	var _arg6 C.guint       // out
	var _arg7 C.GParamFlags // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.guint(minimum)
	_arg5 = C.guint(maximum)
	_arg6 = C.guint(defaultValue)
	_arg7 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_uint(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// ParamSpecUint64 creates a new SpecUInt64 instance specifying a G_TYPE_UINT64
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecUint64(name string, nick string, blurb string, minimum uint64, maximum uint64, defaultValue uint64, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.gchar      // out
	var _arg4 C.guint64     // out
	var _arg5 C.guint64     // out
	var _arg6 C.guint64     // out
	var _arg7 C.GParamFlags // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.guint64(minimum)
	_arg5 = C.guint64(maximum)
	_arg6 = C.guint64(defaultValue)
	_arg7 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_uint64(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// ParamSpecUlong creates a new SpecULong instance specifying a G_TYPE_ULONG
// property.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecUlong(name string, nick string, blurb string, minimum uint32, maximum uint32, defaultValue uint32, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.gchar      // out
	var _arg4 C.gulong      // out
	var _arg5 C.gulong      // out
	var _arg6 C.gulong      // out
	var _arg7 C.GParamFlags // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.gulong(minimum)
	_arg5 = C.gulong(maximum)
	_arg6 = C.gulong(defaultValue)
	_arg7 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_ulong(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// ParamSpecUnichar creates a new SpecUnichar instance specifying a G_TYPE_UINT
// property. #GValue structures for this property can be accessed with
// g_value_set_uint() and g_value_get_uint().
//
// See g_param_spec_internal() for details on property names.
func ParamSpecUnichar(name string, nick string, blurb string, defaultValue uint32, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar      // out
	var _arg2 *C.gchar      // out
	var _arg3 *C.gchar      // out
	var _arg4 C.gunichar    // out
	var _arg5 C.GParamFlags // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.gunichar(defaultValue)
	_arg5 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_unichar(_arg1, _arg2, _arg3, _arg4, _arg5)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}

// ParamSpecVariant creates a new SpecVariant instance specifying a #GVariant
// property.
//
// If @default_value is floating, it is consumed.
//
// See g_param_spec_internal() for details on property names.
func ParamSpecVariant(name string, nick string, blurb string, typ *glib.VariantType, defaultValue *glib.Variant, flags ParamFlags) ParamSpec {
	var _arg1 *C.gchar        // out
	var _arg2 *C.gchar        // out
	var _arg3 *C.gchar        // out
	var _arg4 *C.GVariantType // out
	var _arg5 *C.GVariant     // out
	var _arg6 C.GParamFlags   // out

	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(C.CString(nick))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(C.CString(blurb))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.GVariantType)(unsafe.Pointer(typ.Native()))
	_arg5 = (*C.GVariant)(unsafe.Pointer(defaultValue.Native()))
	_arg6 = (C.GParamFlags)(flags)

	var _cret *C.GParamSpec // in

	_cret = C.g_param_spec_variant(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)

	var _paramSpec ParamSpec // out

	_paramSpec = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(ParamSpec)

	return _paramSpec
}
