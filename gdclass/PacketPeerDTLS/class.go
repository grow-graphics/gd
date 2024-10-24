package PacketPeerDTLS

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
This class represents a DTLS peer connection. It can be used to connect to a DTLS server, and is returned by [method DTLSServer.take_connection].
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
[b]Warning:[/b] TLS certificate revocation and certificate pinning are currently not supported. Revoked certificates are accepted as long as they are otherwise valid. If this is a concern, you may want to use automatically managed certificates with a short validity period.

*/
type Go [1]classdb.PacketPeerDTLS

/*
Poll the connection to check for incoming packets. Call this frequently to update the status and keep the connection working.
*/
func (self Go) Poll() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Poll()
}

/*
Connects a [param packet_peer] beginning the DTLS handshake using the underlying [PacketPeerUDP] which must be connected (see [method PacketPeerUDP.connect_to_host]). You can optionally specify the [param client_options] to be used while verifying the TLS connections. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
*/
func (self Go) ConnectToPeer(packet_peer gdclass.PacketPeerUDP, hostname string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).ConnectToPeer(packet_peer, gc.String(hostname), ([1]gdclass.TLSOptions{}[0])))
}

/*
Returns the status of the connection. See [enum Status] for values.
*/
func (self Go) GetStatus() classdb.PacketPeerDTLSStatus {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.PacketPeerDTLSStatus(class(self).GetStatus())
}

/*
Disconnects this peer, terminating the DTLS session.
*/
func (self Go) DisconnectFromPeer() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).DisconnectFromPeer()
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.PacketPeerDTLS
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("PacketPeerDTLS"))
	return *(*Go)(unsafe.Pointer(&object))
}

/*
Poll the connection to check for incoming packets. Call this frequently to update the status and keep the connection working.
*/
//go:nosplit
func (self class) Poll()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerDTLS.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Connects a [param packet_peer] beginning the DTLS handshake using the underlying [PacketPeerUDP] which must be connected (see [method PacketPeerUDP.connect_to_host]). You can optionally specify the [param client_options] to be used while verifying the TLS connections. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
*/
//go:nosplit
func (self class) ConnectToPeer(packet_peer gdclass.PacketPeerUDP, hostname gd.String, client_options gdclass.TLSOptions) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(packet_peer[0].AsPointer())[0])
	callframe.Arg(frame, mmm.Get(hostname))
	callframe.Arg(frame, mmm.Get(client_options[0].AsPointer())[0])
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerDTLS.Bind_connect_to_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the status of the connection. See [enum Status] for values.
*/
//go:nosplit
func (self class) GetStatus() classdb.PacketPeerDTLSStatus {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.PacketPeerDTLSStatus](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerDTLS.Bind_get_status, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Disconnects this peer, terminating the DTLS session.
*/
//go:nosplit
func (self class) DisconnectFromPeer()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.PacketPeerDTLS.Bind_disconnect_from_peer, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
func (self class) AsPacketPeerDTLS() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsPacketPeerDTLS() Go { return *((*Go)(unsafe.Pointer(&self))) }
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
func init() {classdb.Register("PacketPeerDTLS", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type Status = classdb.PacketPeerDTLSStatus

const (
/*A status representing a [PacketPeerDTLS] that is disconnected.*/
	StatusDisconnected Status = 0
/*A status representing a [PacketPeerDTLS] that is currently performing the handshake with a remote peer.*/
	StatusHandshaking Status = 1
/*A status representing a [PacketPeerDTLS] that is connected to a remote peer.*/
	StatusConnected Status = 2
/*A status representing a [PacketPeerDTLS] in a generic error state.*/
	StatusError Status = 3
/*An error status that shows a mismatch in the DTLS certificate domain presented by the host and the domain requested for validation.*/
	StatusErrorHostnameMismatch Status = 4
)