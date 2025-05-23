package gd_test

import (
	"errors"
	"testing"
	"time"

	"graphics.gd/classdb/GDScript"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/SceneTree"
	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant"
	"graphics.gd/variant/AABB"
	"graphics.gd/variant/Array"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Dictionary"
	"graphics.gd/variant/Enum"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Path"
	"graphics.gd/variant/Plane"
	"graphics.gd/variant/Projection"
	"graphics.gd/variant/Quaternion"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Rect2i"
	"graphics.gd/variant/Signal"
	"graphics.gd/variant/String"
	"graphics.gd/variant/StringName"
	"graphics.gd/variant/Transform2D"
	"graphics.gd/variant/Transform3D"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector3i"
	"graphics.gd/variant/Vector4"
	"graphics.gd/variant/Vector4i"

	_ "embed"
)

//go:embed convert_types_test.gd
var convert_types_test string

var doneConversionsTest = make(chan error, 100)

type Converter struct {
	Node.Extension[Converter]
}

type CustomConverterObject struct {
	Object.Extension[CustomConverterObject]

	Value int
}

type MyEnum Enum.Int[struct {
	A,
	B MyEnum
}]

var MyEnums = Enum.Values[MyEnum]()

func (c Converter) Ready()                { doneConversionsTest <- errors.New("script didn't call Done") }
func (c Converter) AsNode() Node.Instance { return c.Super() }

func (c Converter) Done() { doneConversionsTest <- nil }
func (c Converter) Fail() { doneConversionsTest <- errors.New("test failed") }

func (c Converter) HelloWorld() string     { return "Hello, World!" }
func (c Converter) Echo(s string) string   { return s }
func (c Converter) Int() int               { return 22 }
func (c Converter) Bool() bool             { return true }
func (c Converter) Int8() int8             { return 8 }
func (c Converter) Int16() int16           { return 16 }
func (c Converter) Int32() int32           { return 32 }
func (c Converter) Int64() int64           { return 64 }
func (c Converter) Uint() uint             { return 22 }
func (c Converter) Uint8() uint8           { return 8 }
func (c Converter) Uint16() uint16         { return 16 }
func (c Converter) Uint32() uint32         { return 32 }
func (c Converter) Uint64() uint64         { return 64 }
func (c Converter) Float32() float32       { return 2.0 }
func (c Converter) Float64() float64       { return 2.2 }
func (c Converter) Complex64() complex64   { return 2 + 2i }
func (c Converter) Complex128() complex128 { return 2 + 2i }
func (c Converter) Enum() MyEnum           { return MyEnums.B }

func (c Converter) Float() gd.Float                      { return 2.2 }
func (c Converter) String() String.Readable              { return String.New("testing") }
func (c Converter) Vector2() Vector2.XY                  { return Vector2.New(1, 2) }
func (c Converter) Vector2i() Vector2i.XY                { return Vector2i.New(1, 2) }
func (c Converter) Rect2() Rect2.PositionSize            { return Rect2.New(1, 2, 3, 4) }
func (c Converter) Rect2i() Rect2i.PositionSize          { return Rect2i.New(1, 2, 3, 4) }
func (c Converter) Vector3() Vector3.XYZ                 { return Vector3.New(1, 2, 3) }
func (c Converter) Vector3i() Vector3i.XYZ               { return Vector3i.New(1, 2, 3) }
func (c Converter) GetTransform2D() Transform2D.OriginXY { return Transform2D.Identity }
func (c Converter) Vector4() Vector4.XYZW                { return Vector4.New(1, 2, 3, 4) }
func (c Converter) Vector4i() Vector4i.XYZW              { return Vector4i.New(1, 2, 3, 4) }
func (c Converter) Plane() Plane.NormalD                 { return Plane.New(1, 2, 3, 4) }
func (c Converter) GetQuaternion() Quaternion.IJKX       { return Quaternion.New() }
func (c Converter) AABB() AABB.PositionSize {
	return AABB.PositionSize{Position: Vector3.New(1, 2, 3), Size: Vector3.New(4, 5, 6)}
}
func (c Converter) GetBasis() Basis.XYZ                     { return Basis.New() }
func (c Converter) GetTransform3D() Transform3D.BasisOrigin { return Transform3D.Identity }
func (c Converter) GetProjection() Projection.XYZW          { return Projection.New() }
func (c Converter) Color() Color.RGBA                       { return Color.Bytes(255, 0, 0, 255) }
func (c Converter) StringName() String.Name                 { return StringName.New("testing") }
func (c Converter) NodePath() Path.ToNode                   { return Path.ToNode(String.New("/")) }
func (c Converter) RID() Resource.ID                        { return 22 }
func (c Converter) Object() Object.Instance                 { return Object.New() }
func (c Converter) Callable() Callable.Function {
	return Callable.New(func() string {
		return "Hello, World!"
	})
}
func (c Converter) Signal() Signal.Any {
	sig := gd.NewSignalOf(c.AsObject(), gd.NewStringName("property_list_changed"))
	return Signal.Via(gd.SignalProxy{}, pointers.Pack(sig))
}
func (c Converter) Dictionary() Dictionary.Any {
	var dict = Dictionary.New[string, string]()
	dict.SetIndex("hello", "world")
	return dict.Any()
}
func (c Converter) Array() []int                  { return []int{1, 2, 3} }
func (c Converter) PackedByteArray() []byte       { return []byte{1, 2, 3} }
func (c Converter) PackedInt32Array() []int32     { return []int32{1, 2, 3} }
func (c Converter) PackedInt64Array() []int64     { return []int64{1, 2, 3} }
func (c Converter) PackedFloat32Array() []float32 { return []float32{1, 2, 3} }
func (c Converter) PackedFloat64Array() []float64 { return []float64{1, 2, 3} }
func (c Converter) PackedStringArray() []string   { return []string{"hello", "world"} }
func (c Converter) PackedVector2Array() []Vector2.XY {
	return []Vector2.XY{{1, 2}, {3, 4}}
}
func (c Converter) PackedVector3Array() []Vector3.XYZ {
	return []Vector3.XYZ{{1, 2, 3}, {4, 5, 6}}
}
func (c Converter) PackedColorArray() []Color.RGBA {
	return []Color.RGBA{{1, 2, 3, 4}, {5, 6, 7, 8}}
}
func (c Converter) PackedVector4Array() []Vector4.XYZW {
	return []Vector4.XYZW{{1, 2, 3, 4}, {5, 6, 7, 8}}
}

func (c Converter) ArrayAny() Array.Any {
	var array Array.Any
	array.Append(variant.New(1))
	array.Append(variant.New(2))
	array.Append(variant.New(3))
	return array
}
func (c Converter) ArrayInt() Array.Contains[int] {
	return Array.New(1, 2, 3)
}

func (c Converter) CustomObject() *CustomConverterObject {
	return &CustomConverterObject{Value: 42}
}

func (c Converter) ValidInt(i int) bool               { return i == c.Int() }
func (c Converter) ValidBool(b bool) bool             { return b == c.Bool() }
func (c Converter) ValidInt8(i int8) bool             { return i == c.Int8() }
func (c Converter) ValidInt16(i int16) bool           { return i == c.Int16() }
func (c Converter) ValidInt32(i int32) bool           { return i == c.Int32() }
func (c Converter) ValidInt64(i int64) bool           { return i == c.Int64() }
func (c Converter) ValidUint(i uint) bool             { return i == c.Uint() }
func (c Converter) ValidUint8(i uint8) bool           { return i == c.Uint8() }
func (c Converter) ValidUint16(i uint16) bool         { return i == c.Uint16() }
func (c Converter) ValidUint32(i uint32) bool         { return i == c.Uint32() }
func (c Converter) ValidUint64(i uint64) bool         { return i == c.Uint64() }
func (c Converter) ValidFloat32(f float32) bool       { return f == c.Float32() }
func (c Converter) ValidFloat64(f float64) bool       { return f == c.Float64() }
func (c Converter) ValidComplex64(f complex64) bool   { return f == c.Complex64() }
func (c Converter) ValidComplex128(f complex128) bool { return f == c.Complex128() }

func (c Converter) ValidFloat(f gd.Float) bool                      { return f == c.Float() }
func (c Converter) ValidString(s String.Readable) bool              { return s.String() == c.String().String() }
func (c Converter) ValidVector2(v Vector2.XY) bool                  { return v == c.Vector2() }
func (c Converter) ValidVector2i(v Vector2i.XY) bool                { return v == c.Vector2i() }
func (c Converter) ValidRect2(r Rect2.PositionSize) bool            { return r == c.Rect2() }
func (c Converter) ValidRect2i(r Rect2i.PositionSize) bool          { return r == c.Rect2i() }
func (c Converter) ValidVector3(v Vector3.XYZ) bool                 { return v == c.Vector3() }
func (c Converter) ValidVector3i(v Vector3i.XYZ) bool               { return v == c.Vector3i() }
func (c Converter) ValidTransform2D(t Transform2D.OriginXY) bool    { return t == c.GetTransform2D() }
func (c Converter) ValidVector4(v Vector4.XYZW) bool                { return v == c.Vector4() }
func (c Converter) ValidVector4i(v Vector4i.XYZW) bool              { return v == c.Vector4i() }
func (c Converter) ValidPlane(p Plane.NormalD) bool                 { return p == c.Plane() }
func (c Converter) ValidQuaternion(q Quaternion.IJKX) bool          { return q == c.GetQuaternion() }
func (c Converter) ValidAABB(a AABB.PositionSize) bool              { return a == c.AABB() }
func (c Converter) ValidBasis(b Basis.XYZ) bool                     { return b == c.GetBasis() }
func (c Converter) ValidTransform3D(t Transform3D.BasisOrigin) bool { return t == c.GetTransform3D() }
func (c Converter) ValidProjection(p Projection.XYZW) bool          { return p == c.GetProjection() }
func (c Converter) ValidColor(cc Color.RGBA) bool                   { return cc == c.Color() }
func (c Converter) ValidStringName(s String.Name) bool {
	return s.String() == c.StringName().String()
}
func (c Converter) ValidNodePath(n Path.ToNode) bool { return n.String() == c.NodePath().String() }
func (c Converter) ValidRID(r Resource.ID) bool      { return r == c.RID() }
func (c Converter) ValidObject(o Object.Instance) bool {
	return o.AsObject()[0].GetClass().String() == "Object"
}
func (c Converter) ValidCallable(cc Callable.Function) bool {
	return cc.Call().String() == c.Callable().Call().String()
}
func (c Converter) ValidSignal(s Signal.Any) bool {
	return s.Name().String() == c.Signal().Name().String()
}
func (c Converter) ValidDictionary(d Dictionary.Any) bool {
	return d.Index(variant.New("hello")).String() == "world"
}
func (c Converter) ValidArray(a []int) bool {
	return len(a) == 3 && a[0] == 1 && a[1] == 2 && a[2] == 3
}
func (c Converter) ValidPackedByteArray(a []byte) bool {
	return len(a) == 3 && a[0] == 1 && a[1] == 2 && a[2] == 3
}
func (c Converter) ValidPackedInt32Array(a []int32) bool {
	return len(a) == 3 && a[0] == 1 && a[1] == 2 && a[2] == 3
}
func (c Converter) ValidPackedInt64Array(a []int64) bool {
	return len(a) == 3 && a[0] == 1 && a[1] == 2 && a[2] == 3
}
func (c Converter) ValidPackedFloat32Array(a []float32) bool {
	return len(a) == 3 && a[0] == 1 && a[1] == 2 && a[2] == 3
}
func (c Converter) ValidPackedFloat64Array(a []float64) bool {
	return len(a) == 3 && a[0] == 1 && a[1] == 2 && a[2] == 3
}
func (c Converter) ValidPackedStringArray(a []string) bool {
	return len(a) == 2 && a[0] == "hello" && a[1] == "world"
}
func (c Converter) ValidPackedVector2Array(a []Vector2.XY) bool {
	return len(a) == 2 && a[0] == Vector2.New(1, 2) && a[1] == Vector2.New(3, 4)
}
func (c Converter) ValidPackedVector3Array(a []Vector3.XYZ) bool {
	return len(a) == 2 && a[0] == Vector3.New(1, 2, 3) && a[1] == Vector3.New(4, 5, 6)
}
func (c Converter) ValidPackedColorArray(a []Color.RGBA) bool {
	return len(a) == 2 && a[0] == Color.RGBA{1, 2, 3, 4} && a[1] == Color.RGBA{5, 6, 7, 8}
}
func (c Converter) ValidPackedVector4Array(a []Vector4.XYZW) bool {
	return len(a) == 2 && a[0] == Vector4.New(1, 2, 3, 4) && a[1] == Vector4.New(5, 6, 7, 8)
}
func (c Converter) ValidArrayAny(a Array.Any) bool {
	return a.Index(0).Int() == 1 && a.Index(1).Int() == 2 && a.Index(2).Int() == 3
}
func (c Converter) ValidArrayInt(a Array.Contains[int]) bool {
	return a.Index(0) == 1 && a.Index(1) == 2 && a.Index(2) == 3
}

func (c Converter) ValidEnum(e MyEnum) bool { return e == MyEnums.B }
func (c Converter) ValidCustomObject(a *CustomConverterObject) bool {
	return a.Value == 42
}

func TestConversions(t *testing.T) {
	converter := &Converter{}
	var script = GDScript.New().AsScript()
	script.SetSourceCode(convert_types_test)
	script.Reload()
	Object.Instance(converter.AsObject()).SetScript(script)
	SceneTree.Add(converter)
	select {
	case err := <-doneConversionsTest:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.NewTimer(500 * time.Millisecond).C:
		t.Fatal("timeout")
	}
}
