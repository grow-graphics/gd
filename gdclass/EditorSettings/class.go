package EditorSettings

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
Object that holds the project-independent editor settings. These settings are generally visible in the [b]Editor > Editor Settings[/b] menu.
Property names use slash delimiters to distinguish sections. Setting values can be of any [Variant] type. It's recommended to use [code]snake_case[/code] for editor settings to be consistent with the Godot editor itself.
Accessing the settings can be done using the following methods, such as:
[codeblocks]
[gdscript]
var settings = EditorInterface.get_editor_settings()
# `settings.set("some/property", 10)` also works as this class overrides `_set()` internally.
settings.set_setting("some/property", 10)
# `settings.get("some/property")` also works as this class overrides `_get()` internally.
settings.get_setting("some/property")
var list_of_settings = settings.get_property_list()
[/gdscript]
[csharp]
EditorSettings settings = EditorInterface.Singleton.GetEditorSettings();
// `settings.set("some/property", value)` also works as this class overrides `_set()` internally.
settings.SetSetting("some/property", Value);
// `settings.get("some/property", value)` also works as this class overrides `_get()` internally.
settings.GetSetting("some/property");
Godot.Collections.Array<Godot.Collections.Dictionary> listOfSettings = settings.GetPropertyList();
[/csharp]
[/codeblocks]
[b]Note:[/b] This class shouldn't be instantiated directly. Instead, access the singleton using [method EditorInterface.get_editor_settings].

*/
type Go [1]classdb.EditorSettings

/*
Returns [code]true[/code] if the setting specified by [param name] exists, [code]false[/code] otherwise.
*/
func (self Go) HasSetting(name string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasSetting(gc.String(name)))
}

/*
Sets the [param value] of the setting specified by [param name]. This is equivalent to using [method Object.set] on the EditorSettings instance.
*/
func (self Go) SetSetting(name string, value gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetSetting(gc.String(name), value)
}

/*
Returns the value of the setting specified by [param name]. This is equivalent to using [method Object.get] on the EditorSettings instance.
*/
func (self Go) GetSetting(name string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).GetSetting(gc, gc.String(name)))
}

/*
Erases the setting whose name is specified by [param property].
*/
func (self Go) Erase(property string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Erase(gc.String(property))
}

/*
Sets the initial value of the setting specified by [param name] to [param value]. This is used to provide a value for the Revert button in the Editor Settings. If [param update_current] is true, the current value of the setting will be set to [param value] as well.
*/
func (self Go) SetInitialValue(name string, value gd.Variant, update_current bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetInitialValue(gc.StringName(name), value, update_current)
}

/*
Adds a custom property info to a property. The dictionary must contain:
- [code]name[/code]: [String] (the name of the property)
- [code]type[/code]: [int] (see [enum Variant.Type])
- optionally [code]hint[/code]: [int] (see [enum PropertyHint]) and [code]hint_string[/code]: [String]
[b]Example:[/b]
[codeblocks]
[gdscript]
var settings = EditorInterface.get_editor_settings()
settings.set("category/property_name", 0)

var property_info = {
    "name": "category/property_name",
    "type": TYPE_INT,
    "hint": PROPERTY_HINT_ENUM,
    "hint_string": "one,two,three"
}

settings.add_property_info(property_info)
[/gdscript]
[csharp]
var settings = GetEditorInterface().GetEditorSettings();
settings.Set("category/property_name", 0);

var propertyInfo = new Godot.Collections.Dictionary
{
    {"name", "category/propertyName"},
    {"type", Variant.Type.Int},
    {"hint", PropertyHint.Enum},
    {"hint_string", "one,two,three"}
};

settings.AddPropertyInfo(propertyInfo);
[/csharp]
[/codeblocks]
*/
func (self Go) AddPropertyInfo(info gd.Dictionary) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddPropertyInfo(info)
}

/*
Sets project-specific metadata with the [param section], [param key] and [param data] specified. This metadata is stored outside the project folder and therefore won't be checked into version control. See also [method get_project_metadata].
*/
func (self Go) SetProjectMetadata(section string, key string, data gd.Variant) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetProjectMetadata(gc.String(section), gc.String(key), data)
}

/*
Returns project-specific metadata for the [param section] and [param key] specified. If the metadata doesn't exist, [param default] will be returned instead. See also [method set_project_metadata].
*/
func (self Go) GetProjectMetadata(section string, key string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).GetProjectMetadata(gc, gc.String(section), gc.String(key), gc.Variant(([1]gd.Variant{}[0]))))
}

/*
Sets the list of favorite files and directories for this project.
*/
func (self Go) SetFavorites(dirs []string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetFavorites(gc.PackedStringSlice(dirs))
}

/*
Returns the list of favorite files and directories for this project.
*/
func (self Go) GetFavorites() []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetFavorites(gc).Strings(gc))
}

/*
Sets the list of recently visited folders in the file dialog for this project.
*/
func (self Go) SetRecentDirs(dirs []string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRecentDirs(gc.PackedStringSlice(dirs))
}

/*
Returns the list of recently visited folders in the file dialog for this project.
*/
func (self Go) GetRecentDirs() []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetRecentDirs(gc).Strings(gc))
}

/*
Overrides the built-in editor action [param name] with the input actions defined in [param actions_list].
*/
func (self Go) SetBuiltinActionOverride(name string, actions_list gd.ArrayOf[gdclass.InputEvent]) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBuiltinActionOverride(gc.String(name), actions_list)
}

/*
Checks if any settings with the prefix [param setting_prefix] exist in the set of changed settings. See also [method get_changed_settings].
*/
func (self Go) CheckChangedSettingsInGroup(setting_prefix string) bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).CheckChangedSettingsInGroup(gc.String(setting_prefix)))
}

/*
Gets an array of the settings which have been changed since the last save. Note that internally [code]changed_settings[/code] is cleared after a successful save, so generally the most appropriate place to use this method is when processing [constant NOTIFICATION_EDITOR_SETTINGS_CHANGED].
*/
func (self Go) GetChangedSettings() []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetChangedSettings(gc).Strings(gc))
}

/*
Marks the passed editor setting as being changed, see [method get_changed_settings]. Only settings which exist (see [method has_setting]) will be accepted.
*/
func (self Go) MarkSettingChanged(setting string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).MarkSettingChanged(gc.String(setting))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorSettings
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("EditorSettings"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Returns [code]true[/code] if the setting specified by [param name] exists, [code]false[/code] otherwise.
*/
//go:nosplit
func (self class) HasSetting(name gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSettings.Bind_has_setting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the [param value] of the setting specified by [param name]. This is equivalent to using [method Object.set] on the EditorSettings instance.
*/
//go:nosplit
func (self class) SetSetting(name gd.String, value gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(value))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSettings.Bind_set_setting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the value of the setting specified by [param name]. This is equivalent to using [method Object.get] on the EditorSettings instance.
*/
//go:nosplit
func (self class) GetSetting(ctx gd.Lifetime, name gd.String) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSettings.Bind_get_setting, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Erases the setting whose name is specified by [param property].
*/
//go:nosplit
func (self class) Erase(property gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(property))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSettings.Bind_erase, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the initial value of the setting specified by [param name] to [param value]. This is used to provide a value for the Revert button in the Editor Settings. If [param update_current] is true, the current value of the setting will be set to [param value] as well.
*/
//go:nosplit
func (self class) SetInitialValue(name gd.StringName, value gd.Variant, update_current bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(value))
	callframe.Arg(frame, update_current)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSettings.Bind_set_initial_value, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a custom property info to a property. The dictionary must contain:
- [code]name[/code]: [String] (the name of the property)
- [code]type[/code]: [int] (see [enum Variant.Type])
- optionally [code]hint[/code]: [int] (see [enum PropertyHint]) and [code]hint_string[/code]: [String]
[b]Example:[/b]
[codeblocks]
[gdscript]
var settings = EditorInterface.get_editor_settings()
settings.set("category/property_name", 0)

var property_info = {
    "name": "category/property_name",
    "type": TYPE_INT,
    "hint": PROPERTY_HINT_ENUM,
    "hint_string": "one,two,three"
}

settings.add_property_info(property_info)
[/gdscript]
[csharp]
var settings = GetEditorInterface().GetEditorSettings();
settings.Set("category/property_name", 0);

var propertyInfo = new Godot.Collections.Dictionary
{
    {"name", "category/propertyName"},
    {"type", Variant.Type.Int},
    {"hint", PropertyHint.Enum},
    {"hint_string", "one,two,three"}
};

settings.AddPropertyInfo(propertyInfo);
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) AddPropertyInfo(info gd.Dictionary)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(info))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSettings.Bind_add_property_info, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets project-specific metadata with the [param section], [param key] and [param data] specified. This metadata is stored outside the project folder and therefore won't be checked into version control. See also [method get_project_metadata].
*/
//go:nosplit
func (self class) SetProjectMetadata(section gd.String, key gd.String, data gd.Variant)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(section))
	callframe.Arg(frame, mmm.Get(key))
	callframe.Arg(frame, mmm.Get(data))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSettings.Bind_set_project_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns project-specific metadata for the [param section] and [param key] specified. If the metadata doesn't exist, [param default] will be returned instead. See also [method set_project_metadata].
*/
//go:nosplit
func (self class) GetProjectMetadata(ctx gd.Lifetime, section gd.String, key gd.String, def gd.Variant) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(section))
	callframe.Arg(frame, mmm.Get(key))
	callframe.Arg(frame, mmm.Get(def))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSettings.Bind_get_project_metadata, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the list of favorite files and directories for this project.
*/
//go:nosplit
func (self class) SetFavorites(dirs gd.PackedStringArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(dirs))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSettings.Bind_set_favorites, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the list of favorite files and directories for this project.
*/
//go:nosplit
func (self class) GetFavorites(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSettings.Bind_get_favorites, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Sets the list of recently visited folders in the file dialog for this project.
*/
//go:nosplit
func (self class) SetRecentDirs(dirs gd.PackedStringArray)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(dirs))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSettings.Bind_set_recent_dirs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the list of recently visited folders in the file dialog for this project.
*/
//go:nosplit
func (self class) GetRecentDirs(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSettings.Bind_get_recent_dirs, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Overrides the built-in editor action [param name] with the input actions defined in [param actions_list].
*/
//go:nosplit
func (self class) SetBuiltinActionOverride(name gd.String, actions_list gd.ArrayOf[gdclass.InputEvent])  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	callframe.Arg(frame, mmm.Get(actions_list))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSettings.Bind_set_builtin_action_override, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Checks if any settings with the prefix [param setting_prefix] exist in the set of changed settings. See also [method get_changed_settings].
*/
//go:nosplit
func (self class) CheckChangedSettingsInGroup(setting_prefix gd.String) bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(setting_prefix))
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSettings.Bind_check_changed_settings_in_group, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Gets an array of the settings which have been changed since the last save. Note that internally [code]changed_settings[/code] is cleared after a successful save, so generally the most appropriate place to use this method is when processing [constant NOTIFICATION_EDITOR_SETTINGS_CHANGED].
*/
//go:nosplit
func (self class) GetChangedSettings(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSettings.Bind_get_changed_settings, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Marks the passed editor setting as being changed, see [method get_changed_settings]. Only settings which exist (see [method has_setting]) will be accepted.
*/
//go:nosplit
func (self class) MarkSettingChanged(setting gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(setting))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorSettings.Bind_mark_setting_changed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self Go) OnSettingsChanged(cb func()) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("settings_changed"), gc.Callable(cb), 0)
}


func (self class) AsEditorSettings() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorSettings() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("EditorSettings", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}