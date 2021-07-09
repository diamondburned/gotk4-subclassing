// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/box"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_style_context_print_flags_get_type()), F: marshalStyleContextPrintFlags},
		{T: externglib.Type(C.gtk_style_context_get_type()), F: marshalStyleContext},
	})
}

// StyleContextPrintFlags flags that modify the behavior of
// gtk_style_context_to_string(). New values may be added to this enumeration.
type StyleContextPrintFlags int

const (
	StyleContextPrintFlagsNone StyleContextPrintFlags = 0b0
	// StyleContextPrintFlagsRecurse: print the entire tree of CSS nodes
	// starting at the style context's node
	StyleContextPrintFlagsRecurse StyleContextPrintFlags = 0b1
	// StyleContextPrintFlagsShowStyle: show the values of the CSS properties
	// for each node
	StyleContextPrintFlagsShowStyle StyleContextPrintFlags = 0b10
)

func marshalStyleContextPrintFlags(p uintptr) (interface{}, error) {
	return StyleContextPrintFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// StyleContextOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type StyleContextOverrider interface {
	Changed()
}

// StyleContext is an object that stores styling information affecting a widget
// defined by WidgetPath.
//
// In order to construct the final style information, StyleContext queries
// information from all attached StyleProviders. Style providers can be either
// attached explicitly to the context through gtk_style_context_add_provider(),
// or to the screen through gtk_style_context_add_provider_for_screen(). The
// resulting style is a combination of all providers’ information in priority
// order.
//
// For GTK+ widgets, any StyleContext returned by gtk_widget_get_style_context()
// will already have a WidgetPath, a Screen and RTL/LTR information set. The
// style context will also be updated automatically if any of these settings
// change on the widget.
//
// If you are using the theming layer standalone, you will need to set a widget
// path and a screen yourself to the created style context through
// gtk_style_context_set_path() and possibly gtk_style_context_set_screen(). See
// the “Foreign drawing“ example in gtk3-demo.
//
//
// Style Classes
//
// Widgets can add style classes to their context, which can be used to
// associate different styles by class. The documentation for individual widgets
// lists which style classes it uses itself, and which style classes may be
// added by applications to affect their appearance.
//
// GTK+ defines macros for a number of style classes.
//
//
// Style Regions
//
// Widgets can also add regions with flags to their context. This feature is
// deprecated and will be removed in a future GTK+ update. Please use style
// classes instead.
//
// GTK+ defines macros for a number of style regions.
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
// in mind that the user settings in `XDG_CONFIG_HOME/gtk-3.0/gtk.css` will
// still take precedence over your changes, as it uses the
// GTK_STYLE_PROVIDER_PRIORITY_USER priority.
type StyleContext interface {
	gextras.Objector

	// AddClass adds a style class to @context, so posterior calls to
	// gtk_style_context_get() or any of the gtk_render_*() functions will make
	// use of this new class for styling.
	//
	// In the CSS file format, a Entry defining a “search” class, would be
	// matched by:
	//
	// |[ <!-- language="CSS" --> entry.search { ... } ]|
	//
	// While any widget defining a “search” class would be matched by: |[ <!--
	// language="CSS" --> .search { ... } ]|
	AddClass(className string)
	// AddProvider adds a style provider to @context, to be used in style
	// construction. Note that a style provider added by this function only
	// affects the style of the widget to which @context belongs. If you want to
	// affect the style of all widgets, use
	// gtk_style_context_add_provider_for_screen().
	//
	// Note: If both priorities are the same, a StyleProvider added through this
	// function takes precedence over another added through
	// gtk_style_context_add_provider_for_screen().
	AddProvider(provider StyleProvider, priority uint)
	// CancelAnimations stops all running animations for @region_id and all
	// animatable regions underneath.
	//
	// A nil @region_id will stop all ongoing animations in @context, when
	// dealing with a StyleContext obtained through
	// gtk_widget_get_style_context(), this is normally done for you in all
	// circumstances you would expect all widget to be stopped, so this should
	// be only used in complex widgets with different animatable regions.
	//
	// Deprecated: since version 3.6.
	CancelAnimations(regionId interface{})
	// Direction returns the widget direction used for rendering.
	//
	// Deprecated: since version 3.8.
	Direction() TextDirection
	// FrameClock returns the FrameClock to which @context is attached.
	FrameClock() *gdk.FrameClockClass
	// JunctionSides returns the sides where rendered elements connect visually
	// with others.
	JunctionSides() JunctionSides
	// Parent gets the parent context set via gtk_style_context_set_parent().
	// See that function for details.
	Parent() *StyleContextClass
	// Path returns the widget path used for style matching.
	Path() *WidgetPath
	// Scale returns the scale used for assets.
	Scale() int
	// Screen returns the Screen to which @context is attached.
	Screen() *gdk.ScreenClass
	// Section queries the location in the CSS where @property was defined for
	// the current @context. Note that the state to be queried is taken from
	// gtk_style_context_get_state().
	//
	// If the location is not available, nil will be returned. The location
	// might not be available for various reasons, such as the property being
	// overridden, @property not naming a supported CSS property or tracking of
	// definitions being disabled for performance reasons.
	//
	// Shorthand CSS properties cannot be queried for a location and will always
	// return nil.
	Section(property string) *CSSSection
	// State returns the state used for style matching.
	//
	// This method should only be used to retrieve the StateFlags to pass to
	// StyleContext methods, like gtk_style_context_get_padding(). If you need
	// to retrieve the current state of a Widget, use
	// gtk_widget_get_state_flags().
	State() StateFlags
	// StyleProperty gets the value for a widget style property.
	//
	// When @value is no longer needed, g_value_unset() must be called to free
	// any allocated memory.
	StyleProperty(propertyName string, value *externglib.Value)
	// HasClass returns true if @context currently has defined the given class
	// name.
	HasClass(className string) bool
	// HasRegion returns true if @context has the region defined. If
	// @flags_return is not nil, it is set to the flags affecting the region.
	//
	// Deprecated: since version 3.14.
	HasRegion(regionName string) (RegionFlags, bool)
	// Invalidate invalidates @context style information, so it will be
	// reconstructed again. It is useful if you modify the @context and need the
	// new information immediately.
	//
	// Deprecated: since version 3.12.
	Invalidate()
	// LookupColor looks up and resolves a color name in the @context color map.
	LookupColor(colorName string) (gdk.RGBA, bool)
	// LookupIconSet looks up @stock_id in the icon factories associated to
	// @context and the default icon factory, returning an icon set if found,
	// otherwise nil.
	//
	// Deprecated: since version 3.10.
	LookupIconSet(stockId string) *IconSet
	// PopAnimatableRegion pops an animatable region from @context. See
	// gtk_style_context_push_animatable_region().
	//
	// Deprecated: since version 3.6.
	PopAnimatableRegion()
	// PushAnimatableRegion pushes an animatable region, so all further
	// gtk_render_*() calls between this call and the following
	// gtk_style_context_pop_animatable_region() will potentially show
	// transition animations for this region if
	// gtk_style_context_notify_state_change() is called for a given state, and
	// the current theme/style defines transition animations for state changes.
	//
	// The @region_id used must be unique in @context so the themes can uniquely
	// identify rendered elements subject to a state transition.
	//
	// Deprecated: since version 3.6.
	PushAnimatableRegion(regionId interface{})
	// RemoveClass removes @class_name from @context.
	RemoveClass(className string)
	// RemoveProvider removes @provider from the style providers list in
	// @context.
	RemoveProvider(provider StyleProvider)
	// RemoveRegion removes a region from @context.
	//
	// Deprecated: since version 3.14.
	RemoveRegion(regionName string)
	// Restore restores @context state to a previous stage. See
	// gtk_style_context_save().
	Restore()
	// Save saves the @context state, so temporary modifications done through
	// gtk_style_context_add_class(), gtk_style_context_remove_class(),
	// gtk_style_context_set_state(), etc. can quickly be reverted in one go
	// through gtk_style_context_restore().
	//
	// The matching call to gtk_style_context_restore() must be done before GTK
	// returns to the main loop.
	Save()
	// ScrollAnimations: this function is analogous to gdk_window_scroll(), and
	// should be called together with it so the invalidation areas for any
	// ongoing animation are scrolled together with it.
	//
	// Deprecated: since version 3.6.
	ScrollAnimations(window gdk.Window, dx int, dy int)
	// SetBackground sets the background of @window to the background pattern or
	// color specified in @context for its current state.
	//
	// Deprecated: since version 3.18.
	SetBackground(window gdk.Window)
	// SetFrameClock attaches @context to the given frame clock.
	//
	// The frame clock is used for the timing of animations.
	//
	// If you are using a StyleContext returned from
	// gtk_widget_get_style_context(), you do not need to call this yourself.
	SetFrameClock(frameClock gdk.FrameClock)
	// SetParent sets the parent style context for @context. The parent style
	// context is used to implement inheritance
	// (http://www.w3.org/TR/css3-cascade/#inheritance) of properties.
	//
	// If you are using a StyleContext returned from
	// gtk_widget_get_style_context(), the parent will be set for you.
	SetParent(parent StyleContext)
	// SetPath sets the WidgetPath used for style matching. As a consequence,
	// the style will be regenerated to match the new given path.
	//
	// If you are using a StyleContext returned from
	// gtk_widget_get_style_context(), you do not need to call this yourself.
	SetPath(path *WidgetPath)
	// SetScale sets the scale to use when getting image assets for the style.
	SetScale(scale int)
	// SetScreen attaches @context to the given screen.
	//
	// The screen is used to add style information from “global” style
	// providers, such as the screen’s Settings instance.
	//
	// If you are using a StyleContext returned from
	// gtk_widget_get_style_context(), you do not need to call this yourself.
	SetScreen(screen gdk.Screen)
}

// StyleContextClass implements the StyleContext interface.
type StyleContextClass struct {
	*externglib.Object
}

var _ StyleContext = (*StyleContextClass)(nil)

func wrapStyleContext(obj *externglib.Object) StyleContext {
	return &StyleContextClass{
		Object: obj,
	}
}

func marshalStyleContext(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapStyleContext(obj), nil
}

// NewStyleContext creates a standalone StyleContext, this style context won’t
// be attached to any widget, so you may want to call
// gtk_style_context_set_path() yourself.
//
// This function is only useful when using the theming layer separated from
// GTK+, if you are using StyleContext to theme Widgets, use
// gtk_widget_get_style_context() in order to get a style context ready to theme
// the widget.
func NewStyleContext() *StyleContextClass {
	var _cret *C.GtkStyleContext // in

	_cret = C.gtk_style_context_new()

	var _styleContext *StyleContextClass // out

	_styleContext = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*StyleContextClass)

	return _styleContext
}

// AddClass adds a style class to @context, so posterior calls to
// gtk_style_context_get() or any of the gtk_render_*() functions will make use
// of this new class for styling.
//
// In the CSS file format, a Entry defining a “search” class, would be matched
// by:
//
// |[ <!-- language="CSS" --> entry.search { ... } ]|
//
// While any widget defining a “search” class would be matched by: |[ <!--
// language="CSS" --> .search { ... } ]|
func (c *StyleContextClass) AddClass(className string) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(className))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_style_context_add_class(_arg0, _arg1)
}

// AddProvider adds a style provider to @context, to be used in style
// construction. Note that a style provider added by this function only affects
// the style of the widget to which @context belongs. If you want to affect the
// style of all widgets, use gtk_style_context_add_provider_for_screen().
//
// Note: If both priorities are the same, a StyleProvider added through this
// function takes precedence over another added through
// gtk_style_context_add_provider_for_screen().
func (c *StyleContextClass) AddProvider(provider StyleProvider, priority uint) {
	var _arg0 *C.GtkStyleContext  // out
	var _arg1 *C.GtkStyleProvider // out
	var _arg2 C.guint             // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkStyleProvider)(unsafe.Pointer(provider.Native()))
	_arg2 = C.guint(priority)

	C.gtk_style_context_add_provider(_arg0, _arg1, _arg2)
}

// CancelAnimations stops all running animations for @region_id and all
// animatable regions underneath.
//
// A nil @region_id will stop all ongoing animations in @context, when dealing
// with a StyleContext obtained through gtk_widget_get_style_context(), this is
// normally done for you in all circumstances you would expect all widget to be
// stopped, so this should be only used in complex widgets with different
// animatable regions.
//
// Deprecated: since version 3.6.
func (c *StyleContextClass) CancelAnimations(regionId interface{}) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 C.gpointer         // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	_arg1 = (C.gpointer)(box.Assign(regionId))

	C.gtk_style_context_cancel_animations(_arg0, _arg1)
}

// Direction returns the widget direction used for rendering.
//
// Deprecated: since version 3.8.
func (c *StyleContextClass) Direction() TextDirection {
	var _arg0 *C.GtkStyleContext // out
	var _cret C.GtkTextDirection // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_style_context_get_direction(_arg0)

	var _textDirection TextDirection // out

	_textDirection = (TextDirection)(_cret)

	return _textDirection
}

// FrameClock returns the FrameClock to which @context is attached.
func (c *StyleContextClass) FrameClock() *gdk.FrameClockClass {
	var _arg0 *C.GtkStyleContext // out
	var _cret *C.GdkFrameClock   // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_style_context_get_frame_clock(_arg0)

	var _frameClock *gdk.FrameClockClass // out

	_frameClock = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gdk.FrameClockClass)

	return _frameClock
}

// JunctionSides returns the sides where rendered elements connect visually with
// others.
func (c *StyleContextClass) JunctionSides() JunctionSides {
	var _arg0 *C.GtkStyleContext // out
	var _cret C.GtkJunctionSides // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_style_context_get_junction_sides(_arg0)

	var _junctionSides JunctionSides // out

	_junctionSides = (JunctionSides)(_cret)

	return _junctionSides
}

// Parent gets the parent context set via gtk_style_context_set_parent(). See
// that function for details.
func (c *StyleContextClass) Parent() *StyleContextClass {
	var _arg0 *C.GtkStyleContext // out
	var _cret *C.GtkStyleContext // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_style_context_get_parent(_arg0)

	var _styleContext *StyleContextClass // out

	_styleContext = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*StyleContextClass)

	return _styleContext
}

// Path returns the widget path used for style matching.
func (c *StyleContextClass) Path() *WidgetPath {
	var _arg0 *C.GtkStyleContext // out
	var _cret *C.GtkWidgetPath   // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_style_context_get_path(_arg0)

	var _widgetPath *WidgetPath // out

	_widgetPath = (*WidgetPath)(unsafe.Pointer(_cret))
	C.gtk_widget_path_ref(_cret)
	runtime.SetFinalizer(_widgetPath, func(v *WidgetPath) {
		C.gtk_widget_path_unref((*C.GtkWidgetPath)(unsafe.Pointer(v)))
	})

	return _widgetPath
}

// Scale returns the scale used for assets.
func (c *StyleContextClass) Scale() int {
	var _arg0 *C.GtkStyleContext // out
	var _cret C.gint             // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_style_context_get_scale(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// Screen returns the Screen to which @context is attached.
func (c *StyleContextClass) Screen() *gdk.ScreenClass {
	var _arg0 *C.GtkStyleContext // out
	var _cret *C.GdkScreen       // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_style_context_get_screen(_arg0)

	var _screen *gdk.ScreenClass // out

	_screen = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*gdk.ScreenClass)

	return _screen
}

// Section queries the location in the CSS where @property was defined for the
// current @context. Note that the state to be queried is taken from
// gtk_style_context_get_state().
//
// If the location is not available, nil will be returned. The location might
// not be available for various reasons, such as the property being overridden,
// @property not naming a supported CSS property or tracking of definitions
// being disabled for performance reasons.
//
// Shorthand CSS properties cannot be queried for a location and will always
// return nil.
func (c *StyleContextClass) Section(property string) *CSSSection {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 *C.gchar           // out
	var _cret *C.GtkCssSection   // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(property))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_style_context_get_section(_arg0, _arg1)

	var _cssSection *CSSSection // out

	_cssSection = (*CSSSection)(unsafe.Pointer(_cret))
	C.gtk_css_section_ref(_cret)
	runtime.SetFinalizer(_cssSection, func(v *CSSSection) {
		C.gtk_css_section_unref((*C.GtkCssSection)(unsafe.Pointer(v)))
	})

	return _cssSection
}

// State returns the state used for style matching.
//
// This method should only be used to retrieve the StateFlags to pass to
// StyleContext methods, like gtk_style_context_get_padding(). If you need to
// retrieve the current state of a Widget, use gtk_widget_get_state_flags().
func (c *StyleContextClass) State() StateFlags {
	var _arg0 *C.GtkStyleContext // out
	var _cret C.GtkStateFlags    // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	_cret = C.gtk_style_context_get_state(_arg0)

	var _stateFlags StateFlags // out

	_stateFlags = (StateFlags)(_cret)

	return _stateFlags
}

// StyleProperty gets the value for a widget style property.
//
// When @value is no longer needed, g_value_unset() must be called to free any
// allocated memory.
func (c *StyleContextClass) StyleProperty(propertyName string, value *externglib.Value) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 *C.gchar           // out
	var _arg2 *C.GValue          // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GValue)(unsafe.Pointer(&value.GValue))

	C.gtk_style_context_get_style_property(_arg0, _arg1, _arg2)
}

// HasClass returns true if @context currently has defined the given class name.
func (c *StyleContextClass) HasClass(className string) bool {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 *C.gchar           // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(className))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_style_context_has_class(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasRegion returns true if @context has the region defined. If @flags_return
// is not nil, it is set to the flags affecting the region.
//
// Deprecated: since version 3.14.
func (c *StyleContextClass) HasRegion(regionName string) (RegionFlags, bool) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 *C.gchar           // out
	var _arg2 C.GtkRegionFlags   // in
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(regionName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_style_context_has_region(_arg0, _arg1, &_arg2)

	var _flagsReturn RegionFlags // out
	var _ok bool                 // out

	_flagsReturn = (RegionFlags)(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _flagsReturn, _ok
}

// Invalidate invalidates @context style information, so it will be
// reconstructed again. It is useful if you modify the @context and need the new
// information immediately.
//
// Deprecated: since version 3.12.
func (c *StyleContextClass) Invalidate() {
	var _arg0 *C.GtkStyleContext // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	C.gtk_style_context_invalidate(_arg0)
}

// LookupColor looks up and resolves a color name in the @context color map.
func (c *StyleContextClass) LookupColor(colorName string) (gdk.RGBA, bool) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 *C.gchar           // out
	var _arg2 C.GdkRGBA          // in
	var _cret C.gboolean         // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(colorName))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_style_context_lookup_color(_arg0, _arg1, &_arg2)

	var _color gdk.RGBA // out
	var _ok bool        // out

	_color = *(*gdk.RGBA)(unsafe.Pointer((&_arg2)))
	if _cret != 0 {
		_ok = true
	}

	return _color, _ok
}

// LookupIconSet looks up @stock_id in the icon factories associated to @context
// and the default icon factory, returning an icon set if found, otherwise nil.
//
// Deprecated: since version 3.10.
func (c *StyleContextClass) LookupIconSet(stockId string) *IconSet {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 *C.gchar           // out
	var _cret *C.GtkIconSet      // in

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_style_context_lookup_icon_set(_arg0, _arg1)

	var _iconSet *IconSet // out

	_iconSet = (*IconSet)(unsafe.Pointer(_cret))
	C.gtk_icon_set_ref(_cret)
	runtime.SetFinalizer(_iconSet, func(v *IconSet) {
		C.gtk_icon_set_unref((*C.GtkIconSet)(unsafe.Pointer(v)))
	})

	return _iconSet
}

// PopAnimatableRegion pops an animatable region from @context. See
// gtk_style_context_push_animatable_region().
//
// Deprecated: since version 3.6.
func (c *StyleContextClass) PopAnimatableRegion() {
	var _arg0 *C.GtkStyleContext // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	C.gtk_style_context_pop_animatable_region(_arg0)
}

// PushAnimatableRegion pushes an animatable region, so all further
// gtk_render_*() calls between this call and the following
// gtk_style_context_pop_animatable_region() will potentially show transition
// animations for this region if gtk_style_context_notify_state_change() is
// called for a given state, and the current theme/style defines transition
// animations for state changes.
//
// The @region_id used must be unique in @context so the themes can uniquely
// identify rendered elements subject to a state transition.
//
// Deprecated: since version 3.6.
func (c *StyleContextClass) PushAnimatableRegion(regionId interface{}) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 C.gpointer         // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	_arg1 = (C.gpointer)(box.Assign(regionId))

	C.gtk_style_context_push_animatable_region(_arg0, _arg1)
}

// RemoveClass removes @class_name from @context.
func (c *StyleContextClass) RemoveClass(className string) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(className))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_style_context_remove_class(_arg0, _arg1)
}

// RemoveProvider removes @provider from the style providers list in @context.
func (c *StyleContextClass) RemoveProvider(provider StyleProvider) {
	var _arg0 *C.GtkStyleContext  // out
	var _arg1 *C.GtkStyleProvider // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkStyleProvider)(unsafe.Pointer(provider.Native()))

	C.gtk_style_context_remove_provider(_arg0, _arg1)
}

// RemoveRegion removes a region from @context.
//
// Deprecated: since version 3.14.
func (c *StyleContextClass) RemoveRegion(regionName string) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 *C.gchar           // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.gchar)(C.CString(regionName))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_style_context_remove_region(_arg0, _arg1)
}

// Restore restores @context state to a previous stage. See
// gtk_style_context_save().
func (c *StyleContextClass) Restore() {
	var _arg0 *C.GtkStyleContext // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	C.gtk_style_context_restore(_arg0)
}

// Save saves the @context state, so temporary modifications done through
// gtk_style_context_add_class(), gtk_style_context_remove_class(),
// gtk_style_context_set_state(), etc. can quickly be reverted in one go through
// gtk_style_context_restore().
//
// The matching call to gtk_style_context_restore() must be done before GTK
// returns to the main loop.
func (c *StyleContextClass) Save() {
	var _arg0 *C.GtkStyleContext // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))

	C.gtk_style_context_save(_arg0)
}

// ScrollAnimations: this function is analogous to gdk_window_scroll(), and
// should be called together with it so the invalidation areas for any ongoing
// animation are scrolled together with it.
//
// Deprecated: since version 3.6.
func (c *StyleContextClass) ScrollAnimations(window gdk.Window, dx int, dy int) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 *C.GdkWindow       // out
	var _arg2 C.gint             // out
	var _arg3 C.gint             // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	_arg2 = C.gint(dx)
	_arg3 = C.gint(dy)

	C.gtk_style_context_scroll_animations(_arg0, _arg1, _arg2, _arg3)
}

// SetBackground sets the background of @window to the background pattern or
// color specified in @context for its current state.
//
// Deprecated: since version 3.18.
func (c *StyleContextClass) SetBackground(window gdk.Window) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 *C.GdkWindow       // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))

	C.gtk_style_context_set_background(_arg0, _arg1)
}

// SetFrameClock attaches @context to the given frame clock.
//
// The frame clock is used for the timing of animations.
//
// If you are using a StyleContext returned from gtk_widget_get_style_context(),
// you do not need to call this yourself.
func (c *StyleContextClass) SetFrameClock(frameClock gdk.FrameClock) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 *C.GdkFrameClock   // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkFrameClock)(unsafe.Pointer(frameClock.Native()))

	C.gtk_style_context_set_frame_clock(_arg0, _arg1)
}

// SetParent sets the parent style context for @context. The parent style
// context is used to implement inheritance
// (http://www.w3.org/TR/css3-cascade/#inheritance) of properties.
//
// If you are using a StyleContext returned from gtk_widget_get_style_context(),
// the parent will be set for you.
func (c *StyleContextClass) SetParent(parent StyleContext) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 *C.GtkStyleContext // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkStyleContext)(unsafe.Pointer(parent.Native()))

	C.gtk_style_context_set_parent(_arg0, _arg1)
}

// SetPath sets the WidgetPath used for style matching. As a consequence, the
// style will be regenerated to match the new given path.
//
// If you are using a StyleContext returned from gtk_widget_get_style_context(),
// you do not need to call this yourself.
func (c *StyleContextClass) SetPath(path *WidgetPath) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 *C.GtkWidgetPath   // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GtkWidgetPath)(unsafe.Pointer(path))

	C.gtk_style_context_set_path(_arg0, _arg1)
}

// SetScale sets the scale to use when getting image assets for the style.
func (c *StyleContextClass) SetScale(scale int) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 C.gint             // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	_arg1 = C.gint(scale)

	C.gtk_style_context_set_scale(_arg0, _arg1)
}

// SetScreen attaches @context to the given screen.
//
// The screen is used to add style information from “global” style providers,
// such as the screen’s Settings instance.
//
// If you are using a StyleContext returned from gtk_widget_get_style_context(),
// you do not need to call this yourself.
func (c *StyleContextClass) SetScreen(screen gdk.Screen) {
	var _arg0 *C.GtkStyleContext // out
	var _arg1 *C.GdkScreen       // out

	_arg0 = (*C.GtkStyleContext)(unsafe.Pointer(c.Native()))
	_arg1 = (*C.GdkScreen)(unsafe.Pointer(screen.Native()))

	C.gtk_style_context_set_screen(_arg0, _arg1)
}
