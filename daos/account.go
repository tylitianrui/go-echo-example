package daos

import (
	"errors"
	"github.com/jinzhu/gorm"
	"go-echo-example/models"
)

type AccountDao struct {
	db *gorm.DB
}

func NewAccountDao(db *gorm.DB) *AccountDao {
	return &AccountDao{
		db: db,
	}
}

func (self *AccountDao) Get(id uint) (*models.Account, error) {
	var (
		account models.Account
	)
	self.db.First(&account, id)
	if account.ID == 0 {
		return &account, errors.New("not find")
	}
	return &account, nil
}

func (self *AccountDao) GetByMobile(mobile uint) (acount *models.Account, err error) {
	var (
		account models.Account
	)

	if err = self.db.Where("mobile = ?", mobile).First(&account).Error; err != nil {

		return nil, err
	}

	if account.ID == 0 {
		return &account, errors.New("not find")
	}
	return &account, nil
}

func (self *AccountDao) Create(a *models.Account) (err error) {

	if err = self.db.Create(a).Error; err != nil {
		return err
	}
	return nil
}
