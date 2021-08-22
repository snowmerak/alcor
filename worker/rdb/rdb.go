package rdb

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var store *redis.Client

func init() {
	store = redis.NewClient(&redis.Options{
		Addr:     "rdb:6379",
		Password: "",
		DB:       0,
	})
	if store == nil {
		log.Fatal("cannot connect rdb")
	}
}

func GetLastBundle() ([]byte, error) {
	return store.Get(context.Background(), "last_bundle").Bytes()
}

func SetLastBundle(hash []byte) error {
	return store.Set(context.Background(), "last_bundle", hash, 0).Err()
}

func GetSurveyIndex() (uint64, error) {
	return store.Get(context.Background(), "survey_index").Uint64()
}

func SetSurveyIndex(index uint64) error {
	return store.Set(context.Background(), "survey_index", index, 0).Err()
}
