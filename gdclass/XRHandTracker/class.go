package XRHandTracker

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/XRPositionalTracker"
import "grow.graphics/gd/gdclass/XRTracker"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A hand tracking system will create an instance of this object and add it to the [XRServer]. This tracking system will then obtain skeleton data, convert it to the Godot Humanoid hand skeleton and store this data on the [XRHandTracker] object.
Use [XRHandModifier3D] to animate a hand mesh using hand tracking data.

*/
type Go [1]classdb.XRHandTracker

/*
Sets flags about the validity of the tracking data for the given hand joint.
*/
func (self Go) SetHandJointFlags(joint classdb.XRHandTrackerHandJoint, flags classdb.XRHandTrackerHandJointFlags) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHandJointFlags(joint, flags)
}

/*
Returns flags about the validity of the tracking data for the given hand joint (see [enum XRHandTracker.HandJointFlags]).
*/
func (self Go) GetHandJointFlags(joint classdb.XRHandTrackerHandJoint) classdb.XRHandTrackerHandJointFlags {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.XRHandTrackerHandJointFlags(class(self).GetHandJointFlags(joint))
}

/*
Sets the transform for the given hand joint.
*/
func (self Go) SetHandJointTransform(joint classdb.XRHandTrackerHandJoint, transform gd.Transform3D) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHandJointTransform(joint, transform)
}

/*
Returns the transform for the given hand joint.
*/
func (self Go) GetHandJointTransform(joint classdb.XRHandTrackerHandJoint) gd.Transform3D {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Transform3D(class(self).GetHandJointTransform(joint))
}

/*
Sets the radius of the given hand joint.
*/
func (self Go) SetHandJointRadius(joint classdb.XRHandTrackerHandJoint, radius float64) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHandJointRadius(joint, gd.Float(radius))
}

/*
Returns the radius of the given hand joint.
*/
func (self Go) GetHandJointRadius(joint classdb.XRHandTrackerHandJoint) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetHandJointRadius(joint)))
}

/*
Sets the linear velocity for the given hand joint.
*/
func (self Go) SetHandJointLinearVelocity(joint classdb.XRHandTrackerHandJoint, linear_velocity gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHandJointLinearVelocity(joint, linear_velocity)
}

/*
Returns the linear velocity for the given hand joint.
*/
func (self Go) GetHandJointLinearVelocity(joint classdb.XRHandTrackerHandJoint) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetHandJointLinearVelocity(joint))
}

/*
Sets the angular velocity for the given hand joint.
*/
func (self Go) SetHandJointAngularVelocity(joint classdb.XRHandTrackerHandJoint, angular_velocity gd.Vector3) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHandJointAngularVelocity(joint, angular_velocity)
}

/*
Returns the angular velocity for the given hand joint.
*/
func (self Go) GetHandJointAngularVelocity(joint classdb.XRHandTrackerHandJoint) gd.Vector3 {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Vector3(class(self).GetHandJointAngularVelocity(joint))
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.XRHandTracker
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("XRHandTracker"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) HasTrackingData() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).GetHasTrackingData())
}

func (self Go) SetHasTrackingData(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHasTrackingData(value)
}

func (self Go) HandTrackingSource() classdb.XRHandTrackerHandTrackingSource {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.XRHandTrackerHandTrackingSource(class(self).GetHandTrackingSource())
}

func (self Go) SetHandTrackingSource(value classdb.XRHandTrackerHandTrackingSource) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHandTrackingSource(value)
}

//go:nosplit
func (self class) SetHasTrackingData(has_data bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, has_data)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRHandTracker.Bind_set_has_tracking_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHasTrackingData() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRHandTracker.Bind_get_has_tracking_data, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetHandTrackingSource(source classdb.XRHandTrackerHandTrackingSource)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, source)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRHandTracker.Bind_set_hand_tracking_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetHandTrackingSource() classdb.XRHandTrackerHandTrackingSource {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.XRHandTrackerHandTrackingSource](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRHandTracker.Bind_get_hand_tracking_source, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets flags about the validity of the tracking data for the given hand joint.
*/
//go:nosplit
func (self class) SetHandJointFlags(joint classdb.XRHandTrackerHandJoint, flags classdb.XRHandTrackerHandJointFlags)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, flags)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRHandTracker.Bind_set_hand_joint_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns flags about the validity of the tracking data for the given hand joint (see [enum XRHandTracker.HandJointFlags]).
*/
//go:nosplit
func (self class) GetHandJointFlags(joint classdb.XRHandTrackerHandJoint) classdb.XRHandTrackerHandJointFlags {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[classdb.XRHandTrackerHandJointFlags](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRHandTracker.Bind_get_hand_joint_flags, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the transform for the given hand joint.
*/
//go:nosplit
func (self class) SetHandJointTransform(joint classdb.XRHandTrackerHandJoint, transform gd.Transform3D)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, transform)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRHandTracker.Bind_set_hand_joint_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the transform for the given hand joint.
*/
//go:nosplit
func (self class) GetHandJointTransform(joint classdb.XRHandTrackerHandJoint) gd.Transform3D {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Transform3D](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRHandTracker.Bind_get_hand_joint_transform, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the radius of the given hand joint.
*/
//go:nosplit
func (self class) SetHandJointRadius(joint classdb.XRHandTrackerHandJoint, radius gd.Float)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, radius)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRHandTracker.Bind_set_hand_joint_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the radius of the given hand joint.
*/
//go:nosplit
func (self class) GetHandJointRadius(joint classdb.XRHandTrackerHandJoint) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRHandTracker.Bind_get_hand_joint_radius, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the linear velocity for the given hand joint.
*/
//go:nosplit
func (self class) SetHandJointLinearVelocity(joint classdb.XRHandTrackerHandJoint, linear_velocity gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, linear_velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRHandTracker.Bind_set_hand_joint_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the linear velocity for the given hand joint.
*/
//go:nosplit
func (self class) GetHandJointLinearVelocity(joint classdb.XRHandTrackerHandJoint) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRHandTracker.Bind_get_hand_joint_linear_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the angular velocity for the given hand joint.
*/
//go:nosplit
func (self class) SetHandJointAngularVelocity(joint classdb.XRHandTrackerHandJoint, angular_velocity gd.Vector3)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	callframe.Arg(frame, angular_velocity)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRHandTracker.Bind_set_hand_joint_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the angular velocity for the given hand joint.
*/
//go:nosplit
func (self class) GetHandJointAngularVelocity(joint classdb.XRHandTrackerHandJoint) gd.Vector3 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, joint)
	var r_ret = callframe.Ret[gd.Vector3](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.XRHandTracker.Bind_get_hand_joint_angular_velocity, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsXRHandTracker() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsXRHandTracker() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsXRPositionalTracker() XRPositionalTracker.GD { return *((*XRPositionalTracker.GD)(unsafe.Pointer(&self))) }
func (self Go) AsXRPositionalTracker() XRPositionalTracker.Go { return *((*XRPositionalTracker.Go)(unsafe.Pointer(&self))) }
func (self class) AsXRTracker() XRTracker.GD { return *((*XRTracker.GD)(unsafe.Pointer(&self))) }
func (self Go) AsXRTracker() XRTracker.Go { return *((*XRTracker.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsXRPositionalTracker(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsXRPositionalTracker(), name)
	}
}
func init() {classdb.Register("XRHandTracker", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type HandTrackingSource = classdb.XRHandTrackerHandTrackingSource

const (
/*The source of hand tracking data is unknown.*/
	HandTrackingSourceUnknown HandTrackingSource = 0
/*The source of hand tracking data is unobstructed, meaning that an accurate method of hand tracking is used. These include optical hand tracking, data gloves, etc.*/
	HandTrackingSourceUnobstructed HandTrackingSource = 1
/*The source of hand tracking data is a controller, meaning that joint positions are inferred from controller inputs.*/
	HandTrackingSourceController HandTrackingSource = 2
/*Represents the size of the [enum HandTrackingSource] enum.*/
	HandTrackingSourceMax HandTrackingSource = 3
)
type HandJoint = classdb.XRHandTrackerHandJoint

const (
/*Palm joint.*/
	HandJointPalm HandJoint = 0
/*Wrist joint.*/
	HandJointWrist HandJoint = 1
/*Thumb metacarpal joint.*/
	HandJointThumbMetacarpal HandJoint = 2
/*Thumb phalanx proximal joint.*/
	HandJointThumbPhalanxProximal HandJoint = 3
/*Thumb phalanx distal joint.*/
	HandJointThumbPhalanxDistal HandJoint = 4
/*Thumb tip joint.*/
	HandJointThumbTip HandJoint = 5
/*Index finger metacarpal joint.*/
	HandJointIndexFingerMetacarpal HandJoint = 6
/*Index finger phalanx proximal joint.*/
	HandJointIndexFingerPhalanxProximal HandJoint = 7
/*Index finger phalanx intermediate joint.*/
	HandJointIndexFingerPhalanxIntermediate HandJoint = 8
/*Index finger phalanx distal joint.*/
	HandJointIndexFingerPhalanxDistal HandJoint = 9
/*Index finger tip joint.*/
	HandJointIndexFingerTip HandJoint = 10
/*Middle finger metacarpal joint.*/
	HandJointMiddleFingerMetacarpal HandJoint = 11
/*Middle finger phalanx proximal joint.*/
	HandJointMiddleFingerPhalanxProximal HandJoint = 12
/*Middle finger phalanx intermediate joint.*/
	HandJointMiddleFingerPhalanxIntermediate HandJoint = 13
/*Middle finger phalanx distal joint.*/
	HandJointMiddleFingerPhalanxDistal HandJoint = 14
/*Middle finger tip joint.*/
	HandJointMiddleFingerTip HandJoint = 15
/*Ring finger metacarpal joint.*/
	HandJointRingFingerMetacarpal HandJoint = 16
/*Ring finger phalanx proximal joint.*/
	HandJointRingFingerPhalanxProximal HandJoint = 17
/*Ring finger phalanx intermediate joint.*/
	HandJointRingFingerPhalanxIntermediate HandJoint = 18
/*Ring finger phalanx distal joint.*/
	HandJointRingFingerPhalanxDistal HandJoint = 19
/*Ring finger tip joint.*/
	HandJointRingFingerTip HandJoint = 20
/*Pinky finger metacarpal joint.*/
	HandJointPinkyFingerMetacarpal HandJoint = 21
/*Pinky finger phalanx proximal joint.*/
	HandJointPinkyFingerPhalanxProximal HandJoint = 22
/*Pinky finger phalanx intermediate joint.*/
	HandJointPinkyFingerPhalanxIntermediate HandJoint = 23
/*Pinky finger phalanx distal joint.*/
	HandJointPinkyFingerPhalanxDistal HandJoint = 24
/*Pinky finger tip joint.*/
	HandJointPinkyFingerTip HandJoint = 25
/*Represents the size of the [enum HandJoint] enum.*/
	HandJointMax HandJoint = 26
)
type HandJointFlags = classdb.XRHandTrackerHandJointFlags

const (
/*The hand joint's orientation data is valid.*/
	HandJointFlagOrientationValid HandJointFlags = 1
/*The hand joint's orientation is actively tracked. May not be set if tracking has been temporarily lost.*/
	HandJointFlagOrientationTracked HandJointFlags = 2
/*The hand joint's position data is valid.*/
	HandJointFlagPositionValid HandJointFlags = 4
/*The hand joint's position is actively tracked. May not be set if tracking has been temporarily lost.*/
	HandJointFlagPositionTracked HandJointFlags = 8
/*The hand joint's linear velocity data is valid.*/
	HandJointFlagLinearVelocityValid HandJointFlags = 16
/*The hand joint's angular velocity data is valid.*/
	HandJointFlagAngularVelocityValid HandJointFlags = 32
)