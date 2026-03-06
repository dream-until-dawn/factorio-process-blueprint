package rawDataModels

// 物品放置为地形瓦片时的配置
type PlaceAsTile struct {
	Result        string   `json:"result"`                   // 放置后生成的 tile 类型
	ConditionSize float64  `json:"condition_size"`           // 检测区域大小（用于判断是否允许放置）
	TileCondition []string `json:"tile_condition,omitempty"` // 允许放置的地面 tile 类型
	Invert        *bool    `json:"invert,omitempty"`         // 是否反转条件（true 表示在不满足条件时才能放置）
}

// 游戏物品定义（Item Prototype）
type Item struct {
	Type                 string       `json:"type"`                            // 原型类型（通常为 "item" 或其子类型）
	Name                 string       `json:"name"`                            // 物品唯一名称 ID
	LocalisedName        *interface{} `json:"localised_name,omitempty"`        // 本地化名称（支持多种格式）
	LocalisedDescription []string     `json:"localised_description,omitempty"` // 本地化描述
	Icon                 *string      `json:"icon,omitempty"`                  // 单个图标路径
	Hidden               *bool        `json:"hidden,omitempty"`                // 是否在游戏 UI 中隐藏
	StackSize            float64      `json:"stack_size"`                      // 最大堆叠数量
	Subgroup             *string      `json:"subgroup,omitempty"`              // 所属子分类

	Order                         *string      `json:"order,omitempty"`                            // 排序字符串（控制 UI 排序）
	Parameter                     *bool        `json:"parameter,omitempty"`                        // 是否为参数化物品
	PlaceAsTile                   *PlaceAsTile `json:"place_as_tile,omitempty"`                    // 作为 tile 放置时的行为
	FuelValue                     *string      `json:"fuel_value,omitempty"`                       // 燃料能量值（例如 "8MJ"）
	FuelCategory                  *string      `json:"fuel_category,omitempty"`                    // 燃料类别
	Weight                        *float64     `json:"weight,omitempty"`                           // 物品重量（物流/火箭等系统使用）
	DarkBackgroundIcon            *string      `json:"dark_background_icon,omitempty"`             // 深色背景版本图标
	Pictures                      *interface{} `json:"pictures,omitempty"`                         // 图片定义（可能为单个或数组）
	RandomTintColor               []float64    `json:"random_tint_color,omitempty"`                // 随机染色颜色（RGBA）
	IngredientToWeightCoefficient *float64     `json:"ingredient_to_weight_coefficient,omitempty"` // 原料到重量的换算系数
	PlaceResult                   *string      `json:"place_result,omitempty"`                     // 放置后生成的实体名称
	IconSize                      *float64     `json:"icon_size,omitempty"`                        // 图标尺寸
	Flags                         []string     `json:"flags,omitempty"`                            // 标志位（影响游戏行为）
	AutoRecycle                   *bool        `json:"auto_recycle,omitempty"`                     // 是否自动回收
	FuelAccelerationMultiplier    *float64     `json:"fuel_acceleration_multiplier,omitempty"`     // 燃料加速倍率
	FuelTopSpeedMultiplier        *float64     `json:"fuel_top_speed_multiplier,omitempty"`        // 燃料最高速度倍率
	Icons                         []Icons      `json:"icons,omitempty"`                            // 多层图标定义
	BurntResult                   *string      `json:"burnt_result,omitempty"`                     // 燃烧后生成的物品
	PlaceAsEquipmentResult        *string      `json:"place_as_equipment_result,omitempty"`        // 放置为装备时生成的装备实体
	DefaultImportLocation         *string      `json:"default_import_location,omitempty"`          // 默认导入位置（物流系统）
	SpoilTicks                    *float64     `json:"spoil_ticks,omitempty"`                      // 腐坏时间（tick）
	SpoilResult                   *string      `json:"spoil_result,omitempty"`                     // 腐坏后生成的物品
	PlantResult                   *string      `json:"plant_result,omitempty"`                     // 种植后生成的植物实体
	SpoilLevel                    *float64     `json:"spoil_level,omitempty"`                      // 腐坏等级
	FactoriopediaAlternative      *string      `json:"factoriopedia_alternative,omitempty"`        // Factoriopedia 替代条目
}
