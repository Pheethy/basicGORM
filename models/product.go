package models

import (
	"log"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/Pheethy/psql/helper"
)

type Product struct {
	Id          GUIDWrapper       `json:"id" gorm:"column:id;primaryKey"`
	Title       string            `json:"title" gorm:"column:title"`
	Description string            `json:"description" gorm:"column:description"`
	Price       float64           `json:"price" gorm:"column:price"`
	CreatedAt   *helper.Timestamp `json:"created_at" gorm:"column:created_at"`
	UpdatedAt   *helper.Timestamp `json:"updated_at" gorm:"column:updated_at"`

	Categories *Categories `json:"categories" gorm:"foreignKey:ProductId;references:Id"`
	Images     []*Images   `json:"images" gorm:"foreignKey:ProductId;references:Id"`
}

func (p *Product) NewUUID() {
	id, err := guid.NewV4()
	if err != nil {
		log.Fatal(err)
	}
	p.Id = GUIDWrapper{
		GUID: id,
	}
}
