package domain

import "time"

type Product struct {
	Id           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

func NewProduct(Id uint, name, serialNumber string) *Product {
	return &(Product{Id: Id, CreatedAt: time.Now(), Name: name, SerialNumber: serialNumber})
}
