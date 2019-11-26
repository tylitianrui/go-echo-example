package helper

import (
	"fmt"
	"go-echo-example/models"
	"go-echo-example/pkg/setting"

	//"log"
	"testing"
)

func TestGetDBInstance(t *testing.T) {
	setting.SetupConf("/Users/tyltr/GoSpace/go-echo-example/conf/app.ini")
	db := GetDBInstance()
	if db == nil {
		t.Fail()
	}
	if !db.HasTable("account") {
		t.Fail()
	}
	if !db.HasTable(&models.Account{}) {
		t.Fail()
	}
	db1 := GetDBInstance()
	fmt.Printf("%p\n", db1)
	fmt.Printf("%p\n", db)

}
