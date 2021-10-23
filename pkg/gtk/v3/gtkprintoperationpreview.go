// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

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
		{T: externglib.Type(C.gtk_print_operation_preview_get_type()), F: marshalPrintOperationPreviewer},
	})
}

// PrintOperationPreviewOverrider contains methods that are overridable.
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type PrintOperationPreviewOverrider interface {
	// EndPreview ends a preview.
	//
	// This function must be called to finish a custom print preview.
	EndPreview()
	GotPageSize(context *PrintContext, pageSetup *PageSetup)
	// IsSelected returns whether the given page is included in the set of pages
	// that have been selected for printing.
	IsSelected(pageNr int) bool
	Ready(context *PrintContext)
	// RenderPage renders a page to the preview, using the print context that
	// was passed to the PrintOperation::preview handler together with preview.
	//
	// A custom iprint preview should use this function in its ::expose handler
	// to render the currently selected page.
	//
	// Note that this function requires a suitable cairo context to be
	// associated with the print context.
	RenderPage(pageNr int)
}

type PrintOperationPreview struct {
	*externglib.Object
}

// PrintOperationPreviewer describes PrintOperationPreview's interface methods.
type PrintOperationPreviewer interface {
	externglib.Objector

	// EndPreview ends a preview.
	EndPreview()
	// IsSelected returns whether the given page is included in the set of pages
	// that have been selected for printing.
	IsSelected(pageNr int) bool
	// RenderPage renders a page to the preview, using the print context that
	// was passed to the PrintOperation::preview handler together with preview.
	RenderPage(pageNr int)
}

var _ PrintOperationPreviewer = (*PrintOperationPreview)(nil)

func wrapPrintOperationPreview(obj *externglib.Object) *PrintOperationPreview {
	return &PrintOperationPreview{
		Object: obj,
	}
}

func marshalPrintOperationPreviewer(p uintptr) (interface{}, error) {
	return wrapPrintOperationPreview(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// EndPreview ends a preview.
//
// This function must be called to finish a custom print preview.
func (preview *PrintOperationPreview) EndPreview() {
	var _arg0 *C.GtkPrintOperationPreview // out

	_arg0 = (*C.GtkPrintOperationPreview)(unsafe.Pointer(preview.Native()))

	C.gtk_print_operation_preview_end_preview(_arg0)
	runtime.KeepAlive(preview)
}

// IsSelected returns whether the given page is included in the set of pages
// that have been selected for printing.
//
// The function takes the following parameters:
//
//    - pageNr: page number.
//
func (preview *PrintOperationPreview) IsSelected(pageNr int) bool {
	var _arg0 *C.GtkPrintOperationPreview // out
	var _arg1 C.gint                      // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.GtkPrintOperationPreview)(unsafe.Pointer(preview.Native()))
	_arg1 = C.gint(pageNr)

	_cret = C.gtk_print_operation_preview_is_selected(_arg0, _arg1)
	runtime.KeepAlive(preview)
	runtime.KeepAlive(pageNr)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RenderPage renders a page to the preview, using the print context that was
// passed to the PrintOperation::preview handler together with preview.
//
// A custom iprint preview should use this function in its ::expose handler to
// render the currently selected page.
//
// Note that this function requires a suitable cairo context to be associated
// with the print context.
//
// The function takes the following parameters:
//
//    - pageNr: page to render.
//
func (preview *PrintOperationPreview) RenderPage(pageNr int) {
	var _arg0 *C.GtkPrintOperationPreview // out
	var _arg1 C.gint                      // out

	_arg0 = (*C.GtkPrintOperationPreview)(unsafe.Pointer(preview.Native()))
	_arg1 = C.gint(pageNr)

	C.gtk_print_operation_preview_render_page(_arg0, _arg1)
	runtime.KeepAlive(preview)
	runtime.KeepAlive(pageNr)
}

// ConnectGotPageSize signal is emitted once for each page that gets rendered to
// the preview.
//
// A handler for this signal should update the context according to page_setup
// and set up a suitable cairo context, using
// gtk_print_context_set_cairo_context().
func (preview *PrintOperationPreview) ConnectGotPageSize(f func(context PrintContext, pageSetup PageSetup)) externglib.SignalHandle {
	return preview.Connect("got-page-size", f)
}

// ConnectReady signal gets emitted once per preview operation, before the first
// page is rendered.
//
// A handler for this signal can be used for setup tasks.
func (preview *PrintOperationPreview) ConnectReady(f func(context PrintContext)) externglib.SignalHandle {
	return preview.Connect("ready", f)
}
