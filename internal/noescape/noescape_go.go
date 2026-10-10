//go:build !cgo && !wasm

package noescape

import (
	"sync/atomic"
	"unsafe"

	"graphics.gd/internal/gdextension"
)

func Call[T any](object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args any) T {
	panic("not implemented")
}

func CallThreadSafe[T any](object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args any) T {
	panic("not implemented")
}

func CallThreadSafeIf[T any](safe *atomic.Bool, object gdextension.Object, method gdextension.MethodForClass, shape gdextension.Shape, args any) T {
	panic("not implemented")
}

func (method MethodForClass) Call(self gdextension.Object, args ...gdextension.Variant) (gdextension.Variant, error) {
	panic("not implemented")
}

// ScriptCallResident always reports false without cgo: resident-callback mode
// requires the fastcb runtime patch, which is native-only.
func ScriptCallResident(object gdextension.Object, name gdextension.StringName, args []gdextension.Variant) (gdextension.Variant, gdextension.CallError, bool) {
	return gdextension.Variant{}, gdextension.CallError{}, false
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
