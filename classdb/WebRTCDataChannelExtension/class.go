// Code generated by the generate package DO NOT EDIT

// Package WebRTCDataChannelExtension provides methods for working with WebRTCDataChannelExtension object instances.
package WebRTCDataChannelExtension

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
import "graphics.gd/classdb/PacketPeer"
import "graphics.gd/classdb/WebRTCDataChannel"
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
type Instance [1]gdclass.WebRTCDataChannelExtension

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsWebRTCDataChannelExtension() Instance
}
type Interface interface {
	GetPacket(r_buffer unsafe.Pointer, r_buffer_size *int32) error
	PutPacket(p_buffer unsafe.Pointer, p_buffer_size int) error
	GetAvailablePacketCount() int
	GetMaxPacketSize() int
	Poll() error
	Close()
	SetWriteMode(p_write_mode WebRTCDataChannel.WriteMode)
	GetWriteMode() WebRTCDataChannel.WriteMode
	WasStringPacket() bool
	GetReadyState() WebRTCDataChannel.ChannelState
	GetLabel() string
	IsOrdered() bool
	GetId() int
	GetMaxPacketLifeTime() int
	GetMaxRetransmits() int
	GetProtocol() string
	IsNegotiated() bool
	GetBufferedAmount() int
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetPacket(r_buffer unsafe.Pointer, r_buffer_size *int32) (_ error) { return }
func (self implementation) PutPacket(p_buffer unsafe.Pointer, p_buffer_size int) (_ error)    { return }
func (self implementation) GetAvailablePacketCount() (_ int)                                  { return }
func (self implementation) GetMaxPacketSize() (_ int)                                         { return }
func (self implementation) Poll() (_ error)                                                   { return }
func (self implementation) Close()                                                            { return }
func (self implementation) SetWriteMode(p_write_mode WebRTCDataChannel.WriteMode)             { return }
func (self implementation) GetWriteMode() (_ WebRTCDataChannel.WriteMode)                     { return }
func (self implementation) WasStringPacket() (_ bool)                                         { return }
func (self implementation) GetReadyState() (_ WebRTCDataChannel.ChannelState)                 { return }
func (self implementation) GetLabel() (_ string)                                              { return }
func (self implementation) IsOrdered() (_ bool)                                               { return }
func (self implementation) GetId() (_ int)                                                    { return }
func (self implementation) GetMaxPacketLifeTime() (_ int)                                     { return }
func (self implementation) GetMaxRetransmits() (_ int)                                        { return }
func (self implementation) GetProtocol() (_ string)                                           { return }
func (self implementation) IsNegotiated() (_ bool)                                            { return }
func (self implementation) GetBufferedAmount() (_ int)                                        { return }
func (Instance) _get_packet(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_buffer_size *int32) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var r_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var r_buffer_size = gd.UnsafeGet[*int32](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, r_buffer, r_buffer_size)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _put_packet(impl func(ptr unsafe.Pointer, p_buffer unsafe.Pointer, p_buffer_size int) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var p_buffer_size = gd.UnsafeGet[int64](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_buffer, int(p_buffer_size))
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _get_available_packet_count(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _get_max_packet_size(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _poll(impl func(ptr unsafe.Pointer) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _close(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}
func (Instance) _set_write_mode(impl func(ptr unsafe.Pointer, p_write_mode WebRTCDataChannel.WriteMode)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_write_mode = gd.UnsafeGet[WebRTCDataChannel.WriteMode](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, p_write_mode)
	}
}
func (Instance) _get_write_mode(impl func(ptr unsafe.Pointer) WebRTCDataChannel.WriteMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _was_string_packet(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_ready_state(impl func(ptr unsafe.Pointer) WebRTCDataChannel.ChannelState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_label(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _is_ordered(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_id(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _get_max_packet_life_time(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _get_max_retransmits(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, int64(ret))
	}
}
func (Instance) _get_protocol(impl func(ptr unsafe.Pointer) string) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(String.New(ret)))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _is_negotiated(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_buffered_amount(impl func(ptr unsafe.Pointer) int) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, int64(ret))
	}
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.WebRTCDataChannelExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WebRTCDataChannelExtension"))
	casted := Instance{*(*gdclass.WebRTCDataChannelExtension)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (class) _get_packet(impl func(ptr unsafe.Pointer, r_buffer unsafe.Pointer, r_buffer_size *int32) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var r_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var r_buffer_size = gd.UnsafeGet[*int32](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, r_buffer, r_buffer_size)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _put_packet(impl func(ptr unsafe.Pointer, p_buffer unsafe.Pointer, p_buffer_size int64) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_buffer = gd.UnsafeGet[unsafe.Pointer](p_args, 0)
		var p_buffer_size = gd.UnsafeGet[int64](p_args, 1)
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_buffer, p_buffer_size)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _get_available_packet_count(impl func(ptr unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_max_packet_size(impl func(ptr unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _poll(impl func(ptr unsafe.Pointer) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _close(impl func(ptr unsafe.Pointer)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self)
	}
}

func (class) _set_write_mode(impl func(ptr unsafe.Pointer, p_write_mode WebRTCDataChannel.WriteMode)) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_write_mode = gd.UnsafeGet[WebRTCDataChannel.WriteMode](p_args, 0)
		self := reflect.ValueOf(class).UnsafePointer()
		impl(self, p_write_mode)
	}
}

func (class) _get_write_mode(impl func(ptr unsafe.Pointer) WebRTCDataChannel.WriteMode) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _was_string_packet(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_ready_state(impl func(ptr unsafe.Pointer) WebRTCDataChannel.ChannelState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_label(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _is_ordered(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_id(impl func(ptr unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_max_packet_life_time(impl func(ptr unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_max_retransmits(impl func(ptr unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_protocol(impl func(ptr unsafe.Pointer) String.Readable) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		ptr, ok := pointers.End(gd.InternalString(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _is_negotiated(impl func(ptr unsafe.Pointer) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_buffered_amount(impl func(ptr unsafe.Pointer) int64) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (self class) AsWebRTCDataChannelExtension() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsWebRTCDataChannelExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsWebRTCDataChannelExtension() Instance {
	return self.Super().AsWebRTCDataChannelExtension()
}
func (self class) AsWebRTCDataChannel() WebRTCDataChannel.Advanced {
	return *((*WebRTCDataChannel.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsWebRTCDataChannel() WebRTCDataChannel.Instance {
	return self.Super().AsWebRTCDataChannel()
}
func (self Instance) AsWebRTCDataChannel() WebRTCDataChannel.Instance {
	return *((*WebRTCDataChannel.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsPacketPeer() PacketPeer.Advanced {
	return *((*PacketPeer.Advanced)(unsafe.Pointer(&self)))
}
func (self *Extension[T]) AsPacketPeer() PacketPeer.Instance { return self.Super().AsPacketPeer() }
func (self Instance) AsPacketPeer() PacketPeer.Instance {
	return *((*PacketPeer.Instance)(unsafe.Pointer(&self)))
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
	case "_get_packet":
		return reflect.ValueOf(self._get_packet)
	case "_put_packet":
		return reflect.ValueOf(self._put_packet)
	case "_get_available_packet_count":
		return reflect.ValueOf(self._get_available_packet_count)
	case "_get_max_packet_size":
		return reflect.ValueOf(self._get_max_packet_size)
	case "_poll":
		return reflect.ValueOf(self._poll)
	case "_close":
		return reflect.ValueOf(self._close)
	case "_set_write_mode":
		return reflect.ValueOf(self._set_write_mode)
	case "_get_write_mode":
		return reflect.ValueOf(self._get_write_mode)
	case "_was_string_packet":
		return reflect.ValueOf(self._was_string_packet)
	case "_get_ready_state":
		return reflect.ValueOf(self._get_ready_state)
	case "_get_label":
		return reflect.ValueOf(self._get_label)
	case "_is_ordered":
		return reflect.ValueOf(self._is_ordered)
	case "_get_id":
		return reflect.ValueOf(self._get_id)
	case "_get_max_packet_life_time":
		return reflect.ValueOf(self._get_max_packet_life_time)
	case "_get_max_retransmits":
		return reflect.ValueOf(self._get_max_retransmits)
	case "_get_protocol":
		return reflect.ValueOf(self._get_protocol)
	case "_is_negotiated":
		return reflect.ValueOf(self._is_negotiated)
	case "_get_buffered_amount":
		return reflect.ValueOf(self._get_buffered_amount)
	default:
		return gd.VirtualByName(WebRTCDataChannel.Advanced(self.AsWebRTCDataChannel()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_packet":
		return reflect.ValueOf(self._get_packet)
	case "_put_packet":
		return reflect.ValueOf(self._put_packet)
	case "_get_available_packet_count":
		return reflect.ValueOf(self._get_available_packet_count)
	case "_get_max_packet_size":
		return reflect.ValueOf(self._get_max_packet_size)
	case "_poll":
		return reflect.ValueOf(self._poll)
	case "_close":
		return reflect.ValueOf(self._close)
	case "_set_write_mode":
		return reflect.ValueOf(self._set_write_mode)
	case "_get_write_mode":
		return reflect.ValueOf(self._get_write_mode)
	case "_was_string_packet":
		return reflect.ValueOf(self._was_string_packet)
	case "_get_ready_state":
		return reflect.ValueOf(self._get_ready_state)
	case "_get_label":
		return reflect.ValueOf(self._get_label)
	case "_is_ordered":
		return reflect.ValueOf(self._is_ordered)
	case "_get_id":
		return reflect.ValueOf(self._get_id)
	case "_get_max_packet_life_time":
		return reflect.ValueOf(self._get_max_packet_life_time)
	case "_get_max_retransmits":
		return reflect.ValueOf(self._get_max_retransmits)
	case "_get_protocol":
		return reflect.ValueOf(self._get_protocol)
	case "_is_negotiated":
		return reflect.ValueOf(self._is_negotiated)
	case "_get_buffered_amount":
		return reflect.ValueOf(self._get_buffered_amount)
	default:
		return gd.VirtualByName(WebRTCDataChannel.Instance(self.AsWebRTCDataChannel()), name)
	}
}
func init() {
	gdclass.Register("WebRTCDataChannelExtension", func(ptr gd.Object) any {
		return [1]gdclass.WebRTCDataChannelExtension{*(*gdclass.WebRTCDataChannelExtension)(unsafe.Pointer(&ptr))}
	})
}
