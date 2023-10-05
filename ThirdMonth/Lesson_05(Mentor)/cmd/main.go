package main

import (
	"context"
	"fmt"
	"projectWithGinAndSwagger/api"
	"projectWithGinAndSwagger/api/handler"
	"projectWithGinAndSwagger/config"
	"projectWithGinAndSwagger/pkg/logger"
	"projectWithGinAndSwagger/storage/postgres"
	"projectWithGinAndSwagger/storage/redis"
)

func main() {
	cfg := config.Load()
	log := logger.NewLogger("mini-project", logger.LevelInfo)
	strg, err := postgres.NewStorage(context.Background(), cfg)
	if err != nil {
		return
	}
	redisStrg, err := redis.NewCache(context.Background(), cfg)
	if err != nil {
		return
	}

	h := handler.NewHandler(strg, redisStrg, log)

	r := api.NewServer(h)
	r.Run(fmt.Sprintf(":%s", cfg.Port))

}
