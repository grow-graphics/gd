// Package ShaderGlobalsOverride provides methods for working with ShaderGlobalsOverride object instances.
package ShaderGlobalsOverride

import "unsafe"
import "reflect"
import "slices"
import "graphics.gd/internal/pointers"
import "graphics.gd/internal/callframe"
import gd "graphics.gd/internal"
import "graphics.gd/internal/gdclass"
import "graphics.gd/variant"
import "graphics.gd/classdb/Node"
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

/*
Similar to how a [WorldEnvironment] node can be used to override the environment while a specific scene is loaded, [ShaderGlobalsOverride] can be used to override global shader parameters temporarily. Once the node is removed, the project-wide values for the global shader parameters are restored. See the [RenderingServer] [code]global_shader_parameter_*[/code] methods for more information.
[b]Note:[/b] Only one [ShaderGlobalsOverride] can be used per scene. If there is more than one [ShaderGlobalsOverride] node in the scene tree, only the first node (in tree order) will be taken into account.
[b]Note:[/b] All [ShaderGlobalsOverride] nodes are made part of a [code]"shader_overrides_group"[/code] group when they are added to the scene tree. The currently active [ShaderGlobalsOverride] node also has a [code]"shader_overrides_group_active"[/code] group added to it. You can use this to check which [ShaderGlobalsOverride] node is currently active.
*/
type Instance [1]gdclass.ShaderGlobalsOverride

// Nil is a nil/null instance of the class. Equivalent to the zero value.
var Nil Instance

type Any interface {
	gd.IsClass
	AsShaderGlobalsOverride() Instance
}

// Advanced exposes a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.
type Advanced = class
type class [1]gdclass.ShaderGlobalsOverride

func (self class) AsObject() [1]gd.Object { return self[0].AsObject() }

//go:nosplit
func (self *class) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func (self Instance) AsObject() [1]gd.Object      { return self[0].AsObject() }

//go:nosplit
func (self *Instance) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(self) }
func New() Instance {
	object := gd.Global.ClassDB.ConstructObject(gd.NewStringName("ShaderGlobalsOverride"))
	casted := Instance{*(*gdclass.ShaderGlobalsOverride)(unsafe.Pointer(&object))}
	return casted
}

func (self class) AsShaderGlobalsOverride() Advanced    { return *((*Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsShaderGlobalsOverride() Instance { return *((*Instance)(unsafe.Pointer(&self))) }
func (self class) AsNode() Node.Advanced                { return *((*Node.Advanced)(unsafe.Pointer(&self))) }
func (self Instance) AsNode() Node.Instance             { return *((*Node.Instance)(unsafe.Pointer(&self))) }

func (self class) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Advanced(self.AsNode()), name)
	}
}

func (self Instance) Virtual(name string) reflect.Value {
	switch name {
	default:
		return gd.VirtualByName(Node.Instance(self.AsNode()), name)
	}
}
func init() {
	gdclass.Register("ShaderGlobalsOverride", func(ptr gd.Object) any {
		return [1]gdclass.ShaderGlobalsOverride{*(*gdclass.ShaderGlobalsOverride)(unsafe.Pointer(&ptr))}
	})
}
