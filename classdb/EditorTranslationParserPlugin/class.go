// Code generated by the generate package DO NOT EDIT

// Package EditorTranslationParserPlugin provides methods for working with EditorTranslationParserPlugin object instances.
package EditorTranslationParserPlugin

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Angle"
import "graphics.gd/variant/Euler"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"

var _ Object.ID

type _ gdclass.Node

var _ gd.Object
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ Angle.Radians
var _ Euler.Radians
var _ = slices.Delete[[]struct{}, struct{}]

/*
ID is a typed object ID (reference) to an instance of this class, use it to store references to objects with
unknown lifetimes, as an ID will not panic on use if the underlying object has been destroyed.
*/
type ID Object.ID

func (id ID) Instance() (Instance, bool) { return Object.As[Instance](Object.ID(id).Instance()) }

/*
Extension can be embedded in a new struct to create an extension of this class.
T should be the type that is embedding this [Extension]
*/
type Extension[T gdclass.Interface] struct{ gdclass.Extension[T, Instance] }

/*
[EditorTranslationParserPlugin] is invoked when a file is being parsed to extract strings that require translation. To define the parsing and string extraction logic, override the [method _parse_file] method in script.
The return value should be an [Array] of [PackedStringArray]s, one for each extracted translatable string. Each entry should contain [code][msgid, msgctxt, msgid_plural, comment][/code], where all except [code]msgid[/code] are optional. Empty strings will be ignored.
The extracted strings will be written into a POT file selected by user under "POT Generation" in "Localization" tab in "Project Settings" menu.
Below shows an example of a custom parser that extracts strings from a CSV file to write into a POT.
[codeblocks]
[gdscript]
@tool
extends EditorTranslationParserPlugin

func _parse_file(path):

	var ret: Array[PackedStringArray] = []
	var file = FileAccess.open(path, FileAccess.READ)
	var text = file.get_as_text()
	var split_strs = text.split(",", false)
	for s in split_strs:
	    ret.append(PackedStringArray([s]))
	    #print("Extracted string: " + s)

	return ret

func _get_recognized_extensions():

	return ["csv"]

[/gdscript]
[csharp]
using Godot;

[Tool]
public partial class CustomParser : EditorTranslationParserPlugin

	{
	    public override Godot.Collections.Array<string[]> _ParseFile(string path)
	    {
	        Godot.Collections.Array<string[]> ret;
	        using var file = FileAccess.Open(path, FileAccess.ModeFlags.Read);
	        string text = file.GetAsText();
	        string[] splitStrs = text.Split(",", allowEmpty: false);
	        foreach (string s in splitStrs)
	        {
	            ret.Add([s]);
	            //GD.Print($"Extracted string: {s}");
	        }
	        return ret;
	    }

	    public override string[] _GetRecognizedExtensions()
	    {
	        return ["csv"];
	    }
	}

[/csharp]
[/codeblocks]
To add a translatable string associated with a context, plural, or comment:
[codeblocks]
[gdscript]
# This will add a message with msgid "Test 1", msgctxt "context", msgid_plural "test 1 plurals", and comment "test 1 comment".
ret.append(PackedStringArray(["Test 1", "context", "test 1 plurals", "test 1 comment"]))
# This will add a message with msgid "A test without context" and msgid_plural "plurals".
ret.append(PackedStringArray(["A test without context", "", "plurals"]))
# This will add a message with msgid "Only with context" and msgctxt "a friendly context".
ret.append(PackedStringArray(["Only with context", "a friendly context"]))
[/gdscript]
[csharp]
// This will add a message with msgid "Test 1", msgctxt "context", msgid_plural "test 1 plurals", and comment "test 1 comment".
ret.Add(["Test 1", "context", "test 1 plurals", "test 1 comment"]);
// This will add a message with msgid "A test without context" and msgid_plural "plurals".
ret.Add(["A test without context", "", "plurals"]);
// This will add a message with msgid "Only with context" and msgctxt "a friendly context".
ret.Add(["Only with context", "a friendly context"]);
[/csharp]
[/codeblocks]
[b]Note:[/b] If you override parsing logic for standard script types (GDScript, C#, etc.), it would be better to load the [code]path[/code] argument using [method ResourceLoader.load]. This is because built-in scripts are loaded as [Resource] type, not [FileAccess] type. For example:
[codeblocks]
[gdscript]
func _parse_file(path):

	var res = ResourceLoader.load(path, "Script")
	var text = res.source_code
	# Parsing logic.

func _get_recognized_extensions():

	return ["gd"]

[/gdscript]
[csharp]
public override Godot.Collections.Array<string[]> _ParseFile(string path)

	{
	    var res = ResourceLoader.Load<Script>(path, "Script");
	    string text = res.SourceCode;
	    // Parsing logic.
	}

public override string[] _GetRecognizedExtensions()

	{
	    return ["gd"];
	}

[/csharp]
[/codeblocks]
To use [EditorTranslationParserPlugin], register it using the [method EditorPlugin.add_translation_parser_plugin] method first.

	See [Interface] for methods that can be overridden by a [Class] that extends it.
*/
type Instance [1]gdclass.EditorTranslationParserPlugin

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsEditorTranslationParserPlugin() Instance
}
type Interface interface {
	//Override this method to define a custom parsing logic to extract the translatable strings.
	ParseFile(path string) [][]string
	//Gets the list of file extensions to associate with this parser, e.g. [code]["csv"][/code].
	GetRecognizedExtensions() []string
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) ParseFile(path string) (_ [][]string)  { return }
func (self implementation) GetRecognizedExtensions() (_ []string) { return }

/*
Override this method to define a custom parsing logic to extract the translatable strings.
*/
func (Instance) _parse_file(impl func(ptr unsafe.Pointer, path string) [][]string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(path))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path.String())
		ptr, ok := pointers.End(gd.InternalArray(gd.ArrayFromSlice[Array.Contains[Packed.Strings]](ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Gets the list of file extensions to associate with this parser, e.g. [code]["csv"][/code].
*/
func (Instance) _get_recognized_extensions(impl func(ptr unsafe.Pointer) []string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalPackedStrings(Packed.MakeStrings(ret...)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.EditorTranslationParserPlugin

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("EditorTranslationParserPlugin"))
	casted := Instance{*(*gdclass.EditorTranslationParserPlugin)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Override this method to define a custom parsing logic to extract the translatable strings.
*/
func (class) _parse_file(impl func(ptr unsafe.Pointer, path String.Readable) Array.Contains[Packed.Strings]) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var path = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(path))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, path)
		ptr, ok := pointers.End(gd.InternalArray(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

/*
Gets the list of file extensions to associate with this parser, e.g. [code]["csv"][/code].
*/
func (class) _get_recognized_extensions(impl func(ptr unsafe.Pointer) Packed.Strings) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalPackedStrings(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (self class) AsEditorTranslationParserPlugin() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsEditorTranslationParserPlugin() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsEditorTranslationParserPlugin() Instance {
	return self.Super().AsEditorTranslationParserPlugin()
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsRefCounted() [1]gd.RefCounted { return self.Super().AsRefCounted() }
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_parse_file":
		return reflect.ValueOf(self._parse_file)
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	default:
		return gd.VirtualByName(RefCounted.Advanced(self.AsRefCounted()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_parse_file":
		return reflect.ValueOf(self._parse_file)
	case "_get_recognized_extensions":
		return reflect.ValueOf(self._get_recognized_extensions)
	default:
		return gd.VirtualByName(RefCounted.Instance(self.AsRefCounted()), name)
	}
}
func init() {
	gdclass.Register("EditorTranslationParserPlugin", func(ptr gd.Object) any {
		return [1]gdclass.EditorTranslationParserPlugin{*(*gdclass.EditorTranslationParserPlugin)(unsafe.Pointer(&ptr))}
	})
}
