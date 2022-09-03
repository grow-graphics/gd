//go:build !generate

package gd

import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

var cOnLoad func() = func() {}

var deferred []func()

//export loadExtension
func loadExtension(p_interface, p_library, p_init unsafe.Pointer) uint8 {
	api = (*cInterface)(p_interface)
	classDB = cClassLibrary(p_library)

	ini := (*cInitialization)(p_init)
	ini.cInit()

	return 1
}

//export initialize
func initialize(userdata unsafe.Pointer, level int64) {
	if level == 4 {
		cOnLoad()

		for _, fn := range deferred {
			fn()
		}
	}
}

//export deinitialize
func deinitialize(userdata unsafe.Pointer, level int64) {}

//export create_instance_func
func create_instance_func(userdata uintptr) uintptr {
	return uintptr(extensions[extensionID(userdata)-1].create())
}

//export free_instance_func
func free_instance_func(userdata, instance uintptr) {
	extensions[extensionID(userdata)-1].delete(loadInstanceID(instance))
}

//export get_virtual_index
func get_virtual_index(userdata uintptr, p_name *C.char) uint8 {
	name := C.GoString(p_name)
	method, ok := extensions[extensionID(userdata)-1].lookup(name)
	fmt.Println("lookup", name, ok)
	if !ok {
		return 0
	}
	if method.Index > 255 {
		panic("too many exported methods")
	}
	return uint8(method.Index + 1)
}

//export call_virtual_index
func call_virtual_index(index uint8, instance uintptr, args *unsafe.Pointer, result unsafe.Pointer) {
	id := loadInstanceID(instance)
	extensions[id.class-1].call(id.index, index, args, result)
}

type Class interface {
	owner() cObject
	class() string
	virtual(reflect.Type, string) (reflect.Method, bool)
}

type Loadable interface {
	load()
}

func Load[T Loadable](class T) T {
	class.load()
	return class
}

type extensionID uint32

// Extension to a core class.
type Extension[Type any, Extends Class] struct {
	id extensionID

	vtype reflect.Type // value type.
	ptype reflect.Type // pointer type.

	loader ExtensionLoader[Type, Extends]

	class string             // base class name, cached on register.
	owner string             // owner class name, cached on register.
	slice []instanceOf[Type] // slice of instantiated instances.
}

type extensionInterface interface {
	create() cObject
	delete(instanceID)
	lookup(name string) (reflect.Method, bool)
	call(instanceID uint32, method uint8, args *unsafe.Pointer, result unsafe.Pointer)
}

var extensions []extensionInterface

type ExtensionLoader[Type any, Extends Class] func(Context, *Type) Extends

func Register[Type any, Extends Class](fn ExtensionLoader[Type, Extends]) *Extension[Type, Extends] {
	var zero Type
	var parent Extends

	var rtype = reflect.TypeOf(zero)

	var extension = Extension[Type, Extends]{
		loader: fn,

		vtype: rtype,
		ptype: reflect.PtrTo(rtype),
		class: rtype.Name() + "\000",
		owner: parent.class(),
	}

	extensions = append(extensions, &extension)
	extension.id = extensionID(len(extensions))

	deferred = append(deferred, func() {
		classDB.register_extension_class(extension.class, extension.owner, uintptr(extension.id))
	})

	return &extension
}

type instanceID struct {
	class extensionID
	index uint32
}

func loadInstanceID(id uintptr) instanceID {
	return instanceID{
		class: extensionID(id >> 32),
		index: uint32(id),
	}
}

func (id instanceID) pack() uintptr {
	return uintptr(id.class)<<32 | uintptr(id.index)
}

var extensionIDs uint32

func nextExtensionID() uint32 {
	extensionIDs++
	return extensionIDs
}

type instanceOf[Type any] struct {
	alive bool
	alloc *uint32
	value Type
}

func (i *instanceOf[T]) context() Context {
	if i.alloc == nil {
		i.alloc = new(uint32)
	}
	return Context{
		counter: i.alloc,
	}
}

func (ext *Extension[Type, Extends]) lookup(name string) (reflect.Method, bool) {
	var zero Type
	var core Extends
	return core.virtual(reflect.PtrTo(reflect.TypeOf(zero)), name)
}

func (ext *Extension[Type, Extends]) create() cObject {
	var id instanceID
	id.class = ext.id

	// reuse existing slot if possible.
	for i, slot := range ext.slice {
		if !slot.alive {
			id.index = uint32(i + 1)
			break
		}
	}

	if id.index == 0 {
		ext.slice = append(ext.slice, instanceOf[Type]{})
		id.index = uint32(len(ext.slice))
	}

	instance := &ext.slice[id.index-1]
	instance.alive = true

	obj := ext.loader(instance.context(), &instance.value).owner()
	obj.set_instance(ext.class, id.pack())

	return obj
}

func (ext *Extension[Type, Extends]) delete(id instanceID) {
	var zero Type
	instance := &ext.slice[id.index-1]
	*instance.alloc--
	if *instance.alloc != 0 {
		panic(reflect.TypeOf(zero).String() + " is leaking memory")
	}
	instance.alive = false
}

func (ext *Extension[Type, Extends]) call(instance uint32, index uint8, args *unsafe.Pointer, result unsafe.Pointer) {
	rvalue := reflect.ValueOf(&ext.slice[instance-1].value)
	method := rvalue.Method(int(index) - 1)
	rtype := method.Type()

	buffer := make([]reflect.Value, rtype.NumIn())

	if rtype.NumIn() > 1 {
		slice := unsafe.Slice(args, rtype.NumIn())
		for i := range slice {
			buffer[i] = reflect.New(rtype.In(i + 1)).Elem()
			into(buffer[i].Interface(), slice[i])
		}
	}

	method.Call(buffer)

	// TODO result
}

func (ext *Extension[Type, Extends]) methods() []reflect.Method {
	var methods []reflect.Method
	for i := 0; i < ext.ptype.NumMethod(); i++ {
		methods = append(methods, ext.ptype.Method(i))
	}
	return methods
}
