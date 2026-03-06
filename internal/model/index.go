package model

import rawDataModels "processBlueprint/internal/model/rawData"

type RawData struct {
	Item              map[string]rawDataModels.Item              `json:"item"`               // 物品
	Recipe            map[string]rawDataModels.Recipe            `json:"recipe"`             // 配方
	Technology        map[string]rawDataModels.Technology        `json:"technology"`         // 科技
	Resource          map[string]rawDataModels.Resource          `json:"resource"`           // 资源
	Fluid             map[string]rawDataModels.Fluid             `json:"fluid"`              // 流体
	MiningDrill       map[string]rawDataModels.MiningDrill       `json:"mining-drill"`       // 矿机
	Furnace           map[string]rawDataModels.Furnace           `json:"furnace"`            // 熔炉
	AssemblingMachine map[string]rawDataModels.AssemblingMachine `json:"assembling-machine"` // 工厂
	TransportBelt     map[string]rawDataModels.TransportBelt     `json:"transport-belt"`     // 传送带
	UndergroundBelt   map[string]rawDataModels.UndergroundBelt   `json:"underground-belt"`   // 地下传送带
	Splitter          map[string]rawDataModels.Splitter          `json:"splitter"`           // 分流器
	Pipe              map[string]rawDataModels.Pipe              `json:"pipe"`               // 管道
	Inserter          map[string]rawDataModels.Inserter          `json:"inserter"`           // 机械臂
	Beacon            map[string]rawDataModels.Beacon            `json:"beacon"`             // 插件塔
	Lab               map[string]rawDataModels.Lab               `json:"lab"`                // 实验室
	Roboport          map[string]rawDataModels.Roboport          `json:"roboport"`           // 机器人站点
	Container         map[string]rawDataModels.Container         `json:"container"`          // 箱子
}
