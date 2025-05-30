// Code generated by the generate package DO NOT EDIT

// Package JSONRPC provides methods for working with JSONRPC object instances.
package JSONRPC

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
[url=https://www.jsonrpc.org/]JSON-RPC[/url] is a standard which wraps a method call in a [JSON] object. The object has a particular structure and identifies which method is called, the parameters to that function, and carries an ID to keep track of responses. This class implements that standard on top of [Dictionary]; you will have to convert between a [Dictionary] and [JSON] with other functions.
*/
type Instance [1]gdclass.JSONRPC

func (self Instance) ID() ID { return ID(Object.Instance(self.AsObject()).ID()) }

type Expanded [1]gdclass.JSONRPC

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsJSONRPC() Instance
}

func (self Instance) SetScope(scope string, target Object.Instance) { //gd:JSONRPC.set_scope
	Advanced(self).SetScope(String.New(scope), target)
}

/*
Given a Dictionary which takes the form of a JSON-RPC request: unpack the request and run it. Methods are resolved by looking at the field called "method" and looking for an equivalently named function in the JSONRPC object. If one is found that method is called.
To add new supported methods extend the JSONRPC class and call [method process_action] on your subclass.
[param action]: The action to be run, as a Dictionary in the form of a JSON-RPC request or notification.
*/
func (self Instance) ProcessAction(action any) any { //gd:JSONRPC.process_action
	return any(Advanced(self).ProcessAction(variant.New(action), false).Interface())
}

/*
Given a Dictionary which takes the form of a JSON-RPC request: unpack the request and run it. Methods are resolved by looking at the field called "method" and looking for an equivalently named function in the JSONRPC object. If one is found that method is called.
To add new supported methods extend the JSONRPC class and call [method process_action] on your subclass.
[param action]: The action to be run, as a Dictionary in the form of a JSON-RPC request or notification.
*/
func (self Expanded) ProcessAction(action any, recurse bool) any { //gd:JSONRPC.process_action
	return any(Advanced(self).ProcessAction(variant.New(action), recurse).Interface())
}
func (self Instance) ProcessString(action string) string { //gd:JSONRPC.process_string
	return string(Advanced(self).ProcessString(String.New(action)).String())
}

/*
Returns a dictionary in the form of a JSON-RPC request. Requests are sent to a server with the expectation of a response. The ID field is used for the server to specify which exact request it is responding to.
- [param method]: Name of the method being called.
- [param params]: An array or dictionary of parameters being passed to the method.
- [param id]: Uniquely identifies this request. The server is expected to send a response with the same ID.
*/
func (self Instance) MakeRequest(method string, params any, id any) Request { //gd:JSONRPC.make_request
	return Request(gd.DictionaryAs[Request](Advanced(self).MakeRequest(String.New(method), variant.New(params), variant.New(id))))
}

/*
When a server has received and processed a request, it is expected to send a response. If you did not want a response then you need to have sent a Notification instead.
- [param result]: The return value of the function which was called.
- [param id]: The ID of the request this response is targeted to.
*/
func (self Instance) MakeResponse(result any, id any) Response { //gd:JSONRPC.make_response
	return Response(gd.DictionaryAs[Response](Advanced(self).MakeResponse(variant.New(result), variant.New(id))))
}

/*
Returns a dictionary in the form of a JSON-RPC notification. Notifications are one-shot messages which do not expect a response.
- [param method]: Name of the method being called.
- [param params]: An array or dictionary of parameters being passed to the method.
*/
func (self Instance) MakeNotification(method string, params any) Notification { //gd:JSONRPC.make_notification
	return Notification(gd.DictionaryAs[Notification](Advanced(self).MakeNotification(String.New(method), variant.New(params))))
}

/*
Creates a response which indicates a previous reply has failed in some way.
- [param code]: The error code corresponding to what kind of error this is. See the [enum ErrorCode] constants.
- [param message]: A custom message about this error.
- [param id]: The request this error is a response to.
*/
func (self Instance) MakeResponseError(code int, message string) ResponseError { //gd:JSONRPC.make_response_error
	return ResponseError(gd.DictionaryAs[ResponseError](Advanced(self).MakeResponseError(int64(code), String.New(message), variant.New([1]any{}[0]))))
}

/*
Creates a response which indicates a previous reply has failed in some way.
- [param code]: The error code corresponding to what kind of error this is. See the [enum ErrorCode] constants.
- [param message]: A custom message about this error.
- [param id]: The request this error is a response to.
*/
func (self Expanded) MakeResponseError(code int, message string, id any) ResponseError { //gd:JSONRPC.make_response_error
	return ResponseError(gd.DictionaryAs[ResponseError](Advanced(self).MakeResponseError(int64(code), String.New(message), variant.New(id))))
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.JSONRPC

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self *Extension[T]) AsObject() [1]gd.Object    { return self.Super().AsObject() }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("JSONRPC"))
	casted := Instance{*(*gdclass.JSONRPC)(unsafe.Pointer(&object))}
	return casted
}

//go:nosplit
func (self class) SetScope(scope String.Readable, target [1]gd.Object) { //gd:JSONRPC.set_scope
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(scope)))
	callframe.Arg(frame, pointers.Get(target[0])[0])
	var r_ret = callframe.Nil
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSONRPC.Bind_set_scope, self.AsObject(), frame.Array(0), r_ret.Addr())
	frame.Free()
}

/*
Given a Dictionary which takes the form of a JSON-RPC request: unpack the request and run it. Methods are resolved by looking at the field called "method" and looking for an equivalently named function in the JSONRPC object. If one is found that method is called.
To add new supported methods extend the JSONRPC class and call [method process_action] on your subclass.
[param action]: The action to be run, as a Dictionary in the form of a JSON-RPC request or notification.
*/
//go:nosplit
func (self class) ProcessAction(action variant.Any, recurse bool) variant.Any { //gd:JSONRPC.process_action
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(action)))
	callframe.Arg(frame, recurse)
	var r_ret = callframe.Ret[[3]uint64](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSONRPC.Bind_process_action, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = variant.Implementation(gd.VariantProxy{}, pointers.Pack(pointers.New[gd.Variant](r_ret.Get())))
	frame.Free()
	return ret
}

//go:nosplit
func (self class) ProcessString(action String.Readable) String.Readable { //gd:JSONRPC.process_string
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(action)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSONRPC.Bind_process_string, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = String.Via(gd.StringProxy{}, pointers.Pack(pointers.New[gd.String](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns a dictionary in the form of a JSON-RPC request. Requests are sent to a server with the expectation of a response. The ID field is used for the server to specify which exact request it is responding to.
- [param method]: Name of the method being called.
- [param params]: An array or dictionary of parameters being passed to the method.
- [param id]: Uniquely identifies this request. The server is expected to send a response with the same ID.
*/
//go:nosplit
func (self class) MakeRequest(method String.Readable, params variant.Any, id variant.Any) Dictionary.Any { //gd:JSONRPC.make_request
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(method)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(params)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(id)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSONRPC.Bind_make_request, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
When a server has received and processed a request, it is expected to send a response. If you did not want a response then you need to have sent a Notification instead.
- [param result]: The return value of the function which was called.
- [param id]: The ID of the request this response is targeted to.
*/
//go:nosplit
func (self class) MakeResponse(result variant.Any, id variant.Any) Dictionary.Any { //gd:JSONRPC.make_response
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(result)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(id)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSONRPC.Bind_make_response, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Returns a dictionary in the form of a JSON-RPC notification. Notifications are one-shot messages which do not expect a response.
- [param method]: Name of the method being called.
- [param params]: An array or dictionary of parameters being passed to the method.
*/
//go:nosplit
func (self class) MakeNotification(method String.Readable, params variant.Any) Dictionary.Any { //gd:JSONRPC.make_notification
	var frame = callframe.New()
	callframe.Arg(frame, pointers.Get(gd.InternalString(method)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(params)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSONRPC.Bind_make_notification, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}

/*
Creates a response which indicates a previous reply has failed in some way.
- [param code]: The error code corresponding to what kind of error this is. See the [enum ErrorCode] constants.
- [param message]: A custom message about this error.
- [param id]: The request this error is a response to.
*/
//go:nosplit
func (self class) MakeResponseError(code int64, message String.Readable, id variant.Any) Dictionary.Any { //gd:JSONRPC.make_response_error
	var frame = callframe.New()
	callframe.Arg(frame, code)
	callframe.Arg(frame, pointers.Get(gd.InternalString(message)))
	callframe.Arg(frame, pointers.Get(gd.InternalVariant(id)))
	var r_ret = callframe.Ret[[1]gd.EnginePointer](frame)
	gd.Global.Object.MethodBindPointerCall(gd.Global.Methods.JSONRPC.Bind_make_response_error, self.AsObject(), frame.Array(0), r_ret.Addr())
	var ret = Dictionary.Through(gd.DictionaryProxy[variant.Any, variant.Any]{}, pointers.Pack(pointers.New[gd.Dictionary](r_ret.Get())))
	frame.Free()
	return ret
}
func (self class) AsJSONRPC() Advanced         { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsJSONRPC() Instance      { return *((*Instance)(unsafe.Pointer(&self))) }
func (self *Extension[T]) AsJSONRPC() Instance { return self.Super().AsJSONRPC() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Advanced(self.AsObject()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Object.Instance(self.AsObject()), name)
	}
}
func init() {
	gdclass.Register("JSONRPC", func(ptr gd.Object) any { return [1]gdclass.JSONRPC{*(*gdclass.JSONRPC)(unsafe.Pointer(&ptr))} })
}

type ErrorCode int //gd:JSONRPC.ErrorCode

const (
	/*The request could not be parsed as it was not valid by JSON standard ([method JSON.parse] failed).*/
	ParseError ErrorCode = -32700
	/*A method call was requested but the request's format is not valid.*/
	InvalidRequest ErrorCode = -32600
	/*A method call was requested but no function of that name existed in the JSONRPC subclass.*/
	MethodNotFound ErrorCode = -32601
	/*A method call was requested but the given method parameters are not valid. Not used by the built-in JSONRPC.*/
	InvalidParams ErrorCode = -32602
	/*An internal error occurred while processing the request. Not used by the built-in JSONRPC.*/
	InternalError ErrorCode = -32603
)

type Notification struct {
	Method string      `gd:"method"`
	Params interface{} `gd:"params"`
}
type Request struct {
	Method string      `gd:"method"`
	Params interface{} `gd:"params"`
	ID     string      `gd:"id"`
}
type Response struct {
	Result interface{} `gd:"result"`
	ID     string      `gd:"id"`
}
type ResponseError struct {
	Code    int    `gd:"code"`
	Message string `gd:"message"`
	ID      int    `gd:"id"`
}
