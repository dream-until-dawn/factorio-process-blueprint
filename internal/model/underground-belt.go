package models

// 碰撞层
type CollisionMask struct {
	Layers map[string]bool `json:"layers"` // 碰撞层定义
}

// 地下传送带（Underground Belt）
type UndergroundBelt struct {

	// ===== 基础信息 =====
	Type          string       `json:"type"`                     // 原型类型
	Name          string       `json:"name"`                     // 建筑名称
	LocalisedName *interface{} `json:"localised_name,omitempty"` // 本地化名称（多语言字符串数组）
	Icon          string       `json:"icon"`                     // 图标路径

	// ===== 碰撞盒 =====
	CollisionBox [][]float64 `json:"collision_box"` // 碰撞盒
	SelectionBox [][]float64 `json:"selection_box"` // 选择盒

	// ===== 升级链 =====
	FastReplaceableGroup string  `json:"fast_replaceable_group,omitempty"` // 快速替换组
	NextUpgrade          *string `json:"next_upgrade,omitempty"`           // 下一升级建筑

	// ===== 功能参数 =====
	Speed       float64 `json:"speed"`        // 传送带速度
	MaxDistance float64 `json:"max_distance"` // 地下最大连接距离

	// ===== 渲染资源 =====
	Structure         map[string]Sprite `json:"structure,omitempty"`           // 建筑结构贴图
	UndergroundSprite Sprite            `json:"underground_sprite,omitempty"`  // 地下入口贴图
	RemoveBeltsSprite *Sprite           `json:"remove_belts_sprite,omitempty"` // 移除提示贴图

	// ===== 碰撞层 =====
	UndergroundCollisionMask *CollisionMask `json:"underground_collision_mask,omitempty"`
}
