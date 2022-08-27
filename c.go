package gd

/*
	#include <gdnative_interface.h>

	extern uintptr_t create_instance_func(uintptr_t p_userdata);
	extern void free_instance_func(uintptr_t p_userdata, uintptr_t p_instance);

	extern uint8_t get_virtual_index(void *p_userdata, const char *p_name);
	extern void call_virtual_index(uint8_t index, uintptr_t instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret);

	void call_virtual_index1(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)1, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index2(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)2, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index3(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)3, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index4(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)4, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index5(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)5, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index6(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)6, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index7(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)7, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index8(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)8, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index9(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)9, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index10(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)10, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index11(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)11, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index12(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)12, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index13(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)13, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index14(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)14, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index15(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)15, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index16(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)16, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index17(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)17, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index18(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)18, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index19(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)19, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index20(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)20, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index21(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)21, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index22(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)22, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index23(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)23, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index24(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)24, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index25(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)25, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index26(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)26, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index27(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)27, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index28(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)28, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index29(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)29, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index30(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)30, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index31(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)31, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index32(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)32, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index33(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)33, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index34(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)34, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index35(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)35, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index36(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)36, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index37(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)37, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index38(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)38, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index39(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)39, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index40(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)40, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index41(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)41, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index42(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)42, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index43(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)43, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index44(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)44, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index45(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)45, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index46(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)46, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index47(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)47, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index48(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)48, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index49(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)49, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index50(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)50, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index51(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)51, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index52(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)52, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index53(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)53, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index54(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)54, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index55(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)55, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index56(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)56, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index57(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)57, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index58(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)58, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index59(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)59, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index60(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)60, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index61(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)61, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index62(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)62, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index63(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)63, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index64(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)64, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index65(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)65, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index66(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)66, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index67(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)67, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index68(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)68, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index69(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)69, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index70(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)70, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index71(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)71, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index72(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)72, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index73(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)73, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index74(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)74, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index75(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)75, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index76(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)76, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index77(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)77, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index78(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)78, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index79(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)79, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index80(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)80, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index81(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)81, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index82(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)82, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index83(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)83, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index84(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)84, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index85(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)85, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index86(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)86, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index87(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)87, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index88(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)88, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index89(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)89, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index90(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)90, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index91(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)91, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index92(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)92, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index93(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)93, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index94(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)94, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index95(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)95, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index96(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)96, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index97(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)97, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index98(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)98, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index99(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)99, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index100(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)100, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index101(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)101, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index102(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)102, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index103(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)103, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index104(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)104, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index105(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)105, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index106(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)106, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index107(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)107, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index108(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)108, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index109(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)109, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index110(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)110, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index111(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)111, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index112(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)112, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index113(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)113, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index114(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)114, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index115(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)115, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index116(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)116, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index117(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)117, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index118(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)118, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index119(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)119, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index120(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)120, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index121(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)121, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index122(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)122, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index123(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)123, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index124(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)124, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index125(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)125, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index126(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)126, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index127(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)127, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index128(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)128, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index129(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)129, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index130(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)130, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index131(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)131, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index132(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)132, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index133(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)133, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index134(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)134, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index135(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)135, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index136(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)136, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index137(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)137, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index138(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)138, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index139(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)139, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index140(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)140, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index141(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)141, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index142(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)142, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index143(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)143, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index144(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)144, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index145(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)145, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index146(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)146, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index147(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)147, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index148(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)148, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index149(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)149, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index150(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)150, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index151(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)151, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index152(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)152, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index153(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)153, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index154(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)154, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index155(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)155, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index156(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)156, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index157(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)157, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index158(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)158, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index159(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)159, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index160(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)160, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index161(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)161, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index162(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)162, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index163(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)163, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index164(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)164, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index165(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)165, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index166(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)166, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index167(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)167, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index168(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)168, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index169(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)169, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index170(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)170, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index171(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)171, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index172(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)172, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index173(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)173, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index174(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)174, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index175(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)175, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index176(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)176, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index177(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)177, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index178(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)178, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index179(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)179, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index180(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)180, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index181(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)181, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index182(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)182, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index183(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)183, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index184(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)184, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index185(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)185, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index186(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)186, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index187(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)187, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index188(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)188, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index189(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)189, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index190(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)190, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index191(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)191, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index192(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)192, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index193(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)193, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index194(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)194, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index195(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)195, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index196(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)196, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index197(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)197, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index198(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)198, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index199(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)199, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index200(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)200, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index201(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)201, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index202(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)202, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index203(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)203, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index204(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)204, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index205(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)205, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index206(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)206, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index207(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)207, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index208(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)208, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index209(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)209, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index210(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)210, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index211(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)211, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index212(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)212, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index213(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)213, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index214(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)214, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index215(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)215, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index216(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)216, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index217(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)217, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index218(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)218, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index219(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)219, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index220(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)220, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index221(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)221, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index222(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)222, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index223(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)223, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index224(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)224, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index225(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)225, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index226(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)226, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index227(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)227, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index228(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)228, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index229(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)229, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index230(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)230, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index231(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)231, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index232(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)232, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index233(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)233, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index234(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)234, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index235(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)235, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index236(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)236, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index237(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)237, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index238(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)238, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index239(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)239, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index240(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)240, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index241(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)241, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index242(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)242, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index243(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)243, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index244(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)244, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index245(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)245, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index246(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)246, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index247(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)247, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index248(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)248, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index249(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)249, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index250(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)250, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index251(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)251, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index252(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)252, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index253(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)253, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index254(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)254, (uintptr_t)p_instance,  p_args, r_ret);}
	void call_virtual_index255(GDExtensionClassInstancePtr p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret) {call_virtual_index((uint8_t)255, (uintptr_t)p_instance,  p_args, r_ret);}
	GDNativeExtensionClassCallVirtual get_virtual_func(void *p_userdata, const char *p_name) {
	switch (get_virtual_index(p_userdata, p_name)) {
		case 1: return call_virtual_index1;
		case 2: return call_virtual_index2;
		case 3: return call_virtual_index3;
		case 4: return call_virtual_index4;
		case 5: return call_virtual_index5;
		case 6: return call_virtual_index6;
		case 7: return call_virtual_index7;
		case 8: return call_virtual_index8;
		case 9: return call_virtual_index9;
		case 10: return call_virtual_index10;
		case 11: return call_virtual_index11;
		case 12: return call_virtual_index12;
		case 13: return call_virtual_index13;
		case 14: return call_virtual_index14;
		case 15: return call_virtual_index15;
		case 16: return call_virtual_index16;
		case 17: return call_virtual_index17;
		case 18: return call_virtual_index18;
		case 19: return call_virtual_index19;
		case 20: return call_virtual_index20;
		case 21: return call_virtual_index21;
		case 22: return call_virtual_index22;
		case 23: return call_virtual_index23;
		case 24: return call_virtual_index24;
		case 25: return call_virtual_index25;
		case 26: return call_virtual_index26;
		case 27: return call_virtual_index27;
		case 28: return call_virtual_index28;
		case 29: return call_virtual_index29;
		case 30: return call_virtual_index30;
		case 31: return call_virtual_index31;
		case 32: return call_virtual_index32;
		case 33: return call_virtual_index33;
		case 34: return call_virtual_index34;
		case 35: return call_virtual_index35;
		case 36: return call_virtual_index36;
		case 37: return call_virtual_index37;
		case 38: return call_virtual_index38;
		case 39: return call_virtual_index39;
		case 40: return call_virtual_index40;
		case 41: return call_virtual_index41;
		case 42: return call_virtual_index42;
		case 43: return call_virtual_index43;
		case 44: return call_virtual_index44;
		case 45: return call_virtual_index45;
		case 46: return call_virtual_index46;
		case 47: return call_virtual_index47;
		case 48: return call_virtual_index48;
		case 49: return call_virtual_index49;
		case 50: return call_virtual_index50;
		case 51: return call_virtual_index51;
		case 52: return call_virtual_index52;
		case 53: return call_virtual_index53;
		case 54: return call_virtual_index54;
		case 55: return call_virtual_index55;
		case 56: return call_virtual_index56;
		case 57: return call_virtual_index57;
		case 58: return call_virtual_index58;
		case 59: return call_virtual_index59;
		case 60: return call_virtual_index60;
		case 61: return call_virtual_index61;
		case 62: return call_virtual_index62;
		case 63: return call_virtual_index63;
		case 64: return call_virtual_index64;
		case 65: return call_virtual_index65;
		case 66: return call_virtual_index66;
		case 67: return call_virtual_index67;
		case 68: return call_virtual_index68;
		case 69: return call_virtual_index69;
		case 70: return call_virtual_index70;
		case 71: return call_virtual_index71;
		case 72: return call_virtual_index72;
		case 73: return call_virtual_index73;
		case 74: return call_virtual_index74;
		case 75: return call_virtual_index75;
		case 76: return call_virtual_index76;
		case 77: return call_virtual_index77;
		case 78: return call_virtual_index78;
		case 79: return call_virtual_index79;
		case 80: return call_virtual_index80;
		case 81: return call_virtual_index81;
		case 82: return call_virtual_index82;
		case 83: return call_virtual_index83;
		case 84: return call_virtual_index84;
		case 85: return call_virtual_index85;
		case 86: return call_virtual_index86;
		case 87: return call_virtual_index87;
		case 88: return call_virtual_index88;
		case 89: return call_virtual_index89;
		case 90: return call_virtual_index90;
		case 91: return call_virtual_index91;
		case 92: return call_virtual_index92;
		case 93: return call_virtual_index93;
		case 94: return call_virtual_index94;
		case 95: return call_virtual_index95;
		case 96: return call_virtual_index96;
		case 97: return call_virtual_index97;
		case 98: return call_virtual_index98;
		case 99: return call_virtual_index99;
		case 100: return call_virtual_index100;
		case 101: return call_virtual_index101;
		case 102: return call_virtual_index102;
		case 103: return call_virtual_index103;
		case 104: return call_virtual_index104;
		case 105: return call_virtual_index105;
		case 106: return call_virtual_index106;
		case 107: return call_virtual_index107;
		case 108: return call_virtual_index108;
		case 109: return call_virtual_index109;
		case 110: return call_virtual_index110;
		case 111: return call_virtual_index111;
		case 112: return call_virtual_index112;
		case 113: return call_virtual_index113;
		case 114: return call_virtual_index114;
		case 115: return call_virtual_index115;
		case 116: return call_virtual_index116;
		case 117: return call_virtual_index117;
		case 118: return call_virtual_index118;
		case 119: return call_virtual_index119;
		case 120: return call_virtual_index120;
		case 121: return call_virtual_index121;
		case 122: return call_virtual_index122;
		case 123: return call_virtual_index123;
		case 124: return call_virtual_index124;
		case 125: return call_virtual_index125;
		case 126: return call_virtual_index126;
		case 127: return call_virtual_index127;
		case 128: return call_virtual_index128;
		case 129: return call_virtual_index129;
		case 130: return call_virtual_index130;
		case 131: return call_virtual_index131;
		case 132: return call_virtual_index132;
		case 133: return call_virtual_index133;
		case 134: return call_virtual_index134;
		case 135: return call_virtual_index135;
		case 136: return call_virtual_index136;
		case 137: return call_virtual_index137;
		case 138: return call_virtual_index138;
		case 139: return call_virtual_index139;
		case 140: return call_virtual_index140;
		case 141: return call_virtual_index141;
		case 142: return call_virtual_index142;
		case 143: return call_virtual_index143;
		case 144: return call_virtual_index144;
		case 145: return call_virtual_index145;
		case 146: return call_virtual_index146;
		case 147: return call_virtual_index147;
		case 148: return call_virtual_index148;
		case 149: return call_virtual_index149;
		case 150: return call_virtual_index150;
		case 151: return call_virtual_index151;
		case 152: return call_virtual_index152;
		case 153: return call_virtual_index153;
		case 154: return call_virtual_index154;
		case 155: return call_virtual_index155;
		case 156: return call_virtual_index156;
		case 157: return call_virtual_index157;
		case 158: return call_virtual_index158;
		case 159: return call_virtual_index159;
		case 160: return call_virtual_index160;
		case 161: return call_virtual_index161;
		case 162: return call_virtual_index162;
		case 163: return call_virtual_index163;
		case 164: return call_virtual_index164;
		case 165: return call_virtual_index165;
		case 166: return call_virtual_index166;
		case 167: return call_virtual_index167;
		case 168: return call_virtual_index168;
		case 169: return call_virtual_index169;
		case 170: return call_virtual_index170;
		case 171: return call_virtual_index171;
		case 172: return call_virtual_index172;
		case 173: return call_virtual_index173;
		case 174: return call_virtual_index174;
		case 175: return call_virtual_index175;
		case 176: return call_virtual_index176;
		case 177: return call_virtual_index177;
		case 178: return call_virtual_index178;
		case 179: return call_virtual_index179;
		case 180: return call_virtual_index180;
		case 181: return call_virtual_index181;
		case 182: return call_virtual_index182;
		case 183: return call_virtual_index183;
		case 184: return call_virtual_index184;
		case 185: return call_virtual_index185;
		case 186: return call_virtual_index186;
		case 187: return call_virtual_index187;
		case 188: return call_virtual_index188;
		case 189: return call_virtual_index189;
		case 190: return call_virtual_index190;
		case 191: return call_virtual_index191;
		case 192: return call_virtual_index192;
		case 193: return call_virtual_index193;
		case 194: return call_virtual_index194;
		case 195: return call_virtual_index195;
		case 196: return call_virtual_index196;
		case 197: return call_virtual_index197;
		case 198: return call_virtual_index198;
		case 199: return call_virtual_index199;
		case 200: return call_virtual_index200;
		case 201: return call_virtual_index201;
		case 202: return call_virtual_index202;
		case 203: return call_virtual_index203;
		case 204: return call_virtual_index204;
		case 205: return call_virtual_index205;
		case 206: return call_virtual_index206;
		case 207: return call_virtual_index207;
		case 208: return call_virtual_index208;
		case 209: return call_virtual_index209;
		case 210: return call_virtual_index210;
		case 211: return call_virtual_index211;
		case 212: return call_virtual_index212;
		case 213: return call_virtual_index213;
		case 214: return call_virtual_index214;
		case 215: return call_virtual_index215;
		case 216: return call_virtual_index216;
		case 217: return call_virtual_index217;
		case 218: return call_virtual_index218;
		case 219: return call_virtual_index219;
		case 220: return call_virtual_index220;
		case 221: return call_virtual_index221;
		case 222: return call_virtual_index222;
		case 223: return call_virtual_index223;
		case 224: return call_virtual_index224;
		case 225: return call_virtual_index225;
		case 226: return call_virtual_index226;
		case 227: return call_virtual_index227;
		case 228: return call_virtual_index228;
		case 229: return call_virtual_index229;
		case 230: return call_virtual_index230;
		case 231: return call_virtual_index231;
		case 232: return call_virtual_index232;
		case 233: return call_virtual_index233;
		case 234: return call_virtual_index234;
		case 235: return call_virtual_index235;
		case 236: return call_virtual_index236;
		case 237: return call_virtual_index237;
		case 238: return call_virtual_index238;
		case 239: return call_virtual_index239;
		case 240: return call_virtual_index240;
		case 241: return call_virtual_index241;
		case 242: return call_virtual_index242;
		case 243: return call_virtual_index243;
		case 244: return call_virtual_index244;
		case 245: return call_virtual_index245;
		case 246: return call_virtual_index246;
		case 247: return call_virtual_index247;
		case 248: return call_virtual_index248;
		case 249: return call_virtual_index249;
		case 250: return call_virtual_index250;
		case 251: return call_virtual_index251;
		case 252: return call_virtual_index252;
		case 253: return call_virtual_index253;
		case 254: return call_virtual_index254;
		case 255: return call_virtual_index255;
		default: return NULL;}
	}

	static void *mem_alloc(GDNativeInterface *api, size_t p_bytes)
	{
		return api->mem_alloc(p_bytes);
	}
	static void *mem_free(GDNativeInterface *api, void *p_ptr)
	{
		api->mem_free(p_ptr);
	}
	static void *mem_realloc(GDNativeInterface *api, void *p_ptr, size_t p_bytes)
	{
		return api->mem_realloc(p_ptr, p_bytes);
	}

	static void print_error(GDNativeInterface *api, const char *p_description, const char *p_function, const char *p_file, int32_t p_line)
	{
		api->print_error(p_description, p_function, p_file, p_line);
	}
	static void print_warning(GDNativeInterface *api, const char *p_description, const char *p_function, const char *p_file, int32_t p_line)
	{
		api->print_warning(p_description, p_function, p_file, p_line);
	}
	static void print_script_error(GDNativeInterface *api, const char *p_description, const char *p_function, const char *p_file, int32_t p_line)
	{
		api->print_script_error(p_description, p_function, p_file, p_line);
	}

	static uint64_t get_native_struct_size(GDNativeInterface *api, const char *p_name)
	{
		return api->get_native_struct_size(p_name);
	}

	static void variant_new_copy(GDNativeInterface *api, GDNativeVariantPtr r_dest, const GDNativeVariantPtr p_src)
	{
		api->variant_new_copy(r_dest, p_src);
	}
	static void variant_new_nil(GDNativeInterface *api, GDNativeVariantPtr r_dest)
	{
		api->variant_new_nil(r_dest);
	}
	static void variant_destroy(GDNativeInterface *api, GDNativeVariantPtr p_self)
	{
		api->variant_destroy(p_self);
	}

	typedef struct
	{
		GDNativeBool valid;
		GDNativeBool oob;
	} VariantIndexingStatus;

	typedef struct
	{
		GDNativeBool valid;
		GDNativeBool more;
	} VariantIteratorStatus;

	typedef struct
	{
		GDNativeBool valid;
		GDNativeBool found;
	} VariantLookupStatus;

	static void variant_call(GDNativeInterface *api, GDNativeVariantPtr p_self, const GDNativeStringNamePtr p_method, const GDNativeVariantPtr *p_args, const GDNativeInt p_argument_count, GDNativeVariantPtr r_return, GDNativeCallError *r_error)
	{
		api->variant_call(p_self, p_method, p_args, p_argument_count, r_return, r_error);
	}
	static void variant_call_static(GDNativeInterface *api, GDNativeVariantType p_type, const GDNativeStringNamePtr p_method, const GDNativeVariantPtr *p_args, const GDNativeInt p_argument_count, GDNativeVariantPtr r_return, GDNativeCallError *r_error)
	{
		api->variant_call_static(p_type, p_method, p_args, p_argument_count, r_return, r_error);
	}
	static GDNativeBool variant_evaluate(GDNativeInterface *api, GDNativeVariantOperator p_op, const GDNativeVariantPtr p_a, const GDNativeVariantPtr p_b, GDNativeVariantPtr r_return)
	{
		GDNativeBool r_valid;
		api->variant_evaluate(p_op, p_a, p_b, r_return, &r_valid);
		return r_valid;
	}
	static GDNativeBool variant_set(GDNativeInterface *api, GDNativeVariantPtr p_self, const GDNativeVariantPtr p_key, const GDNativeVariantPtr p_value)
	{
		GDNativeBool r_valid;
		api->variant_set(p_self, p_key, p_value, &r_valid);
		return r_valid;
	}
	static GDNativeBool variant_set_named(GDNativeInterface *api, GDNativeVariantPtr p_self, const GDNativeStringNamePtr p_key, const GDNativeVariantPtr p_value)
	{
		GDNativeBool r_valid;
		api->variant_set_named(p_self, p_key, p_value, &r_valid);
		return r_valid;
	}
	static GDNativeBool variant_set_keyed(GDNativeInterface *api, GDNativeVariantPtr p_self, const GDNativeVariantPtr p_key, const GDNativeVariantPtr p_value)
	{
		GDNativeBool r_valid;
		api->variant_set_keyed(p_self, p_key, p_value, &r_valid);
		return r_valid;
	}
	static VariantIndexingStatus variant_set_indexed(GDNativeInterface *api, GDNativeVariantPtr p_self, GDNativeInt p_index, const GDNativeVariantPtr p_value)
	{
		VariantIndexingStatus status;
		api->variant_set_indexed(p_self, p_index, p_value, &status.valid, &status.oob);
		return status;
	}
	static GDNativeBool variant_get(GDNativeInterface *api, GDNativeVariantPtr p_self, const GDNativeVariantPtr p_key, GDNativeVariantPtr r_return)
	{
		GDNativeBool r_valid;
		api->variant_get(p_self, p_key, r_return, &r_valid);
		return r_valid;
	}
	static GDNativeBool variant_get_named(GDNativeInterface *api, GDNativeVariantPtr p_self, const GDNativeStringNamePtr p_key, GDNativeVariantPtr r_return)
	{
		GDNativeBool r_valid;
		api->variant_get_named(p_self, p_key, r_return, &r_valid);
		return r_valid;
	}
	static GDNativeBool variant_get_keyed(GDNativeInterface *api, GDNativeVariantPtr p_self, const GDNativeVariantPtr p_key, GDNativeVariantPtr r_return)
	{
		GDNativeBool r_valid;
		api->variant_get_keyed(p_self, p_key, r_return, &r_valid);
		return r_valid;
	}
	static VariantIndexingStatus variant_get_indexed(GDNativeInterface *api, GDNativeVariantPtr p_self, GDNativeInt p_index, GDNativeVariantPtr r_return)
	{
		VariantIndexingStatus status;
		api->variant_get_indexed(p_self, p_index, r_return, &status.valid, &status.oob);
		return status;
	}
	static VariantIteratorStatus variant_iter_init(GDNativeInterface *api, const GDNativeVariantPtr p_self, GDNativeVariantPtr r_iter)
	{
		VariantIteratorStatus status;
		status.more = api->variant_iter_init(p_self, r_iter, &status.valid);
		return status;
	}
	static VariantIteratorStatus variant_iter_next(GDNativeInterface *api, const GDNativeVariantPtr p_self, GDNativeVariantPtr r_iter)
	{
		VariantIteratorStatus status;
		status.more = api->variant_iter_next(p_self, r_iter, &status.valid);
		return status;
	}
	static GDNativeBool variant_iter_get(GDNativeInterface *api, const GDNativeVariantPtr p_self, GDNativeVariantPtr r_iter, GDNativeVariantPtr r_ret)
	{
		GDNativeBool r_valid;
		api->variant_iter_get(p_self, r_iter, r_ret, &r_valid);
		return r_valid;
	}
	static GDNativeInt variant_hash(GDNativeInterface *api, const GDNativeVariantPtr p_self)
	{
		return api->variant_hash(p_self);
	}
	static GDNativeInt variant_recursive_hash(GDNativeInterface *api, const GDNativeVariantPtr p_self, GDNativeInt p_recursion_count)
	{
		return api->variant_recursive_hash(p_self, p_recursion_count);
	}
	static GDNativeBool variant_hash_compare(GDNativeInterface *api, const GDNativeVariantPtr p_self, const GDNativeVariantPtr p_other)
	{
		return api->variant_hash_compare(p_self, p_other);
	}
	static GDNativeBool variant_booleanize(GDNativeInterface *api, const GDNativeVariantPtr p_self)
	{
		return api->variant_booleanize(p_self);
	}
	static void variant_sub(GDNativeInterface *api, const GDNativeVariantPtr p_a, const GDNativeVariantPtr p_b, GDNativeVariantPtr r_dst)
	{
		api->variant_sub(p_a, p_b, r_dst);
	}
	static void variant_blend(GDNativeInterface *api, const GDNativeVariantPtr p_a, const GDNativeVariantPtr p_b, float p_c, GDNativeVariantPtr r_dst)
	{
		api->variant_blend(p_a, p_b, p_c, r_dst);
	}
	static void variant_interpolate(GDNativeInterface *api, const GDNativeVariantPtr p_a, const GDNativeVariantPtr p_b, float p_c, GDNativeVariantPtr r_dst)
	{
		api->variant_interpolate(p_a, p_b, p_c, r_dst);
	}
	static void variant_duplicate(GDNativeInterface *api, const GDNativeVariantPtr p_self, GDNativeVariantPtr r_ret, GDNativeBool p_deep)
	{
		api->variant_duplicate(p_self, r_ret, p_deep);
	}
	static void variant_stringify(GDNativeInterface *api, const GDNativeVariantPtr p_self, GDNativeStringPtr r_ret)
	{
		api->variant_stringify(p_self, r_ret);
	}

	static GDNativeVariantType variant_get_type(GDNativeInterface *api, const GDNativeVariantPtr p_self)
	{
		return api->variant_get_type(p_self);
	}
	static GDNativeBool variant_has_method(GDNativeInterface *api, const GDNativeVariantPtr p_self, const GDNativeStringNamePtr p_method)
	{
		return api->variant_has_method(p_self, p_method);
	}
	static GDNativeBool variant_has_member(GDNativeInterface *api, GDNativeVariantType p_type, const GDNativeStringNamePtr p_member) {
		return api->variant_has_member(p_type, p_member);
	}
	static VariantLookupStatus variant_has_key(GDNativeInterface *api, const GDNativeVariantPtr p_self, const GDNativeVariantPtr p_key)
	{
		VariantLookupStatus status;
		status.found = api->variant_has_key(p_self, p_key, &status.valid);
		return status;
	}
	static void variant_get_type_name(GDNativeInterface *api, GDNativeVariantType p_type, GDNativeStringPtr r_name)
	{
		api->variant_get_type_name(p_type, r_name);
	}
	static GDNativeBool variant_can_convert(GDNativeInterface *api, GDNativeVariantType p_from, GDNativeVariantType p_to) {
		return api->variant_can_convert(p_from, p_to);
	}
	static GDNativeBool variant_can_convert_strict(GDNativeInterface *api, GDNativeVariantType p_from, GDNativeVariantType p_to) {
		return api->variant_can_convert_strict(p_from, p_to);
	}

	static GDNativeVariantFromTypeConstructorFunc get_variant_from_type_constructor(GDNativeInterface *api, GDNativeVariantType p_type)
	{
		return api->get_variant_from_type_constructor(p_type);
	}
	static GDNativeTypeFromVariantConstructorFunc get_variant_to_type_constructor(GDNativeInterface *api, GDNativeVariantType p_type)
	{
		return api->get_variant_to_type_constructor(p_type);
	}
	static GDNativePtrOperatorEvaluator variant_get_ptr_operator_evaluator(GDNativeInterface *api, GDNativeVariantOperator p_operator, GDNativeVariantType p_type_a, GDNativeVariantType p_type_b)
	{
		return api->variant_get_ptr_operator_evaluator(p_operator, p_type_a, p_type_b);
	}
	static GDNativePtrBuiltInMethod variant_get_ptr_builtin_method(GDNativeInterface *api, GDNativeVariantType p_type, const char *p_method, GDNativeInt p_hash)
	{
		return api->variant_get_ptr_builtin_method(p_type, p_method, p_hash);
	}
	static GDNativePtrConstructor variant_get_ptr_constructor(GDNativeInterface *api, GDNativeVariantType p_type, int32_t p_constructor)
	{
		return api->variant_get_ptr_constructor(p_type, p_constructor);
	}
	static GDNativePtrDestructor variant_get_ptr_destructor(GDNativeInterface *api, GDNativeVariantType p_type)
	{
		return api->variant_get_ptr_destructor(p_type);
	}
	static void variant_construct(GDNativeInterface *api, GDNativeVariantType p_type, GDNativeVariantPtr p_base, const GDNativeVariantPtr *p_args, int32_t p_argument_count, GDNativeCallError *r_error)
	{
		api->variant_construct(p_type, p_base, p_args, p_argument_count, r_error);
	}
	static GDNativePtrSetter variant_get_ptr_setter(GDNativeInterface *api, GDNativeVariantType p_type, const char *p_member) {
		return api->variant_get_ptr_setter(p_type, p_member);
	}
	static GDNativePtrGetter variant_get_ptr_getter(GDNativeInterface *api, GDNativeVariantType p_type, const char *p_member) {
		return api->variant_get_ptr_getter(p_type, p_member);
	}
	static GDNativePtrIndexedSetter variant_get_ptr_indexed_setter(GDNativeInterface *api, GDNativeVariantType p_type) {
		return api->variant_get_ptr_indexed_setter(p_type);
	}
	static GDNativePtrIndexedGetter variant_get_ptr_indexed_getter(GDNativeInterface *api, GDNativeVariantType p_type) {
		return api->variant_get_ptr_indexed_getter(p_type);
	}
	static GDNativePtrKeyedSetter variant_get_ptr_keyed_setter(GDNativeInterface *api, GDNativeVariantType p_type) {
		return api->variant_get_ptr_keyed_setter(p_type);
	}
	static GDNativePtrKeyedGetter variant_get_ptr_keyed_getter(GDNativeInterface *api, GDNativeVariantType p_type) {
		return api->variant_get_ptr_keyed_getter(p_type);
	}
	static GDNativePtrKeyedChecker variant_get_ptr_keyed_checker(GDNativeInterface *api, GDNativeVariantType p_type) {
		return api->variant_get_ptr_keyed_checker(p_type);
	}
	static void variant_get_constant_value(GDNativeInterface *api, GDNativeVariantType p_type, const char *p_constant, GDNativeVariantPtr r_ret) {
		api->variant_get_constant_value(p_type, p_constant, r_ret);
	}
	static GDNativePtrUtilityFunction variant_get_ptr_utility_function(GDNativeInterface *api, const char *p_function, GDNativeInt p_hash) {
		return api->variant_get_ptr_utility_function(p_function, p_hash);
	}

	static void string_new_with_utf8_chars_and_len(GDNativeInterface *api, GDNativeStringPtr r_dest, const char *p_contents, const GDNativeInt p_size)
	{
		api->string_new_with_utf8_chars_and_len(r_dest, p_contents, p_size);
	}
	static GDNativeInt string_to_utf8_chars(GDNativeInterface *api, const GDNativeStringPtr p_self, char *r_text, GDNativeInt p_max_write_length)
	{
		return api->string_to_utf8_chars(p_self, r_text, p_max_write_length);
	}

	static GDNativeVariantPtr array_operator_index(GDNativeInterface *api, uintptr_t p_self, GDNativeInt p_index)
	{
		return api->array_operator_index((GDNativeTypePtr*)p_self, p_index);
	}
	static GDNativeVariantPtr dictionary_operator_index(GDNativeInterface *api, uintptr_t p_self,  const GDNativeVariantPtr p_key)
	{
		return api->dictionary_operator_index((GDNativeTypePtr*)p_self, p_key);
	}

	static void object_method_bind_call(GDNativeInterface *api, uintptr_t p_method_bind, uintptr_t p_instance, const GDNativeVariantPtr *p_args, GDNativeInt p_arg_count, GDNativeVariantPtr r_ret, GDNativeCallError *r_error)
	{
		api->object_method_bind_call((GDNativeMethodBindPtr)p_method_bind, (GDNativeObjectPtr*)p_instance, p_args, p_arg_count, r_ret, r_error);
	}
	static void object_method_bind_ptrcall(GDNativeInterface *api, const GDNativeMethodBindPtr p_method_bind, uintptr_t p_instance, const GDNativeTypePtr *p_args, GDNativeTypePtr r_ret)
	{
		api->object_method_bind_ptrcall(p_method_bind, (GDNativeObjectPtr*)p_instance, p_args, r_ret);
	}
	static void object_destroy(GDNativeInterface *api, uintptr_t p_self)
	{
		api->object_destroy((GDNativeObjectPtr*)p_self);
	}
	static uintptr_t global_get_singleton(GDNativeInterface *api, const char *p_name)
	{
		return (uintptr_t)(api->global_get_singleton(p_name));
	}
	static uintptr_t object_get_instance_binding(GDNativeInterface *api, uintptr_t p_o, uintptr_t p_token, const GDNativeInstanceBindingCallbacks *p_callbacks)
	{
		return (uintptr_t)(api->object_get_instance_binding((GDNativeObjectPtr*)p_o, (void*)p_token, p_callbacks));
	}
	static void object_set_instance_binding(GDNativeInterface *api, uintptr_t p_o, uintptr_t p_token, uintptr_t p_binding, const GDNativeInstanceBindingCallbacks *p_callbacks)
	{
		api->object_set_instance_binding((GDNativeObjectPtr*)p_o, (void*)p_token, (void*)p_binding, p_callbacks);
	}
	static void object_set_instance(GDNativeInterface *api, uintptr_t p_o, const char *p_classname, uintptr_t p_instance)
	{
		api->object_set_instance((GDNativeObjectPtr*)p_o, p_classname, (GDExtensionClassInstancePtr)p_instance);
	}
	static uintptr_t object_cast_to(GDNativeInterface *api, const uintptr_t p_object, uintptr_t p_class_tag)
	{
		return (uintptr_t)(api->object_cast_to((GDNativeObjectPtr*)p_object, (void*)p_class_tag));
	}
	static uintptr_t object_get_instance_from_id(GDNativeInterface *api, GDObjectInstanceID p_instance_id)
	{
		return (uintptr_t)(api->object_get_instance_from_id(p_instance_id));
	}
	static GDObjectInstanceID object_get_instance_id(GDNativeInterface *api, uintptr_t p_object)
	{
		return api->object_get_instance_id((GDNativeObjectPtr*)p_object);
	}

	static GDNativeScriptInstancePtr script_instance_create(GDNativeInterface *api, const GDNativeExtensionScriptInstanceInfo *p_info, uintptr_t p_instance_data)
	{
		return api->script_instance_create(p_info, (GDNativeExtensionScriptInstanceDataPtr*)p_instance_data);
	}
	static GDNativeObjectPtr classdb_construct_object(GDNativeInterface *api, const char *p_classnam)
	{
		return api->classdb_construct_object(p_classnam);
	}
	static GDNativeMethodBindPtr classdb_get_method_bind(GDNativeInterface *api, const char *p_classname, const char *p_methodname, GDNativeInt p_hash)
	{
		return api->classdb_get_method_bind(p_classname, p_methodname, p_hash);
	}
	static void *classdb_get_class_tag(GDNativeInterface *api, const char *p_classname)
	{
		return api->classdb_get_class_tag(p_classname);
	}

	static void classdb_register_extension_class(GDNativeInterface *api, uintptr_t p_library, const char *p_class_name, const char *p_parent_class_name, uintptr_t userdata) {
		GDNativeExtensionClassCreationInfo p_extension_funcs;
		p_extension_funcs.class_userdata = (void*)userdata;
		p_extension_funcs.create_instance_func = (GDNativeExtensionClassCreateInstance)create_instance_func;
		p_extension_funcs.free_instance_func = (GDNativeExtensionClassFreeInstance)free_instance_func;
		p_extension_funcs.get_virtual_func = (GDNativeExtensionClassGetVirtual)get_virtual_func;
		api->classdb_register_extension_class((GDNativeExtensionClassLibraryPtr*)p_library, p_class_name, p_parent_class_name, &p_extension_funcs);
	}
	void classdb_register_extension_class_method(GDNativeInterface *api, uintptr_t p_library, const char *p_class_name, const GDNativeExtensionClassMethodInfo *p_method_info)
	{
		api->classdb_register_extension_class_method((GDNativeExtensionClassLibraryPtr*)p_library, p_class_name, p_method_info);
	}
	void classdb_register_extension_class_integer_constant(GDNativeInterface *api, uintptr_t p_library, const char *p_class_name, const char *p_enum_name, const char *p_constant_name, GDNativeInt p_constant_value, GDNativeBool p_is_bitfield)
	{
		api->classdb_register_extension_class_integer_constant((GDNativeExtensionClassLibraryPtr*)p_library, p_class_name, p_enum_name, p_constant_name, p_constant_value, p_is_bitfield);
	}
	void classdb_register_extension_class_property(GDNativeInterface *api, uintptr_t p_library, const char *p_class_name, const GDNativePropertyInfo *p_info, const char *p_setter, const char *p_getter) {
		api->classdb_register_extension_class_property((GDNativeExtensionClassLibraryPtr*)p_library, p_class_name, p_info, p_setter, p_getter);
	}
	void classdb_register_extension_class_property_group(GDNativeInterface *api, uintptr_t p_library, const char *p_class_name, const char *p_group_name, const char *p_prefix)
	{
		api->classdb_register_extension_class_property_group((GDNativeExtensionClassLibraryPtr*)p_library, p_class_name, p_group_name, p_prefix);
	}
	void classdb_register_extension_class_property_subgroup(GDNativeInterface *api, uintptr_t p_library, const char *p_class_name, const char *p_subgroup_name, const char *p_prefix)
	{
		api->classdb_register_extension_class_property_subgroup((GDNativeExtensionClassLibraryPtr*)p_library, p_class_name, p_subgroup_name, p_prefix);
	}
	void classdb_register_extension_class_signal(GDNativeInterface *api, uintptr_t p_library, const char *p_class_name, const char *p_signal_name, const GDNativePropertyInfo *p_argument_info, GDNativeInt p_argument_count)
	{
		api->classdb_register_extension_class_signal((GDNativeExtensionClassLibraryPtr*)p_library, p_class_name, p_signal_name, p_argument_info, p_argument_count);
	}
	void classdb_unregister_extension_class(GDNativeInterface *api, uintptr_t p_library, const char *p_class_name) {
		api->classdb_unregister_extension_class((GDNativeExtensionClassLibraryPtr*)p_library, p_class_name);
	}
	void get_library_path(GDNativeInterface *api, uintptr_t p_library, GDNativeStringPtr r_path)
	{
		api->get_library_path((GDNativeExtensionClassLibraryPtr*)p_library, r_path);
	}

	void call_variant_from_type_constructor(uintptr_t p_constructor, GDNativeVariantPtr r_variant, GDNativeTypePtr p_native) {
		((GDNativeVariantFromTypeConstructorFunc)p_constructor)(p_native, r_variant);
	}
	void call_variant_to_type_constructor(uintptr_t p_constructor, GDNativeTypePtr r_native, GDNativeVariantPtr p_variant) {
		((GDNativeTypeFromVariantConstructorFunc)p_constructor)(p_variant, r_native);
	}
	void call_variant_operator_evaluator(uintptr_t p_evaluator, GDNativeVariantPtr p_left, GDNativeVariantPtr p_right, GDNativeVariantPtr r_result) {
		((GDNativePtrOperatorEvaluator)p_evaluator)(p_left, p_right, r_result);
	}
	void call_variant_destructor(uintptr_t p_destructor, GDNativeTypePtr p_base) {
		((GDNativePtrDestructor)p_destructor)(p_base);
	}

	const char *ConstChar(_GoString_ s) {
		return _GoStringPtr(s);
	}
*/
import "C"
import "unsafe"

var (
	api     *cInterface
	classDB cClassLibrary
)

var cOnLoad func() = func() {}

func cInit(p_interface, p_library, r_initialization unsafe.Pointer) {
	api = (*cInterface)(p_interface)
	classDB = cClassLibrary(p_library)
	cOnLoad()
}

/*
	These are the GDNative types and each type and they need to be kept in sync
	with Godot's memory representation of them.

	For each release of Godot, only the default build options are supported.
*/
type (
	cNil     struct{}
	cBool    bool
	cInt     int64
	cFloat   float64
	cString  uintptr
	cVector2 struct {
		X, Y float32
	}
	cVector2i struct {
		X, Y int32
	}
	cRect2 struct {
		Position, Size cVector2
	}
	cRect2i struct {
		Position, Size cVector2i
	}
	cVector3 struct {
		X, Y, Z float32
	}
	cVector3i struct {
		X, Y, Z int32
	}
	cTransform2D struct {
		X, Y, Origin cVector2
	}
	cVector4 struct {
		X, Y, Z, W float32
	}
	cVector4i struct {
		X, Y, Z, W int32
	}
	cPlane struct {
		Matrix cVector4
	}
	cQuaternion struct {
		X, Y, Z, W float32
	}
	cAABB struct {
		Position, Size cVector3
	}
	cBasis struct {
		Rows [3]cVector3
	}
	cTransform3D struct {
		Basis  cBasis
		Origin cVector3
	}
	cProjection struct {
		X, Y, Z, W cVector4
	}
	cColor struct {
		R, G, B, A float32
	}
	cName               uintptr
	cNodePath           uintptr
	cRID                uint64
	cObject             uintptr
	cCallable           [3]uintptr
	cSignal             [2]uintptr
	cDictionary         uintptr
	cArray              uintptr
	cPackedByteArray    [2]uintptr
	cPackedInt32Array   [2]uintptr
	cPackedInt64Array   [2]uintptr
	cPackedFloat32Array [2]uintptr
	cPackedFloat64Array [2]uintptr
	cPackedStringArray  [2]uintptr
	cPackedVector2Array [2]uintptr
	cPackedVector3Array [2]uintptr
	cPackedColorArray   [2]uintptr
	cVariant            [24]byte
)

func cNewString(s string) (*C.char, bool) {
	if s[len(s)-1] == 0 {
		return C.ConstChar(s), true
	}

	length := C.size_t(len(s) + 1) // + null byte

	mem := C.mem_alloc(api, length)

	bytes := unsafe.Slice((*byte)(mem), length)

	copy(bytes, s)
	bytes[len(s)] = 0

	return (*C.char)(mem), false
}

type cVariantType C.GDNativeVariantType

const (
	cVariantTypeNil    = cVariantType(C.GDNATIVE_VARIANT_TYPE_NIL)
	cVariantTypeBool   = cVariantType(C.GDNATIVE_VARIANT_TYPE_BOOL)
	cVariantTypeInt    = cVariantType(C.GDNATIVE_VARIANT_TYPE_INT)
	cVariantTypeFloat  = cVariantType(C.GDNATIVE_VARIANT_TYPE_FLOAT)
	cVariantTypeString = cVariantType(C.GDNATIVE_VARIANT_TYPE_STRING)

	cVariantTypeVector2     = cVariantType(C.GDNATIVE_VARIANT_TYPE_VECTOR2)
	cVariantTypeVector2i    = cVariantType(C.GDNATIVE_VARIANT_TYPE_VECTOR2I)
	cVariantTypeRect2       = cVariantType(C.GDNATIVE_VARIANT_TYPE_RECT2)
	cVariantTypeRect2i      = cVariantType(C.GDNATIVE_VARIANT_TYPE_RECT2I)
	cVariantTypeVector3     = cVariantType(C.GDNATIVE_VARIANT_TYPE_VECTOR3)
	cVariantTypeVector3i    = cVariantType(C.GDNATIVE_VARIANT_TYPE_VECTOR3I)
	cVariantTypeTransform2D = cVariantType(C.GDNATIVE_VARIANT_TYPE_TRANSFORM2D)
	cVariantTypeVector4     = cVariantType(C.GDNATIVE_VARIANT_TYPE_VECTOR4)
	cVariantTypeVector4i    = cVariantType(C.GDNATIVE_VARIANT_TYPE_VECTOR4I)
	cVariantTypePlane       = cVariantType(C.GDNATIVE_VARIANT_TYPE_PLANE)
	cVariantTypeQuaternion  = cVariantType(C.GDNATIVE_VARIANT_TYPE_QUATERNION)
	cVariantTypeAABB        = cVariantType(C.GDNATIVE_VARIANT_TYPE_AABB)
	cVariantTypeBasis       = cVariantType(C.GDNATIVE_VARIANT_TYPE_BASIS)
	cVariantTypeTransform3D = cVariantType(C.GDNATIVE_VARIANT_TYPE_TRANSFORM3D)
	cVariantTypeProjection  = cVariantType(C.GDNATIVE_VARIANT_TYPE_PROJECTION)

	cVariantTypeColor      = cVariantType(C.GDNATIVE_VARIANT_TYPE_COLOR)
	cVariantTypeStringName = cVariantType(C.GDNATIVE_VARIANT_TYPE_STRING_NAME)
	cVariantTypeNodePath   = cVariantType(C.GDNATIVE_VARIANT_TYPE_NODE_PATH)
	cVariantTypeRID        = cVariantType(C.GDNATIVE_VARIANT_TYPE_RID)
	cVariantTypeObject     = cVariantType(C.GDNATIVE_VARIANT_TYPE_OBJECT)
	cVariantTypeCallable   = cVariantType(C.GDNATIVE_VARIANT_TYPE_CALLABLE)
	cVariantTypeSignal     = cVariantType(C.GDNATIVE_VARIANT_TYPE_SIGNAL)
	cVariantTypeDictionary = cVariantType(C.GDNATIVE_VARIANT_TYPE_DICTIONARY)
	cVariantTypeArray      = cVariantType(C.GDNATIVE_VARIANT_TYPE_ARRAY)

	cVariantTypePackedByteArray    = cVariantType(C.GDNATIVE_VARIANT_TYPE_PACKED_BYTE_ARRAY)
	cVariantTypePackedInt32Array   = cVariantType(C.GDNATIVE_VARIANT_TYPE_PACKED_INT32_ARRAY)
	cVariantTypePackedInt64Array   = cVariantType(C.GDNATIVE_VARIANT_TYPE_PACKED_INT64_ARRAY)
	cVariantTypePackedFloat32Array = cVariantType(C.GDNATIVE_VARIANT_TYPE_PACKED_FLOAT32_ARRAY)
	cVariantTypePackedFloat64Array = cVariantType(C.GDNATIVE_VARIANT_TYPE_PACKED_FLOAT64_ARRAY)
	cVariantTypePackedStringArray  = cVariantType(C.GDNATIVE_VARIANT_TYPE_PACKED_STRING_ARRAY)
	cVariantTypePackedVector2Array = cVariantType(C.GDNATIVE_VARIANT_TYPE_PACKED_VECTOR2_ARRAY)
	cVariantTypePackedVector3Array = cVariantType(C.GDNATIVE_VARIANT_TYPE_PACKED_VECTOR3_ARRAY)
	cVariantTypePackedColorArray   = cVariantType(C.GDNATIVE_VARIANT_TYPE_PACKED_COLOR_ARRAY)
)

var cVariantTypes = [...]cVariantType{
	cVariantTypeNil,
	cVariantTypeBool,
	cVariantTypeInt,
	cVariantTypeFloat,
	cVariantTypeString,

	cVariantTypeVector2,
	cVariantTypeVector2i,
	cVariantTypeRect2,
	cVariantTypeRect2i,
	cVariantTypeVector3,
	cVariantTypeVector3i,
	cVariantTypeTransform2D,
	cVariantTypeVector4,
	cVariantTypeVector4i,
	cVariantTypePlane,
	cVariantTypeQuaternion,
	cVariantTypeAABB,
	cVariantTypeBasis,
	cVariantTypeTransform3D,
	cVariantTypeProjection,

	cVariantTypeColor,
	cVariantTypeStringName,
	cVariantTypeNodePath,
	cVariantTypeRID,
	cVariantTypeObject,
	cVariantTypeCallable,
	cVariantTypeSignal,
	cVariantTypeDictionary,
	cVariantTypeArray,

	cVariantTypePackedByteArray,
	cVariantTypePackedInt32Array,
	cVariantTypePackedInt64Array,
	cVariantTypePackedFloat32Array,
	cVariantTypePackedFloat64Array,
	cVariantTypePackedStringArray,
	cVariantTypePackedVector2Array,
	cVariantTypePackedVector3Array,
	cVariantTypePackedColorArray,
}

type cVariantOperator C.GDNativeVariantOperator

const (
	cEqual        = cVariantOperator(C.GDNATIVE_VARIANT_OP_EQUAL)
	cNotEqual     = cVariantOperator(C.GDNATIVE_VARIANT_OP_NOT_EQUAL)
	cLess         = cVariantOperator(C.GDNATIVE_VARIANT_OP_LESS)
	cLessEqual    = cVariantOperator(C.GDNATIVE_VARIANT_OP_LESS_EQUAL)
	cGreater      = cVariantOperator(C.GDNATIVE_VARIANT_OP_GREATER)
	cGreaterEqual = cVariantOperator(C.GDNATIVE_VARIANT_OP_GREATER_EQUAL)

	cAdd = cVariantOperator(C.GDNATIVE_VARIANT_OP_ADD)
	cSub = cVariantOperator(C.GDNATIVE_VARIANT_OP_SUBTRACT)
	cMul = cVariantOperator(C.GDNATIVE_VARIANT_OP_MULTIPLY)
	cDiv = cVariantOperator(C.GDNATIVE_VARIANT_OP_DIVIDE)
	cNeg = cVariantOperator(C.GDNATIVE_VARIANT_OP_NEGATE)
	cPos = cVariantOperator(C.GDNATIVE_VARIANT_OP_POSITIVE)
	cMod = cVariantOperator(C.GDNATIVE_VARIANT_OP_MODULE)
	cPow = cVariantOperator(C.GDNATIVE_VARIANT_OP_POWER)

	cShiftLeft  = cVariantOperator(C.GDNATIVE_VARIANT_OP_SHIFT_LEFT)
	cShiftRight = cVariantOperator(C.GDNATIVE_VARIANT_OP_SHIFT_RIGHT)
	cBitAnd     = cVariantOperator(C.GDNATIVE_VARIANT_OP_BIT_AND)
	cBitOr      = cVariantOperator(C.GDNATIVE_VARIANT_OP_BIT_OR)
	cBitXor     = cVariantOperator(C.GDNATIVE_VARIANT_OP_BIT_XOR)
	cBitNot     = cVariantOperator(C.GDNATIVE_VARIANT_OP_BIT_NEGATE)

	cAnd = cVariantOperator(C.GDNATIVE_VARIANT_OP_AND)
	cOr  = cVariantOperator(C.GDNATIVE_VARIANT_OP_OR)
	cXor = cVariantOperator(C.GDNATIVE_VARIANT_OP_XOR)
	cNot = cVariantOperator(C.GDNATIVE_VARIANT_OP_NOT)

	cIn = C.GDNATIVE_VARIANT_OP_IN
)

type cCallErrorType = C.GDNativeCallErrorType

const (
	cCallOK                    = cCallErrorType(C.GDNATIVE_CALL_OK)
	cCallErrorInvalidMethod    = cCallErrorType(C.GDNATIVE_CALL_ERROR_INVALID_METHOD)
	cCallErrorInvalidArgument  = cCallErrorType(C.GDNATIVE_CALL_ERROR_INVALID_ARGUMENT)
	cCallErrorTooManyArguments = cCallErrorType(C.GDNATIVE_CALL_ERROR_TOO_MANY_ARGUMENTS)
	cCallErrorTooFewArguments  = cCallErrorType(C.GDNATIVE_CALL_ERROR_TOO_FEW_ARGUMENTS)
	cCallErrorInstanceIsNull   = cCallErrorType(C.GDNATIVE_CALL_ERROR_INSTANCE_IS_NULL)
	cCallErrorMethodNotConst   = cCallErrorType(C.GDNATIVE_CALL_ERROR_METHOD_NOT_CONST)
)

type cExtensionClassMethodFlags = C.GDNativeExtensionClassMethodFlags

const (
	cMethodNormal   = cExtensionClassMethodFlags(C.GDNATIVE_EXTENSION_METHOD_FLAG_NORMAL)
	cMethodEditor   = cExtensionClassMethodFlags(C.GDNATIVE_EXTENSION_METHOD_FLAG_EDITOR)
	cMethodConstant = cExtensionClassMethodFlags(C.GDNATIVE_EXTENSION_METHOD_FLAG_CONST)
	cMethodVirtual  = cExtensionClassMethodFlags(C.GDNATIVE_EXTENSION_METHOD_FLAG_VIRTUAL)
	cMethodVariadic = cExtensionClassMethodFlags(C.GDNATIVE_EXTENSION_METHOD_FLAG_VARARG)
	cMethodStatic   = cExtensionClassMethodFlags(C.GDNATIVE_EXTENSION_METHOD_FLAG_STATIC)
)

type cCallError = C.GDNativeCallError

type cInterface = C.GDNativeInterface

func mem_alloc(p_bytes uintptr) unsafe.Pointer {
	return C.mem_alloc(api, C.size_t(p_bytes))
}
func mem_realloc(p_ptr unsafe.Pointer, p_bytes uintptr) unsafe.Pointer {
	return C.mem_realloc(api, p_ptr, C.size_t(p_bytes))
}
func mem_free(p_ptr unsafe.Pointer) {
	C.mem_free(api, p_ptr)
}

func print_error(p_description, p_function, p_file string, p_line int32) {
	description, ok := cNewString(p_description)
	if !ok {
		defer mem_free(unsafe.Pointer(description))
	}
	function, ok := cNewString(p_function)
	if !ok {
		defer mem_free(unsafe.Pointer(function))
	}
	file, ok := cNewString(p_file)
	if !ok {
		defer mem_free(unsafe.Pointer(file))
	}
	C.print_error(api, description, function, file, C.int32_t(p_line))
}
func print_warning(p_description, p_function, p_file string, p_line int32) {
	description, ok := cNewString(p_description)
	if !ok {
		defer mem_free(unsafe.Pointer(description))
	}
	function, ok := cNewString(p_function)
	if !ok {
		defer mem_free(unsafe.Pointer(function))
	}
	file, ok := cNewString(p_file)
	if !ok {
		defer mem_free(unsafe.Pointer(file))
	}
	C.print_warning(api, description, function, file, C.int32_t(p_line))
}
func print_script_error(p_description, p_function, p_file string, p_line int32) {
	description, ok := cNewString(p_description)
	if !ok {
		defer mem_free(unsafe.Pointer(description))
	}
	function, ok := cNewString(p_function)
	if !ok {
		defer mem_free(unsafe.Pointer(function))
	}
	file, ok := cNewString(p_file)
	if !ok {
		defer mem_free(unsafe.Pointer(file))
	}
	C.print_script_error(api, description, function, file, C.int32_t(p_line))
}

func get_native_struct_size(p_name string) uint64 {
	name, ok := cNewString(p_name)
	if !ok {
		defer mem_free(unsafe.Pointer(name))
	}
	return uint64(C.get_native_struct_size(api, name))
}

func variant_new_nil(r_dest *cVariant) {
	C.variant_new_nil(api, C.GDNativeVariantPtr(r_dest))
}
func (p_src *cVariant) new_copy(r_dest *cVariant) {
	C.variant_new_copy(api, C.GDNativeVariantPtr(r_dest), C.GDNativeVariantPtr(p_src))
}
func (p_src *cVariant) destroy(p_self *cVariant) {
	C.variant_destroy(api, C.GDNativeVariantPtr(p_self))
}
func (p_self *cVariant) call(p_method *cName, p_args []cVariant, r_return *cVariant, r_error *cCallError) {
	C.variant_call(api, C.GDNativeVariantPtr(p_self), C.GDNativeStringNamePtr(p_method),
		(*C.GDNativeVariantPtr)(unsafe.Pointer(&p_args[0])), C.GDNativeInt(len(p_args)), C.GDNativeVariantPtr(r_return), r_error)
}
func (p_self *cVariant) call_static(p_type cVariantType, p_method *cName, p_args []cVariant, r_return *cVariant, r_error *cCallError) {
	C.variant_call_static(api, C.GDNativeVariantType(p_type), C.GDNativeStringNamePtr(p_method),
		(*C.GDNativeVariantPtr)(unsafe.Pointer(&p_args[0])), C.GDNativeInt(len(p_args)), C.GDNativeVariantPtr(r_return), r_error)
}
func (p_self *cVariant) evaluate(p_op cVariantOperator, p_b, p_return *cVariant) (valid bool) {
	return C.variant_evaluate(api, C.GDNativeVariantOperator(p_op), C.GDNativeVariantPtr(p_self), C.GDNativeVariantPtr(p_b), C.GDNativeVariantPtr(p_return)) != 0
}
func (p_self *cVariant) set(p_key, p_value *cVariant) (valid bool) {
	return C.variant_set(api, C.GDNativeVariantPtr(p_self), C.GDNativeVariantPtr(p_key), C.GDNativeVariantPtr(p_value)) != 0
}
func (p_self *cVariant) set_named(p_key *cName, p_value *cVariant) (ok bool) {
	return C.variant_set_named(api, C.GDNativeVariantPtr(p_self), C.GDNativeStringNamePtr(p_key), C.GDNativeVariantPtr(p_value)) != 0
}
func (p_self *cVariant) set_keyed(p_key, p_value *cVariant) (ok bool) {
	return C.variant_set_keyed(api, C.GDNativeVariantPtr(p_self), C.GDNativeVariantPtr(p_key), C.GDNativeVariantPtr(p_value)) != 0
}
func (p_self *cVariant) set_indexed(p_index int, p_value *cVariant) (oob, ok bool) {
	status := C.variant_set_indexed(api, C.GDNativeVariantPtr(p_self), C.GDNativeInt(p_index), C.GDNativeVariantPtr(p_value))
	return status.oob != 0, status.valid != 0
}
func (p_self *cVariant) get(p_key *cVariant, r_value *cVariant) (ok bool) {
	return C.variant_get(api, C.GDNativeVariantPtr(p_self), C.GDNativeVariantPtr(p_key), C.GDNativeVariantPtr(r_value)) != 0
}
func (p_self *cVariant) get_named(p_key *cName, r_value *cVariant) (ok bool) {
	return C.variant_get_named(api, C.GDNativeVariantPtr(p_self), C.GDNativeStringNamePtr(p_key), C.GDNativeVariantPtr(r_value)) != 0
}
func (p_self *cVariant) get_keyed(p_key *cVariant, r_value *cVariant) (ok bool) {
	return C.variant_get_keyed(api, C.GDNativeVariantPtr(p_self), C.GDNativeVariantPtr(p_key), C.GDNativeVariantPtr(r_value)) != 0
}
func (p_self *cVariant) get_indexed(p_index int, r_value *cVariant) (oob, ok bool) {
	status := C.variant_get_indexed(api, C.GDNativeVariantPtr(p_self), C.GDNativeInt(p_index), C.GDNativeVariantPtr(r_value))
	return status.oob != 0, status.valid != 0
}
func (p_self *cVariant) iter_init(r_iter *cVariant) (more, ok bool) {
	result := C.variant_iter_init(api, C.GDNativeVariantPtr(p_self), C.GDNativeVariantPtr(r_iter))
	return result.more != 0, result.valid != 0
}
func (p_self *cVariant) iter_next(r_iter *cVariant) (more, ok bool) {
	result := C.variant_iter_next(api, C.GDNativeVariantPtr(p_self), C.GDNativeVariantPtr(r_iter))
	return result.more != 0, result.valid != 0
}
func (p_self *cVariant) iter_get(r_ite, r_ret *cVariant) (ok bool) {
	return C.variant_iter_get(api, C.GDNativeVariantPtr(p_self), C.GDNativeVariantPtr(r_ite), C.GDNativeVariantPtr(r_ret)) != 0
}
func (p_self *cVariant) hash() int64 {
	return int64(C.variant_hash(api, C.GDNativeVariantPtr(p_self)))
}
func (p_self *cVariant) recursive_hash(p_recursion_count int64) int64 {
	return int64(C.variant_recursive_hash(api, C.GDNativeVariantPtr(p_self), C.GDNativeInt(p_recursion_count)))
}
func (p_self *cVariant) hash_compare(p_other *cVariant) bool {
	return C.variant_hash_compare(api, C.GDNativeVariantPtr(p_self), C.GDNativeVariantPtr(p_other)) != 0
}
func (p_self *cVariant) booleanize() bool {
	return C.variant_booleanize(api, C.GDNativeVariantPtr(p_self)) != 0
}
func (p_self *cVariant) sub(p_b, r_dst *cVariant) {
	C.variant_sub(api, C.GDNativeVariantPtr(p_self), C.GDNativeVariantPtr(p_b), C.GDNativeVariantPtr(r_dst))
}
func (p_self *cVariant) blend(p_b *cVariant, p_c float64, r_dst *cVariant) {
	C.variant_blend(api, C.GDNativeVariantPtr(p_self), C.GDNativeVariantPtr(p_b), C.float(p_c), C.GDNativeVariantPtr(r_dst))
}
func (p_self *cVariant) interpolate(p_b *cVariant, p_c float64, r_dst *cVariant) {
	C.variant_interpolate(api, C.GDNativeVariantPtr(p_self), C.GDNativeVariantPtr(p_b), C.float(p_c), C.GDNativeVariantPtr(r_dst))
}
func (p_self *cVariant) duplicate(r_ret *cVariant, p_deep bool) {
	var deep C.GDNativeBool = 0
	if p_deep {
		deep = 1
	}
	C.variant_duplicate(api, C.GDNativeVariantPtr(p_self), C.GDNativeVariantPtr(r_ret), deep)
}
func (p_self *cVariant) stringify(r_ret *cString) {
	C.variant_stringify(api, C.GDNativeVariantPtr(p_self), C.GDNativeStringPtr(r_ret))
}

type cVariantFromTypeConstructorFunc uintptr

func get_variant_from_type_constructor(p_type cVariantType) cVariantFromTypeConstructorFunc {
	return cVariantFromTypeConstructorFunc(unsafe.Pointer(C.get_variant_from_type_constructor(api, C.GDNativeVariantType(p_type))))
}
func (fn cVariantFromTypeConstructorFunc) Call(r_variant *cVariant, p_native unsafe.Pointer) {
	C.call_variant_from_type_constructor(C.uintptr_t(fn), C.GDNativeVariantPtr(r_variant), C.GDNativeTypePtr(p_native))
}

type cNativeTypeFromVariantConstructorFunc uintptr

func get_variant_to_type_constructor(p_type cVariantType) cNativeTypeFromVariantConstructorFunc {
	return cNativeTypeFromVariantConstructorFunc(unsafe.Pointer(C.get_variant_to_type_constructor(api, C.GDNativeVariantType(p_type))))
}
func (fn cNativeTypeFromVariantConstructorFunc) Call(r_native unsafe.Pointer, p_variant *cVariant) {
	C.call_variant_to_type_constructor(C.uintptr_t(fn), C.GDNativeTypePtr(r_native), C.GDNativeVariantPtr(p_variant))
}

type cOperatorEvaluator uintptr

func (p_op cVariantOperator) get_evaluator(p_type_a, p_type_b cVariantType) cOperatorEvaluator {
	return cOperatorEvaluator(unsafe.Pointer(C.variant_get_ptr_operator_evaluator(api, C.GDNativeVariantOperator(p_op), C.GDNativeVariantType(p_type_a), C.GDNativeVariantType(p_type_b))))
}
func (fn cOperatorEvaluator) call(left, right, result *cVariantType) {
	C.call_variant_operator_evaluator(C.uintptr_t(fn), C.GDNativeVariantPtr(left), C.GDNativeVariantPtr(right), C.GDNativeVariantPtr(result))
}

type cBuiltInMethod = C.GDNativePtrBuiltInMethod

func (p_type cVariantType) get_builtin_method(p_method string, hash int64) cBuiltInMethod {
	method, ok := cNewString(p_method)
	if !ok {
		defer mem_free(unsafe.Pointer(method))
	}
	return C.GDNativePtrBuiltInMethod(C.variant_get_ptr_builtin_method(api, C.GDNativeVariantType(p_type), method, C.GDNativeInt(hash)))
}

type cConstructor = C.GDNativePtrConstructor

func (p_type cVariantType) get_constructor(p_constructor int32) cConstructor {
	return C.GDNativePtrConstructor(C.variant_get_ptr_constructor(api, C.GDNativeVariantType(p_type), C.int32_t(p_constructor)))
}

type cDestructor uintptr

func (p_type cVariantType) get_destructor() cDestructor {
	return cDestructor(unsafe.Pointer(C.variant_get_ptr_destructor(api, C.GDNativeVariantType(p_type))))
}
func (d cDestructor) call(p_base unsafe.Pointer) {
	C.call_variant_destructor(C.uintptr_t(d), C.GDNativeTypePtr(p_base))
}

func (p_type cVariantType) construct(p_base *cVariant, p_args []cVariant, r_error *cCallError) {
	C.variant_construct(api, C.GDNativeVariantType(p_type), C.GDNativeVariantPtr(p_base),
		(*C.GDNativeVariantPtr)(unsafe.Pointer(&p_args[0])), C.int32_t(len(p_args)), r_error)
}

type Setter = C.GDNativePtrSetter

func (p_type cVariantType) get_setter(p_member string) Setter {
	member, ok := cNewString(p_member)
	if !ok {
		defer mem_free(unsafe.Pointer(member))
	}
	return C.GDNativePtrSetter(C.variant_get_ptr_setter(api, C.GDNativeVariantType(p_type), member))
}

type cGetter = C.GDNativePtrGetter

func (p_type cVariantType) get_getter(p_member string) cGetter {
	member, ok := cNewString(p_member)
	if !ok {
		defer mem_free(unsafe.Pointer(member))
	}
	return C.GDNativePtrGetter(C.variant_get_ptr_getter(api, C.GDNativeVariantType(p_type), member))
}

type cIndexedSetter = C.GDNativePtrIndexedSetter

func (p_type cVariantType) get_indexed_setter() cIndexedSetter {
	return C.GDNativePtrIndexedSetter(C.variant_get_ptr_indexed_setter(api, C.GDNativeVariantType(p_type)))
}

type cIndexedGetter = C.GDNativePtrIndexedGetter

func (p_type cVariantType) get_indexed_getter() cIndexedGetter {
	return C.GDNativePtrIndexedGetter(C.variant_get_ptr_indexed_getter(api, C.GDNativeVariantType(p_type)))
}

type cKeyedSetter = C.GDNativePtrKeyedSetter

func (p_type cVariantType) get_keyed_setter() cKeyedSetter {
	return C.GDNativePtrKeyedSetter(C.variant_get_ptr_keyed_setter(api, C.GDNativeVariantType(p_type)))
}

type cKeyedGetter = C.GDNativePtrKeyedGetter

func (p_type cVariantType) get_keyed_getter() cKeyedGetter {
	return C.GDNativePtrKeyedGetter(C.variant_get_ptr_keyed_getter(api, C.GDNativeVariantType(p_type)))
}

type cKeyedChecker = C.GDNativePtrKeyedChecker

func (p_type cVariantType) get_keyed_checker() cKeyedChecker {
	return C.GDNativePtrKeyedChecker(C.variant_get_ptr_keyed_checker(api, C.GDNativeVariantType(p_type)))
}

func (p_type cVariantType) get_constant_value(p_constant string, r_ret *cVariant) {
	constant, ok := cNewString(p_constant)
	if !ok {
		defer mem_free(unsafe.Pointer(constant))
	}
	C.variant_get_constant_value(api, C.GDNativeVariantType(p_type), constant, C.GDNativeVariantPtr(r_ret))
}

type cUtilityFunction uintptr

func get_utility_function(p_function string, p_hash int64) cUtilityFunction {
	function, ok := cNewString(p_function)
	if !ok {
		defer mem_free(unsafe.Pointer(function))
	}
	return cUtilityFunction(unsafe.Pointer(C.variant_get_ptr_utility_function(api, function, C.GDNativeInt(p_hash))))
}

func (r_dest *cString) with_utf8_chars(p_string string) {
	s, ok := cNewString(p_string)
	if !ok {
		defer mem_free(unsafe.Pointer(s))
	}
	C.string_new_with_utf8_chars_and_len(api, C.GDNativeStringPtr(r_dest), s, C.GDNativeInt(len(p_string)))
}

func (s *cString) to_utf8_chars() string {
	l := C.string_to_utf8_chars(api, C.GDNativeStringPtr(s), nil, 0)
	b := make([]byte, l)
	C.string_to_utf8_chars(api, C.GDNativeStringPtr(s), (*C.char)(unsafe.Pointer(&b[0])), C.GDNativeInt(len(b)))
	return string(b)
}

func (a cArray) Index(index int) *cVariant {
	return (*cVariant)(C.array_operator_index(api, C.uintptr_t(a), C.GDNativeInt(index)))
}
func (d cDictionary) Index(p_key *cVariant) *cVariant {
	return (*cVariant)(C.dictionary_operator_index(api, C.uintptr_t(d), C.GDNativeVariantPtr(p_key)))
}

type cMethodBind uintptr

func (p_o cObject) call(p_method_bind cMethodBind, p_args []cVariant, r_ret *cVariant, r_error *cCallError) {
	C.object_method_bind_call(api, C.uintptr_t(p_method_bind), C.uintptr_t(p_o),
		(*C.GDNativeVariantPtr)(unsafe.Pointer(&p_args[0])), C.GDNativeInt(len(p_args)), C.GDNativeVariantPtr(r_ret), r_error)
}
func (p_o cObject) destroy() {
	C.object_destroy(api, C.uintptr_t(p_o))
}

func global_get_singleton(p_name string) cObject {
	name, ok := cNewString(p_name)
	if !ok {
		defer mem_free(unsafe.Pointer(name))
	}
	return cObject(C.global_get_singleton(api, name))
}

type InstanceBindingCallbacks = C.GDNativeInstanceBindingCallbacks

func (p_o cObject) get_instance_binding(p_token uintptr, p_callbacks *InstanceBindingCallbacks) uintptr {
	return uintptr(C.object_get_instance_binding(api, C.uintptr_t(p_o), C.uintptr_t(p_token), p_callbacks))
}
func (p_o cObject) set_instance_binding(p_token uintptr, p_binding uintptr, p_callbacks *InstanceBindingCallbacks) {
	C.object_set_instance_binding(api, C.uintptr_t(p_o), C.uintptr_t(p_token), C.uintptr_t(p_binding), p_callbacks)
}
func (p_o cObject) set_instance(p_classname string, p_instance uintptr) {
	classname, ok := cNewString(p_classname)
	if !ok {
		defer mem_free(unsafe.Pointer(classname))
	}
	C.object_set_instance(api, C.uintptr_t(p_o), classname, C.uintptr_t(p_instance))
}
func (p_o cObject) cast_to(p_class_tag cClassTag) cObject {
	return (cObject)(C.object_cast_to(api, C.uintptr_t(p_o), C.uintptr_t(p_class_tag)))
}

type cObjectInstanceID = C.GDObjectInstanceID

func (p_o cObject) get_instance_from_id(p_id cObjectInstanceID) cObject {
	return (cObject)(C.object_get_instance_from_id(api, cObjectInstanceID(p_id)))
}
func (p_o cObject) get_instance_id() cObjectInstanceID {
	return cObjectInstanceID(C.object_get_instance_id(api, C.uintptr_t(p_o)))
}

type cExtensionScriptInstanceInfo = C.GDNativeExtensionScriptInstanceInfo

type cScriptInstance = C.GDNativeScriptInstancePtr

func script_instance_create(p_info *cExtensionScriptInstanceInfo, p_instance_data uintptr) cScriptInstance {
	return C.GDNativeScriptInstancePtr(C.script_instance_create(api, p_info, C.uintptr_t(p_instance_data)))
}

type cClassTag uintptr

type cClassLibrary uintptr

func (cClassLibrary) construct_object(p_classname string) cObject {
	classname, ok := cNewString(p_classname)
	if !ok {
		defer mem_free(unsafe.Pointer(classname))
	}
	return (cObject)(C.classdb_construct_object(api, classname))
}
func (cClassLibrary) get_method_bind(p_classname string, p_method string, hash int64) cMethodBind {
	classname, ok := cNewString(p_classname)
	if !ok {
		defer mem_free(unsafe.Pointer(classname))
	}
	method, ok := cNewString(p_method)
	if !ok {
		defer mem_free(unsafe.Pointer(method))
	}
	return cMethodBind(C.classdb_get_method_bind(api, classname, method, C.GDNativeInt(hash)))
}
func (cClassLibrary) get_class_tag(p_classname string) cClassTag {
	classname, ok := cNewString(p_classname)
	if !ok {
		defer mem_free(unsafe.Pointer(classname))
	}
	return cClassTag(C.classdb_get_class_tag(api, classname))
}

type cExtensionClassCreationInfo = C.GDNativeExtensionClassCreationInfo

func (p_library cClassLibrary) register_extension_class(p_classname, p_parent_class_name string, userdata uintptr) {
	classname, ok := cNewString(p_classname)
	if !ok {
		defer mem_free(unsafe.Pointer(classname))
	}
	parent, ok := cNewString(p_parent_class_name)
	if !ok {
		defer mem_free(unsafe.Pointer(parent))
	}
	C.classdb_register_extension_class(api, C.uintptr_t(p_library), classname, parent, C.uintptr_t(userdata))
}

type cExtensionClassMethodInfo = C.GDNativeExtensionClassMethodInfo

func (p_library cClassLibrary) register_extension_class_method(p_classname string, p_method_info *cExtensionClassMethodInfo) {
	classname, ok := cNewString(p_classname)
	if !ok {
		defer mem_free(unsafe.Pointer(classname))
	}
	C.classdb_register_extension_class_method(api, C.uintptr_t(p_library), classname, p_method_info)
}

func (p_library cClassLibrary) register_extension_class_integer_constant(p_classname string, p_enum_name string, p_constant_name string, p_constant_value int, p_is_bitfield bool) {
	classname, ok := cNewString(p_classname)
	if !ok {
		defer mem_free(unsafe.Pointer(classname))
	}
	enum, ok := cNewString(p_enum_name)
	if !ok {
		defer mem_free(unsafe.Pointer(enum))
	}
	constant, ok := cNewString(p_constant_name)
	if !ok {
		defer mem_free(unsafe.Pointer(constant))
	}
	var bitfield C.GDNativeBool = 0
	if p_is_bitfield {
		bitfield = 1
	}
	C.classdb_register_extension_class_integer_constant(api, C.uintptr_t(p_library), classname, enum, constant, C.GDNativeInt(p_constant_value), bitfield)
}

type cPropertyInfo = C.GDNativePropertyInfo

func (p_library cClassLibrary) register_extension_class_property(p_classname string, p_info *cPropertyInfo, p_setter, p_getter string) {
	classname, ok := cNewString(p_classname)
	if !ok {
		defer mem_free(unsafe.Pointer(classname))
	}
	setter, ok := cNewString(p_setter)
	if !ok {
		defer mem_free(unsafe.Pointer(setter))
	}
	getter, ok := cNewString(p_getter)
	if !ok {
		defer mem_free(unsafe.Pointer(getter))
	}
	C.classdb_register_extension_class_property(api, C.uintptr_t(p_library), classname, p_info, setter, getter)
}

func (p_library cClassLibrary) register_extension_class_property_group(p_classname, p_group_name, p_prefix string) {
	classname, ok := cNewString(p_classname)
	if !ok {
		defer mem_free(unsafe.Pointer(classname))
	}
	group, ok := cNewString(p_group_name)
	if !ok {
		defer mem_free(unsafe.Pointer(group))
	}
	prefix, ok := cNewString(p_prefix)
	if !ok {
		defer mem_free(unsafe.Pointer(prefix))
	}
	C.classdb_register_extension_class_property_group(api, C.uintptr_t(p_library), classname, group, prefix)
}

func (p_library cClassLibrary) register_extension_class_property_subgroup(p_classname, p_subgroup_name, p_prefix string) {
	classname, ok := cNewString(p_classname)
	if !ok {
		defer mem_free(unsafe.Pointer(classname))
	}
	subgroup, ok := cNewString(p_subgroup_name)
	if !ok {
		defer mem_free(unsafe.Pointer(subgroup))
	}
	prefix, ok := cNewString(p_prefix)
	if !ok {
		defer mem_free(unsafe.Pointer(prefix))
	}
	C.classdb_register_extension_class_property_subgroup(api, C.uintptr_t(p_library), classname, subgroup, prefix)
}

func (p_library cClassLibrary) register_extension_class_signal(p_classname string, p_signal_name string, p_argument_info []cPropertyInfo) {
	classname, ok := cNewString(p_classname)
	if !ok {
		defer mem_free(unsafe.Pointer(classname))
	}
	signal, ok := cNewString(p_signal_name)
	if !ok {
		defer mem_free(unsafe.Pointer(signal))
	}
	C.classdb_register_extension_class_signal(api, C.uintptr_t(p_library), classname, signal, (*C.GDNativePropertyInfo)(unsafe.Pointer(&p_argument_info[0])), C.GDNativeInt(len(p_argument_info)))
}

func (p_library cClassLibrary) unregister_extension_class(p_classname string) {
	classname, ok := cNewString(p_classname)
	if !ok {
		defer mem_free(unsafe.Pointer(classname))
	}
	C.classdb_unregister_extension_class(api, C.uintptr_t(p_library), classname)
}

func (p_library cClassLibrary) get_library_path(r_path *cString) {
	C.get_library_path(api, C.uintptr_t(p_library), C.GDNativeStringPtr(r_path))
}

// TODO
type AudioFrame struct{}
type PhysicsServer3DExtensionRayResult struct{}
type PhysicsServer3DExtensionShapeResult struct{}
type PhysicsServer3DExtensionShapeRestInfo struct{}
type PhysicsServer3DExtensionMotionResult struct{}
type ScriptLanguageExtensionProfilingInfo struct{}
type Glyph struct{}
type CaretInfo struct{}
