package handler

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
) 


type RegisterUser struct {
	// Name is the first name of the user.
	Name string `json:"name,omitempty" binding:"required,min=2,max=50"`

	// Email is the email of the user.
	Email string `json:"email,omitempty" binding:"required,email"`

	// PhoneNumber is the phone number of the user.
	PhoneNumber string `json:"phone_number,omitempty" binding:"required,min=10"`

	// Password is the user's password.
	Password string `json:"password,omitempty" binding:"required,min=8"`
}



func (r RegisterUser) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required, validation.Length(2, 50)),
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.PhoneNumber, validation.Required, validation.Length(10, 20)), 
		validation.Field(&r.Password, validation.Required, validation.Length(8, 100)),
	)
}
