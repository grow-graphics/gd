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
	"runtime/cgo"
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

type Method struct {
	ptr C.GDNativeMethodBindPtr
}

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
	Variant struct {
		raw [24]byte
	}
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

	for _, f := range onload {
		f()
	}
}

// New returns a new object of the given class name.
func New(class string) Object {
	return Object(C.classdb_construct_object(api, C.ConstChar(class+"\000")))
}

// InstanceID of an instantiated extension class/struct.
type InstanceID uintptr

func (obj Object) SetInstance(class string, id InstanceID) {
	C.object_set_instance(api, C.uintptr_t(obj), C.ConstChar(class+"\000"), C.uintptr_t(id))
}

func (obj Object) Destroy() {
	C.object_destroy(api, C.uintptr_t(obj))
}

type Object uintptr

type Class interface {

	// Create should call New to instantiate an object of the parent class
	// call SetInstance on it with the ID of the instance that was created
	// and return the object that was created this way.
	Create() Object

	// Delete should free up any resources that are being consumed by the
	// instance with the given ID.
	Delete(id InstanceID)
}

func RegisterClass(name, parent string, class Class) {
	if !loaded {
		onload = append(onload, func() {
			RegisterClass(name, parent, class)
		})
		return
	}
	C.classdb_register_extension_class(api, db, C.ConstChar(name+"\000"), C.ConstChar(parent+"\000"), C.uintptr_t(cgo.NewHandle(class)))
}
