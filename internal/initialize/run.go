package initialize

import (
	"fmt"

	"github.com/danhbuidcn/go_backend_api/global"
	"go.uber.org/zap"
)

func Run() {
	// load configuration
	LoadConfig()
	fmt.Println("Loading configuration mysql", global.Config.Mysql.Username)
	InitLogger()
	global.Logger.Debug("config log ok", zap.String("ok", "success"))
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(":8002")
}
