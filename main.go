package main

import (
	"balanceManager/balancemanager"
	"balanceManager/config"
	"balanceManager/delivery/http"
	"balanceManager/logger"
	"balanceManager/persistent/postgres"
	"go.uber.org/zap"
)

func main() {
	zapLogger, _ := zap.NewDevelopment()
	cfg := config.New()
	if err := cfg.Init(); err != nil {
	}

	postgresClient := postgres.WithConfig(cfg.DB)

	b := balancemanager.New(postgresClient)
	l := logger.New(b, *zapLogger)

	httpServ := http.WithConfig(cfg, l)

	httpServ.Run()
}
