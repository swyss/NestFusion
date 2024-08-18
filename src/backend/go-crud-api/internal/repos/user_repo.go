package repos

import (
	"database/sql"
	"go-crud-api/internal/models"
)

// UserRepository handles the CRUD operations for the User model in the database.
type UserRepository struct {
	DB *sql.DB
}

// NewUserRepository creates a new instance of UserRepository with the provided database connection.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// GetAllUsers retrieves all users from the database.
func (r *UserRepository) GetAllUsers() ([]models.User, error) {
	// Explicitly select relevant columns to ensure clarity.
	rows, err := r.DB.Query("SELECT id, is_active, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			// Handle error when closing rows
		}
	}(rows)

	var users []models.User
	for rows.Next() {
		var u models.User
		// Scan all four fields into the User struct
		if err := rows.Scan(&u.ID, &u.IsActive, &u.Name, &u.Email); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, rows.Err()
}

// GetUserByID retrieves a single user by their ID.
func (r *UserRepository) GetUserByID(id int) (*models.User, error) {
	var u models.User
	// Adjust the query and Scan to match the columns in the table
	err := r.DB.QueryRow("SELECT id, is_active, name, email FROM users WHERE id = $1", id).Scan(&u.ID, &u.IsActive, &u.Name, &u.Email)
	if err != nil {
		return nil, err
	}
	// Return the user if found.
	return &u, nil
}

// CreateUser inserts a new user into the database.
func (r *UserRepository) CreateUser(u *models.User) error {
	// Execute the insert query and retrieve the generated ID.
	return r.DB.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", u.Name, u.Email).Scan(&u.ID)
}

// UpdateUser updates an existing user in the database.
func (r *UserRepository) UpdateUser(u *models.User) error {
	// Execute the update query for the user.
	_, err := r.DB.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", u.Name, u.Email, u.ID)
	return err
}

// DeleteUser deletes a user from the database by their ID.
func (r *UserRepository) DeleteUser(id int) error {
	// Execute the delete query for the user.
	_, err := r.DB.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
