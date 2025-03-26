package dbmodels

import (
	"gorm.io/gorm"
)

// User represents the database model for storing user details
type User struct {
	gorm.Model
	UserName    string `gorm:"column:username;type:varchar(255);not null;unique"`
	EmailID     string `gorm:"column:email;type:varchar(255);not null;unique"`
	PhoneNumber uint64 `gorm:"column:phone_no;type:BIGINT;not null"`
	PANNumber   string `gorm:"column:pan;type:varchar(10);not null;unique"`
	Password    string `gorm:"column:password;type:text;not null"`
}

// TableName overrides the default table name
func (User) TableName() string {
	return "users"
}

