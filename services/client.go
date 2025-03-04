package services

import (
	"context"
	"encoding/json"
	"go-myClient/models"
	"go-myClient/repository"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type ClientServ struct {
	repo  *repository.ClientRepo
	redis *redis.Client
}

func NewClientService(repo *repository.ClientRepo, redis *redis.Client) *ClientServ {
	return &ClientServ{repo: repo, redis: redis}
}

func (s *ClientServ) Create(c *models.Client) error {
	if err := s.repo.Create(c); err != nil {
		return err
	}
	return s.updateRedis(c)
}

func (s *ClientServ) Update(c *models.Client) error {
	if err := s.repo.Update(c); err != nil {
		return err
	}
	return s.updateRedis(c)
}

func (s *ClientServ) Delete(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	return s.redis.Del(context.Background(), strconv.Itoa(id)).Err()
}

func (s *ClientServ) GetByID(id int) (*models.Client, error) {
	var client *models.Client
	val, err := s.redis.Get(context.Background(), strconv.Itoa(id)).Result()
	if err == redis.Nil {
		client, err = s.repo.GetByID(id)
		if err != nil {
			return nil, err
		}
		if err := s.updateRedis(client); err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	} else {
		if err := json.Unmarshal([]byte(val), &client); err != nil {
			return nil, err
		}
	}
	return client, nil
}

func (s *ClientServ) updateRedis(client *models.Client) error {
	data, err := json.Marshal(client)
	if err != nil {
		return err
	}
	return s.redis.Set(context.Background(), client.Slug, data, time.Hour).Err()
}
