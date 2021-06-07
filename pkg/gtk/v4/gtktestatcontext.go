// Code generated by girgen. DO NOT EDIT.

package gtk

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gtk/gtk.h>
import "C"

func TestAccessibleAssertionMessageRole(domain string, file string, line int, fn string, expr string, accessible Accessible, expectedRole AccessibleRole, actualRole AccessibleRole) {
	var arg1 *C.char
	var arg2 *C.char
	var arg3 C.int
	var arg4 *C.char
	var arg5 *C.char
	var arg6 *C.GtkAccessible
	var arg7 C.GtkAccessibleRole
	var arg8 C.GtkAccessibleRole

	arg1 = (*C.char)(C.CString(domain))
	defer C.free(unsafe.Pointer(arg1))
	arg2 = (*C.char)(C.CString(file))
	defer C.free(unsafe.Pointer(arg2))
	arg3 = C.int(line)
	arg4 = (*C.char)(C.CString(fn))
	defer C.free(unsafe.Pointer(arg4))
	arg5 = (*C.char)(C.CString(expr))
	defer C.free(unsafe.Pointer(arg5))
	arg6 = (*C.GtkAccessible)(unsafe.Pointer(accessible.Native()))
	arg7 = (C.GtkAccessibleRole)(expectedRole)
	arg8 = (C.GtkAccessibleRole)(actualRole)

	C.gtk_test_accessible_assertion_message_role(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
}

// TestAccessibleHasProperty checks whether the Accessible has @property set.
func TestAccessibleHasProperty(accessible Accessible, property AccessibleProperty) bool {
	var arg1 *C.GtkAccessible
	var arg2 C.GtkAccessibleProperty

	arg1 = (*C.GtkAccessible)(unsafe.Pointer(accessible.Native()))
	arg2 = (C.GtkAccessibleProperty)(property)

	var cret C.gboolean
	var ok bool

	cret = C.gtk_test_accessible_has_property(arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// TestAccessibleHasRelation checks whether the Accessible has @relation set.
func TestAccessibleHasRelation(accessible Accessible, relation AccessibleRelation) bool {
	var arg1 *C.GtkAccessible
	var arg2 C.GtkAccessibleRelation

	arg1 = (*C.GtkAccessible)(unsafe.Pointer(accessible.Native()))
	arg2 = (C.GtkAccessibleRelation)(relation)

	var cret C.gboolean
	var ok bool

	cret = C.gtk_test_accessible_has_relation(arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// TestAccessibleHasRole checks whether the Accessible:accessible-role of the
// accessible is @role.
func TestAccessibleHasRole(accessible Accessible, role AccessibleRole) bool {
	var arg1 *C.GtkAccessible
	var arg2 C.GtkAccessibleRole

	arg1 = (*C.GtkAccessible)(unsafe.Pointer(accessible.Native()))
	arg2 = (C.GtkAccessibleRole)(role)

	var cret C.gboolean
	var ok bool

	cret = C.gtk_test_accessible_has_role(arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}

// TestAccessibleHasState checks whether the Accessible has @state set.
func TestAccessibleHasState(accessible Accessible, state AccessibleState) bool {
	var arg1 *C.GtkAccessible
	var arg2 C.GtkAccessibleState

	arg1 = (*C.GtkAccessible)(unsafe.Pointer(accessible.Native()))
	arg2 = (C.GtkAccessibleState)(state)

	var cret C.gboolean
	var ok bool

	cret = C.gtk_test_accessible_has_state(arg1, arg2)

	if cret {
		ok = true
	}

	return ok
}
