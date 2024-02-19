package models

import "time"

type Customer struct {
	ID        int       `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"not null; varchar(100)"`
	Email     string    `json:"email" gorm:"not null; unique; varchar(100)"`
	Username  string    `json:"username" gorm:"not null; unique; varchar(50)"`
	Password  string    `json:"password" gorm:"not null; varchar(200)"`
	CreatedAt time.Time `json:"created_at" gorm:"default: NOW()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default: NOW()"`

	CustomerPayments []CustomerPayment `json:"customer_payments" gorm:"foreignKey:CustomerID; cascade; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CustomerAddress []CustomerAddress `json:"customer_address" gorm:"foreignKey:CustomerID; cascade; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type CustomerPayment struct {
	ID          int       `json:"id" gorm:"primary_key"`
	CustomerID  int       `json:"customer_id"`
	PaymentType string    `json:"payment_type"`
	Provider    string    `json:"provider"`
	AccountNo   string    `json:"account_no"`
	ExpireDate  time.Time `json:"expire_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Customer Customer `json:"customer" gorm:"foreignKey:CustomerID"`
}

type CustomerAddress struct {
	ID          int    `json:"id" gorm:"primary_key"`
	CustomerID  int    `json:"customer_id"`
	District    string `json:"district"`
	SubDistrict string `json:"sub_district"`
	City        string `json:"city"`
	ZipCode     string `json:"zip_code"`
	Country     string `json:"country"`
	Telephone   string `json:"telephone"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Customer Customer `json:"customer" gorm:"foreignKey:CustomerID"`
}
