package RenderDataRD

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/RenderData"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This object manages all render data for the rendering device based renderers.
[b]Note:[/b] This is an internal rendering server object only exposed for GDExtension plugins.

*/
type Go [1]classdb.RenderDataRD
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.RenderDataRD
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("RenderDataRD"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsRenderDataRD() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsRenderDataRD() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRenderData() RenderData.GD { return *((*RenderData.GD)(unsafe.Pointer(&self))) }
func (self Go) AsRenderData() RenderData.Go { return *((*RenderData.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRenderData(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRenderData(), name)
	}
}
func init() {classdb.Register("RenderDataRD", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}