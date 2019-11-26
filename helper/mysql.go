package helper

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/gommon/log"
	"go-echo-example/pkg/setting"
	"sync"
)

var (
	err  error
	once sync.Once
	DB   *gorm.DB
)

func newDB() {
	db := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		setting.G_DBConf.User,
		setting.G_DBConf.Password,
		setting.G_DBConf.Host,
		setting.G_DBConf.Port,
		setting.G_DBConf.DataBase,
	)

	if DB, err = gorm.Open(setting.G_DBConf.Type, db); err != nil {
		log.Fatal("err:", err)
		return
	}

}
func GetDBInstance() *gorm.DB {
	once.Do(newDB)
	return DB
}
