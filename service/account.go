package service

import (
	"errors"
	"go-echo-example/daos"
	"go-echo-example/models"
)

type AccountService struct {
	dao *daos.AccountDao
}

func (self *AccountService) Get(id uint) (data interface{}, err error) {
	return self.dao.Get(id)
}

func (*AccountService) Update(interface{}) error {
	panic("implement me")
}

func (self *AccountService) Create(data interface{}) error {

	if d, ok := data.(*models.Account); ok {
		return self.dao.Create(d)
	}
	return errors.New("param require type *models.Account")

}

func (*AccountService) Delete() {
	panic("implement me")
}

func NewAccountService() *AccountService {
	return &AccountService{
		dao: daos.NewAccountDao(),
	}
}
