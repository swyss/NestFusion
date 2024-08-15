package repos

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-crud-api/internal/models"
)

type SettingRepository struct {
	RedisClient *redis.Client
	Ctx         context.Context
}

func NewSettingRepository(redisClient *redis.Client) *SettingRepository {
	return &SettingRepository{
		RedisClient: redisClient,
		Ctx:         context.Background(),
	}
}

func (r *SettingRepository) Create(setting *models.Setting) error {
	data, err := json.Marshal(setting)
	if err != nil {
		return err
	}
	err = r.RedisClient.Set(r.Ctx, fmt.Sprintf("setting:%d", setting.ID), data, 0).Err()
	return err
}

func (r *SettingRepository) FindAll() ([]models.Setting, error) {
	keys, err := r.RedisClient.Keys(r.Ctx, "setting:*").Result()
	if err != nil {
		return nil, err
	}

	var settings []models.Setting
	for _, key := range keys {
		data, err := r.RedisClient.Get(r.Ctx, key).Result()
		if err != nil {
			return nil, err
		}
		var setting models.Setting
		if err = json.Unmarshal([]byte(data), &setting); err != nil {
			return nil, err
		}
		settings = append(settings, setting)
	}
	return settings, nil
}
