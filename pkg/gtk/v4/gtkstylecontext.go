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
		{T: externglib.Type(C.gtk_style_context_get_type()), F: marshalStyleContext},
	})
}

// StyleContext is an object that stores styling information affecting a widget.
//
// In order to construct the final style information, StyleContext queries
// information from all attached StyleProviders. Style providers can be either
// attached explicitly to the context through gtk_style_context_add_provider(),
// or to the display through gtk_style_context_add_provider_for_display(). The
// resulting style is a combination of all providers’ information in priority
// order.
//
// For GTK widgets, any StyleContext returned by gtk_widget_get_style_context()
// will already have a Display and RTL/LTR information set. The style context
// will also be updated automatically if any of these settings change on the
// widget.
//
//
// Style Classes
//
// Widgets can add style classes to their context, which can be used to
// associate different styles by class. The documentation for individual widgets
// lists which style classes it uses itself, and which style classes may be
// added by applications to affect their appearance.
//
// GTK defines macros for a number of style classes.
//
//
// Custom styling in UI libraries and applications
//
// If you are developing a library with custom Widgets that render differently
// than standard components, you may need to add a StyleProvider yourself with
// the GTK_STYLE_PROVIDER_PRIORITY_FALLBACK priority, either a CssProvider or a
// custom object implementing the StyleProvider interface. This way themes may
// still attempt to style your UI elements in a different way if needed so.
//
// If you are using custom styling on an applications, you probably want then to
// make your style information prevail to the theme’s, so you must use a
// StyleProvider with the GTK_STYLE_PROVIDER_PRIORITY_APPLICATION priority, keep
// in mind that the user settings in `XDG_CONFIG_HOME/gtk-4.0/gtk.css` will
// still take precedence over your changes, as it uses the
// GTK_STYLE_PROVIDER_PRIORITY_USER priority.
type StyleContext interface {
	gextras.Objector

	// AddClass adds a style class to @context, so later uses of the style
	// context will make use of this new class for styling.
	//
	// In the CSS file format, a Entry defining a “search” class, would be
	// matched by:
	//
	// |[ <!-- language="CSS" --> entry.search { ... } ]|
	//
	// While any widget defining a “search” class would be matched by: |[ <!--
	// language="CSS" --> .search { ... } ]|
	AddClass(c StyleContext, className string)
	// AddProvider adds a style provider to @context, to be used in style
	// construction. Note that a style provider added by this function only
	// affects the style of the widget to which @context belongs. If you want to
	// affect the style of all widgets, use
	// gtk_style_context_add_provider_for_display().
	//
	// Note: If both priorities are the same, a StyleProvider added through this
	// function takes precedence over another added through
	// gtk_style_context_add_provider_for_display().
	AddProvider(c StyleContext, provider StyleProvider, priority uint)
	// Border gets the border for a given state as a Border.
	Border(c StyleContext) *Border
	// Color gets the foreground color for a given state.
	Color(c StyleContext) *gdk.RGBA
	// Display returns the Display to which @context is attached.
	Display(c StyleContext)
	// Margin gets the margin for a given state as a Border.
	Margin(c StyleContext) *Border
	// Padding gets the padding for a given state as a Border.
	Padding(c StyleContext) *Border
	// Scale returns the scale used for assets.
	Scale(c StyleContext)
	// State returns the state used for style matching.
	//
	// This method should only be used to retrieve the StateFlags to pass to
	// StyleContext methods, like gtk_style_context_get_padding(). If you need
	// to retrieve the current state of a Widget, use
	// gtk_widget_get_state_flags().
	State(c StyleContext)
	// HasClass returns true if @context currently has defined the given class
	// name.
	HasClass(c StyleContext, className string) bool
	// LookupColor looks up and resolves a color name in the @context color map.
	LookupColor(c StyleContext, colorName string) (color *gdk.RGBA, ok bool)
	// RemoveClass removes @class_name from @context.
	RemoveClass(c StyleContext, className string)
	// RemoveProvider removes @provider from the style providers list in
	// @context.
	RemoveProvider(c StyleContext, provider StyleProvider)
	// Restore restores @context state to a previous stage. See
	// gtk_style_context_save().
	Restore(c StyleContext)
	// Save saves the @context state, so temporary modifications done through
	// gtk_style_context_add_class(), gtk_style_context_remove_class(),
	// gtk_style_context_set_state(), etc. can quickly be reverted in one go
	// through gtk_style_context_restore().
	//
	// The matching call to gtk_style_context_restore() must be done before GTK
	// returns to the main loop.
	Save(c StyleContext)
	// SetDisplay attaches @context to the given display.
	//
	// The display is used to add style information from “global” style
	// providers, such as the display's Settings instance.
	//
	// If you are using a StyleContext returned from
	// gtk_widget_get_style_context(), you do not need to call this yourself.
	SetDisplay(c StyleContext, display gdk.Display)
	// SetScale sets the scale to use when getting image assets for the style.
	SetScale(c StyleContext, scale int)
	// SetState sets the state to be used for style matching.
	SetState(c StyleContext, flags StateFlags)
	// String converts the style context into a string representation.
	//
	// The string representation always includes information about the name,
	// state, id, visibility and style classes of the CSS node that is backing
	// @context. Depending on the flags, more information may be included.
	//
	// This function is intended for testing and debugging of the CSS
	// implementation in GTK. There are no guarantees about the format of the
	// returned string, it may change.
	String(c StyleContext, flags StyleContextPrintFlags)
}

// styleContext implements the StyleContext interface.
type styleContext struct {
	gextras.Objector
}

var _ StyleContext = (*styleContext)(nil)

// WrapStyleContext wraps a GObject to the right type. It is
// primarily used internally.
func WrapStyleContext(obj *externglib.Object) StyleContext {
	return StyleContext{
		Objector: obj,
	}
}

func marshalStyleContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapStyleContext(obj), nil
}

// AddClass adds a style class to @context, so later uses of the style
// context will make use of this new class for styling.
//
// In the CSS file format, a Entry defining a “search” class, would be
// matched by:
//
// |[ <!-- language="CSS" --> entry.search { ... } ]|
//
// While any widget defining a “search” class would be matched by: |[ <!--
// language="CSS" --> .search { ... } ]|
func (c styleContext) AddClass(c StyleContext, className string) {
	var arg0 *C.GtkStyleContext
	var arg1 *C.char

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.char)(C.CString(className))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_style_context_add_class(arg0, arg1)
}

// AddProvider adds a style provider to @context, to be used in style
// construction. Note that a style provider added by this function only
// affects the style of the widget to which @context belongs. If you want to
// affect the style of all widgets, use
// gtk_style_context_add_provider_for_display().
//
// Note: If both priorities are the same, a StyleProvider added through this
// function takes precedence over another added through
// gtk_style_context_add_provider_for_display().
func (c styleContext) AddProvider(c StyleContext, provider StyleProvider, priority uint) {
	var arg0 *C.GtkStyleContext
	var arg1 *C.GtkStyleProvider
	var arg2 C.guint

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkStyleProvider)(unsafe.Pointer(provider.Native()))
	arg2 = C.guint(priority)

	C.gtk_style_context_add_provider(arg0, arg1, arg2)
}

// Border gets the border for a given state as a Border.
func (c styleContext) Border(c StyleContext) *Border {
	var arg0 *C.GtkStyleContext

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	var arg1 C.GtkBorder
	var border *Border

	C.gtk_style_context_get_border(arg0, &arg1)

	border = WrapBorder(unsafe.Pointer(&arg1))

	return border
}

// Color gets the foreground color for a given state.
func (c styleContext) Color(c StyleContext) *gdk.RGBA {
	var arg0 *C.GtkStyleContext

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	var arg1 C.GdkRGBA
	var color *gdk.RGBA

	C.gtk_style_context_get_color(arg0, &arg1)

	color = gdk.WrapRGBA(unsafe.Pointer(&arg1))

	return color
}

// Display returns the Display to which @context is attached.
func (c styleContext) Display(c StyleContext) {
	var arg0 *C.GtkStyleContext

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	C.gtk_style_context_get_display(arg0)
}

// Margin gets the margin for a given state as a Border.
func (c styleContext) Margin(c StyleContext) *Border {
	var arg0 *C.GtkStyleContext

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	var arg1 C.GtkBorder
	var margin *Border

	C.gtk_style_context_get_margin(arg0, &arg1)

	margin = WrapBorder(unsafe.Pointer(&arg1))

	return margin
}

// Padding gets the padding for a given state as a Border.
func (c styleContext) Padding(c StyleContext) *Border {
	var arg0 *C.GtkStyleContext

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	var arg1 C.GtkBorder
	var padding *Border

	C.gtk_style_context_get_padding(arg0, &arg1)

	padding = WrapBorder(unsafe.Pointer(&arg1))

	return padding
}

// Scale returns the scale used for assets.
func (c styleContext) Scale(c StyleContext) {
	var arg0 *C.GtkStyleContext

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	C.gtk_style_context_get_scale(arg0)
}

// State returns the state used for style matching.
//
// This method should only be used to retrieve the StateFlags to pass to
// StyleContext methods, like gtk_style_context_get_padding(). If you need
// to retrieve the current state of a Widget, use
// gtk_widget_get_state_flags().
func (c styleContext) State(c StyleContext) {
	var arg0 *C.GtkStyleContext

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	C.gtk_style_context_get_state(arg0)
}

// HasClass returns true if @context currently has defined the given class
// name.
func (c styleContext) HasClass(c StyleContext, className string) bool {
	var arg0 *C.GtkStyleContext
	var arg1 *C.char

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.char)(C.CString(className))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_style_context_has_class(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// LookupColor looks up and resolves a color name in the @context color map.
func (c styleContext) LookupColor(c StyleContext, colorName string) (color *gdk.RGBA, ok bool) {
	var arg0 *C.GtkStyleContext
	var arg1 *C.char

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.char)(C.CString(colorName))
	defer C.free(unsafe.Pointer(arg1))

	var arg2 C.GdkRGBA
	var color *gdk.RGBA
	var cret C.gboolean
	var ok bool

	cret = C.gtk_style_context_lookup_color(arg0, arg1, &arg2)

	color = gdk.WrapRGBA(unsafe.Pointer(&arg2))
	if cret {
		ok = true
	}

	return color, ok
}

// RemoveClass removes @class_name from @context.
func (c styleContext) RemoveClass(c StyleContext, className string) {
	var arg0 *C.GtkStyleContext
	var arg1 *C.char

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.char)(C.CString(className))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_style_context_remove_class(arg0, arg1)
}

// RemoveProvider removes @provider from the style providers list in
// @context.
func (c styleContext) RemoveProvider(c StyleContext, provider StyleProvider) {
	var arg0 *C.GtkStyleContext
	var arg1 *C.GtkStyleProvider

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GtkStyleProvider)(unsafe.Pointer(provider.Native()))

	C.gtk_style_context_remove_provider(arg0, arg1)
}

// Restore restores @context state to a previous stage. See
// gtk_style_context_save().
func (c styleContext) Restore(c StyleContext) {
	var arg0 *C.GtkStyleContext

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	C.gtk_style_context_restore(arg0)
}

// Save saves the @context state, so temporary modifications done through
// gtk_style_context_add_class(), gtk_style_context_remove_class(),
// gtk_style_context_set_state(), etc. can quickly be reverted in one go
// through gtk_style_context_restore().
//
// The matching call to gtk_style_context_restore() must be done before GTK
// returns to the main loop.
func (c styleContext) Save(c StyleContext) {
	var arg0 *C.GtkStyleContext

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	C.gtk_style_context_save(arg0)
}

// SetDisplay attaches @context to the given display.
//
// The display is used to add style information from “global” style
// providers, such as the display's Settings instance.
//
// If you are using a StyleContext returned from
// gtk_widget_get_style_context(), you do not need to call this yourself.
func (c styleContext) SetDisplay(c StyleContext, display gdk.Display) {
	var arg0 *C.GtkStyleContext
	var arg1 *C.GdkDisplay

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gtk_style_context_set_display(arg0, arg1)
}

// SetScale sets the scale to use when getting image assets for the style.
func (c styleContext) SetScale(c StyleContext, scale int) {
	var arg0 *C.GtkStyleContext
	var arg1 C.int

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	arg1 = C.int(scale)

	C.gtk_style_context_set_scale(arg0, arg1)
}

// SetState sets the state to be used for style matching.
func (c styleContext) SetState(c StyleContext, flags StateFlags) {
	var arg0 *C.GtkStyleContext
	var arg1 C.GtkStateFlags

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	arg1 = (C.GtkStateFlags)(flags)

	C.gtk_style_context_set_state(arg0, arg1)
}

// String converts the style context into a string representation.
//
// The string representation always includes information about the name,
// state, id, visibility and style classes of the CSS node that is backing
// @context. Depending on the flags, more information may be included.
//
// This function is intended for testing and debugging of the CSS
// implementation in GTK. There are no guarantees about the format of the
// returned string, it may change.
func (c styleContext) String(c StyleContext, flags StyleContextPrintFlags) {
	var arg0 *C.GtkStyleContext
	var arg1 C.GtkStyleContextPrintFlags

	arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	arg1 = (C.GtkStyleContextPrintFlags)(flags)

	C.gtk_style_context_to_string(arg0, arg1)
}
