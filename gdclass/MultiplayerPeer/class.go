package MultiplayerPeer

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
Manages the connection with one or more remote peers acting as server or client and assigning unique IDs to each of them. See also [MultiplayerAPI].
[b]Note:[/b] The [MultiplayerAPI] protocol is an implementation detail and isn't meant to be used by non-Godot servers. It may change without notice.
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.

*/
type Go [1]classdb.MultiplayerPeer

/*
Sets the peer to which packets will be sent.
The [param id] can be one of: [constant TARGET_PEER_BROADCAST] to send to all connected peers, [constant TARGET_PEER_SERVER] to send to the peer acting as server, a valid peer ID to send to that specific peer, a negative peer ID to send to all peers except that one. By default, the target peer is [constant TARGET_PEER_BROADCAST].
*/
func (self Go) SetTargetPeer(id int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTargetPeer(gd.Int(id))
}

/*
Returns the ID of the [MultiplayerPeer] who sent the next available packet. See [method PacketPeer.get_available_packet_count].
*/
func (self Go) GetPacketPeer() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetPacketPeer()))
}

/*
Returns the channel over which the next available packet was received. See [method PacketPeer.get_available_packet_count].
*/
func (self Go) GetPacketChannel() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetPacketChannel()))
}

/*
Returns the transfer mode the remote peer used to send the next available packet. See [method PacketPeer.get_available_packet_count].
*/
func (self Go) GetPacketMode() classdb.MultiplayerPeerTransferMode {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.MultiplayerPeerTransferMode(class(self).GetPacketMode())
}

/*
Waits up to 1 second to receive a new network event.
*/
func (self Go) Poll() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Poll()
}

/*
Immediately close the multiplayer peer returning to the state [constant CONNECTION_DISCONNECTED]. Connected peers will be dropped without emitting [signal peer_disconnected].
*/
func (self Go) Close() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Close()
}

/*
Disconnects the given [param peer] from this host. If [param force] is [code]true[/code] the [signal peer_disconnected] signal will not be emitted for this peer.
*/
func (self Go) DisconnectPeer(peer int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DisconnectPeer(gd.Int(peer), false)
}

/*
Returns the current state of the connection. See [enum ConnectionStatus].
*/
func (self Go) GetConnectionStatus() classdb.MultiplayerPeerConnectionStatus {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.MultiplayerPeerConnectionStatus(class(self).GetConnectionStatus())
}

/*
Returns the ID of this [MultiplayerPeer].
*/
func (self Go) GetUniqueId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetUniqueId()))
}

/*
Returns a randomly generated integer that can be used as a network unique ID.
*/
func (self Go) GenerateUniqueId() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GenerateUniqueId()))
}

/*
Returns true if the server can act as a relay in the current configuration (i.e. if the higher level [MultiplayerAPI] should notify connected clients of other peers, and implement a relay protocol to allow communication between them).
*/
func (self Go) IsServerRelaySupported() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsServerRelaySupported())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.MultiplayerPeer
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("MultiplayerPeer"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) RefuseNewConnections() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsRefusingNewConnections())
}

func (self Go) SetRefuseNewConnections(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetRefuseNewConnections(value)
}

func (self Go) TransferMode() classdb.MultiplayerPeerTransferMode {
	gc := gd.GarbageCollector(); _ = gc
		return classdb.MultiplayerPeerTransferMode(class(self).GetTransferMode())
}

func (self Go) SetTransferMode(value classdb.MultiplayerPeerTransferMode) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTransferMode(value)
}

func (self Go) TransferChannel() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetTransferChannel()))
}

func (self Go) SetTransferChannel(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetTransferChannel(gd.Int(value))
}

//go:nosplit
func (self class) SetTransferChannel(channel gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, channel)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_set_transfer_channel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTransferChannel() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_get_transfer_channel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetTransferMode(mode classdb.MultiplayerPeerTransferMode)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mode)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_set_transfer_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetTransferMode() classdb.MultiplayerPeerTransferMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.MultiplayerPeerTransferMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_get_transfer_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the peer to which packets will be sent.
The [param id] can be one of: [constant TARGET_PEER_BROADCAST] to send to all connected peers, [constant TARGET_PEER_SERVER] to send to the peer acting as server, a valid peer ID to send to that specific peer, a negative peer ID to send to all peers except that one. By default, the target peer is [constant TARGET_PEER_BROADCAST].
*/
//go:nosplit
func (self class) SetTargetPeer(id gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, id)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_set_target_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the ID of the [MultiplayerPeer] who sent the next available packet. See [method PacketPeer.get_available_packet_count].
*/
//go:nosplit
func (self class) GetPacketPeer() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_get_packet_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the channel over which the next available packet was received. See [method PacketPeer.get_available_packet_count].
*/
//go:nosplit
func (self class) GetPacketChannel() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_get_packet_channel, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the transfer mode the remote peer used to send the next available packet. See [method PacketPeer.get_available_packet_count].
*/
//go:nosplit
func (self class) GetPacketMode() classdb.MultiplayerPeerTransferMode {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.MultiplayerPeerTransferMode](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_get_packet_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Waits up to 1 second to receive a new network event.
*/
//go:nosplit
func (self class) Poll()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Immediately close the multiplayer peer returning to the state [constant CONNECTION_DISCONNECTED]. Connected peers will be dropped without emitting [signal peer_disconnected].
*/
//go:nosplit
func (self class) Close()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_close, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Disconnects the given [param peer] from this host. If [param force] is [code]true[/code] the [signal peer_disconnected] signal will not be emitted for this peer.
*/
//go:nosplit
func (self class) DisconnectPeer(peer gd.Int, force bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, peer)
	callframe.Arg(frame, force)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_disconnect_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Returns the current state of the connection. See [enum ConnectionStatus].
*/
//go:nosplit
func (self class) GetConnectionStatus() classdb.MultiplayerPeerConnectionStatus {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.MultiplayerPeerConnectionStatus](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_get_connection_status, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the ID of this [MultiplayerPeer].
*/
//go:nosplit
func (self class) GetUniqueId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_get_unique_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a randomly generated integer that can be used as a network unique ID.
*/
//go:nosplit
func (self class) GenerateUniqueId() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_generate_unique_id, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetRefuseNewConnections(enable bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enable)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_set_refuse_new_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsRefusingNewConnections() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_is_refusing_new_connections, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns true if the server can act as a relay in the current configuration (i.e. if the higher level [MultiplayerAPI] should notify connected clients of other peers, and implement a relay protocol to allow communication between them).
*/
//go:nosplit
func (self class) IsServerRelaySupported() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.MultiplayerPeer.Bind_is_server_relay_supported, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
func (self Go) OnPeerConnected(cb func(id int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("peer_connected"), gc.Callable(cb), 0)
}


func (self Go) OnPeerDisconnected(cb func(id int)) {
	gc := gd.GarbageCollector(); _ = gc
	self[0].AsObject().Connect(gc.StringName("peer_disconnected"), gc.Callable(cb), 0)
}


func (self class) AsMultiplayerPeer() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsMultiplayerPeer() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("MultiplayerPeer", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type ConnectionStatus = classdb.MultiplayerPeerConnectionStatus

const (
/*The MultiplayerPeer is disconnected.*/
	ConnectionDisconnected ConnectionStatus = 0
/*The MultiplayerPeer is currently connecting to a server.*/
	ConnectionConnecting ConnectionStatus = 1
/*This MultiplayerPeer is connected.*/
	ConnectionConnected ConnectionStatus = 2
)
type TransferMode = classdb.MultiplayerPeerTransferMode

const (
/*Packets are not acknowledged, no resend attempts are made for lost packets. Packets may arrive in any order. Potentially faster than [constant TRANSFER_MODE_UNRELIABLE_ORDERED]. Use for non-critical data, and always consider whether the order matters.*/
	TransferModeUnreliable TransferMode = 0
/*Packets are not acknowledged, no resend attempts are made for lost packets. Packets are received in the order they were sent in. Potentially faster than [constant TRANSFER_MODE_RELIABLE]. Use for non-critical data or data that would be outdated if received late due to resend attempt(s) anyway, for example movement and positional data.*/
	TransferModeUnreliableOrdered TransferMode = 1
/*Packets must be received and resend attempts should be made until the packets are acknowledged. Packets must be received in the order they were sent in. Most reliable transfer mode, but potentially the slowest due to the overhead. Use for critical data that must be transmitted and arrive in order, for example an ability being triggered or a chat message. Consider carefully if the information really is critical, and use sparingly.*/
	TransferModeReliable TransferMode = 2
)