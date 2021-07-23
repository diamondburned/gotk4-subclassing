// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
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
		{T: externglib.Type(C.gtk_builder_error_get_type()), F: marshalBuilderError},
		{T: externglib.Type(C.gtk_builder_get_type()), F: marshalBuilderer},
	})
}

// BuilderError: error codes that identify various errors that can occur while
// using Builder.
type BuilderError int

const (
	// BuilderErrorInvalidTypeFunction: type-func attribute didn’t name a
	// function that returns a #GType.
	BuilderErrorInvalidTypeFunction BuilderError = iota
	// BuilderErrorUnhandledTag: input contained a tag that Builder can’t
	// handle.
	BuilderErrorUnhandledTag
	// BuilderErrorMissingAttribute: attribute that is required by Builder was
	// missing.
	BuilderErrorMissingAttribute
	// BuilderErrorInvalidAttribute found an attribute that it doesn’t
	// understand.
	BuilderErrorInvalidAttribute
	// BuilderErrorInvalidTag found a tag that it doesn’t understand.
	BuilderErrorInvalidTag
	// BuilderErrorMissingPropertyValue: required property value was missing.
	BuilderErrorMissingPropertyValue
	// BuilderErrorInvalidValue couldn’t parse some attribute value.
	BuilderErrorInvalidValue
	// BuilderErrorVersionMismatch: input file requires a newer version of GTK.
	BuilderErrorVersionMismatch
	// BuilderErrorDuplicateID: object id occurred twice.
	BuilderErrorDuplicateID
	// BuilderErrorObjectTypeRefused: specified object type is of the same type
	// or derived from the type of the composite class being extended with
	// builder XML.
	BuilderErrorObjectTypeRefused
	// BuilderErrorTemplateMismatch: wrong type was specified in a composite
	// class’s template XML
	BuilderErrorTemplateMismatch
	// BuilderErrorInvalidProperty: specified property is unknown for the object
	// class.
	BuilderErrorInvalidProperty
	// BuilderErrorInvalidSignal: specified signal is unknown for the object
	// class.
	BuilderErrorInvalidSignal
	// BuilderErrorInvalidID: object id is unknown.
	BuilderErrorInvalidID
	// BuilderErrorInvalidFunction: function could not be found. This often
	// happens when symbols are set to be kept private. Compiling code with
	// -rdynamic or using the gmodule-export-2.0 pkgconfig module can fix this
	// problem.
	BuilderErrorInvalidFunction
)

func marshalBuilderError(p uintptr) (interface{}, error) {
	return BuilderError(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for BuilderError.
func (b BuilderError) String() string {
	switch b {
	case BuilderErrorInvalidTypeFunction:
		return "InvalidTypeFunction"
	case BuilderErrorUnhandledTag:
		return "UnhandledTag"
	case BuilderErrorMissingAttribute:
		return "MissingAttribute"
	case BuilderErrorInvalidAttribute:
		return "InvalidAttribute"
	case BuilderErrorInvalidTag:
		return "InvalidTag"
	case BuilderErrorMissingPropertyValue:
		return "MissingPropertyValue"
	case BuilderErrorInvalidValue:
		return "InvalidValue"
	case BuilderErrorVersionMismatch:
		return "VersionMismatch"
	case BuilderErrorDuplicateID:
		return "DuplicateID"
	case BuilderErrorObjectTypeRefused:
		return "ObjectTypeRefused"
	case BuilderErrorTemplateMismatch:
		return "TemplateMismatch"
	case BuilderErrorInvalidProperty:
		return "InvalidProperty"
	case BuilderErrorInvalidSignal:
		return "InvalidSignal"
	case BuilderErrorInvalidID:
		return "InvalidID"
	case BuilderErrorInvalidFunction:
		return "InvalidFunction"
	default:
		return fmt.Sprintf("BuilderError(%d)", b)
	}
}

// Builder: GtkBuilder reads XML descriptions of a user interface and
// instantiates the described objects.
//
// To create a GtkBuilder from a user interface description, call
// gtk.Builder.NewFromFile, gtk.Builder.NewFromResource or
// gtk.Builder.NewFromString.
//
// In the (unusual) case that you want to add user interface descriptions from
// multiple sources to the same GtkBuilder you can call gtk.Builder.New to get
// an empty builder and populate it by (multiple) calls to
// gtk.Builder.AddFromFile(), gtk.Builder.AddFromResource() or
// gtk.Builder.AddFromString().
//
// A GtkBuilder holds a reference to all objects that it has constructed and
// drops these references when it is finalized. This finalization can cause the
// destruction of non-widget objects or widgets which are not contained in a
// toplevel window. For toplevel windows constructed by a builder, it is the
// responsibility of the user to call gtk.Window.Destroy() to get rid of them
// and all the widgets they contain.
//
// The functions gtk.Builder.GetObject() and gtk.Builder.GetObjects() can be
// used to access the widgets in the interface by the names assigned to them
// inside the UI description. Toplevel windows returned by these functions will
// stay around until the user explicitly destroys them with
// gtk.Window.Destroy(). Other widgets will either be part of a larger hierarchy
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
// The toplevel element is <interface>. It optionally takes a “domain”
// attribute, which will make the builder look for translated strings using
// dgettext() in the domain specified. This can also be done by calling
// gtk.Builder.SetTranslationDomain() on the builder.
//
// Objects are described by <object> elements, which can contain <property>
// elements to set properties, <signal> elements which connect signals to
// handlers, and <child> elements, which describe child objects (most often
// widgets inside a container, but also e.g. actions in an action group, or
// columns in a tree model). A <child> element contains an <object> element
// which describes the child object.
//
// The target toolkit version(s) are described by <requires> elements, the “lib”
// attribute specifies the widget library in question (currently the only
// supported value is “gtk”) and the “version” attribute specifies the target
// version in the form “<major>.<minor>”. GtkBuilder will error out if the
// version requirements are not met.
//
// Typically, the specific kind of object represented by an <object> element is
// specified by the “class” attribute. If the type has not been loaded yet, GTK
// tries to find the get_type() function from the class name by applying
// heuristics. This works in most cases, but if necessary, it is possible to
// specify the name of the get_type() function explicitly with the "type-func"
// attribute.
//
// Objects may be given a name with the “id” attribute, which allows the
// application to retrieve them from the builder with gtk.Builder.GetObject().
// An id is also necessary to use the object as property value in other parts of
// the UI definition. GTK reserves ids starting and ending with ___ (three
// consecutive underscores) for its own purposes.
//
// Setting properties of objects is pretty straightforward with the <property>
// element: the “name” attribute specifies the name of the property, and the
// content of the element specifies the value. If the “translatable” attribute
// is set to a true value, GTK uses gettext() (or dgettext() if the builder has
// a translation domain set) to find a translation for the value. This happens
// before the value is parsed, so it can be used for properties of any type, but
// it is probably most useful for string properties. It is also possible to
// specify a context to disambiguate short strings, and comments which may help
// the translators.
//
// GtkBuilder can parse textual representations for the most common property
// types: characters, strings, integers, floating-point numbers, booleans
// (strings like “TRUE”, “t”, “yes”, “y”, “1” are interpreted as TRUE, strings
// like “FALSE”, “f”, “no”, “n”, “0” are interpreted as FALSE), enumerations
// (can be specified by their name, nick or integer value), flags (can be
// specified by their name, nick, integer value, optionally combined with “|”,
// e.g. “GTK_INPUT_HINT_EMOJI|GTK_INPUT_HINT_LOWERCASE”) and colors (in a format
// understood by gdk.RGBA.Parse()).
//
// GVariants can be specified in the format understood by g_variant_parse(), and
// pixbufs can be specified as a filename of an image file to load.
//
// Objects can be referred to by their name and by default refer to objects
// declared in the local XML fragment and objects exposed via
// gtk.Builder.ExposeObject(). In general, GtkBuilder allows forward references
// to objects — declared in the local XML; an object doesn’t have to be
// constructed before it can be referred to. The exception to this rule is that
// an object has to be constructed before it can be used as the value of a
// construct-only property.
//
// It is also possible to bind a property value to another object's property
// value using the attributes "bind-source" to specify the source object of the
// binding, and optionally, "bind-property" and "bind-flags" to specify the
// source property and source binding flags respectively. Internally, GtkBuilder
// implements this using GBinding objects. For more information see
// g_object_bind_property().
//
// Sometimes it is necessary to refer to widgets which have implicitly been
// constructed by GTK as part of a composite widget, to set properties on them
// or to add further children (e.g. the content area of a GtkDialog). This can
// be achieved by setting the “internal-child” property of the <child> element
// to a true value. Note that Builder still requires an <object> element for the
// internal child, even if it has already been constructed.
//
// A number of widgets have different places where a child can be added (e.g.
// tabs vs. page content in notebooks). This can be reflected in a UI definition
// by specifying the “type” attribute on a <child> The possible values for the
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
// If you rely on GModule support to lookup callbacks in the symbol table, the
// following details should be noted:
//
// When compiling applications for Windows, you must declare signal callbacks
// with G_MODULE_EXPORT, or they will not be put in the symbol table. On Linux
// and Unix, this is not necessary; applications should instead be compiled with
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
// A <template> tag can be used to define a widget class’s components. See the
// GtkWidget documentation
// (class.Widget.html#building-composite-widgets-from-template-xml) for details.
type Builder struct {
	*externglib.Object
}

func wrapBuilder(obj *externglib.Object) *Builder {
	return &Builder{
		Object: obj,
	}
}

func marshalBuilderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapBuilder(obj), nil
}

// NewBuilder creates a new empty builder object.
//
// This function is only useful if you intend to make multiple calls to
// gtk.Builder.AddFromFile(), gtk.Builder.AddFromResource() or
// gtk.Builder.AddFromString() in order to merge multiple UI descriptions into a
// single builder.
func NewBuilder() *Builder {
	var _cret *C.GtkBuilder // in

	_cret = C.gtk_builder_new()

	var _builder *Builder // out

	_builder = wrapBuilder(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _builder
}

// NewBuilderFromFile parses the UI definition in the file filename.
//
// If there is an error opening the file or parsing the description then the
// program will be aborted. You should only ever attempt to parse user interface
// descriptions that are shipped as part of your program.
func NewBuilderFromFile(filename string) *Builder {
	var _arg1 *C.char       // out
	var _cret *C.GtkBuilder // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_builder_new_from_file(_arg1)

	var _builder *Builder // out

	_builder = wrapBuilder(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _builder
}

// NewBuilderFromResource parses the UI definition at resource_path.
//
// If there is an error locating the resource or parsing the description, then
// the program will be aborted.
func NewBuilderFromResource(resourcePath string) *Builder {
	var _arg1 *C.char       // out
	var _cret *C.GtkBuilder // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_builder_new_from_resource(_arg1)

	var _builder *Builder // out

	_builder = wrapBuilder(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _builder
}

// NewBuilderFromString parses the UI definition in string.
//
// If string is NULL-terminated, then length should be -1. If length is not -1,
// then it is the length of string.
//
// If there is an error parsing string then the program will be aborted. You
// should not attempt to parse user interface description from untrusted
// sources.
func NewBuilderFromString(_string string, length int) *Builder {
	var _arg1 *C.char       // out
	var _arg2 C.gssize      // out
	var _cret *C.GtkBuilder // in

	_arg1 = (*C.char)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(length)

	_cret = C.gtk_builder_new_from_string(_arg1, _arg2)

	var _builder *Builder // out

	_builder = wrapBuilder(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _builder
}

// AddFromFile parses a file containing a UI definition and merges it with the
// current contents of builder.
//
// This function is useful if you need to call gtk.Builder.SetCurrentObject())
// to add user data to callbacks before loading GtkBuilder UI. Otherwise, you
// probably want gtk.Builder.NewFromFile instead.
//
// If an error occurs, 0 will be returned and error will be assigned a GError
// from the GTK_BUILDER_ERROR, G_MARKUP_ERROR or G_FILE_ERROR domains.
//
// It’s not really reasonable to attempt to handle failures of this call. You
// should not use this function with untrusted files (ie: files that are not
// part of your application). Broken GtkBuilder files can easily crash your
// program, and it’s possible that memory was leaked leading up to the reported
// failure. The only reasonable thing to do when an error is detected is to call
// g_error().
func (builder *Builder) AddFromFile(filename string) error {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_builder_add_from_file(_arg0, _arg1, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// AddFromResource parses a resource file containing a UI definition and merges
// it with the current contents of builder.
//
// This function is useful if you need to call gtk.Builder.SetCurrentObject() to
// add user data to callbacks before loading GtkBuilder UI. Otherwise, you
// probably want gtk.Builder.NewFromResource instead.
//
// If an error occurs, 0 will be returned and error will be assigned a GError
// from the GTK_BUILDER_ERROR, G_MARKUP_ERROR or G_RESOURCE_ERROR domain.
//
// It’s not really reasonable to attempt to handle failures of this call. The
// only reasonable thing to do when an error is detected is to call g_error().
func (builder *Builder) AddFromResource(resourcePath string) error {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_builder_add_from_resource(_arg0, _arg1, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// AddFromString parses a string containing a UI definition and merges it with
// the current contents of builder.
//
// This function is useful if you need to call gtk.Builder.SetCurrentObject() to
// add user data to callbacks before loading GtkBuilder UI. Otherwise, you
// probably want gtk.Builder.NewFromString instead.
//
// Upon errors FALSE will be returned and error will be assigned a GError from
// the GTK_BUILDER_ERROR, G_MARKUP_ERROR or G_VARIANT_PARSE_ERROR domain.
//
// It’s not really reasonable to attempt to handle failures of this call. The
// only reasonable thing to do when an error is detected is to call g_error().
func (builder *Builder) AddFromString(buffer string, length int) error {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out
	var _arg2 C.gssize      // out
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(buffer)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(length)

	C.gtk_builder_add_from_string(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// AddObjectsFromFile parses a file containing a UI definition building only the
// requested objects and merges them with the current contents of builder.
//
// Upon errors, 0 will be returned and error will be assigned a GError from the
// GTK_BUILDER_ERROR, G_MARKUP_ERROR or G_FILE_ERROR domain.
//
// If you are adding an object that depends on an object that is not its child
// (for instance a GtkTreeView that depends on its GtkTreeModel), you have to
// explicitly list all of them in object_ids.
func (builder *Builder) AddObjectsFromFile(filename string, objectIds []string) error {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out
	var _arg2 **C.char      // out
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		_arg2 = (**C.char)(C.malloc(C.ulong(len(objectIds)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg2))
		{
			out := unsafe.Slice(_arg2, len(objectIds)+1)
			var zero *C.char
			out[len(objectIds)] = zero
			for i := range objectIds {
				out[i] = (*C.char)(unsafe.Pointer(C.CString(objectIds[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.gtk_builder_add_objects_from_file(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// AddObjectsFromResource parses a resource file containing a UI definition,
// building only the requested objects and merges them with the current contents
// of builder.
//
// Upon errors, 0 will be returned and error will be assigned a GError from the
// GTK_BUILDER_ERROR, G_MARKUP_ERROR or G_RESOURCE_ERROR domain.
//
// If you are adding an object that depends on an object that is not its child
// (for instance a GtkTreeView that depends on its GtkTreeModel), you have to
// explicitly list all of them in object_ids.
func (builder *Builder) AddObjectsFromResource(resourcePath string, objectIds []string) error {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out
	var _arg2 **C.char      // out
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(resourcePath)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		_arg2 = (**C.char)(C.malloc(C.ulong(len(objectIds)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg2))
		{
			out := unsafe.Slice(_arg2, len(objectIds)+1)
			var zero *C.char
			out[len(objectIds)] = zero
			for i := range objectIds {
				out[i] = (*C.char)(unsafe.Pointer(C.CString(objectIds[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.gtk_builder_add_objects_from_resource(_arg0, _arg1, _arg2, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// AddObjectsFromString parses a string containing a UI definition, building
// only the requested objects and merges them with the current contents of
// builder.
//
// Upon errors FALSE will be returned and error will be assigned a GError from
// the GTK_BUILDER_ERROR or G_MARKUP_ERROR domain.
//
// If you are adding an object that depends on an object that is not its child
// (for instance a GtkTreeView that depends on its GtkTreeModel), you have to
// explicitly list all of them in object_ids.
func (builder *Builder) AddObjectsFromString(buffer string, length int, objectIds []string) error {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out
	var _arg2 C.gssize      // out
	var _arg3 **C.char      // out
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(buffer)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(length)
	{
		_arg3 = (**C.char)(C.malloc(C.ulong(len(objectIds)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg3))
		{
			out := unsafe.Slice(_arg3, len(objectIds)+1)
			var zero *C.char
			out[len(objectIds)] = zero
			for i := range objectIds {
				out[i] = (*C.char)(unsafe.Pointer(C.CString(objectIds[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	C.gtk_builder_add_objects_from_string(_arg0, _arg1, _arg2, _arg3, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// ExposeObject: add object to the builder object pool so it can be referenced
// just like any other object built by builder.
func (builder *Builder) ExposeObject(name string, object *externglib.Object) {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out
	var _arg2 *C.GObject    // out

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GObject)(unsafe.Pointer(object.Native()))

	C.gtk_builder_expose_object(_arg0, _arg1, _arg2)
}

// ExtendWithTemplate: main private entry point for building composite
// components from template XML.
//
// This is exported purely to let gtk-builder-tool validate templates,
// applications have no need to call this function.
func (builder *Builder) ExtendWithTemplate(object *externglib.Object, templateType externglib.Type, buffer string, length int) error {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.GObject    // out
	var _arg2 C.GType       // out
	var _arg3 *C.char       // out
	var _arg4 C.gssize      // out
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.GObject)(unsafe.Pointer(object.Native()))
	_arg2 = C.GType(templateType)
	_arg3 = (*C.char)(unsafe.Pointer(C.CString(buffer)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = C.gssize(length)

	C.gtk_builder_extend_with_template(_arg0, _arg1, _arg2, _arg3, _arg4, &_cerr)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// CurrentObject gets the current object set via
// gtk_builder_set_current_object().
func (builder *Builder) CurrentObject() *externglib.Object {
	var _arg0 *C.GtkBuilder // out
	var _cret *C.GObject    // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))

	_cret = C.gtk_builder_get_current_object(_arg0)

	var _object *externglib.Object // out

	_object = externglib.Take(unsafe.Pointer(_cret))

	return _object
}

// GetObject gets the object named name.
//
// Note that this function does not increment the reference count of the
// returned object.
func (builder *Builder) GetObject(name string) *externglib.Object {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out
	var _cret *C.GObject    // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_builder_get_object(_arg0, _arg1)

	var _object *externglib.Object // out

	_object = externglib.Take(unsafe.Pointer(_cret))

	return _object
}

// Scope gets the scope in use that was set via gtk_builder_set_scope().
func (builder *Builder) Scope() BuilderScoper {
	var _arg0 *C.GtkBuilder      // out
	var _cret *C.GtkBuilderScope // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))

	_cret = C.gtk_builder_get_scope(_arg0)

	var _builderScope BuilderScoper // out

	_builderScope = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(BuilderScoper)

	return _builderScope
}

// TranslationDomain gets the translation domain of builder.
func (builder *Builder) TranslationDomain() string {
	var _arg0 *C.GtkBuilder // out
	var _cret *C.char       // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))

	_cret = C.gtk_builder_get_translation_domain(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// TypeFromName looks up a type by name.
//
// This is using the virtual function that GtkBuilder has for that purpose. This
// is mainly used when implementing the GtkBuildable interface on a type.
func (builder *Builder) TypeFromName(typeName string) externglib.Type {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out
	var _cret C.GType       // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(typeName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gtk_builder_get_type_from_name(_arg0, _arg1)

	var _gType externglib.Type // out

	_gType = externglib.Type(_cret)

	return _gType
}

// SetCurrentObject sets the current object for the builder.
//
// The current object can be thought of as the this object that the builder is
// working for and will often be used as the default object when an object is
// optional.
//
// gtk.Widget.InitTemplate() for example will set the current object to the
// widget the template is inited for. For functions like
// gtk.Builder.NewFromResource, the current object will be NULL.
func (builder *Builder) SetCurrentObject(currentObject *externglib.Object) {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.GObject    // out

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.GObject)(unsafe.Pointer(currentObject.Native()))

	C.gtk_builder_set_current_object(_arg0, _arg1)
}

// SetScope sets the scope the builder should operate in.
//
// If scope is NULL a new gtk.BuilderCScope will be created.
func (builder *Builder) SetScope(scope BuilderScoper) {
	var _arg0 *C.GtkBuilder      // out
	var _arg1 *C.GtkBuilderScope // out

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = (*C.GtkBuilderScope)(unsafe.Pointer(scope.Native()))

	C.gtk_builder_set_scope(_arg0, _arg1)
}

// SetTranslationDomain sets the translation domain of builder.
func (builder *Builder) SetTranslationDomain(domain string) {
	var _arg0 *C.GtkBuilder // out
	var _arg1 *C.char       // out

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	if domain != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(domain)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_builder_set_translation_domain(_arg0, _arg1)
}

// ValueFromStringType demarshals a value from a string.
//
// Unlike gtk.Builder.ValueFromString(), this function takes a GType instead of
// GParamSpec.
//
// Calls g_value_init() on the value argument, so it need not be initialised
// beforehand.
//
// Upon errors FALSE will be returned and error will be assigned a GError from
// the GTK_BUILDER_ERROR domain.
func (builder *Builder) ValueFromStringType(typ externglib.Type, _string string) (externglib.Value, error) {
	var _arg0 *C.GtkBuilder // out
	var _arg1 C.GType       // out
	var _arg2 *C.char       // out
	var _arg3 C.GValue      // in
	var _cerr *C.GError     // in

	_arg0 = (*C.GtkBuilder)(unsafe.Pointer(builder.Native()))
	_arg1 = C.GType(typ)
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(_string)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_builder_value_from_string_type(_arg0, _arg1, _arg2, &_arg3, &_cerr)

	var _value externglib.Value // out
	var _goerr error            // out

	_value = *externglib.ValueFromNative(unsafe.Pointer((&_arg3)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _value, _goerr
}
