package initialize

import (
	"github.com/danhbuidcn/go_backend_api/global"
	"github.com/danhbuidcn/go_backend_api/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
