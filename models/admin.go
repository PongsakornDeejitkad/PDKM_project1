package models

import (
	"time"

	"gorm.io/gorm"
)

type Admins struct {
	ID        int            `json:"id" gorm:"primary_key"`
	Username  string         `json:"username" gorm:"not null;varchar(50)"`
	Password  string         `json:"password" gorm:"not null;type:text;size:200"`
	Firstname string         `json:"firstname" gorm:"varchar;not null"`
	Lastname  string         `json:"lastname" gorm:"varchar;not null"`
	Email     string         `json:"email" gorm:"varchar(100);not null;unique"`
	TypeID    int            `json:"type_id" gorm:"integer"`
	LastLogin time.Time      `json:"last_login" gorm:"default: NOW()"`
	CreatedAt time.Time      `json:"created_at" gorm:"default: NOW()"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"default: NOW()"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`

	Type AdminType `json:"type" gorm:"foreignKey:TypeID"`
}

type AdminType struct {
	ID         int       `json:"id" gorm:"primary_key"`
	AdminType  string    `json:"admin_type" gorm:"varchar"`
	Permission string    `json:"permission" gorm:"varchar"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:NOW()"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:NOW()"`

	Admin []Admins `json:"admin" gorm:"foreignKey:TypeID"`
}
