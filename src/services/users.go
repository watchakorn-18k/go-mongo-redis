package services

import (
	"encoding/json"
	"fmt"
	"go-mongo-redis/domain/entities"
	"go-mongo-redis/domain/repositories"
)

type usersService struct {
	UsersRepository repositories.IUsersRepository
	RedisRepository repositories.IRedisConnectionRepository
}

type IUsersService interface {
	GetAllUser() ([]entities.UserDataFormat, error)
	GetAllUserRedis() ([]entities.UserDataFormat, error)
	SetAllUserRedis() (string, error)
}

func NewUsersService(repo0 repositories.IUsersRepository, cache0 repositories.IRedisConnectionRepository) IUsersService {
	return &usersService{
		UsersRepository: repo0,
		RedisRepository: cache0,
	}
}

func (sv usersService) GetAllUser() ([]entities.UserDataFormat, error) {
	userData, err := sv.UsersRepository.FindAll()
	if err != nil {
		return nil, err
	}

	return userData, nil

}

func (sv usersService) GetAllUserRedis() ([]entities.UserDataFormat, error) {
	var getRedisData = sv.RedisRepository.GetRedisData()
	if getRedisData == nil {
		userData, err := sv.UsersRepository.FindAll()
		if err != nil {
			return nil, err
		}
		return userData, nil
	} else {
		return getRedisData, nil
	}

}

func (sv usersService) SetAllUserRedis() (string, error) {
	data, err := sv.UsersRepository.FindAll()
	if err != nil {
		return "", err
	}
	dataJSON, _ := json.Marshal(data)
	cache := sv.RedisRepository.SetRedisData(dataJSON)
	if !cache {
		return "", fmt.Errorf("cannot set to redis")
	}

	return "set to redissuccess", nil
}
