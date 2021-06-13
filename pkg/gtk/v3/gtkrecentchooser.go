// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/ptr"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_recent_chooser_get_type()), F: marshalRecentChooser},
	})
}

// RecentChooserOverrider contains methods that are overridable. This
// interface is a subset of the interface RecentChooser.
type RecentChooserOverrider interface {
	// AddFilter adds @filter to the list of RecentFilter objects held by
	// @chooser.
	//
	// If no previous filter objects were defined, this function will call
	// gtk_recent_chooser_set_filter().
	AddFilter(filter RecentFilter)
	// CurrentURI gets the URI currently selected by @chooser.
	CurrentURI() string

	ItemActivated()
	// RemoveFilter removes @filter from the list of RecentFilter objects held
	// by @chooser.
	RemoveFilter(filter RecentFilter)
	// SelectAll selects all the items inside @chooser, if the @chooser supports
	// multiple selection.
	SelectAll()
	// SelectURI selects @uri inside @chooser.
	SelectURI(uri string) error

	SelectionChanged()
	// SetCurrentURI sets @uri as the current URI for @chooser.
	SetCurrentURI(uri string) error
	// UnselectAll unselects all the items inside @chooser.
	UnselectAll()
	// UnselectURI unselects @uri inside @chooser.
	UnselectURI(uri string)
}

// RecentChooser is an interface that can be implemented by widgets displaying
// the list of recently used files. In GTK+, the main objects that implement
// this interface are RecentChooserWidget, RecentChooserDialog and
// RecentChooserMenu.
//
// Recently used files are supported since GTK+ 2.10.
type RecentChooser interface {
	gextras.Objector
	RecentChooserOverrider

	// Limit gets the number of items returned by gtk_recent_chooser_get_items()
	// and gtk_recent_chooser_get_uris().
	Limit() int
	// LocalOnly gets whether only local resources should be shown in the
	// recently used resources selector. See gtk_recent_chooser_set_local_only()
	LocalOnly() bool
	// SelectMultiple gets whether @chooser can select multiple items.
	SelectMultiple() bool
	// ShowIcons retrieves whether @chooser should show an icon near the
	// resource.
	ShowIcons() bool
	// ShowNotFound retrieves whether @chooser should show the recently used
	// resources that were not found.
	ShowNotFound() bool
	// ShowPrivate returns whether @chooser should display recently used
	// resources registered as private.
	ShowPrivate() bool
	// ShowTips gets whether @chooser should display tooltips containing the
	// full path of a recently user resource.
	ShowTips() bool
	// Uris gets the URI of the recently used resources.
	//
	// The return value of this function is affected by the “sort-type” and
	// “limit” properties of @chooser.
	//
	// Since the returned array is nil terminated, @length may be nil.
	Uris() []string
	// SetFilter sets @filter as the current RecentFilter object used by
	// @chooser to affect the displayed recently used resources.
	SetFilter(filter RecentFilter)
	// SetLimit sets the number of items that should be returned by
	// gtk_recent_chooser_get_items() and gtk_recent_chooser_get_uris().
	SetLimit(limit int)
	// SetLocalOnly sets whether only local resources, that is resources using
	// the file:// URI scheme, should be shown in the recently used resources
	// selector. If @local_only is true (the default) then the shown resources
	// are guaranteed to be accessible through the operating system native file
	// system.
	SetLocalOnly(localOnly bool)
	// SetSelectMultiple sets whether @chooser can select multiple items.
	SetSelectMultiple(selectMultiple bool)
	// SetShowIcons sets whether @chooser should show an icon near the resource
	// when displaying it.
	SetShowIcons(showIcons bool)
	// SetShowNotFound sets whether @chooser should display the recently used
	// resources that it didn’t find. This only applies to local resources.
	SetShowNotFound(showNotFound bool)
	// SetShowPrivate: whether to show recently used resources marked registered
	// as private.
	SetShowPrivate(showPrivate bool)
	// SetShowTips sets whether to show a tooltips containing the full path of
	// each recently used resource in a RecentChooser widget.
	SetShowTips(showTips bool)
	// SetSortType changes the sorting order of the recently used resources list
	// displayed by @chooser.
	SetSortType(sortType RecentSortType)
}

// recentChooser implements the RecentChooser interface.
type recentChooser struct {
	gextras.Objector
}

var _ RecentChooser = (*recentChooser)(nil)

// WrapRecentChooser wraps a GObject to a type that implements interface
// RecentChooser. It is primarily used internally.
func WrapRecentChooser(obj *externglib.Object) RecentChooser {
	return RecentChooser{
		Objector: obj,
	}
}

func marshalRecentChooser(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapRecentChooser(obj), nil
}

// AddFilter adds @filter to the list of RecentFilter objects held by
// @chooser.
//
// If no previous filter objects were defined, this function will call
// gtk_recent_chooser_set_filter().
func (c recentChooser) AddFilter(filter RecentFilter) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 *C.GtkRecentFilter  // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkRecentFilter)(unsafe.Pointer(filter.Native()))

	C.gtk_recent_chooser_add_filter(_arg0, _arg1)
}

// CurrentURI gets the URI currently selected by @chooser.
func (c recentChooser) CurrentURI() string {
	var _arg0 *C.GtkRecentChooser // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))

	var _cret *C.gchar // in

	_cret = C.gtk_recent_chooser_get_current_uri(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Limit gets the number of items returned by gtk_recent_chooser_get_items()
// and gtk_recent_chooser_get_uris().
func (c recentChooser) Limit() int {
	var _arg0 *C.GtkRecentChooser // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))

	var _cret C.gint // in

	_cret = C.gtk_recent_chooser_get_limit(_arg0)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// LocalOnly gets whether only local resources should be shown in the
// recently used resources selector. See gtk_recent_chooser_set_local_only()
func (c recentChooser) LocalOnly() bool {
	var _arg0 *C.GtkRecentChooser // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_recent_chooser_get_local_only(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// SelectMultiple gets whether @chooser can select multiple items.
func (c recentChooser) SelectMultiple() bool {
	var _arg0 *C.GtkRecentChooser // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_recent_chooser_get_select_multiple(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// ShowIcons retrieves whether @chooser should show an icon near the
// resource.
func (c recentChooser) ShowIcons() bool {
	var _arg0 *C.GtkRecentChooser // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_recent_chooser_get_show_icons(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// ShowNotFound retrieves whether @chooser should show the recently used
// resources that were not found.
func (c recentChooser) ShowNotFound() bool {
	var _arg0 *C.GtkRecentChooser // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_recent_chooser_get_show_not_found(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// ShowPrivate returns whether @chooser should display recently used
// resources registered as private.
func (c recentChooser) ShowPrivate() bool {
	var _arg0 *C.GtkRecentChooser // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_recent_chooser_get_show_private(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// ShowTips gets whether @chooser should display tooltips containing the
// full path of a recently user resource.
func (c recentChooser) ShowTips() bool {
	var _arg0 *C.GtkRecentChooser // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_recent_chooser_get_show_tips(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Uris gets the URI of the recently used resources.
//
// The return value of this function is affected by the “sort-type” and
// “limit” properties of @chooser.
//
// Since the returned array is nil terminated, @length may be nil.
func (c recentChooser) Uris() []string {
	var _arg0 *C.GtkRecentChooser // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))

	var _cret **C.gchar
	var _arg1 C.gsize // in

	_cret = C.gtk_recent_chooser_get_uris(_arg0, &_arg1)

	var _utf8s []string

	{
		var src []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(_cret), int(_arg1))

		_utf8s = make([]string, _arg1)
		for i := 0; i < uintptr(_arg1); i++ {
			_utf8s = C.GoString(_cret)
			defer C.free(unsafe.Pointer(_cret))
		}
	}

	return _utf8s
}

// RemoveFilter removes @filter from the list of RecentFilter objects held
// by @chooser.
func (c recentChooser) RemoveFilter(filter RecentFilter) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 *C.GtkRecentFilter  // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkRecentFilter)(unsafe.Pointer(filter.Native()))

	C.gtk_recent_chooser_remove_filter(_arg0, _arg1)
}

// SelectAll selects all the items inside @chooser, if the @chooser supports
// multiple selection.
func (c recentChooser) SelectAll() {
	var _arg0 *C.GtkRecentChooser // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))

	C.gtk_recent_chooser_select_all(_arg0)
}

// SelectURI selects @uri inside @chooser.
func (c recentChooser) SelectURI(uri string) error {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(_arg1))

	var _cerr *C.GError // in

	C.gtk_recent_chooser_select_uri(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SetCurrentURI sets @uri as the current URI for @chooser.
func (c recentChooser) SetCurrentURI(uri string) error {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(_arg1))

	var _cerr *C.GError // in

	C.gtk_recent_chooser_set_current_uri(_arg0, _arg1, &_cerr)

	var _goerr error // out

	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _goerr
}

// SetFilter sets @filter as the current RecentFilter object used by
// @chooser to affect the displayed recently used resources.
func (c recentChooser) SetFilter(filter RecentFilter) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 *C.GtkRecentFilter  // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkRecentFilter)(unsafe.Pointer(filter.Native()))

	C.gtk_recent_chooser_set_filter(_arg0, _arg1)
}

// SetLimit sets the number of items that should be returned by
// gtk_recent_chooser_get_items() and gtk_recent_chooser_get_uris().
func (c recentChooser) SetLimit(limit int) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 C.gint              // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))
	_arg1 = C.gint(limit)

	C.gtk_recent_chooser_set_limit(_arg0, _arg1)
}

// SetLocalOnly sets whether only local resources, that is resources using
// the file:// URI scheme, should be shown in the recently used resources
// selector. If @local_only is true (the default) then the shown resources
// are guaranteed to be accessible through the operating system native file
// system.
func (c recentChooser) SetLocalOnly(localOnly bool) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))
	if localOnly {
		_arg1 = C.gboolean(1)
	}

	C.gtk_recent_chooser_set_local_only(_arg0, _arg1)
}

// SetSelectMultiple sets whether @chooser can select multiple items.
func (c recentChooser) SetSelectMultiple(selectMultiple bool) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))
	if selectMultiple {
		_arg1 = C.gboolean(1)
	}

	C.gtk_recent_chooser_set_select_multiple(_arg0, _arg1)
}

// SetShowIcons sets whether @chooser should show an icon near the resource
// when displaying it.
func (c recentChooser) SetShowIcons(showIcons bool) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))
	if showIcons {
		_arg1 = C.gboolean(1)
	}

	C.gtk_recent_chooser_set_show_icons(_arg0, _arg1)
}

// SetShowNotFound sets whether @chooser should display the recently used
// resources that it didn’t find. This only applies to local resources.
func (c recentChooser) SetShowNotFound(showNotFound bool) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))
	if showNotFound {
		_arg1 = C.gboolean(1)
	}

	C.gtk_recent_chooser_set_show_not_found(_arg0, _arg1)
}

// SetShowPrivate: whether to show recently used resources marked registered
// as private.
func (c recentChooser) SetShowPrivate(showPrivate bool) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))
	if showPrivate {
		_arg1 = C.gboolean(1)
	}

	C.gtk_recent_chooser_set_show_private(_arg0, _arg1)
}

// SetShowTips sets whether to show a tooltips containing the full path of
// each recently used resource in a RecentChooser widget.
func (c recentChooser) SetShowTips(showTips bool) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))
	if showTips {
		_arg1 = C.gboolean(1)
	}

	C.gtk_recent_chooser_set_show_tips(_arg0, _arg1)
}

// SetSortType changes the sorting order of the recently used resources list
// displayed by @chooser.
func (c recentChooser) SetSortType(sortType RecentSortType) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 C.GtkRecentSortType // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))
	_arg1 = (C.GtkRecentSortType)(sortType)

	C.gtk_recent_chooser_set_sort_type(_arg0, _arg1)
}

// UnselectAll unselects all the items inside @chooser.
func (c recentChooser) UnselectAll() {
	var _arg0 *C.GtkRecentChooser // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))

	C.gtk_recent_chooser_unselect_all(_arg0)
}

// UnselectURI unselects @uri inside @chooser.
func (c recentChooser) UnselectURI(uri string) {
	var _arg0 *C.GtkRecentChooser // out
	var _arg1 *C.gchar            // out

	_arg0 = (*C.GtkRecentChooser)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(uri))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_recent_chooser_unselect_uri(_arg0, _arg1)
}
