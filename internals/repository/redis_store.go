package repository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(addr string) *RedisStore {
	return &RedisStore{
		client: redis.NewClient(&redis.Options{
			Addr: addr,
		}),
	}
}

func (s *RedisStore) SetInitialStatus(ctx context.Context, txnId string) error {
	return s.client.Set(ctx, "txn:"+txnId, "PENDING", 10*time.Minute).Err()
}

func (s *RedisStore) UpdateFinalStatus(ctx context.Context, txnId string, status string) error {
	return s.client.Set(ctx, "txn:"+txnId, status, 10*time.Minute).Err()
}

func (s *RedisStore) GetStatus(ctx context.Context, txnId string) (string, error){
	val, err := s.client.Get(ctx, "txn:"+txnId).Result()
	if err == redis.Nil{
		return "NOT_FOUND", nil 
	}
	return val, err 
}