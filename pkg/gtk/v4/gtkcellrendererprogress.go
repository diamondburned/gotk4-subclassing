// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_cell_renderer_progress_get_type()), F: marshalCellRendererProgress},
	})
}

// CellRendererProgress renders numbers as progress bars
//
// CellRendererProgress renders a numeric value as a progress par in a cell.
// Additionally, it can display a text on top of the progress bar.
type CellRendererProgress interface {
	gextras.Objector

	privateCellRendererProgressClass()
}

// CellRendererProgressClass implements the CellRendererProgress interface.
type CellRendererProgressClass struct {
	*externglib.Object
	CellRendererClass
	OrientableInterface
}

var _ CellRendererProgress = (*CellRendererProgressClass)(nil)

func wrapCellRendererProgress(obj *externglib.Object) CellRendererProgress {
	return &CellRendererProgressClass{
		Object: obj,
		CellRendererClass: CellRendererClass{
			InitiallyUnowned: externglib.InitiallyUnowned{Object: obj},
		},
		OrientableInterface: OrientableInterface{
			Object: obj,
		},
	}
}

func marshalCellRendererProgress(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapCellRendererProgress(obj), nil
}

// NewCellRendererProgress creates a new CellRendererProgress.
func NewCellRendererProgress() *CellRendererProgressClass {
	var _cret *C.GtkCellRenderer // in

	_cret = C.gtk_cell_renderer_progress_new()

	var _cellRendererProgress *CellRendererProgressClass // out

	_cellRendererProgress = (gextras.CastObject(externglib.Take(unsafe.Pointer(_cret)))).(*CellRendererProgressClass)

	return _cellRendererProgress
}

func (*CellRendererProgressClass) privateCellRendererProgressClass() {}
