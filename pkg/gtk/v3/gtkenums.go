// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_align_get_type()), F: marshalAlign},
		{T: externglib.Type(C.gtk_arrow_type_get_type()), F: marshalArrowType},
		{T: externglib.Type(C.gtk_baseline_position_get_type()), F: marshalBaselinePosition},
		{T: externglib.Type(C.gtk_border_style_get_type()), F: marshalBorderStyle},
		{T: externglib.Type(C.gtk_delete_type_get_type()), F: marshalDeleteType},
		{T: externglib.Type(C.gtk_direction_type_get_type()), F: marshalDirectionType},
		{T: externglib.Type(C.gtk_drag_result_get_type()), F: marshalDragResult},
		{T: externglib.Type(C.gtk_event_sequence_state_get_type()), F: marshalEventSequenceState},
		{T: externglib.Type(C.gtk_im_preedit_style_get_type()), F: marshalIMPreeditStyle},
		{T: externglib.Type(C.gtk_im_status_style_get_type()), F: marshalIMStatusStyle},
		{T: externglib.Type(C.gtk_icon_size_get_type()), F: marshalIconSize},
		{T: externglib.Type(C.gtk_input_purpose_get_type()), F: marshalInputPurpose},
		{T: externglib.Type(C.gtk_justification_get_type()), F: marshalJustification},
		{T: externglib.Type(C.gtk_level_bar_mode_get_type()), F: marshalLevelBarMode},
		{T: externglib.Type(C.gtk_menu_direction_type_get_type()), F: marshalMenuDirectionType},
		{T: externglib.Type(C.gtk_message_type_get_type()), F: marshalMessageType},
		{T: externglib.Type(C.gtk_number_up_layout_get_type()), F: marshalNumberUpLayout},
		{T: externglib.Type(C.gtk_orientation_get_type()), F: marshalOrientation},
		{T: externglib.Type(C.gtk_pack_direction_get_type()), F: marshalPackDirection},
		{T: externglib.Type(C.gtk_pack_type_get_type()), F: marshalPackType},
		{T: externglib.Type(C.gtk_page_orientation_get_type()), F: marshalPageOrientation},
		{T: externglib.Type(C.gtk_page_set_get_type()), F: marshalPageSet},
		{T: externglib.Type(C.gtk_pan_direction_get_type()), F: marshalPanDirection},
		{T: externglib.Type(C.gtk_popover_constraint_get_type()), F: marshalPopoverConstraint},
		{T: externglib.Type(C.gtk_position_type_get_type()), F: marshalPositionType},
		{T: externglib.Type(C.gtk_print_duplex_get_type()), F: marshalPrintDuplex},
		{T: externglib.Type(C.gtk_print_pages_get_type()), F: marshalPrintPages},
		{T: externglib.Type(C.gtk_print_quality_get_type()), F: marshalPrintQuality},
		{T: externglib.Type(C.gtk_propagation_phase_get_type()), F: marshalPropagationPhase},
		{T: externglib.Type(C.gtk_relief_style_get_type()), F: marshalReliefStyle},
		{T: externglib.Type(C.gtk_scroll_type_get_type()), F: marshalScrollType},
		{T: externglib.Type(C.gtk_scrollable_policy_get_type()), F: marshalScrollablePolicy},
		{T: externglib.Type(C.gtk_selection_mode_get_type()), F: marshalSelectionMode},
		{T: externglib.Type(C.gtk_sensitivity_type_get_type()), F: marshalSensitivityType},
		{T: externglib.Type(C.gtk_shadow_type_get_type()), F: marshalShadowType},
		{T: externglib.Type(C.gtk_size_group_mode_get_type()), F: marshalSizeGroupMode},
		{T: externglib.Type(C.gtk_size_request_mode_get_type()), F: marshalSizeRequestMode},
		{T: externglib.Type(C.gtk_sort_type_get_type()), F: marshalSortType},
		{T: externglib.Type(C.gtk_state_type_get_type()), F: marshalStateType},
		{T: externglib.Type(C.gtk_text_direction_get_type()), F: marshalTextDirection},
		{T: externglib.Type(C.gtk_toolbar_style_get_type()), F: marshalToolbarStyle},
		{T: externglib.Type(C.gtk_tree_view_grid_lines_get_type()), F: marshalTreeViewGridLines},
		{T: externglib.Type(C.gtk_unit_get_type()), F: marshalUnit},
		{T: externglib.Type(C.gtk_wrap_mode_get_type()), F: marshalWrapMode},
		{T: externglib.Type(C.gtk_input_hints_get_type()), F: marshalInputHints},
		{T: externglib.Type(C.gtk_junction_sides_get_type()), F: marshalJunctionSides},
		{T: externglib.Type(C.gtk_region_flags_get_type()), F: marshalRegionFlags},
		{T: externglib.Type(C.gtk_state_flags_get_type()), F: marshalStateFlags},
	})
}

// Align controls how a widget deals with extra space in a single (x or y)
// dimension.
//
// Alignment only matters if the widget receives a “too large” allocation, for
// example if you packed the widget with the Widget:expand flag inside a Box,
// then the widget might get extra space. If you have for example a 16x16 icon
// inside a 32x32 space, the icon could be scaled and stretched, it could be
// centered, or it could be positioned to one side of the space.
//
// Note that in horizontal context @GTK_ALIGN_START and @GTK_ALIGN_END are
// interpreted relative to text direction.
//
// GTK_ALIGN_BASELINE support for it is optional for containers and widgets, and
// it is only supported for vertical alignment. When its not supported by a
// child or a container it is treated as @GTK_ALIGN_FILL.
type Align int

const (
	// Fill: stretch to fill all space if possible, center if no meaningful way
	// to stretch
	Fill Align = iota
	// Start: snap to left or top side, leaving space on right or bottom
	Start
	// End: snap to right or bottom side, leaving space on left or top
	End
	// Center natural width of widget inside the allocation
	Center
	// Baseline: align the widget according to the baseline. Since 3.10.
	Baseline
)

func marshalAlign(p uintptr) (interface{}, error) {
	return Align(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ArrowType: used to indicate the direction in which an arrow should point.
type ArrowType int

const (
	// Up represents an upward pointing arrow.
	Up ArrowType = iota
	// Down represents a downward pointing arrow.
	Down
	// Left represents a left pointing arrow.
	Left
	// Right represents a right pointing arrow.
	Right
	// None: no arrow. Since 2.10.
	None
)

func marshalArrowType(p uintptr) (interface{}, error) {
	return ArrowType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// BaselinePosition: whenever a container has some form of natural row it may
// align children in that row along a common typographical baseline. If the
// amount of verical space in the row is taller than the total requested height
// of the baseline-aligned children then it can use a BaselinePosition to select
// where to put the baseline inside the extra availible space.
type BaselinePosition int

const (
	// Top: align the baseline at the top
	Top BaselinePosition = iota
	// Center: center the baseline
	Center
	// Bottom: align the baseline at the bottom
	Bottom
)

func marshalBaselinePosition(p uintptr) (interface{}, error) {
	return BaselinePosition(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// BorderStyle describes how the border of a UI element should be rendered.
type BorderStyle int

const (
	// None: no visible border
	None BorderStyle = iota
	// Solid: single line segment
	Solid
	// Inset looks as if the content is sunken into the canvas
	Inset
	// Outset looks as if the content is coming out of the canvas
	Outset
	// Hidden: same as @GTK_BORDER_STYLE_NONE
	Hidden
	// Dotted series of round dots
	Dotted
	// Dashed series of square-ended dashes
	Dashed
	// Double: two parallel lines with some space between them
	Double
	// Groove looks as if it were carved in the canvas
	Groove
	// Ridge looks as if it were coming out of the canvas
	Ridge
)

func marshalBorderStyle(p uintptr) (interface{}, error) {
	return BorderStyle(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// DeleteType: see also: Entry::delete-from-cursor.
type DeleteType int

const (
	// Chars: delete characters.
	Chars DeleteType = iota
	// WordEnds: delete only the portion of the word to the left/right of cursor
	// if we’re in the middle of a word.
	WordEnds
	// Words: delete words.
	Words
	// DisplayLines: delete display-lines. Display-lines refers to the visible
	// lines, with respect to to the current line breaks. As opposed to
	// paragraphs, which are defined by line breaks in the input.
	DisplayLines
	// DisplayLineEnds: delete only the portion of the display-line to the
	// left/right of cursor.
	DisplayLineEnds
	// ParagraphEnds: delete to the end of the paragraph. Like C-k in Emacs (or
	// its reverse).
	ParagraphEnds
	// Paragraphs: delete entire line. Like C-k in pico.
	Paragraphs
	// Whitespace: delete only whitespace. Like M-\ in Emacs.
	Whitespace
)

func marshalDeleteType(p uintptr) (interface{}, error) {
	return DeleteType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// DirectionType focus movement types.
type DirectionType int

const (
	// TabForward: move forward.
	TabForward DirectionType = iota
	// TabBackward: move backward.
	TabBackward
	// Up: move up.
	Up
	// Down: move down.
	Down
	// Left: move left.
	Left
	// Right: move right.
	Right
)

func marshalDirectionType(p uintptr) (interface{}, error) {
	return DirectionType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// DragResult gives an indication why a drag operation failed. The value can by
// obtained by connecting to the Widget::drag-failed signal.
type DragResult int

const (
	// Success: the drag operation was successful.
	Success DragResult = iota
	// NoTarget: no suitable drag target.
	NoTarget
	// UserCancelled: the user cancelled the drag operation.
	UserCancelled
	// TimeoutExpired: the drag operation timed out.
	TimeoutExpired
	// GrabBroken: the pointer or keyboard grab used for the drag operation was
	// broken.
	GrabBroken
	// Error: the drag operation failed due to some unspecified error.
	Error
)

func marshalDragResult(p uintptr) (interface{}, error) {
	return DragResult(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// EventSequenceState describes the state of a EventSequence in a Gesture.
type EventSequenceState int

const (
	// None: the sequence is handled, but not grabbed.
	None EventSequenceState = iota
	// Claimed: the sequence is handled and grabbed.
	Claimed
	// Denied: the sequence is denied.
	Denied
)

func marshalEventSequenceState(p uintptr) (interface{}, error) {
	return EventSequenceState(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// IMPreeditStyle: style for input method preedit. See also
// Settings:gtk-im-preedit-style
//
// Deprecated: since version 3.10.
type IMPreeditStyle int

const (
	// Nothing: deprecated
	Nothing IMPreeditStyle = iota
	// Callback: deprecated
	Callback
	// None: deprecated
	None
)

func marshalIMPreeditStyle(p uintptr) (interface{}, error) {
	return IMPreeditStyle(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// IMStatusStyle: style for input method status. See also
// Settings:gtk-im-status-style
//
// Deprecated: since version 3.10.
type IMStatusStyle int

const (
	// Nothing: deprecated
	Nothing IMStatusStyle = iota
	// Callback: deprecated
	Callback
	// None: deprecated
	None
)

func marshalIMStatusStyle(p uintptr) (interface{}, error) {
	return IMStatusStyle(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// IconSize: built-in stock icon sizes.
type IconSize int

const (
	// Invalid: invalid size.
	Invalid IconSize = iota
	// Menu: size appropriate for menus (16px).
	Menu
	// SmallToolbar: size appropriate for small toolbars (16px).
	SmallToolbar
	// LargeToolbar: size appropriate for large toolbars (24px)
	LargeToolbar
	// Button: size appropriate for buttons (16px)
	Button
	// Dnd: size appropriate for drag and drop (32px)
	Dnd
	// Dialog: size appropriate for dialogs (48px)
	Dialog
)

func marshalIconSize(p uintptr) (interface{}, error) {
	return IconSize(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// InputPurpose describes primary purpose of the input widget. This information
// is useful for on-screen keyboards and similar input methods to decide which
// keys should be presented to the user.
//
// Note that the purpose is not meant to impose a totally strict rule about
// allowed characters, and does not replace input validation. It is fine for an
// on-screen keyboard to let the user override the character set restriction
// that is expressed by the purpose. The application is expected to validate the
// entry contents, even if it specified a purpose.
//
// The difference between @GTK_INPUT_PURPOSE_DIGITS and
// @GTK_INPUT_PURPOSE_NUMBER is that the former accepts only digits while the
// latter also some punctuation (like commas or points, plus, minus) and “e” or
// “E” as in 3.14E+000.
//
// This enumeration may be extended in the future; input methods should
// interpret unknown values as “free form”.
type InputPurpose int

const (
	// FreeForm: allow any character
	FreeForm InputPurpose = iota
	// Alpha: allow only alphabetic characters
	Alpha
	// Digits: allow only digits
	Digits
	// Number: edited field expects numbers
	Number
	// Phone: edited field expects phone number
	Phone
	// URL: edited field expects URL
	URL
	// Email: edited field expects email address
	Email
	// Name: edited field expects the name of a person
	Name
	// Password: like @GTK_INPUT_PURPOSE_FREE_FORM, but characters are hidden
	Password
	// Pin: like @GTK_INPUT_PURPOSE_DIGITS, but characters are hidden
	Pin
	// Terminal: allow any character, in addition to control codes
	Terminal
)

func marshalInputPurpose(p uintptr) (interface{}, error) {
	return InputPurpose(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Justification: used for justifying the text inside a Label widget. (See also
// Alignment).
type Justification int

const (
	// Left: the text is placed at the left edge of the label.
	Left Justification = iota
	// Right: the text is placed at the right edge of the label.
	Right
	// Center: the text is placed in the center of the label.
	Center
	// Fill: the text is placed is distributed across the label.
	Fill
)

func marshalJustification(p uintptr) (interface{}, error) {
	return Justification(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// LevelBarMode describes how LevelBar contents should be rendered. Note that
// this enumeration could be extended with additional modes in the future.
type LevelBarMode int

const (
	// Continuous: the bar has a continuous mode
	Continuous LevelBarMode = iota
	// Discrete: the bar has a discrete mode
	Discrete
)

func marshalLevelBarMode(p uintptr) (interface{}, error) {
	return LevelBarMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// MenuDirectionType: enumeration representing directional movements within a
// menu.
type MenuDirectionType int

const (
	// Parent: to the parent menu shell
	Parent MenuDirectionType = iota
	// Child: to the submenu, if any, associated with the item
	Child
	// Next: to the next menu item
	Next
	// Prev: to the previous menu item
	Prev
)

func marshalMenuDirectionType(p uintptr) (interface{}, error) {
	return MenuDirectionType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// MessageType: the type of message being displayed in the dialog.
type MessageType int

const (
	// Info: informational message
	Info MessageType = iota
	// Warning: non-fatal warning message
	Warning
	// Question: question requiring a choice
	Question
	// Error: fatal error message
	Error
	// Other: none of the above
	Other
)

func marshalMessageType(p uintptr) (interface{}, error) {
	return MessageType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// NumberUpLayout: used to determine the layout of pages on a sheet when
// printing multiple pages per sheet.
type NumberUpLayout int

const (
	// Lrtb: ! (layout-lrtb.png)
	Lrtb NumberUpLayout = iota
	// Lrbt: ! (layout-lrbt.png)
	Lrbt
	// Rltb: ! (layout-rltb.png)
	Rltb
	// Rlbt: ! (layout-rlbt.png)
	Rlbt
	// Tblr: ! (layout-tblr.png)
	Tblr
	// Tbrl: ! (layout-tbrl.png)
	Tbrl
	// Btlr: ! (layout-btlr.png)
	Btlr
	// Btrl: ! (layout-btrl.png)
	Btrl
)

func marshalNumberUpLayout(p uintptr) (interface{}, error) {
	return NumberUpLayout(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Orientation represents the orientation of widgets and other objects which can
// be switched between horizontal and vertical orientation on the fly, like
// Toolbar or GesturePan.
type Orientation int

const (
	// Horizontal: the element is in horizontal orientation.
	Horizontal Orientation = iota
	// Vertical: the element is in vertical orientation.
	Vertical
)

func marshalOrientation(p uintptr) (interface{}, error) {
	return Orientation(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PackDirection determines how widgets should be packed inside menubars and
// menuitems contained in menubars.
type PackDirection int

const (
	// LTR widgets are packed left-to-right
	LTR PackDirection = iota
	// RTL widgets are packed right-to-left
	RTL
	// Ttb widgets are packed top-to-bottom
	Ttb
	// Btt widgets are packed bottom-to-top
	Btt
)

func marshalPackDirection(p uintptr) (interface{}, error) {
	return PackDirection(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PackType represents the packing location Box children. (See: VBox, HBox, and
// ButtonBox).
type PackType int

const (
	// Start: the child is packed into the start of the box
	Start PackType = iota
	// End: the child is packed into the end of the box
	End
)

func marshalPackType(p uintptr) (interface{}, error) {
	return PackType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PageOrientation: see also gtk_print_settings_set_orientation().
type PageOrientation int

const (
	// Portrait: portrait mode.
	Portrait PageOrientation = iota
	// Landscape: landscape mode.
	Landscape
	// ReversePortrait: reverse portrait mode.
	ReversePortrait
	// ReverseLandscape: reverse landscape mode.
	ReverseLandscape
)

func marshalPageOrientation(p uintptr) (interface{}, error) {
	return PageOrientation(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PageSet: see also gtk_print_job_set_page_set().
type PageSet int

const (
	// All: all pages.
	All PageSet = iota
	// Even: even pages.
	Even
	// Odd: odd pages.
	Odd
)

func marshalPageSet(p uintptr) (interface{}, error) {
	return PageSet(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PanDirection describes the panning direction of a GesturePan
type PanDirection int

const (
	// Left: panned towards the left
	Left PanDirection = iota
	// Right: panned towards the right
	Right
	// Up: panned upwards
	Up
	// Down: panned downwards
	Down
)

func marshalPanDirection(p uintptr) (interface{}, error) {
	return PanDirection(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PopoverConstraint describes constraints to positioning of popovers. More
// values may be added to this enumeration in the future.
type PopoverConstraint int

const (
	// None: don't constrain the popover position beyond what is imposed by the
	// implementation
	None PopoverConstraint = iota
	// Window: constrain the popover to the boundaries of the window that it is
	// attached to
	Window
)

func marshalPopoverConstraint(p uintptr) (interface{}, error) {
	return PopoverConstraint(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PositionType describes which edge of a widget a certain feature is positioned
// at, e.g. the tabs of a Notebook, the handle of a HandleBox or the label of a
// Scale.
type PositionType int

const (
	// Left: the feature is at the left edge.
	Left PositionType = iota
	// Right: the feature is at the right edge.
	Right
	// Top: the feature is at the top edge.
	Top
	// Bottom: the feature is at the bottom edge.
	Bottom
)

func marshalPositionType(p uintptr) (interface{}, error) {
	return PositionType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PrintDuplex: see also gtk_print_settings_set_duplex().
type PrintDuplex int

const (
	// Simplex: no duplex.
	Simplex PrintDuplex = iota
	// Horizontal: horizontal duplex.
	Horizontal
	// Vertical: vertical duplex.
	Vertical
)

func marshalPrintDuplex(p uintptr) (interface{}, error) {
	return PrintDuplex(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PrintPages: see also gtk_print_job_set_pages()
type PrintPages int

const (
	// All: all pages.
	All PrintPages = iota
	// Current: current page.
	Current
	// Ranges: range of pages.
	Ranges
	// Selection: selected pages.
	Selection
)

func marshalPrintPages(p uintptr) (interface{}, error) {
	return PrintPages(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PrintQuality: see also gtk_print_settings_set_quality().
type PrintQuality int

const (
	// Low: low quality.
	Low PrintQuality = iota
	// Normal: normal quality.
	Normal
	// High: high quality.
	High
	// Draft: draft quality.
	Draft
)

func marshalPrintQuality(p uintptr) (interface{}, error) {
	return PrintQuality(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// PropagationPhase describes the stage at which events are fed into a
// EventController.
type PropagationPhase int

const (
	// None events are not delivered automatically. Those can be manually fed
	// through gtk_event_controller_handle_event(). This should only be used
	// when full control about when, or whether the controller handles the event
	// is needed.
	None PropagationPhase = iota
	// Capture events are delivered in the capture phase. The capture phase
	// happens before the bubble phase, runs from the toplevel down to the event
	// widget. This option should only be used on containers that might possibly
	// handle events before their children do.
	Capture
	// Bubble events are delivered in the bubble phase. The bubble phase happens
	// after the capture phase, and before the default handlers are run. This
	// phase runs from the event widget, up to the toplevel.
	Bubble
	// Target events are delivered in the default widget event handlers, note
	// that widget implementations must chain up on button, motion, touch and
	// grab broken handlers for controllers in this phase to be run.
	Target
)

func marshalPropagationPhase(p uintptr) (interface{}, error) {
	return PropagationPhase(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ReliefStyle: indicated the relief to be drawn around a Button.
type ReliefStyle int

const (
	// Normal: draw a normal relief.
	Normal ReliefStyle = iota
	// Half relief. Deprecated in 3.14, does the same as @GTK_RELIEF_NORMAL
	Half
	// None: no relief.
	None
)

func marshalReliefStyle(p uintptr) (interface{}, error) {
	return ReliefStyle(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ScrollType: scrolling types.
type ScrollType int

const (
	// None: no scrolling.
	None ScrollType = iota
	// Jump: jump to new location.
	Jump
	// StepBackward: step backward.
	StepBackward
	// StepForward: step forward.
	StepForward
	// PageBackward: page backward.
	PageBackward
	// PageForward: page forward.
	PageForward
	// StepUp: step up.
	StepUp
	// StepDown: step down.
	StepDown
	// PageUp: page up.
	PageUp
	// PageDown: page down.
	PageDown
	// StepLeft: step to the left.
	StepLeft
	// StepRight: step to the right.
	StepRight
	// PageLeft: page to the left.
	PageLeft
	// PageRight: page to the right.
	PageRight
	// Start: scroll to start.
	Start
	// End: scroll to end.
	End
)

func marshalScrollType(p uintptr) (interface{}, error) {
	return ScrollType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ScrollablePolicy defines the policy to be used in a scrollable widget when
// updating the scrolled window adjustments in a given orientation.
type ScrollablePolicy int

const (
	// Minimum: scrollable adjustments are based on the minimum size
	Minimum ScrollablePolicy = iota
	// Natural: scrollable adjustments are based on the natural size
	Natural
)

func marshalScrollablePolicy(p uintptr) (interface{}, error) {
	return ScrollablePolicy(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// SelectionMode: used to control what selections users are allowed to make.
type SelectionMode int

const (
	// None: no selection is possible.
	None SelectionMode = iota
	// Single: zero or one element may be selected.
	Single
	// Browse: exactly one element is selected. In some circumstances, such as
	// initially or during a search operation, it’s possible for no element to
	// be selected with GTK_SELECTION_BROWSE. What is really enforced is that
	// the user can’t deselect a currently selected element except by selecting
	// another element.
	Browse
	// Multiple: any number of elements may be selected. The Ctrl key may be
	// used to enlarge the selection, and Shift key to select between the focus
	// and the child pointed to. Some widgets may also allow Click-drag to
	// select a range of elements.
	Multiple
)

func marshalSelectionMode(p uintptr) (interface{}, error) {
	return SelectionMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// SensitivityType determines how GTK+ handles the sensitivity of stepper arrows
// at the end of range widgets.
type SensitivityType int

const (
	// Auto: the arrow is made insensitive if the thumb is at the end
	Auto SensitivityType = iota
	// On: the arrow is always sensitive
	On
	// Off: the arrow is always insensitive
	Off
)

func marshalSensitivityType(p uintptr) (interface{}, error) {
	return SensitivityType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ShadowType: used to change the appearance of an outline typically provided by
// a Frame.
//
// Note that many themes do not differentiate the appearance of the various
// shadow types: Either their is no visible shadow (@GTK_SHADOW_NONE), or there
// is (any other value).
type ShadowType int

const (
	// None: no outline.
	None ShadowType = iota
	// In: the outline is bevelled inwards.
	In
	// Out: the outline is bevelled outwards like a button.
	Out
	// EtchedIn: the outline has a sunken 3d appearance.
	EtchedIn
	// EtchedOut: the outline has a raised 3d appearance.
	EtchedOut
)

func marshalShadowType(p uintptr) (interface{}, error) {
	return ShadowType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// SizeGroupMode: the mode of the size group determines the directions in which
// the size group affects the requested sizes of its component widgets.
type SizeGroupMode int

const (
	// None: group has no effect
	None SizeGroupMode = iota
	// Horizontal: group affects horizontal requisition
	Horizontal
	// Vertical: group affects vertical requisition
	Vertical
	// Both: group affects both horizontal and vertical requisition
	Both
)

func marshalSizeGroupMode(p uintptr) (interface{}, error) {
	return SizeGroupMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// SizeRequestMode specifies a preference for height-for-width or
// width-for-height geometry management.
type SizeRequestMode int

const (
	// HeightForWidth: prefer height-for-width geometry management
	HeightForWidth SizeRequestMode = iota
	// WidthForHeight: prefer width-for-height geometry management
	WidthForHeight
	// ConstantSize: don’t trade height-for-width or width-for-height
	ConstantSize
)

func marshalSizeRequestMode(p uintptr) (interface{}, error) {
	return SizeRequestMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// SortType determines the direction of a sort.
type SortType int

const (
	// Ascending: sorting is in ascending order.
	Ascending SortType = iota
	// Descending: sorting is in descending order.
	Descending
)

func marshalSortType(p uintptr) (interface{}, error) {
	return SortType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// StateType: this type indicates the current state of a widget; the state
// determines how the widget is drawn. The StateType enumeration is also used to
// identify different colors in a Style for drawing, so states can be used for
// subparts of a widget as well as entire widgets.
//
// Deprecated: since version 3.14.
type StateType int

const (
	// Normal: state during normal operation.
	Normal StateType = iota
	// Active: state of a currently active widget, such as a depressed button.
	Active
	// Prelight: state indicating that the mouse pointer is over the widget and
	// the widget will respond to mouse clicks.
	Prelight
	// Selected: state of a selected item, such the selected row in a list.
	Selected
	// Insensitive: state indicating that the widget is unresponsive to user
	// actions.
	Insensitive
	// Inconsistent: the widget is inconsistent, such as checkbuttons or
	// radiobuttons that aren’t either set to true nor false, or buttons
	// requiring the user attention.
	Inconsistent
	// Focused: the widget has the keyboard focus.
	Focused
)

func marshalStateType(p uintptr) (interface{}, error) {
	return StateType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// TextDirection: reading directions for text.
type TextDirection int

const (
	// None: no direction.
	None TextDirection = iota
	// LTR: left to right text direction.
	LTR
	// RTL: right to left text direction.
	RTL
)

func marshalTextDirection(p uintptr) (interface{}, error) {
	return TextDirection(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// ToolbarStyle: used to customize the appearance of a Toolbar. Note that
// setting the toolbar style overrides the user’s preferences for the default
// toolbar style. Note that if the button has only a label set and
// GTK_TOOLBAR_ICONS is used, the label will be visible, and vice versa.
type ToolbarStyle int

const (
	// Icons buttons display only icons in the toolbar.
	Icons ToolbarStyle = iota
	// Text buttons display only text labels in the toolbar.
	Text
	// Both buttons display text and icons in the toolbar.
	Both
	// BothHoriz buttons display icons and text alongside each other, rather
	// than vertically stacked
	BothHoriz
)

func marshalToolbarStyle(p uintptr) (interface{}, error) {
	return ToolbarStyle(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// TreeViewGridLines: used to indicate which grid lines to draw in a tree view.
type TreeViewGridLines int

const (
	// None: no grid lines.
	None TreeViewGridLines = iota
	// Horizontal: horizontal grid lines.
	Horizontal
	// Vertical: vertical grid lines.
	Vertical
	// Both: horizontal and vertical grid lines.
	Both
)

func marshalTreeViewGridLines(p uintptr) (interface{}, error) {
	return TreeViewGridLines(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// Unit: see also gtk_print_settings_set_paper_width().
type Unit int

const (
	// None: no units.
	None Unit = iota
	// Points dimensions in points.
	Points
	// Inch dimensions in inches.
	Inch
	// Mm dimensions in millimeters
	Mm
)

func marshalUnit(p uintptr) (interface{}, error) {
	return Unit(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// WrapMode describes a type of line wrapping.
type WrapMode int

const (
	// None: do not wrap lines; just make the text area wider
	None WrapMode = iota
	// Char: wrap text, breaking lines anywhere the cursor can appear (between
	// characters, usually - if you want to be technical, between graphemes, see
	// pango_get_log_attrs())
	Char
	// Word: wrap text, breaking lines in between words
	Word
	// WordChar: wrap text, breaking lines in between words, or if that is not
	// enough, also between graphemes
	WordChar
)

func marshalWrapMode(p uintptr) (interface{}, error) {
	return WrapMode(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// InputHints describes hints that might be taken into account by input methods
// or applications. Note that input methods may already tailor their behaviour
// according to the InputPurpose of the entry.
//
// Some common sense is expected when using these flags - mixing
// @GTK_INPUT_HINT_LOWERCASE with any of the uppercase hints makes no sense.
//
// This enumeration may be extended in the future; input methods should ignore
// unknown values.
type InputHints int

const (
	// InputHintsNone: no special behaviour suggested
	InputHintsNone InputHints = 0b0
	// InputHintsSpellcheck: suggest checking for typos
	InputHintsSpellcheck InputHints = 0b1
	// InputHintsNoSpellcheck: suggest not checking for typos
	InputHintsNoSpellcheck InputHints = 0b10
	// InputHintsWordCompletion: suggest word completion
	InputHintsWordCompletion InputHints = 0b100
	// InputHintsLowercase: suggest to convert all text to lowercase
	InputHintsLowercase InputHints = 0b1000
	// InputHintsUppercaseChars: suggest to capitalize all text
	InputHintsUppercaseChars InputHints = 0b10000
	// InputHintsUppercaseWords: suggest to capitalize the first character of
	// each word
	InputHintsUppercaseWords InputHints = 0b100000
	// InputHintsUppercaseSentences: suggest to capitalize the first word of
	// each sentence
	InputHintsUppercaseSentences InputHints = 0b1000000
	// InputHintsInhibitOsk: suggest to not show an onscreen keyboard (e.g for a
	// calculator that already has all the keys).
	InputHintsInhibitOsk InputHints = 0b10000000
	// InputHintsVerticalWriting: the text is vertical. Since 3.18
	InputHintsVerticalWriting InputHints = 0b100000000
	// InputHintsEmoji: suggest offering Emoji support. Since 3.22.20
	InputHintsEmoji InputHints = 0b1000000000
	// InputHintsNoEmoji: suggest not offering Emoji support. Since 3.22.20
	InputHintsNoEmoji InputHints = 0b10000000000
)

func marshalInputHints(p uintptr) (interface{}, error) {
	return InputHints(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// JunctionSides describes how a rendered element connects to adjacent elements.
type JunctionSides int

const (
	// JunctionSidesNone: no junctions.
	JunctionSidesNone JunctionSides = 0b0
	// JunctionSidesCornerTopleft: element connects on the top-left corner.
	JunctionSidesCornerTopleft JunctionSides = 0b1
	// JunctionSidesCornerTopright: element connects on the top-right corner.
	JunctionSidesCornerTopright JunctionSides = 0b10
	// JunctionSidesCornerBottomleft: element connects on the bottom-left
	// corner.
	JunctionSidesCornerBottomleft JunctionSides = 0b100
	// JunctionSidesCornerBottomright: element connects on the bottom-right
	// corner.
	JunctionSidesCornerBottomright JunctionSides = 0b1000
	// JunctionSidesTop: element connects on the top side.
	JunctionSidesTop JunctionSides = 0b11
	// JunctionSidesBottom: element connects on the bottom side.
	JunctionSidesBottom JunctionSides = 0b1100
	// JunctionSidesLeft: element connects on the left side.
	JunctionSidesLeft JunctionSides = 0b101
	// JunctionSidesRight: element connects on the right side.
	JunctionSidesRight JunctionSides = 0b1010
)

func marshalJunctionSides(p uintptr) (interface{}, error) {
	return JunctionSides(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// RegionFlags describes a region within a widget.
type RegionFlags int

const (
	// RegionFlagsEven: region has an even number within a set.
	RegionFlagsEven RegionFlags = 0b1
	// RegionFlagsOdd: region has an odd number within a set.
	RegionFlagsOdd RegionFlags = 0b10
	// RegionFlagsFirst: region is the first one within a set.
	RegionFlagsFirst RegionFlags = 0b100
	// RegionFlagsLast: region is the last one within a set.
	RegionFlagsLast RegionFlags = 0b1000
	// RegionFlagsOnly: region is the only one within a set.
	RegionFlagsOnly RegionFlags = 0b10000
	// RegionFlagsSorted: region is part of a sorted area.
	RegionFlagsSorted RegionFlags = 0b100000
)

func marshalRegionFlags(p uintptr) (interface{}, error) {
	return RegionFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// StateFlags describes a widget state. Widget states are used to match the
// widget against CSS pseudo-classes. Note that GTK extends the regular CSS
// classes and sometimes uses different names.
type StateFlags int

const (
	// StateFlagsNormal: state during normal operation.
	StateFlagsNormal StateFlags = 0b0
	// StateFlagsActive: widget is active.
	StateFlagsActive StateFlags = 0b1
	// StateFlagsPrelight: widget has a mouse pointer over it.
	StateFlagsPrelight StateFlags = 0b10
	// StateFlagsSelected: widget is selected.
	StateFlagsSelected StateFlags = 0b100
	// StateFlagsInsensitive: widget is insensitive.
	StateFlagsInsensitive StateFlags = 0b1000
	// StateFlagsInconsistent: widget is inconsistent.
	StateFlagsInconsistent StateFlags = 0b10000
	// StateFlagsFocused: widget has the keyboard focus.
	StateFlagsFocused StateFlags = 0b100000
	// StateFlagsBackdrop: widget is in a background toplevel window.
	StateFlagsBackdrop StateFlags = 0b1000000
	// StateFlagsDirLTR: widget is in left-to-right text direction. Since 3.8
	StateFlagsDirLTR StateFlags = 0b10000000
	// StateFlagsDirRTL: widget is in right-to-left text direction. Since 3.8
	StateFlagsDirRTL StateFlags = 0b100000000
	// StateFlagsLink: widget is a link. Since 3.12
	StateFlagsLink StateFlags = 0b1000000000
	// StateFlagsVisited: the location the widget points to has already been
	// visited. Since 3.12
	StateFlagsVisited StateFlags = 0b10000000000
	// StateFlagsChecked: widget is checked. Since 3.14
	StateFlagsChecked StateFlags = 0b100000000000
	// StateFlagsDropActive: widget is highlighted as a drop target for DND.
	// Since 3.20
	StateFlagsDropActive StateFlags = 0b1000000000000
)

func marshalStateFlags(p uintptr) (interface{}, error) {
	return StateFlags(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}
