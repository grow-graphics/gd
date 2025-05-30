// Code generated by the generate package DO NOT EDIT

// Package VideoStreamPlayer provides methods for working with VideoStreamPlayer object instances.
package VideoStreamPlayer

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
import "graphics.gd/classdb/CanvasItem"
import "graphics.gd/classdb/Control"
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/Texture2D"
import "graphics.gd/classdb/VideoStream"
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
A control used for playback of [VideoStream] resources.
Supported video formats are [url=https://www.theora.org/]Ogg Theora[/url] ([code].ogv[/code], [VideoStreamTheora]) and any format exposed via a GDExtension plugin.
[b]Warning:[/b] On Web, video playback [i]will[/i] perform poorly due to missing architecture-specific assembly optimizations.
*/
type Instance [1]gdclass.VideoStreamPlayer

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsVideoStreamPlayer() Instance
}

/*
Starts the video playback from the beginning. If the video is paused, this will not unpause the video.
*/
func (self Instance) Play() { //gd:VideoStreamPlayer.play
	Advanced(self).Play()
}

/*
Stops the video playback and sets the stream position to 0.
[b]Note:[/b] Although the stream position will be set to 0, the first frame of the video stream won't become the current frame.
*/
func (self Instance) Stop() { //gd:VideoStreamPlayer.stop
	Advanced(self).Stop()
}

/*
Returns [code]true[/code] if the video is playing.
[b]Note:[/b] The video is still considered playing if paused during playback.
*/
func (self Instance) IsPlaying() bool { //gd:VideoStreamPlayer.is_playing
	return bool(Advanced(self).IsPlaying())
}

/*
Returns the video stream's name, or [code]"<No Stream>"[/code] if no video stream is assigned.
*/
func (self Instance) GetStreamName() string { //gd:VideoStreamPlayer.get_stream_name
	return string(Advanced(self).GetStreamName().String())
}

/*
The length of the current stream, in seconds.
[b]Note:[/b] For [VideoStreamTheora] streams (the built-in format supported by Godot), this value will always be zero, as getting the stream length is not implemented yet. The feature may be supported by video formats implemented by a GDExtension add-on.
*/
func (self Instance) GetStreamLength() Float.X { //gd:VideoStreamPlayer.get_stream_length
	return Float.X(Float.X(Advanced(self).GetStreamLength()))
}

/*
Returns the current frame as a [Texture2D].
*/
func (self Instance) GetVideoTexture() Texture2D.Instance { //gd:VideoStreamPlayer.get_video_texture
	return Texture2D.Instance(Advanced(self).GetVideoTexture())
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.VideoStreamPlayer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("VideoStreamPlayer"))
	casted := Instance{*(*gdclass.VideoStreamPlayer)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) AudioTrack() int {
	return int(int(class(self).GetAudioTrack()))
}

func (self Instance) SetAudioTrack(value int) {
	class(self).SetAudioTrack(int64(value))
}

func (self Instance) Stream() VideoStream.Instance {
	return VideoStream.Instance(class(self).GetStream())
}

func (self Instance) SetStream(value VideoStream.Instance) {
	class(self).SetStream(value)
}

func (self Instance) VolumeDb() Float.X {
	return Float.X(Float.X(class(self).GetVolumeDb()))
}

func (self Instance) SetVolumeDb(value Float.X) {
	class(self).SetVolumeDb(float64(value))
}

func (self Instance) Volume() Float.X {
	return Float.X(Float.X(class(self).GetVolume()))
}

func (self Instance) SetVolume(value Float.X) {
	class(self).SetVolume(float64(value))
}

func (self Instance) Autoplay() bool {
	return bool(class(self).HasAutoplay())
}

func (self Instance) SetAutoplay(value bool) {
	class(self).SetAutoplay(value)
}

func (self Instance) Paused() bool {
	return bool(class(self).IsPaused())
}

func (self Instance) SetPaused(value bool) {
	class(self).SetPaused(value)
}

func (self Instance) Expand() bool {
	return bool(class(self).HasExpand())
}

func (self Instance) SetExpand(value bool) {
	class(self).SetExpand(value)
}

func (self Instance) Loop() bool {
	return bool(class(self).HasLoop())
}

func (self Instance) SetLoop(value bool) {
	class(self).SetLoop(value)
}

func (self Instance) BufferingMsec() int {
	return int(int(class(self).GetBufferingMsec()))
}

func (self Instance) SetBufferingMsec(value int) {
	class(self).SetBufferingMsec(int64(value))
}

func (self Instance) StreamPosition() Float.X {
	return Float.X(Float.X(class(self).GetStreamPosition()))
}

func (self Instance) SetStreamPosition(value Float.X) {
	class(self).SetStreamPosition(float64(value))
}

func (self Instance) Bus() string {
	return string(class(self).GetBus().String())
}

func (self Instance) SetBus(value string) {
	class(self).SetBus(String.Name(String.New(value)))
}

//go:nosplit
func (self class) SetStream(stream [1]gdclass.VideoStream) { //gd:VideoStreamPlayer.set_stream
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(stream[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_set_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStream() [1]gdclass.VideoStream { //gd:VideoStreamPlayer.get_stream
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_get_stream, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.VideoStream{gd.PointerWithOwnershipTransferredToGo[gdclass.VideoStream](r_ret.Get())}
	frame.Free()
	return ret
}

/*
Starts the video playback from the beginning. If the video is paused, this will not unpause the video.
*/
//go:nosplit
func (self class) Play() { //gd:VideoStreamPlayer.play
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_play, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Stops the video playback and sets the stream position to 0.
[b]Note:[/b] Although the stream position will be set to 0, the first frame of the video stream won't become the current frame.
*/
//go:nosplit
func (self class) Stop() { //gd:VideoStreamPlayer.stop
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_stop, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Returns [code]true[/code] if the video is playing.
[b]Note:[/b] The video is still considered playing if paused during playback.
*/
//go:nosplit
func (self class) IsPlaying() bool { //gd:VideoStreamPlayer.is_playing
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_is_playing, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetPaused(paused bool) { //gd:VideoStreamPlayer.set_paused
	var frame = callframe.New()
	callframe.Arg(frame, paused)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_set_paused, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsPaused() bool { //gd:VideoStreamPlayer.is_paused
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_is_paused, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetLoop(loop bool) { //gd:VideoStreamPlayer.set_loop
	var frame = callframe.New()
	callframe.Arg(frame, loop)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_set_loop, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) HasLoop() bool { //gd:VideoStreamPlayer.has_loop
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_has_loop, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolume(volume float64) { //gd:VideoStreamPlayer.set_volume
	var frame = callframe.New()
	callframe.Arg(frame, volume)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_set_volume, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolume() float64 { //gd:VideoStreamPlayer.get_volume
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_get_volume, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVolumeDb(db float64) { //gd:VideoStreamPlayer.set_volume_db
	var frame = callframe.New()
	callframe.Arg(frame, db)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_set_volume_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVolumeDb() float64 { //gd:VideoStreamPlayer.get_volume_db
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_get_volume_db, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAudioTrack(track int64) { //gd:VideoStreamPlayer.set_audio_track
	var frame = callframe.New()
	callframe.Arg(frame, track)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_set_audio_track, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetAudioTrack() int64 { //gd:VideoStreamPlayer.get_audio_track
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_get_audio_track, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Returns the video stream's name, or [code]"<No Stream>"[/code] if no video stream is assigned.
*/
//go:nosplit
func (self class) GetStreamName() String.Readable { //gd:VideoStreamPlayer.get_stream_name
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_get_stream_name, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
The length of the current stream, in seconds.
[b]Note:[/b] For [VideoStreamTheora] streams (the built-in format supported by Godot), this value will always be zero, as getting the stream length is not implemented yet. The feature may be supported by video formats implemented by a GDExtension add-on.
*/
//go:nosplit
func (self class) GetStreamLength() float64 { //gd:VideoStreamPlayer.get_stream_length
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_get_stream_length, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetStreamPosition(position float64) { //gd:VideoStreamPlayer.set_stream_position
	var frame = callframe.New()
	callframe.Arg(frame, position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_set_stream_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetStreamPosition() float64 { //gd:VideoStreamPlayer.get_stream_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_get_stream_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetAutoplay(enabled bool) { //gd:VideoStreamPlayer.set_autoplay
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_set_autoplay, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) HasAutoplay() bool { //gd:VideoStreamPlayer.has_autoplay
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_has_autoplay, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetExpand(enable bool) { //gd:VideoStreamPlayer.set_expand
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_set_expand, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) HasExpand() bool { //gd:VideoStreamPlayer.has_expand
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_has_expand, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBufferingMsec(msec int64) { //gd:VideoStreamPlayer.set_buffering_msec
	var frame = callframe.New()
	callframe.Arg(frame, msec)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_set_buffering_msec, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBufferingMsec() int64 { //gd:VideoStreamPlayer.get_buffering_msec
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_get_buffering_msec, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetBus(bus String.Name) { //gd:VideoStreamPlayer.set_bus
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalStringName(bus)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_set_bus, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetBus() String.Name { //gd:VideoStreamPlayer.get_bus
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_get_bus, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Name(String.Via(gd.StringNameProxy{}, pointers.Pack(pointers.New[gd.StringName](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
Returns the current frame as a [Texture2D].
*/
//go:nosplit
func (self class) GetVideoTexture() [1]gdclass.Texture2D { //gd:VideoStreamPlayer.get_video_texture
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.VideoStreamPlayer.Bind_get_video_texture, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.Texture2D{gd.PointerWithOwnershipTransferredToGo[gdclass.Texture2D](r_ret.Get())}
	frame.Free()
	return ret
}
func (self Instance) OnFinished(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("finished"), gd.NewCallable(cb), 0)
}

func (self class) AsVideoStreamPlayer() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsVideoStreamPlayer() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsVideoStreamPlayer() Instance { return self.Super().AsVideoStreamPlayer() }
func (self class) AsControl() Control.Advanced           { return *((*Control.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsControl() Control.Instance   { return self.Super().AsControl() }
func (self Instance) AsControl() Control.Instance {
	return *((*Control.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsCanvasItem() CanvasItem.Advanced {
	return *((*CanvasItem.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsCanvasItem() CanvasItem.Instance { return self.Super().AsCanvasItem() }
func (self Instance) AsCanvasItem() CanvasItem.Instance {
	return *((*CanvasItem.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsNode() Node.Advanced         { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsNode() Node.Instance { return self.Super().AsNode() }
func (self Instance) AsNode() Node.Instance      { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Control.Advanced(self.AsControl()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Control.Instance(self.AsControl()), name)
	}
}
func init() {
	gdclass.Register("VideoStreamPlayer", func(ptr gd.Object) any {
		return [1]gdclass.VideoStreamPlayer{*(*gdclass.VideoStreamPlayer)(unsafe.Pointer(&ptr))}
	})
}
