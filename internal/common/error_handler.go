package common

import (
	"github.com/danhbuidcn/go_backend_api/global"
	"go.uber.org/zap"
)

// checkErrorPanic logs the error and panics if the error is not nil
func CheckErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}
