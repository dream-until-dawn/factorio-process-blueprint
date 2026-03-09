package dao

import (
	"fmt"
	rawDataModels "processBlueprint/internal/model/rawData"
	"processBlueprint/internal/utils"
)

func (d *DBDAO) InsertRecipeTable(items map[string]rawDataModels.Recipe) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}

	// 出错自动回滚
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// 清空旧数据
	_, err = tx.Exec(`DELETE FROM recipes`)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`
	INSERT INTO recipes (
		name,
		icon,
		name_zh,
		hidden,
		ingredients,
		results,
		category,
		energy_required,
		allow_productivity,
		allow_quality,
		allow_decomposition,
		reset_freshness_on_craft,
		preserve_products_in_machine_output,
		auto_recycle,
		surface_conditions
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range items {

		// 处理指针字段
		icon := ""
		if item.Icon != nil {
			imgName := utils.ExtractNameFromPath(*item.Icon)
			icon = fmt.Sprintf("https://wiki.factorio.com/images/thumb/%s.png/32px-%s.png", imgName, imgName)
		}

		name_zh := ""
		if item.LocalisedName != nil {
			name_zh = utils.FlattenToString(*item.LocalisedName)
		} else {
			name_zh = fmt.Sprintf("recipe-name.%s", item.Name)
		}

		hidden := 0
		if item.Hidden != nil {
			if *item.Hidden {
				hidden = 1
			} else {
				hidden = -1
			}
		}

		ingredients := rawDataModels.IngredientsList{}
		if item.Ingredients != nil {
			ingredients = *item.Ingredients
		}

		results := rawDataModels.ResultsList{}
		if item.Results != nil {
			results = *item.Results
		}

		category := ""
		if item.Category != nil {
			category = *item.Category
		}

		energy_required := 0.0
		if item.EnergyRequired != nil {
			energy_required = *item.EnergyRequired
		}

		allow_productivity := 0
		if item.AllowProductivity != nil {
			if *item.AllowProductivity {
				allow_productivity = 1
			} else {
				allow_productivity = -1
			}
		}

		allow_quality := 0
		if item.AllowQuality != nil {
			if *item.AllowQuality {
				allow_quality = 1
			} else {
				allow_quality = -1
			}
		}

		allow_decomposition := 0
		if item.AllowDecomposition != nil {
			if *item.AllowDecomposition {
				allow_decomposition = 1
			} else {
				allow_decomposition = -1
			}
		}

		reset_freshness_on_craft := 0
		if item.ResetFreshnessOnCraft != nil {
			if *item.ResetFreshnessOnCraft {
				reset_freshness_on_craft = 1
			} else {
				reset_freshness_on_craft = -1
			}
		}

		preserve_products_in_machine_output := 0
		if item.PreserveProductsInMachineOutput != nil {
			if *item.PreserveProductsInMachineOutput {
				preserve_products_in_machine_output = 1
			} else {
				preserve_products_in_machine_output = -1
			}
		}

		auto_recycle := 0
		if item.AutoRecycle != nil {
			if *item.AutoRecycle {
				auto_recycle = 1
			} else {
				auto_recycle = -1
			}
		}

		surface_conditions := utils.SurfaceMask("111111")
		if item.SurfaceConditions != nil {
			surface_conditions = utils.GetAllowedSurface(*item.SurfaceConditions)
		}

		_, err = stmt.Exec(
			item.Name,
			icon,
			name_zh,
			hidden,
			ingredients,
			results,
			category,
			energy_required,
			allow_productivity,
			allow_quality,
			allow_decomposition,
			reset_freshness_on_craft,
			preserve_products_in_machine_output,
			auto_recycle,
			surface_conditions,
		)

		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
