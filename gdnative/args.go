package gdnative

import (
	"reflect"
	"unsafe"
)

/*
	#include <gdnative_interface.h>

	const char *ConstCharArg(_GoString_ s) {
		return _GoStringPtr(s);
	}

	static GDNativeInt string_to_utf8_chars(
		GDNativeInterface *api,
		const GDNativeStringPtr p_self,
		char *r_text,
		GDNativeInt p_max_write_length
	) {
		return api->string_to_utf8_chars(p_self, r_text, p_max_write_length);
	}

	static void variant_to_type_constructor(
		GDNativeInterface *api,
		GDNativeVariantType p_type,
		GDNativeTypePtr p_native,
		GDNativeVariantPtr p_variant
	) {
		api->get_variant_to_type_constructor(p_type)(p_native, p_variant);
	}
*/
import "C"

func loadArgFromNativeType(as reflect.Type, ptr C.GDNativeTypePtr) reflect.Value {
	switch as.Kind() {
	case reflect.Bool:
		return reflect.ValueOf(*(*bool)(ptr))
	case reflect.Int:
		if as.Size() == 4 {
			return reflect.ValueOf(*(*int32)(ptr))
		}
		return reflect.ValueOf(*(*int64)(ptr))
	case reflect.Int8:
		return reflect.ValueOf(*(*int8)(ptr))
	case reflect.Int16:
		return reflect.ValueOf(*(*int16)(ptr))
	case reflect.Int32:
		return reflect.ValueOf(*(*int32)(ptr))
	case reflect.Int64:
		return reflect.ValueOf(*(*int64)(ptr))
	case reflect.Uint:
		if as.Size() == 4 {
			return reflect.ValueOf(*(*uint32)(ptr))
		}
		return reflect.ValueOf(*(*uint64)(ptr))
	case reflect.Uint8:
		return reflect.ValueOf(*(*uint8)(ptr))
	case reflect.Uint16:
		return reflect.ValueOf(*(*uint16)(ptr))
	case reflect.Uint32:
		return reflect.ValueOf(*(*uint32)(ptr))
	case reflect.Uint64:
		return reflect.ValueOf(*(*uint64)(ptr))
	case reflect.Float32:
		return reflect.ValueOf(*(*float32)(ptr))
	case reflect.Float64:
		return reflect.ValueOf(*(*float64)(ptr))

	case reflect.Complex64:
		return reflect.ValueOf(*(*complex64)(ptr))

	case reflect.UnsafePointer:
		return reflect.ValueOf(unsafe.Pointer(ptr))

	case reflect.String:
		length := C.string_to_utf8_chars(api, C.GDNativeStringPtr(ptr), nil, 0)
		buffer := make([]byte, length)
		if length > 0 {
			C.string_to_utf8_chars(api, C.GDNativeStringPtr(ptr), (*C.char)(unsafe.Pointer(&buffer[0])), length)
		}
		return reflect.ValueOf(string(buffer)) // FIXME avoid copy?
	default:
		panic("loadArgFromNativeType: unsupported type " + as.String())
	}
}

func loadArgFromVariant(as reflect.Type, ptr C.GDNativeVariantPtr) reflect.Value {
	switch as.Kind() {
	case reflect.String:
		var native String
		C.variant_to_type_constructor(api, C.GDNATIVE_VARIANT_TYPE_STRING, C.GDNativeTypePtr(&native.raw), ptr)
		return loadArgFromNativeType(as, C.GDNativeTypePtr(&native.raw))
	default:
		panic("loadArgFromVariant: unsupported type " + as.String())
	}
}

func loadResultFor(dst any) (C.GDNativeTypePtr, bool) {
	switch v := dst.(type) {
	case *int64:
		return C.GDNativeTypePtr(v), true
	case *float64:
		return C.GDNativeTypePtr(v), true
	case *bool:
		return C.GDNativeTypePtr(v), true
	case *string:
		var native String
		return C.GDNativeTypePtr(&native.raw), false
	default:
		panic("loadResultFor: unsupported type " + reflect.TypeOf(dst).String())
	}
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
