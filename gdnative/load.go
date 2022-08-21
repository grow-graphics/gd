package gdnative

// #include <gdnative_interface.h>
import "C"

import (
	"fmt"
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
