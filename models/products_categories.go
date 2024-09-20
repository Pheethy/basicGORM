package models

import "github.com/Microsoft/go-winio/pkg/guid"

type ProductsCatetories struct {
	ProductId  guid.GUID `json:"product_id" gorm:"column:product_id;primaryKey"`
	CategoryId int64     `json:"category_id" gorm:"column:category_id;primaryKey"`
}
