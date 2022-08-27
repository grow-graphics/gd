package gdc

/*
	#include <gdnative_interface.h>
	#include <gdnative_functions.h>

	const char *ConstChar(_GoString_ s) {
		return _GoStringPtr(s);
	}


	static void string_new_with_utf8_chars_and_len(
		GDNativeInterface *api,
		GDNativeStringPtr r_dest,
		const char *p_contents,
		const GDNativeInt p_size
	) {
		api->string_new_with_utf8_chars_and_len(r_dest, p_contents, p_size);
	}

*/
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/readykit/gd/internal/gdunsafe"
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

type cleanupType uint64

const (
	cleanupIgnore = iota
	cleanupObject
	cleanupDirect
)

type Pointer struct {
	self    *Pointer // we use this to stop badly behaved copies.
	address uintptr
	cleanup cleanupType // does this pointer need to be freed?
}

func (p Pointer) Free() {
	switch p.cleanup {
	case cleanupIgnore:
		return
	case cleanupObject:
		C.object_destroy(api, C.uintptr_t(p.address))
	case cleanupDirect:
		C.mem_free(api, C.uintptr_t(p.address))
	default:
		panic("unknown cleanup type")
	}
}

type (
	Nil struct {
		raw [0]byte
	}
	Bool    bool
	Int     int64
	Float   float64
	String  uintptr
	Vector2 struct {
		X, Y float32
	}
	Vector2i struct {
		X, Y int32
	}
	Rect2 struct {
		Position, Size Vector2
	}
	Rect2i struct {
		Position, Size Vector2i
	}
	Vector3 struct {
		X, Y, Z float32
	}
	Vector3i struct {
		X, Y, Z int32
	}
	Transform2D struct {
		X, Y, Origin Vector2
	}
	Vector4 struct {
		X, Y, Z, W float32
	}
	Vector4i struct {
		X, Y, Z, W int32
	}
	Plane struct {
		Matrix Vector4
	}
	Quaternion struct {
		X, Y, Z, W float32
	}
	AABB struct {
		Position, Size Vector3
	}
	Basis struct {
		Rows [3]Vector3
	}
	Transform3D struct {
		Basis  Basis
		Origin Vector3
	}
	Projection struct {
		X, Y, Z, W Vector4
	}
	Color struct {
		R, G, B, A float32
	}
	StringName Pointer
	NodePath   Pointer
	RID        uint64
	Object     Pointer
	Callable   [3]uintptr
	Signal     struct {
		raw [16]byte
	}
	Dictionary Pointer
	Array      struct {
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
		ref uint32
		len uint32
		ptr unsafe.Pointer
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

	C.string_new_with_utf8_chars_and_len(api, C.GDNativeStringPtr(&result), C.ConstChar(s), C.GDNativeInt(len(s)))

	return result
}

func NewStringName(name string) StringName {
	var result StringName
	var s = NewString(name)

	C.constructor(newStringName, C.GDNativeTypePtr(&result.address), (*C.GDNativeTypePtr)(unsafe.Pointer(&s)))

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

/*func Return[Result any](object Object, method Method, args ...any) Result {
	var result Result

	ptr, ok := loadResultFor(&result)

	var first *C.GDNativeTypePtr
	var ptrs = make([]C.GDNativeTypePtr, len(args))
	for i, arg := range args {
		ptrs[i] = sendArg(arg)
	}
	if len(args) > 0 {
		first = &ptrs[0]
	}

	C.object_method_bind_ptrcall(api, C.uintptr_t(method), C.uintptr_t(object), first, ptr)

	if !ok {
		reflect.ValueOf(result).Elem().Set(loadArgFromNativeType(reflect.TypeOf(result), ptr))
	}

	return result
}

func Call(object Object, method Method, args ...any) {
	var first *C.GDNativeTypePtr
	var ptrs = make([]C.GDNativeTypePtr, len(args))
	for i, arg := range args {
		ptrs[i] = sendArg(arg)
	}
	if len(args) > 0 {
		first = &ptrs[0]
	}
	C.object_method_bind_ptrcall(api, C.uintptr_t(method), C.uintptr_t(object), first, nil)
}*/

var loaded bool

var onload []func()

func load(intf, library, init unsafe.Pointer) {
	api = (*C.GDNativeInterface)(intf)
	ini := (*C.GDNativeInitialization)(init)
	db = (C.GDNativeExtensionClassLibraryPtr)(library)

	*(**C.GDNativeInterface)(unsafe.Pointer(&gdunsafe.API)) = api
	*(*C.GDNativeExtensionClassLibraryPtr)(unsafe.Pointer(&gdunsafe.ClassDB)) = db

	println(db)
	C.setInitialise(ini)
	C.setDeinitialize(ini)

	loaded = true
}

// New returns a new object of the given class name.
func New(class string) Object {
	s, ok := cString(class)
	if !ok {
		defer C.mem_free(api, C.uintptr_t(uintptr(unsafe.Pointer(s))))
	}
	obj := Object{address: uintptr(C.classdb_construct_object(api, s))}
	return obj
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

	s, ok := cString(classes[class-1].Name())
	if !ok {
		defer C.mem_free(api, C.uintptr_t(uintptr(unsafe.Pointer(s))))
	}

	C.object_set_instance(api, C.uintptr_t(obj.address), s, C.uintptr_t(ptr.Data()))
}

func (obj Object) Destroy() {
	C.object_destroy(api, C.uintptr_t(obj.address))
}

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

	a, ok := cString(class.Name())
	if !ok {
		defer C.mem_free(api, C.uintptr_t(uintptr(unsafe.Pointer(a))))
	}
	b, ok := cString(class.Godot())
	if !ok {
		defer C.mem_free(api, C.uintptr_t(uintptr(unsafe.Pointer(b))))
	}

	C.classdb_register_extension_class(api, db, a, b, C.uintptr_t(id))

	for _, method := range methods {
		registerMethod(class, id, method)
	}
}

func registerMethod(class Class, id ClassID, method reflect.Method) {
	if !loaded {
		panic("gdnative not loaded")
	}

	name, ok := cString(method.Name)
	if !ok {
		defer C.mem_free(api, C.uintptr_t(uintptr(unsafe.Pointer(name))))
	}

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

	s, ok := cString(class.Name())
	if !ok {
		defer C.mem_free(api, C.uintptr_t(uintptr(unsafe.Pointer(unsafe.Pointer(s)))))
	}

	C.classdb_register_extension_class_method(api, db, s, &info, C.uintptr_t(ptr.Data()))
}

func cString(s string) (*C.char, bool) {
	if s[len(s)-1] == 0 {
		return C.ConstChar(s), true
	}

	length := C.size_t(len(s) + 1) // + null byte

	mem := C.mem_alloc(api, length)

	bytes := unsafe.Slice((*byte)(mem), length)

	copy(bytes, s)
	bytes[len(s)] = 0

	return (*C.char)(mem), false
}

func GetSingleton(name string) Object {
	if !loaded {
		panic("gdnative not loaded")
	}

	s, ok := cString(name)
	if !ok {
		defer C.mem_free(api, C.uintptr_t(uintptr(unsafe.Pointer(unsafe.Pointer(s)))))
	}

	return Object{address: uintptr(C.global_get_singleton(api, s))}
}
