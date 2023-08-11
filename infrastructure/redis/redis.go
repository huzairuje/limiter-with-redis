package redis

import (
	"fmt"
	"time"

	"github.com/test_cache_CQRS/config"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

// NewRedisClient initialize a redis client
func NewRedisClient(conf *config.Config) (redisLib LibInterface, err error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Conf.Redis.Host, config.Conf.Redis.Port),
		Password: conf.Redis.Password,
		DB:       conf.Redis.DB,
	})

	_, err = redisClient.Ping().Result()
	if err != nil {
		log.Fatalf("Open connection to redis, error: %v", err)
		return nil, err
	}
	redisLib = newLib(redisClient)
	log.Printf("Connected to redis on %s (DB: %d)", config.Conf.Redis.Host, config.Conf.Redis.DB)
	return redisLib, nil
}

type client struct {
	redisClient *redis.Client
}

type LibInterface interface {
	SetIdempotencyKey(key string, value interface{}, ttl time.Duration) (err error)
	DeleteKey(key string) (err error)
	Get(key string) (value string)
	Set(key string, value interface{}, ttl time.Duration) (err error)
}

func newLib(redisClient *redis.Client) LibInterface {
	return client{
		redisClient: redisClient,
	}
}

func (r client) SetIdempotencyKey(key string, value interface{}, ttl time.Duration) (err error) {
	valueInRedis := r.redisClient.Get(key).Val()
	if len(valueInRedis) > 0 {
		//err = primitive.ErrFoundInCache
		return
	}
	success, err := r.redisClient.SetNX(key, value, ttl).Result()
	if err != nil {
		return
	}
	if !success {
		//err = primitive.ErrMultipleKeyInCache
	}
	return
}

func (r client) DeleteKey(key string) (err error) {
	val := r.redisClient.Get(key).Val()
	if len(val) > 0 {
		return r.redisClient.Del(key).Err()
	}
	return
}

func (r client) Get(key string) string {
	return r.redisClient.Get(key).Val()
}

func (r client) Set(key string, value interface{}, ttl time.Duration) error {
	return r.redisClient.Set(key, value, ttl).Err()
}
