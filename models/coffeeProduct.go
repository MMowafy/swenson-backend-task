package models

import (
	"time"
)

type CoffeeProduct struct {
	Id                  string    `gorm:"column:id;primary_key:true;type:uuid;default:uuid_generate_v4()" json:"id"`
	ProductType         string    `gorm:"column:productType;index:productType" json:"productType"`
	Sku                 string    `gorm:"column:sku;index:sku" json:"sku"`
	WaterLineCompatible bool      `gorm:"column:waterLineCompatible;default:false;index:waterLineCompatible" json:"waterLineCompatible"`
	Flavor              string    `gorm:"column:flavor;index:flavor" json:"flavor"`
	PackSize            int       `gorm:"column:packSize;index:packSize" json:"packSize"`
	CreatedAt           time.Time `gorm:"column:createdAt;index:createdAt" json:"createdAt" `
	UpdatedAt           time.Time `gorm:"column:updatedAt;index:updatedAt" json:"updatedAt" `
}

func NewCoffeeProduct() *CoffeeProduct {
	return &CoffeeProduct{}
}
