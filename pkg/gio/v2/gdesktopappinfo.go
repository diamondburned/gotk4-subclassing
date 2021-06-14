// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_desktop_app_info_lookup_get_type()), F: marshalDesktopAppInfoLookup},
		{T: externglib.Type(C.g_desktop_app_info_get_type()), F: marshalDesktopAppInfo},
	})
}

// DesktopAppInfoLookupOverrider contains methods that are overridable. This
// interface is a subset of the interface DesktopAppInfoLookup.
type DesktopAppInfoLookupOverrider interface {
	// DefaultForURIScheme gets the default application for launching
	// applications using this URI scheme for a particular AppInfoLookup
	// implementation.
	//
	// The AppInfoLookup interface and this function is used to implement
	// g_app_info_get_default_for_uri_scheme() backends in a GIO module. There
	// is no reason for applications to use it directly. Applications should use
	// g_app_info_get_default_for_uri_scheme().
	DefaultForURIScheme(uriScheme string) AppInfo
}

// DesktopAppInfoLookup is an opaque data structure and can only be accessed
// using the following functions.
type DesktopAppInfoLookup interface {
	gextras.Objector
	DesktopAppInfoLookupOverrider
}

// desktopAppInfoLookup implements the DesktopAppInfoLookup interface.
type desktopAppInfoLookup struct {
	gextras.Objector
}

var _ DesktopAppInfoLookup = (*desktopAppInfoLookup)(nil)

// WrapDesktopAppInfoLookup wraps a GObject to a type that implements interface
// DesktopAppInfoLookup. It is primarily used internally.
func WrapDesktopAppInfoLookup(obj *externglib.Object) DesktopAppInfoLookup {
	return DesktopAppInfoLookup{
		Objector: obj,
	}
}

func marshalDesktopAppInfoLookup(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDesktopAppInfoLookup(obj), nil
}

// DefaultForURIScheme gets the default application for launching
// applications using this URI scheme for a particular AppInfoLookup
// implementation.
//
// The AppInfoLookup interface and this function is used to implement
// g_app_info_get_default_for_uri_scheme() backends in a GIO module. There
// is no reason for applications to use it directly. Applications should use
// g_app_info_get_default_for_uri_scheme().
func (l desktopAppInfoLookup) DefaultForURIScheme(uriScheme string) AppInfo {
	var _arg0 *C.GDesktopAppInfoLookup // out
	var _arg1 *C.char                  // out

	_arg0 = (*C.GDesktopAppInfoLookup)(unsafe.Pointer(l.Native()))
	_arg1 = (*C.char)(C.CString(uriScheme))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.GAppInfo // in

	_cret = C.g_desktop_app_info_lookup_get_default_for_uri_scheme(_arg0, _arg1)

	var _appInfo AppInfo // out

	_appInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(AppInfo)

	return _appInfo
}

// DesktopAppInfo is an implementation of Info based on desktop files.
//
// Note that `<gio/gdesktopappinfo.h>` belongs to the UNIX-specific GIO
// interfaces, thus you have to use the `gio-unix-2.0.pc` pkg-config file when
// using it.
type DesktopAppInfo interface {
	gextras.Objector
	AppInfo

	// ActionName gets the user-visible display name of the "additional
	// application action" specified by @action_name.
	//
	// This corresponds to the "Name" key within the keyfile group for the
	// action.
	ActionName(actionName string) string
	// Boolean looks up a boolean value in the keyfile backing @info.
	//
	// The @key is looked up in the "Desktop Entry" group.
	Boolean(key string) bool
	// Categories gets the categories from the desktop file.
	Categories() string
	// Filename: when @info was created from a known filename, return it. In
	// some situations such as the AppInfo returned from
	// g_desktop_app_info_new_from_keyfile(), this function will return nil.
	Filename() string
	// GenericName gets the generic name from the desktop file.
	GenericName() string
	// IsHidden: a desktop file is hidden if the Hidden key in it is set to
	// True.
	IsHidden() bool
	// Keywords gets the keywords from the desktop file.
	Keywords() []string
	// LocaleString looks up a localized string value in the keyfile backing
	// @info translated to the current locale.
	//
	// The @key is looked up in the "Desktop Entry" group.
	LocaleString(key string) string
	// Nodisplay gets the value of the NoDisplay key, which helps determine if
	// the application info should be shown in menus. See
	// KEY_FILE_DESKTOP_KEY_NO_DISPLAY and g_app_info_should_show().
	Nodisplay() bool
	// ShowIn checks if the application info should be shown in menus that list
	// available applications for a specific name of the desktop, based on the
	// `OnlyShowIn` and `NotShowIn` keys.
	//
	// @desktop_env should typically be given as nil, in which case the
	// `XDG_CURRENT_DESKTOP` environment variable is consulted. If you want to
	// override the default mechanism then you may specify @desktop_env, but
	// this is not recommended.
	//
	// Note that g_app_info_should_show() for @info will include this check
	// (with nil for @desktop_env) as well as additional checks.
	ShowIn(desktopEnv string) bool
	// StartupWmClass retrieves the StartupWMClass field from @info. This
	// represents the WM_CLASS property of the main window of the application,
	// if launched through @info.
	StartupWmClass() string
	// String looks up a string value in the keyfile backing @info.
	//
	// The @key is looked up in the "Desktop Entry" group.
	String(key string) string
	// StringList looks up a string list value in the keyfile backing @info.
	//
	// The @key is looked up in the "Desktop Entry" group.
	StringList(key string) []string
	// HasKey returns whether @key exists in the "Desktop Entry" group of the
	// keyfile backing @info.
	HasKey(key string) bool
	// LaunchAction activates the named application action.
	//
	// You may only call this function on action names that were returned from
	// g_desktop_app_info_list_actions().
	//
	// Note that if the main entry of the desktop file indicates that the
	// application supports startup notification, and @launch_context is
	// non-nil, then startup notification will be used when activating the
	// action (and as such, invocation of the action on the receiving side must
	// signal the end of startup notification when it is completed). This is the
	// expected behaviour of applications declaring additional actions, as per
	// the desktop file specification.
	//
	// As with g_app_info_launch() there is no way to detect failures that occur
	// while using this function.
	LaunchAction(actionName string, launchContext AppLaunchContext)
	// ListActions returns the list of "additional application actions"
	// supported on the desktop file, as per the desktop file specification.
	//
	// As per the specification, this is the list of actions that are explicitly
	// listed in the "Actions" key of the [Desktop Entry] group.
	ListActions() []string
}

// desktopAppInfo implements the DesktopAppInfo class.
type desktopAppInfo struct {
	gextras.Objector
	AppInfo
}

var _ DesktopAppInfo = (*desktopAppInfo)(nil)

// WrapDesktopAppInfo wraps a GObject to the right type. It is
// primarily used internally.
func WrapDesktopAppInfo(obj *externglib.Object) DesktopAppInfo {
	return desktopAppInfo{
		Objector: obj,
		AppInfo:  WrapAppInfo(obj),
	}
}

func marshalDesktopAppInfo(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDesktopAppInfo(obj), nil
}

// NewDesktopAppInfo constructs a class DesktopAppInfo.
func NewDesktopAppInfo(desktopId string) DesktopAppInfo {
	var _arg1 *C.char // out

	_arg1 = (*C.char)(C.CString(desktopId))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GDesktopAppInfo // in

	_cret = C.g_desktop_app_info_new(_arg1)

	var _desktopAppInfo DesktopAppInfo // out

	_desktopAppInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(DesktopAppInfo)

	return _desktopAppInfo
}

// NewDesktopAppInfoFromFilename constructs a class DesktopAppInfo.
func NewDesktopAppInfoFromFilename(filename string) DesktopAppInfo {
	var _arg1 *C.char // out

	_arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GDesktopAppInfo // in

	_cret = C.g_desktop_app_info_new_from_filename(_arg1)

	var _desktopAppInfo DesktopAppInfo // out

	_desktopAppInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(DesktopAppInfo)

	return _desktopAppInfo
}

// NewDesktopAppInfoFromKeyfile constructs a class DesktopAppInfo.
func NewDesktopAppInfoFromKeyfile(keyFile *glib.KeyFile) DesktopAppInfo {
	var _arg1 *C.GKeyFile // out

	_arg1 = (*C.GKeyFile)(unsafe.Pointer(keyFile.Native()))

	var _cret C.GDesktopAppInfo // in

	_cret = C.g_desktop_app_info_new_from_keyfile(_arg1)

	var _desktopAppInfo DesktopAppInfo // out

	_desktopAppInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(DesktopAppInfo)

	return _desktopAppInfo
}

// ActionName gets the user-visible display name of the "additional
// application action" specified by @action_name.
//
// This corresponds to the "Name" key within the keyfile group for the
// action.
func (i desktopAppInfo) ActionName(actionName string) string {
	var _arg0 *C.GDesktopAppInfo // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.gchar // in

	_cret = C.g_desktop_app_info_get_action_name(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Boolean looks up a boolean value in the keyfile backing @info.
//
// The @key is looked up in the "Desktop Entry" group.
func (i desktopAppInfo) Boolean(key string) bool {
	var _arg0 *C.GDesktopAppInfo // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean // in

	_cret = C.g_desktop_app_info_get_boolean(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Categories gets the categories from the desktop file.
func (i desktopAppInfo) Categories() string {
	var _arg0 *C.GDesktopAppInfo // out

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))

	var _cret *C.char // in

	_cret = C.g_desktop_app_info_get_categories(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Filename: when @info was created from a known filename, return it. In
// some situations such as the AppInfo returned from
// g_desktop_app_info_new_from_keyfile(), this function will return nil.
func (i desktopAppInfo) Filename() string {
	var _arg0 *C.GDesktopAppInfo // out

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))

	var _cret *C.char // in

	_cret = C.g_desktop_app_info_get_filename(_arg0)

	var _filename string // out

	_filename = C.GoString(_cret)

	return _filename
}

// GenericName gets the generic name from the desktop file.
func (i desktopAppInfo) GenericName() string {
	var _arg0 *C.GDesktopAppInfo // out

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))

	var _cret *C.char // in

	_cret = C.g_desktop_app_info_get_generic_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// IsHidden: a desktop file is hidden if the Hidden key in it is set to
// True.
func (i desktopAppInfo) IsHidden() bool {
	var _arg0 *C.GDesktopAppInfo // out

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))

	var _cret C.gboolean // in

	_cret = C.g_desktop_app_info_get_is_hidden(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Keywords gets the keywords from the desktop file.
func (i desktopAppInfo) Keywords() []string {
	var _arg0 *C.GDesktopAppInfo // out

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))

	var _cret **C.char

	_cret = C.g_desktop_app_info_get_keywords(_arg0)

	var _utf8s []string

	{
		var i int
		for p := _cret; *p != nil; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
		}
	}

	return _utf8s
}

// LocaleString looks up a localized string value in the keyfile backing
// @info translated to the current locale.
//
// The @key is looked up in the "Desktop Entry" group.
func (i desktopAppInfo) LocaleString(key string) string {
	var _arg0 *C.GDesktopAppInfo // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.char // in

	_cret = C.g_desktop_app_info_get_locale_string(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Nodisplay gets the value of the NoDisplay key, which helps determine if
// the application info should be shown in menus. See
// KEY_FILE_DESKTOP_KEY_NO_DISPLAY and g_app_info_should_show().
func (i desktopAppInfo) Nodisplay() bool {
	var _arg0 *C.GDesktopAppInfo // out

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))

	var _cret C.gboolean // in

	_cret = C.g_desktop_app_info_get_nodisplay(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowIn checks if the application info should be shown in menus that list
// available applications for a specific name of the desktop, based on the
// `OnlyShowIn` and `NotShowIn` keys.
//
// @desktop_env should typically be given as nil, in which case the
// `XDG_CURRENT_DESKTOP` environment variable is consulted. If you want to
// override the default mechanism then you may specify @desktop_env, but
// this is not recommended.
//
// Note that g_app_info_should_show() for @info will include this check
// (with nil for @desktop_env) as well as additional checks.
func (i desktopAppInfo) ShowIn(desktopEnv string) bool {
	var _arg0 *C.GDesktopAppInfo // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.gchar)(C.CString(desktopEnv))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean // in

	_cret = C.g_desktop_app_info_get_show_in(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// StartupWmClass retrieves the StartupWMClass field from @info. This
// represents the WM_CLASS property of the main window of the application,
// if launched through @info.
func (i desktopAppInfo) StartupWmClass() string {
	var _arg0 *C.GDesktopAppInfo // out

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))

	var _cret *C.char // in

	_cret = C.g_desktop_app_info_get_startup_wm_class(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// String looks up a string value in the keyfile backing @info.
//
// The @key is looked up in the "Desktop Entry" group.
func (i desktopAppInfo) String(key string) string {
	var _arg0 *C.GDesktopAppInfo // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.char // in

	_cret = C.g_desktop_app_info_get_string(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// StringList looks up a string list value in the keyfile backing @info.
//
// The @key is looked up in the "Desktop Entry" group.
func (i desktopAppInfo) StringList(key string) []string {
	var _arg0 *C.GDesktopAppInfo // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret **C.gchar
	var _arg2 C.gsize // in

	_cret = C.g_desktop_app_info_get_string_list(_arg0, _arg1, &_arg2)

	var _utf8s []string

	{
		src := unsafe.Slice(_cret, _arg2)
		defer C.free(unsafe.Pointer(_cret))
		_utf8s = make([]string, _arg2)
		for i := 0; i < int(_arg2); i++ {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// HasKey returns whether @key exists in the "Desktop Entry" group of the
// keyfile backing @info.
func (i desktopAppInfo) HasKey(key string) bool {
	var _arg0 *C.GDesktopAppInfo // out
	var _arg1 *C.char            // out

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(key))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean // in

	_cret = C.g_desktop_app_info_has_key(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LaunchAction activates the named application action.
//
// You may only call this function on action names that were returned from
// g_desktop_app_info_list_actions().
//
// Note that if the main entry of the desktop file indicates that the
// application supports startup notification, and @launch_context is
// non-nil, then startup notification will be used when activating the
// action (and as such, invocation of the action on the receiving side must
// signal the end of startup notification when it is completed). This is the
// expected behaviour of applications declaring additional actions, as per
// the desktop file specification.
//
// As with g_app_info_launch() there is no way to detect failures that occur
// while using this function.
func (i desktopAppInfo) LaunchAction(actionName string, launchContext AppLaunchContext) {
	var _arg0 *C.GDesktopAppInfo   // out
	var _arg1 *C.gchar             // out
	var _arg2 *C.GAppLaunchContext // out

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.gchar)(C.CString(actionName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GAppLaunchContext)(unsafe.Pointer(launchContext.Native()))

	C.g_desktop_app_info_launch_action(_arg0, _arg1, _arg2)
}

// ListActions returns the list of "additional application actions"
// supported on the desktop file, as per the desktop file specification.
//
// As per the specification, this is the list of actions that are explicitly
// listed in the "Actions" key of the [Desktop Entry] group.
func (i desktopAppInfo) ListActions() []string {
	var _arg0 *C.GDesktopAppInfo // out

	_arg0 = (*C.GDesktopAppInfo)(unsafe.Pointer(i.Native()))

	var _cret **C.gchar

	_cret = C.g_desktop_app_info_list_actions(_arg0)

	var _utf8s []string

	{
		var i int
		for p := _cret; *p != nil; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
		}
	}

	return _utf8s
}
