package models

// 资源掉落
type ResourceResult struct {
	Name   string   `json:"name"`                 // 产出物品ID
	Amount *float64 `json:"amount,omitempty"`     // 固定产出数量
	Min    *float64 `json:"amount_min,omitempty"` // 最小产出
	Max    *float64 `json:"amount_max,omitempty"` // 最大产出
}

// Resource（矿物、石头、油田等）
type Resource struct {
	Type          string       `json:"type"`                     // 原型类型（resource）
	Name          string       `json:"name"`                     // 资源名称
	LocalisedName *interface{} `json:"localised_name,omitempty"` // 本地化名称（多语言字符串数组）
	Icon          string       `json:"icon"`                     // 图标
	Flags         []string     `json:"flags,omitempty"`          // 标志
	Order         string       `json:"order,omitempty"`          // UI排序

	CollisionBox [][]float64 `json:"collision_box"` // 碰撞盒
	SelectionBox [][]float64 `json:"selection_box"` // 选择盒

	Category *string `json:"category,omitempty"` // 资源类型
	Subgroup *string `json:"subgroup,omitempty"`

	// ======================
	// 无限资源
	Infinite                *bool    `json:"infinite,omitempty"`
	Minimum                 *float64 `json:"minimum,omitempty"`
	Normal                  *float64 `json:"normal,omitempty"`
	InfiniteDepletionAmount *float64 `json:"infinite_depletion_amount,omitempty"`

	// ======================
	// 矿区搜索
	ResourcePatchSearchRadius *float64 `json:"resource_patch_search_radius,omitempty"`
}
