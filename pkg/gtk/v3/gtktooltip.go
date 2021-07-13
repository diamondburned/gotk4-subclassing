// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_tooltip_get_type()), F: marshalTooltiper},
	})
}

// Tooltiper describes Tooltip's methods.
type Tooltiper interface {
	// SetCustom replaces the widget packed into the tooltip with custom_widget.
	SetCustom(customWidget Widgeter)
	// SetIcon sets the icon of the tooltip (which is in front of the text) to
	// be pixbuf.
	SetIcon(pixbuf *gdkpixbuf.Pixbuf)
	// SetIconFromGIcon sets the icon of the tooltip (which is in front of the
	// text) to be the icon indicated by gicon with the size indicated by size.
	SetIconFromGIcon(gicon gio.Iconer, size int)
	// SetIconFromIconName sets the icon of the tooltip (which is in front of
	// the text) to be the icon indicated by icon_name with the size indicated
	// by size.
	SetIconFromIconName(iconName string, size int)
	// SetIconFromStock sets the icon of the tooltip (which is in front of the
	// text) to be the stock item indicated by stock_id with the size indicated
	// by size.
	SetIconFromStock(stockId string, size int)
	// SetMarkup sets the text of the tooltip to be markup, which is marked up
	// with the [Pango text markup language][PangoMarkupFormat].
	SetMarkup(markup string)
	// SetText sets the text of the tooltip to be text.
	SetText(text string)
	// SetTipArea sets the area of the widget, where the contents of this
	// tooltip apply, to be rect (in widget coordinates).
	SetTipArea(rect *gdk.Rectangle)
}

// Tooltip: basic tooltips can be realized simply by using
// gtk_widget_set_tooltip_text() or gtk_widget_set_tooltip_markup() without any
// explicit tooltip object.
//
// When you need a tooltip with a little more fancy contents, like adding an
// image, or you want the tooltip to have different contents per TreeView row or
// cell, you will have to do a little more work:
//
// - Set the Widget:has-tooltip property to TRUE, this will make GTK+ monitor
// the widget for motion and related events which are needed to determine when
// and where to show a tooltip.
//
// - Connect to the Widget::query-tooltip signal. This signal will be emitted
// when a tooltip is supposed to be shown. One of the arguments passed to the
// signal handler is a GtkTooltip object. This is the object that we are about
// to display as a tooltip, and can be manipulated in your callback using
// functions like gtk_tooltip_set_icon(). There are functions for setting the
// tooltip’s markup, setting an image from a named icon, or even putting in a
// custom widget.
//
//    Return TRUE from your query-tooltip handler. This causes the tooltip to be
//    show. If you return FALSE, it will not be shown.
//
// In the probably rare case where you want to have even more control over the
// tooltip that is about to be shown, you can set your own Window which will be
// used as tooltip window. This works as follows:
//
// - Set Widget:has-tooltip and connect to Widget::query-tooltip as before. Use
// gtk_widget_set_tooltip_window() to set a Window created by you as tooltip
// window.
//
// - In the Widget::query-tooltip callback you can access your window using
// gtk_widget_get_tooltip_window() and manipulate as you wish. The semantics of
// the return value are exactly as before, return TRUE to show the window, FALSE
// to not show it.
type Tooltip struct {
	*externglib.Object
}

var (
	_ Tooltiper       = (*Tooltip)(nil)
	_ gextras.Nativer = (*Tooltip)(nil)
)

func wrapTooltip(obj *externglib.Object) *Tooltip {
	return &Tooltip{
		Object: obj,
	}
}

func marshalTooltiper(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapTooltip(obj), nil
}

// SetCustom replaces the widget packed into the tooltip with custom_widget.
// custom_widget does not get destroyed when the tooltip goes away. By default a
// box with a Image and Label is embedded in the tooltip, which can be
// configured using gtk_tooltip_set_markup() and gtk_tooltip_set_icon().
func (tooltip *Tooltip) SetCustom(customWidget Widgeter) {
	var _arg0 *C.GtkTooltip // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkTooltip)(unsafe.Pointer(tooltip.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((customWidget).(gextras.Nativer).Native()))

	C.gtk_tooltip_set_custom(_arg0, _arg1)
}

// SetIcon sets the icon of the tooltip (which is in front of the text) to be
// pixbuf. If pixbuf is NULL, the image will be hidden.
func (tooltip *Tooltip) SetIcon(pixbuf *gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkTooltip // out
	var _arg1 *C.GdkPixbuf  // out

	_arg0 = (*C.GtkTooltip)(unsafe.Pointer(tooltip.Native()))
	_arg1 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gtk_tooltip_set_icon(_arg0, _arg1)
}

// SetIconFromGIcon sets the icon of the tooltip (which is in front of the text)
// to be the icon indicated by gicon with the size indicated by size. If gicon
// is NULL, the image will be hidden.
func (tooltip *Tooltip) SetIconFromGIcon(gicon gio.Iconer, size int) {
	var _arg0 *C.GtkTooltip // out
	var _arg1 *C.GIcon      // out
	var _arg2 C.GtkIconSize // out

	_arg0 = (*C.GtkTooltip)(unsafe.Pointer(tooltip.Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer((gicon).(gextras.Nativer).Native()))
	_arg2 = C.GtkIconSize(size)

	C.gtk_tooltip_set_icon_from_gicon(_arg0, _arg1, _arg2)
}

// SetIconFromIconName sets the icon of the tooltip (which is in front of the
// text) to be the icon indicated by icon_name with the size indicated by size.
// If icon_name is NULL, the image will be hidden.
func (tooltip *Tooltip) SetIconFromIconName(iconName string, size int) {
	var _arg0 *C.GtkTooltip // out
	var _arg1 *C.gchar      // out
	var _arg2 C.GtkIconSize // out

	_arg0 = (*C.GtkTooltip)(unsafe.Pointer(tooltip.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(iconName)))
	_arg2 = C.GtkIconSize(size)

	C.gtk_tooltip_set_icon_from_icon_name(_arg0, _arg1, _arg2)
}

// SetIconFromStock sets the icon of the tooltip (which is in front of the text)
// to be the stock item indicated by stock_id with the size indicated by size.
// If stock_id is NULL, the image will be hidden.
//
// Deprecated: Use gtk_tooltip_set_icon_from_icon_name() instead.
func (tooltip *Tooltip) SetIconFromStock(stockId string, size int) {
	var _arg0 *C.GtkTooltip // out
	var _arg1 *C.gchar      // out
	var _arg2 C.GtkIconSize // out

	_arg0 = (*C.GtkTooltip)(unsafe.Pointer(tooltip.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(stockId)))
	_arg2 = C.GtkIconSize(size)

	C.gtk_tooltip_set_icon_from_stock(_arg0, _arg1, _arg2)
}

// SetMarkup sets the text of the tooltip to be markup, which is marked up with
// the [Pango text markup language][PangoMarkupFormat]. If markup is NULL, the
// label will be hidden.
func (tooltip *Tooltip) SetMarkup(markup string) {
	var _arg0 *C.GtkTooltip // out
	var _arg1 *C.gchar      // out

	_arg0 = (*C.GtkTooltip)(unsafe.Pointer(tooltip.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(markup)))

	C.gtk_tooltip_set_markup(_arg0, _arg1)
}

// SetText sets the text of the tooltip to be text. If text is NULL, the label
// will be hidden. See also gtk_tooltip_set_markup().
func (tooltip *Tooltip) SetText(text string) {
	var _arg0 *C.GtkTooltip // out
	var _arg1 *C.gchar      // out

	_arg0 = (*C.GtkTooltip)(unsafe.Pointer(tooltip.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(text)))

	C.gtk_tooltip_set_text(_arg0, _arg1)
}

// SetTipArea sets the area of the widget, where the contents of this tooltip
// apply, to be rect (in widget coordinates). This is especially useful for
// properly setting tooltips on TreeView rows and cells, IconViews, etc.
//
// For setting tooltips on TreeView, please refer to the convenience functions
// for this: gtk_tree_view_set_tooltip_row() and
// gtk_tree_view_set_tooltip_cell().
func (tooltip *Tooltip) SetTipArea(rect *gdk.Rectangle) {
	var _arg0 *C.GtkTooltip   // out
	var _arg1 *C.GdkRectangle // out

	_arg0 = (*C.GtkTooltip)(unsafe.Pointer(tooltip.Native()))
	_arg1 = (*C.GdkRectangle)(unsafe.Pointer(rect))

	C.gtk_tooltip_set_tip_area(_arg0, _arg1)
}

// TooltipTriggerTooltipQuery triggers a new tooltip query on display, in order
// to update the current visible tooltip, or to show/hide the current tooltip.
// This function is useful to call when, for example, the state of the widget
// changed by a key press.
func TooltipTriggerTooltipQuery(display *gdk.Display) {
	var _arg1 *C.GdkDisplay // out

	_arg1 = (*C.GdkDisplay)(unsafe.Pointer(display.Native()))

	C.gtk_tooltip_trigger_tooltip_query(_arg1)
}
