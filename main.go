package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo"
	"go-echo-example/helper"
	"go-echo-example/pkg/setting"
	"go-echo-example/routers"
	"net/http"
	"os"
	"time"
)

var (
	//ctx context.Context
	err         error
	rootBase    string
	defaultConf string
)

func init() {
	rootBase, _ = os.Getwd()
	defaultConf = fmt.Sprintf("%s/conf/app.ini", rootBase)
	helper.SetupRedis()
}

func main() {
	// 传入配置文件 默认配置文件是当前目录下的/conf/app.ini
	confDir := flag.String("c", defaultConf, "")
	flag.Parse()
	setting.SetupConf(*confDir)
	e := echo.New()
	// 自定义
	s := http.Server{
		Addr:              ":9090",
		ReadTimeout:       600 * time.Millisecond,
		ReadHeaderTimeout: 60 * time.Millisecond,
		WriteTimeout:      600 * time.Millisecond,
		IdleTimeout:       600 * time.Millisecond,
		MaxHeaderBytes:    1 << 20,
	}

	// 初始化路由
	routers.InitRouters(e)

	e.Logger.Fatal(e.StartServer(&s))

}
