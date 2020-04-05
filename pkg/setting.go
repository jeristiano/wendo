package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Config *ini.File

	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeOut time.Duration

	PageSize  int
	JwtSecret string
)

func init() {
	var err error
	Config, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatal("failed to parse 'config/app.ini':%v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()

}

//加载基本配置
func LoadBase() {
	RunMode = Config.Section("").Key("RUN_MODE").MustString("debug")

}

//加载服务器配置
func LoadServer() {
	section, err := Config.GetSection("server")
	if err != nil {
		log.Fatal("failed to get section 'server' :%v", err)
	}

	HTTPPort = section.Key("HTTP_PORT").MustInt(8000)

	ReadTimeout = time.Duration(section.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WriteTimeOut = time.Duration(section.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

func LoadApp() {
	section, err := Config.GetSection("app")
	if err != nil {
		log.Fatal("failed to get section 'app' :%v", err)
	}

	JwtSecret = section.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = section.Key("PAGE_SIZE").MustInt(10)
}
