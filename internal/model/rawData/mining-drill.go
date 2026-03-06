package rawDataModels

// 输入流体箱
type InputFluidBox struct {
	Volume          float64          `json:"volume"`           // 最大容量
	PipeConnections []PipeConnection `json:"pipe_connections"` // 管道连接点
}

// 输出流体箱
type OutputFluidBox struct {
	Volume          float64          `json:"volume"`           // 最大容量
	PipeConnections []PipeConnection `json:"pipe_connections"` // 管道连接点
}

// 挖矿范围可视化
type RadiusVisualisationPicture struct {
	Filename string  `json:"filename"` // 图片路径
	Width    float64 `json:"width"`    // 图片宽度
	Height   float64 `json:"height"`   // 图片高度
}

// 采矿机实体定义
type MiningDrill struct {
	Type          string   `json:"type"`                     // 原型类型
	Name          string   `json:"name"`                     // 实体唯一ID
	LocalisedName []string `json:"localised_name,omitempty"` // 本地化名称（多语言显示）
	Icon          string   `json:"icon"`                     // 图标路径

	Flags              []string `json:"flags"`               // 实体标志
	ResourceCategories []string `json:"resource_categories"` // 可采集资源类别

	CollisionBox [][]float64 `json:"collision_box"` // 碰撞盒
	SelectionBox [][]float64 `json:"selection_box"` // 选择盒

	InputFluidBox  *InputFluidBox  `json:"input_fluid_box,omitempty"`  // 输入流体
	OutputFluidBox *OutputFluidBox `json:"output_fluid_box,omitempty"` // 输出流体

	MiningSpeed             float64 `json:"mining_speed"`              // 挖矿速度
	ResourceSearchingRadius float64 `json:"resource_searching_radius"` // 搜索半径

	VectorToPlaceResult []float64 `json:"vector_to_place_result"` // 挖掘结果掉落位置偏移

	EnergySource EnergySource `json:"energy_source"` // 能源系统
	EnergyUsage  string       `json:"energy_usage"`  // 能耗（如 90kW）

	ModuleSlots *float64 `json:"module_slots,omitempty"` // 模块槽数量

	FastReplaceableGroup string `json:"fast_replaceable_group"` // 快速替换组

	AllowedEffects []string `json:"allowed_effects,omitempty"` // 可使用模块效果

	SurfaceConditions []SurfaceConditions `json:"surface_conditions,omitempty"` // 地表条件限制

	RadiusVisualisationPicture *RadiusVisualisationPicture `json:"radius_visualisation_picture,omitempty"` // 挖矿范围显示

	ResourceDrainRatePercent *float64 `json:"resource_drain_rate_percent,omitempty"` // 资源消耗比例
	DropsFullBeltStacks      *bool    `json:"drops_full_belt_stacks,omitempty"`      // 是否掉落整带物品
}
