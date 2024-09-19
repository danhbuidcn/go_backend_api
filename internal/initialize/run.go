package initialize

import (
	"fmt"

	"github.com/danhbuidcn/go_backend_api/global"
)

func Run() {
	// load configuration
	LoadConfig()
	fmt.Println("Loading configuration mysql", global.Config.Mysql.Username)
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(":8002")
}
