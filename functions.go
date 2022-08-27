package gd

/*
	#include <gdnative_interface.h>

	void method0(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result) {
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, NULL, (GDNativeTypePtr*)result);
	}
	void method1(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a) {
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&a), (GDNativeTypePtr*)result);
	}
	void method2(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b) {
		GDNativeTypePtr stack[2] = {a, b};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method3(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c) {
		GDNativeTypePtr stack[3] = {a, b, c};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method4(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c, void *d) {
		GDNativeTypePtr stack[4] = {a, b, c, d};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method5(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c, void *d, void *e) {
		GDNativeTypePtr stack[5] = {a, b, c, d, e};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method6(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c, void *d, void *e, void *f) {
		GDNativeTypePtr stack[6] = {a, b, c, d, e, f};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method7(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c, void *d, void *e, void *f, void *g) {
		GDNativeTypePtr stack[7] = {a, b, c, d, e, f, g};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method8(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c, void *d, void *e, void *f, void *g, void *h) {
		GDNativeTypePtr stack[8] = {a, b, c, d, e, f, g, h};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method9(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c, void *d, void *e, void *f, void *g, void *h, void *i) {
		GDNativeTypePtr stack[9] = {a, b, c, d, e, f, g, h, i};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method10(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c, void *d, void *e, void *f, void *g, void *h, void *i, void *j) {
		GDNativeTypePtr stack[10] = {a, b, c, d, e, f, g, h, i, j};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method11(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c, void *d, void *e, void *f, void *g, void *h, void *i, void *j, void *k) {
		GDNativeTypePtr stack[11] = {a, b, c, d, e, f, g, h, i, j, k};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method12(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c, void *d, void *e, void *f, void *g, void *h, void *i, void *j, void *k, void *l) {
		GDNativeTypePtr stack[12] = {a, b, c, d, e, f, g, h, i, j, k, l};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}
	void method13(GDNativeInterface *api, uintptr_t method, uintptr_t object, void *result, void *a, void *b, void *c, void *d, void *e, void *f, void *g, void *h, void *i, void *j, void *k, void *l, void *m) {
		GDNativeTypePtr stack[13] = {a, b, c, d, e, f, g, h, i, j, k, l, m};
		((GDNativeInterface*)api)->object_method_bind_ptrcall((GDNativeMethodBindPtr*)method, (GDNativeObjectPtr*)object, (GDNativeTypePtr*)(&stack), (GDNativeTypePtr*)result);
	}

	void utility0(uintptr_t utility, void *result) {
		((GDNativePtrUtilityFunction)utility)(result, NULL, 0);
	}
	void utility1(uintptr_t utility, void *r_ret, void *a) {
		((GDNativePtrUtilityFunction)(utility))(r_ret, &a, 1);
	}
	void utility2(uintptr_t utility, void *r_ret, void *a, void *b) {
		GDNativeTypePtr stack[2] = { a, b };
		((GDNativePtrUtilityFunction)(utility))(r_ret, &stack[0], 2);
	}
	void utility3(uintptr_t utility, void *r_ret, void *a, void *b, void *c) {
		GDNativeTypePtr stack[3] = { a, b, c };
		((GDNativePtrUtilityFunction)(utility))(r_ret, &stack[0], 3);
	}
	void utility4(uintptr_t utility, void *r_ret, void *a, void *b, void *c, void *d) {
		GDNativeTypePtr stack[4] = { a, b, c, d };
		((GDNativePtrUtilityFunction)(utility))(r_ret, &stack[0], 4);
	}
	void utility5(uintptr_t utility, void *r_ret, void *a, void *b, void *c, void *d, void *e) {
		GDNativeTypePtr stack[5] = { a, b, c, d, e };
		((GDNativePtrUtilityFunction)(utility))(r_ret, &stack[0], 5);
	}
*/
import "C"
import (
	"reflect"
	"unsafe"
)

func toUnsafe(val any) (ptr unsafe.Pointer, free func()) {
	switch v := val.(type) {
	case *struct{}:
		return nil, nil
	case *bool:
		return unsafe.Pointer(v), nil
	case *int64:
		return unsafe.Pointer(v), nil
	case *float64:
		return unsafe.Pointer(v), nil
	case *cVector2:
		return unsafe.Pointer(v), nil
	case *string:
		var native cString
		if *v != "" {
			native.with_utf8_chars(*v)
		}
		return unsafe.Pointer(&native), func() {
			cVariantTypeString.get_destructor().call(unsafe.Pointer(&native))
		}
	case *cName:
		return unsafe.Pointer(v), nil
	default:
		panic("loadResultFor: unsupported type " + reflect.TypeOf(val).String())
	}
}

func into(result any, ptr unsafe.Pointer) {
	switch v := result.(type) {
	case *struct{}:
		return
	case *bool:
		*v = *(*bool)(ptr)
	case *int64:
		*v = *(*int64)(ptr)
	case *float64:
		*v = *(*float64)(ptr)
	case *cVector2:
		*v = *(*cVector2)(ptr)
	case *string:
		*v = (*cString)(ptr).to_utf8_chars()
	default:
		panic("into: unsupported type " + reflect.TypeOf(result).String())
	}
}

func methodCall[Result any](object cObject, method cMethodBind, args ...any) Result {
	var result Result
	c_result, free := toUnsafe(&result)

	var c_args = make([]unsafe.Pointer, len(args))
	for i := range args {
		c_args[i], free = toUnsafe(args[i])
		if free != nil {
			defer free()
		}
	}

	method.unsafeReturn(c_result, object, c_args...)

	if free != nil {
		into(&result, c_result)
		free()
	}

	return result
}

func (method cMethodBind) unsafeReturn(result unsafe.Pointer, object cObject, args ...unsafe.Pointer) {
	switch len(args) {
	case 0:
		C.method0(api, C.uintptr_t(method), C.uintptr_t(object), result)
	case 1:
		C.method1(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0])
	case 2:
		C.method2(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1])
	case 3:
		C.method3(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2])
	case 4:
		C.method4(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2], args[3])
	case 5:
		C.method5(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2], args[3],
			args[4])
	case 6:
		C.method6(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2], args[3],
			args[4], args[5])
	case 7:
		C.method7(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2], args[3],
			args[4], args[5], args[6])
	case 8:
		C.method8(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2], args[3],
			args[4], args[5], args[6], args[7])
	case 9:
		C.method9(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2], args[3],
			args[4], args[5], args[6], args[7],
			args[8])
	case 10:
		C.method10(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2], args[3],
			args[4], args[5], args[6], args[7],
			args[8], args[9])
	case 11:
		C.method11(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2], args[3],
			args[4], args[5], args[6], args[7],
			args[8], args[9], args[10])
	case 12:
		C.method12(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2], args[3],
			args[4], args[5], args[6], args[7],
			args[8], args[9], args[10], args[11])
	case 13:
		C.method13(api, C.uintptr_t(method), C.uintptr_t(object), result,
			args[0], args[1], args[2], args[3],
			args[4], args[5], args[6], args[7],
			args[8], args[9], args[10], args[11],
			args[12])
	default:
		panic("too many arguments")
	}
}

func utilityCall[Result any](utility cUtilityFunction, args ...any) Result {
	var result Result
	c_result, free := toUnsafe(&result)

	var c_args = make([]unsafe.Pointer, len(args))
	for i := range args {
		c_args[i], free = toUnsafe(args[i])
		if free != nil {
			defer free()
		}
	}

	utility.unsafeReturn(c_result, c_args...)

	if free != nil {
		into(&result, c_result)
		free()
	}

	return result
}

// Return the result of calling the given utility function with the given arguments, no type-checking
// is performed. So the arguments must be passed in as the correct type.
func (utility cUtilityFunction) unsafeReturn(result unsafe.Pointer, args ...unsafe.Pointer) {
	switch len(args) {
	case 0:
		C.utility0(C.uintptr_t(utility), result)
	case 1:
		C.utility1(C.uintptr_t(utility), result, args[0])
	case 2:
		C.utility2(C.uintptr_t(utility), result, args[0],
			args[1])
	case 3:
		C.utility3(C.uintptr_t(utility), result, args[0],
			args[1], args[2])
	case 4:
		C.utility4(C.uintptr_t(utility), result, args[0],
			args[1], args[2], args[3])
	case 5:
		C.utility5(C.uintptr_t(utility), result, args[0],
			args[1], args[2], args[3], args[4])
	default:
		panic("too many arguments")
	}
}

// Return the result of calling the given builtin method on the given base value with the given arguments, no type-checking
// is performed. So the arguments must be passed in as the correct type.
//func (builtin Builtin) Return(result Value, base Value, args ...Value) {

//}
