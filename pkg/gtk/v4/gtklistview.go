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
		{T: externglib.Type(C.gtk_list_view_get_type()), F: marshalListView},
	})
}

// ListView: gtkListView is a widget to present a view into a large dynamic list
// of items.
//
// GtkListView uses its factory to generate one row widget for each visible item
// and shows them in a linear display, either vertically or horizontally. The
// ListView:show-separators property offers a simple way to display separators
// between the rows.
//
// GtkListView allows the user to select items according to the selection
// characteristics of the model. For models that allow multiple selected items,
// it is possible to turn on _rubberband selection_, using
// ListView:enable-rubberband.
//
// If you need multiple columns with headers, see ColumnView.
//
// To learn more about the list widget framework, see the overview (Widget).
//
// An example of using GtkListView:
//
//    static void
//    setup_listitem_cb (GtkListItemFactory *factory,
//                       GtkListItem        *list_item)
//    {
//      GtkWidget *image;
//
//      image = gtk_image_new ();
//      gtk_image_set_icon_size (GTK_IMAGE (image), GTK_ICON_SIZE_LARGE);
//      gtk_list_item_set_child (list_item, image);
//    }
//
//    static void
//    bind_listitem_cb (GtkListItemFactory *factory,
//                      GtkListItem        *list_item)
//    {
//      GtkWidget *image;
//      GAppInfo *app_info;
//
//      image = gtk_list_item_get_child (list_item);
//      app_info = gtk_list_item_get_item (list_item);
//      gtk_image_set_from_gicon (GTK_IMAGE (image), g_app_info_get_icon (app_info));
//    }
//
//    static void
//    activate_cb (GtkListView  *list,
//                 guint         position,
//                 gpointer      unused)
//    {
//      GAppInfo *app_info;
//
//      app_info = g_list_model_get_item (G_LIST_MODEL (gtk_list_view_get_model (list)), position);
//      g_app_info_launch (app_info, NULL, NULL, NULL);
//      g_object_unref (app_info);
//    }
//
//    ...
//
//      model = create_application_list ();
//
//      factory = gtk_signal_list_item_factory_new ();
//      g_signal_connect (factory, "setup", G_CALLBACK (setup_listitem_cb), NULL);
//      g_signal_connect (factory, "bind", G_CALLBACK (bind_listitem_cb), NULL);
//
//      list = gtk_list_view_new (GTK_SELECTION_MODEL (gtk_single_selection_new (model)), factory);
//
//      g_signal_connect (list, "activate", G_CALLBACK (activate_cb), NULL);
//
//      gtk_scrolled_window_set_child (GTK_SCROLLED_WINDOW (sw), list);
//
// CSS nodes
//
//    listview[.separators][.rich-list][.navigation-sidebar][.data-table]
//    ├── row
//    │
//    ├── row
//    │
//    ┊
//    ╰── [rubberband]
//
//
// GtkListView uses a single CSS node named listview. It may carry the
// .separators style class, when ListView:show-separators property is set. Each
// child widget uses a single CSS node named row. For rubberband selection, a
// node with name rubberband is used.
//
// The main listview node may also carry style classes to select the style of
// list presentation (ListContainers.html#list-styles): .rich-list,
// .navigation-sidebar or .data-table.
//
//
// Accessibility
//
// GtkListView uses the K_ACCESSIBLE_ROLE_LIST role, and the list items use the
// K_ACCESSIBLE_ROLE_LIST_ITEM role.
type ListView interface {
	ListBase
	Accessible
	Buildable
	ConstraintTarget
	Orientable
	Scrollable

	// EnableRubberband returns whether rows can be selected by dragging with
	// the mouse.
	EnableRubberband(s ListView) bool
	// Factory gets the factory that's currently used to populate list items.
	Factory(s ListView)
	// Model gets the model that's currently used to read the items displayed.
	Model(s ListView)
	// ShowSeparators returns whether the list box should show separators
	// between rows.
	ShowSeparators(s ListView) bool
	// SingleClickActivate returns whether rows will be activated on single
	// click and selected on hover.
	SingleClickActivate(s ListView) bool
	// SetEnableRubberband sets whether selections can be changed by dragging
	// with the mouse.
	SetEnableRubberband(s ListView, enableRubberband bool)
	// SetFactory sets the ListItemFactory to use for populating list items.
	SetFactory(s ListView, factory ListItemFactory)
	// SetModel sets the SelectionModel to use.
	SetModel(s ListView, model SelectionModel)
	// SetShowSeparators sets whether the list box should show separators
	// between rows.
	SetShowSeparators(s ListView, showSeparators bool)
	// SetSingleClickActivate sets whether rows should be activated on single
	// click and selected on hover.
	SetSingleClickActivate(s ListView, singleClickActivate bool)
}

// listView implements the ListView interface.
type listView struct {
	ListBase
	Accessible
	Buildable
	ConstraintTarget
	Orientable
	Scrollable
}

var _ ListView = (*listView)(nil)

// WrapListView wraps a GObject to the right type. It is
// primarily used internally.
func WrapListView(obj *externglib.Object) ListView {
	return ListView{
		ListBase:         WrapListBase(obj),
		Accessible:       WrapAccessible(obj),
		Buildable:        WrapBuildable(obj),
		ConstraintTarget: WrapConstraintTarget(obj),
		Orientable:       WrapOrientable(obj),
		Scrollable:       WrapScrollable(obj),
	}
}

func marshalListView(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapListView(obj), nil
}

// NewListView constructs a class ListView.
func NewListView(model SelectionModel, factory ListItemFactory) {
	var arg1 *C.GtkSelectionModel
	var arg2 *C.GtkListItemFactory

	arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))
	arg2 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))

	C.gtk_list_view_new(arg1, arg2)
}

// EnableRubberband returns whether rows can be selected by dragging with
// the mouse.
func (s listView) EnableRubberband(s ListView) bool {
	var arg0 *C.GtkListView

	arg0 = (*C.GtkListView)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_list_view_get_enable_rubberband(arg0)

	if cret {
		ok = true
	}

	return ok
}

// Factory gets the factory that's currently used to populate list items.
func (s listView) Factory(s ListView) {
	var arg0 *C.GtkListView

	arg0 = (*C.GtkListView)(unsafe.Pointer(s.Native()))

	C.gtk_list_view_get_factory(arg0)
}

// Model gets the model that's currently used to read the items displayed.
func (s listView) Model(s ListView) {
	var arg0 *C.GtkListView

	arg0 = (*C.GtkListView)(unsafe.Pointer(s.Native()))

	C.gtk_list_view_get_model(arg0)
}

// ShowSeparators returns whether the list box should show separators
// between rows.
func (s listView) ShowSeparators(s ListView) bool {
	var arg0 *C.GtkListView

	arg0 = (*C.GtkListView)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_list_view_get_show_separators(arg0)

	if cret {
		ok = true
	}

	return ok
}

// SingleClickActivate returns whether rows will be activated on single
// click and selected on hover.
func (s listView) SingleClickActivate(s ListView) bool {
	var arg0 *C.GtkListView

	arg0 = (*C.GtkListView)(unsafe.Pointer(s.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_list_view_get_single_click_activate(arg0)

	if cret {
		ok = true
	}

	return ok
}

// SetEnableRubberband sets whether selections can be changed by dragging
// with the mouse.
func (s listView) SetEnableRubberband(s ListView, enableRubberband bool) {
	var arg0 *C.GtkListView
	var arg1 C.gboolean

	arg0 = (*C.GtkListView)(unsafe.Pointer(s.Native()))
	if enableRubberband {
		arg1 = C.gboolean(1)
	}

	C.gtk_list_view_set_enable_rubberband(arg0, arg1)
}

// SetFactory sets the ListItemFactory to use for populating list items.
func (s listView) SetFactory(s ListView, factory ListItemFactory) {
	var arg0 *C.GtkListView
	var arg1 *C.GtkListItemFactory

	arg0 = (*C.GtkListView)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkListItemFactory)(unsafe.Pointer(factory.Native()))

	C.gtk_list_view_set_factory(arg0, arg1)
}

// SetModel sets the SelectionModel to use.
func (s listView) SetModel(s ListView, model SelectionModel) {
	var arg0 *C.GtkListView
	var arg1 *C.GtkSelectionModel

	arg0 = (*C.GtkListView)(unsafe.Pointer(s.Native()))
	arg1 = (*C.GtkSelectionModel)(unsafe.Pointer(model.Native()))

	C.gtk_list_view_set_model(arg0, arg1)
}

// SetShowSeparators sets whether the list box should show separators
// between rows.
func (s listView) SetShowSeparators(s ListView, showSeparators bool) {
	var arg0 *C.GtkListView
	var arg1 C.gboolean

	arg0 = (*C.GtkListView)(unsafe.Pointer(s.Native()))
	if showSeparators {
		arg1 = C.gboolean(1)
	}

	C.gtk_list_view_set_show_separators(arg0, arg1)
}

// SetSingleClickActivate sets whether rows should be activated on single
// click and selected on hover.
func (s listView) SetSingleClickActivate(s ListView, singleClickActivate bool) {
	var arg0 *C.GtkListView
	var arg1 C.gboolean

	arg0 = (*C.GtkListView)(unsafe.Pointer(s.Native()))
	if singleClickActivate {
		arg1 = C.gboolean(1)
	}

	C.gtk_list_view_set_single_click_activate(arg0, arg1)
}