package ScriptEditor

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/PanelContainer"
import "grow.graphics/gd/gdclass/Container"
import "grow.graphics/gd/gdclass/Control"
import "grow.graphics/gd/gdclass/CanvasItem"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
Godot editor's script editor.
[b]Note:[/b] This class shouldn't be instantiated directly. Instead, access the singleton using [method EditorInterface.get_script_editor].

*/
type Go [1]classdb.ScriptEditor

/*
Returns the [ScriptEditorBase] object that the user is currently editing.
*/
func (self Go) GetCurrentEditor() gdclass.ScriptEditorBase {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.ScriptEditorBase(class(self).GetCurrentEditor(gc))
}

/*
Returns an array with all [ScriptEditorBase] objects which are currently open in editor.
*/
func (self Go) GetOpenScriptEditors() gd.ArrayOf[gdclass.ScriptEditorBase] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gdclass.ScriptEditorBase](class(self).GetOpenScriptEditors(gc))
}

/*
Registers the [EditorSyntaxHighlighter] to the editor, the [EditorSyntaxHighlighter] will be available on all open scripts.
[b]Note:[/b] Does not apply to scripts that are already opened.
*/
func (self Go) RegisterSyntaxHighlighter(syntax_highlighter gdclass.EditorSyntaxHighlighter) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RegisterSyntaxHighlighter(syntax_highlighter)
}

/*
Unregisters the [EditorSyntaxHighlighter] from the editor.
[b]Note:[/b] The [EditorSyntaxHighlighter] will still be applied to scripts that are already opened.
*/
func (self Go) UnregisterSyntaxHighlighter(syntax_highlighter gdclass.EditorSyntaxHighlighter) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).UnregisterSyntaxHighlighter(syntax_highlighter)
}

/*
Goes to the specified line in the current script.
*/
func (self Go) GotoLine(line_number int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).GotoLine(gd.Int(line_number))
}

/*
Returns a [Script] that is currently active in editor.
*/
func (self Go) GetCurrentScript() gdclass.Script {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.Script(class(self).GetCurrentScript(gc))
}

/*
Returns an array with all [Script] objects which are currently open in editor.
*/
func (self Go) GetOpenScripts() gd.ArrayOf[gdclass.Script] {
	gc := gd.GarbageCollector(); _ = gc
	return gd.ArrayOf[gdclass.Script](class(self).GetOpenScripts(gc))
}

/*
Opens the script create dialog. The script will extend [param base_name]. The file extension can be omitted from [param base_path]. It will be added based on the selected scripting language.
*/
func (self Go) OpenScriptCreateDialog(base_name string, base_path string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).OpenScriptCreateDialog(gc.String(base_name), gc.String(base_path))
}

/*
Opens help for the given topic. The [param topic] is an encoded string that controls which class, method, constant, signal, annotation, property, or theme item should be focused.
The supported [param topic] formats include [code]class_name:class[/code], [code]class_method:class:method[/code], [code]class_constant:class:constant[/code], [code]class_signal:class:signal[/code], [code]class_annotation:class:@annotation[/code], [code]class_property:class:property[/code], and [code]class_theme_item:class:item[/code], where [code]class[/code] is the class name, [code]method[/code] is the method name, [code]constant[/code] is the constant name, [code]signal[/code] is the signal name, [code]annotation[/code] is the annotation name, [code]property[/code] is the property name, and [code]item[/code] is the theme item.
[b]Examples:[/b]
[codeblock]
# Shows help for the Node class.
class_name:Node
# Shows help for the global min function.
# Global objects are accessible in the `@GlobalScope` namespace, shown here.
class_method:@GlobalScope:min
# Shows help for get_viewport in the Node class.
class_method:Node:get_viewport
# Shows help for the Input constant MOUSE_BUTTON_MIDDLE.
class_constant:Input:MOUSE_BUTTON_MIDDLE
# Shows help for the BaseButton signal pressed.
class_signal:BaseButton:pressed
# Shows help for the CanvasItem property visible.
class_property:CanvasItem:visible
# Shows help for the GDScript annotation export.
# Annotations should be prefixed with the `@` symbol in the descriptor, as shown here.
class_annotation:@GDScript:@export
# Shows help for the GraphNode theme item named panel_selected.
class_theme_item:GraphNode:panel_selected
[/codeblock]
*/
func (self Go) GotoHelp(topic string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).GotoHelp(gc.String(topic))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ScriptEditor
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ScriptEditor"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Returns the [ScriptEditorBase] object that the user is currently editing.
*/
//go:nosplit
func (self class) GetCurrentEditor(ctx gd.Lifetime) gdclass.ScriptEditorBase {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditor.Bind_get_current_editor, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.ScriptEditorBase
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns an array with all [ScriptEditorBase] objects which are currently open in editor.
*/
//go:nosplit
func (self class) GetOpenScriptEditors(ctx gd.Lifetime) gd.ArrayOf[gdclass.ScriptEditorBase] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditor.Bind_get_open_script_editors, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.ScriptEditorBase](ret)
}
/*
Registers the [EditorSyntaxHighlighter] to the editor, the [EditorSyntaxHighlighter] will be available on all open scripts.
[b]Note:[/b] Does not apply to scripts that are already opened.
*/
//go:nosplit
func (self class) RegisterSyntaxHighlighter(syntax_highlighter gdclass.EditorSyntaxHighlighter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(syntax_highlighter[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditor.Bind_register_syntax_highlighter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Unregisters the [EditorSyntaxHighlighter] from the editor.
[b]Note:[/b] The [EditorSyntaxHighlighter] will still be applied to scripts that are already opened.
*/
//go:nosplit
func (self class) UnregisterSyntaxHighlighter(syntax_highlighter gdclass.EditorSyntaxHighlighter)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(syntax_highlighter[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditor.Bind_unregister_syntax_highlighter, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Goes to the specified line in the current script.
*/
//go:nosplit
func (self class) GotoLine(line_number gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, line_number)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditor.Bind_goto_line, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a [Script] that is currently active in editor.
*/
//go:nosplit
func (self class) GetCurrentScript(ctx gd.Lifetime) gdclass.Script {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditor.Bind_get_current_script, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.Script
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns an array with all [Script] objects which are currently open in editor.
*/
//go:nosplit
func (self class) GetOpenScripts(ctx gd.Lifetime) gd.ArrayOf[gdclass.Script] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditor.Bind_get_open_scripts, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.Script](ret)
}
/*
Opens the script create dialog. The script will extend [param base_name]. The file extension can be omitted from [param base_path]. It will be added based on the selected scripting language.
*/
//go:nosplit
func (self class) OpenScriptCreateDialog(base_name gd.String, base_path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(base_name))
	callframe.Arg(frame, mmm.Get(base_path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditor.Bind_open_script_create_dialog, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Opens help for the given topic. The [param topic] is an encoded string that controls which class, method, constant, signal, annotation, property, or theme item should be focused.
The supported [param topic] formats include [code]class_name:class[/code], [code]class_method:class:method[/code], [code]class_constant:class:constant[/code], [code]class_signal:class:signal[/code], [code]class_annotation:class:@annotation[/code], [code]class_property:class:property[/code], and [code]class_theme_item:class:item[/code], where [code]class[/code] is the class name, [code]method[/code] is the method name, [code]constant[/code] is the constant name, [code]signal[/code] is the signal name, [code]annotation[/code] is the annotation name, [code]property[/code] is the property name, and [code]item[/code] is the theme item.
[b]Examples:[/b]
[codeblock]
# Shows help for the Node class.
class_name:Node
# Shows help for the global min function.
# Global objects are accessible in the `@GlobalScope` namespace, shown here.
class_method:@GlobalScope:min
# Shows help for get_viewport in the Node class.
class_method:Node:get_viewport
# Shows help for the Input constant MOUSE_BUTTON_MIDDLE.
class_constant:Input:MOUSE_BUTTON_MIDDLE
# Shows help for the BaseButton signal pressed.
class_signal:BaseButton:pressed
# Shows help for the CanvasItem property visible.
class_property:CanvasItem:visible
# Shows help for the GDScript annotation export.
# Annotations should be prefixed with the `@` symbol in the descriptor, as shown here.
class_annotation:@GDScript:@export
# Shows help for the GraphNode theme item named panel_selected.
class_theme_item:GraphNode:panel_selected
[/codeblock]
*/
//go:nosplit
func (self class) GotoHelp(topic gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(topic))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ScriptEditor.Bind_goto_help, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnEditorScriptChanged(cb func(script gdclass.Script)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("editor_script_changed"), gc.Callable(cb), 0)
}


func (self Go) OnScriptClose(cb func(script gdclass.Script)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("script_close"), gc.Callable(cb), 0)
}


func (self class) AsScriptEditor() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsScriptEditor() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsPanelContainer() PanelContainer.GD { return *((*PanelContainer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsPanelContainer() PanelContainer.Go { return *((*PanelContainer.Go)(unsafe.Pointer(&self))) }
func (self class) AsContainer() Container.GD { return *((*Container.GD)(unsafe.Pointer(&self))) }
func (self Go) AsContainer() Container.Go { return *((*Container.Go)(unsafe.Pointer(&self))) }
func (self class) AsControl() Control.GD { return *((*Control.GD)(unsafe.Pointer(&self))) }
func (self Go) AsControl() Control.Go { return *((*Control.Go)(unsafe.Pointer(&self))) }
func (self class) AsCanvasItem() CanvasItem.GD { return *((*CanvasItem.GD)(unsafe.Pointer(&self))) }
func (self Go) AsCanvasItem() CanvasItem.Go { return *((*CanvasItem.Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsPanelContainer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsPanelContainer(), name)
	}
}
func init() {classdb.Register("ScriptEditor", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}