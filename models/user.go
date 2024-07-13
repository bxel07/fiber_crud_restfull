package models

import (
	"go_fiber_restfull/validator"
	"time"

	"gorm.io/gorm"
)

type User struct{
	gorm.Model
	Username string `gorm:"uniqueIndex;not null" json:"username"`
	Email string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	ResetToken        *string   `json:"-"`
	ResetTokenExpiry  *time.Time `json:"-"`
}

func (p *User) ValidateRegisterUsers() error {
	v := validator.NewValidator()
	v.AddRule("username", validator.FieldRule{Required: true, Max: 255})
	v.AddRule("email", validator.FieldRule{Required: true})
	v.AddRule("password", validator.FieldRule{Required: true})

	return v.Validate(p)
}

func (p *User) ValidateLoginUsers() error {
	v := validator.NewValidator()
	v.AddRule("email", validator.FieldRule{Required: true})
	v.AddRule("password", validator.FieldRule{Required: true})

	return v.Validate(p)
}