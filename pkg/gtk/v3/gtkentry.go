// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/core/gextras"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/pango"
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
		{T: externglib.Type(C.gtk_entry_icon_position_get_type()), F: marshalEntryIconPosition},
		{T: externglib.Type(C.gtk_entry_get_type()), F: marshalEntry},
	})
}

// EntryIconPosition specifies the side of the entry at which an icon is placed.
type EntryIconPosition int

const (
	// primary: at the beginning of the entry (depending on the text direction).
	EntryIconPositionPrimary EntryIconPosition = 0
	// secondary: at the end of the entry (depending on the text direction).
	EntryIconPositionSecondary EntryIconPosition = 1
)

func marshalEntryIconPosition(p uintptr) (interface{}, error) {
	return EntryIconPosition(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Entry: the Entry widget is a single line text entry widget. A fairly large
// set of key bindings are supported by default. If the entered text is longer
// than the allocation of the widget, the widget will scroll so that the cursor
// position is visible.
//
// When using an entry for passwords and other sensitive information, it can be
// put into “password mode” using gtk_entry_set_visibility(). In this mode,
// entered text is displayed using a “invisible” character. By default, GTK+
// picks the best invisible character that is available in the current font, but
// it can be changed with gtk_entry_set_invisible_char(). Since 2.16, GTK+
// displays a warning when Caps Lock or input methods might interfere with
// entering text in a password entry. The warning can be turned off with the
// Entry:caps-lock-warning property.
//
// Since 2.16, GtkEntry has the ability to display progress or activity
// information behind the text. To make an entry display such information, use
// gtk_entry_set_progress_fraction() or gtk_entry_set_progress_pulse_step().
//
// Additionally, GtkEntry can show icons at either side of the entry. These
// icons can be activatable by clicking, can be set up as drag source and can
// have tooltips. To add an icon, use gtk_entry_set_icon_from_gicon() or one of
// the various other functions that set an icon from a stock id, an icon name or
// a pixbuf. To trigger an action when the user clicks an icon, connect to the
// Entry::icon-press signal. To allow DND operations from an icon, use
// gtk_entry_set_icon_drag_source(). To set a tooltip on an icon, use
// gtk_entry_set_icon_tooltip_text() or the corresponding function for markup.
//
// Note that functionality or information that is only available by clicking on
// an icon in an entry may not be accessible at all to users which are not able
// to use a mouse or other pointing device. It is therefore recommended that any
// such functionality should also be available by other means, e.g. via the
// context menu of the entry.
//
// CSS nodes
//
//    entry[.read-only][.flat][.warning][.error]
//    ├── image.left
//    ├── image.right
//    ├── undershoot.left
//    ├── undershoot.right
//    ├── [selection]
//    ├── [progress[.pulse]]
//    ╰── [window.popup]
//
// GtkEntry has a main node with the name entry. Depending on the properties of
// the entry, the style classes .read-only and .flat may appear. The style
// classes .warning and .error may also be used with entries.
//
// When the entry shows icons, it adds subnodes with the name image and the
// style class .left or .right, depending on where the icon appears.
//
// When the entry has a selection, it adds a subnode with the name selection.
//
// When the entry shows progress, it adds a subnode with the name progress. The
// node has the style class .pulse when the shown progress is pulsing.
//
// The CSS node for a context menu is added as a subnode below entry as well.
//
// The undershoot nodes are used to draw the underflow indication when content
// is scrolled out of view. These nodes get the .left and .right style classes
// added depending on where the indication is drawn.
//
// When touch is used and touch selection handles are shown, they are using CSS
// nodes with name cursor-handle. They get the .top or .bottom style class
// depending on where they are shown in relation to the selection. If there is
// just a single handle for the text cursor, it gets the style class
// .insertion-cursor.
type Entry interface {
	CellEditable
	Editable

	// ActivatesDefault:
	ActivatesDefault() bool
	// Alignment:
	Alignment() float32
	// Attributes:
	Attributes() *pango.AttrList
	// Buffer:
	Buffer() EntryBuffer
	// Completion:
	Completion() EntryCompletion
	// CurrentIconDragSource:
	CurrentIconDragSource() int
	// CursorHAdjustment:
	CursorHAdjustment() Adjustment
	// HasFrame:
	HasFrame() bool
	// IconActivatable:
	IconActivatable(iconPos EntryIconPosition) bool
	// IconArea:
	IconArea(iconPos EntryIconPosition) gdk.Rectangle
	// IconAtPos:
	IconAtPos(x int, y int) int
	// IconGIcon:
	IconGIcon(iconPos EntryIconPosition) gio.Icon
	// IconName:
	IconName(iconPos EntryIconPosition) string
	// IconPixbuf:
	IconPixbuf(iconPos EntryIconPosition) gdkpixbuf.Pixbuf
	// IconSensitive:
	IconSensitive(iconPos EntryIconPosition) bool
	// IconStock:
	IconStock(iconPos EntryIconPosition) string
	// IconStorageType:
	IconStorageType(iconPos EntryIconPosition) ImageType
	// IconTooltipMarkup:
	IconTooltipMarkup(iconPos EntryIconPosition) string
	// IconTooltipText:
	IconTooltipText(iconPos EntryIconPosition) string
	// InnerBorder:
	InnerBorder() *Border
	// InputHints:
	InputHints() InputHints
	// InputPurpose:
	InputPurpose() InputPurpose
	// InvisibleChar:
	InvisibleChar() uint32
	// Layout:
	Layout() pango.Layout
	// LayoutOffsets:
	LayoutOffsets() (x int, y int)
	// MaxLength:
	MaxLength() int
	// MaxWidthChars:
	MaxWidthChars() int
	// OverwriteMode:
	OverwriteMode() bool
	// PlaceholderText:
	PlaceholderText() string
	// ProgressFraction:
	ProgressFraction() float64
	// ProgressPulseStep:
	ProgressPulseStep() float64
	// Tabs:
	Tabs() *pango.TabArray
	// Text:
	Text() string
	// TextArea:
	TextArea() gdk.Rectangle
	// TextLength:
	TextLength() uint16
	// Visibility:
	Visibility() bool
	// WidthChars:
	WidthChars() int
	// GrabFocusWithoutSelectingEntry:
	GrabFocusWithoutSelectingEntry()
	// ImContextFilterKeypressEntry:
	ImContextFilterKeypressEntry(event *gdk.EventKey) bool
	// LayoutIndexToTextIndexEntry:
	LayoutIndexToTextIndexEntry(layoutIndex int) int
	// ProgressPulseEntry:
	ProgressPulseEntry()
	// ResetImContextEntry:
	ResetImContextEntry()
	// SetActivatesDefaultEntry:
	SetActivatesDefaultEntry(setting bool)
	// SetAlignmentEntry:
	SetAlignmentEntry(xalign float32)
	// SetAttributesEntry:
	SetAttributesEntry(attrs *pango.AttrList)
	// SetBufferEntry:
	SetBufferEntry(buffer EntryBuffer)
	// SetCompletionEntry:
	SetCompletionEntry(completion EntryCompletion)
	// SetCursorHAdjustmentEntry:
	SetCursorHAdjustmentEntry(adjustment Adjustment)
	// SetHasFrameEntry:
	SetHasFrameEntry(setting bool)
	// SetIconActivatableEntry:
	SetIconActivatableEntry(iconPos EntryIconPosition, activatable bool)
	// SetIconDragSourceEntry:
	SetIconDragSourceEntry(iconPos EntryIconPosition, targetList *TargetList, actions gdk.DragAction)
	// SetIconFromGIconEntry:
	SetIconFromGIconEntry(iconPos EntryIconPosition, icon gio.Icon)
	// SetIconFromIconNameEntry:
	SetIconFromIconNameEntry(iconPos EntryIconPosition, iconName string)
	// SetIconFromPixbufEntry:
	SetIconFromPixbufEntry(iconPos EntryIconPosition, pixbuf gdkpixbuf.Pixbuf)
	// SetIconFromStockEntry:
	SetIconFromStockEntry(iconPos EntryIconPosition, stockId string)
	// SetIconSensitiveEntry:
	SetIconSensitiveEntry(iconPos EntryIconPosition, sensitive bool)
	// SetIconTooltipMarkupEntry:
	SetIconTooltipMarkupEntry(iconPos EntryIconPosition, tooltip string)
	// SetIconTooltipTextEntry:
	SetIconTooltipTextEntry(iconPos EntryIconPosition, tooltip string)
	// SetInnerBorderEntry:
	SetInnerBorderEntry(border *Border)
	// SetInputHintsEntry:
	SetInputHintsEntry(hints InputHints)
	// SetInputPurposeEntry:
	SetInputPurposeEntry(purpose InputPurpose)
	// SetInvisibleCharEntry:
	SetInvisibleCharEntry(ch uint32)
	// SetMaxLengthEntry:
	SetMaxLengthEntry(max int)
	// SetMaxWidthCharsEntry:
	SetMaxWidthCharsEntry(nChars int)
	// SetOverwriteModeEntry:
	SetOverwriteModeEntry(overwrite bool)
	// SetPlaceholderTextEntry:
	SetPlaceholderTextEntry(text string)
	// SetProgressFractionEntry:
	SetProgressFractionEntry(fraction float64)
	// SetProgressPulseStepEntry:
	SetProgressPulseStepEntry(fraction float64)
	// SetTabsEntry:
	SetTabsEntry(tabs *pango.TabArray)
	// SetTextEntry:
	SetTextEntry(text string)
	// SetVisibilityEntry:
	SetVisibilityEntry(visible bool)
	// SetWidthCharsEntry:
	SetWidthCharsEntry(nChars int)
	// TextIndexToLayoutIndexEntry:
	TextIndexToLayoutIndexEntry(textIndex int) int
	// UnsetInvisibleCharEntry:
	UnsetInvisibleCharEntry()
}

// entry implements the Entry class.
type entry struct {
	Widget
}

// WrapEntry wraps a GObject to the right type. It is
// primarily used internally.
func WrapEntry(obj *externglib.Object) Entry {
	return entry{
		Widget: WrapWidget(obj),
	}
}

func marshalEntry(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapEntry(obj), nil
}

// NewEntry:
func NewEntry() Entry {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_entry_new()

	var _entry Entry // out

	_entry = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Entry)

	return _entry
}

// NewEntryWithBuffer:
func NewEntryWithBuffer(buffer EntryBuffer) Entry {
	var _arg1 *C.GtkEntryBuffer // out
	var _cret *C.GtkWidget      // in

	_arg1 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))

	_cret = C.gtk_entry_new_with_buffer(_arg1)

	var _entry Entry // out

	_entry = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Entry)

	return _entry
}

func (e entry) ActivatesDefault() bool {
	var _arg0 *C.GtkEntry // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_activates_default(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e entry) Alignment() float32 {
	var _arg0 *C.GtkEntry // out
	var _cret C.gfloat    // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_alignment(_arg0)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

func (e entry) Attributes() *pango.AttrList {
	var _arg0 *C.GtkEntry      // out
	var _cret *C.PangoAttrList // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_attributes(_arg0)

	var _attrList *pango.AttrList // out

	_attrList = (*pango.AttrList)(unsafe.Pointer(_cret))

	return _attrList
}

func (e entry) Buffer() EntryBuffer {
	var _arg0 *C.GtkEntry       // out
	var _cret *C.GtkEntryBuffer // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_buffer(_arg0)

	var _entryBuffer EntryBuffer // out

	_entryBuffer = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(EntryBuffer)

	return _entryBuffer
}

func (e entry) Completion() EntryCompletion {
	var _arg0 *C.GtkEntry           // out
	var _cret *C.GtkEntryCompletion // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_completion(_arg0)

	var _entryCompletion EntryCompletion // out

	_entryCompletion = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(EntryCompletion)

	return _entryCompletion
}

func (e entry) CurrentIconDragSource() int {
	var _arg0 *C.GtkEntry // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_current_icon_drag_source(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (e entry) CursorHAdjustment() Adjustment {
	var _arg0 *C.GtkEntry      // out
	var _cret *C.GtkAdjustment // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_cursor_hadjustment(_arg0)

	var _adjustment Adjustment // out

	_adjustment = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(Adjustment)

	return _adjustment
}

func (e entry) HasFrame() bool {
	var _arg0 *C.GtkEntry // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_has_frame(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e entry) IconActivatable(iconPos EntryIconPosition) bool {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)

	_cret = C.gtk_entry_get_icon_activatable(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e entry) IconArea(iconPos EntryIconPosition) gdk.Rectangle {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _arg2 C.GdkRectangle         // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)

	C.gtk_entry_get_icon_area(_arg0, _arg1, &_arg2)

	var _iconArea gdk.Rectangle // out

	{
		var refTmpIn *C.GdkRectangle
		var refTmpOut *gdk.Rectangle

		in0 := &_arg2
		refTmpIn = in0

		refTmpOut = (*gdk.Rectangle)(unsafe.Pointer(refTmpIn))

		_iconArea = *refTmpOut
	}

	return _iconArea
}

func (e entry) IconAtPos(x int, y int) int {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gint      // out
	var _arg2 C.gint      // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)

	_cret = C.gtk_entry_get_icon_at_pos(_arg0, _arg1, _arg2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (e entry) IconGIcon(iconPos EntryIconPosition) gio.Icon {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _cret *C.GIcon               // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)

	_cret = C.gtk_entry_get_icon_gicon(_arg0, _arg1)

	var _icon gio.Icon // out

	_icon = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gio.Icon)

	return _icon
}

func (e entry) IconName(iconPos EntryIconPosition) string {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _cret *C.gchar               // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)

	_cret = C.gtk_entry_get_icon_name(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (e entry) IconPixbuf(iconPos EntryIconPosition) gdkpixbuf.Pixbuf {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _cret *C.GdkPixbuf           // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)

	_cret = C.gtk_entry_get_icon_pixbuf(_arg0, _arg1)

	var _pixbuf gdkpixbuf.Pixbuf // out

	_pixbuf = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(gdkpixbuf.Pixbuf)

	return _pixbuf
}

func (e entry) IconSensitive(iconPos EntryIconPosition) bool {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)

	_cret = C.gtk_entry_get_icon_sensitive(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e entry) IconStock(iconPos EntryIconPosition) string {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _cret *C.gchar               // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)

	_cret = C.gtk_entry_get_icon_stock(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (e entry) IconStorageType(iconPos EntryIconPosition) ImageType {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _cret C.GtkImageType         // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)

	_cret = C.gtk_entry_get_icon_storage_type(_arg0, _arg1)

	var _imageType ImageType // out

	_imageType = ImageType(_cret)

	return _imageType
}

func (e entry) IconTooltipMarkup(iconPos EntryIconPosition) string {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _cret *C.gchar               // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)

	_cret = C.gtk_entry_get_icon_tooltip_markup(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (e entry) IconTooltipText(iconPos EntryIconPosition) string {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _cret *C.gchar               // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)

	_cret = C.gtk_entry_get_icon_tooltip_text(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

func (e entry) InnerBorder() *Border {
	var _arg0 *C.GtkEntry  // out
	var _cret *C.GtkBorder // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_inner_border(_arg0)

	var _border *Border // out

	_border = (*Border)(unsafe.Pointer(_cret))

	return _border
}

func (e entry) InputHints() InputHints {
	var _arg0 *C.GtkEntry     // out
	var _cret C.GtkInputHints // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_input_hints(_arg0)

	var _inputHints InputHints // out

	_inputHints = InputHints(_cret)

	return _inputHints
}

func (e entry) InputPurpose() InputPurpose {
	var _arg0 *C.GtkEntry       // out
	var _cret C.GtkInputPurpose // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_input_purpose(_arg0)

	var _inputPurpose InputPurpose // out

	_inputPurpose = InputPurpose(_cret)

	return _inputPurpose
}

func (e entry) InvisibleChar() uint32 {
	var _arg0 *C.GtkEntry // out
	var _cret C.gunichar  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_invisible_char(_arg0)

	var _gunichar uint32 // out

	_gunichar = uint32(_cret)

	return _gunichar
}

func (e entry) Layout() pango.Layout {
	var _arg0 *C.GtkEntry    // out
	var _cret *C.PangoLayout // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_layout(_arg0)

	var _layout pango.Layout // out

	_layout = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret))).(pango.Layout)

	return _layout
}

func (e entry) LayoutOffsets() (x int, y int) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gint      // in
	var _arg2 C.gint      // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	C.gtk_entry_get_layout_offsets(_arg0, &_arg1, &_arg2)

	var _x int // out
	var _y int // out

	_x = int(_arg1)
	_y = int(_arg2)

	return _x, _y
}

func (e entry) MaxLength() int {
	var _arg0 *C.GtkEntry // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_max_length(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (e entry) MaxWidthChars() int {
	var _arg0 *C.GtkEntry // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_max_width_chars(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (e entry) OverwriteMode() bool {
	var _arg0 *C.GtkEntry // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_overwrite_mode(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e entry) PlaceholderText() string {
	var _arg0 *C.GtkEntry // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_placeholder_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (e entry) ProgressFraction() float64 {
	var _arg0 *C.GtkEntry // out
	var _cret C.gdouble   // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_progress_fraction(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (e entry) ProgressPulseStep() float64 {
	var _arg0 *C.GtkEntry // out
	var _cret C.gdouble   // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_progress_pulse_step(_arg0)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

func (e entry) Tabs() *pango.TabArray {
	var _arg0 *C.GtkEntry      // out
	var _cret *C.PangoTabArray // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_tabs(_arg0)

	var _tabArray *pango.TabArray // out

	_tabArray = (*pango.TabArray)(unsafe.Pointer(_cret))

	return _tabArray
}

func (e entry) Text() string {
	var _arg0 *C.GtkEntry // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_text(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

func (e entry) TextArea() gdk.Rectangle {
	var _arg0 *C.GtkEntry    // out
	var _arg1 C.GdkRectangle // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	C.gtk_entry_get_text_area(_arg0, &_arg1)

	var _textArea gdk.Rectangle // out

	{
		var refTmpIn *C.GdkRectangle
		var refTmpOut *gdk.Rectangle

		in0 := &_arg1
		refTmpIn = in0

		refTmpOut = (*gdk.Rectangle)(unsafe.Pointer(refTmpIn))

		_textArea = *refTmpOut
	}

	return _textArea
}

func (e entry) TextLength() uint16 {
	var _arg0 *C.GtkEntry // out
	var _cret C.guint16   // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_text_length(_arg0)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

func (e entry) Visibility() bool {
	var _arg0 *C.GtkEntry // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_visibility(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e entry) WidthChars() int {
	var _arg0 *C.GtkEntry // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	_cret = C.gtk_entry_get_width_chars(_arg0)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (e entry) GrabFocusWithoutSelectingEntry() {
	var _arg0 *C.GtkEntry // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	C.gtk_entry_grab_focus_without_selecting(_arg0)
}

func (e entry) ImContextFilterKeypressEntry(event *gdk.EventKey) bool {
	var _arg0 *C.GtkEntry    // out
	var _arg1 *C.GdkEventKey // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.GdkEventKey)(unsafe.Pointer(event.Native()))

	_cret = C.gtk_entry_im_context_filter_keypress(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (e entry) LayoutIndexToTextIndexEntry(layoutIndex int) int {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gint      // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.gint(layoutIndex)

	_cret = C.gtk_entry_layout_index_to_text_index(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (e entry) ProgressPulseEntry() {
	var _arg0 *C.GtkEntry // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	C.gtk_entry_progress_pulse(_arg0)
}

func (e entry) ResetImContextEntry() {
	var _arg0 *C.GtkEntry // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	C.gtk_entry_reset_im_context(_arg0)
}

func (e entry) SetActivatesDefaultEntry(setting bool) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_entry_set_activates_default(_arg0, _arg1)
}

func (e entry) SetAlignmentEntry(xalign float32) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gfloat    // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.gfloat(xalign)

	C.gtk_entry_set_alignment(_arg0, _arg1)
}

func (e entry) SetAttributesEntry(attrs *pango.AttrList) {
	var _arg0 *C.GtkEntry      // out
	var _arg1 *C.PangoAttrList // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.PangoAttrList)(unsafe.Pointer(attrs.Native()))

	C.gtk_entry_set_attributes(_arg0, _arg1)
}

func (e entry) SetBufferEntry(buffer EntryBuffer) {
	var _arg0 *C.GtkEntry       // out
	var _arg1 *C.GtkEntryBuffer // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.GtkEntryBuffer)(unsafe.Pointer(buffer.Native()))

	C.gtk_entry_set_buffer(_arg0, _arg1)
}

func (e entry) SetCompletionEntry(completion EntryCompletion) {
	var _arg0 *C.GtkEntry           // out
	var _arg1 *C.GtkEntryCompletion // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.GtkEntryCompletion)(unsafe.Pointer(completion.Native()))

	C.gtk_entry_set_completion(_arg0, _arg1)
}

func (e entry) SetCursorHAdjustmentEntry(adjustment Adjustment) {
	var _arg0 *C.GtkEntry      // out
	var _arg1 *C.GtkAdjustment // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.GtkAdjustment)(unsafe.Pointer(adjustment.Native()))

	C.gtk_entry_set_cursor_hadjustment(_arg0, _arg1)
}

func (e entry) SetHasFrameEntry(setting bool) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	if setting {
		_arg1 = C.TRUE
	}

	C.gtk_entry_set_has_frame(_arg0, _arg1)
}

func (e entry) SetIconActivatableEntry(iconPos EntryIconPosition, activatable bool) {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _arg2 C.gboolean             // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)
	if activatable {
		_arg2 = C.TRUE
	}

	C.gtk_entry_set_icon_activatable(_arg0, _arg1, _arg2)
}

func (e entry) SetIconDragSourceEntry(iconPos EntryIconPosition, targetList *TargetList, actions gdk.DragAction) {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _arg2 *C.GtkTargetList       // out
	var _arg3 C.GdkDragAction        // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)
	_arg2 = (*C.GtkTargetList)(unsafe.Pointer(targetList.Native()))
	_arg3 = C.GdkDragAction(actions)

	C.gtk_entry_set_icon_drag_source(_arg0, _arg1, _arg2, _arg3)
}

func (e entry) SetIconFromGIconEntry(iconPos EntryIconPosition, icon gio.Icon) {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _arg2 *C.GIcon               // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)
	_arg2 = (*C.GIcon)(unsafe.Pointer(icon.Native()))

	C.gtk_entry_set_icon_from_gicon(_arg0, _arg1, _arg2)
}

func (e entry) SetIconFromIconNameEntry(iconPos EntryIconPosition, iconName string) {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _arg2 *C.gchar               // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)
	_arg2 = (*C.gchar)(C.CString(iconName))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_entry_set_icon_from_icon_name(_arg0, _arg1, _arg2)
}

func (e entry) SetIconFromPixbufEntry(iconPos EntryIconPosition, pixbuf gdkpixbuf.Pixbuf) {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _arg2 *C.GdkPixbuf           // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)
	_arg2 = (*C.GdkPixbuf)(unsafe.Pointer(pixbuf.Native()))

	C.gtk_entry_set_icon_from_pixbuf(_arg0, _arg1, _arg2)
}

func (e entry) SetIconFromStockEntry(iconPos EntryIconPosition, stockId string) {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _arg2 *C.gchar               // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)
	_arg2 = (*C.gchar)(C.CString(stockId))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_entry_set_icon_from_stock(_arg0, _arg1, _arg2)
}

func (e entry) SetIconSensitiveEntry(iconPos EntryIconPosition, sensitive bool) {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _arg2 C.gboolean             // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)
	if sensitive {
		_arg2 = C.TRUE
	}

	C.gtk_entry_set_icon_sensitive(_arg0, _arg1, _arg2)
}

func (e entry) SetIconTooltipMarkupEntry(iconPos EntryIconPosition, tooltip string) {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _arg2 *C.gchar               // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)
	_arg2 = (*C.gchar)(C.CString(tooltip))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_entry_set_icon_tooltip_markup(_arg0, _arg1, _arg2)
}

func (e entry) SetIconTooltipTextEntry(iconPos EntryIconPosition, tooltip string) {
	var _arg0 *C.GtkEntry            // out
	var _arg1 C.GtkEntryIconPosition // out
	var _arg2 *C.gchar               // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkEntryIconPosition(iconPos)
	_arg2 = (*C.gchar)(C.CString(tooltip))
	defer C.free(unsafe.Pointer(_arg2))

	C.gtk_entry_set_icon_tooltip_text(_arg0, _arg1, _arg2)
}

func (e entry) SetInnerBorderEntry(border *Border) {
	var _arg0 *C.GtkEntry  // out
	var _arg1 *C.GtkBorder // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.GtkBorder)(unsafe.Pointer(border.Native()))

	C.gtk_entry_set_inner_border(_arg0, _arg1)
}

func (e entry) SetInputHintsEntry(hints InputHints) {
	var _arg0 *C.GtkEntry     // out
	var _arg1 C.GtkInputHints // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkInputHints(hints)

	C.gtk_entry_set_input_hints(_arg0, _arg1)
}

func (e entry) SetInputPurposeEntry(purpose InputPurpose) {
	var _arg0 *C.GtkEntry       // out
	var _arg1 C.GtkInputPurpose // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.GtkInputPurpose(purpose)

	C.gtk_entry_set_input_purpose(_arg0, _arg1)
}

func (e entry) SetInvisibleCharEntry(ch uint32) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gunichar  // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.gunichar(ch)

	C.gtk_entry_set_invisible_char(_arg0, _arg1)
}

func (e entry) SetMaxLengthEntry(max int) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.gint(max)

	C.gtk_entry_set_max_length(_arg0, _arg1)
}

func (e entry) SetMaxWidthCharsEntry(nChars int) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.gint(nChars)

	C.gtk_entry_set_max_width_chars(_arg0, _arg1)
}

func (e entry) SetOverwriteModeEntry(overwrite bool) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	if overwrite {
		_arg1 = C.TRUE
	}

	C.gtk_entry_set_overwrite_mode(_arg0, _arg1)
}

func (e entry) SetPlaceholderTextEntry(text string) {
	var _arg0 *C.GtkEntry // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_entry_set_placeholder_text(_arg0, _arg1)
}

func (e entry) SetProgressFractionEntry(fraction float64) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gdouble   // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.gdouble(fraction)

	C.gtk_entry_set_progress_fraction(_arg0, _arg1)
}

func (e entry) SetProgressPulseStepEntry(fraction float64) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gdouble   // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.gdouble(fraction)

	C.gtk_entry_set_progress_pulse_step(_arg0, _arg1)
}

func (e entry) SetTabsEntry(tabs *pango.TabArray) {
	var _arg0 *C.GtkEntry      // out
	var _arg1 *C.PangoTabArray // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.PangoTabArray)(unsafe.Pointer(tabs.Native()))

	C.gtk_entry_set_tabs(_arg0, _arg1)
}

func (e entry) SetTextEntry(text string) {
	var _arg0 *C.GtkEntry // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = (*C.gchar)(C.CString(text))
	defer C.free(unsafe.Pointer(_arg1))

	C.gtk_entry_set_text(_arg0, _arg1)
}

func (e entry) SetVisibilityEntry(visible bool) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	if visible {
		_arg1 = C.TRUE
	}

	C.gtk_entry_set_visibility(_arg0, _arg1)
}

func (e entry) SetWidthCharsEntry(nChars int) {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gint      // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.gint(nChars)

	C.gtk_entry_set_width_chars(_arg0, _arg1)
}

func (e entry) TextIndexToLayoutIndexEntry(textIndex int) int {
	var _arg0 *C.GtkEntry // out
	var _arg1 C.gint      // out
	var _cret C.gint      // in

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))
	_arg1 = C.gint(textIndex)

	_cret = C.gtk_entry_text_index_to_layout_index(_arg0, _arg1)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

func (e entry) UnsetInvisibleCharEntry() {
	var _arg0 *C.GtkEntry // out

	_arg0 = (*C.GtkEntry)(unsafe.Pointer(e.Native()))

	C.gtk_entry_unset_invisible_char(_arg0)
}

func (c entry) EditingDone() {
	WrapCellEditable(gextras.InternObject(c)).EditingDone()
}

func (c entry) RemoveWidget() {
	WrapCellEditable(gextras.InternObject(c)).RemoveWidget()
}

func (b entry) AddChild(builder Builder, child gextras.Objector, typ string) {
	WrapBuildable(gextras.InternObject(b)).AddChild(builder, child, typ)
}

func (b entry) ConstructChild(builder Builder, name string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).ConstructChild(builder, name)
}

func (b entry) InternalChild(builder Builder, childname string) gextras.Objector {
	return WrapBuildable(gextras.InternObject(b)).InternalChild(builder, childname)
}

func (b entry) Name() string {
	return WrapBuildable(gextras.InternObject(b)).Name()
}

func (b entry) ParserFinished(builder Builder) {
	WrapBuildable(gextras.InternObject(b)).ParserFinished(builder)
}

func (b entry) SetBuildableProperty(builder Builder, name string, value externglib.Value) {
	WrapBuildable(gextras.InternObject(b)).SetBuildableProperty(builder, name, value)
}

func (b entry) SetName(name string) {
	WrapBuildable(gextras.InternObject(b)).SetName(name)
}

func (e entry) CopyClipboard() {
	WrapEditable(gextras.InternObject(e)).CopyClipboard()
}

func (e entry) CutClipboard() {
	WrapEditable(gextras.InternObject(e)).CutClipboard()
}

func (e entry) DeleteSelection() {
	WrapEditable(gextras.InternObject(e)).DeleteSelection()
}

func (e entry) DeleteText(startPos int, endPos int) {
	WrapEditable(gextras.InternObject(e)).DeleteText(startPos, endPos)
}

func (e entry) Chars(startPos int, endPos int) string {
	return WrapEditable(gextras.InternObject(e)).Chars(startPos, endPos)
}

func (e entry) Editable() bool {
	return WrapEditable(gextras.InternObject(e)).Editable()
}

func (e entry) Position() int {
	return WrapEditable(gextras.InternObject(e)).Position()
}

func (e entry) SelectionBounds() (startPos int, endPos int, ok bool) {
	return WrapEditable(gextras.InternObject(e)).SelectionBounds()
}

func (e entry) PasteClipboard() {
	WrapEditable(gextras.InternObject(e)).PasteClipboard()
}

func (e entry) SelectRegion(startPos int, endPos int) {
	WrapEditable(gextras.InternObject(e)).SelectRegion(startPos, endPos)
}

func (e entry) SetEditable(isEditable bool) {
	WrapEditable(gextras.InternObject(e)).SetEditable(isEditable)
}

func (e entry) SetPosition(position int) {
	WrapEditable(gextras.InternObject(e)).SetPosition(position)
}