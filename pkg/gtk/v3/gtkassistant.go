// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"github.com/diamondburned/gotk4/internal/box"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_assistant_get_type()), F: marshalAssistant},
	})
}

// AssistantPageFunc: a function used by gtk_assistant_set_forward_page_func()
// to know which is the next page given a current one. It’s called both for
// computing the next page when the user presses the “forward” button and for
// handling the behavior of the “last” button.
type AssistantPageFunc func(currentPage int) int

//export gotk4_AssistantPageFunc
func gotk4_AssistantPageFunc(arg0 C.gint, arg1 C.gpointer) C.gint {
	v := box.Get(uintptr(arg1))
	if v == nil {
		panic(`callback not found`)
	}

	fn := v.(AssistantPageFunc)
	ret := fn(currentPage, data)

	cret = C.gint(ret)

	return cret
}

// Assistant: a Assistant is a widget used to represent a generally complex
// operation splitted in several steps, guiding the user through its pages and
// controlling the page flow to collect the necessary data.
//
// The design of GtkAssistant is that it controls what buttons to show and to
// make sensitive, based on what it knows about the page sequence and the
// [type][GtkAssistantPageType] of each page, in addition to state information
// like the page [completion][gtk-assistant-set-page-complete] and
// [committed][gtk-assistant-commit] status.
//
// If you have a case that doesn’t quite fit in Assistants way of handling
// buttons, you can use the K_ASSISTANT_PAGE_CUSTOM page type and handle buttons
// yourself.
//
//
// GtkAssistant as GtkBuildable
//
// The GtkAssistant implementation of the Buildable interface exposes the
// @action_area as internal children with the name “action_area”.
//
// To add pages to an assistant in Builder, simply add it as a child to the
// GtkAssistant object, and set its child properties as necessary.
//
//
// CSS nodes
//
// GtkAssistant has a single CSS node with the name assistant.
type Assistant interface {
	Window
	Buildable

	// AddActionWidget adds a widget to the action area of a Assistant.
	AddActionWidget(a Assistant, child Widget)
	// AppendPage appends a page to the @assistant.
	AppendPage(a Assistant, page Widget)
	// Commit erases the visited page history so the back button is not shown on
	// the current page, and removes the cancel button from subsequent pages.
	//
	// Use this when the information provided up to the current page is
	// hereafter deemed permanent and cannot be modified or undone. For example,
	// showing a progress page to track a long-running, unreversible operation
	// after the user has clicked apply on a confirmation page.
	Commit(a Assistant)
	// CurrentPage returns the page number of the current page.
	CurrentPage(a Assistant)
	// NPages returns the number of pages in the @assistant
	NPages(a Assistant)
	// NthPage returns the child widget contained in page number @page_num.
	NthPage(a Assistant, pageNum int)
	// PageComplete gets whether @page is complete.
	PageComplete(a Assistant, page Widget) bool
	// PageHasPadding gets whether page has padding.
	PageHasPadding(a Assistant, page Widget) bool
	// PageHeaderImage gets the header image for @page.
	PageHeaderImage(a Assistant, page Widget)
	// PageSideImage gets the side image for @page.
	PageSideImage(a Assistant, page Widget)
	// PageTitle gets the title for @page.
	PageTitle(a Assistant, page Widget)
	// PageType gets the page type of @page.
	PageType(a Assistant, page Widget)
	// InsertPage inserts a page in the @assistant at a given position.
	InsertPage(a Assistant, page Widget, position int)
	// NextPage: navigate to the next page.
	//
	// It is a programming error to call this function when there is no next
	// page.
	//
	// This function is for use when creating pages of the
	// K_ASSISTANT_PAGE_CUSTOM type.
	NextPage(a Assistant)
	// PrependPage prepends a page to the @assistant.
	PrependPage(a Assistant, page Widget)
	// PreviousPage: navigate to the previous visited page.
	//
	// It is a programming error to call this function when no previous page is
	// available.
	//
	// This function is for use when creating pages of the
	// K_ASSISTANT_PAGE_CUSTOM type.
	PreviousPage(a Assistant)
	// RemoveActionWidget removes a widget from the action area of a Assistant.
	RemoveActionWidget(a Assistant, child Widget)
	// RemovePage removes the @page_num’s page from @assistant.
	RemovePage(a Assistant, pageNum int)
	// SetCurrentPage switches the page to @page_num.
	//
	// Note that this will only be necessary in custom buttons, as the
	// @assistant flow can be set with gtk_assistant_set_forward_page_func().
	SetCurrentPage(a Assistant, pageNum int)
	// SetForwardPageFunc sets the page forwarding function to be @page_func.
	//
	// This function will be used to determine what will be the next page when
	// the user presses the forward button. Setting @page_func to nil will make
	// the assistant to use the default forward function, which just goes to the
	// next visible page.
	SetForwardPageFunc(a Assistant)
	// SetPageComplete sets whether @page contents are complete.
	//
	// This will make @assistant update the buttons state to be able to continue
	// the task.
	SetPageComplete(a Assistant, page Widget, complete bool)
	// SetPageHasPadding sets whether the assistant is adding padding around the
	// page.
	SetPageHasPadding(a Assistant, page Widget, hasPadding bool)
	// SetPageHeaderImage sets a header image for @page.
	SetPageHeaderImage(a Assistant, page Widget, pixbuf gdkpixbuf.Pixbuf)
	// SetPageSideImage sets a side image for @page.
	//
	// This image used to be displayed in the side area of the assistant when
	// @page is the current page.
	SetPageSideImage(a Assistant, page Widget, pixbuf gdkpixbuf.Pixbuf)
	// SetPageTitle sets a title for @page.
	//
	// The title is displayed in the header area of the assistant when @page is
	// the current page.
	SetPageTitle(a Assistant, page Widget, title string)
	// SetPageType sets the page type for @page.
	//
	// The page type determines the page behavior in the @assistant.
	SetPageType(a Assistant, page Widget, typ AssistantPageType)
	// UpdateButtonsState forces @assistant to recompute the buttons state.
	//
	// GTK+ automatically takes care of this in most situations, e.g. when the
	// user goes to a different page, or when the visibility or completeness of
	// a page changes.
	//
	// One situation where it can be necessary to call this function is when
	// changing a value on the current page affects the future page flow of the
	// assistant.
	UpdateButtonsState(a Assistant)
}

// assistant implements the Assistant interface.
type assistant struct {
	Window
	Buildable
}

var _ Assistant = (*assistant)(nil)

// WrapAssistant wraps a GObject to the right type. It is
// primarily used internally.
func WrapAssistant(obj *externglib.Object) Assistant {
	return Assistant{
		Window:    WrapWindow(obj),
		Buildable: WrapBuildable(obj),
	}
}

func marshalAssistant(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapAssistant(obj), nil
}

// NewAssistant constructs a class Assistant.
func NewAssistant() {
	C.gtk_assistant_new()
}

// AddActionWidget adds a widget to the action area of a Assistant.
func (a assistant) AddActionWidget(a Assistant, child Widget) {
	var arg0 *C.GtkAssistant
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_assistant_add_action_widget(arg0, arg1)
}

// AppendPage appends a page to the @assistant.
func (a assistant) AppendPage(a Assistant, page Widget) {
	var arg0 *C.GtkAssistant
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	C.gtk_assistant_append_page(arg0, arg1)
}

// Commit erases the visited page history so the back button is not shown on
// the current page, and removes the cancel button from subsequent pages.
//
// Use this when the information provided up to the current page is
// hereafter deemed permanent and cannot be modified or undone. For example,
// showing a progress page to track a long-running, unreversible operation
// after the user has clicked apply on a confirmation page.
func (a assistant) Commit(a Assistant) {
	var arg0 *C.GtkAssistant

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))

	C.gtk_assistant_commit(arg0)
}

// CurrentPage returns the page number of the current page.
func (a assistant) CurrentPage(a Assistant) {
	var arg0 *C.GtkAssistant

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))

	C.gtk_assistant_get_current_page(arg0)
}

// NPages returns the number of pages in the @assistant
func (a assistant) NPages(a Assistant) {
	var arg0 *C.GtkAssistant

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))

	C.gtk_assistant_get_n_pages(arg0)
}

// NthPage returns the child widget contained in page number @page_num.
func (a assistant) NthPage(a Assistant, pageNum int) {
	var arg0 *C.GtkAssistant
	var arg1 C.gint

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	arg1 = C.gint(pageNum)

	C.gtk_assistant_get_nth_page(arg0, arg1)
}

// PageComplete gets whether @page is complete.
func (a assistant) PageComplete(a Assistant, page Widget) bool {
	var arg0 *C.GtkAssistant
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_assistant_get_page_complete(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// PageHasPadding gets whether page has padding.
func (a assistant) PageHasPadding(a Assistant, page Widget) bool {
	var arg0 *C.GtkAssistant
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	var cret C.gboolean
	var ok bool

	cret = C.gtk_assistant_get_page_has_padding(arg0, arg1)

	if cret {
		ok = true
	}

	return ok
}

// PageHeaderImage gets the header image for @page.
func (a assistant) PageHeaderImage(a Assistant, page Widget) {
	var arg0 *C.GtkAssistant
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	C.gtk_assistant_get_page_header_image(arg0, arg1)
}

// PageSideImage gets the side image for @page.
func (a assistant) PageSideImage(a Assistant, page Widget) {
	var arg0 *C.GtkAssistant
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	C.gtk_assistant_get_page_side_image(arg0, arg1)
}

// PageTitle gets the title for @page.
func (a assistant) PageTitle(a Assistant, page Widget) {
	var arg0 *C.GtkAssistant
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	C.gtk_assistant_get_page_title(arg0, arg1)
}

// PageType gets the page type of @page.
func (a assistant) PageType(a Assistant, page Widget) {
	var arg0 *C.GtkAssistant
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	C.gtk_assistant_get_page_type(arg0, arg1)
}

// InsertPage inserts a page in the @assistant at a given position.
func (a assistant) InsertPage(a Assistant, page Widget, position int) {
	var arg0 *C.GtkAssistant
	var arg1 *C.GtkWidget
	var arg2 C.gint

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	arg2 = C.gint(position)

	C.gtk_assistant_insert_page(arg0, arg1, arg2)
}

// NextPage: navigate to the next page.
//
// It is a programming error to call this function when there is no next
// page.
//
// This function is for use when creating pages of the
// K_ASSISTANT_PAGE_CUSTOM type.
func (a assistant) NextPage(a Assistant) {
	var arg0 *C.GtkAssistant

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))

	C.gtk_assistant_next_page(arg0)
}

// PrependPage prepends a page to the @assistant.
func (a assistant) PrependPage(a Assistant, page Widget) {
	var arg0 *C.GtkAssistant
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))

	C.gtk_assistant_prepend_page(arg0, arg1)
}

// PreviousPage: navigate to the previous visited page.
//
// It is a programming error to call this function when no previous page is
// available.
//
// This function is for use when creating pages of the
// K_ASSISTANT_PAGE_CUSTOM type.
func (a assistant) PreviousPage(a Assistant) {
	var arg0 *C.GtkAssistant

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))

	C.gtk_assistant_previous_page(arg0)
}

// RemoveActionWidget removes a widget from the action area of a Assistant.
func (a assistant) RemoveActionWidget(a Assistant, child Widget) {
	var arg0 *C.GtkAssistant
	var arg1 *C.GtkWidget

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))

	C.gtk_assistant_remove_action_widget(arg0, arg1)
}

// RemovePage removes the @page_num’s page from @assistant.
func (a assistant) RemovePage(a Assistant, pageNum int) {
	var arg0 *C.GtkAssistant
	var arg1 C.gint

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	arg1 = C.gint(pageNum)

	C.gtk_assistant_remove_page(arg0, arg1)
}

// SetCurrentPage switches the page to @page_num.
//
// Note that this will only be necessary in custom buttons, as the
// @assistant flow can be set with gtk_assistant_set_forward_page_func().
func (a assistant) SetCurrentPage(a Assistant, pageNum int) {
	var arg0 *C.GtkAssistant
	var arg1 C.gint

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	arg1 = C.gint(pageNum)

	C.gtk_assistant_set_current_page(arg0, arg1)
}

// SetForwardPageFunc sets the page forwarding function to be @page_func.
//
// This function will be used to determine what will be the next page when
// the user presses the forward button. Setting @page_func to nil will make
// the assistant to use the default forward function, which just goes to the
// next visible page.
func (a assistant) SetForwardPageFunc(a Assistant) {
	var arg0 *C.GtkAssistant

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))

	C.gtk_assistant_set_forward_page_func(arg0, arg1, arg2, arg3)
}

// SetPageComplete sets whether @page contents are complete.
//
// This will make @assistant update the buttons state to be able to continue
// the task.
func (a assistant) SetPageComplete(a Assistant, page Widget, complete bool) {
	var arg0 *C.GtkAssistant
	var arg1 *C.GtkWidget
	var arg2 C.gboolean

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	if complete {
		arg2 = C.gboolean(1)
	}

	C.gtk_assistant_set_page_complete(arg0, arg1, arg2)
}

// SetPageHasPadding sets whether the assistant is adding padding around the
// page.
func (a assistant) SetPageHasPadding(a Assistant, page Widget, hasPadding bool) {
	var arg0 *C.GtkAssistant
	var arg1 *C.GtkWidget
	var arg2 C.gboolean

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	if hasPadding {
		arg2 = C.gboolean(1)
	}

	C.gtk_assistant_set_page_has_padding(arg0, arg1, arg2)
}

// SetPageHeaderImage sets a header image for @page.
func (a assistant) SetPageHeaderImage(a Assistant, page Widget, pixbuf gdkpixbuf.Pixbuf) {
	var arg0 *C.GtkAssistant
	var arg1 *C.GtkWidget
	var arg2 *C.GdkPixbuf

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	arg2 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gtk_assistant_set_page_header_image(arg0, arg1, arg2)
}

// SetPageSideImage sets a side image for @page.
//
// This image used to be displayed in the side area of the assistant when
// @page is the current page.
func (a assistant) SetPageSideImage(a Assistant, page Widget, pixbuf gdkpixbuf.Pixbuf) {
	var arg0 *C.GtkAssistant
	var arg1 *C.GtkWidget
	var arg2 *C.GdkPixbuf

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	arg2 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gtk_assistant_set_page_side_image(arg0, arg1, arg2)
}

// SetPageTitle sets a title for @page.
//
// The title is displayed in the header area of the assistant when @page is
// the current page.
func (a assistant) SetPageTitle(a Assistant, page Widget, title string) {
	var arg0 *C.GtkAssistant
	var arg1 *C.GtkWidget
	var arg2 *C.gchar

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	arg2 = (*C.gchar)(C.CString(title))
	defer C.free(unsafe.Pointer(arg2))

	C.gtk_assistant_set_page_title(arg0, arg1, arg2)
}

// SetPageType sets the page type for @page.
//
// The page type determines the page behavior in the @assistant.
func (a assistant) SetPageType(a Assistant, page Widget, typ AssistantPageType) {
	var arg0 *C.GtkAssistant
	var arg1 *C.GtkWidget
	var arg2 C.GtkAssistantPageType

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))
	arg1 = (*C.GtkWidget)(unsafe.Pointer(page.Native()))
	arg2 = (C.GtkAssistantPageType)(typ)

	C.gtk_assistant_set_page_type(arg0, arg1, arg2)
}

// UpdateButtonsState forces @assistant to recompute the buttons state.
//
// GTK+ automatically takes care of this in most situations, e.g. when the
// user goes to a different page, or when the visibility or completeness of
// a page changes.
//
// One situation where it can be necessary to call this function is when
// changing a value on the current page affects the future page flow of the
// assistant.
func (a assistant) UpdateButtonsState(a Assistant) {
	var arg0 *C.GtkAssistant

	arg0 = (*C.GtkAssistant)(unsafe.Pointer(a.Native()))

	C.gtk_assistant_update_buttons_state(arg0)
}
