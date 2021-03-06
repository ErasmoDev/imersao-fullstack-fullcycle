package model

import {
	"Time"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
}

type Account struct {
	Base `valid:"required"`
	OwnerName string `json:"owner_name" valid:"notnull"`
	Bank *Bank `valid:"-"`
	Number string `json:"number" valid:"notnull"`
	PixKeys []*PixKey `valid:"-"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}
	return nil
}

func newAccount(bank *Bank, ownerName string, number string) (*Account, error){
	account := new Account{
		OwnerName: owneName,
		Number: number,
	}

	account.ID = uuid.newV4().String()
	account.CreatedAt = time.Now()

	err := account.isValid()
	if err != nil {
		return nil, err
	}

	return &Account, nil
}