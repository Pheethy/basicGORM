package models

type Categories struct {
	Id    int64  `json:"id" gorm:"column:id;primaryKey"`
	Title string `json:"title" gorm:"column:title" validate:"required"`

	ProductId GUIDWrapper `json:"-" gorm:"column:product_id"`
}
