package client

import (
	redistimeseries "github.com/RedisTimeSeries/redistimeseries-go"
	"github.com/go-redis/redis/v7"
)

// RedisClient ... return instance of standard Redis client
func RedisClient() *redis.Client {
	// TODO - make configurable from CLI inputs
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// RedisTSClient ... return instance of time series Redis client
func RedisTSClient() *redistimeseries.Client {
	// TODO - make configurable from CLI inputs
	return redistimeseries.NewClient("localhost:6379", "nohelp", nil)
}
