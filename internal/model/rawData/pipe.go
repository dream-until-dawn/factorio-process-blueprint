package rawDataModels

// 管道原型（pipe prototype）
type Pipe struct {
	Type string  `json:"type"`           // 原型类型（pipe）
	Name string  `json:"name"`           // 管道ID
	Icon *string `json:"icon,omitempty"` // 图标

	Flags []string `json:"flags,omitempty"` // 实体标志

	FastReplaceableGroup *string `json:"fast_replaceable_group,omitempty"` // 快速替换组

	CollisionBox [][]float64 `json:"collision_box,omitempty"` // 碰撞盒
	SelectionBox [][]float64 `json:"selection_box,omitempty"` // 选择盒`

	FluidBox      *FluidBox `json:"fluid_box,omitempty"`      // 流体系统
	HeatingEnergy *string   `json:"heating_energy,omitempty"` // 管道加热能量（用于热管系统）
}
