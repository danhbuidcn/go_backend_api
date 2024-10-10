package initialize

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/danhbuidcn/go_backend_api/global"
	"github.com/danhbuidcn/go_backend_api/internal/common"
	"github.com/danhbuidcn/go_backend_api/internal/model"
	"github.com/danhbuidcn/go_backend_api/internal/po"
	"go.uber.org/zap"
	"gorm.io/gen"
)

// InitMysql initializes the MySQL connection
func InitMysqlc() {
	m := global.Config.Mysql
	// Build the Data Source Name (DSN)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username, m.Password, m.Host, m.Port, m.Dbname)

	// Open the connection
	db, err := sql.Open("mysql", dsn)
	common.CheckErrorPanic(err, "Failed to initialize MySQL")

	global.Logger.Info("MySQL Initialized Successfully")
	global.Mdbc = db

	// Set connection pool settings
	// A pool is a set of pre-maintained connections that improve performance.
	setPoolc()

	genTableDAOc()

	// Run migrations
	migrateTablesc()
}

// setPool sets the MySQL connection pool settings
func setPoolc() {
	m := global.Config.Mysql
	sqlDb, err := global.Mdb.DB()
	common.CheckErrorPanic(err, "Failed to get SQL DB from GORM")

	// Set connection pool configurations
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns) * time.Second)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Second)
}

// mysql to model
func genTableDAOc() {
	// Initiate the tables
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/models",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.Mdb) // reuse your gorm db
	g.GenerateAllTable()

	// // Generate basic type-safe DAO API for struct `model.User` following conventions
	// g.ApplyBasic(model.User{})

	// // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	// g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()
}

// migrateTables runs database migrations
func migrateTablesc() {
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
		&model.GoCrmUserV2{}, // model to mysql
	)
	if err != nil {
		global.Logger.Error("Error during table migration", zap.Error(err))
	}
}
