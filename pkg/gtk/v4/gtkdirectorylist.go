// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_directory_list_get_type()), F: marshalDirectoryList},
	})
}

// DirectoryList is a list model that wraps g_file_enumerate_children_async().
// It presents a Model and fills it asynchronously with the Infos returned from
// that function.
//
// Enumeration will start automatically when a the DirectoryList:file property
// is set.
//
// While the DirectoryList is being filled, the DirectoryList:loading property
// will be set to true. You can listen to that property if you want to show
// information like a Spinner or a "Loading..." text.
//
// If loading fails at any point, the DirectoryList:error property will be set
// to give more indication about the failure.
//
// The Infos returned from a DirectoryList have the "standard::file" attribute
// set to the #GFile they refer to. This way you can get at the file that is
// referred to in the same way you would via g_file_enumerator_get_child(). This
// means you do not need access to the DirectoryList but can access the #GFile
// directly from the Info when operating with a ListView or similar.
type DirectoryList interface {
	gextras.Objector
	gio.ListModel

	// Attributes gets the attributes queried on the children.
	Attributes(s DirectoryList)
	// Error gets the loading error, if any.
	//
	// If an error occurs during the loading process, the loading process will
	// finish and this property allows querying the error that happened. This
	// error will persist until a file is loaded again.
	//
	// An error being set does not mean that no files were loaded, and all
	// successfully queried files will remain in the list.
	Error(s DirectoryList)
	// File gets the file whose children are currently enumerated.
	File(s DirectoryList)
	// IOPriority gets the IO priority set via
	// gtk_directory_list_set_io_priority().
	IOPriority(s DirectoryList)
	// Monitored returns whether the directory list is monitoring the directory
	// for changes.
	Monitored(s DirectoryList) bool
	// IsLoading returns true if the children enumeration is currently in
	// progress.
	//
	// Files will be added to @self from time to time while loading is going on.
	// The order in which are added is undefined and may change in between runs.
	IsLoading(s DirectoryList) bool
	// SetAttributes sets the @attributes to be enumerated and starts the
	// enumeration.
	//
	// If @attributes is nil, no attributes will be queried, but a list of Infos
	// will still be created.
	SetAttributes(s DirectoryList, attributes string)
	// SetFile sets the @file to be enumerated and starts the enumeration.
	//
	// If @file is nil, the result will be an empty list.
	SetFile(s DirectoryList, file gio.File)
	// SetIOPriority sets the IO priority to use while loading directories.
	//
	// Setting the priority while @self is loading will reprioritize the ongoing
	// load as soon as possible.
	//
	// The default IO priority is G_PRIORITY_DEFAULT, which is higher than the
	// GTK redraw priority. If you are loading a lot of directories in parallel,
	// lowering it to something like G_PRIORITY_DEFAULT_IDLE may increase
	// responsiveness.
	SetIOPriority(s DirectoryList, ioPriority int)
	// SetMonitored sets whether the directory list will monitor the directory
	// for changes. If monitoring is enabled, the Model::items-changed signal
	// will be emitted when the directory contents change.
	//
	// When monitoring is turned on after the initial creation of the directory
	// list, the directory is reloaded to avoid missing files that appeared
	// between the initial loading and when monitoring was turned on.
	SetMonitored(s DirectoryList, monitored bool)
}

// directoryList implements the DirectoryList interface.
type directoryList struct {
	gextras.Objector
	gio.ListModel
}

var _ DirectoryList = (*directoryList)(nil)

// WrapDirectoryList wraps a GObject to the right type. It is
// primarily used internally.
func WrapDirectoryList(obj *externglib.Object) DirectoryList {
	return DirectoryList{
		Objector:      obj,
		gio.ListModel: gio.WrapListModel(obj),
	}
}

func marshalDirectoryList(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapDirectoryList(obj), nil
}

// NewDirectoryList constructs a class DirectoryList.
func NewDirectoryList(attributes string, file gio.File) {
	var arg1 *C.char
	var arg2 *C.GFile

	arg1 = (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GFile)(unsafe.Pointer(file.Native()))

	C.gtk_directory_list_new(arg1, arg2)
}

// Attributes gets the attributes queried on the children.
func (s directoryList) Attributes(s DirectoryList) {
	var arg0 *C.GtkDirectoryList

	arg0 = (*C.GtkDirectoryList)(unsafe.Pointer(s.Native()))

	C.gtk_directory_list_get_attributes(arg0)
}

// Error gets the loading error, if any.
//
// If an error occurs during the loading process, the loading process will
// finish and this property allows querying the error that happened. This
// error will persist until a file is loaded again.
//
// An error being set does not mean that no files were loaded, and all
// successfully queried files will remain in the list.
func (s directoryList) Error(s DirectoryList) {
	var arg0 *C.GtkDirectoryList

	arg0 = (*C.GtkDirectoryList)(unsafe.Pointer(s.Native()))

	C.gtk_directory_list_get_error(arg0)
}

// File gets the file whose children are currently enumerated.
func (s directoryList) File(s DirectoryList) {
	var arg0 *C.GtkDirectoryList

	arg0 = (*C.GtkDirectoryList)(unsafe.Pointer(s.Native()))

	C.gtk_directory_list_get_file(arg0)
}

// IOPriority gets the IO priority set via
// gtk_directory_list_set_io_priority().
func (s directoryList) IOPriority(s DirectoryList) {
	var arg0 *C.GtkDirectoryList

	arg0 = (*C.GtkDirectoryList)(unsafe.Pointer(s.Native()))

	C.gtk_directory_list_get_io_priority(arg0)
}

// Monitored returns whether the directory list is monitoring the directory
// for changes.
func (s directoryList) Monitored(s DirectoryList) bool {
	var arg0 *C.GtkDirectoryList

	arg0 = (*C.GtkDirectoryList)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_directory_list_get_monitored(arg0)

	if cret {
		ok = true
	}

	return ok
}

// IsLoading returns true if the children enumeration is currently in
// progress.
//
// Files will be added to @self from time to time while loading is going on.
// The order in which are added is undefined and may change in between runs.
func (s directoryList) IsLoading(s DirectoryList) bool {
	var arg0 *C.GtkDirectoryList

	arg0 = (*C.GtkDirectoryList)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_directory_list_is_loading(arg0)

	if cret {
		ok = true
	}

	return ok
}

// SetAttributes sets the @attributes to be enumerated and starts the
// enumeration.
//
// If @attributes is nil, no attributes will be queried, but a list of Infos
// will still be created.
func (s directoryList) SetAttributes(s DirectoryList, attributes string) {
	var arg0 *C.GtkDirectoryList
	var arg1 *C.char

	arg0 = (*C.GtkDirectoryList)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(attributes))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_directory_list_set_attributes(arg0, arg1)
}

// SetFile sets the @file to be enumerated and starts the enumeration.
//
// If @file is nil, the result will be an empty list.
func (s directoryList) SetFile(s DirectoryList, file gio.File) {
	var arg0 *C.GtkDirectoryList
	var arg1 *C.GFile

	arg0 = (*C.GtkDirectoryList)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GFile)(unsafe.Pointer(file.Native()))

	C.gtk_directory_list_set_file(arg0, arg1)
}

// SetIOPriority sets the IO priority to use while loading directories.
//
// Setting the priority while @self is loading will reprioritize the ongoing
// load as soon as possible.
//
// The default IO priority is G_PRIORITY_DEFAULT, which is higher than the
// GTK redraw priority. If you are loading a lot of directories in parallel,
// lowering it to something like G_PRIORITY_DEFAULT_IDLE may increase
// responsiveness.
func (s directoryList) SetIOPriority(s DirectoryList, ioPriority int) {
	var arg0 *C.GtkDirectoryList
	var arg1 C.int

	arg0 = (*C.GtkDirectoryList)(unsafe.Pointer(s.Native()))
	arg1 = C.int(ioPriority)

	C.gtk_directory_list_set_io_priority(arg0, arg1)
}

// SetMonitored sets whether the directory list will monitor the directory
// for changes. If monitoring is enabled, the Model::items-changed signal
// will be emitted when the directory contents change.
//
// When monitoring is turned on after the initial creation of the directory
// list, the directory is reloaded to avoid missing files that appeared
// between the initial loading and when monitoring was turned on.
func (s directoryList) SetMonitored(s DirectoryList, monitored bool) {
	var arg0 *C.GtkDirectoryList
	var arg1 C.gboolean

	arg0 = (*C.GtkDirectoryList)(unsafe.Pointer(s.Native()))
	if monitored {
		arg1 = C.gboolean(1)
	}

	C.gtk_directory_list_set_monitored(arg0, arg1)
}
