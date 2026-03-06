package rawDataModels

// 流体定义（Fluid Prototype）
type Fluid struct {
	Type               string   `json:"type"`                      // 原型类型，通常为 "fluid"
	Name               string   `json:"name"`                      // 流体唯一名称 ID（游戏内部引用）
	LocalisedName      []string `json:"localised_name,omitempty"`  // 本地化名称（多语言显示）
	Icon               string   `json:"icon"`                      // 流体图标文件路径
	Hidden             *bool    `json:"hidden,omitempty"`          // 是否在游戏界面中隐藏该流体
	DefaultTemperature float64  `json:"default_temperature"`       // 默认温度（单位：°C），流体生成时的初始温度
	MaxTemperature     *float64 `json:"max_temperature,omitempty"` // 最大允许温度（超过此温度可能被限制或无法继续加热）
	AutoBarrel         *bool    `json:"auto_barrel,omitempty"`     // 是否允许自动生成装桶/解桶配方
	Subgroup           *string  `json:"subgroup,omitempty"`        // 所属子分类（用于制作菜单分组）
	Parameter          *bool    `json:"parameter,omitempty"`       // 是否为参数化流体（用于某些动态原型系统）
	Order              *string  `json:"order,omitempty"`           // 排序字符串，用于 UI 中排序
	HeatCapacity       *string  `json:"heat_capacity,omitempty"`   // 热容量（例如 "0.2KJ"），影响温度变化所需能量
	GasTemperature     *float64 `json:"gas_temperature,omitempty"` // 气化温度，超过该温度时视为气体状态
	FuelValue          *string  `json:"fuel_value,omitempty"`      // 燃料能量值（若该流体可作为燃料）
}
