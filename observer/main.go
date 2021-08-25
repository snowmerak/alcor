package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "0.0.0.0:6379",
		Password: "",
		DB:       0,
	})

	if err := rdb.Set(context.Background(), "test", uint64(123), 0).Err(); err != nil {
		log.Fatal(err)
	}

	v, err := rdb.Get(context.Background(), "test").Uint64()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(v)
}
