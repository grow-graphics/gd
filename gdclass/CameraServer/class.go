package CameraServer

import "unsafe"
import "sync"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
The [CameraServer] keeps track of different cameras accessible in Godot. These are external cameras such as webcams or the cameras on your phone.
It is notably used to provide AR modules with a video feed from the camera.
[b]Note:[/b] This class is currently only implemented on macOS and iOS. To get a [CameraFeed] on iOS, the camera plugin from [url=https://github.com/godotengine/godot-ios-plugins]godot-ios-plugins[/url] is required. On other platforms, no [CameraFeed]s will be available.

*/
var self gdclass.CameraServer
var once sync.Once
func singleton() {
	gc := gd.Static
	obj := gc.API.Object.GetSingleton(gc, gc.API.Singletons.CameraServer)
	self = *(*gdclass.CameraServer)(unsafe.Pointer(&obj))
}

/*
Returns the [CameraFeed] corresponding to the camera with the given [param index].
*/
func GetFeed(index int) gdclass.CameraFeed {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gdclass.CameraFeed(class(self).GetFeed(gc, gd.Int(index)))
}

/*
Returns the number of [CameraFeed]s registered.
*/
func GetFeedCount() int {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return int(int(class(self).GetFeedCount()))
}

/*
Returns an array of [CameraFeed]s.
*/
func Feeds() gd.ArrayOf[gdclass.CameraFeed] {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	return gd.ArrayOf[gdclass.CameraFeed](class(self).Feeds(gc))
}

/*
Adds the camera [param feed] to the camera server.
*/
func AddFeed(feed gdclass.CameraFeed) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).AddFeed(feed)
}

/*
Removes the specified camera [param feed].
*/
func RemoveFeed(feed gdclass.CameraFeed) {
	gc := gd.GarbageCollector(); _ = gc
	once.Do(singleton)
	class(self).RemoveFeed(feed)
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
func GD() class { once.Do(singleton); return self }
type class [1]classdb.CameraServer
func (self class) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }

/*
Returns the [CameraFeed] corresponding to the camera with the given [param index].
*/
//go:nosplit
func (self class) GetFeed(ctx gd.Lifetime, index gd.Int) gdclass.CameraFeed {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, index)
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraServer.Bind_get_feed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.CameraFeed
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Returns the number of [CameraFeed]s registered.
*/
//go:nosplit
func (self class) GetFeedCount() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraServer.Bind_get_feed_count, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns an array of [CameraFeed]s.
*/
//go:nosplit
func (self class) Feeds(ctx gd.Lifetime) gd.ArrayOf[gdclass.CameraFeed] {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraServer.Bind_feeds, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Array](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return gd.TypedArray[gdclass.CameraFeed](ret)
}
/*
Adds the camera [param feed] to the camera server.
*/
//go:nosplit
func (self class) AddFeed(feed gdclass.CameraFeed)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(feed[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraServer.Bind_add_feed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Removes the specified camera [param feed].
*/
//go:nosplit
func (self class) RemoveFeed(feed gdclass.CameraFeed)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(feed[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.CameraServer.Bind_remove_feed, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func OnCameraFeedAdded(cb func(id int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("camera_feed_added"), gc.Callable(cb), 0)
}


func OnCameraFeedRemoved(cb func(id int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("camera_feed_removed"), gc.Callable(cb), 0)
}


func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsObject(), name)
	}
}
func init() {classdb.Register("CameraServer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type FeedImage = classdb.CameraServerFeedImage

const (
/*The RGBA camera image.*/
	FeedRgbaImage FeedImage = 0
/*The [url=https://en.wikipedia.org/wiki/YCbCr]YCbCr[/url] camera image.*/
	FeedYcbcrImage FeedImage = 0
/*The Y component camera image.*/
	FeedYImage FeedImage = 0
/*The CbCr component camera image.*/
	FeedCbcrImage FeedImage = 1
)