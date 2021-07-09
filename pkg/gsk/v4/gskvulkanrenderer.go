// Code generated by girgen. DO NOT EDIT.

package gsk

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/gotk3/gotk3/glib"
)

// #cgo pkg-config: gtk4
// #cgo CFLAGS: -Wno-deprecated-declarations
//
// #include <glib-object.h>
// #include <gsk/gsk.h>
import "C"

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: externglib.Type(C.gsk_vulkan_renderer_get_type()), F: marshalVulkanRenderer},
	})
}

// VulkanRenderer: GSK renderer that is using Vulkan.
type VulkanRenderer interface {
	gextras.Objector

	privateVulkanRendererClass()
}

// VulkanRendererClass implements the VulkanRenderer interface.
type VulkanRendererClass struct {
	RendererClass
}

var _ VulkanRenderer = (*VulkanRendererClass)(nil)

func wrapVulkanRenderer(obj *externglib.Object) VulkanRenderer {
	return &VulkanRendererClass{
		RendererClass: RendererClass{
			Object: obj,
		},
	}
}

func marshalVulkanRenderer(p uintptr) (interface{}, error) {
	val := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	obj := externglib.Take(unsafe.Pointer(val))
	return wrapVulkanRenderer(obj), nil
}

func NewVulkanRenderer() *VulkanRendererClass {
	var _cret *C.GskRenderer // in

	_cret = C.gsk_vulkan_renderer_new()

	var _vulkanRenderer *VulkanRendererClass // out

	_vulkanRenderer = (gextras.CastObject(externglib.AssumeOwnership(unsafe.Pointer(_cret)))).(*VulkanRendererClass)

	return _vulkanRenderer
}

func (*VulkanRendererClass) privateVulkanRendererClass() {}
