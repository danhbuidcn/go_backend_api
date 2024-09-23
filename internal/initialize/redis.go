package initialize

import (
	"context"
	"fmt"

	"github.com/danhbuidcn/go_backend_api/global"
	"github.com/danhbuidcn/go_backend_api/internal/common"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// ctx is a context used to manage the lifecycle of Redis tasks
var ctx = context.Background()

// InitRedis initializes the connection to the Redis server, checks the connection using Ping, and stores the client in global.Rdb
func InitRedis() {
	r := global.Config.Redis // Retrieves Redis configuration from the config file
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port), // Redis address (Host:Port)
		Password: r.Password,                           // Redis password (if set)
		DB:       r.Database,                           // Redis database number (0 is default)
		PoolSize: r.PoolSize,                           // Maximum number of connections in the pool
	})

	// Ping the Redis server to verify the connection
	_, err := rdb.Ping(ctx).Result()
	common.CheckErrorPanic(err, "Failed to initialize Redis") // Log and panic if the connection fails

	fmt.Println("InitRedis is running") // Log message indicating Redis initialization is running
	global.Rdb = rdb                    // Store the Redis client in the global variable
	redisExample()                      // Call a sample function to demonstrate Redis usage
}

// redisExample is a sample function that sets and retrieves a key-value pair from Redis
func redisExample() {
	// Set the key "score" with the value 100 in Redis
	err := global.Rdb.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		fmt.Println("Error redis setting:", zap.Error(err)) // Log error if setting the value fails
	}

	// Retrieve the value of the key "score" from Redis
	value, err := global.Rdb.Get(ctx, "score").Result()
	if err != nil {
		fmt.Println("Error redis setting: ", zap.Error(err)) // Log error if retrieving the value fails
		return
	}
	// Log the retrieved value of "score" using the logger
	global.Logger.Info("value score is::", zap.String("score", value))
}
