package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jeristiano/wendo/pkg"
	"github.com/jeristiano/wendo/routers"
	"net/http"
)

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	router := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	s := &http.Server{
		Addr:           endPoint,
		Handler:        router,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	_ = s.ListenAndServe()
}
