package dao

type ItemsT struct {
	Name         string `json:"name"`
	Icon         string `json:"icon"`
	NameZh       string `json:"name_zh"`
	Hidden       int    `json:"hidden"`
	StackSize    int    `json:"stack_size"`
	Weight       int    `json:"weight"`
	Subgroup     string `json:"subgroup"`
	PlaceResult  string `json:"place_result"`
	Order        string `json:"item_order"`
	FuelCategory string `json:"fuel_category"`
	FuelValue    string `json:"fuel_value"`
	BurntResult  string `json:"burnt_result"`
	SpoilTicks   int    `json:"spoil_ticks"`
	SpoilResult  string `json:"spoil_result"`
}
