package models

//////////////////////////////////////////////////////////////
// 基础实体（可用于箱子 / 建筑 / 设备）
//////////////////////////////////////////////////////////////

// Container 游戏实体定义（精简版）
// 用于AI工厂规划与蓝图生成
type Container struct {
	Type          string   `json:"type"`                     // 实体类型（container / furnace / assembler 等）
	Name          string   `json:"name"`                     // 实体名称
	LocalisedName []string `json:"localised_name,omitempty"` // 本地化名称
	Icon          *string  `json:"icon,omitempty"`           // 图标
	Icons         []Icons  `json:"icons,omitempty"`          // 多图标
	Flags         []string `json:"flags,omitempty"`          // 实体标志
	Hidden        *bool    `json:"hidden,omitempty"`         // 是否隐藏

	CollisionBox                [][]float64 `json:"collision_box,omitempty"`                  // 碰撞盒（影响建筑摆放）
	SelectionBox                [][]float64 `json:"selection_box,omitempty"`                  // 选择框
	DrawingBoxVerticalExtension *float64    `json:"drawing_box_vertical_extension,omitempty"` // 绘制框额外高度

	InventorySize       *float64             `json:"inventory_size,omitempty"`       // 库存大小
	InventoryType       *string              `json:"inventory_type,omitempty"`       // 库存类型`
	InventoryProperties *InventoryProperties `json:"inventory_properties,omitempty"` // 库存属性

	FastReplaceableGroup *string `json:"fast_replaceable_group,omitempty"` // 快速替换组

	SurfaceConditions []SurfaceConditions `json:"surface_conditions,omitempty"` // 表面建造条件

	Order *string `json:"order,omitempty"` // 排序

}

// InventoryProperties 库存属性
type InventoryProperties struct {
	StackSizeMin float64 `json:"stack_size_min"` // 最小堆叠
}
