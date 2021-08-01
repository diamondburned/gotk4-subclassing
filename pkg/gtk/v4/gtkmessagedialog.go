// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_buttons_type_get_type()), F: marshalButtonsType},
		{T: externglib.Type(C.gtk_message_dialog_get_type()), F: marshalMessageDialogger},
	})
}

// ButtonsType: prebuilt sets of buttons for GtkDialog.
//
// If none of these choices are appropriate, simply use GTK_BUTTONS_NONE and
// call gtk.Dialog.AddButtons().
//
// > Please note that GTK_BUTTONS_OK, GTK_BUTTONS_YES_NO > and
// GTK_BUTTONS_OK_CANCEL are discouraged by the > GNOME Human Interface
// Guidelines (http://library.gnome.org/devel/hig-book/stable/).
type ButtonsType int

const (
	// ButtonsNone: no buttons at all
	ButtonsNone ButtonsType = iota
	// ButtonsOK: OK button
	ButtonsOK
	// ButtonsClose: close button
	ButtonsClose
	// ButtonsCancel: cancel button
	ButtonsCancel
	// ButtonsYesNo yes and No buttons
	ButtonsYesNo
	// ButtonsOKCancel: OK and Cancel buttons
	ButtonsOKCancel
)

func marshalButtonsType(p uintptr) (interface{}, error) {
	return ButtonsType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for ButtonsType.
func (b ButtonsType) String() string {
	switch b {
	case ButtonsNone:
		return "None"
	case ButtonsOK:
		return "OK"
	case ButtonsClose:
		return "Close"
	case ButtonsCancel:
		return "Cancel"
	case ButtonsYesNo:
		return "YesNo"
	case ButtonsOKCancel:
		return "OKCancel"
	default:
		return fmt.Sprintf("ButtonsType(%d)", b)
	}
}

// MessageDialog: GtkMessageDialog presents a dialog with some message text.
//
// !An example GtkMessageDialog (messagedialog.png)
//
// It’s simply a convenience widget; you could construct the equivalent of
// GtkMessageDialog from GtkDialog without too much effort, but GtkMessageDialog
// saves typing.
//
// The easiest way to do a modal message dialog is to use the GTK_DIALOG_MODAL
// flag, which will call gtk.Window.SetModal() internally. The dialog will
// prevent interaction with the parent window until it's hidden or destroyed.
// You can use the gtk.Dialog::response signal to know when the user dismissed
// the dialog.
//
// An example for using a modal dialog:
//
//    GtkDialogFlags flags = GTK_DIALOG_DESTROY_WITH_PARENT | GTK_DIALOG_MODAL;
//    dialog = gtk_message_dialog_new (parent_window,
//                                     flags,
//                                     GTK_MESSAGE_ERROR,
//                                     GTK_BUTTONS_CLOSE,
//                                     "Error reading “s”: s",
//                                     filename,
//                                     g_strerror (errno));
//    // Destroy the dialog when the user responds to it
//    // (e.g. clicks a button)
//
//    g_signal_connect (dialog, "response",
//                      G_CALLBACK (gtk_window_destroy),
//                      NULL);
//
//
// You might do a non-modal GtkMessageDialog simply by omitting the
// GTK_DIALOG_MODAL flag:
//
//    GtkDialogFlags flags = GTK_DIALOG_DESTROY_WITH_PARENT;
//    dialog = gtk_message_dialog_new (parent_window,
//                                     flags,
//                                     GTK_MESSAGE_ERROR,
//                                     GTK_BUTTONS_CLOSE,
//                                     "Error reading “s”: s",
//                                     filename,
//                                     g_strerror (errno));
//
//    // Destroy the dialog when the user responds to it
//    // (e.g. clicks a button)
//    g_signal_connect (dialog, "response",
//                      G_CALLBACK (gtk_window_destroy),
//                      NULL);
//
//
//
// GtkMessageDialog as GtkBuildable
//
// The GtkMessageDialog implementation of the GtkBuildable interface exposes the
// message area as an internal child with the name “message_area”.
type MessageDialog struct {
	Dialog
}

func wrapMessageDialog(obj *externglib.Object) *MessageDialog {
	return &MessageDialog{
		Dialog: Dialog{
			Window: Window{
				Widget: Widget{
					InitiallyUnowned: externglib.InitiallyUnowned{
						Object: obj,
					},
					Accessible: Accessible{
						Object: obj,
					},
					Buildable: Buildable{
						Object: obj,
					},
					ConstraintTarget: ConstraintTarget{
						Object: obj,
					},
					Object: obj,
				},
				Root: Root{
					NativeSurface: NativeSurface{
						Widget: Widget{
							InitiallyUnowned: externglib.InitiallyUnowned{
								Object: obj,
							},
							Accessible: Accessible{
								Object: obj,
							},
							Buildable: Buildable{
								Object: obj,
							},
							ConstraintTarget: ConstraintTarget{
								Object: obj,
							},
							Object: obj,
						},
					},
				},
				ShortcutManager: ShortcutManager{
					Object: obj,
				},
				Object: obj,
			},
		},
	}
}

func marshalMessageDialogger(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapMessageDialog(obj), nil
}

// MessageArea returns the message area of the dialog.
//
// This is the box where the dialog’s primary and secondary labels are packed.
// You can add your own extra content to that box and it will appear below those
// labels. See gtk.Dialog.GetContentArea() for the corresponding function in the
// parent gtk.Dialog.
func (messageDialog *MessageDialog) MessageArea() Widgetter {
	var _arg0 *C.GtkMessageDialog // out
	var _cret *C.GtkWidget        // in

	_arg0 = (*C.GtkMessageDialog)(unsafe.Pointer(messageDialog.Native()))

	_cret = C.gtk_message_dialog_get_message_area(_arg0)

	var _widget Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// SetMarkup sets the text of the message dialog.
func (messageDialog *MessageDialog) SetMarkup(str string) {
	var _arg0 *C.GtkMessageDialog // out
	var _arg1 *C.char             // out

	_arg0 = (*C.GtkMessageDialog)(unsafe.Pointer(messageDialog.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_message_dialog_set_markup(_arg0, _arg1)
}
