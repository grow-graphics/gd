package EditorExportPlugin

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
[EditorExportPlugin]s are automatically invoked whenever the user exports the project. Their most common use is to determine what files are being included in the exported project. For each plugin, [method _export_begin] is called at the beginning of the export process and then [method _export_file] is called for each exported file.
To use [EditorExportPlugin], register it using the [method EditorPlugin.add_export_plugin] method first.
	// EditorExportPlugin methods that can be overridden by a [Class] that extends it.
	type EditorExportPlugin interface {
		//Virtual method to be overridden by the user. Called for each exported file before [method _customize_resource] and [method _customize_scene]. The arguments can be used to identify the file. [param path] is the path of the file, [param type] is the [Resource] represented by the file (e.g. [PackedScene]), and [param features] is the list of features for the export.
		//Calling [method skip] inside this callback will make the file not included in the export.
		ExportFile(path string, atype string, features []string) 
		//Virtual method to be overridden by the user. It is called when the export starts and provides all information about the export. [param features] is the list of features for the export, [param is_debug] is [code]true[/code] for debug builds, [param path] is the target path for the exported project. [param flags] is only used when running a runnable profile, e.g. when using native run on Android.
		ExportBegin(features []string, is_debug bool, path string, flags int) 
		//Virtual method to be overridden by the user. Called when the export is finished.
		ExportEnd() 
		//Return [code]true[/code] if this plugin will customize resources based on the platform and features used.
		//When enabled, [method _get_customization_configuration_hash] and [method _customize_resource] will be called and must be implemented.
		BeginCustomizeResources(platform gdclass.EditorExportPlatform, features []string) bool
		//Customize a resource. If changes are made to it, return the same or a new resource. Otherwise, return [code]null[/code].
		//The [i]path[/i] argument is only used when customizing an actual file, otherwise this means that this resource is part of another one and it will be empty.
		//Implementing this method is required if [method _begin_customize_resources] returns [code]true[/code].
		CustomizeResource(resource gdclass.Resource, path string) gdclass.Resource
		//Return [code]true[/code] if this plugin will customize scenes based on the platform and features used.
		//When enabled, [method _get_customization_configuration_hash] and [method _customize_scene] will be called and must be implemented.
		BeginCustomizeScenes(platform gdclass.EditorExportPlatform, features []string) bool
		//Customize a scene. If changes are made to it, return the same or a new scene. Otherwise, return [code]null[/code]. If a new scene is returned, it is up to you to dispose of the old one.
		//Implementing this method is required if [method _begin_customize_scenes] returns [code]true[/code].
		CustomizeScene(scene gdclass.Node, path string) gdclass.Node
		//Return a hash based on the configuration passed (for both scenes and resources). This helps keep separate caches for separate export configurations.
		//Implementing this method is required if [method _begin_customize_resources] returns [code]true[/code].
		GetCustomizationConfigurationHash() int
		//This is called when the customization process for scenes ends.
		EndCustomizeScenes() 
		//This is called when the customization process for resources ends.
		EndCustomizeResources() 
		//Return a list of export options that can be configured for this export plugin.
		//Each element in the return value is a [Dictionary] with the following keys:
		//- [code]option[/code]: A dictionary with the structure documented by [method Object.get_property_list], but all keys are optional.
		//- [code]default_value[/code]: The default value for this option.
		//- [code]update_visibility[/code]: An optional boolean value. If set to [code]true[/code], the preset will emit [signal Object.property_list_changed] when the option is changed.
		GetExportOptions(platform gdclass.EditorExportPlatform) gd.ArrayOf[gd.Dictionary]
		//Return a [Dictionary] of override values for export options, that will be used instead of user-provided values. Overridden options will be hidden from the user interface.
		//[codeblock]
		//class MyExportPlugin extends EditorExportPlugin:
		//    func _get_name() -> String:
		//        return "MyExportPlugin"
		//
		//    func _supports_platform(platform) -> bool:
		//        if platform is EditorExportPlatformPC:
		//            # Run on all desktop platforms including Windows, MacOS and Linux.
		//            return true
		//        return false
		//
		//    func _get_export_options_overrides(platform) -> Dictionary:
		//        # Override "Embed PCK" to always be enabled.
		//        return {
		//            "binary_format/embed_pck": true,
		//        }
		//[/codeblock]
		GetExportOptionsOverrides(platform gdclass.EditorExportPlatform) gd.Dictionary
		//Return [code]true[/code], if the result of [method _get_export_options] has changed and the export options of preset corresponding to [param platform] should be updated.
		ShouldUpdateExportOptions(platform gdclass.EditorExportPlatform) bool
		//Check the requirements for the given [param option] and return a non-empty warning string if they are not met.
		//[b]Note:[/b] Use [method get_option] to check the value of the export options.
		GetExportOptionWarning(platform gdclass.EditorExportPlatform, option string) string
		//Return a [PackedStringArray] of additional features this preset, for the given [param platform], should have.
		GetExportFeatures(platform gdclass.EditorExportPlatform, debug bool) []string
		//Return the name identifier of this plugin (for future identification by the exporter). The plugins are sorted by name before exporting.
		//Implementing this method is required.
		GetName() string
		//Return [code]true[/code] if the plugin supports the given [param platform].
		SupportsPlatform(platform gdclass.EditorExportPlatform) bool
		//Virtual method to be overridden by the user. This is called to retrieve the set of Android dependencies provided by this plugin. Each returned Android dependency should have the format of an Android remote binary dependency: [code]org.godot.example:my-plugin:0.0.0[/code]
		//For more information see [url=https://developer.android.com/build/dependencies?agpversion=4.1#dependency-types]Android documentation on dependencies[/url].
		//[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
		GetAndroidDependencies(platform gdclass.EditorExportPlatform, debug bool) []string
		//Virtual method to be overridden by the user. This is called to retrieve the URLs of Maven repositories for the set of Android dependencies provided by this plugin.
		//For more information see [url=https://docs.gradle.org/current/userguide/dependency_management.html#sec:maven_repo]Gradle documentation on dependency management[/url].
		//[b]Note:[/b] Google's Maven repo and the Maven Central repo are already included by default.
		//[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
		GetAndroidDependenciesMavenRepos(platform gdclass.EditorExportPlatform, debug bool) []string
		//Virtual method to be overridden by the user. This is called to retrieve the local paths of the Android libraries archive (AAR) files provided by this plugin.
		//[b]Note:[/b] Relative paths [b]must[/b] be relative to Godot's [code]res://addons/[/code] directory. For example, an AAR file located under [code]res://addons/hello_world_plugin/HelloWorld.release.aar[/code] can be returned as an absolute path using [code]res://addons/hello_world_plugin/HelloWorld.release.aar[/code] or a relative path using [code]hello_world_plugin/HelloWorld.release.aar[/code].
		//[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
		GetAndroidLibraries(platform gdclass.EditorExportPlatform, debug bool) []string
		//Virtual method to be overridden by the user. This is used at export time to update the contents of the [code]activity[/code] element in the generated Android manifest.
		//[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
		GetAndroidManifestActivityElementContents(platform gdclass.EditorExportPlatform, debug bool) string
		//Virtual method to be overridden by the user. This is used at export time to update the contents of the [code]application[/code] element in the generated Android manifest.
		//[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
		GetAndroidManifestApplicationElementContents(platform gdclass.EditorExportPlatform, debug bool) string
		//Virtual method to be overridden by the user. This is used at export time to update the contents of the [code]manifest[/code] element in the generated Android manifest.
		//[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
		GetAndroidManifestElementContents(platform gdclass.EditorExportPlatform, debug bool) string
	}

*/
type Go [1]classdb.EditorExportPlugin

/*
Virtual method to be overridden by the user. Called for each exported file before [method _customize_resource] and [method _customize_scene]. The arguments can be used to identify the file. [param path] is the path of the file, [param type] is the [Resource] represented by the file (e.g. [PackedScene]), and [param features] is the list of features for the export.
Calling [method skip] inside this callback will make the file not included in the export.
*/
func (Go) _export_file(impl func(ptr unsafe.Pointer, path string, atype string, features []string) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var path = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,0))
		var atype = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		var features = mmm.Let[gd.PackedStringArray](gc.Lifetime, gc.API, gd.UnsafeGet[[2]uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, path.String(), atype.String(), features.Strings(gc))
		gc.End()
	}
}

/*
Virtual method to be overridden by the user. It is called when the export starts and provides all information about the export. [param features] is the list of features for the export, [param is_debug] is [code]true[/code] for debug builds, [param path] is the target path for the exported project. [param flags] is only used when running a runnable profile, e.g. when using native run on Android.
*/
func (Go) _export_begin(impl func(ptr unsafe.Pointer, features []string, is_debug bool, path string, flags int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var features = mmm.Let[gd.PackedStringArray](gc.Lifetime, gc.API, gd.UnsafeGet[[2]uintptr](p_args,0))
		var is_debug = gd.UnsafeGet[bool](p_args,1)
		var path = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,2))
		var flags = gd.UnsafeGet[gd.Int](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, features.Strings(gc), is_debug, path.String(), int(flags))
		gc.End()
	}
}

/*
Virtual method to be overridden by the user. Called when the export is finished.
*/
func (Go) _export_end(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Return [code]true[/code] if this plugin will customize resources based on the platform and features used.
When enabled, [method _get_customization_configuration_hash] and [method _customize_resource] will be called and must be implemented.
*/
func (Go) _begin_customize_resources(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform, features []string) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var features = mmm.Let[gd.PackedStringArray](gc.Lifetime, gc.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, features.Strings(gc))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Customize a resource. If changes are made to it, return the same or a new resource. Otherwise, return [code]null[/code].
The [i]path[/i] argument is only used when customizing an actual file, otherwise this means that this resource is part of another one and it will be empty.
Implementing this method is required if [method _begin_customize_resources] returns [code]true[/code].
*/
func (Go) _customize_resource(impl func(ptr unsafe.Pointer, resource gdclass.Resource, path string) gdclass.Resource, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var resource gdclass.Resource
		resource[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var path = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, path.String())
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		gc.End()
	}
}

/*
Return [code]true[/code] if this plugin will customize scenes based on the platform and features used.
When enabled, [method _get_customization_configuration_hash] and [method _customize_scene] will be called and must be implemented.
*/
func (Go) _begin_customize_scenes(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform, features []string) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var features = mmm.Let[gd.PackedStringArray](gc.Lifetime, gc.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, features.Strings(gc))
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Customize a scene. If changes are made to it, return the same or a new scene. Otherwise, return [code]null[/code]. If a new scene is returned, it is up to you to dispose of the old one.
Implementing this method is required if [method _begin_customize_scenes] returns [code]true[/code].
*/
func (Go) _customize_scene(impl func(ptr unsafe.Pointer, scene gdclass.Node, path string) gdclass.Node, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var scene gdclass.Node
		scene[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var path = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, scene, path.String())
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		gc.End()
	}
}

/*
Return a hash based on the configuration passed (for both scenes and resources). This helps keep separate caches for separate export configurations.
Implementing this method is required if [method _begin_customize_resources] returns [code]true[/code].
*/
func (Go) _get_customization_configuration_hash(impl func(ptr unsafe.Pointer) int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, gd.Int(ret))
		gc.End()
	}
}

/*
This is called when the customization process for scenes ends.
*/
func (Go) _end_customize_scenes(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
This is called when the customization process for resources ends.
*/
func (Go) _end_customize_resources(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		gc.End()
	}
}

/*
Return a list of export options that can be configured for this export plugin.
Each element in the return value is a [Dictionary] with the following keys:
- [code]option[/code]: A dictionary with the structure documented by [method Object.get_property_list], but all keys are optional.
- [code]default_value[/code]: The default value for this option.
- [code]update_visibility[/code]: An optional boolean value. If set to [code]true[/code], the preset will emit [signal Object.property_list_changed] when the option is changed.
*/
func (Go) _get_export_options(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		gc.End()
	}
}

/*
Return a [Dictionary] of override values for export options, that will be used instead of user-provided values. Overridden options will be hidden from the user interface.
[codeblock]
class MyExportPlugin extends EditorExportPlugin:
    func _get_name() -> String:
        return "MyExportPlugin"

    func _supports_platform(platform) -> bool:
        if platform is EditorExportPlatformPC:
            # Run on all desktop platforms including Windows, MacOS and Linux.
            return true
        return false

    func _get_export_options_overrides(platform) -> Dictionary:
        # Override "Embed PCK" to always be enabled.
        return {
            "binary_format/embed_pck": true,
        }
[/codeblock]
*/
func (Go) _get_export_options_overrides(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform)
		gd.UnsafeSet(p_back, mmm.End(ret))
		gc.End()
	}
}

/*
Return [code]true[/code], if the result of [method _get_export_options] has changed and the export options of preset corresponding to [param platform] should be updated.
*/
func (Go) _should_update_export_options(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Check the requirements for the given [param option] and return a non-empty warning string if they are not met.
[b]Note:[/b] Use [method get_option] to check the value of the export options.
*/
func (Go) _get_export_option_warning(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform, option string) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var option = mmm.Let[gd.String](gc.Lifetime, gc.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, option.String())
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}

/*
Return a [PackedStringArray] of additional features this preset, for the given [param platform], should have.
*/
func (Go) _get_export_features(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform, debug bool) []string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var debug = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		gd.UnsafeSet(p_back, mmm.End(gc.PackedStringSlice(ret)))
		gc.End()
	}
}

/*
Return the name identifier of this plugin (for future identification by the exporter). The plugins are sorted by name before exporting.
Implementing this method is required.
*/
func (Go) _get_name(impl func(ptr unsafe.Pointer) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}

/*
Return [code]true[/code] if the plugin supports the given [param platform].
*/
func (Go) _supports_platform(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform)
		gd.UnsafeSet(p_back, ret)
		gc.End()
	}
}

/*
Virtual method to be overridden by the user. This is called to retrieve the set of Android dependencies provided by this plugin. Each returned Android dependency should have the format of an Android remote binary dependency: [code]org.godot.example:my-plugin:0.0.0[/code]
For more information see [url=https://developer.android.com/build/dependencies?agpversion=4.1#dependency-types]Android documentation on dependencies[/url].
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (Go) _get_android_dependencies(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform, debug bool) []string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var debug = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		gd.UnsafeSet(p_back, mmm.End(gc.PackedStringSlice(ret)))
		gc.End()
	}
}

/*
Virtual method to be overridden by the user. This is called to retrieve the URLs of Maven repositories for the set of Android dependencies provided by this plugin.
For more information see [url=https://docs.gradle.org/current/userguide/dependency_management.html#sec:maven_repo]Gradle documentation on dependency management[/url].
[b]Note:[/b] Google's Maven repo and the Maven Central repo are already included by default.
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (Go) _get_android_dependencies_maven_repos(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform, debug bool) []string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var debug = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		gd.UnsafeSet(p_back, mmm.End(gc.PackedStringSlice(ret)))
		gc.End()
	}
}

/*
Virtual method to be overridden by the user. This is called to retrieve the local paths of the Android libraries archive (AAR) files provided by this plugin.
[b]Note:[/b] Relative paths [b]must[/b] be relative to Godot's [code]res://addons/[/code] directory. For example, an AAR file located under [code]res://addons/hello_world_plugin/HelloWorld.release.aar[/code] can be returned as an absolute path using [code]res://addons/hello_world_plugin/HelloWorld.release.aar[/code] or a relative path using [code]hello_world_plugin/HelloWorld.release.aar[/code].
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (Go) _get_android_libraries(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform, debug bool) []string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var debug = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		gd.UnsafeSet(p_back, mmm.End(gc.PackedStringSlice(ret)))
		gc.End()
	}
}

/*
Virtual method to be overridden by the user. This is used at export time to update the contents of the [code]activity[/code] element in the generated Android manifest.
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (Go) _get_android_manifest_activity_element_contents(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform, debug bool) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var debug = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}

/*
Virtual method to be overridden by the user. This is used at export time to update the contents of the [code]application[/code] element in the generated Android manifest.
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (Go) _get_android_manifest_application_element_contents(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform, debug bool) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var debug = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}

/*
Virtual method to be overridden by the user. This is used at export time to update the contents of the [code]manifest[/code] element in the generated Android manifest.
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (Go) _get_android_manifest_element_contents(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform, debug bool) string, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		gc := gd.NewLifetime(api)
		class.SetTemporary(gc)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](gc.Lifetime, gc.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var debug = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		gd.UnsafeSet(p_back, mmm.End(gc.String(ret)))
		gc.End()
	}
}

/*
Adds a shared object or a directory containing only shared objects with the given [param tags] and destination [param path].
[b]Note:[/b] In case of macOS exports, those shared objects will be added to [code]Frameworks[/code] directory of app bundle.
In case of a directory code-sign will error if you place non code object in directory.
*/
func (self Go) AddSharedObject(path string, tags []string, target string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddSharedObject(gc.String(path), gc.PackedStringSlice(tags), gc.String(target))
}

/*
Adds a static lib from the given [param path] to the iOS project.
*/
func (self Go) AddIosProjectStaticLib(path string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddIosProjectStaticLib(gc.String(path))
}

/*
Adds a custom file to be exported. [param path] is the virtual path that can be used to load the file, [param file] is the binary data of the file.
When called inside [method _export_file] and [param remap] is [code]true[/code], the current file will not be exported, but instead remapped to this custom file. [param remap] is ignored when called in other places.
[param file] will not be imported, so consider using [method _customize_resource] to remap imported resources.
*/
func (self Go) AddFile(path string, file []byte, remap bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddFile(gc.String(path), gc.PackedByteSlice(file), remap)
}

/*
Adds a static library (*.a) or dynamic library (*.dylib, *.framework) to Linking Phase in iOS's Xcode project.
*/
func (self Go) AddIosFramework(path string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddIosFramework(gc.String(path))
}

/*
Adds a dynamic library (*.dylib, *.framework) to Linking Phase in iOS's Xcode project and embeds it into resulting binary.
[b]Note:[/b] For static libraries (*.a) works in same way as [method add_ios_framework].
[b]Note:[/b] This method should not be used for System libraries as they are already present on the device.
*/
func (self Go) AddIosEmbeddedFramework(path string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddIosEmbeddedFramework(gc.String(path))
}

/*
Adds content for iOS Property List files.
*/
func (self Go) AddIosPlistContent(plist_content string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddIosPlistContent(gc.String(plist_content))
}

/*
Adds linker flags for the iOS export.
*/
func (self Go) AddIosLinkerFlags(flags string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddIosLinkerFlags(gc.String(flags))
}

/*
Adds an iOS bundle file from the given [param path] to the exported project.
*/
func (self Go) AddIosBundleFile(path string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddIosBundleFile(gc.String(path))
}

/*
Adds a C++ code to the iOS export. The final code is created from the code appended by each active export plugin.
*/
func (self Go) AddIosCppCode(code string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddIosCppCode(gc.String(code))
}

/*
Adds file or directory matching [param path] to [code]PlugIns[/code] directory of macOS app bundle.
[b]Note:[/b] This is useful only for macOS exports.
*/
func (self Go) AddMacosPluginFile(path string) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).AddMacosPluginFile(gc.String(path))
}

/*
To be called inside [method _export_file]. Skips the current file, so it's not included in the export.
*/
func (self Go) Skip() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Skip()
}

/*
Returns the current value of an export option supplied by [method _get_export_options].
*/
func (self Go) GetOption(name string) gd.Variant {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Variant(class(self).GetOption(gc, gc.StringName(name)))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.EditorExportPlugin
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("EditorExportPlugin"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Virtual method to be overridden by the user. Called for each exported file before [method _customize_resource] and [method _customize_scene]. The arguments can be used to identify the file. [param path] is the path of the file, [param type] is the [Resource] represented by the file (e.g. [PackedScene]), and [param features] is the list of features for the export.
Calling [method skip] inside this callback will make the file not included in the export.
*/
func (class) _export_file(impl func(ptr unsafe.Pointer, path gd.String, atype gd.String, features gd.PackedStringArray) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,0))
		var atype = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		var features = mmm.Let[gd.PackedStringArray](ctx.Lifetime, ctx.API, gd.UnsafeGet[[2]uintptr](p_args,2))
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, path, atype, features)
		ctx.End()
	}
}

/*
Virtual method to be overridden by the user. It is called when the export starts and provides all information about the export. [param features] is the list of features for the export, [param is_debug] is [code]true[/code] for debug builds, [param path] is the target path for the exported project. [param flags] is only used when running a runnable profile, e.g. when using native run on Android.
*/
func (class) _export_begin(impl func(ptr unsafe.Pointer, features gd.PackedStringArray, is_debug bool, path gd.String, flags gd.Int) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var features = mmm.Let[gd.PackedStringArray](ctx.Lifetime, ctx.API, gd.UnsafeGet[[2]uintptr](p_args,0))
		var is_debug = gd.UnsafeGet[bool](p_args,1)
		var path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,2))
		var flags = gd.UnsafeGet[gd.Int](p_args,3)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self, features, is_debug, path, flags)
		ctx.End()
	}
}

/*
Virtual method to be overridden by the user. Called when the export is finished.
*/
func (class) _export_end(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Return [code]true[/code] if this plugin will customize resources based on the platform and features used.
When enabled, [method _get_customization_configuration_hash] and [method _customize_resource] will be called and must be implemented.
*/
func (class) _begin_customize_resources(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform, features gd.PackedStringArray) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var features = mmm.Let[gd.PackedStringArray](ctx.Lifetime, ctx.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, features)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Customize a resource. If changes are made to it, return the same or a new resource. Otherwise, return [code]null[/code].
The [i]path[/i] argument is only used when customizing an actual file, otherwise this means that this resource is part of another one and it will be empty.
Implementing this method is required if [method _begin_customize_resources] returns [code]true[/code].
*/
func (class) _customize_resource(impl func(ptr unsafe.Pointer, resource gdclass.Resource, path gd.String) gdclass.Resource, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var resource gdclass.Resource
		resource[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, resource, path)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}

/*
Return [code]true[/code] if this plugin will customize scenes based on the platform and features used.
When enabled, [method _get_customization_configuration_hash] and [method _customize_scene] will be called and must be implemented.
*/
func (class) _begin_customize_scenes(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform, features gd.PackedStringArray) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var features = mmm.Let[gd.PackedStringArray](ctx.Lifetime, ctx.API, gd.UnsafeGet[[2]uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, features)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Customize a scene. If changes are made to it, return the same or a new scene. Otherwise, return [code]null[/code]. If a new scene is returned, it is up to you to dispose of the old one.
Implementing this method is required if [method _begin_customize_scenes] returns [code]true[/code].
*/
func (class) _customize_scene(impl func(ptr unsafe.Pointer, scene gdclass.Node, path gd.String) gdclass.Node, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var scene gdclass.Node
		scene[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var path = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, scene, path)
		gd.UnsafeSet(p_back, mmm.End(ret[0].AsPointer()))
		ctx.End()
	}
}

/*
Return a hash based on the configuration passed (for both scenes and resources). This helps keep separate caches for separate export configurations.
Implementing this method is required if [method _begin_customize_resources] returns [code]true[/code].
*/
func (class) _get_customization_configuration_hash(impl func(ptr unsafe.Pointer) gd.Int, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
This is called when the customization process for scenes ends.
*/
func (class) _end_customize_scenes(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
This is called when the customization process for resources ends.
*/
func (class) _end_customize_resources(impl func(ptr unsafe.Pointer) , api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
impl(self)
		ctx.End()
	}
}

/*
Return a list of export options that can be configured for this export plugin.
Each element in the return value is a [Dictionary] with the following keys:
- [code]option[/code]: A dictionary with the structure documented by [method Object.get_property_list], but all keys are optional.
- [code]default_value[/code]: The default value for this option.
- [code]update_visibility[/code]: An optional boolean value. If set to [code]true[/code], the preset will emit [signal Object.property_list_changed] when the option is changed.
*/
func (class) _get_export_options(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform) gd.ArrayOf[gd.Dictionary], api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform)
		gd.UnsafeSet(p_back, mmm.End(ret.Array()))
		ctx.End()
	}
}

/*
Return a [Dictionary] of override values for export options, that will be used instead of user-provided values. Overridden options will be hidden from the user interface.
[codeblock]
class MyExportPlugin extends EditorExportPlugin:
    func _get_name() -> String:
        return "MyExportPlugin"

    func _supports_platform(platform) -> bool:
        if platform is EditorExportPlatformPC:
            # Run on all desktop platforms including Windows, MacOS and Linux.
            return true
        return false

    func _get_export_options_overrides(platform) -> Dictionary:
        # Override "Embed PCK" to always be enabled.
        return {
            "binary_format/embed_pck": true,
        }
[/codeblock]
*/
func (class) _get_export_options_overrides(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform) gd.Dictionary, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Return [code]true[/code], if the result of [method _get_export_options] has changed and the export options of preset corresponding to [param platform] should be updated.
*/
func (class) _should_update_export_options(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Check the requirements for the given [param option] and return a non-empty warning string if they are not met.
[b]Note:[/b] Use [method get_option] to check the value of the export options.
*/
func (class) _get_export_option_warning(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform, option gd.String) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var option = mmm.Let[gd.String](ctx.Lifetime, ctx.API, gd.UnsafeGet[uintptr](p_args,1))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, option)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Return a [PackedStringArray] of additional features this preset, for the given [param platform], should have.
*/
func (class) _get_export_features(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform, debug bool) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var debug = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Return the name identifier of this plugin (for future identification by the exporter). The plugins are sorted by name before exporting.
Implementing this method is required.
*/
func (class) _get_name(impl func(ptr unsafe.Pointer) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Return [code]true[/code] if the plugin supports the given [param platform].
*/
func (class) _supports_platform(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform) bool, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform)
		gd.UnsafeSet(p_back, ret)
		ctx.End()
	}
}

/*
Virtual method to be overridden by the user. This is called to retrieve the set of Android dependencies provided by this plugin. Each returned Android dependency should have the format of an Android remote binary dependency: [code]org.godot.example:my-plugin:0.0.0[/code]
For more information see [url=https://developer.android.com/build/dependencies?agpversion=4.1#dependency-types]Android documentation on dependencies[/url].
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (class) _get_android_dependencies(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform, debug bool) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var debug = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Virtual method to be overridden by the user. This is called to retrieve the URLs of Maven repositories for the set of Android dependencies provided by this plugin.
For more information see [url=https://docs.gradle.org/current/userguide/dependency_management.html#sec:maven_repo]Gradle documentation on dependency management[/url].
[b]Note:[/b] Google's Maven repo and the Maven Central repo are already included by default.
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (class) _get_android_dependencies_maven_repos(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform, debug bool) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var debug = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Virtual method to be overridden by the user. This is called to retrieve the local paths of the Android libraries archive (AAR) files provided by this plugin.
[b]Note:[/b] Relative paths [b]must[/b] be relative to Godot's [code]res://addons/[/code] directory. For example, an AAR file located under [code]res://addons/hello_world_plugin/HelloWorld.release.aar[/code] can be returned as an absolute path using [code]res://addons/hello_world_plugin/HelloWorld.release.aar[/code] or a relative path using [code]hello_world_plugin/HelloWorld.release.aar[/code].
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (class) _get_android_libraries(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform, debug bool) gd.PackedStringArray, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var debug = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Virtual method to be overridden by the user. This is used at export time to update the contents of the [code]activity[/code] element in the generated Android manifest.
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (class) _get_android_manifest_activity_element_contents(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform, debug bool) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var debug = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Virtual method to be overridden by the user. This is used at export time to update the contents of the [code]application[/code] element in the generated Android manifest.
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (class) _get_android_manifest_application_element_contents(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform, debug bool) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var debug = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Virtual method to be overridden by the user. This is used at export time to update the contents of the [code]manifest[/code] element in the generated Android manifest.
[b]Note:[/b] Only supported on Android and requires [member EditorExportPlatformAndroid.gradle_build/use_gradle_build] to be enabled.
*/
func (class) _get_android_manifest_element_contents(impl func(ptr unsafe.Pointer, platform gdclass.EditorExportPlatform, debug bool) gd.String, api *gd.API) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class gd.ExtensionClass, p_args gd.UnsafeArgs, p_back gd.UnsafeBack) {
		ctx := gd.NewLifetime(api)
		class.SetTemporary(ctx)
		var platform gdclass.EditorExportPlatform
		platform[0].SetPointer(mmm.Let[gd.Pointer](ctx.Lifetime, ctx.API, [2]uintptr{gd.UnsafeGet[uintptr](p_args,0)}))
		var debug = gd.UnsafeGet[bool](p_args,1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, platform, debug)
		gd.UnsafeSet(p_back, mmm.End(ret))
		ctx.End()
	}
}

/*
Adds a shared object or a directory containing only shared objects with the given [param tags] and destination [param path].
[b]Note:[/b] In case of macOS exports, those shared objects will be added to [code]Frameworks[/code] directory of app bundle.
In case of a directory code-sign will error if you place non code object in directory.
*/
//go:nosplit
func (self class) AddSharedObject(path gd.String, tags gd.PackedStringArray, target gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, mmm.Get(tags))
	callframe.Arg(frame, mmm.Get(target))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorExportPlugin.Bind_add_shared_object, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a static lib from the given [param path] to the iOS project.
*/
//go:nosplit
func (self class) AddIosProjectStaticLib(path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorExportPlugin.Bind_add_ios_project_static_lib, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a custom file to be exported. [param path] is the virtual path that can be used to load the file, [param file] is the binary data of the file.
When called inside [method _export_file] and [param remap] is [code]true[/code], the current file will not be exported, but instead remapped to this custom file. [param remap] is ignored when called in other places.
[param file] will not be imported, so consider using [method _customize_resource] to remap imported resources.
*/
//go:nosplit
func (self class) AddFile(path gd.String, file gd.PackedByteArray, remap bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	callframe.Arg(frame, mmm.Get(file))
	callframe.Arg(frame, remap)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorExportPlugin.Bind_add_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a static library (*.a) or dynamic library (*.dylib, *.framework) to Linking Phase in iOS's Xcode project.
*/
//go:nosplit
func (self class) AddIosFramework(path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorExportPlugin.Bind_add_ios_framework, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a dynamic library (*.dylib, *.framework) to Linking Phase in iOS's Xcode project and embeds it into resulting binary.
[b]Note:[/b] For static libraries (*.a) works in same way as [method add_ios_framework].
[b]Note:[/b] This method should not be used for System libraries as they are already present on the device.
*/
//go:nosplit
func (self class) AddIosEmbeddedFramework(path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorExportPlugin.Bind_add_ios_embedded_framework, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds content for iOS Property List files.
*/
//go:nosplit
func (self class) AddIosPlistContent(plist_content gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(plist_content))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorExportPlugin.Bind_add_ios_plist_content, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds linker flags for the iOS export.
*/
//go:nosplit
func (self class) AddIosLinkerFlags(flags gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(flags))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorExportPlugin.Bind_add_ios_linker_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds an iOS bundle file from the given [param path] to the exported project.
*/
//go:nosplit
func (self class) AddIosBundleFile(path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorExportPlugin.Bind_add_ios_bundle_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds a C++ code to the iOS export. The final code is created from the code appended by each active export plugin.
*/
//go:nosplit
func (self class) AddIosCppCode(code gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(code))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorExportPlugin.Bind_add_ios_cpp_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Adds file or directory matching [param path] to [code]PlugIns[/code] directory of macOS app bundle.
[b]Note:[/b] This is useful only for macOS exports.
*/
//go:nosplit
func (self class) AddMacosPluginFile(path gd.String)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(path))
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorExportPlugin.Bind_add_macos_plugin_file, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
To be called inside [method _export_file]. Skips the current file, so it's not included in the export.
*/
//go:nosplit
func (self class) Skip()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorExportPlugin.Bind_skip, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the current value of an export option supplied by [method _get_export_options].
*/
//go:nosplit
func (self class) GetOption(ctx gd.Lifetime, name gd.StringName) gd.Variant {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(name))
	var r_ret = callframe.Ret[[3]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.EditorExportPlugin.Bind_get_option, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Variant](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsEditorExportPlugin() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsEditorExportPlugin() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_export_file": return reflect.ValueOf(self._export_file);
	case "_export_begin": return reflect.ValueOf(self._export_begin);
	case "_export_end": return reflect.ValueOf(self._export_end);
	case "_begin_customize_resources": return reflect.ValueOf(self._begin_customize_resources);
	case "_customize_resource": return reflect.ValueOf(self._customize_resource);
	case "_begin_customize_scenes": return reflect.ValueOf(self._begin_customize_scenes);
	case "_customize_scene": return reflect.ValueOf(self._customize_scene);
	case "_get_customization_configuration_hash": return reflect.ValueOf(self._get_customization_configuration_hash);
	case "_end_customize_scenes": return reflect.ValueOf(self._end_customize_scenes);
	case "_end_customize_resources": return reflect.ValueOf(self._end_customize_resources);
	case "_get_export_options": return reflect.ValueOf(self._get_export_options);
	case "_get_export_options_overrides": return reflect.ValueOf(self._get_export_options_overrides);
	case "_should_update_export_options": return reflect.ValueOf(self._should_update_export_options);
	case "_get_export_option_warning": return reflect.ValueOf(self._get_export_option_warning);
	case "_get_export_features": return reflect.ValueOf(self._get_export_features);
	case "_get_name": return reflect.ValueOf(self._get_name);
	case "_supports_platform": return reflect.ValueOf(self._supports_platform);
	case "_get_android_dependencies": return reflect.ValueOf(self._get_android_dependencies);
	case "_get_android_dependencies_maven_repos": return reflect.ValueOf(self._get_android_dependencies_maven_repos);
	case "_get_android_libraries": return reflect.ValueOf(self._get_android_libraries);
	case "_get_android_manifest_activity_element_contents": return reflect.ValueOf(self._get_android_manifest_activity_element_contents);
	case "_get_android_manifest_application_element_contents": return reflect.ValueOf(self._get_android_manifest_application_element_contents);
	case "_get_android_manifest_element_contents": return reflect.ValueOf(self._get_android_manifest_element_contents);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	case "_export_file": return reflect.ValueOf(self._export_file);
	case "_export_begin": return reflect.ValueOf(self._export_begin);
	case "_export_end": return reflect.ValueOf(self._export_end);
	case "_begin_customize_resources": return reflect.ValueOf(self._begin_customize_resources);
	case "_customize_resource": return reflect.ValueOf(self._customize_resource);
	case "_begin_customize_scenes": return reflect.ValueOf(self._begin_customize_scenes);
	case "_customize_scene": return reflect.ValueOf(self._customize_scene);
	case "_get_customization_configuration_hash": return reflect.ValueOf(self._get_customization_configuration_hash);
	case "_end_customize_scenes": return reflect.ValueOf(self._end_customize_scenes);
	case "_end_customize_resources": return reflect.ValueOf(self._end_customize_resources);
	case "_get_export_options": return reflect.ValueOf(self._get_export_options);
	case "_get_export_options_overrides": return reflect.ValueOf(self._get_export_options_overrides);
	case "_should_update_export_options": return reflect.ValueOf(self._should_update_export_options);
	case "_get_export_option_warning": return reflect.ValueOf(self._get_export_option_warning);
	case "_get_export_features": return reflect.ValueOf(self._get_export_features);
	case "_get_name": return reflect.ValueOf(self._get_name);
	case "_supports_platform": return reflect.ValueOf(self._supports_platform);
	case "_get_android_dependencies": return reflect.ValueOf(self._get_android_dependencies);
	case "_get_android_dependencies_maven_repos": return reflect.ValueOf(self._get_android_dependencies_maven_repos);
	case "_get_android_libraries": return reflect.ValueOf(self._get_android_libraries);
	case "_get_android_manifest_activity_element_contents": return reflect.ValueOf(self._get_android_manifest_activity_element_contents);
	case "_get_android_manifest_application_element_contents": return reflect.ValueOf(self._get_android_manifest_application_element_contents);
	case "_get_android_manifest_element_contents": return reflect.ValueOf(self._get_android_manifest_element_contents);
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("EditorExportPlugin", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}