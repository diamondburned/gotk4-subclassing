// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/internal/ptr"
)

// #cgo pkg-config: gdk-3.0 gtk+-3.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// QueryDepths: this function returns the available bit depths for the default
// screen. It’s equivalent to listing the visuals (gdk_list_visuals()) and then
// looking at the depth field in each visual, removing duplicates.
//
// The array returned by this function should not be freed.
func QueryDepths() []*int {
	var _arg1 *C.gint
	var _arg2 C.gint // in

	C.gdk_query_depths(&_arg1, &_arg2)

	var _depths []*int

	{
		var src []*C.gint
		ptr.SetSlice(unsafe.Pointer(&src), unsafe.Pointer(_arg1), int(_arg2))

		_depths = make([]*int, _arg2)
		for i := 0; i < uintptr(_arg2); i++ {
			_depths = (*int)(_arg1)
		}
	}

	return _depths
}
