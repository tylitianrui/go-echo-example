package main

import (
	"flag"
	"fmt"
	"go-echo-example/runtime/setting"
	"os"
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

}
