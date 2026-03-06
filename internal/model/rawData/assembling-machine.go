package rawDataModels

// =============================
// 建筑主体
// =============================

// Entity 游戏建筑原型
type AssemblingMachine struct {
	Type                 string       `json:"type"`                             // 原型类型
	Name                 string       `json:"name"`                             // 建筑名称
	Icon                 string       `json:"icon"`                             // 图标
	Flags                []string     `json:"flags"`                            // 标记
	CollisionBox         Box          `json:"collision_box"`                    // 碰撞盒
	SelectionBox         Box          `json:"selection_box"`                    // 选择盒
	FastReplaceableGroup string       `json:"fast_replaceable_group,omitempty"` // 快速替换组
	LocalisedName        *interface{} `json:"localised_name,omitempty"`         // 本地化名称（多语言字符串数组）
	NextUpgrade          string       `json:"next_upgrade,omitempty"`           // 升级目标
	CraftingCategories   []string     `json:"crafting_categories,omitempty"`    // 可制作配方类别
	CraftingSpeed        float64      `json:"crafting_speed,omitempty"`         // 制作速度
	ModuleSlots          int          `json:"module_slots,omitempty"`           // 模块槽
	// 能源
	EnergySource EnergySource `json:"energy_source"` // 能源来源
	EnergyUsage  string       `json:"energy_usage"`  // 能源消耗
	// 流体系统
	FluidBoxes []FluidBox `json:"fluid_boxes,omitempty"` // 流体接口
	// 地形限制
	SurfaceConditions []SurfaceConditions `json:"surface_conditions,omitempty"` // 表面条件
	// 分类
	Subgroup string `json:"subgroup,omitempty"`
}
