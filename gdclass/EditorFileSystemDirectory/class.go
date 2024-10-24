package EditorFileSystemDirectory

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
A more generalized, low-level variation of the directory concept.

*/
type Go [1]classdb.EditorFileSystemDirectory

/*
Returns the number of subdirectories in this directory.
*/
func (self Go) GetSubdirCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetSubdirCount()))
}

/*
Returns the subdirectory at index [param idx].
*/
func (self Go) GetSubdir(idx int) gdclass.EditorFileSystemDirectory {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.EditorFileSystemDirectory(class(self).GetSubdir(gc, gd.Int(idx)))
}

/*
Returns the number of files in this directory.
*/
func (self Go) GetFileCount() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetFileCount()))
}

/*
Returns the name of the file at index [param idx].
*/
func (self Go) GetFile(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetFile(gc, gd.Int(idx)).String())
}

/*
Returns the path to the file at index [param idx].
*/
func (self Go) GetFilePath(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetFilePath(gc, gd.Int(idx)).String())
}

/*
Returns the resource type of the file at index [param idx]. This returns a string such as [code]"Resource"[/code] or [code]"GDScript"[/code], [i]not[/i] a file extension such as [code]".gd"[/code].
*/
func (self Go) GetFileType(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetFileType(gc, gd.Int(idx)).String())
}

/*
Returns the name of the script class defined in the file at index [param idx]. If the file doesn't define a script class using the [code]class_name[/code] syntax, this will return an empty string.
*/
func (self Go) GetFileScriptClassName(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetFileScriptClassName(gc, gd.Int(idx)).String())
}

/*
Returns the base class of the script class defined in the file at index [param idx]. If the file doesn't define a script class using the [code]class_name[/code] syntax, this will return an empty string.
*/
func (self Go) GetFileScriptClassExtends(idx int) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetFileScriptClassExtends(gc, gd.Int(idx)).String())
}

/*
Returns [code]true[/code] if the file at index [param idx] imported properly.
*/
func (self Go) GetFileImportIsValid(idx int) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).GetFileImportIsValid(gd.Int(idx)))
}

/*
Returns the name of this directory.
*/
func (self Go) GetName() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetName(gc).String())
}

/*
Returns the path to this directory.
*/
func (self Go) GetPath() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetPath(gc).String())
}

/*
Returns the parent directory for this directory or [code]null[/code] if called on a directory at [code]res://[/code] or [code]user://[/code].
*/
func (self Go) GetParent() gdclass.EditorFileSystemDirectory {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.EditorFileSystemDirectory(class(self).GetParent(gc))
}

/*
Returns the index of the file with name [param name] or [code]-1[/code] if not found.
*/
func (self Go) FindFileIndex(name string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).FindFileIndex(gc.String(name))))
}

/*
Returns the index of the directory with name [param name] or [code]-1[/code] if not found.
*/
func (self Go) FindDirIndex(name string) int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).FindDirIndex(gc.String(name))))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorFileSystemDirectory
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("EditorFileSystemDirectory"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Returns the number of subdirectories in this directory.
*/
//go:nosplit
func (self class) GetSubdirCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_subdir_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the subdirectory at index [param idx].
*/
//go:nosplit
func (self class) GetSubdir(ctx gd.Lifetime, idx gd.Int) gdclass.EditorFileSystemDirectory {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_subdir, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.EditorFileSystemDirectory
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the number of files in this directory.
*/
//go:nosplit
func (self class) GetFileCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_file_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the name of the file at index [param idx].
*/
//go:nosplit
func (self class) GetFile(ctx gd.Lifetime, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the path to the file at index [param idx].
*/
//go:nosplit
func (self class) GetFilePath(ctx gd.Lifetime, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_file_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the resource type of the file at index [param idx]. This returns a string such as [code]"Resource"[/code] or [code]"GDScript"[/code], [i]not[/i] a file extension such as [code]".gd"[/code].
*/
//go:nosplit
func (self class) GetFileType(ctx gd.Lifetime, idx gd.Int) gd.StringName {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_file_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.StringName](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the name of the script class defined in the file at index [param idx]. If the file doesn't define a script class using the [code]class_name[/code] syntax, this will return an empty string.
*/
//go:nosplit
func (self class) GetFileScriptClassName(ctx gd.Lifetime, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_file_script_class_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the base class of the script class defined in the file at index [param idx]. If the file doesn't define a script class using the [code]class_name[/code] syntax, this will return an empty string.
*/
//go:nosplit
func (self class) GetFileScriptClassExtends(ctx gd.Lifetime, idx gd.Int) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_file_script_class_extends, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the file at index [param idx] imported properly.
*/
//go:nosplit
func (self class) GetFileImportIsValid(idx gd.Int) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, idx)
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_file_import_is_valid, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the name of this directory.
*/
//go:nosplit
func (self class) GetName(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_name, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the path to this directory.
*/
//go:nosplit
func (self class) GetPath(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the parent directory for this directory or [code]null[/code] if called on a directory at [code]res://[/code] or [code]user://[/code].
*/
//go:nosplit
func (self class) GetParent(ctx gd.Lifetime) gdclass.EditorFileSystemDirectory {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_get_parent, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.EditorFileSystemDirectory
	ret[0].SetPointer(gd.PointerMustAssertInstanceID(ctx, r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the index of the file with name [param name] or [code]-1[/code] if not found.
*/
//go:nosplit
func (self class) FindFileIndex(name gd.String) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_find_file_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the index of the directory with name [param name] or [code]-1[/code] if not found.
*/
//go:nosplit
func (self class) FindDirIndex(name gd.String) gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystemDirectory.Bind_find_dir_index, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsEditorFileSystemDirectory() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorFileSystemDirectory() Go { return *((*Go)(unsafe.Pointer(&self))) }

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
func init() {classdb.Register("EditorFileSystemDirectory", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}