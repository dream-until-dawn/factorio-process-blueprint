package models

// 电路信号
type Signal struct {
	Type string `json:"type"` // 信号类型 item/fluid/virtual
	Name string `json:"name"` // 信号名称
}

// 电路条件（统一结构）
type CircuitCondition struct {
	First      Signal  `json:"first"`      // 左侧信号
	Comparator string  `json:"comparator"` // 比较符号 > < =
	Second     float64 `json:"second"`     // 右侧值
}

// 电路条件集合
type CircuitConditions struct {
	InputLeft   *CircuitCondition `json:"input_left,omitempty"`
	InputRight  *CircuitCondition `json:"input_right,omitempty"`
	OutputLeft  *CircuitCondition `json:"output_left,omitempty"`
	OutputRight *CircuitCondition `json:"output_right,omitempty"`
}

// Splitter（分流器建筑）
type Splitter struct {
	// ===== 基础信息 =====
	Type          string       `json:"type"`                     // 原型类型
	Name          string       `json:"name"`                     // 建筑名称
	LocalisedName *interface{} `json:"localised_name,omitempty"` // 本地化名称（多语言字符串数组）
	Icon          string       `json:"icon"`                     // 图标路径
	// ===== 建筑属性 =====
	MaxHealth float64 `json:"max_health"` // 最大生命值
	// ===== 碰撞与选择 =====
	CollisionBox [][]float64 `json:"collision_box"` // 碰撞盒
	SelectionBox [][]float64 `json:"selection_box"` // 选择盒
	// ===== 升级链 =====
	FastReplaceableGroup string  `json:"fast_replaceable_group,omitempty"` // 快速替换组
	NextUpgrade          *string `json:"next_upgrade,omitempty"`           // 下一升级建筑
	// ===== 功能参数 =====
	Speed                float64 `json:"speed"`                  // 传送带速度
	RelatedTransportBelt string  `json:"related_transport_belt"` // 关联传送带
	// ===== 电路网络 =====
	CircuitWireMaxDistance float64            `json:"circuit_wire_max_distance,omitempty"` // 电路线最大距离
	CircuitConditions      *CircuitConditions `json:"circuit_conditions,omitempty"`        // 默认电路条件
}
