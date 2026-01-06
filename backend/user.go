package backend

import "github.com/asaskevich/govalidator"

type Address struct {
	City     string `valid:"required,alpha"`
	PostCode string `valid:"required,numeric"`
}
type Role struct {
	RoleName string `valid:"required,alpha"` // เช่น admin, editor, viewer
}
type User struct {
	Sut_id  string    `valid:"matches(^[BCM]\\d{7}$)"`
	Name    string    `valid:"required,alpha,matches(^[A-Z])"`
	Email   string    `valid:"required,email"`
	Address []Address `valid:"required,valid"`
	Roles   []Role    `valid:"required,valid"`
}

func ValidateUser(user User) (bool, error) {
	return govalidator.ValidateStruct(user)
}
