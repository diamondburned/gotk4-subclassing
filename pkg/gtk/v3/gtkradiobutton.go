// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_radio_button_get_type()), F: marshalRadioButtonner},
	})
}

// RadioButtonOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type RadioButtonOverrider interface {
	GroupChanged()
}

// RadioButton: single radio button performs the same basic function as a
// CheckButton, as its position in the object hierarchy reflects. It is only
// when multiple radio buttons are grouped together that they become a different
// user interface component in their own right.
//
// Every radio button is a member of some group of radio buttons. When one is
// selected, all other radio buttons in the same group are deselected. A
// RadioButton is one way of giving the user a choice from many options.
//
// Radio button widgets are created with gtk_radio_button_new(), passing NULL as
// the argument if this is the first radio button in a group. In subsequent
// calls, the group you wish to add this button to should be passed as an
// argument. Optionally, gtk_radio_button_new_with_label() can be used if you
// want a text label on the radio button.
//
// Alternatively, when adding widgets to an existing group of radio buttons, use
// gtk_radio_button_new_from_widget() with a RadioButton that already has a
// group assigned to it. The convenience function
// gtk_radio_button_new_with_label_from_widget() is also provided.
//
// To retrieve the group a RadioButton is assigned to, use
// gtk_radio_button_get_group().
//
// To remove a RadioButton from one group and make it part of a new one, use
// gtk_radio_button_set_group().
//
// The group list does not need to be freed, as each RadioButton will remove
// itself and its list item when it is destroyed.
//
// CSS nodes
//
//    void create_radio_buttons (void) {
//
//       GtkWidget *window, *radio1, *radio2, *box, *entry;
//       window = gtk_window_new (GTK_WINDOW_TOPLEVEL);
//       box = gtk_box_new (GTK_ORIENTATION_VERTICAL, 2);
//       gtk_box_set_homogeneous (GTK_BOX (box), TRUE);
//
//       // Create a radio button with a GtkEntry widget
//       radio1 = gtk_radio_button_new (NULL);
//       entry = gtk_entry_new ();
//       gtk_container_add (GTK_CONTAINER (radio1), entry);
//
//
//       // Create a radio button with a label
//       radio2 = gtk_radio_button_new_with_label_from_widget (GTK_RADIO_BUTTON (radio1),
//                                                             "I’m the second radio button.");
//
//       // Pack them into a box, then show all the widgets
//       gtk_box_pack_start (GTK_BOX (box), radio1);
//       gtk_box_pack_start (GTK_BOX (box), radio2);
//       gtk_container_add (GTK_CONTAINER (window), box);
//       gtk_widget_show_all (window);
//       return;
//    }
//
// When an unselected button in the group is clicked the clicked button receives
// the ToggleButton::toggled signal, as does the previously selected button.
// Inside the ToggleButton::toggled handler, gtk_toggle_button_get_active() can
// be used to determine if the button has been selected or deselected.
type RadioButton struct {
	CheckButton
}

func wrapRadioButton(obj *externglib.Object) *RadioButton {
	return &RadioButton{
		CheckButton: CheckButton{
			ToggleButton: ToggleButton{
				Button: Button{
					Bin: Bin{
						Container: Container{
							Widget: Widget{
								InitiallyUnowned: externglib.InitiallyUnowned{
									Object: obj,
								},
								ImplementorIface: atk.ImplementorIface{
									Object: obj,
								},
								Buildable: Buildable{
									Object: obj,
								},
								Object: obj,
							},
						},
					},
					Actionable: Actionable{
						Widget: Widget{
							InitiallyUnowned: externglib.InitiallyUnowned{
								Object: obj,
							},
							ImplementorIface: atk.ImplementorIface{
								Object: obj,
							},
							Buildable: Buildable{
								Object: obj,
							},
							Object: obj,
						},
					},
					Activatable: Activatable{
						Object: obj,
					},
					Object: obj,
				},
			},
		},
	}
}

func marshalRadioButtonner(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapRadioButton(obj), nil
}

// NewRadioButton creates a new RadioButton. To be of any practical value, a
// widget should then be packed into the radio button.
func NewRadioButton(group []RadioButton) *RadioButton {
	var _arg1 *C.GSList    // out
	var _cret *C.GtkWidget // in

	if group != nil {
		for i := len(group) - 1; i >= 0; i-- {
			src := group[i]
			var dst *C.GtkRadioButton // out
			dst = (*C.GtkRadioButton)(unsafe.Pointer((&src).Native()))
			_arg1 = C.g_slist_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
		}
		defer C.g_slist_free(_arg1)
	}

	_cret = C.gtk_radio_button_new(_arg1)
	runtime.KeepAlive(group)

	var _radioButton *RadioButton // out

	_radioButton = wrapRadioButton(externglib.Take(unsafe.Pointer(_cret)))

	return _radioButton
}

// NewRadioButtonFromWidget creates a new RadioButton, adding it to the same
// group as radio_group_member. As with gtk_radio_button_new(), a widget should
// be packed into the radio button.
func NewRadioButtonFromWidget(radioGroupMember *RadioButton) *RadioButton {
	var _arg1 *C.GtkRadioButton // out
	var _cret *C.GtkWidget      // in

	if radioGroupMember != nil {
		_arg1 = (*C.GtkRadioButton)(unsafe.Pointer(radioGroupMember.Native()))
	}

	_cret = C.gtk_radio_button_new_from_widget(_arg1)
	runtime.KeepAlive(radioGroupMember)

	var _radioButton *RadioButton // out

	_radioButton = wrapRadioButton(externglib.Take(unsafe.Pointer(_cret)))

	return _radioButton
}

// NewRadioButtonWithLabel creates a new RadioButton with a text label.
func NewRadioButtonWithLabel(group []RadioButton, label string) *RadioButton {
	var _arg1 *C.GSList    // out
	var _arg2 *C.gchar     // out
	var _cret *C.GtkWidget // in

	if group != nil {
		for i := len(group) - 1; i >= 0; i-- {
			src := group[i]
			var dst *C.GtkRadioButton // out
			dst = (*C.GtkRadioButton)(unsafe.Pointer((&src).Native()))
			_arg1 = C.g_slist_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
		}
		defer C.g_slist_free(_arg1)
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_radio_button_new_with_label(_arg1, _arg2)
	runtime.KeepAlive(group)
	runtime.KeepAlive(label)

	var _radioButton *RadioButton // out

	_radioButton = wrapRadioButton(externglib.Take(unsafe.Pointer(_cret)))

	return _radioButton
}

// NewRadioButtonWithLabelFromWidget creates a new RadioButton with a text
// label, adding it to the same group as radio_group_member.
func NewRadioButtonWithLabelFromWidget(radioGroupMember *RadioButton, label string) *RadioButton {
	var _arg1 *C.GtkRadioButton // out
	var _arg2 *C.gchar          // out
	var _cret *C.GtkWidget      // in

	if radioGroupMember != nil {
		_arg1 = (*C.GtkRadioButton)(unsafe.Pointer(radioGroupMember.Native()))
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_radio_button_new_with_label_from_widget(_arg1, _arg2)
	runtime.KeepAlive(radioGroupMember)
	runtime.KeepAlive(label)

	var _radioButton *RadioButton // out

	_radioButton = wrapRadioButton(externglib.Take(unsafe.Pointer(_cret)))

	return _radioButton
}

// NewRadioButtonWithMnemonic creates a new RadioButton containing a label,
// adding it to the same group as group. The label will be created using
// gtk_label_new_with_mnemonic(), so underscores in label indicate the mnemonic
// for the button.
func NewRadioButtonWithMnemonic(group []RadioButton, label string) *RadioButton {
	var _arg1 *C.GSList    // out
	var _arg2 *C.gchar     // out
	var _cret *C.GtkWidget // in

	if group != nil {
		for i := len(group) - 1; i >= 0; i-- {
			src := group[i]
			var dst *C.GtkRadioButton // out
			dst = (*C.GtkRadioButton)(unsafe.Pointer((&src).Native()))
			_arg1 = C.g_slist_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
		}
		defer C.g_slist_free(_arg1)
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_radio_button_new_with_mnemonic(_arg1, _arg2)
	runtime.KeepAlive(group)
	runtime.KeepAlive(label)

	var _radioButton *RadioButton // out

	_radioButton = wrapRadioButton(externglib.Take(unsafe.Pointer(_cret)))

	return _radioButton
}

// NewRadioButtonWithMnemonicFromWidget creates a new RadioButton containing a
// label. The label will be created using gtk_label_new_with_mnemonic(), so
// underscores in label indicate the mnemonic for the button.
func NewRadioButtonWithMnemonicFromWidget(radioGroupMember *RadioButton, label string) *RadioButton {
	var _arg1 *C.GtkRadioButton // out
	var _arg2 *C.gchar          // out
	var _cret *C.GtkWidget      // in

	if radioGroupMember != nil {
		_arg1 = (*C.GtkRadioButton)(unsafe.Pointer(radioGroupMember.Native()))
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(label)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gtk_radio_button_new_with_mnemonic_from_widget(_arg1, _arg2)
	runtime.KeepAlive(radioGroupMember)
	runtime.KeepAlive(label)

	var _radioButton *RadioButton // out

	_radioButton = wrapRadioButton(externglib.Take(unsafe.Pointer(_cret)))

	return _radioButton
}

// Group retrieves the group assigned to a radio button.
func (radioButton *RadioButton) Group() []RadioButton {
	var _arg0 *C.GtkRadioButton // out
	var _cret *C.GSList         // in

	_arg0 = (*C.GtkRadioButton)(unsafe.Pointer(radioButton.Native()))

	_cret = C.gtk_radio_button_get_group(_arg0)
	runtime.KeepAlive(radioButton)

	var _sList []RadioButton // out

	_sList = make([]RadioButton, 0, gextras.SListSize(unsafe.Pointer(_cret)))
	gextras.MoveSList(unsafe.Pointer(_cret), false, func(v unsafe.Pointer) {
		src := (*C.GtkRadioButton)(v)
		var dst RadioButton // out
		dst = *wrapRadioButton(externglib.Take(unsafe.Pointer(src)))
		_sList = append(_sList, dst)
	})

	return _sList
}

// JoinGroup joins a RadioButton object to the group of another RadioButton
// object
//
// Use this in language bindings instead of the gtk_radio_button_get_group() and
// gtk_radio_button_set_group() methods
//
// A common way to set up a group of radio buttons is the following:
//
//      GtkRadioButton *radio_button;
//      GtkRadioButton *last_button;
//
//      while (some_condition)
//        {
//           radio_button = gtk_radio_button_new (NULL);
//
//           gtk_radio_button_join_group (radio_button, last_button);
//           last_button = radio_button;
//        }.
func (radioButton *RadioButton) JoinGroup(groupSource *RadioButton) {
	var _arg0 *C.GtkRadioButton // out
	var _arg1 *C.GtkRadioButton // out

	_arg0 = (*C.GtkRadioButton)(unsafe.Pointer(radioButton.Native()))
	if groupSource != nil {
		_arg1 = (*C.GtkRadioButton)(unsafe.Pointer(groupSource.Native()))
	}

	C.gtk_radio_button_join_group(_arg0, _arg1)
	runtime.KeepAlive(radioButton)
	runtime.KeepAlive(groupSource)
}

// SetGroup sets a RadioButton’s group. It should be noted that this does not
// change the layout of your interface in any way, so if you are changing the
// group, it is likely you will need to re-arrange the user interface to reflect
// these changes.
func (radioButton *RadioButton) SetGroup(group []RadioButton) {
	var _arg0 *C.GtkRadioButton // out
	var _arg1 *C.GSList         // out

	_arg0 = (*C.GtkRadioButton)(unsafe.Pointer(radioButton.Native()))
	if group != nil {
		for i := len(group) - 1; i >= 0; i-- {
			src := group[i]
			var dst *C.GtkRadioButton // out
			dst = (*C.GtkRadioButton)(unsafe.Pointer((&src).Native()))
			_arg1 = C.g_slist_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
		}
		defer C.g_slist_free(_arg1)
	}

	C.gtk_radio_button_set_group(_arg0, _arg1)
	runtime.KeepAlive(radioButton)
	runtime.KeepAlive(group)
}

// ConnectGroupChanged: emitted when the group of radio buttons that a radio
// button belongs to changes. This is emitted when a radio button switches from
// being alone to being part of a group of 2 or more buttons, or vice-versa, and
// when a button is moved from one group of 2 or more buttons to a different
// one, but not when the composition of the group that a button belongs to
// changes.
func (r *RadioButton) ConnectGroupChanged(f func()) glib.SignalHandle {
	return r.Connect("group-changed", f)
}
