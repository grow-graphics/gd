package gdc

/*
	#include <gdnative_interface.h>

	void call0(GDNativeInterface *api, uintptr_t p_method_bind, uintptr_t p_instance, GDNativeTypePtr r_ret) {
		api->object_method_bind_ptrcall((GDNativeMethodBindPtr)p_method_bind, (GDNativeObjectPtr)p_instance, NULL, r_ret);
	}
	void call1(GDNativeInterface *api, uintptr_t p_method_bind, uintptr_t p_instance, GDNativeTypePtr r_ret,
		GDNativeTypePtr a
	) {
		api->object_method_bind_ptrcall((GDNativeMethodBindPtr)p_method_bind, (GDNativeObjectPtr)p_instance, &a, r_ret);
	}
	void call2(GDNativeInterface *api, uintptr_t p_method_bind, uintptr_t p_instance, GDNativeTypePtr r_ret,
		GDNativeTypePtr a, GDNativeTypePtr b
	) {
		GDNativeTypePtr stack[2] = { a, b };
		api->object_method_bind_ptrcall((GDNativeMethodBindPtr)p_method_bind, (GDNativeObjectPtr)p_instance, &stack[0], r_ret);
	}
	void call3(GDNativeInterface *api, uintptr_t p_method_bind, uintptr_t p_instance, GDNativeTypePtr r_ret,
		GDNativeTypePtr a, GDNativeTypePtr b, GDNativeTypePtr c
	) {
		GDNativeTypePtr stack[3] = { a, b, c };
		api->object_method_bind_ptrcall((GDNativeMethodBindPtr)p_method_bind, (GDNativeObjectPtr)p_instance, &stack[0], r_ret);
	}
	void call4(GDNativeInterface *api, uintptr_t p_method_bind, uintptr_t p_instance, GDNativeTypePtr r_ret,
		GDNativeTypePtr a, GDNativeTypePtr b, GDNativeTypePtr c, GDNativeTypePtr d
	) {
		GDNativeTypePtr stack[4] = { a, b, c, d };
		api->object_method_bind_ptrcall((GDNativeMethodBindPtr)p_method_bind, (GDNativeObjectPtr)p_instance, &stack[0], r_ret);
	}
	void call5(GDNativeInterface *api, uintptr_t p_method_bind, uintptr_t p_instance, GDNativeTypePtr r_ret,
		GDNativeTypePtr a, GDNativeTypePtr b, GDNativeTypePtr c, GDNativeTypePtr d, GDNativeTypePtr e
	) {
		GDNativeTypePtr stack[5] = { a, b, c, d, e };
		api->object_method_bind_ptrcall((GDNativeMethodBindPtr)p_method_bind, (GDNativeObjectPtr)p_instance, &stack[0], r_ret);
	}
	void call6(GDNativeInterface *api, uintptr_t p_method_bind, uintptr_t p_instance, GDNativeTypePtr r_ret,
		GDNativeTypePtr a, GDNativeTypePtr b, GDNativeTypePtr c, GDNativeTypePtr d, GDNativeTypePtr e, GDNativeTypePtr f
	) {
		GDNativeTypePtr stack[6] = { a, b, c, d, e, f };
		api->object_method_bind_ptrcall((GDNativeMethodBindPtr)p_method_bind, (GDNativeObjectPtr)p_instance, &stack[0], r_ret);
	}
	void call7(GDNativeInterface *api, uintptr_t p_method_bind, uintptr_t p_instance, GDNativeTypePtr r_ret,
		GDNativeTypePtr a, GDNativeTypePtr b, GDNativeTypePtr c, GDNativeTypePtr d, GDNativeTypePtr e, GDNativeTypePtr f, GDNativeTypePtr g
	) {
		GDNativeTypePtr stack[7] = { a, b, c, d, e, f, g };
		api->object_method_bind_ptrcall((GDNativeMethodBindPtr)p_method_bind, (GDNativeObjectPtr)p_instance, &stack[0], r_ret);
	}
	void call8(GDNativeInterface *api, uintptr_t p_method_bind, uintptr_t p_instance, GDNativeTypePtr r_ret,
		GDNativeTypePtr a, GDNativeTypePtr b, GDNativeTypePtr c, GDNativeTypePtr d, GDNativeTypePtr e, GDNativeTypePtr f, GDNativeTypePtr g, GDNativeTypePtr h
	) {
		GDNativeTypePtr stack[8] = { a, b, c, d, e, f, g, h };
		api->object_method_bind_ptrcall((GDNativeMethodBindPtr)p_method_bind, (GDNativeObjectPtr)p_instance, &stack[0], r_ret);
	}
	void call9(GDNativeInterface *api, uintptr_t p_method_bind, uintptr_t p_instance, GDNativeTypePtr r_ret,
		GDNativeTypePtr a, GDNativeTypePtr b, GDNativeTypePtr c, GDNativeTypePtr d, GDNativeTypePtr e, GDNativeTypePtr f,
		GDNativeTypePtr g, GDNativeTypePtr h, GDNativeTypePtr i
	) {
		GDNativeTypePtr stack[9] = { a, b, c, d, e, f, g, h, i };
		api->object_method_bind_ptrcall((GDNativeMethodBindPtr)p_method_bind, (GDNativeObjectPtr)p_instance, &stack[0], r_ret);
	}
	void call10(GDNativeInterface *api, uintptr_t p_method_bind, uintptr_t p_instance, GDNativeTypePtr r_ret,
		GDNativeTypePtr a, GDNativeTypePtr b, GDNativeTypePtr c, GDNativeTypePtr d, GDNativeTypePtr e, GDNativeTypePtr f,
		GDNativeTypePtr g, GDNativeTypePtr h, GDNativeTypePtr i, GDNativeTypePtr j
	) {
		GDNativeTypePtr stack[10] = { a, b, c, d, e, f, g, h, i, j };
		api->object_method_bind_ptrcall((GDNativeMethodBindPtr)p_method_bind, (GDNativeObjectPtr)p_instance, &stack[0], r_ret);
	}
	void call11(GDNativeInterface *api, uintptr_t p_method_bind, uintptr_t p_instance, GDNativeTypePtr r_ret,
		GDNativeTypePtr a, GDNativeTypePtr b, GDNativeTypePtr c, GDNativeTypePtr d, GDNativeTypePtr e, GDNativeTypePtr f,
		GDNativeTypePtr g, GDNativeTypePtr h, GDNativeTypePtr i, GDNativeTypePtr j, GDNativeTypePtr k
	) {
		GDNativeTypePtr stack[11] = { a, b, c, d, e, f, g, h, i, j, k };
		api->object_method_bind_ptrcall((GDNativeMethodBindPtr)p_method_bind, (GDNativeObjectPtr)p_instance, &stack[0], r_ret);
	}
	void call12(GDNativeInterface *api, uintptr_t p_method_bind, uintptr_t p_instance, GDNativeTypePtr r_ret,
		GDNativeTypePtr a, GDNativeTypePtr b, GDNativeTypePtr c, GDNativeTypePtr d, GDNativeTypePtr e, GDNativeTypePtr f,
		GDNativeTypePtr g, GDNativeTypePtr h, GDNativeTypePtr i, GDNativeTypePtr j, GDNativeTypePtr k, GDNativeTypePtr l
	) {
		GDNativeTypePtr stack[12] = { a, b, c, d, e, f, g, h, i, j, k, l };
		api->object_method_bind_ptrcall((GDNativeMethodBindPtr)p_method_bind, (GDNativeObjectPtr)p_instance, &stack[0], r_ret);
	}
	void call13(GDNativeInterface *api, uintptr_t p_method_bind, uintptr_t p_instance, GDNativeTypePtr r_ret,
		GDNativeTypePtr a, GDNativeTypePtr b, GDNativeTypePtr c, GDNativeTypePtr d, GDNativeTypePtr e, GDNativeTypePtr f,
		GDNativeTypePtr g, GDNativeTypePtr h, GDNativeTypePtr i, GDNativeTypePtr j, GDNativeTypePtr k, GDNativeTypePtr l, GDNativeTypePtr m
	) {
		GDNativeTypePtr stack[13] = { a, b, c, d, e, f, g, h, i, j, k, l, m };
		api->object_method_bind_ptrcall((GDNativeMethodBindPtr)p_method_bind, (GDNativeObjectPtr)p_instance, &stack[0], r_ret);
	}

	void util0(uintptr_t utility, GDNativeTypePtr r_ret) {
		((GDNativePtrUtilityFunction)(utility))(r_ret, NULL, 0);
	}
	void util1(uintptr_t utility, GDNativeTypePtr r_ret, GDNativeTypePtr a) {
		((GDNativePtrUtilityFunction)(utility))(r_ret, &a, 1);
	}
	void util2(uintptr_t utility, GDNativeTypePtr r_ret, GDNativeTypePtr a, GDNativeTypePtr b) {
		GDNativeTypePtr stack[2] = { a, b };
		((GDNativePtrUtilityFunction)(utility))(r_ret, &stack[0], 2);
	}
	void util3(uintptr_t utility, GDNativeTypePtr r_ret, GDNativeTypePtr a, GDNativeTypePtr b, GDNativeTypePtr c) {
		GDNativeTypePtr stack[3] = { a, b, c };
		((GDNativePtrUtilityFunction)(utility))(r_ret, &stack[0], 3);
	}
	void util4(uintptr_t utility, GDNativeTypePtr r_ret, GDNativeTypePtr a, GDNativeTypePtr b, GDNativeTypePtr c, GDNativeTypePtr d) {
		GDNativeTypePtr stack[4] = { a, b, c, d };
		((GDNativePtrUtilityFunction)(utility))(r_ret, &stack[0], 4);
	}
	void util5(uintptr_t utility, GDNativeTypePtr r_ret, GDNativeTypePtr a, GDNativeTypePtr b, GDNativeTypePtr c, GDNativeTypePtr d, GDNativeTypePtr e) {
		GDNativeTypePtr stack[5] = { a, b, c, d, e };
		((GDNativePtrUtilityFunction)(utility))(r_ret, &stack[0], 5);
	}

	const char *ConstArgChar(_GoString_ s) {
		return _GoStringPtr(s);
	}

	static uintptr_t classdb_get_method_bind(
		GDNativeInterface *api,
		const char *p_classname,
		const char *p_methodname,
		GDNativeInt p_hash
	) {
		return (uintptr_t)(api->classdb_get_method_bind(p_classname, p_methodname, p_hash));
	}

	static uintptr_t variant_get_ptr_utility_function(
		GDNativeInterface *api,
		const char *p_function,
		GDNativeInt p_hash
	) {
		return (uintptr_t)(api->variant_get_ptr_utility_function(p_function, p_hash));
	}

*/
import "C"
import (
	"fmt"
	"reflect"
)

// MethodOf arguments must be null-terminated.
func MethodOf(class, method string, hash int64) Method {
	result := Method(C.classdb_get_method_bind(api, C.ConstArgChar(class), C.ConstArgChar(method), C.GDNativeInt(hash)))
	if result == 0 {
		fmt.Println("Method not found:", class, method)
	}
	return result
}

// LookupUtility looks up a utility function.
func LookupUtility(name string, hash int64) Utility {
	result := Utility(C.variant_get_ptr_utility_function(api, C.ConstArgChar(name), C.GDNativeInt(hash)))
	if result == 0 {
		fmt.Println("Utility not found:", name)
	}
	return result
}

func sendArg(arg any) C.GDNativeTypePtr {
	switch v := arg.(type) {
	case bool:
		return C.GDNativeTypePtr(&v)
	case int:
		return C.GDNativeTypePtr(&v)
	case int8:
		return C.GDNativeTypePtr(&v)
	case int16:
		return C.GDNativeTypePtr(&v)
	case int32:
		return C.GDNativeTypePtr(&v)
	case int64:
		return C.GDNativeTypePtr(&v)
	case uint:
		return C.GDNativeTypePtr(&v)
	case uint8:
		return C.GDNativeTypePtr(&v)
	case uint16:
		return C.GDNativeTypePtr(&v)
	case uint32:
		return C.GDNativeTypePtr(&v)
	case uint64:
		return C.GDNativeTypePtr(&v)
	case float32:
		return C.GDNativeTypePtr(&v)
	case float64:
		return C.GDNativeTypePtr(&v)
	case complex64:
		return C.GDNativeTypePtr(&v)
	case string:
		var native = NewString(v)
		return C.GDNativeTypePtr(&native.raw)
	default:
		panic("sendArg: unsupported type " + reflect.TypeOf(arg).String())
	}
}

func Return0[Result any](object Object, method Method) Result {
	var result Result
	ret, ok := loadResultFor(&result)
	C.call0(api, C.uintptr_t(method), C.uintptr_t(object), ret)
	if !ok {
		result = loadArgFromNativeType(reflect.TypeOf(result), ret).Interface().(Result)
	}
	return result
}

func Return1[Result any](object Object, method Method, a any) Result {
	var result Result
	ret, ok := loadResultFor(&result)
	C.call1(api, C.uintptr_t(method), C.uintptr_t(object), ret,
		sendArg(a),
	)
	if !ok {
		result = loadArgFromNativeType(reflect.TypeOf(result), ret).Interface().(Result)
	}
	return result
}

func Return2[Result any](object Object, method Method, a, b any) Result {
	var result Result
	ret, ok := loadResultFor(&result)
	C.call2(api, C.uintptr_t(method), C.uintptr_t(object), ret,
		sendArg(a), sendArg(b),
	)
	if !ok {
		result = loadArgFromNativeType(reflect.TypeOf(result), ret).Interface().(Result)
	}
	return result
}

func Return3[Result any](object Object, method Method, a, b, c any) Result {
	var result Result
	ret, ok := loadResultFor(&result)
	C.call3(api, C.uintptr_t(method), C.uintptr_t(object), ret,
		sendArg(a), sendArg(b), sendArg(c),
	)
	if !ok {
		result = loadArgFromNativeType(reflect.TypeOf(result), ret).Interface().(Result)
	}
	return result
}

func Return4[Result any](object Object, method Method, a, b, c, d any) Result {
	var result Result
	ret, ok := loadResultFor(&result)
	C.call4(api, C.uintptr_t(method), C.uintptr_t(object), ret,
		sendArg(a), sendArg(b), sendArg(c), sendArg(d),
	)
	if !ok {
		result = loadArgFromNativeType(reflect.TypeOf(result), ret).Interface().(Result)
	}
	return result
}

func Return5[Result any](object Object, method Method, a, b, c, d, e any) Result {
	var result Result
	ret, ok := loadResultFor(&result)
	C.call5(api, C.uintptr_t(method), C.uintptr_t(object), ret,
		sendArg(a), sendArg(b), sendArg(c), sendArg(d), sendArg(e),
	)
	if !ok {
		result = loadArgFromNativeType(reflect.TypeOf(result), ret).Interface().(Result)
	}
	return result
}

func Return6[Result any](object Object, method Method, a, b, c, d, e, f any) Result {
	var result Result
	ret, ok := loadResultFor(&result)
	C.call6(api, C.uintptr_t(method), C.uintptr_t(object), ret,
		sendArg(a), sendArg(b), sendArg(c), sendArg(d), sendArg(e), sendArg(f),
	)
	if !ok {
		result = loadArgFromNativeType(reflect.TypeOf(result), ret).Interface().(Result)
	}
	return result
}

func Return7[Result any](object Object, method Method, a, b, c, d, e, f, g any) Result {
	var result Result
	ret, ok := loadResultFor(&result)
	C.call7(api, C.uintptr_t(method), C.uintptr_t(object), ret,
		sendArg(a), sendArg(b), sendArg(c), sendArg(d), sendArg(e), sendArg(f), sendArg(g),
	)
	if !ok {
		result = loadArgFromNativeType(reflect.TypeOf(result), ret).Interface().(Result)
	}
	return result
}

func Return8[Result any](object Object, method Method, a, b, c, d, e, f, g, h any) Result {
	var result Result
	ret, ok := loadResultFor(&result)
	C.call8(api, C.uintptr_t(method), C.uintptr_t(object), ret,
		sendArg(a), sendArg(b), sendArg(c), sendArg(d), sendArg(e), sendArg(f), sendArg(g), sendArg(h),
	)
	if !ok {
		result = loadArgFromNativeType(reflect.TypeOf(result), ret).Interface().(Result)
	}
	return result
}

func Return9[Result any](object Object, method Method, a, b, c, d, e, f, g, h, i any) Result {
	var result Result
	ret, ok := loadResultFor(&result)
	C.call9(api, C.uintptr_t(method), C.uintptr_t(object), ret,
		sendArg(a), sendArg(b), sendArg(c), sendArg(d), sendArg(e), sendArg(f), sendArg(g), sendArg(h), sendArg(i),
	)
	if !ok {
		result = loadArgFromNativeType(reflect.TypeOf(result), ret).Interface().(Result)
	}
	return result
}

func Return10[Result any](object Object, method Method, a, b, c, d, e, f, g, h, i, j any) Result {
	var result Result
	ret, ok := loadResultFor(&result)
	C.call10(api, C.uintptr_t(method), C.uintptr_t(object), ret,
		sendArg(a), sendArg(b), sendArg(c), sendArg(d), sendArg(e), sendArg(f), sendArg(g), sendArg(h), sendArg(i), sendArg(j),
	)
	if !ok {
		result = loadArgFromNativeType(reflect.TypeOf(result), ret).Interface().(Result)
	}
	return result
}

func Return11[Result any](object Object, method Method, a, b, c, d, e, f, g, h, i, j, k any) Result {
	var result Result
	ret, ok := loadResultFor(&result)
	C.call11(api, C.uintptr_t(method), C.uintptr_t(object), ret,
		sendArg(a), sendArg(b), sendArg(c), sendArg(d), sendArg(e), sendArg(f), sendArg(g), sendArg(h), sendArg(i), sendArg(j), sendArg(k),
	)
	if !ok {
		result = loadArgFromNativeType(reflect.TypeOf(result), ret).Interface().(Result)
	}
	return result
}

func Return12[Result any](object Object, method Method, a, b, c, d, e, f, g, h, i, j, k, l any) Result {
	var result Result
	ret, ok := loadResultFor(&result)
	C.call12(api, C.uintptr_t(method), C.uintptr_t(object), ret,
		sendArg(a), sendArg(b), sendArg(c), sendArg(d), sendArg(e), sendArg(f), sendArg(g), sendArg(h), sendArg(i), sendArg(j), sendArg(k), sendArg(l),
	)
	if !ok {
		result = loadArgFromNativeType(reflect.TypeOf(result), ret).Interface().(Result)
	}
	return result
}

func Return13[Result any](object Object, method Method, a, b, c, d, e, f, g, h, i, j, k, l, m any) Result {
	var result Result
	ret, ok := loadResultFor(&result)
	C.call13(api, C.uintptr_t(method), C.uintptr_t(object), ret,
		sendArg(a), sendArg(b), sendArg(c), sendArg(d), sendArg(e), sendArg(f), sendArg(g), sendArg(h), sendArg(i), sendArg(j), sendArg(k), sendArg(l), sendArg(m),
	)
	if !ok {
		result = loadArgFromNativeType(reflect.TypeOf(result), ret).Interface().(Result)
	}
	return result
}

type Utility uintptr

func ReturnUtility0[Result any](utility Utility) Result {
	var result Result
	ret, ok := loadResultFor(&result)
	C.util0(C.uintptr_t(utility), ret)
	if !ok {
		result = loadArgFromNativeType(reflect.TypeOf(result), ret).Interface().(Result)
	}
	return result
}

func ReturnUtility1[Result any](utility Utility, a any) Result {
	var result Result
	ret, ok := loadResultFor(&result)
	C.util1(C.uintptr_t(utility), ret, sendArg(a))
	if !ok {
		result = loadArgFromNativeType(reflect.TypeOf(result), ret).Interface().(Result)
	}
	return result
}

func ReturnUtility2[Result any](utility Utility, a, b any) Result {
	var result Result
	ret, ok := loadResultFor(&result)
	C.util2(C.uintptr_t(utility), ret, sendArg(a), sendArg(b))
	if !ok {
		result = loadArgFromNativeType(reflect.TypeOf(result), ret).Interface().(Result)
	}
	return result
}

func ReturnUtility3[Result any](utility Utility, a, b, c any) Result {
	var result Result
	ret, ok := loadResultFor(&result)
	C.util3(C.uintptr_t(utility), ret, sendArg(a), sendArg(b), sendArg(c))
	if !ok {
		result = loadArgFromNativeType(reflect.TypeOf(result), ret).Interface().(Result)
	}
	return result
}

func ReturnUtility4[Result any](utility Utility, a, b, c, d any) Result {
	var result Result
	ret, ok := loadResultFor(&result)
	C.util4(C.uintptr_t(utility), ret, sendArg(a), sendArg(b), sendArg(c), sendArg(d))
	if !ok {
		result = loadArgFromNativeType(reflect.TypeOf(result), ret).Interface().(Result)
	}
	return result
}

func ReturnUtility5[Result any](utility Utility, a, b, c, d, e any) Result {
	var result Result
	ret, ok := loadResultFor(&result)
	C.util5(C.uintptr_t(utility), ret, sendArg(a), sendArg(b), sendArg(c), sendArg(d), sendArg(e))
	if !ok {
		result = loadArgFromNativeType(reflect.TypeOf(result), ret).Interface().(Result)
	}
	return result
}
