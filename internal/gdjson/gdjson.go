package gdjson

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
)

type Enum struct {
	Name   string `json:"name"`
	Values []struct {
		Name        string `json:"name"`
		Value       int    `json:"value"`
		Description string `json:"description"`
	} `json:"values"`
}

type Signal struct {
	Name      string `json:"name"`
	Arguments []struct {
		Name string `json:"name"`
		Type string `json:"type"`
		Meta string `json:"meta"`
	} `json:"arguments,omitempty"`
}

type Method struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsConst     bool   `json:"is_const"`
	IsVararg    bool   `json:"is_vararg"`
	IsStatic    bool   `json:"is_static"`
	IsVirtual   bool   `json:"is_virtual"`
	Hash        int64  `json:"hash"`
	ReturnType  string `json:"return_type"` // builtin class
	ReturnValue struct {
		Type string `json:"type"`
		Meta string `json:"meta"`
	} `json:"return_value,omitempty"`
	Arguments         []Argument `json:"arguments,omitempty"`
	ArgumentsRemapped []Argument
}
type Argument struct {
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	Meta         string  `json:"meta"`
	DefaultValue *string `json:"default_value"`
}

type Class struct {
	Name    string `json:"name"`
	Rename  string `json:"rename,omitempty"`
	Package string `json:"package,omitempty"`
	IsEnum  bool   `json:"is_enum"`

	Description    string `json:"description"`
	IsRefcounted   bool   `json:"is_refcounted"`
	IsInstantiable bool   `json:"is_instantiable"`
	IsSingleton    bool
	IsEphemeral    bool

	IsKeyed   bool `json:"is_keyed"` // builtin class
	Operators []struct {
		Name       string `json:"name"`
		RightType  string `json:"right_type,omitempty"`
		ReturnType string `json:"return_type"`
	} `json:"operators"` // builtin class
	Constructors []struct {
		Index     int `json:"index"`
		Arguments []struct {
			Name string `json:"name"`
			Type string `json:"type"`
			Meta string `json:"meta"`
		} `json:"arguments,omitempty"`
	} `json:"constructors"` // builtin class
	HasDestructor      bool   `json:"has_destructor"`                 // builtin class
	IndexingReturnType string `json:"indexing_return_type,omitempty"` // builtin class

	Inherits string   `json:"inherits,omitempty"`
	APIType  string   `json:"api_type"`
	Enums    []Enum   `json:"enums,omitempty"`
	Methods  []Method `json:"methods,omitempty"`

	Members []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"members,omitempty"` // builtin class

	Signals    []Signal `json:"signals,omitempty"`
	Properties []struct {
		Type        string `json:"type"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Setter      string `json:"setter"`
		Getter      string `json:"getter"`
		Index       *int   `json:"index"`
	} `json:"properties,omitempty"`
	Constants []struct {
		Name  string `json:"name"`
		Type  string `json:"type"` // builtin class
		Value any    `json:"value"`
	} `json:"constants,omitempty"`
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
	BuiltinClasses []Class `json:"builtin_classes"`
	Classes        []Class `json:"classes"`
	Singletons     []struct {
		Name string `json:"name"`
		Type string `json:"type"`
	} `json:"singletons"`
	NativeStructures []struct {
		Name   string `json:"name"`
		Format string `json:"format"`
	} `json:"native_structures"`
}

// LoadSpecification, either from a local file or by downloading
// it from the Godot Github repository.
func LoadSpecification() (*Specification, error) {
	file, err := os.Open("./extension_api.json")
	if os.IsNotExist(err) {
		req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/godotengine/godot-headers/master/extension_api.json", nil)
		if err != nil {
			return nil, err
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		file, err = os.Create("./extension_api.json")
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return nil, err
		}
		file.Seek(0, 0)
		resp.Body.Close()
	}
	var spec Specification
	if err := json.NewDecoder(file).Decode(&spec); err != nil {
		return nil, err
	}
	var singletons = make(map[string]bool)
	for _, class := range spec.Singletons {
		singletons[class.Name] = true
	}
	for i := range spec.Classes {
		class := &spec.Classes[i]
		class.IsSingleton = singletons[class.Name]
	}
	return &spec, nil
}

func ConvertName(fnName string) string {
	if fnName == "seek" {
		return "SeekTo"
	}
	if fnName == "type_string" {
		return "TypeToString"
	}

	fnName = strings.ToLower(fnName)

	joins := []string{}
	split := strings.Split(fnName, "_")
	for _, word := range split {
		joins = append(joins, strings.Title(word))
	}
	/*if joins[0] == "Get" {
		backup := joins
		joins = joins[1:]

		if len(joins) == 0 {
			joins = backup
		} else {
			if _, err := strconv.Atoi(joins[0]); err == nil {
				joins = backup
			}
		}
	}*/
	return strings.Join(joins, "")
}
