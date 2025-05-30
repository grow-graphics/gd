// Code generated by the generate package DO NOT EDIT

// Package MultiplayerSynchronizer provides methods for working with MultiplayerSynchronizer object instances.
package MultiplayerSynchronizer

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
import "graphics.gd/classdb/Node"
import "graphics.gd/classdb/SceneReplicationConfig"
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
By default, [MultiplayerSynchronizer] synchronizes configured properties to all peers.
Visibility can be handled directly with [method set_visibility_for] or as-needed with [method add_visibility_filter] and [method update_visibility].
[MultiplayerSpawner]s will handle nodes according to visibility of synchronizers as long as the node at [member root_path] was spawned by one.
Internally, [MultiplayerSynchronizer] uses [method MultiplayerAPI.object_configuration_add] to notify synchronization start passing the [Node] at [member root_path] as the [code]object[/code] and itself as the [code]configuration[/code], and uses [method MultiplayerAPI.object_configuration_remove] to notify synchronization end in a similar way.
[b]Note:[/b] Synchronization is not supported for [Object] type properties, like [Resource]. Properties that are unique to each peer, like the instance IDs of [Object]s (see [method Object.get_instance_id]) or [RID]s, will also not work in synchronization.
*/
type Instance [1]gdclass.MultiplayerSynchronizer

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

type Expanded [1]gdclass.MultiplayerSynchronizer

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsMultiplayerSynchronizer() Instance
}

/*
Updates the visibility of [param for_peer] according to visibility filters. If [param for_peer] is [code]0[/code] (the default), all peers' visibilties are updated.
*/
func (self Instance) UpdateVisibility() { //gd:MultiplayerSynchronizer.update_visibility
	Advanced(self).UpdateVisibility(int64(0))
}

/*
Updates the visibility of [param for_peer] according to visibility filters. If [param for_peer] is [code]0[/code] (the default), all peers' visibilties are updated.
*/
func (self Expanded) UpdateVisibility(for_peer int) { //gd:MultiplayerSynchronizer.update_visibility
	Advanced(self).UpdateVisibility(int64(for_peer))
}

/*
Adds a peer visibility filter for this synchronizer.
[param filter] should take a peer ID [int] and return a [bool].
*/
func (self Instance) AddVisibilityFilter(filter func(peer_id int) bool) { //gd:MultiplayerSynchronizer.add_visibility_filter
	Advanced(self).AddVisibilityFilter(Callable.New(filter))
}

/*
Removes a peer visibility filter from this synchronizer.
*/
func (self Instance) RemoveVisibilityFilter(filter Callable.Function) { //gd:MultiplayerSynchronizer.remove_visibility_filter
	Advanced(self).RemoveVisibilityFilter(Callable.New(filter))
}

/*
Sets the visibility of [param peer] to [param visible]. If [param peer] is [code]0[/code], the value of [member public_visibility] will be updated instead.
*/
func (self Instance) SetVisibilityFor(peer int, visible bool) { //gd:MultiplayerSynchronizer.set_visibility_for
	Advanced(self).SetVisibilityFor(int64(peer), visible)
}

/*
Queries the current visibility for peer [param peer].
*/
func (self Instance) GetVisibilityFor(peer int) bool { //gd:MultiplayerSynchronizer.get_visibility_for
	return bool(Advanced(self).GetVisibilityFor(int64(peer)))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.MultiplayerSynchronizer

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("MultiplayerSynchronizer"))
	casted := Instance{*(*gdclass.MultiplayerSynchronizer)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) RootPath() string {
	return string(class(self).GetRootPath().String())
}

func (self Instance) SetRootPath(value string) {
	class(self).SetRootPath(Path.ToNode(String.New(value)))
}

func (self Instance) ReplicationInterval() Float.X {
	return Float.X(Float.X(class(self).GetReplicationInterval()))
}

func (self Instance) SetReplicationInterval(value Float.X) {
	class(self).SetReplicationInterval(float64(value))
}

func (self Instance) DeltaInterval() Float.X {
	return Float.X(Float.X(class(self).GetDeltaInterval()))
}

func (self Instance) SetDeltaInterval(value Float.X) {
	class(self).SetDeltaInterval(float64(value))
}

func (self Instance) ReplicationConfig() SceneReplicationConfig.Instance {
	return SceneReplicationConfig.Instance(class(self).GetReplicationConfig())
}

func (self Instance) SetReplicationConfig(value SceneReplicationConfig.Instance) {
	class(self).SetReplicationConfig(value)
}

func (self Instance) VisibilityUpdateMode() VisibilityUpdateMode {
	return VisibilityUpdateMode(class(self).GetVisibilityUpdateMode())
}

func (self Instance) SetVisibilityUpdateMode(value VisibilityUpdateMode) {
	class(self).SetVisibilityUpdateMode(value)
}

func (self Instance) PublicVisibility() bool {
	return bool(class(self).IsVisibilityPublic())
}

func (self Instance) SetPublicVisibility(value bool) {
	class(self).SetVisibilityPublic(value)
}

//go:nosplit
func (self class) SetRootPath(path Path.ToNode) { //gd:MultiplayerSynchronizer.set_root_path
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_set_root_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRootPath() Path.ToNode { //gd:MultiplayerSynchronizer.get_root_path
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_get_root_path, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetReplicationInterval(milliseconds float64) { //gd:MultiplayerSynchronizer.set_replication_interval
	var frame = callframe.New()
	callframe.Arg(frame, milliseconds)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_set_replication_interval, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetReplicationInterval() float64 { //gd:MultiplayerSynchronizer.get_replication_interval
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_get_replication_interval, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetDeltaInterval(milliseconds float64) { //gd:MultiplayerSynchronizer.set_delta_interval
	var frame = callframe.New()
	callframe.Arg(frame, milliseconds)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_set_delta_interval, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetDeltaInterval() float64 { //gd:MultiplayerSynchronizer.get_delta_interval
	var frame = callframe.New()
	var r_ret = callframe.Ret[float64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_get_delta_interval, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetReplicationConfig(config [1]gdclass.SceneReplicationConfig) { //gd:MultiplayerSynchronizer.set_replication_config
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(config[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_set_replication_config, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetReplicationConfig() [1]gdclass.SceneReplicationConfig { //gd:MultiplayerSynchronizer.get_replication_config
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_get_replication_config, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = [1]gdclass.SceneReplicationConfig{gd.PointerWithOwnershipTransferredToGo[gdclass.SceneReplicationConfig](r_ret.Get())}
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetVisibilityUpdateMode(mode VisibilityUpdateMode) { //gd:MultiplayerSynchronizer.set_visibility_update_mode
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_set_visibility_update_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetVisibilityUpdateMode() VisibilityUpdateMode { //gd:MultiplayerSynchronizer.get_visibility_update_mode
	var frame = callframe.New()
	var r_ret = callframe.Ret[VisibilityUpdateMode](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_get_visibility_update_mode, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Updates the visibility of [param for_peer] according to visibility filters. If [param for_peer] is [code]0[/code] (the default), all peers' visibilties are updated.
*/
//go:nosplit
func (self class) UpdateVisibility(for_peer int64) { //gd:MultiplayerSynchronizer.update_visibility
	var frame = callframe.New()
	callframe.Arg(frame, for_peer)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_update_visibility, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetVisibilityPublic(visible bool) { //gd:MultiplayerSynchronizer.set_visibility_public
	var frame = callframe.New()
	callframe.Arg(frame, visible)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_set_visibility_public, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) IsVisibilityPublic() bool { //gd:MultiplayerSynchronizer.is_visibility_public
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_is_visibility_public, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

/*
Adds a peer visibility filter for this synchronizer.
[param filter] should take a peer ID [int] and return a [bool].
*/
//go:nosplit
func (self class) AddVisibilityFilter(filter Callable.Function) { //gd:MultiplayerSynchronizer.add_visibility_filter
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(filter)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_add_visibility_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Removes a peer visibility filter from this synchronizer.
*/
//go:nosplit
func (self class) RemoveVisibilityFilter(filter Callable.Function) { //gd:MultiplayerSynchronizer.remove_visibility_filter
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalCallable(filter)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_remove_visibility_filter, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Sets the visibility of [param peer] to [param visible]. If [param peer] is [code]0[/code], the value of [member public_visibility] will be updated instead.
*/
//go:nosplit
func (self class) SetVisibilityFor(peer int64, visible bool) { //gd:MultiplayerSynchronizer.set_visibility_for
	var frame = callframe.New()
	callframe.Arg(frame, peer)
	callframe.Arg(frame, visible)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_set_visibility_for, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Queries the current visibility for peer [param peer].
*/
//go:nosplit
func (self class) GetVisibilityFor(peer int64) bool { //gd:MultiplayerSynchronizer.get_visibility_for
	var frame = callframe.New()
	callframe.Arg(frame, peer)
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.MultiplayerSynchronizer.Bind_get_visibility_for, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Instance) OnSynchronized(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("synchronized"), gd.NewCallable(cb), 0)
}

func (self Instance) OnDeltaSynchronized(cb func()) {
	self[0].AsObject()[0].Connect(gd.NewStringName("delta_synchronized"), gd.NewCallable(cb), 0)
}

func (self Instance) OnVisibilityChanged(cb func(for_peer int)) {
	self[0].AsObject()[0].Connect(gd.NewStringName("visibility_changed"), gd.NewCallable(cb), 0)
}

func (self class) AsMultiplayerSynchronizer() Advanced { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsMultiplayerSynchronizer() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsMultiplayerSynchronizer() Instance {
	return self.Super().AsMultiplayerSynchronizer()
}
func (self class) AsNode() Node.Advanced         { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsNode() Node.Instance { return self.Super().AsNode() }
func (self Instance) AsNode() Node.Instance      { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Advanced(self.AsNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Instance(self.AsNode()), name)
	}
}
func init() {
	gdclass.Register("MultiplayerSynchronizer", func(ptr gd.Object) any {
		return [1]gdclass.MultiplayerSynchronizer{*(*gdclass.MultiplayerSynchronizer)(unsafe.Pointer(&ptr))}
	})
}

type VisibilityUpdateMode int //gd:MultiplayerSynchronizer.VisibilityUpdateMode

const (
	/*Visibility filters are updated during process frames (see [constant Node.NOTIFICATION_INTERNAL_PROCESS]).*/
	VisibilityProcessIdle VisibilityUpdateMode = 0
	/*Visibility filters are updated during physics frames (see [constant Node.NOTIFICATION_INTERNAL_PHYSICS_PROCESS]).*/
	VisibilityProcessPhysics VisibilityUpdateMode = 1
	/*Visibility filters are not updated automatically, and must be updated manually by calling [method update_visibility].*/
	VisibilityProcessNone VisibilityUpdateMode = 2
)
