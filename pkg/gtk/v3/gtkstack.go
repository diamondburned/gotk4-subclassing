// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"fmt"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_stack_transition_type_get_type()), F: marshalStackTransitionType},
		{T: externglib.Type(C.gtk_stack_get_type()), F: marshalStacker},
	})
}

// StackTransitionType: these enumeration values describe the possible
// transitions between pages in a Stack widget.
//
// New values may be added to this enumeration over time.
type StackTransitionType int

const (
	// StackTransitionTypeNone: no transition
	StackTransitionTypeNone StackTransitionType = iota
	// StackTransitionTypeCrossfade: cross-fade
	StackTransitionTypeCrossfade
	// StackTransitionTypeSlideRight: slide from left to right
	StackTransitionTypeSlideRight
	// StackTransitionTypeSlideLeft: slide from right to left
	StackTransitionTypeSlideLeft
	// StackTransitionTypeSlideUp: slide from bottom up
	StackTransitionTypeSlideUp
	// StackTransitionTypeSlideDown: slide from top down
	StackTransitionTypeSlideDown
	// StackTransitionTypeSlideLeftRight: slide from left or right according to
	// the children order
	StackTransitionTypeSlideLeftRight
	// StackTransitionTypeSlideUpDown: slide from top down or bottom up
	// according to the order
	StackTransitionTypeSlideUpDown
	// StackTransitionTypeOverUp: cover the old page by sliding up. Since 3.12
	StackTransitionTypeOverUp
	// StackTransitionTypeOverDown: cover the old page by sliding down. Since:
	// 3.12
	StackTransitionTypeOverDown
	// StackTransitionTypeOverLeft: cover the old page by sliding to the left.
	// Since: 3.12
	StackTransitionTypeOverLeft
	// StackTransitionTypeOverRight: cover the old page by sliding to the right.
	// Since: 3.12
	StackTransitionTypeOverRight
	// StackTransitionTypeUnderUp: uncover the new page by sliding up. Since
	// 3.12
	StackTransitionTypeUnderUp
	// StackTransitionTypeUnderDown: uncover the new page by sliding down.
	// Since: 3.12
	StackTransitionTypeUnderDown
	// StackTransitionTypeUnderLeft: uncover the new page by sliding to the
	// left. Since: 3.12
	StackTransitionTypeUnderLeft
	// StackTransitionTypeUnderRight: uncover the new page by sliding to the
	// right. Since: 3.12
	StackTransitionTypeUnderRight
	// StackTransitionTypeOverUpDown: cover the old page sliding up or uncover
	// the new page sliding down, according to order. Since: 3.12
	StackTransitionTypeOverUpDown
	// StackTransitionTypeOverDownUp: cover the old page sliding down or uncover
	// the new page sliding up, according to order. Since: 3.14
	StackTransitionTypeOverDownUp
	// StackTransitionTypeOverLeftRight: cover the old page sliding left or
	// uncover the new page sliding right, according to order. Since: 3.14
	StackTransitionTypeOverLeftRight
	// StackTransitionTypeOverRightLeft: cover the old page sliding right or
	// uncover the new page sliding left, according to order. Since: 3.14
	StackTransitionTypeOverRightLeft
)

func marshalStackTransitionType(p uintptr) (interface{}, error) {
	return StackTransitionType(C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))), nil
}

// String returns the name in string for StackTransitionType.
func (s StackTransitionType) String() string {
	switch s {
	case StackTransitionTypeNone:
		return "None"
	case StackTransitionTypeCrossfade:
		return "Crossfade"
	case StackTransitionTypeSlideRight:
		return "SlideRight"
	case StackTransitionTypeSlideLeft:
		return "SlideLeft"
	case StackTransitionTypeSlideUp:
		return "SlideUp"
	case StackTransitionTypeSlideDown:
		return "SlideDown"
	case StackTransitionTypeSlideLeftRight:
		return "SlideLeftRight"
	case StackTransitionTypeSlideUpDown:
		return "SlideUpDown"
	case StackTransitionTypeOverUp:
		return "OverUp"
	case StackTransitionTypeOverDown:
		return "OverDown"
	case StackTransitionTypeOverLeft:
		return "OverLeft"
	case StackTransitionTypeOverRight:
		return "OverRight"
	case StackTransitionTypeUnderUp:
		return "UnderUp"
	case StackTransitionTypeUnderDown:
		return "UnderDown"
	case StackTransitionTypeUnderLeft:
		return "UnderLeft"
	case StackTransitionTypeUnderRight:
		return "UnderRight"
	case StackTransitionTypeOverUpDown:
		return "OverUpDown"
	case StackTransitionTypeOverDownUp:
		return "OverDownUp"
	case StackTransitionTypeOverLeftRight:
		return "OverLeftRight"
	case StackTransitionTypeOverRightLeft:
		return "OverRightLeft"
	default:
		return fmt.Sprintf("StackTransitionType(%d)", s)
	}
}

// Stack widget is a container which only shows one of its children at a time.
// In contrast to GtkNotebook, GtkStack does not provide a means for users to
// change the visible child. Instead, the StackSwitcher widget can be used with
// GtkStack to provide this functionality.
//
// Transitions between pages can be animated as slides or fades. This can be
// controlled with gtk_stack_set_transition_type(). These animations respect the
// Settings:gtk-enable-animations setting.
//
// The GtkStack widget was added in GTK+ 3.10.
//
//
// CSS nodes
//
// GtkStack has a single CSS node named stack.
type Stack struct {
	Container
}

var _ gextras.Nativer = (*Stack)(nil)

func wrapStack(obj *externglib.Object) *Stack {
	return &Stack{
		Container: Container{
			Widget: Widget{
				InitiallyUnowned: externglib.InitiallyUnowned{
					Object: obj,
				},
				ImplementorIface: atk.ImplementorIface{
					Object: obj,
				},
				Buildable: Buildable{
					Object: obj,
				},
			},
		},
	}
}

func marshalStacker(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapStack(obj), nil
}

// NewStack creates a new Stack container.
func NewStack() *Stack {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_stack_new()

	var _stack *Stack // out

	_stack = wrapStack(externglib.Take(unsafe.Pointer(_cret)))

	return _stack
}

// AddNamed adds a child to stack. The child is identified by the name.
func (stack *Stack) AddNamed(child Widgetter, name string) {
	var _arg0 *C.GtkStack  // out
	var _arg1 *C.GtkWidget // out
	var _arg2 *C.gchar     // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(name)))

	C.gtk_stack_add_named(_arg0, _arg1, _arg2)
}

// AddTitled adds a child to stack. The child is identified by the name. The
// title will be used by StackSwitcher to represent child in a tab bar, so it
// should be short.
func (stack *Stack) AddTitled(child Widgetter, name string, title string) {
	var _arg0 *C.GtkStack  // out
	var _arg1 *C.GtkWidget // out
	var _arg2 *C.gchar     // out
	var _arg3 *C.gchar     // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(title)))

	C.gtk_stack_add_titled(_arg0, _arg1, _arg2, _arg3)
}

// ChildByName finds the child of the Stack with the name given as the argument.
// Returns NULL if there is no child with this name.
func (stack *Stack) ChildByName(name string) *Widget {
	var _arg0 *C.GtkStack  // out
	var _arg1 *C.gchar     // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))

	_cret = C.gtk_stack_get_child_by_name(_arg0, _arg1)

	var _widget *Widget // out

	_widget = wrapWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _widget
}

// Hhomogeneous gets whether stack is horizontally homogeneous. See
// gtk_stack_set_hhomogeneous().
func (stack *Stack) Hhomogeneous() bool {
	var _arg0 *C.GtkStack // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	_cret = C.gtk_stack_get_hhomogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Homogeneous gets whether stack is homogeneous. See
// gtk_stack_set_homogeneous().
func (stack *Stack) Homogeneous() bool {
	var _arg0 *C.GtkStack // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	_cret = C.gtk_stack_get_homogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InterpolateSize returns wether the Stack is set up to interpolate between the
// sizes of children on page switch.
func (stack *Stack) InterpolateSize() bool {
	var _arg0 *C.GtkStack // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	_cret = C.gtk_stack_get_interpolate_size(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TransitionDuration returns the amount of time (in milliseconds) that
// transitions between pages in stack will take.
func (stack *Stack) TransitionDuration() uint {
	var _arg0 *C.GtkStack // out
	var _cret C.guint     // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	_cret = C.gtk_stack_get_transition_duration(_arg0)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// TransitionRunning returns whether the stack is currently in a transition from
// one page to another.
func (stack *Stack) TransitionRunning() bool {
	var _arg0 *C.GtkStack // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	_cret = C.gtk_stack_get_transition_running(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TransitionType gets the type of animation that will be used for transitions
// between pages in stack.
func (stack *Stack) TransitionType() StackTransitionType {
	var _arg0 *C.GtkStack              // out
	var _cret C.GtkStackTransitionType // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	_cret = C.gtk_stack_get_transition_type(_arg0)

	var _stackTransitionType StackTransitionType // out

	_stackTransitionType = StackTransitionType(_cret)

	return _stackTransitionType
}

// Vhomogeneous gets whether stack is vertically homogeneous. See
// gtk_stack_set_vhomogeneous().
func (stack *Stack) Vhomogeneous() bool {
	var _arg0 *C.GtkStack // out
	var _cret C.gboolean  // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	_cret = C.gtk_stack_get_vhomogeneous(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// VisibleChild gets the currently visible child of stack, or NULL if there are
// no visible children.
func (stack *Stack) VisibleChild() *Widget {
	var _arg0 *C.GtkStack  // out
	var _cret *C.GtkWidget // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	_cret = C.gtk_stack_get_visible_child(_arg0)

	var _widget *Widget // out

	_widget = wrapWidget(externglib.Take(unsafe.Pointer(_cret)))

	return _widget
}

// VisibleChildName returns the name of the currently visible child of stack, or
// NULL if there is no visible child.
func (stack *Stack) VisibleChildName() string {
	var _arg0 *C.GtkStack // out
	var _cret *C.gchar    // in

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))

	_cret = C.gtk_stack_get_visible_child_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// SetHhomogeneous sets the Stack to be horizontally homogeneous or not. If it
// is homogeneous, the Stack will request the same width for all its children.
// If it isn't, the stack may change width when a different child becomes
// visible.
func (stack *Stack) SetHhomogeneous(hhomogeneous bool) {
	var _arg0 *C.GtkStack // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	if hhomogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_stack_set_hhomogeneous(_arg0, _arg1)
}

// SetHomogeneous sets the Stack to be homogeneous or not. If it is homogeneous,
// the Stack will request the same size for all its children. If it isn't, the
// stack may change size when a different child becomes visible.
//
// Since 3.16, homogeneity can be controlled separately for horizontal and
// vertical size, with the Stack:hhomogeneous and Stack:vhomogeneous.
func (stack *Stack) SetHomogeneous(homogeneous bool) {
	var _arg0 *C.GtkStack // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	if homogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_stack_set_homogeneous(_arg0, _arg1)
}

// SetInterpolateSize sets whether or not stack will interpolate its size when
// changing the visible child. If the Stack:interpolate-size property is set to
// TRUE, stack will interpolate its size between the current one and the one
// it'll take after changing the visible child, according to the set transition
// duration.
func (stack *Stack) SetInterpolateSize(interpolateSize bool) {
	var _arg0 *C.GtkStack // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	if interpolateSize {
		_arg1 = C.TRUE
	}

	C.gtk_stack_set_interpolate_size(_arg0, _arg1)
}

// SetTransitionDuration sets the duration that transitions between pages in
// stack will take.
func (stack *Stack) SetTransitionDuration(duration uint) {
	var _arg0 *C.GtkStack // out
	var _arg1 C.guint     // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	_arg1 = C.guint(duration)

	C.gtk_stack_set_transition_duration(_arg0, _arg1)
}

// SetTransitionType sets the type of animation that will be used for
// transitions between pages in stack. Available types include various kinds of
// fades and slides.
//
// The transition type can be changed without problems at runtime, so it is
// possible to change the animation based on the page that is about to become
// current.
func (stack *Stack) SetTransitionType(transition StackTransitionType) {
	var _arg0 *C.GtkStack              // out
	var _arg1 C.GtkStackTransitionType // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	_arg1 = C.GtkStackTransitionType(transition)

	C.gtk_stack_set_transition_type(_arg0, _arg1)
}

// SetVhomogeneous sets the Stack to be vertically homogeneous or not. If it is
// homogeneous, the Stack will request the same height for all its children. If
// it isn't, the stack may change height when a different child becomes visible.
func (stack *Stack) SetVhomogeneous(vhomogeneous bool) {
	var _arg0 *C.GtkStack // out
	var _arg1 C.gboolean  // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	if vhomogeneous {
		_arg1 = C.TRUE
	}

	C.gtk_stack_set_vhomogeneous(_arg0, _arg1)
}

// SetVisibleChild makes child the visible child of stack.
//
// If child is different from the currently visible child, the transition
// between the two will be animated with the current transition type of stack.
//
// Note that the child widget has to be visible itself (see gtk_widget_show())
// in order to become the visible child of stack.
func (stack *Stack) SetVisibleChild(child Widgetter) {
	var _arg0 *C.GtkStack  // out
	var _arg1 *C.GtkWidget // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer((child).(gextras.Nativer).Native()))

	C.gtk_stack_set_visible_child(_arg0, _arg1)
}

// SetVisibleChildFull makes the child with the given name visible.
//
// Note that the child widget has to be visible itself (see gtk_widget_show())
// in order to become the visible child of stack.
func (stack *Stack) SetVisibleChildFull(name string, transition StackTransitionType) {
	var _arg0 *C.GtkStack              // out
	var _arg1 *C.gchar                 // out
	var _arg2 C.GtkStackTransitionType // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	_arg2 = C.GtkStackTransitionType(transition)

	C.gtk_stack_set_visible_child_full(_arg0, _arg1, _arg2)
}

// SetVisibleChildName makes the child with the given name visible.
//
// If child is different from the currently visible child, the transition
// between the two will be animated with the current transition type of stack.
//
// Note that the child widget has to be visible itself (see gtk_widget_show())
// in order to become the visible child of stack.
func (stack *Stack) SetVisibleChildName(name string) {
	var _arg0 *C.GtkStack // out
	var _arg1 *C.gchar    // out

	_arg0 = (*C.GtkStack)(unsafe.Pointer(stack.Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))

	C.gtk_stack_set_visible_child_name(_arg0, _arg1)
}
