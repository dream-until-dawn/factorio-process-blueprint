package rawDataModels

// RecipeData 表示游戏中的一个配方定义（核心生产数据）
type RecipeData struct {
	Type                            string              `json:"type"`                                          // 原型类型，通常为 "recipe"
	Name                            string              `json:"name"`                                          // 配方唯一ID，例如 electronic-circuit
	LocalisedName                   *interface{}        `json:"localised_name,omitempty"`                      // 本地化名称（多语言字符串数组）
	Icon                            *string             `json:"icon,omitempty"`                                // 单图标路径
	IconSize                        *float64            `json:"icon_size,omitempty"`                           // 图标尺寸
	Hidden                          *bool               `json:"hidden,omitempty"`                              // 是否隐藏（不会在玩家UI中显示）
	Ingredients                     *Ingredients        `json:"ingredients,omitempty"`                         // 输入材料
	Results                         *Results            `json:"results,omitempty"`                             // 输出产品
	Subgroup                        *string             `json:"subgroup,omitempty"`                            // UI子分类
	Category                        *string             `json:"category,omitempty"`                            // 制造类别，例如 crafting / smelting / chemistry
	Order                           *string             `json:"order,omitempty"`                               // UI排序权重
	Parameter                       *bool               `json:"parameter,omitempty"`                           // 是否为参数化配方（用于动态配方）
	AllowProductivity               *bool               `json:"allow_productivity,omitempty"`                  // 是否允许生产力模块影响
	Enabled                         *bool               `json:"enabled,omitempty"`                             // 游戏开始时是否已解锁
	EnergyRequired                  *float64            `json:"energy_required,omitempty"`                     // 制作时间（秒）
	MainProduct                     *string             `json:"main_product,omitempty"`                        // 主产物（用于UI显示）
	AllowQuality                    *bool               `json:"allow_quality,omitempty"`                       // 是否允许质量系统影响（Space Age）
	AllowDecomposition              *bool               `json:"allow_decomposition,omitempty"`                 // 是否允许被分解系统拆分
	AutoRecycle                     *bool               `json:"auto_recycle,omitempty"`                        // 是否允许自动回收
	RequesterPasteMultiplier        *float64            `json:"requester_paste_multiplier,omitempty"`          // 请求箱粘贴倍率
	HideFromPlayerCrafting          *bool               `json:"hide_from_player_crafting,omitempty"`           // 不允许玩家手工制作
	SurfaceConditions               []SurfaceConditions `json:"surface_conditions,omitempty"`                  // 星球表面条件限制
	ShowAmountInTitle               *bool               `json:"show_amount_in_title,omitempty"`                // UI标题显示产量
	ResetFreshnessOnCraft           *bool               `json:"reset_freshness_on_craft,omitempty"`            // 制作时重置新鲜度（食物系统）
	HideFromSignalGui               *bool               `json:"hide_from_signal_gui,omitempty"`                // 不在电路信号GUI中显示
	AlwaysShowProducts              *bool               `json:"always_show_products,omitempty"`                // 始终显示产物
	PreserveProductsInMachineOutput *bool               `json:"preserve_products_in_machine_output,omitempty"` // 保留机器输出中的产物
	Icons                           []Icons             `json:"icons,omitempty"`                               // 多图标定义
	AlwaysShowMadeIn                *bool               `json:"always_show_made_in,omitempty"`                 // UI始终显示制造机器
	FactoriopediaAlternative        *string             `json:"factoriopedia_alternative,omitempty"`           // Factoriopedia百科替代条目
	UnlockResults                   *bool               `json:"unlock_results,omitempty"`                      // 解锁配方时同时解锁产物
}

// Ingredients 表示配方输入材料
type Ingredients struct {
	Type               *string  `json:"type,omitempty"`                // 类型：item 或 fluid
	Name               *string  `json:"name,omitempty"`                // 材料名称
	Amount             *float64 `json:"amount,omitempty"`              // 消耗数量
	FluidboxIndex      *float64 `json:"fluidbox_index,omitempty"`      // 液体输入口索引
	IgnoredByStats     *float64 `json:"ignored_by_stats,omitempty"`    // 不计入生产统计
	FluidboxMultiplier *float64 `json:"fluidbox_multiplier,omitempty"` // 液体输入倍率
}

// Results 表示配方输出产品
type Results struct {
	Type                       *string  `json:"type,omitempty"`                           // 类型 item 或 fluid
	Name                       *string  `json:"name,omitempty"`                           // 产物名称
	Amount                     *float64 `json:"amount,omitempty"`                         // 产出数量
	FluidboxIndex              *float64 `json:"fluidbox_index,omitempty"`                 // 液体输出口索引
	IgnoredByStats             *float64 `json:"ignored_by_stats,omitempty"`               // 不计入生产统计
	IgnoredByProductivity      *float64 `json:"ignored_by_productivity,omitempty"`        // 不受生产力模块影响
	Probability                *float64 `json:"probability,omitempty"`                    // 产出概率（随机产物）
	PercentSpoiled             *float64 `json:"percent_spoiled,omitempty"`                // 腐坏比例
	Temperature                *float64 `json:"temperature,omitempty"`                    // 输出液体温度
	ShowDetailsInRecipeTooltip *bool    `json:"show_details_in_recipe_tooltip,omitempty"` // tooltip显示详细信息
	ExtraCountFraction         *float64 `json:"extra_count_fraction,omitempty"`           // 额外数量小数部分
}
