package redis

import (
	"context"
	"fmt"
	redisDb "github.com/redis/go-redis/v9"
	"rating-bot/configs"
)

type RedisImpl struct {
	redisClient *redisDb.Client
	context     context.Context
}

func New(cfg *configs.Config) RedisApi {
	ctx := context.Background()
	client := getClient(
		cfg.RedisProps.RedisAddress,
		cfg.RedisProps.RedisPassword,
	)
	err := client.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}
	return &RedisImpl{
		redisClient: client,
		context:     ctx,
	}
}

func getClient(
	address string,
	password string,
) *redisDb.Client {
	return redisDb.NewClient(&redisDb.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})
}

func (adapter *RedisImpl) SetData(key string, value string) error {
	err := adapter.redisClient.Set(adapter.context, key, value, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (adapter *RedisImpl) GetData(key string) string {
	result, err := adapter.redisClient.Get(adapter.context, key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return result
}

func (adapter *RedisImpl) DeleteData(key string) error {
	err := adapter.redisClient.Del(adapter.context, key).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (adapter *RedisImpl) SAdd(key string, member string) error {
	err := adapter.redisClient.SAdd(adapter.context, key, member).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (adapter *RedisImpl) SMembers(key string) []string {
	result, err := adapter.redisClient.SMembers(adapter.context, key).Result()
	if err != nil {
		fmt.Println(err)
		return []string{}
	}
	return result
}

func (adapter *RedisImpl) SISMembers(key string, value string) bool {
	result, err := adapter.redisClient.SIsMember(adapter.context, key, value).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return result
}

func (adapter *RedisImpl) SRem(key string, member string) error {
	_, err := adapter.redisClient.SRem(adapter.context, key, member).Result()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func (adapter *RedisImpl) LPush(key string, value string) error {
	err := adapter.redisClient.LPush(adapter.context, key, value).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (adapter *RedisImpl) RPush(key string, value string) error {
	err := adapter.redisClient.RPush(adapter.context, key, value).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (adapter *RedisImpl) LPop(key string) string {
	result, err := adapter.redisClient.LPop(adapter.context, key).Result()
	if err != nil {
		return ""
	}
	return result
}

func (adapter *RedisImpl) GetSize(key string) int64 {
	result, err := adapter.redisClient.LLen(adapter.context, key).Result()
	if err != nil {
		return 0
	}
	return result
}

func (adapter *RedisImpl) LRange(key string) []string {
	result, err := adapter.redisClient.LRange(adapter.context, key, 0, 100).Result()
	if err != nil {
		fmt.Println(err)
		return []string{}
	}

	return result
}

func (adapter *RedisImpl) LTrim(key string, startIndex int64) error {
	_, err := adapter.redisClient.LTrim(adapter.context, key, startIndex, -1).Result()
	if err != nil {
		fmt.Println(err)
	}
	return err
}
