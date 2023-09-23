package boot

import (
	"os"

	"github.com/test_cache_CQRS/config"
	"github.com/test_cache_CQRS/infrastructure/limiter"
	logger "github.com/test_cache_CQRS/infrastructure/log"
	"github.com/test_cache_CQRS/infrastructure/postgres"
	"github.com/test_cache_CQRS/infrastructure/redis"
	"github.com/test_cache_CQRS/module/article"
	"github.com/test_cache_CQRS/module/health"
	"github.com/test_cache_CQRS/utils"

	redisThirdPartyClient "github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

type HandlerSetup struct {
	Limiter     *limiter.RateLimiter
	HealthHttp  health.InterfaceHttp
	ArticleHttp article.InterfaceHttp
}

func MakeHandler() HandlerSetup {
	//initiate config
	config.Initialize()

	//initiate logger
	if config.Conf.LogMode {
		logger.Init(config.Conf.LogFormat, config.Conf.LogLevel)
	}

	var err error

	var redisClient *redisThirdPartyClient.Client
	var redisLibInterface redis.LibInterface
	if config.Conf.Redis.EnableRedis {
		//initiate a redis client
		redisClient, err = redis.NewRedisClient(&config.Conf)
		if err != nil {
			log.Fatalf("failed initiate redis: %v", err)
			os.Exit(1)
		}
		//initiate a redis library interface
		redisLibInterface, err = redis.NewRedisLibInterface(redisClient)
		if err != nil {
			log.Fatalf("failed initiate redis library: %v", err)
			os.Exit(1)
		}
	}

	//setup infrastructure postgres
	db, err := postgres.NewDatabaseClient(&config.Conf)
	if err != nil {
		log.Fatalf("failed initiate database postgres: %v", err)
		os.Exit(1)
	}

	//add limiter
	interval := utils.StringUnitToDuration(config.Conf.Interval)
	middlewareWithLimiter := limiter.NewRateLimiter(int(config.Conf.Rate), interval)

	//health module
	healthRepository := health.NewRepository(db.DbConn)
	healthService := health.NewService(healthRepository, redisClient)
	healthModule := health.NewHttp(healthService)

	//article module
	articleRepository := article.NewRepository(db.DbConn)
	articleService := article.NewService(articleRepository, redisLibInterface)
	articleModule := article.NewHttp(articleService)

	return HandlerSetup{
		Limiter:     middlewareWithLimiter,
		HealthHttp:  healthModule,
		ArticleHttp: articleModule,
	}
}
