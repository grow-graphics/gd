package gd

import "C"

import (
	"reflect"
	"unsafe"
)

//export create_instance_func
func create_instance_func(userdata uintptr) uintptr {
	return uintptr(extensions[extensionID(userdata)].create())
}

//export free_instance_func
func free_instance_func(userdata, instance uintptr) {
	extensions[extensionID(userdata)].delete(loadInstanceID(instance))
}

//export get_virtual_index
func get_virtual_index(userdata uintptr, name *C.char) uint8 {
	method, ok := extensions[extensionID(userdata)].lookup(C.GoString(name))
	if !ok {
		return 0
	}
	if method.Index > 255 {
		panic("too many exported methods")
	}
	return uint8(method.Index)
}

//export call_virtual_index
func call_virtual_index(index uint8, instance uintptr, args *unsafe.Pointer, result unsafe.Pointer) {
	id := loadInstanceID(instance)
	extensions[id.class].call(id.index, index, args, result)
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

type ExtensionLoader[Type any, Extends Class] func(*Type) Extends

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

	// defer on load..

	classDB.register_extension_class(extension.class, extension.owner, uintptr(extension.id))

	// register methods.

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
	value Type
}

func (ext *Extension[Type, Extends]) lookup(name string) (reflect.Method, bool) {
	var zero Type
	var core Extends
	return core.virtual(reflect.TypeOf(zero), name)
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
		id.index = uint32(len(ext.slice))
	}

	instance := &ext.slice[id.index-1]
	instance.alive = true

	obj := ext.loader(&instance.value).owner()
	obj.set_instance(ext.class, id.pack())

	return obj
}

func (ext *Extension[Type, Extends]) delete(id instanceID) {
	ext.slice[id.index-1].alive = false
}

func (ext *Extension[Type, Extends]) call(instance uint32, index uint8, args *unsafe.Pointer, result unsafe.Pointer) {
	rvalue := reflect.ValueOf(&ext.slice[instance-1])
	method := rvalue.Method(int(index))
	rtype := method.Type()

	//fmt.Println(method.Name, " METHOD CALLED DIRECTLY")

	buffer := make([]reflect.Value, rtype.NumIn())

	buffer[0] = rvalue
	if rtype.NumIn() > 1 {
		slice := unsafe.Slice(args, rtype.NumIn()-1)
		value := buffer[1:]
		for i := range value {
			value[i] = reflect.New(rtype.In(i + 1)).Elem()
			into(value[i].Interface(), slice[i])
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
