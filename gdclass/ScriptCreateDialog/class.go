package ScriptCreateDialog

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/ConfirmationDialog"
import "grow.graphics/gd/gdclass/AcceptDialog"
import "grow.graphics/gd/gdclass/Window"
import "grow.graphics/gd/gdclass/Viewport"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The [ScriptCreateDialog] creates script files according to a given template for a given scripting language. The standard use is to configure its fields prior to calling one of the [method Window.popup] methods.
[codeblocks]
[gdscript]
func _ready():
    var dialog = ScriptCreateDialog.new();
    dialog.config("Node", "res://new_node.gd") # For in-engine types.
    dialog.config("\"res://base_node.gd\"", "res://derived_node.gd") # For script types.
    dialog.popup_centered()
[/gdscript]
[csharp]
public override void _Ready()
{
    var dialog = new ScriptCreateDialog();
    dialog.Config("Node", "res://NewNode.cs"); // For in-engine types.
    dialog.Config("\"res://BaseNode.cs\"", "res://DerivedNode.cs"); // For script types.
    dialog.PopupCentered();
}
[/csharp]
[/codeblocks]

*/
type Go [1]classdb.ScriptCreateDialog

/*
Prefills required fields to configure the ScriptCreateDialog for use.
*/
func (self Go) Config(inherits string, path string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Config(gc.String(inherits), gc.String(path), true, true)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ScriptCreateDialog
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ScriptCreateDialog"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Prefills required fields to configure the ScriptCreateDialog for use.
*/
//go:nosplit
func (self class) Config(inherits gd.String, path gd.String, built_in_enabled bool, load_enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(inherits))
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, built_in_enabled)
	callframe.Arg(frame, load_enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptCreateDialog.Bind_config, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnScriptCreated(cb func(script gdclass.Script)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("script_created"), gc.Callable(cb), 0)
}


func (self class) AsScriptCreateDialog() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsScriptCreateDialog() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsConfirmationDialog() ConfirmationDialog.GD { return *((*ConfirmationDialog.GD)(unsafe.Pointer(&self))) }
func (self Go) AsConfirmationDialog() ConfirmationDialog.Go { return *((*ConfirmationDialog.Go)(unsafe.Pointer(&self))) }
func (self class) AsAcceptDialog() AcceptDialog.GD { return *((*AcceptDialog.GD)(unsafe.Pointer(&self))) }
func (self Go) AsAcceptDialog() AcceptDialog.Go { return *((*AcceptDialog.Go)(unsafe.Pointer(&self))) }
func (self class) AsWindow() Window.GD { return *((*Window.GD)(unsafe.Pointer(&self))) }
func (self Go) AsWindow() Window.Go { return *((*Window.Go)(unsafe.Pointer(&self))) }
func (self class) AsViewport() Viewport.GD { return *((*Viewport.GD)(unsafe.Pointer(&self))) }
func (self Go) AsViewport() Viewport.Go { return *((*Viewport.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsConfirmationDialog(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsConfirmationDialog(), name)
	}
}
func init() {classdb.Register("ScriptCreateDialog", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}