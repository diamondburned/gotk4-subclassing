// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/atk"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
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
		{T: externglib.Type(C.gtk_separator_tool_item_get_type()), F: marshalSeparatorToolItemmer},
	})
}

// SeparatorToolItem is a ToolItem that separates groups of other ToolItems.
// Depending on the theme, a SeparatorToolItem will often look like a vertical
// line on horizontally docked toolbars.
//
// If the Toolbar child property “expand” is TRUE and the property
// SeparatorToolItem:draw is FALSE, a SeparatorToolItem will act as a “spring”
// that forces other items to the ends of the toolbar.
//
// Use gtk_separator_tool_item_new() to create a new SeparatorToolItem.
//
//
// CSS nodes
//
// GtkSeparatorToolItem has a single CSS node with name separator.
type SeparatorToolItem struct {
	ToolItem
}

func wrapSeparatorToolItem(obj *externglib.Object) *SeparatorToolItem {
	return &SeparatorToolItem{
		ToolItem: ToolItem{
			Bin: Bin{
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
						Object: obj,
					},
				},
			},
			Activatable: Activatable{
				Object: obj,
			},
			Object: obj,
		},
	}
}

func marshalSeparatorToolItemmer(p uintptr) (interface{}, error) {
	return wrapSeparatorToolItem(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewSeparatorToolItem: create a new SeparatorToolItem.
func NewSeparatorToolItem() *SeparatorToolItem {
	var _cret *C.GtkToolItem // in

	_cret = C.gtk_separator_tool_item_new()

	var _separatorToolItem *SeparatorToolItem // out

	_separatorToolItem = wrapSeparatorToolItem(externglib.Take(unsafe.Pointer(_cret)))

	return _separatorToolItem
}

// Draw returns whether item is drawn as a line, or just blank. See
// gtk_separator_tool_item_set_draw().
func (item *SeparatorToolItem) Draw() bool {
	var _arg0 *C.GtkSeparatorToolItem // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GtkSeparatorToolItem)(unsafe.Pointer(item.Native()))

	_cret = C.gtk_separator_tool_item_get_draw(_arg0)
	runtime.KeepAlive(item)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetDraw: whether item is drawn as a vertical line, or just blank. Setting
// this to FALSE along with gtk_tool_item_set_expand() is useful to create an
// item that forces following items to the end of the toolbar.
//
// The function takes the following parameters:
//
//    - draw: whether item is drawn as a vertical line.
//
func (item *SeparatorToolItem) SetDraw(draw bool) {
	var _arg0 *C.GtkSeparatorToolItem // out
	var _arg1 C.gboolean              // out

	_arg0 = (*C.GtkSeparatorToolItem)(unsafe.Pointer(item.Native()))
	if draw {
		_arg1 = C.TRUE
	}

	C.gtk_separator_tool_item_set_draw(_arg0, _arg1)
	runtime.KeepAlive(item)
	runtime.KeepAlive(draw)
}
