// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
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
		{T: externglib.Type(C.gtk_overlay_get_type()), F: marshalOverlay},
	})
}

// Overlay is a container which contains a single main child, on top of which it
// can place “overlay” widgets. The position of each overlay widget is
// determined by its Widget:halign and Widget:valign properties. E.g. a widget
// with both alignments set to GTK_ALIGN_START will be placed at the top left
// corner of the GtkOverlay container, whereas an overlay with halign set to
// GTK_ALIGN_CENTER and valign set to GTK_ALIGN_END will be placed a the bottom
// edge of the GtkOverlay, horizontally centered. The position can be adjusted
// by setting the margin properties of the child to non-zero values.
//
// More complicated placement of overlays is possible by connecting to the
// Overlay::get-child-position signal.
//
// An overlay’s minimum and natural sizes are those of its main child. The sizes
// of overlay children are not considered when measuring these preferred sizes.
//
//
// GtkOverlay as GtkBuildable
//
// The GtkOverlay implementation of the GtkBuildable interface supports placing
// a child as an overlay by specifying “overlay” as the “type” attribute of a
// `<child>` element.
//
//
// CSS nodes
//
// GtkOverlay has a single CSS node with the name “overlay”. Overlay children
// whose alignments cause them to be positioned at an edge get the style classes
// “.left”, “.right”, “.top”, and/or “.bottom” according to their position.
type Overlay interface {
	gextras.Objector

	// AddOverlay adds @widget to @overlay.
	//
	// The widget will be stacked on top of the main widget added with
	// gtk_container_add().
	//
	// The position at which @widget is placed is determined from its
	// Widget:halign and Widget:valign properties.
	AddOverlay(widget Widget)
	// OverlayPassThrough: convenience function to get the value of the
	// Overlay:pass-through child property for @widget.
	OverlayPassThrough(widget Widget) bool
	// ReorderOverlay moves @child to a new @index in the list of @overlay
	// children. The list contains overlays in the order that these were added
	// to @overlay by default. See also Overlay:index.
	//
	// A widget’s index in the @overlay children list determines which order the
	// children are drawn if they overlap. The first child is drawn at the
	// bottom. It also affects the default focus chain order.
	ReorderOverlay(child Widget, index_ int)
	// SetOverlayPassThrough: convenience function to set the value of the
	// Overlay:pass-through child property for @widget.
	SetOverlayPassThrough(widget Widget, passThrough bool)
}

// OverlayClass implements the Overlay interface.
type OverlayClass struct {
	*externglib.Object
	BinClass
	BuildableInterface
}

var _ Overlay = (*OverlayClass)(nil)

func wrapOverlay(obj *externglib.Object) Overlay {
	return &OverlayClass{
		Object: obj,
		BinClass: BinClass{
			Object: obj,
			ContainerClass: ContainerClass{
				Object: obj,
				WidgetClass: WidgetClass{
					InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
					BuildableInterface: BuildableInterface{
						Object: obj,
					},
				},
				BuildableInterface: BuildableInterface{
					Object: obj,
				},
			},
			BuildableInterface: BuildableInterface{
				Object: obj,
			},
		},
		BuildableInterface: BuildableInterface{
			Object: obj,
		},
	}
}

func marshalOverlay(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapOverlay(obj), nil
}

// NewOverlay creates a new Overlay.
func NewOverlay() *OverlayClass {
	var _cret *C.GtkWidget // in

	_cret = C.gtk_overlay_new()

	var _overlay *OverlayClass // out

	_overlay = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*OverlayClass)

	return _overlay
}

// AddOverlay adds @widget to @overlay.
//
// The widget will be stacked on top of the main widget added with
// gtk_container_add().
//
// The position at which @widget is placed is determined from its Widget:halign
// and Widget:valign properties.
func (o *OverlayClass) AddOverlay(widget Widget) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_overlay_add_overlay(_arg0, _arg1)
}

// OverlayPassThrough: convenience function to get the value of the
// Overlay:pass-through child property for @widget.
func (o *OverlayClass) OverlayPassThrough(widget Widget) bool {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _cret C.gboolean    // in

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	_cret = C.gtk_overlay_get_overlay_pass_through(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ReorderOverlay moves @child to a new @index in the list of @overlay children.
// The list contains overlays in the order that these were added to @overlay by
// default. See also Overlay:index.
//
// A widget’s index in the @overlay children list determines which order the
// children are drawn if they overlap. The first child is drawn at the bottom.
// It also affects the default focus chain order.
func (o *OverlayClass) ReorderOverlay(child Widget, index_ int) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.int         // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(child.Native()))
	_arg2 = C.int(index_)

	C.gtk_overlay_reorder_overlay(_arg0, _arg1, _arg2)
}

// SetOverlayPassThrough: convenience function to set the value of the
// Overlay:pass-through child property for @widget.
func (o *OverlayClass) SetOverlayPassThrough(widget Widget, passThrough bool) {
	var _arg0 *C.GtkOverlay // out
	var _arg1 *C.GtkWidget  // out
	var _arg2 C.gboolean    // out

	_arg0 = (*C.GtkOverlay)(unsafe.Pointer(o.Native()))
	_arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))
	if passThrough {
		_arg2 = C.TRUE
	}

	C.gtk_overlay_set_overlay_pass_through(_arg0, _arg1, _arg2)
}
