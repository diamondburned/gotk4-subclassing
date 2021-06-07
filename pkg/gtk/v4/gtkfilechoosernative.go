// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_file_chooser_native_get_type()), F: marshalFileChooserNative},
	})
}

// FileChooserNative is an abstraction of a dialog box suitable for use with
// “File Open” or “File Save as” commands. By default, this just uses a
// FileChooserDialog to implement the actual dialog. However, on certain
// platforms, such as Windows and macOS, the native platform file chooser is
// used instead. When the application is running in a sandboxed environment
// without direct filesystem access (such as Flatpak), FileChooserNative may
// call the proper APIs (portals) to let the user choose a file and make it
// available to the application.
//
// While the API of FileChooserNative closely mirrors FileChooserDialog, the
// main difference is that there is no access to any Window or Widget for the
// dialog. This is required, as there may not be one in the case of a platform
// native dialog.
//
// Showing, hiding and running the dialog is handled by the NativeDialog
// functions.
//
// Note that unlike FileChooserDialog, FileChooserNative objects are not
// toplevel widgets, and GTK does not keep them alive. It is your responsibility
// to keep a reference until you are done with the object.
//
//
// Typical usage
//
// In the simplest of cases, you can the following code to use FileChooserDialog
// to select a file for opening:
//
//    static void
//    on_response (GtkNativeDialog *native,
//                 int              response)
//    {
//      if (response == GTK_RESPONSE_ACCEPT)
//        {
//          GtkFileChooser *chooser = GTK_FILE_CHOOSER (native);
//          GFile *file = gtk_file_chooser_get_file (chooser);
//
//          save_to_file (file);
//
//          g_object_unref (file);
//        }
//
//      g_object_unref (native);
//    }
//
//      // ...
//      GtkFileChooserNative *native;
//      GtkFileChooser *chooser;
//      GtkFileChooserAction action = GTK_FILE_CHOOSER_ACTION_SAVE;
//
//      native = gtk_file_chooser_native_new ("Save File",
//                                            parent_window,
//                                            action,
//                                            "_Save",
//                                            "_Cancel");
//      chooser = GTK_FILE_CHOOSER (native);
//
//      if (user_edited_a_new_document)
//        gtk_file_chooser_set_current_name (chooser, _("Untitled document"));
//      else
//        gtk_file_chooser_set_file (chooser, existing_file, NULL);
//
//      g_signal_connect (native, "response", G_CALLBACK (on_response), NULL);
//      gtk_native_dialog_show (GTK_NATIVE_DIALOG (native));
//
// For more information on how to best set up a file dialog, see
// FileChooserDialog.
//
//
// Response Codes
//
// FileChooserNative inherits from NativeDialog, which means it will return
// K_RESPONSE_ACCEPT if the user accepted, and K_RESPONSE_CANCEL if he pressed
// cancel. It can also return K_RESPONSE_DELETE_EVENT if the window was
// unexpectedly closed.
//
// Differences from FileChooserDialog ##
// {#gtkfilechooserdialognative-differences}
//
// There are a few things in the GtkFileChooser API that are not possible to use
// with FileChooserNative, as such use would prohibit the use of a native
// dialog.
//
// No operations that change the dialog work while the dialog is visible. Set
// all the properties that are required before showing the dialog.
//
//
// Win32 details
//
// On windows the IFileDialog implementation (added in Windows Vista) is used.
// It supports many of the features that FileChooserDialog does, but there are
// some things it does not handle:
//
// * Any FileFilter added using a mimetype
//
// If any of these features are used the regular FileChooserDialog will be used
// in place of the native one.
//
//
// Portal details
//
// When the org.freedesktop.portal.FileChooser portal is available on the
// session bus, it is used to bring up an out-of-process file chooser. Depending
// on the kind of session the application is running in, this may or may not be
// a GTK file chooser.
//
// macOS details
//
// On macOS the NSSavePanel and NSOpenPanel classes are used to provide native
// file chooser dialogs. Some features provided by FileChooserDialog are not
// supported:
//
// * Shortcut folders.
type FileChooserNative interface {
	NativeDialog
	FileChooser

	// AcceptLabel retrieves the custom label text for the accept button.
	AcceptLabel(s FileChooserNative)
	// CancelLabel retrieves the custom label text for the cancel button.
	CancelLabel(s FileChooserNative)
	// SetAcceptLabel sets the custom label text for the accept button.
	//
	// If characters in @label are preceded by an underscore, they are
	// underlined. If you need a literal underscore character in a label, use
	// “__” (two underscores). The first underlined character represents a
	// keyboard accelerator called a mnemonic. Pressing Alt and that key
	// activates the button.
	SetAcceptLabel(s FileChooserNative, acceptLabel string)
	// SetCancelLabel sets the custom label text for the cancel button.
	//
	// If characters in @label are preceded by an underscore, they are
	// underlined. If you need a literal underscore character in a label, use
	// “__” (two underscores). The first underlined character represents a
	// keyboard accelerator called a mnemonic. Pressing Alt and that key
	// activates the button.
	SetCancelLabel(s FileChooserNative, cancelLabel string)
}

// fileChooserNative implements the FileChooserNative interface.
type fileChooserNative struct {
	NativeDialog
	FileChooser
}

var _ FileChooserNative = (*fileChooserNative)(nil)

// WrapFileChooserNative wraps a GObject to the right type. It is
// primarily used internally.
func WrapFileChooserNative(obj *externglib.Object) FileChooserNative {
	return FileChooserNative{
		NativeDialog: WrapNativeDialog(obj),
		FileChooser:  WrapFileChooser(obj),
	}
}

func marshalFileChooserNative(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileChooserNative(obj), nil
}

// NewFileChooserNative constructs a class FileChooserNative.
func NewFileChooserNative(title string, parent Window, action FileChooserAction, acceptLabel string, cancelLabel string) {
	var arg1 *C.char
	var arg2 *C.GtkWindow
	var arg3 C.GtkFileChooserAction
	var arg4 *C.char
	var arg5 *C.char

	arg1 = (*C.char)(C.CString(title))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GtkWindow)(unsafe.Pointer(parent.Native()))
	arg3 = (C.GtkFileChooserAction)(action)
	arg4 = (*C.char)(C.CString(acceptLabel))
	defer C.free(unsafe.Pointer(arg4))
	arg5 = (*C.char)(C.CString(cancelLabel))
	defer C.free(unsafe.Pointer(arg5))

	C.gtk_file_chooser_native_new(arg1, arg2, arg3, arg4, arg5)
}

// AcceptLabel retrieves the custom label text for the accept button.
func (s fileChooserNative) AcceptLabel(s FileChooserNative) {
	var arg0 *C.GtkFileChooserNative

	arg0 = (*C.GtkFileChooserNative)(unsafe.Pointer(s.Native()))

	C.gtk_file_chooser_native_get_accept_label(arg0)
}

// CancelLabel retrieves the custom label text for the cancel button.
func (s fileChooserNative) CancelLabel(s FileChooserNative) {
	var arg0 *C.GtkFileChooserNative

	arg0 = (*C.GtkFileChooserNative)(unsafe.Pointer(s.Native()))

	C.gtk_file_chooser_native_get_cancel_label(arg0)
}

// SetAcceptLabel sets the custom label text for the accept button.
//
// If characters in @label are preceded by an underscore, they are
// underlined. If you need a literal underscore character in a label, use
// “__” (two underscores). The first underlined character represents a
// keyboard accelerator called a mnemonic. Pressing Alt and that key
// activates the button.
func (s fileChooserNative) SetAcceptLabel(s FileChooserNative, acceptLabel string) {
	var arg0 *C.GtkFileChooserNative
	var arg1 *C.char

	arg0 = (*C.GtkFileChooserNative)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(acceptLabel))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_file_chooser_native_set_accept_label(arg0, arg1)
}

// SetCancelLabel sets the custom label text for the cancel button.
//
// If characters in @label are preceded by an underscore, they are
// underlined. If you need a literal underscore character in a label, use
// “__” (two underscores). The first underlined character represents a
// keyboard accelerator called a mnemonic. Pressing Alt and that key
// activates the button.
func (s fileChooserNative) SetCancelLabel(s FileChooserNative, cancelLabel string) {
	var arg0 *C.GtkFileChooserNative
	var arg1 *C.char

	arg0 = (*C.GtkFileChooserNative)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(cancelLabel))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_file_chooser_native_set_cancel_label(arg0, arg1)
}
