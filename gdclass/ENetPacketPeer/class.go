package ENetPacketPeer

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import "grow.graphics/gd/gdclass"
import classdb "grow.graphics/gd/internal/classdb"
import "grow.graphics/gd/gdclass/PacketPeer"

var _ unsafe.Pointer
var _ gdclass.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
A PacketPeer implementation representing a peer of an [ENetConnection].
This class cannot be instantiated directly but can be retrieved during [method ENetConnection.service] or via [method ENetConnection.get_peers].
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.

*/
type Go [1]classdb.ENetPacketPeer

/*
Request a disconnection from a peer. An [constant ENetConnection.EVENT_DISCONNECT] will be generated during [method ENetConnection.service] once the disconnection is complete.
*/
func (self Go) PeerDisconnect() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PeerDisconnect(gd.Int(0))
}

/*
Request a disconnection from a peer, but only after all queued outgoing packets are sent. An [constant ENetConnection.EVENT_DISCONNECT] will be generated during [method ENetConnection.service] once the disconnection is complete.
*/
func (self Go) PeerDisconnectLater() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PeerDisconnectLater(gd.Int(0))
}

/*
Force an immediate disconnection from a peer. No [constant ENetConnection.EVENT_DISCONNECT] will be generated. The foreign peer is not guaranteed to receive the disconnect notification, and is reset immediately upon return from this function.
*/
func (self Go) PeerDisconnectNow() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PeerDisconnectNow(gd.Int(0))
}

/*
Sends a ping request to a peer. ENet automatically pings all connected peers at regular intervals, however, this function may be called to ensure more frequent ping requests.
*/
func (self Go) Ping() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Ping()
}

/*
Sets the [param ping_interval] in milliseconds at which pings will be sent to a peer. Pings are used both to monitor the liveness of the connection and also to dynamically adjust the throttle during periods of low traffic so that the throttle has reasonable responsiveness during traffic spikes. The default ping interval is [code]500[/code] milliseconds.
*/
func (self Go) PingInterval(ping_interval int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).PingInterval(gd.Int(ping_interval))
}

/*
Forcefully disconnects a peer. The foreign host represented by the peer is not notified of the disconnection and will timeout on its connection to the local host.
*/
func (self Go) Reset() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Reset()
}

/*
Queues a [param packet] to be sent over the specified [param channel]. See [code]FLAG_*[/code] constants for available packet flags.
*/
func (self Go) Send(channel int, packet []byte, flags int) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Send(gd.Int(channel), gc.PackedByteSlice(packet), gd.Int(flags)))
}

/*
Configures throttle parameter for a peer.
Unreliable packets are dropped by ENet in response to the varying conditions of the Internet connection to the peer. The throttle represents a probability that an unreliable packet should not be dropped and thus sent by ENet to the peer. By measuring fluctuations in round trip times of reliable packets over the specified [param interval], ENet will either increase the probability by the amount specified in the [param acceleration] parameter, or decrease it by the amount specified in the [param deceleration] parameter (both are ratios to [constant PACKET_THROTTLE_SCALE]).
When the throttle has a value of [constant PACKET_THROTTLE_SCALE], no unreliable packets are dropped by ENet, and so 100% of all unreliable packets will be sent.
When the throttle has a value of [code]0[/code], all unreliable packets are dropped by ENet, and so 0% of all unreliable packets will be sent.
Intermediate values for the throttle represent intermediate probabilities between 0% and 100% of unreliable packets being sent. The bandwidth limits of the local and foreign hosts are taken into account to determine a sensible limit for the throttle probability above which it should not raise even in the best of conditions.
*/
func (self Go) ThrottleConfigure(interval int, acceleration int, deceleration int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).ThrottleConfigure(gd.Int(interval), gd.Int(acceleration), gd.Int(deceleration))
}

/*
Sets the timeout parameters for a peer. The timeout parameters control how and when a peer will timeout from a failure to acknowledge reliable traffic. Timeout values are expressed in milliseconds.
The [param timeout] is a factor that, multiplied by a value based on the average round trip time, will determine the timeout limit for a reliable packet. When that limit is reached, the timeout will be doubled, and the peer will be disconnected if that limit has reached [param timeout_min]. The [param timeout_max] parameter, on the other hand, defines a fixed timeout for which any packet must be acknowledged or the peer will be dropped.
*/
func (self Go) SetTimeout(timeout int, timeout_min int, timeout_max int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTimeout(gd.Int(timeout), gd.Int(timeout_min), gd.Int(timeout_max))
}

/*
Returns the IP address of this peer.
*/
func (self Go) GetRemoteAddress() string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).GetRemoteAddress(gc).String())
}

/*
Returns the remote port of this peer.
*/
func (self Go) GetRemotePort() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetRemotePort()))
}

/*
Returns the requested [param statistic] for this peer. See [enum PeerStatistic].
*/
func (self Go) GetStatistic(statistic classdb.ENetPacketPeerPeerStatistic) float64 {
	gc := gd.GarbageCollector(); _ = gc
	return float64(float64(class(self).GetStatistic(statistic)))
}

/*
Returns the current peer state. See [enum PeerState].
*/
func (self Go) GetState() classdb.ENetPacketPeerPeerState {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.ENetPacketPeerPeerState(class(self).GetState())
}

/*
Returns the number of channels allocated for communication with peer.
*/
func (self Go) GetChannels() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetChannels()))
}

/*
Returns [code]true[/code] if the peer is currently active (i.e. the associated [ENetConnection] is still valid).
*/
func (self Go) IsActive() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsActive())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.ENetPacketPeer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("ENetPacketPeer"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Request a disconnection from a peer. An [constant ENetConnection.EVENT_DISCONNECT] will be generated during [method ENetConnection.service] once the disconnection is complete.
*/
//go:nosplit
func (self class) PeerDisconnect(data gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, data)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetPacketPeer.Bind_peer_disconnect, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Request a disconnection from a peer, but only after all queued outgoing packets are sent. An [constant ENetConnection.EVENT_DISCONNECT] will be generated during [method ENetConnection.service] once the disconnection is complete.
*/
//go:nosplit
func (self class) PeerDisconnectLater(data gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, data)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetPacketPeer.Bind_peer_disconnect_later, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Force an immediate disconnection from a peer. No [constant ENetConnection.EVENT_DISCONNECT] will be generated. The foreign peer is not guaranteed to receive the disconnect notification, and is reset immediately upon return from this function.
*/
//go:nosplit
func (self class) PeerDisconnectNow(data gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, data)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetPacketPeer.Bind_peer_disconnect_now, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sends a ping request to a peer. ENet automatically pings all connected peers at regular intervals, however, this function may be called to ensure more frequent ping requests.
*/
//go:nosplit
func (self class) Ping()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetPacketPeer.Bind_ping, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the [param ping_interval] in milliseconds at which pings will be sent to a peer. Pings are used both to monitor the liveness of the connection and also to dynamically adjust the throttle during periods of low traffic so that the throttle has reasonable responsiveness during traffic spikes. The default ping interval is [code]500[/code] milliseconds.
*/
//go:nosplit
func (self class) PingInterval(ping_interval gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, ping_interval)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetPacketPeer.Bind_ping_interval, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Forcefully disconnects a peer. The foreign host represented by the peer is not notified of the disconnection and will timeout on its connection to the local host.
*/
//go:nosplit
func (self class) Reset()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetPacketPeer.Bind_reset, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Queues a [param packet] to be sent over the specified [param channel]. See [code]FLAG_*[/code] constants for available packet flags.
*/
//go:nosplit
func (self class) Send(channel gd.Int, packet gd.PackedByteArray, flags gd.Int) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, channel)
	callframe.Arg(frame, mmm.Get(packet))
	callframe.Arg(frame, flags)
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetPacketPeer.Bind_send, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Configures throttle parameter for a peer.
Unreliable packets are dropped by ENet in response to the varying conditions of the Internet connection to the peer. The throttle represents a probability that an unreliable packet should not be dropped and thus sent by ENet to the peer. By measuring fluctuations in round trip times of reliable packets over the specified [param interval], ENet will either increase the probability by the amount specified in the [param acceleration] parameter, or decrease it by the amount specified in the [param deceleration] parameter (both are ratios to [constant PACKET_THROTTLE_SCALE]).
When the throttle has a value of [constant PACKET_THROTTLE_SCALE], no unreliable packets are dropped by ENet, and so 100% of all unreliable packets will be sent.
When the throttle has a value of [code]0[/code], all unreliable packets are dropped by ENet, and so 0% of all unreliable packets will be sent.
Intermediate values for the throttle represent intermediate probabilities between 0% and 100% of unreliable packets being sent. The bandwidth limits of the local and foreign hosts are taken into account to determine a sensible limit for the throttle probability above which it should not raise even in the best of conditions.
*/
//go:nosplit
func (self class) ThrottleConfigure(interval gd.Int, acceleration gd.Int, deceleration gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, interval)
	callframe.Arg(frame, acceleration)
	callframe.Arg(frame, deceleration)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetPacketPeer.Bind_throttle_configure, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the timeout parameters for a peer. The timeout parameters control how and when a peer will timeout from a failure to acknowledge reliable traffic. Timeout values are expressed in milliseconds.
The [param timeout] is a factor that, multiplied by a value based on the average round trip time, will determine the timeout limit for a reliable packet. When that limit is reached, the timeout will be doubled, and the peer will be disconnected if that limit has reached [param timeout_min]. The [param timeout_max] parameter, on the other hand, defines a fixed timeout for which any packet must be acknowledged or the peer will be dropped.
*/
//go:nosplit
func (self class) SetTimeout(timeout gd.Int, timeout_min gd.Int, timeout_max gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, timeout)
	callframe.Arg(frame, timeout_min)
	callframe.Arg(frame, timeout_max)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetPacketPeer.Bind_set_timeout, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the IP address of this peer.
*/
//go:nosplit
func (self class) GetRemoteAddress(ctx gd.Lifetime) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetPacketPeer.Bind_get_remote_address, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the remote port of this peer.
*/
//go:nosplit
func (self class) GetRemotePort() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetPacketPeer.Bind_get_remote_port, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the requested [param statistic] for this peer. See [enum PeerStatistic].
*/
//go:nosplit
func (self class) GetStatistic(statistic classdb.ENetPacketPeerPeerStatistic) gd.Float {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, statistic)
	var r_ret = callframe.Ret[gd.Float](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetPacketPeer.Bind_get_statistic, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the current peer state. See [enum PeerState].
*/
//go:nosplit
func (self class) GetState() classdb.ENetPacketPeerPeerState {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.ENetPacketPeerPeerState](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetPacketPeer.Bind_get_state, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the number of channels allocated for communication with peer.
*/
//go:nosplit
func (self class) GetChannels() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetPacketPeer.Bind_get_channels, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns [code]true[/code] if the peer is currently active (i.e. the associated [ENetConnection] is still valid).
*/
//go:nosplit
func (self class) IsActive() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.ENetPacketPeer.Bind_is_active, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self class) AsENetPacketPeer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsENetPacketPeer() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsPacketPeer() PacketPeer.GD { return *((*PacketPeer.GD)(unsafe.Pointer(&self))) }
func (self Go) AsPacketPeer() PacketPeer.Go { return *((*PacketPeer.Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsPacketPeer(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsPacketPeer(), name)
	}
}
func init() {classdb.Register("ENetPacketPeer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type PeerState = classdb.ENetPacketPeerPeerState

const (
/*The peer is disconnected.*/
	StateDisconnected PeerState = 0
/*The peer is currently attempting to connect.*/
	StateConnecting PeerState = 1
/*The peer has acknowledged the connection request.*/
	StateAcknowledgingConnect PeerState = 2
/*The peer is currently connecting.*/
	StateConnectionPending PeerState = 3
/*The peer has successfully connected, but is not ready to communicate with yet ([constant STATE_CONNECTED]).*/
	StateConnectionSucceeded PeerState = 4
/*The peer is currently connected and ready to communicate with.*/
	StateConnected PeerState = 5
/*The peer is slated to disconnect after it has no more outgoing packets to send.*/
	StateDisconnectLater PeerState = 6
/*The peer is currently disconnecting.*/
	StateDisconnecting PeerState = 7
/*The peer has acknowledged the disconnection request.*/
	StateAcknowledgingDisconnect PeerState = 8
/*The peer has lost connection, but is not considered truly disconnected (as the peer didn't acknowledge the disconnection request).*/
	StateZombie PeerState = 9
)
type PeerStatistic = classdb.ENetPacketPeerPeerStatistic

const (
/*Mean packet loss of reliable packets as a ratio with respect to the [constant PACKET_LOSS_SCALE].*/
	PeerPacketLoss PeerStatistic = 0
/*Packet loss variance.*/
	PeerPacketLossVariance PeerStatistic = 1
/*The time at which packet loss statistics were last updated (in milliseconds since the connection started). The interval for packet loss statistics updates is 10 seconds, and at least one packet must have been sent since the last statistics update.*/
	PeerPacketLossEpoch PeerStatistic = 2
/*Mean packet round trip time for reliable packets.*/
	PeerRoundTripTime PeerStatistic = 3
/*Variance of the mean round trip time.*/
	PeerRoundTripTimeVariance PeerStatistic = 4
/*Last recorded round trip time for a reliable packet.*/
	PeerLastRoundTripTime PeerStatistic = 5
/*Variance of the last trip time recorded.*/
	PeerLastRoundTripTimeVariance PeerStatistic = 6
/*The peer's current throttle status.*/
	PeerPacketThrottle PeerStatistic = 7
/*The maximum number of unreliable packets that should not be dropped. This value is always greater than or equal to [code]1[/code]. The initial value is equal to [constant PACKET_THROTTLE_SCALE].*/
	PeerPacketThrottleLimit PeerStatistic = 8
/*Internal value used to increment the packet throttle counter. The value is hardcoded to [code]7[/code] and cannot be changed. You probably want to look at [constant PEER_PACKET_THROTTLE_ACCELERATION] instead.*/
	PeerPacketThrottleCounter PeerStatistic = 9
/*The time at which throttle statistics were last updated (in milliseconds since the connection started). The interval for throttle statistics updates is [constant PEER_PACKET_THROTTLE_INTERVAL].*/
	PeerPacketThrottleEpoch PeerStatistic = 10
/*The throttle's acceleration factor. Higher values will make ENet adapt to fluctuating network conditions faster, causing unrelaible packets to be sent [i]more[/i] often. The default value is [code]2[/code].*/
	PeerPacketThrottleAcceleration PeerStatistic = 11
/*The throttle's deceleration factor. Higher values will make ENet adapt to fluctuating network conditions faster, causing unrelaible packets to be sent [i]less[/i] often. The default value is [code]2[/code].*/
	PeerPacketThrottleDeceleration PeerStatistic = 12
/*The interval over which the lowest mean round trip time should be measured for use by the throttle mechanism (in milliseconds). The default value is [code]5000[/code].*/
	PeerPacketThrottleInterval PeerStatistic = 13
)