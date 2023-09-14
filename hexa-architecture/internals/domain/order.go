package domain

import "time"

type Order struct {
	Id           uint `json:"id" gorm:"primaryKey"`
	CreatedAt    time.Time
	ProductRefer int     `json:"product_id"`
	Product      Product `gorm:"foreignKey:ProductRefer"`
	UserRefer    int     `json:"user_id"`
	User         User    `gorm:"foreignKey:UserRefer"`
}

func NewOrder(Id uint, product Product, user User) *Order {
	return &(Order{Id: Id, CreatedAt: time.Now(), Product: product, ProductRefer: int(product.Id), User: user, UserRefer: int(user.Id)})
}
