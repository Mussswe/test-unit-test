package backend

import "github.com/asaskevich/govalidator"

type User struct {
	Name  string `valid:"required,alpha"`
	Email string `valid:"required,email"`
}

func ValidateUser(user User) (bool, error) {
	return govalidator.ValidateStruct(user)
}
