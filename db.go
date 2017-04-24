package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var client *redis.Client

func makeRedisConn() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if pong, err := client.Ping().Result(); err != nil {
		panic(err)
	} else {
		fmt.Printf("pong is %s\n", pong)
	}
	client.Del("mytest")
	if str, err := client.Set("mytest", 15, 0).Result(); err != nil {
		panic(err)
	} else {
		fmt.Printf("result for setting mytest is: %s\n", str)
	}

}
