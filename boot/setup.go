package boot

import (
	"os"

	"github.com/test_cache_CQRS/config"
	logger "github.com/test_cache_CQRS/infrastructure/log"
	"github.com/test_cache_CQRS/infrastructure/postgres"
	"github.com/test_cache_CQRS/infrastructure/redis"
	"github.com/test_cache_CQRS/module/article"
	"github.com/test_cache_CQRS/module/health"

	log "github.com/sirupsen/logrus"
)

type HandlerSetup struct {
	HealthHttp  health.InterfaceHttp
	ArticleHttp article.InterfaceHttp
}

func MakeHandler() HandlerSetup {
	config.Initialize()
	logger.Init(config.Conf.LogFormat, config.Conf.LogLevel)

	var err error

	redisClient, err := redis.NewRedisClient(&config.Conf)
	if err != nil {
		log.Fatalf("failed initiate redis: %v", err)
		os.Exit(1)
	}

	//setup infrastructure postgres
	db, err := postgres.NewDatabaseClient(&config.Conf)
	if err != nil {
		log.Fatalf("failed initiate database postgres: %v", err)
		os.Exit(1)
	}

	//health module
	healthRepository := health.NewRepository(db.DbConn)
	healthService := health.NewService(healthRepository)
	healthModule := health.NewHttp(healthService)

	//article module
	articleRepository := article.NewRepository(db.DbConn)
	articleService := article.NewService(articleRepository, redisClient)
	articleModule := article.NewHttp(articleService)

	return HandlerSetup{
		HealthHttp:  healthModule,
		ArticleHttp: articleModule,
	}
}
