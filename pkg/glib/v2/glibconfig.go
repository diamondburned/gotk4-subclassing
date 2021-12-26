// Code generated by girgen. DO NOT EDIT.

package glib

// #include <stdlib.h>
// #include <glib.h>
import "C"

// DIR_SEPARATOR: directory separator character. This is '/' on UNIX machines
// and '\' under Windows.
const DIR_SEPARATOR = 47

// DIR_SEPARATOR_S: directory separator as a string. This is "/" on UNIX
// machines and "\" under Windows.
const DIR_SEPARATOR_S = "/"

// GINT16_FORMAT: this is the platform dependent conversion specifier for
// scanning and printing values of type #gint16. It is a string literal, but
// doesn't include the percent-sign, such that you can add precision and length
// modifiers between percent-sign and conversion specifier.
//
//    gint16 in;
//    gint32 out;
//    sscanf ("42", "%" G_GINT16_FORMAT, &in)
//    out = in * 1000;
//    g_print ("%" G_GINT32_FORMAT, out);.
const GINT16_FORMAT = "hi"

// GINT16_MODIFIER: platform dependent length modifier for conversion specifiers
// for scanning and printing values of type #gint16 or #guint16. It is a string
// literal, but doesn't include the percent-sign, such that you can add
// precision and length modifiers between percent-sign and conversion specifier
// and append a conversion specifier.
//
// The following example prints "0x7b";
//
//    gint16 value = 123;
//    g_print ("%#" G_GINT16_MODIFIER "x", value);.
const GINT16_MODIFIER = "h"

// GINT32_FORMAT: this is the platform dependent conversion specifier for
// scanning and printing values of type #gint32. See also GINT16_FORMAT.
const GINT32_FORMAT = "i"

// GINT32_MODIFIER: platform dependent length modifier for conversion specifiers
// for scanning and printing values of type #gint32 or #guint32. It is a string
// literal. See also GINT16_MODIFIER.
const GINT32_MODIFIER = ""

// GINT64_FORMAT: this is the platform dependent conversion specifier for
// scanning and printing values of type #gint64. See also GINT16_FORMAT.
//
// Some platforms do not support scanning and printing 64-bit integers, even
// though the types are supported. On such platforms G_GINT64_FORMAT is not
// defined. Note that scanf() may not support 64-bit integers, even if
// G_GINT64_FORMAT is defined. Due to its weak error handling, scanf() is not
// recommended for parsing anyway; consider using g_ascii_strtoull() instead.
const GINT64_FORMAT = "li"

// GINT64_MODIFIER: platform dependent length modifier for conversion specifiers
// for scanning and printing values of type #gint64 or #guint64. It is a string
// literal.
//
// Some platforms do not support printing 64-bit integers, even though the types
// are supported. On such platforms G_GINT64_MODIFIER is not defined.
const GINT64_MODIFIER = "l"

// GINTPTR_FORMAT: this is the platform dependent conversion specifier for
// scanning and printing values of type #gintptr.
const GINTPTR_FORMAT = "li"

// GINTPTR_MODIFIER: platform dependent length modifier for conversion
// specifiers for scanning and printing values of type #gintptr or #guintptr. It
// is a string literal.
const GINTPTR_MODIFIER = "l"

// GSIZE_FORMAT: this is the platform dependent conversion specifier for
// scanning and printing values of type #gsize. See also GINT16_FORMAT.
const GSIZE_FORMAT = "lu"

// GSIZE_MODIFIER: platform dependent length modifier for conversion specifiers
// for scanning and printing values of type #gsize. It is a string literal.
const GSIZE_MODIFIER = "l"

// GSSIZE_FORMAT: this is the platform dependent conversion specifier for
// scanning and printing values of type #gssize. See also GINT16_FORMAT.
const GSSIZE_FORMAT = "li"

// GSSIZE_MODIFIER: platform dependent length modifier for conversion specifiers
// for scanning and printing values of type #gssize. It is a string literal.
const GSSIZE_MODIFIER = "l"

// GUINT16_FORMAT: this is the platform dependent conversion specifier for
// scanning and printing values of type #guint16. See also GINT16_FORMAT.
const GUINT16_FORMAT = "hu"

// GUINT32_FORMAT: this is the platform dependent conversion specifier for
// scanning and printing values of type #guint32. See also GINT16_FORMAT.
const GUINT32_FORMAT = "u"

// GUINT64_FORMAT: this is the platform dependent conversion specifier for
// scanning and printing values of type #guint64. See also GINT16_FORMAT.
//
// Some platforms do not support scanning and printing 64-bit integers, even
// though the types are supported. On such platforms G_GUINT64_FORMAT is not
// defined. Note that scanf() may not support 64-bit integers, even if
// G_GINT64_FORMAT is defined. Due to its weak error handling, scanf() is not
// recommended for parsing anyway; consider using g_ascii_strtoull() instead.
const GUINT64_FORMAT = "lu"

// GUINTPTR_FORMAT: this is the platform dependent conversion specifier for
// scanning and printing values of type #guintptr.
const GUINTPTR_FORMAT = "lu"
const HAVE_GINT64 = 1
const HAVE_GNUC_VARARGS = 1

// HAVE_GNUC_VISIBILITY: defined to 1 if gcc-style visibility handling is
// supported.
const HAVE_GNUC_VISIBILITY = 1
const HAVE_GROWING_STACK = 0
const HAVE_ISO_VARARGS = 1

// MAJOR_VERSION: major version number of the GLib library.
//
// Like #glib_major_version, but from the headers used at application compile
// time, rather than from the library linked against at application run time.
const MAJOR_VERSION = 2

// MICRO_VERSION: micro version number of the GLib library.
//
// Like #gtk_micro_version, but from the headers used at application compile
// time, rather than from the library linked against at application run time.
const MICRO_VERSION = 2

// MINOR_VERSION: minor version number of the GLib library.
//
// Like #gtk_minor_version, but from the headers used at application compile
// time, rather than from the library linked against at application run time.
const MINOR_VERSION = 68
const MODULE_SUFFIX = "so"

// PID_FORMAT: format specifier that can be used in printf()-style format
// strings when printing a #GPid.
const PID_FORMAT = "i"

// POLLFD_FORMAT: format specifier that can be used in printf()-style format
// strings when printing the fd member of a FD.
const POLLFD_FORMAT = "%d"

// SEARCHPATH_SEPARATOR: search path separator character. This is ':' on UNIX
// machines and ';' under Windows.
const SEARCHPATH_SEPARATOR = 58

// SEARCHPATH_SEPARATOR_S: search path separator as a string. This is ":" on
// UNIX machines and ";" under Windows.
const SEARCHPATH_SEPARATOR_S = ":"
const SIZEOF_LONG = 8
const SIZEOF_SIZE_T = 8
const SIZEOF_SSIZE_T = 8
const SIZEOF_VOID_P = 8
const SYSDEF_AF_INET = 2
const SYSDEF_AF_INET6 = 10
const SYSDEF_AF_UNIX = 1
const SYSDEF_MSG_DONTROUTE = 4
const SYSDEF_MSG_OOB = 1
const SYSDEF_MSG_PEEK = 2
const VA_COPY_AS_ARRAY = 1
