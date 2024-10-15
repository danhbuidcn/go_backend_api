package global

import (
	"database/sql"

	"github.com/danhbuidcn/go_backend_api/pkg/logger"
	"github.com/danhbuidcn/go_backend_api/pkg/setting"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

var (
	Config        setting.Config
	Logger        *logger.LoggerZap
	Rdb           *redis.Client
	Mdb           *gorm.DB
	Mdbc          *sql.DB
	KafkaProducer *kafka.Writer
)
