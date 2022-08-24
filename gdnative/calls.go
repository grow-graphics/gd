package gdnative

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
*/
import "C"
import "reflect"

func Call0(object Object, method Method) {
	C.call0(api, C.uintptr_t(method), C.uintptr_t(object), nil)
}

func Call1(object Object, method Method, a any) {
	C.call1(api, C.uintptr_t(method), C.uintptr_t(object), nil,
		sendArg(a))
}

func Call2(object Object, method Method, a, b any) {
	C.call2(api, C.uintptr_t(method), C.uintptr_t(object), nil,
		sendArg(a), sendArg(b))
}

func Call3(object Object, method Method, a, b, c any) {
	C.call3(api, C.uintptr_t(method), C.uintptr_t(object), nil,
		sendArg(a), sendArg(b), sendArg(c))
}

func Call4(object Object, method Method, a, b, c, d any) {
	C.call4(api, C.uintptr_t(method), C.uintptr_t(object), nil,
		sendArg(a), sendArg(b), sendArg(c), sendArg(d))
}

func Call5(object Object, method Method, a, b, c, d, e any) {
	C.call5(api, C.uintptr_t(method), C.uintptr_t(object), nil,
		sendArg(a), sendArg(b), sendArg(c), sendArg(d), sendArg(e))
}

func Call6(object Object, method Method, a, b, c, d, e, f any) {
	C.call6(api, C.uintptr_t(method), C.uintptr_t(object), nil,
		sendArg(a), sendArg(b), sendArg(c), sendArg(d), sendArg(e), sendArg(f))
}

func Call7(object Object, method Method, a, b, c, d, e, f, g any) {
	C.call7(api, C.uintptr_t(method), C.uintptr_t(object), nil,
		sendArg(a), sendArg(b), sendArg(c), sendArg(d), sendArg(e), sendArg(f), sendArg(g))
}

func Call8(object Object, method Method, a, b, c, d, e, f, g, h any) {
	C.call8(api, C.uintptr_t(method), C.uintptr_t(object), nil,
		sendArg(a), sendArg(b), sendArg(c), sendArg(d), sendArg(e), sendArg(f), sendArg(g), sendArg(h))
}

func Call9(object Object, method Method, a, b, c, d, e, f, g, h, i any) {
	C.call9(api, C.uintptr_t(method), C.uintptr_t(object), nil,
		sendArg(a), sendArg(b), sendArg(c), sendArg(d), sendArg(e), sendArg(f), sendArg(g), sendArg(h), sendArg(i))
}

func Call10(object Object, method Method, a, b, c, d, e, f, g, h, i, j any) {
	C.call10(api, C.uintptr_t(method), C.uintptr_t(object), nil,
		sendArg(a), sendArg(b), sendArg(c), sendArg(d), sendArg(e), sendArg(f), sendArg(g), sendArg(h), sendArg(i), sendArg(j))
}

func Call11(object Object, method Method, a, b, c, d, e, f, g, h, i, j, k any) {
	C.call11(api, C.uintptr_t(method), C.uintptr_t(object), nil,
		sendArg(a), sendArg(b), sendArg(c), sendArg(d), sendArg(e), sendArg(f), sendArg(g), sendArg(h), sendArg(i), sendArg(j), sendArg(k))
}

func Call12(object Object, method Method, a, b, c, d, e, f, g, h, i, j, k, l any) {
	C.call12(api, C.uintptr_t(method), C.uintptr_t(object), nil,
		sendArg(a), sendArg(b), sendArg(c), sendArg(d), sendArg(e), sendArg(f), sendArg(g), sendArg(h), sendArg(i), sendArg(j), sendArg(k), sendArg(l))
}

func Call13(object Object, method Method, a, b, c, d, e, f, g, h, i, j, k, l, m any) {
	C.call13(api, C.uintptr_t(method), C.uintptr_t(object), nil,
		sendArg(a), sendArg(b), sendArg(c), sendArg(d), sendArg(e), sendArg(f), sendArg(g), sendArg(h), sendArg(i), sendArg(j), sendArg(k), sendArg(l), sendArg(m))
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
