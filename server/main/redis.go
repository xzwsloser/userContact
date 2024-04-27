package main

import (
	"github.com/redis/go-redis/v9"
)

var client *redis.Client

func initClient() {
	client = redis.NewClient(&redis.Options{
		Addr:     "192.168.59.132:6379",
		Password: "808453",
		DB:       0,
	})
}
