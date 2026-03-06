package models

// =============================
// 基础结构
// =============================

// Vector2 二维坐标
type Vector2 []float64

// Box 碰撞盒 / 选择盒
type Box [][]float64

// 图标定义（支持多层图标）
type Icons struct {
	Icon     string       `json:"icon"`                // 图标文件路径
	Scale    *float64     `json:"scale,omitempty"`     // 图标缩放比例
	Shift    []float64    `json:"shift,omitempty"`     // 图标偏移量 [x,y]
	Tint     *interface{} `json:"tint,omitempty"`      // 图标颜色着色（可能为RGBA数组或其他格式）
	IconSize *float64     `json:"icon_size,omitempty"` // 图标尺寸
	Floating *bool        `json:"floating,omitempty"`
}

// Sprite 精灵贴图
type Sprite struct {
	Filename   string   `json:"filename"`              // 图片路径
	Width      float64  `json:"width"`                 // 宽度
	Height     float64  `json:"height"`                // 高度
	Scale      float64  `json:"scale"`                 // 缩放
	Shift      Vector2  `json:"shift"`                 // 偏移
	Priority   *string  `json:"priority,omitempty"`    // 渲染优先级
	LineLength *float64 `json:"line_length,omitempty"` // 动画行长度
}

// SpriteLayer 多层贴图
type SpriteLayer struct {
	Layers []Sprite `json:"layers"`
}

// SurfaceConditions 表示配方只能在特定星球或表面环境下使用（Space Age扩展新增）
type SurfaceConditions struct {
	Property string   `json:"property"`      // 条件属性，例如 gravity、pressure 等
	Min      *float64 `json:"min,omitempty"` // 最小值限制
	Max      *float64 `json:"max,omitempty"` // 最大值限制
}

// 管道连接点
type PipeConnection struct {
	Position []float64 `json:"position"`       // 管道连接坐标
	Type     *string   `json:"type,omitempty"` // 连接类型（input/output）
}

// 能源来源定义
type EnergySource struct {
	Type              string   `json:"type"`                          // 能源类型（electric / burner / fluid）
	UsagePriority     *string  `json:"usage_priority,omitempty"`      // 用电优先级
	FuelCategories    []string `json:"fuel_categories,omitempty"`     // 可使用燃料类型
	Effectivity       *float64 `json:"effectivity,omitempty"`         // 能源效率
	FuelInventorySize *float64 `json:"fuel_inventory_size,omitempty"` // 燃料槽大小
}

// 流体箱（流体网络核心）
type FluidBox struct {
	// 输入 / 输出类型
	ProductionType  string           `json:"production_type"`
	Volume          *float64         `json:"volume,omitempty"`           // 容量
	BaseArea        *float64         `json:"base_area,omitempty"`        // 基础面积
	BaseLevel       *float64         `json:"base_level,omitempty"`       // 基础流体高度
	PipeConnections []PipeConnection `json:"pipe_connections,omitempty"` // 管道连接
}
