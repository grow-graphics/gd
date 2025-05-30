// Code generated by the generate package DO NOT EDIT

// Package Noise provides methods for working with Noise object instances.
package Noise

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
import "graphics.gd/classdb/Image"
import "graphics.gd/classdb/Resource"
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
import "graphics.gd/variant/Vector2"
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
This class defines the interface for noise generation libraries to inherit from.
A default [method get_seamless_image] implementation is provided for libraries that do not provide seamless noise. This function requests a larger image from the [method get_image] method, reverses the quadrants of the image, then uses the strips of extra width to blend over the seams.
Inheriting noise classes can optionally override this function to provide a more optimal algorithm.
*/
type Instance [1]gdclass.Noise

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

type Expanded [1]gdclass.Noise

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsNoise() Instance
}

/*
Returns the 1D noise value at the given (x) coordinate.
*/
func (self Instance) GetNoise1d(x Float.X) Float.X { //gd:Noise.get_noise_1d
	return Float.X(Float.X(Advanced(self).GetNoise1d(float64(x))))
}

/*
Returns the 2D noise value at the given position.
*/
func (self Instance) GetNoise2d(x Float.X, y Float.X) Float.X { //gd:Noise.get_noise_2d
	return Float.X(Float.X(Advanced(self).GetNoise2d(float64(x), float64(y))))
}

/*
Returns the 2D noise value at the given position.
*/
func (self Instance) GetNoise2dv(v Vector2.XY) Float.X { //gd:Noise.get_noise_2dv
	return Float.X(Float.X(Advanced(self).GetNoise2dv(Vector2.XY(v))))
}

/*
Returns the 3D noise value at the given position.
*/
func (self Instance) GetNoise3d(x Float.X, y Float.X, z Float.X) Float.X { //gd:Noise.get_noise_3d
	return Float.X(Float.X(Advanced(self).GetNoise3d(float64(x), float64(y), float64(z))))
}

/*
Returns the 3D noise value at the given position.
*/
func (self Instance) GetNoise3dv(v Vector3.XYZ) Float.X { //gd:Noise.get_noise_3dv
	return Float.X(Float.X(Advanced(self).GetNoise3dv(Vector3.XYZ(v))))
}

/*
Returns an [Image] containing 2D noise values.
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
func (self Instance) GetImage(width int, height int) Image.Instance { //gd:Noise.get_image
	return Image.Instance(Advanced(self).GetImage(int64(width), int64(height), false, false, true))
}

/*
Returns an [Image] containing 2D noise values.
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
func (self Expanded) GetImage(width int, height int, invert bool, in_3d_space bool, normalize bool) Image.Instance { //gd:Noise.get_image
	return Image.Instance(Advanced(self).GetImage(int64(width), int64(height), invert, in_3d_space, normalize))
}

/*
Returns an [Image] containing seamless 2D noise values.
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
func (self Instance) GetSeamlessImage(width int, height int) Image.Instance { //gd:Noise.get_seamless_image
	return Image.Instance(Advanced(self).GetSeamlessImage(int64(width), int64(height), false, false, float64(0.1), true))
}

/*
Returns an [Image] containing seamless 2D noise values.
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
func (self Expanded) GetSeamlessImage(width int, height int, invert bool, in_3d_space bool, skirt Float.X, normalize bool) Image.Instance { //gd:Noise.get_seamless_image
	return Image.Instance(Advanced(self).GetSeamlessImage(int64(width), int64(height), invert, in_3d_space, float64(skirt), normalize))
}

/*
Returns an [Array] of [Image]s containing 3D noise values for use with [method ImageTexture3D.create].
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
func (self Instance) GetImage3d(width int, height int, depth int) []Image.Instance { //gd:Noise.get_image_3d
	return []Image.Instance(gd.ArrayAs[[]Image.Instance](gd.InternalArray(Advanced(self).GetImage3d(int64(width), int64(height), int64(depth), false, true))))
}

/*
Returns an [Array] of [Image]s containing 3D noise values for use with [method ImageTexture3D.create].
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
func (self Expanded) GetImage3d(width int, height int, depth int, invert bool, normalize bool) []Image.Instance { //gd:Noise.get_image_3d
	return []Image.Instance(gd.ArrayAs[[]Image.Instance](gd.InternalArray(Advanced(self).GetImage3d(int64(width), int64(height), int64(depth), invert, normalize))))
}

/*
Returns an [Array] of [Image]s containing seamless 3D noise values for use with [method ImageTexture3D.create].
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
func (self Instance) GetSeamlessImage3d(width int, height int, depth int) []Image.Instance { //gd:Noise.get_seamless_image_3d
	return []Image.Instance(gd.ArrayAs[[]Image.Instance](gd.InternalArray(Advanced(self).GetSeamlessImage3d(int64(width), int64(height), int64(depth), false, float64(0.1), true))))
}

/*
Returns an [Array] of [Image]s containing seamless 3D noise values for use with [method ImageTexture3D.create].
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
func (self Expanded) GetSeamlessImage3d(width int, height int, depth int, invert bool, skirt Float.X, normalize bool) []Image.Instance { //gd:Noise.get_seamless_image_3d
	return []Image.Instance(gd.ArrayAs[[]Image.Instance](gd.InternalArray(Advanced(self).GetSeamlessImage3d(int64(width), int64(height), int64(depth), invert, float64(skirt), normalize))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.Noise

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("Noise"))
	casted := Instance{*(*gdclass.Noise)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

/*
Returns the 1D noise value at the given (x) coordinate.
*/
//go:nosplit
func (self class) GetNoise1d(x float64) float64 { //gd:Noise.get_noise_1d
	var frame = callframe.New()
	callframe.Arg(frame, x)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Noise.Bind_get_noise_1d, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the 2D noise value at the given position.
*/
//go:nosplit
func (self class) GetNoise2d(x float64, y float64) float64 { //gd:Noise.get_noise_2d
	var frame = callframe.New()
	callframe.Arg(frame, x)
	callframe.Arg(frame, y)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Noise.Bind_get_noise_2d, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the 2D noise value at the given position.
*/
//go:nosplit
func (self class) GetNoise2dv(v Vector2.XY) float64 { //gd:Noise.get_noise_2dv
	var frame = callframe.New()
	callframe.Arg(frame, v)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Noise.Bind_get_noise_2dv, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the 3D noise value at the given position.
*/
//go:nosplit
func (self class) GetNoise3d(x float64, y float64, z float64) float64 { //gd:Noise.get_noise_3d
	var frame = callframe.New()
	callframe.Arg(frame, x)
	callframe.Arg(frame, y)
	callframe.Arg(frame, z)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Noise.Bind_get_noise_3d, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the 3D noise value at the given position.
*/
//go:nosplit
func (self class) GetNoise3dv(v Vector3.XYZ) float64 { //gd:Noise.get_noise_3dv
	var frame = callframe.New()
	callframe.Arg(frame, v)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Noise.Bind_get_noise_3dv, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns an [Image] containing 2D noise values.
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
//go:nosplit
func (self class) GetImage(width int64, height int64, invert bool, in_3d_space bool, normalize bool) [1]gdclass.Image { //gd:Noise.get_image
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, invert)
	callframe.Arg(frame, in_3d_space)
	callframe.Arg(frame, normalize)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Noise.Bind_get_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns an [Image] containing seamless 2D noise values.
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
//go:nosplit
func (self class) GetSeamlessImage(width int64, height int64, invert bool, in_3d_space bool, skirt float64, normalize bool) [1]gdclass.Image { //gd:Noise.get_seamless_image
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, invert)
	callframe.Arg(frame, in_3d_space)
	callframe.Arg(frame, skirt)
	callframe.Arg(frame, normalize)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Noise.Bind_get_seamless_image, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Image{gd.PointerWithOwnershipTransferredToGo[gdclass.Image](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Returns an [Array] of [Image]s containing 3D noise values for use with [method ImageTexture3D.create].
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
//go:nosplit
func (self class) GetImage3d(width int64, height int64, depth int64, invert bool, normalize bool) Array.Contains[[1]gdclass.Image] { //gd:Noise.get_image_3d
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, depth)
	callframe.Arg(frame, invert)
	callframe.Arg(frame, normalize)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Noise.Bind_get_image_3d, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.Image]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns an [Array] of [Image]s containing seamless 3D noise values for use with [method ImageTexture3D.create].
[b]Note:[/b] With [param normalize] set to [code]false[/code], the default implementation expects the noise generator to return values in the range [code]-1.0[/code] to [code]1.0[/code].
*/
//go:nosplit
func (self class) GetSeamlessImage3d(width int64, height int64, depth int64, invert bool, skirt float64, normalize bool) Array.Contains[[1]gdclass.Image] { //gd:Noise.get_seamless_image_3d
	var frame = callframe.New()
	callframe.Arg(frame, width)
	callframe.Arg(frame, height)
	callframe.Arg(frame, depth)
	callframe.Arg(frame, invert)
	callframe.Arg(frame, skirt)
	callframe.Arg(frame, normalize)
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.Noise.Bind_get_seamless_image_3d, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Array.Through(gd.ArrayProxy[[1]gdclass.Image]{}, pointers.Pack(pointers.New[gd.Array](r_ret.Get())))
	frame.Free()
	return ret
}
func (self class) AsNoise() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNoise() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsNoise() Instance { return self.Super().AsNoise() }
func (self class) AsResource() Resource.Advanced {
	return *((*Resource.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsResource() Resource.Instance { return self.Super().AsResource() }
func (self Instance) AsResource() Resource.Instance {
	return *((*Resource.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsRefCounted() [1]gd.RefCounted { return self.Super().AsRefCounted() }
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Advanced(self.AsResource()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Resource.Instance(self.AsResource()), name)
	}
}
func init() {
	gdclass.Register("Noise", func(ptr gd.Object) any { return [1]gdclass.Noise{*(*gdclass.Noise)(unsafe.Pointer(&ptr))} })
}
