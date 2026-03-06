package rawDataModels

type Unit struct {
	Count        *float64        `json:"count,omitempty"`
	Ingredients  [][]interface{} `json:"ingredients"`
	Time         float64         `json:"time"`
	CountFormula *string         `json:"count_formula,omitempty"`
}

type Effects struct {
	Type                   string       `json:"type"`
	Recipe                 *string      `json:"recipe,omitempty"`
	AmmoCategory           *string      `json:"ammo_category,omitempty"`
	Modifier               *interface{} `json:"modifier,omitempty"` // conflict: float64(94.4%) bool(5.6%)
	TurretId               *string      `json:"turret_id,omitempty"`
	Icon                   *string      `json:"icon,omitempty"`
	Hidden                 *bool        `json:"hidden,omitempty"`
	Quality                *string      `json:"quality,omitempty"`
	IconSize               *float64     `json:"icon_size,omitempty"`
	SpaceLocation          *string      `json:"space_location,omitempty"`
	UseIconOverlayConstant *bool        `json:"use_icon_overlay_constant,omitempty"`
	Change                 *float64     `json:"change,omitempty"`
}

type ResearchTrigger struct {
	Type   string   `json:"type"`
	Item   *string  `json:"item,omitempty"`
	Count  *float64 `json:"count,omitempty"`
	Entity *string  `json:"entity,omitempty"`
}

type Technology struct {
	Type                     string           `json:"type"`
	Name                     string           `json:"name"`
	LocalisedName            *interface{}     `json:"localised_name,omitempty"` // 本地化名称（多语言字符串数组）
	Icon                     *string          `json:"icon,omitempty"`
	Effects                  []Effects        `json:"effects,omitempty"`
	ResearchTrigger          *ResearchTrigger `json:"research_trigger,omitempty"`
	Essential                *bool            `json:"essential,omitempty"`
	Prerequisites            []string         `json:"prerequisites,omitempty"`
	Unit                     *Unit            `json:"unit,omitempty"`
	Icons                    []Icons          `json:"icons,omitempty"`
	Upgrade                  *bool            `json:"upgrade,omitempty"`
	MaxLevel                 *string          `json:"max_level,omitempty"`
	IgnoreTechCostMultiplier *bool            `json:"ignore_tech_cost_multiplier,omitempty"`
	LocalisedDescription     []string         `json:"localised_description,omitempty"`
	Order                    *string          `json:"order,omitempty"`
}
