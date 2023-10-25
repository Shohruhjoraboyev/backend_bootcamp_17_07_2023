package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gitlab.com/market3723841/api-gateway-service/api"
	"gitlab.com/market3723841/api-gateway-service/api/handler"
	"gitlab.com/market3723841/api-gateway-service/config"
	"gitlab.com/market3723841/api-gateway-service/pkg/logger"
	"gitlab.com/market3723841/api-gateway-service/services"
)

func main() {
	cfg := config.Load()

	fmt.Printf("config: %+v/n", cfg)

	// Setup Logger
	loggerLevel := logger.LevelDebug
	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelDebug
	default:
		loggerLevel = logger.LevelInfo
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer logger.Cleanup(log)

	grpcSrvc, err := services.NewGrpcClients(cfg)
	if err != nil {
		panic(err)
	}

	r := gin.New()

	h := handler.NewHandler(cfg, log, grpcSrvc)

	api.SetUpApi(r, h, cfg)

	fmt.Println("Start api gateway...")

	err = r.Run(cfg.HTTPPort)
	if err != nil {
		return
	}
}
