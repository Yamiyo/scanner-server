package main

import (
	"context"
	"fmt"
	"runtime/debug"

	"portto-homework/internal/utils/logger"
	"portto-homework/service/api-service/app"
)

func main() {
	defer func() {
		logger.SysLog().Info(context.Background(), "Server shutdown...")
		if err := recover(); err != nil {
			logger.SysLog().Error(context.Background(), fmt.Sprintf("error: %v", err))
			logger.SysLog().Error(context.Background(), string(debug.Stack()))
		}
	}()

	app.Run()
}
