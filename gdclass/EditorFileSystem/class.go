package EditorFileSystem

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/Node"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
This object holds information of all resources in the filesystem, their types, etc.
[b]Note:[/b] This class shouldn't be instantiated directly. Instead, access the singleton using [method EditorInterface.get_resource_filesystem].

*/
type Go [1]classdb.EditorFileSystem

/*
Gets the root directory object.
*/
func (self Go) GetFilesystem() gdclass.EditorFileSystemDirectory {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.EditorFileSystemDirectory(class(self).GetFilesystem(gc))
}

/*
Returns [code]true[/code] if the filesystem is being scanned.
*/
func (self Go) IsScanning() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsScanning())
}

/*
Returns the scan progress for 0 to 1 if the FS is being scanned.
*/
func (self Go) GetScanningProgress() float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetScanningProgress()))
}

/*
Scan the filesystem for changes.
*/
func (self Go) Scan() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Scan()
}

/*
Check if the source of any imported resource changed.
*/
func (self Go) ScanSources() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ScanSources()
}

/*
Add a file in an existing directory, or schedule file information to be updated on editor restart. Can be used to update text files saved by an external program.
This will not import the file. To reimport, call [method reimport_files] or [method scan] methods.
*/
func (self Go) UpdateFile(path string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).UpdateFile(gc.String(path))
}

/*
Returns a view into the filesystem at [param path].
*/
func (self Go) GetFilesystemPath(path string) gdclass.EditorFileSystemDirectory {
	gc := gd.GarbageCollector(); _ = gc
	return gdclass.EditorFileSystemDirectory(class(self).GetFilesystemPath(gc, gc.String(path)))
}

/*
Returns the resource type of the file, given the full path. This returns a string such as [code]"Resource"[/code] or [code]"GDScript"[/code], [i]not[/i] a file extension such as [code]".gd"[/code].
*/
func (self Go) GetFileType(path string) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetFileType(gc, gc.String(path)).String())
}

/*
Reimports a set of files. Call this if these files or their [code].import[/code] files were directly edited by script or an external program.
If the file type changed or the file was newly created, use [method update_file] or [method scan].
[b]Note:[/b] This function blocks until the import is finished. However, the main loop iteration, including timers and [method Node._process], will occur during the import process due to progress bar updates. Avoid calls to [method reimport_files] or [method scan] while an import is in progress.
*/
func (self Go) ReimportFiles(files []string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ReimportFiles(gc.PackedStringSlice(files))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorFileSystem
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("EditorFileSystem"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Gets the root directory object.
*/
//go:nosplit
func (self class) GetFilesystem(ctx gd.Lifetime) gdclass.EditorFileSystemDirectory {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystem.Bind_get_filesystem, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.EditorFileSystemDirectory
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the filesystem is being scanned.
*/
//go:nosplit
func (self class) IsScanning() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystem.Bind_is_scanning, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the scan progress for 0 to 1 if the FS is being scanned.
*/
//go:nosplit
func (self class) GetScanningProgress() gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystem.Bind_get_scanning_progress, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Scan the filesystem for changes.
*/
//go:nosplit
func (self class) Scan()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystem.Bind_scan, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Check if the source of any imported resource changed.
*/
//go:nosplit
func (self class) ScanSources()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystem.Bind_scan_sources, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Add a file in an existing directory, or schedule file information to be updated on editor restart. Can be used to update text files saved by an external program.
This will not import the file. To reimport, call [method reimport_files] or [method scan] methods.
*/
//go:nosplit
func (self class) UpdateFile(path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystem.Bind_update_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns a view into the filesystem at [param path].
*/
//go:nosplit
func (self class) GetFilesystemPath(ctx gd.Lifetime, path gd.String) gdclass.EditorFileSystemDirectory {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystem.Bind_get_filesystem_path, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.EditorFileSystemDirectory
	ret[0].SetPointer(gd.PointerLifetimeBoundTo(ctx, self.AsObject(), r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the resource type of the file, given the full path. This returns a string such as [code]"Resource"[/code] or [code]"GDScript"[/code], [i]not[/i] a file extension such as [code]".gd"[/code].
*/
//go:nosplit
func (self class) GetFileType(ctx gd.Lifetime, path gd.String) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystem.Bind_get_file_type, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Reimports a set of files. Call this if these files or their [code].import[/code] files were directly edited by script or an external program.
If the file type changed or the file was newly created, use [method update_file] or [method scan].
[b]Note:[/b] This function blocks until the import is finished. However, the main loop iteration, including timers and [method Node._process], will occur during the import process due to progress bar updates. Avoid calls to [method reimport_files] or [method scan] while an import is in progress.
*/
//go:nosplit
func (self class) ReimportFiles(files gd.PackedStringArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(files))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorFileSystem.Bind_reimport_files, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnFilesystemChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("filesystem_changed"), gc.Callable(cb), 0)
}


func (self Go) OnScriptClassesUpdated(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("script_classes_updated"), gc.Callable(cb), 0)
}


func (self Go) OnSourcesChanged(cb func(exist bool)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("sources_changed"), gc.Callable(cb), 0)
}


func (self Go) OnResourcesReimporting(cb func(resources []string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("resources_reimporting"), gc.Callable(cb), 0)
}


func (self Go) OnResourcesReimported(cb func(resources []string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("resources_reimported"), gc.Callable(cb), 0)
}


func (self Go) OnResourcesReload(cb func(resources []string)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("resources_reload"), gc.Callable(cb), 0)
}


func (self class) AsEditorFileSystem() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorFileSystem() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.GD { return *((*Node.GD)(unsafe.Pointer(&self))) }
func (self Go) AsNode() Node.Go { return *((*Node.Go)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsNode(), name)
	}
}
func init() {classdb.Register("EditorFileSystem", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}