package app

import (
	"context"
	"errors"
	"github.com/stickpro/vpn-sass/internal/config"
	"github.com/stickpro/vpn-sass/internal/repository"
	"github.com/stickpro/vpn-sass/internal/router"
	"github.com/stickpro/vpn-sass/internal/server"
	"github.com/stickpro/vpn-sass/internal/service"
	"github.com/stickpro/vpn-sass/pkg/database/pgsql"
	"github.com/stickpro/vpn-sass/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	cfg, err := config.Init()

	if err != nil {
		logger.Error(err)
		return
	}

	db, _ := pgsql.ConnectionDataBase(cfg.DB.Host, cfg.DB.Username, cfg.DB.Password, cfg.DB.DBName, cfg.DB.Port)

	repos := repository.NewRepositories(db)
	repos.Migrate()

	services := service.NewServices(service.Deps{
		Repository: repos,
	})

	newRouter := router.NewRouter(services)

	srv := server.NewServer(cfg.HTTP, newRouter.Init())

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("error occurred while running http server: %s\n", err.Error())
		}
	}()

	logger.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		logger.Errorf("failed to stop server: %v", err)
	}
}
