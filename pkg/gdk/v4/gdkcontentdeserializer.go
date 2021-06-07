// Code generated by girgen. DO NOT EDIT.

package gdk

// #cgo pkg-config:
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gdk/gdk.h>
import "C"

// ContentDeserializeAsync: read content from the given input stream and
// deserialize it, asynchronously. When the operation is finished, @callback
// will be called. You can then call gdk_content_deserialize_finish() to get the
// result of the operation.
func ContentDeserializeAsync() {
	C.gdk_content_deserialize_async(arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}

// ContentDeserializeFinish finishes a content deserialization operation.
func ContentDeserializeFinish(result gio.AsyncResult, value *externglib.Value) error {
	var arg1 *C.GAsyncResult
	var arg2 *C.GValue

	arg1 = (*C.GAsyncResult)(unsafe.Pointer(result.Native()))
	arg2 = (*C.GValue)(value.GValue)

	var errout *C.GError
	var err error

	C.gdk_content_deserialize_finish(arg1, arg2, &errout)

	err = gerror.Take(unsafe.Pointer(errout))

	return err
}

// ContentRegisterDeserializer registers a function to create objects of a given
// @type from a serialized representation with the given mime type.
func ContentRegisterDeserializer() {
	C.gdk_content_register_deserializer(arg1, arg2, arg3, arg4, arg5)
}