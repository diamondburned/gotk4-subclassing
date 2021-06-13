// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/box"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4 glib-2.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_area_get_type()), F: marshalCellArea},
	})
}

// CellArea: an abstract class for laying out GtkCellRenderers
//
// The CellArea is an abstract class for CellLayout widgets (also referred to as
// "layouting widgets") to interface with an arbitrary number of CellRenderers
// and interact with the user for a given TreeModel row.
//
// The cell area handles events, focus navigation, drawing and size requests and
// allocations for a given row of data.
//
// Usually users dont have to interact with the CellArea directly unless they
// are implementing a cell-layouting widget themselves.
//
//
// Requesting area sizes
//
// As outlined in [GtkWidget’s geometry management
// section][geometry-management], GTK uses a height-for-width geometry
// management system to compute the sizes of widgets and user interfaces.
// CellArea uses the same semantics to calculate the size of an area for an
// arbitrary number of TreeModel rows.
//
// When requesting the size of a cell area one needs to calculate the size for a
// handful of rows, and this will be done differently by different layouting
// widgets. For instance a TreeViewColumn always lines up the areas from top to
// bottom while a IconView on the other hand might enforce that all areas
// received the same width and wrap the areas around, requesting height for more
// cell areas when allocated less width.
//
// It’s also important for areas to maintain some cell alignments with areas
// rendered for adjacent rows (cells can appear “columnized” inside an area even
// when the size of cells are different in each row). For this reason the
// CellArea uses a CellAreaContext object to store the alignments and sizes
// along the way (as well as the overall largest minimum and natural size for
// all the rows which have been calculated with the said context).
//
// The CellAreaContext is an opaque object specific to the CellArea which
// created it (see gtk_cell_area_create_context()). The owning cell-layouting
// widget can create as many contexts as it wishes to calculate sizes of rows
// which should receive the same size in at least one orientation (horizontally
// or vertically), However, it’s important that the same CellAreaContext which
// was used to request the sizes for a given TreeModel row be used when
// rendering or processing events for that row.
//
// In order to request the width of all the rows at the root level of a
// TreeModel one would do the following:
//
//    static gboolean
//    foo_focus (GtkWidget       *widget,
//               GtkDirectionType direction)
//    {
//      Foo        *foo  = FOO (widget);
//      FooPrivate *priv = foo->priv;
//      int         focus_row;
//      gboolean    have_focus = FALSE;
//
//      focus_row = priv->focus_row;
//
//      if (!gtk_widget_has_focus (widget))
//        gtk_widget_grab_focus (widget);
//
//      valid = gtk_tree_model_iter_nth_child (priv->model, &iter, NULL, priv->focus_row);
//      while (valid)
//        {
//          gtk_cell_area_apply_attributes (priv->area, priv->model, &iter, FALSE, FALSE);
//
//          if (gtk_cell_area_focus (priv->area, direction))
//            {
//               priv->focus_row = focus_row;
//               have_focus = TRUE;
//               break;
//            }
//          else
//            {
//              if (direction == GTK_DIR_RIGHT ||
//                  direction == GTK_DIR_LEFT)
//                break;
//              else if (direction == GTK_DIR_UP ||
//                       direction == GTK_DIR_TAB_BACKWARD)
//               {
//                  if (focus_row == 0)
//                    break;
//                  else
//                   {
//                      focus_row--;
//                      valid = gtk_tree_model_iter_nth_child (priv->model, &iter, NULL, focus_row);
//                   }
//                }
//              else
//                {
//                  if (focus_row == last_row)
//                    break;
//                  else
//                    {
//                      focus_row++;
//                      valid = gtk_tree_model_iter_next (priv->model, &iter);
//                    }
//                }
//            }
//        }
//        return have_focus;
//    }
//
// Note that the layouting widget is responsible for matching the
// GtkDirectionType values to the way it lays out its cells.
//
//
// Cell Properties
//
// The CellArea introduces cell properties for CellRenderers. This provides some
// general interfaces for defining the relationship cell areas have with their
// cells. For instance in a CellAreaBox a cell might “expand” and receive extra
// space when the area is allocated more than its full natural request, or a
// cell might be configured to “align” with adjacent rows which were requested
// and rendered with the same CellAreaContext.
//
// Use gtk_cell_area_class_install_cell_property() to install cell properties
// for a cell area class and gtk_cell_area_class_find_cell_property() or
// gtk_cell_area_class_list_cell_properties() to get information about existing
// cell properties.
//
// To set the value of a cell property, use gtk_cell_area_cell_set_property(),
// gtk_cell_area_cell_set() or gtk_cell_area_cell_set_valist(). To obtain the
// value of a cell property, use gtk_cell_area_cell_get_property(),
// gtk_cell_area_cell_get() or gtk_cell_area_cell_get_valist().
type CellArea interface {
	gextras.Objector
	Buildable
	CellLayout

	// Activate activates @area, usually by activating the currently focused
	// cell, however some subclasses which embed widgets in the area can also
	// activate a widget if it currently has the focus.
	Activate(context CellAreaContext, widget Widget, cellArea *gdk.Rectangle, flags CellRendererState, editOnly bool) bool
	// ActivateCell: this is used by CellArea subclasses when handling events to
	// activate cells, the base CellArea class activates cells for keyboard
	// events for free in its own GtkCellArea->activate() implementation.
	ActivateCell(widget Widget, renderer CellRenderer, event gdk.Event, cellArea *gdk.Rectangle, flags CellRendererState) bool
	// Add adds @renderer to @area with the default child cell properties.
	Add(renderer CellRenderer)
	// AddFocusSibling adds @sibling to @renderer’s focusable area, focus will
	// be drawn around @renderer and all of its siblings if @renderer can focus
	// for a given row.
	//
	// Events handled by focus siblings can also activate the given focusable
	// @renderer.
	AddFocusSibling(renderer CellRenderer, sibling CellRenderer)
	// ApplyAttributes applies any connected attributes to the renderers in
	// @area by pulling the values from @tree_model.
	ApplyAttributes(treeModel TreeModel, iter *TreeIter, isExpander bool, isExpanded bool)
	// AttributeConnect connects an @attribute to apply values from @column for
	// the TreeModel in use.
	AttributeConnect(renderer CellRenderer, attribute string, column int)
	// AttributeDisconnect disconnects @attribute for the @renderer in @area so
	// that attribute will no longer be updated with values from the model.
	AttributeDisconnect(renderer CellRenderer, attribute string)
	// AttributeGetColumn returns the model column that an attribute has been
	// mapped to, or -1 if the attribute is not mapped.
	AttributeGetColumn(renderer CellRenderer, attribute string) int
	// CellGetProperty gets the value of a cell property for @renderer in @area.
	CellGetProperty(renderer CellRenderer, propertyName string, value **externglib.Value)
	// CellSetProperty sets a cell property for @renderer in @area.
	CellSetProperty(renderer CellRenderer, propertyName string, value **externglib.Value)
	// Event delegates event handling to a CellArea.
	Event(context CellAreaContext, widget Widget, event gdk.Event, cellArea *gdk.Rectangle, flags CellRendererState) int
	// Focus: this should be called by the @area’s owning layout widget when
	// focus is to be passed to @area, or moved within @area for a given
	// @direction and row data.
	//
	// Implementing CellArea classes should implement this method to receive and
	// navigate focus in its own way particular to how it lays out cells.
	Focus(direction DirectionType) bool
	// Foreach calls @callback for every CellRenderer in @area.
	Foreach(callback CellCallback)
	// ForeachAlloc calls @callback for every CellRenderer in @area with the
	// allocated rectangle inside @cell_area.
	ForeachAlloc(context CellAreaContext, widget Widget, cellArea *gdk.Rectangle, backgroundArea *gdk.Rectangle, callback CellAllocCallback)
	// CellAllocation derives the allocation of @renderer inside @area if @area
	// were to be renderered in @cell_area.
	CellAllocation(context CellAreaContext, widget Widget, renderer CellRenderer, cellArea *gdk.Rectangle) gdk.Rectangle
	// CurrentPathString gets the current TreePath string for the currently
	// applied TreeIter, this is implicitly updated when
	// gtk_cell_area_apply_attributes() is called and can be used to interact
	// with renderers from CellArea subclasses.
	CurrentPathString() string
	// PreferredHeight retrieves a cell area’s initial minimum and natural
	// height.
	//
	// @area will store some geometrical information in @context along the way;
	// when requesting sizes over an arbitrary number of rows, it’s not
	// important to check the @minimum_height and @natural_height of this call
	// but rather to consult gtk_cell_area_context_get_preferred_height() after
	// a series of requests.
	PreferredHeight(context CellAreaContext, widget Widget) (minimumHeight int, naturalHeight int)
	// PreferredHeightForWidth retrieves a cell area’s minimum and natural
	// height if it would be given the specified @width.
	//
	// @area stores some geometrical information in @context along the way while
	// calling gtk_cell_area_get_preferred_width(). It’s important to perform a
	// series of gtk_cell_area_get_preferred_width() requests with @context
	// first and then call gtk_cell_area_get_preferred_height_for_width() on
	// each cell area individually to get the height for width of each fully
	// requested row.
	//
	// If at some point, the width of a single row changes, it should be
	// requested with gtk_cell_area_get_preferred_width() again and then the
	// full width of the requested rows checked again with
	// gtk_cell_area_context_get_preferred_width().
	PreferredHeightForWidth(context CellAreaContext, widget Widget, width int) (minimumHeight int, naturalHeight int)
	// PreferredWidth retrieves a cell area’s initial minimum and natural width.
	//
	// @area will store some geometrical information in @context along the way;
	// when requesting sizes over an arbitrary number of rows, it’s not
	// important to check the @minimum_width and @natural_width of this call but
	// rather to consult gtk_cell_area_context_get_preferred_width() after a
	// series of requests.
	PreferredWidth(context CellAreaContext, widget Widget) (minimumWidth int, naturalWidth int)
	// PreferredWidthForHeight retrieves a cell area’s minimum and natural width
	// if it would be given the specified @height.
	//
	// @area stores some geometrical information in @context along the way while
	// calling gtk_cell_area_get_preferred_height(). It’s important to perform a
	// series of gtk_cell_area_get_preferred_height() requests with @context
	// first and then call gtk_cell_area_get_preferred_width_for_height() on
	// each cell area individually to get the height for width of each fully
	// requested row.
	//
	// If at some point, the height of a single row changes, it should be
	// requested with gtk_cell_area_get_preferred_height() again and then the
	// full height of the requested rows checked again with
	// gtk_cell_area_context_get_preferred_height().
	PreferredWidthForHeight(context CellAreaContext, widget Widget, height int) (minimumWidth int, naturalWidth int)
	// HasRenderer checks if @area contains @renderer.
	HasRenderer(renderer CellRenderer) bool
	// InnerCellArea: this is a convenience function for CellArea
	// implementations to get the inner area where a given CellRenderer will be
	// rendered. It removes any padding previously added by
	// gtk_cell_area_request_renderer().
	InnerCellArea(widget Widget, cellArea *gdk.Rectangle) gdk.Rectangle
	// IsActivatable returns whether the area can do anything when activated,
	// after applying new attributes to @area.
	IsActivatable() bool
	// IsFocusSibling returns whether @sibling is one of @renderer’s focus
	// siblings (see gtk_cell_area_add_focus_sibling()).
	IsFocusSibling(renderer CellRenderer, sibling CellRenderer) bool
	// Remove removes @renderer from @area.
	Remove(renderer CellRenderer)
	// RemoveFocusSibling removes @sibling from @renderer’s focus sibling list
	// (see gtk_cell_area_add_focus_sibling()).
	RemoveFocusSibling(renderer CellRenderer, sibling CellRenderer)
	// RequestRenderer: this is a convenience function for CellArea
	// implementations to request size for cell renderers. It’s important to use
	// this function to request size and then use
	// gtk_cell_area_inner_cell_area() at render and event time since this
	// function will add padding around the cell for focus painting.
	RequestRenderer(renderer CellRenderer, orientation Orientation, widget Widget, forSize int) (minimumSize int, naturalSize int)
	// SetFocusCell: explicitly sets the currently focused cell to @renderer.
	//
	// This is generally called by implementations of CellAreaClass.focus() or
	// CellAreaClass.event(), however it can also be used to implement functions
	// such as gtk_tree_view_set_cursor_on_cell().
	SetFocusCell(renderer CellRenderer)
	// Snapshot snapshots @area’s cells according to @area’s layout onto at the
	// given coordinates.
	Snapshot(context CellAreaContext, widget Widget, snapshot Snapshot, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState, paintFocus bool)
	// StopEditing: explicitly stops the editing of the currently edited cell.
	//
	// If @canceled is true, the currently edited cell renderer will emit the
	// ::editing-canceled signal, otherwise the the ::editing-done signal will
	// be emitted on the current edit widget.
	//
	// See gtk_cell_area_get_edited_cell() and gtk_cell_area_get_edit_widget().
	StopEditing(canceled bool)
}

// cellArea implements the CellArea interface.
type cellArea struct {
	gextras.Objector
	Buildable
	CellLayout
}

var _ CellArea = (*cellArea)(nil)

// WrapCellArea wraps a GObject to the right type. It is
// primarily used internally.
func WrapCellArea(obj *externglib.Object) CellArea {
	return CellArea{
		Objector:   obj,
		Buildable:  WrapBuildable(obj),
		CellLayout: WrapCellLayout(obj),
	}
}

func marshalCellArea(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapCellArea(obj), nil
}

// Activate activates @area, usually by activating the currently focused
// cell, however some subclasses which embed widgets in the area can also
// activate a widget if it currently has the focus.
func (a cellArea) Activate(context CellAreaContext, widget Widget, cellArea *gdk.Rectangle, flags CellRendererState, editOnly bool) bool {
	var _arg0 *C.GtkCellArea         // out
	var _arg1 *C.GtkCellAreaContext  // out
	var _arg2 *C.GtkWidget           // out
	var _arg3 *C.GdkRectangle        // out
	var _arg4 C.GtkCellRendererState // out
	var _arg5 C.gboolean             // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = (*C.GdkRectangle)(unsafe.Pointer(cellArea.Native()))
	_arg4 = (C.GtkCellRendererState)(flags)
	if editOnly {
		_arg5 = C.gboolean(1)
	}

	var _cret C.gboolean // in

	_cret = C.gtk_cell_area_activate(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// ActivateCell: this is used by CellArea subclasses when handling events to
// activate cells, the base CellArea class activates cells for keyboard
// events for free in its own GtkCellArea->activate() implementation.
func (a cellArea) ActivateCell(widget Widget, renderer CellRenderer, event gdk.Event, cellArea *gdk.Rectangle, flags CellRendererState) bool {
	var _arg0 *C.GtkCellArea         // out
	var _arg1 *C.GtkWidget           // out
	var _arg2 *C.GtkCellRenderer     // out
	var _arg3 *C.GdkEvent            // out
	var _arg4 *C.GdkRectangle        // out
	var _arg5 C.GtkCellRendererState // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg3 = (*C.GdkEvent)(unsafe.Pointer(event.Native()))
	_arg4 = (*C.GdkRectangle)(unsafe.Pointer(cellArea.Native()))
	_arg5 = (C.GtkCellRendererState)(flags)

	var _cret C.gboolean // in

	_cret = C.gtk_cell_area_activate_cell(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Add adds @renderer to @area with the default child cell properties.
func (a cellArea) Add(renderer CellRenderer) {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))

	C.gtk_cell_area_add(_arg0, _arg1)
}

// AddFocusSibling adds @sibling to @renderer’s focusable area, focus will
// be drawn around @renderer and all of its siblings if @renderer can focus
// for a given row.
//
// Events handled by focus siblings can also activate the given focusable
// @renderer.
func (a cellArea) AddFocusSibling(renderer CellRenderer, sibling CellRenderer) {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 *C.GtkCellRenderer // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg2 = (*C.GtkCellRenderer)(unsafe.Pointer(sibling.Native()))

	C.gtk_cell_area_add_focus_sibling(_arg0, _arg1, _arg2)
}

// ApplyAttributes applies any connected attributes to the renderers in
// @area by pulling the values from @tree_model.
func (a cellArea) ApplyAttributes(treeModel TreeModel, iter *TreeIter, isExpander bool, isExpanded bool) {
	var _arg0 *C.GtkCellArea  // out
	var _arg1 *C.GtkTreeModel // out
	var _arg2 *C.GtkTreeIter  // out
	var _arg3 C.gboolean      // out
	var _arg4 C.gboolean      // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkTreeModel)(unsafe.Pointer(treeModel.Native()))
	_arg2 = (*C.GtkTreeIter)(unsafe.Pointer(iter.Native()))
	if isExpander {
		_arg3 = C.gboolean(1)
	}
	if isExpanded {
		_arg4 = C.gboolean(1)
	}

	C.gtk_cell_area_apply_attributes(_arg0, _arg1, _arg2, _arg3, _arg4)
}

// AttributeConnect connects an @attribute to apply values from @column for
// the TreeModel in use.
func (a cellArea) AttributeConnect(renderer CellRenderer, attribute string, column int) {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 *C.char            // out
	var _arg3 C.int              // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg2 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.int(column)

	C.gtk_cell_area_attribute_connect(_arg0, _arg1, _arg2, _arg3)
}

// AttributeDisconnect disconnects @attribute for the @renderer in @area so
// that attribute will no longer be updated with values from the model.
func (a cellArea) AttributeDisconnect(renderer CellRenderer, attribute string) {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 *C.char            // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg2 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_cell_area_attribute_disconnect(_arg0, _arg1, _arg2)
}

// AttributeGetColumn returns the model column that an attribute has been
// mapped to, or -1 if the attribute is not mapped.
func (a cellArea) AttributeGetColumn(renderer CellRenderer, attribute string) int {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 *C.char            // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg2 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg2))

	var _cret C.int // in

	_cret = C.gtk_cell_area_attribute_get_column(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// CellGetProperty gets the value of a cell property for @renderer in @area.
func (a cellArea) CellGetProperty(renderer CellRenderer, propertyName string, value **externglib.Value) {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 *C.char            // out
	var _arg3 *C.GValue          // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg2 = (*C.char)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GValue)(value.GValue)

	C.gtk_cell_area_cell_get_property(_arg0, _arg1, _arg2, _arg3)
}

// CellSetProperty sets a cell property for @renderer in @area.
func (a cellArea) CellSetProperty(renderer CellRenderer, propertyName string, value **externglib.Value) {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 *C.char            // out
	var _arg3 *C.GValue          // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg2 = (*C.char)(C.CString(propertyName))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GValue)(value.GValue)

	C.gtk_cell_area_cell_set_property(_arg0, _arg1, _arg2, _arg3)
}

// Event delegates event handling to a CellArea.
func (a cellArea) Event(context CellAreaContext, widget Widget, event gdk.Event, cellArea *gdk.Rectangle, flags CellRendererState) int {
	var _arg0 *C.GtkCellArea         // out
	var _arg1 *C.GtkCellAreaContext  // out
	var _arg2 *C.GtkWidget           // out
	var _arg3 *C.GdkEvent            // out
	var _arg4 *C.GdkRectangle        // out
	var _arg5 C.GtkCellRendererState // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = (*C.GdkEvent)(unsafe.Pointer(event.Native()))
	_arg4 = (*C.GdkRectangle)(unsafe.Pointer(cellArea.Native()))
	_arg5 = (C.GtkCellRendererState)(flags)

	var _cret C.int // in

	_cret = C.gtk_cell_area_event(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5)

	var _gint int // out

	_gint = (int)(_cret)

	return _gint
}

// Focus: this should be called by the @area’s owning layout widget when
// focus is to be passed to @area, or moved within @area for a given
// @direction and row data.
//
// Implementing CellArea classes should implement this method to receive and
// navigate focus in its own way particular to how it lays out cells.
func (a cellArea) Focus(direction DirectionType) bool {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 C.GtkDirectionType // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (C.GtkDirectionType)(direction)

	var _cret C.gboolean // in

	_cret = C.gtk_cell_area_focus(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Foreach calls @callback for every CellRenderer in @area.
func (a cellArea) Foreach(callback CellCallback) {
	var _arg0 *C.GtkCellArea    // out
	var _arg1 C.GtkCellCallback // out
	var _arg2 C.gpointer

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*[0]byte)(C.gotk4_CellCallback)
	_arg2 = C.gpointer(box.Assign(callback))

	C.gtk_cell_area_foreach(_arg0, _arg1, _arg2)
}

// ForeachAlloc calls @callback for every CellRenderer in @area with the
// allocated rectangle inside @cell_area.
func (a cellArea) ForeachAlloc(context CellAreaContext, widget Widget, cellArea *gdk.Rectangle, backgroundArea *gdk.Rectangle, callback CellAllocCallback) {
	var _arg0 *C.GtkCellArea         // out
	var _arg1 *C.GtkCellAreaContext  // out
	var _arg2 *C.GtkWidget           // out
	var _arg3 *C.GdkRectangle        // out
	var _arg4 *C.GdkRectangle        // out
	var _arg5 C.GtkCellAllocCallback // out
	var _arg6 C.gpointer

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = (*C.GdkRectangle)(unsafe.Pointer(cellArea.Native()))
	_arg4 = (*C.GdkRectangle)(unsafe.Pointer(backgroundArea.Native()))
	_arg5 = (*[0]byte)(C.gotk4_CellAllocCallback)
	_arg6 = C.gpointer(box.Assign(callback))

	C.gtk_cell_area_foreach_alloc(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
}

// CellAllocation derives the allocation of @renderer inside @area if @area
// were to be renderered in @cell_area.
func (a cellArea) CellAllocation(context CellAreaContext, widget Widget, renderer CellRenderer, cellArea *gdk.Rectangle) gdk.Rectangle {
	var _arg0 *C.GtkCellArea        // out
	var _arg1 *C.GtkCellAreaContext // out
	var _arg2 *C.GtkWidget          // out
	var _arg3 *C.GtkCellRenderer    // out
	var _arg4 *C.GdkRectangle       // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg4 = (*C.GdkRectangle)(unsafe.Pointer(cellArea.Native()))

	var _allocation gdk.Rectangle

	C.gtk_cell_area_get_cell_allocation(_arg0, _arg1, _arg2, _arg3, _arg4, (*C.GdkRectangle)(unsafe.Pointer(&_allocation)))

	return _allocation
}

// CurrentPathString gets the current TreePath string for the currently
// applied TreeIter, this is implicitly updated when
// gtk_cell_area_apply_attributes() is called and can be used to interact
// with renderers from CellArea subclasses.
func (a cellArea) CurrentPathString() string {
	var _arg0 *C.GtkCellArea // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))

	var _cret *C.char // in

	_cret = C.gtk_cell_area_get_current_path_string(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// PreferredHeight retrieves a cell area’s initial minimum and natural
// height.
//
// @area will store some geometrical information in @context along the way;
// when requesting sizes over an arbitrary number of rows, it’s not
// important to check the @minimum_height and @natural_height of this call
// but rather to consult gtk_cell_area_context_get_preferred_height() after
// a series of requests.
func (a cellArea) PreferredHeight(context CellAreaContext, widget Widget) (minimumHeight int, naturalHeight int) {
	var _arg0 *C.GtkCellArea        // out
	var _arg1 *C.GtkCellAreaContext // out
	var _arg2 *C.GtkWidget          // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	var _arg3 C.int // in
	var _arg4 C.int // in

	C.gtk_cell_area_get_preferred_height(_arg0, _arg1, _arg2, &_arg3, &_arg4)

	var _minimumHeight int // out
	var _naturalHeight int // out

	_minimumHeight = (int)(_arg3)
	_naturalHeight = (int)(_arg4)

	return _minimumHeight, _naturalHeight
}

// PreferredHeightForWidth retrieves a cell area’s minimum and natural
// height if it would be given the specified @width.
//
// @area stores some geometrical information in @context along the way while
// calling gtk_cell_area_get_preferred_width(). It’s important to perform a
// series of gtk_cell_area_get_preferred_width() requests with @context
// first and then call gtk_cell_area_get_preferred_height_for_width() on
// each cell area individually to get the height for width of each fully
// requested row.
//
// If at some point, the width of a single row changes, it should be
// requested with gtk_cell_area_get_preferred_width() again and then the
// full width of the requested rows checked again with
// gtk_cell_area_context_get_preferred_width().
func (a cellArea) PreferredHeightForWidth(context CellAreaContext, widget Widget, width int) (minimumHeight int, naturalHeight int) {
	var _arg0 *C.GtkCellArea        // out
	var _arg1 *C.GtkCellAreaContext // out
	var _arg2 *C.GtkWidget          // out
	var _arg3 C.int                 // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = C.int(width)

	var _arg4 C.int // in
	var _arg5 C.int // in

	C.gtk_cell_area_get_preferred_height_for_width(_arg0, _arg1, _arg2, _arg3, &_arg4, &_arg5)

	var _minimumHeight int // out
	var _naturalHeight int // out

	_minimumHeight = (int)(_arg4)
	_naturalHeight = (int)(_arg5)

	return _minimumHeight, _naturalHeight
}

// PreferredWidth retrieves a cell area’s initial minimum and natural width.
//
// @area will store some geometrical information in @context along the way;
// when requesting sizes over an arbitrary number of rows, it’s not
// important to check the @minimum_width and @natural_width of this call but
// rather to consult gtk_cell_area_context_get_preferred_width() after a
// series of requests.
func (a cellArea) PreferredWidth(context CellAreaContext, widget Widget) (minimumWidth int, naturalWidth int) {
	var _arg0 *C.GtkCellArea        // out
	var _arg1 *C.GtkCellAreaContext // out
	var _arg2 *C.GtkWidget          // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	var _arg3 C.int // in
	var _arg4 C.int // in

	C.gtk_cell_area_get_preferred_width(_arg0, _arg1, _arg2, &_arg3, &_arg4)

	var _minimumWidth int // out
	var _naturalWidth int // out

	_minimumWidth = (int)(_arg3)
	_naturalWidth = (int)(_arg4)

	return _minimumWidth, _naturalWidth
}

// PreferredWidthForHeight retrieves a cell area’s minimum and natural width
// if it would be given the specified @height.
//
// @area stores some geometrical information in @context along the way while
// calling gtk_cell_area_get_preferred_height(). It’s important to perform a
// series of gtk_cell_area_get_preferred_height() requests with @context
// first and then call gtk_cell_area_get_preferred_width_for_height() on
// each cell area individually to get the height for width of each fully
// requested row.
//
// If at some point, the height of a single row changes, it should be
// requested with gtk_cell_area_get_preferred_height() again and then the
// full height of the requested rows checked again with
// gtk_cell_area_context_get_preferred_height().
func (a cellArea) PreferredWidthForHeight(context CellAreaContext, widget Widget, height int) (minimumWidth int, naturalWidth int) {
	var _arg0 *C.GtkCellArea        // out
	var _arg1 *C.GtkCellAreaContext // out
	var _arg2 *C.GtkWidget          // out
	var _arg3 C.int                 // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = C.int(height)

	var _arg4 C.int // in
	var _arg5 C.int // in

	C.gtk_cell_area_get_preferred_width_for_height(_arg0, _arg1, _arg2, _arg3, &_arg4, &_arg5)

	var _minimumWidth int // out
	var _naturalWidth int // out

	_minimumWidth = (int)(_arg4)
	_naturalWidth = (int)(_arg5)

	return _minimumWidth, _naturalWidth
}

// HasRenderer checks if @area contains @renderer.
func (a cellArea) HasRenderer(renderer CellRenderer) bool {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_cell_area_has_renderer(_arg0, _arg1)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// InnerCellArea: this is a convenience function for CellArea
// implementations to get the inner area where a given CellRenderer will be
// rendered. It removes any padding previously added by
// gtk_cell_area_request_renderer().
func (a cellArea) InnerCellArea(widget Widget, cellArea *gdk.Rectangle) gdk.Rectangle {
	var _arg0 *C.GtkCellArea  // out
	var _arg1 *C.GtkWidget    // out
	var _arg2 *C.GdkRectangle // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg2 = (*C.GdkRectangle)(unsafe.Pointer(cellArea.Native()))

	var _innerArea gdk.Rectangle

	C.gtk_cell_area_inner_cell_area(_arg0, _arg1, _arg2, (*C.GdkRectangle)(unsafe.Pointer(&_innerArea)))

	return _innerArea
}

// IsActivatable returns whether the area can do anything when activated,
// after applying new attributes to @area.
func (a cellArea) IsActivatable() bool {
	var _arg0 *C.GtkCellArea // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_cell_area_is_activatable(_arg0)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// IsFocusSibling returns whether @sibling is one of @renderer’s focus
// siblings (see gtk_cell_area_add_focus_sibling()).
func (a cellArea) IsFocusSibling(renderer CellRenderer, sibling CellRenderer) bool {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 *C.GtkCellRenderer // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg2 = (*C.GtkCellRenderer)(unsafe.Pointer(sibling.Native()))

	var _cret C.gboolean // in

	_cret = C.gtk_cell_area_is_focus_sibling(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret {
		_ok = true
	}

	return _ok
}

// Remove removes @renderer from @area.
func (a cellArea) Remove(renderer CellRenderer) {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))

	C.gtk_cell_area_remove(_arg0, _arg1)
}

// RemoveFocusSibling removes @sibling from @renderer’s focus sibling list
// (see gtk_cell_area_add_focus_sibling()).
func (a cellArea) RemoveFocusSibling(renderer CellRenderer, sibling CellRenderer) {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 *C.GtkCellRenderer // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg2 = (*C.GtkCellRenderer)(unsafe.Pointer(sibling.Native()))

	C.gtk_cell_area_remove_focus_sibling(_arg0, _arg1, _arg2)
}

// RequestRenderer: this is a convenience function for CellArea
// implementations to request size for cell renderers. It’s important to use
// this function to request size and then use
// gtk_cell_area_inner_cell_area() at render and event time since this
// function will add padding around the cell for focus painting.
func (a cellArea) RequestRenderer(renderer CellRenderer, orientation Orientation, widget Widget, forSize int) (minimumSize int, naturalSize int) {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out
	var _arg2 C.GtkOrientation   // out
	var _arg3 *C.GtkWidget       // out
	var _arg4 C.int              // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))
	_arg2 = (C.GtkOrientation)(orientation)
	_arg3 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg4 = C.int(forSize)

	var _arg5 C.int // in
	var _arg6 C.int // in

	C.gtk_cell_area_request_renderer(_arg0, _arg1, _arg2, _arg3, _arg4, &_arg5, &_arg6)

	var _minimumSize int // out
	var _naturalSize int // out

	_minimumSize = (int)(_arg5)
	_naturalSize = (int)(_arg6)

	return _minimumSize, _naturalSize
}

// SetFocusCell: explicitly sets the currently focused cell to @renderer.
//
// This is generally called by implementations of CellAreaClass.focus() or
// CellAreaClass.event(), however it can also be used to implement functions
// such as gtk_tree_view_set_cursor_on_cell().
func (a cellArea) SetFocusCell(renderer CellRenderer) {
	var _arg0 *C.GtkCellArea     // out
	var _arg1 *C.GtkCellRenderer // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellRenderer)(unsafe.Pointer(renderer.Native()))

	C.gtk_cell_area_set_focus_cell(_arg0, _arg1)
}

// Snapshot snapshots @area’s cells according to @area’s layout onto at the
// given coordinates.
func (a cellArea) Snapshot(context CellAreaContext, widget Widget, snapshot Snapshot, backgroundArea *gdk.Rectangle, cellArea *gdk.Rectangle, flags CellRendererState, paintFocus bool) {
	var _arg0 *C.GtkCellArea         // out
	var _arg1 *C.GtkCellAreaContext  // out
	var _arg2 *C.GtkWidget           // out
	var _arg3 *C.GtkSnapshot         // out
	var _arg4 *C.GdkRectangle        // out
	var _arg5 *C.GdkRectangle        // out
	var _arg6 C.GtkCellRendererState // out
	var _arg7 C.gboolean             // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	_arg1 = (*C.GtkCellAreaContext)(unsafe.Pointer(context.Native()))
	_arg2 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	_arg3 = (*C.GtkSnapshot)(unsafe.Pointer(snapshot.Native()))
	_arg4 = (*C.GdkRectangle)(unsafe.Pointer(backgroundArea.Native()))
	_arg5 = (*C.GdkRectangle)(unsafe.Pointer(cellArea.Native()))
	_arg6 = (C.GtkCellRendererState)(flags)
	if paintFocus {
		_arg7 = C.gboolean(1)
	}

	C.gtk_cell_area_snapshot(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
}

// StopEditing: explicitly stops the editing of the currently edited cell.
//
// If @canceled is true, the currently edited cell renderer will emit the
// ::editing-canceled signal, otherwise the the ::editing-done signal will
// be emitted on the current edit widget.
//
// See gtk_cell_area_get_edited_cell() and gtk_cell_area_get_edit_widget().
func (a cellArea) StopEditing(canceled bool) {
	var _arg0 *C.GtkCellArea // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GtkCellArea)(unsafe.Pointer(a.Native()))
	if canceled {
		_arg1 = C.gboolean(1)
	}

	C.gtk_cell_area_stop_editing(_arg0, _arg1)
}
