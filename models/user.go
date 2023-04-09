package models

import (
	"go-jwt/helpers"

	"github.com/asaskevich/govalidator"

	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Full name is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Email is not valid"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,minstringlength(6)~Password must be at least 6 characters"`
	Products []Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"products"`
}


func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}