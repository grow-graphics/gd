// Code generated by the generate package DO NOT EDIT

// Package TranslationServer provides methods for working with TranslationServer object instances.
package TranslationServer

import "unsafe"
import "sync"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Angle"
import "graphics.gd/variant/Euler"
import "graphics.gd/classdb/Translation"
import "graphics.gd/classdb/TranslationDomain"
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
The translation server is the API backend that manages all language translations.
Translations are stored in [TranslationDomain]s, which can be accessed by name. The most commonly used translation domain is the main translation domain. It always exists and can be accessed using an empty [StringName]. The translation server provides wrapper methods for accessing the main translation domain directly, without having to fetch the translation domain first. Custom translation domains are mainly for advanced usages like editor plugins. Names starting with [code]godot.[/code] are reserved for engine internals.
*/
type Instance [1]gdclass.TranslationServer

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

var self [1]gdclass.TranslationServer
var once sync.Once

func singleton() {
	obj := gd.Global.Object.GetSingleton(gd.Global.Singletons.TranslationServer)
	self = *(*[1]gdclass.TranslationServer)(unsafe.Pointer(&obj))
}

/*
Sets the locale of the project. The [param locale] string will be standardized to match known locales (e.g. [code]en-US[/code] would be matched to [code]en_US[/code]).
If translations have been loaded beforehand for the new locale, they will be applied.
*/
func SetLocale(locale string) { //gd:TranslationServer.set_locale
	once.Do(singleton)
	Advanced().SetLocale(String.New(locale))
}

/*
Returns the current locale of the project.
See also [method OS.get_locale] and [method OS.get_locale_language] to query the locale of the user system.
*/
func GetLocale() string { //gd:TranslationServer.get_locale
	once.Do(singleton)
	return string(Advanced().GetLocale().String())
}

/*
Returns the current locale of the editor.
[b]Note:[/b] When called from an exported project returns the same value as [method get_locale].
*/
func GetToolLocale() string { //gd:TranslationServer.get_tool_locale
	once.Do(singleton)
	return string(Advanced().GetToolLocale().String())
}

/*
Compares two locales and returns a similarity score between [code]0[/code] (no match) and [code]10[/code] (full match).
*/
func CompareLocales(locale_a string, locale_b string) int { //gd:TranslationServer.compare_locales
	once.Do(singleton)
	return int(int(Advanced().CompareLocales(String.New(locale_a), String.New(locale_b))))
}

/*
Returns a [param locale] string standardized to match known locales (e.g. [code]en-US[/code] would be matched to [code]en_US[/code]). If [param add_defaults] is [code]true[/code], the locale may have a default script or country added.
*/
func StandardizeLocale(locale string, add_defaults bool) string { //gd:TranslationServer.standardize_locale
	once.Do(singleton)
	return string(Advanced().StandardizeLocale(String.New(locale), add_defaults).String())
}

/*
Returns a [param locale] string standardized to match known locales (e.g. [code]en-US[/code] would be matched to [code]en_US[/code]). If [param add_defaults] is [code]true[/code], the locale may have a default script or country added.
*/
func StandardizeLocaleOptions(locale string, add_defaults bool) string { //gd:TranslationServer.standardize_locale
	once.Do(singleton)
	return string(Advanced().StandardizeLocale(String.New(locale), add_defaults).String())
}

/*
Returns array of known language codes.
*/
func GetAllLanguages() []string { //gd:TranslationServer.get_all_languages
	once.Do(singleton)
	return []string(Advanced().GetAllLanguages().Strings())
}

/*
Returns a readable language name for the [param language] code.
*/
func GetLanguageName(language string) string { //gd:TranslationServer.get_language_name
	once.Do(singleton)
	return string(Advanced().GetLanguageName(String.New(language)).String())
}

/*
Returns an array of known script codes.
*/
func GetAllScripts() []string { //gd:TranslationServer.get_all_scripts
	once.Do(singleton)
	return []string(Advanced().GetAllScripts().Strings())
}

/*
Returns a readable script name for the [param script] code.
*/
func GetScriptName(script string) string { //gd:TranslationServer.get_script_name
	once.Do(singleton)
	return string(Advanced().GetScriptName(String.New(script)).String())
}

/*
Returns an array of known country codes.
*/
func GetAllCountries() []string { //gd:TranslationServer.get_all_countries
	once.Do(singleton)
	return []string(Advanced().GetAllCountries().Strings())
}

/*
Returns a readable country name for the [param country] code.
*/
func GetCountryName(country string) string { //gd:TranslationServer.get_country_name
	once.Do(singleton)
	return string(Advanced().GetCountryName(String.New(country)).String())
}

/*
Returns a locale's language and its variant (e.g. [code]"en_US"[/code] would return [code]"English (United States)"[/code]).
*/
func GetLocaleName(locale string) string { //gd:TranslationServer.get_locale_name
	once.Do(singleton)
	return string(Advanced().GetLocaleName(String.New(locale)).String())
}

/*
Returns the current locale's translation for the given message and context.
[b]Note:[/b] This method always uses the main translation domain.
*/
func Translate(message string, context string) string { //gd:TranslationServer.translate
	once.Do(singleton)
	return string(Advanced().Translate(String.Name(String.New(message)), String.Name(String.New(context))).String())
}

/*
Returns the current locale's translation for the given message and context.
[b]Note:[/b] This method always uses the main translation domain.
*/
func TranslateOptions(message string, context string) string { //gd:TranslationServer.translate
	once.Do(singleton)
	return string(Advanced().Translate(String.Name(String.New(message)), String.Name(String.New(context))).String())
}

/*
Returns the current locale's translation for the given message, plural message and context.
The number [param n] is the number or quantity of the plural object. It will be used to guide the translation system to fetch the correct plural form for the selected language.
[b]Note:[/b] This method always uses the main translation domain.
*/
func TranslatePlural(message string, plural_message string, n int, context string) string { //gd:TranslationServer.translate_plural
	once.Do(singleton)
	return string(Advanced().TranslatePlural(String.Name(String.New(message)), String.Name(String.New(plural_message)), int64(n), String.Name(String.New(context))).String())
}

/*
Returns the current locale's translation for the given message, plural message and context.
The number [param n] is the number or quantity of the plural object. It will be used to guide the translation system to fetch the correct plural form for the selected language.
[b]Note:[/b] This method always uses the main translation domain.
*/
func TranslatePluralOptions(message string, plural_message string, n int, context string) string { //gd:TranslationServer.translate_plural
	once.Do(singleton)
	return string(Advanced().TranslatePlural(String.Name(String.New(message)), String.Name(String.New(plural_message)), int64(n), String.Name(String.New(context))).String())
}

/*
Adds a translation to the main translation domain.
*/
func AddTranslation(translation Translation.Instance) { //gd:TranslationServer.add_translation
	once.Do(singleton)
	Advanced().AddTranslation(translation)
}

/*
Removes the given translation from the main translation domain.
*/
func RemoveTranslation(translation Translation.Instance) { //gd:TranslationServer.remove_translation
	once.Do(singleton)
	Advanced().RemoveTranslation(translation)
}

/*
Returns the [Translation] instance that best matches [param locale] in the main translation domain. Returns [code]null[/code] if there are no matches.
*/
func GetTranslationObject(locale string) Translation.Instance { //gd:TranslationServer.get_translation_object
	once.Do(singleton)
	return Translation.Instance(Advanced().GetTranslationObject(String.New(locale)))
}

/*
Returns [code]true[/code] if a translation domain with the specified name exists.
*/
func HasDomain(domain string) bool { //gd:TranslationServer.has_domain
	once.Do(singleton)
	return bool(Advanced().HasDomain(String.Name(String.New(domain))))
}

/*
Returns the translation domain with the specified name. An empty translation domain will be created and added if it does not exist.
*/
func GetOrAddDomain(domain string) TranslationDomain.Instance { //gd:TranslationServer.get_or_add_domain
	once.Do(singleton)
	return TranslationDomain.Instance(Advanced().GetOrAddDomain(String.Name(String.New(domain))))
}

/*
Removes the translation domain with the specified name.
[b]Note:[/b] Trying to remove the main translation domain is an error.
*/
func RemoveDomain(domain string) { //gd:TranslationServer.remove_domain
	once.Do(singleton)
	Advanced().RemoveDomain(String.Name(String.New(domain)))
}

/*
Removes all translations from the main translation domain.
*/
func Clear() { //gd:TranslationServer.clear
	once.Do(singleton)
	Advanced().Clear()
}

/*
Returns an array of all loaded locales of the project.
*/
func GetLoadedLocales() []string { //gd:TranslationServer.get_loaded_locales
	once.Do(singleton)
	return []string(Advanced().GetLoadedLocales().Strings())
}

/*
Reparses the pseudolocalization options and reloads the translation for the main translation domain.
*/
func ReloadPseudolocalization() { //gd:TranslationServer.reload_pseudolocalization
	once.Do(singleton)
	Advanced().ReloadPseudolocalization()
}

/*
Returns the pseudolocalized string based on the [param message] passed in.
[b]Note:[/b] This method always uses the main translation domain.
*/
func Pseudolocalize(message string) string { //gd:TranslationServer.pseudolocalize
	once.Do(singleton)
	return string(Advanced().Pseudolocalize(String.Name(String.New(message))).String())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func Advanced() class { once.Do(singleton); return self }

type class [1]gdclass.TranslationServer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }

func PseudolocalizationEnabled() bool {
	once.Do(singleton)
	return bool(class(self).IsPseudolocalizationEnabled())
}

func SetPseudolocalizationEnabled(value bool) {
	once.Do(singleton)
	class(self).SetPseudolocalizationEnabled(value)
}

/*
Sets the locale of the project. The [param locale] string will be standardized to match known locales (e.g. [code]en-US[/code] would be matched to [code]en_US[/code]).
If translations have been loaded beforehand for the new locale, they will be applied.
*/
//go:nosplit
func (self class) SetLocale(locale String.Readable) { //gd:TranslationServer.set_locale
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(locale)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_set_locale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the current locale of the project.
See also [method OS.get_locale] and [method OS.get_locale_language] to query the locale of the user system.
*/
//go:nosplit
func (self class) GetLocale() String.Readable { //gd:TranslationServer.get_locale
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_locale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the current locale of the editor.
[b]Note:[/b] When called from an exported project returns the same value as [method get_locale].
*/
//go:nosplit
func (self class) GetToolLocale() String.Readable { //gd:TranslationServer.get_tool_locale
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_tool_locale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Compares two locales and returns a similarity score between [code]0[/code] (no match) and [code]10[/code] (full match).
*/
//go:nosplit
func (self class) CompareLocales(locale_a String.Readable, locale_b String.Readable) int64 { //gd:TranslationServer.compare_locales
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(locale_a)))
	callframe.Arg(frame, pointers.Get(gd.InternalString(locale_b)))
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_compare_locales, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns a [param locale] string standardized to match known locales (e.g. [code]en-US[/code] would be matched to [code]en_US[/code]). If [param add_defaults] is [code]true[/code], the locale may have a default script or country added.
*/
//go:nosplit
func (self class) StandardizeLocale(locale String.Readable, add_defaults bool) String.Readable { //gd:TranslationServer.standardize_locale
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(locale)))
	callframe.Arg(frame, add_defaults)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_standardize_locale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns array of known language codes.
*/
//go:nosplit
func (self class) GetAllLanguages() Packed.Strings { //gd:TranslationServer.get_all_languages
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_all_languages, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns a readable language name for the [param language] code.
*/
//go:nosplit
func (self class) GetLanguageName(language String.Readable) String.Readable { //gd:TranslationServer.get_language_name
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(language)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_language_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an array of known script codes.
*/
//go:nosplit
func (self class) GetAllScripts() Packed.Strings { //gd:TranslationServer.get_all_scripts
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_all_scripts, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns a readable script name for the [param script] code.
*/
//go:nosplit
func (self class) GetScriptName(script String.Readable) String.Readable { //gd:TranslationServer.get_script_name
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(script)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_script_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an array of known country codes.
*/
//go:nosplit
func (self class) GetAllCountries() Packed.Strings { //gd:TranslationServer.get_all_countries
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_all_countries, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns a readable country name for the [param country] code.
*/
//go:nosplit
func (self class) GetCountryName(country String.Readable) String.Readable { //gd:TranslationServer.get_country_name
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(country)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_country_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns a locale's language and its variant (e.g. [code]"en_US"[/code] would return [code]"English (United States)"[/code]).
*/
//go:nosplit
func (self class) GetLocaleName(locale String.Readable) String.Readable { //gd:TranslationServer.get_locale_name
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(locale)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_locale_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns the current locale's translation for the given message and context.
[b]Note:[/b] This method always uses the main translation domain.
*/
//go:nosplit
func (self class) Translate(message String.Name, context String.Name) String.Name { //gd:TranslationServer.translate
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(message)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(context)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_translate, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the current locale's translation for the given message, plural message and context.
The number [param n] is the number or quantity of the plural object. It will be used to guide the translation system to fetch the correct plural form for the selected language.
[b]Note:[/b] This method always uses the main translation domain.
*/
//go:nosplit
func (self class) TranslatePlural(message String.Name, plural_message String.Name, n int64, context String.Name) String.Name { //gd:TranslationServer.translate_plural
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(message)))
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(plural_message)))
	callframe.Arg(frame, n)
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(context)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_translate_plural, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Adds a translation to the main translation domain.
*/
//go:nosplit
func (self class) AddTranslation(translation [1]gdclass.Translation) { //gd:TranslationServer.add_translation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(translation[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_add_translation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes the given translation from the main translation domain.
*/
//go:nosplit
func (self class) RemoveTranslation(translation [1]gdclass.Translation) { //gd:TranslationServer.remove_translation
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(translation[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_remove_translation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the [Translation] instance that best matches [param locale] in the main translation domain. Returns [code]null[/code] if there are no matches.
*/
//go:nosplit
func (self class) GetTranslationObject(locale String.Readable) [1]gdclass.Translation { //gd:TranslationServer.get_translation_object
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(locale)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_translation_object, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Translation{gd.PointerWithOwnershipTransferredToGo[gdclass.Translation](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns [code]true[/code] if a translation domain with the specified name exists.
*/
//go:nosplit
func (self class) HasDomain(domain String.Name) bool { //gd:TranslationServer.has_domain
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(domain)))
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_has_domain, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the translation domain with the specified name. An empty translation domain will be created and added if it does not exist.
*/
//go:nosplit
func (self class) GetOrAddDomain(domain String.Name) [1]gdclass.TranslationDomain { //gd:TranslationServer.get_or_add_domain
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(domain)))
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_or_add_domain, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.TranslationDomain{gd.PointerWithOwnershipTransferredToGo[gdclass.TranslationDomain](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Removes the translation domain with the specified name.
[b]Note:[/b] Trying to remove the main translation domain is an error.
*/
//go:nosplit
func (self class) RemoveDomain(domain String.Name) { //gd:TranslationServer.remove_domain
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(domain)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_remove_domain, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes all translations from the main translation domain.
*/
//go:nosplit
func (self class) Clear() { //gd:TranslationServer.clear
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_clear, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns an array of all loaded locales of the project.
*/
//go:nosplit
func (self class) GetLoadedLocales() Packed.Strings { //gd:TranslationServer.get_loaded_locales
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.PackedPointers](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_get_loaded_locales, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Packed.Strings(Array.Through(gd.PackedStringArrayProxy{}, pointers.Pack(pointers.New[gd.PackedStringArray](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) IsPseudolocalizationEnabled() bool { //gd:TranslationServer.is_pseudolocalization_enabled
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_is_pseudolocalization_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPseudolocalizationEnabled(enabled bool) { //gd:TranslationServer.set_pseudolocalization_enabled
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_set_pseudolocalization_enabled, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Reparses the pseudolocalization options and reloads the translation for the main translation domain.
*/
//go:nosplit
func (self class) ReloadPseudolocalization() { //gd:TranslationServer.reload_pseudolocalization
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_reload_pseudolocalization, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the pseudolocalized string based on the [param message] passed in.
[b]Note:[/b] This method always uses the main translation domain.
*/
//go:nosplit
func (self class) Pseudolocalize(message String.Name) String.Name { //gd:TranslationServer.pseudolocalize
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(message)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.TranslationServer.Bind_pseudolocalize, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}
func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Instance(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("TranslationServer", func(ptr gd.Object) any {
		return [1]gdclass.TranslationServer{*(*gdclass.TranslationServer)(unsafe.Pointer(&ptr))}
	})
}
