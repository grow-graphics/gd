// Code generated by the generate package DO NOT EDIT

// Package RemoteTransform3D provides methods for working with RemoteTransform3D object instances.
package RemoteTransform3D

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
import "graphics.gd/classdb/Node3D"
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
RemoteTransform3D pushes its own [Transform3D] to another [Node3D] derived Node (called the remote node) in the scene.
It can be set to update another Node's position, rotation and/or scale. It can use either global or local coordinates.
*/
type Instance [1]gdclass.RemoteTransform3D

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsRemoteTransform3D() Instance
}

/*
[RemoteTransform3D] caches the remote node. It may not notice if the remote node disappears; [method force_update_cache] forces it to update the cache again.
*/
func (self Instance) ForceUpdateCache() { //gd:RemoteTransform3D.force_update_cache
	Advanced(self).ForceUpdateCache()
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.RemoteTransform3D

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("RemoteTransform3D"))
	casted := Instance{*(*gdclass.RemoteTransform3D)(unsafe.Pointer(&object))}
	return casted
}

func (self Instance) RemotePath() string {
	return string(class(self).GetRemoteNode().String())
}

func (self Instance) SetRemotePath(value string) {
	class(self).SetRemoteNode(Path.ToNode(String.New(value)))
}

func (self Instance) UseGlobalCoordinates() bool {
	return bool(class(self).GetUseGlobalCoordinates())
}

func (self Instance) SetUseGlobalCoordinates(value bool) {
	class(self).SetUseGlobalCoordinates(value)
}

func (self Instance) UpdatePosition() bool {
	return bool(class(self).GetUpdatePosition())
}

func (self Instance) SetUpdatePosition(value bool) {
	class(self).SetUpdatePosition(value)
}

func (self Instance) UpdateRotation() bool {
	return bool(class(self).GetUpdateRotation())
}

func (self Instance) SetUpdateRotation(value bool) {
	class(self).SetUpdateRotation(value)
}

func (self Instance) UpdateScale() bool {
	return bool(class(self).GetUpdateScale())
}

func (self Instance) SetUpdateScale(value bool) {
	class(self).SetUpdateScale(value)
}

//go:nosplit
func (self class) SetRemoteNode(path Path.ToNode) { //gd:RemoteTransform3D.set_remote_node
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalNodePath(path)))
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_set_remote_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetRemoteNode() Path.ToNode { //gd:RemoteTransform3D.get_remote_node
	var frame = callframe.New()
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_get_remote_node, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Path.ToNode(String.Via(gd.NodePathProxy{}, pointers.Pack(pointers.New[gd.NodePath](r_ret.Get()))))
	frame.Free()
	return ret
}

/*
[RemoteTransform3D] caches the remote node. It may not notice if the remote node disappears; [method force_update_cache] forces it to update the cache again.
*/
//go:nosplit
func (self class) ForceUpdateCache() { //gd:RemoteTransform3D.force_update_cache
	var frame = callframe.New()
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_force_update_cache, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) SetUseGlobalCoordinates(use_global_coordinates bool) { //gd:RemoteTransform3D.set_use_global_coordinates
	var frame = callframe.New()
	callframe.Arg(frame, use_global_coordinates)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_set_use_global_coordinates, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUseGlobalCoordinates() bool { //gd:RemoteTransform3D.get_use_global_coordinates
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_get_use_global_coordinates, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUpdatePosition(update_remote_position bool) { //gd:RemoteTransform3D.set_update_position
	var frame = callframe.New()
	callframe.Arg(frame, update_remote_position)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_set_update_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUpdatePosition() bool { //gd:RemoteTransform3D.get_update_position
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_get_update_position, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUpdateRotation(update_remote_rotation bool) { //gd:RemoteTransform3D.set_update_rotation
	var frame = callframe.New()
	callframe.Arg(frame, update_remote_rotation)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_set_update_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUpdateRotation() bool { //gd:RemoteTransform3D.get_update_rotation
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_get_update_rotation, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}

//go:nosplit
func (self class) SetUpdateScale(update_remote_scale bool) { //gd:RemoteTransform3D.set_update_scale
	var frame = callframe.New()
	callframe.Arg(frame, update_remote_scale)
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_set_update_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

//go:nosplit
func (self class) GetUpdateScale() bool { //gd:RemoteTransform3D.get_update_scale
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.RemoteTransform3D.Bind_get_update_scale, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsRemoteTransform3D() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsRemoteTransform3D() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsRemoteTransform3D() Instance { return self.Super().AsRemoteTransform3D() }
func (self class) AsNode3D() Node3D.Advanced             { return *((*Node3D.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsNode3D() Node3D.Instance     { return self.Super().AsNode3D() }
func (self Instance) AsNode3D() Node3D.Instance          { return *((*Node3D.Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced                 { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsNode() Node.Instance         { return self.Super().AsNode() }
func (self Instance) AsNode() Node.Instance              { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Advanced(self.AsNode3D()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node3D.Instance(self.AsNode3D()), name)
	}
}
func init() {
	gdclass.Register("RemoteTransform3D", func(ptr gd.Object) any {
		return [1]gdclass.RemoteTransform3D{*(*gdclass.RemoteTransform3D)(unsafe.Pointer(&ptr))}
	})
}
