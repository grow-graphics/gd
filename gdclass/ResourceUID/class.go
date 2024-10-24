package ResourceUID

import "unsafe"
import "sync"
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
Resource UIDs (Unique IDentifiers) allow the engine to keep references between resources intact, even if files can renamed or moved. They can be accessed with [code]uid://[/code].
[ResourceUID] keeps track of all registered resource UIDs in a project, generates new UIDs, and converts between their string and integer representations.

*/
var self gdclass.ResourceUID
var once sync.Once
func singleton() {
	gc := gd.Static
	obj := gc.API.Object.GetSingleton(gc, gc.API.Singletons.ResourceUID)
	self = *(*gdclass.ResourceUID)(unsafe.Pointer(&obj))
}

/*
Converts the given UID to a [code]uid://[/code] string value.
*/
func IdToText(id int) string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).IdToText(gc, gd.Int(id)).String())
}

/*
Extracts the UID value from the given [code]uid://[/code] string.
*/
func TextToId(text_id string) int {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return int(int(class(self).TextToId(gc.String(text_id))))
}

/*
Generates a random resource UID which is guaranteed to be unique within the list of currently loaded UIDs.
In order for this UID to be registered, you must call [method add_id] or [method set_id].
*/
func CreateId() int {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return int(int(class(self).CreateId()))
}

/*
Returns whether the given UID value is known to the cache.
*/
func HasId(id int) bool {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return bool(class(self).HasId(gd.Int(id)))
}

/*
Adds a new UID value which is mapped to the given resource path.
Fails with an error if the UID already exists, so be sure to check [method has_id] beforehand, or use [method set_id] instead.
*/
func AddId(id int, path string) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).AddId(gd.Int(id), gc.String(path))
}

/*
Updates the resource path of an existing UID.
Fails with an error if the UID does not exist, so be sure to check [method has_id] beforehand, or use [method add_id] instead.
*/
func SetId(id int, path string) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).SetId(gd.Int(id), gc.String(path))
}

/*
Returns the path that the given UID value refers to.
Fails with an error if the UID does not exist, so be sure to check [method has_id] beforehand.
*/
func GetIdPath(id int) string {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return string(class(self).GetIdPath(gc, gd.Int(id)).String())
}

/*
Removes a loaded UID value from the cache.
Fails with an error if the UID does not exist, so be sure to check [method has_id] beforehand.
*/
func RemoveId(id int) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).RemoveId(gd.Int(id))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func GD() class { once.Do(singleton); return self }
type class [1]classdb.ResourceUID
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Converts the given UID to a [code]uid://[/code] string value.
*/
//go:nosplit
func (self class) IdToText(ctx gd.Lifetime, id gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ResourceUID.Bind_id_to_text, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Extracts the UID value from the given [code]uid://[/code] string.
*/
//go:nosplit
func (self class) TextToId(text_id gd.String) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(text_id))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ResourceUID.Bind_text_to_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Generates a random resource UID which is guaranteed to be unique within the list of currently loaded UIDs.
In order for this UID to be registered, you must call [method add_id] or [method set_id].
*/
//go:nosplit
func (self class) CreateId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ResourceUID.Bind_create_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns whether the given UID value is known to the cache.
*/
//go:nosplit
func (self class) HasId(id gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ResourceUID.Bind_has_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Adds a new UID value which is mapped to the given resource path.
Fails with an error if the UID already exists, so be sure to check [method has_id] beforehand, or use [method set_id] instead.
*/
//go:nosplit
func (self class) AddId(id gd.Int, path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ResourceUID.Bind_add_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Updates the resource path of an existing UID.
Fails with an error if the UID does not exist, so be sure to check [method has_id] beforehand, or use [method add_id] instead.
*/
//go:nosplit
func (self class) SetId(id gd.Int, path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ResourceUID.Bind_set_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the path that the given UID value refers to.
Fails with an error if the UID does not exist, so be sure to check [method has_id] beforehand.
*/
//go:nosplit
func (self class) GetIdPath(ctx gd.Lifetime, id gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ResourceUID.Bind_get_id_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Removes a loaded UID value from the cache.
Fails with an error if the UID does not exist, so be sure to check [method has_id] beforehand.
*/
//go:nosplit
func (self class) RemoveId(id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ResourceUID.Bind_remove_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("ResourceUID", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}