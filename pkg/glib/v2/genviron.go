// Code generated by girgen. DO NOT EDIT.

package glib

import (
	"runtime"
	"unsafe"
)

// #cgo pkg-config: glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <glib.h>
import "C"

// EnvironGetenv returns the value of the environment variable variable in the
// provided list envp.
//
// The function takes the following parameters:
//
//    - envp: an environment list (eg, as returned from g_get_environ()), or
//    NULL for an empty environment list.
//    - variable: environment variable to get.
//
func EnvironGetenv(envp []string, variable string) string {
	var _arg1 **C.gchar // out
	var _arg2 *C.gchar  // out
	var _cret *C.gchar  // in

	{
		_arg1 = (**C.gchar)(C.malloc(C.ulong(len(envp)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(envp)+1)
			var zero *C.gchar
			out[len(envp)] = zero
			for i := range envp {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(envp[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(variable)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_environ_getenv(_arg1, _arg2)
	runtime.KeepAlive(envp)
	runtime.KeepAlive(variable)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
}

// EnvironSetenv sets the environment variable variable in the provided list
// envp to value.
//
// The function takes the following parameters:
//
//    - envp: an environment list that can be freed using g_strfreev() (e.g.,
//    as returned from g_get_environ()), or NULL for an empty environment list.
//    - variable: environment variable to set, must not contain '='.
//    - value for to set the variable to.
//    - overwrite: whether to change the variable if it already exists.
//
func EnvironSetenv(envp []string, variable, value string, overwrite bool) []string {
	var _arg1 **C.gchar  // out
	var _arg2 *C.gchar   // out
	var _arg3 *C.gchar   // out
	var _arg4 C.gboolean // out
	var _cret **C.gchar  // in

	{
		_arg1 = (**C.gchar)(C.malloc(C.ulong(len(envp)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		{
			out := unsafe.Slice(_arg1, len(envp)+1)
			var zero *C.gchar
			out[len(envp)] = zero
			for i := range envp {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(envp[i])))
			}
		}
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(variable)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg3))
	if overwrite {
		_arg4 = C.TRUE
	}

	_cret = C.g_environ_setenv(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(envp)
	runtime.KeepAlive(variable)
	runtime.KeepAlive(value)
	runtime.KeepAlive(overwrite)

	var _filenames []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_filenames = make([]string, i)
		for i := range src {
			_filenames[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _filenames
}

// EnvironUnsetenv removes the environment variable variable from the provided
// environment envp.
//
// The function takes the following parameters:
//
//    - envp: an environment list that can be freed using g_strfreev() (e.g.,
//    as returned from g_get_environ()), or NULL for an empty environment list.
//    - variable: environment variable to remove, must not contain '='.
//
func EnvironUnsetenv(envp []string, variable string) []string {
	var _arg1 **C.gchar // out
	var _arg2 *C.gchar  // out
	var _cret **C.gchar // in

	{
		_arg1 = (**C.gchar)(C.malloc(C.ulong(len(envp)+1) * C.ulong(unsafe.Sizeof(uint(0)))))
		{
			out := unsafe.Slice(_arg1, len(envp)+1)
			var zero *C.gchar
			out[len(envp)] = zero
			for i := range envp {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(envp[i])))
			}
		}
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(variable)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_environ_unsetenv(_arg1, _arg2)
	runtime.KeepAlive(envp)
	runtime.KeepAlive(variable)

	var _filenames []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_filenames = make([]string, i)
		for i := range src {
			_filenames[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _filenames
}

// GetEnviron gets the list of environment variables for the current process.
//
// The list is NULL terminated and each item in the list is of the form
// 'NAME=VALUE'.
//
// This is equivalent to direct access to the 'environ' global variable, except
// portable.
//
// The return value is freshly allocated and it should be freed with
// g_strfreev() when it is no longer needed.
func GetEnviron() []string {
	var _cret **C.gchar // in

	_cret = C.g_get_environ()

	var _filenames []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_filenames = make([]string, i)
		for i := range src {
			_filenames[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _filenames
}

// Getenv returns the value of an environment variable.
//
// On UNIX, the name and value are byte strings which might or might not be in
// some consistent character set and encoding. On Windows, they are in UTF-8. On
// Windows, in case the environment variable's value contains references to
// other environment variables, they are expanded.
//
// The function takes the following parameters:
//
//    - variable: environment variable to get.
//
func Getenv(variable string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(variable)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_getenv(_arg1)
	runtime.KeepAlive(variable)

	var _filename string // out

	_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _filename
}

// Listenv gets the names of all variables set in the environment.
//
// Programs that want to be portable to Windows should typically use this
// function and g_getenv() instead of using the environ array from the C library
// directly. On Windows, the strings in the environ array are in system codepage
// encoding, while in most of the typical use cases for environment variables in
// GLib-using programs you want the UTF-8 encoding that this function and
// g_getenv() provide.
func Listenv() []string {
	var _cret **C.gchar // in

	_cret = C.g_listenv()

	var _filenames []string // out

	defer C.free(unsafe.Pointer(_cret))
	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_filenames = make([]string, i)
		for i := range src {
			_filenames[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _filenames
}

// Setenv sets an environment variable. On UNIX, both the variable's name and
// value can be arbitrary byte strings, except that the variable's name cannot
// contain '='. On Windows, they should be in UTF-8.
//
// Note that on some systems, when variables are overwritten, the memory used
// for the previous variables and its value isn't reclaimed.
//
// You should be mindful of the fact that environment variable handling in UNIX
// is not thread-safe, and your program may crash if one thread calls g_setenv()
// while another thread is calling getenv(). (And note that many functions, such
// as gettext(), call getenv() internally.) This function is only safe to use at
// the very start of your program, before creating any other threads (or
// creating objects that create worker threads of their own).
//
// If you need to set up the environment for a child process, you can use
// g_get_environ() to get an environment array, modify that with
// g_environ_setenv() and g_environ_unsetenv(), and then pass that array
// directly to execvpe(), g_spawn_async(), or the like.
//
// The function takes the following parameters:
//
//    - variable: environment variable to set, must not contain '='.
//    - value for to set the variable to.
//    - overwrite: whether to change the variable if it already exists.
//
func Setenv(variable, value string, overwrite bool) bool {
	var _arg1 *C.gchar   // out
	var _arg2 *C.gchar   // out
	var _arg3 C.gboolean // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(variable)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
	defer C.free(unsafe.Pointer(_arg2))
	if overwrite {
		_arg3 = C.TRUE
	}

	_cret = C.g_setenv(_arg1, _arg2, _arg3)
	runtime.KeepAlive(variable)
	runtime.KeepAlive(value)
	runtime.KeepAlive(overwrite)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Unsetenv removes an environment variable from the environment.
//
// Note that on some systems, when variables are overwritten, the memory used
// for the previous variables and its value isn't reclaimed.
//
// You should be mindful of the fact that environment variable handling in UNIX
// is not thread-safe, and your program may crash if one thread calls
// g_unsetenv() while another thread is calling getenv(). (And note that many
// functions, such as gettext(), call getenv() internally.) This function is
// only safe to use at the very start of your program, before creating any other
// threads (or creating objects that create worker threads of their own).
//
// If you need to set up the environment for a child process, you can use
// g_get_environ() to get an environment array, modify that with
// g_environ_setenv() and g_environ_unsetenv(), and then pass that array
// directly to execvpe(), g_spawn_async(), or the like.
//
// The function takes the following parameters:
//
//    - variable: environment variable to remove, must not contain '='.
//
func Unsetenv(variable string) {
	var _arg1 *C.gchar // out

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(variable)))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_unsetenv(_arg1)
	runtime.KeepAlive(variable)
}
