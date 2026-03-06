package models

// Beacon 插件塔
// 用于向周围机器传递模块效果
type Beacon struct {
	// ===== 基础信息 =====
	Type          string       `json:"type"`                     // 实体类型
	Name          string       `json:"name"`                     // 实体名称
	Icon          string       `json:"icon"`                     // UI图标
	LocalisedName *interface{} `json:"localised_name,omitempty"` // 本地化名称（多语言字符串数组）
	// ===== 占地 =====
	CollisionBox [][]float64 `json:"collision_box"` // 碰撞盒
	SelectionBox [][]float64 `json:"selection_box"` // 选择盒
	// ===== Beacon核心 =====
	SupplyAreaDistance      float64  `json:"supply_area_distance"`     // 影响半径
	DistributionEffectivity float64  `json:"distribution_effectivity"` // 效果传递效率
	AllowedEffects          []string `json:"allowed_effects"`          // 允许的模块效果
	ModuleSlots             float64  `json:"module_slots"`             // 模块槽数量
	// ===== 能源 =====
	EnergySource EnergySource `json:"energy_source"` // 能源类型
	EnergyUsage  string       `json:"energy_usage"`  // 功耗
	// ===== 升级链 =====
	FastReplaceableGroup *string `json:"fast_replaceable_group,omitempty"` // 快速替换组
}
