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
		{T: externglib.Type(C.gtk_picture_get_type()), F: marshalPicture},
	})
}

// Picture: the Picture widget displays a Paintable. Many convenience functions
// are provided to make pictures simple to use. For example, if you want to load
// an image from a file, and then display that, there’s a convenience function
// to do this:
//
//    GtkWidget *widget;
//    widget = gtk_picture_new_for_filename ("myfile.png");
//
// If the file isn’t loaded successfully, the picture will contain a “broken
// image” icon similar to that used in many web browsers. If you want to handle
// errors in loading the file yourself, for example by displaying an error
// message, then load the image with gdk_texture_new_from_file(), then create
// the Picture with gtk_picture_new_for_paintable().
//
// Sometimes an application will want to avoid depending on external data files,
// such as image files. See the documentation of #GResource for details. In this
// case, gtk_picture_new_for_resource() and gtk_picture_set_resource() should be
// used.
//
// GtkPicture displays an image at its natural size. See Image if you want to
// display a fixed-size image, such as an icon.
//
//
// Sizing the paintable
//
// You can influence how the paintable is displayed inside the Picture. By
// turning off Picture:keep-aspect-ratio you can allow the paintable to get
// stretched. Picture:can-shrink can be unset to make sure that paintables are
// never made smaller than their ideal size - but be careful if you do not know
// the size of the paintable in use (like when displaying user-loaded images).
// This can easily cause the picture to grow larger than the screen. And
// Widget:halign and Widget:valign can be used to make sure the paintable
// doesn't fill all available space but is instead displayed at its original
// size.
//
//
// CSS nodes
//
// GtkPicture has a single CSS node with the name picture.
//
//
// Accessibility
//
// GtkPicture uses the K_ACCESSIBLE_ROLE_IMG role.
type Picture interface {
	Widget
	Accessible
	Buildable
	ConstraintTarget

	// AlternativeText gets the alternative textual description of the picture
	// or returns nil if the picture cannot be described textually.
	AlternativeText(s Picture)
	// CanShrink gets the value set via gtk_picture_set_can_shrink().
	CanShrink(s Picture) bool
	// File gets the #GFile currently displayed if @self is displaying a file.
	// If @self is not displaying a file, for example when
	// gtk_picture_set_paintable() was used, then nil is returned.
	File(s Picture)
	// KeepAspectRatio gets the value set via
	// gtk_picture_set_keep_aspect_ratio().
	KeepAspectRatio(s Picture) bool
	// Paintable gets the Paintable being displayed by the Picture.
	Paintable(s Picture)
	// SetAlternativeText sets an alternative textual description for the
	// picture contents. It is equivalent to the "alt" attribute for images on
	// websites.
	//
	// This text will be made available to accessibility tools.
	//
	// If the picture cannot be described textually, set this property to nil.
	SetAlternativeText(s Picture, alternativeText string)
	// SetCanShrink: if set to true, the @self can be made smaller than its
	// contents. The contents will then be scaled down when rendering.
	//
	// If you want to still force a minimum size manually, consider using
	// gtk_widget_set_size_request().
	//
	// Also of note is that a similar function for growing does not exist
	// because the grow behavior can be controlled via gtk_widget_set_halign()
	// and gtk_widget_set_valign().
	SetCanShrink(s Picture, canShrink bool)
	// SetFile makes @self load and display @file.
	//
	// See gtk_picture_new_for_file() for details.
	SetFile(s Picture, file gio.File)
	// SetFilename makes @self load and display the given @filename.
	//
	// This is a utility function that calls gtk_picture_set_file().
	SetFilename(s Picture, filename string)
	// SetKeepAspectRatio: if set to true, the @self will render its contents
	// according to their aspect ratio. That means that empty space may show up
	// at the top/bottom or left/right of @self.
	//
	// If set to false or if the contents provide no aspect ratio, the contents
	// will be stretched over the picture's whole area.
	SetKeepAspectRatio(s Picture, keepAspectRatio bool)
	// SetPaintable makes @self display the given @paintable. If @paintable is
	// nil, nothing will be displayed.
	//
	// See gtk_picture_new_for_paintable() for details.
	SetPaintable(s Picture, paintable gdk.Paintable)
	// SetPixbuf: see gtk_picture_new_for_pixbuf() for details.
	//
	// This is a utility function that calls gtk_picture_set_paintable(),
	SetPixbuf(s Picture, pixbuf gdkpixbuf.Pixbuf)
	// SetResource makes @self load and display the resource at the given
	// @resource_path.
	//
	// This is a utility function that calls gtk_picture_set_file(),
	SetResource(s Picture, resourcePath string)
}

// picture implements the Picture interface.
type picture struct {
	Widget
	Accessible
	Buildable
	ConstraintTarget
}

var _ Picture = (*picture)(nil)

// WrapPicture wraps a GObject to the right type. It is
// primarily used internally.
func WrapPicture(obj *externglib.Object) Picture {
	return Picture{
		Widget:           WrapWidget(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
	}
}

func marshalPicture(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPicture(obj), nil
}

// NewPicture constructs a class Picture.
func NewPicture() {
	C.gtk_picture_new()
}

// NewPictureForFile constructs a class Picture.
func NewPictureForFile(file gio.File) {
	var arg1 *C.GFile

	arg1 = (*C.GFile)(unsafe.Pointer(file.Native()))

	C.gtk_picture_new_for_file(arg1)
}

// NewPictureForFilename constructs a class Picture.
func NewPictureForFilename(filename string) {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_picture_new_for_filename(arg1)
}

// NewPictureForPaintable constructs a class Picture.
func NewPictureForPaintable(paintable gdk.Paintable) {
	var arg1 *C.GdkPaintable

	arg1 = (*C.GdkPaintable)(unsafe.Pointer(paintable.Native()))

	C.gtk_picture_new_for_paintable(arg1)
}

// NewPictureForPixbuf constructs a class Picture.
func NewPictureForPixbuf(pixbuf gdkpixbuf.Pixbuf) {
	var arg1 *C.GdkPixbuf

	arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gtk_picture_new_for_pixbuf(arg1)
}

// NewPictureForResource constructs a class Picture.
func NewPictureForResource(resourcePath string) {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_picture_new_for_resource(arg1)
}

// AlternativeText gets the alternative textual description of the picture
// or returns nil if the picture cannot be described textually.
func (s picture) AlternativeText(s Picture) {
	var arg0 *C.GtkPicture

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))

	C.gtk_picture_get_alternative_text(arg0)
}

// CanShrink gets the value set via gtk_picture_set_can_shrink().
func (s picture) CanShrink(s Picture) bool {
	var arg0 *C.GtkPicture

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_picture_get_can_shrink(arg0)

	if cret {
		ok = true
	}

	return ok
}

// File gets the #GFile currently displayed if @self is displaying a file.
// If @self is not displaying a file, for example when
// gtk_picture_set_paintable() was used, then nil is returned.
func (s picture) File(s Picture) {
	var arg0 *C.GtkPicture

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))

	C.gtk_picture_get_file(arg0)
}

// KeepAspectRatio gets the value set via
// gtk_picture_set_keep_aspect_ratio().
func (s picture) KeepAspectRatio(s Picture) bool {
	var arg0 *C.GtkPicture

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_picture_get_keep_aspect_ratio(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Paintable gets the Paintable being displayed by the Picture.
func (s picture) Paintable(s Picture) {
	var arg0 *C.GtkPicture

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))

	C.gtk_picture_get_paintable(arg0)
}

// SetAlternativeText sets an alternative textual description for the
// picture contents. It is equivalent to the "alt" attribute for images on
// websites.
//
// This text will be made available to accessibility tools.
//
// If the picture cannot be described textually, set this property to nil.
func (s picture) SetAlternativeText(s Picture, alternativeText string) {
	var arg0 *C.GtkPicture
	var arg1 *C.char

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(alternativeText))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_picture_set_alternative_text(arg0, arg1)
}

// SetCanShrink: if set to true, the @self can be made smaller than its
// contents. The contents will then be scaled down when rendering.
//
// If you want to still force a minimum size manually, consider using
// gtk_widget_set_size_request().
//
// Also of note is that a similar function for growing does not exist
// because the grow behavior can be controlled via gtk_widget_set_halign()
// and gtk_widget_set_valign().
func (s picture) SetCanShrink(s Picture, canShrink bool) {
	var arg0 *C.GtkPicture
	var arg1 C.gboolean

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	if canShrink {
		arg1 = C.gboolean(1)
	}

	C.gtk_picture_set_can_shrink(arg0, arg1)
}

// SetFile makes @self load and display @file.
//
// See gtk_picture_new_for_file() for details.
func (s picture) SetFile(s Picture, file gio.File) {
	var arg0 *C.GtkPicture
	var arg1 *C.GFile

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GFile)(unsafe.Pointer(file.Native()))

	C.gtk_picture_set_file(arg0, arg1)
}

// SetFilename makes @self load and display the given @filename.
//
// This is a utility function that calls gtk_picture_set_file().
func (s picture) SetFilename(s Picture, filename string) {
	var arg0 *C.GtkPicture
	var arg1 *C.char

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_picture_set_filename(arg0, arg1)
}

// SetKeepAspectRatio: if set to true, the @self will render its contents
// according to their aspect ratio. That means that empty space may show up
// at the top/bottom or left/right of @self.
//
// If set to false or if the contents provide no aspect ratio, the contents
// will be stretched over the picture's whole area.
func (s picture) SetKeepAspectRatio(s Picture, keepAspectRatio bool) {
	var arg0 *C.GtkPicture
	var arg1 C.gboolean

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	if keepAspectRatio {
		arg1 = C.gboolean(1)
	}

	C.gtk_picture_set_keep_aspect_ratio(arg0, arg1)
}

// SetPaintable makes @self display the given @paintable. If @paintable is
// nil, nothing will be displayed.
//
// See gtk_picture_new_for_paintable() for details.
func (s picture) SetPaintable(s Picture, paintable gdk.Paintable) {
	var arg0 *C.GtkPicture
	var arg1 *C.GdkPaintable

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GdkPaintable)(unsafe.Pointer(paintable.Native()))

	C.gtk_picture_set_paintable(arg0, arg1)
}

// SetPixbuf: see gtk_picture_new_for_pixbuf() for details.
//
// This is a utility function that calls gtk_picture_set_paintable(),
func (s picture) SetPixbuf(s Picture, pixbuf gdkpixbuf.Pixbuf) {
	var arg0 *C.GtkPicture
	var arg1 *C.GdkPixbuf

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gtk_picture_set_pixbuf(arg0, arg1)
}

// SetResource makes @self load and display the resource at the given
// @resource_path.
//
// This is a utility function that calls gtk_picture_set_file(),
func (s picture) SetResource(s Picture, resourcePath string) {
	var arg0 *C.GtkPicture
	var arg1 *C.char

	arg0 = (*C.GtkPicture)(unsafe.Pointer(s.Native()))
	arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_picture_set_resource(arg0, arg1)
}