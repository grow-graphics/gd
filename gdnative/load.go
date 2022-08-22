package gdnative

// #include <gdnative_interface.h>
import "C"

import (
	"fmt"
	"runtime/cgo"
	"unsafe"
)

//export loadExtension
func loadExtension(api, library, init unsafe.Pointer) uint8 {
	load(api, library, init)
	return 1
}

//export goInitialise
func goInitialise(userdata unsafe.Pointer, level InitializationLevel) {
	fmt.Println("initialising ", level)
}

//export goDeinitialize
func goDeinitialize(userdata unsafe.Pointer, level InitializationLevel) {
	fmt.Println("deinitialize ", level)
}

//export goClassCreateInstance
func goClassCreateInstance(userdata uintptr) uintptr {
	return uintptr(cgo.Handle(userdata).Value().(Class).Create())
}

//export goClassFreeInstance
func goClassFreeInstance(userdata, instance uintptr) {
	cgo.Handle(userdata).Value().(Class).Delete(InstanceID(instance))
}

// Main should be called in your extension's main function.
func Main() {}
