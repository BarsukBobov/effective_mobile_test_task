package app

import (
	"context"
	"effective_mobile_test_task/internal/handler"
	"effective_mobile_test_task/internal/models"
	"effective_mobile_test_task/internal/repository/em"
	"effective_mobile_test_task/internal/repository/sql"
	"effective_mobile_test_task/internal/service"
	"effective_mobile_test_task/pkg/misc"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var logger = misc.GetLogger()

type App struct {
	envConf *models.EnvConfig
	dbPool  *sql.DbPool
	srv     *misc.Server
}

func NewApp(envConf *models.EnvConfig) *App {
	return &App{envConf: envConf}
}

func (a *App) Init() {
	appConf, err := models.NewAppConfig(a.envConf.AppConfigPath)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	err = a.start(appConf)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Info("EFFECTIVE MOBILE TEST TASK API started successfully")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	a.shutdown()

	logger.Info("EFFECTIVE MOBILE TEST TASK API shutting down successfully")
}

func (a *App) start(appConf *models.AppConfig) error {
	var err error
	var errMessage string

	a.dbPool, err = sql.NewDbPool(
		appConf.PostgreSQL.Dsn,
		a.envConf.MinConns,
		a.envConf.MaxConns,
		a.envConf.MaxConnLifetime,
		a.envConf.MaxConnIdleTime,
		a.envConf.HealthCheckPeriod,
	)
	if err != nil {
		errMessage = fmt.Sprintf("failed to initialize postgreSQL: %s", err.Error())
		return errors.New(errMessage)
	}

	logger.Info("postgreSQL connection pool opened successfully")

	sqlObj := sql.NewSQL(a.dbPool)

	emHttp, err := em.NewEmHttp(a.envConf.EmApi)
	if err != nil {
		errMessage = fmt.Sprintf("failed to migrate schema to postgreSQL: %s", err.Error())
		return errors.New(errMessage)
	}

	logger.Info("EmHttp created successfully")

	newService := service.NewService(sqlObj, emHttp)

	err = misc.DbMigrate(appConf.PostgreSQL.Dsn, appConf.DbmateMigrationsDir)
	if err != nil {
		errMessage = fmt.Sprintf("failed to migrate schema to postgreSQL: %s", err.Error())
		return errors.New(errMessage)
	}

	logger.Info("Migration was successfully")

	a.srv = new(misc.Server)

	go func() {
		if err = a.srv.Run(handler.InitRoutes(newService, appConf, a.envConf.Production)); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				logger.Errorf("error occured while running http server: %s", err.Error())
			}
		}
	}()
	return nil
}

// shutdown function for graceful close. The function will not run on develop mode
func (a *App) shutdown() {
	if err := a.srv.Shutdown(context.Background()); err != nil {
		logger.Errorf("error occured on server shutting down: %s", err.Error())
	} else {
		logger.Info("HTTP server closed")
	}

	if err := a.dbPool.Close(); err != nil {
		logger.Errorf("error occured on redis connection close: %s", err.Error())
	} else {
		logger.Info("postgreSQL connection pool closed successfully")
	}

}
