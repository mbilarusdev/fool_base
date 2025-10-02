package cache

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/mbilarusdev/fool_base/v2/log"
)

func PingRDB() *redis.Client {
	segments := strings.Split(os.Getenv("RDB"), "&")
	addr := segments[0]
	password := segments[1]

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	log.Info("Connected to redis")

	ctx := context.Background()

	pingCmd := rdb.Ping(ctx)

	log.Info(fmt.Sprintf("From Redis: %v", pingCmd.Val()))

	return rdb
}
