package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	redis.Client
}

func (rdb *Client) Get(ctx context.Context, key string, dest any) (err error) {
	var data string
	if data, err = rdb.Client.Get(ctx, key).Result(); err != nil {
		return err
	}

	if err = json.Unmarshal([]byte(data), dest); err != nil {
		return err
	}
	return nil
}

func (rdb *Client) Set(ctx context.Context, key string, data any, exp time.Duration) (err error) {
	var bytes []byte
	if bytes, err = json.Marshal(data); err != nil {
		return err
	}

	if err = rdb.Client.Set(ctx, key, bytes, exp).Err(); err != nil {
		return err
	}
	return nil
}
