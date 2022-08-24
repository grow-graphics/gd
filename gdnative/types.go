package gdnative

/*
	#include <gdnative_interface.h>
	#include <gdnative_functions.h>

	const char *ConstChar(_GoString_ s) {
		return _GoStringPtr(s);
	}
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

type InitializationLevel C.GDNativeInitializationLevel

var api *C.GDNativeInterface
var db C.GDNativeExtensionClassLibraryPtr

var (
	newStringName C.GDNativePtrConstructor
)

func Init(bridge *C.GDNativeInterface) {
	api = bridge
	newStringName = C.variant_get_ptr_constructor(api, C.GDNATIVE_VARIANT_TYPE_STRING_NAME, 2)
}

type Method uintptr

type (
	Nil struct {
		raw [0]byte
	}
	Bool struct {
		raw [1]byte
	}
	Int struct {
		raw [8]byte
	}
	Float struct {
		raw [8]byte
	}
	String struct {
		raw [8]byte
	}
	Vector2 struct {
		raw [8]byte
	}
	Vector2i struct {
		raw [8]byte
	}
	Rect2 struct {
		raw [16]byte
	}
	Rect2i struct {
		raw [16]byte
	}
	Vector3 struct {
		raw [12]byte
	}
	Vector3i struct {
		raw [12]byte
	}
	Transform2D struct {
		raw [24]byte
	}
	Vector4 struct {
		raw [16]byte
	}
	Vector4i struct {
		raw [16]byte
	}
	Plane struct {
		raw [16]byte
	}
	Quaternion struct {
		raw [16]byte
	}
	AABB struct {
		raw [24]byte
	}
	Basis struct {
		raw [36]byte
	}
	Transform3D struct {
		raw [48]byte
	}
	Projection struct {
		raw [64]byte
	}
	Color struct {
		raw [16]byte
	}
	StringName struct {
		raw [8]byte
	}
	NodePath struct {
		raw [8]byte
	}
	RID struct {
		raw [8]byte
	}
	/*Object struct {
		raw [8]byte
	}*/
	Callable struct {
		raw [16]byte
	}
	Signal struct {
		raw [16]byte
	}
	Dictionary struct {
		raw [8]byte
	}
	Array struct {
		raw [8]byte
	}
	PackedByteArray struct {
		raw [16]byte
	}
	PackedInt32Array struct {
		raw [16]byte
	}
	PackedInt64Array struct {
		raw [16]byte
	}
	PackedFloat32Array struct {
		raw [16]byte
	}
	PackedFloat64Array struct {
		raw [16]byte
	}
	PackedStringArray struct {
		raw [16]byte
	}
	PackedVector2Array struct {
		raw [16]byte
	}
	PackedVector3Array struct {
		raw [16]byte
	}
	PackedColorArray struct {
		raw [16]byte
	}
	Variant [24]byte
)

func NewString(s string) String {
	var result String

	C.string_new_with_utf8_chars_and_len(api, C.GDNativeStringPtr(&result.raw), C.ConstChar(s), C.GDNativeInt(len(s)))

	return result
}

func NewStringName(name string) StringName {
	var result StringName
	var s = NewString(name)

	C.constructor(newStringName, C.GDNativeTypePtr(&result.raw), (*C.GDNativeTypePtr)(unsafe.Pointer(&s)))

	return result
}

func (val Variant) CallMethod(name string, args ...any) any {

	NewStringName(name)

	//C.variant_call(api, obj.ptr, C.GDNativeStringNamePtr(C.CString(name)), C.GDNativeVariantPtrPtr(unsafe.Pointer(&args)), C.GDNativeInt(len(args)), nil, nil)

	return nil
}

func (ob Object) Call(method Method, args ...any) any {
	return nil
}

func Return[Result any](object Object, method Method, args ...any) Result {
	var result Result
	return result
}

func Call(object Object, method Method, args ...any) {
}

var loaded bool

var onload []func()

func load(intf, library, init unsafe.Pointer) {
	api = (*C.GDNativeInterface)(intf)
	ini := (*C.GDNativeInitialization)(init)
	db = (C.GDNativeExtensionClassLibraryPtr)(library)
	println(db)
	C.setInitialise(ini)
	C.setDeinitialize(ini)

	loaded = true
}

// New returns a new object of the given class name.
func New(class string) Object {
	return Object(C.classdb_construct_object(api, cString(class)))
}

// InstancePointer points to both the extension class ID
// (the first value) and the instance ID (the second value).
type InstancePointer struct {
	ClassID    ClassID
	InstanceID InstanceID
}

func (ip InstancePointer) Data() uintptr {
	var array [2]uint32
	array[0] = uint32(ip.ClassID)
	array[1] = uint32(ip.InstanceID)
	return *(*uintptr)(unsafe.Pointer(&array[0]))
}

func (ip *InstancePointer) Load(userdata uintptr) {
	var array [2]uint32
	*(*uintptr)(unsafe.Pointer(&array[0])) = userdata
	ip.ClassID = ClassID(array[0])
	ip.InstanceID = InstanceID(array[1])
}

type MethodID uint32

type MethodPointer struct {
	MethodID MethodID
	ClassID  ClassID
}

func (mp MethodPointer) Data() uintptr {
	var array [2]uint32
	array[1] = uint32(mp.ClassID)
	array[0] = uint32(mp.MethodID)
	return *(*uintptr)(unsafe.Pointer(&array[0]))
}

func (mp *MethodPointer) Load(userdata uintptr) {
	var array [2]uint32
	*(*uintptr)(unsafe.Pointer(&array[0])) = userdata
	mp.ClassID = ClassID(array[1])
	mp.MethodID = MethodID(array[0])
}

// ClassID of the extension.
type ClassID uint32

// classes that have been registered.
var classes []registeredClass

type registeredClass struct {
	Class

	methods []reflect.Method
}

func (class registeredClass) GetMethod(i MethodID) reflect.Method {
	return class.methods[i-1]
}

// InstanceID of an instantiated extension class/struct.
type InstanceID uint32

func (obj Object) SetInstance(class ClassID, instance InstanceID) {
	var ptr InstancePointer
	ptr.ClassID = class
	ptr.InstanceID = instance

	C.object_set_instance(api, C.uintptr_t(obj), cString(classes[class-1].Name()), C.uintptr_t(ptr.Data()))
}

func (obj Object) Destroy() {
	C.object_destroy(api, C.uintptr_t(obj))
}

type Object uintptr

type Class interface {
	// Name should return the name of the class. As a null
	// terminated string if possible.
	Name() string

	// Inherits returns the name of the base Godot class.
	Godot() string

	// Create should call New to instantiate an object of the parent class
	// call SetInstance on it with the ID of the instance that was created
	// and return the object that was created this way.
	Create(id ClassID) Object

	// Lookup shuold return a ptr reflect.Value of the corresponding value
	// to the given instance ID.
	Lookup(id InstanceID) reflect.Value

	// Delete should free up any resources that are being consumed by the
	// instance with the given ID.
	Delete(id InstanceID)

	// GetVirtual should return the corresponding method for the given
	// godot virtual method name.
	GetVirtual(name string) (reflect.Method, bool)

	Methods() []reflect.Method
}

func RegisterClass(class Class) {
	if !loaded {
		panic("gdnative not loaded")
	}
	methods := class.Methods()

	classes = append(classes, registeredClass{
		Class:   class,
		methods: methods,
	})

	id := ClassID(len(classes))

	C.classdb_register_extension_class(api, db, cString(class.Name()), cString(class.Godot()), C.uintptr_t(id))

	for _, method := range methods {
		registerMethod(class, id, method)
	}
}

func registerMethod(class Class, id ClassID, method reflect.Method) {
	if !loaded {
		panic("gdnative not loaded")
	}

	name := cString(method.Name)
	defer C.mem_free(api, unsafe.Pointer(name))

	var ptr = MethodPointer{
		ClassID:  id,
		MethodID: MethodID(method.Index + 1), // zero is reserved as invalid
	}

	var info = C.GDNativeExtensionClassMethodInfo{
		name:                   name,
		argument_count:         C.uint32_t(method.Type.NumIn() - 1), // exclude reciever.
		has_return_value:       C.GDNativeBool(method.Type.NumOut()),
		default_argument_count: 0, // Go doesn't have default arguments.
	}

	C.classdb_register_extension_class_method(api, db, cString(class.Name()), &info, C.uintptr_t(ptr.Data()))
}

// MethodOf arguments must be null-terminated.
func MethodOf(class, method string, hash int64) Method {
	result := Method(C.classdb_get_method_bind(api, C.ConstChar(class), C.ConstChar(method), C.GDNativeInt(hash)))
	if result == 0 {
		fmt.Println("Method not found:", class, method)
	}
	return result
}

func cString(s string) *C.char {
	if s[len(s)-1] == 0 {
		return C.ConstChar(s)
	}
	length := C.size_t(len(s) + 1) // + null byte

	mem := C.mem_alloc(api, length)

	bytes := unsafe.Slice((*byte)(mem), length)

	copy(bytes, s)
	bytes[len(s)] = 0

	return (*C.char)(mem)
}

func loadArgFromNativeType(as reflect.Type, ptr C.GDNativeTypePtr) reflect.Value {
	switch as.Kind() {
	case reflect.String:
		length := C.string_to_utf8_chars(api, C.GDNativeStringPtr(ptr), nil, 0)
		buffer := make([]byte, length)
		C.string_to_utf8_chars(api, C.GDNativeStringPtr(ptr), (*C.char)(unsafe.Pointer(&buffer[0])), length)
		return reflect.ValueOf(string(buffer)) // FIXME avoid copy?
	default:
		panic("loadArgFromNativeType: unsupported type " + as.String())
	}
}

func loadArgFromVariant(as reflect.Type, ptr C.GDNativeVariantPtr) reflect.Value {
	switch as.Kind() {
	case reflect.String:
		var native String
		C.variant_to_type_constructor(api, C.GDNATIVE_VARIANT_TYPE_STRING, C.GDNativeTypePtr(&native.raw), ptr)
		return loadArgFromNativeType(as, C.GDNativeTypePtr(&native.raw))
	default:
		panic("loadArgFromVariant: unsupported type " + as.String())
	}
}
