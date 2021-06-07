// Code generated by girgen. DO NOT EDIT.

package gtk

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk.h>
import "C"

// TestListAllTypes: return the type ids that have been registered after calling
// gtk_test_register_all_types().
func TestListAllTypes() uint {
	var arg1 C.guint
	var nTypes uint

	C.gtk_test_list_all_types(&arg1)

	nTypes = uint(&arg1)

	return nTypes
}

// TestRegisterAllTypes: force registration of all core GTK object types.
//
// This allowes to refer to any of those object types via g_type_from_name()
// after calling this function.
func TestRegisterAllTypes() {
	C.gtk_test_register_all_types()
}

// TestWidgetWaitForDraw enters the main loop and waits for @widget to be
// “drawn”. In this context that means it waits for the frame clock of @widget
// to have run a full styling, layout and drawing cycle.
//
// This function is intended to be used for syncing with actions that depend on
// @widget relayouting or on interaction with the display server.
func TestWidgetWaitForDraw(widget Widget) {
	var arg1 *C.GtkWidget

	arg1 = (*C.GtkWidget)(unsafe.Pointer(widget.Native()))

	C.gtk_test_widget_wait_for_draw(arg1)
}