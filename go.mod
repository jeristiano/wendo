module github.com/jeristiano/wendo

go 1.14

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1 // indirect
	github.com/gin-gonic/gin v1.6.2
	github.com/go-ini/ini v1.55.0
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/tealeg/xlsx v1.0.5 // indirect
	golang.org/x/sys v0.0.0-20200331124033-c3d80250170d // indirect
)

replace (
	github.com/jerisitano/wendo/config => ./wendo/pkg/config
	github.com/jerisitano/wendo/kg/error => ./wendo/pkg/error
	github.com/jerisitano/wendo/pkg/export => ./wendo/pkg/export
	github.com/jerisitano/wendo/pkg/logging => ./wendo/pkg/logging
	//github.com/jerisitano/wendo/middleware => ./wendo/middleware
	//github.com/jerisitano/wendo/models => ./wendo/models
	github.com/jerisitano/wendo/pkg/setting => ./wendo/pkg/setting
	github.com/jerisitano/wendo/routers => ./wendo/routers
	github.com/jerisitano/wendo/runtime => ./wendo/runtime
	github.com/jerisitano/wendo/service => ./wendo/service
)
