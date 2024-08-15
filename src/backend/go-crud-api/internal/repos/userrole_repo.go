package repos

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-crud-api/internal/models"
)

type UserRoleRepository struct {
	RedisClient *redis.Client
	Ctx         context.Context
}

func NewUserRoleRepository(redisClient *redis.Client) *UserRoleRepository {
	return &UserRoleRepository{
		RedisClient: redisClient,
		Ctx:         context.Background(),
	}
}

func (r *UserRoleRepository) Create(role *models.UserRole) error {
	data, err := json.Marshal(role)
	if err != nil {
		return err
	}
	err = r.RedisClient.Set(r.Ctx, fmt.Sprintf("userrole:%d", role.ID), data, 0).Err()
	return err
}

func (r *UserRoleRepository) FindAll() ([]models.UserRole, error) {
	keys, err := r.RedisClient.Keys(r.Ctx, "userrole:*").Result()
	if err != nil {
		return nil, err
	}

	var roles []models.UserRole
	for _, key := range keys {
		data, err := r.RedisClient.Get(r.Ctx, key).Result()
		if err != nil {
			return nil, err
		}
		var role models.UserRole
		if err = json.Unmarshal([]byte(data), &role); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}
