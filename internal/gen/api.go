package main

type Enum struct {
	Name   string `json:"name"`
	Values []struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"values"`
}

type Method struct {
	Name        string `json:"name"`
	IsConst     bool   `json:"is_const"`
	IsVararg    bool   `json:"is_vararg"`
	IsStatic    bool   `json:"is_static"`
	IsVirtual   bool   `json:"is_virtual"`
	Hash        int64  `json:"hash"`
	ReturnValue struct {
		Type string `json:"type"`
		Meta string `json:"meta"`
	} `json:"return_value,omitempty"`
	Arguments []struct {
		Name         string `json:"name"`
		Type         string `json:"type"`
		Meta         string `json:"meta"`
		DefaultValue string `json:"default_value,omitempty"`
	} `json:"arguments,omitempty"`
}

/*
	Specification of the Godot Extension API.
	Created with https://mholt.github.io/json-to-go/
*/
type Specification struct {
	Header struct {
		VersionMajor    int    `json:"version_major"`
		VersionMinor    int    `json:"version_minor"`
		VersionPatch    int    `json:"version_patch"`
		VersionStatus   string `json:"version_status"`
		VersionBuild    string `json:"version_build"`
		VersionFullName string `json:"version_full_name"`
	} `json:"header"`
	BuiltinClassSizes []struct {
		BuildConfiguration string `json:"build_configuration"`
		Sizes              []struct {
			Name string `json:"name"`
			Size int    `json:"size"`
		} `json:"sizes"`
	} `json:"builtin_class_sizes"`
	BuiltinClassMemberOffsets []struct {
		BuildConfiguration string `json:"build_configuration"`
		Classes            []struct {
			Name    string `json:"name"`
			Members []struct {
				Member string `json:"member"`
				Offset int    `json:"offset"`
			} `json:"members"`
		} `json:"classes"`
	} `json:"builtin_class_member_offsets"`
	GlobalConstants  []interface{} `json:"global_constants"`
	GlobalEnums      []Enum        `json:"global_enums"`
	UtilityFunctions []struct {
		Name       string `json:"name"`
		ReturnType string `json:"return_type,omitempty"`
		Category   string `json:"category"`
		IsVararg   bool   `json:"is_vararg"`
		Hash       int    `json:"hash"`
		Arguments  []struct {
			Name string `json:"name"`
			Type string `json:"type"`
			Meta string `json:"meta"`
		} `json:"arguments,omitempty"`
	} `json:"utility_functions"`
	BuiltinClasses []struct {
		Name      string `json:"name"`
		IsKeyed   bool   `json:"is_keyed"`
		Operators []struct {
			Name       string `json:"name"`
			RightType  string `json:"right_type,omitempty"`
			ReturnType string `json:"return_type"`
		} `json:"operators"`
		Constructors []struct {
			Index     int `json:"index"`
			Arguments []struct {
				Name string `json:"name"`
				Type string `json:"type"`
				Meta string `json:"meta"`
			} `json:"arguments,omitempty"`
		} `json:"constructors"`
		HasDestructor      bool   `json:"has_destructor"`
		IndexingReturnType string `json:"indexing_return_type,omitempty"`
		Methods            []struct {
			Name       string `json:"name"`
			ReturnType string `json:"return_type"`
			IsVararg   bool   `json:"is_vararg"`
			IsConst    bool   `json:"is_const"`
			IsStatic   bool   `json:"is_static"`
			Hash       int64  `json:"hash"`
			Arguments  []struct {
				Name string `json:"name"`
				Type string `json:"type"`
				Meta string `json:"meta"`
			} `json:"arguments,omitempty"`
		} `json:"methods,omitempty"`
		Members []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"members,omitempty"`
		Constants []struct {
			Name  string `json:"name"`
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"constants,omitempty"`
		Enums []Enum `json:"enums,omitempty"`
	} `json:"builtin_classes"`
	Classes []struct {
		Name           string   `json:"name"`
		IsRefcounted   bool     `json:"is_refcounted"`
		IsInstantiable bool     `json:"is_instantiable"`
		Inherits       string   `json:"inherits,omitempty"`
		APIType        string   `json:"api_type"`
		Enums          []Enum   `json:"enums,omitempty"`
		Methods        []Method `json:"methods,omitempty"`
		Signals        []struct {
			Name      string `json:"name"`
			Arguments []struct {
				Name string `json:"name"`
				Type string `json:"type"`
				Meta string `json:"meta"`
			} `json:"arguments,omitempty"`
		} `json:"signals,omitempty"`
		Properties []struct {
			Type   string `json:"type"`
			Name   string `json:"name"`
			Setter string `json:"setter"`
			Getter string `json:"getter"`
			Index  int    `json:"index"`
		} `json:"properties,omitempty"`
		Constants []struct {
			Name  string `json:"name"`
			Value int    `json:"value"`
		} `json:"constants,omitempty"`
	} `json:"classes"`
	Singletons []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"singletons"`
	NativeStructures []struct {
		Name   string `json:"name"`
		Format string `json:"format"`
	} `json:"native_structures"`
}
