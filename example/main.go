package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

type Product struct {
	ID        uint      `gorm:"column:id;type:int(11);primaryKey"`
	Code      string    `gorm:"column:code;type:varchar(10);not null;default:'';comment:编号"`
	Price     uint      `gorm:"column:price;type:int(11);not null;default:8;comment:价格"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:13306)/pyrethrum?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	DB.AutoMigrate(Product{})

	// Create
	DB.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	DB.First(&product)                    // 根据整型主键查找
	DB.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	DB.Model(&product).Update("Price", 200)
	// Update - 更新多个字段
	DB.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	DB.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//DB.Delete(&product, 1)
}
