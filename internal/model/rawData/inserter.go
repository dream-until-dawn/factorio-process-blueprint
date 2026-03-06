package rawDataModels

// 机械臂
type Inserter struct {

	// ===== 基础信息 =====
	Type string `json:"type"` // 实体类型
	Name string `json:"name"` // 名称
	Icon string `json:"icon"` // 图标
	// ===== 碰撞 =====
	CollisionBox [][]float64 `json:"collision_box"` // 碰撞盒
	SelectionBox [][]float64 `json:"selection_box"` // 选择盒
	// ===== 能源 =====
	EnergySource      EnergySource `json:"energy_source"`                 // 能源系统
	EnergyPerMovement string       `json:"energy_per_movement,omitempty"` // 每次移动能量消耗
	EnergyPerRotation string       `json:"energy_per_rotation,omitempty"` // 每次旋转能量消耗
	// ===== 速度 =====
	ExtensionSpeed float64 `json:"extension_speed"` // 伸展速度
	RotationSpeed  float64 `json:"rotation_speed"`  // 旋转速度
	// ===== 功能参数 =====
	FilterCount      float64  `json:"filter_count,omitempty"`      // 过滤器数量
	StackSizeBonus   *float64 `json:"stack_size_bonus,omitempty"`  // 堆叠抓取加成
	HandSize         *float64 `json:"hand_size,omitempty"`         // 机械臂手长
	StartingDistance *float64 `json:"starting_distance,omitempty"` // 起始抓取距离
	// ===== 抓取位置 =====
	PickupPosition []float64 `json:"pickup_position"` // 抓取位置
	InsertPosition []float64 `json:"insert_position"` // 放置位置
	// ===== 升级链 =====
	FastReplaceableGroup string  `json:"fast_replaceable_group,omitempty"` // 快速替换组
	NextUpgrade          *string `json:"next_upgrade,omitempty"`           // 下一升级
}
