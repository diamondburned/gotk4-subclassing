// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_header_bar_get_type()), F: marshalHeaderBar},
	})
}

// HeaderBar: gtkHeaderBar is similar to a horizontal Box. It allows children to
// be placed at the start or the end. In addition, it allows a title and
// subtitle to be displayed. The title will be centered with respect to the
// width of the box, even if the children at either side take up different
// amounts of space. The height of the titlebar will be set to provide
// sufficient space for the subtitle, even if none is currently set. If a
// subtitle is not needed, the space reservation can be turned off with
// gtk_header_bar_set_has_subtitle().
//
// GtkHeaderBar can add typical window frame controls, such as minimize,
// maximize and close buttons, or the window icon.
//
// For these reasons, GtkHeaderBar is the natural choice for use as the custom
// titlebar widget of a Window (see gtk_window_set_titlebar()), as it gives
// features typical of titlebars while allowing the addition of child widgets.
type HeaderBar interface {
	Container
	Buildable

	// DecorationLayout gets the decoration layout set with
	// gtk_header_bar_set_decoration_layout().
	DecorationLayout() string
	// HasSubtitle retrieves whether the header bar reserves space for a
	// subtitle, regardless if one is currently set or not.
	HasSubtitle() bool
	// ShowCloseButton returns whether this header bar shows the standard window
	// decorations.
	ShowCloseButton() bool
	// Subtitle retrieves the subtitle of the header. See
	// gtk_header_bar_set_subtitle().
	Subtitle() string
	// Title retrieves the title of the header. See gtk_header_bar_set_title().
	Title() string
	// PackEnd adds @child to @bar, packed with reference to the end of the
	// @bar.
	PackEnd(child Widget)
	// PackStart adds @child to @bar, packed with reference to the start of the
	// @bar.
	PackStart(child Widget)
	// SetCustomTitle sets a custom title for the HeaderBar.
	//
	// The title should help a user identify the current view. This supersedes
	// any title set by gtk_header_bar_set_title() or
	// gtk_header_bar_set_subtitle(). To achieve the same style as the builtin
	// title and subtitle, use the “title” and “subtitle” style classes.
	//
	// You should set the custom title to nil, for the header title label to be
	// visible again.
	SetCustomTitle(titleWidget Widget)
	// SetDecorationLayout sets the decoration layout for this header bar,
	// overriding the Settings:gtk-decoration-layout setting.
	//
	// There can be valid reasons for overriding the setting, such as a header
	// bar design that does not allow for buttons to take room on the right, or
	// only offers room for a single close button. Split header bars are another
	// example for overriding the setting.
	//
	// The format of the string is button names, separated by commas. A colon
	// separates the buttons that should appear on the left from those on the
	// right. Recognized button names are minimize, maximize, close, icon (the
	// window icon) and menu (a menu button for the fallback app menu).
	//
	// For example, “menu:minimize,maximize,close” specifies a menu on the left,
	// and minimize, maximize and close buttons on the right.
	SetDecorationLayout(layout string)
	// SetHasSubtitle sets whether the header bar should reserve space for a
	// subtitle, even if none is currently set.
	SetHasSubtitle(setting bool)
	// SetShowCloseButton sets whether this header bar shows the standard window
	// decorations, including close, maximize, and minimize.
	SetShowCloseButton(setting bool)
	// SetSubtitle sets the subtitle of the HeaderBar. The title should give a
	// user an additional detail to help him identify the current view.
	//
	// Note that GtkHeaderBar by default reserves room for the subtitle, even if
	// none is currently set. If this is not desired, set the
	// HeaderBar:has-subtitle property to false.
	SetSubtitle(subtitle string)
	// SetTitle sets the title of the HeaderBar. The title should help a user
	// identify the current view. A good title should not include the
	// application name.
	SetTitle(title string)
}

// headerBar implements the HeaderBar interface.
type headerBar struct {
	Container
	Buildable
}

var _ HeaderBar = (*headerBar)(nil)

// WrapHeaderBar wraps a GObject to the right type. It is
// primarily used internally.
func WrapHeaderBar(obj *externglib.Object) HeaderBar {
	return HeaderBar{
		Container: WrapContainer(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalHeaderBar(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapHeaderBar(obj), nil
}

// DecorationLayout gets the decoration layout set with
// gtk_header_bar_set_decoration_layout().
func (b headerBar) DecorationLayout() string {
	var _arg0 *C.GtkHeaderBar // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))

	var _cret *C.gchar // in

	_cret = C.gtk_header_bar_get_decoration_layout(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// HasSubtitle retrieves whether the header bar reserves space for a
// subtitle, regardless if one is currently set or not.
func (b headerBar) HasSubtitle() bool {
	var _arg0 *C.GtkHeaderBar // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_header_bar_get_has_subtitle(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// ShowCloseButton returns whether this header bar shows the standard window
// decorations.
func (b headerBar) ShowCloseButton() bool {
	var _arg0 *C.GtkHeaderBar // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_header_bar_get_show_close_button(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Subtitle retrieves the subtitle of the header. See
// gtk_header_bar_set_subtitle().
func (b headerBar) Subtitle() string {
	var _arg0 *C.GtkHeaderBar // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))

	var _cret *C.gchar // in

	_cret = C.gtk_header_bar_get_subtitle(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Title retrieves the title of the header. See gtk_header_bar_set_title().
func (b headerBar) Title() string {
	var _arg0 *C.GtkHeaderBar // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))

	var _cret *C.gchar // in

	_cret = C.gtk_header_bar_get_title(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// PackEnd adds @child to @bar, packed with reference to the end of the
// @bar.
func (b headerBar) PackEnd(child Widget) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_header_bar_pack_end(_arg0, _arg1)
}

// PackStart adds @child to @bar, packed with reference to the start of the
// @bar.
func (b headerBar) PackStart(child Widget) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_header_bar_pack_start(_arg0, _arg1)
}

// SetCustomTitle sets a custom title for the HeaderBar.
//
// The title should help a user identify the current view. This supersedes
// any title set by gtk_header_bar_set_title() or
// gtk_header_bar_set_subtitle(). To achieve the same style as the builtin
// title and subtitle, use the “title” and “subtitle” style classes.
//
// You should set the custom title to nil, for the header title label to be
// visible again.
func (b headerBar) SetCustomTitle(titleWidget Widget) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.GtkWidget    // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(titleWidget.Native()))

	C.gtk_header_bar_set_custom_title(_arg0, _arg1)
}

// SetDecorationLayout sets the decoration layout for this header bar,
// overriding the Settings:gtk-decoration-layout setting.
//
// There can be valid reasons for overriding the setting, such as a header
// bar design that does not allow for buttons to take room on the right, or
// only offers room for a single close button. Split header bars are another
// example for overriding the setting.
//
// The format of the string is button names, separated by commas. A colon
// separates the buttons that should appear on the left from those on the
// right. Recognized button names are minimize, maximize, close, icon (the
// window icon) and menu (a menu button for the fallback app menu).
//
// For example, “menu:minimize,maximize,close” specifies a menu on the left,
// and minimize, maximize and close buttons on the right.
func (b headerBar) SetDecorationLayout(layout string) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.gchar)(C.CString(layout))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_header_bar_set_decoration_layout(_arg0, _arg1)
}

// SetHasSubtitle sets whether the header bar should reserve space for a
// subtitle, even if none is currently set.
func (b headerBar) SetHasSubtitle(setting bool) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	if setting {
		_arg1 = C.gboolean(1)
	}

	C.gtk_header_bar_set_has_subtitle(_arg0, _arg1)
}

// SetShowCloseButton sets whether this header bar shows the standard window
// decorations, including close, maximize, and minimize.
func (b headerBar) SetShowCloseButton(setting bool) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	if setting {
		_arg1 = C.gboolean(1)
	}

	C.gtk_header_bar_set_show_close_button(_arg0, _arg1)
}

// SetSubtitle sets the subtitle of the HeaderBar. The title should give a
// user an additional detail to help him identify the current view.
//
// Note that GtkHeaderBar by default reserves room for the subtitle, even if
// none is currently set. If this is not desired, set the
// HeaderBar:has-subtitle property to false.
func (b headerBar) SetSubtitle(subtitle string) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.gchar)(C.CString(subtitle))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_header_bar_set_subtitle(_arg0, _arg1)
}

// SetTitle sets the title of the HeaderBar. The title should help a user
// identify the current view. A good title should not include the
// application name.
func (b headerBar) SetTitle(title string) {
	var _arg0 *C.GtkHeaderBar // out
	var _arg1 *C.gchar        // out

	_arg0 = (*C.GtkHeaderBar)(unsafe.Pointer(b.Native()))
	_arg1 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_header_bar_set_title(_arg0, _arg1)
}
