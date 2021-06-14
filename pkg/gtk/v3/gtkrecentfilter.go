// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: glib-2.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_recent_filter_flags_get_type()), F: marshalRecentFilterFlags},
		{T: externglib.Type(C.gtk_recent_filter_get_type()), F: marshalRecentFilter},
	})
}

// RecentFilterFlags: these flags indicate what parts of a RecentFilterInfo
// struct are filled or need to be filled.
type RecentFilterFlags int

const (
	// RecentFilterFlagsURI: the URI of the file being tested
	RecentFilterFlagsURI RecentFilterFlags = 1
	// RecentFilterFlagsDisplayName: the string that will be used to display the
	// file in the recent chooser
	RecentFilterFlagsDisplayName RecentFilterFlags = 2
	// RecentFilterFlagsMIMEType: the mime type of the file
	RecentFilterFlagsMIMEType RecentFilterFlags = 4
	// RecentFilterFlagsApplication: the list of applications that have
	// registered the file
	RecentFilterFlagsApplication RecentFilterFlags = 8
	// RecentFilterFlagsGroup: the groups to which the file belongs to
	RecentFilterFlagsGroup RecentFilterFlags = 16
	// RecentFilterFlagsAge: the number of days elapsed since the file has been
	// registered
	RecentFilterFlagsAge RecentFilterFlags = 32
)

func marshalRecentFilterFlags(p uintptr) (interface{}, error) {
	return RecentFilterFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// RecentFilter: a RecentFilter can be used to restrict the files being shown in
// a RecentChooser. Files can be filtered based on their name (with
// gtk_recent_filter_add_pattern()), on their mime type (with
// gtk_file_filter_add_mime_type()), on the application that has registered them
// (with gtk_recent_filter_add_application()), or by a custom filter function
// (with gtk_recent_filter_add_custom()).
//
// Filtering by mime type handles aliasing and subclassing of mime types; e.g. a
// filter for text/plain also matches a file with mime type application/rtf,
// since application/rtf is a subclass of text/plain. Note that RecentFilter
// allows wildcards for the subtype of a mime type, so you can e.g. filter for
// image/\*.
//
// Normally, filters are used by adding them to a RecentChooser, see
// gtk_recent_chooser_add_filter(), but it is also possible to manually use a
// filter on a file with gtk_recent_filter_filter().
//
// Recently used files are supported since GTK+ 2.10.
//
//
// GtkRecentFilter as GtkBuildable
//
// The GtkRecentFilter implementation of the GtkBuildable interface supports
// adding rules using the <mime-types>, <patterns> and <applications> elements
// and listing the rules within. Specifying a <mime-type>, <pattern> or
// <application> has the same effect as calling
// gtk_recent_filter_add_mime_type(), gtk_recent_filter_add_pattern() or
// gtk_recent_filter_add_application().
//
// An example of a UI definition fragment specifying GtkRecentFilter rules:
//
//    <object class="GtkRecentFilter">
//      <mime-types>
//        <mime-type>text/plain</mime-type>
//        <mime-type>image/png</mime-type>
//      </mime-types>
//      <patterns>
//        <pattern>*.txt</pattern>
//        <pattern>*.png</pattern>
//      </patterns>
//      <applications>
//        <application>gimp</application>
//        <application>gedit</application>
//        <application>glade</application>
//      </applications>
//    </object>
type RecentFilter interface {
	gextras.Objector
	Buildable

	// AddAge adds a rule that allows resources based on their age - that is,
	// the number of days elapsed since they were last modified.
	AddAge(days int)
	// AddApplication adds a rule that allows resources based on the name of the
	// application that has registered them.
	AddApplication(application string)
	// AddGroup adds a rule that allows resources based on the name of the group
	// to which they belong
	AddGroup(group string)
	// AddMIMEType adds a rule that allows resources based on their registered
	// MIME type.
	AddMIMEType(mimeType string)
	// AddPattern adds a rule that allows resources based on a pattern matching
	// their display name.
	AddPattern(pattern string)
	// AddPixbufFormats adds a rule allowing image files in the formats
	// supported by GdkPixbuf.
	AddPixbufFormats()
	// Filter tests whether a file should be displayed according to @filter. The
	// RecentFilterInfo @filter_info should include the fields returned from
	// gtk_recent_filter_get_needed(), and must set the
	// RecentFilterInfo.contains field of @filter_info to indicate which fields
	// have been set.
	//
	// This function will not typically be used by applications; it is intended
	// principally for use in the implementation of RecentChooser.
	Filter(filterInfo *RecentFilterInfo) bool
	// Name gets the human-readable name for the filter. See
	// gtk_recent_filter_set_name().
	Name() string
	// Needed gets the fields that need to be filled in for the RecentFilterInfo
	// passed to gtk_recent_filter_filter()
	//
	// This function will not typically be used by applications; it is intended
	// principally for use in the implementation of RecentChooser.
	Needed() RecentFilterFlags
	// SetName sets the human-readable name of the filter; this is the string
	// that will be displayed in the recently used resources selector user
	// interface if there is a selectable list of filters.
	SetName(name string)
}

// recentFilter implements the RecentFilter class.
type recentFilter struct {
	gextras.Objector
	Buildable
}

var _ RecentFilter = (*recentFilter)(nil)

// WrapRecentFilter wraps a GObject to the right type. It is
// primarily used internally.
func WrapRecentFilter(obj *externglib.Object) RecentFilter {
	return recentFilter{
		Objector:  obj,
		Buildable: WrapBuildable(obj),
	}
}

func marshalRecentFilter(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRecentFilter(obj), nil
}

// NewRecentFilter constructs a class RecentFilter.
func NewRecentFilter() RecentFilter {
	var _cret C.GtkRecentFilter // in

	_cret = C.gtk_recent_filter_new()

	var _recentFilter RecentFilter // out

	_recentFilter = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(RecentFilter)

	return _recentFilter
}

// AddAge adds a rule that allows resources based on their age - that is,
// the number of days elapsed since they were last modified.
func (f recentFilter) AddAge(days int) {
	var _arg0 *C.GtkRecentFilter // out
	var _arg1 C.gint             // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))
	_arg1 = C.gint(days)

	C.gtk_recent_filter_add_age(_arg0, _arg1)
}

// AddApplication adds a rule that allows resources based on the name of the
// application that has registered them.
func (f recentFilter) AddApplication(application string) {
	var _arg0 *C.GtkRecentFilter // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(application))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_filter_add_application(_arg0, _arg1)
}

// AddGroup adds a rule that allows resources based on the name of the group
// to which they belong
func (f recentFilter) AddGroup(group string) {
	var _arg0 *C.GtkRecentFilter // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(group))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_filter_add_group(_arg0, _arg1)
}

// AddMIMEType adds a rule that allows resources based on their registered
// MIME type.
func (f recentFilter) AddMIMEType(mimeType string) {
	var _arg0 *C.GtkRecentFilter // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(mimeType))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_filter_add_mime_type(_arg0, _arg1)
}

// AddPattern adds a rule that allows resources based on a pattern matching
// their display name.
func (f recentFilter) AddPattern(pattern string) {
	var _arg0 *C.GtkRecentFilter // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(pattern))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_filter_add_pattern(_arg0, _arg1)
}

// AddPixbufFormats adds a rule allowing image files in the formats
// supported by GdkPixbuf.
func (f recentFilter) AddPixbufFormats() {
	var _arg0 *C.GtkRecentFilter // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))

	C.gtk_recent_filter_add_pixbuf_formats(_arg0)
}

// Filter tests whether a file should be displayed according to @filter. The
// RecentFilterInfo @filter_info should include the fields returned from
// gtk_recent_filter_get_needed(), and must set the
// RecentFilterInfo.contains field of @filter_info to indicate which fields
// have been set.
//
// This function will not typically be used by applications; it is intended
// principally for use in the implementation of RecentChooser.
func (f recentFilter) Filter(filterInfo *RecentFilterInfo) bool {
	var _arg0 *C.GtkRecentFilter     // out
	var _arg1 *C.GtkRecentFilterInfo // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.GtkRecentFilterInfo)(unsafe.Pointer(filterInfo.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_recent_filter_filter(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Name gets the human-readable name for the filter. See
// gtk_recent_filter_set_name().
func (f recentFilter) Name() string {
	var _arg0 *C.GtkRecentFilter // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))

	var _cret *C.gchar // in

	_cret = C.gtk_recent_filter_get_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Needed gets the fields that need to be filled in for the RecentFilterInfo
// passed to gtk_recent_filter_filter()
//
// This function will not typically be used by applications; it is intended
// principally for use in the implementation of RecentChooser.
func (f recentFilter) Needed() RecentFilterFlags {
	var _arg0 *C.GtkRecentFilter // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))

	var _cret C.GtkRecentFilterFlags // in

	_cret = C.gtk_recent_filter_get_needed(_arg0)

	var _recentFilterFlags RecentFilterFlags // out

	_recentFilterFlags = RecentFilterFlags(_cret)

	return _recentFilterFlags
}

// SetName sets the human-readable name of the filter; this is the string
// that will be displayed in the recently used resources selector user
// interface if there is a selectable list of filters.
func (f recentFilter) SetName(name string) {
	var _arg0 *C.GtkRecentFilter // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GtkRecentFilter)(unsafe.Pointer(f.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_filter_set_name(_arg0, _arg1)
}

// RecentFilterInfo: a GtkRecentFilterInfo struct is used to pass information
// about the tested file to gtk_recent_filter_filter().
type RecentFilterInfo struct {
	native C.GtkRecentFilterInfo
}

// WrapRecentFilterInfo wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapRecentFilterInfo(ptr unsafe.Pointer) *RecentFilterInfo {
	if ptr == nil {
		return nil
	}

	return (*RecentFilterInfo)(ptr)
}

// Native returns the underlying C source pointer.
func (r *RecentFilterInfo) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// Contains gets the field inside the struct.
func (r *RecentFilterInfo) Contains() RecentFilterFlags {
	var v RecentFilterFlags // out
	v = RecentFilterFlags(r.native.contains)
	return v
}

// URI gets the field inside the struct.
func (r *RecentFilterInfo) URI() string {
	var v string // out
	v = C.GoString(r.native.uri)
	return v
}

// DisplayName gets the field inside the struct.
func (r *RecentFilterInfo) DisplayName() string {
	var v string // out
	v = C.GoString(r.native.display_name)
	return v
}

// MIMEType gets the field inside the struct.
func (r *RecentFilterInfo) MIMEType() string {
	var v string // out
	v = C.GoString(r.native.mime_type)
	return v
}

// Applications gets the field inside the struct.
func (r *RecentFilterInfo) Applications() []string {
	var v []string
	{
		var i int
		for p := r.native.applications; *p != nil; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(r.native.applications, i)
		v = make([]string, i)
		for i := range src {
			v[i] = C.GoString(src[i])
		}
	}
	return v
}

// Groups gets the field inside the struct.
func (r *RecentFilterInfo) Groups() []string {
	var v []string
	{
		var i int
		for p := r.native.groups; *p != nil; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(r.native.groups, i)
		v = make([]string, i)
		for i := range src {
			v[i] = C.GoString(src[i])
		}
	}
	return v
}

// Age gets the field inside the struct.
func (r *RecentFilterInfo) Age() int {
	var v int // out
	v = (int)(r.native.age)
	return v
}
