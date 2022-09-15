package app

import (
	"dovran/mascot/config"
	v1 "dovran/mascot/internal/controller/v1"
	"dovran/mascot/internal/usecase"
	"dovran/mascot/internal/usecase/balance"
	"dovran/mascot/internal/usecase/transaction"
	"dovran/mascot/pkg/httpserver"
	"dovran/mascot/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {

	var err error

	l := logger.New(cfg.Level)

	uc := usecase.NewUCase(balance.NewStorageBalance(), transaction.NewStorageTransaction())

	handler := gin.New()

	httpServer := httpserver.New(handler, httpserver.Port(cfg.HttpPort))

	v1.NewRouter(handler, uc, l)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
