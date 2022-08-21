// Since Go can't call C function pointers directly, we need to create wrapper functions for them.
void variant_call(
    GDNativeInterface *api,
    GDNativeVariantPtr p_self,
    const GDNativeStringNamePtr p_method,
    const GDNativeVariantPtr *p_args,
    const GDNativeInt p_argument_count,
    GDNativeVariantPtr r_return,
    GDNativeCallError *r_error
) {
    api->variant_call(p_self, p_method, p_args, p_argument_count, r_return, r_error);
}

void string_new_with_utf8_chars_and_len(
    GDNativeInterface *api,
    GDNativeStringPtr r_dest,
    const char *p_contents,
    const GDNativeInt p_size
) {
    api->string_new_with_utf8_chars_and_len(r_dest, p_contents, p_size);
}

void *mem_alloc(GDNativeInterface *api, size_t p_bytes) {
    api->mem_alloc(p_bytes);
}

void *mem_free(GDNativeInterface *api, void *ptr) {
    api->mem_free(ptr);
}

GDNativePtrConstructor variant_get_ptr_constructor(
    GDNativeInterface *api,
    GDNativeVariantType p_type, 
    int32_t p_constructor
) {
    api->variant_get_ptr_constructor(p_type, p_constructor);
}

void constructor(GDNativePtrConstructor ptr, GDNativeTypePtr base, GDNativeTypePtr *args) {
    ptr(base, args);
}