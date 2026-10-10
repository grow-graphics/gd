//go:build wasm

package noescape

import (
	"sync/atomic"
	"unsafe"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/ring"
)

// CallThreadSafe behaves like Call on wasm, which is single-threaded.
func CallThreadSafe[T any](object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args any) T {
	return Call[T](object, method, shape, args)
}

// CallThreadSafeIf behaves like Call on wasm, which is single-threaded.
func CallThreadSafeIf[T any](safe *atomic.Bool, object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args any) T {
	return Call[T](object, method, shape, args)
}

func Call[T any](object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args any) T {
	var argptr unsafe.Pointer = nil
	var result T
	if args != nil {
		argptr = argPointer(args)
	}
	if unsafe.Sizeof(result) == 0 {
		ring.Main.Buffer(uintptr(object), uintptr(method), uint64(shape), argptr, 0)
		return result
	}
	if ring.Main.Pending() {
		ring.Main.Flush()
	}
	call_noescape(object, method, unsafe.Pointer(&result), shape, argptr)
	return result
}

//go:noescape
func call_noescape(object gdextension.Object, method gdextension.MethodForClass, result unsafe.Pointer, shape gdextension.Shape, args unsafe.Pointer)

//go:linkname call graphics.gd/internal/noescape.call_noescape
func call(object gdextension.Object, method gdextension.MethodForClass, result unsafe.Pointer, shape gdextension.Shape, args unsafe.Pointer) {
	gdextension.Host.Objects.Unsafe.Call(object, method, gdextension.CallReturns[any](result), shape, gdextension.CallAccepts[any](args))
}

func (method MethodForClass) Call(self gdextension.Object, args ...gdextension.Variant) (gdextension.Variant, error) {
	var result gdextension.Variant
	var err gdextension.CallError
	object_method_call_noescape(self, gdextension.MethodForClass(method), &result, args, &err)
	return result, err.Err()
}

// ScriptCallResident always reports false on wasm: resident-callback mode
// requires the fastcb runtime patch, which is native-only.
func ScriptCallResident(object gdextension.Object, name gdextension.StringName, args []gdextension.Variant) (gdextension.Variant, gdextension.CallError, bool) {
	return gdextension.Variant{}, gdextension.CallError{}, false
}

//go:noescape
func object_method_call_noescape(object gdextension.Object, method gdextension.MethodForClass, result *gdextension.Variant, args []gdextension.Variant, err *gdextension.CallError)

//go:linkname object_method_call graphics.gd/internal/noescape.object_method_call_noescape
func object_method_call(object gdextension.Object, method gdextension.MethodForClass, result *gdextension.Variant, args []gdextension.Variant, err *gdextension.CallError) {
	gdextension.Host.Objects.Call(object, method, gdextension.CallReturns[gdextension.Variant](result), len(args), gdextension.CallAccepts[gdextension.Variant](unsafe.SliceData(args)), gdextension.CallReturns[gdextension.CallError](err))
}

// CallableCallv is Callable.callv. The callable, the argument and the result
// are handed to the Host function, which makes them escape to the heap: the
// function being called can be Go code that moves the calling goroutine's
// stack, and the engine writes the result after it has returned.
func CallableCallv(callv gdextension.MethodForBuiltinType, callable gdextension.Callable, args gdextension.Array) gdextension.Variant {
	frame := new(struct {
		self   gdextension.Callable
		result gdextension.Variant
		args   struct{ gdextension.Array }
	})
	frame.self = callable
	frame.args.Array = args
	gdextension.Host.Builtin.Types.Unsafe.Call(
		gdextension.CallMutates[any](unsafe.Pointer(&frame.self)),
		callv,
		gdextension.CallReturns[any](unsafe.Pointer(&frame.result)),
		gdextension.SizeVariant|gdextension.SizeCallable<<4|gdextension.SizeArray<<8,
		gdextension.CallAccepts[any](unsafe.Pointer(&frame.args)),
	)
	return frame.result
}
