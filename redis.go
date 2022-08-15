package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var (
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", cfg.InMemoryDb.Hostname, cfg.InMemoryDb.Port),
	})
	ctx = context.Background()
)

func getAll(topic string) map[string]any {
	output := make(map[string]any)
	raw := must(rdb.HGetAll(ctx, topic).Result())

	for k, v := range raw {
		output[k] = must(strconv.Atoi(v))
	}

	return output
}

func incrementCount(topic, eventType string) {
	count := 0
	val, _ := rdb.HGet(ctx, topic, eventType).Result()
	if val != "" {
		count = must(strconv.Atoi(val))
	}

	must(rdb.HSet(ctx, topic, eventType, count+1).Result())
}
