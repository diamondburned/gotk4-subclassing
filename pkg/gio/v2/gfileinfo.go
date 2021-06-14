// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/internal/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gio-2.0 gio-unix-2.0 glib-2.0 gobject-introspection-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <gio/gdesktopappinfo.h>
// #include <gio/gfiledescriptorbased.h>
// #include <gio/gio.h>
// #include <gio/gunixconnection.h>
// #include <gio/gunixcredentialsmessage.h>
// #include <gio/gunixfdlist.h>
// #include <gio/gunixfdmessage.h>
// #include <gio/gunixinputstream.h>
// #include <gio/gunixmounts.h>
// #include <gio/gunixoutputstream.h>
// #include <gio/gunixsocketaddress.h>
// #include <glib-object.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.g_file_info_get_type()), F: marshalFileInfo},
	})
}

// FileInfo: functionality for manipulating basic metadata for files. Info
// implements methods for getting information that all files should contain, and
// allows for manipulation of extended attributes.
//
// See [GFileAttribute][gio-GFileAttribute] for more information on how GIO
// handles file attributes.
//
// To obtain a Info for a #GFile, use g_file_query_info() (or its async
// variant). To obtain a Info for a file input or output stream, use
// g_file_input_stream_query_info() or g_file_output_stream_query_info() (or
// their async variants).
//
// To change the actual attributes of a file, you should then set the attribute
// in the Info and call g_file_set_attributes_from_info() or
// g_file_set_attributes_async() on a GFile.
//
// However, not all attributes can be changed in the file. For instance, the
// actual size of a file cannot be changed via g_file_info_set_size(). You may
// call g_file_query_settable_attributes() and
// g_file_query_writable_namespaces() to discover the settable attributes of a
// particular file at runtime.
//
// AttributeMatcher allows for searching through a Info for attributes.
type FileInfo interface {
	gextras.Objector

	// ClearStatus clears the status information from @info.
	ClearStatus()
	// CopyInto: first clears all of the [GFileAttribute][gio-GFileAttribute] of
	// @dest_info, and then copies all of the file attributes from @src_info to
	// @dest_info.
	CopyInto(destInfo FileInfo)
	// Dup duplicates a file info structure.
	Dup() FileInfo
	// AttributeAsString gets the value of a attribute, formatted as a string.
	// This escapes things as needed to make the string valid UTF-8.
	AttributeAsString(attribute string) string
	// AttributeBoolean gets the value of a boolean attribute. If the attribute
	// does not contain a boolean value, false will be returned.
	AttributeBoolean(attribute string) bool
	// AttributeByteString gets the value of a byte string attribute. If the
	// attribute does not contain a byte string, nil will be returned.
	AttributeByteString(attribute string) string
	// AttributeInt32 gets a signed 32-bit integer contained within the
	// attribute. If the attribute does not contain a signed 32-bit integer, or
	// is invalid, 0 will be returned.
	AttributeInt32(attribute string) int32
	// AttributeInt64 gets a signed 64-bit integer contained within the
	// attribute. If the attribute does not contain a signed 64-bit integer, or
	// is invalid, 0 will be returned.
	AttributeInt64(attribute string) int64
	// AttributeObject gets the value of a #GObject attribute. If the attribute
	// does not contain a #GObject, nil will be returned.
	AttributeObject(attribute string) gextras.Objector
	// AttributeStatus gets the attribute status for an attribute key.
	AttributeStatus(attribute string) FileAttributeStatus
	// AttributeString gets the value of a string attribute. If the attribute
	// does not contain a string, nil will be returned.
	AttributeString(attribute string) string
	// AttributeStringv gets the value of a stringv attribute. If the attribute
	// does not contain a stringv, nil will be returned.
	AttributeStringv(attribute string) []string
	// AttributeType gets the attribute type for an attribute key.
	AttributeType(attribute string) FileAttributeType
	// AttributeUint32 gets an unsigned 32-bit integer contained within the
	// attribute. If the attribute does not contain an unsigned 32-bit integer,
	// or is invalid, 0 will be returned.
	AttributeUint32(attribute string) uint32
	// AttributeUint64 gets a unsigned 64-bit integer contained within the
	// attribute. If the attribute does not contain an unsigned 64-bit integer,
	// or is invalid, 0 will be returned.
	AttributeUint64(attribute string) uint64
	// ContentType gets the file's content type.
	ContentType() string
	// DisplayName gets a display name for a file. This is guaranteed to always
	// be set.
	DisplayName() string
	// EditName gets the edit name for a file.
	EditName() string
	// Etag gets the [entity tag][gfile-etag] for a given Info. See
	// G_FILE_ATTRIBUTE_ETAG_VALUE.
	Etag() string
	// FileType gets a file's type (whether it is a regular file, symlink, etc).
	// This is different from the file's content type, see
	// g_file_info_get_content_type().
	FileType() FileType
	// Icon gets the icon for a file.
	Icon() Icon
	// IsBackup checks if a file is a backup file.
	IsBackup() bool
	// IsHidden checks if a file is hidden.
	IsHidden() bool
	// IsSymlink checks if a file is a symlink.
	IsSymlink() bool
	// ModificationTime gets the modification time of the current @info and sets
	// it in @result.
	ModificationTime() glib.TimeVal
	// Name gets the name for a file. This is guaranteed to always be set.
	Name() string
	// Size gets the file's size (in bytes). The size is retrieved through the
	// value of the G_FILE_ATTRIBUTE_STANDARD_SIZE attribute and is converted
	// from #guint64 to #goffset before returning the result.
	Size() int64
	// SortOrder gets the value of the sort_order attribute from the Info. See
	// G_FILE_ATTRIBUTE_STANDARD_SORT_ORDER.
	SortOrder() int32
	// SymbolicIcon gets the symbolic icon for a file.
	SymbolicIcon() Icon
	// SymlinkTarget gets the symlink target for a given Info.
	SymlinkTarget() string
	// HasAttribute checks if a file info structure has an attribute named
	// @attribute.
	HasAttribute(attribute string) bool
	// HasNamespace checks if a file info structure has an attribute in the
	// specified @name_space.
	HasNamespace(nameSpace string) bool
	// ListAttributes lists the file info structure's attributes.
	ListAttributes(nameSpace string) []string
	// RemoveAttribute removes all cases of @attribute from @info if it exists.
	RemoveAttribute(attribute string)
	// SetAttributeBoolean sets the @attribute to contain the given @attr_value,
	// if possible.
	SetAttributeBoolean(attribute string, attrValue bool)
	// SetAttributeByteString sets the @attribute to contain the given
	// @attr_value, if possible.
	SetAttributeByteString(attribute string, attrValue string)
	// SetAttributeInt32 sets the @attribute to contain the given @attr_value,
	// if possible.
	SetAttributeInt32(attribute string, attrValue int32)
	// SetAttributeInt64 sets the @attribute to contain the given @attr_value,
	// if possible.
	SetAttributeInt64(attribute string, attrValue int64)
	// SetAttributeMask sets @mask on @info to match specific attribute types.
	SetAttributeMask(mask *FileAttributeMatcher)
	// SetAttributeObject sets the @attribute to contain the given @attr_value,
	// if possible.
	SetAttributeObject(attribute string, attrValue gextras.Objector)
	// SetAttributeStatus sets the attribute status for an attribute key. This
	// is only needed by external code that implement
	// g_file_set_attributes_from_info() or similar functions.
	//
	// The attribute must exist in @info for this to work. Otherwise false is
	// returned and @info is unchanged.
	SetAttributeStatus(attribute string, status FileAttributeStatus) bool
	// SetAttributeString sets the @attribute to contain the given @attr_value,
	// if possible.
	SetAttributeString(attribute string, attrValue string)
	// SetAttributeStringv sets the @attribute to contain the given @attr_value,
	// if possible.
	//
	// Sinze: 2.22
	SetAttributeStringv(attribute string, attrValue []string)
	// SetAttributeUint32 sets the @attribute to contain the given @attr_value,
	// if possible.
	SetAttributeUint32(attribute string, attrValue uint32)
	// SetAttributeUint64 sets the @attribute to contain the given @attr_value,
	// if possible.
	SetAttributeUint64(attribute string, attrValue uint64)
	// SetContentType sets the content type attribute for a given Info. See
	// G_FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE.
	SetContentType(contentType string)
	// SetDisplayName sets the display name for the current Info. See
	// G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME.
	SetDisplayName(displayName string)
	// SetEditName sets the edit name for the current file. See
	// G_FILE_ATTRIBUTE_STANDARD_EDIT_NAME.
	SetEditName(editName string)
	// SetFileType sets the file type in a Info to @type. See
	// G_FILE_ATTRIBUTE_STANDARD_TYPE.
	SetFileType(typ FileType)
	// SetIcon sets the icon for a given Info. See
	// G_FILE_ATTRIBUTE_STANDARD_ICON.
	SetIcon(icon Icon)
	// SetIsHidden sets the "is_hidden" attribute in a Info according to
	// @is_hidden. See G_FILE_ATTRIBUTE_STANDARD_IS_HIDDEN.
	SetIsHidden(isHidden bool)
	// SetIsSymlink sets the "is_symlink" attribute in a Info according to
	// @is_symlink. See G_FILE_ATTRIBUTE_STANDARD_IS_SYMLINK.
	SetIsSymlink(isSymlink bool)
	// SetModificationTime sets the G_FILE_ATTRIBUTE_TIME_MODIFIED and
	// G_FILE_ATTRIBUTE_TIME_MODIFIED_USEC attributes in the file info to the
	// given time value.
	SetModificationTime(mtime *glib.TimeVal)
	// SetName sets the name attribute for the current Info. See
	// G_FILE_ATTRIBUTE_STANDARD_NAME.
	SetName(name string)
	// SetSize sets the G_FILE_ATTRIBUTE_STANDARD_SIZE attribute in the file
	// info to the given size.
	SetSize(size int64)
	// SetSortOrder sets the sort order attribute in the file info structure.
	// See G_FILE_ATTRIBUTE_STANDARD_SORT_ORDER.
	SetSortOrder(sortOrder int32)
	// SetSymbolicIcon sets the symbolic icon for a given Info. See
	// G_FILE_ATTRIBUTE_STANDARD_SYMBOLIC_ICON.
	SetSymbolicIcon(icon Icon)
	// SetSymlinkTarget sets the G_FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET
	// attribute in the file info to the given symlink target.
	SetSymlinkTarget(symlinkTarget string)
	// UnsetAttributeMask unsets a mask set by g_file_info_set_attribute_mask(),
	// if one is set.
	UnsetAttributeMask()
}

// fileInfo implements the FileInfo class.
type fileInfo struct {
	gextras.Objector
}

var _ FileInfo = (*fileInfo)(nil)

// WrapFileInfo wraps a GObject to the right type. It is
// primarily used internally.
func WrapFileInfo(obj *externglib.Object) FileInfo {
	return fileInfo{
		Objector: obj,
	}
}

func marshalFileInfo(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return WrapFileInfo(obj), nil
}

// NewFileInfo constructs a class FileInfo.
func NewFileInfo() FileInfo {
	var _cret C.GFileInfo // in

	_cret = C.g_file_info_new()

	var _fileInfo FileInfo // out

	_fileInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(FileInfo)

	return _fileInfo
}

// ClearStatus clears the status information from @info.
func (i fileInfo) ClearStatus() {
	var _arg0 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	C.g_file_info_clear_status(_arg0)
}

// CopyInto: first clears all of the [GFileAttribute][gio-GFileAttribute] of
// @dest_info, and then copies all of the file attributes from @src_info to
// @dest_info.
func (s fileInfo) CopyInto(destInfo FileInfo) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(s.Native()))
	_arg1 = (*C.GFileInfo)(unsafe.Pointer(destInfo.Native()))

	C.g_file_info_copy_into(_arg0, _arg1)
}

// Dup duplicates a file info structure.
func (o fileInfo) Dup() FileInfo {
	var _arg0 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(o.Native()))

	var _cret *C.GFileInfo // in

	_cret = C.g_file_info_dup(_arg0)

	var _fileInfo FileInfo // out

	_fileInfo = gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret.Native()))).(FileInfo)

	return _fileInfo
}

// AttributeAsString gets the value of a attribute, formatted as a string.
// This escapes things as needed to make the string valid UTF-8.
func (i fileInfo) AttributeAsString(attribute string) string {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.char // in

	_cret = C.g_file_info_get_attribute_as_string(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// AttributeBoolean gets the value of a boolean attribute. If the attribute
// does not contain a boolean value, false will be returned.
func (i fileInfo) AttributeBoolean(attribute string) bool {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean // in

	_cret = C.g_file_info_get_attribute_boolean(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AttributeByteString gets the value of a byte string attribute. If the
// attribute does not contain a byte string, nil will be returned.
func (i fileInfo) AttributeByteString(attribute string) string {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.char // in

	_cret = C.g_file_info_get_attribute_byte_string(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// AttributeInt32 gets a signed 32-bit integer contained within the
// attribute. If the attribute does not contain a signed 32-bit integer, or
// is invalid, 0 will be returned.
func (i fileInfo) AttributeInt32(attribute string) int32 {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gint32 // in

	_cret = C.g_file_info_get_attribute_int32(_arg0, _arg1)

	var _gint32 int32 // out

	_gint32 = (int32)(_cret)

	return _gint32
}

// AttributeInt64 gets a signed 64-bit integer contained within the
// attribute. If the attribute does not contain a signed 64-bit integer, or
// is invalid, 0 will be returned.
func (i fileInfo) AttributeInt64(attribute string) int64 {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gint64 // in

	_cret = C.g_file_info_get_attribute_int64(_arg0, _arg1)

	var _gint64 int64 // out

	_gint64 = (int64)(_cret)

	return _gint64
}

// AttributeObject gets the value of a #GObject attribute. If the attribute
// does not contain a #GObject, nil will be returned.
func (i fileInfo) AttributeObject(attribute string) gextras.Objector {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.GObject // in

	_cret = C.g_file_info_get_attribute_object(_arg0, _arg1)

	var _object gextras.Objector // out

	_object = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(gextras.Objector)

	return _object
}

// AttributeStatus gets the attribute status for an attribute key.
func (i fileInfo) AttributeStatus(attribute string) FileAttributeStatus {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GFileAttributeStatus // in

	_cret = C.g_file_info_get_attribute_status(_arg0, _arg1)

	var _fileAttributeStatus FileAttributeStatus // out

	_fileAttributeStatus = FileAttributeStatus(_cret)

	return _fileAttributeStatus
}

// AttributeString gets the value of a string attribute. If the attribute
// does not contain a string, nil will be returned.
func (i fileInfo) AttributeString(attribute string) string {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret *C.char // in

	_cret = C.g_file_info_get_attribute_string(_arg0, _arg1)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// AttributeStringv gets the value of a stringv attribute. If the attribute
// does not contain a stringv, nil will be returned.
func (i fileInfo) AttributeStringv(attribute string) []string {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret **C.char

	_cret = C.g_file_info_get_attribute_stringv(_arg0, _arg1)

	var _utf8s []string

	{
		var i int
		for p := _cret; *p != nil; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
		}
	}

	return _utf8s
}

// AttributeType gets the attribute type for an attribute key.
func (i fileInfo) AttributeType(attribute string) FileAttributeType {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.GFileAttributeType // in

	_cret = C.g_file_info_get_attribute_type(_arg0, _arg1)

	var _fileAttributeType FileAttributeType // out

	_fileAttributeType = FileAttributeType(_cret)

	return _fileAttributeType
}

// AttributeUint32 gets an unsigned 32-bit integer contained within the
// attribute. If the attribute does not contain an unsigned 32-bit integer,
// or is invalid, 0 will be returned.
func (i fileInfo) AttributeUint32(attribute string) uint32 {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.guint32 // in

	_cret = C.g_file_info_get_attribute_uint32(_arg0, _arg1)

	var _guint32 uint32 // out

	_guint32 = (uint32)(_cret)

	return _guint32
}

// AttributeUint64 gets a unsigned 64-bit integer contained within the
// attribute. If the attribute does not contain an unsigned 64-bit integer,
// or is invalid, 0 will be returned.
func (i fileInfo) AttributeUint64(attribute string) uint64 {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.guint64 // in

	_cret = C.g_file_info_get_attribute_uint64(_arg0, _arg1)

	var _guint64 uint64 // out

	_guint64 = (uint64)(_cret)

	return _guint64
}

// ContentType gets the file's content type.
func (i fileInfo) ContentType() string {
	var _arg0 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	var _cret *C.char // in

	_cret = C.g_file_info_get_content_type(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// DisplayName gets a display name for a file. This is guaranteed to always
// be set.
func (i fileInfo) DisplayName() string {
	var _arg0 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	var _cret *C.char // in

	_cret = C.g_file_info_get_display_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// EditName gets the edit name for a file.
func (i fileInfo) EditName() string {
	var _arg0 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	var _cret *C.char // in

	_cret = C.g_file_info_get_edit_name(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// Etag gets the [entity tag][gfile-etag] for a given Info. See
// G_FILE_ATTRIBUTE_ETAG_VALUE.
func (i fileInfo) Etag() string {
	var _arg0 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	var _cret *C.char // in

	_cret = C.g_file_info_get_etag(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// FileType gets a file's type (whether it is a regular file, symlink, etc).
// This is different from the file's content type, see
// g_file_info_get_content_type().
func (i fileInfo) FileType() FileType {
	var _arg0 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	var _cret C.GFileType // in

	_cret = C.g_file_info_get_file_type(_arg0)

	var _fileType FileType // out

	_fileType = FileType(_cret)

	return _fileType
}

// Icon gets the icon for a file.
func (i fileInfo) Icon() Icon {
	var _arg0 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	var _cret *C.GIcon // in

	_cret = C.g_file_info_get_icon(_arg0)

	var _icon Icon // out

	_icon = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Icon)

	return _icon
}

// IsBackup checks if a file is a backup file.
func (i fileInfo) IsBackup() bool {
	var _arg0 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	var _cret C.gboolean // in

	_cret = C.g_file_info_get_is_backup(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsHidden checks if a file is hidden.
func (i fileInfo) IsHidden() bool {
	var _arg0 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	var _cret C.gboolean // in

	_cret = C.g_file_info_get_is_hidden(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsSymlink checks if a file is a symlink.
func (i fileInfo) IsSymlink() bool {
	var _arg0 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	var _cret C.gboolean // in

	_cret = C.g_file_info_get_is_symlink(_arg0)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ModificationTime gets the modification time of the current @info and sets
// it in @result.
func (i fileInfo) ModificationTime() glib.TimeVal {
	var _arg0 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	var _result glib.TimeVal

	C.g_file_info_get_modification_time(_arg0, (*C.GTimeVal)(unsafe.Pointer(&_result)))

	return _result
}

// Name gets the name for a file. This is guaranteed to always be set.
func (i fileInfo) Name() string {
	var _arg0 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	var _cret *C.char // in

	_cret = C.g_file_info_get_name(_arg0)

	var _filename string // out

	_filename = C.GoString(_cret)

	return _filename
}

// Size gets the file's size (in bytes). The size is retrieved through the
// value of the G_FILE_ATTRIBUTE_STANDARD_SIZE attribute and is converted
// from #guint64 to #goffset before returning the result.
func (i fileInfo) Size() int64 {
	var _arg0 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	var _cret C.goffset // in

	_cret = C.g_file_info_get_size(_arg0)

	var _gint64 int64 // out

	_gint64 = (int64)(_cret)

	return _gint64
}

// SortOrder gets the value of the sort_order attribute from the Info. See
// G_FILE_ATTRIBUTE_STANDARD_SORT_ORDER.
func (i fileInfo) SortOrder() int32 {
	var _arg0 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	var _cret C.gint32 // in

	_cret = C.g_file_info_get_sort_order(_arg0)

	var _gint32 int32 // out

	_gint32 = (int32)(_cret)

	return _gint32
}

// SymbolicIcon gets the symbolic icon for a file.
func (i fileInfo) SymbolicIcon() Icon {
	var _arg0 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	var _cret *C.GIcon // in

	_cret = C.g_file_info_get_symbolic_icon(_arg0)

	var _icon Icon // out

	_icon = gextras.CastObject(externglib.Take(unsafe.Pointer(_cret.Native()))).(Icon)

	return _icon
}

// SymlinkTarget gets the symlink target for a given Info.
func (i fileInfo) SymlinkTarget() string {
	var _arg0 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	var _cret *C.char // in

	_cret = C.g_file_info_get_symlink_target(_arg0)

	var _utf8 string // out

	_utf8 = C.GoString(_cret)

	return _utf8
}

// HasAttribute checks if a file info structure has an attribute named
// @attribute.
func (i fileInfo) HasAttribute(attribute string) bool {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean // in

	_cret = C.g_file_info_has_attribute(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasNamespace checks if a file info structure has an attribute in the
// specified @name_space.
func (i fileInfo) HasNamespace(nameSpace string) bool {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(nameSpace))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret C.gboolean // in

	_cret = C.g_file_info_has_namespace(_arg0, _arg1)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ListAttributes lists the file info structure's attributes.
func (i fileInfo) ListAttributes(nameSpace string) []string {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(nameSpace))
	defer C.free(unsafe.Pointer(_arg1))

	var _cret **C.char

	_cret = C.g_file_info_list_attributes(_arg0, _arg1)

	var _utf8s []string

	{
		var i int
		for p := _cret; *p != nil; p = &unsafe.Slice(p, i+1)[i] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString(src[i])
			defer C.free(unsafe.Pointer(src[i]))
		}
	}

	return _utf8s
}

// RemoveAttribute removes all cases of @attribute from @info if it exists.
func (i fileInfo) RemoveAttribute(attribute string) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_file_info_remove_attribute(_arg0, _arg1)
}

// SetAttributeBoolean sets the @attribute to contain the given @attr_value,
// if possible.
func (i fileInfo) SetAttributeBoolean(attribute string, attrValue bool) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _arg2 C.gboolean   // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	if attrValue {
		_arg2 = C.TRUE
	}

	C.g_file_info_set_attribute_boolean(_arg0, _arg1, _arg2)
}

// SetAttributeByteString sets the @attribute to contain the given
// @attr_value, if possible.
func (i fileInfo) SetAttributeByteString(attribute string, attrValue string) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _arg2 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(attrValue))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_file_info_set_attribute_byte_string(_arg0, _arg1, _arg2)
}

// SetAttributeInt32 sets the @attribute to contain the given @attr_value,
// if possible.
func (i fileInfo) SetAttributeInt32(attribute string, attrValue int32) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _arg2 C.gint32     // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint32(attrValue)

	C.g_file_info_set_attribute_int32(_arg0, _arg1, _arg2)
}

// SetAttributeInt64 sets the @attribute to contain the given @attr_value,
// if possible.
func (i fileInfo) SetAttributeInt64(attribute string, attrValue int64) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _arg2 C.gint64     // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gint64(attrValue)

	C.g_file_info_set_attribute_int64(_arg0, _arg1, _arg2)
}

// SetAttributeMask sets @mask on @info to match specific attribute types.
func (i fileInfo) SetAttributeMask(mask *FileAttributeMatcher) {
	var _arg0 *C.GFileInfo             // out
	var _arg1 *C.GFileAttributeMatcher // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GFileAttributeMatcher)(unsafe.Pointer(mask.Native()))

	C.g_file_info_set_attribute_mask(_arg0, _arg1)
}

// SetAttributeObject sets the @attribute to contain the given @attr_value,
// if possible.
func (i fileInfo) SetAttributeObject(attribute string, attrValue gextras.Objector) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _arg2 *C.GObject   // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GObject)(unsafe.Pointer(attrValue.Native()))

	C.g_file_info_set_attribute_object(_arg0, _arg1, _arg2)
}

// SetAttributeStatus sets the attribute status for an attribute key. This
// is only needed by external code that implement
// g_file_set_attributes_from_info() or similar functions.
//
// The attribute must exist in @info for this to work. Otherwise false is
// returned and @info is unchanged.
func (i fileInfo) SetAttributeStatus(attribute string, status FileAttributeStatus) bool {
	var _arg0 *C.GFileInfo           // out
	var _arg1 *C.char                // out
	var _arg2 C.GFileAttributeStatus // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (C.GFileAttributeStatus)(status)

	var _cret C.gboolean // in

	_cret = C.g_file_info_set_attribute_status(_arg0, _arg1, _arg2)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetAttributeString sets the @attribute to contain the given @attr_value,
// if possible.
func (i fileInfo) SetAttributeString(attribute string, attrValue string) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _arg2 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(C.CString(attrValue))
	defer C.free(unsafe.Pointer(_arg2))

	C.g_file_info_set_attribute_string(_arg0, _arg1, _arg2)
}

// SetAttributeStringv sets the @attribute to contain the given @attr_value,
// if possible.
//
// Sinze: 2.22
func (i fileInfo) SetAttributeStringv(attribute string, attrValue []string) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _arg2 **C.char

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (**C.char)(C.malloc(C.ulong((len(attrValue) + 1)) * C.ulong(unsafe.Sizeof(uint(0)))))
	defer C.free(unsafe.Pointer(_arg2))

	{
		out := unsafe.Slice(_arg2, len(attrValue))
		for i := range attrValue {
			out[i] = (*C.gchar)(C.CString(attrValue[i]))
			defer C.free(unsafe.Pointer(out[i]))
		}
	}

	C.g_file_info_set_attribute_stringv(_arg0, _arg1, _arg2)
}

// SetAttributeUint32 sets the @attribute to contain the given @attr_value,
// if possible.
func (i fileInfo) SetAttributeUint32(attribute string, attrValue uint32) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _arg2 C.guint32    // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint32(attrValue)

	C.g_file_info_set_attribute_uint32(_arg0, _arg1, _arg2)
}

// SetAttributeUint64 sets the @attribute to contain the given @attr_value,
// if possible.
func (i fileInfo) SetAttributeUint64(attribute string, attrValue uint64) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out
	var _arg2 C.guint64    // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(attribute))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.guint64(attrValue)

	C.g_file_info_set_attribute_uint64(_arg0, _arg1, _arg2)
}

// SetContentType sets the content type attribute for a given Info. See
// G_FILE_ATTRIBUTE_STANDARD_CONTENT_TYPE.
func (i fileInfo) SetContentType(contentType string) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(contentType))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_file_info_set_content_type(_arg0, _arg1)
}

// SetDisplayName sets the display name for the current Info. See
// G_FILE_ATTRIBUTE_STANDARD_DISPLAY_NAME.
func (i fileInfo) SetDisplayName(displayName string) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(displayName))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_file_info_set_display_name(_arg0, _arg1)
}

// SetEditName sets the edit name for the current file. See
// G_FILE_ATTRIBUTE_STANDARD_EDIT_NAME.
func (i fileInfo) SetEditName(editName string) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(editName))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_file_info_set_edit_name(_arg0, _arg1)
}

// SetFileType sets the file type in a Info to @type. See
// G_FILE_ATTRIBUTE_STANDARD_TYPE.
func (i fileInfo) SetFileType(typ FileType) {
	var _arg0 *C.GFileInfo // out
	var _arg1 C.GFileType  // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (C.GFileType)(typ)

	C.g_file_info_set_file_type(_arg0, _arg1)
}

// SetIcon sets the icon for a given Info. See
// G_FILE_ATTRIBUTE_STANDARD_ICON.
func (i fileInfo) SetIcon(icon Icon) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.GIcon     // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))

	C.g_file_info_set_icon(_arg0, _arg1)
}

// SetIsHidden sets the "is_hidden" attribute in a Info according to
// @is_hidden. See G_FILE_ATTRIBUTE_STANDARD_IS_HIDDEN.
func (i fileInfo) SetIsHidden(isHidden bool) {
	var _arg0 *C.GFileInfo // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	if isHidden {
		_arg1 = C.TRUE
	}

	C.g_file_info_set_is_hidden(_arg0, _arg1)
}

// SetIsSymlink sets the "is_symlink" attribute in a Info according to
// @is_symlink. See G_FILE_ATTRIBUTE_STANDARD_IS_SYMLINK.
func (i fileInfo) SetIsSymlink(isSymlink bool) {
	var _arg0 *C.GFileInfo // out
	var _arg1 C.gboolean   // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	if isSymlink {
		_arg1 = C.TRUE
	}

	C.g_file_info_set_is_symlink(_arg0, _arg1)
}

// SetModificationTime sets the G_FILE_ATTRIBUTE_TIME_MODIFIED and
// G_FILE_ATTRIBUTE_TIME_MODIFIED_USEC attributes in the file info to the
// given time value.
func (i fileInfo) SetModificationTime(mtime *glib.TimeVal) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.GTimeVal  // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GTimeVal)(unsafe.Pointer(mtime.Native()))

	C.g_file_info_set_modification_time(_arg0, _arg1)
}

// SetName sets the name attribute for the current Info. See
// G_FILE_ATTRIBUTE_STANDARD_NAME.
func (i fileInfo) SetName(name string) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(name))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_file_info_set_name(_arg0, _arg1)
}

// SetSize sets the G_FILE_ATTRIBUTE_STANDARD_SIZE attribute in the file
// info to the given size.
func (i fileInfo) SetSize(size int64) {
	var _arg0 *C.GFileInfo // out
	var _arg1 C.goffset    // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = C.goffset(size)

	C.g_file_info_set_size(_arg0, _arg1)
}

// SetSortOrder sets the sort order attribute in the file info structure.
// See G_FILE_ATTRIBUTE_STANDARD_SORT_ORDER.
func (i fileInfo) SetSortOrder(sortOrder int32) {
	var _arg0 *C.GFileInfo // out
	var _arg1 C.gint32     // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = C.gint32(sortOrder)

	C.g_file_info_set_sort_order(_arg0, _arg1)
}

// SetSymbolicIcon sets the symbolic icon for a given Info. See
// G_FILE_ATTRIBUTE_STANDARD_SYMBOLIC_ICON.
func (i fileInfo) SetSymbolicIcon(icon Icon) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.GIcon     // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.GIcon)(unsafe.Pointer(icon.Native()))

	C.g_file_info_set_symbolic_icon(_arg0, _arg1)
}

// SetSymlinkTarget sets the G_FILE_ATTRIBUTE_STANDARD_SYMLINK_TARGET
// attribute in the file info to the given symlink target.
func (i fileInfo) SetSymlinkTarget(symlinkTarget string) {
	var _arg0 *C.GFileInfo // out
	var _arg1 *C.char      // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))
	_arg1 = (*C.char)(C.CString(symlinkTarget))
	defer C.free(unsafe.Pointer(_arg1))

	C.g_file_info_set_symlink_target(_arg0, _arg1)
}

// UnsetAttributeMask unsets a mask set by g_file_info_set_attribute_mask(),
// if one is set.
func (i fileInfo) UnsetAttributeMask() {
	var _arg0 *C.GFileInfo // out

	_arg0 = (*C.GFileInfo)(unsafe.Pointer(i.Native()))

	C.g_file_info_unset_attribute_mask(_arg0)
}
