// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"unsafe"

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
		{T: externglib.Type(C.gtk_print_operation_preview_get_type()), F: marshalPrintOperationPreview},
	})
}

// PrintOperationPreviewOverrider contains methods that are overridable .
//
// As of right now, interface overriding and subclassing is not supported
// yet, so the interface currently has no use.
type PrintOperationPreviewOverrider interface {
	// EndPreview ends a preview.
	//
	// This function must be called to finish a custom print preview.
	EndPreview()
	GotPageSize(context PrintContext, pageSetup PageSetup)
	// IsSelected returns whether the given page is included in the set of pages
	// that have been selected for printing.
	IsSelected(pageNr int) bool
	Ready(context PrintContext)
	// RenderPage renders a page to the preview.
	//
	// This is using the print context that was passed to the
	// [signal@Gtk.PrintOperation::preview] handler together with @preview.
	//
	// A custom print preview should use this function to render the currently
	// selected page.
	//
	// Note that this function requires a suitable cairo context to be
	// associated with the print context.
	RenderPage(pageNr int)
}

// PrintOperationPreview: `GtkPrintOperationPreview` is the interface that is
// used to implement print preview.
//
// A `GtkPrintOperationPreview` object is passed to the
// [signal@Gtk.PrintOperation::preview] signal by [class@Gtk.PrintOperation].
type PrintOperationPreview interface {
	gextras.Objector

	// EndPreview ends a preview.
	//
	// This function must be called to finish a custom print preview.
	EndPreview()
	// IsSelected returns whether the given page is included in the set of pages
	// that have been selected for printing.
	IsSelected(pageNr int) bool
	// RenderPage renders a page to the preview.
	//
	// This is using the print context that was passed to the
	// [signal@Gtk.PrintOperation::preview] handler together with @preview.
	//
	// A custom print preview should use this function to render the currently
	// selected page.
	//
	// Note that this function requires a suitable cairo context to be
	// associated with the print context.
	RenderPage(pageNr int)
}

// printOperationPreview implements the PrintOperationPreview interface.
type printOperationPreview struct {
	*externglib.Object
}

var _ PrintOperationPreview = (*printOperationPreview)(nil)

// WrapPrintOperationPreview wraps a GObject to a type that implements
// interface PrintOperationPreview. It is primarily used internally.
func WrapPrintOperationPreview(obj *externglib.Object) PrintOperationPreview {
	return printOperationPreview{obj}
}

func marshalPrintOperationPreview(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapPrintOperationPreview(obj), nil
}

func (p printOperationPreview) EndPreview() {
	var _arg0 *C.GtkPrintOperationPreview // out

	_arg0 = (*C.GtkPrintOperationPreview)(unsafe.Pointer(p.Native()))

	C.gtk_print_operation_preview_end_preview(_arg0)
}

func (p printOperationPreview) IsSelected(pageNr int) bool {
	var _arg0 *C.GtkPrintOperationPreview // out
	var _arg1 C.int                       // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.GtkPrintOperationPreview)(unsafe.Pointer(p.Native()))
	_arg1 = C.int(pageNr)

	_cret = C.gtk_print_operation_preview_is_selected(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (p printOperationPreview) RenderPage(pageNr int) {
	var _arg0 *C.GtkPrintOperationPreview // out
	var _arg1 C.int                       // out

	_arg0 = (*C.GtkPrintOperationPreview)(unsafe.Pointer(p.Native()))
	_arg1 = C.int(pageNr)

	C.gtk_print_operation_preview_render_page(_arg0, _arg1)
}
