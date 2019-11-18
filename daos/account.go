package daos

import (
	"crypto/md5"
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
	md5.New()
	var (
		account models.Account
	)
	self.db.First(&account, id)
	if account.ID == 0 {
		return &account, errors.New("not find")
	}
	return &account, nil
}
