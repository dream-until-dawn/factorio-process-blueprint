package rawDataModels

// TransportBelt 传送带原型
type TransportBelt struct {
	Type          string       `json:"type"`                     // 原型类型
	Name          string       `json:"name"`                     // 名称
	LocalisedName *interface{} `json:"localised_name,omitempty"` // 本地化名称（多语言字符串数组）
	Icon          string       `json:"icon"`                     // 图标
	Flags         []string     `json:"flags"`                    // 标志

	CollisionBox Box `json:"collision_box"` // 碰撞盒
	SelectionBox Box `json:"selection_box"` // 选择盒

	Speed float64 `json:"speed"` // 传送带速度 Factorio单位：tile/tick

	FastReplaceableGroup   string `json:"fast_replaceable_group"`   // 快速替换组
	NextUpgrade            string `json:"next_upgrade,omitempty"`   // 下一个升级
	RelatedUndergroundBelt string `json:"related_underground_belt"` // 对应地下传送带

	HeatingEnergy string `json:"heating_energy,omitempty"` // 加热能量（极寒星球）
}
