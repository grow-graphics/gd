// Package WebRTCPeerConnectionExtension provides methods for working with WebRTCPeerConnectionExtension object instances.
package WebRTCPeerConnectionExtension

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/WebRTCPeerConnection"
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
var _ = slices.Delete[[]struct{}, struct{}]

type Instance [1]gdclass.WebRTCPeerConnectionExtension

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsWebRTCPeerConnectionExtension() Instance
}
type Interface interface {
	GetConnectionState() gdclass.WebRTCPeerConnectionConnectionState
	GetGatheringState() gdclass.WebRTCPeerConnectionGatheringState
	GetSignalingState() gdclass.WebRTCPeerConnectionSignalingState
	Initialize(p_config map[any]any) error
	CreateDataChannel(p_label string, p_config map[any]any) [1]gdclass.WebRTCDataChannel
	CreateOffer() error
	SetRemoteDescription(p_type string, p_sdp string) error
	SetLocalDescription(p_type string, p_sdp string) error
	AddIceCandidate(p_sdp_mid_name string, p_sdp_mline_index int, p_sdp_name string) error
	Poll() error
	Close()
}

// Implementation implements [Interface] with empty methods.
type Implementation = implementation

type implementation struct{}

func (self implementation) GetConnectionState() (_ gdclass.WebRTCPeerConnectionConnectionState) {
	return
}
func (self implementation) GetGatheringState() (_ gdclass.WebRTCPeerConnectionGatheringState) { return }
func (self implementation) GetSignalingState() (_ gdclass.WebRTCPeerConnectionSignalingState) { return }
func (self implementation) Initialize(p_config map[any]any) (_ error)                         { return }
func (self implementation) CreateDataChannel(p_label string, p_config map[any]any) (_ [1]gdclass.WebRTCDataChannel) {
	return
}
func (self implementation) CreateOffer() (_ error)                                     { return }
func (self implementation) SetRemoteDescription(p_type string, p_sdp string) (_ error) { return }
func (self implementation) SetLocalDescription(p_type string, p_sdp string) (_ error)  { return }
func (self implementation) AddIceCandidate(p_sdp_mid_name string, p_sdp_mline_index int, p_sdp_name string) (_ error) {
	return
}
func (self implementation) Poll() (_ error) { return }
func (self implementation) Close()          { return }
func (Instance) _get_connection_state(impl func(ptr unsafe.Pointer) gdclass.WebRTCPeerConnectionConnectionState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_gathering_state(impl func(ptr unsafe.Pointer) gdclass.WebRTCPeerConnectionGatheringState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _get_signaling_state(impl func(ptr unsafe.Pointer) gdclass.WebRTCPeerConnectionSignalingState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}
func (Instance) _initialize(impl func(ptr unsafe.Pointer, p_config map[any]any) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_config = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalDictionary(p_config))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, gd.DictionaryAs[map[any]any](p_config))
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _create_data_channel(impl func(ptr unsafe.Pointer, p_label string, p_config map[any]any) [1]gdclass.WebRTCDataChannel) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_label = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(p_label))
		var p_config = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalDictionary(p_config))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_label.String(), gd.DictionaryAs[map[any]any](p_config))
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _create_offer(impl func(ptr unsafe.Pointer) error) (cb gd.ExtensionClassCallVirtualFunc) {
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
func (Instance) _set_remote_description(impl func(ptr unsafe.Pointer, p_type string, p_sdp string) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_type = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(p_type))
		var p_sdp = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(p_sdp))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_type.String(), p_sdp.String())
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _set_local_description(impl func(ptr unsafe.Pointer, p_type string, p_sdp string) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_type = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(p_type))
		var p_sdp = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(p_sdp))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_type.String(), p_sdp.String())
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}
func (Instance) _add_ice_candidate(impl func(ptr unsafe.Pointer, p_sdp_mid_name string, p_sdp_mline_index int, p_sdp_name string) error) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_sdp_mid_name = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(p_sdp_mid_name))
		var p_sdp_mline_index = gd.UnsafeGet[int64](p_args, 1)
		var p_sdp_name = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(p_sdp_name))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_sdp_mid_name.String(), int(p_sdp_mline_index), p_sdp_name.String())
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(Error.New(ret))

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
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

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.WebRTCPeerConnectionExtension

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("WebRTCPeerConnectionExtension"))
	casted := Instance{*(*gdclass.WebRTCPeerConnectionExtension)(unsafe.Pointer(&object))}
	casted.AsRefCounted()[0].Reference()
	return casted
}

func (class) _get_connection_state(impl func(ptr unsafe.Pointer) gdclass.WebRTCPeerConnectionConnectionState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_gathering_state(impl func(ptr unsafe.Pointer) gdclass.WebRTCPeerConnectionGatheringState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _get_signaling_state(impl func(ptr unsafe.Pointer) gdclass.WebRTCPeerConnectionSignalingState) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self)
		gd.UnsafeSet(p_back, ret)
	}
}

func (class) _initialize(impl func(ptr unsafe.Pointer, p_config Dictionary.Any) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_config = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalDictionary(p_config))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_config)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _create_data_channel(impl func(ptr unsafe.Pointer, p_label String.Readable, p_config Dictionary.Any) [1]gdclass.WebRTCDataChannel) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_label = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(p_label))
		var p_config = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalDictionary(p_config))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_label, p_config)
		ptr, ok := pointers.End(ret[0])

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _create_offer(impl func(ptr unsafe.Pointer) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
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

func (class) _set_remote_description(impl func(ptr unsafe.Pointer, p_type String.Readable, p_sdp String.Readable) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_type = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(p_type))
		var p_sdp = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(p_sdp))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_type, p_sdp)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _set_local_description(impl func(ptr unsafe.Pointer, p_type String.Readable, p_sdp String.Readable) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_type = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(p_type))
		var p_sdp = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 1))))
		defer pointers.End(gd.InternalString(p_sdp))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_type, p_sdp)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
	}
}

func (class) _add_ice_candidate(impl func(ptr unsafe.Pointer, p_sdp_mid_name String.Readable, p_sdp_mline_index int64, p_sdp_name String.Readable) Error.Code) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args gd.Address, p_back gd.Address) {
		var p_sdp_mid_name = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 0))))
		defer pointers.End(gd.InternalString(p_sdp_mid_name))
		var p_sdp_mline_index = gd.UnsafeGet[int64](p_args, 1)
		var p_sdp_name = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](gd.UnsafeGet[[1]gd.EnginePointer](p_args, 2))))
		defer pointers.End(gd.InternalString(p_sdp_name))
		self := reflect.ValueOf(class).UnsafePointer()
		ret := impl(self, p_sdp_mid_name, p_sdp_mline_index, p_sdp_name)
		ptr, ok := func(e Error.Code) (int64, bool) { return int64(e), true }(ret)

		if !ok {
			return
		}
		gd.UnsafeSet(p_back, ptr)
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

func (self class) AsWebRTCPeerConnectionExtension() Advanced {
	return *((*Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsWebRTCPeerConnectionExtension() Instance {
	return *((*Instance)(unsafe.Pointer(&self)))
}
func (self class) AsWebRTCPeerConnection() WebRTCPeerConnection.Advanced {
	return *((*WebRTCPeerConnection.Advanced)(unsafe.Pointer(&self)))
}
func (self Instance) AsWebRTCPeerConnection() WebRTCPeerConnection.Instance {
	return *((*WebRTCPeerConnection.Instance)(unsafe.Pointer(&self)))
}
func (self class) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}
func (self Instance) AsRefCounted() [1]gd.RefCounted {
	return *((*[1]gd.RefCounted)(unsafe.Pointer(&self)))
}

func (self class) Virtual(name string) reflect.Value {
	switch name {
	case "_get_connection_state":
		return reflect.ValueOf(self._get_connection_state)
	case "_get_gathering_state":
		return reflect.ValueOf(self._get_gathering_state)
	case "_get_signaling_state":
		return reflect.ValueOf(self._get_signaling_state)
	case "_initialize":
		return reflect.ValueOf(self._initialize)
	case "_create_data_channel":
		return reflect.ValueOf(self._create_data_channel)
	case "_create_offer":
		return reflect.ValueOf(self._create_offer)
	case "_set_remote_description":
		return reflect.ValueOf(self._set_remote_description)
	case "_set_local_description":
		return reflect.ValueOf(self._set_local_description)
	case "_add_ice_candidate":
		return reflect.ValueOf(self._add_ice_candidate)
	case "_poll":
		return reflect.ValueOf(self._poll)
	case "_close":
		return reflect.ValueOf(self._close)
	default:
		return gd.VirtualByName(WebRTCPeerConnection.Advanced(self.AsWebRTCPeerConnection()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	case "_get_connection_state":
		return reflect.ValueOf(self._get_connection_state)
	case "_get_gathering_state":
		return reflect.ValueOf(self._get_gathering_state)
	case "_get_signaling_state":
		return reflect.ValueOf(self._get_signaling_state)
	case "_initialize":
		return reflect.ValueOf(self._initialize)
	case "_create_data_channel":
		return reflect.ValueOf(self._create_data_channel)
	case "_create_offer":
		return reflect.ValueOf(self._create_offer)
	case "_set_remote_description":
		return reflect.ValueOf(self._set_remote_description)
	case "_set_local_description":
		return reflect.ValueOf(self._set_local_description)
	case "_add_ice_candidate":
		return reflect.ValueOf(self._add_ice_candidate)
	case "_poll":
		return reflect.ValueOf(self._poll)
	case "_close":
		return reflect.ValueOf(self._close)
	default:
		return gd.VirtualByName(WebRTCPeerConnection.Instance(self.AsWebRTCPeerConnection()), name)
	}
}
func init() {
	gdclass.Register("WebRTCPeerConnectionExtension", func(ptr gd.Object) any {
		return [1]gdclass.WebRTCPeerConnectionExtension{*(*gdclass.WebRTCPeerConnectionExtension)(unsafe.Pointer(&ptr))}
	})
}
