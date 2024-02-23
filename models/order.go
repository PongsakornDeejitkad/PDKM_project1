package models

import (
	"time"

	"gorm.io/gorm"
)

type Orders struct {
	ID           int            `json:"id" gorm:"primary_key"`
	CustomerId   int            `json:"customer_id" gorm:"integer ; not null"`
	CourierId    int            `json:"courier_id" gorm:"integer"`
	PaymentId    int            `json:"payment_id" gorm:"integer"`
	OrderAddress string         `json:"order_address" gorm:"varchar"`
	OrderEmail   string         `json:"order_email" gorm:"varchar"`
	OrderStatus  string         `json:"order_status" gorm:"varchar"`
	OrderDate    time.Time      `json:"order_date" gorm:"default:NOW()"`
	CreatedAt    time.Time      `json:"created_at" gorm:"default:NOW()"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"default:Now()"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	OrderItems []OrderItems `json:"-" gorm:"foreignKey:OrderId"`
}

type OrderItems struct {
	ID        int       `json:"id" gorm:"primary_key"`
	OrderId   int       `json:"order_id" gorm:"integer"`
	ProductId int       `json:"product_id" gorm:"integer"`
	Quantity  int       `json:"quantity" gorm:"integer"`
	CreatedAt time.Time `json:"created_at" gorm:"default:NOW()"`

	Order Orders `json:"-" gorm:"foreignKey:OrderId"`
}

type PaymentDetails struct {
	ID           int       `json:"id" gorm:"primary_key"`
	CustomerId   int       `json:"customer_id" gorm:"integer"`
	Amount       int       `json:"amount" gorm:"integer"`
	Provider     int       `json:"provider" gorm:"integer"`
	Status       string    `json:"status" gorm:"varchar"`
	ModifiedDate time.Time `json:"updated_at" gorm:"default:Now()"`
	CreatedAt    time.Time `json:"created_at" gorm:"default:NOW()"`

	Order Orders `json:"-" gorm:"foreignKey:CustomerId"`
}

type CourierDetails struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Name        string `json:"name" gorm:"varchar"`
	CourierType string `json:"courier_type" gorm:"varchar"`

	Order Orders `json:"-" gorm:"foreignKey:CourierId"`
}
