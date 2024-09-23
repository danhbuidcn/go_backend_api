package global

import (
	"github.com/danhbuidcn/go_backend_api/pkg/logger"
	"github.com/danhbuidcn/go_backend_api/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
