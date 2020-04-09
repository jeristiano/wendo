module github.com/jeristiano/wendo

go 1.14

require (
	github.com/astaxie/beego v1.12.1
	github.com/denisenkom/go-mssqldb v0.0.0-20200206145737-bbfc9a55622e // indirect
	github.com/gin-gonic/gin v1.6.2
	github.com/go-ini/ini v1.55.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/jinzhu/now v1.1.1 // indirect
	github.com/lib/pq v1.3.0 // indirect
	github.com/mattn/go-sqlite3 v2.0.3+incompatible // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/tealeg/xlsx v1.0.5
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20200406173513-056763e48d71 // indirect
	golang.org/x/sys v0.0.0-20200408040146-ea54a3c99b9b // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/ini.v1 v1.55.0 // indirect
)

replace (
	github.com/jerisitano/wendo/config => ./wendo/pkg/config
	github.com/jerisitano/wendo/kg/error => ./wendo/pkg/error
	//github.com/jerisitano/wendo/middod leware => ./wendo/middleware
	github.com/jerisitano/wendo/model => ./wendo/model
	github.com/jerisitano/wendo/pkg/export => ./wendo/pkg/export
	github.com/jerisitano/wendo/pkg/logging => ./wendo/pkg/logging
	github.com/jerisitano/wendo/pkg/setting => ./wendo/pkg/setting
	github.com/jerisitano/wendo/routers => ./wendo/routers
	github.com/jerisitano/wendo/runtime => ./wendo/runtime
	github.com/jerisitano/wendo/service => ./wendo/service
)
