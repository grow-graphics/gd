package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Node3D"
	"graphics.gd/variant/Vector3"
)

type Player struct {
	classdb.Extension[Player, Node3D.Instance] `gd:"Player"`
}

func (p Player) AsNode3D() Node3D.Instance {
	return p.Super().AsNode3D()
}

func (p Player) CollectCoin() {}

func (p Player) Damage(impact_point Vector3.XYZ, velocity Vector3.XYZ) {
	// Handle damage logic here
}
