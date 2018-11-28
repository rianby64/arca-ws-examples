package examples

import (
	"errors"

	"../arca"
)

// Good whatever
type Good struct {
	ID          int
	Description string
	Price       int
}

// Goods whatever
type Goods []Good

// This is my Data-Base!
var goods = Goods{
	Good{1, "Computer", 1000},
	Good{2, "Smartphone", 2000},
	Good{3, "Wine", 1500},
}
var lastGoodsID = len(goods)

var goodsCRUD = arca.DIRUD{
	Read: func(requestParams *interface{}, context *interface{},
		_ *arca.ResponseHandler) (interface{}, error) {
		return goods, nil
	},
	Update: func(requestParams *interface{}, context *interface{},
		_ *arca.ResponseHandler) (interface{}, error) {
		params := (*requestParams).(map[string]interface{})
		preid, ok := params["ID"]
		if !ok {
			return nil, errors.New("params in request doesn't contain ID")
		}
		preid2, ok := preid.(float64)
		if !ok {
			return nil, errors.New("ID in params isn't int")
		}

		id := int(preid2)
		for index, good := range goods {
			if good.ID == id {
				if description, ok := params["Description"]; ok {
					goods[index].Description = description.(string)
				}
				if price, ok := params["Price"]; ok && price != nil {
					preprice := price.(float64)
					goods[index].Price = int(preprice)
				}
				return goods[index], nil
			}
		}
		return nil, errors.New("nothing")
	},
	Insert: func(requestParams *interface{}, context *interface{},
		_ *arca.ResponseHandler) (interface{}, error) {
		params := (*requestParams).(map[string]interface{})
		lastGoodsID++
		newGood := Good{ID: lastGoodsID}
		if description, ok := params["Description"]; ok {
			newGood.Description = description.(string)
		}
		if price, ok := params["Price"]; ok && price != nil {
			preprice := price.(float64)
			newGood.Price = int(preprice)
		}
		goods = append(goods, newGood)
		return newGood, nil
	},
	Delete: func(requestParams *interface{}, context *interface{},
		_ *arca.ResponseHandler) (interface{}, error) {
		params := (*requestParams).(map[string]interface{})
		preid, ok := params["ID"]
		if !ok {
			return nil, errors.New("params in request doesn't contain ID")
		}
		preid2, ok := preid.(float64)
		if !ok {
			return nil, errors.New("ID in params isn't int")
		}

		id := int(preid2)
		deletedGood := Good{ID: id}
		for i, good := range goods {
			if good.ID == id {
				goods = append(goods[:i], goods[i+1:]...)
				return deletedGood, nil
			}
		}
		return nil, errors.New("nothing")
	},
}
