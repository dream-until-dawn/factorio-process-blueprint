package rawDataModels

// Roboport 机器人网络基站（物流机器人/建设机器人停靠与充电中心）
type Roboport struct {
	Type  *string  `json:"type,omitempty"`  // 原型类型（通常为 "roboport"）
	Name  *string  `json:"name,omitempty"`  // 实体名称
	Icon  *string  `json:"icon,omitempty"`  // 图标路径（仅前端展示用）
	Flags []string `json:"flags,omitempty"` // 实体标志（placeable-player / player-creation 等）

	FastReplaceableGroup *string `json:"fast_replaceable_group,omitempty"` // 可快速替换分组

	CollisionBox [][]float64 `json:"collision_box,omitempty"` // 碰撞盒（影响建筑摆放）
	SelectionBox [][]float64 `json:"selection_box,omitempty"` // 选择框（玩家选中范围）

	EnergySource    *EnergySource `json:"energy_source,omitempty"`    // 能源来源
	RechargeMinimum *string       `json:"recharge_minimum,omitempty"` // 最低充电能量
	EnergyUsage     *string       `json:"energy_usage,omitempty"`     // 能源使用量

	ChargingEnergy     *string  `json:"charging_energy,omitempty"`     // 单机器人充电功率
	LogisticsRadius    *float64 `json:"logistics_radius,omitempty"`    // 物流网络范围
	ConstructionRadius *float64 `json:"construction_radius,omitempty"` // 建造网络范围

	ChargeApproachDistance *float64    `json:"charge_approach_distance,omitempty"` // 机器人接近充电距离
	RobotSlotsCount        *float64    `json:"robot_slots_count,omitempty"`        // 机器人槽位数量
	MaterialSlotsCount     *float64    `json:"material_slots_count,omitempty"`     // 材料槽位数量
	StationingOffset       []float64   `json:"stationing_offset,omitempty"`        // 机器人停靠位置偏移
	ChargingOffsets        [][]float64 `json:"charging_offsets,omitempty"`         // 机器人充电位置

	DefaultAvailableLogisticOutputSignal     *SignalID `json:"default_available_logistic_output_signal,omitempty"`     // 可用物流机器人信号
	DefaultTotalLogisticOutputSignal         *SignalID `json:"default_total_logistic_output_signal,omitempty"`         // 总物流机器人信号
	DefaultAvailableConstructionOutputSignal *SignalID `json:"default_available_construction_output_signal,omitempty"` // 可用建设机器人信号
	DefaultTotalConstructionOutputSignal     *SignalID `json:"default_total_construction_output_signal,omitempty"`     // 总建设机器人信号
	DefaultRoboportCountOutputSignal         *SignalID `json:"default_roboport_count_output_signal,omitempty"`         // roboport 数量信号

	HeatingEnergy     *string             `json:"heating_energy,omitempty"`     // 加热能量
	SurfaceConditions []SurfaceConditions `json:"surface_conditions,omitempty"` // 表面条件（某些星球可否建造）
}

// SignalID 信号ID
type SignalID struct {
	Type string `json:"type"` // 信号类型（item / fluid / virtual）
	Name string `json:"name"` // 信号名称
}
