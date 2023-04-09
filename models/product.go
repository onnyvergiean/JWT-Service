package models

import (
	_ "github.com/asaskevich/govalidator"
)

type Product struct {
	GormModel
	Title 	 string `json:"title" form:"title" valid:"required~Title of your product is required"`
	Description string `json:"description" form:"description" valid:"required~Description of your product is required"`
	UserID 	 uint 
	User *User
}