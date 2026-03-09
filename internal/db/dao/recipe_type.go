package dao

import (
	rawDataModels "processBlueprint/internal/model/rawData"
)

type RecipeT struct {
	Name                            string                        `json:"name"`
	Icon                            string                        `json:"icon"`
	NameZh                          string                        `json:"name_zh"`
	Hidden                          int                           `json:"hidden"`
	Ingredients                     rawDataModels.IngredientsList `json:"ingredients"`
	Results                         rawDataModels.ResultsList     `json:"results"`
	Category                        string                        `json:"category"`
	EnergyRequired                  float64                       `json:"energy_required"`
	AllowProductivity               int                           `json:"allow_productivity"`
	AllowQuality                    int                           `json:"allow_quality"`
	AllowDecomposition              int                           `json:"allow_decomposition"`
	ResetFreshnessOnCraft           int                           `json:"reset_freshness_on_craft"`
	PreserveProductsInMachineOutput int                           `json:"preserve_products_in_machine_output"`
	AutoRecycle                     int                           `json:"auto_recycle"`
	SurfaceConditions               int                           `json:"surface_conditions"`
}
