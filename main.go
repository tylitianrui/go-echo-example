package main

import (
	"flag"
	"fmt"
	"github.com/labstack/echo"
	"go-echo-example/routers"
	"go-echo-example/runtime/setting"
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
}

func main() {
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
