// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_notebook_tab_get_type()), F: marshalNotebookTab},
		{T: externglib.Type(C.gtk_notebook_get_type()), F: marshalNotebooker},
		{T: externglib.Type(C.gtk_notebook_page_get_type()), F: marshalNotebookPager},
	})
}

// NotebookTab: parameter used in the action signals of GtkNotebook.
type NotebookTab int

const (
	// NotebookTabFirst tab in the notebook
	NotebookTabFirst NotebookTab = iota
	// NotebookTabLast tab in the notebook
	NotebookTabLast
)

func marshalNotebookTab(p uintptr) (interface{}, error) {
	return NotebookTab(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for NotebookTab.
func (n NotebookTab) String() string {
	switch n {
	case NotebookTabFirst:
		return "First"
	case NotebookTabLast:
		return "Last"
	default:
		return fmt.Sprintf("NotebookTab(%d)", n)
	}
}

// Notebook: GtkNotebook is a container whose children are pages switched
// between using tabs.
//
// !An example GtkNotebook (notebook.png)
//
// There are many configuration options for GtkNotebook. Among other things, you
// can choose on which edge the tabs appear (see gtk.Notebook.SetTabPos()),
// whether, if there are too many tabs to fit the notebook should be made bigger
// or scrolling arrows added (see gtk.Notebook.SetScrollable()), and whether
// there will be a popup menu allowing the users to switch pages. (see
// gtk.Notebook.PopupEnable()).
//
//
// GtkNotebook as GtkBuildable
//
// The GtkNotebook implementation of the GtkBuildable interface supports placing
// children into tabs by specifying “tab” as the “type” attribute of a <child>
// element. Note that the content of the tab must be created before the tab can
// be filled. A tab child can be specified without specifying a <child> type
// attribute.
//
// To add a child widget in the notebooks action area, specify "action-start" or
// “action-end” as the “type” attribute of the <child> element.
//
// An example of a UI definition fragment with GtkNotebook:
//
//    <object class="GtkNotebook">
//      <child>
//        <object class="GtkLabel" id="notebook-content">
//          <property name="label">Content</property>
//        </object>
//      </child>
//      <child type="tab">
//        <object class="GtkLabel" id="notebook-tab">
//          <property name="label">Tab</property>
//        </object>
//      </child>
//    </object>
//
//
// CSS nodes
//
//    notebook
//    ├── header.top
//    │   ├── [<action widget>]
//    │   ├── tabs
//    │   │   ├── [arrow]
//    │   │   ├── tab
//    │   │   │   ╰── <tab label>
//    ┊   ┊   ┊
//    │   │   ├── tab[.reorderable-page]
//    │   │   │   ╰── <tab label>
//    │   │   ╰── [arrow]
//    │   ╰── [<action widget>]
//    │
//    ╰── stack
//        ├── <child>
//        ┊
//        ╰── <child>
//
//
// GtkNotebook has a main CSS node with name notebook, a subnode with name
// header and below that a subnode with name tabs which contains one subnode per
// tab with name tab.
//
// If action widgets are present, their CSS nodes are placed next to the tabs
// node. If the notebook is scrollable, CSS nodes with name arrow are placed as
// first and last child of the tabs node.
//
// The main node gets the .frame style class when the notebook has a border (see
// gtk.Notebook.SetShowBorder()).
//
// The header node gets one of the style class .top, .bottom, .left or .right,
// depending on where the tabs are placed. For reorderable pages, the tab node
// gets the .reorderable-page class.
//
// A tab node gets the .dnd style class while it is moved with drag-and-drop.
//
// The nodes are always arranged from left-to-right, regardless of text
// direction.
//
//
// Accessibility
//
// GtkNotebook uses the following roles:
//
//    - GTK_ACCESSIBLE_ROLE_GROUP for the notebook widget
//    - GTK_ACCESSIBLE_ROLE_TAB_LIST for the list of tabs
//    - GTK_ACCESSIBLE_ROLE_TAB role for each tab
//    - GTK_ACCESSIBLE_ROLE_TAB_PANEL for each page
type Notebook struct {
	Widget
}

func wrapNotebook(obj *externglib.Object) *Notebook {
	return &Notebook{
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
	}
}

func marshalNotebooker(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapNotebook(obj), nil
}

// NewNotebook creates a new GtkNotebook widget with no pages.
func NewNotebook() *Notebook {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_notebook_new()

	var _notebook *Notebook // out

	_notebook = wrapNotebook(externglib.Take(unsafe.Pointer(_cret)))

	return _notebook
}

// AppendPage appends a page to notebook.
func (notebook *Notebook) AppendPage(child Widgetter, tabLabel Widgetter) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _cret C.int          // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))

	_cret = C.gtk_notebook_append_page(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// AppendPageMenu appends a page to notebook, specifying the widget to use as
// the label in the popup menu.
func (notebook *Notebook) AppendPageMenu(child Widgetter, tabLabel Widgetter, menuLabel Widgetter) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _arg3 *C.GtkWidget   // out
	var _cret C.int          // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))
	_arg3 = (*C.GtkWidget)(unsafe.Pointer(menuLabel.Native()))

	_cret = C.gtk_notebook_append_page_menu(_arg0, _arg1, _arg2, _arg3)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// DetachTab removes the child from the notebook.
//
// This function is very similar to gtk.Notebook.RemovePage(), but additionally
// informs the notebook that the removal is happening as part of a tab DND
// operation, which should not be cancelled.
func (notebook *Notebook) DetachTab(child Widgetter) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_notebook_detach_tab(_arg0, _arg1)
}

// ActionWidget gets one of the action widgets.
//
// See gtk.Notebook.SetActionWidget().
func (notebook *Notebook) ActionWidget(packType PackType) Widgetter {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.GtkPackType  // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = C.GtkPackType(packType)

	_cret = C.gtk_notebook_get_action_widget(_arg0, _arg1)

	var _widget Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// CurrentPage returns the page number of the current page.
func (notebook *Notebook) CurrentPage() int {
	var _arg0 *C.GtkNotebook // out
	var _cret C.int          // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_current_page(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// GroupName gets the current group name for notebook.
func (notebook *Notebook) GroupName() string {
	var _arg0 *C.GtkNotebook // out
	var _cret *C.char        // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_group_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// MenuLabel retrieves the menu label widget of the page containing child.
func (notebook *Notebook) MenuLabel(child Widgetter) Widgetter {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_get_menu_label(_arg0, _arg1)

	var _widget Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// MenuLabelText retrieves the text of the menu label for the page containing
// child.
func (notebook *Notebook) MenuLabelText(child Widgetter) string {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret *C.char        // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_get_menu_label_text(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// NPages gets the number of pages in a notebook.
func (notebook *Notebook) NPages() int {
	var _arg0 *C.GtkNotebook // out
	var _cret C.int          // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_n_pages(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NthPage returns the child widget contained in page number page_num.
func (notebook *Notebook) NthPage(pageNum int) Widgetter {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.int          // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = C.int(pageNum)

	_cret = C.gtk_notebook_get_nth_page(_arg0, _arg1)

	var _widget Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// Page returns the GtkNotebookPage for child.
func (notebook *Notebook) Page(child Widgetter) *NotebookPage {
	var _arg0 *C.GtkNotebook     // out
	var _arg1 *C.GtkWidget       // out
	var _cret *C.GtkNotebookPage // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_get_page(_arg0, _arg1)

	var _notebookPage *NotebookPage // out

	_notebookPage = wrapNotebookPage(externglib.Take(unsafe.Pointer(_cret)))

	return _notebookPage
}

// Pages returns a GListModel that contains the pages of the notebook.
//
// This can be used to keep an up-to-date view. The model also implements
// gtk.SelectionModel and can be used to track and modify the visible page.
func (notebook *Notebook) Pages() gio.ListModeller {
	var _arg0 *C.GtkNotebook // out
	var _cret *C.GListModel  // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_pages(_arg0)

	var _listModel gio.ListModeller // out

	_listModel = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(gio.ListModeller)

	return _listModel
}

// Scrollable returns whether the tab label area has arrows for scrolling.
func (notebook *Notebook) Scrollable() bool {
	var _arg0 *C.GtkNotebook // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_scrollable(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowBorder returns whether a bevel will be drawn around the notebook pages.
func (notebook *Notebook) ShowBorder() bool {
	var _arg0 *C.GtkNotebook // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_show_border(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ShowTabs returns whether the tabs of the notebook are shown.
func (notebook *Notebook) ShowTabs() bool {
	var _arg0 *C.GtkNotebook // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_show_tabs(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TabDetachable returns whether the tab contents can be detached from notebook.
func (notebook *Notebook) TabDetachable(child Widgetter) bool {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_get_tab_detachable(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TabLabel returns the tab label widget for the page child.
//
// NULL is returned if child is not in notebook or if no tab label has
// specifically been set for child.
func (notebook *Notebook) TabLabel(child Widgetter) Widgetter {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret *C.GtkWidget   // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_get_tab_label(_arg0, _arg1)

	var _widget Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}

// TabLabelText retrieves the text of the tab label for the page containing
// child.
func (notebook *Notebook) TabLabelText(child Widgetter) string {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret *C.char        // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_get_tab_label_text(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// TabPos gets the edge at which the tabs are drawn.
func (notebook *Notebook) TabPos() PositionType {
	var _arg0 *C.GtkNotebook    // out
	var _cret C.GtkPositionType // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	_cret = C.gtk_notebook_get_tab_pos(_arg0)

	var _positionType PositionType // out

	_positionType = PositionType(_cret)

	return _positionType
}

// TabReorderable gets whether the tab can be reordered via drag and drop or
// not.
func (notebook *Notebook) TabReorderable(child Widgetter) bool {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_get_tab_reorderable(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InsertPage: insert a page into notebook at the given position.
func (notebook *Notebook) InsertPage(child Widgetter, tabLabel Widgetter, position int) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _arg3 C.int          // out
	var _cret C.int          // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))
	_arg3 = C.int(position)

	_cret = C.gtk_notebook_insert_page(_arg0, _arg1, _arg2, _arg3)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// InsertPageMenu: insert a page into notebook at the given position, specifying
// the widget to use as the label in the popup menu.
func (notebook *Notebook) InsertPageMenu(child Widgetter, tabLabel Widgetter, menuLabel Widgetter, position int) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _arg3 *C.GtkWidget   // out
	var _arg4 C.int          // out
	var _cret C.int          // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))
	_arg3 = (*C.GtkWidget)(unsafe.Pointer(menuLabel.Native()))
	_arg4 = C.int(position)

	_cret = C.gtk_notebook_insert_page_menu(_arg0, _arg1, _arg2, _arg3, _arg4)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// NextPage switches to the next page.
//
// Nothing happens if the current page is the last page.
func (notebook *Notebook) NextPage() {
	var _arg0 *C.GtkNotebook // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	C.gtk_notebook_next_page(_arg0)
}

// PageNum finds the index of the page which contains the given child widget.
func (notebook *Notebook) PageNum(child Widgetter) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _cret C.int          // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	_cret = C.gtk_notebook_page_num(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PopupDisable disables the popup menu.
func (notebook *Notebook) PopupDisable() {
	var _arg0 *C.GtkNotebook // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	C.gtk_notebook_popup_disable(_arg0)
}

// PopupEnable enables the popup menu.
//
// If the user clicks with the right mouse button on the tab labels, a menu with
// all the pages will be popped up.
func (notebook *Notebook) PopupEnable() {
	var _arg0 *C.GtkNotebook // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	C.gtk_notebook_popup_enable(_arg0)
}

// PrependPage prepends a page to notebook.
func (notebook *Notebook) PrependPage(child Widgetter, tabLabel Widgetter) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _cret C.int          // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))

	_cret = C.gtk_notebook_prepend_page(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PrependPageMenu prepends a page to notebook, specifying the widget to use as
// the label in the popup menu.
func (notebook *Notebook) PrependPageMenu(child Widgetter, tabLabel Widgetter, menuLabel Widgetter) int {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out
	var _arg3 *C.GtkWidget   // out
	var _cret C.int          // in

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))
	_arg3 = (*C.GtkWidget)(unsafe.Pointer(menuLabel.Native()))

	_cret = C.gtk_notebook_prepend_page_menu(_arg0, _arg1, _arg2, _arg3)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// PrevPage switches to the previous page.
//
// Nothing happens if the current page is the first page.
func (notebook *Notebook) PrevPage() {
	var _arg0 *C.GtkNotebook // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))

	C.gtk_notebook_prev_page(_arg0)
}

// RemovePage removes a page from the notebook given its index in the notebook.
func (notebook *Notebook) RemovePage(pageNum int) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = C.int(pageNum)

	C.gtk_notebook_remove_page(_arg0, _arg1)
}

// ReorderChild reorders the page containing child, so that it appears in
// position position.
//
// If position is greater than or equal to the number of children in the list or
// negative, child will be moved to the end of the list.
func (notebook *Notebook) ReorderChild(child Widgetter, position int) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 C.int          // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.int(position)

	C.gtk_notebook_reorder_child(_arg0, _arg1, _arg2)
}

// SetActionWidget sets widget as one of the action widgets.
//
// Depending on the pack type the widget will be placed before or after the
// tabs. You can use a GtkBox if you need to pack more than one widget on the
// same side.
func (notebook *Notebook) SetActionWidget(widget Widgetter, packType PackType) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 C.GtkPackType  // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = C.GtkPackType(packType)

	C.gtk_notebook_set_action_widget(_arg0, _arg1, _arg2)
}

// SetCurrentPage switches to the page number page_num.
//
// Note that due to historical reasons, GtkNotebook refuses to switch to a page
// unless the child widget is visible. Therefore, it is recommended to show
// child widgets before adding them to a notebook.
func (notebook *Notebook) SetCurrentPage(pageNum int) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.int          // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = C.int(pageNum)

	C.gtk_notebook_set_current_page(_arg0, _arg1)
}

// SetGroupName sets a group name for notebook.
//
// Notebooks with the same name will be able to exchange tabs via drag and drop.
// A notebook with a NULL group name will not be able to exchange tabs with any
// other notebook.
func (notebook *Notebook) SetGroupName(groupName string) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.char        // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	if groupName != "" {
		_arg1 = (*C.char)(unsafe.Pointer(C.CString(groupName)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	C.gtk_notebook_set_group_name(_arg0, _arg1)
}

// SetMenuLabel changes the menu label for the page containing child.
func (notebook *Notebook) SetMenuLabel(child Widgetter, menuLabel Widgetter) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(menuLabel.Native()))

	C.gtk_notebook_set_menu_label(_arg0, _arg1, _arg2)
}

// SetMenuLabelText creates a new label and sets it as the menu label of child.
func (notebook *Notebook) SetMenuLabelText(child Widgetter, menuText string) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.char        // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(menuText)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_notebook_set_menu_label_text(_arg0, _arg1, _arg2)
}

// SetScrollable sets whether the tab label area will have arrows for scrolling
// if there are too many tabs to fit in the area.
func (notebook *Notebook) SetScrollable(scrollable bool) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	if scrollable {
		_arg1 = C.TRUE
	}

	C.gtk_notebook_set_scrollable(_arg0, _arg1)
}

// SetShowBorder sets whether a bevel will be drawn around the notebook pages.
//
// This only has a visual effect when the tabs are not shown.
func (notebook *Notebook) SetShowBorder(showBorder bool) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	if showBorder {
		_arg1 = C.TRUE
	}

	C.gtk_notebook_set_show_border(_arg0, _arg1)
}

// SetShowTabs sets whether to show the tabs for the notebook or not.
func (notebook *Notebook) SetShowTabs(showTabs bool) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	if showTabs {
		_arg1 = C.TRUE
	}

	C.gtk_notebook_set_show_tabs(_arg0, _arg1)
}

// SetTabDetachable sets whether the tab can be detached from notebook to
// another notebook or widget.
//
// Note that two notebooks must share a common group identificator (see
// gtk.Notebook.SetGroupName()) to allow automatic tabs interchange between
// them.
//
// If you want a widget to interact with a notebook through DnD (i.e.: accept
// dragged tabs from it) it must be set as a drop destination and accept the
// target “GTK_NOTEBOOK_TAB”. The notebook will fill the selection with a
// GtkWidget** pointing to the child widget that corresponds to the dropped tab.
//
// Note that you should use gtk.Notebook.DetachTab() instead of
// gtk.Notebook.RemovePage() if you want to remove the tab from the source
// notebook as part of accepting a drop. Otherwise, the source notebook will
// think that the dragged tab was removed from underneath the ongoing drag
// operation, and will initiate a drag cancel animation.
//
//    static void
//    on_drag_data_received (GtkWidget        *widget,
//                           GdkDrop          *drop,
//                           GtkSelectionData *data,
//                           guint             time,
//                           gpointer          user_data)
//    {
//      GtkDrag *drag;
//      GtkWidget *notebook;
//      GtkWidget **child;
//
//      drag = gtk_drop_get_drag (drop);
//      notebook = g_object_get_data (drag, "gtk-notebook-drag-origin");
//      child = (void*) gtk_selection_data_get_data (data);
//
//      // process_widget (*child);
//
//      gtk_notebook_detach_tab (GTK_NOTEBOOK (notebook), *child);
//    }
//
//
// If you want a notebook to accept drags from other widgets, you will have to
// set your own DnD code to do it.
func (notebook *Notebook) SetTabDetachable(child Widgetter, detachable bool) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 C.gboolean     // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if detachable {
		_arg2 = C.TRUE
	}

	C.gtk_notebook_set_tab_detachable(_arg0, _arg1, _arg2)
}

// SetTabLabel changes the tab label for child.
//
// If NULL is specified for tab_label, then the page will have the label “page
// N”.
func (notebook *Notebook) SetTabLabel(child Widgetter, tabLabel Widgetter) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.GtkWidget   // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(tabLabel.Native()))

	C.gtk_notebook_set_tab_label(_arg0, _arg1, _arg2)
}

// SetTabLabelText creates a new label and sets it as the tab label for the page
// containing child.
func (notebook *Notebook) SetTabLabelText(child Widgetter, tabText string) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 *C.char        // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(tabText)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_notebook_set_tab_label_text(_arg0, _arg1, _arg2)
}

// SetTabPos sets the edge at which the tabs are drawn.
func (notebook *Notebook) SetTabPos(pos PositionType) {
	var _arg0 *C.GtkNotebook    // out
	var _arg1 C.GtkPositionType // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = C.GtkPositionType(pos)

	C.gtk_notebook_set_tab_pos(_arg0, _arg1)
}

// SetTabReorderable sets whether the notebook tab can be reordered via drag and
// drop or not.
func (notebook *Notebook) SetTabReorderable(child Widgetter, reorderable bool) {
	var _arg0 *C.GtkNotebook // out
	var _arg1 *C.GtkWidget   // out
	var _arg2 C.gboolean     // out

	_arg0 = (*C.GtkNotebook)(unsafe.Pointer(notebook.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	if reorderable {
		_arg2 = C.TRUE
	}

	C.gtk_notebook_set_tab_reorderable(_arg0, _arg1, _arg2)
}

// NotebookPage: GtkNotebookPage is an auxiliary object used by GtkNotebook.
type NotebookPage struct {
	*externglib.Object
}

func wrapNotebookPage(obj *externglib.Object) *NotebookPage {
	return &NotebookPage{
		Object: obj,
	}
}

func marshalNotebookPager(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapNotebookPage(obj), nil
}

// Child returns the notebook child to which page belongs.
func (page *NotebookPage) Child() Widgetter {
	var _arg0 *C.GtkNotebookPage // out
	var _cret *C.GtkWidget       // in

	_arg0 = (*C.GtkNotebookPage)(unsafe.Pointer(page.Native()))

	_cret = C.gtk_notebook_page_get_child(_arg0)

	var _widget Widgetter // out

	_widget = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(Widgetter)

	return _widget
}
