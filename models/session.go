package models

import (
	"time"
)

type Sessions struct {
	ID        int       `json:"id" gorm:"primary_key"`
	UserId    int       `json:"user_id" gorm:"integer"`
	Total     int       `json:"total" gorm:"Decimal"`
	CreatedAt time.Time `json:"created_at" gorm:"default:NOW()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:Now()"`

	CartItems []CartItems `json:"-" gorm:"foreignKey:SessionId"`
	Customers Customers   `json:"-" gorm:"foreignKey:UserId"`
}

type CartItems struct {
	ID        int       `json:"id" gorm:"primary_key"`
	SessionId int       `json:"session_id" gorm:"integer"`
	ProductId int       `json:"product_id" gorm:"integer"`
	Quantity  int       `json:"quantity" gorm:"integer"`
	CreatedAt time.Time `json:"created_at" gorm:"default:NOW()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:Now()"`

	Session Sessions `json:"-" gorm:"foreignKey:SessionId"`
	Product Products `json:"-" gorm:"foreignKey:ProductId"`
}
