package OpenXRActionMap

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Resource"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
OpenXR uses an action system similar to Godots Input map system to bind inputs and outputs on various types of XR controllers to named actions. OpenXR specifies more detail on these inputs and outputs than Godot supports.
Another important distinction is that OpenXR offers no control over these bindings. The bindings we register are suggestions, it is up to the XR runtime to offer users the ability to change these bindings. This allows the XR runtime to fill in the gaps if new hardware becomes available.
The action map therefore needs to be loaded at startup and can't be changed afterwards. This resource is a container for the entire action map.

*/
type Go [1]classdb.OpenXRActionMap

/*
Retrieve the number of actions sets in our action map.
*/
func (self Go) GetActionSetCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetActionSetCount()))
}

/*
Retrieve an action set by name.
*/
func (self Go) FindActionSet(name string) gdclass.OpenXRActionSet {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.OpenXRActionSet(class(self).FindActionSet(gc, gc.String(name)))
}

/*
Retrieve the action set at this index.
*/
func (self Go) GetActionSet(idx int) gdclass.OpenXRActionSet {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.OpenXRActionSet(class(self).GetActionSet(gc, gd.Int(idx)))
}

/*
Add an action set.
*/
func (self Go) AddActionSet(action_set gdclass.OpenXRActionSet) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddActionSet(action_set)
}

/*
Remove an action set.
*/
func (self Go) RemoveActionSet(action_set gdclass.OpenXRActionSet) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveActionSet(action_set)
}

/*
Retrieve the number of interaction profiles in our action map.
*/
func (self Go) GetInteractionProfileCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetInteractionProfileCount()))
}

/*
Find an interaction profile by its name (path).
*/
func (self Go) FindInteractionProfile(name string) gdclass.OpenXRInteractionProfile {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.OpenXRInteractionProfile(class(self).FindInteractionProfile(gc, gc.String(name)))
}

/*
Get the interaction profile at this index.
*/
func (self Go) GetInteractionProfile(idx int) gdclass.OpenXRInteractionProfile {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.OpenXRInteractionProfile(class(self).GetInteractionProfile(gc, gd.Int(idx)))
}

/*
Add an interaction profile.
*/
func (self Go) AddInteractionProfile(interaction_profile gdclass.OpenXRInteractionProfile) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddInteractionProfile(interaction_profile)
}

/*
Remove an interaction profile.
*/
func (self Go) RemoveInteractionProfile(interaction_profile gdclass.OpenXRInteractionProfile) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).RemoveInteractionProfile(interaction_profile)
}

/*
Setup this action set with our default actions.
*/
func (self Go) CreateDefaultActionSets() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).CreateDefaultActionSets()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.OpenXRActionMap
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("OpenXRActionMap"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) ActionSets() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Array(class(self).GetActionSets(gc))
}

func (self Go) SetActionSets(value gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetActionSets(value)
}

func (self Go) InteractionProfiles() gd.Array {
	gc := gd.GarbageCollector(); _ = gc
		return gd.Array(class(self).GetInteractionProfiles(gc))
}

func (self Go) SetInteractionProfiles(value gd.Array) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetInteractionProfiles(value)
}

//go:nosplit
func (self class) SetActionSets(action_sets gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action_sets))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_set_action_sets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetActionSets(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_get_action_sets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Retrieve the number of actions sets in our action map.
*/
//go:nosplit
func (self class) GetActionSetCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_get_action_set_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Retrieve an action set by name.
*/
//go:nosplit
func (self class) FindActionSet(ctx gd.Lifetime, name gd.String) gdclass.OpenXRActionSet {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_find_action_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.OpenXRActionSet
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Retrieve the action set at this index.
*/
//go:nosplit
func (self class) GetActionSet(ctx gd.Lifetime, idx gd.Int) gdclass.OpenXRActionSet {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_get_action_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.OpenXRActionSet
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Add an action set.
*/
//go:nosplit
func (self class) AddActionSet(action_set gdclass.OpenXRActionSet)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action_set[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_add_action_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Remove an action set.
*/
//go:nosplit
func (self class) RemoveActionSet(action_set gdclass.OpenXRActionSet)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(action_set[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_remove_action_set, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) SetInteractionProfiles(interaction_profiles gd.Array)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(interaction_profiles))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_set_interaction_profiles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetInteractionProfiles(ctx gd.Lifetime) gd.Array {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_get_interaction_profiles, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Retrieve the number of interaction profiles in our action map.
*/
//go:nosplit
func (self class) GetInteractionProfileCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_get_interaction_profile_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Find an interaction profile by its name (path).
*/
//go:nosplit
func (self class) FindInteractionProfile(ctx gd.Lifetime, name gd.String) gdclass.OpenXRInteractionProfile {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_find_interaction_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.OpenXRInteractionProfile
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Get the interaction profile at this index.
*/
//go:nosplit
func (self class) GetInteractionProfile(ctx gd.Lifetime, idx gd.Int) gdclass.OpenXRInteractionProfile {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_get_interaction_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.OpenXRInteractionProfile
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Add an interaction profile.
*/
//go:nosplit
func (self class) AddInteractionProfile(interaction_profile gdclass.OpenXRInteractionProfile)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(interaction_profile[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_add_interaction_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Remove an interaction profile.
*/
//go:nosplit
func (self class) RemoveInteractionProfile(interaction_profile gdclass.OpenXRInteractionProfile)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(interaction_profile[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_remove_interaction_profile, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Setup this action set with our default actions.
*/
//go:nosplit
func (self class) CreateDefaultActionSets()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.OpenXRActionMap.Bind_create_default_action_sets, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsOpenXRActionMap() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsOpenXRActionMap() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsResource() Resource.GD { return *((*Resource.GD)(unsafe.Pointer(&self))) }
func (self Go) AsResource() Resource.Go { return *((*Resource.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsResource(), name)
	}
}
func init() {classdb.Register("OpenXRActionMap", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}