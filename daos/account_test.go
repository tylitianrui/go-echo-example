package daos

import (
	"go-echo-example/helper"
	"go-echo-example/models"
	"go-echo-example/pkg/setting"
	"testing"
)

func TestAccountDao_Get(t *testing.T) {
	setting.SetupConf("/Users/tyltr/GoSpace/go-echo-example/conf/app.ini")
	db := helper.GetDBInstance()
	dao := NewAccountDao(db)
	//
	//a, _ := dao.Get(1)
	//fmt.Println(a.Mobile, a.Status)
	//
	//if a.User != "tyltr" {
	//	t.Fail()
	//}

	b := &models.Account{
		//ID:       111,
		User:     "qqq",
		Password: "qqqqqq",
		Mobile:   "113331",
		Status:   0,
		Created:  0,
	}
	dao.Create(b)

}
