package model

import {
	"Time"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
}

type User struct {
	Base `valid:"required"`
	name string `json:"name" valid:"notnull"`
	email string `json:"email" valid:"notnull"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}
	return nil
}

func newUser(name string, email string) (*User, error){
	user := new User{
		Name: name,
		Email: email,
	}

	user.ID = uuid.newV4().String()
	user.CreatedAt = time.Now()

	err := user.isValid()
	if err != nil {
		return nil, err
	}

	return &User, nil
}