#include <gdnative_virtuals.h>

#include <assert.h>
static_assert(sizeof(void*) == 8, "Pointers are assumed to be exactly 4 bytes");
static_assert(sizeof(uintptr_t) == 8, "Pointers are assumed to be exactly 4 bytes");

extern void goInitialise(void *userdata, GDNativeInitializationLevel level);
extern void goDeinitialize(void *userdata, GDNativeInitializationLevel level);

static void setInitialise(GDNativeInitialization *init) {
    init->initialize = goInitialise;
}
static void setDeinitialize(GDNativeInitialization *init) {
    init->deinitialize = goDeinitialize;
}

GDNativeObjectPtr global_get_singleton(GDNativeInterface *api, const char *name) {
    return api->global_get_singleton(name);
}


// Since Go can't call C function pointers directly, we need to create wrapper functions for them.
static void variant_call(
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

static GDNativeObjectPtr classdb_construct_object(
    GDNativeInterface *api,
    const char *p_classname
) {
    return api->classdb_construct_object(p_classname);
}

static void object_set_instance(
    GDNativeInterface *api,
    uintptr_t p_o, 
    const char *p_classname, 
    uintptr_t p_instance
) {
    api->object_set_instance((GDNativeObjectPtr)(p_o), p_classname, (GDExtensionClassInstancePtr)(p_instance));
}



static void *mem_alloc(GDNativeInterface *api, size_t p_bytes) {
    return api->mem_alloc(p_bytes);
}

static void *mem_free(GDNativeInterface *api, uintptr_t ptr) {
    api->mem_free((void*)ptr);
}

static GDNativePtrConstructor variant_get_ptr_constructor(
    GDNativeInterface *api,
    GDNativeVariantType p_type, 
    int32_t p_constructor
) {
    api->variant_get_ptr_constructor(p_type, p_constructor);
}

static void constructor(GDNativePtrConstructor ptr, GDNativeTypePtr base, GDNativeTypePtr *args) {
    ptr(base, args);
}



static void object_method_bind_ptrcall(
    GDNativeInterface *api,
    uintptr_t p_method_bind, 
    uintptr_t p_instance,
    const GDNativeTypePtr *p_args, 
    GDNativeTypePtr r_ret
) {
    api->object_method_bind_ptrcall((GDNativeMethodBindPtr)p_method_bind, (GDNativeObjectPtr)p_instance, p_args, r_ret);
}

//
// Class Registration.
//

extern uintptr_t goClassCreateInstance(uintptr_t userdata);
extern void goClassFreeInstance(uintptr_t userdata, uintptr_t instance);

static GDNativeObjectPtr create_instance_func(void *p_userdata) {
    return (GDNativeObjectPtr)(goClassCreateInstance((uintptr_t)p_userdata));
}

static void free_instance_func(void *p_userdata, GDExtensionClassInstancePtr p_instance) {
    goClassFreeInstance((uintptr_t)p_userdata, (uintptr_t)p_instance);
}

static void object_destroy(GDNativeInterface *api, uintptr_t p_o) {
    api->object_destroy((GDNativeObjectPtr)p_o);
}

static void classdb_register_extension_class(
    GDNativeInterface *api,
    const GDNativeExtensionClassLibraryPtr p_library, 
    const char *p_class_name, 
    const char *p_parent_class_name, 
    uintptr_t userdata // cgo handle
) {

    GDNativeExtensionClassCreationInfo functions = {
        0, // GDNativeExtensionClassSet set_func;
		0, // GDNativeExtensionClassGet get_func;
		0, // GDNativeExtensionClassGetPropertyList get_property_list_func;
		0, // GDNativeExtensionClassFreePropertyList free_property_list_func;
		0, // GDNativeExtensionClassPropertyCanRevert property_can_revert_func;
		0, // GDNativeExtensionClassPropertyGetRevert property_get_revert_func;
		0, // GDNativeExtensionClassNotification notification_func;
		0, // GDNativeExtensionClassToString to_string_func;
		0, // GDNativeExtensionClassReference reference_func;
		0, // GDNativeExtensionClassUnreference unreference_func;
		create_instance_func, // GDNativeExtensionClassCreateInstance create_instance_func; /* this one is mandatory */
		free_instance_func, // GDNativeExtensionClassFreeInstance free_instance_func; /* this one is mandatory */
		get_virtual_func, // GDNativeExtensionClassGetVirtual get_virtual_func;
		0, // GDNativeExtensionClassGetRID get_rid;
		(void *)userdata, // void *class_userdata;
    };

    api->classdb_register_extension_class(p_library, p_class_name, p_parent_class_name, &functions);
}

//
// Method Registration.
//

extern void goMethodCall(uintptr_t userdata, uintptr_t instance, const GDNativeVariantPtr *p_args, const GDNativeInt p_argument_count, GDNativeVariantPtr r_return, GDNativeCallError *r_error);
extern GDNativeVariantType goMethodGetArgumentType(uintptr_t userdata, int32_t argument);
extern GDNativeExtensionClassMethodArgumentMetadata goMethodGetArgumentMetadata(uintptr_t userdata, int32_t argument);
extern void goMethodGetArgumentInfo(uintptr_t userdata, int32_t argument, GDNativePropertyInfo *r_info);

static void call_func(
    void *method_userdata, 
    GDExtensionClassInstancePtr p_instance, 
    const GDNativeVariantPtr *p_args, 
    const GDNativeInt p_argument_count, 
    GDNativeVariantPtr r_return, 
    GDNativeCallError *r_error
) {
    goMethodCall((uintptr_t)method_userdata, (uintptr_t)p_instance, p_args, p_argument_count, r_return, r_error);
}

extern void goMethodCallDirect(uintptr_t userdata, uintptr_t instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret);

static void ptrcall_func(
    void *method_userdata, 
    GDExtensionClassInstancePtr p_instance, 
    const GDNativeTypePtr *p_args, 
    GDNativeTypePtr r_ret
) {
    goMethodCallDirect((uintptr_t)method_userdata, (uintptr_t)p_instance, p_args, r_ret);
}

static GDNativeVariantType get_argument_type_func(
    void *p_method_userdata, int32_t p_argument
) {
    return goMethodGetArgumentType((uintptr_t)p_method_userdata, p_argument);
}

static void get_argument_info_func(
    void *p_method_userdata, 
    int32_t p_argument, 
    GDNativePropertyInfo *r_info
) {
    *r_info = (GDNativePropertyInfo){0};
    return goMethodGetArgumentInfo((uintptr_t)p_method_userdata, p_argument, r_info);
}

static GDNativeExtensionClassMethodArgumentMetadata get_argument_metadata_func(
    void *p_method_userdata, int32_t p_argument
) {
    return goMethodGetArgumentMetadata((uintptr_t)p_method_userdata, p_argument);
}

static void classdb_register_extension_class_method(
    GDNativeInterface *api,
    const GDNativeExtensionClassLibraryPtr p_library, 
    const char *p_class_name,
    GDNativeExtensionClassMethodInfo *p_method_info,
    uintptr_t userdata // cgo handle
) {
    p_method_info->method_userdata = (void*)userdata;
    p_method_info->call_func = call_func;
    p_method_info->ptrcall_func = ptrcall_func;
    p_method_info->get_argument_type_func = get_argument_type_func;
    p_method_info->get_argument_info_func = get_argument_info_func;
    p_method_info->get_argument_metadata_func = get_argument_metadata_func;
    api->classdb_register_extension_class_method(p_library, p_class_name, p_method_info);
}
