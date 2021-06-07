// Code generated by girgen. DO NOT EDIT.

package gio

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
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
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_app_info_get_type()), F: marshalAppInfo},
		{T: externglib.Type(C.g_app_launch_context_get_type()), F: marshalAppLaunchContext},
	})
}

// AppInfoCreateFromCommandline creates a new Info from the given information.
//
// Note that for @commandline, the quoting rules of the Exec key of the
// freedesktop.org Desktop Entry Specification
// (http://freedesktop.org/Standards/desktop-entry-spec) are applied. For
// example, if the @commandline contains percent-encoded URIs, the
// percent-character must be doubled in order to prevent it from being swallowed
// by Exec key unquoting. See the specification for exact quoting rules.
func AppInfoCreateFromCommandline(commandline string, applicationName string, flags AppInfoCreateFlags) error {
	var arg1 *C.char
	var arg2 *C.char
	var arg3 C.GAppInfoCreateFlags

	arg1 = (*C.char)(C.CString(commandline))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.char)(C.CString(applicationName))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = (C.GAppInfoCreateFlags)(flags)

	var errout *C.GError
	var err error

	C.g_app_info_create_from_commandline(arg1, arg2, arg3, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// AppInfoGetAll gets a list of all of the applications currently registered on
// this system.
//
// For desktop files, this includes applications that have `NoDisplay=true` set
// or are excluded from display by means of `OnlyShowIn` or `NotShowIn`. See
// g_app_info_should_show(). The returned list does not include applications
// which have the `Hidden` key set.
func AppInfoGetAll() {
	C.g_app_info_get_all()
}

// AppInfoGetAllForType gets a list of all Infos for a given content type,
// including the recommended and fallback Infos. See
// g_app_info_get_recommended_for_type() and g_app_info_get_fallback_for_type().
func AppInfoGetAllForType(contentType string) {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(arg1))

	C.g_app_info_get_all_for_type(arg1)
}

// AppInfoGetDefaultForType gets the default Info for a given content type.
func AppInfoGetDefaultForType(contentType string, mustSupportUris bool) {
	var arg1 *C.char
	var arg2 C.gboolean

	arg1 = (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(arg1))
	if mustSupportUris {
		arg2 = C.gboolean(1)
	}

	C.g_app_info_get_default_for_type(arg1, arg2)
}

// AppInfoGetDefaultForURIScheme gets the default application for handling URIs
// with the given URI scheme. A URI scheme is the initial part of the URI, up to
// but not including the ':', e.g. "http", "ftp" or "sip".
func AppInfoGetDefaultForURIScheme(uriScheme string) {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(uriScheme))
	defer C.free(unsafe.Pointer(arg1))

	C.g_app_info_get_default_for_uri_scheme(arg1)
}

// AppInfoGetFallbackForType gets a list of fallback Infos for a given content
// type, i.e. those applications which claim to support the given content type
// by MIME type subclassing and not directly.
func AppInfoGetFallbackForType(contentType string) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(contentType))
	defer C.free(unsafe.Pointer(arg1))

	C.g_app_info_get_fallback_for_type(arg1)
}

// AppInfoGetRecommendedForType gets a list of recommended Infos for a given
// content type, i.e. those applications which claim to support the given
// content type exactly, and not by MIME type subclassing. Note that the first
// application of the list is the last used one, i.e. the last one for which
// g_app_info_set_as_last_used_for_type() has been called.
func AppInfoGetRecommendedForType(contentType string) {
	var arg1 *C.gchar

	arg1 = (*C.gchar)(C.CString(contentType))
	defer C.free(unsafe.Pointer(arg1))

	C.g_app_info_get_recommended_for_type(arg1)
}

// AppInfoLaunchDefaultForURI: utility function that launches the default
// application registered to handle the specified uri. Synchronous I/O is done
// on the uri to detect the type of the file if required.
//
// The D-Bus–activated applications don't have to be started if your application
// terminates too soon after this function. To prevent this, use
// g_app_info_launch_default_for_uri_async() instead.
func AppInfoLaunchDefaultForURI(uri string, context AppLaunchContext) error {
	var arg1 *C.char
	var arg2 *C.GAppLaunchContext

	arg1 = (*C.char)(C.CString(uri))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GAppLaunchContext)(unsafe.Pointer(context.Native()))

	var errout *C.GError
	var err error

	C.g_app_info_launch_default_for_uri(arg1, arg2, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// AppInfoLaunchDefaultForURIAsync: async version of
// g_app_info_launch_default_for_uri().
//
// This version is useful if you are interested in receiving error information
// in the case where the application is sandboxed and the portal may present an
// application chooser dialog to the user.
//
// This is also useful if you want to be sure that the D-Bus–activated
// applications are really started before termination and if you are interested
// in receiving error information from their activation.
func AppInfoLaunchDefaultForURIAsync() {
	C.g_app_info_launch_default_for_uri_async(arg1, arg2, arg3, arg4, arg5)
}

// AppInfoLaunchDefaultForURIFinish finishes an asynchronous
// launch-default-for-uri operation.
func AppInfoLaunchDefaultForURIFinish(result AsyncResult) error {
	var arg1 *C.GAsyncResult

	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var errout *C.GError
	var err error

	C.g_app_info_launch_default_for_uri_finish(arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// AppInfoResetTypeAssociations removes all changes to the type associations
// done by g_app_info_set_as_default_for_type(),
// g_app_info_set_as_default_for_extension(), g_app_info_add_supports_type() or
// g_app_info_remove_supports_type().
func AppInfoResetTypeAssociations(contentType string) {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(arg1))

	C.g_app_info_reset_type_associations(arg1)
}

// AppInfoOverrider contains methods that are overridable. This
// interface is a subset of the interface AppInfo.
type AppInfoOverrider interface {
	// AddSupportsType adds a content type to the application information to
	// indicate the application is capable of opening files with the given
	// content type.
	AddSupportsType(a AppInfo, contentType string) error
	// CanDelete obtains the information whether the Info can be deleted. See
	// g_app_info_delete().
	CanDelete(a AppInfo) bool
	// CanRemoveSupportsType checks if a supported content type can be removed
	// from an application.
	CanRemoveSupportsType(a AppInfo) bool
	// DoDelete tries to delete a Info.
	//
	// On some platforms, there may be a difference between user-defined Infos
	// which can be deleted, and system-wide ones which cannot. See
	// g_app_info_can_delete().
	DoDelete(a AppInfo) bool
	// Dup creates a duplicate of a Info.
	Dup(a AppInfo)
	// Equal checks if two Infos are equal.
	//
	// Note that the check <emphasis>may not</emphasis> compare each individual
	// field, and only does an identity check. In case detecting changes in the
	// contents is needed, program code must additionally compare relevant
	// fields.
	Equal(a AppInfo, appinfo2 AppInfo) bool

	Commandline(a AppInfo)
	// Description gets a human-readable description of an installed
	// application.
	Description(a AppInfo)
	// DisplayName gets the display name of the application. The display name is
	// often more descriptive to the user than the name itself.
	DisplayName(a AppInfo)

	Executable(a AppInfo)
	// Icon gets the icon for the application.
	Icon(a AppInfo)
	// ID gets the ID of an application. An id is a string that identifies the
	// application. The exact format of the id is platform dependent. For
	// instance, on Unix this is the desktop file id from the xdg menu
	// specification.
	//
	// Note that the returned ID may be nil, depending on how the @appinfo has
	// been constructed.
	ID(a AppInfo)
	// Name gets the installed name of the application.
	Name(a AppInfo)
	// SupportedTypes retrieves the list of content types that @app_info claims
	// to support. If this information is not provided by the environment, this
	// function will return nil. This function does not take in consideration
	// associations added with g_app_info_add_supports_type(), but only those
	// exported directly by the application.
	SupportedTypes(a AppInfo)
	// Launch launches the application. Passes @files to the launched
	// application as arguments, using the optional @context to get information
	// about the details of the launcher (like what screen it is on). On error,
	// @error will be set accordingly.
	//
	// To launch the application without arguments pass a nil @files list.
	//
	// Note that even if the launch is successful the application launched can
	// fail to start if it runs into problems during startup. There is no way to
	// detect this.
	//
	// Some URIs can be changed when passed through a GFile (for instance
	// unsupported URIs with strange formats like mailto:), so if you have a
	// textual URI you want to pass in as argument, consider using
	// g_app_info_launch_uris() instead.
	//
	// The launched application inherits the environment of the launching
	// process, but it can be modified with g_app_launch_context_setenv() and
	// g_app_launch_context_unsetenv().
	//
	// On UNIX, this function sets the `GIO_LAUNCHED_DESKTOP_FILE` environment
	// variable with the path of the launched desktop file and
	// `GIO_LAUNCHED_DESKTOP_FILE_PID` to the process id of the launched
	// process. This can be used to ignore `GIO_LAUNCHED_DESKTOP_FILE`, should
	// it be inherited by further processes. The `DISPLAY` and
	// `DESKTOP_STARTUP_ID` environment variables are also set, based on
	// information provided in @context.
	Launch(a AppInfo, files *glib.List, context AppLaunchContext) error
	// LaunchUris launches the application. This passes the @uris to the
	// launched application as arguments, using the optional @context to get
	// information about the details of the launcher (like what screen it is
	// on). On error, @error will be set accordingly.
	//
	// To launch the application without arguments pass a nil @uris list.
	//
	// Note that even if the launch is successful the application launched can
	// fail to start if it runs into problems during startup. There is no way to
	// detect this.
	LaunchUris(a AppInfo, uris *glib.List, context AppLaunchContext) error
	// LaunchUrisAsync: async version of g_app_info_launch_uris().
	//
	// The @callback is invoked immediately after the application launch, but it
	// waits for activation in case of D-Bus–activated applications and also
	// provides extended error information for sandboxed applications, see notes
	// for g_app_info_launch_default_for_uri_async().
	LaunchUrisAsync(a AppInfo)
	// LaunchUrisFinish finishes a g_app_info_launch_uris_async() operation.
	LaunchUrisFinish(a AppInfo, result AsyncResult) error
	// RemoveSupportsType removes a supported type from an application, if
	// possible.
	RemoveSupportsType(a AppInfo, contentType string) error
	// SetAsDefaultForExtension sets the application as the default handler for
	// the given file extension.
	SetAsDefaultForExtension(a AppInfo, extension string) error
	// SetAsDefaultForType sets the application as the default handler for a
	// given type.
	SetAsDefaultForType(a AppInfo, contentType string) error
	// SetAsLastUsedForType sets the application as the last used application
	// for a given type. This will make the application appear as first in the
	// list returned by g_app_info_get_recommended_for_type(), regardless of the
	// default application for that content type.
	SetAsLastUsedForType(a AppInfo, contentType string) error
	// ShouldShow checks if the application info should be shown in menus that
	// list available applications.
	ShouldShow(a AppInfo) bool
	// SupportsFiles checks if the application accepts files as arguments.
	SupportsFiles(a AppInfo) bool
	// SupportsUris checks if the application supports reading files and
	// directories from URIs.
	SupportsUris(a AppInfo) bool
}

// AppInfo: Info and LaunchContext are used for describing and launching
// applications installed on the system.
//
// As of GLib 2.20, URIs will always be converted to POSIX paths (using
// g_file_get_path()) when using g_app_info_launch() even if the application
// requested an URI and not a POSIX path. For example for a desktop-file based
// application with Exec key `totem U` and a single URI, `sftp://foo/file.avi`,
// then `/home/user/.gvfs/sftp on foo/file.avi` will be passed. This will only
// work if a set of suitable GIO extensions (such as gvfs 2.26 compiled with
// FUSE support), is available and operational; if this is not the case, the URI
// will be passed unmodified to the application. Some URIs, such as `mailto:`,
// of course cannot be mapped to a POSIX path (in gvfs there's no FUSE mount for
// it); such URIs will be passed unmodified to the application.
//
// Specifically for gvfs 2.26 and later, the POSIX URI will be mapped back to
// the GIO URI in the #GFile constructors (since gvfs implements the #GVfs
// extension point). As such, if the application needs to examine the URI, it
// needs to use g_file_get_uri() or similar on #GFile. In other words, an
// application cannot assume that the URI passed to e.g.
// g_file_new_for_commandline_arg() is equal to the result of g_file_get_uri().
// The following snippet illustrates this:
//
//    GFile *f;
//    char *uri;
//
//    file = g_file_new_for_commandline_arg (uri_from_commandline);
//
//    uri = g_file_get_uri (file);
//    strcmp (uri, uri_from_commandline) == 0;
//    g_free (uri);
//
//    if (g_file_has_uri_scheme (file, "cdda"))
//      {
//        // do something special with uri
//      }
//    g_object_unref (file);
//
// This code will work when both `cdda://sr0/Track 1.wav` and
// `/home/user/.gvfs/cdda on sr0/Track 1.wav` is passed to the application. It
// should be noted that it's generally not safe for applications to rely on the
// format of a particular URIs. Different launcher applications (e.g. file
// managers) may have different ideas of what a given URI means.
type AppInfo interface {
	gextras.Objector
	AppInfoOverrider

	// Delete tries to delete a Info.
	//
	// On some platforms, there may be a difference between user-defined Infos
	// which can be deleted, and system-wide ones which cannot. See
	// g_app_info_can_delete().
	Delete(a AppInfo) bool
}

// appInfo implements the AppInfo interface.
type appInfo struct {
	gextras.Objector
}

var _ AppInfo = (*appInfo)(nil)

// WrapAppInfo wraps a GObject to a type that implements interface
// AppInfo. It is primarily used internally.
func WrapAppInfo(obj *externglib.Object) AppInfo {
	return AppInfo{
		Objector: obj,
	}
}

func marshalAppInfo(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAppInfo(obj), nil
}

// AddSupportsType adds a content type to the application information to
// indicate the application is capable of opening files with the given
// content type.
func (a appInfo) AddSupportsType(a AppInfo, contentType string) error {
	var arg0 *C.GAppInfo
	var arg1 *C.char

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))
	arg1 = (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(arg1))

	var errout *C.GError
	var err error

	C.g_app_info_add_supports_type(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// CanDelete obtains the information whether the Info can be deleted. See
// g_app_info_delete().
func (a appInfo) CanDelete(a AppInfo) bool {
	var arg0 *C.GAppInfo

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_app_info_can_delete(arg0)

	if cret {
		ok = true
	}

	return ok
}

// CanRemoveSupportsType checks if a supported content type can be removed
// from an application.
func (a appInfo) CanRemoveSupportsType(a AppInfo) bool {
	var arg0 *C.GAppInfo

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_app_info_can_remove_supports_type(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Delete tries to delete a Info.
//
// On some platforms, there may be a difference between user-defined Infos
// which can be deleted, and system-wide ones which cannot. See
// g_app_info_can_delete().
func (a appInfo) Delete(a AppInfo) bool {
	var arg0 *C.GAppInfo

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_app_info_delete(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Dup creates a duplicate of a Info.
func (a appInfo) Dup(a AppInfo) {
	var arg0 *C.GAppInfo

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	C.g_app_info_dup(arg0)
}

// Equal checks if two Infos are equal.
//
// Note that the check <emphasis>may not</emphasis> compare each individual
// field, and only does an identity check. In case detecting changes in the
// contents is needed, program code must additionally compare relevant
// fields.
func (a appInfo) Equal(a AppInfo, appinfo2 AppInfo) bool {
	var arg0 *C.GAppInfo
	var arg1 *C.GAppInfo

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GAppInfo)(unsafe.Pointer(appinfo2.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_app_info_equal(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// Commandline gets the commandline with which the application will be
// started.
func (a appInfo) Commandline(a AppInfo) {
	var arg0 *C.GAppInfo

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	C.g_app_info_get_commandline(arg0)
}

// Description gets a human-readable description of an installed
// application.
func (a appInfo) Description(a AppInfo) {
	var arg0 *C.GAppInfo

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	C.g_app_info_get_description(arg0)
}

// DisplayName gets the display name of the application. The display name is
// often more descriptive to the user than the name itself.
func (a appInfo) DisplayName(a AppInfo) {
	var arg0 *C.GAppInfo

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	C.g_app_info_get_display_name(arg0)
}

// Executable gets the executable's name for the installed application.
func (a appInfo) Executable(a AppInfo) {
	var arg0 *C.GAppInfo

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	C.g_app_info_get_executable(arg0)
}

// Icon gets the icon for the application.
func (a appInfo) Icon(a AppInfo) {
	var arg0 *C.GAppInfo

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	C.g_app_info_get_icon(arg0)
}

// ID gets the ID of an application. An id is a string that identifies the
// application. The exact format of the id is platform dependent. For
// instance, on Unix this is the desktop file id from the xdg menu
// specification.
//
// Note that the returned ID may be nil, depending on how the @appinfo has
// been constructed.
func (a appInfo) ID(a AppInfo) {
	var arg0 *C.GAppInfo

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	C.g_app_info_get_id(arg0)
}

// Name gets the installed name of the application.
func (a appInfo) Name(a AppInfo) {
	var arg0 *C.GAppInfo

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	C.g_app_info_get_name(arg0)
}

// SupportedTypes retrieves the list of content types that @app_info claims
// to support. If this information is not provided by the environment, this
// function will return nil. This function does not take in consideration
// associations added with g_app_info_add_supports_type(), but only those
// exported directly by the application.
func (a appInfo) SupportedTypes(a AppInfo) {
	var arg0 *C.GAppInfo

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	C.g_app_info_get_supported_types(arg0)
}

// Launch launches the application. Passes @files to the launched
// application as arguments, using the optional @context to get information
// about the details of the launcher (like what screen it is on). On error,
// @error will be set accordingly.
//
// To launch the application without arguments pass a nil @files list.
//
// Note that even if the launch is successful the application launched can
// fail to start if it runs into problems during startup. There is no way to
// detect this.
//
// Some URIs can be changed when passed through a GFile (for instance
// unsupported URIs with strange formats like mailto:), so if you have a
// textual URI you want to pass in as argument, consider using
// g_app_info_launch_uris() instead.
//
// The launched application inherits the environment of the launching
// process, but it can be modified with g_app_launch_context_setenv() and
// g_app_launch_context_unsetenv().
//
// On UNIX, this function sets the `GIO_LAUNCHED_DESKTOP_FILE` environment
// variable with the path of the launched desktop file and
// `GIO_LAUNCHED_DESKTOP_FILE_PID` to the process id of the launched
// process. This can be used to ignore `GIO_LAUNCHED_DESKTOP_FILE`, should
// it be inherited by further processes. The `DISPLAY` and
// `DESKTOP_STARTUP_ID` environment variables are also set, based on
// information provided in @context.
func (a appInfo) Launch(a AppInfo, files *glib.List, context AppLaunchContext) error {
	var arg0 *C.GAppInfo
	var arg1 *C.GList
	var arg2 *C.GAppLaunchContext

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GList)(unsafe.Pointer(files.Native()))
	arg2 = (*C.GAppLaunchContext)(unsafe.Pointer(context.Native()))

	var errout *C.GError
	var err error

	C.g_app_info_launch(arg0, arg1, arg2, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// LaunchUris launches the application. This passes the @uris to the
// launched application as arguments, using the optional @context to get
// information about the details of the launcher (like what screen it is
// on). On error, @error will be set accordingly.
//
// To launch the application without arguments pass a nil @uris list.
//
// Note that even if the launch is successful the application launched can
// fail to start if it runs into problems during startup. There is no way to
// detect this.
func (a appInfo) LaunchUris(a AppInfo, uris *glib.List, context AppLaunchContext) error {
	var arg0 *C.GAppInfo
	var arg1 *C.GList
	var arg2 *C.GAppLaunchContext

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GList)(unsafe.Pointer(uris.Native()))
	arg2 = (*C.GAppLaunchContext)(unsafe.Pointer(context.Native()))

	var errout *C.GError
	var err error

	C.g_app_info_launch_uris(arg0, arg1, arg2, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// LaunchUrisAsync: async version of g_app_info_launch_uris().
//
// The @callback is invoked immediately after the application launch, but it
// waits for activation in case of D-Bus–activated applications and also
// provides extended error information for sandboxed applications, see notes
// for g_app_info_launch_default_for_uri_async().
func (a appInfo) LaunchUrisAsync(a AppInfo) {
	var arg0 *C.GAppInfo

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	C.g_app_info_launch_uris_async(arg0, arg1, arg2, arg3, arg4, arg5)
}

// LaunchUrisFinish finishes a g_app_info_launch_uris_async() operation.
func (a appInfo) LaunchUrisFinish(a AppInfo, result AsyncResult) error {
	var arg0 *C.GAppInfo
	var arg1 *C.GAsyncResult

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))

	var errout *C.GError
	var err error

	C.g_app_info_launch_uris_finish(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// RemoveSupportsType removes a supported type from an application, if
// possible.
func (a appInfo) RemoveSupportsType(a AppInfo, contentType string) error {
	var arg0 *C.GAppInfo
	var arg1 *C.char

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))
	arg1 = (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(arg1))

	var errout *C.GError
	var err error

	C.g_app_info_remove_supports_type(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// SetAsDefaultForExtension sets the application as the default handler for
// the given file extension.
func (a appInfo) SetAsDefaultForExtension(a AppInfo, extension string) error {
	var arg0 *C.GAppInfo
	var arg1 *C.char

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))
	arg1 = (*C.char)(C.CString(extension))
	defer C.free(unsafe.Pointer(arg1))

	var errout *C.GError
	var err error

	C.g_app_info_set_as_default_for_extension(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// SetAsDefaultForType sets the application as the default handler for a
// given type.
func (a appInfo) SetAsDefaultForType(a AppInfo, contentType string) error {
	var arg0 *C.GAppInfo
	var arg1 *C.char

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))
	arg1 = (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(arg1))

	var errout *C.GError
	var err error

	C.g_app_info_set_as_default_for_type(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// SetAsLastUsedForType sets the application as the last used application
// for a given type. This will make the application appear as first in the
// list returned by g_app_info_get_recommended_for_type(), regardless of the
// default application for that content type.
func (a appInfo) SetAsLastUsedForType(a AppInfo, contentType string) error {
	var arg0 *C.GAppInfo
	var arg1 *C.char

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))
	arg1 = (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(arg1))

	var errout *C.GError
	var err error

	C.g_app_info_set_as_last_used_for_type(arg0, arg1, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// ShouldShow checks if the application info should be shown in menus that
// list available applications.
func (a appInfo) ShouldShow(a AppInfo) bool {
	var arg0 *C.GAppInfo

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_app_info_should_show(arg0)

	if cret {
		ok = true
	}

	return ok
}

// SupportsFiles checks if the application accepts files as arguments.
func (a appInfo) SupportsFiles(a AppInfo) bool {
	var arg0 *C.GAppInfo

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_app_info_supports_files(arg0)

	if cret {
		ok = true
	}

	return ok
}

// SupportsUris checks if the application supports reading files and
// directories from URIs.
func (a appInfo) SupportsUris(a AppInfo) bool {
	var arg0 *C.GAppInfo

	arg0 = (*C.GAppInfo)(unsafe.Pointer(a.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.g_app_info_supports_uris(arg0)

	if cret {
		ok = true
	}

	return ok
}

// AppLaunchContext: integrating the launch with the launching application. This
// is used to handle for instance startup notification and launching the new
// application on the same screen as the launching window.
type AppLaunchContext interface {
	gextras.Objector

	// Display gets the display string for the @context. This is used to ensure
	// new applications are started on the same display as the launching
	// application, by setting the `DISPLAY` environment variable.
	Display(c AppLaunchContext, info AppInfo, files *glib.List)
	// Environment gets the complete environment variable list to be passed to
	// the child process when @context is used to launch an application. This is
	// a nil-terminated array of strings, where each string has the form
	// `KEY=VALUE`.
	Environment(c AppLaunchContext)
	// StartupNotifyID initiates startup notification for the application and
	// returns the `DESKTOP_STARTUP_ID` for the launched operation, if
	// supported.
	//
	// Startup notification IDs are defined in the FreeDesktop.Org Startup
	// Notifications standard
	// (http://standards.freedesktop.org/startup-notification-spec/startup-notification-latest.txt).
	StartupNotifyID(c AppLaunchContext, info AppInfo, files *glib.List)
	// LaunchFailed: called when an application has failed to launch, so that it
	// can cancel the application startup notification started in
	// g_app_launch_context_get_startup_notify_id().
	LaunchFailed(c AppLaunchContext, startupNotifyID string)
	// Setenv arranges for @variable to be set to @value in the child's
	// environment when @context is used to launch an application.
	Setenv(c AppLaunchContext, variable string, value string)
	// Unsetenv arranges for @variable to be unset in the child's environment
	// when @context is used to launch an application.
	Unsetenv(c AppLaunchContext, variable string)
}

// appLaunchContext implements the AppLaunchContext interface.
type appLaunchContext struct {
	gextras.Objector
}

var _ AppLaunchContext = (*appLaunchContext)(nil)

// WrapAppLaunchContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapAppLaunchContext(obj *externglib.Object) AppLaunchContext {
	return AppLaunchContext{
		Objector: obj,
	}
}

func marshalAppLaunchContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAppLaunchContext(obj), nil
}

// NewAppLaunchContext constructs a class AppLaunchContext.
func NewAppLaunchContext() {
	C.g_app_launch_context_new()
}

// Display gets the display string for the @context. This is used to ensure
// new applications are started on the same display as the launching
// application, by setting the `DISPLAY` environment variable.
func (c appLaunchContext) Display(c AppLaunchContext, info AppInfo, files *glib.List) {
	var arg0 *C.GAppLaunchContext
	var arg1 *C.GAppInfo
	var arg2 *C.GList

	arg0 = (*C.GAppLaunchContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GAppInfo)(unsafe.Pointer(info.Native()))
	arg2 = (*C.GList)(unsafe.Pointer(files.Native()))

	C.g_app_launch_context_get_display(arg0, arg1, arg2)
}

// Environment gets the complete environment variable list to be passed to
// the child process when @context is used to launch an application. This is
// a nil-terminated array of strings, where each string has the form
// `KEY=VALUE`.
func (c appLaunchContext) Environment(c AppLaunchContext) {
	var arg0 *C.GAppLaunchContext

	arg0 = (*C.GAppLaunchContext)(unsafe.Pointer(c.Native()))

	C.g_app_launch_context_get_environment(arg0)
}

// StartupNotifyID initiates startup notification for the application and
// returns the `DESKTOP_STARTUP_ID` for the launched operation, if
// supported.
//
// Startup notification IDs are defined in the FreeDesktop.Org Startup
// Notifications standard
// (http://standards.freedesktop.org/startup-notification-spec/startup-notification-latest.txt).
func (c appLaunchContext) StartupNotifyID(c AppLaunchContext, info AppInfo, files *glib.List) {
	var arg0 *C.GAppLaunchContext
	var arg1 *C.GAppInfo
	var arg2 *C.GList

	arg0 = (*C.GAppLaunchContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GAppInfo)(unsafe.Pointer(info.Native()))
	arg2 = (*C.GList)(unsafe.Pointer(files.Native()))

	C.g_app_launch_context_get_startup_notify_id(arg0, arg1, arg2)
}

// LaunchFailed: called when an application has failed to launch, so that it
// can cancel the application startup notification started in
// g_app_launch_context_get_startup_notify_id().
func (c appLaunchContext) LaunchFailed(c AppLaunchContext, startupNotifyID string) {
	var arg0 *C.GAppLaunchContext
	var arg1 *C.char

	arg0 = (*C.GAppLaunchContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.char)(C.CString(startupNotifyID))
	defer C.free(unsafe.Pointer(arg1))

	C.g_app_launch_context_launch_failed(arg0, arg1)
}

// Setenv arranges for @variable to be set to @value in the child's
// environment when @context is used to launch an application.
func (c appLaunchContext) Setenv(c AppLaunchContext, variable string, value string) {
	var arg0 *C.GAppLaunchContext
	var arg1 *C.char
	var arg2 *C.char

	arg0 = (*C.GAppLaunchContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.char)(C.CString(variable))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.char)(C.CString(value))
	defer C.free(unsafe.Pointer(arg2))

	C.g_app_launch_context_setenv(arg0, arg1, arg2)
}

// Unsetenv arranges for @variable to be unset in the child's environment
// when @context is used to launch an application.
func (c appLaunchContext) Unsetenv(c AppLaunchContext, variable string) {
	var arg0 *C.GAppLaunchContext
	var arg1 *C.char

	arg0 = (*C.GAppLaunchContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.char)(C.CString(variable))
	defer C.free(unsafe.Pointer(arg1))

	C.g_app_launch_context_unsetenv(arg0, arg1)
}