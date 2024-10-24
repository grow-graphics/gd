package EditorExportPlatformWeb

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/EditorExportPlatform"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The Web exporter customizes how a web build is handled. In the editor's "Export" window, it is created when adding a new "Web" preset.
[b]Note:[/b] Godot on Web is rendered inside a [code]<canvas>[/code] tag. Normally, the canvas cannot be positioned or resized manually, but otherwise acts as the main [Window] of the application.

*/
type Go [1]classdb.EditorExportPlatformWeb
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorExportPlatformWeb
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("EditorExportPlatformWeb"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self class) AsEditorExportPlatformWeb() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorExportPlatformWeb() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsEditorExportPlatform() EditorExportPlatform.GD { return *((*EditorExportPlatform.GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorExportPlatform() EditorExportPlatform.Go { return *((*EditorExportPlatform.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsEditorExportPlatform(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsEditorExportPlatform(), name)
	}
}
func init() {classdb.Register("EditorExportPlatformWeb", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}