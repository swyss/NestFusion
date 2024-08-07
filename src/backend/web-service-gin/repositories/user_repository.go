package repositories

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"web-service-gin/models"
)

type UserRepository struct {
	Rdb *redis.Client
}

func NewUserRepository(rdb *redis.Client) *UserRepository {
	return &UserRepository{Rdb: rdb}
}

func (repo *UserRepository) Save(user *models.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}
	return repo.Rdb.Set(context.Background(), user.ID, data, 0).Err()
}

func (repo *UserRepository) FindById(id string) (*models.User, error) {
	data, err := repo.Rdb.Get(context.Background(), id).Result()
	if err != nil {
		return nil, err
	}
	var user models.User
	err = json.Unmarshal([]byte(data), &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	keys, err := repo.Rdb.Keys(context.Background(), "*").Result()
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		user, err := repo.FindById(key)
		if err != nil {
			continue
		}
		users = append(users, *user)
	}
	return users, nil
}

func (repo *UserRepository) Update(user *models.User) error {
	return repo.Save(user)
}

func (repo *UserRepository) Delete(id string) error {
	return repo.Rdb.Del(context.Background(), id).Err()
}
