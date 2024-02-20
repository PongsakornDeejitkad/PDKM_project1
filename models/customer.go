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
	CustomerAddress  []CustomerAddress `json:"customer_address" gorm:"foreignKey:CustomerID; cascade; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type CustomerPayment struct {
	ID          int       `json:"id" gorm:"primary_key"`
	CustomerID  int       `json:"customer_id" gorm:"interger"`
	PaymentType string    `json:"payment_type" gorm:"varchar(100)"`
	Provider    string    `json:"provider" gorm:"varchar(100)"`
	AccountNo   int       `json:"account_no" gorm:"integer"`
	ExpireDate  time.Time `json:"expire_date" gorm:"integer"`
	CreatedAt   time.Time `json:"created_at" gorm:"default: NOW()"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default: NOW()"`

	Customer Customer `json:"customer" gorm:"foreignKey:CustomerID"`
}

type CustomerAddress struct {
	ID          int       `json:"id" gorm:"primary_key"`
	CustomerID  int       `json:"customer_id" gorm:"interger"`
	District    string    `json:"district" gorm:"varchar(100)"`
	SubDistrict string    `json:"sub_district" gorm:"varchar(100)"`
	City        string    `json:"city" gorm:"varchar(100)"`
	ZipCode     string    `json:"zip_code" gorm:"varchar(100)"`
	Country     string    `json:"country" gorm:"varchar(100)"`
	Telephone   int       `json:"telephone" gorm:"integer"`
	CreatedAt   time.Time `json:"created_at" gorm:"default: NOW()"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default: NOW()"`

	Customer Customer `json:"customer" gorm:"foreignKey:CustomerID"`
}
