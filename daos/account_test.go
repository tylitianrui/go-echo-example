package daos

import (
	"fmt"
	"go-echo-example/helper"
	"go-echo-example/runtime/setting"
	"testing"
)

func TestAccountDao_Get(t *testing.T) {
	setting.SetupConf("/Users/tyltr/GoSpace/go-echo-example/conf/app.ini")
	db := helper.GetDBInstance()
	dao := NewAccountDao(db)

	a, _ := dao.Get(1)
	fmt.Println(a.Mobile, a.Status)

	if a.User != "tyltr" {
		t.Fail()
	}

}
