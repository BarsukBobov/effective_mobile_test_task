package main

import (
	"effective_mobile_test_task/internal/app"
	"effective_mobile_test_task/internal/models"
	"effective_mobile_test_task/pkg/misc"
)

// @title EFFECTIVE MOBILE TEST TASK API
// @version 1.0
// @description API Server for effective mobile's test task
// @host localhost:3000
// @BasePath /api/v1
func main() {
	envConf, err := models.NewEnvConfig()
	if err != nil {
		panic(err.Error())
		return
	}

	logger := misc.GetLogger()
	logger.SetParams(envConf.LogLevel)

	appObj := app.NewApp(envConf)
	appObj.Init()
}
