package model

import {
	"Time"
	"errors"
	"github.com/asaskevich/govalidator"
}

type IPixKeyRepository interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAccount(account *Account) error
	FindAccount(id string) (*Account, error)
}

type PixKey struct {
	Base `valid:"required"`
	Kind string `json:"kind" valid:"notnull"`
	Key string `json:"key" valid:"notnull"`
	AccountId string `json:"account_id" valid:"notnull"`
	Account *Account `valid:"-"`
	Status string `json:"status" valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {
	_, err:= govalidator.ValidateStruct(pixKey)

	if pixKey.Kind != "email" && pixKey.Kind != "cpf" {
		return errors.New("Invalid type of key")
	}
	
	if pixKey.Status != "active" && pixKey.Status != "inactive" {
		return errors.New("Invalid status")
	}

	if err != nil {
		return err
	}
}