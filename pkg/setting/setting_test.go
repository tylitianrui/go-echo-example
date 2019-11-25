package setting

import (
	"fmt"
	"testing"
)

func TestSetupConf(t *testing.T) {
	SetupConf("/Users/tyltr/GoSpace/go-echo-example/conf/app.ini")
	fmt.Println(G_DBConf)
	if G_DBConf.Password == "" {
		t.Log("password", G_DBConf.Password)
	} else {
		t.Fail()
	}

	if G_DBConf.Host == "127.0.0.1" {
		t.Log("Host", G_DBConf.Host)
	} else {
		t.Fail()
	}

	if G_DBConf.Type == "mysql" {
		t.Log("Host", G_DBConf.Type)
	} else {
		t.Fail()
	}

}
