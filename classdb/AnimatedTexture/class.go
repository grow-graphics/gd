// Code generated by the generate package DO NOT EDIT

// Package AnimatedTexture provides methods for working with AnimatedTexture object instances.
package AnimatedTexture

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
import "graphics.gd/classdb/Resource"
import "graphics.gd/classdb/Texture"
import "graphics.gd/classdb/Texture2D"
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
[AnimatedTexture] is a resource format for frame-based animations, where multiple textures can be chained automatically with a predefined delay for each frame. Unlike [AnimationPlayer] or [AnimatedSprite2D], it isn't a [Node], but has the advantage of being usable anywhere a [Texture2D] resource can be used, e.g. in a [TileSet].
The playback of the animation is controlled by the [member speed_scale] property, as well as each frame's duration (see [method set_frame_duration]). The animation loops, i.e. it will restart at frame 0 automatically after playing the last frame.
[AnimatedTexture] currently requires all frame textures to have the same size, otherwise the bigger ones will be cropped to match the smallest one.
[b]Note:[/b] AnimatedTexture doesn't support using [AtlasTexture]s. Each frame needs to be a separate [Texture2D].
[b]Warning:[/b] The current implementation is not efficient for the modern renderers.
*/
type Instance [1]gdclass.AnimatedTexture

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsAnimatedTexture() Instance
}

/*
Assigns a [Texture2D] to the given frame. Frame IDs start at 0, so the first frame has ID 0, and the last frame of the animation has ID [member frames] - 1.
You can define any number of textures up to [constant MAX_FRAMES], but keep in mind that only frames from 0 to [member frames] - 1 will be part of the animation.
*/
func (self Instance) SetFrameTexture(frame_ int, texture Texture2D.Instance) { //gd:AnimatedTexture.set_frame_texture
	Advanced(self).SetFrameTexture(int64(frame_), texture)
}

/*
Returns the given frame's [Texture2D].
*/
func (self Instance) GetFrameTexture(frame_ int) Texture2D.Instance { //gd:AnimatedTexture.get_frame_texture
	return Texture2D.Instance(Advanced(self).GetFrameTexture(int64(frame_)))
}

/*
Sets the duration of any given [param frame]. The final duration is affected by the [member speed_scale]. If set to [code]0[/code], the frame is skipped during playback.
*/
func (self Instance) SetFrameDuration(frame_ int, duration Float.X) { //gd:AnimatedTexture.set_frame_duration
	Advanced(self).SetFrameDuration(int64(frame_), float64(duration))
}

/*
Returns the given [param frame]'s duration, in seconds.
*/
func (self Instance) GetFrameDuration(frame_ int) Float.X { //gd:AnimatedTexture.get_frame_duration
	return Float.X(Float.X(Advanced(self).GetFrameDuration(int64(frame_))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.AnimatedTexture

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("AnimatedTexture"))
	casted := Instance{*(*gdclass.AnimatedTexture)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (self Instance) Frames() int {
	return int(int(class(self).GetFrames()))
}

func (self Instance) SetFrames(value int) {
	class(self).SetFrames(int64(value))
}

func (self Instance) CurrentFrame() int {
	return int(int(class(self).GetCurrentFrame()))
}

func (self Instance) SetCurrentFrame(value int) {
	class(self).SetCurrentFrame(int64(value))
}

func (self Instance) Pause() bool {
	return bool(class(self).GetPause())
}

func (self Instance) SetPause(value bool) {
	class(self).SetPause(value)
}

func (self Instance) OneShot() bool {
	return bool(class(self).GetOneShot())
}

func (self Instance) SetOneShot(value bool) {
	class(self).SetOneShot(value)
}

func (self Instance) SpeedScale() Float.X {
	return Float.X(Float.X(class(self).GetSpeedScale()))
}

func (self Instance) SetSpeedScale(value Float.X) {
	class(self).SetSpeedScale(float64(value))
}

//go:nosplit
func (self class) SetFrames(frames int64) { //gd:AnimatedTexture.set_frames
	var frame = callframe.New()
	callframe.Arg(frame, frames)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedTexture.Bind_set_frames, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetFrames() int64 { //gd:AnimatedTexture.get_frames
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedTexture.Bind_get_frames, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetCurrentFrame(frame_ int64) { //gd:AnimatedTexture.set_current_frame
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedTexture.Bind_set_current_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetCurrentFrame() int64 { //gd:AnimatedTexture.get_current_frame
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedTexture.Bind_get_current_frame, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPause(pause bool) { //gd:AnimatedTexture.set_pause
	var frame = callframe.New()
	callframe.Arg(frame, pause)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedTexture.Bind_set_pause, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetPause() bool { //gd:AnimatedTexture.get_pause
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedTexture.Bind_get_pause, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetOneShot(one_shot bool) { //gd:AnimatedTexture.set_one_shot
	var frame = callframe.New()
	callframe.Arg(frame, one_shot)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedTexture.Bind_set_one_shot, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetOneShot() bool { //gd:AnimatedTexture.get_one_shot
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedTexture.Bind_get_one_shot, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetSpeedScale(scale float64) { //gd:AnimatedTexture.set_speed_scale
	var frame = callframe.New()
	callframe.Arg(frame, scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedTexture.Bind_set_speed_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetSpeedScale() float64 { //gd:AnimatedTexture.get_speed_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedTexture.Bind_get_speed_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Assigns a [Texture2D] to the given frame. Frame IDs start at 0, so the first frame has ID 0, and the last frame of the animation has ID [member frames] - 1.
You can define any number of textures up to [constant MAX_FRAMES], but keep in mind that only frames from 0 to [member frames] - 1 will be part of the animation.
*/
//go:nosplit
func (self class) SetFrameTexture(frame_ int64, texture [1]gdclass.Texture2D) { //gd:AnimatedTexture.set_frame_texture
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	callframe.Arg(frame, pointers.Get(texture[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedTexture.Bind_set_frame_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the given frame's [Texture2D].
*/
//go:nosplit
func (self class) GetFrameTexture(frame_ int64) [1]gdclass.Texture2D { //gd:AnimatedTexture.get_frame_texture
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedTexture.Bind_get_frame_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Sets the duration of any given [param frame]. The final duration is affected by the [member speed_scale]. If set to [code]0[/code], the frame is skipped during playback.
*/
//go:nosplit
func (self class) SetFrameDuration(frame_ int64, duration float64) { //gd:AnimatedTexture.set_frame_duration
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	callframe.Arg(frame, duration)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedTexture.Bind_set_frame_duration, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns the given [param frame]'s duration, in seconds.
*/
//go:nosplit
func (self class) GetFrameDuration(frame_ int64) float64 { //gd:AnimatedTexture.get_frame_duration
	var frame = callframe.New()
	callframe.Arg(frame, frame_)
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.AnimatedTexture.Bind_get_frame_duration, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsAnimatedTexture() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsAnimatedTexture() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsAnimatedTexture() Instance { return self.Super().AsAnimatedTexture() }
func (self class) AsTexture2D() Texture2D.Advanced {
	return *((*Texture2D.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsTexture2D() Texture2D.Instance { return self.Super().AsTexture2D() }
func (self Instance) AsTexture2D() Texture2D.Instance {
	return *((*Texture2D.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsTexture() Texture.Advanced         { return *((*Texture.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsTexture() Texture.Instance { return self.Super().AsTexture() }
func (self Instance) AsTexture() Texture.Instance {
	return *((*Texture.Instance)(unsafe.Pointer(&self)))
}
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
		return gd.VirtualByName(Texture2D.Advanced(self.AsTexture2D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Texture2D.Instance(self.AsTexture2D()), name)
	}
}
func init() {
	gdclass.Register("AnimatedTexture", func(ptr gd.Object) any {
		return [1]gdclass.AnimatedTexture{*(*gdclass.AnimatedTexture)(unsafe.Pointer(&ptr))}
	})
}

const MaxFrames int = 256 //gd:AnimatedTexture.MAX_FRAMES
