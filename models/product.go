package models

import "time"

type Product struct {
	ID        uint      `gorm:"column:id;type:int(11);primaryKey"`
	Code      string    `gorm:"column:code;type:varchar(10);not null;default:'';comment:编号"`
	Price     uint      `gorm:"column:price;type:int(11);not null;default:8;comment:价格"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
