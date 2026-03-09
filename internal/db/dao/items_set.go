package dao

import (
	"fmt"
	rawDataModels "processBlueprint/internal/model/rawData"
	"processBlueprint/internal/utils"
)

func (d *DBDAO) InsertItemsTable(items map[string]rawDataModels.Item) error {

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
	_, err = tx.Exec(`DELETE FROM items`)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`
	INSERT INTO items (
		name,
		icon,
		name_zh,
		hidden,
		stack_size,
		weight,
		subgroup,
		place_result,
		item_order,
		fuel_category,
		fuel_value,
		burnt_result,
		spoil_ticks,
		spoil_result
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
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
			name_zh = fmt.Sprintf("item-name.%s", item.Name)
		}

		hidden := 0
		if item.Hidden != nil {
			if *item.Hidden {
				hidden = 1
			} else {
				hidden = -1
			}
		}

		subgroup := ""
		if item.Subgroup != nil {
			subgroup = *item.Subgroup
		}

		placeResult := ""
		if item.PlaceResult != nil {
			placeResult = *item.PlaceResult
		}

		order := ""
		if item.Order != nil {
			order = *item.Order
		}

		fuelCategory := ""
		if item.FuelCategory != nil {
			fuelCategory = *item.FuelCategory
		}

		fuelValue := ""
		if item.FuelValue != nil {
			fuelValue = *item.FuelValue
		}

		burntResult := ""
		if item.BurntResult != nil {
			burntResult = *item.BurntResult
		}

		spoilTicks := 0
		if item.SpoilTicks != nil {
			spoilTicks = int(*item.SpoilTicks)
		}

		spoilResult := ""
		if item.SpoilResult != nil {
			spoilResult = *item.SpoilResult
		}

		weight := -1
		if item.Weight != nil {
			weight = int(*item.Weight)
		}

		_, err = stmt.Exec(
			item.Name,
			icon,
			name_zh,
			hidden,
			int(item.StackSize),
			weight,
			subgroup,
			placeResult,
			order,
			fuelCategory,
			fuelValue,
			burntResult,
			spoilTicks,
			spoilResult,
		)

		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
