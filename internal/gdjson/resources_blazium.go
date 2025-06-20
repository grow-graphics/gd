package gdjson

import (
	"reflect"

	"graphics.gd/variant/RID"
)

func init() {
	Resources["PhysicsServer2D.space_step.space"] = reflect.TypeFor[RID.Space2D]()
	Resources["PhysicsServer2D.space_flush_queries.space"] = reflect.TypeFor[RID.Space2D]()
	Resources["PhysicsServer2D.joint_set_enabled.joint"] = reflect.TypeFor[RID.Joint2D]()
	Resources["PhysicsServer2D.joint_is_enabled.joint"] = reflect.TypeFor[RID.Joint2D]()
	Resources["PhysicsServer2D.space_get_last_process_info.space"] = reflect.TypeFor[RID.Space2D]()
	Resources["PhysicsServer3D.space_step.space"] = reflect.TypeFor[RID.Space3D]()
	Resources["PhysicsServer3D.space_flush_queries.space"] = reflect.TypeFor[RID.Space3D]()
	Resources["PhysicsServer3D.joint_set_enabled.joint"] = reflect.TypeFor[RID.Joint3D]()
	Resources["PhysicsServer3D.joint_is_enabled.joint"] = reflect.TypeFor[RID.Joint3D]()
	Resources["PhysicsServer3D.space_get_last_process_info.space"] = reflect.TypeFor[RID.Space3D]()
	Resources["TextServer.font_set_lcd_subpixel_layout.font_rid"] = reflect.TypeFor[RID.Font]()
	Resources["TextServer.font_get_lcd_subpixel_layout.font_rid"] = reflect.TypeFor[RID.Font]()
}
