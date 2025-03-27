package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Node3D"
)

type Player struct {
	classdb.Extension[Player, Node3D.Instance] `gd:"Player"`
}

func (p Player) AsNode3D() Node3D.Instance {
	return p.Super().AsNode3D()
}

func (p Player) CollectCoin() {}
