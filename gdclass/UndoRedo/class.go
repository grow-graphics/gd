package UndoRedo

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
UndoRedo works by registering methods and property changes inside "actions". You can create an action, then provide ways to do and undo this action using function calls and property changes, then commit the action.
When an action is committed, all of the [code]do_*[/code] methods will run. If the [method undo] method is used, the [code]undo_*[/code] methods will run. If the [method redo] method is used, once again, all of the [code]do_*[/code] methods will run.
Here's an example on how to add an action:
[codeblocks]
[gdscript]
var undo_redo = UndoRedo.new()

func do_something():
    pass # Put your code here.

func undo_something():
    pass # Put here the code that reverts what's done by "do_something()".

func _on_my_button_pressed():
    var node = get_node("MyNode2D")
    undo_redo.create_action("Move the node")
    undo_redo.add_do_method(do_something)
    undo_redo.add_undo_method(undo_something)
    undo_redo.add_do_property(node, "position", Vector2(100,100))
    undo_redo.add_undo_property(node, "position", node.position)
    undo_redo.commit_action()
[/gdscript]
[csharp]
private UndoRedo _undoRedo;

public override void _Ready()
{
    _undoRedo = new UndoRedo();
}

public void DoSomething()
{
    // Put your code here.
}

public void UndoSomething()
{
    // Put here the code that reverts what's done by "DoSomething()".
}

private void OnMyButtonPressed()
{
    var node = GetNode<Node2D>("MyNode2D");
    _undoRedo.CreateAction("Move the node");
    _undoRedo.AddDoMethod(new Callable(this, MethodName.DoSomething));
    _undoRedo.AddUndoMethod(new Callable(this, MethodName.UndoSomething));
    _undoRedo.AddDoProperty(node, "position", new Vector2(100, 100));
    _undoRedo.AddUndoProperty(node, "position", node.Position);
    _undoRedo.CommitAction();
}
[/csharp]
[/codeblocks]
Before calling any of the [code]add_(un)do_*[/code] methods, you need to first call [method create_action]. Afterwards you need to call [method commit_action].
If you don't need to register a method, you can leave [method add_do_method] and [method add_undo_method] out; the same goes for properties. You can also register more than one method/property.
If you are making an [EditorPlugin] and want to integrate into the editor's undo history, use [EditorUndoRedoManager] instead.
If you are registering multiple properties/method which depend on one another, be aware that by default undo operation are called in the same order they have been added. Therefore instead of grouping do operation with their undo operations it is better to group do on one side and undo on the other as shown below.
[codeblocks]
[gdscript]
undo_redo.create_action("Add object")

# DO
undo_redo.add_do_method(_create_object)
undo_redo.add_do_method(_add_object_to_singleton)

# UNDO
undo_redo.add_undo_method(_remove_object_from_singleton)
undo_redo.add_undo_method(_destroy_that_object)

undo_redo.commit_action()
[/gdscript]
[csharp]
_undo_redo.CreateAction("Add object");

// DO
_undo_redo.AddDoMethod(new Callable(this, MethodName.CreateObject));
_undo_redo.AddDoMethod(new Callable(this, MethodName.AddObjectToSingleton));

// UNDO
_undo_redo.AddUndoMethod(new Callable(this, MethodName.RemoveObjectFromSingleton));
_undo_redo.AddUndoMethod(new Callable(this, MethodName.DestroyThatObject));

_undo_redo.CommitAction();
[/csharp]
[/codeblocks]

*/
type Go [1]classdb.UndoRedo

/*
Create a new action. After this is called, do all your calls to [method add_do_method], [method add_undo_method], [method add_do_property], and [method add_undo_property], then commit the action with [method commit_action].
The way actions are merged is dictated by [param merge_mode]. See [enum MergeMode] for details.
The way undo operation are ordered in actions is dictated by [param backward_undo_ops]. When [param backward_undo_ops] is [code]false[/code] undo option are ordered in the same order they were added. Which means the first operation to be added will be the first to be undone.
*/
func (self Go) CreateAction(name string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).CreateAction(gc.String(name), 0, false)
}

/*
Commit the action. If [param execute] is [code]true[/code] (which it is by default), all "do" methods/properties are called/set when this function is called.
*/
func (self Go) CommitAction() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).CommitAction(true)
}

/*
Returns [code]true[/code] if the [UndoRedo] is currently committing the action, i.e. running its "do" method or property change (see [method commit_action]).
*/
func (self Go) IsCommittingAction() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsCommittingAction())
}

/*
Register a [Callable] that will be called when the action is committed.
*/
func (self Go) AddDoMethod(callable gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddDoMethod(callable)
}

/*
Register a [Callable] that will be called when the action is undone.
*/
func (self Go) AddUndoMethod(callable gd.Callable) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddUndoMethod(callable)
}

/*
Register a [param property] that would change its value to [param value] when the action is committed.
*/
func (self Go) AddDoProperty(obj gd.Object, property string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddDoProperty(obj, gc.StringName(property), value)
}

/*
Register a [param property] that would change its value to [param value] when the action is undone.
*/
func (self Go) AddUndoProperty(obj gd.Object, property string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddUndoProperty(obj, gc.StringName(property), value)
}

/*
Register a reference to an object that will be erased if the "do" history is deleted. This is useful for objects added by the "do" action and removed by the "undo" action.
When the "do" history is deleted, if the object is a [RefCounted], it will be unreferenced. Otherwise, it will be freed. Do not use for resources.
[codeblock]
var node = Node2D.new()
undo_redo.create_action("Add node")
undo_redo.add_do_method(add_child.bind(node))
undo_redo.add_do_reference(node)
undo_redo.add_undo_method(remove_child.bind(node))
undo_redo.commit_action()
[/codeblock]
*/
func (self Go) AddDoReference(obj gd.Object) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddDoReference(obj)
}

/*
Register a reference to an object that will be erased if the "undo" history is deleted. This is useful for objects added by the "undo" action and removed by the "do" action.
When the "undo" history is deleted, if the object is a [RefCounted], it will be unreferenced. Otherwise, it will be freed. Do not use for resources.
[codeblock]
var node = $Node2D
undo_redo.create_action("Remove node")
undo_redo.add_do_method(remove_child.bind(node))
undo_redo.add_undo_method(add_child.bind(node))
undo_redo.add_undo_reference(node)
undo_redo.commit_action()
[/codeblock]
*/
func (self Go) AddUndoReference(obj gd.Object) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddUndoReference(obj)
}

/*
Marks the next "do" and "undo" operations to be processed even if the action gets merged with another in the [constant MERGE_ENDS] mode. Return to normal operation using [method end_force_keep_in_merge_ends].
*/
func (self Go) StartForceKeepInMergeEnds() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).StartForceKeepInMergeEnds()
}

/*
Stops marking operations as to be processed even if the action gets merged with another in the [constant MERGE_ENDS] mode. See [method start_force_keep_in_merge_ends].
*/
func (self Go) EndForceKeepInMergeEnds() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).EndForceKeepInMergeEnds()
}

/*
Returns how many elements are in the history.
*/
func (self Go) GetHistoryCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetHistoryCount()))
}

/*
Gets the index of the current action.
*/
func (self Go) GetCurrentAction() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetCurrentAction()))
}

/*
Gets the action name from its index.
*/
func (self Go) GetActionName(id int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetActionName(gc, gd.Int(id)).String())
}

/*
Clear the undo/redo history and associated references.
Passing [code]false[/code] to [param increase_version] will prevent the version number from increasing when the history is cleared.
*/
func (self Go) ClearHistory() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ClearHistory(true)
}

/*
Gets the name of the current action, equivalent to [code]get_action_name(get_current_action())[/code].
*/
func (self Go) GetCurrentActionName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetCurrentActionName(gc).String())
}

/*
Returns [code]true[/code] if an "undo" action is available.
*/
func (self Go) HasUndo() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasUndo())
}

/*
Returns [code]true[/code] if a "redo" action is available.
*/
func (self Go) HasRedo() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasRedo())
}

/*
Gets the version. Every time a new action is committed, the [UndoRedo]'s version number is increased automatically.
This is useful mostly to check if something changed from a saved version.
*/
func (self Go) GetVersion() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetVersion()))
}

/*
Redo the last action.
*/
func (self Go) Redo() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).Redo())
}

/*
Undo the last action.
*/
func (self Go) Undo() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).Undo())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.UndoRedo
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("UndoRedo"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) MaxSteps() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetMaxSteps()))
}

func (self Go) SetMaxSteps(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetMaxSteps(gd.Int(value))
}

/*
Create a new action. After this is called, do all your calls to [method add_do_method], [method add_undo_method], [method add_do_property], and [method add_undo_property], then commit the action with [method commit_action].
The way actions are merged is dictated by [param merge_mode]. See [enum MergeMode] for details.
The way undo operation are ordered in actions is dictated by [param backward_undo_ops]. When [param backward_undo_ops] is [code]false[/code] undo option are ordered in the same order they were added. Which means the first operation to be added will be the first to be undone.
*/
//go:nosplit
func (self class) CreateAction(name gd.String, merge_mode classdb.UndoRedoMergeMode, backward_undo_ops bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, merge_mode)
	callframe.Arg(frame, backward_undo_ops)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_create_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Commit the action. If [param execute] is [code]true[/code] (which it is by default), all "do" methods/properties are called/set when this function is called.
*/
//go:nosplit
func (self class) CommitAction(execute bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, execute)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_commit_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns [code]true[/code] if the [UndoRedo] is currently committing the action, i.e. running its "do" method or property change (see [method commit_action]).
*/
//go:nosplit
func (self class) IsCommittingAction() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_is_committing_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Register a [Callable] that will be called when the action is committed.
*/
//go:nosplit
func (self class) AddDoMethod(callable gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(callable))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_add_do_method, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Register a [Callable] that will be called when the action is undone.
*/
//go:nosplit
func (self class) AddUndoMethod(callable gd.Callable)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(callable))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_add_undo_method, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Register a [param property] that would change its value to [param value] when the action is committed.
*/
//go:nosplit
func (self class) AddDoProperty(obj gd.Object, property gd.StringName, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	callframe.Arg(frame, mmm.Get(property))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_add_do_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Register a [param property] that would change its value to [param value] when the action is undone.
*/
//go:nosplit
func (self class) AddUndoProperty(obj gd.Object, property gd.StringName, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	callframe.Arg(frame, mmm.Get(property))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_add_undo_property, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Register a reference to an object that will be erased if the "do" history is deleted. This is useful for objects added by the "do" action and removed by the "undo" action.
When the "do" history is deleted, if the object is a [RefCounted], it will be unreferenced. Otherwise, it will be freed. Do not use for resources.
[codeblock]
var node = Node2D.new()
undo_redo.create_action("Add node")
undo_redo.add_do_method(add_child.bind(node))
undo_redo.add_do_reference(node)
undo_redo.add_undo_method(remove_child.bind(node))
undo_redo.commit_action()
[/codeblock]
*/
//go:nosplit
func (self class) AddDoReference(obj gd.Object)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_add_do_reference, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Register a reference to an object that will be erased if the "undo" history is deleted. This is useful for objects added by the "undo" action and removed by the "do" action.
When the "undo" history is deleted, if the object is a [RefCounted], it will be unreferenced. Otherwise, it will be freed. Do not use for resources.
[codeblock]
var node = $Node2D
undo_redo.create_action("Remove node")
undo_redo.add_do_method(remove_child.bind(node))
undo_redo.add_undo_method(add_child.bind(node))
undo_redo.add_undo_reference(node)
undo_redo.commit_action()
[/codeblock]
*/
//go:nosplit
func (self class) AddUndoReference(obj gd.Object)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.End(obj.AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_add_undo_reference, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Marks the next "do" and "undo" operations to be processed even if the action gets merged with another in the [constant MERGE_ENDS] mode. Return to normal operation using [method end_force_keep_in_merge_ends].
*/
//go:nosplit
func (self class) StartForceKeepInMergeEnds()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_start_force_keep_in_merge_ends, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Stops marking operations as to be processed even if the action gets merged with another in the [constant MERGE_ENDS] mode. See [method start_force_keep_in_merge_ends].
*/
//go:nosplit
func (self class) EndForceKeepInMergeEnds()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_end_force_keep_in_merge_ends, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns how many elements are in the history.
*/
//go:nosplit
func (self class) GetHistoryCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_get_history_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets the index of the current action.
*/
//go:nosplit
func (self class) GetCurrentAction() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_get_current_action, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets the action name from its index.
*/
//go:nosplit
func (self class) GetActionName(ctx gd.Lifetime, id gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_get_action_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Clear the undo/redo history and associated references.
Passing [code]false[/code] to [param increase_version] will prevent the version number from increasing when the history is cleared.
*/
//go:nosplit
func (self class) ClearHistory(increase_version bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, increase_version)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_clear_history, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Gets the name of the current action, equivalent to [code]get_action_name(get_current_action())[/code].
*/
//go:nosplit
func (self class) GetCurrentActionName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_get_current_action_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if an "undo" action is available.
*/
//go:nosplit
func (self class) HasUndo() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_has_undo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if a "redo" action is available.
*/
//go:nosplit
func (self class) HasRedo() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_has_redo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets the version. Every time a new action is committed, the [UndoRedo]'s version number is increased automatically.
This is useful mostly to check if something changed from a saved version.
*/
//go:nosplit
func (self class) GetVersion() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_get_version, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetMaxSteps(max_steps gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, max_steps)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_set_max_steps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetMaxSteps() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_get_max_steps, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Redo the last action.
*/
//go:nosplit
func (self class) Redo() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_redo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Undo the last action.
*/
//go:nosplit
func (self class) Undo() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.UndoRedo.Bind_undo, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnVersionChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("version_changed"), gc.Callable(cb), 0)
}


func (self class) AsUndoRedo() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsUndoRedo() Go { return *((*Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("UndoRedo", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type MergeMode = classdb.UndoRedoMergeMode

const (
/*Makes "do"/"undo" operations stay in separate actions.*/
	MergeDisable MergeMode = 0
/*Merges this action with the previous one if they have the same name. Keeps only the first action's "undo" operations and the last action's "do" operations. Useful for sequential changes to a single value.*/
	MergeEnds MergeMode = 1
/*Merges this action with the previous one if they have the same name.*/
	MergeAll MergeMode = 2
)