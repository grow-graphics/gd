package JavaScriptObject

import "unsafe"
import "reflect"
import "grow.graphics/gd/internal/mmm"
import "grow.graphics/gd/internal/callframe"
import gd "grow.graphics/gd/internal"
import object "grow.graphics/gd/object"
import classdb "grow.graphics/gd/internal/classdb"

var _ unsafe.Pointer
var _ object.Engine
var _ reflect.Type
var _ callframe.Frame
var _ mmm.Lifetime

/*
JavaScriptObject is used to interact with JavaScript objects retrieved or created via [method JavaScriptBridge.get_interface], [method JavaScriptBridge.create_object], or [method JavaScriptBridge.create_callback].
[b]Example:[/b]
[codeblock]
extends Node

var _my_js_callback = JavaScriptBridge.create_callback(myCallback) # This reference must be kept
var console = JavaScriptBridge.get_interface("console")

func _init():
    var buf = JavaScriptBridge.create_object("ArrayBuffer", 10) # new ArrayBuffer(10)
    print(buf) # prints [JavaScriptObject:OBJECT_ID]
    var uint8arr = JavaScriptBridge.create_object("Uint8Array", buf) # new Uint8Array(buf)
    uint8arr[1] = 255
    prints(uint8arr[1], uint8arr.byteLength) # prints 255 10
    console.log(uint8arr) # prints in browser console "Uint8Array(10) [ 0, 255, 0, 0, 0, 0, 0, 0, 0, 0 ]"

    # Equivalent of JavaScriptBridge: Array.from(uint8arr).forEach(myCallback)
    JavaScriptBridge.get_interface("Array").from(uint8arr).forEach(_my_js_callback)

func myCallback(args):
    # Will be called with the parameters passed to the "forEach" callback
    # [0, 0, [JavaScriptObject:1173]]
    # [255, 1, [JavaScriptObject:1173]]
    # ...
    # [0, 9, [JavaScriptObject:1180]]
    print(args)
[/codeblock]
[b]Note:[/b] Only available in the Web platform.

*/
type Simple [1]classdb.JavaScriptObject
// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Expert = class
type class [1]classdb.JavaScriptObject
func (self class) AsObject() gd.Object { return self[0].AsObject() }
func (self Simple) AsObject() gd.Object { return self[0].AsObject() }


//go:nosplit
func (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }


//go:nosplit
func (self class) AsJavaScriptObject() Expert { return self[0].AsJavaScriptObject() }


//go:nosplit
func (self Simple) AsJavaScriptObject() Simple { return self[0].AsJavaScriptObject() }


//go:nosplit
func (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }


//go:nosplit
func (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}

func (self Simple) Virtual(name string) reflect.Value {
	switch name {
	default: return gd.VirtualByName(self[0].Super()[0], name)
	}
}
func init() {classdb.Register("JavaScriptObject", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })}
