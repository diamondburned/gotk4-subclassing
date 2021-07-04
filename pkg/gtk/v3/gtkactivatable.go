// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
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
		{T: externglib.Type(C.gtk_activatable_get_type()), F: marshalActivatable},
	})
}

// Activatable: activatable widgets can be connected to a Action and reflects
// the state of its action. A Activatable can also provide feedback through its
// action, as they are responsible for activating their related actions.
//
//
// Implementing GtkActivatable
//
// When extending a class that is already Activatable; it is only necessary to
// implement the Activatable->sync_action_properties() and Activatable->update()
// methods and chain up to the parent implementation, however when introducing a
// new Activatable class; the Activatable:related-action and
// Activatable:use-action-appearance properties need to be handled by the
// implementor. Handling these properties is mostly a matter of installing the
// action pointer and boolean flag on your instance, and calling
// gtk_activatable_do_set_related_action() and
// gtk_activatable_sync_action_properties() at the appropriate times.
//
// A class fragment implementing Activatable
//
//
//    enum {
//    ...
//
//    PROP_ACTIVATABLE_RELATED_ACTION,
//    PROP_ACTIVATABLE_USE_ACTION_APPEARANCE
//    }
//
//    struct _FooBarPrivate
//    {
//
//      ...
//
//      GtkAction      *action;
//      gboolean        use_action_appearance;
//    };
//
//    ...
//
//    static void foo_bar_activatable_interface_init         (GtkActivatableIface  *iface);
//    static void foo_bar_activatable_update                 (GtkActivatable       *activatable,
//    						           GtkAction            *action,
//    						           const gchar          *property_name);
//    static void foo_bar_activatable_sync_action_properties (GtkActivatable       *activatable,
//    						           GtkAction            *action);
//    ...
//
//
//    static void
//    foo_bar_class_init (FooBarClass *klass)
//    {
//
//      ...
//
//      g_object_class_override_property (gobject_class, PROP_ACTIVATABLE_RELATED_ACTION, "related-action");
//      g_object_class_override_property (gobject_class, PROP_ACTIVATABLE_USE_ACTION_APPEARANCE, "use-action-appearance");
//
//      ...
//    }
//
//
//    static void
//    foo_bar_activatable_interface_init (GtkActivatableIface  *iface)
//    {
//      iface->update = foo_bar_activatable_update;
//      iface->sync_action_properties = foo_bar_activatable_sync_action_properties;
//    }
//
//    ... Break the reference using gtk_activatable_do_set_related_action()...
//
//    static void
//    foo_bar_dispose (GObject *object)
//    {
//      FooBar *bar = FOO_BAR (object);
//      FooBarPrivate *priv = FOO_BAR_GET_PRIVATE (bar);
//
//      ...
//
//      if (priv->action)
//        {
//          gtk_activatable_do_set_related_action (GTK_ACTIVATABLE (bar), NULL);
//          priv->action = NULL;
//        }
//      G_OBJECT_CLASS (foo_bar_parent_class)->dispose (object);
//    }
//
//    ... Handle the “related-action” and “use-action-appearance” properties ...
//
//    static void
//    foo_bar_set_property (GObject         *object,
//                          guint            prop_id,
//                          const GValue    *value,
//                          GParamSpec      *pspec)
//    {
//      FooBar *bar = FOO_BAR (object);
//      FooBarPrivate *priv = FOO_BAR_GET_PRIVATE (bar);
//
//      switch (prop_id)
//        {
//
//          ...
//
//        case PROP_ACTIVATABLE_RELATED_ACTION:
//          foo_bar_set_related_action (bar, g_value_get_object (value));
//          break;
//        case PROP_ACTIVATABLE_USE_ACTION_APPEARANCE:
//          foo_bar_set_use_action_appearance (bar, g_value_get_boolean (value));
//          break;
//        default:
//          G_OBJECT_WARN_INVALID_PROPERTY_ID (object, prop_id, pspec);
//          break;
//        }
//    }
//
//    static void
//    foo_bar_get_property (GObject         *object,
//                             guint            prop_id,
//                             GValue          *value,
//                             GParamSpec      *pspec)
//    {
//      FooBar *bar = FOO_BAR (object);
//      FooBarPrivate *priv = FOO_BAR_GET_PRIVATE (bar);
//
//      switch (prop_id)
//        {
//
//          ...
//
//        case PROP_ACTIVATABLE_RELATED_ACTION:
//          g_value_set_object (value, priv->action);
//          break;
//        case PROP_ACTIVATABLE_USE_ACTION_APPEARANCE:
//          g_value_set_boolean (value, priv->use_action_appearance);
//          break;
//        default:
//          G_OBJECT_WARN_INVALID_PROPERTY_ID (object, prop_id, pspec);
//          break;
//        }
//    }
//
//
//    static void
//    foo_bar_set_use_action_appearance (FooBar   *bar,
//    				   gboolean  use_appearance)
//    {
//      FooBarPrivate *priv = FOO_BAR_GET_PRIVATE (bar);
//
//      if (priv->use_action_appearance != use_appearance)
//        {
//          priv->use_action_appearance = use_appearance;
//
//          gtk_activatable_sync_action_properties (GTK_ACTIVATABLE (bar), priv->action);
//        }
//    }
//
//    ... call gtk_activatable_do_set_related_action() and then assign the action pointer,
//    no need to reference the action here since gtk_activatable_do_set_related_action() already
//    holds a reference here for you...
//    static void
//    foo_bar_set_related_action (FooBar    *bar,
//    			    GtkAction *action)
//    {
//      FooBarPrivate *priv = FOO_BAR_GET_PRIVATE (bar);
//
//      if (priv->action == action)
//        return;
//
//      gtk_activatable_do_set_related_action (GTK_ACTIVATABLE (bar), action);
//
//      priv->action = action;
//    }
//
//    ... Selectively reset and update activatable depending on the use-action-appearance property ...
//    static void
//    gtk_button_activatable_sync_action_properties (GtkActivatable       *activatable,
//    		                                  GtkAction            *action)
//    {
//      GtkButtonPrivate *priv = GTK_BUTTON_GET_PRIVATE (activatable);
//
//      if (!action)
//        return;
//
//      if (gtk_action_is_visible (action))
//        gtk_widget_show (GTK_WIDGET (activatable));
//      else
//        gtk_widget_hide (GTK_WIDGET (activatable));
//
//      gtk_widget_set_sensitive (GTK_WIDGET (activatable), gtk_action_is_sensitive (action));
//
//      ...
//
//      if (priv->use_action_appearance)
//        {
//          if (gtk_action_get_stock_id (action))
//    	foo_bar_set_stock (button, gtk_action_get_stock_id (action));
//          else if (gtk_action_get_label (action))
//    	foo_bar_set_label (button, gtk_action_get_label (action));
//
//          ...
//
//        }
//    }
//
//    static void
//    foo_bar_activatable_update (GtkActivatable       *activatable,
//    			       GtkAction            *action,
//    			       const gchar          *property_name)
//    {
//      FooBarPrivate *priv = FOO_BAR_GET_PRIVATE (activatable);
//
//      if (strcmp (property_name, "visible") == 0)
//        {
//          if (gtk_action_is_visible (action))
//    	gtk_widget_show (GTK_WIDGET (activatable));
//          else
//    	gtk_widget_hide (GTK_WIDGET (activatable));
//        }
//      else if (strcmp (property_name, "sensitive") == 0)
//        gtk_widget_set_sensitive (GTK_WIDGET (activatable), gtk_action_is_sensitive (action));
//
//      ...
//
//      if (!priv->use_action_appearance)
//        return;
//
//      if (strcmp (property_name, "stock-id") == 0)
//        foo_bar_set_stock (button, gtk_action_get_stock_id (action));
//      else if (strcmp (property_name, "label") == 0)
//        foo_bar_set_label (button, gtk_action_get_label (action));
//
//      ...
//    }
type Activatable interface {
	gextras.Objector

	// DoSetRelatedAction: this is called to update the activatable completely,
	// this is called internally when the Activatable:related-action property is
	// set or unset and by the implementing class when
	// Activatable:use-action-appearance changes.
	//
	// Deprecated: since version 3.10.
	DoSetRelatedAction(action Action)
	// RelatedAction: this is called to update the activatable completely, this
	// is called internally when the Activatable:related-action property is set
	// or unset and by the implementing class when
	// Activatable:use-action-appearance changes.
	//
	// Deprecated: since version 3.10.
	RelatedAction() Action
	// UseActionAppearance: this is called to update the activatable completely,
	// this is called internally when the Activatable:related-action property is
	// set or unset and by the implementing class when
	// Activatable:use-action-appearance changes.
	//
	// Deprecated: since version 3.10.
	UseActionAppearance() bool
	// SetRelatedAction: this is called to update the activatable completely,
	// this is called internally when the Activatable:related-action property is
	// set or unset and by the implementing class when
	// Activatable:use-action-appearance changes.
	//
	// Deprecated: since version 3.10.
	SetRelatedAction(action Action)
	// SetUseActionAppearance: this is called to update the activatable
	// completely, this is called internally when the Activatable:related-action
	// property is set or unset and by the implementing class when
	// Activatable:use-action-appearance changes.
	//
	// Deprecated: since version 3.10.
	SetUseActionAppearance(useAppearance bool)
	// SyncActionProperties: this is called to update the activatable
	// completely, this is called internally when the Activatable:related-action
	// property is set or unset and by the implementing class when
	// Activatable:use-action-appearance changes.
	//
	// Deprecated: since version 3.10.
	SyncActionProperties(action Action)
}

// activatable implements the Activatable interface.
type activatable struct {
	gextras.Objector
}

var _ Activatable = (*activatable)(nil)

// WrapActivatable wraps a GObject to a type that implements
// interface Activatable. It is primarily used internally.
func WrapActivatable(obj *externglib.Object) Activatable {
	return activatable{
		Objector: obj,
	}
}

func marshalActivatable(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapActivatable(obj), nil
}

func (a activatable) DoSetRelatedAction(action Action) {
	var _arg0 *C.GtkActivatable // out
	var _arg1 *C.GtkAction      // out

	_arg0 = (*C.GtkActivatable)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	C.gtk_activatable_do_set_related_action(_arg0, _arg1)
}

func (a activatable) RelatedAction() Action {
	var _arg0 *C.GtkActivatable // out
	var _cret *C.GtkAction      // in

	_arg0 = (*C.GtkActivatable)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_activatable_get_related_action(_arg0)

	var _action Action // out

	_action = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Action)

	return _action
}

func (a activatable) UseActionAppearance() bool {
	var _arg0 *C.GtkActivatable // out
	var _cret C.gboolean        // in

	_arg0 = (*C.GtkActivatable)(unsafe.Pointer(a.Native()))

	_cret = C.gtk_activatable_get_use_action_appearance(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (a activatable) SetRelatedAction(action Action) {
	var _arg0 *C.GtkActivatable // out
	var _arg1 *C.GtkAction      // out

	_arg0 = (*C.GtkActivatable)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	C.gtk_activatable_set_related_action(_arg0, _arg1)
}

func (a activatable) SetUseActionAppearance(useAppearance bool) {
	var _arg0 *C.GtkActivatable // out
	var _arg1 C.gboolean        // out

	_arg0 = (*C.GtkActivatable)(unsafe.Pointer(a.Native()))
	if useAppearance {
		_arg1 = C.TRUE
	}

	C.gtk_activatable_set_use_action_appearance(_arg0, _arg1)
}

func (a activatable) SyncActionProperties(action Action) {
	var _arg0 *C.GtkActivatable // out
	var _arg1 *C.GtkAction      // out

	_arg0 = (*C.GtkActivatable)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkAction)(unsafe.Pointer(action.Native()))

	C.gtk_activatable_sync_action_properties(_arg0, _arg1)
}