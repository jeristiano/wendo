package main

import (
	"fmt"
	"github.com/jeristiano/wendo/pkg"
	"github.com/jeristiano/wendo/routers"
	"net/http"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	_ = s.ListenAndServe()
}
