package setting

import (
	"fmt"
	"gopkg.in/ini.v1"
)

type MysqlConf struct {
	Type     string
	Host     string
	Port     int
	User     string
	Password string
	DataBase string
}
type AppConf struct {
	JwtSecret       string
	MobileVerifyNum int
}

var (
	G_AppConf = &AppConf{}
	// 全局
	G_DBConf = &MysqlConf{}
)

func SetupConf(conf string) {
	fmt.Println(fmt.Sprintf("conf file: %s ,loading...", conf))
	parseIni(conf)

}

func parseIni(conf string) {
	var (
		err error
		cfg *ini.File
	)

	if cfg, err = ini.Load(conf); err != nil {
		return
	}
	iniMapTo(cfg, "database", G_DBConf)
	iniMapTo(cfg, "app", G_AppConf)

}

func iniMapTo(cfg *ini.File, section string, x interface{}) {
	var (
		err error
	)
	if err = cfg.Section(section).MapTo(x); err != nil {
		fmt.Println(err)
		return
	}

}
