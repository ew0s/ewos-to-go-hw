package main

type param = string

const (
	name        param = "name"
	description param = "description"
	price       param = "price"
	location    param = "location"
	delivery    param = "delivery"
)

type paramData = string

const (
	paramDataName        paramData = "💬 Название: станок"
	paramDataDescription paramData = "📖 Описание: станок для дерева"
	paramDataPrice       paramData = "💵 Цена: 100$"
	paramDataLocation    paramData = "📍 Локация: Казань"
	paramDataDelivery    paramData = "📦 Доставка: имеется"
)

func getParamData(p param) (*paramData, error) {
	data, ok := map[param]paramData{
		name:        paramDataName,
		description: paramDataDescription,
		price:       paramDataPrice,
		location:    paramDataLocation,
		delivery:    paramDataDelivery,
	}[p]
	if !ok {
		return nil, errorInvalidParamData
	}

	if len([]rune(data)) > width {
		return nil, errorParamDataLen
	}

	return &data, nil
}
