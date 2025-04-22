package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type DbAccess struct {
	rdb *redis.Client
}

func NewDbAccess(dsn string) DbAccess {
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		log.Fatal(err)
	}

	rdb := redis.NewClient(opt)
	return DbAccess{rdb: rdb}
}

func (db *DbAccess) AddTask(list string, data []byte) {
	db.rdb.RPush(context.Background(), list, data)
}

func (db *DbAccess) GetTask(list string) ([]byte, error) {
	result, err := db.rdb.BLPop(context.Background(), 1*time.Second, list).Result()
	return []byte(result[1]), err
}

func (db *DbAccess) AddResult(key string, data []byte) {
	db.rdb.Set(context.Background(), key, data, 0)
}
func (db *DbAccess) GetResult(key string) ([]byte, error) {
	result, err := db.rdb.GetDel(context.Background(), key).Bytes()
	return result, err
}
