package rawDataModels

// =============================
// 效果接收器
// =============================

// EffectReceiver 模块 / 信标效果接收
type EffectReceiver struct {
	UsesModuleEffects  bool `json:"uses_module_effects"`  // 是否接受模块效果
	UsesBeaconEffects  bool `json:"uses_beacon_effects"`  // 是否接受信标效果
	UsesSurfaceEffects bool `json:"uses_surface_effects"` // 是否接受地表效果
}

// Furnace 熔炉类建筑
type Furnace struct {
	Type string `json:"type"` // 原型类型
	Name string `json:"name"` // 名称
	Icon string `json:"icon"` // 图标

	Flags        []string `json:"flags"`         // 标记
	CollisionBox Box      `json:"collision_box"` // 碰撞盒
	SelectionBox Box      `json:"selection_box"` // 选择盒

	FastReplaceableGroup string `json:"fast_replaceable_group"` // 快速替换组
	NextUpgrade          string `json:"next_upgrade,omitempty"` // 下一升级

	// 制造系统
	CraftingCategories  []string `json:"crafting_categories"`   // 支持配方类别
	CraftingSpeed       float64  `json:"crafting_speed"`        // 制作速度
	SourceInventorySize int      `json:"source_inventory_size"` // 输入库存
	ResultInventorySize int      `json:"result_inventory_size"` // 输出库存

	// 能源系统
	EnergySource  EnergySource `json:"energy_source"`            // 能源来源
	EnergyUsage   string       `json:"energy_usage"`             // 能源消耗
	HeatingEnergy string       `json:"heating_energy,omitempty"` // 加热能量

	// 模块系统
	ModuleSlots                            int             `json:"module_slots,omitempty"`                                 // 模块槽
	FastTransferModulesIntoModuleSlotsOnly bool            `json:"fast_transfer_modules_into_module_slots_only,omitempty"` // 是否只能快速插入模块
	EffectReceiver                         *EffectReceiver `json:"effect_receiver,omitempty"`                              // 效果接收器

	SurfaceConditions []SurfaceConditions `json:"surface_conditions,omitempty"` // 表面条件
}
