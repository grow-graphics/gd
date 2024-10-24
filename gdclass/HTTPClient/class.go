package HTTPClient

import "unsafe"
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
Hyper-text transfer protocol client (sometimes called "User Agent"). Used to make HTTP requests to download web content, upload files and other data or to communicate with various services, among other use cases.
See the [HTTPRequest] node for a higher-level alternative.
[b]Note:[/b] This client only needs to connect to a host once (see [method connect_to_host]) to send multiple requests. Because of this, methods that take URLs usually take just the part after the host instead of the full URL, as the client is already connected to a host. See [method request] for a full example and to get started.
A [HTTPClient] should be reused between multiple requests or to connect to different hosts instead of creating one client per request. Supports Transport Layer Security (TLS), including server certificate verification. HTTP status codes in the 2xx range indicate success, 3xx redirection (i.e. "try again, but over here"), 4xx something was wrong with the request, and 5xx something went wrong on the server's side.
For more information on HTTP, see [url=https://developer.mozilla.org/en-US/docs/Web/HTTP]MDN's documentation on HTTP[/url] (or read [url=https://tools.ietf.org/html/rfc2616]RFC 2616[/url] to get it straight from the source).
[b]Note:[/b] When exporting to Android, make sure to enable the [code]INTERNET[/code] permission in the Android export preset before exporting the project or using one-click deploy. Otherwise, network communication of any kind will be blocked by Android.
[b]Note:[/b] It's recommended to use transport encryption (TLS) and to avoid sending sensitive information (such as login credentials) in HTTP GET URL parameters. Consider using HTTP POST requests or HTTP headers for such information instead.
[b]Note:[/b] When performing HTTP requests from a project exported to Web, keep in mind the remote server may not allow requests from foreign origins due to [url=https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS]CORS[/url]. If you host the server in question, you should modify its backend to allow requests from foreign origins by adding the [code]Access-Control-Allow-Origin: *[/code] HTTP header.
[b]Note:[/b] TLS support is currently limited to TLS 1.0, TLS 1.1, and TLS 1.2. Attempting to connect to a TLS 1.3-only server will return an error.
[b]Warning:[/b] TLS certificate revocation and certificate pinning are currently not supported. Revoked certificates are accepted as long as they are otherwise valid. If this is a concern, you may want to use automatically managed certificates with a short validity period.

*/
type Go [1]classdb.HTTPClient

/*
Connects to a host. This needs to be done before any requests are sent.
If no [param port] is specified (or [code]-1[/code] is used), it is automatically set to 80 for HTTP and 443 for HTTPS. You can pass the optional [param tls_options] parameter to customize the trusted certification authorities, or the common name verification when using HTTPS. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
*/
func (self Go) ConnectToHost(host string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).ConnectToHost(gc.String(host), gd.Int(-1), ([1]gdclass.TLSOptions{}[0])))
}

/*
Sends a raw request to the connected host.
The URL parameter is usually just the part after the host, so for [code]https://somehost.com/index.php[/code], it is [code]/index.php[/code]. When sending requests to an HTTP proxy server, it should be an absolute URL. For [constant HTTPClient.METHOD_OPTIONS] requests, [code]*[/code] is also allowed. For [constant HTTPClient.METHOD_CONNECT] requests, it should be the authority component ([code]host:port[/code]).
Headers are HTTP request headers. For available HTTP methods, see [enum Method].
Sends the body data raw, as a byte array and does not encode it in any way.
*/
func (self Go) RequestRaw(method classdb.HTTPClientMethod, url string, headers []string, body []byte) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).RequestRaw(method, gc.String(url), gc.PackedStringSlice(headers), gc.PackedByteSlice(body)))
}

/*
Sends a request to the connected host.
The URL parameter is usually just the part after the host, so for [code]https://somehost.com/index.php[/code], it is [code]/index.php[/code]. When sending requests to an HTTP proxy server, it should be an absolute URL. For [constant HTTPClient.METHOD_OPTIONS] requests, [code]*[/code] is also allowed. For [constant HTTPClient.METHOD_CONNECT] requests, it should be the authority component ([code]host:port[/code]).
Headers are HTTP request headers. For available HTTP methods, see [enum Method].
To create a POST request with query strings to push to the server, do:
[codeblocks]
[gdscript]
var fields = {"username" : "user", "password" : "pass"}
var query_string = http_client.query_string_from_dict(fields)
var headers = ["Content-Type: application/x-www-form-urlencoded", "Content-Length: " + str(query_string.length())]
var result = http_client.request(http_client.METHOD_POST, "/index.php", headers, query_string)
[/gdscript]
[csharp]
var fields = new Godot.Collections.Dictionary { { "username", "user" }, { "password", "pass" } };
string queryString = new HttpClient().QueryStringFromDict(fields);
string[] headers = { "Content-Type: application/x-www-form-urlencoded", $"Content-Length: {queryString.Length}" };
var result = new HttpClient().Request(HttpClient.Method.Post, "index.php", headers, queryString);
[/csharp]
[/codeblocks]
[b]Note:[/b] The [param body] parameter is ignored if [param method] is [constant HTTPClient.METHOD_GET]. This is because GET methods can't contain request data. As a workaround, you can pass request data as a query string in the URL. See [method String.uri_encode] for an example.
*/
func (self Go) Request(method classdb.HTTPClientMethod, url string, headers []string) gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Request(method, gc.String(url), gc.PackedStringSlice(headers), gc.String("")))
}

/*
Closes the current connection, allowing reuse of this [HTTPClient].
*/
func (self Go) Close() {
	gc := gd.GarbageCollector(); _ = gc
	class(self).Close()
}

/*
If [code]true[/code], this [HTTPClient] has a response available.
*/
func (self Go) HasResponse() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).HasResponse())
}

/*
If [code]true[/code], this [HTTPClient] has a response that is chunked.
*/
func (self Go) IsResponseChunked() bool {
	gc := gd.GarbageCollector(); _ = gc
	return bool(class(self).IsResponseChunked())
}

/*
Returns the response's HTTP status code.
*/
func (self Go) GetResponseCode() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetResponseCode()))
}

/*
Returns the response headers.
*/
func (self Go) GetResponseHeaders() []string {
	gc := gd.GarbageCollector(); _ = gc
	return []string(class(self).GetResponseHeaders(gc).Strings(gc))
}

/*
Returns all response headers as a Dictionary of structure [code]{ "key": "value1; value2" }[/code] where the case-sensitivity of the keys and values is kept like the server delivers it. A value is a simple String, this string can have more than one value where "; " is used as separator.
[b]Example:[/b]
[codeblock]
{
    "content-length": 12,
    "Content-Type": "application/json; charset=UTF-8",
}
[/codeblock]
*/
func (self Go) GetResponseHeadersAsDictionary() gd.Dictionary {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Dictionary(class(self).GetResponseHeadersAsDictionary(gc))
}

/*
Returns the response's body length.
[b]Note:[/b] Some Web servers may not send a body length. In this case, the value returned will be [code]-1[/code]. If using chunked transfer encoding, the body length will also be [code]-1[/code].
[b]Note:[/b] This function always returns [code]-1[/code] on the Web platform due to browsers limitations.
*/
func (self Go) GetResponseBodyLength() int {
	gc := gd.GarbageCollector(); _ = gc
	return int(int(class(self).GetResponseBodyLength()))
}

/*
Reads one chunk from the response.
*/
func (self Go) ReadResponseBodyChunk() []byte {
	gc := gd.GarbageCollector(); _ = gc
	return []byte(class(self).ReadResponseBodyChunk(gc).Bytes())
}

/*
Returns a [enum Status] constant. Need to call [method poll] in order to get status updates.
*/
func (self Go) GetStatus() classdb.HTTPClientStatus {
	gc := gd.GarbageCollector(); _ = gc
	return classdb.HTTPClientStatus(class(self).GetStatus())
}

/*
This needs to be called in order to have any request processed. Check results with [method get_status].
*/
func (self Go) Poll() gd.Error {
	gc := gd.GarbageCollector(); _ = gc
	return gd.Error(class(self).Poll())
}

/*
Sets the proxy server for HTTP requests.
The proxy server is unset if [param host] is empty or [param port] is -1.
*/
func (self Go) SetHttpProxy(host string, port int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHttpProxy(gc.String(host), gd.Int(port))
}

/*
Sets the proxy server for HTTPS requests.
The proxy server is unset if [param host] is empty or [param port] is -1.
*/
func (self Go) SetHttpsProxy(host string, port int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetHttpsProxy(gc.String(host), gd.Int(port))
}

/*
Generates a GET/POST application/x-www-form-urlencoded style query string from a provided dictionary, e.g.:
[codeblocks]
[gdscript]
var fields = {"username": "user", "password": "pass"}
var query_string = http_client.query_string_from_dict(fields)
# Returns "username=user&password=pass"
[/gdscript]
[csharp]
var fields = new Godot.Collections.Dictionary { { "username", "user" }, { "password", "pass" } };
string queryString = httpClient.QueryStringFromDict(fields);
// Returns "username=user&password=pass"
[/csharp]
[/codeblocks]
Furthermore, if a key has a [code]null[/code] value, only the key itself is added, without equal sign and value. If the value is an array, for each value in it a pair with the same key is added.
[codeblocks]
[gdscript]
var fields = {"single": 123, "not_valued": null, "multiple": [22, 33, 44]}
var query_string = http_client.query_string_from_dict(fields)
# Returns "single=123&not_valued&multiple=22&multiple=33&multiple=44"
[/gdscript]
[csharp]
var fields = new Godot.Collections.Dictionary
{
    { "single", 123 },
    { "notValued", default },
    { "multiple", new Godot.Collections.Array { 22, 33, 44 } },
};
string queryString = httpClient.QueryStringFromDict(fields);
// Returns "single=123&not_valued&multiple=22&multiple=33&multiple=44"
[/csharp]
[/codeblocks]
*/
func (self Go) QueryStringFromDict(fields gd.Dictionary) string {
	gc := gd.GarbageCollector(); _ = gc
	return string(class(self).QueryStringFromDict(gc, fields).String())
}
// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type GD = class
type class [1]classdb.HTTPClient
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Go) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Go) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }
func New() Go {
	gc := gd.GarbageCollector()
	object := gc.API.ClassDB.ConstructObject(gc, gc.StringName("HTTPClient"))
	return *(*Go)(unsafe.Pointer(&object))
}

func (self Go) BlockingModeEnabled() bool {
	gc := gd.GarbageCollector(); _ = gc
		return bool(class(self).IsBlockingModeEnabled())
}

func (self Go) SetBlockingModeEnabled(value bool) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetBlockingMode(value)
}

func (self Go) Connection() gdclass.StreamPeer {
	gc := gd.GarbageCollector(); _ = gc
		return gdclass.StreamPeer(class(self).GetConnection(gc))
}

func (self Go) SetConnection(value gdclass.StreamPeer) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetConnection(value)
}

func (self Go) ReadChunkSize() int {
	gc := gd.GarbageCollector(); _ = gc
		return int(int(class(self).GetReadChunkSize()))
}

func (self Go) SetReadChunkSize(value int) {
	gc := gd.GarbageCollector(); _ = gc
	class(self).SetReadChunkSize(gd.Int(value))
}

/*
Connects to a host. This needs to be done before any requests are sent.
If no [param port] is specified (or [code]-1[/code] is used), it is automatically set to 80 for HTTP and 443 for HTTPS. You can pass the optional [param tls_options] parameter to customize the trusted certification authorities, or the common name verification when using HTTPS. See [method TLSOptions.client] and [method TLSOptions.client_unsafe].
*/
//go:nosplit
func (self class) ConnectToHost(host gd.String, port gd.Int, tls_options gdclass.TLSOptions) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(host))
	callframe.Arg(frame, port)
	callframe.Arg(frame, mmm.Get(tls_options[0].AsPointer())[0])
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_connect_to_host, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetConnection(connection gdclass.StreamPeer)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(connection[0].AsPointer())[0])
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_set_connection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetConnection(ctx gd.Lifetime) gdclass.StreamPeer {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_get_connection, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret gdclass.StreamPeer
	ret[0].SetPointer(gd.PointerWithOwnershipTransferredToGo(ctx,r_ret.Get()))
	frame.Free()
	return ret
}
/*
Sends a raw request to the connected host.
The URL parameter is usually just the part after the host, so for [code]https://somehost.com/index.php[/code], it is [code]/index.php[/code]. When sending requests to an HTTP proxy server, it should be an absolute URL. For [constant HTTPClient.METHOD_OPTIONS] requests, [code]*[/code] is also allowed. For [constant HTTPClient.METHOD_CONNECT] requests, it should be the authority component ([code]host:port[/code]).
Headers are HTTP request headers. For available HTTP methods, see [enum Method].
Sends the body data raw, as a byte array and does not encode it in any way.
*/
//go:nosplit
func (self class) RequestRaw(method classdb.HTTPClientMethod, url gd.String, headers gd.PackedStringArray, body gd.PackedByteArray) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, method)
	callframe.Arg(frame, mmm.Get(url))
	callframe.Arg(frame, mmm.Get(headers))
	callframe.Arg(frame, mmm.Get(body))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_request_raw, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sends a request to the connected host.
The URL parameter is usually just the part after the host, so for [code]https://somehost.com/index.php[/code], it is [code]/index.php[/code]. When sending requests to an HTTP proxy server, it should be an absolute URL. For [constant HTTPClient.METHOD_OPTIONS] requests, [code]*[/code] is also allowed. For [constant HTTPClient.METHOD_CONNECT] requests, it should be the authority component ([code]host:port[/code]).
Headers are HTTP request headers. For available HTTP methods, see [enum Method].
To create a POST request with query strings to push to the server, do:
[codeblocks]
[gdscript]
var fields = {"username" : "user", "password" : "pass"}
var query_string = http_client.query_string_from_dict(fields)
var headers = ["Content-Type: application/x-www-form-urlencoded", "Content-Length: " + str(query_string.length())]
var result = http_client.request(http_client.METHOD_POST, "/index.php", headers, query_string)
[/gdscript]
[csharp]
var fields = new Godot.Collections.Dictionary { { "username", "user" }, { "password", "pass" } };
string queryString = new HttpClient().QueryStringFromDict(fields);
string[] headers = { "Content-Type: application/x-www-form-urlencoded", $"Content-Length: {queryString.Length}" };
var result = new HttpClient().Request(HttpClient.Method.Post, "index.php", headers, queryString);
[/csharp]
[/codeblocks]
[b]Note:[/b] The [param body] parameter is ignored if [param method] is [constant HTTPClient.METHOD_GET]. This is because GET methods can't contain request data. As a workaround, you can pass request data as a query string in the URL. See [method String.uri_encode] for an example.
*/
//go:nosplit
func (self class) Request(method classdb.HTTPClientMethod, url gd.String, headers gd.PackedStringArray, body gd.String) int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, method)
	callframe.Arg(frame, mmm.Get(url))
	callframe.Arg(frame, mmm.Get(headers))
	callframe.Arg(frame, mmm.Get(body))
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_request, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Closes the current connection, allowing reuse of this [HTTPClient].
*/
//go:nosplit
func (self class) Close()  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_close, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
If [code]true[/code], this [HTTPClient] has a response available.
*/
//go:nosplit
func (self class) HasResponse() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_has_response, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
If [code]true[/code], this [HTTPClient] has a response that is chunked.
*/
//go:nosplit
func (self class) IsResponseChunked() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_is_response_chunked, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the response's HTTP status code.
*/
//go:nosplit
func (self class) GetResponseCode() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_get_response_code, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns the response headers.
*/
//go:nosplit
func (self class) GetResponseHeaders(ctx gd.Lifetime) gd.PackedStringArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_get_response_headers, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedStringArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns all response headers as a Dictionary of structure [code]{ "key": "value1; value2" }[/code] where the case-sensitivity of the keys and values is kept like the server delivers it. A value is a simple String, this string can have more than one value where "; " is used as separator.
[b]Example:[/b]
[codeblock]
{
    "content-length": 12,
    "Content-Type": "application/json; charset=UTF-8",
}
[/codeblock]
*/
//go:nosplit
func (self class) GetResponseHeadersAsDictionary(ctx gd.Lifetime) gd.Dictionary {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_get_response_headers_as_dictionary, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.Dictionary](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
/*
Returns the response's body length.
[b]Note:[/b] Some Web servers may not send a body length. In this case, the value returned will be [code]-1[/code]. If using chunked transfer encoding, the body length will also be [code]-1[/code].
[b]Note:[/b] This function always returns [code]-1[/code] on the Web platform due to browsers limitations.
*/
//go:nosplit
func (self class) GetResponseBodyLength() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_get_response_body_length, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Reads one chunk from the response.
*/
//go:nosplit
func (self class) ReadResponseBodyChunk(ctx gd.Lifetime) gd.PackedByteArray {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[[2]uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_read_response_body_chunk, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.PackedByteArray](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetReadChunkSize(bytes gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, bytes)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_set_read_chunk_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) GetReadChunkSize() gd.Int {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[gd.Int](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_get_read_chunk_size, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
//go:nosplit
func (self class) SetBlockingMode(enabled bool)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, enabled)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_set_blocking_mode, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
//go:nosplit
func (self class) IsBlockingModeEnabled() bool {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[bool](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_is_blocking_mode_enabled, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Returns a [enum Status] constant. Need to call [method poll] in order to get status updates.
*/
//go:nosplit
func (self class) GetStatus() classdb.HTTPClientStatus {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[classdb.HTTPClientStatus](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_get_status, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
This needs to be called in order to have any request processed. Check results with [method get_status].
*/
//go:nosplit
func (self class) Poll() int64 {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	var r_ret = callframe.Ret[int64](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_poll, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = r_ret.Get()
	frame.Free()
	return ret
}
/*
Sets the proxy server for HTTP requests.
The proxy server is unset if [param host] is empty or [param port] is -1.
*/
//go:nosplit
func (self class) SetHttpProxy(host gd.String, port gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(host))
	callframe.Arg(frame, port)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_set_http_proxy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Sets the proxy server for HTTPS requests.
The proxy server is unset if [param host] is empty or [param port] is -1.
*/
//go:nosplit
func (self class) SetHttpsProxy(host gd.String, port gd.Int)  {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(host))
	callframe.Arg(frame, port)
	var r_ret callframe.Nil
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_set_https_proxy, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	frame.Free()
}
/*
Generates a GET/POST application/x-www-form-urlencoded style query string from a provided dictionary, e.g.:
[codeblocks]
[gdscript]
var fields = {"username": "user", "password": "pass"}
var query_string = http_client.query_string_from_dict(fields)
# Returns "username=user&password=pass"
[/gdscript]
[csharp]
var fields = new Godot.Collections.Dictionary { { "username", "user" }, { "password", "pass" } };
string queryString = httpClient.QueryStringFromDict(fields);
// Returns "username=user&password=pass"
[/csharp]
[/codeblocks]
Furthermore, if a key has a [code]null[/code] value, only the key itself is added, without equal sign and value. If the value is an array, for each value in it a pair with the same key is added.
[codeblocks]
[gdscript]
var fields = {"single": 123, "not_valued": null, "multiple": [22, 33, 44]}
var query_string = http_client.query_string_from_dict(fields)
# Returns "single=123&not_valued&multiple=22&multiple=33&multiple=44"
[/gdscript]
[csharp]
var fields = new Godot.Collections.Dictionary
{
    { "single", 123 },
    { "notValued", default },
    { "multiple", new Godot.Collections.Array { 22, 33, 44 } },
};
string queryString = httpClient.QueryStringFromDict(fields);
// Returns "single=123&not_valued&multiple=22&multiple=33&multiple=44"
[/csharp]
[/codeblocks]
*/
//go:nosplit
func (self class) QueryStringFromDict(ctx gd.Lifetime, fields gd.Dictionary) gd.String {
	var selfPtr = self[0].AsPointer()
	var frame = callframe.New()
	callframe.Arg(frame, mmm.Get(fields))
	var r_ret = callframe.Ret[uintptr](frame)
	mmm.API(selfPtr).Object.MethodBindPointerCall(mmm.API(selfPtr).Methods.HTTPClient.Bind_query_string_from_dict, self.AsObject(), frame.Array(0), r_ret.Uintptr())
	var ret = mmm.New[gd.String](ctx.Lifetime, ctx.API, r_ret.Get())
	frame.Free()
	return ret
}
func (self class) AsHTTPClient() GD { return *((*GD)(unsafe.Pointer(&self))) }
func (self Go) AsHTTPClient() Go { return *((*Go)(unsafe.Pointer(&self))) }
func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }
func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}

func (self Go) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self.AsRefCounted(), name)
	}
}
func init() {classdb.Register("HTTPClient", func(ptr gd.Pointer) any {var class class; class[0].SetPointer(ptr); return class })}
type Method = classdb.HTTPClientMethod

const (
/*HTTP GET method. The GET method requests a representation of the specified resource. Requests using GET should only retrieve data.*/
	MethodGet Method = 0
/*HTTP HEAD method. The HEAD method asks for a response identical to that of a GET request, but without the response body. This is useful to request metadata like HTTP headers or to check if a resource exists.*/
	MethodHead Method = 1
/*HTTP POST method. The POST method is used to submit an entity to the specified resource, often causing a change in state or side effects on the server. This is often used for forms and submitting data or uploading files.*/
	MethodPost Method = 2
/*HTTP PUT method. The PUT method asks to replace all current representations of the target resource with the request payload. (You can think of POST as "create or update" and PUT as "update", although many services tend to not make a clear distinction or change their meaning).*/
	MethodPut Method = 3
/*HTTP DELETE method. The DELETE method requests to delete the specified resource.*/
	MethodDelete Method = 4
/*HTTP OPTIONS method. The OPTIONS method asks for a description of the communication options for the target resource. Rarely used.*/
	MethodOptions Method = 5
/*HTTP TRACE method. The TRACE method performs a message loop-back test along the path to the target resource. Returns the entire HTTP request received in the response body. Rarely used.*/
	MethodTrace Method = 6
/*HTTP CONNECT method. The CONNECT method establishes a tunnel to the server identified by the target resource. Rarely used.*/
	MethodConnect Method = 7
/*HTTP PATCH method. The PATCH method is used to apply partial modifications to a resource.*/
	MethodPatch Method = 8
/*Represents the size of the [enum Method] enum.*/
	MethodMax Method = 9
)
type Status = classdb.HTTPClientStatus

const (
/*Status: Disconnected from the server.*/
	StatusDisconnected Status = 0
/*Status: Currently resolving the hostname for the given URL into an IP.*/
	StatusResolving Status = 1
/*Status: DNS failure: Can't resolve the hostname for the given URL.*/
	StatusCantResolve Status = 2
/*Status: Currently connecting to server.*/
	StatusConnecting Status = 3
/*Status: Can't connect to the server.*/
	StatusCantConnect Status = 4
/*Status: Connection established.*/
	StatusConnected Status = 5
/*Status: Currently sending request.*/
	StatusRequesting Status = 6
/*Status: HTTP body received.*/
	StatusBody Status = 7
/*Status: Error in HTTP connection.*/
	StatusConnectionError Status = 8
/*Status: Error in TLS handshake.*/
	StatusTlsHandshakeError Status = 9
)
type ResponseCode = classdb.HTTPClientResponseCode

const (
/*HTTP status code [code]100 Continue[/code]. Interim response that indicates everything so far is OK and that the client should continue with the request (or ignore this status if already finished).*/
	ResponseContinue ResponseCode = 100
/*HTTP status code [code]101 Switching Protocol[/code]. Sent in response to an [code]Upgrade[/code] request header by the client. Indicates the protocol the server is switching to.*/
	ResponseSwitchingProtocols ResponseCode = 101
/*HTTP status code [code]102 Processing[/code] (WebDAV). Indicates that the server has received and is processing the request, but no response is available yet.*/
	ResponseProcessing ResponseCode = 102
/*HTTP status code [code]200 OK[/code]. The request has succeeded. Default response for successful requests. Meaning varies depending on the request. GET: The resource has been fetched and is transmitted in the message body. HEAD: The entity headers are in the message body. POST: The resource describing the result of the action is transmitted in the message body. TRACE: The message body contains the request message as received by the server.*/
	ResponseOk ResponseCode = 200
/*HTTP status code [code]201 Created[/code]. The request has succeeded and a new resource has been created as a result of it. This is typically the response sent after a PUT request.*/
	ResponseCreated ResponseCode = 201
/*HTTP status code [code]202 Accepted[/code]. The request has been received but not yet acted upon. It is non-committal, meaning that there is no way in HTTP to later send an asynchronous response indicating the outcome of processing the request. It is intended for cases where another process or server handles the request, or for batch processing.*/
	ResponseAccepted ResponseCode = 202
/*HTTP status code [code]203 Non-Authoritative Information[/code]. This response code means returned meta-information set is not exact set as available from the origin server, but collected from a local or a third party copy. Except this condition, 200 OK response should be preferred instead of this response.*/
	ResponseNonAuthoritativeInformation ResponseCode = 203
/*HTTP status code [code]204 No Content[/code]. There is no content to send for this request, but the headers may be useful. The user-agent may update its cached headers for this resource with the new ones.*/
	ResponseNoContent ResponseCode = 204
/*HTTP status code [code]205 Reset Content[/code]. The server has fulfilled the request and desires that the client resets the "document view" that caused the request to be sent to its original state as received from the origin server.*/
	ResponseResetContent ResponseCode = 205
/*HTTP status code [code]206 Partial Content[/code]. This response code is used because of a range header sent by the client to separate download into multiple streams.*/
	ResponsePartialContent ResponseCode = 206
/*HTTP status code [code]207 Multi-Status[/code] (WebDAV). A Multi-Status response conveys information about multiple resources in situations where multiple status codes might be appropriate.*/
	ResponseMultiStatus ResponseCode = 207
/*HTTP status code [code]208 Already Reported[/code] (WebDAV). Used inside a DAV: propstat response element to avoid enumerating the internal members of multiple bindings to the same collection repeatedly.*/
	ResponseAlreadyReported ResponseCode = 208
/*HTTP status code [code]226 IM Used[/code] (WebDAV). The server has fulfilled a GET request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance.*/
	ResponseImUsed ResponseCode = 226
/*HTTP status code [code]300 Multiple Choice[/code]. The request has more than one possible responses and there is no standardized way to choose one of the responses. User-agent or user should choose one of them.*/
	ResponseMultipleChoices ResponseCode = 300
/*HTTP status code [code]301 Moved Permanently[/code]. Redirection. This response code means the URI of requested resource has been changed. The new URI is usually included in the response.*/
	ResponseMovedPermanently ResponseCode = 301
/*HTTP status code [code]302 Found[/code]. Temporary redirection. This response code means the URI of requested resource has been changed temporarily. New changes in the URI might be made in the future. Therefore, this same URI should be used by the client in future requests.*/
	ResponseFound ResponseCode = 302
/*HTTP status code [code]303 See Other[/code]. The server is redirecting the user agent to a different resource, as indicated by a URI in the Location header field, which is intended to provide an indirect response to the original request.*/
	ResponseSeeOther ResponseCode = 303
/*HTTP status code [code]304 Not Modified[/code]. A conditional GET or HEAD request has been received and would have resulted in a 200 OK response if it were not for the fact that the condition evaluated to [code]false[/code].*/
	ResponseNotModified ResponseCode = 304
/*HTTP status code [code]305 Use Proxy[/code].*/
	ResponseUseProxy ResponseCode = 305
/*HTTP status code [code]306 Switch Proxy[/code].*/
	ResponseSwitchProxy ResponseCode = 306
/*HTTP status code [code]307 Temporary Redirect[/code]. The target resource resides temporarily under a different URI and the user agent MUST NOT change the request method if it performs an automatic redirection to that URI.*/
	ResponseTemporaryRedirect ResponseCode = 307
/*HTTP status code [code]308 Permanent Redirect[/code]. The target resource has been assigned a new permanent URI and any future references to this resource ought to use one of the enclosed URIs.*/
	ResponsePermanentRedirect ResponseCode = 308
/*HTTP status code [code]400 Bad Request[/code]. The request was invalid. The server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, invalid request contents, or deceptive request routing).*/
	ResponseBadRequest ResponseCode = 400
/*HTTP status code [code]401 Unauthorized[/code]. Credentials required. The request has not been applied because it lacks valid authentication credentials for the target resource.*/
	ResponseUnauthorized ResponseCode = 401
/*HTTP status code [code]402 Payment Required[/code]. This response code is reserved for future use. Initial aim for creating this code was using it for digital payment systems, however this is not currently used.*/
	ResponsePaymentRequired ResponseCode = 402
/*HTTP status code [code]403 Forbidden[/code]. The client does not have access rights to the content, i.e. they are unauthorized, so server is rejecting to give proper response. Unlike [code]401[/code], the client's identity is known to the server.*/
	ResponseForbidden ResponseCode = 403
/*HTTP status code [code]404 Not Found[/code]. The server can not find requested resource. Either the URL is not recognized or the endpoint is valid but the resource itself does not exist. May also be sent instead of 403 to hide existence of a resource if the client is not authorized.*/
	ResponseNotFound ResponseCode = 404
/*HTTP status code [code]405 Method Not Allowed[/code]. The request's HTTP method is known by the server but has been disabled and cannot be used. For example, an API may forbid DELETE-ing a resource. The two mandatory methods, GET and HEAD, must never be disabled and should not return this error code.*/
	ResponseMethodNotAllowed ResponseCode = 405
/*HTTP status code [code]406 Not Acceptable[/code]. The target resource does not have a current representation that would be acceptable to the user agent, according to the proactive negotiation header fields received in the request. Used when negotiation content.*/
	ResponseNotAcceptable ResponseCode = 406
/*HTTP status code [code]407 Proxy Authentication Required[/code]. Similar to 401 Unauthorized, but it indicates that the client needs to authenticate itself in order to use a proxy.*/
	ResponseProxyAuthenticationRequired ResponseCode = 407
/*HTTP status code [code]408 Request Timeout[/code]. The server did not receive a complete request message within the time that it was prepared to wait.*/
	ResponseRequestTimeout ResponseCode = 408
/*HTTP status code [code]409 Conflict[/code]. The request could not be completed due to a conflict with the current state of the target resource. This code is used in situations where the user might be able to resolve the conflict and resubmit the request.*/
	ResponseConflict ResponseCode = 409
/*HTTP status code [code]410 Gone[/code]. The target resource is no longer available at the origin server and this condition is likely permanent.*/
	ResponseGone ResponseCode = 410
/*HTTP status code [code]411 Length Required[/code]. The server refuses to accept the request without a defined Content-Length header.*/
	ResponseLengthRequired ResponseCode = 411
/*HTTP status code [code]412 Precondition Failed[/code]. One or more conditions given in the request header fields evaluated to [code]false[/code] when tested on the server.*/
	ResponsePreconditionFailed ResponseCode = 412
/*HTTP status code [code]413 Entity Too Large[/code]. The server is refusing to process a request because the request payload is larger than the server is willing or able to process.*/
	ResponseRequestEntityTooLarge ResponseCode = 413
/*HTTP status code [code]414 Request-URI Too Long[/code]. The server is refusing to service the request because the request-target is longer than the server is willing to interpret.*/
	ResponseRequestUriTooLong ResponseCode = 414
/*HTTP status code [code]415 Unsupported Media Type[/code]. The origin server is refusing to service the request because the payload is in a format not supported by this method on the target resource.*/
	ResponseUnsupportedMediaType ResponseCode = 415
/*HTTP status code [code]416 Requested Range Not Satisfiable[/code]. None of the ranges in the request's Range header field overlap the current extent of the selected resource or the set of ranges requested has been rejected due to invalid ranges or an excessive request of small or overlapping ranges.*/
	ResponseRequestedRangeNotSatisfiable ResponseCode = 416
/*HTTP status code [code]417 Expectation Failed[/code]. The expectation given in the request's Expect header field could not be met by at least one of the inbound servers.*/
	ResponseExpectationFailed ResponseCode = 417
/*HTTP status code [code]418 I'm A Teapot[/code]. Any attempt to brew coffee with a teapot should result in the error code "418 I'm a teapot". The resulting entity body MAY be short and stout.*/
	ResponseImATeapot ResponseCode = 418
/*HTTP status code [code]421 Misdirected Request[/code]. The request was directed at a server that is not able to produce a response. This can be sent by a server that is not configured to produce responses for the combination of scheme and authority that are included in the request URI.*/
	ResponseMisdirectedRequest ResponseCode = 421
/*HTTP status code [code]422 Unprocessable Entity[/code] (WebDAV). The server understands the content type of the request entity (hence a 415 Unsupported Media Type status code is inappropriate), and the syntax of the request entity is correct (thus a 400 Bad Request status code is inappropriate) but was unable to process the contained instructions.*/
	ResponseUnprocessableEntity ResponseCode = 422
/*HTTP status code [code]423 Locked[/code] (WebDAV). The source or destination resource of a method is locked.*/
	ResponseLocked ResponseCode = 423
/*HTTP status code [code]424 Failed Dependency[/code] (WebDAV). The method could not be performed on the resource because the requested action depended on another action and that action failed.*/
	ResponseFailedDependency ResponseCode = 424
/*HTTP status code [code]426 Upgrade Required[/code]. The server refuses to perform the request using the current protocol but might be willing to do so after the client upgrades to a different protocol.*/
	ResponseUpgradeRequired ResponseCode = 426
/*HTTP status code [code]428 Precondition Required[/code]. The origin server requires the request to be conditional.*/
	ResponsePreconditionRequired ResponseCode = 428
/*HTTP status code [code]429 Too Many Requests[/code]. The user has sent too many requests in a given amount of time (see "rate limiting"). Back off and increase time between requests or try again later.*/
	ResponseTooManyRequests ResponseCode = 429
/*HTTP status code [code]431 Request Header Fields Too Large[/code]. The server is unwilling to process the request because its header fields are too large. The request MAY be resubmitted after reducing the size of the request header fields.*/
	ResponseRequestHeaderFieldsTooLarge ResponseCode = 431
/*HTTP status code [code]451 Response Unavailable For Legal Reasons[/code]. The server is denying access to the resource as a consequence of a legal demand.*/
	ResponseUnavailableForLegalReasons ResponseCode = 451
/*HTTP status code [code]500 Internal Server Error[/code]. The server encountered an unexpected condition that prevented it from fulfilling the request.*/
	ResponseInternalServerError ResponseCode = 500
/*HTTP status code [code]501 Not Implemented[/code]. The server does not support the functionality required to fulfill the request.*/
	ResponseNotImplemented ResponseCode = 501
/*HTTP status code [code]502 Bad Gateway[/code]. The server, while acting as a gateway or proxy, received an invalid response from an inbound server it accessed while attempting to fulfill the request. Usually returned by load balancers or proxies.*/
	ResponseBadGateway ResponseCode = 502
/*HTTP status code [code]503 Service Unavailable[/code]. The server is currently unable to handle the request due to a temporary overload or scheduled maintenance, which will likely be alleviated after some delay. Try again later.*/
	ResponseServiceUnavailable ResponseCode = 503
/*HTTP status code [code]504 Gateway Timeout[/code]. The server, while acting as a gateway or proxy, did not receive a timely response from an upstream server it needed to access in order to complete the request. Usually returned by load balancers or proxies.*/
	ResponseGatewayTimeout ResponseCode = 504
/*HTTP status code [code]505 HTTP Version Not Supported[/code]. The server does not support, or refuses to support, the major version of HTTP that was used in the request message.*/
	ResponseHttpVersionNotSupported ResponseCode = 505
/*HTTP status code [code]506 Variant Also Negotiates[/code]. The server has an internal configuration error: the chosen variant resource is configured to engage in transparent content negotiation itself, and is therefore not a proper end point in the negotiation process.*/
	ResponseVariantAlsoNegotiates ResponseCode = 506
/*HTTP status code [code]507 Insufficient Storage[/code]. The method could not be performed on the resource because the server is unable to store the representation needed to successfully complete the request.*/
	ResponseInsufficientStorage ResponseCode = 507
/*HTTP status code [code]508 Loop Detected[/code]. The server terminated an operation because it encountered an infinite loop while processing a request with "Depth: infinity". This status indicates that the entire operation failed.*/
	ResponseLoopDetected ResponseCode = 508
/*HTTP status code [code]510 Not Extended[/code]. The policy for accessing the resource has not been met in the request. The server should send back all the information necessary for the client to issue an extended request.*/
	ResponseNotExtended ResponseCode = 510
/*HTTP status code [code]511 Network Authentication Required[/code]. The client needs to authenticate to gain network access.*/
	ResponseNetworkAuthRequired ResponseCode = 511
)