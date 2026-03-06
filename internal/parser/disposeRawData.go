package parser

import (
	"encoding/json"
	"fmt"
	"os"

	"processBlueprint/internal/model"
	"processBlueprint/internal/utils"
)

func readJSONFile(filename string) (*model.RawData, error) {
	// 读取文件
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("读取文件失败: %w", err)
	}

	// 解析 JSON
	var rawData model.RawData
	err = json.Unmarshal(data, &rawData)
	if err != nil {
		return nil, fmt.Errorf("解析 JSON 失败: %w", err)
	}

	return &rawData, nil
}

func DisposeRawData() {
	rawData, err := readJSONFile("rawData/data-raw-dump.json")
	if err != nil {
		fmt.Printf("错误: %v\n", err)
		return
	}

	fmt.Printf("解析成功!\n")
	// fmt.Printf("Item: %+v\n", rawData.Item)
	utils.SaveToJSON("env/Item.json", rawData.Item, true)
	utils.SaveToJSON("env/Recipe.json", rawData.Recipe, true)
	utils.SaveToJSON("env/Technology.json", rawData.Technology, true)
	utils.SaveToJSON("env/Resource.json", rawData.Resource, true)
	utils.SaveToJSON("env/Fluid.json", rawData.Fluid, true)
	utils.SaveToJSON("env/MiningDrill.json", rawData.MiningDrill, true)
	utils.SaveToJSON("env/Furnace.json", rawData.Furnace, true)
	utils.SaveToJSON("env/AssemblingMachine.json", rawData.AssemblingMachine, true)
	utils.SaveToJSON("env/TransportBelt.json", rawData.TransportBelt, true)
	utils.SaveToJSON("env/UndergroundBelt.json", rawData.UndergroundBelt, true)
	utils.SaveToJSON("env/Splitter.json", rawData.Splitter, true)
	utils.SaveToJSON("env/Pipe.json", rawData.Pipe, true)
	utils.SaveToJSON("env/Inserter.json", rawData.Inserter, true)
	utils.SaveToJSON("env/Beacon.json", rawData.Beacon, true)
	utils.SaveToJSON("env/Lab.json", rawData.Lab, true)
	utils.SaveToJSON("env/Roboport.json", rawData.Roboport, true)
	utils.SaveToJSON("env/Container.json", rawData.Container, true)
}
