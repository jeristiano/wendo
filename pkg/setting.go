package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

var (
	Config *ini.File
)

func Setup() {
	var err error
	Config, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatal("failed to parse 'config/app.ini':%v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := Config.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

//
////加载基本配置
//func LoadBase() {
//	RunMode = Config.Section("").Key("RUN_MODE").MustString("debug")
//
//}
//
////加载服务器配置
//func LoadServer() {
//	section, err := Config.GetSection("server")
//	if err != nil {
//		log.Fatal("failed to get section 'server' :%v", err)
//	}
//
//	HTTPPort = section.Key("HTTP_PORT").MustInt(8000)
//
//	ReadTimeout = time.Duration(section.Key("READ_TIMEOUT").MustInt(60)) * time.Second
//	WriteTimeOut = time.Duration(section.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
//}
//
//func LoadApp() {
//	section, err := Config.GetSection("app")
//	if err != nil {
//		log.Fatal("failed to get section 'app' :%v", err)
//	}
//
//	JwtSecret = section.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
//	PageSize = section.Key("PAGE_SIZE").MustInt(10)
//}

//获得app instance
var AppSetting = &App{}
