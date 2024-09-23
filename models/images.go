package models

import (
	"time"

	"github.com/Microsoft/go-winio/pkg/guid"
	"github.com/Pheethy/psql/helper"
)

type Images struct {
	Id        GUIDWrapper       `json:"id" gorm:"column:id;primaryKey" type:"uuid"`
	FileName  string            `json:"filename" gorm:"column:filename" type:"string"`
	URL       string            `json:"url" gorm:"column:url" type:"string"`
	ProductId GUIDWrapper       `json:"-" gorm:"column:product_id" type:"uuid"`
	CreatedAt *helper.Timestamp `json:"created_at" gorm:"column:created_at" type:"timestamp"`
	UpdatedAt *helper.Timestamp `json:"updated_at" gorm:"column:updated_at" type:"timestamp"`
}

func (p *Images) NewUUID() {
	id, _ := guid.NewV4()
	p.Id = GUIDWrapper{
		GUID: id,
	}
}

func (p *Images) SetCreatedAt() {
	ti := helper.NewTimestampFromTime(time.Now())
	p.CreatedAt = &ti
}

func (p *Images) SetUpdatedAt() {
	ti := helper.NewTimestampFromTime(time.Now())
	p.UpdatedAt = &ti
}
