package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// 1.
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name:%s, age:%d", "TipsGo", 40) // like fmt.Printf

	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "TipsGo"), zap.Int("age", 40))

	// // 2.
	// // {"level":"info","msg":"Hello NewExample"}
	// logger := zap.NewExample()
	// logger.Info("Hello NewExample")

	// // development
	// // 2024-09-18T16:23:36.166+0700    INFO    cli/main.log.go:17      Hello NewDevelopment
	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello NewDevelopment")

	// // production
	// // {"level":"info","ts":1726651416.1669984,"caller":"cli/main.log.go:21","msg":"Hello NewProduction"}
	// logger, _ = zap.NewProduction()
	// logger.Info("Hello NewProduction")

	// 3.
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info log", zap.Int("line", 1))
	logger.Error("Error log", zap.Int("line", 2))
}

// Customize log information
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// 1726651416.1669984 -> 2024-09-18T16:23:36.166+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// ts -> time
	encodeConfig.TimeKey = "time"
	// from info INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// "caller":"cli/main.log.go:21"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

// log to file and console
func getWriterSync() zapcore.WriteSyncer {
	// create folder if not exist
	os.MkdirAll("./log", os.ModePerm)

	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm) // file name, open file mode, permissions
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)
}
