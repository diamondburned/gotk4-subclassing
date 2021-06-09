// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gerror"
	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/internal/ptr"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	"github.com/diamondburned/gotk4/pkg/gobject/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_builder_get_type()), F: marshalBuilder},
	})
}

// Builder: a GtkBuilder is an auxiliary object that reads textual descriptions
// of a user interface and instantiates the described objects. To create a
// GtkBuilder from a user interface description, call
// gtk_builder_new_from_file(), gtk_builder_new_from_resource() or
// gtk_builder_new_from_string().
//
// In the (unusual) case that you want to add user interface descriptions from
// multiple sources to the same GtkBuilder you can call gtk_builder_new() to get
// an empty builder and populate it by (multiple) calls to
// gtk_builder_add_from_file(), gtk_builder_add_from_resource() or
// gtk_builder_add_from_string().
//
// A GtkBuilder holds a reference to all objects that it has constructed and
// drops these references when it is finalized. This finalization can cause the
// destruction of non-widget objects or widgets which are not contained in a
// toplevel window. For toplevel windows constructed by a builder, it is the
// responsibility of the user to call gtk_window_destroy() to get rid of them
// and all the widgets they contain.
//
// The functions gtk_builder_get_object() and gtk_builder_get_objects() can be
// used to access the widgets in the interface by the names assigned to them
// inside the UI description. Toplevel windows returned by these functions will
// stay around until the user explicitly destroys them with
// gtk_window_destroy(). Other widgets will either be part of a larger hierarchy
// constructed by the builder (in which case you should not have to worry about
// their lifecycle), or without a parent, in which case they have to be added to
// some container to make use of them. Non-widget objects need to be reffed with
// g_object_ref() to keep them beyond the lifespan of the builder.
//
//
// GtkBuilder UI Definitions
//
// GtkBuilder parses textual descriptions of user interfaces which are specified
// in XML format. We refer to these descriptions as “GtkBuilder UI definitions”
// or just “UI definitions” if the context is clear.
//
// The toplevel element is `<interface>`. It optionally takes a “domain”
// attribute, which will make the builder look for translated strings using
// `dgettext()` in the domain specified. This can also be done by calling
// gtk_builder_set_translation_domain() on the builder. Objects are described by
// `<object>` elements, which can contain <property> elements to set properties,
// `<signal>` elements which connect signals to handlers, and `<child>`
// elements, which describe child objects (most often widgets inside a
// container, but also e.g. actions in an action group, or columns in a tree
// model). A `<child>` element contains an `<object>` element which describes
// the child object. The target toolkit version(s) are described by <requires>
// elements, the “lib” attribute specifies the widget library in question
// (currently the only supported value is “gtk”) and the “version” attribute
// specifies the target version in the form “`<major>`.`<minor>`”. The builder
// will error out if the version requirements are not met.
//
// Typically, the specific kind of object represented by an `<object>` element
// is specified by the “class” attribute. If the type has not been loaded yet,
// GTK tries to find the `get_type()` function from the class name by applying
// heuristics. This works in most cases, but if necessary, it is possible to
// specify the name of the `get_type()` function explicitly with the "type-func"
// attribute.
//
// Objects may be given a name with the “id” attribute, which allows the
// application to retrieve them from the builder with gtk_builder_get_object().
// An id is also necessary to use the object as property value in other parts of
// the UI definition. GTK reserves ids starting and ending with `___` (three
// consecutive underscores) for its own purposes.
//
// Setting properties of objects is pretty straightforward with the <property>
// element: the “name” attribute specifies the name of the property, and the
// content of the element specifies the value. If the “translatable” attribute
// is set to a true value, GTK uses `gettext()` (or `dgettext()` if the builder
// has a translation domain set) to find a translation for the value. This
// happens before the value is parsed, so it can be used for properties of any
// type, but it is probably most useful for string properties. It is also
// possible to specify a context to disambiguate short strings, and comments
// which may help the translators.
//
// Builder can parse textual representations for the most common property types:
// characters, strings, integers, floating-point numbers, booleans (strings like
// “TRUE”, “t”, “yes”, “y”, “1” are interpreted as true, strings like “FALSE”,
// “f”, “no”, “n”, “0” are interpreted as false), enumerations (can be specified
// by their name, nick or integer value), flags (can be specified by their name,
// nick, integer value, optionally combined with “|”, e.g.
// “GTK_INPUT_HINT_EMOJI|GTK_INPUT_HINT_LOWERCASE”) and colors (in a format
// understood by gdk_rgba_parse()).
//
// GVariants can be specified in the format understood by g_variant_parse(), and
// pixbufs can be specified as a filename of an image file to load.
//
// Objects can be referred to by their name and by default refer to objects
// declared in the local XML fragment and objects exposed via
// gtk_builder_expose_object(). In general, GtkBuilder allows forward references
// to objects — declared in the local XML; an object doesn’t have to be
// constructed before it can be referred to. The exception to this rule is that
// an object has to be constructed before it can be used as the value of a
// construct-only property.
//
// It is also possible to bind a property value to another object's property
// value using the attributes "bind-source" to specify the source object of the
// binding, and optionally, "bind-property" and "bind-flags" to specify the
// source property and source binding flags respectively. Internally builder
// implements this using #GBinding objects. For more information see
// g_object_bind_property()
//
// Sometimes it is necessary to refer to widgets which have implicitly been
// constructed by GTK as part of a composite widget, to set properties on them
// or to add further children (e.g. the content area of a Dialog). This can be
// achieved by setting the “internal-child” property of the `<child>` element to
// a true value. Note that Builder still requires an `<object>` element for the
// internal child, even if it has already been constructed.
//
// A number of widgets have different places where a child can be added (e.g.
// tabs vs. page content in notebooks). This can be reflected in a UI definition
// by specifying the “type” attribute on a `<child>` The possible values for the
// “type” attribute are described in the sections describing the widget-specific
// portions of UI definitions.
//
//
// Signal handlers and function pointers
//
// Signal handlers are set up with the <signal> element. The “name” attribute
// specifies the name of the signal, and the “handler” attribute specifies the
// function to connect to the signal. The remaining attributes, “after”,
// “swapped” and “object”, have the same meaning as the corresponding parameters
// of the g_signal_connect_object() or g_signal_connect_data() functions. A
// “last_modification_time” attribute is also allowed, but it does not have a
// meaning to the builder.
//
// If you rely on #GModule support to lookup callbacks in the symbol table, the
// following details should be noted:
//
// When compiling applications for Windows, you must declare signal callbacks
// with MODULE_EXPORT, or they will not be put in the symbol table. On Linux and
// Unices, this is not necessary; applications should instead be compiled with
// the -Wl,--export-dynamic CFLAGS, and linked against gmodule-export-2.0.
//
// A GtkBuilder UI Definition
//
//    <interface>
//      <object class="GtkDialog" id="dialog1">
//        <child internal-child="vbox">
//          <object class="GtkBox" id="vbox1">
//            <child internal-child="action_area">
//              <object class="GtkBox" id="hbuttonbox1">
//                <child>
//                  <object class="GtkButton" id="ok_button">
//                    <property name="label">gtk-ok</property>
//                    <signal name="clicked" handler="ok_button_clicked"/>
//                  </object>
//                </child>
//              </object>
//            </child>
//          </object>
//        </child>
//      </object>
//    </interface>
//
// Beyond this general structure, several object classes define their own XML
// DTD fragments for filling in the ANY placeholders in the DTD above. Note that
// a custom element in a <child> element gets parsed by the custom tag handler
// of the parent object, while a custom element in an <object> element gets
// parsed by the custom tag handler of the object.
//
// These XML fragments are explained in the documentation of the respective
// objects.
//
// Additionally, since 3.10 a special <template> tag has been added to the
// format allowing one to define a widget class’s components. See the [GtkWidget
// documentation][composite-templates] for details.
type Builder interface {
	gextras.Objector

	// AddFromFile parses a file containing a [GtkBuilder UI
	// definition][BUILDER-UI] and merges it with the current contents of
	// @builder.
	//
	// This function is useful if you need to call
	// gtk_builder_set_current_object() to add user data to callbacks before
	// loading GtkBuilder UI. Otherwise, you probably want
	// gtk_builder_new_from_file() instead.
	//
	// If an error occurs, 0 will be returned and @error will be assigned a
	// #GError from the K_BUILDER_ERROR, MARKUP_ERROR or FILE_ERROR domain.
	//
	// It’s not really reasonable to attempt to handle failures of this call.
	// You should not use this function with untrusted files (ie: files that are
	// not part of your application). Broken Builder files can easily crash your
	// program, and it’s possible that memory was leaked leading up to the
	// reported failure. The only reasonable thing to do when an error is
	// detected is to call g_error().
	AddFromFile(filename string) error
	// AddFromResource parses a resource file containing a [GtkBuilder UI
	// definition][BUILDER-UI] and merges it with the current contents of
	// @builder.
	//
	// This function is useful if you need to call
	// gtk_builder_set_current_object() to add user data to callbacks before
	// loading GtkBuilder UI. Otherwise, you probably want
	// gtk_builder_new_from_resource() instead.
	//
	// If an error occurs, 0 will be returned and @error will be assigned a
	// #GError from the K_BUILDER_ERROR, MARKUP_ERROR or RESOURCE_ERROR domain.
	//
	// It’s not really reasonable to attempt to handle failures of this call.
	// The only reasonable thing to do when an error is detected is to call
	// g_error().
	AddFromResource(resourcePath string) error
	// AddFromString parses a string containing a [GtkBuilder UI
	// definition][BUILDER-UI] and merges it with the current contents of
	// @builder.
	//
	// This function is useful if you need to call
	// gtk_builder_set_current_object() to add user data to callbacks before
	// loading GtkBuilder UI. Otherwise, you probably want
	// gtk_builder_new_from_string() instead.
	//
	// Upon errors false will be returned and @error will be assigned a #GError
	// from the K_BUILDER_ERROR, MARKUP_ERROR or VARIANT_PARSE_ERROR domain.
	//
	// It’s not really reasonable to attempt to handle failures of this call.
	// The only reasonable thing to do when an error is detected is to call
	// g_error().
	AddFromString(buffer string, length int) error
	// AddObjectsFromFile parses a file containing a [GtkBuilder UI
	// definition][BUILDER-UI] building only the requested objects and merges
	// them with the current contents of @builder.
	//
	// Upon errors 0 will be returned and @error will be assigned a #GError from
	// the K_BUILDER_ERROR, MARKUP_ERROR or FILE_ERROR domain.
	//
	// If you are adding an object that depends on an object that is not its
	// child (for instance a TreeView that depends on its TreeModel), you have
	// to explicitly list all of them in @object_ids.
	AddObjectsFromFile(filename string, objectIds []string) error
	// AddObjectsFromResource parses a resource file containing a [GtkBuilder UI
	// definition][BUILDER-UI] building only the requested objects and merges
	// them with the current contents of @builder.
	//
	// Upon errors 0 will be returned and @error will be assigned a #GError from
	// the K_BUILDER_ERROR, MARKUP_ERROR or RESOURCE_ERROR domain.
	//
	// If you are adding an object that depends on an object that is not its
	// child (for instance a TreeView that depends on its TreeModel), you have
	// to explicitly list all of them in @object_ids.
	AddObjectsFromResource(resourcePath string, objectIds []string) error
	// AddObjectsFromString parses a string containing a [GtkBuilder UI
	// definition][BUILDER-UI] building only the requested objects and merges
	// them with the current contents of @builder.
	//
	// Upon errors false will be returned and @error will be assigned a #GError
	// from the K_BUILDER_ERROR or MARKUP_ERROR domain.
	//
	// If you are adding an object that depends on an object that is not its
	// child (for instance a TreeView that depends on its TreeModel), you have
	// to explicitly list all of them in @object_ids.
	AddObjectsFromString(buffer string, length int, objectIds []string) error
	// ExposeObject: add @object to the @builder object pool so it can be
	// referenced just like any other object built by builder.
	ExposeObject(name string, object gextras.Objector)
	// ExtendWithTemplate: main private entry point for building composite
	// container components from template XML.
	//
	// This is exported purely to let gtk-builder-tool validate templates,
	// applications have no need to call this function.
	ExtendWithTemplate(object gextras.Objector, templateType externglib.Type, buffer string, length int) error
	// CurrentObject gets the current object set via
	// gtk_builder_set_current_object().
	CurrentObject() gextras.Objector
	// Object gets the object named @name. Note that this function does not
	// increment the reference count of the returned object.
	Object(name string) gextras.Objector
	// Objects gets all objects that have been constructed by @builder. Note
	// that this function does not increment the reference counts of the
	// returned objects.
	Objects() *glib.SList
	// Scope gets the scope in use that was set via gtk_builder_set_scope().
	//
	// See the BuilderScope documentation for details.
	Scope() BuilderScope
	// TranslationDomain gets the translation domain of @builder.
	TranslationDomain() string
	// TypeFromName looks up a type by name, using the virtual function that
	// Builder has for that purpose. This is mainly used when implementing the
	// Buildable interface on a type.
	TypeFromName(typeName string) externglib.Type
	// SetCurrentObject sets the current object for the @builder. The current
	// object can be thought of as the `this` object that the builder is working
	// for and will often be used as the default object when an object is
	// optional.
	//
	// gtk_widget_init_template() for example will set the current object to the
	// widget the template is inited for. For functions like
	// gtk_builder_new_from_resource(), the current object will be nil.
	SetCurrentObject(currentObject gextras.Objector)
	// SetScope sets the scope the builder should operate in.
	//
	// If @scope is nil a new BuilderCScope will be created.
	//
	// See the BuilderScope documentation for details.
	SetScope(scope BuilderScope)
	// SetTranslationDomain sets the translation domain of @builder. See
	// Builder:translation-domain.
	SetTranslationDomain(domain string)
	// ValueFromString: this function demarshals a value from a string. This
	// function calls g_value_init() on the @value argument, so it need not be
	// initialised beforehand.
	//
	// This function can handle char, uchar, boolean, int, uint, long, ulong,
	// enum, flags, float, double, string, RGBA and Adjustment type values.
	// Support for Widget type values is still to come.
	//
	// Upon errors false will be returned and @error will be assigned a #GError
	// from the K_BUILDER_ERROR domain.
	ValueFromString(pspec gobject.ParamSpec, string string) (value *externglib.Value, goerr error)
	// ValueFromStringType: like gtk_builder_value_from_string(), this function
	// demarshals a value from a string, but takes a #GType instead of Spec.
	// This function calls g_value_init() on the @value argument, so it need not
	// be initialised beforehand.
	//
	// Upon errors false will be returned and @error will be assigned a #GError
	// from the K_BUILDER_ERROR domain.
	ValueFromStringType(typ externglib.Type, string string) (value *externglib.Value, goerr error)
}

// builder implements the Builder interface.
type builder struct {
	gextras.Objector
}

var _ Builder = (*builder)(nil)

// WrapBuilder wraps a GObject to the right type. It is
// primarily used internally.
func WrapBuilder(obj *externglib.Object) Builder {
	return Builder{
		Objector: obj,
	}
}

func marshalBuilder(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapBuilder(obj), nil
}

// NewBuilder constructs a class Builder.
func NewBuilder() Builder {
	var cret C.GtkBuilder

	cret = C.gtk_builder_new()

	var builder Builder

	builder = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Builder)

	return builder
}

// NewBuilderFromFile constructs a class Builder.
func NewBuilderFromFile(filename string) Builder {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkBuilder

	cret = C.gtk_builder_new_from_file(arg1)

	var builder Builder

	builder = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Builder)

	return builder
}

// NewBuilderFromResource constructs a class Builder.
func NewBuilderFromResource(resourcePath string) Builder {
	var arg1 *C.char

	arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GtkBuilder

	cret = C.gtk_builder_new_from_resource(arg1)

	var builder Builder

	builder = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Builder)

	return builder
}

// NewBuilderFromString constructs a class Builder.
func NewBuilderFromString(string string, length int) Builder {
	var arg1 *C.char
	var arg2 C.gssize

	arg1 = (*C.char)(C.CString(string))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gssize(length)

	var cret C.GtkBuilder

	cret = C.gtk_builder_new_from_string(arg1, arg2)

	var builder Builder

	builder = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(cret.Native()))).(Builder)

	return builder
}

// AddFromFile parses a file containing a [GtkBuilder UI
// definition][BUILDER-UI] and merges it with the current contents of
// @builder.
//
// This function is useful if you need to call
// gtk_builder_set_current_object() to add user data to callbacks before
// loading GtkBuilder UI. Otherwise, you probably want
// gtk_builder_new_from_file() instead.
//
// If an error occurs, 0 will be returned and @error will be assigned a
// #GError from the K_BUILDER_ERROR, MARKUP_ERROR or FILE_ERROR domain.
//
// It’s not really reasonable to attempt to handle failures of this call.
// You should not use this function with untrusted files (ie: files that are
// not part of your application). Broken Builder files can easily crash your
// program, and it’s possible that memory was leaked leading up to the
// reported failure. The only reasonable thing to do when an error is
// detected is to call g_error().
func (b builder) AddFromFile(filename string) error {
	var arg0 *C.GtkBuilder
	var arg1 *C.char

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))

	var cerr *C.GError

	C.gtk_builder_add_from_file(arg0, arg1, cerr)

	var goerr error

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// AddFromResource parses a resource file containing a [GtkBuilder UI
// definition][BUILDER-UI] and merges it with the current contents of
// @builder.
//
// This function is useful if you need to call
// gtk_builder_set_current_object() to add user data to callbacks before
// loading GtkBuilder UI. Otherwise, you probably want
// gtk_builder_new_from_resource() instead.
//
// If an error occurs, 0 will be returned and @error will be assigned a
// #GError from the K_BUILDER_ERROR, MARKUP_ERROR or RESOURCE_ERROR domain.
//
// It’s not really reasonable to attempt to handle failures of this call.
// The only reasonable thing to do when an error is detected is to call
// g_error().
func (b builder) AddFromResource(resourcePath string) error {
	var arg0 *C.GtkBuilder
	var arg1 *C.char

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(arg1))

	var cerr *C.GError

	C.gtk_builder_add_from_resource(arg0, arg1, cerr)

	var goerr error

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// AddFromString parses a string containing a [GtkBuilder UI
// definition][BUILDER-UI] and merges it with the current contents of
// @builder.
//
// This function is useful if you need to call
// gtk_builder_set_current_object() to add user data to callbacks before
// loading GtkBuilder UI. Otherwise, you probably want
// gtk_builder_new_from_string() instead.
//
// Upon errors false will be returned and @error will be assigned a #GError
// from the K_BUILDER_ERROR, MARKUP_ERROR or VARIANT_PARSE_ERROR domain.
//
// It’s not really reasonable to attempt to handle failures of this call.
// The only reasonable thing to do when an error is detected is to call
// g_error().
func (b builder) AddFromString(buffer string, length int) error {
	var arg0 *C.GtkBuilder
	var arg1 *C.char
	var arg2 C.gssize

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(buffer))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gssize(length)

	var cerr *C.GError

	C.gtk_builder_add_from_string(arg0, arg1, arg2, cerr)

	var goerr error

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// AddObjectsFromFile parses a file containing a [GtkBuilder UI
// definition][BUILDER-UI] building only the requested objects and merges
// them with the current contents of @builder.
//
// Upon errors 0 will be returned and @error will be assigned a #GError from
// the K_BUILDER_ERROR, MARKUP_ERROR or FILE_ERROR domain.
//
// If you are adding an object that depends on an object that is not its
// child (for instance a TreeView that depends on its TreeModel), you have
// to explicitly list all of them in @object_ids.
func (b builder) AddObjectsFromFile(filename string, objectIds []string) error {
	var arg0 *C.GtkBuilder
	var arg1 *C.char
	var arg2 **C.char

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(filename))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (**C.char)(C.malloc((len(objectIds) + 1) * unsafe.Sizeof(int(0))))
	defer C.free(unsafe.Pointer(arg2))

	{
		var out []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(arg2), int(len(objectIds)))

		for i := range objectIds {
			arg2 = (*C.gchar)(C.CString(objectIds))
			defer C.free(unsafe.Pointer(arg2))
		}
	}

	var cerr *C.GError

	C.gtk_builder_add_objects_from_file(arg0, arg1, arg2, cerr)

	var goerr error

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// AddObjectsFromResource parses a resource file containing a [GtkBuilder UI
// definition][BUILDER-UI] building only the requested objects and merges
// them with the current contents of @builder.
//
// Upon errors 0 will be returned and @error will be assigned a #GError from
// the K_BUILDER_ERROR, MARKUP_ERROR or RESOURCE_ERROR domain.
//
// If you are adding an object that depends on an object that is not its
// child (for instance a TreeView that depends on its TreeModel), you have
// to explicitly list all of them in @object_ids.
func (b builder) AddObjectsFromResource(resourcePath string, objectIds []string) error {
	var arg0 *C.GtkBuilder
	var arg1 *C.char
	var arg2 **C.char

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(resourcePath))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (**C.char)(C.malloc((len(objectIds) + 1) * unsafe.Sizeof(int(0))))
	defer C.free(unsafe.Pointer(arg2))

	{
		var out []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(arg2), int(len(objectIds)))

		for i := range objectIds {
			arg2 = (*C.gchar)(C.CString(objectIds))
			defer C.free(unsafe.Pointer(arg2))
		}
	}

	var cerr *C.GError

	C.gtk_builder_add_objects_from_resource(arg0, arg1, arg2, cerr)

	var goerr error

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// AddObjectsFromString parses a string containing a [GtkBuilder UI
// definition][BUILDER-UI] building only the requested objects and merges
// them with the current contents of @builder.
//
// Upon errors false will be returned and @error will be assigned a #GError
// from the K_BUILDER_ERROR or MARKUP_ERROR domain.
//
// If you are adding an object that depends on an object that is not its
// child (for instance a TreeView that depends on its TreeModel), you have
// to explicitly list all of them in @object_ids.
func (b builder) AddObjectsFromString(buffer string, length int, objectIds []string) error {
	var arg0 *C.GtkBuilder
	var arg1 *C.char
	var arg2 C.gssize
	var arg3 **C.char

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(buffer))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = C.gssize(length)
	arg3 = (**C.char)(C.malloc((len(objectIds) + 1) * unsafe.Sizeof(int(0))))
	defer C.free(unsafe.Pointer(arg3))

	{
		var out []*C.gchar
		ptr.SetSlice(unsafe.Pointer(&dst), unsafe.Pointer(arg3), int(len(objectIds)))

		for i := range objectIds {
			arg3 = (*C.gchar)(C.CString(objectIds))
			defer C.free(unsafe.Pointer(arg3))
		}
	}

	var cerr *C.GError

	C.gtk_builder_add_objects_from_string(arg0, arg1, arg2, arg3, cerr)

	var goerr error

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// ExposeObject: add @object to the @builder object pool so it can be
// referenced just like any other object built by builder.
func (b builder) ExposeObject(name string, object gextras.Objector) {
	var arg0 *C.GtkBuilder
	var arg1 *C.char
	var arg2 *C.GObject

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.GObject)(unsafe.Pointer(object.Native()))

	C.gtk_builder_expose_object(arg0, arg1, arg2)
}

// ExtendWithTemplate: main private entry point for building composite
// container components from template XML.
//
// This is exported purely to let gtk-builder-tool validate templates,
// applications have no need to call this function.
func (b builder) ExtendWithTemplate(object gextras.Objector, templateType externglib.Type, buffer string, length int) error {
	var arg0 *C.GtkBuilder
	var arg1 *C.GObject
	var arg2 C.GType
	var arg3 *C.char
	var arg4 C.gssize

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GObject)(unsafe.Pointer(object.Native()))
	arg2 = C.GType(templateType)
	arg3 = (*C.char)(C.CString(buffer))
	defer C.free(unsafe.Pointer(arg3))
	arg4 = C.gssize(length)

	var cerr *C.GError

	C.gtk_builder_extend_with_template(arg0, arg1, arg2, arg3, arg4, cerr)

	var goerr error

	goerr = gerror.Take(unsafe.Pointer(cerr))

	return goerr
}

// CurrentObject gets the current object set via
// gtk_builder_set_current_object().
func (b builder) CurrentObject() gextras.Objector {
	var arg0 *C.GtkBuilder

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))

	var cret *C.GObject

	cret = C.gtk_builder_get_current_object(arg0)

	var object gextras.Objector

	object = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gextras.Objector)

	return object
}

// Object gets the object named @name. Note that this function does not
// increment the reference count of the returned object.
func (b builder) Object(name string) gextras.Objector {
	var arg0 *C.GtkBuilder
	var arg1 *C.char

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(arg1))

	var cret *C.GObject

	cret = C.gtk_builder_get_object(arg0, arg1)

	var object gextras.Objector

	object = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(gextras.Objector)

	return object
}

// Objects gets all objects that have been constructed by @builder. Note
// that this function does not increment the reference counts of the
// returned objects.
func (b builder) Objects() *glib.SList {
	var arg0 *C.GtkBuilder

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))

	var cret *C.GSList

	cret = C.gtk_builder_get_objects(arg0)

	var sList *glib.SList

	sList = glib.WrapSList(unsafe.Pointer(cret))
	runtime.SetFinalizer(sList, func(v *glib.SList) {
		C.free(unsafe.Pointer(v.Native()))
	})

	return sList
}

// Scope gets the scope in use that was set via gtk_builder_set_scope().
//
// See the BuilderScope documentation for details.
func (b builder) Scope() BuilderScope {
	var arg0 *C.GtkBuilder

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))

	var cret *C.GtkBuilderScope

	cret = C.gtk_builder_get_scope(arg0)

	var builderScope BuilderScope

	builderScope = gextras.CastObject(externglib.Take(unsafe.Pointer(cret.Native()))).(BuilderScope)

	return builderScope
}

// TranslationDomain gets the translation domain of @builder.
func (b builder) TranslationDomain() string {
	var arg0 *C.GtkBuilder

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))

	var cret *C.char

	cret = C.gtk_builder_get_translation_domain(arg0)

	var utf8 string

	utf8 = C.GoString(cret)

	return utf8
}

// TypeFromName looks up a type by name, using the virtual function that
// Builder has for that purpose. This is mainly used when implementing the
// Buildable interface on a type.
func (b builder) TypeFromName(typeName string) externglib.Type {
	var arg0 *C.GtkBuilder
	var arg1 *C.char

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(typeName))
	defer C.free(unsafe.Pointer(arg1))

	var cret C.GType

	cret = C.gtk_builder_get_type_from_name(arg0, arg1)

	var gType externglib.Type

	gType = externglib.Type(cret)

	return gType
}

// SetCurrentObject sets the current object for the @builder. The current
// object can be thought of as the `this` object that the builder is working
// for and will often be used as the default object when an object is
// optional.
//
// gtk_widget_init_template() for example will set the current object to the
// widget the template is inited for. For functions like
// gtk_builder_new_from_resource(), the current object will be nil.
func (b builder) SetCurrentObject(currentObject gextras.Objector) {
	var arg0 *C.GtkBuilder
	var arg1 *C.GObject

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GObject)(unsafe.Pointer(currentObject.Native()))

	C.gtk_builder_set_current_object(arg0, arg1)
}

// SetScope sets the scope the builder should operate in.
//
// If @scope is nil a new BuilderCScope will be created.
//
// See the BuilderScope documentation for details.
func (b builder) SetScope(scope BuilderScope) {
	var arg0 *C.GtkBuilder
	var arg1 *C.GtkBuilderScope

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GtkBuilderScope)(unsafe.Pointer(scope.Native()))

	C.gtk_builder_set_scope(arg0, arg1)
}

// SetTranslationDomain sets the translation domain of @builder. See
// Builder:translation-domain.
func (b builder) SetTranslationDomain(domain string) {
	var arg0 *C.GtkBuilder
	var arg1 *C.char

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.char)(C.CString(domain))
	defer C.free(unsafe.Pointer(arg1))

	C.gtk_builder_set_translation_domain(arg0, arg1)
}

// ValueFromString: this function demarshals a value from a string. This
// function calls g_value_init() on the @value argument, so it need not be
// initialised beforehand.
//
// This function can handle char, uchar, boolean, int, uint, long, ulong,
// enum, flags, float, double, string, RGBA and Adjustment type values.
// Support for Widget type values is still to come.
//
// Upon errors false will be returned and @error will be assigned a #GError
// from the K_BUILDER_ERROR domain.
func (b builder) ValueFromString(pspec gobject.ParamSpec, string string) (value *externglib.Value, goerr error) {
	var arg0 *C.GtkBuilder
	var arg1 *C.GParamSpec
	var arg2 *C.char

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = (*C.GParamSpec)(unsafe.Pointer(pspec.Native()))
	arg2 = (*C.char)(C.CString(string))
	defer C.free(unsafe.Pointer(arg2))

	var arg3 C.GValue
	var cerr *C.GError

	C.gtk_builder_value_from_string(arg0, arg1, arg2, &arg3, cerr)

	var value *externglib.Value
	var goerr error

	value = externglib.ValueFromNative(unsafe.Pointer(arg3))
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return value, goerr
}

// ValueFromStringType: like gtk_builder_value_from_string(), this function
// demarshals a value from a string, but takes a #GType instead of Spec.
// This function calls g_value_init() on the @value argument, so it need not
// be initialised beforehand.
//
// Upon errors false will be returned and @error will be assigned a #GError
// from the K_BUILDER_ERROR domain.
func (b builder) ValueFromStringType(typ externglib.Type, string string) (value *externglib.Value, goerr error) {
	var arg0 *C.GtkBuilder
	var arg1 C.GType
	var arg2 *C.char

	arg0 = (*C.GtkBuilder)(unsafe.Pointer(b.Native()))
	arg1 = C.GType(typ)
	arg2 = (*C.char)(C.CString(string))
	defer C.free(unsafe.Pointer(arg2))

	var arg3 C.GValue
	var cerr *C.GError

	C.gtk_builder_value_from_string_type(arg0, arg1, arg2, &arg3, cerr)

	var value *externglib.Value
	var goerr error

	value = externglib.ValueFromNative(unsafe.Pointer(arg3))
	goerr = gerror.Take(unsafe.Pointer(cerr))

	return value, goerr
}
