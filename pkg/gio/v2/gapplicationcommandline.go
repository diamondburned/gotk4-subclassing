// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
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
		{T: externglib.Type(C.g_application_command_line_get_type()), F: marshalApplicationCommandLine},
	})
}

// ApplicationCommandLineOverrider contains methods that are overridable .
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type ApplicationCommandLineOverrider interface {
	// Stdin gets the stdin of the invoking process.
	//
	// The Stream can be used to read data passed to the standard input of the
	// invoking process. This doesn't work on all platforms. Presently, it is
	// only available on UNIX when using a D-Bus daemon capable of passing file
	// descriptors. If stdin is not available then nil will be returned. In the
	// future, support may be expanded to other platforms.
	//
	// You must only call this function once per commandline invocation.
	Stdin() InputStream
	PrintLiteral(message string)
	PrinterrLiteral(message string)
}

// ApplicationCommandLine represents a command-line invocation of an
// application. It is created by #GApplication and emitted in the
// #GApplication::command-line signal and virtual function.
//
// The class contains the list of arguments that the program was invoked with.
// It is also possible to query if the commandline invocation was local (ie: the
// current process is running in direct response to the invocation) or remote
// (ie: some other process forwarded the commandline to this process).
//
// The GApplicationCommandLine object can provide the @argc and @argv parameters
// for use with the Context command-line parsing API, with the
// g_application_command_line_get_arguments() function. See
// [gapplication-example-cmdline3.c][gapplication-example-cmdline3] for an
// example.
//
// The exit status of the originally-invoked process may be set and messages can
// be printed to stdout or stderr of that process. The lifecycle of the
// originally-invoked process is tied to the lifecycle of this object (ie: the
// process exits when the last reference is dropped).
//
// The main use for CommandLine (and the #GApplication::command-line signal) is
// 'Emacs server' like use cases: You can set the `EDITOR` environment variable
// to have e.g. git use your favourite editor to edit commit messages, and if
// you already have an instance of the editor running, the editing will happen
// in the running instance, instead of opening a new one. An important aspect of
// this use case is that the process that gets started by git does not return
// until the editing is done.
//
// Normally, the commandline is completely handled in the
// #GApplication::command-line handler. The launching instance exits once the
// signal handler in the primary instance has returned, and the return value of
// the signal handler becomes the exit status of the launching instance.
//
//    static gboolean
//    my_cmdline_handler (gpointer data)
//    {
//      GApplicationCommandLine *cmdline = data;
//
//      // do the heavy lifting in an idle
//
//      g_application_command_line_set_exit_status (cmdline, 0);
//      g_object_unref (cmdline); // this releases the application
//
//      return G_SOURCE_REMOVE;
//    }
//
//    static int
//    command_line (GApplication            *application,
//                  GApplicationCommandLine *cmdline)
//    {
//      // keep the application running until we are done with this commandline
//      g_application_hold (application);
//
//      g_object_set_data_full (G_OBJECT (cmdline),
//                              "application", application,
//                              (GDestroyNotify)g_application_release);
//
//      g_object_ref (cmdline);
//      g_idle_add (my_cmdline_handler, cmdline);
//
//      return 0;
//    }
//
// In this example the commandline is not completely handled before the
// #GApplication::command-line handler returns. Instead, we keep a reference to
// the CommandLine object and handle it later (in this example, in an idle).
// Note that it is necessary to hold the application until you are done with the
// commandline.
//
// The complete example can be found here: gapplication-example-cmdline3.c
// (https://git.gnome.org/browse/glib/tree/gio/tests/gapplication-example-cmdline3.c)
type ApplicationCommandLine interface {
	gextras.Objector

	// CreateFileForArg creates a #GFile corresponding to a filename that was
	// given as part of the invocation of @cmdline.
	//
	// This differs from g_file_new_for_commandline_arg() in that it resolves
	// relative pathnames using the current working directory of the invoking
	// process rather than the local process.
	CreateFileForArg(arg string) File
	// Cwd gets the working directory of the command line invocation. The string
	// may contain non-utf8 data.
	//
	// It is possible that the remote application did not send a working
	// directory, so this may be nil.
	//
	// The return value should not be modified or freed and is valid for as long
	// as @cmdline exists.
	Cwd() string
	// Environ gets the contents of the 'environ' variable of the command line
	// invocation, as would be returned by g_get_environ(), ie as a
	// nil-terminated list of strings in the form 'NAME=VALUE'. The strings may
	// contain non-utf8 data.
	//
	// The remote application usually does not send an environment. Use
	// G_APPLICATION_SEND_ENVIRONMENT to affect that. Even with this flag set it
	// is possible that the environment is still not available (due to
	// invocation messages from other applications).
	//
	// The return value should not be modified or freed and is valid for as long
	// as @cmdline exists.
	//
	// See g_application_command_line_getenv() if you are only interested in the
	// value of a single environment variable.
	Environ() []string
	// ExitStatus gets the exit status of @cmdline. See
	// g_application_command_line_set_exit_status() for more information.
	ExitStatus() int
	// IsRemote determines if @cmdline represents a remote invocation.
	IsRemote() bool
	// OptionsDict gets the options there were passed to
	// g_application_command_line().
	//
	// If you did not override local_command_line() then these are the same
	// options that were parsed according to the Entrys added to the application
	// with g_application_add_main_option_entries() and possibly modified from
	// your GApplication::handle-local-options handler.
	//
	// If no options were sent then an empty dictionary is returned so that you
	// don't need to check for nil.
	OptionsDict() *glib.VariantDict
	// PlatformData gets the platform data associated with the invocation of
	// @cmdline.
	//
	// This is a #GVariant dictionary containing information about the context
	// in which the invocation occurred. It typically contains information like
	// the current working directory and the startup notification ID.
	//
	// For local invocation, it will be nil.
	PlatformData() *glib.Variant
	// Stdin gets the stdin of the invoking process.
	//
	// The Stream can be used to read data passed to the standard input of the
	// invoking process. This doesn't work on all platforms. Presently, it is
	// only available on UNIX when using a D-Bus daemon capable of passing file
	// descriptors. If stdin is not available then nil will be returned. In the
	// future, support may be expanded to other platforms.
	//
	// You must only call this function once per commandline invocation.
	Stdin() InputStream
	// Env gets the value of a particular environment variable of the command
	// line invocation, as would be returned by g_getenv(). The strings may
	// contain non-utf8 data.
	//
	// The remote application usually does not send an environment. Use
	// G_APPLICATION_SEND_ENVIRONMENT to affect that. Even with this flag set it
	// is possible that the environment is still not available (due to
	// invocation messages from other applications).
	//
	// The return value should not be modified or freed and is valid for as long
	// as @cmdline exists.
	env(name string) string
	// SetExitStatus sets the exit status that will be used when the invoking
	// process exits.
	//
	// The return value of the #GApplication::command-line signal is passed to
	// this function when the handler returns. This is the usual way of setting
	// the exit status.
	//
	// In the event that you want the remote invocation to continue running and
	// want to decide on the exit status in the future, you can use this call.
	// For the case of a remote invocation, the remote process will typically
	// exit when the last reference is dropped on @cmdline. The exit status of
	// the remote process will be equal to the last value that was set with this
	// function.
	//
	// In the case that the commandline invocation is local, the situation is
	// slightly more complicated. If the commandline invocation results in the
	// mainloop running (ie: because the use-count of the application increased
	// to a non-zero value) then the application is considered to have been
	// 'successful' in a certain sense, and the exit status is always zero. If
	// the application use count is zero, though, the exit status of the local
	// CommandLine is used.
	SetExitStatus(exitStatus int)
}

// applicationCommandLine implements the ApplicationCommandLine interface.
type applicationCommandLine struct {
	*externglib.Object
}

var _ ApplicationCommandLine = (*applicationCommandLine)(nil)

// WrapApplicationCommandLine wraps a GObject to a type that implements
// interface ApplicationCommandLine. It is primarily used internally.
func WrapApplicationCommandLine(obj *externglib.Object) ApplicationCommandLine {
	return applicationCommandLine{obj}
}

func marshalApplicationCommandLine(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapApplicationCommandLine(obj), nil
}

func (c applicationCommandLine) CreateFileForArg(arg string) File {
	var _arg0 *C.GApplicationCommandLine // out
	var _arg1 *C.gchar                   // out
	var _cret *C.GFile                   // in

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(arg))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_application_command_line_create_file_for_arg(_arg0, _arg1)

	var _file File // out

	_file = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(File)

	return _file
}

func (c applicationCommandLine) Cwd() string {
	var _arg0 *C.GApplicationCommandLine // out
	var _cret *C.gchar                   // in

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))

	_cret = C.g_application_command_line_get_cwd(_arg0)

	var _filename string // out

	_filename = C.GoString(_cret)

	return _filename
}

func (c applicationCommandLine) Environ() []string {
	var _arg0 *C.GApplicationCommandLine // out
	var _cret **C.gchar

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))

	_cret = C.g_application_command_line_get_environ(_arg0)

	var _filenames []string

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_filenames = make([]string, i)
		for i := range src {
			_filenames[i] = C.GoString(src[i])
		}
	}

	return _filenames
}

func (c applicationCommandLine) ExitStatus() int {
	var _arg0 *C.GApplicationCommandLine // out
	var _cret C.int                      // in

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))

	_cret = C.g_application_command_line_get_exit_status(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (c applicationCommandLine) IsRemote() bool {
	var _arg0 *C.GApplicationCommandLine // out
	var _cret C.gboolean                 // in

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))

	_cret = C.g_application_command_line_get_is_remote(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (c applicationCommandLine) OptionsDict() *glib.VariantDict {
	var _arg0 *C.GApplicationCommandLine // out
	var _cret *C.GVariantDict            // in

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))

	_cret = C.g_application_command_line_get_options_dict(_arg0)

	var _variantDict *glib.VariantDict // out

	_variantDict = (*glib.VariantDict)(unsafe.Pointer(_cret))
	C.g_variant_dict_ref(_cret)
	runtime.SetFinalizer(_variantDict, func(v *glib.VariantDict) {
		C.g_variant_dict_unref((*C.GVariantDict)(unsafe.Pointer(v)))
	})

	return _variantDict
}

func (c applicationCommandLine) PlatformData() *glib.Variant {
	var _arg0 *C.GApplicationCommandLine // out
	var _cret *C.GVariant                // in

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))

	_cret = C.g_application_command_line_get_platform_data(_arg0)

	var _variant *glib.Variant // out

	_variant = (*glib.Variant)(unsafe.Pointer(_cret))
	C.g_variant_ref(_cret)
	runtime.SetFinalizer(_variant, func(v *glib.Variant) {
		C.g_variant_unref((*C.GVariant)(unsafe.Pointer(v)))
	})

	return _variant
}

func (c applicationCommandLine) Stdin() InputStream {
	var _arg0 *C.GApplicationCommandLine // out
	var _cret *C.GInputStream            // in

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))

	_cret = C.g_application_command_line_get_stdin(_arg0)

	var _inputStream InputStream // out

	_inputStream = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(InputStream)

	return _inputStream
}

func (c applicationCommandLine) env(name string) string {
	var _arg0 *C.GApplicationCommandLine // out
	var _arg1 *C.gchar                   // out
	var _cret *C.gchar                   // in

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_application_command_line_getenv(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (c applicationCommandLine) SetExitStatus(exitStatus int) {
	var _arg0 *C.GApplicationCommandLine // out
	var _arg1 C.int                      // out

	_arg0 = (*C.GApplicationCommandLine)(unsafe.Pointer(c.Native()))
	_arg1 = C.int(exitStatus)

	C.g_application_command_line_set_exit_status(_arg0, _arg1)
}
