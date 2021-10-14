// Code generated by girgen. DO NOT EDIT.

package gtk

import (
	"runtime"
	"unsafe"

	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib-object.h>
// #include <gtk/gtk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gtk_numeric_sorter_get_type()), F: marshalNumericSorterer},
	})
}

// NumericSorter: GtkNumericSorter is a GtkSorter that compares numbers.
//
// To obtain the numbers to compare, this sorter evaluates a gtk.Expression.
type NumericSorter struct {
	Sorter
}

func wrapNumericSorter(obj *externglib.Object) *NumericSorter {
	return &NumericSorter{
		Sorter: Sorter{
			Object: obj,
		},
	}
}

func marshalNumericSorterer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapNumericSorter(obj), nil
}

// NewNumericSorter creates a new numeric sorter using the given expression.
//
// Smaller numbers will be sorted first. You can call
// gtk.NumericSorter.SetSortOrder() to change this.
//
// The function takes the following parameters:
//
//    - expression to evaluate.
//
func NewNumericSorter(expression Expressioner) *NumericSorter {
	var _arg1 *C.GtkExpression    // out
	var _cret *C.GtkNumericSorter // in

	if expression != nil {
		_arg1 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))
		C.g_object_ref(C.gpointer(expression.Native()))
	}

	_cret = C.gtk_numeric_sorter_new(_arg1)
	runtime.KeepAlive(expression)

	var _numericSorter *NumericSorter // out

	_numericSorter = wrapNumericSorter(externglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _numericSorter
}

// Expression gets the expression that is evaluated to obtain numbers from
// items.
func (self *NumericSorter) Expression() Expressioner {
	var _arg0 *C.GtkNumericSorter // out
	var _cret *C.GtkExpression    // in

	_arg0 = (*C.GtkNumericSorter)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_numeric_sorter_get_expression(_arg0)
	runtime.KeepAlive(self)

	var _expression Expressioner // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := externglib.Take(objptr)
			rv, ok := (externglib.CastObject(object)).(Expressioner)
			if !ok {
				panic("object of type " + object.TypeFromInstance().String() + " is not gtk.Expressioner")
			}
			_expression = rv
		}
	}

	return _expression
}

// SortOrder gets whether this sorter will sort smaller numbers first.
func (self *NumericSorter) SortOrder() SortType {
	var _arg0 *C.GtkNumericSorter // out
	var _cret C.GtkSortType       // in

	_arg0 = (*C.GtkNumericSorter)(unsafe.Pointer(self.Native()))

	_cret = C.gtk_numeric_sorter_get_sort_order(_arg0)
	runtime.KeepAlive(self)

	var _sortType SortType // out

	_sortType = SortType(_cret)

	return _sortType
}

// SetExpression sets the expression that is evaluated to obtain numbers from
// items.
//
// Unless an expression is set on self, the sorter will always compare items as
// invalid.
//
// The expression must have a return type that can be compared numerically, such
// as G_TYPE_INT or G_TYPE_DOUBLE.
//
// The function takes the following parameters:
//
//    - expression: GtkExpression, or NULL.
//
func (self *NumericSorter) SetExpression(expression Expressioner) {
	var _arg0 *C.GtkNumericSorter // out
	var _arg1 *C.GtkExpression    // out

	_arg0 = (*C.GtkNumericSorter)(unsafe.Pointer(self.Native()))
	if expression != nil {
		_arg1 = (*C.GtkExpression)(unsafe.Pointer(expression.Native()))
	}

	C.gtk_numeric_sorter_set_expression(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(expression)
}

// SetSortOrder sets whether to sort smaller numbers before larger ones.
//
// The function takes the following parameters:
//
//    - sortOrder: whether to sort smaller numbers first.
//
func (self *NumericSorter) SetSortOrder(sortOrder SortType) {
	var _arg0 *C.GtkNumericSorter // out
	var _arg1 C.GtkSortType       // out

	_arg0 = (*C.GtkNumericSorter)(unsafe.Pointer(self.Native()))
	_arg1 = C.GtkSortType(sortOrder)

	C.gtk_numeric_sorter_set_sort_order(_arg0, _arg1)
	runtime.KeepAlive(self)
	runtime.KeepAlive(sortOrder)
}
