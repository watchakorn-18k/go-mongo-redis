package repositories

import (
	"context"
	"encoding/json"
	"go-mongo-redis/domain/datasources"
	"go-mongo-redis/domain/entities"
	"log"

	"github.com/go-redis/redis/v8"
)

type redisConnectionRepository struct {
	Context   context.Context
	RedisWR   *redis.Client
	RedisRead *redis.Client
}

type IRedisConnectionRepository interface {
	GetRedisData() []entities.UserDataFormat
	SetRedisData(dataByte []byte) bool
}

func NewRedisRepository(redis *datasources.RedisConnection) IRedisConnectionRepository {
	return &redisConnectionRepository{
		Context:   redis.Context,
		RedisWR:   redis.RedisWR,
		RedisRead: redis.RedisRead,
	}
}

func (repo redisConnectionRepository) GetRedisData() []entities.UserDataFormat {
	dataSpeakerRedis, err := repo.RedisRead.Get(repo.RedisRead.Context(), "UsersMock").Result()
	if err != nil {
		log.Println("error GetUsersData ", err.Error())
		return nil
	}
	var data []entities.UserDataFormat
	json.Unmarshal([]byte(dataSpeakerRedis), &data)
	log.Println("Get users data to redis success!")
	return data
}

func (repo redisConnectionRepository) SetRedisData(dataByte []byte) bool {
	err := repo.RedisWR.Set(repo.RedisWR.Context(), "UsersMock", dataByte, 0).Err()
	if err != nil {
		log.Println("error SetUsersName ", err.Error())
		return false
	}
	log.Println("Set new users data to redis success!")
	return true
}
