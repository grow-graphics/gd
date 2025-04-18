// Package ClassDB provides methods for working with ClassDB object instances.
package ClassDB

import "unsafe"
import "sync"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
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
var _ = slices.Delete[[]struct{}, struct{}]

/*
Provides access to metadata stored for every available class.
*/
var self [1]gdclass.ClassDB
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.ClassDB)
	self = *(*[1]gdclass.ClassDB)(unsafe.Pointer(&obj))
}

/*
Returns the names of all the classes available.
*/
func GetClassList() []string { //gd:ClassDB.get_class_list
	once.Do(singleton)
	return []string(Advanced().GetClassList().Strings())
}

/*
Returns the names of all the classes that directly or indirectly inherit from [param class].
*/
func GetInheritersFromClass(class_ string) []string { //gd:ClassDB.get_inheriters_from_class
	once.Do(singleton)
	return []string(Advanced().GetInheritersFromClass(String.Name(String.New(class_))).Strings())
}

/*
Returns the parent class of [param class].
*/
func GetParentClass(class_ string) string { //gd:ClassDB.get_parent_class
	once.Do(singleton)
	return string(Advanced().GetParentClass(String.Name(String.New(class_))).String())
}

/*
Returns whether the specified [param class] is available or not.
*/
func ClassExists(class_ string) bool { //gd:ClassDB.class_exists
	once.Do(singleton)
	return bool(Advanced().ClassExists(String.Name(String.New(class_))))
}

/*
Returns whether [param inherits] is an ancestor of [param class] or not.
*/
func IsParentClass(class_ string, inherits string) bool { //gd:ClassDB.is_parent_class
	once.Do(singleton)
	return bool(Advanced().IsParentClass(String.Name(String.New(class_)), String.Name(String.New(inherits))))
}

/*
Returns [code]true[/code] if objects can be instantiated from the specified [param class], otherwise returns [code]false[/code].
*/
func CanInstantiate(class_ string) bool { //gd:ClassDB.can_instantiate
	once.Do(singleton)
	return bool(Advanced().CanInstantiate(String.Name(String.New(class_))))
}

/*
Creates an instance of [param class].
*/
func Instantiate(class_ string) any { //gd:ClassDB.instantiate
	once.Do(singleton)
	return any(Advanced().Instantiate(String.Name(String.New(class_))).Interface())
}

/*
Returns the API type of [param class]. See [enum APIType].
*/
func ClassGetApiType(class_ string) gdclass.ClassDBAPIType { //gd:ClassDB.class_get_api_type
	once.Do(singleton)
	return gdclass.ClassDBAPIType(Advanced().ClassGetApiType(String.Name(String.New(class_))))
}

/*
Returns whether [param class] or its ancestry has a signal called [param signal] or not.
*/
func ClassHasSignal(class_ string, signal string) bool { //gd:ClassDB.class_has_signal
	once.Do(singleton)
	return bool(Advanced().ClassHasSignal(String.Name(String.New(class_)), String.Name(String.New(signal))))
}

/*
Returns the [param signal] data of [param class] or its ancestry. The returned value is a [Dictionary] with the following keys: [code]args[/code], [code]default_args[/code], [code]flags[/code], [code]id[/code], [code]name[/code], [code]return: (class_name, hint, hint_string, name, type, usage)[/code].
*/
func ClassGetSignal(class_ string, signal string) SignalInfo { //gd:ClassDB.class_get_signal
	once.Do(singleton)
	return SignalInfo(gd.DictionaryAs[SignalInfo](Advanced().ClassGetSignal(String.Name(String.New(class_)), String.Name(String.New(signal)))))
}

/*
Returns an array with all the signals of [param class] or its ancestry if [param no_inheritance] is [code]false[/code]. Every element of the array is a [Dictionary] as described in [method class_get_signal].
*/
func ClassGetSignalList(class_ string, no_inheritance bool) []SignalInfo { //gd:ClassDB.class_get_signal_list
	once.Do(singleton)
	return []SignalInfo(gd.ArrayAs[[]SignalInfo](gd.InternalArray(Advanced().ClassGetSignalList(String.Name(String.New(class_)), no_inheritance))))
}

/*
Returns an array with all the signals of [param class] or its ancestry if [param no_inheritance] is [code]false[/code]. Every element of the array is a [Dictionary] as described in [method class_get_signal].
*/
func ClassGetSignalListOptions(class_ string, no_inheritance bool) []SignalInfo { //gd:ClassDB.class_get_signal_list
	once.Do(singleton)
	return []SignalInfo(gd.ArrayAs[[]SignalInfo](gd.InternalArray(Advanced().ClassGetSignalList(String.Name(String.New(class_)), no_inheritance))))
}

/*
Returns an array with all the properties of [param class] or its ancestry if [param no_inheritance] is [code]false[/code].
*/
func ClassGetPropertyList(class_ string, no_inheritance bool) []PropertyInfo { //gd:ClassDB.class_get_property_list
	once.Do(singleton)
	return []PropertyInfo(gd.ArrayAs[[]PropertyInfo](gd.InternalArray(Advanced().ClassGetPropertyList(String.Name(String.New(class_)), no_inheritance))))
}

/*
Returns an array with all the properties of [param class] or its ancestry if [param no_inheritance] is [code]false[/code].
*/
func ClassGetPropertyListOptions(class_ string, no_inheritance bool) []PropertyInfo { //gd:ClassDB.class_get_property_list
	once.Do(singleton)
	return []PropertyInfo(gd.ArrayAs[[]PropertyInfo](gd.InternalArray(Advanced().ClassGetPropertyList(String.Name(String.New(class_)), no_inheritance))))
}

/*
Returns the getter method name of [param property] of [param class].
*/
func ClassGetPropertyGetter(class_ string, property string) string { //gd:ClassDB.class_get_property_getter
	once.Do(singleton)
	return string(Advanced().ClassGetPropertyGetter(String.Name(String.New(class_)), String.Name(String.New(property))).String())
}

/*
Returns the setter method name of [param property] of [param class].
*/
func ClassGetPropertySetter(class_ string, property string) string { //gd:ClassDB.class_get_property_setter
	once.Do(singleton)
	return string(Advanced().ClassGetPropertySetter(String.Name(String.New(class_)), String.Name(String.New(property))).String())
}

/*
Returns the value of [param property] of [param object] or its ancestry.
*/
func ClassGetProperty(obj Object.Instance, property string) any { //gd:ClassDB.class_get_property
	once.Do(singleton)
	return any(Advanced().ClassGetProperty(obj, String.Name(String.New(property))).Interface())
}

/*
Sets [param property] value of [param object] to [param value].
*/
func ClassSetProperty(obj Object.Instance, property string, value any) error { //gd:ClassDB.class_set_property
	once.Do(singleton)
	return error(gd.ToError(Advanced().ClassSetProperty(obj, String.Name(String.New(property)), variant.New(value))))
}

/*
Returns the default value of [param property] of [param class] or its ancestor classes.
*/
func ClassGetPropertyDefaultValue(class_ string, property string) any { //gd:ClassDB.class_get_property_default_value
	once.Do(singleton)
	return any(Advanced().ClassGetPropertyDefaultValue(String.Name(String.New(class_)), String.Name(String.New(property))).Interface())
}

/*
Returns whether [param class] (or its ancestry if [param no_inheritance] is [code]false[/code]) has a method called [param method] or not.
*/
func ClassHasMethod(class_ string, method string, no_inheritance bool) bool { //gd:ClassDB.class_has_method
	once.Do(singleton)
	return bool(Advanced().ClassHasMethod(String.Name(String.New(class_)), String.Name(String.New(method)), no_inheritance))
}

/*
Returns whether [param class] (or its ancestry if [param no_inheritance] is [code]false[/code]) has a method called [param method] or not.
*/
func ClassHasMethodOptions(class_ string, method string, no_inheritance bool) bool { //gd:ClassDB.class_has_method
	once.Do(singleton)
	return bool(Advanced().ClassHasMethod(String.Name(String.New(class_)), String.Name(String.New(method)), no_inheritance))
}

/*
Returns the number of arguments of the method [param method] of [param class] or its ancestry if [param no_inheritance] is [code]false[/code].
*/
func ClassGetMethodArgumentCount(class_ string, method string, no_inheritance bool) int { //gd:ClassDB.class_get_method_argument_count
	once.Do(singleton)
	return int(int(Advanced().ClassGetMethodArgumentCount(String.Name(String.New(class_)), String.Name(String.New(method)), no_inheritance)))
}

/*
Returns the number of arguments of the method [param method] of [param class] or its ancestry if [param no_inheritance] is [code]false[/code].
*/
func ClassGetMethodArgumentCountOptions(class_ string, method string, no_inheritance bool) int { //gd:ClassDB.class_get_method_argument_count
	once.Do(singleton)
	return int(int(Advanced().ClassGetMethodArgumentCount(String.Name(String.New(class_)), String.Name(String.New(method)), no_inheritance)))
}

/*
Returns an array with all the methods of [param class] or its ancestry if [param no_inheritance] is [code]false[/code]. Every element of the array is a [Dictionary] with the following keys: [code]args[/code], [code]default_args[/code], [code]flags[/code], [code]id[/code], [code]name[/code], [code]return: (class_name, hint, hint_string, name, type, usage)[/code].
[b]Note:[/b] In exported release builds the debug info is not available, so the returned dictionaries will contain only method names.
*/
func ClassGetMethodList(class_ string, no_inheritance bool) []PropertyInfo { //gd:ClassDB.class_get_method_list
	once.Do(singleton)
	return []PropertyInfo(gd.ArrayAs[[]PropertyInfo](gd.InternalArray(Advanced().ClassGetMethodList(String.Name(String.New(class_)), no_inheritance))))
}

/*
Returns an array with all the methods of [param class] or its ancestry if [param no_inheritance] is [code]false[/code]. Every element of the array is a [Dictionary] with the following keys: [code]args[/code], [code]default_args[/code], [code]flags[/code], [code]id[/code], [code]name[/code], [code]return: (class_name, hint, hint_string, name, type, usage)[/code].
[b]Note:[/b] In exported release builds the debug info is not available, so the returned dictionaries will contain only method names.
*/
func ClassGetMethodListOptions(class_ string, no_inheritance bool) []PropertyInfo { //gd:ClassDB.class_get_method_list
	once.Do(singleton)
	return []PropertyInfo(gd.ArrayAs[[]PropertyInfo](gd.InternalArray(Advanced().ClassGetMethodList(String.Name(String.New(class_)), no_inheritance))))
}

/*
Returns an array with the names all the integer constants of [param class] or its ancestry.
*/
func ClassGetIntegerConstantList(class_ string, no_inheritance bool) []string { //gd:ClassDB.class_get_integer_constant_list
	once.Do(singleton)
	return []string(Advanced().ClassGetIntegerConstantList(String.Name(String.New(class_)), no_inheritance).Strings())
}

/*
Returns an array with the names all the integer constants of [param class] or its ancestry.
*/
func ClassGetIntegerConstantListOptions(class_ string, no_inheritance bool) []string { //gd:ClassDB.class_get_integer_constant_list
	once.Do(singleton)
	return []string(Advanced().ClassGetIntegerConstantList(String.Name(String.New(class_)), no_inheritance).Strings())
}

/*
Returns whether [param class] or its ancestry has an integer constant called [param name] or not.
*/
func ClassHasIntegerConstant(class_ string, name string) bool { //gd:ClassDB.class_has_integer_constant
	once.Do(singleton)
	return bool(Advanced().ClassHasIntegerConstant(String.Name(String.New(class_)), String.Name(String.New(name))))
}

/*
Returns the value of the integer constant [param name] of [param class] or its ancestry. Always returns 0 when the constant could not be found.
*/
func ClassGetIntegerConstant(class_ string, name string) int { //gd:ClassDB.class_get_integer_constant
	once.Do(singleton)
	return int(int(Advanced().ClassGetIntegerConstant(String.Name(String.New(class_)), String.Name(String.New(name)))))
}

/*
Returns whether [param class] or its ancestry has an enum called [param name] or not.
*/
func ClassHasEnum(class_ string, name string, no_inheritance bool) bool { //gd:ClassDB.class_has_enum
	once.Do(singleton)
	return bool(Advanced().ClassHasEnum(String.Name(String.New(class_)), String.Name(String.New(name)), no_inheritance))
}

/*
Returns whether [param class] or its ancestry has an enum called [param name] or not.
*/
func ClassHasEnumOptions(class_ string, name string, no_inheritance bool) bool { //gd:ClassDB.class_has_enum
	once.Do(singleton)
	return bool(Advanced().ClassHasEnum(String.Name(String.New(class_)), String.Name(String.New(name)), no_inheritance))
}

/*
Returns an array with all the enums of [param class] or its ancestry.
*/
func ClassGetEnumList(class_ string, no_inheritance bool) []string { //gd:ClassDB.class_get_enum_list
	once.Do(singleton)
	return []string(Advanced().ClassGetEnumList(String.Name(String.New(class_)), no_inheritance).Strings())
}

/*
Returns an array with all the enums of [param class] or its ancestry.
*/
func ClassGetEnumListOptions(class_ string, no_inheritance bool) []string { //gd:ClassDB.class_get_enum_list
	once.Do(singleton)
	return []string(Advanced().ClassGetEnumList(String.Name(String.New(class_)), no_inheritance).Strings())
}

/*
Returns an array with all the keys in [param enum] of [param class] or its ancestry.
*/
func ClassGetEnumConstants(class_ string, enum string, no_inheritance bool) []string { //gd:ClassDB.class_get_enum_constants
	once.Do(singleton)
	return []string(Advanced().ClassGetEnumConstants(String.Name(String.New(class_)), String.Name(String.New(enum)), no_inheritance).Strings())
}

/*
Returns an array with all the keys in [param enum] of [param class] or its ancestry.
*/
func ClassGetEnumConstantsOptions(class_ string, enum string, no_inheritance bool) []string { //gd:ClassDB.class_get_enum_constants
	once.Do(singleton)
	return []string(Advanced().ClassGetEnumConstants(String.Name(String.New(class_)), String.Name(String.New(enum)), no_inheritance).Strings())
}

/*
Returns which enum the integer constant [param name] of [param class] or its ancestry belongs to.
*/
func ClassGetIntegerConstantEnum(class_ string, name string, no_inheritance bool) string { //gd:ClassDB.class_get_integer_constant_enum
	once.Do(singleton)
	return string(Advanced().ClassGetIntegerConstantEnum(String.Name(String.New(class_)), String.Name(String.New(name)), no_inheritance).String())
}

/*
Returns which enum the integer constant [param name] of [param class] or its ancestry belongs to.
*/
func ClassGetIntegerConstantEnumOptions(class_ string, name string, no_inheritance bool) string { //gd:ClassDB.class_get_integer_constant_enum
	once.Do(singleton)
	return string(Advanced().ClassGetIntegerConstantEnum(String.Name(String.New(class_)), String.Name(String.New(name)), no_inheritance).String())
}

/*
Returns whether [param class] (or its ancestor classes if [param no_inheritance] is [code]false[/code]) has an enum called [param enum] that is a bitfield.
*/
func IsClassEnumBitfield(class_ string, enum string, no_inheritance bool) bool { //gd:ClassDB.is_class_enum_bitfield
	once.Do(singleton)
	return bool(Advanced().IsClassEnumBitfield(String.Name(String.New(class_)), String.Name(String.New(enum)), no_inheritance))
}

/*
Returns whether [param class] (or its ancestor classes if [param no_inheritance] is [code]false[/code]) has an enum called [param enum] that is a bitfield.
*/
func IsClassEnumBitfieldOptions(class_ string, enum string, no_inheritance bool) bool { //gd:ClassDB.is_class_enum_bitfield
	once.Do(singleton)
	return bool(Advanced().IsClassEnumBitfield(String.Name(String.New(class_)), String.Name(String.New(enum)), no_inheritance))
}

/*
Returns whether this [param class] is enabled or not.
*/
func IsClassEnabled(class_ string) bool { //gd:ClassDB.is_class_enabled
	once.Do(singleton)
	return bool(Advanced().IsClassEnabled(String.Name(String.New(class_))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.ClassDB

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }

/*
Returns the names of all the classes available.
*/
//go:nosplit
func (self class) GetClassList() Packed.Strings { //gd:ClassDB.get_class_list
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_get_class_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the names of all the classes that directly or indirectly inherit from [param class].
*/
//go:nosplit
func (self class) GetInheritersFromClass(class_ String.Name) Packed.Strings { //gd:ClassDB.get_inheriters_from_class
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_get_inheriters_from_class, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the parent class of [param class].
*/
//go:nosplit
func (self class) GetParentClass(class_ String.Name) String.Name { //gd:ClassDB.get_parent_class
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_get_parent_class, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns whether the specified [param class] is available or not.
*/
//go:nosplit
func (self class) ClassExists(class_ String.Name) bool { //gd:ClassDB.class_exists
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_exists, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether [param inherits] is an ancestor of [param class] or not.
*/
//go:nosplit
func (self class) IsParentClass(class_ String.Name, inherits String.Name) bool { //gd:ClassDB.is_parent_class
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(inherits)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_is_parent_class, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if objects can be instantiated from the specified [param class], otherwise returns [code]false[/code].
*/
//go:nosplit
func (self class) CanInstantiate(class_ String.Name) bool { //gd:ClassDB.can_instantiate
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_can_instantiate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Creates an instance of [param class].
*/
//go:nosplit
func (self class) Instantiate(class_ String.Name) variant.Any { //gd:ClassDB.instantiate
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_instantiate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the API type of [param class]. See [enum APIType].
*/
//go:nosplit
func (self class) ClassGetApiType(class_ String.Name) gdclass.ClassDBAPIType { //gd:ClassDB.class_get_api_type
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	var r_ret = callframe.Ret[gdclass.ClassDBAPIType](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_api_type, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether [param class] or its ancestry has a signal called [param signal] or not.
*/
//go:nosplit
func (self class) ClassHasSignal(class_ String.Name, signal String.Name) bool { //gd:ClassDB.class_has_signal
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(signal)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_has_signal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the [param signal] data of [param class] or its ancestry. The returned value is a [Dictionary] with the following keys: [code]args[/code], [code]default_args[/code], [code]flags[/code], [code]id[/code], [code]name[/code], [code]return: (class_name, hint, hint_string, name, type, usage)[/code].
*/
//go:nosplit
func (self class) ClassGetSignal(class_ String.Name, signal String.Name) Dictionary.Any { //gd:ClassDB.class_get_signal
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(signal)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_signal, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an array with all the signals of [param class] or its ancestry if [param no_inheritance] is [code]false[/code]. Every element of the array is a [Dictionary] as described in [method class_get_signal].
*/
//go:nosplit
func (self class) ClassGetSignalList(class_ String.Name, no_inheritance bool) Array.Contains[Dictionary.Any] { //gd:ClassDB.class_get_signal_list
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_signal_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Dictionary.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an array with all the properties of [param class] or its ancestry if [param no_inheritance] is [code]false[/code].
*/
//go:nosplit
func (self class) ClassGetPropertyList(class_ String.Name, no_inheritance bool) Array.Contains[Dictionary.Any] { //gd:ClassDB.class_get_property_list
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_property_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Dictionary.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the getter method name of [param property] of [param class].
*/
//go:nosplit
func (self class) ClassGetPropertyGetter(class_ String.Name, property String.Name) String.Name { //gd:ClassDB.class_get_property_getter
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(property)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_property_getter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the setter method name of [param property] of [param class].
*/
//go:nosplit
func (self class) ClassGetPropertySetter(class_ String.Name, property String.Name) String.Name { //gd:ClassDB.class_get_property_setter
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(property)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_property_setter, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the value of [param property] of [param object] or its ancestry.
*/
//go:nosplit
func (self class) ClassGetProperty(obj [1]gd.Object, property String.Name) variant.Any { //gd:ClassDB.class_get_property
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(obj[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(property)))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_property, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Sets [param property] value of [param object] to [param value].
*/
//go:nosplit
func (self class) ClassSetProperty(obj [1]gd.Object, property String.Name, value variant.Any) Error.Code { //gd:ClassDB.class_set_property
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(obj[0])[0])
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(property)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(value)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_set_property, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Error.Code(r_ret.Get())
	frame.Free()
	return ret
}

/*
Returns the default value of [param property] of [param class] or its ancestor classes.
*/
//go:nosplit
func (self class) ClassGetPropertyDefaultValue(class_ String.Name, property String.Name) variant.Any { //gd:ClassDB.class_get_property_default_value
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(property)))
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_property_default_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns whether [param class] (or its ancestry if [param no_inheritance] is [code]false[/code]) has a method called [param method] or not.
*/
//go:nosplit
func (self class) ClassHasMethod(class_ String.Name, method String.Name, no_inheritance bool) bool { //gd:ClassDB.class_has_method
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(method)))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_has_method, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the number of arguments of the method [param method] of [param class] or its ancestry if [param no_inheritance] is [code]false[/code].
*/
//go:nosplit
func (self class) ClassGetMethodArgumentCount(class_ String.Name, method String.Name, no_inheritance bool) int64 { //gd:ClassDB.class_get_method_argument_count
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(method)))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_method_argument_count, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array with all the methods of [param class] or its ancestry if [param no_inheritance] is [code]false[/code]. Every element of the array is a [Dictionary] with the following keys: [code]args[/code], [code]default_args[/code], [code]flags[/code], [code]id[/code], [code]name[/code], [code]return: (class_name, hint, hint_string, name, type, usage)[/code].
[b]Note:[/b] In exported release builds the debug info is not available, so the returned dictionaries will contain only method names.
*/
//go:nosplit
func (self class) ClassGetMethodList(class_ String.Name, no_inheritance bool) Array.Contains[Dictionary.Any] { //gd:ClassDB.class_get_method_list
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_method_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[Dictionary.Any]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an array with the names all the integer constants of [param class] or its ancestry.
*/
//go:nosplit
func (self class) ClassGetIntegerConstantList(class_ String.Name, no_inheritance bool) Packed.Strings { //gd:ClassDB.class_get_integer_constant_list
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_integer_constant_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns whether [param class] or its ancestry has an integer constant called [param name] or not.
*/
//go:nosplit
func (self class) ClassHasIntegerConstant(class_ String.Name, name String.Name) bool { //gd:ClassDB.class_has_integer_constant
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_has_integer_constant, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the value of the integer constant [param name] of [param class] or its ancestry. Always returns 0 when the constant could not be found.
*/
//go:nosplit
func (self class) ClassGetIntegerConstant(class_ String.Name, name String.Name) int64 { //gd:ClassDB.class_get_integer_constant
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_integer_constant, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether [param class] or its ancestry has an enum called [param name] or not.
*/
//go:nosplit
func (self class) ClassHasEnum(class_ String.Name, name String.Name, no_inheritance bool) bool { //gd:ClassDB.class_has_enum
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_has_enum, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an array with all the enums of [param class] or its ancestry.
*/
//go:nosplit
func (self class) ClassGetEnumList(class_ String.Name, no_inheritance bool) Packed.Strings { //gd:ClassDB.class_get_enum_list
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_enum_list, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns an array with all the keys in [param enum] of [param class] or its ancestry.
*/
//go:nosplit
func (self class) ClassGetEnumConstants(class_ String.Name, enum String.Name, no_inheritance bool) Packed.Strings { //gd:ClassDB.class_get_enum_constants
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(enum)))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_enum_constants, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns which enum the integer constant [param name] of [param class] or its ancestry belongs to.
*/
//go:nosplit
func (self class) ClassGetIntegerConstantEnum(class_ String.Name, name String.Name, no_inheritance bool) String.Name { //gd:ClassDB.class_get_integer_constant_enum
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(name)))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_class_get_integer_constant_enum, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns whether [param class] (or its ancestor classes if [param no_inheritance] is [code]false[/code]) has an enum called [param enum] that is a bitfield.
*/
//go:nosplit
func (self class) IsClassEnumBitfield(class_ String.Name, enum String.Name, no_inheritance bool) bool { //gd:ClassDB.is_class_enum_bitfield
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(enum)))
	callframe.Arg(frame, no_inheritance)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_is_class_enum_bitfield, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns whether this [param class] is enabled or not.
*/
//go:nosplit
func (self class) IsClassEnabled(class_ String.Name) bool { //gd:ClassDB.is_class_enabled
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(class_)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.ClassDB.Bind_is_class_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("ClassDB", func(ptr gd.Object) any { return [1]gdclass.ClassDB{*(*gdclass.ClassDB)(unsafe.Pointer(&ptr))} })
}

type APIType = gdclass.ClassDBAPIType //gd:ClassDB.APIType

const (
	/*Native Core class type.*/
	ApiCore APIType = 0
	/*Native Editor class type.*/
	ApiEditor APIType = 1
	/*GDExtension class type.*/
	ApiExtension APIType = 2
	/*GDExtension Editor class type.*/
	ApiEditorExtension APIType = 3
	/*Unknown class type.*/
	ApiNone APIType = 4
)

type PropertyInfo struct {
	ClassName  string       `gd:"class_name"`
	Name       string       `gd:"name"`
	Hint       int          `gd:"hint"`
	HintString string       `gd:"hint_string"`
	Type       reflect.Type `gd:"type"`
	Usage      int          `gd:"usage"`
}
type SignalInfo struct {
	Name        string         `gd:"name"`
	Flags       int            `gd:"flags"`
	ID          int            `gd:"id"`
	DefaultArgs []interface{}  `gd:"default_args"`
	Args        []PropertyInfo `gd:"args"`
}
