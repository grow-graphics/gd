// Code generated by the generate package DO NOT EDIT

// Package GPUParticlesCollisionSDF3D provides methods for working with GPUParticlesCollisionSDF3D object instances.
package GPUParticlesCollisionSDF3D

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/variant/Angle"
import "graphics.gd/variant/Euler"
import "graphics.gd/classdb/GPUParticlesCollision3D"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Node3D"
import "graphics.gd/classdb/Texture3D"
import "graphics.gd/classdb/VisualInstance3D"
import "graphics.gd/variant/Array"
import "graphics.gd/variant/Callable"
import "graphics.gd/variant/Dictionary"
import "graphics.gd/variant/Error"
import "graphics.gd/variant/Float"
import "graphics.gd/variant/Object"
import "graphics.gd/variant/Packed"
import "graphics.gd/variant/Path"
import "graphics.gd/variant/RID"
import "graphics.gd/variant/RefCounted"
import "graphics.gd/variant/String"
import "graphics.gd/variant/Vector3"

var _ Object.ID

type _ gdclass.Node

var _ gd.Object
var _ RefCounted.Instance
var _ unsafe.Pointer
var _ reflect.Type
var _ callframe.Frame
var _ = pointers.Cycle
var _ = Array.Nil
var _ variant.Any
var _ Callable.Function
var _ Dictionary.Any
var _ RID.Any
var _ String.Readable
var _ Path.ToNode
var _ Packed.Bytes
var _ Error.Code
var _ Float.X
var _ Angle.Radians
var _ Euler.Radians
var _ = slices.Delete[[]struct{}, struct{}]

/*
ID is a typed object ID (reference) to an instance of this class, use it to store references to objects with
unknown lifetimes, as an ID will not panic on use if the underlying object has been destroyed.
*/
type ID Object.ID

func (id ID) Instance() (Instance, bool) { return Object.As[Instance](Object.ID(id).Instance()) }

/*
Extension can be embedded in a new struct to create an extension of this class.
T should be the type that is embedding this [Extension]
*/
type Extension[T gdclass.Interface] struct{ gdclass.Extension[T, Instance] }

/*
A baked signed distance field 3D particle collision shape affecting [GPUParticles3D] nodes.
Signed distance fields (SDF) allow for efficiently representing approximate collision shapes for convex and concave objects of any shape. This is more flexible than [GPUParticlesCollisionHeightField3D], but it requires a baking step.
[b]Baking:[/b] The signed distance field texture can be baked by selecting the [GPUParticlesCollisionSDF3D] node in the editor, then clicking [b]Bake SDF[/b] at the top of the 3D viewport. Any [i]visible[/i] [MeshInstance3D]s within the [member size] will be taken into account for baking, regardless of their [member GeometryInstance3D.gi_mode].
[b]Note:[/b] Baking a [GPUParticlesCollisionSDF3D]'s [member texture] is only possible within the editor, as there is no bake method exposed for use in exported projects. However, it's still possible to load pre-baked [Texture3D]s into its [member texture] property in an exported project.
[b]Note:[/b] [member ParticleProcessMaterial.collision_mode] must be [constant ParticleProcessMaterial.COLLISION_RIGID] or [constant ParticleProcessMaterial.COLLISION_HIDE_ON_CONTACT] on the [GPUParticles3D]'s process material for collision to work.
[b]Note:[/b] Particle collision only affects [GPUParticles3D], not [CPUParticles3D].
*/
type Instance [1]gdclass.GPUParticlesCollisionSDF3D

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsGPUParticlesCollisionSDF3D() Instance
}

/*
Based on [param value], enables or disables the specified layer in the [member bake_mask], given a [param layer_number] between 1 and 32.
*/
func (self Instance) SetBakeMaskValue(layer_number int, value bool) { //gd:GPUParticlesCollisionSDF3D.set_bake_mask_value
	Advanced(self).SetBakeMaskValue(int64(layer_number), value)
}

/*
Returns whether or not the specified layer of the [member bake_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
func (self Instance) GetBakeMaskValue(layer_number int) bool { //gd:GPUParticlesCollisionSDF3D.get_bake_mask_value
	return bool(Advanced(self).GetBakeMaskValue(int64(layer_number)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.GPUParticlesCollisionSDF3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("GPUParticlesCollisionSDF3D"))
	casted := Instance{*(*gdclass.GPUParticlesCollisionSDF3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) Size() Vector3.XYZ {
	return Vector3.XYZ(class(self).GetSize())
}

func (self Instance) SetSize(value Vector3.XYZ) {
	class(self).SetSize(Vector3.XYZ(value))
}

func (self Instance) Resolution() Resolution {
	return Resolution(class(self).GetResolution())
}

func (self Instance) SetResolution(value Resolution) {
	class(self).SetResolution(value)
}

func (self Instance) Thickness() Float.X {
	return Float.X(Float.X(class(self).GetThickness()))
}

func (self Instance) SetThickness(value Float.X) {
	class(self).SetThickness(float64(value))
}

func (self Instance) BakeMask() int {
	return int(int(class(self).GetBakeMask()))
}

func (self Instance) SetBakeMask(value int) {
	class(self).SetBakeMask(int64(value))
}

func (self Instance) Texture() Texture3D.Instance {
	return Texture3D.Instance(class(self).GetTexture())
}

func (self Instance) SetTexture(value Texture3D.Instance) {
	class(self).SetTexture(value)
}

//go:nosplit
func (self class) SetSize(size Vector3.XYZ) { //gd:GPUParticlesCollisionSDF3D.set_size
	var frame = callframe.New()
	callframe.Arg(frame, size)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_set_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSize() Vector3.XYZ { //gd:GPUParticlesCollisionSDF3D.get_size
	var frame = callframe.New()
	var r_ret = callframe.Ret[Vector3.XYZ](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_get_size, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetResolution(resolution Resolution) { //gd:GPUParticlesCollisionSDF3D.set_resolution
	var frame = callframe.New()
	callframe.Arg(frame, resolution)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_set_resolution, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetResolution() Resolution { //gd:GPUParticlesCollisionSDF3D.get_resolution
	var frame = callframe.New()
	var r_ret = callframe.Ret[Resolution](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_get_resolution, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetTexture(texture [1]gdclass.Texture3D) { //gd:GPUParticlesCollisionSDF3D.set_texture
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_set_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetTexture() [1]gdclass.Texture3D { //gd:GPUParticlesCollisionSDF3D.get_texture
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_get_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture3D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture3D](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetThickness(thickness float64) { //gd:GPUParticlesCollisionSDF3D.set_thickness
	var frame = callframe.New()
	callframe.Arg(frame, thickness)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_set_thickness, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetThickness() float64 { //gd:GPUParticlesCollisionSDF3D.get_thickness
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_get_thickness, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBakeMask(mask int64) { //gd:GPUParticlesCollisionSDF3D.set_bake_mask
	var frame = callframe.New()
	callframe.Arg(frame, mask)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_set_bake_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBakeMask() int64 { //gd:GPUParticlesCollisionSDF3D.get_bake_mask
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_get_bake_mask, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Based on [param value], enables or disables the specified layer in the [member bake_mask], given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) SetBakeMaskValue(layer_number int64, value bool) { //gd:GPUParticlesCollisionSDF3D.set_bake_mask_value
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	callframe.Arg(frame, value)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_set_bake_mask_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns whether or not the specified layer of the [member bake_mask] is enabled, given a [param layer_number] between 1 and 32.
*/
//go:nosplit
func (self class) GetBakeMaskValue(layer_number int64) bool { //gd:GPUParticlesCollisionSDF3D.get_bake_mask_value
	var frame = callframe.New()
	callframe.Arg(frame, layer_number)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.GPUParticlesCollisionSDF3D.Bind_get_bake_mask_value, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsGPUParticlesCollisionSDF3D() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsGPUParticlesCollisionSDF3D() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsGPUParticlesCollisionSDF3D() Instance {
	return self.Super().AsGPUParticlesCollisionSDF3D()
}
func (self class) AsGPUParticlesCollision3D() GPUParticlesCollision3D.Advanced {
	return *((*GPUParticlesCollision3D.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsGPUParticlesCollision3D() GPUParticlesCollision3D.Instance {
	return self.Super().AsGPUParticlesCollision3D()
}
func (self Instance) AsGPUParticlesCollision3D() GPUParticlesCollision3D.Instance {
	return *((*GPUParticlesCollision3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsVisualInstance3D() VisualInstance3D.Advanced {
	return *((*VisualInstance3D.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsVisualInstance3D() VisualInstance3D.Instance {
	return self.Super().AsVisualInstance3D()
}
func (self Instance) AsVisualInstance3D() VisualInstance3D.Instance {
	return *((*VisualInstance3D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode3D() Node3D.Advanced         { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsNode3D() Node3D.Instance { return self.Super().AsNode3D() }
func (self Instance) AsNode3D() Node3D.Instance      { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced             { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsNode() Node.Instance     { return self.Super().AsNode() }
func (self Instance) AsNode() Node.Instance          { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(GPUParticlesCollision3D.Advanced(self.AsGPUParticlesCollision3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(GPUParticlesCollision3D.Instance(self.AsGPUParticlesCollision3D()), name)
	}
}
func init() {
	gdclass.Register("GPUParticlesCollisionSDF3D", func(ptr gd.Object) any {
		return [1]gdclass.GPUParticlesCollisionSDF3D{*(*gdclass.GPUParticlesCollisionSDF3D)(unsafe.Pointer(&ptr))}
	})
}

type Resolution int //gd:GPUParticlesCollisionSDF3D.Resolution

const (
	/*Bake a 16×16×16 signed distance field. This is the fastest option, but also the least precise.*/
	Resolution16 Resolution = 0
	/*Bake a 32×32×32 signed distance field.*/
	Resolution32 Resolution = 1
	/*Bake a 64×64×64 signed distance field.*/
	Resolution64 Resolution = 2
	/*Bake a 128×128×128 signed distance field.*/
	Resolution128 Resolution = 3
	/*Bake a 256×256×256 signed distance field.*/
	Resolution256 Resolution = 4
	/*Bake a 512×512×512 signed distance field. This is the slowest option, but also the most precise.*/
	Resolution512 Resolution = 5
	/*Represents the size of the [enum Resolution] enum.*/
	ResolutionMax Resolution = 6
)
