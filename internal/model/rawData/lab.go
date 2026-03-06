package rawDataModels

// Lab 研究中心
type Lab struct {
	// ===== 基础信息 =====
	Type string `json:"type"` // 实体类型
	Name string `json:"name"` // 实体名称
	Icon string `json:"icon"` // 图标
	// ===== 碰撞 =====
	CollisionBox [][]float64 `json:"collision_box"` // 碰撞盒
	SelectionBox [][]float64 `json:"selection_box"` // 选择盒
	// ===== 能源 =====
	EnergySource EnergySource `json:"energy_source"` // 能源类型
	EnergyUsage  string       `json:"energy_usage"`  // 能耗
	// ===== 研究属性 =====
	ResearchingSpeed float64  `json:"researching_speed"` // 研究速度倍率
	Inputs           []string `json:"inputs"`            // 支持的科技包
	ModuleSlots      float64  `json:"module_slots"`      // 模块槽数量
	// ===== 升级链 =====
	FastReplaceableGroup *string `json:"fast_replaceable_group,omitempty"` // 快速替换组
}
