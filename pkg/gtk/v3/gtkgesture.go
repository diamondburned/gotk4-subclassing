// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v3"
)

// #cgo pkg-config: gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gtk/gtk-a11y.h>
// #include <gtk/gtk.h>
// #include <gtk/gtkx.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_gesture_get_type()), F: marshalGesturer},
	})
}

// Gesture is the base object for gesture recognition, although this object is
// quite generalized to serve as a base for multi-touch gestures, it is suitable
// to implement single-touch and pointer-based gestures (using the special NULL
// EventSequence value for these).
//
// The number of touches that a Gesture need to be recognized is controlled by
// the Gesture:n-points property, if a gesture is keeping track of less or more
// than that number of sequences, it won't check wether the gesture is
// recognized.
//
// As soon as the gesture has the expected number of touches, the gesture will
// run the Gesture::check signal regularly on input events until the gesture is
// recognized, the criteria to consider a gesture as "recognized" is left to
// Gesture subclasses.
//
// A recognized gesture will then emit the following signals:
//
// - Gesture::begin when the gesture is recognized.
//
// - A number of Gesture::update, whenever an input event is processed.
//
// - Gesture::end when the gesture is no longer recognized.
//
//
// Event propagation
//
// In order to receive events, a gesture needs to either set a propagation phase
// through gtk_event_controller_set_propagation_phase(), or feed those manually
// through gtk_event_controller_handle_event().
//
// In the capture phase, events are propagated from the toplevel down to the
// target widget, and gestures that are attached to containers above the widget
// get a chance to interact with the event before it reaches the target.
//
// After the capture phase, GTK+ emits the traditional
// Widget::button-press-event, Widget::button-release-event,
// Widget::touch-event, etc signals. Gestures with the GTK_PHASE_TARGET phase
// are fed events from the default Widget::event handlers.
//
// In the bubble phase, events are propagated up from the target widget to the
// toplevel, and gestures that are attached to containers above the widget get a
// chance to interact with events that have not been handled yet.
//
//
// States of a sequence
//
// Whenever input interaction happens, a single event may trigger a cascade of
// Gestures, both across the parents of the widget receiving the event and in
// parallel within an individual widget. It is a responsibility of the widgets
// using those gestures to set the state of touch sequences accordingly in order
// to enable cooperation of gestures around the EventSequences triggering those.
//
// Within a widget, gestures can be grouped through gtk_gesture_group(), grouped
// gestures synchronize the state of sequences, so calling
// gtk_gesture_set_sequence_state() on one will effectively propagate the state
// throughout the group.
//
// By default, all sequences start out in the K_EVENT_SEQUENCE_NONE state,
// sequences in this state trigger the gesture event handler, but event
// propagation will continue unstopped by gestures.
//
// If a sequence enters into the K_EVENT_SEQUENCE_DENIED state, the gesture
// group will effectively ignore the sequence, letting events go unstopped
// through the gesture, but the "slot" will still remain occupied while the
// touch is active.
//
// If a sequence enters in the K_EVENT_SEQUENCE_CLAIMED state, the gesture group
// will grab all interaction on the sequence, by:
//
// - Setting the same sequence to K_EVENT_SEQUENCE_DENIED on every other gesture
// group within the widget, and every gesture on parent widgets in the
// propagation chain.
//
// - calling Gesture::cancel on every gesture in widgets underneath in the
// propagation chain.
//
// - Stopping event propagation after the gesture group handles the event.
//
// Note: if a sequence is set early to K_EVENT_SEQUENCE_CLAIMED on
// K_TOUCH_BEGIN/K_BUTTON_PRESS (so those events are captured before reaching
// the event widget, this implies K_PHASE_CAPTURE), one similar event will
// emulated if the sequence changes to K_EVENT_SEQUENCE_DENIED. This way event
// coherence is preserved before event propagation is unstopped again.
//
// Sequence states can't be changed freely, see gtk_gesture_set_sequence_state()
// to know about the possible lifetimes of a EventSequence.
//
//
// Touchpad gestures
//
// On the platforms that support it, Gesture will handle transparently touchpad
// gesture events. The only precautions users of Gesture should do to enable
// this support are:
//
// - Enabling GDK_TOUCHPAD_GESTURE_MASK on their Windows
//
// - If the gesture has GTK_PHASE_NONE, ensuring events of type
// GDK_TOUCHPAD_SWIPE and GDK_TOUCHPAD_PINCH are handled by the Gesture.
type Gesture struct {
	EventController
}

// Gesturer describes types inherited from class Gesture.
// To get the original type, the caller must assert this to an interface or
// another type.
type Gesturer interface {
	externglib.Objector

	// BaseGesture returns the underlying base class.
	BaseGesture() *Gesture
}

var _ Gesturer = (*Gesture)(nil)

func wrapGesture(obj *externglib.Object) *Gesture {
	return &Gesture{
		EventController: EventController{
			Object: obj,
		},
	}
}

func marshalGesturer(p uintptr) (interface{}, error) {
	return wrapGesture(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// BoundingBox: if there are touch sequences being currently handled by gesture,
// this function returns TRUE and fills in rect with the bounding box containing
// all active touches. Otherwise, FALSE will be returned.
//
// Note: This function will yield unexpected results on touchpad gestures. Since
// there is no correlation between physical and pixel distances, these will look
// as if constrained in an infinitely small area, rect width and height will
// thus be 0 regardless of the number of touchpoints.
func (gesture *Gesture) BoundingBox() (*gdk.Rectangle, bool) {
	var _arg0 *C.GtkGesture  // out
	var _arg1 C.GdkRectangle // in
	var _cret C.gboolean     // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_get_bounding_box(_arg0, &_arg1)
	runtime.KeepAlive(gesture)

	var _rect *gdk.Rectangle // out
	var _ok bool             // out

	_rect = (*gdk.Rectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	if _cret != 0 {
		_ok = true
	}

	return _rect, _ok
}

// BoundingBoxCenter: if there are touch sequences being currently handled by
// gesture, this function returns TRUE and fills in x and y with the center of
// the bounding box containing all active touches. Otherwise, FALSE will be
// returned.
func (gesture *Gesture) BoundingBoxCenter() (x float64, y float64, ok bool) {
	var _arg0 *C.GtkGesture // out
	var _arg1 C.gdouble     // in
	var _arg2 C.gdouble     // in
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_get_bounding_box_center(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(gesture)

	var _x float64 // out
	var _y float64 // out
	var _ok bool   // out

	_x = float64(_arg1)
	_y = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _x, _y, _ok
}

// Device returns the master Device that is currently operating on gesture, or
// NULL if the gesture is not being interacted.
func (gesture *Gesture) Device() gdk.Devicer {
	var _arg0 *C.GtkGesture // out
	var _cret *C.GdkDevice  // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_get_device(_arg0)
	runtime.KeepAlive(gesture)

	var _device gdk.Devicer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(gdk.Devicer)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Devicer")
			}
			_device = rv
		}
	}

	return _device
}

// GetGroup returns all gestures in the group of gesture.
func (gesture *Gesture) GetGroup() []Gesturer {
	var _arg0 *C.GtkGesture // out
	var _cret *C.GList      // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_get_group(_arg0)
	runtime.KeepAlive(gesture)

	var _list []Gesturer // out

	_list = make([]Gesturer, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GtkGesture)(v)
		var dst Gesturer // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type gtk.Gesturer is nil")
			}

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(Gesturer)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Gesturer")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})

	return _list
}

// LastUpdatedSequence returns the EventSequence that was last updated on
// gesture.
func (gesture *Gesture) LastUpdatedSequence() *gdk.EventSequence {
	var _arg0 *C.GtkGesture       // out
	var _cret *C.GdkEventSequence // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_get_last_updated_sequence(_arg0)
	runtime.KeepAlive(gesture)

	var _eventSequence *gdk.EventSequence // out

	if _cret != nil {
		_eventSequence = (*gdk.EventSequence)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _eventSequence
}

// Point: if sequence is currently being interpreted by gesture, this function
// returns TRUE and fills in x and y with the last coordinates stored for that
// event sequence. The coordinates are always relative to the widget allocation.
//
// The function takes the following parameters:
//
//    - sequence or NULL for pointer events.
//
func (gesture *Gesture) Point(sequence *gdk.EventSequence) (x float64, y float64, ok bool) {
	var _arg0 *C.GtkGesture       // out
	var _arg1 *C.GdkEventSequence // out
	var _arg2 C.gdouble           // in
	var _arg3 C.gdouble           // in
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))
	if sequence != nil {
		_arg1 = (*C.GdkEventSequence)(gextras.StructNative(unsafe.Pointer(sequence)))
	}

	_cret = C.gtk_gesture_get_point(_arg0, _arg1, &_arg2, &_arg3)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(sequence)

	var _x float64 // out
	var _y float64 // out
	var _ok bool   // out

	_x = float64(_arg2)
	_y = float64(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _x, _y, _ok
}

// SequenceState returns the sequence state, as seen by gesture.
//
// The function takes the following parameters:
//
//    - sequence: EventSequence.
//
func (gesture *Gesture) SequenceState(sequence *gdk.EventSequence) EventSequenceState {
	var _arg0 *C.GtkGesture           // out
	var _arg1 *C.GdkEventSequence     // out
	var _cret C.GtkEventSequenceState // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))
	_arg1 = (*C.GdkEventSequence)(gextras.StructNative(unsafe.Pointer(sequence)))

	_cret = C.gtk_gesture_get_sequence_state(_arg0, _arg1)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(sequence)

	var _eventSequenceState EventSequenceState // out

	_eventSequenceState = EventSequenceState(_cret)

	return _eventSequenceState
}

// Sequences returns the list of EventSequences currently being interpreted by
// gesture.
func (gesture *Gesture) Sequences() []*gdk.EventSequence {
	var _arg0 *C.GtkGesture // out
	var _cret *C.GList      // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_get_sequences(_arg0)
	runtime.KeepAlive(gesture)

	var _list []*gdk.EventSequence // out

	_list = make([]*gdk.EventSequence, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GdkEventSequence)(v)
		var dst *gdk.EventSequence // out
		dst = (*gdk.EventSequence)(gextras.NewStructNative(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}

// Window returns the user-defined window that receives the events handled by
// gesture. See gtk_gesture_set_window() for more information.
func (gesture *Gesture) Window() gdk.Windower {
	var _arg0 *C.GtkGesture // out
	var _cret *C.GdkWindow  // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_get_window(_arg0)
	runtime.KeepAlive(gesture)

	var _window gdk.Windower // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(gdk.Windower)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gdk.Windower")
			}
			_window = rv
		}
	}

	return _window
}

// Group adds gesture to the same group than group_gesture. Gestures are by
// default isolated in their own groups.
//
// When gestures are grouped, the state of EventSequences is kept in sync for
// all of those, so calling gtk_gesture_set_sequence_state(), on one will
// transfer the same value to the others.
//
// Groups also perform an "implicit grabbing" of sequences, if a EventSequence
// state is set to K_EVENT_SEQUENCE_CLAIMED on one group, every other gesture
// group attached to the same Widget will switch the state for that sequence to
// K_EVENT_SEQUENCE_DENIED.
//
// The function takes the following parameters:
//
//    - gesture: Gesture.
//
func (groupGesture *Gesture) Group(gesture Gesturer) {
	var _arg0 *C.GtkGesture // out
	var _arg1 *C.GtkGesture // out

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(groupGesture.Native()))
	_arg1 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	C.gtk_gesture_group(_arg0, _arg1)
	runtime.KeepAlive(groupGesture)
	runtime.KeepAlive(gesture)
}

// HandlesSequence returns TRUE if gesture is currently handling events
// corresponding to sequence.
//
// The function takes the following parameters:
//
//    - sequence or NULL.
//
func (gesture *Gesture) HandlesSequence(sequence *gdk.EventSequence) bool {
	var _arg0 *C.GtkGesture       // out
	var _arg1 *C.GdkEventSequence // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))
	if sequence != nil {
		_arg1 = (*C.GdkEventSequence)(gextras.StructNative(unsafe.Pointer(sequence)))
	}

	_cret = C.gtk_gesture_handles_sequence(_arg0, _arg1)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(sequence)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsActive returns TRUE if the gesture is currently active. A gesture is active
// meanwhile there are touch sequences interacting with it.
func (gesture *Gesture) IsActive() bool {
	var _arg0 *C.GtkGesture // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_is_active(_arg0)
	runtime.KeepAlive(gesture)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsGroupedWith returns TRUE if both gestures pertain to the same group.
//
// The function takes the following parameters:
//
//    - other Gesture.
//
func (gesture *Gesture) IsGroupedWith(other Gesturer) bool {
	var _arg0 *C.GtkGesture // out
	var _arg1 *C.GtkGesture // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))
	_arg1 = (*C.GtkGesture)(unsafe.Pointer(other.Native()))

	_cret = C.gtk_gesture_is_grouped_with(_arg0, _arg1)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(other)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsRecognized returns TRUE if the gesture is currently recognized. A gesture
// is recognized if there are as many interacting touch sequences as required by
// gesture, and Gesture::check returned TRUE for the sequences being currently
// interpreted.
func (gesture *Gesture) IsRecognized() bool {
	var _arg0 *C.GtkGesture // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	_cret = C.gtk_gesture_is_recognized(_arg0)
	runtime.KeepAlive(gesture)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetSequenceState sets the state of sequence in gesture. Sequences start in
// state K_EVENT_SEQUENCE_NONE, and whenever they change state, they can never
// go back to that state. Likewise, sequences in state K_EVENT_SEQUENCE_DENIED
// cannot turn back to a not denied state. With these rules, the lifetime of an
// event sequence is constrained to the next four:
//
// * None * None → Denied * None → Claimed * None → Claimed → Denied
//
// Note: Due to event handling ordering, it may be unsafe to set the state on
// another gesture within a Gesture::begin signal handler, as the callback might
// be executed before the other gesture knows about the sequence. A safe way to
// perform this could be:
//
//    static void
//    first_gesture_begin_cb (GtkGesture       *first_gesture,
//                            GdkEventSequence *sequence,
//                            gpointer          user_data)
//    {
//      gtk_gesture_set_sequence_state (first_gesture, sequence, GTK_EVENT_SEQUENCE_CLAIMED);
//      gtk_gesture_set_sequence_state (second_gesture, sequence, GTK_EVENT_SEQUENCE_DENIED);
//    }
//
//    static void
//    second_gesture_begin_cb (GtkGesture       *second_gesture,
//                             GdkEventSequence *sequence,
//                             gpointer          user_data)
//    {
//      if (gtk_gesture_get_sequence_state (first_gesture, sequence) == GTK_EVENT_SEQUENCE_CLAIMED)
//        gtk_gesture_set_sequence_state (second_gesture, sequence, GTK_EVENT_SEQUENCE_DENIED);
//    }
//
// If both gestures are in the same group, just set the state on the gesture
// emitting the event, the sequence will be already be initialized to the
// group's global state when the second gesture processes the event.
//
// The function takes the following parameters:
//
//    - sequence: EventSequence.
//    - state: sequence state.
//
func (gesture *Gesture) SetSequenceState(sequence *gdk.EventSequence, state EventSequenceState) bool {
	var _arg0 *C.GtkGesture           // out
	var _arg1 *C.GdkEventSequence     // out
	var _arg2 C.GtkEventSequenceState // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))
	_arg1 = (*C.GdkEventSequence)(gextras.StructNative(unsafe.Pointer(sequence)))
	_arg2 = C.GtkEventSequenceState(state)

	_cret = C.gtk_gesture_set_sequence_state(_arg0, _arg1, _arg2)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(sequence)
	runtime.KeepAlive(state)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetState sets the state of all sequences that gesture is currently
// interacting with. See gtk_gesture_set_sequence_state() for more details on
// sequence states.
//
// The function takes the following parameters:
//
//    - state: sequence state.
//
func (gesture *Gesture) SetState(state EventSequenceState) bool {
	var _arg0 *C.GtkGesture           // out
	var _arg1 C.GtkEventSequenceState // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))
	_arg1 = C.GtkEventSequenceState(state)

	_cret = C.gtk_gesture_set_state(_arg0, _arg1)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(state)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetWindow sets a specific window to receive events about, so gesture will
// effectively handle only events targeting window, or a child of it. window
// must pertain to gtk_event_controller_get_widget().
//
// The function takes the following parameters:
//
//    - window or NULL.
//
func (gesture *Gesture) SetWindow(window gdk.Windower) {
	var _arg0 *C.GtkGesture // out
	var _arg1 *C.GdkWindow  // out

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))
	if window != nil {
		_arg1 = (*C.GdkWindow)(unsafe.Pointer(window.Native()))
	}

	C.gtk_gesture_set_window(_arg0, _arg1)
	runtime.KeepAlive(gesture)
	runtime.KeepAlive(window)
}

// Ungroup separates gesture into an isolated group.
func (gesture *Gesture) Ungroup() {
	var _arg0 *C.GtkGesture // out

	_arg0 = (*C.GtkGesture)(unsafe.Pointer(gesture.Native()))

	C.gtk_gesture_ungroup(_arg0)
	runtime.KeepAlive(gesture)
}

// BaseGesture returns gesture.
func (gesture *Gesture) BaseGesture() *Gesture {
	return gesture
}

// ConnectBegin: this signal is emitted when the gesture is recognized. This
// means the number of touch sequences matches Gesture:n-points, and the
// Gesture::check handler(s) returned UE.
//
// Note: These conditions may also happen when an extra touch (eg. a third touch
// on a 2-touches gesture) is lifted, in that situation sequence won't pertain
// to the current set of active touches, so don't rely on this being true.
func (gesture *Gesture) ConnectBegin(f func(sequence *gdk.EventSequence)) externglib.SignalHandle {
	return gesture.Connect("begin", f)
}

// ConnectCancel: this signal is emitted whenever a sequence is cancelled. This
// usually happens on active touches when gtk_event_controller_reset() is called
// on gesture (manually, due to grabs...), or the individual sequence was
// claimed by parent widgets' controllers (see
// gtk_gesture_set_sequence_state()).
//
// gesture must forget everything about sequence as a reaction to this signal.
func (gesture *Gesture) ConnectCancel(f func(sequence *gdk.EventSequence)) externglib.SignalHandle {
	return gesture.Connect("cancel", f)
}

// ConnectEnd: this signal is emitted when gesture either stopped recognizing
// the event sequences as something to be handled (the Gesture::check handler
// returned FALSE), or the number of touch sequences became higher or lower than
// Gesture:n-points.
//
// Note: sequence might not pertain to the group of sequences that were
// previously triggering recognition on gesture (ie. a just pressed touch
// sequence that exceeds Gesture:n-points). This situation may be detected by
// checking through gtk_gesture_handles_sequence().
func (gesture *Gesture) ConnectEnd(f func(sequence *gdk.EventSequence)) externglib.SignalHandle {
	return gesture.Connect("end", f)
}

// ConnectSequenceStateChanged: this signal is emitted whenever a sequence state
// changes. See gtk_gesture_set_sequence_state() to know more about the
// expectable sequence lifetimes.
func (gesture *Gesture) ConnectSequenceStateChanged(f func(sequence *gdk.EventSequence, state EventSequenceState)) externglib.SignalHandle {
	return gesture.Connect("sequence-state-changed", f)
}

// ConnectUpdate: this signal is emitted whenever an event is handled while the
// gesture is recognized. sequence is guaranteed to pertain to the set of active
// touches.
func (gesture *Gesture) ConnectUpdate(f func(sequence *gdk.EventSequence)) externglib.SignalHandle {
	return gesture.Connect("update", f)
}
