// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
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
		{T: externglib.Type(C.g_file_attribute_matcher_get_type()), F: marshalFileAttributeMatcher},
		{T: externglib.Type(C.g_resource_get_type()), F: marshalResource},
		{T: externglib.Type(C.g_srv_target_get_type()), F: marshalSrvTarget},
	})
}

// AsyncReadyCallback: type definition for a function that will be called back
// when an asynchronous operation within GIO has been completed. ReadyCallback
// callbacks from #GTask are guaranteed to be invoked in a later iteration of
// the [thread-default main context][g-main-context-push-thread-default] where
// the #GTask was created. All other users of ReadyCallback must likewise call
// it asynchronously in a later iteration of the main context.
type AsyncReadyCallback func(sourceObject gextras.Objector, res AsyncResult)

//export gotk4_AsyncReadyCallback
func gotk4_AsyncReadyCallback(arg0 *C.GObject, arg1 *C.GAsyncResult, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var sourceObject gextras.Objector // out
	var res AsyncResult               // out

	sourceObject = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(gextras.Objector)
	res = gextras.CastObject(externglib.Take(unsafe.Pointer(arg1))).(AsyncResult)

	fn := v.(AsyncReadyCallback)
	fn(sourceObject, res)
}

// CancellableSourceFunc: this is the function type of the callback used for the
// #GSource returned by g_cancellable_source_new().
type CancellableSourceFunc func(cancellable Cancellable) (ok bool)

//export gotk4_CancellableSourceFunc
func gotk4_CancellableSourceFunc(arg0 *C.GCancellable, arg1 C.gpointer) (cret C.gboolean) {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var cancellable Cancellable // out

	cancellable = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(Cancellable)

	fn := v.(CancellableSourceFunc)
	ok := fn(cancellable)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// DBusProxyTypeFunc: function signature for a function used to determine the
// #GType to use for an interface proxy (if @interface_name is not nil) or
// object proxy (if @interface_name is nil).
//
// This function is called in the [thread-default main
// loop][g-main-context-push-thread-default] that @manager was constructed in.
type DBusProxyTypeFunc func(manager DBusObjectManagerClient, objectPath string, interfaceName string) (gType externglib.Type)

//export gotk4_DBusProxyTypeFunc
func gotk4_DBusProxyTypeFunc(arg0 *C.GDBusObjectManagerClient, arg1 *C.gchar, arg2 *C.gchar, arg3 C.gpointer) (cret C.GType) {
	v := box.Get(uintptr(arg3))
	if v == nil {
		panic(`callback not found`)
	}

	var manager DBusObjectManagerClient // out
	var objectPath string               // out
	var interfaceName string            // out

	manager = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(DBusObjectManagerClient)
	objectPath = C.GoString(arg1)
	interfaceName = C.GoString(arg2)

	fn := v.(DBusProxyTypeFunc)
	gType := fn(manager, objectPath, interfaceName)

	cret = (C.GType)(gType)

	return cret
}

// DatagramBasedSourceFunc: this is the function type of the callback used for
// the #GSource returned by g_datagram_based_create_source().
type DatagramBasedSourceFunc func(datagramBased DatagramBased, condition glib.IOCondition) (ok bool)

//export gotk4_DatagramBasedSourceFunc
func gotk4_DatagramBasedSourceFunc(arg0 *C.GDatagramBased, arg1 C.GIOCondition, arg2 C.gpointer) (cret C.gboolean) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var datagramBased DatagramBased // out
	var condition glib.IOCondition  // out

	datagramBased = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(DatagramBased)
	condition = glib.IOCondition(arg1)

	fn := v.(DatagramBasedSourceFunc)
	ok := fn(datagramBased, condition)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// FileMeasureProgressCallback: this callback type is used by
// g_file_measure_disk_usage() to make periodic progress reports when measuring
// the amount of disk spaced used by a directory.
//
// These calls are made on a best-effort basis and not all types of #GFile will
// support them. At the minimum, however, one call will always be made
// immediately.
//
// In the case that there is no support, @reporting will be set to false (and
// the other values undefined) and no further calls will be made. Otherwise, the
// @reporting will be true and the other values all-zeros during the first
// (immediate) call. In this way, you can know which type of progress UI to show
// without a delay.
//
// For g_file_measure_disk_usage() the callback is made directly. For
// g_file_measure_disk_usage_async() the callback is made via the default main
// context of the calling thread (ie: the same way that the final async result
// would be reported).
//
// @current_size is in the same units as requested by the operation (see
// G_FILE_MEASURE_APPARENT_SIZE).
//
// The frequency of the updates is implementation defined, but is ideally about
// once every 200ms.
//
// The last progress callback may or may not be equal to the final result.
// Always check the async result to get the final value.
type FileMeasureProgressCallback func(reporting bool, currentSize uint64, numDirs uint64, numFiles uint64)

//export gotk4_FileMeasureProgressCallback
func gotk4_FileMeasureProgressCallback(arg0 C.gboolean, arg1 C.guint64, arg2 C.guint64, arg3 C.guint64, arg4 C.gpointer) {
	v := box.Get(uintptr(arg4))
	if v == nil {
		panic(`callback not found`)
	}

	var reporting bool     // out
	var currentSize uint64 // out
	var numDirs uint64     // out
	var numFiles uint64    // out

	if arg0 != 0 {
		reporting = true
	}
	currentSize = uint64(arg1)
	numDirs = uint64(arg2)
	numFiles = uint64(arg3)

	fn := v.(FileMeasureProgressCallback)
	fn(reporting, currentSize, numDirs, numFiles)
}

// FileProgressCallback: when doing file operations that may take a while, such
// as moving a file or copying a file, a progress callback is used to pass how
// far along that operation is to the application.
type FileProgressCallback func(currentNumBytes int64, totalNumBytes int64)

//export gotk4_FileProgressCallback
func gotk4_FileProgressCallback(arg0 C.goffset, arg1 C.goffset, arg2 C.gpointer) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var currentNumBytes int64 // out
	var totalNumBytes int64   // out

	currentNumBytes = int64(arg0)
	totalNumBytes = int64(arg1)

	fn := v.(FileProgressCallback)
	fn(currentNumBytes, totalNumBytes)
}

// FileReadMoreCallback: when loading the partial contents of a file with
// g_file_load_partial_contents_async(), it may become necessary to determine if
// any more data from the file should be loaded. A ReadMoreCallback function
// facilitates this by returning true if more data should be read, or false
// otherwise.
type FileReadMoreCallback func(fileContents string, fileSize int64) (ok bool)

//export gotk4_FileReadMoreCallback
func gotk4_FileReadMoreCallback(arg0 *C.char, arg1 C.goffset, arg2 C.gpointer) (cret C.gboolean) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var fileContents string // out
	var fileSize int64      // out

	fileContents = C.GoString(arg0)
	fileSize = int64(arg1)

	fn := v.(FileReadMoreCallback)
	ok := fn(fileContents, fileSize)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// PollableSourceFunc: this is the function type of the callback used for the
// #GSource returned by g_pollable_input_stream_create_source() and
// g_pollable_output_stream_create_source().
type PollableSourceFunc func(pollableStream gextras.Objector) (ok bool)

//export gotk4_PollableSourceFunc
func gotk4_PollableSourceFunc(arg0 *C.GObject, arg1 C.gpointer) (cret C.gboolean) {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	var pollableStream gextras.Objector // out

	pollableStream = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(gextras.Objector)

	fn := v.(PollableSourceFunc)
	ok := fn(pollableStream)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// SocketSourceFunc: this is the function type of the callback used for the
// #GSource returned by g_socket_create_source().
type SocketSourceFunc func(socket Socket, condition glib.IOCondition) (ok bool)

//export gotk4_SocketSourceFunc
func gotk4_SocketSourceFunc(arg0 *C.GSocket, arg1 C.GIOCondition, arg2 C.gpointer) (cret C.gboolean) {
	v := box.Get(uintptr(arg2))
	if v == nil {
		panic(`callback not found`)
	}

	var socket Socket              // out
	var condition glib.IOCondition // out

	socket = gextras.CastObject(externglib.Take(unsafe.Pointer(arg0))).(Socket)
	condition = glib.IOCondition(arg1)

	fn := v.(SocketSourceFunc)
	ok := fn(socket, condition)

	if ok {
		cret = C.TRUE
	}

	return cret
}

// FileAttributeMatcher determines if a string matches a file attribute.
type FileAttributeMatcher struct {
	native C.GFileAttributeMatcher
}

// WrapFileAttributeMatcher wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapFileAttributeMatcher(ptr unsafe.Pointer) *FileAttributeMatcher {
	return (*FileAttributeMatcher)(ptr)
}

func marshalFileAttributeMatcher(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*FileAttributeMatcher)(unsafe.Pointer(b)), nil
}

// NewFileAttributeMatcher constructs a struct FileAttributeMatcher.
func NewFileAttributeMatcher(attributes string) *FileAttributeMatcher {
	var _arg1 *C.char                  // out
	var _cret *C.GFileAttributeMatcher // in

	_arg1 = (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_attribute_matcher_new(_arg1)

	var _fileAttributeMatcher *FileAttributeMatcher // out

	_fileAttributeMatcher = (*FileAttributeMatcher)(unsafe.Pointer(_cret))
	C.g_file_attribute_matcher_ref(_cret)
	runtime.SetFinalizer(_fileAttributeMatcher, func(v *FileAttributeMatcher) {
		C.g_file_attribute_matcher_unref((*C.GFileAttributeMatcher)(unsafe.Pointer(v)))
	})

	return _fileAttributeMatcher
}

// Native returns the underlying C source pointer.
func (f *FileAttributeMatcher) Native() unsafe.Pointer {
	return unsafe.Pointer(&f.native)
}

// EnumerateNamespace checks if the matcher will match all of the keys in a
// given namespace. This will always return true if a wildcard character is in
// use (e.g. if matcher was created with "standard::*" and @ns is "standard", or
// if matcher was created using "*" and namespace is anything.)
//
// TODO: this is awkwardly worded.
func (m *FileAttributeMatcher) EnumerateNamespace(ns string) bool {
	var _arg0 *C.GFileAttributeMatcher // out
	var _arg1 *C.char                  // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GFileAttributeMatcher)(unsafe.Pointer(m))
	_arg1 = (*C.char)(C.CString(ns))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_attribute_matcher_enumerate_namespace(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// EnumerateNext gets the next matched attribute from a AttributeMatcher.
func (m *FileAttributeMatcher) EnumerateNext() string {
	var _arg0 *C.GFileAttributeMatcher // out
	var _cret *C.char                  // in

	_arg0 = (*C.GFileAttributeMatcher)(unsafe.Pointer(m))

	_cret = C.g_file_attribute_matcher_enumerate_next(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Matches checks if an attribute will be matched by an attribute matcher. If
// the matcher was created with the "*" matching string, this function will
// always return true.
func (m *FileAttributeMatcher) Matches(attribute string) bool {
	var _arg0 *C.GFileAttributeMatcher // out
	var _arg1 *C.char                  // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GFileAttributeMatcher)(unsafe.Pointer(m))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_attribute_matcher_matches(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MatchesOnly checks if a attribute matcher only matches a given attribute.
// Always returns false if "*" was used when creating the matcher.
func (m *FileAttributeMatcher) MatchesOnly(attribute string) bool {
	var _arg0 *C.GFileAttributeMatcher // out
	var _arg1 *C.char                  // out
	var _cret C.gboolean               // in

	_arg0 = (*C.GFileAttributeMatcher)(unsafe.Pointer(m))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_file_attribute_matcher_matches_only(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Ref references a file attribute matcher.
func (m *FileAttributeMatcher) ref() *FileAttributeMatcher {
	var _arg0 *C.GFileAttributeMatcher // out
	var _cret *C.GFileAttributeMatcher // in

	_arg0 = (*C.GFileAttributeMatcher)(unsafe.Pointer(m))

	_cret = C.g_file_attribute_matcher_ref(_arg0)

	var _fileAttributeMatcher *FileAttributeMatcher // out

	_fileAttributeMatcher = (*FileAttributeMatcher)(unsafe.Pointer(_cret))
	C.g_file_attribute_matcher_ref(_cret)
	runtime.SetFinalizer(_fileAttributeMatcher, func(v *FileAttributeMatcher) {
		C.g_file_attribute_matcher_unref((*C.GFileAttributeMatcher)(unsafe.Pointer(v)))
	})

	return _fileAttributeMatcher
}

// Subtract subtracts all attributes of @subtract from @matcher and returns a
// matcher that supports those attributes.
//
// Note that currently it is not possible to remove a single attribute when the
// @matcher matches the whole namespace - or remove a namespace or attribute
// when the matcher matches everything. This is a limitation of the current
// implementation, but may be fixed in the future.
func (m *FileAttributeMatcher) Subtract(subtract *FileAttributeMatcher) *FileAttributeMatcher {
	var _arg0 *C.GFileAttributeMatcher // out
	var _arg1 *C.GFileAttributeMatcher // out
	var _cret *C.GFileAttributeMatcher // in

	_arg0 = (*C.GFileAttributeMatcher)(unsafe.Pointer(m))
	_arg1 = (*C.GFileAttributeMatcher)(unsafe.Pointer(subtract))

	_cret = C.g_file_attribute_matcher_subtract(_arg0, _arg1)

	var _fileAttributeMatcher *FileAttributeMatcher // out

	_fileAttributeMatcher = (*FileAttributeMatcher)(unsafe.Pointer(_cret))
	C.g_file_attribute_matcher_ref(_cret)
	runtime.SetFinalizer(_fileAttributeMatcher, func(v *FileAttributeMatcher) {
		C.g_file_attribute_matcher_unref((*C.GFileAttributeMatcher)(unsafe.Pointer(v)))
	})

	return _fileAttributeMatcher
}

// String prints what the matcher is matching against. The format will be equal
// to the format passed to g_file_attribute_matcher_new(). The output however,
// might not be identical, as the matcher may decide to use a different order or
// omit needless parts.
func (m *FileAttributeMatcher) String() string {
	var _arg0 *C.GFileAttributeMatcher // out
	var _cret *C.char                  // in

	_arg0 = (*C.GFileAttributeMatcher)(unsafe.Pointer(m))

	_cret = C.g_file_attribute_matcher_to_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// Unref unreferences @matcher. If the reference count falls below 1, the
// @matcher is automatically freed.
func (m *FileAttributeMatcher) unref() {
	var _arg0 *C.GFileAttributeMatcher // out

	_arg0 = (*C.GFileAttributeMatcher)(unsafe.Pointer(m))

	C.g_file_attribute_matcher_unref(_arg0)
}

// InputMessage: structure used for scatter/gather data input when receiving
// multiple messages or packets in one go. You generally pass in an array of
// empty Vectors and the operation will use all the buffers as if they were one
// buffer, and will set @bytes_received to the total number of bytes received
// across all Vectors.
//
// This structure closely mirrors `struct mmsghdr` and `struct msghdr` from the
// POSIX sockets API (see `man 2 recvmmsg`).
//
// If @address is non-nil then it is set to the source address the message was
// received from, and the caller must free it afterwards.
//
// If @control_messages is non-nil then it is set to an array of control
// messages received with the message (if any), and the caller must free it
// afterwards. @num_control_messages is set to the number of elements in this
// array, which may be zero.
//
// Flags relevant to this message will be returned in @flags. For example,
// `MSG_EOR` or `MSG_TRUNC`.
type InputMessage struct {
	native C.GInputMessage
}

// WrapInputMessage wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapInputMessage(ptr unsafe.Pointer) *InputMessage {
	return (*InputMessage)(ptr)
}

// Native returns the underlying C source pointer.
func (i *InputMessage) Native() unsafe.Pointer {
	return unsafe.Pointer(&i.native)
}

// InputVector: structure used for scatter/gather data input. You generally pass
// in an array of Vectors and the operation will store the read data starting in
// the first buffer, switching to the next as needed.
type InputVector struct {
	native C.GInputVector
}

// WrapInputVector wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapInputVector(ptr unsafe.Pointer) *InputVector {
	return (*InputVector)(ptr)
}

// Native returns the underlying C source pointer.
func (i *InputVector) Native() unsafe.Pointer {
	return unsafe.Pointer(&i.native)
}

// OutputMessage: structure used for scatter/gather data output when sending
// multiple messages or packets in one go. You generally pass in an array of
// Vectors and the operation will use all the buffers as if they were one
// buffer.
//
// If @address is nil then the message is sent to the default receiver (as
// previously set by g_socket_connect()).
type OutputMessage struct {
	native C.GOutputMessage
}

// WrapOutputMessage wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapOutputMessage(ptr unsafe.Pointer) *OutputMessage {
	return (*OutputMessage)(ptr)
}

// Native returns the underlying C source pointer.
func (o *OutputMessage) Native() unsafe.Pointer {
	return unsafe.Pointer(&o.native)
}

// OutputVector: structure used for scatter/gather data output. You generally
// pass in an array of Vectors and the operation will use all the buffers as if
// they were one buffer.
type OutputVector struct {
	native C.GOutputVector
}

// WrapOutputVector wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapOutputVector(ptr unsafe.Pointer) *OutputVector {
	return (*OutputVector)(ptr)
}

// Native returns the underlying C source pointer.
func (o *OutputVector) Native() unsafe.Pointer {
	return unsafe.Pointer(&o.native)
}

// Resource applications and libraries often contain binary or textual data that
// is really part of the application, rather than user data. For instance
// Builder .ui files, splashscreen images, GMenu markup XML, CSS files, icons,
// etc. These are often shipped as files in `$datadir/appname`, or manually
// included as literal strings in the code.
//
// The #GResource API and the [glib-compile-resources][glib-compile-resources]
// program provide a convenient and efficient alternative to this which has some
// nice properties. You maintain the files as normal files, so its easy to edit
// them, but during the build the files are combined into a binary bundle that
// is linked into the executable. This means that loading the resource files are
// efficient (as they are already in memory, shared with other instances) and
// simple (no need to check for things like I/O errors or locate the files in
// the filesystem). It also makes it easier to create relocatable applications.
//
// Resource files can also be marked as compressed. Such files will be included
// in the resource bundle in a compressed form, but will be automatically
// uncompressed when the resource is used. This is very useful e.g. for larger
// text files that are parsed once (or rarely) and then thrown away.
//
// Resource files can also be marked to be preprocessed, by setting the value of
// the `preprocess` attribute to a comma-separated list of preprocessing
// options. The only options currently supported are:
//
// `xml-stripblanks` which will use the xmllint command to strip ignorable
// whitespace from the XML file. For this to work, the `XMLLINT` environment
// variable must be set to the full path to the xmllint executable, or xmllint
// must be in the `PATH`; otherwise the preprocessing step is skipped.
//
// `to-pixdata` (deprecated since gdk-pixbuf 2.32) which will use the
// `gdk-pixbuf-pixdata` command to convert images to the Pixdata format, which
// allows you to create pixbufs directly using the data inside the resource
// file, rather than an (uncompressed) copy of it. For this, the
// `gdk-pixbuf-pixdata` program must be in the `PATH`, or the
// `GDK_PIXBUF_PIXDATA` environment variable must be set to the full path to the
// `gdk-pixbuf-pixdata` executable; otherwise the resource compiler will abort.
// `to-pixdata` has been deprecated since gdk-pixbuf 2.32, as #GResource
// supports embedding modern image formats just as well. Instead of using it,
// embed a PNG or SVG file in your #GResource.
//
// `json-stripblanks` which will use the `json-glib-format` command to strip
// ignorable whitespace from the JSON file. For this to work, the
// `JSON_GLIB_FORMAT` environment variable must be set to the full path to the
// `json-glib-format` executable, or it must be in the `PATH`; otherwise the
// preprocessing step is skipped. In addition, at least version 1.6 of
// `json-glib-format` is required.
//
// Resource files will be exported in the GResource namespace using the
// combination of the given `prefix` and the filename from the `file` element.
// The `alias` attribute can be used to alter the filename to expose them at a
// different location in the resource namespace. Typically, this is used to
// include files from a different source directory without exposing the source
// directory in the resource namespace, as in the example below.
//
// Resource bundles are created by the
// [glib-compile-resources][glib-compile-resources] program which takes an XML
// file that describes the bundle, and a set of files that the XML references.
// These are combined into a binary resource bundle.
//
// An example resource description:
//
//    <?xml version="1.0" encoding="UTF-8"?>
//    <gresources>
//      <gresource prefix="/org/gtk/Example">
//        <file>data/splashscreen.png</file>
//        <file compressed="true">dialog.ui</file>
//        <file preprocess="xml-stripblanks">menumarkup.xml</file>
//        <file alias="example.css">data/example.css</file>
//      </gresource>
//    </gresources>
//
// This will create a resource bundle with the following files:
//
//    /org/gtk/Example/data/splashscreen.png
//    /org/gtk/Example/dialog.ui
//    /org/gtk/Example/menumarkup.xml
//    /org/gtk/Example/example.css
//
// Note that all resources in the process share the same namespace, so use
// Java-style path prefixes (like in the above example) to avoid conflicts.
//
// You can then use [glib-compile-resources][glib-compile-resources] to compile
// the XML to a binary bundle that you can load with g_resource_load(). However,
// its more common to use the --generate-source and --generate-header arguments
// to create a source file and header to link directly into your application.
// This will generate `get_resource()`, `register_resource()` and
// `unregister_resource()` functions, prefixed by the `--c-name` argument passed
// to [glib-compile-resources][glib-compile-resources]. `get_resource()` returns
// the generated #GResource object. The register and unregister functions
// register the resource so its files can be accessed using
// g_resources_lookup_data().
//
// Once a #GResource has been created and registered all the data in it can be
// accessed globally in the process by using API calls like
// g_resources_open_stream() to stream the data or g_resources_lookup_data() to
// get a direct pointer to the data. You can also use URIs like
// "resource:///org/gtk/Example/data/splashscreen.png" with #GFile to access the
// resource data.
//
// Some higher-level APIs, such as Application, will automatically load
// resources from certain well-known paths in the resource namespace as a
// convenience. See the documentation for those APIs for details.
//
// There are two forms of the generated source, the default version uses the
// compiler support for constructor and destructor functions (where available)
// to automatically create and register the #GResource on startup or library
// load time. If you pass `--manual-register`, two functions to
// register/unregister the resource are created instead. This requires an
// explicit initialization call in your application/library, but it works on all
// platforms, even on the minor ones where constructors are not supported.
// (Constructor support is available for at least Win32, Mac OS and Linux.)
//
// Note that resource data can point directly into the data segment of e.g. a
// library, so if you are unloading libraries during runtime you need to be very
// careful with keeping around pointers to data from a resource, as this goes
// away when the library is unloaded. However, in practice this is not generally
// a problem, since most resource accesses are for your own resources, and
// resource data is often used once, during parsing, and then released.
//
// When debugging a program or testing a change to an installed version, it is
// often useful to be able to replace resources in the program or library,
// without recompiling, for debugging or quick hacking and testing purposes.
// Since GLib 2.50, it is possible to use the `G_RESOURCE_OVERLAYS` environment
// variable to selectively overlay resources with replacements from the
// filesystem. It is a G_SEARCHPATH_SEPARATOR-separated list of substitutions to
// perform during resource lookups. It is ignored when running in a setuid
// process.
//
// A substitution has the form
//
//    /org/gtk/libgtk=/home/desrt/gtk-overlay
//
// The part before the `=` is the resource subpath for which the overlay
// applies. The part after is a filesystem path which contains files and
// subdirectories as you would like to be loaded as resources with the
// equivalent names.
//
// In the example above, if an application tried to load a resource with the
// resource path `/org/gtk/libgtk/ui/gtkdialog.ui` then GResource would check
// the filesystem path `/home/desrt/gtk-overlay/ui/gtkdialog.ui`. If a file was
// found there, it would be used instead. This is an overlay, not an outright
// replacement, which means that if a file is not found at that path, the
// built-in version will be used instead. Whiteouts are not currently supported.
//
// Substitutions must start with a slash, and must not contain a trailing slash
// before the '='. The path after the slash should ideally be absolute, but this
// is not strictly required. It is possible to overlay the location of a single
// resource with an individual file.
type Resource struct {
	native C.GResource
}

// WrapResource wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapResource(ptr unsafe.Pointer) *Resource {
	return (*Resource)(ptr)
}

func marshalResource(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*Resource)(unsafe.Pointer(b)), nil
}

// Native returns the underlying C source pointer.
func (r *Resource) Native() unsafe.Pointer {
	return unsafe.Pointer(&r.native)
}

// EnumerateChildren returns all the names of children at the specified @path in
// the resource. The return result is a nil terminated list of strings which
// should be released with g_strfreev().
//
// If @path is invalid or does not exist in the #GResource,
// G_RESOURCE_ERROR_NOT_FOUND will be returned.
//
// @lookup_flags controls the behaviour of the lookup.
func (r *Resource) EnumerateChildren(path string, lookupFlags ResourceLookupFlags) ([]string, error) {
	var _arg0 *C.GResource           // out
	var _arg1 *C.char                // out
	var _arg2 C.GResourceLookupFlags // out
	var _cret **C.char
	var _cerr *C.GError // in

	_arg0 = (*C.GResource)(unsafe.Pointer(r))
	_arg1 = (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GResourceLookupFlags(lookupFlags)

	_cret = C.g_resource_enumerate_children(_arg0, _arg1, _arg2, &_cerr)

	var _utf8s []string
	var _goerr error // out

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _utf8s, _goerr
}

// Info looks for a file at the specified @path in the resource and if found
// returns information about it.
//
// @lookup_flags controls the behaviour of the lookup.
func (r *Resource) Info(path string, lookupFlags ResourceLookupFlags) (uint, uint32, error) {
	var _arg0 *C.GResource           // out
	var _arg1 *C.char                // out
	var _arg2 C.GResourceLookupFlags // out
	var _arg3 C.gsize                // in
	var _arg4 C.guint32              // in
	var _cerr *C.GError              // in

	_arg0 = (*C.GResource)(unsafe.Pointer(r))
	_arg1 = (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GResourceLookupFlags(lookupFlags)

	C.g_resource_get_info(_arg0, _arg1, _arg2, &_arg3, &_arg4, &_cerr)

	var _size uint    // out
	var _flags uint32 // out
	var _goerr error  // out

	_size = uint(_arg3)
	_flags = uint32(_arg4)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _size, _flags, _goerr
}

// OpenStream looks for a file at the specified @path in the resource and
// returns a Stream that lets you read the data.
//
// @lookup_flags controls the behaviour of the lookup.
func (r *Resource) OpenStream(path string, lookupFlags ResourceLookupFlags) (InputStream, error) {
	var _arg0 *C.GResource           // out
	var _arg1 *C.char                // out
	var _arg2 C.GResourceLookupFlags // out
	var _cret *C.GInputStream        // in
	var _cerr *C.GError              // in

	_arg0 = (*C.GResource)(unsafe.Pointer(r))
	_arg1 = (*C.char)(C.CString(path))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.GResourceLookupFlags(lookupFlags)

	_cret = C.g_resource_open_stream(_arg0, _arg1, _arg2, &_cerr)

	var _inputStream InputStream // out
	var _goerr error             // out

	_inputStream = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret))).(InputStream)
	_goerr = gerror.Take(unsafe.Pointer(_cerr))

	return _inputStream, _goerr
}

// Ref: atomically increments the reference count of @resource by one. This
// function is MT-safe and may be called from any thread.
func (r *Resource) ref() *Resource {
	var _arg0 *C.GResource // out
	var _cret *C.GResource // in

	_arg0 = (*C.GResource)(unsafe.Pointer(r))

	_cret = C.g_resource_ref(_arg0)

	var _ret *Resource // out

	_ret = (*Resource)(unsafe.Pointer(_cret))
	C.g_resource_ref(_cret)
	runtime.SetFinalizer(_ret, func(v *Resource) {
		C.g_resource_unref((*C.GResource)(unsafe.Pointer(v)))
	})

	return _ret
}

// Unref: atomically decrements the reference count of @resource by one. If the
// reference count drops to 0, all memory allocated by the resource is released.
// This function is MT-safe and may be called from any thread.
func (r *Resource) unref() {
	var _arg0 *C.GResource // out

	_arg0 = (*C.GResource)(unsafe.Pointer(r))

	C.g_resource_unref(_arg0)
}

// SrvTarget: SRV (service) records are used by some network protocols to
// provide service-specific aliasing and load-balancing. For example, XMPP
// (Jabber) uses SRV records to locate the XMPP server for a domain; rather than
// connecting directly to "example.com" or assuming a specific server hostname
// like "xmpp.example.com", an XMPP client would look up the "xmpp-client" SRV
// record for "example.com", and then connect to whatever host was pointed to by
// that record.
//
// You can use g_resolver_lookup_service() or g_resolver_lookup_service_async()
// to find the Targets for a given service. However, if you are simply planning
// to connect to the remote service, you can use Service's Connectable interface
// and not need to worry about Target at all.
type SrvTarget struct {
	native C.GSrvTarget
}

// WrapSrvTarget wraps the C unsafe.Pointer to be the right type. It is
// primarily used internally.
func WrapSrvTarget(ptr unsafe.Pointer) *SrvTarget {
	return (*SrvTarget)(ptr)
}

func marshalSrvTarget(p uintptr) (interface{}, error) {
	b := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return (*SrvTarget)(unsafe.Pointer(b)), nil
}

// NewSrvTarget constructs a struct SrvTarget.
func NewSrvTarget(hostname string, port uint16, priority uint16, weight uint16) *SrvTarget {
	var _arg1 *C.gchar      // out
	var _arg2 C.guint16     // out
	var _arg3 C.guint16     // out
	var _arg4 C.guint16     // out
	var _cret *C.GSrvTarget // in

	_arg1 = (*C.gchar)(C.CString(hostname))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint16(port)
	_arg3 = C.guint16(priority)
	_arg4 = C.guint16(weight)

	_cret = C.g_srv_target_new(_arg1, _arg2, _arg3, _arg4)

	var _srvTarget *SrvTarget // out

	_srvTarget = (*SrvTarget)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_srvTarget, func(v *SrvTarget) {
		C.free(unsafe.Pointer(v))
	})

	return _srvTarget
}

// Native returns the underlying C source pointer.
func (s *SrvTarget) Native() unsafe.Pointer {
	return unsafe.Pointer(&s.native)
}

// Copy copies @target
func (t *SrvTarget) Copy() *SrvTarget {
	var _arg0 *C.GSrvTarget // out
	var _cret *C.GSrvTarget // in

	_arg0 = (*C.GSrvTarget)(unsafe.Pointer(t))

	_cret = C.g_srv_target_copy(_arg0)

	var _srvTarget *SrvTarget // out

	_srvTarget = (*SrvTarget)(unsafe.Pointer(_cret))
	runtime.SetFinalizer(_srvTarget, func(v *SrvTarget) {
		C.free(unsafe.Pointer(v))
	})

	return _srvTarget
}

// Free frees @target
func (t *SrvTarget) free() {
	var _arg0 *C.GSrvTarget // out

	_arg0 = (*C.GSrvTarget)(unsafe.Pointer(t))

	C.g_srv_target_free(_arg0)
}

// Hostname gets @target's hostname (in ASCII form; if you are going to present
// this to the user, you should use g_hostname_is_ascii_encoded() to check if it
// contains encoded Unicode segments, and use g_hostname_to_unicode() to convert
// it if it does.)
func (t *SrvTarget) Hostname() string {
	var _arg0 *C.GSrvTarget // out
	var _cret *C.gchar      // in

	_arg0 = (*C.GSrvTarget)(unsafe.Pointer(t))

	_cret = C.g_srv_target_get_hostname(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Port gets @target's port
func (t *SrvTarget) Port() uint16 {
	var _arg0 *C.GSrvTarget // out
	var _cret C.guint16     // in

	_arg0 = (*C.GSrvTarget)(unsafe.Pointer(t))

	_cret = C.g_srv_target_get_port(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// Priority gets @target's priority. You should not need to look at this;
// #GResolver already sorts the targets according to the algorithm in RFC 2782.
func (t *SrvTarget) Priority() uint16 {
	var _arg0 *C.GSrvTarget // out
	var _cret C.guint16     // in

	_arg0 = (*C.GSrvTarget)(unsafe.Pointer(t))

	_cret = C.g_srv_target_get_priority(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// Weight gets @target's weight. You should not need to look at this; #GResolver
// already sorts the targets according to the algorithm in RFC 2782.
func (t *SrvTarget) Weight() uint16 {
	var _arg0 *C.GSrvTarget // out
	var _cret C.guint16     // in

	_arg0 = (*C.GSrvTarget)(unsafe.Pointer(t))

	_cret = C.g_srv_target_get_weight(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}
