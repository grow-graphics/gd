package main

import (
	"reflect"
	"strings"

	"graphics.gd/internal/gdjson"
	"graphics.gd/internal/gdtype"
	"graphics.gd/variant/RID"
)

type ClassDB map[string]gdjson.Class

func (db ClassDB) nameOf(pkg, original string) string {
	class := db[original]
	if pkg == class.Package {
		return class.Name
	}
	if class.Package == "internal" {
		return "gd." + class.Name
	}
	return class.Package + "." + class.Name
}

func pascal(name string) string {
	words := strings.Split(name, "_")
	for i := range words {
		words[i] = strings.Title(words[i])
	}
	return strings.Join(words, "")
}

func fixReserved(name string) string {
	switch name {
	case "bool":
		return "b"
	case "type":
		return "atype"
	case "range":
		return "arange"
	case "default":
		return "def"
	case "class":
		return "class_"
	case "func":
		return "fn"
	case "frame":
		return "frame_"
	case "interface":
		return "intf"
	case "internal":
		return "internal_"
	case "map":
		return "mapping"
	case "var":
		return "v"
	case "object":
		return "obj"
	case "string":
		return "s"
	case "RID":
		return "rid"
	case "variant":
		return "v"
	default:
		return name
	}
}

var StructablesInThisPackageGlobalHack = make(map[reflect.Type]bool)

func (classDB ClassDB) convertTypeSimple(class gdjson.Class, lookup, meta string, gdType string) string {
	distinctions := gdjson.Distinctions[class.Name]
	for _, distinction := range distinctions {
		pattern := distinction[0]
		if gdType != distinction[1] {
			continue
		}
		matchFunction, matchArgument, hasFunctionMatch := strings.Cut(pattern, " ")
		if !hasFunctionMatch {
			matchArgument = pattern
		}
		_, thisFunc, _ := strings.Cut(lookup, ".")
		thisFunc, thisArg, _ := strings.Cut(thisFunc, ".")
		if hasFunctionMatch {
			if !strings.Contains(thisFunc, matchFunction) {
				continue
			}
		}
		if (matchArgument == "" && thisArg != "") || !strings.Contains(thisArg, matchArgument) {
			continue
		}
		return distinction[2]
	}
	if strings.HasPrefix(gdType, "typedarray::") {
		gdType = strings.TrimPrefix(gdType, "typedarray::")
		meta, rest, ok := strings.Cut(gdType, ":")
		if ok {
			gdType = rest
		}
		return "[]" + classDB.convertTypeSimple(class, lookup, meta, gdType)
	}
	switch gdType {
	case "int", "Int":
		return "int"
	case "float", "Float":
		return "Float.X"
	case "bool", "Bool":
		return "bool"
	case "String", "StringName":
		return "string"
	case "PackedByteArray":
		return "[]byte"
	case "PackedStringArray":
		return "[]string"
	case "PackedInt32Array":
		return "[]int32"
	case "PackedInt64Array":
		return "[]int64"
	case "PackedFloat32Array":
		return "[]float32"
	case "PackedFloat64Array":
		return "[]float64"
	case "PackedVector2Array":
		return "[]Vector2.XY"
	case "PackedVector3Array":
		return "[]Vector3.XYZ"
	case "PackedVector4Array":
		return "[]Vector4.XYZW"
	case "PackedColorArray":
		return "[]Color.RGBA"
	case "enum::Error":
		return "error"
	case "Vector2":
		return "Vector2.XY"
	case "Vector2i":
		return "Vector2i.XY"
	case "Rect2":
		return "Rect2.PositionSize"
	case "Rect2i":
		return "Rect2i.PositionSize"
	case "Vector3":
		return "Vector3.XYZ"
	case "Vector3i":
		return "Vector3i.XYZ"
	case "Transform2D":
		return "Transform2D.OriginXY"
	case "Vector4":
		return "Vector4.XYZW"
	case "Vector4i":
		return "Vector4i.XYZW"
	case "Plane":
		return "Plane.NormalD"
	case "Quaternion":
		return "Quaternion.IJKX"
	case "AABB":
		return "AABB.PositionSize"
	case "Basis":
		return "Basis.XYZ"
	case "Transform3D":
		return "Transform3D.BasisOrigin"
	case "Projection":
		return "Projection.XYZW"
	case "Color":
		return "Color.RGBA"
	case "NodePath":
		return "string"
	case "RID", "gd.RID":
		if class.Name == "Resource" {
			return "ID"
		}
		if typed, ok := gdjson.Resources[lookup]; ok {
			if typed == reflect.TypeFor[RID.Any]() {
				return "RID.Any"
			}
			return typed.String()
		}
		return "RID.Any"
	case "ObjectID":
		return "Object.ID"
	case "Dictionary":
		rtype, ok := gdjson.Structables[lookup]
		if !ok {
			return "map[any]any"
		}
		registerStructables(rtype)
		if rtype.Name() == "" {
			return strings.Replace(rtype.String(), "gdjson.", "", -1)
		}
		return strings.Replace(rtype.Name(), "gdjson.", "", -1)
	case "Array":
		rtype, ok := gdjson.Structables[lookup]
		if !ok {
			return "[]any"
		}
		registerStructables(rtype)
		if rtype.Name() == "" {
			return strings.Replace(rtype.String(), "gdjson.", "", -1)
		}
		return strings.Replace(rtype.Name(), "gdjson.", "", -1)
	case "Variant":
		return "any"
	case "Object":
		return "Object.Instance"
	case "Callable":
		details, ok := gdjson.Callables[lookup]
		if !ok || len(details) == 0 || strings.HasSuffix(lookup, ".") {
			return "Callable.Function"
		}
		var ftype string = "func("
		for i, arg := range details[1:] {
			if i > 0 {
				ftype += ", "
			}
			atype, name, ok := strings.Cut(arg, " ")
			if ok {
				ftype += name + " " + classDB.convertTypeSimple(class, "", "", atype)
			} else {
				ftype += classDB.convertTypeSimple(class, "", "", arg)
			}
		}
		ftype += ")"
		if details[0] != "void" {
			ftype += " " + classDB.convertTypeSimple(class, "", "", details[0])
		}
		return ftype
	default:
		object, ok := classDB[gdType]
		if ok {
			if class.Name == object.Name {
				return "Instance"
			}
			return object.Name + ".Instance"
		}
		return gdtype.EngineTypeAsGoType(class.Name, meta, gdType)
	}
}

func mapOperator(name string) string {
	switch name {
	case "==":
		return "Equals"
	case "!=":
		return "NotEqual"
	case "<":
		return "Less"
	case "<=":
		return "LessEqual"
	case ">":
		return "Greater"
	case ">=":
		return "GreaterEqual"
	case "+":
		return "Add"
	case "-":
		return "Subtract"
	case "*":
		return "Multiply"
	case "/":
		return "Divide"
	case "unary-":
		return "Negate"
	case "%":
		return "Module"
	case "**":
		return "Power"
	case "<<":
		return "ShiftLeft"
	case ">>":
		return "ShiftRight"
	case "&":
		return "BitAnd"
	case "|":
		return "BitOr"
	case "^":
		return "BitXor"
	case "~":
		return "BitNegate"
	case "and":
		return "And"
	case "or":
		return "Or"
	case "xor":
		return "Xor"
	case "not":
		return "Not"
	case "in":
		return "In"
	default:
		panic("unknown operator: " + name)
	}
}

func convertName(fnName string) string {
	return gdjson.ConvertName(fnName)
}
