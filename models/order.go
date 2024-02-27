package models

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	ID         int            `json:"id" gorm:"primary_key"`
	CustomerId int            `json:"customer_id" gorm:"not null;integer"`
	CourierId  int            `json:"courier_id" gorm:"integer"`
	PaymentId  int            `json:"payment_id" gorm:"integer"`
	Address    string         `json:"address" gorm:"varchar"`
	Email      string         `json:"email" gorm:"varchar"`
	Status     string         `json:"status" gorm:"varchar"`
	Date       time.Time      `json:"date" gorm:"default:NOW()"`
	CreatedAt  time.Time      `json:"created_at" gorm:"default:NOW()"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"default:NOW()"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	OrderItem []OrderItems `json:"-" gorm:"foreignKey:OrderId"`
	Customer  Customers    `json:"-" gorm:"foreignKey:CustomerId"`
}

type OrderItems struct {
	ID        int       `json:"id" gorm:"primary_key"`
	OrderId   int       `json:"order_id" gorm:"integer"`
	ProductId int       `json:"product_id" gorm:"integer"`
	Quantity  int       `json:"quantity" gorm:"integer"`
	CreatedAt time.Time `json:"created_at" gorm:"default:NOW()"`

	Order   Orders   `json:"-" gorm:"foreignKey:OrderId"`
	Product Products `json:"-" gorm:"foreignKey:ProductId"`
}

type PaymentDetails struct {
	ID        int       `json:"id" gorm:"primary_key"`
	PaymentId int       `json:"payment_id" gorm:"unique;not null;integer"`
	Amount    int       `json:"amount" gorm:"integer"`
	Provider  int       `json:"provider" gorm:"integer"`
	Status    string    `json:"status" gorm:"varchar"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:NOW()"`
	CreatedAt time.Time `json:"created_at" gorm:"default:NOW()"`

	Order Orders `json:"-" gorm:"foreignKey:PaymentId"`
}

type CourierDetails struct {
	ID          int    `json:"id" gorm:"primary_key"`
	CourierId   int    `json:"courier_id" gorm:"integer"`
	Name        string `json:"name" gorm:"varchar"`
	CourierType string `json:"courier_type" gorm:"varchar"`

	Order Orders `json:"-" gorm:"foreignKey:CourierId"`
}
