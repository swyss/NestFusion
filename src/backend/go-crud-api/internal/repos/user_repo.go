package repos

import (
	"database/sql"
	"go-crud-api/internal/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	rows, err := r.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	var u models.User
	err := r.DB.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) CreateUser(u *models.User) error {
	return r.DB.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", u.Name, u.Email).Scan(&u.ID)
}

func (r *UserRepository) UpdateUser(u *models.User) error {
	_, err := r.DB.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", u.Name, u.Email, u.ID)
	return err
}

func (r *UserRepository) DeleteUser(id int) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
